package services

import (
	"MSC2021/src/global"
	"MSC2021/src/models"
	"container/list"
	"errors"

	"gorm.io/gorm"
)

type Interviewee struct {
	UserID uint   `json:"user-id"`
	Name   string `json:"name"`
	Active bool   `json:"active"`
}

func (i *Interviewee) Activate() {
	i.Active = true
}

func (i *Interviewee) Deactivate() {
	i.Active = false
}

type InterviewQueue struct {
	GroupID           uint
	GroupName         string
	ActiveInterviewee *Interviewee
	QueuedInterviewee list.List
}

func (queue *InterviewQueue) Add(interviewee *Interviewee) {
	queue.QueuedInterviewee.PushBack(interviewee)
}

func (queue *InterviewQueue) Next() *Interviewee {
	queue.ActiveInterviewee.Deactivate()
	if queue.QueuedInterviewee.Len() == 0 {
		return nil
	}

	for e := queue.QueuedInterviewee.Front(); e != nil; e = e.Next() {
		if !e.Value.(*Interviewee).Active {
			queue.QueuedInterviewee.Remove(e)
			res := e.Value.(*Interviewee)
			res.Activate()
			queue.ActiveInterviewee = res
			return res
		}
	}
	return nil
}

func (queue *InterviewQueue) MaybeNext() *Interviewee {
	if queue.QueuedInterviewee.Len() == 0 {
		return nil
	}

	for e := queue.QueuedInterviewee.Front(); e != nil; e = e.Next() {
		if !e.Value.(*Interviewee).Active {
			res := e.Value.(*Interviewee)
			return res
		}
	}
	return nil
}

func (queue *InterviewQueue) GetActiveInterviewee() *Interviewee {
	return queue.ActiveInterviewee
}

func (queue *InterviewQueue) GetInterviewees() []*Interviewee {
	res := make([]*Interviewee, 0)
	for e := queue.QueuedInterviewee.Front(); e != nil; e = e.Next() {
		res = append(res, e.Value.(*Interviewee))
	}
	return res
}

func (queue *InterviewQueue) Status() map[string]interface{} {
	return map[string]interface{}{
		"group-id":   queue.GroupID,
		"group-name": queue.GroupName,
		"active":     queue.ActiveInterviewee,
		"queue":      queue.GetInterviewees(),
	}
}

type InterviewService struct {
	Started bool
	Groups  map[uint]*InterviewQueue
}

var InterviewServiceInstance InterviewService

type ServiceNotStartedError struct {
}

func (e *ServiceNotStartedError) Error() string {
	return "Interview service not started"
}

type GroupAlreadyExistsError struct {
}

func (e *GroupAlreadyExistsError) Error() string {
	return "Group already exists"
}

type GroupNotExistsError struct {
}

func (e *GroupNotExistsError) Error() string {
	return "Group does not exist"
}

func (service *InterviewService) AddGroup(groupID uint, groupName string) error {
	if !service.Started {
		return &ServiceNotStartedError{}
	}
	if _, ok := service.Groups[groupID]; ok {
		return &GroupAlreadyExistsError{}
	}
	service.Groups[groupID] = &InterviewQueue{
		GroupID:           groupID,
		GroupName:         groupName,
		ActiveInterviewee: nil,
		QueuedInterviewee: list.List{},
	}
	return nil
}

func StartInterview() error {
	InterviewServiceInstance.Started = true
	var groups []models.Group
	err := global.DATABASE.Model(&models.Group{}).Find(&groups).Error
	if err != nil {
		return err
	}
	for _, group := range groups {
		InterviewServiceInstance.AddGroup(group.ID, group.Name)
	}
	return nil
}

func StopInterview() error {
	InterviewServiceInstance.Started = false
	InterviewServiceInstance.Groups = make(map[uint]*InterviewQueue)
	return nil
}

func GetInterviewStatus() (map[string]interface{}, error) {
	if !InterviewServiceInstance.Started {
		return nil, &ServiceNotStartedError{}
	}
	res := make(map[string]interface{})
	for _, group := range InterviewServiceInstance.Groups {
		res[group.GroupName] = group.Status()
	}
	return res, nil
}

func UserSignInInterview(id uint) error {
	if !InterviewServiceInstance.Started {
		return &ServiceNotStartedError{}
	}
	var user models.User
	err := global.DATABASE.Select([]string{"id", "name"}).Where("id = ?", id).First(&user).Error
	if err != nil {
		return err
	}
	var groups []models.Group
	err = global.DATABASE.Model(&models.User{}).Where("id = ?", id).Association("WantedGroups").Find(&groups)
	if err != nil {
		return err
	}
	for _, group := range groups {
		if _, ok := InterviewServiceInstance.Groups[group.ID]; ok {
			InterviewServiceInstance.Groups[group.ID].Add(&Interviewee{
				UserID: id,
				Name:   user.Name,
			})
		}
	}
	return nil
}

func AdminGetSuggestInterviewUser(groupId uint) (map[string]interface{}, error) {
	if !InterviewServiceInstance.Started {
		return nil, &ServiceNotStartedError{}
	}
	if _, ok := InterviewServiceInstance.Groups[groupId]; !ok {
		return nil, &GroupNotExistsError{}
	}
	res := InterviewServiceInstance.Groups[groupId].MaybeNext()
	if res == nil {
		return nil, nil
	}
	return map[string]interface{}{
		"user-id": res.UserID,
		"name":    res.Name,
	}, nil
}

func AdminGetNextInterviewUser(groupId uint) (map[string]interface{}, error) {
	if !InterviewServiceInstance.Started {
		return nil, &ServiceNotStartedError{}
	}
	if _, ok := InterviewServiceInstance.Groups[groupId]; !ok {
		return nil, &GroupNotExistsError{}
	}
	res := InterviewServiceInstance.Groups[groupId].Next()
	if res == nil {
		return nil, nil
	}
	return map[string]interface{}{
		"user-id": res.UserID,
		"name":    res.Name,
	}, nil
}

func AdminGetInterviewContent(groupId uint, userId uint) (models.Interview, error) {
	var interview models.Interview
	err := global.DATABASE.Where("group_id = ? AND user_id = ?", groupId, userId).First(&interview).Error
	if err != nil {
		return models.Interview{}, err
	}
	return interview, nil
}

func AdminUpdateInterviewContent(groupId uint, userId uint, content string) error {
	interview := models.Interview{
		GroupID: groupId,
		UserID:  userId,
	}
	var prevInterview models.Interview
	err := global.DATABASE.Where("group_id = ? AND user_id = ?").First(&prevInterview).Error
	if err != nil && errors.Is(err, gorm.ErrRecordNotFound) {
		interview.Record = content
		err = global.DATABASE.Create(&interview).Error
		if err != nil {
			return err
		}
	} else if err != nil {
		return err
	}
	prevInterview.Record = content
	err = global.DATABASE.Save(&prevInterview).Error
	if err != nil {
		return err
	}
	return nil
}

func AdminDeleteInterviewContent(groupId uint, userId uint) error {
	err := global.DATABASE.Where("group_id = ? AND user_id = ?").Delete(models.Interview{}).Error
	if err != nil {
		return err
	}
	return nil
}
