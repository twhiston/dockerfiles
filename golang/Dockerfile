FROM fedora

LABEL RUN='docker run --rm -v $(pwd):/home/developer -e USER=${SUDO_UID} -e GROUP=${SUDO_GID} --security-opt label=type:container_runtime_t registry.twhiston.cloud/tomwhiston/golang:latest'

ENV GO_VERSION=1.10 

RUN dnf install -y gcc git && adduser -u 1000 developer

RUN curl https://dl.google.com/go/go${GO_VERSION}.linux-amd64.tar.gz | tar -C /usr/local -xz 
RUN mkdir -p /home/developer/go/src && chown -R developer /home/developer/go

ENV PATH=$PATH:/usr/local/go/bin:/home/developer/go/bin

USER developer 
WORKDIR /home/developer/go
VOLUME /home/developer/go/src/app

ENTRYPOINT ["/usr/local/go/bin/go"]

