#!/bin/sh
Red='\033[0;31m'          # Red
Green='\033[0;32m'        # Green
set -e
echo "$ACCESS_KEY:$SECRET_KEY" | tee /etc/passwd-s3fs
chmod 600 /etc/passwd-s3fs

s3fs $BUCKETNAME /s3mnt -o passwd_file=/etc/passwd-s3fs -o use_cache=/tmp/s3cache -o allow_other -o url=$S3_URL -o use_path_request_style

bash
