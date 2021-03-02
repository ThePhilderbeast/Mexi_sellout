FROM golang:1.16

RUN go get -u github.com/ThePhilderbeast/twitch_booli/cmd/bot

ENV APP_HOME /twitch_booli

WORKDIR $APP_HOME

CMD ["bot"]