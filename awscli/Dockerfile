FROM fedora
LABEL maintainer "Tom Whiston <tom.whiston@gmail.com>"

RUN dnf install -y \
	groff \
	less \
	python \
	&& pip install awscli \ 
  && dnf clean all

ENTRYPOINT [ "aws" ]
