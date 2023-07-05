package dbconn

import (
	"fmt"
	"github.com/spf13/viper"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"gorm.io/gorm/schema"
	"time"
)

type DB struct {
	*gorm.DB
}

var defaultDB DB

func PersistDB() *DB {
	return &defaultDB

}

func DoPersistConnect() error {
	return PersistDB().PersistConnect()
}

func (db *DB) PersistConnect() error {

	if db.DB != nil {
		return nil
	}

	con, err := NewConnection()

	if err != nil {
		return err
	}
	db.DB = con
	return nil
}

func NewConnection() (*gorm.DB, error) {
	dbHost := viper.GetString("database.host")
	dbPort := viper.GetInt("database.port")
	dbName := viper.GetString("database.name")
	dbUser := viper.GetString("database.user")
	dbPassword := viper.GetString("database.password")
	dbSSLMode := viper.GetString("database.ssl_mode")
	dbTimeZone := viper.GetString("database.time_zone")
	schemaName := viper.GetString("database.schema_name")
	tablePrefix := viper.GetString("database.table_prefix")
	level := viper.GetString("log.level")

	if len(schemaName) == 0 {
		schemaName = "pay,public"
	}

	dsn := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=%s TimeZone=%s search_path=%s",
		dbHost, dbPort, dbUser, dbPassword, dbName, dbSSLMode, dbTimeZone, schemaName)

	gConf := &gorm.Config{
		Logger: logger.Default.LogMode(ParseGormLevel(level)),
	}

	if len(tablePrefix) > 0 {
		gConf.NamingStrategy = schema.NamingStrategy{
			TablePrefix: fmt.Sprintf("%s", tablePrefix), // schema name, e.g. "pay."
		}
	}

	con, err := gorm.Open(postgres.New(postgres.Config{
		DSN: dsn,
	}), gConf)

	if err != nil {
		return nil, err
	}

	d, err := con.DB()
	if err != nil {
		return nil, err
	}

	dbMaxOpenConn := viper.GetInt("database.max_open_connections")
	dbMaxIdleConn := viper.GetInt("database.max_idle_connections")

	if dbMaxOpenConn == 0 {
		dbMaxOpenConn = 10
	}
	if dbMaxIdleConn == 0 {
		dbMaxIdleConn = 5
	}
	d.SetConnMaxLifetime(time.Second * 10)
	d.SetMaxOpenConns(dbMaxOpenConn)
	d.SetMaxIdleConns(dbMaxIdleConn)

	return con, err
}
