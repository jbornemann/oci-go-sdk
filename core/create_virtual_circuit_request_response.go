// Copyright (c) 2016, 2017, Oracle and/or its affiliates. All rights reserved.
// Code generated. DO NOT EDIT.

package core

import (
	"github.com/oci-go-sdk/common"
	"net/http"
)

// CreateVirtualCircuitRequest wrapper for the CreateVirtualCircuit operation
type CreateVirtualCircuitRequest struct {

	// Details to create a VirtualCircuit.
	CreateVirtualCircuitDetails `contributesTo:"body"`

	// A token that uniquely identifies a request so it can be retried in case of a timeout or
	// server error without risk of executing that same action again. Retry tokens expire after 24
	// hours, but can be invalidated before then due to conflicting operations (for example, if a resource
	// has been deleted and purged from the system, then a retry of the original creation request
	// may be rejected).
	OpcRetryToken *string `mandatory:"false" contributesTo:"header" name:"opc-retry-token"`
}

func (request CreateVirtualCircuitRequest) String() string {
	return common.PointerString(request)
}

// CreateVirtualCircuitResponse wrapper for the CreateVirtualCircuit operation
type CreateVirtualCircuitResponse struct {

	// The underlying http response
	RawResponse *http.Response

	// The VirtualCircuit instance
	VirtualCircuit `presentIn:"body"`

	// For optimistic concurrency control. See `if-match`.
	Etag *string `presentIn:"header" name:"etag"`

	// Unique Oracle-assigned identifier for the request. If you need to contact Oracle about
	// a particular request, please provide the request ID.
	OpcRequestID *string `presentIn:"header" name:"opc-request-id"`
}

func (response CreateVirtualCircuitResponse) String() string {
	return common.PointerString(response)
}
