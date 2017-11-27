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

// BackendSetHealth. The health status details for a backend set.
// This object does not explicitly enumerate backend servers with a status of `OK`. However, they are included in the
// `totalBackendCount` sum.
type BackendSetHealth struct {

	// A list of backend servers that are currently in the `CRITICAL` health state. The list identifies each backend server by
	// IP address and port.
	// Example: `1.1.1.1:80`
	CriticalStateBackendNames *[]string `mandatory:"true" json:"criticalStateBackendNames,omitempty"`

	// Overall health status of the backend set.
	// *  **OK:** All backend servers in the backend set return a status of `OK`.
	// *  **WARNING:** Half or more of the backend set's backend servers return a status of `OK` and at least one backend
	// server returns a status of `WARNING`, `CRITICAL`, or `UNKNOWN`.
	// *  **CRITICAL:** Fewer than half of the backend set's backend servers return a status of `OK`.
	// *  **UNKNOWN:** More than half of the backend set's backend servers return a status of `UNKNOWN`, the system was
	// unable to retrieve metrics, or the backend set does not have a listener attached.
	Status BackendSetHealthStatusEnum `mandatory:"true" json:"status,omitempty"`

	// The total number of backend servers in this backend set.
	// Example: `5`
	TotalBackendCount *int `mandatory:"true" json:"totalBackendCount,omitempty"`

	// A list of backend servers that are currently in the `UNKNOWN` health state. The list identifies each backend server by
	// IP address and port.
	// Example: `1.1.1.5:80`
	UnknownStateBackendNames *[]string `mandatory:"true" json:"unknownStateBackendNames,omitempty"`

	// A list of backend servers that are currently in the `WARNING` health state. The list identifies each backend server by
	// IP address and port.
	// Example: `1.1.1.7:42`
	WarningStateBackendNames *[]string `mandatory:"true" json:"warningStateBackendNames,omitempty"`
}

func (model BackendSetHealth) String() string {
	return common.PointerString(model)
}

type BackendSetHealthStatusEnum string

const (
	BACKEND_SET_HEALTH_STATUS_OK       BackendSetHealthStatusEnum = "OK"
	BACKEND_SET_HEALTH_STATUS_WARNING  BackendSetHealthStatusEnum = "WARNING"
	BACKEND_SET_HEALTH_STATUS_CRITICAL BackendSetHealthStatusEnum = "CRITICAL"
	BACKEND_SET_HEALTH_STATUS_UNKNOWN  BackendSetHealthStatusEnum = "UNKNOWN"
)

var mapping_backendsethealth_status = map[string]BackendSetHealthStatusEnum{
	"OK":       BACKEND_SET_HEALTH_STATUS_OK,
	"WARNING":  BACKEND_SET_HEALTH_STATUS_WARNING,
	"CRITICAL": BACKEND_SET_HEALTH_STATUS_CRITICAL,
	"UNKNOWN":  BACKEND_SET_HEALTH_STATUS_UNKNOWN,
}

func GetBackendSetHealthStatusEnumValues() []BackendSetHealthStatusEnum {
	values := make([]BackendSetHealthStatusEnum, 0)
	for _, v := range mapping_backendsethealth_status {
		values = append(values, v)
	}
	return values
}
