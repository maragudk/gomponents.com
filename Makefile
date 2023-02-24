.PHONY: build
build: build-css
	go run .

.PHONY: build-css
build-css: tailwindcss
	./tailwindcss -i tailwind.css -o docs/styles/app.css --minify

.PHONY: lint
lint:
	golangci-lint run

tailwindcss:
	curl -sLO https://github.com/tailwindlabs/tailwindcss/releases/latest/download/tailwindcss-macos-arm64
	chmod +x tailwindcss-macos-arm64
	mv tailwindcss-macos-arm64 tailwindcss

.PHONY: watch-css
watch-css: tailwindcss
	./tailwindcss -i tailwind.css -o docs/styles/app.css --watch
