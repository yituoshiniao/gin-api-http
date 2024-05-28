# \ExampleAPI

All URIs are relative to *http://127.0.0.1:3013/goodsCenterLogic*

Method | HTTP request | Description
------------- | ------------- | -------------
[**V1ExampleGetGet**](ExampleAPI.md#V1ExampleGetGet) | **Get** /v1/exampleGet | get接口示例
[**V1ExampleGetOneGet**](ExampleAPI.md#V1ExampleGetOneGet) | **Get** /v1/exampleGetOne | getOne接口示例
[**V1ExamplePostPost**](ExampleAPI.md#V1ExamplePostPost) | **Post** /v1/examplePost | post 接口 示例



## V1ExampleGetGet

> InternalApiDtoExampleGetResponse V1ExampleGetGet(ctx).CreateTime(createTime).QueryId(queryId).UserId(userId).Execute()

get接口示例



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
	createTime := int32(56) // int32 | create_time
	queryId := "queryId_example" // string | query_id
	userId := "userId_example" // string | user_id 用户id

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ExampleAPI.V1ExampleGetGet(context.Background()).CreateTime(createTime).QueryId(queryId).UserId(userId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ExampleAPI.V1ExampleGetGet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `V1ExampleGetGet`: InternalApiDtoExampleGetResponse
	fmt.Fprintf(os.Stdout, "Response from `ExampleAPI.V1ExampleGetGet`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiV1ExampleGetGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **createTime** | **int32** | create_time | 
 **queryId** | **string** | query_id | 
 **userId** | **string** | user_id 用户id | 

### Return type

[**InternalApiDtoExampleGetResponse**](InternalApiDtoExampleGetResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## V1ExampleGetOneGet

> InternalApiDtoExampleGetOneResponse V1ExampleGetOneGet(ctx).CreateTime(createTime).QueryId(queryId).UserId(userId).Execute()

getOne接口示例



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
	createTime := int32(56) // int32 | create_time
	queryId := "queryId_example" // string | query_id
	userId := "userId_example" // string | user_id 用户id

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ExampleAPI.V1ExampleGetOneGet(context.Background()).CreateTime(createTime).QueryId(queryId).UserId(userId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ExampleAPI.V1ExampleGetOneGet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `V1ExampleGetOneGet`: InternalApiDtoExampleGetOneResponse
	fmt.Fprintf(os.Stdout, "Response from `ExampleAPI.V1ExampleGetOneGet`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiV1ExampleGetOneGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **createTime** | **int32** | create_time | 
 **queryId** | **string** | query_id | 
 **userId** | **string** | user_id 用户id | 

### Return type

[**InternalApiDtoExampleGetOneResponse**](InternalApiDtoExampleGetOneResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## V1ExamplePostPost

> InternalApiDtoExamplePostResponse V1ExamplePostPost(ctx).Body(body).Execute()

post 接口 示例



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
	body := *openapiclient.NewInternalApiDtoExamplePostRequest("Title_example", "AppId_example", "Question_example") // InternalApiDtoExamplePostRequest | 请求参数

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ExampleAPI.V1ExamplePostPost(context.Background()).Body(body).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ExampleAPI.V1ExamplePostPost``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `V1ExamplePostPost`: InternalApiDtoExamplePostResponse
	fmt.Fprintf(os.Stdout, "Response from `ExampleAPI.V1ExamplePostPost`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiV1ExamplePostPostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **body** | [**InternalApiDtoExamplePostRequest**](InternalApiDtoExamplePostRequest.md) | 请求参数 | 

### Return type

[**InternalApiDtoExamplePostResponse**](InternalApiDtoExamplePostResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

