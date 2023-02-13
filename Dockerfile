# base image
FROM golang:latest
WORKDIR /
COPY gin-sql-admin /usr/local/bin/
RUN apt-get update
RUN apt-get -y insatll net-tools
EXPOSE 8081
CMD ["go-sql-admin"]
ENTRYPOINT ["tail", "-f", "/dev/null"]