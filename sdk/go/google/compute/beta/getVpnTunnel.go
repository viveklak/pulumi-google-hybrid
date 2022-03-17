// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package beta

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Returns the specified VpnTunnel resource. Gets a list of available VPN tunnels by making a list() request.
func LookupVpnTunnel(ctx *pulumi.Context, args *LookupVpnTunnelArgs, opts ...pulumi.InvokeOption) (*LookupVpnTunnelResult, error) {
	var rv LookupVpnTunnelResult
	err := ctx.Invoke("google-native:compute/beta:getVpnTunnel", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupVpnTunnelArgs struct {
	Project   *string `pulumi:"project"`
	Region    string  `pulumi:"region"`
	VpnTunnel string  `pulumi:"vpnTunnel"`
}

type LookupVpnTunnelResult struct {
	// Creation timestamp in RFC3339 text format.
	CreationTimestamp string `pulumi:"creationTimestamp"`
	// An optional description of this resource. Provide this property when you create the resource.
	Description string `pulumi:"description"`
	// Detailed status message for the VPN tunnel.
	DetailedStatus string `pulumi:"detailedStatus"`
	// IKE protocol version to use when establishing the VPN tunnel with the peer VPN gateway. Acceptable IKE versions are 1 or 2. The default version is 2.
	IkeVersion int `pulumi:"ikeVersion"`
	// Type of resource. Always compute#vpnTunnel for VPN tunnels.
	Kind string `pulumi:"kind"`
	// A fingerprint for the labels being applied to this VpnTunnel, which is essentially a hash of the labels set used for optimistic locking. The fingerprint is initially generated by Compute Engine and changes after every request to modify or update labels. You must always provide an up-to-date fingerprint hash in order to update or change labels, otherwise the request will fail with error 412 conditionNotMet. To see the latest fingerprint, make a get() request to retrieve a VpnTunnel.
	LabelFingerprint string `pulumi:"labelFingerprint"`
	// Labels for this resource. These can only be added or modified by the setLabels method. Each label key/value pair must comply with RFC1035. Label values may be empty.
	Labels map[string]string `pulumi:"labels"`
	// Local traffic selector to use when establishing the VPN tunnel with the peer VPN gateway. The value should be a CIDR formatted string, for example: 192.168.0.0/16. The ranges must be disjoint. Only IPv4 is supported.
	LocalTrafficSelector []string `pulumi:"localTrafficSelector"`
	// Name of the resource. Provided by the client when the resource is created. The name must be 1-63 characters long, and comply with RFC1035. Specifically, the name must be 1-63 characters long and match the regular expression `[a-z]([-a-z0-9]*[a-z0-9])?` which means the first character must be a lowercase letter, and all following characters must be a dash, lowercase letter, or digit, except the last character, which cannot be a dash.
	Name string `pulumi:"name"`
	// URL of the peer side external VPN gateway to which this VPN tunnel is connected. Provided by the client when the VPN tunnel is created. This field is exclusive with the field peerGcpGateway.
	PeerExternalGateway string `pulumi:"peerExternalGateway"`
	// The interface ID of the external VPN gateway to which this VPN tunnel is connected. Provided by the client when the VPN tunnel is created.
	PeerExternalGatewayInterface int `pulumi:"peerExternalGatewayInterface"`
	// URL of the peer side HA GCP VPN gateway to which this VPN tunnel is connected. Provided by the client when the VPN tunnel is created. This field can be used when creating highly available VPN from VPC network to VPC network, the field is exclusive with the field peerExternalGateway. If provided, the VPN tunnel will automatically use the same vpnGatewayInterface ID in the peer GCP VPN gateway.
	PeerGcpGateway string `pulumi:"peerGcpGateway"`
	// IP address of the peer VPN gateway. Only IPv4 is supported.
	PeerIp string `pulumi:"peerIp"`
	// URL of the region where the VPN tunnel resides. You must specify this field as part of the HTTP request URL. It is not settable as a field in the request body.
	Region string `pulumi:"region"`
	// Remote traffic selectors to use when establishing the VPN tunnel with the peer VPN gateway. The value should be a CIDR formatted string, for example: 192.168.0.0/16. The ranges should be disjoint. Only IPv4 is supported.
	RemoteTrafficSelector []string `pulumi:"remoteTrafficSelector"`
	// URL of the router resource to be used for dynamic routing.
	Router string `pulumi:"router"`
	// Server-defined URL for the resource.
	SelfLink string `pulumi:"selfLink"`
	// Shared secret used to set the secure session between the Cloud VPN gateway and the peer VPN gateway.
	SharedSecret string `pulumi:"sharedSecret"`
	// Hash of the shared secret.
	SharedSecretHash string `pulumi:"sharedSecretHash"`
	// The status of the VPN tunnel, which can be one of the following: - PROVISIONING: Resource is being allocated for the VPN tunnel. - WAITING_FOR_FULL_CONFIG: Waiting to receive all VPN-related configs from the user. Network, TargetVpnGateway, VpnTunnel, ForwardingRule, and Route resources are needed to setup the VPN tunnel. - FIRST_HANDSHAKE: Successful first handshake with the peer VPN. - ESTABLISHED: Secure session is successfully established with the peer VPN. - NETWORK_ERROR: Deprecated, replaced by NO_INCOMING_PACKETS - AUTHORIZATION_ERROR: Auth error (for example, bad shared secret). - NEGOTIATION_FAILURE: Handshake failed. - DEPROVISIONING: Resources are being deallocated for the VPN tunnel. - FAILED: Tunnel creation has failed and the tunnel is not ready to be used. - NO_INCOMING_PACKETS: No incoming packets from peer. - REJECTED: Tunnel configuration was rejected, can be result of being denied access. - ALLOCATING_RESOURCES: Cloud VPN is in the process of allocating all required resources. - STOPPED: Tunnel is stopped due to its Forwarding Rules being deleted for Classic VPN tunnels or the project is in frozen state. - PEER_IDENTITY_MISMATCH: Peer identity does not match peer IP, probably behind NAT. - TS_NARROWING_NOT_ALLOWED: Traffic selector narrowing not allowed for an HA-VPN tunnel.
	Status string `pulumi:"status"`
	// URL of the Target VPN gateway with which this VPN tunnel is associated. Provided by the client when the VPN tunnel is created.
	TargetVpnGateway string `pulumi:"targetVpnGateway"`
	// URL of the VPN gateway with which this VPN tunnel is associated. Provided by the client when the VPN tunnel is created. This must be used (instead of target_vpn_gateway) if a High Availability VPN gateway resource is created.
	VpnGateway string `pulumi:"vpnGateway"`
	// The interface ID of the VPN gateway with which this VPN tunnel is associated.
	VpnGatewayInterface int `pulumi:"vpnGatewayInterface"`
}

func LookupVpnTunnelOutput(ctx *pulumi.Context, args LookupVpnTunnelOutputArgs, opts ...pulumi.InvokeOption) LookupVpnTunnelResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupVpnTunnelResult, error) {
			args := v.(LookupVpnTunnelArgs)
			r, err := LookupVpnTunnel(ctx, &args, opts...)
			return *r, err
		}).(LookupVpnTunnelResultOutput)
}

