# A GinBro RESTful APIs

## 安装依赖
	go get github.com/gin-contrib/cors
	go get github.com/gin-contrib/static
	go get github.com/gin-gonic/autotls
	go get github.com/gin-gonic/gin
	go get github.com/sirupsen/logrus
	go get github.com/spf13/viper
    go get github.com/go-redis/redis
    go get github.com/go-sql-driver/mysql
    go get github.com/jinzhu/gorm
    
## Usage
- [swagger DOC ](http://{{.AppAddr}}/swagger/)`http://{{.AppAddr}}/swagger/`
- [static ](http://{{.AppAddr}})`http://{{.AppAddr}}`
- [app INFO ](http://1{{.AppAddr}}/app/info)`http://{{.AppAddr}}/app/info`
- API baseURL : `http://{{.AppAddr}}/api/v1`

## 注意
- table'schema which has no "ID","id","Id'" or "iD" will not generate model or route.
- the column which type is json value must be a string which is able to decode to a JSON,when call POST or PATCH.