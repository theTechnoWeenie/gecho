gamer-tag
=====

Search several popular games for a user.

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

## / (root)
// TODO build and document me


# Build instructions
// TODO document me

## Notes
I recommend running this in a container as the port is currently not configurable. Running it in the container will allow you to NAT 8080 (the port it runs on) to whatever you'd like.
If you are running mac, or windows you will need boot2docker (boot2docker.com) to build the docker container.

Sample command to run the echo service:
```
docker build -t echoservice .
docker run -d -P -e REGION=HOME echoservice
```

# Roadmap

