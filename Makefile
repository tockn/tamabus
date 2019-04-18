.PHONY: run webapp batch

mysql:
	mysql -u root --protocol=tcp -D tamabus -p

run: webapp batch

webapp:
	go run webapp/main.go

batch:
	python3 batch/main.py

test:
	go test -v ./...

docker_up:
	docker-compose up -d

deps:
	which dep || go get -v -u github.com/golang/dep/cmd/dep
	cd webapp; dep ensure
	which sql-migrate || go get github.com/rubenv/sql-migrate/...
