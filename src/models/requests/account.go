package requests

type ChangePasswordRequest struct {
	Password    string `json:"password"`
	NewPassword string `json:"new-password"`
}

type ChangeProfileRequest struct {
	Name      string `json:"name"`
	Email     string `json:"email"`
	Password  string `json:"password"`
	Major     string `json:"major"`
	Telephone string `json:"telephone"`
	QQ        string `json:"qq"`
	Level     int32  `json:"level"`
	Wanted    string `json:"wanted"`
	Intro     string `json:"intro"`
}

type LoginRequest struct {
	Account  string `json:"account"`
	Password string `json:"password"`
}

type RegisterRequest struct {
	Name      string `json:"name"`
	Email     string `json:"email"`
	Password  string `json:"password"`
	Major     string `json:"major"`
	Telephone string `json:"telephone"`
	QQ        string `json:"qq"`
	Level     int32  `json:"level"`
	Wanted    string `json:"wanted"`
	Intro     string `json:"intro"`
}
