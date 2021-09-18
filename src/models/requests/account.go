package requests

type ChangePasswordRequest struct {
	Password    string `json:"password"`
	NewPassword string `json:"new-password"`
}

type ChangeProfileRequest struct {
	Name      string `json:"name"`
	Email     string `json:"email"`
	StudentID string `json:"student-id"`
	Major     string `json:"major"`
	Telephone string `json:"telephone"`
	QQ        string `json:"qq"`
	Level     string `json:"level"`
	Intro     string `json:"intro"`
}

type LoginRequest struct {
	Account  string `json:"account"`
	Password string `json:"password"`
	Token    string `json:"token"`
}

type RegisterRequest struct {
	Name      string `json:"name"`
	Email     string `json:"email"`
	Password  string `json:"password"`
	StudentID string `json:"student-id"`
	Major     string `json:"major"`
	Telephone string `json:"telephone"`
	QQ        string `json:"qq"`
	Level     string `json:"level"`
	Intro     string `json:"intro"`
	Token     string `json:"token"`
}
