.PHONY: build
build: build-css
	go run .

.PHONY: build-css
build-css: tailwindcss
	./tailwindcss -i tailwind.css -o docs/styles/app.css --minify

.PHONY: lint
lint:
	golangci-lint run

.PHONY: serve
serve:
	go run ./cmd/server

tailwindcss:
	curl -sL -o tailwindcss https://github.com/tailwindlabs/tailwindcss/releases/latest/download/tailwindcss-macos-arm64
	chmod a+x tailwindcss

.PHONY: watch-css
watch-css: tailwindcss
	./tailwindcss -i tailwind.css -o docs/styles/app.css --watch
