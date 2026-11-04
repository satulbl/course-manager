package v1

import "course-manager/internal/models"

type registerRequest struct {
	Email    string `json:"email"`
	Password string `json:"password"`
	Name     string `json:"name"`
}

func (req registerRequest) toInput() models.RegisterInput {
	return models.RegisterInput{
		Email:    req.Email,
		Password: req.Password,
		Name:     req.Name,
	}
}

type loginRequest struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

type tokenPairResponse struct {
	AccessToken  string `json:"accessToken"`
	RefreshToken string `json:"refreshToken"`
}

func newTokenPairResponse(pair *models.TokenPair) tokenPairResponse {
	return tokenPairResponse{
		AccessToken:  pair.AccessToken,
		RefreshToken: pair.RefreshToken,
	}
}
