// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package v1beta1

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Creates a new channel in the specified site.
type Channel struct {
	pulumi.CustomResourceState

	// The time at which the channel was created.
	CreateTime pulumi.StringOutput `pulumi:"createTime"`
	// The time at which the channel will be automatically deleted. If null, the channel will not be automatically deleted. This field is present in the output whether it's set directly or via the `ttl` field.
	ExpireTime pulumi.StringOutput `pulumi:"expireTime"`
	// Text labels used for extra metadata and/or filtering.
	Labels pulumi.StringMapOutput `pulumi:"labels"`
	// The fully-qualified resource name for the channel, in the format: sites/ SITE_ID/channels/CHANNEL_ID
	Name pulumi.StringOutput `pulumi:"name"`
	// The current release for the channel, if any.
	Release ReleaseResponseOutput `pulumi:"release"`
	// The number of previous releases to retain on the channel for rollback or other purposes. Must be a number between 1-100. Defaults to 10 for new channels.
	RetainedReleaseCount pulumi.IntOutput `pulumi:"retainedReleaseCount"`
	// Input only. A time-to-live for this channel. Sets `expire_time` to the provided duration past the time of the request.
	Ttl pulumi.StringOutput `pulumi:"ttl"`
	// The time at which the channel was last updated.
	UpdateTime pulumi.StringOutput `pulumi:"updateTime"`
	// The URL at which the content of this channel's current release can be viewed. This URL is a Firebase-provided subdomain of `web.app`. The content of this channel's current release can also be viewed at the Firebase-provided subdomain of `firebaseapp.com`. If this channel is the `live` channel for the Hosting site, then the content of this channel's current release can also be viewed at any connected custom domains.
	Url pulumi.StringOutput `pulumi:"url"`
}

// NewChannel registers a new resource with the given unique name, arguments, and options.
func NewChannel(ctx *pulumi.Context,
	name string, args *ChannelArgs, opts ...pulumi.ResourceOption) (*Channel, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.ChannelId == nil {
		return nil, errors.New("invalid value for required argument 'ChannelId'")
	}
	if args.SiteId == nil {
		return nil, errors.New("invalid value for required argument 'SiteId'")
	}
	var resource Channel
	err := ctx.RegisterResource("google-native:firebasehosting/v1beta1:Channel", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetChannel gets an existing Channel resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetChannel(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *ChannelState, opts ...pulumi.ResourceOption) (*Channel, error) {
	var resource Channel
	err := ctx.ReadResource("google-native:firebasehosting/v1beta1:Channel", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering Channel resources.
type channelState struct {
}

type ChannelState struct {
}

func (ChannelState) ElementType() reflect.Type {
	return reflect.TypeOf((*channelState)(nil)).Elem()
}

type channelArgs struct {
	ChannelId string `pulumi:"channelId"`
	// The time at which the channel will be automatically deleted. If null, the channel will not be automatically deleted. This field is present in the output whether it's set directly or via the `ttl` field.
	ExpireTime *string `pulumi:"expireTime"`
	// Text labels used for extra metadata and/or filtering.
	Labels map[string]string `pulumi:"labels"`
	// The fully-qualified resource name for the channel, in the format: sites/ SITE_ID/channels/CHANNEL_ID
	Name    *string `pulumi:"name"`
	Project *string `pulumi:"project"`
	// The number of previous releases to retain on the channel for rollback or other purposes. Must be a number between 1-100. Defaults to 10 for new channels.
	RetainedReleaseCount *int   `pulumi:"retainedReleaseCount"`
	SiteId               string `pulumi:"siteId"`
	// Input only. A time-to-live for this channel. Sets `expire_time` to the provided duration past the time of the request.
	Ttl *string `pulumi:"ttl"`
}

// The set of arguments for constructing a Channel resource.
type ChannelArgs struct {
	ChannelId pulumi.StringInput
	// The time at which the channel will be automatically deleted. If null, the channel will not be automatically deleted. This field is present in the output whether it's set directly or via the `ttl` field.
	ExpireTime pulumi.StringPtrInput
	// Text labels used for extra metadata and/or filtering.
	Labels pulumi.StringMapInput
	// The fully-qualified resource name for the channel, in the format: sites/ SITE_ID/channels/CHANNEL_ID
	Name    pulumi.StringPtrInput
	Project pulumi.StringPtrInput
	// The number of previous releases to retain on the channel for rollback or other purposes. Must be a number between 1-100. Defaults to 10 for new channels.
	RetainedReleaseCount pulumi.IntPtrInput
	SiteId               pulumi.StringInput
	// Input only. A time-to-live for this channel. Sets `expire_time` to the provided duration past the time of the request.
	Ttl pulumi.StringPtrInput
}

func (ChannelArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*channelArgs)(nil)).Elem()
}

type ChannelInput interface {
	pulumi.Input

	ToChannelOutput() ChannelOutput
	ToChannelOutputWithContext(ctx context.Context) ChannelOutput
}

func (*Channel) ElementType() reflect.Type {
	return reflect.TypeOf((**Channel)(nil)).Elem()
}

func (i *Channel) ToChannelOutput() ChannelOutput {
	return i.ToChannelOutputWithContext(context.Background())
}

func (i *Channel) ToChannelOutputWithContext(ctx context.Context) ChannelOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ChannelOutput)
}

type ChannelOutput struct{ *pulumi.OutputState }

func (ChannelOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**Channel)(nil)).Elem()
}

func (o ChannelOutput) ToChannelOutput() ChannelOutput {
	return o
}

func (o ChannelOutput) ToChannelOutputWithContext(ctx context.Context) ChannelOutput {
	return o
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*ChannelInput)(nil)).Elem(), &Channel{})
	pulumi.RegisterOutputType(ChannelOutput{})
}
