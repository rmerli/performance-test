package inputs

type startAuthenticationInput struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

var StartAuthenticationInputValues = startAuthenticationInput{
	Email:    "bsupp@example.com",
	Password: "Madonna!123",
}
