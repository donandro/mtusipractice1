package models

type SignUpRequest struct {
	Name     *string `json:"name"`
	Login    *string `json:"login"`
	Password *string `json:"password"`
}

type LoginRequest struct {
	Login    *string `json:"login"`
	Password *string `json:"password"`
}

type AuthResponse struct {
	AccessToken string `json:"accessToken"`
	Expiry      int    `json:"expiry"`
}
