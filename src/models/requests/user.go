package requests

type ChangeUserProfileRequest struct {
	Name      string `json:"name"`
	Email     string `json:"email"`
	StudentID string `json:"student-id"`
	Major     string `json:"major"`
	Admin     bool   `json:"admin"`
	Telephone string `json:"telephone"`
	QQ        string `json:"qq"`
	Level     int32  `json:"level"`
	Wanted    string `json:"wanted"`
	Intro     string `json:"intro"`
}

type ChangeUserPasswordRequest struct {
	Password    string `json:"password"`
	NewPassword string `json:"new-password"`
}

type CreateUserRequest struct {
	Name      string `json:"name"`
	Email     string `json:"email"`
	Password  string `json:"password"`
	Admin     bool   `json:"admin"`
	StudentID string `json:"student-id"`
	Major     string `json:"major"`
	Telephone string `json:"telephone"`
	QQ        string `json:"qq"`
	Level     int32  `json:"level"`
	Wanted    string `json:"wanted"`
	Intro     string `json:"intro"`
}
