// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package v1

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Use this method to create a connection profile in a project and location.
// Auto-naming is currently not supported for this resource.
type ConnectionProfile struct {
	pulumi.CustomResourceState

	// The create time of the resource.
	CreateTime pulumi.StringOutput `pulumi:"createTime"`
	// Display name.
	DisplayName pulumi.StringOutput `pulumi:"displayName"`
	// Forward SSH tunnel connectivity.
	ForwardSshConnectivity ForwardSshTunnelConnectivityResponseOutput `pulumi:"forwardSshConnectivity"`
	// Cloud Storage ConnectionProfile configuration.
	GcsProfile GcsProfileResponseOutput `pulumi:"gcsProfile"`
	// Labels.
	Labels pulumi.StringMapOutput `pulumi:"labels"`
	// MySQL ConnectionProfile configuration.
	MysqlProfile MysqlProfileResponseOutput `pulumi:"mysqlProfile"`
	// The resource's name.
	Name pulumi.StringOutput `pulumi:"name"`
	// Oracle ConnectionProfile configuration.
	OracleProfile OracleProfileResponseOutput `pulumi:"oracleProfile"`
	// Private connectivity.
	PrivateConnectivity PrivateConnectivityResponseOutput `pulumi:"privateConnectivity"`
	// Static Service IP connectivity.
	StaticServiceIpConnectivity StaticServiceIpConnectivityResponseOutput `pulumi:"staticServiceIpConnectivity"`
	// The update time of the resource.
	UpdateTime pulumi.StringOutput `pulumi:"updateTime"`
}

// NewConnectionProfile registers a new resource with the given unique name, arguments, and options.
func NewConnectionProfile(ctx *pulumi.Context,
	name string, args *ConnectionProfileArgs, opts ...pulumi.ResourceOption) (*ConnectionProfile, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.ConnectionProfileId == nil {
		return nil, errors.New("invalid value for required argument 'ConnectionProfileId'")
	}
	if args.DisplayName == nil {
		return nil, errors.New("invalid value for required argument 'DisplayName'")
	}
	var resource ConnectionProfile
	err := ctx.RegisterResource("google-native:datastream/v1:ConnectionProfile", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetConnectionProfile gets an existing ConnectionProfile resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetConnectionProfile(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *ConnectionProfileState, opts ...pulumi.ResourceOption) (*ConnectionProfile, error) {
	var resource ConnectionProfile
	err := ctx.ReadResource("google-native:datastream/v1:ConnectionProfile", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering ConnectionProfile resources.
type connectionProfileState struct {
}

type ConnectionProfileState struct {
}

func (ConnectionProfileState) ElementType() reflect.Type {
	return reflect.TypeOf((*connectionProfileState)(nil)).Elem()
}

type connectionProfileArgs struct {
	ConnectionProfileId string `pulumi:"connectionProfileId"`
	// Display name.
	DisplayName string  `pulumi:"displayName"`
	Force       *string `pulumi:"force"`
	// Forward SSH tunnel connectivity.
	ForwardSshConnectivity *ForwardSshTunnelConnectivity `pulumi:"forwardSshConnectivity"`
	// Cloud Storage ConnectionProfile configuration.
	GcsProfile *GcsProfile `pulumi:"gcsProfile"`
	// Labels.
	Labels   map[string]string `pulumi:"labels"`
	Location *string           `pulumi:"location"`
	// MySQL ConnectionProfile configuration.
	MysqlProfile *MysqlProfile `pulumi:"mysqlProfile"`
	// Oracle ConnectionProfile configuration.
	OracleProfile *OracleProfile `pulumi:"oracleProfile"`
	// Private connectivity.
	PrivateConnectivity *PrivateConnectivity `pulumi:"privateConnectivity"`
	Project             *string              `pulumi:"project"`
	RequestId           *string              `pulumi:"requestId"`
	// Static Service IP connectivity.
	StaticServiceIpConnectivity *StaticServiceIpConnectivity `pulumi:"staticServiceIpConnectivity"`
}

// The set of arguments for constructing a ConnectionProfile resource.
type ConnectionProfileArgs struct {
	ConnectionProfileId pulumi.StringInput
	// Display name.
	DisplayName pulumi.StringInput
	Force       pulumi.StringPtrInput
	// Forward SSH tunnel connectivity.
	ForwardSshConnectivity ForwardSshTunnelConnectivityPtrInput
	// Cloud Storage ConnectionProfile configuration.
	GcsProfile GcsProfilePtrInput
	// Labels.
	Labels   pulumi.StringMapInput
	Location pulumi.StringPtrInput
	// MySQL ConnectionProfile configuration.
	MysqlProfile MysqlProfilePtrInput
	// Oracle ConnectionProfile configuration.
	OracleProfile OracleProfilePtrInput
	// Private connectivity.
	PrivateConnectivity PrivateConnectivityPtrInput
	Project             pulumi.StringPtrInput
	RequestId           pulumi.StringPtrInput
	// Static Service IP connectivity.
	StaticServiceIpConnectivity StaticServiceIpConnectivityPtrInput
}

func (ConnectionProfileArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*connectionProfileArgs)(nil)).Elem()
}

type ConnectionProfileInput interface {
	pulumi.Input

	ToConnectionProfileOutput() ConnectionProfileOutput
	ToConnectionProfileOutputWithContext(ctx context.Context) ConnectionProfileOutput
}

func (*ConnectionProfile) ElementType() reflect.Type {
	return reflect.TypeOf((*ConnectionProfile)(nil))
}

func (i *ConnectionProfile) ToConnectionProfileOutput() ConnectionProfileOutput {
	return i.ToConnectionProfileOutputWithContext(context.Background())
}

func (i *ConnectionProfile) ToConnectionProfileOutputWithContext(ctx context.Context) ConnectionProfileOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ConnectionProfileOutput)
}

type ConnectionProfileOutput struct{ *pulumi.OutputState }

func (ConnectionProfileOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*ConnectionProfile)(nil))
}

func (o ConnectionProfileOutput) ToConnectionProfileOutput() ConnectionProfileOutput {
	return o
}

func (o ConnectionProfileOutput) ToConnectionProfileOutputWithContext(ctx context.Context) ConnectionProfileOutput {
	return o
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*ConnectionProfileInput)(nil)).Elem(), &ConnectionProfile{})
	pulumi.RegisterOutputType(ConnectionProfileOutput{})
}