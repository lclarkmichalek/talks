VERSION:=$(shell head -c 1024 /dev/urandom | md5sum | cut -f1 -d' ')

build:
	docker build -t bluepeppers/demo:$(VERSION) .

push: build
	docker push bluepeppers/demo:$(VERSION)

deploy: push
	kubectl set image deployments/demo-3 *=bluepeppers/demo:$(VERSION)
