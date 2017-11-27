// Copyright (c) 2016, 2017, Oracle and/or its affiliates. All rights reserved.
// Code generated. DO NOT EDIT.

package loadbalancer

import (
	"github.com/oci-go-sdk/common"
	"net/http"
)

// GetBackendSetHealthRequest wrapper for the GetBackendSetHealth operation
type GetBackendSetHealthRequest struct {

	// The [OCID]({{DOC_SERVER_URL}}/Content/General/Concepts/identifiers.htm) of the load balancer associated with the backend set health status to be retrieved.
	LoadBalancerID *string `mandatory:"true" contributesTo:"path" name:"loadBalancerId"`

	// The name of the backend set to retrieve the health status for.
	// Example: `My_backend_set`
	BackendSetName *string `mandatory:"true" contributesTo:"path" name:"backendSetName"`

	// The unique Oracle-assigned identifier for the request. If you need to contact Oracle about a
	// particular request, please provide the request ID.
	OpcRequestID *string `mandatory:"false" contributesTo:"header" name:"opc-request-id"`
}

func (request GetBackendSetHealthRequest) String() string {
	return common.PointerString(request)
}

// GetBackendSetHealthResponse wrapper for the GetBackendSetHealth operation
type GetBackendSetHealthResponse struct {

	// The underlying http response
	RawResponse *http.Response

	// The BackendSetHealth instance
	BackendSetHealth `presentIn:"body"`

	// Unique Oracle-assigned identifier for the request. If you need to contact Oracle about
	// a particular request, please provide the request ID.
	OpcRequestID *string `presentIn:"header" name:"opc-request-id"`
}

func (response GetBackendSetHealthResponse) String() string {
	return common.PointerString(response)
}
