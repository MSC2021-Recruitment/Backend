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

type IntervieweeQueueElement struct {
	status int
	/*
		0 - waiting
		1 - interviewed
		2 - interviewing in other group
		3 - interviewing
	*/
	Interviewees Interviewee
}

type GroupInterviewQueueStatus struct {
	Name          string                    `json:"name"`
	GroupID       uint                      `json:"group-id"`
	CurrentUserID uint                      `json:"current-user-id"`
	Interviewees  []IntervieweeQueueElement `json:"interviewees"`
}

func (queue *GroupInterviewQueue) Data() *GroupInterviewQueueStatus {
	if queue.List.Len() == 0 {
		return nil
	}
	var currentInterviewees []IntervieweeQueueElement
	for e := queue.List.Front(); e != nil; e = e.Next() {
		currentInterviewees = append(currentInterviewees, e.Value.(IntervieweeQueueElement))
	}
	return &GroupInterviewQueueStatus{
		Name:          queue.Name,
		GroupID:       queue.GroupID,
		CurrentUserID: queue.CurrentUserID,
		Interviewees:  currentInterviewees,
	}
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
			groups.Groups[wantedGroup.ID].List.PushBack(IntervieweeQueueElement{
				status: 0,
				Interviewees: Interviewee{
					Name:   user.Name,
					UserID: user.ID,
				}})
		}
	}
	return
}

func (groups *InterviewManager) CanInterview(userId uint) int {
	if !groups.Started {
		return -2
	}
	for i, queue := range groups.Groups {
		if queue.CurrentUserID == userId {
			return int(i)
		}
	}
	return -1
}

/*
更新队列元素状态, 并且把现在正在面试的人的状态标记为已面试
*/
func (groups *InterviewManager) UpdateIntervieweeStatus(ee *IntervieweeQueueElement) {
	if ee == nil {
		return
	}
	if ee.status == 3 {
		ee.status = 1
	}
	if ee.status != 1 {
		if groups.CanInterview(ee.Interviewees.UserID) != -1 { //被人占用
			ee.status = 2
		} else {
			ee.status = 0
		}
	}
	return
}

func (groups *InterviewManager) RequestInterview(groupId uint) *Interviewee {
	if !groups.Started {
		return nil
	}
	currentGroup := groups.Groups[groupId]
	for true {
		var ansInterviewer *IntervieweeQueueElement
		currentGroup.CurrentUserID = 0
		for e := currentGroup.List.Front(); e != nil; e = e.Next() {
			currentElement := e.Value.(*IntervieweeQueueElement)
			groups.UpdateIntervieweeStatus(currentElement)
			if currentElement.status == 0 && ansInterviewer == nil { //第一个为0的人称为下一个被面试的 不要break是因为要顺便更新下面的
				ansInterviewer = currentElement
			}
		}
		if ansInterviewer == nil {
			return nil
		}
		currentGroup.CurrentUserID = ansInterviewer.Interviewees.UserID
		ansInterviewer.status = 3

		//如果两个线协程同时找到了这个人, 编号较小的主动放弃.
		if groups.CanInterview(currentGroup.CurrentUserID) > int(groupId) {
			currentGroup.CurrentUserID = 0
			ansInterviewer.status = 2
		} else {
			return &ansInterviewer.Interviewees
		}
	}

	return nil
}

//一个组面试一个人结束了
func (groups *InterviewManager) CloseGroupInterview(groupId uint) {
	if !groups.Started {
		return
	}
	currentGroup := groups.Groups[groupId]
	for e := currentGroup.List.Front(); e != nil; e.Next() {
		if e.Value.(*IntervieweeQueueElement).status == 3 {
			e.Value.(*IntervieweeQueueElement).status = 1
		}
	}
	groups.Groups[groupId].CurrentUserID = 0

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
