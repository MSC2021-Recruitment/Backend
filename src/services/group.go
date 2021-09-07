package services

import (
	"MSC2021/src/global"
	"MSC2021/src/models"
)

func GetGroupList() ([]models.Group, error) {
	var groups []models.Group
	err := global.DATABASE.Select([]string{"id", "name"}).Find(&groups).Error
	return groups, err
}

type AdminGroupRes struct {
	ID    uint   `json:"id"`
	Name  string `json:"name"`
	Count int    `json:"count"`
}

func AdminGetGroupList() ([]AdminGroupRes, error) {
	var groups []models.Group
	err := global.DATABASE.Select([]string{"id", "name"}).Find(&groups).Error
	if err != nil {
		return nil, err
	}
	var res []AdminGroupRes
	for _, group := range groups {
		var count = global.DATABASE.Model(&group).Association("Interviewees").Count()
		res = append(res, AdminGroupRes{ID: group.ID, Name: group.Name, Count: int(count)})
	}
	return res, nil
}

func AdminGetGroupDetail(groupId uint) (models.Group, error) {
	group := models.Group{ID: groupId}
	err := global.DATABASE.First(&group).Error
	return group, err
}

func JoinGroup(groupId uint, userId uint) error {
	group := models.Group{ID: groupId}
	err := global.DATABASE.First(&group).Error
	if err != nil {
		return err
	}
	user := models.User{ID: userId}
	err = global.DATABASE.First(&user).Error
	if err != nil {
		return err
	}
	err = global.DATABASE.Model(&group).Association("Interviewees").Append(user)
	if err != nil {
		return err
	}
	err = global.DATABASE.Model(&user).Association("WantedGroups").Append(group)
	return err
}

func LeaveGroup(groupId uint, userId uint) error {
	group := models.Group{ID: groupId}
	err := global.DATABASE.First(&group).Error
	if err != nil {
		return err
	}
	user := models.User{ID: userId}
	err = global.DATABASE.First(&user).Error
	if err != nil {
		return err
	}
	err = global.DATABASE.Model(&group).Association("Interviewees").Delete(user)
	if err != nil {
		return err
	}
	err = global.DATABASE.Model(&user).Association("WantedGroups").Delete(group)
	return err
}

func ChangeGroupDetail(id uint, name string) error {
	group := models.Group{ID: id}
	err := global.DATABASE.First(&group).Error
	if err != nil {
		return err
	}
	group.Name = name
	err = global.DATABASE.Save(&group).Error
	return err
}

func DeleteGroup(id uint) error {
	group := models.Group{ID: id}
	err := global.DATABASE.First(&group).Error
	if err != nil {
		return err
	}
	err = global.DATABASE.Delete(&group).Error
	return err
}

func CreateGroup(name string) error {
	group := models.Group{Name: name}
	err := global.DATABASE.Create(&group).Error
	return err
}
