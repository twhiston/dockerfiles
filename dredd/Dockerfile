FROM registry.twhiston.cloud/tomwhiston/golang:latest

LABEL maintainer="Tom Whiston <tom.whiston@gmail.com>"
LABEL RUN="docker run -it --rm -v $(pwd):/home/developer/src --network host tomwhiston/dredd"

USER 0
RUN dnf install -y gcc-c++ python nodejs && mkdir -p /home/developer/go/src && chown -R developer /home/developer/go

USER developer
WORKDIR /home/developer
RUN npm i -D dredd 

WORKDIR /home/developer
RUN go get github.com/snikch/goodman/cmd/goodman

WORKDIR /home/developer/src
ENTRYPOINT ["/home/developer/node_modules/.bin/dredd"]
