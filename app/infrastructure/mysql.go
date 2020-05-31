package infra

import (
	"fmt"

	"github.com/jmoiron/sqlx"
	"github.com/kelseyhightower/envconfig"

	// mysql driver
	_ "github.com/go-sql-driver/mysql"
)

// MySQLConfig MySQLの設定情報
type MySQLConfig struct {
	Host     string `envconfig:"MYSQL_HOST" required:"true"`
	Port     string `envconfig:"MYSQL_PORT" required:"true"`
	User     string `envconfig:"MYSQL_USER" required:"true"`
	Password string `envconfig:"MYSQL_PASSWORD" required:"true"`
	Database string `envconfig:"MYSQL_DATABASE" required:"true"`
}

// LoadMySQLConfig 設定情報読み込み
func LoadMySQLConfig() (*MySQLConfig, error) {
	config := &MySQLConfig{}
	if err := envconfig.Process("", config); err != nil {
		return nil, err
	}
	return config, nil
}

// NewMySQLDB DBクライアントの取得
func NewMySQLDB(config *MySQLConfig) (*sqlx.DB, error) {
	conf := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s", config.User, config.Password, config.Host, config.Port, config.Database) + "?parseTime=true&loc=Asia%2FTokyo&charset=utf8mb4"
	return sqlx.Connect("mysql", conf)
}

// NewDefaultMySQLDB デフォルト設定を使用したクライアントの取得
func NewDefaultMySQLDB() (*sqlx.DB, error) {
	config, err := LoadMySQLConfig()
	if err != nil {
		return nil, err
	}
	return NewMySQLDB(config)
}
