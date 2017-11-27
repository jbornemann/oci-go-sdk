// Copyright (c) 2016, 2017, Oracle and/or its affiliates. All rights reserved.
// Code generated. DO NOT EDIT.

// Load Balancing Service API
//
// API for the Load Balancing Service
//

package loadbalancer

import (
	"github.com/oci-go-sdk/common"
)

// UpdateHealthCheckerDetails. The health checker's configuration details.
type UpdateHealthCheckerDetails struct {

	// The interval between health checks, in milliseconds.
	// Example: `30000`
	IntervalInMillis *int `mandatory:"true" json:"intervalInMillis,omitempty"`

	// The backend server port against which to run the health check.
	// Example: `8080`
	Port *int `mandatory:"true" json:"port,omitempty"`

	// The protocol the health check must use; either HTTP or TCP.
	// Example: `HTTP`
	Protocol *string `mandatory:"true" json:"protocol,omitempty"`

	// A regular expression for parsing the response body from the backend server.
	// Example: `^(500|40[1348])$`
	ResponseBodyRegex *string `mandatory:"true" json:"responseBodyRegex,omitempty"`

	// The number of retries to attempt before a backend server is considered "unhealthy".
	// Example: `3`
	Retries *int `mandatory:"true" json:"retries,omitempty"`

	// The status code a healthy backend server should return.
	// Example: `200`
	ReturnCode *int `mandatory:"true" json:"returnCode,omitempty"`

	// The maximum time, in milliseconds, to wait for a reply to a health check. A health check is successful only if a reply
	// returns within this timeout period.
	// Example: `6000`
	TimeoutInMillis *int `mandatory:"true" json:"timeoutInMillis,omitempty"`

	// The path against which to run the health check.
	// Example: `/healthcheck`
	UrlPath *string `mandatory:"false" json:"urlPath,omitempty"`
}

func (model UpdateHealthCheckerDetails) String() string {
	return common.PointerString(model)
}
