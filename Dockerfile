FROM google/golang:1.3
MAINTAINER "Sam Jantz<sjantz0@gmail.com>"
ADD main.go /app/main.go
WORKDIR /app/
RUN go build main.go && chmod +x main

EXPOSE 8080
CMD /app/main
