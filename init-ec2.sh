#!/bin/bash
set -e

yum update -y
yum install -y docker

systemctl start docker
systemctl enable docker
usermod -a -G docker ec2-user

curl -L "https://github.com/docker/compose/releases/download/v2.24.5/docker-compose-$(uname -s)-$(uname -m)" -o /usr/local/bin/docker-compose
chmod +x /usr/local/bin/docker-compose
ln -sf /usr/local/bin/docker-compose /usr/bin/docker-compose

mkdir -p /home/ec2-user/app
cd /home/ec2-user/app

aws s3 cp s3://s3-config-to-ec2-try/public/docker-compose.prod.yml docker-compose.prod.yml

docker-compose -f docker-compose.prod.yml up -d