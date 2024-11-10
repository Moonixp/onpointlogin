# building a custom image from alpine is better than the official image 
# built with archlinux which tallies at over 800MB
# this will be ~700MB during build then back to ~25MB after build
# issue is this might leave a container behind the build container so delete it with a label

# build with 
# docker build --rm . -t onpointserver && docker rmi `docker images --filter label=builder=true -q`

FROM alpine:latest AS builder
label builder=true

RUN apk update && apk add go
WORKDIR /server

COPY . . 

RUN go mod init server
RUN  go mod tidy

RUN go build -o /server/

# create a prod container 
from alpine:latest 

COPY --from=builder /server/server /usr/bin/server
CMD ["server"]

