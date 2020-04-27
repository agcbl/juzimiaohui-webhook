DATE = $(shell date +%s)

docker:
	tar -cvf ./dockerfile/chibaole-feishu-bot.tar.gz .
	docker build -t feishubot:$(VERSION) ./dockerfile --build-arg CACHEBUST=$(DATE)