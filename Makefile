dev-setup:
	cd public; npm install; npm run build
	cp dev_config.yml config.yml

run-dev-backend:
	go run cmd/main.go

run-dev-frontend:
	cd public; npm run dev

swag:
	${GOPATH}/bin/swag init -d cmd,internal -g ../internal/app/server.go

IMG ?= accweb/accweb
VERSION ?= dev
TAG ?= latest
docker-build:
	docker build . -t $(IMG):$(TAG) --progress plain --build-arg VERSION=$(VERSION) --build-arg COMMIT=`git rev-parse HEAD` 