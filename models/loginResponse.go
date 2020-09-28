package models

/*
	ResponseLogin contains the login token
 */
type LoginResponse struct{
	Token string `json:"token,omitempty"`
}

