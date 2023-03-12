package responses

type Registration struct {
	ID                int    `json:"ID"`
	Email             string `json:"Email"`
	Name              string `json:"Name"`
	SeccondName       string `json:"SeccondName"`
	Password          string `json:"Password,omitempty"`
	Isadmin           bool   `json:"Isadmin"`
	EncryptedPassword string `json:"-"`
}

type Login struct {
	Token string `json:"token"`
	Email string `json:"Email"`
	Name  string `json:"Name"`
}
