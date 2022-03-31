// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package v1

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Returns the specified network endpoint group. Gets a list of available network endpoint groups by making a list() request.
func LookupNetworkEndpointGroup(ctx *pulumi.Context, args *LookupNetworkEndpointGroupArgs, opts ...pulumi.InvokeOption) (*LookupNetworkEndpointGroupResult, error) {
	var rv LookupNetworkEndpointGroupResult
	err := ctx.Invoke("google-native:compute/v1:getNetworkEndpointGroup", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupNetworkEndpointGroupArgs struct {
	NetworkEndpointGroup string  `pulumi:"networkEndpointGroup"`
	Project              *string `pulumi:"project"`
	Zone                 string  `pulumi:"zone"`
}

type LookupNetworkEndpointGroupResult struct {
	// Metadata defined as annotations on the network endpoint group.
	Annotations map[string]string `pulumi:"annotations"`
	// Only valid when networkEndpointType is "SERVERLESS". Only one of cloudRun, appEngine or cloudFunction may be set.
	AppEngine NetworkEndpointGroupAppEngineResponse `pulumi:"appEngine"`
	// Only valid when networkEndpointType is "SERVERLESS". Only one of cloudRun, appEngine or cloudFunction may be set.
	CloudFunction NetworkEndpointGroupCloudFunctionResponse `pulumi:"cloudFunction"`
	// Only valid when networkEndpointType is "SERVERLESS". Only one of cloudRun, appEngine or cloudFunction may be set.
	CloudRun NetworkEndpointGroupCloudRunResponse `pulumi:"cloudRun"`
	// Creation timestamp in RFC3339 text format.
	CreationTimestamp string `pulumi:"creationTimestamp"`
	// The default port used if the port number is not specified in the network endpoint.
	DefaultPort int `pulumi:"defaultPort"`
	// An optional description of this resource. Provide this property when you create the resource.
	Description string `pulumi:"description"`
	// Type of the resource. Always compute#networkEndpointGroup for network endpoint group.
	Kind string `pulumi:"kind"`
	// Name of the resource; provided by the client when the resource is created. The name must be 1-63 characters long, and comply with RFC1035. Specifically, the name must be 1-63 characters long and match the regular expression `[a-z]([-a-z0-9]*[a-z0-9])?` which means the first character must be a lowercase letter, and all following characters must be a dash, lowercase letter, or digit, except the last character, which cannot be a dash.
	Name string `pulumi:"name"`
	// The URL of the network to which all network endpoints in the NEG belong. Uses "default" project network if unspecified.
	Network string `pulumi:"network"`
	// Type of network endpoints in this network endpoint group. Can be one of GCE_VM_IP, GCE_VM_IP_PORT, NON_GCP_PRIVATE_IP_PORT, INTERNET_FQDN_PORT, INTERNET_IP_PORT, SERVERLESS, PRIVATE_SERVICE_CONNECT.
	NetworkEndpointType string `pulumi:"networkEndpointType"`
	// The target service url used to set up private service connection to a Google API. An example value is: "asia-northeast3-cloudkms.googleapis.com"
	PscTargetService string `pulumi:"pscTargetService"`
	// The URL of the region where the network endpoint group is located.
	Region string `pulumi:"region"`
	// Server-defined URL for the resource.
	SelfLink string `pulumi:"selfLink"`
	// [Output only] Number of network endpoints in the network endpoint group.
	Size int `pulumi:"size"`
	// Optional URL of the subnetwork to which all network endpoints in the NEG belong.
	Subnetwork string `pulumi:"subnetwork"`
	// The URL of the zone where the network endpoint group is located.
	Zone string `pulumi:"zone"`
}

func LookupNetworkEndpointGroupOutput(ctx *pulumi.Context, args LookupNetworkEndpointGroupOutputArgs, opts ...pulumi.InvokeOption) LookupNetworkEndpointGroupResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupNetworkEndpointGroupResult, error) {
			args := v.(LookupNetworkEndpointGroupArgs)
			r, err := LookupNetworkEndpointGroup(ctx, &args, opts...)
			return *r, err
		}).(LookupNetworkEndpointGroupResultOutput)
}

type LookupNetworkEndpointGroupOutputArgs struct {
	NetworkEndpointGroup pulumi.StringInput    `pulumi:"networkEndpointGroup"`
	Project              pulumi.StringPtrInput `pulumi:"project"`
	Zone                 pulumi.StringInput    `pulumi:"zone"`
}

func (LookupNetworkEndpointGroupOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupNetworkEndpointGroupArgs)(nil)).Elem()
}

type LookupNetworkEndpointGroupResultOutput struct{ *pulumi.OutputState }

func (LookupNetworkEndpointGroupResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupNetworkEndpointGroupResult)(nil)).Elem()
}

func (o LookupNetworkEndpointGroupResultOutput) ToLookupNetworkEndpointGroupResultOutput() LookupNetworkEndpointGroupResultOutput {
	return o
}

func (o LookupNetworkEndpointGroupResultOutput) ToLookupNetworkEndpointGroupResultOutputWithContext(ctx context.Context) LookupNetworkEndpointGroupResultOutput {
	return o
}

// Metadata defined as annotations on the network endpoint group.
func (o LookupNetworkEndpointGroupResultOutput) Annotations() pulumi.StringMapOutput {
	return o.ApplyT(func(v LookupNetworkEndpointGroupResult) map[string]string { return v.Annotations }).(pulumi.StringMapOutput)
}

