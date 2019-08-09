USER=actionloop
IMAGE=docker.io/$(USER)/kwhisk-controller

.PHONY: image
image: controller
	docker build . -t  $(IMAGE)

.PHONY: push
push: image
	docker login
	docker push $(IMAGE)

.PHONY: clean
clean:
	-rm controller

controller:
	cd ../../controller ;\
	CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build . ;\
	mv controller ../images/kwhisk-controller/