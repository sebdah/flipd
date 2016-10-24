PROJECT:=sebdah
SERVICE:=flipd
DOCKER_IMAGE:=$(PROJECT)/$(SERVICE)
LANGS:=\
	--gen go:thrift_import=github.com/apache/thrift/lib/go/thrift,package_prefix=github.com/sebdah/flipd/thrift/generated-code/gen-go/
GEN_DEST:=thrift/generated-code

.PHONY: build
build:
	$(info --- Building Docker image $(DOCKER_IMAGE))
	@docker build -t $(DOCKER_IMAGE) .

.PHONY: build-test
build-test:
	$(info --- Building Docker compose images)
	@docker-compose build

.PHONY: clean
clean: teardown
	$(info --- Cleaning Docker images)
	@docker-compose rm -vf

.PHONY: environment
environment:
	$(info --- Setting up environment)
	@docker-compose up -d

.PHONY: generate
generate-thrift:
	rm -rf $(GEN_DEST)
	mkdir -p $(GEN_DEST)
	thrift -o $(GEN_DEST) -r $(LANGS) flipd.thrift

.PHONY: lint
lint:
	$(info --- Running lint checks)
	@docker-compose run $(SERVICE) gometalinter --enable=misspell --cyclo-over=15 --disable=gotype --disable=ineffassign --fast ./source/... ; exit 0

.PHONY: teardown
teardown:
	$(info --- Tearing down the environment)
	@docker-compose stop -t 0

.PHONY: test
test: build-test environment lint
	$(info --- Running tests)
	@docker-compose run -e ENV=test $(SERVICE) go test -cover ./source/...