type LookupVpnTunnelOutputArgs struct {
	Project   pulumi.StringPtrInput `pulumi:"project"`
	Region    pulumi.StringInput    `pulumi:"region"`
	VpnTunnel pulumi.StringInput    `pulumi:"vpnTunnel"`
}

func (LookupVpnTunnelOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupVpnTunnelArgs)(nil)).Elem()
}

type LookupVpnTunnelResultOutput struct{ *pulumi.OutputState }

func (LookupVpnTunnelResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupVpnTunnelResult)(nil)).Elem()
}

func (o LookupVpnTunnelResultOutput) ToLookupVpnTunnelResultOutput() LookupVpnTunnelResultOutput {
	return o
}

func (o LookupVpnTunnelResultOutput) ToLookupVpnTunnelResultOutputWithContext(ctx context.Context) LookupVpnTunnelResultOutput {
	return o
}

// Creation timestamp in RFC3339 text format.
func (o LookupVpnTunnelResultOutput) CreationTimestamp() pulumi.StringOutput {
	return o.ApplyT(func(v LookupVpnTunnelResult) string { return v.CreationTimestamp }).(pulumi.StringOutput)
}

// An optional description of this resource. Provide this property when you create the resource.
func (o LookupVpnTunnelResultOutput) Description() pulumi.StringOutput {
	return o.ApplyT(func(v LookupVpnTunnelResult) string { return v.Description }).(pulumi.StringOutput)
}

// Detailed status message for the VPN tunnel.
func (o LookupVpnTunnelResultOutput) DetailedStatus() pulumi.StringOutput {
	return o.ApplyT(func(v LookupVpnTunnelResult) string { return v.DetailedStatus }).(pulumi.StringOutput)
}

