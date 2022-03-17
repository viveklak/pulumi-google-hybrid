// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package alpha

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Returns the specified Router resource. Gets a list of available routers by making a list() request.
func LookupRouter(ctx *pulumi.Context, args *LookupRouterArgs, opts ...pulumi.InvokeOption) (*LookupRouterResult, error) {
	var rv LookupRouterResult
	err := ctx.Invoke("google-native:compute/alpha:getRouter", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupRouterArgs struct {
	Project *string `pulumi:"project"`
	Region  string  `pulumi:"region"`
	Router  string  `pulumi:"router"`
}

type LookupRouterResult struct {
	// BGP information specific to this router.
	Bgp RouterBgpResponse `pulumi:"bgp"`
	// BGP information that must be configured into the routing stack to establish BGP peering. This information must specify the peer ASN and either the interface name, IP address, or peer IP address. Please refer to RFC4273.
	BgpPeers []RouterBgpPeerResponse `pulumi:"bgpPeers"`
	// Creation timestamp in RFC3339 text format.
	CreationTimestamp string `pulumi:"creationTimestamp"`
	// An optional description of this resource. Provide this property when you create the resource.
	Description string `pulumi:"description"`
	// Indicates if a router is dedicated for use with encrypted VLAN attachments (interconnectAttachments). Not currently available publicly.
	EncryptedInterconnectRouter bool `pulumi:"encryptedInterconnectRouter"`
	// Router interfaces. Each interface requires either one linked resource, (for example, linkedVpnTunnel), or IP address and IP address range (for example, ipRange), or both.
	Interfaces []RouterInterfaceResponse `pulumi:"interfaces"`
	// Type of resource. Always compute#router for routers.
	Kind string `pulumi:"kind"`
	// Keys used for MD5 authentication.
	Md5AuthenticationKeys []RouterMd5AuthenticationKeyResponse `pulumi:"md5AuthenticationKeys"`
	// Name of the resource. Provided by the client when the resource is created. The name must be 1-63 characters long, and comply with RFC1035. Specifically, the name must be 1-63 characters long and match the regular expression `[a-z]([-a-z0-9]*[a-z0-9])?` which means the first character must be a lowercase letter, and all following characters must be a dash, lowercase letter, or digit, except the last character, which cannot be a dash.
	Name string `pulumi:"name"`
	// A list of NAT services created in this router.
	Nats []RouterNatResponse `pulumi:"nats"`
	// URI of the network to which this router belongs.
	Network string `pulumi:"network"`
	// URI of the region where the router resides. You must specify this field as part of the HTTP request URL. It is not settable as a field in the request body.
	Region string `pulumi:"region"`
	// Server-defined URL for the resource.
	SelfLink string `pulumi:"selfLink"`
	// Server-defined URL for this resource with the resource id.
	SelfLinkWithId string `pulumi:"selfLinkWithId"`
}

func LookupRouterOutput(ctx *pulumi.Context, args LookupRouterOutputArgs, opts ...pulumi.InvokeOption) LookupRouterResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupRouterResult, error) {
			args := v.(LookupRouterArgs)
			r, err := LookupRouter(ctx, &args, opts...)
			return *r, err
		}).(LookupRouterResultOutput)
}

type LookupRouterOutputArgs struct {
	Project pulumi.StringPtrInput `pulumi:"project"`
	Region  pulumi.StringInput    `pulumi:"region"`
	Router  pulumi.StringInput    `pulumi:"router"`
}

func (LookupRouterOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupRouterArgs)(nil)).Elem()
}

type LookupRouterResultOutput struct{ *pulumi.OutputState }

func (LookupRouterResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupRouterResult)(nil)).Elem()
}

func (o LookupRouterResultOutput) ToLookupRouterResultOutput() LookupRouterResultOutput {
	return o
}

func (o LookupRouterResultOutput) ToLookupRouterResultOutputWithContext(ctx context.Context) LookupRouterResultOutput {
	return o
}

// BGP information specific to this router.
func (o LookupRouterResultOutput) Bgp() RouterBgpResponseOutput {
	return o.ApplyT(func(v LookupRouterResult) RouterBgpResponse { return v.Bgp }).(RouterBgpResponseOutput)
}

