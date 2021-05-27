// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package v1alpha2

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Registers a new domain name and creates a corresponding `Registration` resource. Call `RetrieveRegisterParameters` first to check availability of the domain name and determine parameters like price that are needed to build a call to this method. A successful call creates a `Registration` resource in state `REGISTRATION_PENDING`, which resolves to `ACTIVE` within 1-2 minutes, indicating that the domain was successfully registered. If the resource ends up in state `REGISTRATION_FAILED`, it indicates that the domain was not registered successfully, and you can safely delete the resource and retry registration.
type Registration struct {
	pulumi.CustomResourceState

	// Required. Settings for contact information linked to the `Registration`. You cannot update these with the `UpdateRegistration` method. To update these settings, use the `ConfigureContactSettings` method.
	ContactSettings ContactSettingsResponseOutput `pulumi:"contactSettings"`
	// The creation timestamp of the `Registration` resource.
	CreateTime pulumi.StringOutput `pulumi:"createTime"`
	// Settings controlling the DNS configuration of the `Registration`. You cannot update these with the `UpdateRegistration` method. To update these settings, use the `ConfigureDnsSettings` method.
	DnsSettings DnsSettingsResponseOutput `pulumi:"dnsSettings"`
	// Required. Immutable. The domain name. Unicode domain names must be expressed in Punycode format.
	DomainName pulumi.StringOutput `pulumi:"domainName"`
	// The expiration timestamp of the `Registration`.
	ExpireTime pulumi.StringOutput `pulumi:"expireTime"`
	// The set of issues with the `Registration` that require attention.
	Issues pulumi.StringArrayOutput `pulumi:"issues"`
	// Set of labels associated with the `Registration`.
	Labels pulumi.StringMapOutput `pulumi:"labels"`
	// Settings for management of the `Registration`, including renewal, billing, and transfer. You cannot update these with the `UpdateRegistration` method. To update these settings, use the `ConfigureManagementSettings` method.
	ManagementSettings ManagementSettingsResponseOutput `pulumi:"managementSettings"`
	// Name of the `Registration` resource, in the format `projects/*/locations/*/registrations/`.
	Name pulumi.StringOutput `pulumi:"name"`
	// Pending contact settings for the `Registration`. Updates to the `contact_settings` field that change its `registrant_contact` or `privacy` fields require email confirmation by the `registrant_contact` before taking effect. This field is set only if there are pending updates to the `contact_settings` that have not yet been confirmed. To confirm the changes, the `registrant_contact` must follow the instructions in the email they receive.
	PendingContactSettings ContactSettingsResponseOutput `pulumi:"pendingContactSettings"`
	// The state of the `Registration`
	State pulumi.StringOutput `pulumi:"state"`
	// Set of options for the `contact_settings.privacy` field that this `Registration` supports.
	SupportedPrivacy pulumi.StringArrayOutput `pulumi:"supportedPrivacy"`
}

