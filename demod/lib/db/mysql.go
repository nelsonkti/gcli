package db

import (
	"database/sql"
	"demod/config"
	pb "demod/config/pb"
	applogger "demod/lib/logger"
	"fmt"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"gorm.io/plugin/dbresolver"
	"sync"
	"time"
)

var Databases sync.Map

const defaultConfig = "?parseTime=true&charset=utf8mb4&loc=Asia%2FShanghai"

func InitMysql() {
	for _, conf := range config.AppConf.Data.Connection.Database {
		newDatabase().connect(conf)
	}
}

func Mysql(name string) *gorm.DB {
	db, _ := Databases.Load(name)
	return db.(*gorm.DB)
}

func DisconnectMysql() {

	Databases.Range(func(key, value interface{}) bool {
		db, _ := value.(*gorm.DB)
		sqlDB, _ := db.DB()
		_ = sqlDB.Close()
		return true
	})

	applogger.Sugar.Info("disconnect mysql")
}

func newDatabase() *database {
	db := &database{}

	db.slave = &slave{}
	db.dbConfig = db.defaultConfig
	db.ping = db.defaultPing
	db.DB = db.copyDb

	return db
}

type nilFunc func()

type database struct {
	db       *gorm.DB
	sqlDb    *sql.DB
	database *pb.Data_Database
	dbConfig nilFunc
	slave    *slave
	ping     nilFunc
	DB       func() *database
}

func (d *database) connect(database *pb.Data_Database) {

	d.database = database
	dsn := tcpSprint(database, database.Host)

	var err error
	d.db, err = gorm.Open(mysql.Open(dsn), d.config())

	if err != nil {
		applogger.Sugar.Info(err)
		panic(err)
	}

	d.slave.database(d).connect(dsn)

	d.DB().ping()

	d.dbConfig()

	Databases.Store(database.Database, d.db)
}

func (d *database) config() *gorm.Config {
	logLevel := logger.Silent

	if config.AppConf.App.Env == "local" {
		logLevel = logger.Info
	}

	return &gorm.Config{
		Logger: logger.Default.LogMode(logLevel),
	}
}

func (d *database) connectSlave(dsn string) {
	if d.database.Read == nil {
		return
	}

	var dialect []gorm.Dialector
	for _, v := range d.database.Read {
		dialect = append(dialect, mysql.Open(tcpSprint(d.database, v)))
	}

	err := d.db.Use(dbresolver.Register(dbresolver.Config{
		Sources:  []gorm.Dialector{mysql.Open(dsn)},
		Replicas: dialect,
		//  负载均衡策略
		Policy: dbresolver.RandomPolicy{},
	}))

	if err != nil {
		panic(err)
	}

}

func (d *database) copyDb() *database {
	var err error
	d.sqlDb, err = d.db.DB()

	if err != nil {
		applogger.Sugar.Info("failed to connect mysql:" + d.database.Database)
		panic("failed to connect mysql:" + d.database.Database)
	}

	return d
}

func (d *database) defaultPing() {
	err := d.sqlDb.Ping()
	if err != nil {
		applogger.Sugar.Info("failed to connect mysql:" + d.database.Database)
		panic("failed to connect mysql:" + d.database.Database)
	}
}

func (d *database) defaultConfig() {

	d.sqlDb.SetMaxIdleConns(1024)

	d.sqlDb.SetMaxOpenConns(1024)

	d.sqlDb.SetConnMaxLifetime(time.Minute * 10)

}

func tcpSprint(conf *pb.Data_Database, network string) string {
	return fmt.Sprintf("%s:%s@tcp(%s:%d)/%s%s", conf.Username, conf.Password, network, conf.Port, conf.Database, defaultConfig)
}

type slave struct {
	db *database
}

func (s *slave) database(d *database) (slave *slave) {
	s.db = d
	return s
}

func (s *slave) connect(dsn string) {
	db := *s.db
	if db.database.Read == nil {
		return
	}

	var dialect []gorm.Dialector
	for _, v := range db.database.Read {
		dialect = append(dialect, mysql.Open(tcpSprint(db.database, v)))
	}

	err := db.db.Use(dbresolver.Register(dbresolver.Config{
		Sources:  []gorm.Dialector{mysql.Open(dsn)},
		Replicas: dialect,
		//  负载均衡策略
		Policy: dbresolver.RandomPolicy{},
	}))

	if err != nil {
		panic(err)
	}
}
