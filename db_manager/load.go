package db_manager

import "ec/model"

func (dm *DbInstance) GetUserByUserName(username string) (*model.User, error) {
	user := new(model.User)
	if err := dm.Db.First(&user, model.User{UserName: username}).Error; err != nil {
		return nil, err
	}

	return user, nil
}
