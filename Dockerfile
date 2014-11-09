FROM google/golang:1.3
MAINTAINER "Evan Cobb<cobb.evan@gmail.com>"
ADD main.go /app/main.go
WORKDIR /app/
RUN go get github.com/icbat/gamer-tag/tagservice
RUN go build main.go && chmod +x main

EXPOSE 8080
CMD /app/main
