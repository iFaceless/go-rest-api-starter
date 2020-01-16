FROM golang:latest

WORKDIR /apps/

RUN apt-get update
RUN apt-get install git
RUN git clone https://github.com/iFaceless/go-rest-api-starter.git go-starter

WORKDIR go-starter
RUN go mod download
RUN make web
RUN rm .env

#CMD ["/apps/go-rest-api-starter/bin/go-starter-web"]