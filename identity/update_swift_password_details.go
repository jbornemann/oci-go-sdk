// Copyright (c) 2016, 2017, Oracle and/or its affiliates. All rights reserved.
// Code generated. DO NOT EDIT.

// Identity and Access Management Service API
//
// APIs for managing users, groups, compartments, and policies.
//

package identity

import (
	"github.com/oci-go-sdk/common"
)

type UpdateSwiftPasswordDetails struct {

	// The description you assign to the Swift password. Does not have to be unique, and it's changeable.
	Description *string `mandatory:"false" json:"description,omitempty"`
}

func (model UpdateSwiftPasswordDetails) String() string {
	return common.PointerString(model)
}
