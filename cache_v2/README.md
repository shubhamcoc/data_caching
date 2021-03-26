# data_caching
data caching using redis in GoLang.
This is a sample project to save key value pairs(i.e. like empid: empname) in database, fetch the results of a particular key from database (if not available in cache) and return the value as output. The value stored in Cache stay for 10 mins and once the data is deleted backend will get the data from DB and store it in cache.

## Components used
1. Swagger
2. MariaDB
3. Kafka

## Pre-requisite

To run the application, it require following tools to be installed in the system.

1. Docker
2. docker-compose

## Steps to run the application using docker

Use the following steps to run the application using docker.

1. Go to deploy folder
   ```
   $ cd deploy
   ```

2. To build the docker images.
   ```
   $ docker-compose build
   ```

3. To start the application.
   ```
   $ docker-compose up -d
   ```

## Steps to test the application

1. Use the curl commands given in [deploy/curl_command.txt](deploy/curl_command.txt)