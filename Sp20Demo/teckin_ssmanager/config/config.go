package config

import (
	"fmt"
	"gopkg.in/yaml.v2"
	"io/ioutil"
	"os"
	"path/filepath"
)

var Conf *Config

type Config struct {
	Kafka  *KafKaConfig      `yaml:"kafka"`
	Mysql  *MysqlConfig      `yaml:"mysql"`
	Redis  *RedisConfig      `yaml:"redis"`
	Http   *HttpServerConfig `yaml:"http"`
	Awsiot *AwsIotConfig     `yaml:"awsiot"`
	Grpc   *GrpcServerConfig `yaml:"grpc"`
}

//kafka配置信息
type KafKaConfig struct {
	Brokers           string    `yaml:"brokers"` //kafka集群连接地址，多个地址使用分号(;)隔开
	Version           string    `yaml:"version"`
	Group             string    `yaml:"group"`
	ChannelBufferSize int       `yaml:"channelbuffersize"`
	Ready             chan bool `yaml:"ready"`
}

type MysqlConfig struct {
	Addr     string `yaml:"addr"`
	Port     int    `yaml:"port"`
	User     string `yaml:"user"`
	Password string `yaml:"password"`
	DBName   string `yaml:"dbname"`
}

type RedisConfig struct {
	Addr     string `yaml:"addr"`
	Password string `yaml:"password"`
	DB       int    `yaml:"db"`
}

type HttpServerConfig struct {
	Addr string `yaml:"addr"`
}

type GrpcServerConfig struct {
	Addr string `yaml:"addr"`
}

type AwsIotConfig struct {
	AkId      string `yaml:"akid"`      // aws秘钥key
	SecretKey string `yaml:"secretKey"` // aws秘钥secret
	Region    string `yaml:"region"`    // 服务区域
	Endpoint  string `yaml:"endpoint"`  // 地址
}

func Init() error {
	gen, _ := os.Getwd()
	path := filepath.Join(gen, "config/teckin.yaml")
	yamlFile, err := ioutil.ReadFile(path)
	if err != nil {
		return err
	}
	Conf = new(Config)
	fmt.Println(string(yamlFile))
	err = yaml.Unmarshal(yamlFile, Conf)
	return err
}

func DefaultConfig() {
	Conf = &Config{
		Kafka: &KafKaConfig{
			Brokers:           "b-1.teckin.jopm4n.c2.kafka.cn-north-1.amazonaws.com.cn:9092;b-2.teckin.jopm4n.c2.kafka.cn-north-1.amazonaws.com.cn:9092",
			Version:           "2.6.1",
			Group:             "teckin_test",
			ChannelBufferSize: 256,
		},
		Mysql: &MysqlConfig{
			Addr:     "127.0.0.1",
			Port:     3306,
			User:     "root",
			Password: "root",
			DBName:   "teckin",
		},
		Redis: &RedisConfig{
			Addr: "127.0.0.1:6379",
			DB:   0,
		},
		Http: &HttpServerConfig{
			Addr: ":15000",
		},
		Grpc: &GrpcServerConfig{
			Addr: ":21000",
		},
	}
}
