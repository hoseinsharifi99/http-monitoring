package db_manager

import "ec/model"

func (d *DbInstance) AddUser(user *model.User) error {
	if err := d.Db.Create(user).Error; err != nil {
		return err
	}
	return nil
}
