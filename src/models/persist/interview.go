package persist

import (
	"MSC2021/src/models"
	"container/list"
)

type Interviewee struct {
	Name   string `json:"name"`
	UserID uint   `json:"user-id"`
}

type GroupInterviewQueue struct {
	Name          string
	GroupID       uint
	CurrentUserID uint
	List          *list.List
}

type GroupInterviewQueueStatus struct {
	Name          string        `json:"name"`
	GroupID       uint          `json:"group-id"`
	CurrentUserID uint          `json:"current-user-id"`
	Interviewees  []Interviewee `json:"interviewees"`
}

func (queue *GroupInterviewQueue) Data() *GroupInterviewQueueStatus {
	if queue.List.Len() == 0 {
		return nil
	}
	var currentInterviewees []Interviewee
	for e := queue.List.Front(); e != nil; e = e.Next() {
		currentInterviewees = append(currentInterviewees, e.Value.(Interviewee))
	}
	return &GroupInterviewQueueStatus{
		Name:          queue.Name,
		GroupID:       queue.GroupID,
		CurrentUserID: queue.CurrentUserID,
		Interviewees:  currentInterviewees,
	}
}

func (queue *GroupInterviewQueue) DelayInterviewee() bool {
	afterUser := queue.List.Front()
	if afterUser == nil {
		return false
	}
	if afterUser = afterUser.Next(); afterUser == nil {
		return false
	} else if afterUser = afterUser.Next(); afterUser == nil {
		afterUser = queue.List.Front().Next()
	}
	queue.List.MoveAfter(queue.List.Front(), afterUser)
	return true
}

func (queue *GroupInterviewQueue) FirstInterviewee() *Interviewee {
	if queue.List.Front() == nil {
		return nil
	}
	return queue.List.Front().Value.(*Interviewee)
}

type InterviewManager struct {
	Started bool
	Groups  map[uint]*GroupInterviewQueue
}

func (groups *InterviewManager) AddGroup(id uint, name string) {
	groups.Groups[id] = &GroupInterviewQueue{GroupID: id, Name: name, List: list.New()}
}

func (groups *InterviewManager) AddInterviewee(user models.User) {
	if !groups.Started {
		return
	}
	for _, wantedGroup := range user.WantedGroups {
		if _, ok := groups.Groups[wantedGroup.ID]; ok {
			groups.Groups[wantedGroup.ID].List.PushBack(user)
		}
	}
}

func (groups *InterviewManager) CanInterview(userId uint) bool {
	if !groups.Started {
		return false
	}
	for _, queue := range groups.Groups {
		if queue.CurrentUserID == userId {
			return false
		}
	}
	return true
}

func (groups *InterviewManager) RequestInterview(groupId uint) *Interviewee {
	if !groups.Started {
		return nil
	}
	currentGroup := groups.Groups[groupId]
	userHead := currentGroup.FirstInterviewee()
	if userHead == nil {
		return nil
	}
	if groups.CanInterview(userHead.UserID) {
		currentGroup.CurrentUserID = userHead.UserID
		return userHead
	} else {
		currentGroup.DelayInterviewee()
	}
	for user := currentGroup.FirstInterviewee(); user.UserID != userHead.UserID; user = currentGroup.FirstInterviewee() {
		if groups.CanInterview(userHead.UserID) {
			currentGroup.CurrentUserID = userHead.UserID
			return userHead
		} else {
			currentGroup.DelayInterviewee()
		}
	}
	return nil
}

func (groups *InterviewManager) Start() {
	groups.Started = true
}

func (groups *InterviewManager) Stop() {
	groups.Started = false
	groups.Groups = nil
}

func (groups *InterviewManager) Status() map[string]interface{} {
	if !groups.Started {
		return nil
	}
	var status []*GroupInterviewQueueStatus
	for _, queue := range groups.Groups {
		status = append(status, queue.Data())
	}
	return map[string]interface{}{
		"status": status,
	}
}
