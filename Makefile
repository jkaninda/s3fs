BINARY_NAME=s3fs
include .env
export
run:
	go run .

build:
	go build -o bin/${BINARY_NAME} .

compile:
	GOOS=darwin GOARCH=arm64 go build -o bin/${BINARY_NAME}-darwin-arm64 .
	GOOS=darwin GOARCH=amd64 go build -o bin/${BINARY_NAME}-darwin-amd64 .
	GOOS=linux GOARCH=arm64 go build -o bin/${BINARY_NAME}-linux-arm64 .
	GOOS=linux GOARCH=amd64 go build -o bin/${BINARY_NAME}-linux-amd64 .

docker-build:
	docker build -f docker/Dockerfile  -t jkaninda/s3fs:latest .


docker-run: docker-build
	docker run --rm --privileged --device /dev/fuse --name s3fs -e "ACCESS_KEY=${ACCESS_KEY}" -e "SECRET_KEY=${SECRET_KEY}" -e "BUCKET_NAME=${BUCKET_NAME}" -e "S3_ENDPOINT=${S3_ENDPOINT}" -v ./backup:/backup jkaninda/s3fs

docker-run-copy:
	docker exec -it s3fs cp /backup/backup.gz /s3mnt