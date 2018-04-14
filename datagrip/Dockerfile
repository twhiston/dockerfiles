FROM fedora

LABEL RUN='docker run --rm -v /tmp/.X11-unix:/tmp/.X11-unix -v /home/\${USER}:/home/developer -e USER=\${SUDO_UID} -e GROUP=\${SUDO_GID} -e DISPLAY=\${DISPLAY} -v /var/run/docker.sock:/var/run/docker.sock --security-opt label=type:container_runtime_t --network host registry.twhiston.cloud/tomwhiston/datagrip:latest'

RUN dnf install -y sudo  which git findutils java-9-openjdk  zsh util-linux-user hostname gcc dnf-plugins-core && \
    dnf clean all && \
    adduser -u 1000 -s /usr/bin/zsh developer && \
    echo '%developer ALL=(ALL) NOPASSWD:ALL' >> /etc/sudoers

ENV DATAGRIP_VERSION=2017.3
RUN mkdir /opt/datagrip && curl https://download-cf.jetbrains.com/datagrip/datagrip-${DATAGRIP_VERSION}.tar.gz | tar -C /opt/datagrip -xz --strip 1

USER 1000

WORKDIR /home/developer
VOLUME /home/developer

CMD /opt/datagrip/bin/datagrip.sh
