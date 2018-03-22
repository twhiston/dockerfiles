FROM fedora

RUN dnf install -y git && \
    adduser -u 1000 developer

VOLUME /home/developer /git
WORKDIR /git
USER 1000
ENTRYPOINT ["git"]

