# A GinBro RESTful APIs

## Packages
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

## Info
- table'schema which has no "ID","id","Id'" or "iD" will not generate model or route.
- the column which type is json value must be a string which is able to decode to a JSON,when call POST or PATCH.
## Thanks
- [swagger Specification](https://swagger.io/specification/)
- [gin-gonic/gin](https://github.com/gin-gonic/gin)
- [GORM](http://gorm.io/)
- [viper](https://github.com/spf13/viper)
- [cobra](https://github.com/spf13/cobra#getting-started)