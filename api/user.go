package api

import (
	"time"

	"github.com/gin-gonic/gin"
)

type createUserRequest struct {
	Username string `json:"username" binding:"required,alphanum"`
	Password string `json:"password" binding:"required,min=6"`
	FullName string `json:"full_name" binding:"required,alphanum"`
	Email    string `json:"email" binding:"required,email"`
}

type userResponse struct {
	Username          string    `json:"username"`
	FullName          string    `json:"full_name"`
	Email             string    `json:"email" binding:"required,email"`
	PasswordChangedAt time.Time `json:"passed_changed_at"`
	CreatedAt         time.Time `json:"created_at"`
}

func newUserResponse(user interface{}) userResponse { return userResponse{} }

func (server *Server) createUser(ctx *gin.Context) {

}

type loginUserRequest struct {
}
