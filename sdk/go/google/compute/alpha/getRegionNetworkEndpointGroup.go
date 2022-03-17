// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package alpha

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Returns the specified network endpoint group. Gets a list of available network endpoint groups by making a list() request.
func LookupRegionNetworkEndpointGroup(ctx *pulumi.Context, args *LookupRegionNetworkEndpointGroupArgs, opts ...pulumi.InvokeOption) (*LookupRegionNetworkEndpointGroupResult, error) {
	var rv LookupRegionNetworkEndpointGroupResult
	err := ctx.Invoke("google-native:compute/alpha:getRegionNetworkEndpointGroup", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupRegionNetworkEndpointGroupArgs struct {
	NetworkEndpointGroup string  `pulumi:"networkEndpointGroup"`
	Project              *string `pulumi:"project"`
	Region               string  `pulumi:"region"`
}

type LookupRegionNetworkEndpointGroupResult struct {
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
	// This field is only valid when the network endpoint group is used for load balancing. [Deprecated] This field is deprecated.
	//
	// Deprecated: This field is only valid when the network endpoint group is used for load balancing. [Deprecated] This field is deprecated.
	LoadBalancer NetworkEndpointGroupLbNetworkEndpointGroupResponse `pulumi:"loadBalancer"`
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
	// Server-defined URL for this resource with the resource id.
	SelfLinkWithId string `pulumi:"selfLinkWithId"`
	// Only valid when networkEndpointType is "SERVERLESS". Only one of cloudRun, appEngine, cloudFunction or serverlessDeployment may be set.
	ServerlessDeployment NetworkEndpointGroupServerlessDeploymentResponse `pulumi:"serverlessDeployment"`
	// [Output only] Number of network endpoints in the network endpoint group.
	Size int `pulumi:"size"`
	// Optional URL of the subnetwork to which all network endpoints in the NEG belong.
	Subnetwork string `pulumi:"subnetwork"`
	// Specify the type of this network endpoint group. Only LOAD_BALANCING is valid for now.
	Type string `pulumi:"type"`
	// The URL of the zone where the network endpoint group is located.
	Zone string `pulumi:"zone"`
}

func LookupRegionNetworkEndpointGroupOutput(ctx *pulumi.Context, args LookupRegionNetworkEndpointGroupOutputArgs, opts ...pulumi.InvokeOption) LookupRegionNetworkEndpointGroupResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupRegionNetworkEndpointGroupResult, error) {
			args := v.(LookupRegionNetworkEndpointGroupArgs)
			r, err := LookupRegionNetworkEndpointGroup(ctx, &args, opts...)
			return *r, err
		}).(LookupRegionNetworkEndpointGroupResultOutput)
}

type LookupRegionNetworkEndpointGroupOutputArgs struct {
	NetworkEndpointGroup pulumi.StringInput    `pulumi:"networkEndpointGroup"`
	Project              pulumi.StringPtrInput `pulumi:"project"`
	Region               pulumi.StringInput    `pulumi:"region"`
}

func (LookupRegionNetworkEndpointGroupOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupRegionNetworkEndpointGroupArgs)(nil)).Elem()
}

type LookupRegionNetworkEndpointGroupResultOutput struct{ *pulumi.OutputState }

func (LookupRegionNetworkEndpointGroupResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupRegionNetworkEndpointGroupResult)(nil)).Elem()
}

func (o LookupRegionNetworkEndpointGroupResultOutput) ToLookupRegionNetworkEndpointGroupResultOutput() LookupRegionNetworkEndpointGroupResultOutput {
	return o
}

func (o LookupRegionNetworkEndpointGroupResultOutput) ToLookupRegionNetworkEndpointGroupResultOutputWithContext(ctx context.Context) LookupRegionNetworkEndpointGroupResultOutput {
	return o
}

// Metadata defined as annotations on the network endpoint group.
func (o LookupRegionNetworkEndpointGroupResultOutput) Annotations() pulumi.StringMapOutput {
	return o.ApplyT(func(v LookupRegionNetworkEndpointGroupResult) map[string]string { return v.Annotations }).(pulumi.StringMapOutput)
}

