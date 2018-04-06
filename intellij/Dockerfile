FROM fedora

LABEL RUN='docker run --rm -v /tmp/.X11-unix:/tmp/.X11-unix -v /home/\${USER}:/home/developer -e USER=\${SUDO_UID} -e GROUP=\${SUDO_GID} -e DISPLAY=\${DISPLAY} -v /var/run/docker.sock:/var/run/docker.sock --security-opt label=type:container_runtime_t --security-opt seccomp=unconfined tomwhiston/intellij:latest'

ENV INTELLIJ_VERSION=U-2017.3.5

RUN dnf install -y sudo ansible docker-compose which git findutils java-9-openjdk nodejs zsh util-linux-user hostname gcc dnf-plugins-core && \
    dnf config-manager --add-repo https://download.docker.com/linux/fedora/docker-ce.repo && \
    dnf install -y docker-ce && \
    dnf clean all && \
    pip install virtualenv && \
    adduser -u 1000 -s /usr/bin/zsh developer && \
    echo '%developer ALL=(ALL) NOPASSWD:ALL' >> /etc/sudoers

RUN mkdir /opt/intellij && curl https://download-cf.jetbrains.com/idea/ideaI${INTELLIJ_VERSION}.tar.gz | tar -C /opt/intellij -xz --strip 1

USER 1000

WORKDIR /home/developer
VOLUME /home/developer

CMD /opt/intellij/bin/idea.sh
