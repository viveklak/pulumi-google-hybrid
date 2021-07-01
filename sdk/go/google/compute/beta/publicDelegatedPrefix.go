// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package beta

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Creates a PublicDelegatedPrefix in the specified project in the given region using the parameters that are included in the request.
type PublicDelegatedPrefix struct {
	pulumi.CustomResourceState

	// Creation timestamp in RFC3339 text format.
	CreationTimestamp pulumi.StringOutput `pulumi:"creationTimestamp"`
	// An optional description of this resource. Provide this property when you create the resource.
	Description pulumi.StringOutput `pulumi:"description"`
	// Fingerprint of this resource. A hash of the contents stored in this object. This field is used in optimistic locking. This field will be ignored when inserting a new PublicDelegatedPrefix. An up-to-date fingerprint must be provided in order to update the PublicDelegatedPrefix, otherwise the request will fail with error 412 conditionNotMet.
	//
	// To see the latest fingerprint, make a get() request to retrieve a PublicDelegatedPrefix.
	Fingerprint pulumi.StringOutput `pulumi:"fingerprint"`
	// The IPv4 address range, in CIDR format, represented by this public delegated prefix.
	IpCidrRange pulumi.StringOutput `pulumi:"ipCidrRange"`
	// If true, the prefix will be live migrated.
	IsLiveMigration pulumi.BoolOutput `pulumi:"isLiveMigration"`
	// Type of the resource. Always compute#publicDelegatedPrefix for public delegated prefixes.
	Kind pulumi.StringOutput `pulumi:"kind"`
	// Name of the resource. Provided by the client when the resource is created. The name must be 1-63 characters long, and comply with RFC1035. Specifically, the name must be 1-63 characters long and match the regular expression `[a-z]([-a-z0-9]*[a-z0-9])?` which means the first character must be a lowercase letter, and all following characters must be a dash, lowercase letter, or digit, except the last character, which cannot be a dash.
	Name pulumi.StringOutput `pulumi:"name"`
	// The URL of parent prefix. Either PublicAdvertisedPrefix or PublicDelegatedPrefix.
	ParentPrefix pulumi.StringOutput `pulumi:"parentPrefix"`
	// The list of sub public delegated prefixes that exist for this public delegated prefix.
	PublicDelegatedSubPrefixs PublicDelegatedPrefixPublicDelegatedSubPrefixResponseArrayOutput `pulumi:"publicDelegatedSubPrefixs"`
	// URL of the region where the public delegated prefix resides. This field applies only to the region resource. You must specify this field as part of the HTTP request URL. It is not settable as a field in the request body.
	Region pulumi.StringOutput `pulumi:"region"`
	// Server-defined URL for the resource.
	SelfLink pulumi.StringOutput `pulumi:"selfLink"`
	// The status of the public delegated prefix.
	Status pulumi.StringOutput `pulumi:"status"`
}