// Only valid when networkEndpointType is "SERVERLESS". Only one of cloudRun, appEngine or cloudFunction may be set.
func (o LookupRegionNetworkEndpointGroupResultOutput) AppEngine() NetworkEndpointGroupAppEngineResponseOutput {
	return o.ApplyT(func(v LookupRegionNetworkEndpointGroupResult) NetworkEndpointGroupAppEngineResponse {
		return v.AppEngine
	}).(NetworkEndpointGroupAppEngineResponseOutput)
}

// Only valid when networkEndpointType is "SERVERLESS". Only one of cloudRun, appEngine or cloudFunction may be set.
func (o LookupRegionNetworkEndpointGroupResultOutput) CloudFunction() NetworkEndpointGroupCloudFunctionResponseOutput {
	return o.ApplyT(func(v LookupRegionNetworkEndpointGroupResult) NetworkEndpointGroupCloudFunctionResponse {
		return v.CloudFunction
	}).(NetworkEndpointGroupCloudFunctionResponseOutput)
}

// Only valid when networkEndpointType is "SERVERLESS". Only one of cloudRun, appEngine or cloudFunction may be set.
func (o LookupRegionNetworkEndpointGroupResultOutput) CloudRun() NetworkEndpointGroupCloudRunResponseOutput {
	return o.ApplyT(func(v LookupRegionNetworkEndpointGroupResult) NetworkEndpointGroupCloudRunResponse { return v.CloudRun }).(NetworkEndpointGroupCloudRunResponseOutput)
}

// Creation timestamp in RFC3339 text format.
func (o LookupRegionNetworkEndpointGroupResultOutput) CreationTimestamp() pulumi.StringOutput {
	return o.ApplyT(func(v LookupRegionNetworkEndpointGroupResult) string { return v.CreationTimestamp }).(pulumi.StringOutput)
}

// The default port used if the port number is not specified in the network endpoint.
func (o LookupRegionNetworkEndpointGroupResultOutput) DefaultPort() pulumi.IntOutput {
	return o.ApplyT(func(v LookupRegionNetworkEndpointGroupResult) int { return v.DefaultPort }).(pulumi.IntOutput)
}

// An optional description of this resource. Provide this property when you create the resource.
func (o LookupRegionNetworkEndpointGroupResultOutput) Description() pulumi.StringOutput {
	return o.ApplyT(func(v LookupRegionNetworkEndpointGroupResult) string { return v.Description }).(pulumi.StringOutput)
}

// Type of the resource. Always compute#networkEndpointGroup for network endpoint group.
func (o LookupRegionNetworkEndpointGroupResultOutput) Kind() pulumi.StringOutput {
	return o.ApplyT(func(v LookupRegionNetworkEndpointGroupResult) string { return v.Kind }).(pulumi.StringOutput)
}

// This field is only valid when the network endpoint group is used for load balancing. [Deprecated] This field is deprecated.
//
// Deprecated: This field is only valid when the network endpoint group is used for load balancing. [Deprecated] This field is deprecated.
func (o LookupRegionNetworkEndpointGroupResultOutput) LoadBalancer() NetworkEndpointGroupLbNetworkEndpointGroupResponseOutput {
	return o.ApplyT(func(v LookupRegionNetworkEndpointGroupResult) NetworkEndpointGroupLbNetworkEndpointGroupResponse {
		return v.LoadBalancer
	}).(NetworkEndpointGroupLbNetworkEndpointGroupResponseOutput)
}

// Name of the resource; provided by the client when the resource is created. The name must be 1-63 characters long, and comply with RFC1035. Specifically, the name must be 1-63 characters long and match the regular expression `[a-z]([-a-z0-9]*[a-z0-9])?` which means the first character must be a lowercase letter, and all following characters must be a dash, lowercase letter, or digit, except the last character, which cannot be a dash.
func (o LookupRegionNetworkEndpointGroupResultOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v LookupRegionNetworkEndpointGroupResult) string { return v.Name }).(pulumi.StringOutput)
}

