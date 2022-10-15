FROM archlinux:base

RUN pacman -Syu --noconfirm \
    && pacman -S --noconfirm \
        wget

WORKDIR /opt/ss-deployer

RUN wget https://github.com/tmota900/ss-deployer/releases/download/1.0.1/ss-deployer-1.0.1-linux-386.tar.gz \
    && tar -xvf ss-deployer-1.0.1-linux-386.tar.gz \
    && rm ss-deployer-1.0.1-linux-386.tar.gz

CMD ["/opt/ss-deployer/ss-deployer", "deployer"]