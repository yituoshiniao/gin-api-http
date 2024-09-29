# \DefaultApi

All URIs are relative to *http://127.0.0.1:3013/goodsCenterLogic*

Method | HTTP request | Description
------------- | ------------- | -------------
[**V1CommonGenerateIdGet**](DefaultApi.md#V1CommonGenerateIdGet) | **Get** /v1/common/generateId | 雪花ID生成



## V1CommonGenerateIdGet

> InternalHandlerServicev1HttpGenerateIDResponse V1CommonGenerateIdGet(ctx, id).Num(num).Authorization(authorization).Execute()

雪花ID生成



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
	num := int32(56) // int32 | 生成id数量 1-1000
	id := int32(56) // int32 | ID
	authorization := "authorization_example" // string | 认证信息 eg:xxxx-xxxx-xxxx-xxx (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DefaultApi.V1CommonGenerateIdGet(context.Background(), id).Num(num).Authorization(authorization).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.V1CommonGenerateIdGet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `V1CommonGenerateIdGet`: InternalHandlerServicev1HttpGenerateIDResponse
	fmt.Fprintf(os.Stdout, "Response from `DefaultApi.V1CommonGenerateIdGet`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int32** | ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiV1CommonGenerateIdGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **num** | **int32** | 生成id数量 1-1000 | 

 **authorization** | **string** | 认证信息 eg:xxxx-xxxx-xxxx-xxx | 

### Return type

[**InternalHandlerServicev1HttpGenerateIDResponse**](InternalHandlerServicev1HttpGenerateIDResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

