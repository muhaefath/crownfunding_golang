package campaign

import (
	"errors"
	"fmt"

	"github.com/gosimple/slug"
)

type Service interface {
	FindCampaigns(user_id int) ([]Campaign, error)
	FindCampaign(input GetCampaignInput) (Campaign, error)
	CreateCampaign(input CreateCampaignInput) (Campaign, error)
	UpdateCampaign(inputId GetCampaignInput, inputData CreateCampaignInput) (Campaign, error)
}

type service struct {
	repository Repository
}

func NewService(repository Repository) *service {

	return &service{repository}
}

func (s *service) FindCampaigns(user_id int) ([]Campaign, error) {

	if user_id != 0 {
		campaigns, err := s.repository.FindByUserId(user_id)
		if err != nil {
			return campaigns, err
		}

		return campaigns, nil
	}

	campaigns, err := s.repository.FindAll()
	if err != nil {
		return campaigns, err
	}

	return campaigns, nil
}

func (s *service) FindCampaign(input GetCampaignInput) (Campaign, error) {

	campaigns, err := s.repository.FindById(input.Id)
	if err != nil {
		return campaigns, err
	}

	return campaigns, nil
}

func (s *service) CreateCampaign(input CreateCampaignInput) (Campaign, error) {

	campaign := Campaign{}
	campaign.Name = input.Name
	campaign.ShortDescription = input.ShortDescription
	campaign.Description = input.Description
	campaign.Perks = input.Perks
	campaign.GoalAmmount = input.GoalAmmount

	campaign.UserId = input.User.Id

	slugCandidate := fmt.Sprintf("%s %s", input.Name, input.User.Id)
	campaign.Slug = slug.Make(slugCandidate)

	newCampaign, err := s.repository.Save(campaign)
	if err != nil {
		return newCampaign, err
	}

	return newCampaign, nil

}

func (s *service) UpdateCampaign(inputId GetCampaignInput, inputData CreateCampaignInput) (Campaign, error) {

	campaign, err := s.repository.FindById(inputId.Id)
	if err != nil {
		return campaign, err
	}

	if inputData.User.Id != campaign.UserId {
		return campaign, errors.New("not owner of this campaign")
	}

	campaign.Name = inputData.Name
	campaign.ShortDescription = inputData.ShortDescription
	campaign.Description = inputData.Description
	campaign.Perks = inputData.Perks
	campaign.GoalAmmount = inputData.GoalAmmount

	updatedCampaign, err := s.repository.Update(campaign)
	if err != nil {
		return updatedCampaign, err
	}

	return updatedCampaign, nil

}
