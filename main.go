package main

import (
	"fmt"
	"time"

	"github.com/golang-jwt/jwt/v5"
	"github.com/google/uuid"
)

// Color is used with go:generate stringer directive
//
//go:generate stringer -type=Color
type Color int

const (
	Red Color = iota
	Green
	Blue
)

func main() {
	id := uuid.New()
	fmt.Println("UUID:", id)

	// Create a JWT token
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"sub": id.String(),
		"exp": time.Now().Add(time.Hour).Unix(),
	})
	signed, err := token.SignedString([]byte("secret"))
	if err != nil {
		fmt.Println("error:", err)
		return
	}
	fmt.Println("Token:", signed)
}