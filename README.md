# juzimiaohui-webhook
句子秒回 Webhook

## Config

```toml
[database]
name = "dbname"
user = "user"
password = "password"
host = "db host"

[word]
words = ["test"]

[lark]
path = "bot path"
```

## Run

### Init room

```shell script
go run scripts/room/main.go -token Token -config /path/to/webhook.toml
```

### Run webhook service

```shell script
docker run -p 8000:8000 -v /path/to/webhook.toml:/etc/webhook.toml fatelei/juzhimiaohui-webhook:1.1
```