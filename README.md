[![build-ci](https://github.com/jkaninda/s3fs/actions/workflows/build.yml/badge.svg)](https://github.com/jkaninda/s3fs/actions/workflows/build.yml)

# s3fs

 🐳 AWS S3 client, Object Storage client  

[![Build](https://github.com/jkaninda/s3fs/actions/workflows/build.yml/badge.svg)](https://github.com/jkaninda/mysql-bkup/actions/workflows/build.yml)
[![Go Report](https://goreportcard.com/badge/github.com/jkaninda/s3fs)](https://goreportcard.com/report/github.com/jkaninda/s3fs)
![Docker Image Size (latest by date)](https://img.shields.io/docker/image-size/jkaninda/s3fs?style=flat-square)
![Docker Pulls](https://img.shields.io/docker/pulls/jkaninda/s3fs?style=flat-square)

> Path:

AWS S3 Storage mounting path: `/s3mnt`

## Supported
- AWS S3
- MinIO
- And all AWS S3 alternative object storage solution


## Simple docker compose usage:

```yaml
version: '3.7'
services:
 s3fs:
  container_name: s3fs
  image: jkaninda/s3fs
  privileged: true
  devices:
    - "/dev/fuse"
  environment:
   - ACCESS_KEY=${ACCESS_KEY}
   - SECRET_KEY=${SECRET_KEY}
   - BUCKET_NAME=${BUCKET_NAME}
   - S3_ENDPOINT=${S3_ENDPOINT}
```
## Copy a file to S3
This is a simple example of copying a file from your local storage to S3, and after the copy container will stop running.
To keep the container running you need add `--keep` flag.

`s3fsctl mount --keep`


```yaml
version: '3.7'
services:
 s3fs:
  container_name: s3fs
  image: jkaninda/s3fs
  privileged: true
  volumes:
    - ./backup/:/backup
  devices:
    - "/dev/fuse"
    ## Mount S3 Storage and copy a file
  command:
   - /bin/sh
   - -c
   - |
    s3fsctl mount 
    cp /backup/my_file.gz /s3mnt/my_file.gz 
  environment:
   - ACCESS_KEY=${ACCESS_KEY}
   - SECRET_KEY=${SECRET_KEY}
   - BUCKET_NAME=${BUCKET_NAME}
   - S3_ENDPOINT=${S3_ENDPOINT}
```


> P.S. please give a star if you like it :wink:


