# Licensed to the Apache Software Foundation (ASF) under one or more contributor
# license agreements; and to You under the Apache License, Version 2.0.

build:
	go build main.go

test:
	go test -v ./kw

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
