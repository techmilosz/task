.PHONY: run
run:
	go run ./main.go

.PHONY: builder-docker
build-docker:
	docker build -t recruitment:1.0 .

.PHONY: run-docker
run-docker: build-docker
	docker run \
	-dp 8080:8080 \
	recruitment:1.0

.PHONY: generate
generate:
	go generate ./...

.PHONY: install
install:
	go install go.uber.org/mock/mockgen@v0.4.0

.PHONY: test
test:
	go test -race ./...
