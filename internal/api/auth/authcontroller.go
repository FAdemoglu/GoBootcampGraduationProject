package auth

import (
	"github.com/FAdemoglu/graduationproject/internal/config"
	"github.com/FAdemoglu/graduationproject/internal/domain/users"
	jwtHelper "github.com/FAdemoglu/graduationproject/pkg/jwt"
	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	"net/http"
	"os"
	"time"
)

type AuthController struct {
	appConfig *config.Configuration
	r         users.UserRepository
}

func NewAuthController(appConfig *config.Configuration, r users.UserRepository) *AuthController {
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
// @Router /auth/login [post]
func (c *AuthController) Login(g *gin.Context) {
	var req LoginRequest

	if err := g.ShouldBind(&req); err != nil {
		g.JSON(http.StatusBadRequest, gin.H{
			"error_message": "Check your request body.",
		})
		return
	}
	user := c.r.GetByUsernameAndPassword(req.Username, req.Password)

	if user.Username == "" {
		g.JSON(http.StatusNotFound, gin.H{
			"error_message": "User not found!",
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
	token := jwtHelper.GenerateToken(jwtClaims, c.appConfig.JwtSettings.SecretKey)
	g.JSON(http.StatusOK, token)
}

// Register godoc
// @Summary Register into system with username and password
// @Tags Auth
// @Accept  json
// @Produce  json
// @Param loginRequest body RegisterRequest true "login informations"
// @Success 200
// @Failure 400 {object} map[string]string
// @Failure 401 {object} map[string]string
// @Failure 404 {object} map[string]string
// @Failure 500 {object} map[string]string
// @Router /auth/register [post]
func (c *AuthController) Register(g *gin.Context) {
	var req RegisterRequest

	if err := g.ShouldBind(&req); err != nil {
		g.JSON(http.StatusBadRequest, gin.H{
			"error_message": "Check your request body.",
		})
		return
	}

	user := c.r.GetByUsername(req.Username)
	if user.Username != "" {
		g.JSON(http.StatusBadRequest, gin.H{
			"error_message": "Bu kullanıcı adı ile kullanıcı bulunuyor",
		})
		return
	}

	if req.Password != req.ConfirmPassword {
		g.JSON(http.StatusForbidden, gin.H{
			"error_message": "Şifreler uyuşmuyor",
		})
		return
	}
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
	g.JSON(http.StatusOK, token)

}

func (c *AuthController) VerifyToken(g *gin.Context) {
	token := g.GetHeader("Authorization")
	decodedClaims := jwtHelper.VerifyToken(token, c.appConfig.JwtSettings.SecretKey, os.Getenv("ENV"))

	g.JSON(http.StatusOK, decodedClaims)
}