// BGP information that must be configured into the routing stack to establish BGP peering. This information must specify the peer ASN and either the interface name, IP address, or peer IP address. Please refer to RFC4273.
func (o LookupRouterResultOutput) BgpPeers() RouterBgpPeerResponseArrayOutput {
	return o.ApplyT(func(v LookupRouterResult) []RouterBgpPeerResponse { return v.BgpPeers }).(RouterBgpPeerResponseArrayOutput)
}

// Creation timestamp in RFC3339 text format.
func (o LookupRouterResultOutput) CreationTimestamp() pulumi.StringOutput {
	return o.ApplyT(func(v LookupRouterResult) string { return v.CreationTimestamp }).(pulumi.StringOutput)
}

// An optional description of this resource. Provide this property when you create the resource.
func (o LookupRouterResultOutput) Description() pulumi.StringOutput {
	return o.ApplyT(func(v LookupRouterResult) string { return v.Description }).(pulumi.StringOutput)
}

// Indicates if a router is dedicated for use with encrypted VLAN attachments (interconnectAttachments). Not currently available publicly.
func (o LookupRouterResultOutput) EncryptedInterconnectRouter() pulumi.BoolOutput {
	return o.ApplyT(func(v LookupRouterResult) bool { return v.EncryptedInterconnectRouter }).(pulumi.BoolOutput)
}

// Router interfaces. Each interface requires either one linked resource, (for example, linkedVpnTunnel), or IP address and IP address range (for example, ipRange), or both.
func (o LookupRouterResultOutput) Interfaces() RouterInterfaceResponseArrayOutput {
	return o.ApplyT(func(v LookupRouterResult) []RouterInterfaceResponse { return v.Interfaces }).(RouterInterfaceResponseArrayOutput)
}

// Type of resource. Always compute#router for routers.
func (o LookupRouterResultOutput) Kind() pulumi.StringOutput {
	return o.ApplyT(func(v LookupRouterResult) string { return v.Kind }).(pulumi.StringOutput)
}

// Keys used for MD5 authentication.
func (o LookupRouterResultOutput) Md5AuthenticationKeys() RouterMd5AuthenticationKeyResponseArrayOutput {
	return o.ApplyT(func(v LookupRouterResult) []RouterMd5AuthenticationKeyResponse { return v.Md5AuthenticationKeys }).(RouterMd5AuthenticationKeyResponseArrayOutput)
}

// Name of the resource. Provided by the client when the resource is created. The name must be 1-63 characters long, and comply with RFC1035. Specifically, the name must be 1-63 characters long and match the regular expression `[a-z]([-a-z0-9]*[a-z0-9])?` which means the first character must be a lowercase letter, and all following characters must be a dash, lowercase letter, or digit, except the last character, which cannot be a dash.
func (o LookupRouterResultOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v LookupRouterResult) string { return v.Name }).(pulumi.StringOutput)
}

// A list of NAT services created in this router.
func (o LookupRouterResultOutput) Nats() RouterNatResponseArrayOutput {
	return o.ApplyT(func(v LookupRouterResult) []RouterNatResponse { return v.Nats }).(RouterNatResponseArrayOutput)
}

// URI of the network to which this router belongs.
func (o LookupRouterResultOutput) Network() pulumi.StringOutput {
	return o.ApplyT(func(v LookupRouterResult) string { return v.Network }).(pulumi.StringOutput)
}

// URI of the region where the router resides. You must specify this field as part of the HTTP request URL. It is not settable as a field in the request body.
func (o LookupRouterResultOutput) Region() pulumi.StringOutput {
	return o.ApplyT(func(v LookupRouterResult) string { return v.Region }).(pulumi.StringOutput)
}

// Server-defined URL for the resource.
func (o LookupRouterResultOutput) SelfLink() pulumi.StringOutput {
	return o.ApplyT(func(v LookupRouterResult) string { return v.SelfLink }).(pulumi.StringOutput)
}

// Server-defined URL for this resource with the resource id.
func (o LookupRouterResultOutput) SelfLinkWithId() pulumi.StringOutput {
	return o.ApplyT(func(v LookupRouterResult) string { return v.SelfLinkWithId }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupRouterResultOutput{})
}
