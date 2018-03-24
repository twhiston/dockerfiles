# qmk firmware build toolchain image
# by default builds my workman flavoured layout, which only exists in my fork
# deal with it!

FROM fedora

RUN dnf install -y  make gcc unzip wget sudo zip dfu-util dfu-programmer avr-gcc \
                    avr-libc binutils-avr32-linux-gnu arm-none-eabi-gcc-cs \
                    arm-none-eabi-binutils-cs arm-none-eabi-newlib

RUN adduser -u 1000 -g 0  developer && \
    echo '%root ALL=(ALL) NOPASSWD:ALL' >> /etc/sudoers


USER 1000
WORKDIR /home/developer

CMD ["make", "ergodox_infinity:workman"]
