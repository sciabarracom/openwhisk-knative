USER=actionloop
NAME=openwhisk-knative-operator
IMG=$(USER)/$(NAME):latest
BRANCH=$(shell git rev-parse --abbrev-ref HEAD)
REPO=$(shell git config --get remote.origin.url)

build: setup.jl
	docker rmi -f $(IMG)
	docker build -t $(IMG) . --build-arg REPO="$(REPO)" --build-arg BRANCH="$(BRANCH)"
	docker push $(IMG)

setup.jl: $(find src test -name \*.jl)
	touch setup.jl

clean:
	docker rmi -f $(IMG)

shell:
	docker run -ti -p 8000:8000 $(IMG) bash

run:
	docker run -p 8000:8000 $(IMG)

