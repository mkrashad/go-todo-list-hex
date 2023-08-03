package auth

import (
	"errors"
	"github.com/golang-jwt/jwt"
	"github.com/mkrashad/go-todo/api-gw/pb"
	"time"
)

func GenerateToken(user *pb.User) (string, error) {
	timeNow := time.Now()
	token := jwt.New(jwt.SigningMethodHS256)
	claims := token.Claims.(jwt.MapClaims)
	claims["uid"] = user.Id
	claims["iat"] = timeNow.Unix()
	claims["exp"] = timeNow.Add(time.Minute * 10).Unix()

	tokenString, err := token.SignedString([]byte("secret"))
	if err != nil {
		return "", err
	}
	return tokenString, nil
}

func ParseToken(tokenString string) (jwt.MapClaims, error) {
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, jwt.ErrSignatureInvalid
		}
		return []byte("secret"), nil
	})

	if err != nil {
		return nil, err
	}

	claims, ok := token.Claims.(jwt.MapClaims)
	if ok && token.Valid {
		exp := claims["exp"]
		expTime := time.Unix(int64(exp.(float64)), 0)
		timeNow := time.Now()
		if timeNow.After(expTime) {
			return nil, errors.New("token expired")
		}

		return claims, nil
	}
	return nil, errors.New("could not extract claims")
}

