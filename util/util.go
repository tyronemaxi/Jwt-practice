package util

import (
	"crypto/md5"
	"fmt"
	"github.com/gofrs/uuid"
	"github.com/golang-jwt/jwt"
	"strconv"
	"time"
)

const
(
	PROJECTISSUE = "jwt-practice"
	SHORTID_DIGITS = "abcdefghijklmnopqrstuvwxyz0123456789"
)

var projectSecret = []byte("you-will-nerver-guess")

type Claims struct {
	Username string `json:"username"`
	Password string `json:"password"`
	jwt.StandardClaims
}

// Generate JWT token
func GenerateJwtToken(username, password string) (string, error) {
	nowTime := time.Now()
	expireTime := nowTime.Add(2 * time.Hour)

	claims := Claims{
		username,
		password,
		jwt.StandardClaims{
			ExpiresAt: expireTime.Unix(),
			Issuer:    PROJECTISSUE,
		},
	}

	tokenClaims := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	token, err := tokenClaims.SignedString(projectSecret)
	return token, err
}

func ParseJwtToken(token string) (*Claims, error) {
	tokenClaims, err := jwt.ParseWithClaims(token, &Claims{}, func(token *jwt.Token) (interface{}, error) {
		return projectSecret, nil
	})

	if tokenClaims != nil {
		if claims, ok := tokenClaims.Claims.(*Claims); ok && tokenClaims.Valid {
			return claims, nil
		}
	}

	return nil, err
}

func GenerateUUid(prefix string) (string, error) {
	Uid, err := uuid.NewV4()
	if err != nil {
		return "", err
	}

	shortidStr := UUIDToShortID(Uid.String())
	if prefix != "" {
		shortidStr = prefix + "-" + shortidStr
	}

	return shortidStr, nil

}

func UUIDToShortID(UUID string) string {
	// 32uuid -> 32md5 hex
	data := []byte(UUID)
	hash := md5.Sum(data)
	md5str := fmt.Sprintf("%x", hash)

	var result []byte
	for i := 0; i < 16; i++ {
		// parse 2bit char from 16base to 10base
		index, _ := strconv.ParseUint(md5str[2*i:2*i+2], 16, 32)
		result = append(result, SHORTID_DIGITS[index%34])
	}
	return string(result)
}
