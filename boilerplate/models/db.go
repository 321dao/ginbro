package models

import (
	"errors"
	"fmt"
	_ "github.com/dejavuzhou/ginbro/boilerplate/config"
	"github.com/go-redis/redis"
	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
	"github.com/sirupsen/logrus"
	"github.com/spf13/viper"
	"reflect"
	"strconv"
	"strings"
)

//redis client
var redisDB *redis.Client
var mysqlDB *gorm.DB

const redisPrefix = "ginbro:"

func init() {
	//initializing redis client
	redisAddr, redisPassword := viper.GetString("redis.addr"), viper.GetString("redis.password")
	if redisAddr != "" {
		redisDB = redis.NewClient(&redis.Options{
			Addr:     redisAddr,
			Password: redisPassword,                // no password set
			DB:       viper.GetInt("redis.db_idx"), // use default DB
		})
		if pong, err := redisDB.Ping().Result(); err != nil || pong != "PONG" {
			logrus.WithError(err).Fatal("could not connect to the redis server")
		}
	}

	//init mysql
	conn := fmt.Sprintf("%s:%s@(%s)/%s?charset=%s&parseTime=True&loc=Local", viper.GetString("mysql.user"),
		viper.GetString("mysql.password"), viper.GetString("mysql.addr"), viper.GetString("mysql.database"),
		viper.GetString("mysql.charset"))
	if db, err := gorm.Open("mysql", conn); err == nil {
		mysqlDB = db
	} else {
		logrus.WithError(err).Fatalln("initialize mysql database failed")
	}
	//enable Gorm mysql log
	if flag := viper.GetBool("app.enable_sql_log"); flag {
		mysqlDB.LogMode(flag)
		//f, err := os.OpenFile("mysql_gorm.log", os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)
		//if err != nil {
		//	logrus.WithError(err).Fatalln("could not create mysql gorm log file")
		//}
		//logger :=  New(f,"", Ldate)
		//mysqlDB.SetLogger(logger)
	}
	//mysqlDB.AutoMigrate()

}

func Close() {
	if mysqlDB != nil {
		mysqlDB.Close()
	}
	if redisDB != nil {
		redisDB.Close()
	}
}

type PaginationQuery struct {
	Where  string `form:"where"`
	Fields string `form:"fields"`
	Order  string `form:"order"`
	Offset uint   `form:"offset"`
	Limit  uint   `form:"limit"`
}

func (pq *PaginationQuery) String() string {
	return fmt.Sprintf("w=%v_f=%s_o=%s_of=%d_l=%d", pq.Where, pq.Fields, pq.Order, pq.Offset, pq.Limit)
}

func crudAll(m interface{}, q *PaginationQuery, list interface{}) (total uint, err error) {
	var tx *gorm.DB
	total, tx = getResourceCount(m, q)
	if q.Fields != "" {
		columns := strings.Split(q.Fields, ",")
		if len(columns) > 0 {
			tx = tx.Select(q.Fields)
		}
	}
	if q.Order != "" {
		tx = tx.Order(q.Order)
	}
	if q.Offset > 0 {
		tx = tx.Offset(q.Offset)
	}
	if q.Limit <= 0 {
		q.Limit = 15
	}
	err = tx.Limit(q.Limit).Find(list).Error
	return
}

func crudOne(m interface{}, one interface{}) (err error) {
	if mysqlDB.Where(m).First(one).RecordNotFound() {
		return errors.New("resource is not found")
	}
	return nil
}

func crudUpdate(m interface{}, where interface{}) (err error) {
	db := mysqlDB.Model(where).Updates(m)
	if err = db.Error; err != nil {
		return
	}
	if db.RowsAffected != 1 {
		return errors.New("id is invalid and resource is not found")
	}
	return nil
}

func crudDelete(m interface{}) (err error) {
	//WARNING When delete a record, you need to ensure it’s primary field has value, and GORM will use the primary key to delete the record, if primary field’s blank, GORM will delete all records for the model
	//primary key must be not zero value
	db := mysqlDB.Delete(m)
	if err = db.Error; err != nil {
		return
	}
	if db.RowsAffected != 1 {
		return errors.New("resource is not found to destroy")
	}
	return nil
}
func getResourceCount(m interface{}, q *PaginationQuery) (uint, *gorm.DB) {
	var tx = mysqlDB.Model(m)
	conditions := strings.Split(q.Where, ",")
	for _, val := range conditions {
		w := strings.SplitN(val, ":", 2)
		if len(w) == 2 {
			bindKey, bindValue := w[0], w[1]
			if intV, err := strconv.ParseInt(bindValue, 10, 64); err == nil {
				// bind value is int
				field := fmt.Sprintf("`%s` > ?", bindKey)
				tx = tx.Where(field, intV)
			} else if fV, err := strconv.ParseFloat(bindValue, 64); err == nil {
				// bind value is float
				field := fmt.Sprintf("`%s` > ?", bindKey)
				tx = tx.Where(field, fV)
			} else if bindValue != "" {
				// bind value is string
				field := fmt.Sprintf("`%s` LIKE ?", bindKey)
				sV := fmt.Sprintf("%%%s%%", bindValue)
				tx = tx.Where(field, sV)
			}
		}
	}
	modelName := getType(m)
	rKey := redisPrefix + modelName + q.String() + "_count"

	if v, err := mem.GetUint(rKey); err != nil {
		var count uint = 0
		tx.Count(&count)
		mem.Set(rKey, count)
		return count, tx
	} else {
		return v, tx
	}
}

func getType(v interface{}) string {
	if t := reflect.TypeOf(v); t.Kind() == reflect.Ptr {
		return "*" + t.Elem().Name()
	} else {
		return t.Name()
	}
}
