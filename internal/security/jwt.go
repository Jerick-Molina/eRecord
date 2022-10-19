package security

import (
	"fmt"
	"time"

	"github.com/dgrijalva/jwt-go/v4"
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

func CreateCompanyValidToken(compName string, uniqueId string) (string, error) {
	claims := token.Claims.(jwt.MapClaims)
	claims["exp"] = time.Now().Add(time.Minute * 5).Unix()
	claims["iss"] = "localhost"
	claims["companyName"] = compName
	claims["uniqueId"] = uniqueId
	tokenString, err := token.SignedString(secretKey)

	if err != nil {
		//
		return "", nil
	}

	return tokenString, nil
}

func CreateInvitationToken() {

}
func TokenReader(token string) (jwt.MapClaims, error) {
	var err error
	if token != "" {
		token = token[7:]
		token, err := jwt.Parse(token, func(t *jwt.Token) (interface{}, error) {
			if _, ok := t.Method.(*jwt.SigningMethodHMAC); !ok {
				return nil, nil
			}
			return secretKey, nil
		})
		if err != nil {
			fmt.Println(err.Error())
			return nil, err
		}
		if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {

			return claims, nil

		}
		return nil, err
	}
	return nil, err
}
