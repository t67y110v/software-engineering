package requests

type Registration struct {
	Email       string `json:"email"`
	Password    string `json:"password"`
	Name        string `json:"name"`
	SeccondName string `json:"seccond_name"`
}

type Login struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

type CheckToken struct {
	Cookie string `json:"token"`
}
