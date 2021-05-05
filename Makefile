.PHONY: build build-css build-css-dev lint

build: build-css
	go run *.go

build-css:
	NODE_ENV=production ./node_modules/.bin/tailwindcss build app.css -o docs/styles/app.css

build-css-dev:
	NODE_ENV=development ./node_modules/.bin/tailwindcss build app.css -o docs/styles/app.css

lint:
	golangci-lint run
