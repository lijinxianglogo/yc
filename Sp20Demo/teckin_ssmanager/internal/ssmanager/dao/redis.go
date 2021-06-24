package dao

import (
	"encoding/json"
	"fmt"
	"teckin_ssmanager/internal/ssmanager/model"
	"time"
)

const (
	teckin_email_token = "teckin_tokenmanage_%s" //format: email
	teckin_userinfo    = "teckin_amin_%d"        //format uid
)

func getUserInfo(uid int64) string {
	return fmt.Sprintf(teckin_userinfo, uid)
}

func getEmailTokenMap(email string) string {
	return fmt.Sprintf(teckin_email_token, email)
}

func (d *Dao) SetUserToken(email, token string) error {
	return d.Redis.Set(getEmailTokenMap(email), token, time.Hour*24)
}

func (d *Dao) GetUserToken(email string) (string, error) {
	return d.Redis.Get(getEmailTokenMap(email))
}

func (d *Dao) DelUserToken(email string) error {
	return d.Redis.Del(getEmailTokenMap(email))
}

func (d *Dao) InsertUserInfo(uid int64, user model.User) error {
	b, _ := json.Marshal(user)
	return d.Redis.Set(getUserInfo(uid), string(b), 0)
}

func (d *Dao) GetUserInfoByUid(uid int64) (model.User, error){
	data, err := d.Redis.Get(getUserInfo(uid))
	if err != nil{
		return model.User{}, err
	}
	var user model.User
	err = json.Unmarshal([]byte(data), &user)
	return user, err
}
