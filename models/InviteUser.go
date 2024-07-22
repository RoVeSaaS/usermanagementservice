package models

type UserInvite struct {
	EmailID string `json:"email_id" binding:"required"`
	Role    string `json:"role" binding:"required"`
}
