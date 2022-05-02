package token

import (
	"invite-token/src/daos"
	"invite-token/utils/drivers/db"
	"time"

	"invite-token/config"
	"invite-token/src/models"

	"invite-token/config/constants"

	"gorm.io/gorm"
)

type InviteToken struct {
	db          *gorm.DB
	inviteToken *daos.InviteToken
}

func New() *InviteToken {
	return &InviteToken{
		db:          db.DB,
		inviteToken: daos.NewInviteToken(),
	}
}

func (i *InviteToken) Generate() (string, error) {
	token := &models.InviteToken{
		ID:        randStringBytes(config.TokenLen),
		Active:    true,
		ExpiresAt: time.Now().Add(24 * time.Duration(config.TokenExpiry) * time.Hour),
	}

	err := i.inviteToken.Upsert(token)
	if err != nil {
		return "", err
	}

	return token.ID, nil
}

func (i *InviteToken) GetAll() ([]*models.InviteToken, error) {
	return i.inviteToken.GetAll()
}

func (i *InviteToken) Validate(token string) (bool, error) {
	inviteToken, err := i.inviteToken.Get(token)
	if err != nil {
		return false, err
	}

	if !inviteToken.Active || (inviteToken.ExpiresAt.Before(time.Now())) {
		return false, constants.ErrInvalidInviteToken
	}

	return true, nil
}

func (i *InviteToken) InValidate(token string) error {
	inviteToken, err := i.inviteToken.Get(token)
	if err != nil {
		return err
	}

	inviteToken.Active = false
	err = i.inviteToken.Upsert(inviteToken)
	if err != nil {
		return err
	}

	return nil
}
