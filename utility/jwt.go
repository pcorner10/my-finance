package utility

import (
	"errors"
	"fmt"
	"os"
	"strconv"
	"strings"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v4"
	"github.com/pcorner10/my-finance/api/models"
)

var privateKey = []byte(os.Getenv("JWT_PRIVATE_KEY"))

func GenerateJWT(user models.User) (string, error) {
	tokenTTL, _ := strconv.Atoi(os.Getenv("TOKEN_TTL"))
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"id":  user.ID,
		"iat": time.Now().Unix(),
		"eat": time.Now().Add(time.Second * time.Duration(tokenTTL)).Unix(),
	})
	return token.SignedString(privateKey)
}

func ValidateJWT(context *gin.Context) error {
	token, err := getToken(context)
	if err != nil {
		return err
	}
	_, ok := token.Claims.(jwt.MapClaims)
	if ok && token.Valid {
		return nil
	}
	return errors.New("invalid token provided")
}

func CurrentUser(context *gin.Context) (models.User, error) {
	err := ValidateJWT(context)
	if err != nil {
		return models.User{}, err
	}
	token, _ := getToken(context)
	
	var claims jwt.MapClaims

	if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		fmt.Println(claims["id"], claims["eat"])
	} else {
		fmt.Println(err)
	}

	userId := uint(claims["id"].(float64))

	user, err := models.FindUserById(userId)
	if err != nil {
		return models.User{}, err
	}
	return user, nil
}

func getToken(context *gin.Context) (*jwt.Token, error) {
	tokenString := getTokenFromRequest(context)
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		// Don't forget to validate the alg is what you expect:
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}

		// hmacSampleSecret is a []byte containing your secret, e.g. []byte("my_secret_key")
		return privateKey, nil
	})
	return token, err
}

func getTokenFromRequest(context *gin.Context) string {
	bearerToken := context.Request.Header.Get("Authorization")

	splitToken := strings.Split(bearerToken, " ")
	if len(splitToken) == 2 {
		fmt.Println(splitToken[1], "samina")
		return splitToken[1]
	}
	return ""
}
