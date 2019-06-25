# Go API client for dotcom-monitor

dotcom-monitor API

## Overview
Golang API client for https://www.dotcom-monitor.com/. Only cover Device and Task, other endpoints might be added in the future. 

## Installation
Put the package under your project folder and add the following in import:
```golang
import "github.com/gfleury/go-dotcom-monitor"
```
There is an example under **cmd/main.go**(cmd/main.go).

## Documentation for API Endpoints

All URIs are relative to *https://api.dotcom-monitor.com/config_api_v1*

Class | Method | HTTP request | Description
------------ | ------------- | ------------- | -------------
*DeviceApi* | [**AddDevice**](docs/DeviceApi.md#adddevice) | **Put** /devices | Get device list by platform.
*DeviceApi* | [**DeleteDevice**](docs/DeviceApi.md#deletedevice) | **Delete** /device/{device_id} | Get Device info.
*DeviceApi* | [**DisableDeviceAlert**](docs/DeviceApi.md#disabledevicealert) | **Post** /device/{device_id}/DisableAlert/ | Get Device info.
*DeviceApi* | [**EditDevice**](docs/DeviceApi.md#editdevice) | **Post** /device/{device_id} | Edit device list by platform.
*DeviceApi* | [**GetDevice**](docs/DeviceApi.md#getdevice) | **Get** /device/{device_id} | Get Device info.
*DeviceApi* | [**GetDevicesPlataform**](docs/DeviceApi.md#getdevicesplataform) | **Get** /devices/{platform} | Get device list by platform.
*DeviceApi* | [**GetTasks**](docs/DeviceApi.md#gettasks) | **Get** /device/{device_id}/tasks | Get Device tasks.
*PlatformsApi* | [**GetPlataforms**](docs/PlatformsApi.md#getplataforms) | **Get** /platforms | Return list of available platforms
*TaskApi* | [**AddTask**](docs/TaskApi.md#addtask) | **Put** /tasks | Create new task.
*TaskApi* | [**DeleteTask**](docs/TaskApi.md#deletetask) | **Delete** /task/{task_id} | Delete Task info.
*TaskApi* | [**EditTask**](docs/TaskApi.md#edittask) | **Post** /task/{task_id} | Edit task.
*TaskApi* | [**GetTask**](docs/TaskApi.md#gettask) | **Get** /task/{task_id} | Get Task info.
*TaskApi* | [**GetTasks**](docs/TaskApi.md#gettasks) | **Get** /device/{device_id}/tasks | Get Device tasks.


## Documentation For Models

 - [AlertSilencePeriod](docs/AlertSilencePeriod.md)
 - [Device](docs/Device.md)
 - [DeviceNotifications](docs/DeviceNotifications.md)
 - [DeviceNotificationsNotificationGroups](docs/DeviceNotificationsNotificationGroups.md)
 - [InlineResponse200](docs/InlineResponse200.md)
 - [Platform](docs/Platform.md)
 - [PlatformPackages](docs/PlatformPackages.md)
 - [Result](docs/Result.md)
 - [Task](docs/Task.md)
 - [TaskGetParams](docs/TaskGetParams.md)


## Documentation For Authentication

## Authentication
- **Type**: User/Password

Example
```golang
basicAuth := dotcommonitor.BasicAuth{UserName: "", Password: ""}
response, _, err := client.LoginApi.Login(ctx, basicAuth)
...
r, err := client.Service.Operation(ctx, args)
```

## Author

George Fleury

