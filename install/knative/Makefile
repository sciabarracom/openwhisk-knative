USER=actionloop
NAME=knative-installer
IMG=$(USER)/$(NAME):latest

.PHONY: push
push: build
	docker login
	docker push $(IMG)

.PHONY: build
build: installer.jl install.jl util.jl
	-docker rmi -f $(IMG)
	docker build -t $(IMG) .

.PHONY: clean
clean:
	docker rmi -f $(IMG)

.PHONY: shell
shell:
	docker run -ti -p 8000:8000 $(IMG) bash

.PHONY: run
run:
	docker run -p 8000:8000 $(IMG)

