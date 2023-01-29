package api

import (
	"fmt"

	zabbixgosdk "github.com/Spartan0nix/zabbix-go-sdk"
)

// InitService is used to return a new ZabbixService after executing connectivity test.
func InitService(url string) (*zabbixgosdk.ZabbixService, error) {
	client := zabbixgosdk.NewZabbixService()

	client.Auth.Client.Url = url
	client.HostGroup.Client.Url = url

	if err := client.Auth.Client.CheckConnectivity(); err != nil {
		return nil, err
	}

	return client, nil
}

// Authenticate is used to retrieve an Api token for the HostGroup service.
func Authenticate(client *zabbixgosdk.ZabbixService, user string, password string) error {
	u := &zabbixgosdk.ApiUser{
		User: user,
		Pwd:  password,
	}

	res, err := client.Auth.GetCredentials(u.User, u.Pwd)
	if err != nil {
		return err
	}

	if len(res.Result) == 0 {
		return fmt.Errorf("no token were returned during the authentification phase")
	}

	var token string
	err = client.Auth.Client.ConvertResponse(*res, &token)
	if err != nil {
		return err
	}

	client.HostGroup.Client.Token = token

	return nil
}
