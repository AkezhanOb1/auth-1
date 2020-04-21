package services

import (
	"context"
	pb "github.com/AkezhanOb1/auth/api"
	"github.com/AkezhanOb1/auth/configs"
	"github.com/dgrijalva/jwt-go"
	"time"
)

type TokenServer struct {}

//GenerateToken is
func (*TokenServer) GenerateToken(ctx context.Context, request *pb.GenerateTokenRequest) (*pb.GenerateTokenResponse, error) {
	now := time.Now()
	//determining token's lifetime
	expirationTime := int64((5 * time.Hour).Seconds())

	//create the JWT claims, which includes the username and expiry time
	//issue time and the issuer
	claims := jwt.StandardClaims{
		Audience: request.Credentials.GetEmail(),
		// In JWT, the expiry time is expressed as unix milliseconds
		ExpiresAt: now.Add(4 * time.Hour).Unix(),
		IssuedAt:  now.Unix(),
		Issuer:    "http://127.0.0.1:9001/token",
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	//create a token
	tokenString, err := token.SignedString([]byte(configs.JwtKey))
	if err != nil {
		return nil, err
	}

	return &pb.GenerateTokenResponse{
		Token: &pb.Token{
			AccessToken:  tokenString,
			TokenType:    "bearer",
			ExpiresIn:    expirationTime,
		},
	}, nil
}
