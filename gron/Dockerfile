FROM fedora

ENV GRON_VERSION=0.5.2

WORKDIR /usr/local/bin 
RUN curl -L https://github.com/tomnomnom/gron/releases/download/v${GRON_VERSION}/gron-linux-amd64-${GRON_VERSION}.tgz | tar xz

WORKDIR ~
ENTRYPOINT ["gron"]

