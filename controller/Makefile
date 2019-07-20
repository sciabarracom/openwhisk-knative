build:
	go build main.go


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
	rm -Rvf gen

#-P knativewhisk.auth.Auth -t ./gen 
