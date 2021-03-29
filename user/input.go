package user

type RegisterUserInput struct {
	Name     string `json:"name" binding:"required"`
	Email    string `json:"name" binding:"required,email"`
	Password string `json:"name" binding:"required"`
	// Token    string
}
