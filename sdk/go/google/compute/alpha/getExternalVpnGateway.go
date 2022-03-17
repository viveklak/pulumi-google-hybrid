// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package alpha

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Returns the specified externalVpnGateway. Get a list of available externalVpnGateways by making a list() request.
func LookupExternalVpnGateway(ctx *pulumi.Context, args *LookupExternalVpnGatewayArgs, opts ...pulumi.InvokeOption) (*LookupExternalVpnGatewayResult, error) {
	var rv LookupExternalVpnGatewayResult
	err := ctx.Invoke("google-native:compute/alpha:getExternalVpnGateway", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupExternalVpnGatewayArgs struct {
	ExternalVpnGateway string  `pulumi:"externalVpnGateway"`
	Project            *string `pulumi:"project"`
}

type LookupExternalVpnGatewayResult struct {
	// Creation timestamp in RFC3339 text format.
	CreationTimestamp string `pulumi:"creationTimestamp"`
	// An optional description of this resource. Provide this property when you create the resource.
	Description string `pulumi:"description"`
	// A list of interfaces for this external VPN gateway. If your peer-side gateway is an on-premises gateway and non-AWS cloud providers' gateway, at most two interfaces can be provided for an external VPN gateway. If your peer side is an AWS virtual private gateway, four interfaces should be provided for an external VPN gateway.
	Interfaces []ExternalVpnGatewayInterfaceResponse `pulumi:"interfaces"`
	// Type of the resource. Always compute#externalVpnGateway for externalVpnGateways.
	Kind string `pulumi:"kind"`
	// A fingerprint for the labels being applied to this ExternalVpnGateway, which is essentially a hash of the labels set used for optimistic locking. The fingerprint is initially generated by Compute Engine and changes after every request to modify or update labels. You must always provide an up-to-date fingerprint hash in order to update or change labels, otherwise the request will fail with error 412 conditionNotMet. To see the latest fingerprint, make a get() request to retrieve an ExternalVpnGateway.
	LabelFingerprint string `pulumi:"labelFingerprint"`
	// Labels for this resource. These can only be added or modified by the setLabels method. Each label key/value pair must comply with RFC1035. Label values may be empty.
	Labels map[string]string `pulumi:"labels"`
	// Name of the resource. Provided by the client when the resource is created. The name must be 1-63 characters long, and comply with RFC1035. Specifically, the name must be 1-63 characters long and match the regular expression `[a-z]([-a-z0-9]*[a-z0-9])?` which means the first character must be a lowercase letter, and all following characters must be a dash, lowercase letter, or digit, except the last character, which cannot be a dash.
	Name string `pulumi:"name"`
	// Indicates the user-supplied redundancy type of this external VPN gateway.
	RedundancyType string `pulumi:"redundancyType"`
	// Server-defined URL for the resource.
	SelfLink string `pulumi:"selfLink"`
}

func LookupExternalVpnGatewayOutput(ctx *pulumi.Context, args LookupExternalVpnGatewayOutputArgs, opts ...pulumi.InvokeOption) LookupExternalVpnGatewayResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupExternalVpnGatewayResult, error) {
			args := v.(LookupExternalVpnGatewayArgs)
			r, err := LookupExternalVpnGateway(ctx, &args, opts...)
			return *r, err
		}).(LookupExternalVpnGatewayResultOutput)
}

type LookupExternalVpnGatewayOutputArgs struct {
	ExternalVpnGateway pulumi.StringInput    `pulumi:"externalVpnGateway"`
	Project            pulumi.StringPtrInput `pulumi:"project"`
}

func (LookupExternalVpnGatewayOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupExternalVpnGatewayArgs)(nil)).Elem()
}

type LookupExternalVpnGatewayResultOutput struct{ *pulumi.OutputState }

func (LookupExternalVpnGatewayResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupExternalVpnGatewayResult)(nil)).Elem()
}

func (o LookupExternalVpnGatewayResultOutput) ToLookupExternalVpnGatewayResultOutput() LookupExternalVpnGatewayResultOutput {
	return o
}

