#

## 安装依赖
	go get -u github.com/gin-contrib/cors
	go get -u github.com/gin-contrib/static
	go get -u github.com/gin-gonic/autotls
	go get -u github.com/gin-gonic/gin
	go get -u github.com/sirupsen/logrus
	go get -u github.com/spf13/viper
    go get -u github.com/go-redis/redis
    go get -u github.com/go-sql-driver/mysql
    go get -u github.com/jinzhu/gorm
    
## 使用
- [swagger DOC ](http://{{.AppAddr}}/swagger/)`http://{{.AppAddr}}/swagger/`
- [static ](http://{{.AppAddr}})`http://{{.AppAddr}}`
- [app INFO ](http://1{{.AppAddr}}/app/info)`http://{{.AppAddr}}/app/info`
- API baseURL : `http://{{.AppAddr}}/api/v1`