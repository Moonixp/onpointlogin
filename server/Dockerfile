# building a custom image from alpine is better than the official image 
# built with archlinux which tallies at over 800MB
# this will be ~450MB during build then back to ~15MB after build
FROM alpine:latest

RUN apk update && apk add go
WORKDIR /server

COPY . . 

RUN go mod init server
RUN  go mod tidy

RUN go build -o /usr/bin/

# removing build tools
RUN rm -rf /server 
RUN apk del go

CMD ["server"]

