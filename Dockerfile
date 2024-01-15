FROM golang:latest

WORKDIR /go/app

RUN apt-get update && apt-get install -y librdkafka-dev && apt-get clean

CMD [ "tail", "-f", "/dev/null" ]