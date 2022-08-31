#

- install packages
  ```
  go mod tidy
  ```

- run the consul docker image
    ```
    docker run -d -p 8500:8500 -p 8600:8600/udp --name=dev-consul consul:1.15.4 agent -server -ui -node=server-1 -bootstrap-expect=1 -client=0.0.0.0
    ```

- run each microservice
    ```
    go run *.go
    ```

- install the protobuf compiler
    ```
    go install google.golang.org/protobuf/cmd/protoc-gen-go@latest
    export PATH="$PATH:$(go env GOPATH)/bin"
    ```

- generate the protobuf code for the structures
    ```
    protoc -I=api --go_out=. --go-grpc_out=. movie.proto
    ```

- run testing benchmarks
    ```
    go test -bench=.
    ```

- install the gRPC plugin
    ```
    go install google.golang.org/grpc/cmd/protoc-gen-go-grpc@latest
    export PATH="$PATH:$(go env GOPATH)/bin"
    ```

- run the db docker instance
  ```
  docker run --name movieexample_db -e MYSQL_ROOT_PASSWORD=password -e MYSQL_DATABASE=movieexample -p 3306:3306 -d mysql:latest
  ```

- connect to the db docker instance and initialize the schema
  ```
  docker exec -i movieexample_db mysql movieexample -h localhost -P 3306 --protocol=tcp -uroot -ppassword < schema/schema.sql
  ```

- hit the service using grpcurl
  ```
  grpcurl -plaintext -d '{"record_id":"1", "record_type":"movie"}' localhost:8082 RatingService/GetAggregatedRating
  grpcurl -plaintext -d '{"record_id":"1", "record_type":"movie", "user_id": "alex", "rating_value": 5}' localhost:8082 RatingService/PutRating
  ```  

- build each service into an executable
  ```
  GOOS=linux go build -o main cmd/*.go
  ```

- build, tag, and publish docker images for each service
  ```
  docker build -t <tag> .
  docker tag <tag> <docker-hub-username>/<tag>:1.0.0
  docker push <docker-hub-username>/<tag>:1.0.0
  ```

- start minikube
  ```
  minikube start
  ```

- apply kubernenets deployment locally for each service
  ```
  kubectl apply -f kubernetes-deployment.yml
  kubectl get deployments
  kubectl get pods
  kubectl logs -f <POD_ID>
  ```

- run prometheus docker instance
  ```
  docker run -p 9090:9090 -v configs:/etc/prometheus prom/prometheus --config.file=/etc/prometheus/prometheus.yml  --web.enable-lifecycle
  ```

- run alertmanager
  ```
  docker run -p 9093:9093 -v /configs:/etc/alertmanager prom/alertmanager --config.file=/etc/alertmanager/alertmanager.yml
  ```