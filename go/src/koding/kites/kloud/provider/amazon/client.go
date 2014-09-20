package amazon

import (
	"errors"
	"fmt"
	"strings"

	aws "koding/kites/kloud/api/amazon"
	"koding/kites/kloud/machinestate"
	"koding/kites/kloud/protocol"
	"koding/kites/kloud/waitstate"

	"github.com/koding/logging"
	"github.com/mitchellh/goamz/ec2"
)

type AmazonClient struct {
	*aws.Amazon
	Log  logging.Logger
	Push func(string, int, machinestate.State)

	// Used for customization
	InfoLog func(string, ...interface{})
}

func (a *AmazonClient) Build(instanceName string) (*protocol.Artifact, error) {
	infoLog := func(format string, args ...interface{}) {
		a.Log.Info(format, args...)
	}

	// Check wether a new infoLog was passed, and use it
	if a.InfoLog != nil {
		infoLog = a.InfoLog
	}

	// Don't build anything without this, otherwise ec2 complains about it as a
	// missing paramater.
	if a.Builder.SecurityGroupId == "" {
		return nil, errors.New("security group id is empty.")
	}

	// Make sure AMI exists
	infoLog("Checking if image '%s' exists", a.Builder.SourceAmi)
	if _, err := a.Image(a.Builder.SourceAmi); err != nil {
		if err != nil {
			return nil, err
		}
	}

	// Get the necessary keynames that we are going to provide with Amazon. If
	// it doesn't exist a new one will be created.  check if the key exist, if
	// yes return the keyname
	infoLog("Checking if keypair '%s' does exist", a.Builder.KeyPair)
	keyName, err := a.DeployKey()
	if err != nil {
		return nil, err
	}

	// Create instance with this keypair
	a.Builder.KeyPair = keyName

	infoLog("Creating instance with type: '%s' based on AMI: '%s'",
		a.Builder.InstanceType, a.Builder.SourceAmi)
	resp, err := a.CreateInstance()
	if err != nil {
		return nil, err
	}

	// we do not check intentionally, because CreateInstance() is designed to
	// create only one instance. If it creates something else we catch it here
	// by panicing
	instance := resp.Instances[0]

	stateFunc := func(currentPercentage int) (machinestate.State, error) {
		instance, err = a.Instance(instance.InstanceId)
		if err != nil {
			return 0, err
		}

		a.Push(fmt.Sprintf("Launching instance '%s'. Current state: %s",
			instanceName, instance.State.Name),
			currentPercentage, machinestate.Building)
		return statusToState(instance.State.Name), nil
	}

	ws := waitstate.WaitState{
		StateFunc:    stateFunc,
		DesiredState: machinestate.Running,
		Start:        45,
		Finish:       60,
	}
	if err := ws.Wait(); err != nil {
		return nil, err
	}

	return &protocol.Artifact{
		IpAddress:     instance.PublicIpAddress,
		InstanceName:  instanceName,
		InstanceId:    instance.InstanceId,
		SSHPrivateKey: a.Builder.PrivateKey,
		SSHUsername:   "", // deploy with root
	}, nil
}

func (a *AmazonClient) DeployKey() (string, error) {
	resp, err := a.Showkey(a.Builder.KeyPair)
	if err == nil {
		// key is found
		return resp.Keys[0].Name, nil
	}

	// not a ec2 error, return it
	ec2Err, ok := err.(*ec2.Error)
	if !ok {
		return "", err
	}

	// the key has another problem
	if ec2Err.Code != "InvalidKeyPair.NotFound" {
		return "", err
	}

	// ok now the key is not found, means it needs to be created
	a.Log.Info("Keypair '%s' doesn't exist, creating a new one", a.Builder.KeyPair)

	if a.Builder.PublicKey == "" {
		return "", errors.New("PublicKey is not defined. Can't create key")
	}

	key, err := a.CreateKey(a.Builder.KeyPair, a.Builder.PublicKey)
	if err != nil {
		return "", err
	}

	return key.KeyName, nil
}

func (a *AmazonClient) Start() (*protocol.Artifact, error) {
	a.Push("Starting machine", 10, machinestate.Starting)
	_, err := a.Client.StartInstances(a.Id())
	if err != nil {
		return nil, err
	}

	var instance ec2.Instance
	stateFunc := func(currentPercentage int) (machinestate.State, error) {
		instance, err = a.Instance(a.Id())
		if err != nil {
			return 0, err
		}

		a.Push(fmt.Sprintf("Starting instance '%s'. Current state: %s",
			a.Builder.InstanceName, instance.State.Name),
			currentPercentage, machinestate.Starting)

		return statusToState(instance.State.Name), nil
	}

	ws := waitstate.WaitState{
		StateFunc:    stateFunc,
		DesiredState: machinestate.Running,
		Start:        25,
		Finish:       60,
	}

	if err := ws.Wait(); err != nil {
		return nil, err
	}

	return &protocol.Artifact{
		InstanceId: instance.InstanceId,
		IpAddress:  instance.PublicIpAddress,
	}, nil
}

