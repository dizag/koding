AccountMixin = do ->
  init:(api)->
    {JAccount, JGuest} = api

    JAccount::tellKite = do->
      {Scrubber, Store} = Bongo.dnodeProtocol

      localStore = new Store
      remoteStore = new Store

      listenerId = 0

      channels = {}

      scrub =(method, args, callback) ->
        scrubber = new Scrubber localStore
        scrubber.scrub args, =>
          scrubbed = scrubber.toDnodeProtocol()
          scrubbed.method or= method
          callback scrubbed

      request =(kiteName, method, args)->
        scrub method, args, (scrubbed) ->
          fetchChannel kiteName, (channel)->
            messageString = JSON.stringify(scrubbed)
            channel.publish messageString

      response =(kiteName, method, args) ->
        scrub method, args, (scrubbed) ->
          fetchChannel kiteName, (channel)=>
            channel.publis "client-message", JSON.stringify(scrubbed)

      readyChannels = {}

      onReady = (channel, callback)->
        {name} = channel
        if readyChannels[name] then callback()
        else
          readyChannels[name] = channel
          channel.once 'ready', callback

      ready =(resourceName)->
        @exchange = resourceName
        @emit 'ready'

      messageHandler =(kiteName, args) ->
        if args.method is 'ready'
          callback = ready.bind this
        else
          callback = localStore.get(args.method)
        scrubber = new Scrubber localStore
        unscrubbed = scrubber.unscrub args, (callbackId)->
          unless remoteStore.has callbackId
            remoteStore.add callbackId, ->
              response kiteName, callbackId, [].slice.call(arguments)
          remoteStore.get callbackId
        callback.apply this, unscrubbed

      getChannelName = do ->
        namesCache = {}
        (kiteName)->
          return namesCache[kiteName]  if namesCache[kiteName]
          delegate  = KD.whoami()
          nickname  = delegate?.profile.nickname ?
                      if delegate.guestId then "guest#{delegate.guestId}" ?
                      'unknown'
          channelName = "#{Bongo.createId 128}.#{nickname}.kite-#{kiteName}"
          namesCache[kiteName] = channelName
          return channelName

      fetchChannel =(kiteName, callback)-> 
        channelName = getChannelName(kiteName)
        return callback channels[channelName]  if channels[channelName]
        channel = KD.remote.mq.subscribe channelName
        channels[channelName] = channel
        channel.on 'broker.subscribed', ->
          onReady channel, -> callback channel
        channel.setAuthenticationInfo
          serviceType : 'kite'
          name        : "kite-#{kiteName}"
          clientId    : KD.remote.getSessionToken()
        channel.on "message", messageHandler.bind channel, kiteName

      tellKite =(options, callback=->)->
        scrubber = new Scrubber localStore
        args = [options, callback]
        {method} = options
        delete options.autoCull
        request options.kiteName, method, args
