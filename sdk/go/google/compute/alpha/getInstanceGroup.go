// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package alpha

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Returns the specified zonal instance group. Get a list of available zonal instance groups by making a list() request. For managed instance groups, use the instanceGroupManagers or regionInstanceGroupManagers methods instead.
func LookupInstanceGroup(ctx *pulumi.Context, args *LookupInstanceGroupArgs, opts ...pulumi.InvokeOption) (*LookupInstanceGroupResult, error) {
	var rv LookupInstanceGroupResult
	err := ctx.Invoke("google-native:compute/alpha:getInstanceGroup", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupInstanceGroupArgs struct {
	InstanceGroup string  `pulumi:"instanceGroup"`
	Project       *string `pulumi:"project"`
	Zone          string  `pulumi:"zone"`
}

type LookupInstanceGroupResult struct {
	// The creation timestamp for this instance group in RFC3339 text format.
	CreationTimestamp string `pulumi:"creationTimestamp"`
	// An optional description of this resource. Provide this property when you create the resource.
	Description string `pulumi:"description"`
	// The fingerprint of the named ports. The system uses this fingerprint to detect conflicts when multiple users change the named ports concurrently.
	Fingerprint string `pulumi:"fingerprint"`
	// The resource type, which is always compute#instanceGroup for instance groups.
	Kind string `pulumi:"kind"`
	// The name of the instance group. The name must be 1-63 characters long, and comply with RFC1035.
	Name string `pulumi:"name"`
	//  Assigns a name to a port number. For example: {name: "http", port: 80} This allows the system to reference ports by the assigned name instead of a port number. Named ports can also contain multiple ports. For example: [{name: "app1", port: 8080}, {name: "app1", port: 8081}, {name: "app2", port: 8082}] Named ports apply to all instances in this instance group.
	NamedPorts []NamedPortResponse `pulumi:"namedPorts"`
	// The URL of the network to which all instances in the instance group belong. If your instance has multiple network interfaces, then the network and subnetwork fields only refer to the network and subnet used by your primary interface (nic0).
	Network string `pulumi:"network"`
	// The URL of the region where the instance group is located (for regional resources).
	Region string `pulumi:"region"`
	// The URL for this instance group. The server generates this URL.
	SelfLink string `pulumi:"selfLink"`
	// Server-defined URL for this resource with the resource id.
	SelfLinkWithId string `pulumi:"selfLinkWithId"`
	// The total number of instances in the instance group.
	Size int `pulumi:"size"`
	// The URL of the subnetwork to which all instances in the instance group belong. If your instance has multiple network interfaces, then the network and subnetwork fields only refer to the network and subnet used by your primary interface (nic0).
	Subnetwork string `pulumi:"subnetwork"`
	// The URL of the zone where the instance group is located (for zonal resources).
	Zone string `pulumi:"zone"`
}

func LookupInstanceGroupOutput(ctx *pulumi.Context, args LookupInstanceGroupOutputArgs, opts ...pulumi.InvokeOption) LookupInstanceGroupResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupInstanceGroupResult, error) {
			args := v.(LookupInstanceGroupArgs)
			r, err := LookupInstanceGroup(ctx, &args, opts...)
			return *r, err
		}).(LookupInstanceGroupResultOutput)
}

type LookupInstanceGroupOutputArgs struct {
	InstanceGroup pulumi.StringInput    `pulumi:"instanceGroup"`
	Project       pulumi.StringPtrInput `pulumi:"project"`
	Zone          pulumi.StringInput    `pulumi:"zone"`
}

func (LookupInstanceGroupOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupInstanceGroupArgs)(nil)).Elem()
}

type LookupInstanceGroupResultOutput struct{ *pulumi.OutputState }

func (LookupInstanceGroupResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupInstanceGroupResult)(nil)).Elem()
}

func (o LookupInstanceGroupResultOutput) ToLookupInstanceGroupResultOutput() LookupInstanceGroupResultOutput {
	return o
}

func (o LookupInstanceGroupResultOutput) ToLookupInstanceGroupResultOutputWithContext(ctx context.Context) LookupInstanceGroupResultOutput {
	return o
}

