# steps 
## devops
- install nginx
- remove all default conf files
- create files:
  - base config
  - client dn config
  - empty mtls server config
  - empty clear channel server config 
  - mtls base config yaml
  - clear_channel base config yaml
- bonus: validate server yaml files to schema

## software
- update config to support multiple files
- when clear channel
  - all ports added to clear channel server conf
    - parse clear channel yaml
    - 
  - write clear channel conf to file
  - write null to mtls conf file
- when mtls enabled
  - all port are added to mtls server conf
  - write mtls conf to file
  - write null to clear channel conf file
  - 