FROM golang:1.16-alpine as builder
WORKDIR /build
COPY . .
RUN go build -o ss-deployer .

FROM archlinux:base

RUN pacman -Syu --noconfirm \
    && pacman -S --noconfirm \
        kubectl

WORKDIR /opt/ss-deployer

COPY --from=builder /build/ss-deployer .

RUN chmod +x ss-deployer

CMD ["/opt/ss-deployer/ss-deployer", "deployer"]