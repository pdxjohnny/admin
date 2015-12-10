FROM ubuntu:14.04
MAINTAINER John Andersen

RUN apt-get update -y && \
    apt-get install -qqy openssh-server git python && \
    apt-get clean && \
    rm -rf /var/cache/apt/* && \
    mkdir /var/run/sshd

COPY ./adduserserver /adduserserver
COPY ./startup.sh /startup.sh

CMD ["/startup.sh"]
