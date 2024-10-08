/*
gin-http API

gin-http服务文档

API version: 1.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi

import (
	"bytes"
	"context"
	"io"
	"net/http"
	"net/url"
)


// AuthAPIService AuthAPI service
type AuthAPIService service

type ApiAuthV1TokenGenerateGetRequest struct {
	ctx context.Context
	ApiService *AuthAPIService
	userId *string
	userName *string
	env *string
}

// UserID 用户id
func (r ApiAuthV1TokenGenerateGetRequest) UserId(userId string) ApiAuthV1TokenGenerateGetRequest {
	r.userId = &userId
	return r
}

// UserName 用户名  example:张三
func (r ApiAuthV1TokenGenerateGetRequest) UserName(userName string) ApiAuthV1TokenGenerateGetRequest {
	r.userName = &userName
	return r
}

// 环境变量,默认线上; sandbox 沙箱环境, production 生产环境
func (r ApiAuthV1TokenGenerateGetRequest) Env(env string) ApiAuthV1TokenGenerateGetRequest {
	r.env = &env
	return r
}

func (r ApiAuthV1TokenGenerateGetRequest) Execute() (*InternalApiDtoAppJwtTokenSwgResponse, *http.Response, error) {
	return r.ApiService.AuthV1TokenGenerateGetExecute(r)
}

/*
AuthV1TokenGenerateGet jwt-token生成及校验

jwt-token生成及校验

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @return ApiAuthV1TokenGenerateGetRequest
*/
func (a *AuthAPIService) AuthV1TokenGenerateGet(ctx context.Context) ApiAuthV1TokenGenerateGetRequest {
	return ApiAuthV1TokenGenerateGetRequest{
		ApiService: a,
		ctx: ctx,
	}
}

// Execute executes the request
//  @return InternalApiDtoAppJwtTokenSwgResponse
func (a *AuthAPIService) AuthV1TokenGenerateGetExecute(r ApiAuthV1TokenGenerateGetRequest) (*InternalApiDtoAppJwtTokenSwgResponse, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodGet
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *InternalApiDtoAppJwtTokenSwgResponse
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "AuthAPIService.AuthV1TokenGenerateGet")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/auth/v1/token/generate"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}
	if r.userId == nil {
		return localVarReturnValue, nil, reportError("userId is required and must be specified")
	}
	if r.userName == nil {
		return localVarReturnValue, nil, reportError("userName is required and must be specified")
	}

	if r.env != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "env", r.env, "", "")
	}
	parameterAddToHeaderOrQuery(localVarQueryParams, "userId", r.userId, "", "")
	parameterAddToHeaderOrQuery(localVarQueryParams, "userName", r.userName, "", "")
	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{}

	// set Content-Type header
	localVarHTTPContentType := selectHeaderContentType(localVarHTTPContentTypes)
	if localVarHTTPContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHTTPContentType
	}

	// to determine the Accept header
	localVarHTTPHeaderAccepts := []string{"*/*"}

	// set Accept header
	localVarHTTPHeaderAccept := selectHeaderAccept(localVarHTTPHeaderAccepts)
	if localVarHTTPHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHTTPHeaderAccept
	}
	req, err := a.client.prepareRequest(r.ctx, localVarPath, localVarHTTPMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, formFiles)
	if err != nil {
		return localVarReturnValue, nil, err
	}

	localVarHTTPResponse, err := a.client.callAPI(req)
	if err != nil || localVarHTTPResponse == nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	localVarBody, err := io.ReadAll(localVarHTTPResponse.Body)
	localVarHTTPResponse.Body.Close()
	localVarHTTPResponse.Body = io.NopCloser(bytes.NewBuffer(localVarBody))
	if err != nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	if localVarHTTPResponse.StatusCode >= 300 {
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: localVarHTTPResponse.Status,
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	err = a.client.decode(&localVarReturnValue, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
	if err != nil {
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: err.Error(),
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	return localVarReturnValue, localVarHTTPResponse, nil
}
