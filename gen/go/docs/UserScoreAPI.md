# \UserScoreAPI

All URIs are relative to *http://127.0.0.1:3013/goodsCenterLogic*

Method | HTTP request | Description
------------- | ------------- | -------------
[**V1UserScoreAddPost**](UserScoreAPI.md#V1UserScoreAddPost) | **Post** /v1/userScore/add | 添加数据
[**V1UserScoreDelPost**](UserScoreAPI.md#V1UserScoreDelPost) | **Post** /v1/userScore/del | 删除数据
[**V1UserScoreFindGet**](UserScoreAPI.md#V1UserScoreFindGet) | **Get** /v1/userScore/find | 查询一条数据
[**V1UserScoreListGet**](UserScoreAPI.md#V1UserScoreListGet) | **Get** /v1/userScore/list | 查询列表
[**V1UserScoreUpdatePost**](UserScoreAPI.md#V1UserScoreUpdatePost) | **Post** /v1/userScore/update | 更新数据



## V1UserScoreAddPost

> InternalHandlerResponseData V1UserScoreAddPost(ctx).Execute()

添加数据



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

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.UserScoreAPI.V1UserScoreAddPost(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `UserScoreAPI.V1UserScoreAddPost``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `V1UserScoreAddPost`: InternalHandlerResponseData
	fmt.Fprintf(os.Stdout, "Response from `UserScoreAPI.V1UserScoreAddPost`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiV1UserScoreAddPostRequest struct via the builder pattern


### Return type

[**InternalHandlerResponseData**](InternalHandlerResponseData.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## V1UserScoreDelPost

> InternalHandlerResponseData V1UserScoreDelPost(ctx).Execute()

删除数据



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

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.UserScoreAPI.V1UserScoreDelPost(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `UserScoreAPI.V1UserScoreDelPost``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `V1UserScoreDelPost`: InternalHandlerResponseData
	fmt.Fprintf(os.Stdout, "Response from `UserScoreAPI.V1UserScoreDelPost`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiV1UserScoreDelPostRequest struct via the builder pattern


### Return type

[**InternalHandlerResponseData**](InternalHandlerResponseData.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## V1UserScoreFindGet

> InternalApiDtoUserScoreFindResponse V1UserScoreFindGet(ctx).Execute()

查询一条数据



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

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.UserScoreAPI.V1UserScoreFindGet(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `UserScoreAPI.V1UserScoreFindGet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `V1UserScoreFindGet`: InternalApiDtoUserScoreFindResponse
	fmt.Fprintf(os.Stdout, "Response from `UserScoreAPI.V1UserScoreFindGet`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiV1UserScoreFindGetRequest struct via the builder pattern


### Return type

[**InternalApiDtoUserScoreFindResponse**](InternalApiDtoUserScoreFindResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## V1UserScoreListGet

> InternalApiDtoUserScoreListResponse V1UserScoreListGet(ctx).Execute()

查询列表



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

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.UserScoreAPI.V1UserScoreListGet(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `UserScoreAPI.V1UserScoreListGet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `V1UserScoreListGet`: InternalApiDtoUserScoreListResponse
	fmt.Fprintf(os.Stdout, "Response from `UserScoreAPI.V1UserScoreListGet`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiV1UserScoreListGetRequest struct via the builder pattern


### Return type

[**InternalApiDtoUserScoreListResponse**](InternalApiDtoUserScoreListResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## V1UserScoreUpdatePost

> InternalHandlerResponseData V1UserScoreUpdatePost(ctx).Execute()

更新数据



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

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.UserScoreAPI.V1UserScoreUpdatePost(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `UserScoreAPI.V1UserScoreUpdatePost``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `V1UserScoreUpdatePost`: InternalHandlerResponseData
	fmt.Fprintf(os.Stdout, "Response from `UserScoreAPI.V1UserScoreUpdatePost`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiV1UserScoreUpdatePostRequest struct via the builder pattern


### Return type

[**InternalHandlerResponseData**](InternalHandlerResponseData.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