// The URL of the network to which all network endpoints in the NEG belong. Uses "default" project network if unspecified.
func (o LookupRegionNetworkEndpointGroupResultOutput) Network() pulumi.StringOutput {
	return o.ApplyT(func(v LookupRegionNetworkEndpointGroupResult) string { return v.Network }).(pulumi.StringOutput)
}

// Type of network endpoints in this network endpoint group. Can be one of GCE_VM_IP, GCE_VM_IP_PORT, NON_GCP_PRIVATE_IP_PORT, INTERNET_FQDN_PORT, INTERNET_IP_PORT, SERVERLESS, PRIVATE_SERVICE_CONNECT.
func (o LookupRegionNetworkEndpointGroupResultOutput) NetworkEndpointType() pulumi.StringOutput {
	return o.ApplyT(func(v LookupRegionNetworkEndpointGroupResult) string { return v.NetworkEndpointType }).(pulumi.StringOutput)
}

// The target service url used to set up private service connection to a Google API. An example value is: "asia-northeast3-cloudkms.googleapis.com"
func (o LookupRegionNetworkEndpointGroupResultOutput) PscTargetService() pulumi.StringOutput {
	return o.ApplyT(func(v LookupRegionNetworkEndpointGroupResult) string { return v.PscTargetService }).(pulumi.StringOutput)
}

// The URL of the region where the network endpoint group is located.
func (o LookupRegionNetworkEndpointGroupResultOutput) Region() pulumi.StringOutput {
	return o.ApplyT(func(v LookupRegionNetworkEndpointGroupResult) string { return v.Region }).(pulumi.StringOutput)
}

// Server-defined URL for the resource.
func (o LookupRegionNetworkEndpointGroupResultOutput) SelfLink() pulumi.StringOutput {
	return o.ApplyT(func(v LookupRegionNetworkEndpointGroupResult) string { return v.SelfLink }).(pulumi.StringOutput)
}

// Server-defined URL for this resource with the resource id.
func (o LookupRegionNetworkEndpointGroupResultOutput) SelfLinkWithId() pulumi.StringOutput {
	return o.ApplyT(func(v LookupRegionNetworkEndpointGroupResult) string { return v.SelfLinkWithId }).(pulumi.StringOutput)
}

// Only valid when networkEndpointType is "SERVERLESS". Only one of cloudRun, appEngine, cloudFunction or serverlessDeployment may be set.
func (o LookupRegionNetworkEndpointGroupResultOutput) ServerlessDeployment() NetworkEndpointGroupServerlessDeploymentResponseOutput {
	return o.ApplyT(func(v LookupRegionNetworkEndpointGroupResult) NetworkEndpointGroupServerlessDeploymentResponse {
		return v.ServerlessDeployment
	}).(NetworkEndpointGroupServerlessDeploymentResponseOutput)
}

// [Output only] Number of network endpoints in the network endpoint group.
func (o LookupRegionNetworkEndpointGroupResultOutput) Size() pulumi.IntOutput {
	return o.ApplyT(func(v LookupRegionNetworkEndpointGroupResult) int { return v.Size }).(pulumi.IntOutput)
}

// Optional URL of the subnetwork to which all network endpoints in the NEG belong.
func (o LookupRegionNetworkEndpointGroupResultOutput) Subnetwork() pulumi.StringOutput {
	return o.ApplyT(func(v LookupRegionNetworkEndpointGroupResult) string { return v.Subnetwork }).(pulumi.StringOutput)
}

// Specify the type of this network endpoint group. Only LOAD_BALANCING is valid for now.
func (o LookupRegionNetworkEndpointGroupResultOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v LookupRegionNetworkEndpointGroupResult) string { return v.Type }).(pulumi.StringOutput)
}

// The URL of the zone where the network endpoint group is located.
func (o LookupRegionNetworkEndpointGroupResultOutput) Zone() pulumi.StringOutput {
	return o.ApplyT(func(v LookupRegionNetworkEndpointGroupResult) string { return v.Zone }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupRegionNetworkEndpointGroupResultOutput{})
}
