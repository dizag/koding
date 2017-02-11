debug = (require 'debug') 'nse:stackeditor'

kd = require 'kd'
bowser = require 'bowser'
Encoder = require 'htmlencode'

Events = require '../events'

FlexSplit = require './flexsplit'
FlexSplitStorage = require './flexsplit/storage'
AppStorageAdapter = require './adapters/appstorageadapter'

Toolbar = require './toolbar'
Editor = require './editor'
Statusbar = require './statusbar'
SideView = require './sideview'

LogsController = require '../controllers/logs'
VariablesController = require '../controllers/variables'
CredentialsController = require '../controllers/credentials'

{ yamlToJson } = require 'app/util/stacks/yamlutils'
Help = require './help'


module.exports = class StackEditor extends kd.View

  EDITORS = ['editor', 'readme', 'variables', 'logs']

  constructor: (options = {}, data = {}) ->

    super options, data

    # In-Memory Snapshot Storage
    @_snapshots = {}
    @_current = null

    # Storage
    @layoutStorage = new FlexSplitStorage
      adapter: AppStorageAdapter

    # Toolbar
    @toolbar = new Toolbar
    @forwardEvent @toolbar, Events.InitializeRequested
    @toolbar.on Events.ToggleCredentials, => @sideView.toggle 'credentials'

    # Status bar
    @statusbar = new Statusbar

    # Editor views
    @editor = new Editor {
      cssClass: 'editor'
      help: Help.stack
      filename: 'template.yaml'
      @statusbar
    }

    @logs = new Editor {
      cssClass: 'logs'
      title: 'Logs'
      filename: 'logs.sh'
      showgutter: no
      readonly: yes
      closable: yes
      @statusbar
    }

    @variables = new Editor {
      cssClass: 'variables'
      title: 'Custom Variables'
      filename: 'variables.yaml'
      help: Help.variables
      @statusbar
    }

    @readme = new Editor {
      cssClass: 'readme'
      title: 'Readme'
      filename: 'readme.md'
      help: Help.readme
      @statusbar
    }

    @controllers = {}

    @controllers.logs = new LogsController
      editor: @logs

    @controllers.variables = new VariablesController
      editor: @variables
      logs  : @controllers.logs

    # SideView for Search and Credentials
    @controllers.credentials = new CredentialsController
      logs  : @controllers.logs

    @controllers.credentials.on Events.WarnAboutMissingCredentials, =>
      @toolbar.showMissingCredentialWarning()

    @sideView       = new SideView
      title         : yes
      views         :
        credentials :
          title     : 'Credentials'
          cssClass  : 'credentials show-controls has-markdown'
          view      : @controllers.credentials.getView()
          controls  :
            plus    : =>
              @controllers.credentials.getCredentialAddButton()

    for _, controller of @controllers
      controller.on Events.TemplateDataChanged, @bound 'setTemplateData'

    @emit 'ready'


  setTemplateData: (data, reset = no) ->

    debug 'setTemplateData with args:', data, reset

    { _id: id, title, description, template } = data
    unless id or description or template
      throw { message: 'A valid JStackTemplate is required!' }

    @setData data
    @toolbar.setData data
    @controllers.variables.setData data
    @controllers.credentials.setData data

    @_saveSnapshot @_current  if @_current
    @_deleteSnapshot id  if reset

    @editor.setOption 'title', title

    unless @_loadSnapshot id

      @editor.setContent Encoder.htmlDecode template.rawContent
      @readme.setContent Encoder.htmlDecode description
      @variables.setContent ''
      @controllers.logs.set 'stack template loaded'

      @_saveSnapshot id
      @_current = id

    kd.utils.defer @editor.bound 'focus'


  check: (callback) ->

    [ err ] = @getConvertedContent()
    callback err


  getConvertedContent: ->

    convertedDoc = yamlToJson @editor.getContent()

    return if convertedDoc.err
    then [ 'Failed to convert YAML to JSON, fix the document and try again.' ]
    else [ null, convertedDoc ]


  _loadSnapshot: (id) ->

    return no  unless id
    return no  unless snapshot = @_snapshots[id]

    for view in EDITORS
      @[view]._restore snapshot[view]
    @_current = id

    return yes


  _saveSnapshot: (id) ->

    return  unless id

    @_snapshots[id] ?= {}
    for view in EDITORS
      @_snapshots[id][view] = @[view]._dump()


  _deleteSnapshot: (id) -> delete @_snapshots[id]


  viewAppended: ->

    # Layout
    @addSubView new FlexSplit
      cssClass            : 'mainview'
      resizable           : no
      views               : [
        @toolbar          # Toolbar on top, fixed height
        new FlexSplit
          resizable       : no
          views           : [
            contentView   = new FlexSplit
              name        : 'contentView'
              cssClass    : 'content'
              views       : [
                new FlexSplit
                  name    : 'leftColumn'
                  views   : [@editor, @logs]
                  sizes   : [90, 10]
                  storage : @layoutStorage
                new FlexSplit
                  name    : 'rightColumn'
                  sizes   : [50, 50]
                  views   : [@variables, @readme]
                  storage : @layoutStorage
              ]
              sizes       : [55, 45]
              type        : FlexSplit.VERTICAL
              storage     : @layoutStorage
            @statusbar    # Statusbar on bottom, fixed height
          ]
      ]

    contentView.setClass 'safari-flex-fix'  if bowser.safari
    contentView.addSubView @sideView
