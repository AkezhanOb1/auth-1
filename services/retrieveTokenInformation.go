package services

import (
	"context"
	pb "github.com/AkezhanOb1/auth/api"
	"github.com/AkezhanOb1/auth/configs"
	"github.com/dgrijalva/jwt-go"
)


//RetrieveTokenInformation is
func (*TokenServer) RetrieveTokenInformation(ctx context.Context, request *pb.RetrieveTokenInformationRequest) (*pb.RetrieveTokenInformationResponse, error) {

	hmacSecret := []byte(configs.JwtKey)

	//validating and parsing token
	token, err := jwt.Parse(request.GetAccessToken(), func(token *jwt.Token) (interface{}, error) {
		// check token signing method etc
		return hmacSecret, nil
	})

	if err != nil {
		return nil, err
	}

	claim := token.Claims.(jwt.MapClaims)

	return &pb.RetrieveTokenInformationResponse{
		Email:     claim["aud"].(string),
		ExpiresAt: int64(claim["exp"].(float64)),
		IssuedAt:  int64(claim["iat"].(float64)),
	}, err
}
