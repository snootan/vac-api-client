/*
Catalog Governor Service REST API

This is the service to track assets deployed in customer clusters

API version: 0.1.0
Contact: content-building-ecosystem@vmware.com
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi

import (
	"bytes"
	"context"
	"io/ioutil"
	"net/http"
	"net/url"
	"strings"
)


type WorkloadsApi interface {

	/*
	GetCluster Get cluster and workload telemetry

	Returns information about a cluster and its worklooads deployed for an organization

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@param orgId A string with UUID format as the identifier of an organization
	@param clusterId A string as the identifier of a cluster
	@return ApiGetClusterRequest
	*/
	GetCluster(ctx context.Context, orgId string, clusterId string) ApiGetClusterRequest

	// GetClusterExecute executes the request
	//  @return Cluster
	GetClusterExecute(r ApiGetClusterRequest) (*Cluster, *http.Response, error)

	/*
	GetClusters Get information about all clusters and their workloads

	Returns all clusters and their workloads for an organization

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@param orgId A string with UUID format as the identifier of an organization
	@return ApiGetClustersRequest
	*/
	GetClusters(ctx context.Context, orgId string) ApiGetClustersRequest

	// GetClustersExecute executes the request
	//  @return []Cluster
	GetClustersExecute(r ApiGetClustersRequest) ([]Cluster, *http.Response, error)

	/*
	SaveCluster Save/update cluster and workload telemetry

	Stores information about a cluster and its workloads deployed for an organization

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@param orgId A string with UUID format as the identifier of an organization
	@param clusterId A string as the identifier of a cluster
	@return ApiSaveClusterRequest
	*/
	SaveCluster(ctx context.Context, orgId string, clusterId string) ApiSaveClusterRequest

	// SaveClusterExecute executes the request
	SaveClusterExecute(r ApiSaveClusterRequest) (*http.Response, error)
}

// WorkloadsApiService WorkloadsApi service
type WorkloadsApiService service

type ApiGetClusterRequest struct {
	ctx context.Context
	ApiService WorkloadsApi
	orgId string
	clusterId string
}

func (r ApiGetClusterRequest) Execute() (*Cluster, *http.Response, error) {
	return r.ApiService.GetClusterExecute(r)
}

/*
GetCluster Get cluster and workload telemetry

Returns information about a cluster and its worklooads deployed for an organization

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param orgId A string with UUID format as the identifier of an organization
 @param clusterId A string as the identifier of a cluster
 @return ApiGetClusterRequest
*/
func (a *WorkloadsApiService) GetCluster(ctx context.Context, orgId string, clusterId string) ApiGetClusterRequest {
	return ApiGetClusterRequest{
		ApiService: a,
		ctx: ctx,
		orgId: orgId,
		clusterId: clusterId,
	}
}

// Execute executes the request
//  @return Cluster
func (a *WorkloadsApiService) GetClusterExecute(r ApiGetClusterRequest) (*Cluster, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodGet
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *Cluster
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "WorkloadsApiService.GetCluster")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/orgs/{org_id}/clusters/{cluster_id}"
	localVarPath = strings.Replace(localVarPath, "{"+"org_id"+"}", url.PathEscape(parameterToString(r.orgId, "")), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"cluster_id"+"}", url.PathEscape(parameterToString(r.clusterId, "")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{}

	// set Content-Type header
	localVarHTTPContentType := selectHeaderContentType(localVarHTTPContentTypes)
	if localVarHTTPContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHTTPContentType
	}

	// to determine the Accept header
	localVarHTTPHeaderAccepts := []string{"application/json", "application/problem+json"}

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

	localVarBody, err := ioutil.ReadAll(localVarHTTPResponse.Body)
	localVarHTTPResponse.Body.Close()
	localVarHTTPResponse.Body = ioutil.NopCloser(bytes.NewBuffer(localVarBody))
	if err != nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	if localVarHTTPResponse.StatusCode >= 300 {
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: localVarHTTPResponse.Status,
		}
		if localVarHTTPResponse.StatusCode == 400 {
			var v ConstraintsViolationsError
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarReturnValue, localVarHTTPResponse, newErr
			}
			newErr.model = v
			return localVarReturnValue, localVarHTTPResponse, newErr
		}
		if localVarHTTPResponse.StatusCode == 404 {
			var v Error
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarReturnValue, localVarHTTPResponse, newErr
			}
			newErr.model = v
			return localVarReturnValue, localVarHTTPResponse, newErr
		}
		if localVarHTTPResponse.StatusCode == 500 {
			var v Error
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarReturnValue, localVarHTTPResponse, newErr
			}
			newErr.model = v
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

type ApiGetClustersRequest struct {
	ctx context.Context
	ApiService WorkloadsApi
	orgId string
}

func (r ApiGetClustersRequest) Execute() ([]Cluster, *http.Response, error) {
	return r.ApiService.GetClustersExecute(r)
}

/*
GetClusters Get information about all clusters and their workloads

Returns all clusters and their workloads for an organization

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param orgId A string with UUID format as the identifier of an organization
 @return ApiGetClustersRequest
*/
func (a *WorkloadsApiService) GetClusters(ctx context.Context, orgId string) ApiGetClustersRequest {
	return ApiGetClustersRequest{
		ApiService: a,
		ctx: ctx,
		orgId: orgId,
	}
}

