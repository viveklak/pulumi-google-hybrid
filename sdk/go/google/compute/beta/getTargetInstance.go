// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package beta

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Returns the specified TargetInstance resource. Gets a list of available target instances by making a list() request.
func LookupTargetInstance(ctx *pulumi.Context, args *LookupTargetInstanceArgs, opts ...pulumi.InvokeOption) (*LookupTargetInstanceResult, error) {
	var rv LookupTargetInstanceResult
	err := ctx.Invoke("google-native:compute/beta:getTargetInstance", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupTargetInstanceArgs struct {
	Project        *string `pulumi:"project"`
	TargetInstance string  `pulumi:"targetInstance"`
	Zone           string  `pulumi:"zone"`
}

type LookupTargetInstanceResult struct {
	// Creation timestamp in RFC3339 text format.
	CreationTimestamp string `pulumi:"creationTimestamp"`
	// An optional description of this resource. Provide this property when you create the resource.
	Description string `pulumi:"description"`
	// A URL to the virtual machine instance that handles traffic for this target instance. When creating a target instance, you can provide the fully-qualified URL or a valid partial URL to the desired virtual machine. For example, the following are all valid URLs: - https://www.googleapis.com/compute/v1/projects/project/zones/zone /instances/instance - projects/project/zones/zone/instances/instance - zones/zone/instances/instance
	Instance string `pulumi:"instance"`
	// The type of the resource. Always compute#targetInstance for target instances.
	Kind string `pulumi:"kind"`
	// Name of the resource. Provided by the client when the resource is created. The name must be 1-63 characters long, and comply with RFC1035. Specifically, the name must be 1-63 characters long and match the regular expression `[a-z]([-a-z0-9]*[a-z0-9])?` which means the first character must be a lowercase letter, and all following characters must be a dash, lowercase letter, or digit, except the last character, which cannot be a dash.
	Name string `pulumi:"name"`
	// NAT option controlling how IPs are NAT'ed to the instance. Currently only NO_NAT (default value) is supported.
	NatPolicy string `pulumi:"natPolicy"`
	// The URL of the network this target instance uses to forward traffic. If not specified, the traffic will be forwarded to the network that the default network interface belongs to.
	Network string `pulumi:"network"`
	// Server-defined URL for the resource.
	SelfLink string `pulumi:"selfLink"`
	// URL of the zone where the target instance resides. You must specify this field as part of the HTTP request URL. It is not settable as a field in the request body.
	Zone string `pulumi:"zone"`
}

func LookupTargetInstanceOutput(ctx *pulumi.Context, args LookupTargetInstanceOutputArgs, opts ...pulumi.InvokeOption) LookupTargetInstanceResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupTargetInstanceResult, error) {
			args := v.(LookupTargetInstanceArgs)
			r, err := LookupTargetInstance(ctx, &args, opts...)
			return *r, err
		}).(LookupTargetInstanceResultOutput)
}

type LookupTargetInstanceOutputArgs struct {
	Project        pulumi.StringPtrInput `pulumi:"project"`
	TargetInstance pulumi.StringInput    `pulumi:"targetInstance"`
	Zone           pulumi.StringInput    `pulumi:"zone"`
}

func (LookupTargetInstanceOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupTargetInstanceArgs)(nil)).Elem()
}

type LookupTargetInstanceResultOutput struct{ *pulumi.OutputState }

func (LookupTargetInstanceResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupTargetInstanceResult)(nil)).Elem()
}

func (o LookupTargetInstanceResultOutput) ToLookupTargetInstanceResultOutput() LookupTargetInstanceResultOutput {
	return o
}

func (o LookupTargetInstanceResultOutput) ToLookupTargetInstanceResultOutputWithContext(ctx context.Context) LookupTargetInstanceResultOutput {
	return o
}

// Creation timestamp in RFC3339 text format.
func (o LookupTargetInstanceResultOutput) CreationTimestamp() pulumi.StringOutput {
	return o.ApplyT(func(v LookupTargetInstanceResult) string { return v.CreationTimestamp }).(pulumi.StringOutput)
}

// An optional description of this resource. Provide this property when you create the resource.
func (o LookupTargetInstanceResultOutput) Description() pulumi.StringOutput {
	return o.ApplyT(func(v LookupTargetInstanceResult) string { return v.Description }).(pulumi.StringOutput)
}

// A URL to the virtual machine instance that handles traffic for this target instance. When creating a target instance, you can provide the fully-qualified URL or a valid partial URL to the desired virtual machine. For example, the following are all valid URLs: - https://www.googleapis.com/compute/v1/projects/project/zones/zone /instances/instance - projects/project/zones/zone/instances/instance - zones/zone/instances/instance
func (o LookupTargetInstanceResultOutput) Instance() pulumi.StringOutput {
	return o.ApplyT(func(v LookupTargetInstanceResult) string { return v.Instance }).(pulumi.StringOutput)
}

// The type of the resource. Always compute#targetInstance for target instances.
func (o LookupTargetInstanceResultOutput) Kind() pulumi.StringOutput {
	return o.ApplyT(func(v LookupTargetInstanceResult) string { return v.Kind }).(pulumi.StringOutput)
}

// Name of the resource. Provided by the client when the resource is created. The name must be 1-63 characters long, and comply with RFC1035. Specifically, the name must be 1-63 characters long and match the regular expression `[a-z]([-a-z0-9]*[a-z0-9])?` which means the first character must be a lowercase letter, and all following characters must be a dash, lowercase letter, or digit, except the last character, which cannot be a dash.
func (o LookupTargetInstanceResultOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v LookupTargetInstanceResult) string { return v.Name }).(pulumi.StringOutput)
}

// NAT option controlling how IPs are NAT'ed to the instance. Currently only NO_NAT (default value) is supported.
func (o LookupTargetInstanceResultOutput) NatPolicy() pulumi.StringOutput {
	return o.ApplyT(func(v LookupTargetInstanceResult) string { return v.NatPolicy }).(pulumi.StringOutput)
}

// The URL of the network this target instance uses to forward traffic. If not specified, the traffic will be forwarded to the network that the default network interface belongs to.
func (o LookupTargetInstanceResultOutput) Network() pulumi.StringOutput {
	return o.ApplyT(func(v LookupTargetInstanceResult) string { return v.Network }).(pulumi.StringOutput)
}

// Server-defined URL for the resource.
func (o LookupTargetInstanceResultOutput) SelfLink() pulumi.StringOutput {
	return o.ApplyT(func(v LookupTargetInstanceResult) string { return v.SelfLink }).(pulumi.StringOutput)
}

// URL of the zone where the target instance resides. You must specify this field as part of the HTTP request URL. It is not settable as a field in the request body.
func (o LookupTargetInstanceResultOutput) Zone() pulumi.StringOutput {
	return o.ApplyT(func(v LookupTargetInstanceResult) string { return v.Zone }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupTargetInstanceResultOutput{})
}
