package auth

import (
	"fmt"
	"log"
	"time"

	"github.com/akekapong78/workflow/internal/constant"
	"github.com/golang-jwt/jwt/v5"
	"golang.org/x/crypto/bcrypt"
)

func GenerateToken(username string, role constant.UserRole, secret string) (string, error) {
	t := jwt.NewWithClaims(jwt.SigningMethodHS256, &jwt.RegisteredClaims{
		Audience:  jwt.ClaimStrings{username, string(role)},
		ExpiresAt: jwt.NewNumericDate(time.Now().Add(15 * time.Minute)),
	})

	signedToken, err := t.SignedString([]byte(secret))
	if err != nil {
		log.Println("error signing key")
		return signedToken, err
	}

	return signedToken, nil
}

func VerifyToken(tokenString string, secret string) (*jwt.Token, error) {
	t, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
			// the algorithm used in the token’s header (alg) is a type of HMAC (Hash-based Message Authentication Code)
		}

		// hmacSampleSecret is a []byte containing your secret, e.g. []byte("my_secret_key")
		return []byte(secret), nil
	})

	if err != nil {
		return nil, err
	}

	// Return the verified token
	return t, nil
}


func CheckPasswordHash(password, hash string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	return err == nil
}

func HashPassword(password string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), 14)
	return string(bytes), err
}