// Execute executes the request
//  @return []Cluster
func (a *WorkloadsApiService) GetClustersExecute(r ApiGetClustersRequest) ([]Cluster, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodGet
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  []Cluster
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "WorkloadsApiService.GetClusters")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/orgs/{org_id}/clusters"
	localVarPath = strings.Replace(localVarPath, "{"+"org_id"+"}", url.PathEscape(parameterToString(r.orgId, "")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{}

	// set Content-Type header
	localVarHTTPContentType := selectHeaderContentType(localVarHTTPContentTypes)
	if localVarHTTPContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHTTPContentType
	}

	// to determine the Accept header
	localVarHTTPHeaderAccepts := []string{"application/json", "application/problem+json"}

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

	localVarBody, err := ioutil.ReadAll(localVarHTTPResponse.Body)
	localVarHTTPResponse.Body.Close()
	localVarHTTPResponse.Body = ioutil.NopCloser(bytes.NewBuffer(localVarBody))
	if err != nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	if localVarHTTPResponse.StatusCode >= 300 {
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: localVarHTTPResponse.Status,
		}
		if localVarHTTPResponse.StatusCode == 400 {
			var v ConstraintsViolationsError
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarReturnValue, localVarHTTPResponse, newErr
			}
			newErr.model = v
			return localVarReturnValue, localVarHTTPResponse, newErr
		}
		if localVarHTTPResponse.StatusCode == 404 {
			var v Error
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarReturnValue, localVarHTTPResponse, newErr
			}
			newErr.model = v
			return localVarReturnValue, localVarHTTPResponse, newErr
		}
		if localVarHTTPResponse.StatusCode == 500 {
			var v Error
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarReturnValue, localVarHTTPResponse, newErr
			}
			newErr.model = v
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

type ApiSaveClusterRequest struct {
	ctx context.Context
	ApiService WorkloadsApi
	orgId string
	clusterId string
	cluster *Cluster
}

// Details of the cluster stored
func (r ApiSaveClusterRequest) Cluster(cluster Cluster) ApiSaveClusterRequest {
	r.cluster = &cluster
	return r
}

func (r ApiSaveClusterRequest) Execute() (*http.Response, error) {
	return r.ApiService.SaveClusterExecute(r)
}

/*
SaveCluster Save/update cluster and workload telemetry

Stores information about a cluster and its workloads deployed for an organization

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param orgId A string with UUID format as the identifier of an organization
 @param clusterId A string as the identifier of a cluster
 @return ApiSaveClusterRequest
*/
func (a *WorkloadsApiService) SaveCluster(ctx context.Context, orgId string, clusterId string) ApiSaveClusterRequest {
	return ApiSaveClusterRequest{
		ApiService: a,
		ctx: ctx,
		orgId: orgId,
		clusterId: clusterId,
	}
}

// Execute executes the request
func (a *WorkloadsApiService) SaveClusterExecute(r ApiSaveClusterRequest) (*http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodPut
		localVarPostBody     interface{}
		formFiles            []formFile
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "WorkloadsApiService.SaveCluster")
	if err != nil {
		return nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/orgs/{org_id}/clusters/{cluster_id}"
	localVarPath = strings.Replace(localVarPath, "{"+"org_id"+"}", url.PathEscape(parameterToString(r.orgId, "")), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"cluster_id"+"}", url.PathEscape(parameterToString(r.clusterId, "")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}
	if r.cluster == nil {
		return nil, reportError("cluster is required and must be specified")
	}

	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{"application/json"}

	// set Content-Type header
	localVarHTTPContentType := selectHeaderContentType(localVarHTTPContentTypes)
	if localVarHTTPContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHTTPContentType
	}

	// to determine the Accept header
	localVarHTTPHeaderAccepts := []string{"application/problem+json"}

	// set Accept header
	localVarHTTPHeaderAccept := selectHeaderAccept(localVarHTTPHeaderAccepts)
	if localVarHTTPHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHTTPHeaderAccept
	}
	// body params
	localVarPostBody = r.cluster
	req, err := a.client.prepareRequest(r.ctx, localVarPath, localVarHTTPMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, formFiles)
	if err != nil {
		return nil, err
	}

	localVarHTTPResponse, err := a.client.callAPI(req)
	if err != nil || localVarHTTPResponse == nil {
		return localVarHTTPResponse, err
	}

	localVarBody, err := ioutil.ReadAll(localVarHTTPResponse.Body)
	localVarHTTPResponse.Body.Close()
	localVarHTTPResponse.Body = ioutil.NopCloser(bytes.NewBuffer(localVarBody))
	if err != nil {
		return localVarHTTPResponse, err
	}

	if localVarHTTPResponse.StatusCode >= 300 {
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: localVarHTTPResponse.Status,
		}
		if localVarHTTPResponse.StatusCode == 400 {
			var v ConstraintsViolationsError
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarHTTPResponse, newErr
			}
			newErr.model = v
			return localVarHTTPResponse, newErr
		}
		if localVarHTTPResponse.StatusCode == 404 {
			var v Error
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarHTTPResponse, newErr
			}
			newErr.model = v
			return localVarHTTPResponse, newErr
		}
		if localVarHTTPResponse.StatusCode == 500 {
			var v Error
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarHTTPResponse, newErr
			}
			newErr.model = v
		}
		return localVarHTTPResponse, newErr
	}

	return localVarHTTPResponse, nil
}
