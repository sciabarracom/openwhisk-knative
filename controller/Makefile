USER=actionloop
IMAGE=docker.io/$(USER)/kwhisk-controller

.PHONY:
run:
	go run main.go

.PHONY: gen
gen: 
	-mkdir gen
	swagger generate server -f apiv1swagger.json -t gen -P models.Auth --exclude-main 
	cp auth.src gen/models/auth.go
	
.PHONY: clean
clean:
	-rm -Rvf gen
	-rm controller


.PHONY: image
image: controller
	docker build . -t  $(IMAGE)

.PHONY: push
push: image
	docker login
	docker push $(IMAGE)

controller:
	CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build . 