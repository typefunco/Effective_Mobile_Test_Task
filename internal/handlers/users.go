package handlers

import (
	"effectiveMobile/internal/entity"
	"effectiveMobile/internal/repo"
	"effectiveMobile/internal/utils"
	"fmt"
	"log/slog"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
)

// @Summary User sign up
// @Description Registers a new user
// @Tags users
// @Accept json
// @Produce json
// @Param user body entity.User true "User sign up information"
// @Success 200 {object} map[string]interface{}
// @Failure 400 {object} map[string]string
// @Failure 500 {object} map[string]string
// @Router /sign-up [post]
func signUp(ctx *gin.Context) {
	var user entity.User
	err := ctx.ShouldBindJSON(&user)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"Error": "Failed to create"})
		return
	}

	err = repo.SaveUser(user.Username, user.Password)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"Error": "Failed to save user"})
		return
	}

	jwt, err := utils.GenerateJWT(user.Username)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"Error": "Failed to create JWT"})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{"Object saved": user, "JWT": jwt})

}

// @Summary Become admin
// @Description Upgrades a user to admin status
// @Tags users
// @Produce json
// @Security Bearer
// @Success 200 {string} string "USER UPDATED TO ADMIN"
// @Failure 401 {object} map[string]string
// @Failure 500 {object} map[string]string
// @Router /update [patch]
func becomeAdmin(ctx *gin.Context) {
	authHeader := ctx.GetHeader("Authorization")
	if authHeader == "" {
		ctx.JSON(http.StatusUnauthorized, gin.H{"error": "Authorization header is missing"})
		return
	}

	splitToken := strings.Split(authHeader, " ")
	if len(splitToken) != 2 || splitToken[0] != "Bearer" {
		ctx.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid authorization format"})
		return
	}

	tokenString := splitToken[1]

	username, err := utils.GetUsernameFromJWT(tokenString)
	if err != nil {
		ctx.JSON(http.StatusUnauthorized, gin.H{"error": fmt.Sprintf("Invalid token: %v", err)})
		return
	}

	isAdmin, err := repo.CheckIsAdmin(username)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"Error": "Can't check is user admin"})
		return
	}

	fmt.Println(isAdmin)
	if isAdmin == true {
		ctx.JSON(200, "User already admin")
		return
	}

	err = repo.UpdateToAdmin(username)
	if err != nil {
		slog.Warn("Can't update user to ADMIN")
		ctx.JSON(500, "Can't update user to ADMIN")
		return
	}

	ctx.JSON(200, "USER UPDATED TO ADMIN")

}
