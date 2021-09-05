package requests

type PublishQuestionRequest struct {
	Group   string `json:"group"`
	Title   string `json:"title"`
	Content string `json:"content"`
}

type UpdateQuestionRequest struct {
	Id      uint   `json:"id"`
	Title   string `json:"title"`
	Content string `json:"content"`
}

type AnswerQuestionRequest struct {
	Id     uint   `json:"id"` // question id
	Answer string `json:"answer"`
}