// IKE protocol version to use when establishing the VPN tunnel with the peer VPN gateway. Acceptable IKE versions are 1 or 2. The default version is 2.
func (o LookupVpnTunnelResultOutput) IkeVersion() pulumi.IntOutput {
	return o.ApplyT(func(v LookupVpnTunnelResult) int { return v.IkeVersion }).(pulumi.IntOutput)
}

// Type of resource. Always compute#vpnTunnel for VPN tunnels.
func (o LookupVpnTunnelResultOutput) Kind() pulumi.StringOutput {
	return o.ApplyT(func(v LookupVpnTunnelResult) string { return v.Kind }).(pulumi.StringOutput)
}

// A fingerprint for the labels being applied to this VpnTunnel, which is essentially a hash of the labels set used for optimistic locking. The fingerprint is initially generated by Compute Engine and changes after every request to modify or update labels. You must always provide an up-to-date fingerprint hash in order to update or change labels, otherwise the request will fail with error 412 conditionNotMet. To see the latest fingerprint, make a get() request to retrieve a VpnTunnel.
func (o LookupVpnTunnelResultOutput) LabelFingerprint() pulumi.StringOutput {
	return o.ApplyT(func(v LookupVpnTunnelResult) string { return v.LabelFingerprint }).(pulumi.StringOutput)
}

// Labels for this resource. These can only be added or modified by the setLabels method. Each label key/value pair must comply with RFC1035. Label values may be empty.
func (o LookupVpnTunnelResultOutput) Labels() pulumi.StringMapOutput {
	return o.ApplyT(func(v LookupVpnTunnelResult) map[string]string { return v.Labels }).(pulumi.StringMapOutput)
}

// Local traffic selector to use when establishing the VPN tunnel with the peer VPN gateway. The value should be a CIDR formatted string, for example: 192.168.0.0/16. The ranges must be disjoint. Only IPv4 is supported.
func (o LookupVpnTunnelResultOutput) LocalTrafficSelector() pulumi.StringArrayOutput {
	return o.ApplyT(func(v LookupVpnTunnelResult) []string { return v.LocalTrafficSelector }).(pulumi.StringArrayOutput)
}

// Name of the resource. Provided by the client when the resource is created. The name must be 1-63 characters long, and comply with RFC1035. Specifically, the name must be 1-63 characters long and match the regular expression `[a-z]([-a-z0-9]*[a-z0-9])?` which means the first character must be a lowercase letter, and all following characters must be a dash, lowercase letter, or digit, except the last character, which cannot be a dash.
func (o LookupVpnTunnelResultOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v LookupVpnTunnelResult) string { return v.Name }).(pulumi.StringOutput)
}

// URL of the peer side external VPN gateway to which this VPN tunnel is connected. Provided by the client when the VPN tunnel is created. This field is exclusive with the field peerGcpGateway.
func (o LookupVpnTunnelResultOutput) PeerExternalGateway() pulumi.StringOutput {
	return o.ApplyT(func(v LookupVpnTunnelResult) string { return v.PeerExternalGateway }).(pulumi.StringOutput)
}

// The interface ID of the external VPN gateway to which this VPN tunnel is connected. Provided by the client when the VPN tunnel is created.
func (o LookupVpnTunnelResultOutput) PeerExternalGatewayInterface() pulumi.IntOutput {
	return o.ApplyT(func(v LookupVpnTunnelResult) int { return v.PeerExternalGatewayInterface }).(pulumi.IntOutput)
}

// URL of the peer side HA GCP VPN gateway to which this VPN tunnel is connected. Provided by the client when the VPN tunnel is created. This field can be used when creating highly available VPN from VPC network to VPC network, the field is exclusive with the field peerExternalGateway. If provided, the VPN tunnel will automatically use the same vpnGatewayInterface ID in the peer GCP VPN gateway.
func (o LookupVpnTunnelResultOutput) PeerGcpGateway() pulumi.StringOutput {
	return o.ApplyT(func(v LookupVpnTunnelResult) string { return v.PeerGcpGateway }).(pulumi.StringOutput)
}

// IP address of the peer VPN gateway. Only IPv4 is supported.
func (o LookupVpnTunnelResultOutput) PeerIp() pulumi.StringOutput {
	return o.ApplyT(func(v LookupVpnTunnelResult) string { return v.PeerIp }).(pulumi.StringOutput)
}

// URL of the region where the VPN tunnel resides. You must specify this field as part of the HTTP request URL. It is not settable as a field in the request body.
func (o LookupVpnTunnelResultOutput) Region() pulumi.StringOutput {
	return o.ApplyT(func(v LookupVpnTunnelResult) string { return v.Region }).(pulumi.StringOutput)
}

