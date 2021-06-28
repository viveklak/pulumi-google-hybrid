// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package beta

import (
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Retrieves information about the specified reservation.
func LookupReservation(ctx *pulumi.Context, args *LookupReservationArgs, opts ...pulumi.InvokeOption) (*LookupReservationResult, error) {
	var rv LookupReservationResult
	err := ctx.Invoke("google-native:compute/beta:getReservation", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupReservationArgs struct {
	Project     string `pulumi:"project"`
	Reservation string `pulumi:"reservation"`
	Zone        string `pulumi:"zone"`
}

type LookupReservationResult struct {
	// Full or partial URL to a parent commitment. This field displays for reservations that are tied to a commitment.
	Commitment string `pulumi:"commitment"`
	// Creation timestamp in RFC3339 text format.
	CreationTimestamp string `pulumi:"creationTimestamp"`
	// An optional description of this resource. Provide this property when you create the resource.
	Description string `pulumi:"description"`
	// Type of the resource. Always compute#reservations for reservations.
	Kind string `pulumi:"kind"`
	// The name of the resource, provided by the client when initially creating the resource. The resource name must be 1-63 characters long, and comply with RFC1035. Specifically, the name must be 1-63 characters long and match the regular expression `[a-z]([-a-z0-9]*[a-z0-9])?` which means the first character must be a lowercase letter, and all following characters must be a dash, lowercase letter, or digit, except the last character, which cannot be a dash.
	Name string `pulumi:"name"`
	// Reserved for future use.
	SatisfiesPzs bool `pulumi:"satisfiesPzs"`
	// Server-defined fully-qualified URL for this resource.
	SelfLink string `pulumi:"selfLink"`
	// Reservation for instances with specific machine shapes.
	SpecificReservation AllocationSpecificSKUReservationResponse `pulumi:"specificReservation"`
	// Indicates whether the reservation can be consumed by VMs with affinity for "any" reservation. If the field is set, then only VMs that target the reservation by name can consume from this reservation.
	SpecificReservationRequired bool `pulumi:"specificReservationRequired"`
	// The status of the reservation.
	Status string `pulumi:"status"`
	// Zone in which the reservation resides. A zone must be provided if the reservation is created within a commitment.
	Zone string `pulumi:"zone"`
}
