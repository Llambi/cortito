package domain

import "time"

type CortitoUser struct {
	Ip        string    `json:"ip" binding:"required"`
	Quota     uint64    `json:"quota" binding:"required"`
	RefreshAt time.Time `json:"RefreshAt" binding:"required"`
}

type FindCortitoUserRequest struct {
	Ip string `json:"ip" binding:"required"`
}

type CreateCortitoUserRequest struct {
	Ip string `json:"ip" binding:"required"`
}

type CortitoUserResponse struct {
	Quota     uint64    `json:"quota" binding:"required"`
	RefreshAt time.Time `json:"RefreshAt" binding:"required"`
}

func (createRequest *CreateCortitoUserRequest) Into() CortitoUser {
	cortitoUser := CortitoUser{
		Ip: createRequest.Ip,
	}
	return cortitoUser
}

func (cortitoUser *CortitoUser) Into() CortitoUserResponse {
	return CortitoUserResponse{
		Quota:     cortitoUser.Quota,
		RefreshAt: cortitoUser.RefreshAt,
	}
}