// NewPublicDelegatedPrefix registers a new resource with the given unique name, arguments, and options.
func NewPublicDelegatedPrefix(ctx *pulumi.Context,
	name string, args *PublicDelegatedPrefixArgs, opts ...pulumi.ResourceOption) (*PublicDelegatedPrefix, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.Project == nil {
		return nil, errors.New("invalid value for required argument 'Project'")
	}
	if args.Region == nil {
		return nil, errors.New("invalid value for required argument 'Region'")
	}
	var resource PublicDelegatedPrefix
	err := ctx.RegisterResource("google-native:compute/beta:PublicDelegatedPrefix", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetPublicDelegatedPrefix gets an existing PublicDelegatedPrefix resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetPublicDelegatedPrefix(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *PublicDelegatedPrefixState, opts ...pulumi.ResourceOption) (*PublicDelegatedPrefix, error) {
	var resource PublicDelegatedPrefix
	err := ctx.ReadResource("google-native:compute/beta:PublicDelegatedPrefix", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering PublicDelegatedPrefix resources.
type publicDelegatedPrefixState struct {
}

type PublicDelegatedPrefixState struct {
}

func (PublicDelegatedPrefixState) ElementType() reflect.Type {
	return reflect.TypeOf((*publicDelegatedPrefixState)(nil)).Elem()
}

type publicDelegatedPrefixArgs struct {
	// An optional description of this resource. Provide this property when you create the resource.
	Description *string `pulumi:"description"`
	// The IPv4 address range, in CIDR format, represented by this public delegated prefix.
	IpCidrRange *string `pulumi:"ipCidrRange"`
	// If true, the prefix will be live migrated.
	IsLiveMigration *bool `pulumi:"isLiveMigration"`
	// Name of the resource. Provided by the client when the resource is created. The name must be 1-63 characters long, and comply with RFC1035. Specifically, the name must be 1-63 characters long and match the regular expression `[a-z]([-a-z0-9]*[a-z0-9])?` which means the first character must be a lowercase letter, and all following characters must be a dash, lowercase letter, or digit, except the last character, which cannot be a dash.
	Name *string `pulumi:"name"`
	// The URL of parent prefix. Either PublicAdvertisedPrefix or PublicDelegatedPrefix.
	ParentPrefix *string `pulumi:"parentPrefix"`
	Project      string  `pulumi:"project"`
	// The list of sub public delegated prefixes that exist for this public delegated prefix.
	PublicDelegatedSubPrefixs []PublicDelegatedPrefixPublicDelegatedSubPrefix `pulumi:"publicDelegatedSubPrefixs"`
	Region                    string                                          `pulumi:"region"`
	RequestId                 *string                                         `pulumi:"requestId"`
}

// The set of arguments for constructing a PublicDelegatedPrefix resource.
type PublicDelegatedPrefixArgs struct {
	// An optional description of this resource. Provide this property when you create the resource.
	Description pulumi.StringPtrInput
	// The IPv4 address range, in CIDR format, represented by this public delegated prefix.
	IpCidrRange pulumi.StringPtrInput
	// If true, the prefix will be live migrated.
	IsLiveMigration pulumi.BoolPtrInput
	// Name of the resource. Provided by the client when the resource is created. The name must be 1-63 characters long, and comply with RFC1035. Specifically, the name must be 1-63 characters long and match the regular expression `[a-z]([-a-z0-9]*[a-z0-9])?` which means the first character must be a lowercase letter, and all following characters must be a dash, lowercase letter, or digit, except the last character, which cannot be a dash.
	Name pulumi.StringPtrInput
	// The URL of parent prefix. Either PublicAdvertisedPrefix or PublicDelegatedPrefix.
	ParentPrefix pulumi.StringPtrInput
	Project      pulumi.StringInput
	// The list of sub public delegated prefixes that exist for this public delegated prefix.
	PublicDelegatedSubPrefixs PublicDelegatedPrefixPublicDelegatedSubPrefixArrayInput
	Region                    pulumi.StringInput
	RequestId                 pulumi.StringPtrInput
}

func (PublicDelegatedPrefixArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*publicDelegatedPrefixArgs)(nil)).Elem()
}

type PublicDelegatedPrefixInput interface {
	pulumi.Input

	ToPublicDelegatedPrefixOutput() PublicDelegatedPrefixOutput
	ToPublicDelegatedPrefixOutputWithContext(ctx context.Context) PublicDelegatedPrefixOutput
}

func (*PublicDelegatedPrefix) ElementType() reflect.Type {
	return reflect.TypeOf((*PublicDelegatedPrefix)(nil))
}

func (i *PublicDelegatedPrefix) ToPublicDelegatedPrefixOutput() PublicDelegatedPrefixOutput {
	return i.ToPublicDelegatedPrefixOutputWithContext(context.Background())
}

func (i *PublicDelegatedPrefix) ToPublicDelegatedPrefixOutputWithContext(ctx context.Context) PublicDelegatedPrefixOutput {
	return pulumi.ToOutputWithContext(ctx, i).(PublicDelegatedPrefixOutput)
}

type PublicDelegatedPrefixOutput struct {
	*pulumi.OutputState
}

func (PublicDelegatedPrefixOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*PublicDelegatedPrefix)(nil))
}

func (o PublicDelegatedPrefixOutput) ToPublicDelegatedPrefixOutput() PublicDelegatedPrefixOutput {
	return o
}

func (o PublicDelegatedPrefixOutput) ToPublicDelegatedPrefixOutputWithContext(ctx context.Context) PublicDelegatedPrefixOutput {
	return o
}

func init() {
	pulumi.RegisterOutputType(PublicDelegatedPrefixOutput{})
}
