FROM golang:1.18

ENV APP_USER app
ENV APP_HOME /go/src/go-arbitrage-bot/bin

RUN groupadd $APP_USER && useradd -m -g $APP_USER -l $APP_USER
RUN mkdir -p $APP_HOME
WORKDIR $APP_HOME

COPY --chown=0:0 --from=builder $APP_HOME/arbotgo-linux $APP_HOME/arbotgo
COPY --chown=0:0 --from=builder $APP_HOME/.version $APP_HOME/.version

USER $APP_USER
ENTRYPOINT ["./arbotgo"]