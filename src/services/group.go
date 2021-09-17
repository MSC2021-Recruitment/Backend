package services

import (
	"MSC2021/src/global"
	"MSC2021/src/models"
	"gorm.io/gorm"
)

type GroupRes struct {
	ID          uint   `json:"id"`
	Name        string `json:"name"`
	Description string `json:"description"`
	Applied     bool   `json:"applied"`
}

func GetGroupList(uid int) ([]GroupRes, error) {
	var groups []models.Group
	ans := make([]GroupRes, 0)
	err := global.DATABASE.Select([]string{"id", "name", "description"}).Find(&groups).Error
	if err != nil {
		return nil, err
	}
	for _, group := range groups {
		ans = append(ans, GroupRes{
			ID:          group.ID,
			Name:        group.Name,
			Description: group.Description,
			Applied:     false,
		})
	}
	if uid != -1 {
		var wantedGroups []models.Group
		err := global.DATABASE.Model(&models.User{ID: uint(uid)}).Association("WantedGroups").Find(&wantedGroups)
		global.LOGGER.Sugar().Info("wanted groups: ", wantedGroups)
		if err != nil && err != gorm.ErrRecordNotFound {
			return nil, err
		}
		for i := 0; i < len(ans); i++ {
			id := ans[i].ID
			for _, g := range wantedGroups {
				if g.ID == id {
					ans[i].Applied = true
				}
			}
		}
	}

	return ans, err
}

type AdminGroupRes struct {
	ID          uint   `json:"id"`
	Name        string `json:"name"`
	Description string `json:"description"`
	Count       int    `json:"count"`
}

func AdminGetGroupList() ([]AdminGroupRes, error) {
	var groups []models.Group
	err := global.DATABASE.Select([]string{"id", "name", "description"}).Find(&groups).Error
	if err != nil {
		return nil, err
	}
	var res []AdminGroupRes
	for _, group := range groups {
		var count = global.DATABASE.Model(&group).Association("Interviewees").Count()
		res = append(res, AdminGroupRes{ID: group.ID, Name: group.Name, Description: group.Description, Count: int(count)})
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
	err = global.DATABASE.Model(&group).Association("Interviewees").Append(&user)
	if err != nil {
		return err
	}
	err = global.DATABASE.Model(&user).Association("WantedGroups").Append(&group)
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

func ChangeGroupDetail(id uint, name string, description string) error {
	group := models.Group{ID: id}
	err := global.DATABASE.First(&group).Error
	if err != nil {
		return err
	}
	group.Name = name
	group.Description = description
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

func CreateGroup(name string, description string) error {
	group := models.Group{Name: name, Description: description}
	err := global.DATABASE.Create(&group).Error
	return err
}