// NewRegistration registers a new resource with the given unique name, arguments, and options.
func NewRegistration(ctx *pulumi.Context,
	name string, args *RegistrationArgs, opts ...pulumi.ResourceOption) (*Registration, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.Location == nil {
		return nil, errors.New("invalid value for required argument 'Location'")
	}
	if args.Project == nil {
		return nil, errors.New("invalid value for required argument 'Project'")
	}
	var resource Registration
	err := ctx.RegisterResource("google-native:domains/v1alpha2:Registration", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetRegistration gets an existing Registration resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetRegistration(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *RegistrationState, opts ...pulumi.ResourceOption) (*Registration, error) {
	var resource Registration
	err := ctx.ReadResource("google-native:domains/v1alpha2:Registration", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering Registration resources.
type registrationState struct {
	// Required. Settings for contact information linked to the `Registration`. You cannot update these with the `UpdateRegistration` method. To update these settings, use the `ConfigureContactSettings` method.
	ContactSettings *ContactSettingsResponse `pulumi:"contactSettings"`
	// The creation timestamp of the `Registration` resource.
	CreateTime *string `pulumi:"createTime"`
	// Settings controlling the DNS configuration of the `Registration`. You cannot update these with the `UpdateRegistration` method. To update these settings, use the `ConfigureDnsSettings` method.
	DnsSettings *DnsSettingsResponse `pulumi:"dnsSettings"`
	// Required. Immutable. The domain name. Unicode domain names must be expressed in Punycode format.
	DomainName *string `pulumi:"domainName"`
	// The expiration timestamp of the `Registration`.
	ExpireTime *string `pulumi:"expireTime"`
	// The set of issues with the `Registration` that require attention.
	Issues []string `pulumi:"issues"`
	// Set of labels associated with the `Registration`.
	Labels map[string]string `pulumi:"labels"`
	// Settings for management of the `Registration`, including renewal, billing, and transfer. You cannot update these with the `UpdateRegistration` method. To update these settings, use the `ConfigureManagementSettings` method.
	ManagementSettings *ManagementSettingsResponse `pulumi:"managementSettings"`
	// Name of the `Registration` resource, in the format `projects/*/locations/*/registrations/`.
	Name *string `pulumi:"name"`
	// Pending contact settings for the `Registration`. Updates to the `contact_settings` field that change its `registrant_contact` or `privacy` fields require email confirmation by the `registrant_contact` before taking effect. This field is set only if there are pending updates to the `contact_settings` that have not yet been confirmed. To confirm the changes, the `registrant_contact` must follow the instructions in the email they receive.
	PendingContactSettings *ContactSettingsResponse `pulumi:"pendingContactSettings"`
	// The state of the `Registration`
	State *string `pulumi:"state"`
	// Set of options for the `contact_settings.privacy` field that this `Registration` supports.
	SupportedPrivacy []string `pulumi:"supportedPrivacy"`
}

type RegistrationState struct {
	// Required. Settings for contact information linked to the `Registration`. You cannot update these with the `UpdateRegistration` method. To update these settings, use the `ConfigureContactSettings` method.
	ContactSettings ContactSettingsResponsePtrInput
	// The creation timestamp of the `Registration` resource.
	CreateTime pulumi.StringPtrInput
	// Settings controlling the DNS configuration of the `Registration`. You cannot update these with the `UpdateRegistration` method. To update these settings, use the `ConfigureDnsSettings` method.
	DnsSettings DnsSettingsResponsePtrInput
	// Required. Immutable. The domain name. Unicode domain names must be expressed in Punycode format.
	DomainName pulumi.StringPtrInput
	// The expiration timestamp of the `Registration`.
	ExpireTime pulumi.StringPtrInput
	// The set of issues with the `Registration` that require attention.
	Issues pulumi.StringArrayInput
	// Set of labels associated with the `Registration`.
	Labels pulumi.StringMapInput
	// Settings for management of the `Registration`, including renewal, billing, and transfer. You cannot update these with the `UpdateRegistration` method. To update these settings, use the `ConfigureManagementSettings` method.
	ManagementSettings ManagementSettingsResponsePtrInput
	// Name of the `Registration` resource, in the format `projects/*/locations/*/registrations/`.
	Name pulumi.StringPtrInput
	// Pending contact settings for the `Registration`. Updates to the `contact_settings` field that change its `registrant_contact` or `privacy` fields require email confirmation by the `registrant_contact` before taking effect. This field is set only if there are pending updates to the `contact_settings` that have not yet been confirmed. To confirm the changes, the `registrant_contact` must follow the instructions in the email they receive.
	PendingContactSettings ContactSettingsResponsePtrInput
	// The state of the `Registration`
	State pulumi.StringPtrInput
	// Set of options for the `contact_settings.privacy` field that this `Registration` supports.
	SupportedPrivacy pulumi.StringArrayInput
}

func (RegistrationState) ElementType() reflect.Type {
	return reflect.TypeOf((*registrationState)(nil)).Elem()
}

type registrationArgs struct {
	// The list of contact notices that the caller acknowledges. The notices needed here depend on the values specified in `registration.contact_settings`.
	ContactNotices []string `pulumi:"contactNotices"`
	// Required. Settings for contact information linked to the `Registration`. You cannot update these with the `UpdateRegistration` method. To update these settings, use the `ConfigureContactSettings` method.
	ContactSettings *ContactSettings `pulumi:"contactSettings"`
	// Settings controlling the DNS configuration of the `Registration`. You cannot update these with the `UpdateRegistration` method. To update these settings, use the `ConfigureDnsSettings` method.
	DnsSettings *DnsSettings `pulumi:"dnsSettings"`
	// Required. Immutable. The domain name. Unicode domain names must be expressed in Punycode format.
	DomainName *string `pulumi:"domainName"`
	// The list of domain notices that you acknowledge. Call `RetrieveRegisterParameters` to see the notices that need acknowledgement.
	DomainNotices []string `pulumi:"domainNotices"`
	// Set of labels associated with the `Registration`.
	Labels   map[string]string `pulumi:"labels"`
	Location string            `pulumi:"location"`
	// Settings for management of the `Registration`, including renewal, billing, and transfer. You cannot update these with the `UpdateRegistration` method. To update these settings, use the `ConfigureManagementSettings` method.
	ManagementSettings *ManagementSettings `pulumi:"managementSettings"`
	Project            string              `pulumi:"project"`
	// When true, only validation will be performed, without actually registering the domain. Follows: https://cloud.google.com/apis/design/design_patterns#request_validation
	ValidateOnly *bool `pulumi:"validateOnly"`
	// Required. Yearly price to register or renew the domain. The value that should be put here can be obtained from RetrieveRegisterParameters or SearchDomains calls.
	YearlyPrice *Money `pulumi:"yearlyPrice"`
}

// The set of arguments for constructing a Registration resource.
type RegistrationArgs struct {
	// The list of contact notices that the caller acknowledges. The notices needed here depend on the values specified in `registration.contact_settings`.
	ContactNotices pulumi.StringArrayInput
	// Required. Settings for contact information linked to the `Registration`. You cannot update these with the `UpdateRegistration` method. To update these settings, use the `ConfigureContactSettings` method.
	ContactSettings ContactSettingsPtrInput
	// Settings controlling the DNS configuration of the `Registration`. You cannot update these with the `UpdateRegistration` method. To update these settings, use the `ConfigureDnsSettings` method.
	DnsSettings DnsSettingsPtrInput
	// Required. Immutable. The domain name. Unicode domain names must be expressed in Punycode format.
	DomainName pulumi.StringPtrInput
	// The list of domain notices that you acknowledge. Call `RetrieveRegisterParameters` to see the notices that need acknowledgement.
	DomainNotices pulumi.StringArrayInput
	// Set of labels associated with the `Registration`.
	Labels   pulumi.StringMapInput
	Location pulumi.StringInput
	// Settings for management of the `Registration`, including renewal, billing, and transfer. You cannot update these with the `UpdateRegistration` method. To update these settings, use the `ConfigureManagementSettings` method.
	ManagementSettings ManagementSettingsPtrInput
	Project            pulumi.StringInput
	// When true, only validation will be performed, without actually registering the domain. Follows: https://cloud.google.com/apis/design/design_patterns#request_validation
	ValidateOnly pulumi.BoolPtrInput
	// Required. Yearly price to register or renew the domain. The value that should be put here can be obtained from RetrieveRegisterParameters or SearchDomains calls.
	YearlyPrice MoneyPtrInput
}

func (RegistrationArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*registrationArgs)(nil)).Elem()
}

type RegistrationInput interface {
	pulumi.Input

	ToRegistrationOutput() RegistrationOutput
	ToRegistrationOutputWithContext(ctx context.Context) RegistrationOutput
}

func (*Registration) ElementType() reflect.Type {
	return reflect.TypeOf((*Registration)(nil))
}

func (i *Registration) ToRegistrationOutput() RegistrationOutput {
	return i.ToRegistrationOutputWithContext(context.Background())
}

func (i *Registration) ToRegistrationOutputWithContext(ctx context.Context) RegistrationOutput {
	return pulumi.ToOutputWithContext(ctx, i).(RegistrationOutput)
}

type RegistrationOutput struct {
	*pulumi.OutputState
}

func (RegistrationOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*Registration)(nil))
}

func (o RegistrationOutput) ToRegistrationOutput() RegistrationOutput {
	return o
}

func (o RegistrationOutput) ToRegistrationOutputWithContext(ctx context.Context) RegistrationOutput {
	return o
}

func init() {
	pulumi.RegisterOutputType(RegistrationOutput{})
}
