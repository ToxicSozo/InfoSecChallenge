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

.PHONY: build-css
build-css:
	npm --prefix web run build:css

.PHONY: run
run: build
	$(BINPATH)/$(BINNAME)

.PHONY watch:
watch:
	$(MAKE) -j3 watch-app watch-templ watch-css

.PHONY: watch-app
watch-app:
	go run github.com/air-verse/air@latest \
	--build.cmd "$(MAKE) build-app" \
	--build.bin "$(BINPATH)/$(BINNAME)" \
	--build.include_ext "go"

.PHONY: watch-templ
watch-templ:
	templ generate \
	--watch \
	--proxy="http://localhost:8080/" \
	--open-browser=false

.PHONY: watch-css
watch-css:
	npm --prefix web run watch:css

.PHONY: fmt
fmt:
	templ fmt ./internal/view