// Remote traffic selectors to use when establishing the VPN tunnel with the peer VPN gateway. The value should be a CIDR formatted string, for example: 192.168.0.0/16. The ranges should be disjoint. Only IPv4 is supported.
func (o LookupVpnTunnelResultOutput) RemoteTrafficSelector() pulumi.StringArrayOutput {
	return o.ApplyT(func(v LookupVpnTunnelResult) []string { return v.RemoteTrafficSelector }).(pulumi.StringArrayOutput)
}

// URL of the router resource to be used for dynamic routing.
func (o LookupVpnTunnelResultOutput) Router() pulumi.StringOutput {
	return o.ApplyT(func(v LookupVpnTunnelResult) string { return v.Router }).(pulumi.StringOutput)
}

// Server-defined URL for the resource.
func (o LookupVpnTunnelResultOutput) SelfLink() pulumi.StringOutput {
	return o.ApplyT(func(v LookupVpnTunnelResult) string { return v.SelfLink }).(pulumi.StringOutput)
}

// Shared secret used to set the secure session between the Cloud VPN gateway and the peer VPN gateway.
func (o LookupVpnTunnelResultOutput) SharedSecret() pulumi.StringOutput {
	return o.ApplyT(func(v LookupVpnTunnelResult) string { return v.SharedSecret }).(pulumi.StringOutput)
}

// Hash of the shared secret.
func (o LookupVpnTunnelResultOutput) SharedSecretHash() pulumi.StringOutput {
	return o.ApplyT(func(v LookupVpnTunnelResult) string { return v.SharedSecretHash }).(pulumi.StringOutput)
}

// The status of the VPN tunnel, which can be one of the following: - PROVISIONING: Resource is being allocated for the VPN tunnel. - WAITING_FOR_FULL_CONFIG: Waiting to receive all VPN-related configs from the user. Network, TargetVpnGateway, VpnTunnel, ForwardingRule, and Route resources are needed to setup the VPN tunnel. - FIRST_HANDSHAKE: Successful first handshake with the peer VPN. - ESTABLISHED: Secure session is successfully established with the peer VPN. - NETWORK_ERROR: Deprecated, replaced by NO_INCOMING_PACKETS - AUTHORIZATION_ERROR: Auth error (for example, bad shared secret). - NEGOTIATION_FAILURE: Handshake failed. - DEPROVISIONING: Resources are being deallocated for the VPN tunnel. - FAILED: Tunnel creation has failed and the tunnel is not ready to be used. - NO_INCOMING_PACKETS: No incoming packets from peer. - REJECTED: Tunnel configuration was rejected, can be result of being denied access. - ALLOCATING_RESOURCES: Cloud VPN is in the process of allocating all required resources. - STOPPED: Tunnel is stopped due to its Forwarding Rules being deleted for Classic VPN tunnels or the project is in frozen state. - PEER_IDENTITY_MISMATCH: Peer identity does not match peer IP, probably behind NAT. - TS_NARROWING_NOT_ALLOWED: Traffic selector narrowing not allowed for an HA-VPN tunnel.
func (o LookupVpnTunnelResultOutput) Status() pulumi.StringOutput {
	return o.ApplyT(func(v LookupVpnTunnelResult) string { return v.Status }).(pulumi.StringOutput)
}

// URL of the Target VPN gateway with which this VPN tunnel is associated. Provided by the client when the VPN tunnel is created.
func (o LookupVpnTunnelResultOutput) TargetVpnGateway() pulumi.StringOutput {
	return o.ApplyT(func(v LookupVpnTunnelResult) string { return v.TargetVpnGateway }).(pulumi.StringOutput)
}

// URL of the VPN gateway with which this VPN tunnel is associated. Provided by the client when the VPN tunnel is created. This must be used (instead of target_vpn_gateway) if a High Availability VPN gateway resource is created.
func (o LookupVpnTunnelResultOutput) VpnGateway() pulumi.StringOutput {
	return o.ApplyT(func(v LookupVpnTunnelResult) string { return v.VpnGateway }).(pulumi.StringOutput)
}

// The interface ID of the VPN gateway with which this VPN tunnel is associated.
func (o LookupVpnTunnelResultOutput) VpnGatewayInterface() pulumi.IntOutput {
	return o.ApplyT(func(v LookupVpnTunnelResult) int { return v.VpnGatewayInterface }).(pulumi.IntOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupVpnTunnelResultOutput{})
}
