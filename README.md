## 准备
- 安装etcd
- 安装goctl
- 在mysql数据库中执行`service/user/model/user.sql`

## 快速开始
```bash
# 启动etcd
etcd
# 下载相关依赖
go mod tidy
go mod download
# 使用 make 命令启动rpc
make user-rpc
make book-rpc
# 使用 make 命令启动api
make user-api
make book-api
make search-api
```