FROM alpine:3.20

RUN apk add --no-cache sqlite

RUN mkdir /data

WORKDIR /data

VOLUME ["/data"]

ENTRYPOINT ["sqlite3"]
CMD ["/data/database.db"]
