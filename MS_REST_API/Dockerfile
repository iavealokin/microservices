# STEP 1 build executable binary
FROM golang:latest

LABEL maintainer="Nikolaev <iavelokin@gmail.com>"
WORKDIR /home/remote/Go/src/Microservices/MS_REST_API/


COPY go.mod .
COPY go.sum .
RUN go mod download
COPY . . 

#get dependancies
RUN    go get 
RUN    go build 
RUN chmod ugo+x entrypoint.sh
CMD ["./entrypoint.sh"]