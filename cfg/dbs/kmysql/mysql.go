package kmysql

import (
	"fmt"
	"time"

	_ "github.com/go-sql-driver/mysql" // MySQL Driver
	"github.com/jmoiron/sqlx"
	"github.com/spf13/viper"
)

// Pool sqlx.DB
var (
	Pool *sqlx.DB
)

// DBConfig Db Config struct
type DBConfig struct {
	Host     string
	Port     string
	Username string
	Password string
	DBName   string
}

// NewConfig generates a mysql configuration object which
// This mysql instance will becomes the single source of
// truth for the app configuration.
func NewConfig() {
	conn, err := getConnection()
	if err != nil {
		panic(err.Error())
	}
	connection := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s", conn.Username, conn.Password, conn.Host, conn.Port, conn.DBName)
	Pool, err = sqlx.Connect("mysql", connection)
	if err != nil {
		panic(err.Error())
	}
	Pool.SetMaxIdleConns(viper.GetInt("APP.DB.MYSQL.IDLECONNECTION"))
	Pool.SetMaxOpenConns(viper.GetInt("APP.DB.MYSQL.MAXOPENCONNECTION"))
	Pool.SetConnMaxLifetime(time.Hour)

	err = Pool.Ping()
	if err != nil {
		panic(err.Error())
	}
}

// GetConn Get Pool
func GetConn() *sqlx.DB {
	return Pool
}

// DBstats Get DB Stats
func DBstats() interface{} {
	return Pool.DB.Stats()
}

// DBstatus Get Conn Ping
func DBstatus() error {
	return Pool.Ping()
}

func getConnection() (conn DBConfig, err error) {
	host := viper.GetString("APP.DB.MYSQL.HOST")
	port := viper.GetString("APP.DB.MYSQL.PORT")
	username := viper.GetString("APP.DB.MYSQL.USERNAME")
	password := viper.GetString("APP.DB.MYSQL.PASSWORD")
	dbName := viper.GetString("APP.DB.MYSQL.NAME")
	conn = DBConfig{
		Host:     host,
		Port:     port,
		Username: username,
		Password: password,
		DBName:   dbName,
	}
	return conn, nil
}
