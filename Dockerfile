# base image
FROM golang:latest
WORKDIR /
COPY gin-sql-admin /usr/local/bin/
RUN mkdir /configs
COPY configs/mysql.yaml /configs/
RUN apt-get update
RUN apt-get -y install net-tools vim iputils-ping telnet
# Install GO debugger: delve
RUN go install github.com/go-delve/delve/cmd/dlv@latest
EXPOSE 8081
CMD ["go-sql-admin"]
ENTRYPOINT ["tail", "-f", "/dev/null"]