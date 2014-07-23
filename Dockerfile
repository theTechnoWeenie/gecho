FROM google/golang:1.3
MAINTAINER "Sam Jantz<sjantz0@gmail.com>"
ADD echoService.go /app/echoService.go
WORKDIR /app/
RUN go build echoService.go
RUN chmod +x echoService

EXPOSE 8080
CMD /app/echoService
