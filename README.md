[![build-ci](https://github.com/jkaninda/s3fs/actions/workflows/build.yml/badge.svg)](https://github.com/jkaninda/s3fs/actions/workflows/build.yml)

# s3fs

> 🐳 AWS S3 client, Object Storage client , Contabo Object Storage client 

* [Docker Hub](https://hub.docker.com/r/jkaninda/s3fs)
* [Github](https://github.com/jkaninda/s3fs)

## Supported
- AWS S3
- Contabo Object Storage
- ...


## Simple docker-compose usage:

```yaml
version: '3.7'
services:
 s3fs:
  container_name: s3fs
  image: jkaninda/s3fs
  privileged: true
  restart: always
  tty: true
  devices:
    - "/dev/fuse"
  environment:
   - ACCESS_KEY=${ACCESS_KEY}
   - SECRET_KEY=${SECRET_KEY}
   - BUCKETNAME=${BUCKETNAME}
   - S3_ENDPOINT=${S3_ENDPOINT}
```

> P.S. please give a star if you like it :wink:
