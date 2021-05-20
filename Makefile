.PHONY: build build-css lint watch-css

build: build-css
	go run *.go

build-css:
	NODE_ENV=production ./node_modules/.bin/postcss build tailwind.css -o docs/styles/app.css

lint:
	golangci-lint run

watch-css:
	NODE_ENV=development ./node_modules/.bin/postcss build tailwind.css -o docs/styles/app.css -w
