FROM golang:1.15-alpine
LABEL maintainer="Vishalkprabhu"


ENV WAIT_VERSION 2.7.2
ADD https://github.com/ufoscout/docker-compose-wait/releases/download/$WAIT_VERSION/wait /wait
RUN chmod +x /wait

# Setting up Dev environment

WORKDIR /web_app/
EXPOSE 1200
CMD  "/web_app/go-boilerplate"