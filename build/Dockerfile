FROM golang:1.16

ARG version

RUN go get -u github.com/ThePhilderbeast/twitch_booli/cmd/bot@v1.0.3

ENV APP_HOME /twitch_booli

WORKDIR $APP_HOME

CMD ["bot"]