package resources

type User struct {
	Id int
	Email string 
	PasswordHash string 
	FirstName string 
	LastName string 
}

type SignupRequestBody struct {
	Email string `json:"email"`
	Password string `json:"password"`
	FirstName string `json:"firstName"`
	LastName string `json:"lastName"`
}


type SigninRequestBody struct {
	Email string `json:"email"`
	Password string `json:"password"`
}

