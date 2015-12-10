FROM ubuntu:14.04
MAINTAINER John Andersen

RUN apt-get update -y && \
    apt-get install -qqy openssh-server git python sed && \
    apt-get clean && \
    rm -rf /var/cache/apt/* && \
    sed -i 's/#PasswordAuthentication yes/PasswordAuthentication yes/g' /etc/ssh/sshd_config  && \
    sed -i 's/PermitRootLogin without-password/PermitRootLogin no/g' /etc/ssh/sshd_config && \
    mkdir /var/run/sshd

COPY ./adduserserver /adduserserver
COPY ./startup.sh /startup.sh
COPY static /static

CMD ["/startup.sh"]
