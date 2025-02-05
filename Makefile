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
	curl -sLO https://github.com/tailwindlabs/tailwindcss/releases/download/v3.4.15/tailwindcss-macos-arm64
	mv tailwindcss-macos-arm64 tailwindcss
	chmod +x tailwindcss
	mkdir -p node_modules/tailwindcss/lib && ln -s tailwindcss node_modules/tailwindcss/lib/cli.js
	echo '{"devDependencies": {"tailwindcss": "3.4.15"}}' >package.json

.PHONY: watch-css
watch-css: tailwindcss
	./tailwindcss -i tailwind.css -o docs/styles/app.css --watch
