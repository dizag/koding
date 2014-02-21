module.exports = (options={})->
  options.favicon or= "/a/images/favicon.ico"
  """
  <meta charset="utf-8" />
  <meta http-equiv="X-UA-Compatible" content="IE=edge,chrome=1" />
  <meta name="description" content="" />
  <meta name="author" content="" />
  <meta name="apple-mobile-web-app-capable" content="yes">
  <meta name="apple-mobile-web-app-status-bar-style" content="black">
  <meta name="apple-mobile-web-app-title" content="Koding" />
  <meta name="viewport" content="user-scalable=no, width=device-width, initial-scale=1" />
  <link rel="shortcut icon" href="#{options.favicon}" />
  <link rel="fluid-icon" href="/a/images/logos/fluid512.png" title="Koding" />
  <link rel="stylesheet" href="/a/css/kd.#{KONFIG.version}.css" />
  <link rel="stylesheet" href="/a/css/introapp.#{KONFIG.version}.css" />
  <link rel="stylesheet" href="/a/css/koding.#{KONFIG.version}.css" />
  """
