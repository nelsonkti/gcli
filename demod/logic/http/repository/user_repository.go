/**
** @创建时间 : 2022/1/6 10:51
** @作者 : fzy
 */
package repository

import "demod/logic/http/model"

type UserRepository struct {
}

func (ur *UserRepository) Create(data model.UserModel) error {
	var userModel model.UserModel
	err := userModel.Model().Create(&data).Error
	if err != nil {
		return err
	}
	return nil
}