func (a *AmazonClient) Stop() error {
	a.Push("Stopping machine", 10, machinestate.Stopping)
	_, err := a.Client.StopInstances(a.Id())
	if err != nil {
		return err
	}

	stateFunc := func(currentPercentage int) (machinestate.State, error) {
		instance, err := a.Instance(a.Id())
		if err != nil {
			return 0, err
		}

		a.Push(fmt.Sprintf("Stopping instance '%s'. Current state: %s",
			a.Builder.InstanceName, instance.State.Name),
			currentPercentage, machinestate.Stopping)

		return statusToState(instance.State.Name), nil
	}

	ws := waitstate.WaitState{
		StateFunc:    stateFunc,
		DesiredState: machinestate.Stopped,
		Start:        25,
		Finish:       60,
	}

	return ws.Wait()
}

func (a *AmazonClient) Restart() error {
	a.Push("Restarting machine", 10, machinestate.Rebooting)
	_, err := a.Client.RebootInstances(a.Id())
	if err != nil {
		return err
	}

	stateFunc := func(currentPercentage int) (machinestate.State, error) {
		instance, err := a.Instance(a.Id())
		if err != nil {
			return 0, err
		}

		a.Push(fmt.Sprintf("Rebooting instance '%s'. Current state: %s",
			a.Builder.InstanceName, instance.State.Name),
			currentPercentage, machinestate.Rebooting)

		return statusToState(instance.State.Name), nil
	}

	ws := waitstate.WaitState{
		StateFunc:    stateFunc,
		DesiredState: machinestate.Running,
		Start:        25,
		Finish:       60,
	}

	return ws.Wait()
}

func (a *AmazonClient) Destroy() error {
	a.Push("Terminating machine", 10, machinestate.Terminating)
	_, err := a.Client.TerminateInstances([]string{a.Id()})
	if err != nil {
		return err
	}

	stateFunc := func(currentPercentage int) (machinestate.State, error) {
		instance, err := a.Instance(a.Id())
		if err != nil {
			return 0, err
		}

		a.Push(fmt.Sprintf("Terminating instance '%s'. Current state: %s",
			a.Builder.InstanceName, instance.State.Name),
			currentPercentage, machinestate.Terminating)

		return statusToState(instance.State.Name), nil
	}

	ws := waitstate.WaitState{
		StateFunc:    stateFunc,
		DesiredState: machinestate.Terminated,
		Start:        25,
		Finish:       60,
	}

	return ws.Wait()
}

func (a *AmazonClient) Info() (*protocol.InfoArtifact, error) {
	if a.Id() == "" {
		return &protocol.InfoArtifact{
			State: machinestate.NotInitialized,
			Name:  "not-existing-instance",
		}, nil
	}

	instance, err := a.Instance(a.Id())
	if err == aws.ErrNoInstances {
		return &protocol.InfoArtifact{
			State: machinestate.Terminated,
			Name:  "terminated-instance",
		}, nil
	}

	// if it's something else, return it back
	if err != nil {
		return nil, err
	}

	if statusToState(instance.State.Name) == machinestate.Unknown {
		a.Log.Warning("Unknown amazon status: %+v. This needs to be fixed.", instance.State)
	}

	var instanceName string
	for _, tag := range instance.Tags {
		if tag.Key == "Name" {
			instanceName = tag.Value
		}
	}

	// this shouldn't happen
	if instanceName == "" {
		a.Log.Warning("instance %s doesn't have a name tag. needs to be fixed!", a.Id())
	}

	return &protocol.InfoArtifact{
		State: statusToState(instance.State.Name),
		Name:  instanceName,
	}, nil

}

// statusToState converts a amazon status to a sensible machinestate.State
// format
func statusToState(status string) machinestate.State {
	status = strings.ToLower(status)

	// Valid values: pending | running | shutting-down | terminated | stopping | stopped

	switch status {
	case "pending":
		return machinestate.Starting
	case "running":
		return machinestate.Running
	case "stopped":
		return machinestate.Stopped
	case "stopping":
		return machinestate.Stopping
	case "shutting-down":
		return machinestate.Terminating
	case "terminated":
		return machinestate.Terminated
	default:
		return machinestate.Unknown
	}
}
