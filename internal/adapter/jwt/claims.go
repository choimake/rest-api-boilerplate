package jwt

type Claims struct {
	Subject   string `json:"sub,omitempty"`
	FirstName string `json:"first_name,omitempty"`
	LastName  string `json:"last_name,omitempty"`
	Email     string `json:"email,omitempty"`
}

func (c Claims) Valid() error {
	return nil
}
