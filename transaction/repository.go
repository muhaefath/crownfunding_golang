package transaction

import "gorm.io/gorm"

type repositoy struct {
	db *gorm.DB
}

type Repository interface {
	GetCampaignById(campaignId int) ([]Transaction, error)
}

func NewRepository(db *gorm.DB) *repositoy {

	return &repositoy{db}
}

func (r *repositoy) GetCampaignById(campaignId int) ([]Transaction, error) {

	var transactions []Transaction

	err := r.db.Preload("User").Where("campaign_id = ?", campaignId).Order("id desc").Find(&transactions).Error
	if err != nil {
		return transactions, err
	}

	return transactions, nil
}
