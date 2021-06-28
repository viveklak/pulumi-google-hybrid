// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package beta

import (
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Returns the specified NotificationEndpoint resource in the given region.
func LookupRegionNotificationEndpoint(ctx *pulumi.Context, args *LookupRegionNotificationEndpointArgs, opts ...pulumi.InvokeOption) (*LookupRegionNotificationEndpointResult, error) {
	var rv LookupRegionNotificationEndpointResult
	err := ctx.Invoke("google-native:compute/beta:getRegionNotificationEndpoint", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupRegionNotificationEndpointArgs struct {
	NotificationEndpoint string `pulumi:"notificationEndpoint"`
	Project              string `pulumi:"project"`
	Region               string `pulumi:"region"`
}

type LookupRegionNotificationEndpointResult struct {
	// Creation timestamp in RFC3339 text format.
	CreationTimestamp string `pulumi:"creationTimestamp"`
	// An optional description of this resource. Provide this property when you create the resource.
	Description string `pulumi:"description"`
	// Settings of the gRPC notification endpoint including the endpoint URL and the retry duration.
	GrpcSettings NotificationEndpointGrpcSettingsResponse `pulumi:"grpcSettings"`
	// Type of the resource. Always compute#notificationEndpoint for notification endpoints.
	Kind string `pulumi:"kind"`
	// Name of the resource. Provided by the client when the resource is created. The name must be 1-63 characters long, and comply with RFC1035. Specifically, the name must be 1-63 characters long and match the regular expression `[a-z]([-a-z0-9]*[a-z0-9])?` which means the first character must be a lowercase letter, and all following characters must be a dash, lowercase letter, or digit, except the last character, which cannot be a dash.
	Name string `pulumi:"name"`
	// URL of the region where the notification endpoint resides. This field applies only to the regional resource. You must specify this field as part of the HTTP request URL. It is not settable as a field in the request body.
	Region string `pulumi:"region"`
	// Server-defined URL for the resource.
	SelfLink string `pulumi:"selfLink"`
}
