package dto

type CreateUserRequest struct {
	ID        string  `json:"id"`
	Email     string  `json:"email"`
	Password  string  `json:"password"`
	FirstName *string `json:"first_name"`
	LastName  *string `json:"last_name"`
	Phone     *string `json:"phone"`
}
