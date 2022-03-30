FROM golang:1.18 as builder

ENV APP_USER app
ENV APP_HOME /arbotgo

RUN groupadd $APP_USER && useradd -m -g $APP_USER -l $APP_USER
RUN mkdir -p $APP_HOME && chown -R $APP_USER:$APP_USER $APP_HOME

WORKDIR $APP_HOME
USER $APP_USER
COPY . .

RUN go mod download
RUN go mod verify
RUN go build -o arbotgo

FROM golang:1.18

ENV APP_USER app
ENV APP_HOME /arbotgo

RUN groupadd $APP_USER && useradd -m -g $APP_USER -l $APP_USER
RUN mkdir -p $APP_HOME
WORKDIR $APP_HOME

COPY --chown=0:0 --from=builder $APP_HOME/arbotgo $APP_HOME

USER $APP_USER
ENTRYPOINT ["./arbotgo"]