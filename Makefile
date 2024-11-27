BINPATH = $(HOME)/Desktop/cyrsovaia
BINNAME = cyrsovaia

.PHONY: build
build: build-templ build-app

.PHONY: build-app
build-app:
	go build -o $(BINPATH)/$(BINNAME) ./cmd/main.go

.PHONY: build-templ
build-templ:
	templ generate

.PHONY: run
run: build
	$(BINPATH)/$(BINNAME)

.PHONY: fmt
fmt:
	templ fmt ./internal/view