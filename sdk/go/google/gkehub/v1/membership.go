// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package v1

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Creates a new Membership. **This is currently only supported for GKE clusters on Google Cloud**. To register other clusters, follow the instructions at https://cloud.google.com/anthos/multicluster-management/connect/registering-a-cluster.
// Auto-naming is currently not supported for this resource.
type Membership struct {
	pulumi.CustomResourceState

	// Optional. How to identify workloads from this Membership. See the documentation on Workload Identity for more details: https://cloud.google.com/kubernetes-engine/docs/how-to/workload-identity
	Authority AuthorityResponseOutput `pulumi:"authority"`
	// When the Membership was created.
	CreateTime pulumi.StringOutput `pulumi:"createTime"`
	// When the Membership was deleted.
	DeleteTime pulumi.StringOutput `pulumi:"deleteTime"`
	// Description of this membership, limited to 63 characters. Must match the regex: `a-zA-Z0-9*` This field is present for legacy purposes.
	Description pulumi.StringOutput `pulumi:"description"`
	// Optional. Endpoint information to reach this member.
	Endpoint MembershipEndpointResponseOutput `pulumi:"endpoint"`
	// Optional. An externally-generated and managed ID for this Membership. This ID may be modified after creation, but this is not recommended. The ID must match the regex: `a-zA-Z0-9*` If this Membership represents a Kubernetes cluster, this value should be set to the UID of the `kube-system` namespace object.
	ExternalId pulumi.StringOutput `pulumi:"externalId"`
	// Optional. GCP labels for this membership.
	Labels pulumi.StringMapOutput `pulumi:"labels"`
	// For clusters using Connect, the timestamp of the most recent connection established with Google Cloud. This time is updated every several minutes, not continuously. For clusters that do not use GKE Connect, or that have never connected successfully, this field will be unset.
	LastConnectionTime pulumi.StringOutput `pulumi:"lastConnectionTime"`
	// The full, unique name of this Membership resource in the format `projects/*/locations/*/memberships/{membership_id}`, set during creation. `membership_id` must be a valid RFC 1123 compliant DNS label: 1. At most 63 characters in length 2. It must consist of lower case alphanumeric characters or `-` 3. It must start and end with an alphanumeric character Which can be expressed as the regex: `[a-z0-9]([-a-z0-9]*[a-z0-9])?`, with a maximum length of 63 characters.
	Name pulumi.StringOutput `pulumi:"name"`
	// State of the Membership resource.
	State MembershipStateResponseOutput `pulumi:"state"`
	// Google-generated UUID for this resource. This is unique across all Membership resources. If a Membership resource is deleted and another resource with the same name is created, it gets a different unique_id.
	UniqueId pulumi.StringOutput `pulumi:"uniqueId"`
	// When the Membership was last updated.
	UpdateTime pulumi.StringOutput `pulumi:"updateTime"`
}

