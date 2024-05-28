# \AsynqAPI

All URIs are relative to *http://127.0.0.1:3013/goodsCenterLogic*

Method | HTTP request | Description
------------- | ------------- | -------------
[**AsynqV1AddAggTaskGet**](AsynqAPI.md#AsynqV1AddAggTaskGet) | **Get** /asynq/v1/addAggTask | asynq-添加聚合任务
[**AsynqV1AddTaskGet**](AsynqAPI.md#AsynqV1AddTaskGet) | **Get** /asynq/v1/addTask | asynq-add异步任务



## AsynqV1AddAggTaskGet

> InternalApiDtoGroupDeliveryTaskAddResponse AsynqV1AddAggTaskGet(ctx).Env(env).Execute()

asynq-添加聚合任务



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/yituoshiniao/openapi-client-go"
)

func main() {
	env := "env_example" // string | 环境变量,默认线上; sandbox 沙箱环境, production 生产环境 (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AsynqAPI.AsynqV1AddAggTaskGet(context.Background()).Env(env).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AsynqAPI.AsynqV1AddAggTaskGet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `AsynqV1AddAggTaskGet`: InternalApiDtoGroupDeliveryTaskAddResponse
	fmt.Fprintf(os.Stdout, "Response from `AsynqAPI.AsynqV1AddAggTaskGet`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiAsynqV1AddAggTaskGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **env** | **string** | 环境变量,默认线上; sandbox 沙箱环境, production 生产环境 | 

### Return type

[**InternalApiDtoGroupDeliveryTaskAddResponse**](InternalApiDtoGroupDeliveryTaskAddResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## AsynqV1AddTaskGet

> InternalApiDtoAsynqEmailDeliveryTaskAddResponse AsynqV1AddTaskGet(ctx).Env(env).Execute()

asynq-add异步任务



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/yituoshiniao/openapi-client-go"
)

func main() {
	env := "env_example" // string | 环境变量,默认线上; sandbox 沙箱环境, production 生产环境 (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AsynqAPI.AsynqV1AddTaskGet(context.Background()).Env(env).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AsynqAPI.AsynqV1AddTaskGet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `AsynqV1AddTaskGet`: InternalApiDtoAsynqEmailDeliveryTaskAddResponse
	fmt.Fprintf(os.Stdout, "Response from `AsynqAPI.AsynqV1AddTaskGet`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiAsynqV1AddTaskGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **env** | **string** | 环境变量,默认线上; sandbox 沙箱环境, production 生产环境 | 

### Return type

[**InternalApiDtoAsynqEmailDeliveryTaskAddResponse**](InternalApiDtoAsynqEmailDeliveryTaskAddResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

