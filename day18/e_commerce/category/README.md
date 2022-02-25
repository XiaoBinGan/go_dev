### 课程中的 
git.imooc.com/`cap1573`/category 
需要更改为 
git.imooc.com/`coding-447`/category
# Category service 
Go微服务定制容器
名称为 Category 类型 service 


## Getting Started

- [Configuration](#configuration)
- [Dependencies](#dependencies)
- [Usage](#usage)

## Configuration

- FQDN: go.micro.service.category
- Type: service
- Alias: category

## Dependencies

Micro services depend on service discovery. The default is multicast DNS, a zeroconf system.

In the event you need a resilient multi-host setup we recommend etcd.

 

## 构建

```
make build
```

Run the service
```
./category-service
```

Build a docker image
```
make docker
```