func (o LookupExternalVpnGatewayResultOutput) ToLookupExternalVpnGatewayResultOutputWithContext(ctx context.Context) LookupExternalVpnGatewayResultOutput {
	return o
}

// Creation timestamp in RFC3339 text format.
func (o LookupExternalVpnGatewayResultOutput) CreationTimestamp() pulumi.StringOutput {
	return o.ApplyT(func(v LookupExternalVpnGatewayResult) string { return v.CreationTimestamp }).(pulumi.StringOutput)
}

// An optional description of this resource. Provide this property when you create the resource.
func (o LookupExternalVpnGatewayResultOutput) Description() pulumi.StringOutput {
	return o.ApplyT(func(v LookupExternalVpnGatewayResult) string { return v.Description }).(pulumi.StringOutput)
}

// A list of interfaces for this external VPN gateway. If your peer-side gateway is an on-premises gateway and non-AWS cloud providers' gateway, at most two interfaces can be provided for an external VPN gateway. If your peer side is an AWS virtual private gateway, four interfaces should be provided for an external VPN gateway.
func (o LookupExternalVpnGatewayResultOutput) Interfaces() ExternalVpnGatewayInterfaceResponseArrayOutput {
	return o.ApplyT(func(v LookupExternalVpnGatewayResult) []ExternalVpnGatewayInterfaceResponse { return v.Interfaces }).(ExternalVpnGatewayInterfaceResponseArrayOutput)
}

// Type of the resource. Always compute#externalVpnGateway for externalVpnGateways.
func (o LookupExternalVpnGatewayResultOutput) Kind() pulumi.StringOutput {
	return o.ApplyT(func(v LookupExternalVpnGatewayResult) string { return v.Kind }).(pulumi.StringOutput)
}

// A fingerprint for the labels being applied to this ExternalVpnGateway, which is essentially a hash of the labels set used for optimistic locking. The fingerprint is initially generated by Compute Engine and changes after every request to modify or update labels. You must always provide an up-to-date fingerprint hash in order to update or change labels, otherwise the request will fail with error 412 conditionNotMet. To see the latest fingerprint, make a get() request to retrieve an ExternalVpnGateway.
func (o LookupExternalVpnGatewayResultOutput) LabelFingerprint() pulumi.StringOutput {
	return o.ApplyT(func(v LookupExternalVpnGatewayResult) string { return v.LabelFingerprint }).(pulumi.StringOutput)
}

// Labels for this resource. These can only be added or modified by the setLabels method. Each label key/value pair must comply with RFC1035. Label values may be empty.
func (o LookupExternalVpnGatewayResultOutput) Labels() pulumi.StringMapOutput {
	return o.ApplyT(func(v LookupExternalVpnGatewayResult) map[string]string { return v.Labels }).(pulumi.StringMapOutput)
}

// Name of the resource. Provided by the client when the resource is created. The name must be 1-63 characters long, and comply with RFC1035. Specifically, the name must be 1-63 characters long and match the regular expression `[a-z]([-a-z0-9]*[a-z0-9])?` which means the first character must be a lowercase letter, and all following characters must be a dash, lowercase letter, or digit, except the last character, which cannot be a dash.
func (o LookupExternalVpnGatewayResultOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v LookupExternalVpnGatewayResult) string { return v.Name }).(pulumi.StringOutput)
}

// Indicates the user-supplied redundancy type of this external VPN gateway.
func (o LookupExternalVpnGatewayResultOutput) RedundancyType() pulumi.StringOutput {
	return o.ApplyT(func(v LookupExternalVpnGatewayResult) string { return v.RedundancyType }).(pulumi.StringOutput)
}

// Server-defined URL for the resource.
func (o LookupExternalVpnGatewayResultOutput) SelfLink() pulumi.StringOutput {
	return o.ApplyT(func(v LookupExternalVpnGatewayResult) string { return v.SelfLink }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupExternalVpnGatewayResultOutput{})
}
