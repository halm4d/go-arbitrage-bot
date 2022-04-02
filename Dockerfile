FROM golang:1.18 as builder

ENV APP_HOME_SRC /go/src/go-arbitrage-bot/src
RUN mkdir -p $APP_HOME_SRC
WORKDIR $APP_HOME_SRC
COPY ./src .

RUN go mod download
RUN go mod verify
RUN go build -o arbotgo

FROM golang:1.18

ENV APP_USER app
ENV APP_HOME_SRC /go/src/go-arbitrage-bot/src

RUN groupadd $APP_USER && useradd -m -g $APP_USER -l $APP_USER
RUN mkdir -p APP_HOME_SRC
WORKDIR $APP_HOME_SRC

COPY --chown=0:0 --from=builder $APP_HOME_SRC/arbotgo arbotgo

USER $APP_USER
ENTRYPOINT ["./arbotgo"]