package dbs

import (
	"errors"
	"fmt"
	"teckin_ssmanager/internal/ssmanager/model"
	scv "teckin_ssmanager/pkg/strconv"
	"time"
)

//创建用户
func (s *MysqlDB) CreateUser(user model.User) (id int64, err error) {
	sqlStr := fmt.Sprintf("insert into t_user(name, code, email, password, add_time, is_del) values (" +
		"'%s', '%s', '%s', '%s', %d, 1);", user.Name, user.Key, user.Email, user.Password, time.Now().Unix())
	id, err = s.conn.Insert(sqlStr)
	return id, err
}

func(s *MysqlDB) GetUserInfoByEmailAndPwd(email, pwd string) (model.User, error){
	sqlStr := fmt.Sprintf("select id,name, code from t_user where email='%s' and password='%s';", email, pwd)
	mData, err := s.conn.Query(sqlStr)
	if err != nil{
		return model.User{}, err
	}
	if len(mData) <= 0{
		return model.User{}, errors.New("密码或邮箱错误")
	}
	for _, data := range mData{
		var user model.User
		user.ID = scv.S2I(data["id"])
		user.Name = data["name"]
		user.Key = data["code"]
		return user, nil
	}
	return model.User{}, nil
}

func(s *MysqlDB)CheckUserExist(email string) error {
	sqlStr := fmt.Sprintf("select id,name, code from t_user where email='%s';", email)
	mData, err := s.conn.Query(sqlStr)
	if err != nil{
		return errors.New("未知错误：" + err.Error())
	}
	if len(mData) == 0{
		return nil
	}
	return errors.New("邮箱已注册")

}

func(s *MysqlDB) GetUserInfoByUserID(id int) (model.User, error){
	sqlStr := fmt.Sprintf("select email, name, password, code from t_user where id=%d;", id)
	mData, err := s.conn.Query(sqlStr)
	if err != nil{
		return model.User{}, err
	}
	if len(mData) <= 0{
		return model.User{}, errors.New("没有此用户信息")
	}
	for _, data := range mData{
		var user model.User
		user.Password = data["password"]
		user.Name = data["name"]
		user.Email = data["email"]
		user.Key = data["code"]
		return user, nil
	}
	return model.User{}, nil
}

func (s *MysqlDB) UpdateUserPassword(newPwd string, id int) error{
	sqlStr := fmt.Sprintf("update t_user set password='%s' where id=%d;", newPwd, id)
	return s.conn.Update(sqlStr)
}

func (s *MysqlDB) UpdateUserInfo(name string, id int) error{
	sqlStr := fmt.Sprintf("update t_user set name='%s' where id=%d;", name, id)
	return s.conn.Update(sqlStr)
}