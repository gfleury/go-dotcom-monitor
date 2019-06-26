package main

import (
	"context"
	"encoding/json"
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
			fmt.Printf("Getting devices %s\n", fmt.Sprintf("%d", devices[0]))
			device, _, err := client.DeviceApi.GetDevice(ctx, fmt.Sprintf("%d", devices[0]))
			if err != nil {
				fmt.Printf("Device: %s\n", err.Error())
			}

			fmt.Println(device)
		}
	}

	newDevice := dotcommonitor.Device{
		AvoidSimultaneousChecks: true,
		FalsePositiveCheck:      false,
		Locations:               []int32{1, 2, 3, 4, 6, 11, 13, 14, 15, 18, 19, 23, 43, 68, 97, 113, 118, 138, 153, 233},
		SendUptimeAlert:         true,
		StatusDescription:       "ACTIVE",
		Postpone:                false,
		Frequency:               14400,
		Name:                    "www.google.com",
		PlatformId:              1,
	}

	result, _, err := client.DeviceApi.AddDevice(ctx, newDevice)
	if err != nil {
		fmt.Printf("%s\n", err.Error())
	}

	fmt.Println(result)

	newTask := dotcommonitor.Task{
		RequestType:      "GET",
		Url:              "http://www.google.com",
		FullPageDownload: true,
		DeviceId:         result.Result,
		TaskTypeId:       1,
		Name:             "www.google.com",
		Timeout:          30,
		GetParams:        []dotcommonitor.TaskGetParams{},
		PostParams:       []dotcommonitor.TaskGetParams{},
		HeaderParams:     []dotcommonitor.TaskGetParams{},
		DNSResolveMode:   "Device Cached",
		DNSserverIP:      nil,
	}

	result, _, err = client.TaskApi.AddTask(ctx, newTask)

	if err != nil {
		b, _ := json.Marshal(newTask)
		fmt.Println(string(b))
		fmt.Printf("%s\n", err.Error())
	}

	fmt.Println(result)

}
