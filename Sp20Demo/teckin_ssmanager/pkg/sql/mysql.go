package sql

import (
	"database/sql"
	"errors"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"teckin_ssmanager/config"
	"time"
)

type Mysql struct {
	rdconn *sql.DB
	wdconn *sql.DB
	close  chan struct{}
}

func (s *Mysql) Open(conf *config.MysqlConfig) (err error) {
	err = s.initRd(conf)
	err = s.initWd(conf)
	go s.checkPing(conf)
	return err
}

func (s *Mysql) initRd(conf *config.MysqlConfig) error {
	var err error
	//dsn := fmt.Sprint("%s:%s@tcp(%s:%d)/%s?charset=utf8mb4&parseTime=True&loc=Local", s.User, s.Password, s.Address, s.Port, s.DB)

	// TODO 测试用，直接写好连接地址，日后删掉
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=utf8mb4&parseTime=True&loc=Local", conf.User, conf.Password, conf.Addr, conf.Port, conf.DBName)

	s.rdconn, err = sql.Open("mysql", dsn)
	if err != nil {
		return errors.New(fmt.Sprintf("读数据库连接失败：%v", err))
	}
	return nil
}

func (s *Mysql) initWd(conf *config.MysqlConfig) error {
	var err error
	//dsn := fmt.Sprint("%s:%s@tcp(%s:%d)/%s?charset=utf8mb4&parseTime=True&loc=Local", s.User, s.Password, s.Address, s.Port, s.DB)

	// 测试用，直接写好连接地址
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=utf8mb4&parseTime=True&loc=Local", conf.User, conf.Password, conf.Addr, conf.Port, conf.DBName)
	s.wdconn, err = sql.Open("mysql", dsn)
	if err != nil {
		return errors.New(fmt.Sprintf("写数据库连接失败：%v", err))
	}
	return nil
}

//每隔五秒检查读写对象连接
func (s *Mysql) checkPing(conf *config.MysqlConfig) {
	for {
		select {
		case <-s.close:
			break
		default:
			s.rdPing(conf)
			s.wdPing(conf)
		}
		time.Sleep(time.Second * 5)
	}
}

func (s *Mysql) rdPing(conf *config.MysqlConfig) {
	if err := s.rdconn.Ping(); err != nil {
		_ = s.initRd(conf)
	}
}

func (s *Mysql) wdPing(conf *config.MysqlConfig) {
	if err := s.rdconn.Ping(); err != nil {
		_ = s.initRd(conf)
	}
}

func (s *Mysql) Query(sqlStr string) ([]map[string]string, error) {
	if s.rdconn == nil {
		return nil, errors.New("读连接对象为空")
	}
	rows, err := s.rdconn.Query(sqlStr)
	if err != nil {
		return nil, errors.New(fmt.Sprintf("Query SQL %s Faild !", sqlStr))
	}
	defer rows.Close()
	columns, _ := rows.Columns()
	scanArgs := make([]interface{}, len(columns))
	values := make([]interface{}, len(columns))
	for i := range values {
		scanArgs[i] = &values[i]
	}
	var Result []map[string]string
	for rows.Next() {
		//将行数据保存到record字典
		err = rows.Scan(scanArgs...)
		record := make(map[string]string)
		for i, col := range values {
			if col != nil {
				record[columns[i]] = string(col.([]byte))
			}
		}
		Result = append(Result, record)
	}
	return Result, nil
}

func (s *Mysql) Update(sqlStr string) error {
	_, err := s.wdconn.Exec(sqlStr)
	if err != nil {
		return err
	}
	return nil
}

func (s *Mysql) Delete(sqlStr string) error {
	_, err := s.wdconn.Exec(sqlStr)
	if err != nil {
		return err
	}
	return nil
}

func (s *Mysql) Insert(sqlStr string) (int64, error) {
	result, err := s.wdconn.Exec(sqlStr)
	if err != nil {
		return 0, err
	}
	id, _ := result.LastInsertId()
	return id, nil
}

func (s *Mysql) Close() {
	s.close <- struct{}{}
}
