// Copyright (c) 2016, 2017, Oracle and/or its affiliates. All rights reserved.
// Code generated. DO NOT EDIT.

// Core Services API
//
// APIs for Networking Service, Compute Service, and Block Volume Service.
//

package core

import (
	"github.com/oci-go-sdk/common"
)

// Subnet. A logical subdivision of a VCN. Each subnet exists in a single Availability Domain and
// consists of a contiguous range of IP addresses that do not overlap with
// other subnets in the VCN. Example: 172.16.1.0/24. For more information, see
// [Overview of the Networking Service]({{DOC_SERVER_URL}}/Content/Network/Concepts/overview.htm) and
// [VCNs and Subnets]({{DOC_SERVER_URL}}/Content/Network/Tasks/managingVCNs.htm).
// To use any of the API operations, you must be authorized in an IAM policy. If you're not authorized,
// talk to an administrator. If you're an administrator who needs to write policies to give users access, see
// [Getting Started with Policies]({{DOC_SERVER_URL}}/Content/Identity/Concepts/policygetstarted.htm).
type Subnet struct {

	// The subnet's Availability Domain.
	// Example: `Uocm:PHX-AD-1`
	AvailabilityDomain *string `mandatory:"true" json:"availabilityDomain,omitempty"`

	// The subnet's CIDR block.
	// Example: `172.16.1.0/24`
	CidrBlock *string `mandatory:"true" json:"cidrBlock,omitempty"`

	// The OCID of the compartment containing the subnet.
	CompartmentID *string `mandatory:"true" json:"compartmentId,omitempty"`

	// The subnet's Oracle ID (OCID).
	ID *string `mandatory:"true" json:"id,omitempty"`

	// The subnet's current state.
	LifecycleState SubnetLifecycleStateEnum `mandatory:"true" json:"lifecycleState,omitempty"`

	// The OCID of the route table the subnet is using.
	RouteTableID *string `mandatory:"true" json:"routeTableId,omitempty"`

	// The OCID of the VCN the subnet is in.
	VcnID *string `mandatory:"true" json:"vcnId,omitempty"`

	// The IP address of the virtual router.
	// Example: `10.0.14.1`
	VirtualRouterIp *string `mandatory:"true" json:"virtualRouterIp,omitempty"`

	// The MAC address of the virtual router.
	// Example: `00:00:17:B6:4D:DD`
	VirtualRouterMac *string `mandatory:"true" json:"virtualRouterMac,omitempty"`

	// The OCID of the set of DHCP options associated with the subnet.
	DhcpOptionsID *string `mandatory:"false" json:"dhcpOptionsId,omitempty"`

	// A user-friendly name. Does not have to be unique, and it's changeable.
	// Avoid entering confidential information.
	DisplayName *string `mandatory:"false" json:"displayName,omitempty"`

	// A DNS label for the subnet, used in conjunction with the VNIC's hostname and
	// VCN's DNS label to form a fully qualified domain name (FQDN) for each VNIC
	// within this subnet (for example, `bminstance-1.subnet123.vcn1.oraclevcn.com`).
	// Must be an alphanumeric string that begins with a letter and is unique within the VCN.
	// The value cannot be changed.
	// The absence of this parameter means the Internet and VCN Resolver
	// will not resolve hostnames of instances in this subnet.
	// For more information, see
	// [DNS in Your Virtual Cloud Network]({{DOC_SERVER_URL}}/Content/Network/Concepts/dns.htm).
	// Example: `subnet123`
	DnsLabel *string `mandatory:"false" json:"dnsLabel,omitempty"`

	// Whether VNICs within this subnet can have public IP addresses.
	// Defaults to false, which means VNICs created in this subnet will
	// automatically be assigned public IP addresses unless specified
	// otherwise during instance launch or VNIC creation (with the
	// `assignPublicIp` flag in
	// CreateVnicDetails).
	// If `prohibitPublicIpOnVnic` is set to true, VNICs created in this
	// subnet cannot have public IP addresses (that is, it's a private
	// subnet).
	// Example: `true`
	ProhibitPublicIpOnVnic *bool `mandatory:"false" json:"prohibitPublicIpOnVnic,omitempty"`

	// OCIDs for the security lists to use for VNICs in this subnet.
	SecurityListIds *[]string `mandatory:"false" json:"securityListIds,omitempty"`

	// The subnet's domain name, which consists of the subnet's DNS label,
	// the VCN's DNS label, and the `oraclevcn.com` domain.
	// For more information, see
	// [DNS in Your Virtual Cloud Network]({{DOC_SERVER_URL}}/Content/Network/Concepts/dns.htm).
	// Example: `subnet123.vcn1.oraclevcn.com`
	SubnetDomainName *string `mandatory:"false" json:"subnetDomainName,omitempty"`

	// The date and time the subnet was created, in the format defined by RFC3339.
	// Example: `2016-08-25T21:10:29.600Z`
	TimeCreated *common.SDKTime `mandatory:"false" json:"timeCreated,omitempty"`
}

func (model Subnet) String() string {
	return common.PointerString(model)
}

type SubnetLifecycleStateEnum string

const (
	SUBNET_LIFECYCLE_STATE_PROVISIONING SubnetLifecycleStateEnum = "PROVISIONING"
	SUBNET_LIFECYCLE_STATE_AVAILABLE    SubnetLifecycleStateEnum = "AVAILABLE"
	SUBNET_LIFECYCLE_STATE_TERMINATING  SubnetLifecycleStateEnum = "TERMINATING"
	SUBNET_LIFECYCLE_STATE_TERMINATED   SubnetLifecycleStateEnum = "TERMINATED"
	SUBNET_LIFECYCLE_STATE_UNKNOWN      SubnetLifecycleStateEnum = "UNKNOWN"
)

var mapping_subnet_lifecycleState = map[string]SubnetLifecycleStateEnum{
	"PROVISIONING": SUBNET_LIFECYCLE_STATE_PROVISIONING,
	"AVAILABLE":    SUBNET_LIFECYCLE_STATE_AVAILABLE,
	"TERMINATING":  SUBNET_LIFECYCLE_STATE_TERMINATING,
	"TERMINATED":   SUBNET_LIFECYCLE_STATE_TERMINATED,
	"UNKNOWN":      SUBNET_LIFECYCLE_STATE_UNKNOWN,
}

func GetSubnetLifecycleStateEnumValues() []SubnetLifecycleStateEnum {
	values := make([]SubnetLifecycleStateEnum, 0)
	for _, v := range mapping_subnet_lifecycleState {
		if v != SUBNET_LIFECYCLE_STATE_UNKNOWN {
			values = append(values, v)
		}
	}
	return values
}
