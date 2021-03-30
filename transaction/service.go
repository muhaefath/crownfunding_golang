package transaction

import (
	"errors"
	"golang_project/campaign"
)

type service struct {
	repositoy          Repository
	CampaignRepository campaign.Repository
}

type Service interface {
	GetTransactionByCampaignId(input GetTransactionDetailInput) ([]Transaction, error)
}

func NewService(repository Repository, campaignRepository campaign.Repository) *service {

	return &service{repository, campaignRepository}
}

func (s *service) GetTransactionByCampaignId(input GetTransactionDetailInput) ([]Transaction, error) {

	campaign, err := s.CampaignRepository.FindById(input.Id)
	if err != nil {
		return []Transaction{}, err
	}

	if campaign.UserId != input.User.Id {
		return []Transaction{}, errors.New("not owner of this campaign")
	}

	transactions, err := s.repositoy.GetCampaignById(input.Id)
	if err != nil {

		return transactions, err
	}

	return transactions, nil
}
