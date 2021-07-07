package middleware

import (
	"github.com/car12o/audio-mastery/api/generated/models"
	"github.com/car12o/audio-mastery/api/generated/restapi/operations"
)

type AuthHandler func(string) (*models.Principal, error)

func setupAuths(api *operations.AudioMasteryAPI) {
	api.BearerAuth = bearerAuth(api)
}

func bearerAuth(api *operations.AudioMasteryAPI) AuthHandler {
	return func(bearerToken string) (*models.Principal, error) {
		// implement logic to authenticate the user
		return &models.Principal{
			UserID: "test-userID",
			Email:  "test-email@emai.com",
		}, nil
	}
}
