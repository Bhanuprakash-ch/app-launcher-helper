[![Build Status](https://travis-ci.org/trustedanalytics/app-launcher-helper.svg)](https://travis-ci.org/trustedanalytics/app-launcher-helper)

App Launcher Helper
==================

App Launcher Helper is a service that provides a list of services created by a specific instance of [Application Broker](https://github.com/trustedanalytics/application-broker).

Usage
=====

The problem with Cloud Foundry and Application Broker is that there's no direct connection between a service and a related application for now. They're bound by naming convention only. App Launcher Helper is trying to fill this gap by providing a list of entries on a REST call:

```
http://hostname/rest/orgs/:orgId/atkinstances
```

Example response body:

```
{  
  "instances": [
    {
	  "name": "Name of a service instance",
  	  "url": "Url of an application related to a service instance",
	  "state": "current state of an app - STARTED, STOPPED, etc ...."
    }
  ]
}
```

AL Helper finds all service instances started in a specific organization for a service with a name configured by env variable. 
It's worth to mention that the application is an OAuth2 Resource Server, which means that there's access token in Authorization header needed. When deployed on Cloud Foundry, the application can be queried this way:

```
curl -H "Authorization: \`cf oauth-token|grep bearer\`" http://applauncher-helper.54.154.194.181.xip.io/rest/orgs/:orgId/atkinstances
```

Development
===========

To locally develop this application you'll need `godep` tool to manage dependencies and build the project. You will also need to have `gccgo-go` installed, as well as Go workspace created and GOPATH environment variable exported. If you don't meet these prerequisites, please refer to Development Environment Setup instructions in the project [Wiki] (https://github.com/trustedanalytics/project-wiki) for further instructions.

Clone app-launcher-helper using `gccgo-go`:
```
$ go get github.com/trustedanalytics/app-launcher-helper
``` 

Navigate to the project directory:
```
cd $GOPATH/src/github.com/trustedanalytics/app-launcher-helper
``` 

Build and test app-launcher-helper using `godep`:

```
$ godep go build
$ godep go test ./...
```


Deployment
==========

Before pushing the app to the Cloud Foundry, there're three env variables to be set:

* `TOKEN_KEY_URL` - an address of a key, to validate user's access token;
* `API_URL` - Cloud Foundry API address;
* `SERVICE_NAME` - a service name provided by App Launching Service Broker;

They are defined in manifest.yml, but they can be set by a `cf set-env` command as well.
When environment is ready, there's only one command needed:

```
$ cf push
```

Versioning
==========
`Bumpversion` tools is used to manage project version number, which is kept in two places: .bumpversion.cfg and manifest.yml. The first one is for bumpversion itself,
while the second one helps to identify the version of an application deployed in Cloud Foundry.

There's no need to use bumpversion manually - it's being used by CI.                                                                                             
