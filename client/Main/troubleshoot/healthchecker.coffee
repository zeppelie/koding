class HealthChecker extends KDObject

  constructor: (options={}, @cb) ->
    super options

    @identifier = options.identifier or Date.now()
    @status = "not started"

  run: ->
    @emit "healthCheckStarted"
    @status = "waiting"
    @startTime = Date.now()
    @setPingTimeout()
    @cb @finish.bind(this)

  setPingTimeout: ->
    @pingTimeout = setTimeout =>
      @status = "down"
      @emit "healthCheckCompleted"
    , 5000

  finish: (data)->
    # some services (e.g. kite controller) does return callback with error parameter
    # hence we are having
    unless @status is "down"
      @status = "success"
      @finishTime = Date.now()
      clearTimeout @pingTimeout
      @pingTimeout = null
      @emit "healthCheckCompleted"

  reset: ->
    @status = "waiting"
    @finishTime = null
    @startTime = null

  getResponseTime: ->
    if @status is "success" then @finishTime - @startTime else 0