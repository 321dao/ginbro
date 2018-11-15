#
## 环境
- 我的开发环境
    - Windows 10 专业版 64位
    - go version go1.11.1 windows/amd64
    - mysql 数据库 <= 5.7

## 依赖

	"github.com/gin-contrib/cors"
	"github.com/gin-contrib/static"
	"github.com/gin-gonic/autotls"
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
	"github.com/spf13/viper"
    "github.com/go-redis/redis"
    _ "github.com/go-sql-driver/mysql"
    "github.com/jinzhu/gorm"
- 使用

`ginbo gen -u root -p PASSWORD -a "127.0.0.1:3306" -d "dbname" -o "github.com/mojocn/apiapp" `


```shell
$ ./ginbo gen -h
generate a RESTful APIs app with gin and gorm for gophers. For example:
        ginbo gen -u eric -p password -a "127.0.0.1:3306" -d "mydb"

Usage:
  create gen [flags]

Flags:
  -a, --address string    mysql host:port (default "dev.mojotv.com:3306")
  -l, --appAddr string    app listen Address eg:mojotv.cn, use domain will support gin-TLS (default "127.0.0.1:5555")
  -c, --charset string    database charset (default "utf8")
  -d, --database string   database name (default "dbname")
  -h, --help              help for gen
  -o, --out string        golang project package name of your output project. eg: github.com/awesome/my_project, the project will be created at $GOPATH/src/github.com/awesome/my_project (default "github.
com/dejavuzhou/gin-project")
  -p, --password string   database password (default "Password")
  -u, --user string       database user name (default "root")


```

## 开发计划

- [ ] 支持PostgreSQL数据库
- [ ] 支持一键生产jwt密码验证
- [ ] 支持MongoDB数据库
- [ ] 更具数据映射关联模型
- [ ] 分页总数做redis缓存
- [ ] 支持生成gRPC服务


## 致谢
- [gin-gonic/gin框架](https://github.com/gin-gonic/gin)
- [GORM数据库ORM](http://gorm.io/)
