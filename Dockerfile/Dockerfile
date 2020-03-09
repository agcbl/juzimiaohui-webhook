FROM golang:latest
LABEL maintainer="fatelei@gmail.com"
LABEL version="1.0"
LABEL description="juzhimiaohui webhook"
RUN mkdir -p $GOPATH/src/github.com/fatelei
ARG CACHEBUST=1
RUN cd $GOPATH/src/github.com/fatelei && git clone https://github.com/fatelei/juzimiaohui-webhook.git
WORKDIR $GOPATH/src/github.com/fatelei/juzimiaohui-webhook
RUN ls -l $GOPATH/src/github.com/fatelei/juzimiaohui-webhook
RUN go mod tidy
RUN go build -o bin/web -i cmd/web/main.go
ENV GIN_MODE=release
CMD ["bin/web", "-config", "/etc/webhook.toml"]
EXPOSE 8000
