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
ci:	test docker-build docker-push-ci
build:	deps
		$(GOBUILD) -o $(DIST) ./...
test:
		$(GOTEST) -v ./...
clean:
		$(GOCLEAN)
		rm -f $(BINARY)
start:	build
		./$(BINARY)
deps:
		$(GOGET) ./...
docker-login:
		@echo "$(DOCKER_PASSWORD)" | docker login --username "$(DOCKER_USERNAME)" --password-stdin
docker-build:
		docker build -t $(SERVICE) .
docker-push-ci:	docker-login
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
