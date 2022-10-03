# go-opentelemetry

## Distributed Tracing with OpenTelemetry

### Go apps

Test proppagation through:

1. gRPC
2. Kafka
3. Rest

### Benthos pipelines

Test propagation through inputs/processors/outputs

1. Kafka
3. Rest
4. gRPC

## Getting started


## Run applications

### Start `kafka` and `jaeger`


#### Zookeeper

```
docker run --rm --name zookeeper -p 2181:2181 -e ALLOW_ANONYMOUS_LOGIN=yes -d wurstmeister/zookeeper:latest 
```

#### Kafka
```
docker run --name kafka --rm -p 9092:9092 -e ADVERTISED_HOST=localhost -e KAFKA_ZOOKEEPER_CONNECT=host.docker.internal:2181 -e KAFKA_ADVERTISED_LISTENERS=OUTSIDE://host.docker.internal:9093,INSIDE://localhost:9092 -e KAFKA_LISTENERS=OUTSIDE://:9093,INSIDE://:9092 -e KAFKA_LISTENER_SECURITY_PROTOCOL_MAP=OUTSIDE:PLAINTEXT,INSIDE:PLAINTEXT -e KAFKA_INTER_BROKER_LISTENER_NAME=INSIDE -d wurstmeister/kafka:latest
```

#### Jaeger
```
docker run -d --name jaeger -e COLLECTOR_ZIPKIN_HTTP_PORT=9411 -p 5775:5775/udp -p 6831:6831/udp -p 6832:6832/udp -p 5778:5778 -p 16686:16686 -p 14268:14268 -p 9411:9411 jaegertracing/all-in-one:1.6
```

### Start both Benthos pipelines

```
 go run cmd/pipeline_a.go -c config/pipeline-a.yaml
```
```
 go run cmd/pipeline_b.go -c config/pipeline-b.yaml
```

### See messages arriving to `kafka`

```
kcat -b localhost:9092 -t messages -C -o end
```


### See `traces` on `Jaeger` UI

```
open http://localhost:16686
```

# gRPC

```sh
protoc -I=. --go-grpc_out=require_unimplemented_servers=false:. --go_out=:. service.proto
```
