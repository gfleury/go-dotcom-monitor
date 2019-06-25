package main

import (
	"context"
	"fmt"
	"strings"
	"time"

	"github.com/gfleury/go-dotcom-monitor"
)

func main() {
	basicAuth := dotcommonitor.BasicAuth{UserName: "", Password: ""}
	ctx, cancel := context.WithTimeout(context.Background(), 60000*time.Millisecond)

	defer cancel()

	client := dotcommonitor.NewAPIClient(
		dotcommonitor.NewConfiguration(),
	)

	response, _, err := client.LoginApi.Login(ctx, basicAuth)
	if err != nil {
		fmt.Printf("%s\n", err.Error())
	}

	fmt.Println(response)

	platforms, _, err := client.PlatformsApi.GetPlataforms(ctx)
	if err != nil {
		fmt.Printf("%s\n", err.Error())
	}

	fmt.Println(platforms)

	for _, platform := range platforms {
		if platform.Available {
			fmt.Printf("Showing devices for platform, %v\n", platform.Name)
			devices, _, err := client.DeviceApi.GetDevicesPlataform(ctx, strings.ToLower(platform.Name))
			if err != nil {
				fmt.Printf("%s\n", err.Error())
			}

			fmt.Println(devices)
		}
	}
}
