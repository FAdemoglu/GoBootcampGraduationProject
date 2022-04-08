package auth

import (
	"github.com/FAdemoglu/graduationproject/internal/config"
	"github.com/FAdemoglu/graduationproject/internal/domain/users"
	jwtHelper "github.com/FAdemoglu/graduationproject/pkg/jwt"
	"github.com/FAdemoglu/graduationproject/shared"
	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	"net/http"
	"os"
	"time"
)

type AuthController struct {
	appConfig *config.Configuration
	r         users.UserService
}

func NewAuthController(appConfig *config.Configuration, r users.UserService) *AuthController {
	return &AuthController{
		appConfig: appConfig,
		r:         r,
	}
}

// Login godoc
// @Summary Login into system with username and password
// @Tags Auth
// @Accept  json
// @Produce  json
// @Param loginRequest body LoginRequest true "login informations"
// @Success 200
// @Failure 400 {object} map[string]string
// @Failure 401 {object} map[string]string
// @Failure 404 {object} map[string]string
// @Failure 500 {object} map[string]string
// @Router /user/login [post]
func (c *AuthController) Login(g *gin.Context) {
	var req LoginRequest

	//If request doesnt match with login request struct send error
	if err := g.ShouldBind(&req); err != nil {
		g.JSON(http.StatusBadRequest, gin.H{
			"error_message": "Check your request body.",
		})
		return
	}
	//We are comparing request in database user table
	user := c.r.GetByUserNameAndPassword(req.Username, req.Password)

	//If user is can not find by system send the client error
	if user.Username == "" {
		g.JSON(http.StatusNotFound, gin.H{
			"error_message": "Check your username and password",
		})
		return
	}
	jwtClaims := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"userId":   user.Id,
		"username": user.Username,
		"iat":      time.Now().Unix(),
		"iss":      os.Getenv("ENV"),
		"exp":      time.Now().Add(24 * time.Hour).Unix(),
		"roles":    user.Roles,
	})
	//Token generate with jwt helper
	token := jwtHelper.GenerateToken(jwtClaims, c.appConfig.JwtSettings.SecretKey)
	//If all success send the client jwt token
	data := shared.ApiResponse{IsSuccess: true, Message: "Başarılı", Data: token}
	g.JSON(http.StatusOK, data)
}

// Register godoc
// @Summary Register into system with username and password
// @Tags Auth
// @Accept  json
// @Produce  json
// @Param registerRequest body RegisterRequest true "register informations"
// @Success 200
// @Failure 400 {object} map[string]string
// @Failure 401 {object} map[string]string
// @Failure 404 {object} map[string]string
// @Failure 500 {object} map[string]string
// @Router /user/register [post]
func (c *AuthController) Register(g *gin.Context) {
	var req RegisterRequest

	//Request should match with register request
	if err := g.ShouldBind(&req); err != nil {
		g.JSON(http.StatusBadRequest, gin.H{
			"error_message": "Check your request body.",
		})
		return
	}

	//We are search the username in database if exist send the client error with message
	user := c.r.GetByUsername(req.Username)
	if user.Username != "" {
		g.JSON(http.StatusBadRequest, gin.H{
			"error_message": "Username exist in database",
		})
		return
	}

	//Compare passwords
	if req.Password != req.ConfirmPassword {
		g.JSON(http.StatusForbidden, gin.H{
			"error_message": "Passwords do not match",
		})
		return
	}

	//If all errors are nil system create user in database and return jwt token to user
	c.r.CreateUser(req.Username, req.Password)
	jwtClaims := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"userId":   user.Id,
		"username": user.Username,
		"iat":      time.Now().Unix(),
		"iss":      os.Getenv("ENV"),
		"exp":      time.Now().Add(24 * time.Hour).Unix(),
		"roles":    user.Roles,
	})
	token := jwtHelper.GenerateToken(jwtClaims, c.appConfig.JwtSettings.SecretKey)
	data := shared.ApiResponse{IsSuccess: true, Message: "Başarılı", Data: token}
	g.JSON(http.StatusOK, data)

}

func (c *AuthController) VerifyToken(g *gin.Context) {
	token := g.GetHeader("Authorization")
	decodedClaims := jwtHelper.VerifyToken(token, c.appConfig.JwtSettings.SecretKey, os.Getenv("ENV"))

	g.JSON(http.StatusOK, decodedClaims)
}
