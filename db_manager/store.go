package db_manager

import "ec/model"

//user func

func (d *DbInstance) AddUser(user *model.User) error {
	if err := d.Db.Create(user).Error; err != nil {
		return err
	}
	return nil
}

// //url func
func (d *DbInstance) AddUrl(url *model.Url) error {
	return d.Db.Create(url).Error
}
