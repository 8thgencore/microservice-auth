FROM alpine:3.20

ARG ENV=$ENV

RUN apk update && \
    apk upgrade && \
    apk add bash && \
    rm -rf /var/cache/apk/*

ADD https://github.com/pressly/goose/releases/download/v3.21.1/goose_linux_x86_64 /bin/goose
RUN chmod +x /bin/goose

WORKDIR /opt/app

COPY migrations/*.sql migrations/
COPY migration.sh ./migration.sh
COPY .env.${ENV} ./.env

RUN chmod +x migration.sh

ENTRYPOINT ["bash", "migration.sh"]