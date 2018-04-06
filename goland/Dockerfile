FROM fedora

LABEL RUN='docker run --rm -v /tmp/.X11-unix:/tmp/.X11-unix -v /home/${USER}:/home/developer -e USER=${SUDO_UID} -e GROUP=${SUDO_GID} -e DISPLAY=${DISPLAY} -v /var/run/docker.sock:/var/run/docker.sock --security-opt label=type:container_runtime_t --security-opt seccomp=unconfined tomwhiston/goland:latest'

ENV GOLAND_VERSION=2018.1 \
    GO_VERSION=1.10 \
    GOSS_VERSION=0.3.5

RUN dnf install -y ansible docker-compose which sudo git findutils java-9-openjdk nodejs zsh util-linux-user hostname gcc dnf-plugins-core && \
    dnf config-manager --add-repo https://download.docker.com/linux/fedora/docker-ce.repo && \
    dnf install -y docker-ce && \
    dnf clean all && \
    adduser -u 1000 -s /usr/bin/zsh developer && \
    echo '%developer ALL=(ALL) NOPASSWD:ALL' >> /etc/sudoers 

RUN curl https://dl.google.com/go/go${GO_VERSION}.linux-amd64.tar.gz | tar -C /usr/local -xz && \
    mkdir /opt/goland && curl https://download-cf.jetbrains.com/go/goland-${GOLAND_VERSION}.tar.gz | tar -C /opt/goland -xz && \
    curl -L https://github.com/aelsabbahy/goss/releases/download/v${GOSS_VERSION}/goss-linux-amd64 -o /usr/local/bin/goss && \
    curl -L https://raw.githubusercontent.com/aelsabbahy/goss/v${GOSS_VERSION}/extras/dgoss/dgoss -o /usr/local/bin/dgoss && \
    chmod +rx /usr/local/bin/goss /usr/local/bin/dgoss

USER 1000
ENV PATH=$PATH:/usr/local/go/bin:/home/developer/Code/golang/bin \
    GOPATH=/home/developer/Code/golang

WORKDIR /home/developer
VOLUME /home/developer

CMD /opt/goland/GoLand-${GOLAND_VERSION}/bin/goland.sh
