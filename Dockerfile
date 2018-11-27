FROM ubuntu:18.10

COPY booklist-api /bin/api

EXPOSE 8080

ENTRYPOINT ["/bin/api"]
