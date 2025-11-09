FROM alpine:3.20

RUN apk add --no-cache curl tar

ARG MIGRATE_VERSION
ENV MIGRATE_VERSION=${MIGRATE_VERSION}

RUN curl -L https://github.com/golang-migrate/migrate/releases/download/${MIGRATE_VERSION}/migrate.linux-amd64.tar.gz \
    | tar -xz -C /usr/local/bin \
    && chmod +x /usr/local/bin/migrate

ENTRYPOINT ["migrate"]
