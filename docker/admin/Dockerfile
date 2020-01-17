FROM golang:latest

WORKDIR /apps/

RUN apt-get update
RUN apt-get install git
RUN git clone https://github.com/iFaceless/go-rest-api-starter.git go-starter

WORKDIR /apps/go-starter
RUN go mod download
RUN make admin
RUN rm .env

CMD ["bin/starter-admin"]