// The creation timestamp for this instance group in RFC3339 text format.
func (o LookupInstanceGroupResultOutput) CreationTimestamp() pulumi.StringOutput {
	return o.ApplyT(func(v LookupInstanceGroupResult) string { return v.CreationTimestamp }).(pulumi.StringOutput)
}

// An optional description of this resource. Provide this property when you create the resource.
func (o LookupInstanceGroupResultOutput) Description() pulumi.StringOutput {
	return o.ApplyT(func(v LookupInstanceGroupResult) string { return v.Description }).(pulumi.StringOutput)
}

// The fingerprint of the named ports. The system uses this fingerprint to detect conflicts when multiple users change the named ports concurrently.
func (o LookupInstanceGroupResultOutput) Fingerprint() pulumi.StringOutput {
	return o.ApplyT(func(v LookupInstanceGroupResult) string { return v.Fingerprint }).(pulumi.StringOutput)
}

// The resource type, which is always compute#instanceGroup for instance groups.
func (o LookupInstanceGroupResultOutput) Kind() pulumi.StringOutput {
	return o.ApplyT(func(v LookupInstanceGroupResult) string { return v.Kind }).(pulumi.StringOutput)
}

// The name of the instance group. The name must be 1-63 characters long, and comply with RFC1035.
func (o LookupInstanceGroupResultOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v LookupInstanceGroupResult) string { return v.Name }).(pulumi.StringOutput)
}

//  Assigns a name to a port number. For example: {name: "http", port: 80} This allows the system to reference ports by the assigned name instead of a port number. Named ports can also contain multiple ports. For example: [{name: "app1", port: 8080}, {name: "app1", port: 8081}, {name: "app2", port: 8082}] Named ports apply to all instances in this instance group.
func (o LookupInstanceGroupResultOutput) NamedPorts() NamedPortResponseArrayOutput {
	return o.ApplyT(func(v LookupInstanceGroupResult) []NamedPortResponse { return v.NamedPorts }).(NamedPortResponseArrayOutput)
}

// The URL of the network to which all instances in the instance group belong. If your instance has multiple network interfaces, then the network and subnetwork fields only refer to the network and subnet used by your primary interface (nic0).
func (o LookupInstanceGroupResultOutput) Network() pulumi.StringOutput {
	return o.ApplyT(func(v LookupInstanceGroupResult) string { return v.Network }).(pulumi.StringOutput)
}

// The URL of the region where the instance group is located (for regional resources).
func (o LookupInstanceGroupResultOutput) Region() pulumi.StringOutput {
	return o.ApplyT(func(v LookupInstanceGroupResult) string { return v.Region }).(pulumi.StringOutput)
}

// The URL for this instance group. The server generates this URL.
func (o LookupInstanceGroupResultOutput) SelfLink() pulumi.StringOutput {
	return o.ApplyT(func(v LookupInstanceGroupResult) string { return v.SelfLink }).(pulumi.StringOutput)
}

// Server-defined URL for this resource with the resource id.
func (o LookupInstanceGroupResultOutput) SelfLinkWithId() pulumi.StringOutput {
	return o.ApplyT(func(v LookupInstanceGroupResult) string { return v.SelfLinkWithId }).(pulumi.StringOutput)
}

// The total number of instances in the instance group.
func (o LookupInstanceGroupResultOutput) Size() pulumi.IntOutput {
	return o.ApplyT(func(v LookupInstanceGroupResult) int { return v.Size }).(pulumi.IntOutput)
}

// The URL of the subnetwork to which all instances in the instance group belong. If your instance has multiple network interfaces, then the network and subnetwork fields only refer to the network and subnet used by your primary interface (nic0).
func (o LookupInstanceGroupResultOutput) Subnetwork() pulumi.StringOutput {
	return o.ApplyT(func(v LookupInstanceGroupResult) string { return v.Subnetwork }).(pulumi.StringOutput)
}

// The URL of the zone where the instance group is located (for zonal resources).
func (o LookupInstanceGroupResultOutput) Zone() pulumi.StringOutput {
	return o.ApplyT(func(v LookupInstanceGroupResult) string { return v.Zone }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupInstanceGroupResultOutput{})
}
