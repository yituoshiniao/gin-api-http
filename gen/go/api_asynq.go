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


// AsynqAPIService AsynqAPI service
type AsynqAPIService service

type ApiAsynqV1AddAggTaskGetRequest struct {
	ctx context.Context
	ApiService *AsynqAPIService
	env *string
}

// 环境变量,默认线上; sandbox 沙箱环境, production 生产环境
func (r ApiAsynqV1AddAggTaskGetRequest) Env(env string) ApiAsynqV1AddAggTaskGetRequest {
	r.env = &env
	return r
}

func (r ApiAsynqV1AddAggTaskGetRequest) Execute() (*InternalApiDtoGroupDeliveryTaskAddResponse, *http.Response, error) {
	return r.ApiService.AsynqV1AddAggTaskGetExecute(r)
}

/*
AsynqV1AddAggTaskGet asynq-添加聚合任务

asynq-添加聚合任务

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @return ApiAsynqV1AddAggTaskGetRequest
*/
func (a *AsynqAPIService) AsynqV1AddAggTaskGet(ctx context.Context) ApiAsynqV1AddAggTaskGetRequest {
	return ApiAsynqV1AddAggTaskGetRequest{
		ApiService: a,
		ctx: ctx,
	}
}

// Execute executes the request
//	@return	InternalApiDtoGroupDeliveryTaskAddResponse
func (a *AsynqAPIService) AsynqV1AddAggTaskGetExecute(r ApiAsynqV1AddAggTaskGetRequest) (*InternalApiDtoGroupDeliveryTaskAddResponse, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodGet
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *InternalApiDtoGroupDeliveryTaskAddResponse
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "AsynqAPIService.AsynqV1AddAggTaskGet")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/asynq/v1/addAggTask"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	if r.env != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "env", r.env, "")
	}
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

type ApiAsynqV1AddTaskGetRequest struct {
	ctx context.Context
	ApiService *AsynqAPIService
	env *string
}

// 环境变量,默认线上; sandbox 沙箱环境, production 生产环境
func (r ApiAsynqV1AddTaskGetRequest) Env(env string) ApiAsynqV1AddTaskGetRequest {
	r.env = &env
	return r
}

func (r ApiAsynqV1AddTaskGetRequest) Execute() (*InternalApiDtoAsynqEmailDeliveryTaskAddResponse, *http.Response, error) {
	return r.ApiService.AsynqV1AddTaskGetExecute(r)
}

/*
AsynqV1AddTaskGet asynq-add异步任务

asynq-异步任务,可通过： http://localhost:7013/monitoring/ 查看dashbord报表

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @return ApiAsynqV1AddTaskGetRequest
*/
func (a *AsynqAPIService) AsynqV1AddTaskGet(ctx context.Context) ApiAsynqV1AddTaskGetRequest {
	return ApiAsynqV1AddTaskGetRequest{
		ApiService: a,
		ctx: ctx,
	}
}

// Execute executes the request
//	@return	InternalApiDtoAsynqEmailDeliveryTaskAddResponse
func (a *AsynqAPIService) AsynqV1AddTaskGetExecute(r ApiAsynqV1AddTaskGetRequest) (*InternalApiDtoAsynqEmailDeliveryTaskAddResponse, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodGet
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *InternalApiDtoAsynqEmailDeliveryTaskAddResponse
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "AsynqAPIService.AsynqV1AddTaskGet")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/asynq/v1/addTask"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	if r.env != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "env", r.env, "")
	}
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
