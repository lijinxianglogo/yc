package dbs

import (
	"fmt"
	"teckin_ssmanager/internal/ssmanager/model"
	"time"
)

const (
	deviceTable = "t_device"
)

func (s *MysqlDB) InsertDeviceInfo(info model.DeviceInfo, uid int) error {
	sqlStr := fmt.Sprintf("insert into t_device (user_id, code, name, model, version, add_time) values ("+
		"%d, '%s', '%s', '%s', '%s', %d);", uid, info.Code, info.Name, info.Model, info.Version, time.Now().Unix())
	_, err := s.conn.Insert(sqlStr)
	return err
}

func (s *MysqlDB) UpdateDeviceName(devname, devcode string, uid int) error {
	sqlStr := fmt.Sprintf("update t_device set name='%s' where code='%s' and user_id=%d;", devname, devcode, uid)
	return s.conn.Update(sqlStr)
}

func (s *MysqlDB) GetUserAllRelateDeviceInfo(uid int) (devList model.DeviceList, err error) {
	sqlStr := fmt.Sprintf("select code, name, model, version from t_device where user_id = %d;", uid)
	mData, err := s.conn.Query(sqlStr)
	if err != nil {
		return model.DeviceList{}, err
	}
	for _, data := range mData {
		var devInfo model.DeviceInfo
		devInfo.Code = data["code"]
		devInfo.Name = data["name"]
		devInfo.Model = data["model"]
		devInfo.Version = data["version"]
		devList.Info = append(devList.Info, devInfo)
	}
	return
}

func (s *MysqlDB) GetUserRelateDeviceInfoByDeviceCode(uid int, devCode string) (devList model.DeviceList, err error) {
	sqlStr := fmt.Sprintf("select code, name, model, version from t_device where user_id = %d and code in (%s);", uid, devCode)
	mData, err := s.conn.Query(sqlStr)
	if err != nil {
		return model.DeviceList{}, err
	}
	for _, data := range mData {
		var devInfo model.DeviceInfo
		devInfo.Code = data["code"]
		devInfo.Name = data["name"]
		devInfo.Model = data["model"]
		devInfo.Version = data["version"]
		devList.Info = append(devList.Info, devInfo)
	}
	return
}

func (s *MysqlDB) BindDevice(uuid string, uid int) error {
	sqlStr := fmt.Sprintf("update t_device set user_id=%d where code='%s';", uid, uuid)
	return s.conn.Update(sqlStr)
}
