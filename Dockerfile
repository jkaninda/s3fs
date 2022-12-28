FROM ubuntu:20.04
ENV ACCESS_KEY=""
ENV SECRET_KEY=""
ENV BUCKETNAME=""
ENV S3_URL=""
ENV TZ=Africa/Lubumbashi
ARG DEBIAN_FRONTEND=noninteractive

RUN apt-get update -qq
RUN apt-get install build-essential libcurl4-openssl-dev libxml2-dev mime-support -y
RUN apt install s3fs -y
RUN mkdir /s3mnt
RUN mkdir /tmp/s3cache
RUN chmod 777 /s3mnt
RUN chmod 777 /tmp/s3cache

COPY ./entrypoint.sh /usr/local/bin/
RUN chmod +x /usr/local/bin/entrypoint.sh
RUN ln -s /usr/local/bin/entrypoint.sh /

ENTRYPOINT ["entrypoint.sh"]



WORKDIR /root