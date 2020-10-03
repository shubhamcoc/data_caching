# data_caching
data caching using redis in GoLang.
This is a sample project to save key value pairs(i.e. like empid: empname) in database, fetch the results of a particular key from database (if not available in cache) and return the value as output. It can also search top 10 results and next 10 results and give the "key: value" as json output. The value stored in Cache stay for 10 mins and once the data is deleted backend will get the data from DB and store it in cache.

## Pre-requisite

To run the application, it require following tools to be installed in the system.

1. Docker
2. docker-compose
2. python (for testing purpose)

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

## Steps to run the application using Kubernetes

Use the following steps to run the application using kubernetes.
 
1. Go to deploy folder
   ```
   $ cd deploy
   ```

2. To build the docker images.
   ```
   $ docker-compose build
   ```

3. Go the kubernetes_deployment folder
   ```
   $ cd kubernetes_deployment
   ```

4. Launch the deployments and services
   ```
   $ kubectl apply -f broker_deployment.yaml
   ```

   ```
   $ kubectl apply -f broker_service.yaml
   ```

   ```
   $ kubectl apply -f backend_deployment.yaml
   ```

   ```
   $ kubectl apply -f backend_service.yaml
   ```

   ```
   $ kubectl apply -f ext_backend_service.yaml
   ```

   ```
   $ kubectl apply -f cache_deployment.yaml
   ```

   ```
   $ kubectl apply -f cache_service.yaml
   ```

   ```
   $ kubectl apply -f database_deployment.yaml
   ```

   ```
   $ kubectl apply -f database_service.yaml
   ```

## Steps to test the application

1. Run the server.py code kept in deploy/test folder.

2. Use the following URL to send data to the application. It can be done through browser or using curl in linux

   ```
   $ curl -v http://127.0.0.1:8082/submit
   ```

3. To see the last 10 results, use the following URL.

   ```
   $ curl http://127.0.0.1:10000/api/search
   ```

4. To see the value of a particular key, use the following URL.

   ```
   $ curl http://127.0.0.1:10000/api/searchbyid?key=123401
   ```
