#!/bin/sh
Red='\033[0;31m'          # Red
Green='\033[0;32m'        # Green
set -e
if [ -z "${ACCESS_KEY}"] ||  [ -z "${SECRET_KEY}"]; then
echo "Please make sure all environment variables are set "
echo "BUCKETNAME=$BUCKETNAME \nACCESS_KEY=$nACCESS_KEY \nSECRET_KEY=$SECRET_KEY"
else
echo "$ACCESS_KEY:$SECRET_KEY" | tee /etc/passwd-s3fs
chmod 600 /etc/passwd-s3fs
echo "Mounting Object storage in /s3mnt .... "
s3fs $BUCKETNAME /s3mnt -o passwd_file=/etc/passwd-s3fs -o use_cache=/tmp/s3cache -o allow_other -o url=$S3_ENDPOINT -o use_path_request_style
ls /s3mnt | wc -l
fi

bash
