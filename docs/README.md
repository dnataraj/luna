Luna - A Simple Distributed Logger
============


Initial components being considered:

  * Client (obviously)
  * Logger API
  * NSQ
  * ElasticSearch
  * Log Viewer (Browser / Electron App)

### 1. Simple Architecture

![Architecture](./logger.png)

### 2. Installing gRPC / Protobuf

https://grpc.io/docs/quickstart/go/

### 3. protoc and build 

protoc -I logger/ logger/logger.proto --go_out=plugins=grpc:logger

### 4. Building Starting up services

  * Working directory : `luna/`	

  * NSQ
    
    ```
	$> cd nsq
	$> docker-compose up -d
    ```

    Make sure NSQ is running with
    `docker-compose ps`
  
  * Luna Server
   
    ```
	$> cd server
	$> go run server.go
    ```
    
  * Sample client

    ```
	$> cd examples/client_go
	$> go run client.go
    ```

#### TODO

  - Ingestion key for clients (rather than client id)
  - Log consumers
  - Client log durability
  - Connection retry
  -
