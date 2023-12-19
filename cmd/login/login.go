package login

import (
	"fmt"
	"time"

	"github.com/dgrijalva/jwt-go"
)

func CreateJWT() string {
	claims := jwt.MapClaims{
		"userID": "test@gmail.com",
		"exp":    time.Now().Add(time.Hour * 24).Unix(),
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, _ := token.SignedString([]byte("secret"))
	fmt.Println(tokenString)
	return tokenString
}

// トークンを複合化
func ParseJWT(tokenString string) (jwt.MapClaims, error) {
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok { // 署名の検証
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}
		return []byte("secret"), nil // 署名の検証に成功したらシークレットキーを返す
	})
	if err != nil {
		return nil, err
	}
	if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid { // トークンの有効性の検証
		fmt.Println(claims["userID"], claims["exp"])
		return claims, nil
	}
	return nil, err
}
