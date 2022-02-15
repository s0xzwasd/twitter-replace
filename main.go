package main

import (
	"context"
	"github.com/Nerzal/gocloak/v10"
)

func main() {
	client := gocloak.NewClient("https://mycool.keycloak.instance")
	ctx := context.Background()
	token, err := client.LoginAdmin(ctx, "user", "password", "realmName")
	if err != nil {
		panic("Something wrong with the credentials or url")
	}

	user := gocloak.User{
		FirstName: gocloak.StringP("Bob"),
		LastName:  gocloak.StringP("Uncle"),
		Email:     gocloak.StringP("something@really.wrong"),
		Enabled:   gocloak.BoolP(true),
		Username:  gocloak.StringP("CoolGuy"),
	}

	_, err = client.CreateUser(ctx, token.AccessToken, "realm", user)
	_ = client.UpdateRequiredAction(ctx, "123", "str", gocloak.RequiredActionProviderRepresentation{
		Alias:         nil,
		Config:        nil,
		DefaultAction: nil,
		Enabled:       nil,
		Name:          nil,
		Priority:      nil,
		ProviderID:    nil,
	})
	if err != nil {
		return
	}
	if err != nil {
		panic("Oh no!, failed to create user :(")
	}
}
