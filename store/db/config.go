package db

import (
	"fmt"
	"log"
	"os"
	"time"

	mysql "gorm.io/driver/mysql"
	postgres "gorm.io/driver/postgres"
	sqlite "gorm.io/driver/sqlite"
	gorm "gorm.io/gorm"
	gormlog "gorm.io/gorm/logger"
)

type Config struct {
	Dialect  string `json:"dialect"` // mysql,postgres,sqlite3
	Host     string `json:"host"`    // if Dialect is `sqlite3`, host should be db file path
	ReadHost string `json:"read_host"`
	Port     int    `json:"port"`
	User     string `json:"user"`
	Password string `json:"password"`
	Database string `json:"database"`
	Location string `json:"location,omitempty"`
	SSLMode  string `json:"sslmode,omitempty"` // postgres has this property, which could be "disable"
	Debug    bool   `json:"debug,omitempty"`
}

func SqliteInMemory() Config {
	return Config{
		Dialect: "sqlite3",
		Host:    ":memory:",
	}
}

func Connect(dialect, uri string) (*DB, error) {
	coon, err := openWithDSN(dialect, uri)
	if err != nil {
		return nil, err
	}

	return &DB{
		write: coon,
		read:  coon,
	}, nil
}

func gormLog() gormlog.Interface {
	logger := gormlog.New(
		log.New(os.Stdout, "\r\n", log.LstdFlags), // io writer
		gormlog.Config{
			SlowThreshold: time.Second,    // 慢 SQL 阈值
			LogLevel:      gormlog.Silent, // Log level
			Colorful:      true,           // 禁用彩色打印
		},
	)

	return logger
}

func open(dialect string, host string, port int, user, password, database, loc, sslmode string) (*gorm.DB, error) {
	var dsn string
	switch dialect {
	case "mysql":
		dsn = fmt.Sprintf("%s:%s@%s(%s:%d)/%s?parseTime=True&charset=utf8mb4,utf8&loc=%s",
			user,
			password,
			"tcp",
			host,
			port,
			database,
			loc,
		)
	case "postgres":
		dsn = fmt.Sprintf("host=%s port=%d user=%s dbname=%s password=%s sslmode=%s",
			host,
			port,
			user,
			database,
			password,
			sslmode,
		)
	case "sqlite3", "sqlite":
		dialect = "sqlite3"
		dsn = host
	default:
		return nil, fmt.Errorf("unkonow db dialect: %s", dialect)
	}

	return openWithDSN(dialect, dsn)
}

func openWithDSN(dialect, dsn string) (*gorm.DB, error) {
	var dialector gorm.Dialector
	switch dialect {
	case "mysql":
		dialector = mysql.Open(dsn)
	case "postgres":
		dialector = postgres.Open(dsn)
	case "sqlite3", "sqlite":
		dialector = sqlite.Open(dsn)
	default:
		return nil, fmt.Errorf("unkonow db dialect: %s", dialect)
	}

	db, err := gorm.Open(dialector, &gorm.Config{
		Logger: gormLog(),
	})
	if err != nil {
		return nil, err
	}

	return db, nil
}

func Open(cfg Config) (*DB, error) {
	write, err := open(cfg.Dialect, cfg.Host, cfg.Port, cfg.User, cfg.Password, cfg.Database, cfg.Location, cfg.SSLMode)
	if err != nil {
		return nil, err
	}

	db := &DB{
		write: write,
		read:  write,
	}

	if cfg.ReadHost != "" && cfg.ReadHost != cfg.Host {
		db.read, err = open(cfg.Dialect, cfg.Host, cfg.Port, cfg.User, cfg.Password, cfg.Database, cfg.Location, cfg.SSLMode)
		if err != nil {
			return nil, err
		}
	}

	if cfg.Debug {
		db = db.Debug()
	}

	return db, nil
}

func MustOpen(cfg Config) *DB {
	db, err := Open(cfg)
	if err != nil {
		panic(fmt.Errorf("open db failed: %w", err))
	}

	return db
}
