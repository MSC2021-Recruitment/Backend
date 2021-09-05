package requests

type CreateQuestionRequest struct {
	Group   string `json:"group"`
	Title   string `json:"title"`
	Content string `json:"content"`
}

type UpdateQuestionRequest struct {
	Title   string `json:"title"`
	Content string `json:"content"`
}

type AnswerQuestionRequest struct {
	Answer string `json:"answer"`
}
