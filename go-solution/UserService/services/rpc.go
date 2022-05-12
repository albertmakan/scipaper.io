package services

import (
	"fmt"

	"github.com/albertmakan/scipaper.io/go-solution/UserService/helpers"
	"github.com/dgrijalva/jwt-go"
)

type RPC struct {
	UserService *UserService
}

func (rpc *RPC) IsLoggedIn(token string, reply *bool) error {
	_, err := helpers.VerifyJwtToken(token)
	*reply = err == nil
	return err
}

func (rpc *RPC) GetName(token string, reply *string) error {
	claims, err := helpers.VerifyJwtToken(token)
	if err != nil { return err }
	username := fmt.Sprintf("%v", claims.(jwt.MapClaims)["name"])
	user := rpc.UserService.FindByUsername(username)
	*reply = user.FirstName+" "+user.LastName+" "+username
	return err
}