FROM amacneil/dbmate:1.12.1

RUN apk update && apk add bash
COPY sql /db/migrations

COPY migrate.sh /usr/local/bin/

ENTRYPOINT ["migrate.sh"]

CMD ["up"]