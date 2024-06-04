package domain

const (
	ErrorMsg01 = "username must be fifteen characters or fewer."
	ErrorMsg02 = "username must be at least five characters long."
	ErrorMsg03 = "password must be fifteen characters or fewer."
	ErrorMsg04 = "password must be at least five characters long."
	ErrorMsg05 = "mailadress must be thirty characters or fewer."
	ErrorMsg06 = "mailadress must be at least five characters long."
	ErrorMsg07 = "mailadress is invalid."
	ErrorMsg08 = "signup failed."
	ErrorMsg09 = "user does not exist."
)

type Message struct {
	Message string `json:"message"`
}
