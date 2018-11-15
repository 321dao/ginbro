[app]
    name = "ginBro"
    env = "local" # only allows local/dev/test/prod
    secret = "{{.AppSecret}}"
    log_level = "error" # only allows debug info warn error fatal panic
    enable_cors = false  # true will case 403 error in swaggerUI  may cause api perform decrease
    enable_sql_log = true # show gorm sql in terminal
    enable_https = false # if addr is a domain enable_https will works
    addr ="{{.AppAddr}}" # eg1: www.mojotv.cn     eg:localhost:3333 eg1:127.0.0.1:88
    time_zone = "Asia/Shanghai"
    api_prefix = "v1" #  api_prefix could be empty string,            the api uri will be api/v1/resource
    enable_swagger = true
    static_path = "./static/"  # path must be an absolute path or relative to the go-build-executable file, may cause api perform decrease
    enable_not_found = true # if true and static_path is not empty string, all not found route will serve static/index.html
[mysql]
    addr = "{{.MysqlAddr}}"
    user = "{{.MysqlUser}}"
    password = "{{.MysqlPassword}}"
    database = "{{.MysqlDatabase}}"
    charset = "{{.MysqlCharset}}"
[redis]
    addr = "" # 127.0.0.1:6379 empty string will not init the redis db in models package
    password = ""
    db_idx = 0
