package config

import (
	"github.com/jinzhu/configor"
	"github.com/kataras/iris/v12"
	"log"
	"time"
)

var Setting *Config

const (
	DefaultBindAddress       = "0.0.0.0"
	DefaultListenPort        = 8080
	DefaultMysqlHost         = "127.0.0.1"
	DefaultMysqlPort         = "3306"
	DefaultMysqlDatabase     = "test"
	DefaultMysqlUser         = "root"
	DefaultMysqlPassword     = "123456"
	DefaultMysqlCharset      = "utf8"
	DefaultMysqlMaxIdleConns = 20
	DefaultMysqlMaxOpenConns = 200
	DefaultMysqlMaxConnLifetime = 60
	DefaultRedisHost         = "127.0.0.1:6379"
	//Redis password
	DefaultRedisPassword = ""
	//Number of idle connections in redis connection pool
	DefaultRedisMaxIdle = 10
	//The maximum number of active connections in redis connection pool, 0 is unlimited
	DefaultRedisMaxActive = 0
	//idle time out
	DefaultRedisDIdleTimeout = 0
	//maximum connection
	DefaultRedisMaxConnLifetime = 0 //Active time
)

// Config application configuration
type Config struct {
	App      AppConfig `yaml:"app"`
	Mysql    MysqlConfig `yaml:"mysql"` // mysql configuration
	Redis    RedisConfig `yaml:"redis"` // redis configuration
	Iris     iris.Configuration `yaml:"iris"`
	Log      LogConfig `yaml:"log"`
}

type AppConfig struct {
	Debug       bool `yaml:"debug"`
	Timezone    string `yaml:"timezone"`
	BindAddress string `yaml:"bindAddress"` // Server listening address
	Port        int `yaml:"port"`    // Server listening port
}
type MysqlConfig struct {
	Host         string `yaml:"host"`
	Port         string `yaml:"port"`
	Database     string `yaml:"database"`
	User         string `yaml:"user"`
	Password     string `yaml:"password"`
	Charset      string `yaml:"charset"`
	MaxIdleConns int `yaml:"maxIdleConns"`
	MaxOpenConns int `yaml:"maxOpenConns"`
	MaxConnLifetime time.Duration `yaml:"maxConnLifeTime"`
	SqlLog       bool `yaml:"sqlLog"`
}
type RedisConfig struct {
	Host            string `yaml:"host"`
	Password        string `yaml:"password"`
	MaxIdle         int `yaml:"maxidle"` // Connection pool maximum idle connections
	MaxActive       int `yaml:"maxactive"` // Connection pool maximum active connections
	IdleTimeout     time.Duration `yaml:"idletimeout"`
	Wait            bool `yaml:"wait"`          //wait for connection
	MaxConnLifetime time.Duration `yaml:"maxconnlifetime"` //Active time
	Debug           bool `yaml:"debug"`          //debug open
}

type LogConfig struct {
	Level    string `yaml:"level" default:"debug"`
	Path     string `yaml:"path"  default:"storage/logs"`
	FileName string `yaml:"filename"  default:"iris-demo-new"`
}


func Init(path string) {
	Setting = &Config{}
	if path == "" {
		Setting.initDefaultConfig()
		return
	}

	Setting.parse(path)
}

// Resolve profile
func (config *Config) parse(path string) {
	err := configor.Load(config, path)
	if err != nil {
		log.Fatalf("Unable to resolve profile#%s", err.Error())
	}
}

// Initialize default configuration
func (config *Config) initDefaultConfig() {
	config.App.Debug = false
	config.App.Timezone = "UTC"
	config.App.BindAddress = DefaultBindAddress
	config.App.Port = DefaultListenPort
	//Default MySQL configuration
	config.Mysql.Host = DefaultMysqlHost
	config.Mysql.Port = DefaultMysqlPort
	config.Mysql.Database = DefaultMysqlDatabase
	config.Mysql.User = DefaultMysqlUser
	config.Mysql.Password = DefaultMysqlPassword
	config.Mysql.Charset = DefaultMysqlCharset
	config.Mysql.MaxIdleConns = DefaultMysqlMaxIdleConns
	config.Mysql.MaxOpenConns = DefaultMysqlMaxOpenConns
	config.Mysql.MaxConnLifetime = DefaultMysqlMaxConnLifetime
	config.Mysql.SqlLog = false
	//Default redis configuration
	config.Redis.Host = DefaultRedisHost
	config.Redis.Password = DefaultRedisPassword
	config.Redis.MaxIdle = DefaultRedisMaxIdle
	config.Redis.MaxIdle = DefaultRedisMaxIdle
	config.Redis.IdleTimeout = DefaultRedisDIdleTimeout
	config.Redis.Wait = false
	config.Redis.MaxConnLifetime = DefaultRedisMaxConnLifetime
	config.Redis.Debug = false
	config.Log.Path = "/storage/logs"
	config.Log.Level = "error"
	config.Log.FileName = "iris-demo-new"
}