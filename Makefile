DATE = $(shell date +%s)

docker:
	docker build . -t registry.cn-hangzhou.aliyuncs.com/fatelei/juzhimiaohui-webhook:$(VERSION) --build-arg CACHEBUST=$(DATE)
