FROM fedora

LABEL RUN='docker run -it -v /tmp/.X11-unix:/tmp/.X11-unix -e DISPLAY=$DISPLAY --device /dev/snd -v ~/.tidal:/root -it --privileged registry.twhiston.cloud/tomwhiston/tidal-cli'

RUN dnf install -y http://mirror.switch.ch/ftp/mirror/rpmfusion/free/fedora/rpmfusion-free-release-$(rpm -E %fedora).noarch.rpm 

RUN curl --silent --location https://rpm.nodesource.com/setup_8.x | bash -
RUN dnf install -y libstdc++
RUN dnf install -y git nodejs mpv w3m w3m-img

RUN npm -g i tidal-cli-client@latest

WORKDIR /root 
CMD tidal-cli
