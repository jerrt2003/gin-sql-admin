# base image
FROM golang:latest

WORKDIR /

COPY gin-sql-admin /

EXPOSE 8081

CMD ["/go-sql-admin"]

ENTRYPOINT ["tail", "-f", "/dev/null"]