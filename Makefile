GOCMD=go
GOBUILD=$(GOCMD) build
GOCLEAN=$(GOCMD) clean
GOTEST=$(GOCMD) test
GOGET=$(GOCMD) get ./...
SERVICE=devices
DIST=dist
BINARY=$(DIST)/$(SERVICE)
DOCKER_REPOSITORY=tidepool/$(SERVICE)

all: test build
ci:	test docker-login docker-build docker-push-ci
dist:
		mkdir -p dist
generate:
			protoc \
        		-I ./api \
        		-I `go list -m -f "{{.Dir}}" github.com/grpc-ecosystem/grpc-gateway/v2`/third_party/googleapis \
        		--go_out=plugins=grpc,paths=source_relative:./api \
        		--grpc-gateway_out=./api \
        		api/api.proto

			mv ./api/github.com/tidepool-org/devices/api/* ./api/ && \
            	rm -r ./api/github.com
install:
		go get \
			github.com/golang/protobuf/protoc-gen-go \
			github.com/grpc-ecosystem/grpc-gateway/v2/protoc-gen-grpc-gateway \
			github.com/grpc-ecosystem/grpc-gateway/v2/protoc-gen-openapiv2 \
			github.com/rakyll/statik
build:	dist
		$(GOBUILD) -o $(DIST) ./...
test:
		$(GOTEST) -v ./...
clean:
		$(GOCLEAN)
		rm -rf $(DIST)
start:	build
		TIDEPOOL_DEVICES_CONFIG_FILENAME=`pwd`/devices.yaml ./$(BINARY)
deps:
		$(GOGET) ./...
docker-login:
		@echo "$(DOCKER_PASSWORD)" | docker login --username "$(DOCKER_USERNAME)" --password-stdin
docker-build:
		docker build -t $(SERVICE) .
docker-push-ci:
ifdef TRAVIS_BRANCH
ifdef TRAVIS_COMMIT
ifdef TRAVIS_PULL_REQUEST_BRANCH
	docker tag $(SERVICE) $(DOCKER_REPOSITORY):PR-$(subst /,-,$(TRAVIS_BRANCH))-$(TRAVIS_COMMIT)
	docker push $(DOCKER_REPOSITORY):PR-$(subst /,-,$(TRAVIS_BRANCH))-$(TRAVIS_COMMIT)
else
	docker tag $(SERVICE) $(DOCKER_REPOSITORY):$(subst /,-,$(TRAVIS_BRANCH))-$(TRAVIS_COMMIT)
	docker tag $(SERVICE) $(DOCKER_REPOSITORY):$(subst /,-,$(TRAVIS_BRANCH))-latest
	docker push $(DOCKER_REPOSITORY):$(subst /,-,$(TRAVIS_BRANCH))-$(TRAVIS_COMMIT)
	docker push $(DOCKER_REPOSITORY):$(subst /,-,$(TRAVIS_BRANCH))-latest
endif
endif
endif

.PHONY: ci-generate
ci-generate: generate

.PHONY: ci-build
ci-build: build

.PHONY: ci-test
ci-test: test
