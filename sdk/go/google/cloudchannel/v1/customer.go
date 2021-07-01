// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package v1

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Creates a new Customer resource under the reseller or distributor account. Possible error codes: * PERMISSION_DENIED: The reseller account making the request is different from the reseller account in the API request. * INVALID_ARGUMENT: * Required request parameters are missing or invalid. * Domain field value doesn't match the primary email domain. Return value: The newly created Customer resource.
type Customer struct {
	pulumi.CustomResourceState

	// Secondary contact email. Alternate email and primary contact email are required to have different domains if primary contact email is present. When creating admin.google.com accounts, users get notified credentials at this email. This email address is also used as a recovery email.
	AlternateEmail pulumi.StringOutput `pulumi:"alternateEmail"`
	// Cloud Identity ID of the customer's channel partner. Populated only if a channel partner exists for this customer.
	ChannelPartnerId pulumi.StringOutput `pulumi:"channelPartnerId"`
	// Customer's cloud_identity_id. Populated only if a Cloud Identity resource exists for this customer.
	CloudIdentityId pulumi.StringOutput `pulumi:"cloudIdentityId"`
	// Cloud Identity information for the customer. Populated only if a Cloud Identity account exists for this customer.
	CloudIdentityInfo GoogleCloudChannelV1CloudIdentityInfoResponseOutput `pulumi:"cloudIdentityInfo"`
	// The time at which the customer is created.
	CreateTime pulumi.StringOutput `pulumi:"createTime"`
	// Required. Primary domain used by the customer. Domain of primary contact email is required to be same as the provided domain.
	Domain pulumi.StringOutput `pulumi:"domain"`
	// Optional. The BCP-47 language code, such as "en-US" or "sr-Latn". For more information, see https://www.unicode.org/reports/tr35/#Unicode_locale_identifier.
	LanguageCode pulumi.StringOutput `pulumi:"languageCode"`
	// Resource name of the customer. Format: accounts/{account_id}/customers/{customer_id}
	Name pulumi.StringOutput `pulumi:"name"`
	// Required. Name of the organization that the customer entity represents.
	OrgDisplayName pulumi.StringOutput `pulumi:"orgDisplayName"`
	// Required. Address of the organization of the customer entity. Region and zip codes are required to enforce US laws and embargoes. Valid address lines are required for all customers. Language code is discarded. Use the Customer-level language code to set the customer's language.
	OrgPostalAddress GoogleTypePostalAddressResponseOutput `pulumi:"orgPostalAddress"`
	// Primary contact info.
	PrimaryContactInfo GoogleCloudChannelV1ContactInfoResponseOutput `pulumi:"primaryContactInfo"`
	// The time at which the customer is updated.
	UpdateTime pulumi.StringOutput `pulumi:"updateTime"`
}

