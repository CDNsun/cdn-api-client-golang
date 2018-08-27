package main

import (
	"./cdnsun"
	"fmt"
)

func main() {
	username := "YOUR_API_USERNAME"
	password := "YOUR_API_PASSWORD"
	ID := "YOUR_CDN_SERVICE_ID"

	apiClient, _ := cdnsun.New(username, password)

	result, err := apiClient.Get(&cdnsun.Options{
		Url: "cdns",
	})
	fmt.Printf("%+v\n", string(result))
	fmt.Printf("%+v\n", err)

	result, err = apiClient.Get(&cdnsun.Options{
		Url: "cdns/" + ID + "/reports",
		Data: map[string]interface{}{
			"type":   "GB",
			"period": "4h",
		},
	})
	fmt.Printf("%+v\n", string(result))
	fmt.Printf("%+v\n", err)

	result, err = apiClient.Post(&cdnsun.Options{
		Url: "cdns/" + ID + "/purge",
		Data: map[string]interface{}{
			"purge_paths": []string{
				"/path1.img",
				"/path2.img",
			},
		},
	})
	fmt.Printf("%+v\n", string(result))
	fmt.Printf("%+v\n", err)
}
