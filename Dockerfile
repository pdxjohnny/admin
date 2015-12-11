FROM ubuntu:14.04

WORKDIR /app

COPY ./static /app/static
COPY ./admin_linux-amd64 /app/run

CMD ["/app/run"]
