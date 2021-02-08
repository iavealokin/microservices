# STEP 1 build executable binary
FROM golang:latest

LABEL maintainer="Nikolaev <iavelokin@gmail.com>"
WORKDIR /home/remote/Go/src/Microservices/MS_Presistence/
RUN apt-get update
RUN apt-get install -y curl

COPY go.mod .
COPY go.sum .
RUN go mod download
COPY . . 

RUN    go get 
RUN    go build 
RUN chmod ugo+x entrypoint.sh
CMD ["./entrypoint.sh"]