package model

type User struct {
	ID       int
	Key     string
	Name     string
	Mobile   int64
	Email    string
	Idcard   string
	Sex      int32
	Password string
	AddTime  int64
	LastTime int64
	IsDel    int32
}

type UserLoginResponse struct {
	ID       int    `json:"id"`
	Key     string  `json:"key"`
	Name     string `json:"name"`
	Email    string `json:"email"`
	Token    string `json:"token"`
}