// Only valid when networkEndpointType is "SERVERLESS". Only one of cloudRun, appEngine or cloudFunction may be set.
func (o LookupNetworkEndpointGroupResultOutput) AppEngine() NetworkEndpointGroupAppEngineResponseOutput {
	return o.ApplyT(func(v LookupNetworkEndpointGroupResult) NetworkEndpointGroupAppEngineResponse { return v.AppEngine }).(NetworkEndpointGroupAppEngineResponseOutput)
}

// Only valid when networkEndpointType is "SERVERLESS". Only one of cloudRun, appEngine or cloudFunction may be set.
func (o LookupNetworkEndpointGroupResultOutput) CloudFunction() NetworkEndpointGroupCloudFunctionResponseOutput {
	return o.ApplyT(func(v LookupNetworkEndpointGroupResult) NetworkEndpointGroupCloudFunctionResponse {
		return v.CloudFunction
	}).(NetworkEndpointGroupCloudFunctionResponseOutput)
}

// Only valid when networkEndpointType is "SERVERLESS". Only one of cloudRun, appEngine or cloudFunction may be set.
func (o LookupNetworkEndpointGroupResultOutput) CloudRun() NetworkEndpointGroupCloudRunResponseOutput {
	return o.ApplyT(func(v LookupNetworkEndpointGroupResult) NetworkEndpointGroupCloudRunResponse { return v.CloudRun }).(NetworkEndpointGroupCloudRunResponseOutput)
}

// Creation timestamp in RFC3339 text format.
func (o LookupNetworkEndpointGroupResultOutput) CreationTimestamp() pulumi.StringOutput {
	return o.ApplyT(func(v LookupNetworkEndpointGroupResult) string { return v.CreationTimestamp }).(pulumi.StringOutput)
}

// The default port used if the port number is not specified in the network endpoint.
func (o LookupNetworkEndpointGroupResultOutput) DefaultPort() pulumi.IntOutput {
	return o.ApplyT(func(v LookupNetworkEndpointGroupResult) int { return v.DefaultPort }).(pulumi.IntOutput)
}

// An optional description of this resource. Provide this property when you create the resource.
func (o LookupNetworkEndpointGroupResultOutput) Description() pulumi.StringOutput {
	return o.ApplyT(func(v LookupNetworkEndpointGroupResult) string { return v.Description }).(pulumi.StringOutput)
}

// Type of the resource. Always compute#networkEndpointGroup for network endpoint group.
func (o LookupNetworkEndpointGroupResultOutput) Kind() pulumi.StringOutput {
	return o.ApplyT(func(v LookupNetworkEndpointGroupResult) string { return v.Kind }).(pulumi.StringOutput)
}

// Name of the resource; provided by the client when the resource is created. The name must be 1-63 characters long, and comply with RFC1035. Specifically, the name must be 1-63 characters long and match the regular expression `[a-z]([-a-z0-9]*[a-z0-9])?` which means the first character must be a lowercase letter, and all following characters must be a dash, lowercase letter, or digit, except the last character, which cannot be a dash.
func (o LookupNetworkEndpointGroupResultOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v LookupNetworkEndpointGroupResult) string { return v.Name }).(pulumi.StringOutput)
}

// The URL of the network to which all network endpoints in the NEG belong. Uses "default" project network if unspecified.
func (o LookupNetworkEndpointGroupResultOutput) Network() pulumi.StringOutput {
	return o.ApplyT(func(v LookupNetworkEndpointGroupResult) string { return v.Network }).(pulumi.StringOutput)
}

// Type of network endpoints in this network endpoint group. Can be one of GCE_VM_IP, GCE_VM_IP_PORT, NON_GCP_PRIVATE_IP_PORT, INTERNET_FQDN_PORT, INTERNET_IP_PORT, SERVERLESS, PRIVATE_SERVICE_CONNECT.
func (o LookupNetworkEndpointGroupResultOutput) NetworkEndpointType() pulumi.StringOutput {
	return o.ApplyT(func(v LookupNetworkEndpointGroupResult) string { return v.NetworkEndpointType }).(pulumi.StringOutput)
}

// The target service url used to set up private service connection to a Google API. An example value is: "asia-northeast3-cloudkms.googleapis.com"
func (o LookupNetworkEndpointGroupResultOutput) PscTargetService() pulumi.StringOutput {
	return o.ApplyT(func(v LookupNetworkEndpointGroupResult) string { return v.PscTargetService }).(pulumi.StringOutput)
}

// The URL of the region where the network endpoint group is located.
func (o LookupNetworkEndpointGroupResultOutput) Region() pulumi.StringOutput {
	return o.ApplyT(func(v LookupNetworkEndpointGroupResult) string { return v.Region }).(pulumi.StringOutput)
}

// Server-defined URL for the resource.
func (o LookupNetworkEndpointGroupResultOutput) SelfLink() pulumi.StringOutput {
	return o.ApplyT(func(v LookupNetworkEndpointGroupResult) string { return v.SelfLink }).(pulumi.StringOutput)
}

// [Output only] Number of network endpoints in the network endpoint group.
func (o LookupNetworkEndpointGroupResultOutput) Size() pulumi.IntOutput {
	return o.ApplyT(func(v LookupNetworkEndpointGroupResult) int { return v.Size }).(pulumi.IntOutput)
}

// Optional URL of the subnetwork to which all network endpoints in the NEG belong.
func (o LookupNetworkEndpointGroupResultOutput) Subnetwork() pulumi.StringOutput {
	return o.ApplyT(func(v LookupNetworkEndpointGroupResult) string { return v.Subnetwork }).(pulumi.StringOutput)
}

// The URL of the zone where the network endpoint group is located.
func (o LookupNetworkEndpointGroupResultOutput) Zone() pulumi.StringOutput {
	return o.ApplyT(func(v LookupNetworkEndpointGroupResult) string { return v.Zone }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupNetworkEndpointGroupResultOutput{})
}
