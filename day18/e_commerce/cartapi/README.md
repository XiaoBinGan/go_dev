# 课程中的 
git.imooc.com/`cap1573`/cartApi 
需要更改为 
git.imooc.com/`coding-447`/cartApi
# CartApi api 
慕课网 ap 老师 课程购买地址：http://www.imooc.com/t/6512963

当前服务 名称为 CartApi 类型 api 

创建初始化模版请使用

```
sudo docker run --rm -v $(pwd): $(pwd) -w  $(pwd) -e ICODE=xxxxxx cap1573/cap-micro new git.imooc.com/coding-447/cartApi
```
以上命令中 "xxxxxx" 为个人购买的 icode 码，请勿多人使用（会被慕课网检测失效）。

## 快速开始

- [配置信息](#配置信息)
- [使用](#使用)

## 配置信息

- 服务名称: go.micro.api.cartApi
- 类型: api
- 简称: cartApi

 

## 使用
根据 proto 自动生成
```
make proto
```

编译
```
make proto
```

构建镜像
```
make docker
```