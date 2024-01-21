FROM golang:1.21.0 AS build
WORKDIR /app

# Copy the source code.
COPY . .
# Installs Go dependencies
RUN go mod download

# Build
RUN CGO_ENABLED=0 GOOS=linux go build -o /app/s3fsctl

FROM ubuntu:24.04
ENV BUCKETNAME=""
ENV ACCESS_KEY=""
ENV SECRET_KEY=""
ENV S3_PATH=""
ENV S3_ENDPOINT=https://s3.amazonaws.com
ARG DEBIAN_FRONTEND=noninteractive
ENV VERSION="v0.1"
LABEL authors="Jonas Kaninda"

RUN apt-get update -qq
RUN apt install s3fs mysql-client supervisor -y

# Clear cache
RUN apt-get clean && rm -rf /var/lib/apt/lists/*

RUN mkdir /s3mnt
RUN mkdir /tmp/s3cache
RUN chmod 777 /s3mnt
RUN chmod 777 /tmp/s3cache

COPY --from=build /app/s3fsctl /usr/local/bin/s3fsctl
RUN chmod +x /usr/local/bin/s3fsctl

ADD docker/supervisord.conf /etc/supervisor/supervisord.conf


WORKDIR /root

CMD ["/bin/bash", "-c", "/usr/local/bin/s3fsctl mount --keep"]
