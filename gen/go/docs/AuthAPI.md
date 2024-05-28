# \AuthAPI

All URIs are relative to *http://127.0.0.1:3013/goodsCenterLogic*

Method | HTTP request | Description
------------- | ------------- | -------------
[**AuthV1TokenGenerateGet**](AuthAPI.md#AuthV1TokenGenerateGet) | **Get** /auth/v1/token/generate | jwt-token生成及校验



## AuthV1TokenGenerateGet

> InternalApiDtoAppJwtTokenSwgResponse AuthV1TokenGenerateGet(ctx).UserId(userId).UserName(userName).Env(env).Execute()

jwt-token生成及校验



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
	userId := "userId_example" // string | UserID 用户id
	userName := "userName_example" // string | UserName 用户名  example:张三
	env := "env_example" // string | 环境变量,默认线上; sandbox 沙箱环境, production 生产环境 (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AuthAPI.AuthV1TokenGenerateGet(context.Background()).UserId(userId).UserName(userName).Env(env).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AuthAPI.AuthV1TokenGenerateGet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `AuthV1TokenGenerateGet`: InternalApiDtoAppJwtTokenSwgResponse
	fmt.Fprintf(os.Stdout, "Response from `AuthAPI.AuthV1TokenGenerateGet`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiAuthV1TokenGenerateGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **userId** | **string** | UserID 用户id | 
 **userName** | **string** | UserName 用户名  example:张三 | 
 **env** | **string** | 环境变量,默认线上; sandbox 沙箱环境, production 生产环境 | 

### Return type

[**InternalApiDtoAppJwtTokenSwgResponse**](InternalApiDtoAppJwtTokenSwgResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

