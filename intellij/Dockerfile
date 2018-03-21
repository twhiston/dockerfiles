FROM fedora

ARG USERNAME

ENV INTELLIJ_VERSION=U-2017.3.5

RUN dnf install -y which git findutils java-9-openjdk nodejs zsh util-linux-user hostname gcc dnf-plugins-core && \
    dnf config-manager --add-repo https://download.docker.com/linux/fedora/docker-ce.repo && \
    dnf install -y docker-ce && \
    dnf clean all

RUN mkdir /opt/intellij && curl https://download-cf.jetbrains.com/idea/ideaI${INTELLIJ_VERSION}.tar.gz | tar -C /opt/intellij -xz --strip 1 && \
    adduser -u 1000 -s /usr/bin/zsh -g 0 ${USERNAME}

USER 1000

WORKDIR /home/${USERNAME}

CMD /opt/intellij/bin/idea.sh
