package config

import (
	"context"
	"fmt"
	"os"
	"path/filepath"

	"github.com/fsnotify/fsnotify"
	"github.com/google/uuid"
	"github.com/spf13/viper"
	v1 "github.com/yituoshiniao/kit/xdb/v1"
	"github.com/yituoshiniao/kit/xlog"
	"github.com/yituoshiniao/kit/xrds"
	"github.com/yituoshiniao/kit/xtask"
	"github.com/yituoshiniao/kit/xtrace"
)

type LoggerConfig struct {
	Network string `yaml:"network" json:"network"`
	Host    string `yaml:"host" json:"host"`
	Port    int    `yaml:"port" json:"port"`
	Tag     string `yaml:"tag" json:"tag"`
}

type ServiceConfig struct {
	Name    string   `yaml:"name" json:"name"`
	Backups []string `yaml:"backups" json:"backups"`
}

type RegistryConfig struct {
	Addrs  string `yaml:"addrs" json:"addrs"`
	Prefix string `yaml:"prefix" json:"prefix"`
	Token  string `yaml:"token" json:"token"`
}

type MonitorConfig struct {
	Enable bool   `yaml:"enable" json:"enable"`
	Host   string `yaml:"host" json:"host"`
	Port   int    `yaml:"port" json:"port"`
}

type XTraceConfig struct {
	// #上报配置类型
	SamplerType string
	// #常量配置
	SamplerParam float64
	// #数据服务器地址
	ReporterLocalAgentHostPort string
	// 服务名
	ServerName string
}

type TraceConfig struct {
	Enable   bool    `yaml:"enable" json:"enable"`
	Host     string  `yaml:"host" json:"host"`
	Port     int     `yaml:"port" json:"port"`
	Sampling float64 `yaml:"sampling" json:"sampling"`
}

type Config struct {
	Mode       string            `yaml:"mode" json:"mode"`
	LocalCache LocalCache        `yaml:"localCache" json:"localCache"`
	Bind       string            `yaml:"bind" json:"bind"`
	Name       string            `yaml:"name" json:"name"`
	Monitor    MonitorConfig     `yaml:"monitor" json:"monitor"`
	Asynq      xtask.AsynqConfig `yaml:"asynq" json:"asynq"`
	Trace      TraceConfig       `yaml:"trace" json:"trace"`
	Area       string            `yaml:"area" json:"area"`
	GobPath    string            `yaml:"gobPath" json:"gobPath"`

	XLogger                 xlog.Config  `yaml:"xLogger" json:"xLogger"`
	XTrace                  XTraceConfig `yaml:"xTrace" json:"xTrace"`
	IDService               Upstream     `yaml:"idService" json:"idService"`
	XRedis                  xrds.Config  `yaml:"xRedis" json:"xRedis"`
	Database                Database     `yaml:"database" json:"-"`
	AppStoreConnectApi      Upstream     `yaml:"appStoreConnectApi" json:"appStoreConnectApi"`
	AppStoreServerApi       Upstream     `yaml:"appStoreServerApi" json:"appStoreServerApi"`
	DbProxyApi              Upstream     `yaml:"dbProxyApi" json:"dbProxyApi"`
	ImportTerritoryCodeFile string       `yaml:"importTerritoryCodeFile" json:"importTerritoryCodeFile"`
	GpConnectApi            Upstream     `yaml:"gpConnectApi" json:"gpConnectApi"`
	GpConnectApiTokenURL    string       `yaml:"gpConnectApiTokenURL" json:"gpConnectApiTokenURL"`
	SentinelConfFile        string       `yaml:"sentinelConfFile" json:"sentinelConfFile"`
	CsUserIp                []string     `yaml:"csUserIp" json:"csUserIp"`
}

type Database struct {
	GoodsCenter v1.Config `yaml:"goodsCenter"`
	Camexam     v1.Config `yaml:"camexam"`
	// 解决viper读取yaml配置存在下划线时无法映射 添加tag: mapstructure
	DbToken v1.Config `yaml:"db_token" mapstructure:"db_token"`
}

type LocalCache struct {
	Enable           bool
	LifeWindow       int64
	CleanWindow      int64
	HardMaxCacheSize int
}

type Upstream struct {
	Host    string
	Timeout int64
	Prefix  string
}

var (
	// GlobalConfig 全局配置变量
	GlobalConfig Config
)

func ParseConfig(envPath string) Config {
	pwd, _ := os.Getwd()
	dirName := filepath.Base(pwd)
	// read config
	viper.SetConfigFile("./env.yaml") // 指定配置文件路径
	viper.SetConfigName("env")        // 配置文件名称(无扩展名)

	// 多环境变量
	if envPath != "" {
		viper.SetConfigFile("./env_" + envPath + ".yaml") // 指定配置文件路径
		viper.SetConfigName("env_" + envPath)             // 配置文件名称(无扩展名)
	}
	viper.SetConfigType("yaml")                                 // 如果配置文件的名称中没有扩展名，则需要配置此项
	viper.AddConfigPath(fmt.Sprintf("/apps/conf/%s/", dirName)) // 查找配置文件所在的路径
	viper.AddConfigPath("./config")                             // 还可以在config目录中查找配置

	err := viper.ReadInConfig() // 读取配置数据
	if err != nil {
		panic(fmt.Sprintf("Fatal error config file: %s", err))
		// log.Fatalf("Fatal error config file: %s", err)
	}

	dynamicReloadConfig(context.Background()) // 动态加载配置

	var cfg Config
	err = viper.Unmarshal(&cfg) // 查找并读取配置文件
	if err != nil {
		// 处理读取配置文件的错误
		panic(fmt.Sprintf("Unmarshal error config file: %s", err))
		// log.Fatal(ctx, err)
	}
	GlobalConfig = cfg
	return cfg
}

func dynamicReloadConfig(ctx context.Context) {
	ctx = xtrace.NewCtxWithTraceId(ctx)
	uuid := uuid.New().String()
	viper.WatchConfig()
	viper.OnConfigChange(func(event fsnotify.Event) {
		xlog.S(ctx).Infow("配置发生改变", "event", event.String(), "uuid", uuid)
		err := viper.ReadInConfig() // 读取配置数据
		if err != nil {
			xlog.S(ctx).Errorw("读取配置文件错误", "err", err, "uuid", uuid)
			return
		}

		var tmpCfg Config
		if err := viper.Unmarshal(&tmpCfg); err != nil {
			xlog.S(ctx).Errorw("加载配置错误", "err", err, "uuid", uuid)
			// panic(fmt.Errorf("配置重载失败:%s\n", err))
			return
		}
		GlobalConfig = tmpCfg
		xlog.S(ctx).Infow("GlobalConfig-配置发生改变", "config", GlobalConfig, "uuid", uuid)
	})
}
