// Copyright (c) 2016, 2017, Oracle and/or its affiliates. All rights reserved.
// Code generated. DO NOT EDIT.

// Identity and Access Management Service API
//
// APIs for managing users, groups, compartments, and policies.
//

package identity

import (
	"github.com/oci-go-sdk/common"
	"encoding/json"
)

// IdentityProvider. The resulting base object when you add an identity provider to your tenancy. A
// Saml2IdentityProvider
// is a specific type of `IdentityProvider` that supports the SAML 2.0 protocol. Each
// `IdentityProvider` object has its own OCID. For more information, see
// [Identity Providers and Federation]({{DOC_SERVER_URL}}/Content/Identity/Concepts/federation.htm).
// To use any of the API operations, you must be authorized in an IAM policy. If you're not authorized,
// talk to an administrator. If you're an administrator who needs to write policies to give users access,
// see [Getting Started with Policies]({{DOC_SERVER_URL}}/Content/Identity/Concepts/policygetstarted.htm).
type IdentityProvider interface {

	// The OCID of the `IdentityProvider`.
	GetID() *string

	// The OCID of the tenancy containing the `IdentityProvider`.
	GetCompartmentID() *string

	// The name you assign to the `IdentityProvider` during creation. The name
	// must be unique across all `IdentityProvider` objects in the tenancy and
	// cannot be changed. This is the name federated users see when choosing
	// which identity provider to use when signing in to the Oracle Cloud Infrastructure
	// Console.
	GetName() *string

	// The description you assign to the `IdentityProvider` during creation. Does
	// not have to be unique, and it's changeable.
	GetDescription() *string

	// The identity provider service or product.
	// Supported identity providers are Oracle Identity Cloud Service (IDCS) and Microsoft
	// Active Directory Federation Services (ADFS).
	// Allowed values are:
	// - `ADFS`
	// - `IDCS`
	// Example: `IDCS`
	GetProductType() *string

	// Date and time the `IdentityProvider` was created, in the format defined by RFC3339.
	// Example: `2016-08-25T21:10:29.600Z`
	GetTimeCreated() *common.SDKTime

	// The current state. After creating an `IdentityProvider`, make sure its
	// `lifecycleState` changes from CREATING to ACTIVE before using it.
	GetLifecycleState() IdentityProviderLifecycleStateEnum

	// The detailed status of INACTIVE lifecycleState.
	GetInactiveStatus() *int
}

type identityprovider struct {
	ID             *string                            `mandatory:"true" json:"id,omitempty"`
	CompartmentID  *string                            `mandatory:"true" json:"compartmentId,omitempty"`
	Name           *string                            `mandatory:"true" json:"name,omitempty"`
	Description    *string                            `mandatory:"true" json:"description,omitempty"`
	ProductType    *string                            `mandatory:"true" json:"productType,omitempty"`
	TimeCreated    *common.SDKTime                    `mandatory:"true" json:"timeCreated,omitempty"`
	LifecycleState IdentityProviderLifecycleStateEnum `mandatory:"true" json:"lifecycleState,omitempty"`
	InactiveStatus *int                               `mandatory:"false" json:"inactiveStatus,omitempty"`
	Protocol       string                             `json:"protocol"`
}

func (m *identityprovider) UnmarshalPolymorphicJSON(data []byte) (interface{}, error) {
	err := json.Unmarshal(data, m)
	if err != nil {
		return nil, err
	}

	switch m.Protocol {
	case "SAML2":
		mm := Saml2IdentityProvider{}
		err = json.Unmarshal(data, &mm)
		return mm, err
	default:
		return m, nil
	}
}

func (m identityprovider) GetID() *string {
	return m.ID
}
func (m identityprovider) GetCompartmentID() *string {
	return m.CompartmentID
}
func (m identityprovider) GetName() *string {
	return m.Name
}
func (m identityprovider) GetDescription() *string {
	return m.Description
}
func (m identityprovider) GetProductType() *string {
	return m.ProductType
}
func (m identityprovider) GetTimeCreated() *common.SDKTime {
	return m.TimeCreated
}
func (m identityprovider) GetLifecycleState() IdentityProviderLifecycleStateEnum {
	return m.LifecycleState
}
func (m identityprovider) GetInactiveStatus() *int {
	return m.InactiveStatus
}

func (model identityprovider) String() string {
	return common.PointerString(model)
}

type IdentityProviderLifecycleStateEnum string

const (
	IDENTITY_PROVIDER_LIFECYCLE_STATE_CREATING IdentityProviderLifecycleStateEnum = "CREATING"
	IDENTITY_PROVIDER_LIFECYCLE_STATE_ACTIVE   IdentityProviderLifecycleStateEnum = "ACTIVE"
	IDENTITY_PROVIDER_LIFECYCLE_STATE_INACTIVE IdentityProviderLifecycleStateEnum = "INACTIVE"
	IDENTITY_PROVIDER_LIFECYCLE_STATE_DELETING IdentityProviderLifecycleStateEnum = "DELETING"
	IDENTITY_PROVIDER_LIFECYCLE_STATE_DELETED  IdentityProviderLifecycleStateEnum = "DELETED"
	IDENTITY_PROVIDER_LIFECYCLE_STATE_UNKNOWN  IdentityProviderLifecycleStateEnum = "UNKNOWN"
)

var mapping_identityprovider_lifecycleState = map[string]IdentityProviderLifecycleStateEnum{
	"CREATING": IDENTITY_PROVIDER_LIFECYCLE_STATE_CREATING,
	"ACTIVE":   IDENTITY_PROVIDER_LIFECYCLE_STATE_ACTIVE,
	"INACTIVE": IDENTITY_PROVIDER_LIFECYCLE_STATE_INACTIVE,
	"DELETING": IDENTITY_PROVIDER_LIFECYCLE_STATE_DELETING,
	"DELETED":  IDENTITY_PROVIDER_LIFECYCLE_STATE_DELETED,
	"UNKNOWN":  IDENTITY_PROVIDER_LIFECYCLE_STATE_UNKNOWN,
}

func GetIdentityProviderLifecycleStateEnumValues() []IdentityProviderLifecycleStateEnum {
	values := make([]IdentityProviderLifecycleStateEnum, 0)
	for _, v := range mapping_identityprovider_lifecycleState {
		if v != IDENTITY_PROVIDER_LIFECYCLE_STATE_UNKNOWN {
			values = append(values, v)
		}
	}
	return values
}
