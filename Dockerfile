FROM golang:1.22-alpine3.19 as builder
ENV APP_HOME /usr/src/app
WORKDIR "$APP_HOME"
RUN go install github.com/cosmtrek/air@latest
COPY [".", "$APP_HOME"]
RUN go mod download
RUN go build -C ./src -o "$APP_HOME"/src/app-go

FROM golang:1.22-alpine3.19
ENV APP_HOME /usr/src/app
WORKDIR "$APP_HOME"
COPY --from=builder ["$APP_HOME/docs", "$APP_HOME/docs"]
COPY --from=builder ["$APP_HOME/src/app-go", "$APP_HOME"]
CMD ["./app-go"]
