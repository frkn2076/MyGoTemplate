# MyGoTemplate

docker-compose --env-file infra/environments/DEV.env up

docker run --name postgre-instance -p 5432:5432 -e POSTGRES_USER=uSeR1 -e POSTGRES_PASSWORD=12345 -e POSTGRES_DB=SampleDB -d postgres:13-alpine

docker run --name mongo-instance -p 27017:27017 -e MONGO_INITDB_ROOT_USERNAME=uSeRrr -e MONGO_INITDB_ROOT_PASSWORD=12345 -d mongo