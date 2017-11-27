// Copyright (c) 2016, 2017, Oracle and/or its affiliates. All rights reserved.
// Code generated. DO NOT EDIT.

package core

import (
	"github.com/oci-go-sdk/common"
	"net/http"
)

// GetVnicAttachmentRequest wrapper for the GetVnicAttachment operation
type GetVnicAttachmentRequest struct {

	// The OCID of the VNIC attachment.
	VnicAttachmentID *string `mandatory:"true" contributesTo:"path" name:"vnicAttachmentId"`
}

func (request GetVnicAttachmentRequest) String() string {
	return common.PointerString(request)
}

// GetVnicAttachmentResponse wrapper for the GetVnicAttachment operation
type GetVnicAttachmentResponse struct {

	// The underlying http response
	RawResponse *http.Response

	// The VnicAttachment instance
	VnicAttachment `presentIn:"body"`

	// For optimistic concurrency control. See `if-match`.
	Etag *string `presentIn:"header" name:"etag"`

	// Unique Oracle-assigned identifier for the request. If you need to contact Oracle about
	// a particular request, please provide the request ID.
	OpcRequestID *string `presentIn:"header" name:"opc-request-id"`
}

func (response GetVnicAttachmentResponse) String() string {
	return common.PointerString(response)
}
