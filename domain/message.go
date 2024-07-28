package domain

const (
	ErrorMsg01 = "user name must be fifteen characters or fewer."
	ErrorMsg02 = "user name must be at least five characters long."
	ErrorMsg03 = "password must be fifteen characters or fewer."
	ErrorMsg04 = "password must be at least five characters long."
	ErrorMsg05 = "mailadress must be thirty characters or fewer."
	ErrorMsg06 = "mailadress must be at least five characters long."
	ErrorMsg07 = "mailadress is invalid."
	ErrorMsg08 = "signup failed."
	ErrorMsg09 = "user does not exist."
	ErrorMsg10 = "no permission."
	ErrorMsg11 = "project name must be at least thirty characters long."
	ErrorMsg12 = "please enter a numeric user id."
)

type Message struct {
	Message string `json:"message"`
}
