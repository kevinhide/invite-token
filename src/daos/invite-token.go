package daos

import (
	"invite-token/src/models"
	"invite-token/utils/drivers/db"

	"gorm.io/gorm"
)

type InviteToken struct {
	db *gorm.DB
}

func NewInviteToken() *InviteToken {
	return &InviteToken{
		db: db.DB,
	}
}

func (i *InviteToken) Get(id string) (*models.InviteToken, error) {
	inviteToken := models.InviteToken{}
	result := i.db.First(&inviteToken, "id = ?", id)
	if result.Error != nil {
		return &inviteToken, result.Error
	}

	return &inviteToken, nil
}

func (i *InviteToken) Upsert(inviteToken *models.InviteToken) error {
	result := i.db.Save(inviteToken)
	if result.Error != nil {
		return result.Error
	}

	return nil
}

func (i *InviteToken) GetAll() ([]*models.InviteToken, error) {
	var inviteTokens []*models.InviteToken

	result := i.db.Find(&inviteTokens)
	if result.Error != nil {
		return nil, result.Error
	}

	return inviteTokens, nil
}
