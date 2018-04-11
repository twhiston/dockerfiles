FROM fedora

LABEL maintainer="Tom Whiston <tom.whiston@gmail.com>"
LABEL RUN="docker run -it --rm -v $(pwd):/home/developer/src --network host tomwhiston/dredd"

RUN dnf install -y git python nodejs && adduser -u 1000 developer

USER developer
WORKDIR /home/developer
RUN npm i -D dredd 
WORKDIR /home/developer/src
ENTRYPOINT ["/home/developer/node_modules/.bin/dredd"]
