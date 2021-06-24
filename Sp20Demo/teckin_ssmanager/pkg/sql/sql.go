package sql

import "teckin_ssmanager/config"

const (
	MysqlType = 0
)

type SqlFactory interface {
	Open(conf *config.MysqlConfig) error
	Query(sqlStr string)([]map[string]string, error)
	Update(sqlStr string) error
	Delete(sqlStr string) error
	Insert(sqlStr string) (int64, error)
	Close()
}

func NewSql(nType int) SqlFactory {
	switch nType {
	case MysqlType:
		return new(Mysql)
	}
	return nil
}