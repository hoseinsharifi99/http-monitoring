package db_manager

import "ec/model"

func (dm *DbInstance) GetUserByUserName(username string) (*model.User, error) {
	user := new(model.User)
	if err := dm.Db.First(&user, model.User{UserName: username}).Error; err != nil {
		return nil, err
	}

	return user, nil
}

func (d *DbInstance) GetUrlsByUserId(id uint) ([]model.Url, error) {
	var urls []model.Url
	if err := d.Db.Model(&model.Url{}).Where("user_id = ?", id).Find(&urls).Error; err != nil {
		return nil, err
	}
	return urls, nil
}

func (d *DbInstance) GetAllUrls() ([]*model.Url, error) {
	var urls []*model.Url
	if err := d.Db.Model(&model.Url{}).Find(&urls).Error; err != nil {
		return nil, err
	}
	return urls, nil
}

func (d *DbInstance) GetUrlById(id uint) (*model.Url, error) {
	url := new(model.Url)
	if err := d.Db.First(url, id).Error; err != nil {
		return nil, err
	}
	requests := make([]model.Request, 0)
	d.Db.Model(url).Association("Requests").Find(&requests)
	url.Requests = requests
	return url, nil
}

func (s *DbInstance) AddRequest(req *model.Request) error {
	return s.Db.Create(req).Error
}

func (s *DbInstance) IncrementFailed(url *model.Url) error {
	url.FailedTimes += 1
	return s.UpdateUrl(url)
}

func (s *DbInstance) IncrementSuccess(url *model.Url) error {
	url.SuccessTimes += 1
	return s.UpdateUrl(url)
}

func (s *DbInstance) UpdateUrl(url *model.Url) error {
	return s.Db.Model(url).Update(url).Error
}