// NewMembership registers a new resource with the given unique name, arguments, and options.
func NewMembership(ctx *pulumi.Context,
	name string, args *MembershipArgs, opts ...pulumi.ResourceOption) (*Membership, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.MembershipId == nil {
		return nil, errors.New("invalid value for required argument 'MembershipId'")
	}
	var resource Membership
	err := ctx.RegisterResource("google-native:gkehub/v1:Membership", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetMembership gets an existing Membership resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetMembership(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *MembershipState, opts ...pulumi.ResourceOption) (*Membership, error) {
	var resource Membership
	err := ctx.ReadResource("google-native:gkehub/v1:Membership", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering Membership resources.
type membershipState struct {
}

type MembershipState struct {
}

func (MembershipState) ElementType() reflect.Type {
	return reflect.TypeOf((*membershipState)(nil)).Elem()
}

type membershipArgs struct {
	// Optional. How to identify workloads from this Membership. See the documentation on Workload Identity for more details: https://cloud.google.com/kubernetes-engine/docs/how-to/workload-identity
	Authority *Authority `pulumi:"authority"`
	// Optional. Endpoint information to reach this member.
	Endpoint *MembershipEndpoint `pulumi:"endpoint"`
	// Optional. An externally-generated and managed ID for this Membership. This ID may be modified after creation, but this is not recommended. The ID must match the regex: `a-zA-Z0-9*` If this Membership represents a Kubernetes cluster, this value should be set to the UID of the `kube-system` namespace object.
	ExternalId *string `pulumi:"externalId"`
	// Optional. GCP labels for this membership.
	Labels   map[string]string `pulumi:"labels"`
	Location *string           `pulumi:"location"`
	// Required. Client chosen ID for the membership. `membership_id` must be a valid RFC 1123 compliant DNS label: 1. At most 63 characters in length 2. It must consist of lower case alphanumeric characters or `-` 3. It must start and end with an alphanumeric character Which can be expressed as the regex: `[a-z0-9]([-a-z0-9]*[a-z0-9])?`, with a maximum length of 63 characters.
	MembershipId string  `pulumi:"membershipId"`
	Project      *string `pulumi:"project"`
	// Optional. A request ID to identify requests. Specify a unique request ID so that if you must retry your request, the server will know to ignore the request if it has already been completed. The server will guarantee that for at least 60 minutes after the first request. For example, consider a situation where you make an initial request and the request times out. If you make the request again with the same request ID, the server can check if original operation with the same request ID was received, and if so, will ignore the second request. This prevents clients from accidentally creating duplicate commitments. The request ID must be a valid UUID with the exception that zero UUID is not supported (00000000-0000-0000-0000-000000000000).
	RequestId *string `pulumi:"requestId"`
}

// The set of arguments for constructing a Membership resource.
type MembershipArgs struct {
	// Optional. How to identify workloads from this Membership. See the documentation on Workload Identity for more details: https://cloud.google.com/kubernetes-engine/docs/how-to/workload-identity
	Authority AuthorityPtrInput
	// Optional. Endpoint information to reach this member.
	Endpoint MembershipEndpointPtrInput
	// Optional. An externally-generated and managed ID for this Membership. This ID may be modified after creation, but this is not recommended. The ID must match the regex: `a-zA-Z0-9*` If this Membership represents a Kubernetes cluster, this value should be set to the UID of the `kube-system` namespace object.
	ExternalId pulumi.StringPtrInput
	// Optional. GCP labels for this membership.
	Labels   pulumi.StringMapInput
	Location pulumi.StringPtrInput
	// Required. Client chosen ID for the membership. `membership_id` must be a valid RFC 1123 compliant DNS label: 1. At most 63 characters in length 2. It must consist of lower case alphanumeric characters or `-` 3. It must start and end with an alphanumeric character Which can be expressed as the regex: `[a-z0-9]([-a-z0-9]*[a-z0-9])?`, with a maximum length of 63 characters.
	MembershipId pulumi.StringInput
	Project      pulumi.StringPtrInput
	// Optional. A request ID to identify requests. Specify a unique request ID so that if you must retry your request, the server will know to ignore the request if it has already been completed. The server will guarantee that for at least 60 minutes after the first request. For example, consider a situation where you make an initial request and the request times out. If you make the request again with the same request ID, the server can check if original operation with the same request ID was received, and if so, will ignore the second request. This prevents clients from accidentally creating duplicate commitments. The request ID must be a valid UUID with the exception that zero UUID is not supported (00000000-0000-0000-0000-000000000000).
	RequestId pulumi.StringPtrInput
}

func (MembershipArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*membershipArgs)(nil)).Elem()
}

type MembershipInput interface {
	pulumi.Input

	ToMembershipOutput() MembershipOutput
	ToMembershipOutputWithContext(ctx context.Context) MembershipOutput
}

func (*Membership) ElementType() reflect.Type {
	return reflect.TypeOf((**Membership)(nil)).Elem()
}

func (i *Membership) ToMembershipOutput() MembershipOutput {
	return i.ToMembershipOutputWithContext(context.Background())
}

func (i *Membership) ToMembershipOutputWithContext(ctx context.Context) MembershipOutput {
	return pulumi.ToOutputWithContext(ctx, i).(MembershipOutput)
}

type MembershipOutput struct{ *pulumi.OutputState }

func (MembershipOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**Membership)(nil)).Elem()
}

func (o MembershipOutput) ToMembershipOutput() MembershipOutput {
	return o
}

func (o MembershipOutput) ToMembershipOutputWithContext(ctx context.Context) MembershipOutput {
	return o
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*MembershipInput)(nil)).Elem(), &Membership{})
	pulumi.RegisterOutputType(MembershipOutput{})
}
