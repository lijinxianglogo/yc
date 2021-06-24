package sql

import (
	"context"
	"errors"
	"github.com/go-redis/redis/v8"
	"github.com/sirupsen/logrus"
	"teckin_ssmanager/config"
	"time"
)

type Redis struct {
	Addr         string
	Password     string
	conn         *redis.Client
	PoolSize     int // 连接池最大socket连接数，默认为4倍CPU数， 4 * runtime.NumCPU
	MinIdleConns int //在启动阶段创建指定数量的Idle连接，并长期维持idle状态的连接数不少于指定数量；。

	DialTimeout  int //连接建立超时时间，默认5秒。
	ReadTimeout  int //读超时，默认3秒， -1表示取消读超时
	WriteTimeout int //写超时，默认等于读超时
	PoolTimeout  int //当所有连接都处在繁忙状态时，客户端等待可用连接的最大等待时长，默认为读超时+1秒
	DB           int

	logger  *logrus.Entry
	CloseCh chan struct{}
}

func NewRedis(conf *config.RedisConfig, l *logrus.Entry) *Redis {
	return &Redis{
		Addr:     conf.Addr,
		Password: conf.Password,
		conn:     nil,
		DB:       conf.DB,
		logger:   l,
	}
}

func (r *Redis) Open() error {
	if err := r.dial(); err != nil {
		panic(err)
	}
	go r.checkDial()
	return nil
}

func (r *Redis) dial() error {
	if r.PoolSize == 0 {
		r.PoolSize = 15
	}
	r.conn = redis.NewClient(&redis.Options{
		Addr:     r.Addr,
		Password: r.Password,
		DB:       r.DB,
		PoolSize: r.PoolSize,
	})

	if err := r.ping(); err != nil {
		return err
	}
	return nil
}

func (r *Redis) checkDial() {
	for {
		select {
		case <-r.CloseCh:
			r.conn.Close()
			return
		default:
			if err := r.ping(); err != nil {
				r.logger.Errorf(err.Error())
			}
		}
		time.Sleep(time.Second * 5)
	}
}

func (r *Redis) ping() error {
	if pong, _ := r.Ping(); pong != "PONG" {
		return errors.New("redis 连接失败")
	}
	return nil
}

func (r *Redis) checkConn() error {
	if r.conn == nil {
		return errors.New("连接对象为空")
	}
	return nil
}

func (r *Redis) Set(key string, value interface{}, expire time.Duration) (err error) {
	_, err = r.conn.Set(context.Background(), key, value, expire).Result()
	return
}

func (r *Redis) Get(key string) (string, error) {
	return r.conn.Get(context.Background(), key).Result()
}

func (r *Redis) Del(key string) error {
	_, err := r.conn.Del(context.Background(), key).Result()
	return err
}

func (r *Redis) HSet(key string, values ...interface{}) error {
	_, err := r.conn.HSet(context.Background(), key, values).Result()
	return err
}

func (r *Redis) HGetAll(key string) (map[string]string, error) {
	data, err := r.conn.HGetAll(context.Background(), key).Result()
	return data, err
}

//key不存在返回-2,存在没有ttl返回-1, 否则返回剩余ttl秒数
func (r *Redis) TTL(key string) (time.Duration, error) {
	return r.conn.TTL(context.Background(), key).Result()
}

//key不存在返回false,否则返回true
func (r *Redis) Expire(key string, expire time.Duration) (bool, error) {
	return r.conn.Expire(context.Background(), key, expire).Result()
}

func (r *Redis) Ping() (string, error) {
	return r.conn.Ping(context.Background()).Result()
}

func (r *Redis) Close() {
	r.CloseCh <- struct{}{}
}
