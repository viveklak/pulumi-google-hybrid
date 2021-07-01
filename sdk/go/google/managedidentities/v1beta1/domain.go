// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package v1beta1

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Creates a Microsoft AD domain.
type Domain struct {
	pulumi.CustomResourceState

	// Optional. The name of delegated administrator account used to perform Active Directory operations. If not specified, `setupadmin` will be used.
	Admin pulumi.StringOutput `pulumi:"admin"`
	// Optional. Configuration for audit logs. True if audit logs are enabled, else false. Default is audit logs disabled.
	AuditLogsEnabled pulumi.BoolOutput `pulumi:"auditLogsEnabled"`
	// Optional. The full names of the Google Compute Engine [networks](/compute/docs/networks-and-firewalls#networks) the domain instance is connected to. Networks can be added using UpdateDomain. The domain is only available on networks listed in `authorized_networks`. If CIDR subnets overlap between networks, domain creation will fail.
	AuthorizedNetworks pulumi.StringArrayOutput `pulumi:"authorizedNetworks"`
	// The time the instance was created.
	CreateTime pulumi.StringOutput `pulumi:"createTime"`
	// The fully-qualified domain name of the exposed domain used by clients to connect to the service. Similar to what would be chosen for an Active Directory set up on an internal network.
	Fqdn pulumi.StringOutput `pulumi:"fqdn"`
	// Optional. Resource labels that can contain user-provided metadata.
	Labels pulumi.StringMapOutput `pulumi:"labels"`
	// Required. Locations where domain needs to be provisioned. regions e.g. us-west1 or us-east4 Service supports up to 4 locations at once. Each location will use a /26 block.
	Locations pulumi.StringArrayOutput `pulumi:"locations"`
	// The unique name of the domain using the form: `projects/{project_id}/locations/global/domains/{domain_name}`.
	Name pulumi.StringOutput `pulumi:"name"`
	// Required. The CIDR range of internal addresses that are reserved for this domain. Reserved networks must be /24 or larger. Ranges must be unique and non-overlapping with existing subnets in [Domain].[authorized_networks].
	ReservedIpRange pulumi.StringOutput `pulumi:"reservedIpRange"`
	// The current state of this domain.
	State pulumi.StringOutput `pulumi:"state"`
	// Additional information about the current status of this domain, if available.
	StatusMessage pulumi.StringOutput `pulumi:"statusMessage"`
	// The current trusts associated with the domain.
	Trusts TrustResponseArrayOutput `pulumi:"trusts"`
	// The last update time.
	UpdateTime pulumi.StringOutput `pulumi:"updateTime"`
}

// NewDomain registers a new resource with the given unique name, arguments, and options.
func NewDomain(ctx *pulumi.Context,
	name string, args *DomainArgs, opts ...pulumi.ResourceOption) (*Domain, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.DomainName == nil {
		return nil, errors.New("invalid value for required argument 'DomainName'")
	}
	if args.Project == nil {
		return nil, errors.New("invalid value for required argument 'Project'")
	}
	var resource Domain
	err := ctx.RegisterResource("google-native:managedidentities/v1beta1:Domain", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetDomain gets an existing Domain resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetDomain(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *DomainState, opts ...pulumi.ResourceOption) (*Domain, error) {
	var resource Domain
	err := ctx.ReadResource("google-native:managedidentities/v1beta1:Domain", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering Domain resources.
type domainState struct {
}

type DomainState struct {
}

func (DomainState) ElementType() reflect.Type {
	return reflect.TypeOf((*domainState)(nil)).Elem()
}

type domainArgs struct {
	// Optional. The name of delegated administrator account used to perform Active Directory operations. If not specified, `setupadmin` will be used.
	Admin *string `pulumi:"admin"`
	// Optional. Configuration for audit logs. True if audit logs are enabled, else false. Default is audit logs disabled.
	AuditLogsEnabled *bool `pulumi:"auditLogsEnabled"`
	// Optional. The full names of the Google Compute Engine [networks](/compute/docs/networks-and-firewalls#networks) the domain instance is connected to. Networks can be added using UpdateDomain. The domain is only available on networks listed in `authorized_networks`. If CIDR subnets overlap between networks, domain creation will fail.
	AuthorizedNetworks []string `pulumi:"authorizedNetworks"`
	DomainName         string   `pulumi:"domainName"`
	// Optional. Resource labels that can contain user-provided metadata.
	Labels map[string]string `pulumi:"labels"`
	// Required. Locations where domain needs to be provisioned. regions e.g. us-west1 or us-east4 Service supports up to 4 locations at once. Each location will use a /26 block.
	Locations []string `pulumi:"locations"`
	Project   string   `pulumi:"project"`
	// Required. The CIDR range of internal addresses that are reserved for this domain. Reserved networks must be /24 or larger. Ranges must be unique and non-overlapping with existing subnets in [Domain].[authorized_networks].
	ReservedIpRange *string `pulumi:"reservedIpRange"`
}

// The set of arguments for constructing a Domain resource.
type DomainArgs struct {
	// Optional. The name of delegated administrator account used to perform Active Directory operations. If not specified, `setupadmin` will be used.
	Admin pulumi.StringPtrInput
	// Optional. Configuration for audit logs. True if audit logs are enabled, else false. Default is audit logs disabled.
	AuditLogsEnabled pulumi.BoolPtrInput
	// Optional. The full names of the Google Compute Engine [networks](/compute/docs/networks-and-firewalls#networks) the domain instance is connected to. Networks can be added using UpdateDomain. The domain is only available on networks listed in `authorized_networks`. If CIDR subnets overlap between networks, domain creation will fail.
	AuthorizedNetworks pulumi.StringArrayInput
	DomainName         pulumi.StringInput
	// Optional. Resource labels that can contain user-provided metadata.
	Labels pulumi.StringMapInput
	// Required. Locations where domain needs to be provisioned. regions e.g. us-west1 or us-east4 Service supports up to 4 locations at once. Each location will use a /26 block.
	Locations pulumi.StringArrayInput
	Project   pulumi.StringInput
	// Required. The CIDR range of internal addresses that are reserved for this domain. Reserved networks must be /24 or larger. Ranges must be unique and non-overlapping with existing subnets in [Domain].[authorized_networks].
	ReservedIpRange pulumi.StringPtrInput
}

func (DomainArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*domainArgs)(nil)).Elem()
}

type DomainInput interface {
	pulumi.Input

	ToDomainOutput() DomainOutput
	ToDomainOutputWithContext(ctx context.Context) DomainOutput
}

func (*Domain) ElementType() reflect.Type {
	return reflect.TypeOf((*Domain)(nil))
}

func (i *Domain) ToDomainOutput() DomainOutput {
	return i.ToDomainOutputWithContext(context.Background())
}

func (i *Domain) ToDomainOutputWithContext(ctx context.Context) DomainOutput {
	return pulumi.ToOutputWithContext(ctx, i).(DomainOutput)
}

type DomainOutput struct {
	*pulumi.OutputState
}

func (DomainOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*Domain)(nil))
}

func (o DomainOutput) ToDomainOutput() DomainOutput {
	return o
}

func (o DomainOutput) ToDomainOutputWithContext(ctx context.Context) DomainOutput {
	return o
}

func init() {
	pulumi.RegisterOutputType(DomainOutput{})
}
