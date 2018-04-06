# Bleeding edge living dangerously ansible container
# for Fedora atom. Override entrypoint with ansible-playbook in alias

FROM fedora

LABEL maintainer="Tom Whiston <tom.whiston@gmail.com>"

RUN dnf install -y git zsh sshpass sudo net-tools which python2-dnf libselinux-python  && dnf clean all && \
    adduser -u 1000 -s /usr/bin/zsh -g 0 developer && \
    sed -i -e 's/^\(Defaults\s*requiretty\)/#--- \1/'  /etc/sudoers && \ 
    mkdir /etc/ansible && touch /etc/ansible/hosts && echo -e '[local]\nlocalhost ansible_connection=local' > /etc/ansible/hosts && \
    mkdir /etc/ansible/dev

# living dangerously
RUN pip install git+https://github.com/ansible/ansible.git@devel openshift virtualenv

USER 1000
WORKDIR /home/developer

ENTRYPOINT [ "ansible" ]


