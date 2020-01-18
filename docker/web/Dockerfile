FROM golang:latest

WORKDIR /apps/

RUN apt-get update
RUN apt-get install git
RUN git clone https://github.com/iFaceless/go-rest-api-starter.git go-starter

WORKDIR /apps/go-starter
RUN go env -w GOPROXY=https://goproxy.cn,direct
RUN go mod download && make web && rm .env

RUN wget http://raw.githubusercontent.com/vishnubob/wait-for-it/master/wait-for-it.sh && chmod a+x ./wait-for-it.sh