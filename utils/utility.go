package utils

import (
	"crypto/rand"
	"errors"
	"fmt"
	"math/big"
	"os"
	"time"

	"github.com/golang-jwt/jwt/v5"
	"golang.org/x/crypto/bcrypt"
)

func HashPassword(password string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), 14)
	if err != nil {
		return "", err
	}
	return string(bytes), nil
}

func VerifyPassword(hashedPassword string, plainPassword string) (bool, error) {
	err := bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(plainPassword))
	if err != nil {
		return false, err
	}
	return true, nil
}

func GenerateOTP(length int) (string, error) {
	if length <= 0 {
		return "", fmt.Errorf("invalid length: must be > 0")
	}

	// Lower bound = 10^(length-1), Upper bound = 10^length - 1
	min := new(big.Int).Exp(big.NewInt(10), big.NewInt(int64(length-1)), nil)
	max := new(big.Int).Exp(big.NewInt(10), big.NewInt(int64(length)), nil)
	rangeNum := new(big.Int).Sub(max, min) // (10^length - 10^(length-1))

	// Random number in [0, rangeNum)
	n, err := rand.Int(rand.Reader, rangeNum)
	if err != nil {
		return "", errors.New("something went wrong")
	}

	// Shift into desired range [min, max)
	n.Add(n, min)

	return n.String(), nil
}

func JwtSignHour(userinfo interface{}, hournum int) (string ,error) {
	JWT_SECRET_KEY := os.Getenv("JWT_SECRET_KEY")
	token := jwt.NewWithClaims(jwt.SigningMethodHS256,jwt.MapClaims{
		"userinfo":userinfo,
		"expr": time.Now().Add(time.Hour *  time.Duration(hournum)).Unix(),
	})

	tokenString,err := token.SignedString([]byte(JWT_SECRET_KEY))
	if err != nil {
		return "" ,err 
	} 

	return tokenString ,nil
}

func JwtSignMinutes(userinfo interface{}, minutenum int) (string ,error) {
	JWT_SECRET_KEY := os.Getenv("JWT_SECRET_KEY")
	token := jwt.NewWithClaims(jwt.SigningMethodHS256,jwt.MapClaims{
		"userinfo":userinfo,
		"expr": time.Now().Add(time.Minute *  time.Duration(minutenum)).Unix(),
	})
    tokenString,err := token.SignedString([]byte(JWT_SECRET_KEY))
	if err != nil {
		return "",err 
	} 

	return tokenString ,nil
}


