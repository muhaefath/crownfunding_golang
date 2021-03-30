package transaction

import "golang_project/user"

type GetTransactionDetailInput struct {
	Id   int `uri:"id" binding:"required"`
	User user.User
}

type CreateTransaction struct {
	Amount     int `json:"amount" binding:"required"`
	CampaignId int `json:"campaign_id" binding:"required"`
	User       user.User
}
