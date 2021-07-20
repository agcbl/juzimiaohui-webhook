FROM golang:latest
LABEL maintainer="fatelei@gmail.com"
LABEL version="1.0"
LABEL description="juzhimiaohui webhook"
RUN mkdir -p $GOPATH/src/github.com/agcbl/juzimiaohui-webhook
ARG CACHEBUST=1
COPY . $GOPATH/src/github.com/agcbl/juzimiaohui-webhook
WORKDIR $GOPATH/src/github.com/agcbl/juzimiaohui-webhook
RUN go mod tidy && go build -o bin/web -i cmd/web/main.go
ENV GIN_MODE=release
CMD ["bin/web", "-config", "/etc/webhook.toml"]
EXPOSE 8000
