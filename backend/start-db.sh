#!/bin/bash
# ローカル環境でデータベースコンテナを構築・立ち上げる用の簡単なスクリプト
# Linux/macOS -> `./start-database.sh`

DB_CONTAINER_NAME="qr-order-postgres"

if ! [ -x "$(command -v docker)" ]; then
	echo "Docker is not installed. Please install docker and try again.\nDocker install guide: https://docs.docker.com/engine/install/"
	exit 1
fi

if [ "$(docker ps -q -f name=$DB_CONTAINER_NAME)" ]; then
	docker start $DB_CONTAINER_NAME
	echo "Database container started"
	exit 0
fi

# import env variables from .env
set -a
source .env

if [ "$DB_PASSWORD" = "password" ]; then
	echo "You are using the default database password"
	read -p "Should we generate a random password for you? [y/N]: " -r REPLY
	if ! [[ $REPLY =~ ^[Yy]$ ]]; then
		echo "Please set a password in the .env file and try again"
		exit 1
	fi
	DB_PASSWORD=$(openssl rand -base64 12)
	sed -i -e "s/:password@/:$DB_PASSWORD@/" .env
	echo $DB_PASSWORD
fi

docker run --name $DB_CONTAINER_NAME -e POSTGRES_USER=$DB_USER -e POSTGRES_PASSWORD=$DB_PASSWORD -e POSTGRES_DB=$DB_NAME -d -p 5432:5432 docker.io/postgres

echo "Database container was succesfuly created"
