// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package v1

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// This method creates [billing subaccounts](https://cloud.google.com/billing/docs/concepts#subaccounts). Google Cloud resellers should use the Channel Services APIs, [accounts.customers.create](https://cloud.google.com/channel/docs/reference/rest/v1/accounts.customers/create) and [accounts.customers.entitlements.create](https://cloud.google.com/channel/docs/reference/rest/v1/accounts.customers.entitlements/create). When creating a subaccount, the current authenticated user must have the `billing.accounts.update` IAM permission on the parent account, which is typically given to billing account [administrators](https://cloud.google.com/billing/docs/how-to/billing-access). This method will return an error if the parent account has not been provisioned as a reseller account.
type BillingAccount struct {
	pulumi.CustomResourceState

	// The display name given to the billing account, such as `My Billing Account`. This name is displayed in the Google Cloud Console.
	DisplayName pulumi.StringOutput `pulumi:"displayName"`
	// If this account is a [subaccount](https://cloud.google.com/billing/docs/concepts), then this will be the resource name of the parent billing account that it is being resold through. Otherwise this will be empty.
	MasterBillingAccount pulumi.StringOutput `pulumi:"masterBillingAccount"`
	// The resource name of the billing account. The resource name has the form `billingAccounts/{billing_account_id}`. For example, `billingAccounts/012345-567890-ABCDEF` would be the resource name for billing account `012345-567890-ABCDEF`.
	Name pulumi.StringOutput `pulumi:"name"`
	// True if the billing account is open, and will therefore be charged for any usage on associated projects. False if the billing account is closed, and therefore projects associated with it will be unable to use paid services.
	Open pulumi.BoolOutput `pulumi:"open"`
}

// NewBillingAccount registers a new resource with the given unique name, arguments, and options.
func NewBillingAccount(ctx *pulumi.Context,
	name string, args *BillingAccountArgs, opts ...pulumi.ResourceOption) (*BillingAccount, error) {
	if args == nil {
		args = &BillingAccountArgs{}
	}

	var resource BillingAccount
	err := ctx.RegisterResource("google-native:cloudbilling/v1:BillingAccount", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetBillingAccount gets an existing BillingAccount resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetBillingAccount(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *BillingAccountState, opts ...pulumi.ResourceOption) (*BillingAccount, error) {
	var resource BillingAccount
	err := ctx.ReadResource("google-native:cloudbilling/v1:BillingAccount", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering BillingAccount resources.
type billingAccountState struct {
	// The display name given to the billing account, such as `My Billing Account`. This name is displayed in the Google Cloud Console.
	DisplayName *string `pulumi:"displayName"`
	// If this account is a [subaccount](https://cloud.google.com/billing/docs/concepts), then this will be the resource name of the parent billing account that it is being resold through. Otherwise this will be empty.
	MasterBillingAccount *string `pulumi:"masterBillingAccount"`
	// The resource name of the billing account. The resource name has the form `billingAccounts/{billing_account_id}`. For example, `billingAccounts/012345-567890-ABCDEF` would be the resource name for billing account `012345-567890-ABCDEF`.
	Name *string `pulumi:"name"`
	// True if the billing account is open, and will therefore be charged for any usage on associated projects. False if the billing account is closed, and therefore projects associated with it will be unable to use paid services.
	Open *bool `pulumi:"open"`
}

type BillingAccountState struct {
	// The display name given to the billing account, such as `My Billing Account`. This name is displayed in the Google Cloud Console.
	DisplayName pulumi.StringPtrInput
	// If this account is a [subaccount](https://cloud.google.com/billing/docs/concepts), then this will be the resource name of the parent billing account that it is being resold through. Otherwise this will be empty.
	MasterBillingAccount pulumi.StringPtrInput
	// The resource name of the billing account. The resource name has the form `billingAccounts/{billing_account_id}`. For example, `billingAccounts/012345-567890-ABCDEF` would be the resource name for billing account `012345-567890-ABCDEF`.
	Name pulumi.StringPtrInput
	// True if the billing account is open, and will therefore be charged for any usage on associated projects. False if the billing account is closed, and therefore projects associated with it will be unable to use paid services.
	Open pulumi.BoolPtrInput
}

func (BillingAccountState) ElementType() reflect.Type {
	return reflect.TypeOf((*billingAccountState)(nil)).Elem()
}

type billingAccountArgs struct {
	// The display name given to the billing account, such as `My Billing Account`. This name is displayed in the Google Cloud Console.
	DisplayName *string `pulumi:"displayName"`
	// If this account is a [subaccount](https://cloud.google.com/billing/docs/concepts), then this will be the resource name of the parent billing account that it is being resold through. Otherwise this will be empty.
	MasterBillingAccount *string `pulumi:"masterBillingAccount"`
}

// The set of arguments for constructing a BillingAccount resource.
type BillingAccountArgs struct {
	// The display name given to the billing account, such as `My Billing Account`. This name is displayed in the Google Cloud Console.
	DisplayName pulumi.StringPtrInput
	// If this account is a [subaccount](https://cloud.google.com/billing/docs/concepts), then this will be the resource name of the parent billing account that it is being resold through. Otherwise this will be empty.
	MasterBillingAccount pulumi.StringPtrInput
}

func (BillingAccountArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*billingAccountArgs)(nil)).Elem()
}

type BillingAccountInput interface {
	pulumi.Input

	ToBillingAccountOutput() BillingAccountOutput
	ToBillingAccountOutputWithContext(ctx context.Context) BillingAccountOutput
}

func (*BillingAccount) ElementType() reflect.Type {
	return reflect.TypeOf((*BillingAccount)(nil))
}

func (i *BillingAccount) ToBillingAccountOutput() BillingAccountOutput {
	return i.ToBillingAccountOutputWithContext(context.Background())
}

func (i *BillingAccount) ToBillingAccountOutputWithContext(ctx context.Context) BillingAccountOutput {
	return pulumi.ToOutputWithContext(ctx, i).(BillingAccountOutput)
}

type BillingAccountOutput struct {
	*pulumi.OutputState
}

func (BillingAccountOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*BillingAccount)(nil))
}

func (o BillingAccountOutput) ToBillingAccountOutput() BillingAccountOutput {
	return o
}

func (o BillingAccountOutput) ToBillingAccountOutputWithContext(ctx context.Context) BillingAccountOutput {
	return o
}

func init() {
	pulumi.RegisterOutputType(BillingAccountOutput{})
}
