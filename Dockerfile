FROM golang:1.16-alpine as builder
WORKDIR /build
COPY . .
RUN go build -o ss-deployer .

FROM alpine:latest

RUN apk add curl
RUN apk add bash

RUN curl -LO "https://dl.k8s.io/release/$(curl -L -s https://dl.k8s.io/release/stable.txt)/bin/linux/$(arch | sed s/aarch64/arm64/ | sed s/x86_64/amd64/)/kubectl"

RUN chmod +x ./kubectl
RUN mv ./kubectl /usr/local/bin/

WORKDIR /opt/ss-deployer

COPY --from=builder /build/ss-deployer .

RUN chmod +x ss-deployer

CMD ["/opt/ss-deployer/ss-deployer", "deployer"]