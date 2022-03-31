// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package v1

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Returns the requested ChannelPartnerLink resource. You must be a distributor to call this method. Possible error codes: * PERMISSION_DENIED: The reseller account making the request is different from the reseller account in the API request. * INVALID_ARGUMENT: Required request parameters are missing or invalid. * NOT_FOUND: ChannelPartnerLink resource not found because of an invalid channel partner link name. Return value: The ChannelPartnerLink resource.
func LookupChannelPartnerLink(ctx *pulumi.Context, args *LookupChannelPartnerLinkArgs, opts ...pulumi.InvokeOption) (*LookupChannelPartnerLinkResult, error) {
	var rv LookupChannelPartnerLinkResult
	err := ctx.Invoke("google-native:cloudchannel/v1:getChannelPartnerLink", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupChannelPartnerLinkArgs struct {
	AccountId            string  `pulumi:"accountId"`
	ChannelPartnerLinkId string  `pulumi:"channelPartnerLinkId"`
	View                 *string `pulumi:"view"`
}

type LookupChannelPartnerLinkResult struct {
	// Cloud Identity info of the channel partner (IR).
	ChannelPartnerCloudIdentityInfo GoogleCloudChannelV1CloudIdentityInfoResponse `pulumi:"channelPartnerCloudIdentityInfo"`
	// Timestamp of when the channel partner link is created.
	CreateTime string `pulumi:"createTime"`
	// URI of the web page where partner accepts the link invitation.
	InviteLinkUri string `pulumi:"inviteLinkUri"`
	// State of the channel partner link.
	LinkState string `pulumi:"linkState"`
	// Resource name for the channel partner link, in the format accounts/{account_id}/channelPartnerLinks/{id}.
	Name string `pulumi:"name"`
	// Public identifier that a customer must use to generate a transfer token to move to this distributor-reseller combination.
	PublicId string `pulumi:"publicId"`
	// Cloud Identity ID of the linked reseller.
	ResellerCloudIdentityId string `pulumi:"resellerCloudIdentityId"`
	// Timestamp of when the channel partner link is updated.
	UpdateTime string `pulumi:"updateTime"`
}

func LookupChannelPartnerLinkOutput(ctx *pulumi.Context, args LookupChannelPartnerLinkOutputArgs, opts ...pulumi.InvokeOption) LookupChannelPartnerLinkResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupChannelPartnerLinkResult, error) {
			args := v.(LookupChannelPartnerLinkArgs)
			r, err := LookupChannelPartnerLink(ctx, &args, opts...)
			return *r, err
		}).(LookupChannelPartnerLinkResultOutput)
}

type LookupChannelPartnerLinkOutputArgs struct {
	AccountId            pulumi.StringInput    `pulumi:"accountId"`
	ChannelPartnerLinkId pulumi.StringInput    `pulumi:"channelPartnerLinkId"`
	View                 pulumi.StringPtrInput `pulumi:"view"`
}

func (LookupChannelPartnerLinkOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupChannelPartnerLinkArgs)(nil)).Elem()
}

type LookupChannelPartnerLinkResultOutput struct{ *pulumi.OutputState }

func (LookupChannelPartnerLinkResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupChannelPartnerLinkResult)(nil)).Elem()
}

func (o LookupChannelPartnerLinkResultOutput) ToLookupChannelPartnerLinkResultOutput() LookupChannelPartnerLinkResultOutput {
	return o
}

func (o LookupChannelPartnerLinkResultOutput) ToLookupChannelPartnerLinkResultOutputWithContext(ctx context.Context) LookupChannelPartnerLinkResultOutput {
	return o
}

// Cloud Identity info of the channel partner (IR).
func (o LookupChannelPartnerLinkResultOutput) ChannelPartnerCloudIdentityInfo() GoogleCloudChannelV1CloudIdentityInfoResponseOutput {
	return o.ApplyT(func(v LookupChannelPartnerLinkResult) GoogleCloudChannelV1CloudIdentityInfoResponse {
		return v.ChannelPartnerCloudIdentityInfo
	}).(GoogleCloudChannelV1CloudIdentityInfoResponseOutput)
}

// Timestamp of when the channel partner link is created.
func (o LookupChannelPartnerLinkResultOutput) CreateTime() pulumi.StringOutput {
	return o.ApplyT(func(v LookupChannelPartnerLinkResult) string { return v.CreateTime }).(pulumi.StringOutput)
}

// URI of the web page where partner accepts the link invitation.
func (o LookupChannelPartnerLinkResultOutput) InviteLinkUri() pulumi.StringOutput {
	return o.ApplyT(func(v LookupChannelPartnerLinkResult) string { return v.InviteLinkUri }).(pulumi.StringOutput)
}

// State of the channel partner link.
func (o LookupChannelPartnerLinkResultOutput) LinkState() pulumi.StringOutput {
	return o.ApplyT(func(v LookupChannelPartnerLinkResult) string { return v.LinkState }).(pulumi.StringOutput)
}

// Resource name for the channel partner link, in the format accounts/{account_id}/channelPartnerLinks/{id}.
func (o LookupChannelPartnerLinkResultOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v LookupChannelPartnerLinkResult) string { return v.Name }).(pulumi.StringOutput)
}

// Public identifier that a customer must use to generate a transfer token to move to this distributor-reseller combination.
func (o LookupChannelPartnerLinkResultOutput) PublicId() pulumi.StringOutput {
	return o.ApplyT(func(v LookupChannelPartnerLinkResult) string { return v.PublicId }).(pulumi.StringOutput)
}

// Cloud Identity ID of the linked reseller.
func (o LookupChannelPartnerLinkResultOutput) ResellerCloudIdentityId() pulumi.StringOutput {
	return o.ApplyT(func(v LookupChannelPartnerLinkResult) string { return v.ResellerCloudIdentityId }).(pulumi.StringOutput)
}

// Timestamp of when the channel partner link is updated.
func (o LookupChannelPartnerLinkResultOutput) UpdateTime() pulumi.StringOutput {
	return o.ApplyT(func(v LookupChannelPartnerLinkResult) string { return v.UpdateTime }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupChannelPartnerLinkResultOutput{})
}
