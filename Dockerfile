FROM google/golang:1.3
MAINTAINER "Sam Jantz<sjantz0@gmail.com>"
ADD echoservice.go /app/echoservice.go
WORKDIR /app/
RUN go build echoservice.go
RUN chmod +x echoservice

EXPOSE 8080
CMD /app/echoservice
