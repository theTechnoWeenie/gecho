gecho
=====

Simple HTTP echo service written in go

# Configuration
Currently this is very static, and has only 2 configuration option.  
* Setting the environment variable `REGION` will change the greeting on the root page.
* Setting the environment variable `GREG_ADDRESS` will set the address for the the greg service registry (www.github.com/theTechnoWeenie/greg)

# Endpoints

## /uptime
This will return a JSON blob describing the uptime of the server.

### Example:
Request:
```
curl localhost:8080/uptime
```
Response:
```
{"Miliseconds":13804,"HourMinuteSecond":"00:00:13"}
```

## /echo
To use the echo service, either do a GET request with query parameters, or a post request with JSON in the body.

### Example:
Request:
```
curl localhost:8080/echo?param=value
```
Response:
```
{"param":["value"]}
```

Doing a post is similar except that the data that is sent in the body is returned as is. No conversion to JSON is done, so GIGO.

## /healthCheck
This is the endpoint that will be used by greg to determine if the service is still responsive.
Since this is such a simple service it  will only respond with

```
{ "status":"RUNNING" }
```
As defined in the `greg/server` package

## / (root)
This endpoint displays a very static page that uses the environemnt variable `REGION` to display where this server is being run.


# Build instructions
To build for local use
```
go build echoservice.go
```

To build in a docker container run

```
docker build .
```

## Notes
I recommend running this in a container as the port is currently not configurable. Running it in the container will allow you to NAT 8080 (the port it runs on) to whatever you'd like.
If you are running mac, or windows you will need boot2docker (boot2docker.com) to build the docker container.

Sample command to run the echo service:
```
docker build -t echoservice .
docker run -d -P -e REGION=HOME echoservice
```

# Roadmap

Although this is meant to be a simple service, the reason it was developed was to provide a sane basic service for a microservice environment.

## TODO:
- Add in registration with service discovery. (Service discovery method TBD, but probably something with SRV)
- Add in health metrics. No service is too small to monitor!
-- This would be adding endpoints such as /status, /metrics or really anything else that would be informative
