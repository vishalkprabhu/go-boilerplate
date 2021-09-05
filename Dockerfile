FROM golang:1.15-alpine
LABEL maintainer="Vishalkprabhu"


ENV WAIT_VERSION 2.7.2
ADD https://github.com/ufoscout/docker-compose-wait/releases/download/$WAIT_VERSION/wait /wait
RUN chmod +x /wait

RUN apk --no-cache add \
    libc6-compat \
    tzdata \
    curl && \
    curl -L https://github.com/golang-migrate/migrate/releases/download/v4.14.1/migrate.linux-amd64.tar.gz | tar xvz && \
    mv migrate.linux-amd64 /bin/migrate


# Setting up Dev environment

EXPOSE 1200
CMD  ["/go-boilerplate"]