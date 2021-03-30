package transaction

import "time"

type CampaignTransactionFormatter struct {
	Id        int       `json:"id"`
	Name      string    `json:"name"`
	Amount    int       `json:"amount"`
	CreatedAt time.Time `json:"created_at"`
}

type UserTransactionFormatter struct {
	Id        int               `json:"id"`
	Status    string            `json:"status"`
	Amount    int               `json:"amount"`
	CreatedAt time.Time         `json:"created_at"`
	Campaign  CampaignFormatter `json:"campaign"`
}

type CampaignFormatter struct {
	Name     string `json:"name"`
	ImageUrl string `json:"image_url"`
}

type TransactionFormatter struct {
	Id         int    `json:"id"`
	CampaignId int    `json:"campaign_id"`
	UserId     int    `json:"user_id"`
	Amount     int    `json:"amount"`
	Status     string `json:"status"`
	Code       string `json:"code"`
}

func FormatCampaignTransaction(transaction Transaction) CampaignTransactionFormatter {

	formatter := CampaignTransactionFormatter{}
	formatter.Id = transaction.Id
	formatter.Name = transaction.User.Name
	formatter.Amount = transaction.Amount
	formatter.CreatedAt = transaction.CreatedAt

	return formatter
}

func FormatCampaignTransactions(transactions []Transaction) []CampaignTransactionFormatter {

	if len(transactions) == 0 {

		return []CampaignTransactionFormatter{}
	}

	var transactionFormatter []CampaignTransactionFormatter

	for _, transaction := range transactions {

		formatter := FormatCampaignTransaction(transaction)
		transactionFormatter = append(transactionFormatter, formatter)
	}

	return transactionFormatter
}

func FormatUserTransaction(transactions Transaction) UserTransactionFormatter {

	formatter := UserTransactionFormatter{}
	formatter.Id = transactions.Id
	formatter.Amount = transactions.Amount
	formatter.Status = transactions.Status
	formatter.CreatedAt = transactions.CreatedAt

	campaignFormatter := CampaignFormatter{}
	campaignFormatter.Name = transactions.Campaign.Name
	campaignFormatter.ImageUrl = ""
	if len(transactions.Campaign.CampaignImages) > 0 {
		campaignFormatter.ImageUrl = transactions.Campaign.CampaignImages[0].FileName
	}

	formatter.Campaign = campaignFormatter
	return formatter
}

func FormatUserTransactions(transactions []Transaction) []UserTransactionFormatter {

	if len(transactions) == 0 {

		return []UserTransactionFormatter{}
	}

	var transactionFormatter []UserTransactionFormatter

	for _, transaction := range transactions {

		formatter := FormatUserTransaction(transaction)
		transactionFormatter = append(transactionFormatter, formatter)
	}

	return transactionFormatter
}

func FormatTransaction(transactions Transaction) TransactionFormatter {

	formatter := TransactionFormatter{}
	formatter.Id = transactions.Id
	formatter.CampaignId = transactions.CampaignId
	formatter.UserId = transactions.UserId
	formatter.Amount = transactions.Amount
	formatter.Status = transactions.Status
	formatter.Code = transactions.Code

	return formatter
}
