# [Converting a MySQL database to a RESTful golang APIs app in the fastest way](https://github.com/dejavuzhou/ginbro)
##### ginbro,**GinBro**,Gimbo,GimBro,**Jimbo**,GinOrm or GinGorm

## [中文](readme_zh.md)
## Feature
- powered with Swagger document and SwaggerUI
- capable of serve VueJs app's static files
- fastest way to generate a RESTful APIs application in Go
    
## Ginbro Installation
you can install it by go get command：
```shell
go get github.com/dejavuzhou/ginbro
```
the Ginbro executable binary will locate in $GOPATH/bin
check GOBIN is in your environment PATH




## Usage
`ginbro gen -u root -p PASSWORD -a "127.0.0.1:3306" -d dbname -o "github.com/mojocn/apiapp"` -o api_app_project
- cd $GOPATH/src/api_app_project
- `go build`  and `run`
- Visit[`http://127.0.0.1:5555/swagger`](http://127.0.0.1:5555/swagger)

### the generated project directory [ginbro-son DEMO](https://github.com/dejavuzhou/ginbro-son)
```shell
$GOPATH/src/apiapp>tree /f /a
Folder PATH listing
Volume serial number is 8452-D575
C:.
|   2018-11-15-app.log
|   config.toml
|   main.go
|   readme.md
|
+---config
|       viper.go
|
+---handlers
|       gin.go
|       handler_wp_litespeed_img_optm.go
|       handler_wp_litespeed_optimizer.go
|       handler_wp_posts.go
|       handler_wp_users.go
|       handler_wp_yoast_seo_links.go
|
+---models
|       db.go
|       model_wp_litespeed_img_optm.go
|       model_wp_litespeed_optimizer.go
|       model_wp_posts.go
|       model_wp_users.go
|       model_wp_yoast_seo_links.go
|
+---static
|   |   .gitignore
|   |   index.html
|   |   readme.md
|   |
|   \---index_files
|           jquery.js.download
|           style.css
|           syntax.css
|
\---swagger
        .gitignore
        doc.yml
        favicon-16x16.png
        favicon-32x32.png
        index.html
        oauth2-redirect.html
        readme.md
        swagger-ui-bundle.js
        swagger-ui-standalone-preset.js
        swagger-ui.css
        swagger-ui.js
```
### Command help
```shell
ginbro gen -h
generate a RESTful APIs app with gin and gorm for gophers. For example:
        ginbro gen -u eric -p password -a "127.0.0.1:3306" -d "mydb"

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
## environment
- my development environment
    - Windows 10 pro 64
    - go version go1.11.1 windows/amd64
    - mysql version <= 5.7

## go packages
```shell
go get github.com/gin-contrib/cors
go get github.com/gin-contrib/static
go get github.com/gin-gonic/autotls
go get github.com/gin-gonic/gin
go get github.com/sirupsen/logrus
go get github.com/spf13/viper
go get github.com/spf13/cobra
go get github.com/go-redis/redis
go get github.com/go-sql-driver/mysql
go get github.com/jinzhu/gorm
```

## Info
- table'schema which has no "ID","id","Id'" or "iD" will not generate model or route.
- the column which type is json value must be a string which is able to decode to a JSON,when call POST or PATCH.
## Thanks
- [swagger Specification](https://swagger.io/specification/)
- [gin-gonic/gin](https://github.com/gin-gonic/gin)
- [GORM](http://gorm.io/)
- [viper](https://github.com/spf13/viper)
- [cobra](https://github.com/spf13/cobra#getting-started)

- [base64captcha](https://github.com/mojocn/base64Captcha)

## Please feedback your [`issue`](https://github.com/dejavuzhou/ginbro/issues) with database schema file