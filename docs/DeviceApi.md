# \DeviceApi

All URIs are relative to *https://api.dotcom-monitor.com/config_api_v1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**AddDevice**](DeviceApi.md#AddDevice) | **Put** /devices | Get device list by platform.
[**DeleteDevice**](DeviceApi.md#DeleteDevice) | **Delete** /device/{device_id} | Get Device info.
[**DisableDeviceAlert**](DeviceApi.md#DisableDeviceAlert) | **Post** /device/{device_id}/DisableAlert/ | Get Device info.
[**EditDevice**](DeviceApi.md#EditDevice) | **Post** /device/{device_id} | Edit device list by platform.
[**GetDevice**](DeviceApi.md#GetDevice) | **Get** /device/{device_id} | Get Device info.
[**GetDevicesPlataform**](DeviceApi.md#GetDevicesPlataform) | **Get** /devices/{platform} | Get device list by platform.
[**GetTasks**](DeviceApi.md#GetTasks) | **Get** /device/{device_id}/tasks | Get Device tasks.


# **AddDevice**
> Result AddDevice(ctx, device)
Get device list by platform.

Get device list by platform.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **device** | [**Device**](Device.md)| Device Object | 

### Return type

[**Result**](Result.md)

### Authorization

[api_key](../README.md#api_key)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **DeleteDevice**
> Result DeleteDevice(ctx, deviceId)
Get Device info.

Get Device info.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **deviceId** | **string**| device id | 

### Return type

[**Result**](Result.md)

### Authorization

[api_key](../README.md#api_key)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **DisableDeviceAlert**
> InlineResponse200 DisableDeviceAlert(ctx, deviceId, optional)
Get Device info.

Get Device info.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **deviceId** | **string**| device id | 
 **optional** | ***DisableDeviceAlertOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a DisableDeviceAlertOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **alertSilencePeriod** | [**optional.Interface of AlertSilencePeriod**](AlertSilencePeriod.md)|  | 

### Return type

[**InlineResponse200**](inline_response_200.md)

### Authorization

[api_key](../README.md#api_key)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **EditDevice**
> Result EditDevice(ctx, deviceId, device)
Edit device list by platform.

Edit device list by platform.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **deviceId** | **string**| Device id | 
  **device** | [**Device**](Device.md)| Device Object | 

### Return type

[**Result**](Result.md)

### Authorization

[api_key](../README.md#api_key)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetDevice**
> interface{} GetDevice(ctx, deviceId)
Get Device info.

Get Device info.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **deviceId** | **string**| device id | 

### Return type

[**interface{}**](interface{}.md)

### Authorization

[api_key](../README.md#api_key)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetDevicesPlataform**
> []int32 GetDevicesPlataform(ctx, platform)
Get device list by platform.

Get device list by platform.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **platform** | **string**| Platform name | 

### Return type

**[]int32**

### Authorization

[api_key](../README.md#api_key)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetTasks**
> []int32 GetTasks(ctx, deviceId)
Get Device tasks.

Get Device tasks.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **deviceId** | **string**| device id | 

### Return type

**[]int32**

### Authorization

[api_key](../README.md#api_key)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

