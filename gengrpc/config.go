package gengrpc

var yamlConfig = `
# postgres configuration
mysql:
  username: root
  password: 'Kalifun677520'
  path: '127.0.0.1'
  port: 3306
  db-name: 'casbin'
  config: 'charset=utf8&parseTime=True&loc=Local'
  max-idle-conns: 10
  max-open-conns: 10

system:
  RPCPORT: ':8100'

  # logger configuration
log:
  prefix: '[auth-center]'
  log-file: true
  stdout: 'DEBUG'
  file: 'DEBUG'
`
