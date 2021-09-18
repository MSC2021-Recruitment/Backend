package requests

type ChangeUserProfileRequest struct {
	Name      string `json:"name"`
	Email     string `json:"email"`
	StudentID string `json:"student-id"`
	Major     string `json:"major"`
	Admin     bool   `json:"admin"`
	Telephone string `json:"telephone"`
	QQ        string `json:"qq"`
	Level     string `json:"level"`
	Wanted    string `json:"wanted"`
	Intro     string `json:"intro"`
}

type ChangeUserPasswordRequest struct {
	Password    string `json:"password"`
	NewPassword string `json:"new-password"`
}
