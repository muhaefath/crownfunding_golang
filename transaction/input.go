package transaction

import "golang_project/user"

type GetTransactionDetailInput struct {
	Id   int `uri:"id" binding:"required"`
	User user.User
}
