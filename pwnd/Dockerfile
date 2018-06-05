FROM fedora

RUN dnf install -y perl-Digest-SHA findutils

COPY pwnd.sh /pwnd.sh
RUN chmod +x /pwnd.sh
USER 1000
ENTRYPOINT /pwnd.sh