// NewCustomer registers a new resource with the given unique name, arguments, and options.
func NewCustomer(ctx *pulumi.Context,
	name string, args *CustomerArgs, opts ...pulumi.ResourceOption) (*Customer, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.AccountId == nil {
		return nil, errors.New("invalid value for required argument 'AccountId'")
	}
	if args.ChannelPartnerLinkId == nil {
		return nil, errors.New("invalid value for required argument 'ChannelPartnerLinkId'")
	}
	var resource Customer
	err := ctx.RegisterResource("google-native:cloudchannel/v1:Customer", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetCustomer gets an existing Customer resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetCustomer(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *CustomerState, opts ...pulumi.ResourceOption) (*Customer, error) {
	var resource Customer
	err := ctx.ReadResource("google-native:cloudchannel/v1:Customer", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering Customer resources.
type customerState struct {
}

type CustomerState struct {
}

func (CustomerState) ElementType() reflect.Type {
	return reflect.TypeOf((*customerState)(nil)).Elem()
}

type customerArgs struct {
	AccountId string `pulumi:"accountId"`
	// Secondary contact email. Alternate email and primary contact email are required to have different domains if primary contact email is present. When creating admin.google.com accounts, users get notified credentials at this email. This email address is also used as a recovery email.
	AlternateEmail *string `pulumi:"alternateEmail"`
	// Cloud Identity ID of the customer's channel partner. Populated only if a channel partner exists for this customer.
	ChannelPartnerId     *string `pulumi:"channelPartnerId"`
	ChannelPartnerLinkId string  `pulumi:"channelPartnerLinkId"`
	// Required. Primary domain used by the customer. Domain of primary contact email is required to be same as the provided domain.
	Domain *string `pulumi:"domain"`
	// Optional. The BCP-47 language code, such as "en-US" or "sr-Latn". For more information, see https://www.unicode.org/reports/tr35/#Unicode_locale_identifier.
	LanguageCode *string `pulumi:"languageCode"`
	// Required. Name of the organization that the customer entity represents.
	OrgDisplayName *string `pulumi:"orgDisplayName"`
	// Required. Address of the organization of the customer entity. Region and zip codes are required to enforce US laws and embargoes. Valid address lines are required for all customers. Language code is discarded. Use the Customer-level language code to set the customer's language.
	OrgPostalAddress *GoogleTypePostalAddress `pulumi:"orgPostalAddress"`
	// Primary contact info.
	PrimaryContactInfo *GoogleCloudChannelV1ContactInfo `pulumi:"primaryContactInfo"`
}

// The set of arguments for constructing a Customer resource.
type CustomerArgs struct {
	AccountId pulumi.StringInput
	// Secondary contact email. Alternate email and primary contact email are required to have different domains if primary contact email is present. When creating admin.google.com accounts, users get notified credentials at this email. This email address is also used as a recovery email.
	AlternateEmail pulumi.StringPtrInput
	// Cloud Identity ID of the customer's channel partner. Populated only if a channel partner exists for this customer.
	ChannelPartnerId     pulumi.StringPtrInput
	ChannelPartnerLinkId pulumi.StringInput
	// Required. Primary domain used by the customer. Domain of primary contact email is required to be same as the provided domain.
	Domain pulumi.StringPtrInput
	// Optional. The BCP-47 language code, such as "en-US" or "sr-Latn". For more information, see https://www.unicode.org/reports/tr35/#Unicode_locale_identifier.
	LanguageCode pulumi.StringPtrInput
	// Required. Name of the organization that the customer entity represents.
	OrgDisplayName pulumi.StringPtrInput
	// Required. Address of the organization of the customer entity. Region and zip codes are required to enforce US laws and embargoes. Valid address lines are required for all customers. Language code is discarded. Use the Customer-level language code to set the customer's language.
	OrgPostalAddress GoogleTypePostalAddressPtrInput
	// Primary contact info.
	PrimaryContactInfo GoogleCloudChannelV1ContactInfoPtrInput
}

func (CustomerArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*customerArgs)(nil)).Elem()
}

type CustomerInput interface {
	pulumi.Input

	ToCustomerOutput() CustomerOutput
	ToCustomerOutputWithContext(ctx context.Context) CustomerOutput
}

func (*Customer) ElementType() reflect.Type {
	return reflect.TypeOf((*Customer)(nil))
}

func (i *Customer) ToCustomerOutput() CustomerOutput {
	return i.ToCustomerOutputWithContext(context.Background())
}

func (i *Customer) ToCustomerOutputWithContext(ctx context.Context) CustomerOutput {
	return pulumi.ToOutputWithContext(ctx, i).(CustomerOutput)
}

type CustomerOutput struct {
	*pulumi.OutputState
}

func (CustomerOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*Customer)(nil))
}

func (o CustomerOutput) ToCustomerOutput() CustomerOutput {
	return o
}

func (o CustomerOutput) ToCustomerOutputWithContext(ctx context.Context) CustomerOutput {
	return o
}

func init() {
	pulumi.RegisterOutputType(CustomerOutput{})
}
