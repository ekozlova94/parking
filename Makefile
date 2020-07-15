.PHONY: build
build:
	go build cmd/parking/main.go

.PHONY: run
run:
	go run cmd/parking/main.go

.PHONY: test
test:
	go test ./... -v -count 2

.PHONY: e2e-test
e2e-test:
	cd tests && go test -v -tags e2e -count 1
