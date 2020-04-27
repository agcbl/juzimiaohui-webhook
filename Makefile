DATE = $(shell date +%s)

docker:
	tar -cvf ./dockerfile/juzimiaohui-webhook.tar.gz .
	docker build -t fatelei/juzhimiaohui-webhook:$(VERSION) ./dockerfile --build-arg CACHEBUST=$(DATE)