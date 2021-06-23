// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package v1

import (
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// View a notification configuration.
func LookupNotification(ctx *pulumi.Context, args *LookupNotificationArgs, opts ...pulumi.InvokeOption) (*LookupNotificationResult, error) {
	var rv LookupNotificationResult
	err := ctx.Invoke("google-native:storage/v1:getNotification", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupNotificationArgs struct {
	Bucket                 string  `pulumi:"bucket"`
	Notification           string  `pulumi:"notification"`
	ProvisionalUserProject *string `pulumi:"provisionalUserProject"`
	UserProject            *string `pulumi:"userProject"`
}

type LookupNotificationResult struct {
	// An optional list of additional attributes to attach to each Cloud PubSub message published for this notification subscription.
	Custom_attributes map[string]string `pulumi:"custom_attributes"`
	// HTTP 1.1 Entity tag for this subscription notification.
	Etag string `pulumi:"etag"`
	// If present, only send notifications about listed event types. If empty, sent notifications for all event types.
	Event_types []string `pulumi:"event_types"`
	// The kind of item this is. For notifications, this is always storage#notification.
	Kind string `pulumi:"kind"`
	// If present, only apply this notification configuration to object names that begin with this prefix.
	Object_name_prefix string `pulumi:"object_name_prefix"`
	// The desired content of the Payload.
	Payload_format string `pulumi:"payload_format"`
	// The canonical URL of this notification.
	SelfLink string `pulumi:"selfLink"`
	// The Cloud PubSub topic to which this subscription publishes. Formatted as: '//pubsub.googleapis.com/projects/{project-identifier}/topics/{my-topic}'
	Topic string `pulumi:"topic"`
}