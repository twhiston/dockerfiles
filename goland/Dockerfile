FROM fedora

RUN dnf install -y which git findutils java-9-openjdk nodejs zsh util-linux-user hostname && \
    dnf clean all

RUN curl https://dl.google.com/go/go1.10.linux-amd64.tar.gz | tar -C /usr/local -xz
ADD goland-2017.3.3.tar.gz /opt/goland

RUN adduser -u 1000 -s /usr/bin/zsh -g 0 twhiston
USER 1000
ENV PATH=$PATH:/usr/local/go/bin:~/Code/golang/bin \
    GOPATH=~/Code/golang
ENTRYPOINT /opt/goland/GoLand-2017.3.3/bin/goland.sh
