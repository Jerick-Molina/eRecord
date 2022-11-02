package security

import (
	"eRecord/util"
	"fmt"
	"reflect"
	"time"

	"github.com/dgrijalva/jwt-go/v4"
	"github.com/gin-gonic/gin"
)

var secretKey = []byte("p8cafxzquew4juy1rk9f")

var token = jwt.New(jwt.SigningMethodHS256)

func CreateAccessToken(usrId int, role string, companyId int) (string, error) {

	claims := token.Claims.(jwt.MapClaims)
	claims["exp"] = time.Now().Add(time.Hour * 10).Unix()
	claims["iss"] = "localhost"
	claims["userId"] = usrId
	claims["role"] = role
	claims["companyId"] = companyId

	tokenString, err := token.SignedString(secretKey)

	if err != nil {
		//
		return "", nil
	}

	return tokenString, nil
}

func CreateCompanyInviteToken(compName string, givenRole string) (string, error) {
	claims := token.Claims.(jwt.MapClaims)
	claims["exp"] = time.Now().Add(time.Minute * 5).Unix()
	claims["iss"] = "localhost"
	claims["invitationCode"] = util.RandomChars(15)
	claims["givenRole"] = givenRole
	tokenString, err := token.SignedString(secretKey)

	if err != nil {
		//
		return "", err
	}

	return tokenString, nil
}

func TokenReader(token string) (jwt.MapClaims, error) {
	var err error
	if token != "" {
		bearer := token[:6]
		token = token[7:]
		if bearer != "Bearer" {
			return nil, fmt.Errorf("Not bearer token")
		}
		token, err := jwt.Parse(token, func(t *jwt.Token) (interface{}, error) {
			if _, ok := t.Method.(*jwt.SigningMethodHMAC); !ok {
				return nil, fmt.Errorf("Token not valid")
			}

			return secretKey, nil
		})

		if err != nil {

			return nil, err
		}
		if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {

			return claims, nil

		}

		return nil, err
	}

	return nil, err
}

func GetJwtMap(ctx *gin.Context) (jwt.MapClaims, error) {
	claims, exist := ctx.Get("claims")
	if exist == false {
		err := fmt.Errorf("MAJOR INTERNAL ERR")

		return nil, err
	}
	ref := reflect.ValueOf(claims)
	c := ref.Interface()
	return c.(jwt.MapClaims), nil
}
