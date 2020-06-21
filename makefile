all:
	make build-front
	make build-lambdas

build-front:
	yarn run lint:netlify
	yarn build

build-lambdas:
	mkdir -p functions/
	cd src/functions/products && go get ./... && go build -o ../../../functions/products ./...