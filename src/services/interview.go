package services

import (
	"MSC2021/src/global"
	"MSC2021/src/models"
)

func StartInterview() error {
	var groups []models.Group
	err := global.DATABASE.Model(&models.Group{}).Find(&groups).Error
	if err != nil {
		return err
	}
	for _, group := range groups {
		global.INTERVIEW_QUEUE.AddGroup(group.ID, group.Name)
	}
	global.INTERVIEW_QUEUE.Start()
	return nil
}

func StopInterview() error {
	global.INTERVIEW_QUEUE.Stop()
	return nil
}

func GetInterviewStatus() (map[string]interface{}, error) {
	res := global.INTERVIEW_QUEUE.Status()
	return res, nil
}
