// Copyright (c) 2016, 2017, Oracle and/or its affiliates. All rights reserved.
// Code generated. DO NOT EDIT.

package core

import (
	"github.com/oci-go-sdk/common"
	"net/http"
)

// GetVcnRequest wrapper for the GetVcn operation
type GetVcnRequest struct {

	// The OCID of the VCN.
	VcnID *string `mandatory:"true" contributesTo:"path" name:"vcnId"`
}

func (request GetVcnRequest) String() string {
	return common.PointerString(request)
}

// GetVcnResponse wrapper for the GetVcn operation
type GetVcnResponse struct {

	// The underlying http response
	RawResponse *http.Response

	// The Vcn instance
	Vcn `presentIn:"body"`

	// For optimistic concurrency control. See `if-match`.
	Etag *string `presentIn:"header" name:"etag"`

	// Unique Oracle-assigned identifier for the request. If you need to contact Oracle about
	// a particular request, please provide the request ID.
	OpcRequestID *string `presentIn:"header" name:"opc-request-id"`
}

func (response GetVcnResponse) String() string {
	return common.PointerString(response)
}
