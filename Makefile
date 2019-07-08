USER=actionloop
NAME=openwhisk-knative-operator
OPERATOR=$(USER)/$(NAME):latest
INIT=$(USER)/$(NAME)-init:latest
BRANCH=$(shell git rev-parse --abbrev-ref HEAD)
REPO=$(shell git config --get remote.origin.url)

build: operator init

init: $(find knative test -name \*.jl)
	-docker rmi -f $(INIT)
	docker build -t $(INIT) init
	docker push $(INIT)

operator: setup.jl
	-docker rmi -f $(OPERATOR)
	docker build -t $(OPERATOR) . --build-arg REPO="$(REPO)" --build-arg BRANCH="$(BRANCH)"
	docker push $(OPERATOR)

setup.jl: $(find src test -name \*.jl)
	touch setup.jl

clean:
	docker rmi -f $(IMG)

shell:
	docker run -ti -p 8000:8000 $(IMG) bash

run:
	docker run -p 8000:8000 $(IMG)

