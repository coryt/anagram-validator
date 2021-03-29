build:
	docker-compose build

run: build
	docker-compose up

web-build: 
	cd web/anagram && yarn install

web-dev: web-build
	cd web/anagram && yarn start

server-dev:
	go run cmd/web/main.go

test:
	go test ./... -coverprofile=coverage.out

coverage: test
	go tool cover -html=coverage.out

seed:
	./tools/generate_data.sh