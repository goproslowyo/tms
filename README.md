# Simple Talos Metadata Service

This is a simple HTTP service that listens on `MDS_LISTEN` (defaults to `127.0.0.1:8088`) to respond with YAML files based on the MAC address in the request. The service will respond with a corresponding YAML file if it is found in the `MDS_CONFIG_DIR` (defaults to `./configs`) directory otherwise it will return 404 not found.

## Setup

There are two environmental variables that can be set to change the behavior of the service.

* `MDS_CONFIG_DIR` - The directory to look for YAML files in. Defaults to `./configs`.
* `MDS_LISTEN` - The address and port to listen on. Defaults to `127.0.01:8088`.
* `MDS_LOG_LEVEL` - The log level to use. Defaults to `info`.

The service will look for a YAML file with the name of the MAC address in the `MDS_CONFIG_DIR` directory. For example, if the MAC address is `6c:4b:90:26:e4:4d` then the service will look for a file named `6c4b9026e44d.yaml` in the `MDS_CONFIG_DIR` directory. If the file is found then the contents of the file will be returned. If the file is not found then a 404 not found will be returned.

## Usage

```shell
$ MDS_LOG_LEVEL=debug MDS_LISTEN=0.0.0.0:1234 MDS_CONFIG_DIR=/path/to/configs ./metadata-service
DEBU[0000] Log level set to debug
DEBU[0000] Listen interface: 0.0.0.0:1234
DEBU[0000] Config directory: /path/to/configs
INFO[0000] Starting Talos Metadata Config Service
INFO[0000] Listening on 0.0.0.0:1234
DEBU[0014] Raw MAC: aa:bb:cc:00:11:22
DEBU[0014] Sanitized MAC aabbcc001122
INFO[0014] Attempting to load from /path/to/configs/aabbcc001122.yaml
ERRO[0014] Config file not found: open /path/to/configs/aabbcc001122.yaml: no such file or directory

$ curl 'localhost:1234/talos/config?mac=aa:bb:cc:00:11:22'
Not Found
```
