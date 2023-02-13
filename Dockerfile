# base image
FROM golang:latest
WORKDIR /
COPY gin-sql-admin /usr/local/bin/
RUN mkdir /configs
COPY configs/mysql.yaml /configs/
RUN apt-get update
RUN apt-get -y install net-tools vim iputils-ping telnet
EXPOSE 8081
ENTRYPOINT ["gin-sql-admin","/configs/mysql.yaml"]