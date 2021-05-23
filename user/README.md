# User Server

### Create Proto
```bash
docker run --rm -v $(pwd):$(pwd) -w $(pwd) -e ICODE=EE953D449C12DEEC cap1573/cap-protoc -I ./ --go_out=./ --micro_out=./ ./user.proto
```

### Consul
```bash
docker pull consul

docker run -it -d -p 8500:8500 --name go-shop-consul consul
```