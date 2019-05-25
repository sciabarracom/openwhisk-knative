USER=actionloop
NAME=openwhisk-knative-operator
VERS=latest
IMG=$(USER)/$(NAME):$(VERS)

build: setup.jl
	docker build -t $(IMG) .

setup.jl: $(find src test -name \*.jl)
	touch setup.jl

clean:
	docker rmi -f $(IMG)

shell:
	docker run -ti -p 8000:8000 $(IMG) bash

run:
	docker run -p 8000:8000 $(IMG)
