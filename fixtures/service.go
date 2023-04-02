package main

import (
	"encoding/json"
	"flag"
	"log"
	"os"

	zabbixgosdk "github.com/Spartan0nix/zabbix-go-sdk/v2"
)

var URL string
var USER string
var PASSWORD string
var ACTION string

func init() {
	flag.StringVar(&URL, "url", "", "zabbix server api url")
	flag.StringVar(&USER, "user", "", "zabbix user used to authenticate")
	flag.StringVar(&PASSWORD, "password", "", "zabbix user password")
	flag.StringVar(&ACTION, "action", "", "action to execute (export | import)")
	flag.Parse()
}

// export is used to export the list of services configured on a Zabbix server
func export(service *zabbixgosdk.ZabbixService) error {
	res, err := service.Service.Get(&zabbixgosdk.ServiceGetParameters{
		SelectParents: []string{
			"serviceid",
			"name",
		},
		SelectTags: []string{
			"tag",
			"value",
		},
		SelectProblemTags: []string{
			"tag",
			"operator",
			"value",
		},
		CommonGetParameters: zabbixgosdk.CommonGetParameters{
			Output: []string{
				"algorithm",
				"name",
				"sortorder",
				"weight",
				"propagation_rule",
				"propagation_value",
			},
			Sortfield: []string{
				"serviceid",
			},
			SortOrder: []string{
				"ASC",
			},
		},
	})
	if err != nil {
		return err
	}

	b, err := json.Marshal(&res)
	if err != nil {
		return err
	}

	err = os.WriteFile("fixtures/export_services.json", b, 0644)
	if err != nil {
		return err
	}

	return nil
}

func _import(service *zabbixgosdk.ZabbixService) error {
	b, err := os.ReadFile("fixtures/export_services.json")
	if err != nil {
		return err
	}

	services := make([]*zabbixgosdk.ServiceGetResponse, 0)
	err = json.Unmarshal(b, &services)
	if err != nil {
		return err
	}

	for _, s := range services {
		params := &zabbixgosdk.ServiceCreateParameters{
			Algorithm:        s.Algorithm,
			Name:             s.Name,
			SortOrder:        s.SortOrder,
			Weight:           s.Weight,
			PropagationRule:  s.PropagationRule,
			PropagationValue: s.PropagationValue,
		}

		if len(s.Parents) > 0 {
			parentsName := make([]string, 0)

			for _, p := range s.Parents {
				parentsName = append(parentsName, p.Name)
			}

			res, err := service.Service.Get(&zabbixgosdk.ServiceGetParameters{
				CommonGetParameters: zabbixgosdk.CommonGetParameters{
					Search: map[string][]string{
						"name": parentsName,
					},
					SearchWildcardsEnabled: true,
				},
			})

			if err != nil {
				return err
			}

			parents := make([]*zabbixgosdk.ServiceId, 0)
			for _, s := range res {
				parents = append(parents, &zabbixgosdk.ServiceId{
					ServiceId: s.ServiceId,
				})
			}

			params.Parents = parents
		}

		_, err := service.Service.Create(params)
		if err != nil {
			return err
		}
	}

	return nil
}

func main() {
	service := zabbixgosdk.NewZabbixService()

	service.Auth.Client.Url = URL
	service.Service.Client.Url = URL
	service.Auth.User = &zabbixgosdk.ApiUser{
		User: USER,
		Pwd:  PASSWORD,
	}

	err := service.Authenticate()
	if err != nil {
		log.Fatalf("error during authentication.\nReason : %v", err)
	}

	switch ACTION {
	case "export":
		err = export(service)
		if err != nil {
			log.Fatalf("error during export process.\nReason : %v", err)
		}
	case "import":
		err = _import(service)
		if err != nil {
			log.Fatalf("error during import process.\nReason : %v", err)
		}
	default:
		log.Printf("unsupported action '%s'\nSupported actions are : 'export' and 'import'", ACTION)
	}
}
