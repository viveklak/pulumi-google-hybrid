// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package v1

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Gets the configuration details of a subscription.
func LookupSubscription(ctx *pulumi.Context, args *LookupSubscriptionArgs, opts ...pulumi.InvokeOption) (*LookupSubscriptionResult, error) {
	var rv LookupSubscriptionResult
	err := ctx.Invoke("google-native:pubsub/v1:getSubscription", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupSubscriptionArgs struct {
	Project        *string `pulumi:"project"`
	SubscriptionId string  `pulumi:"subscriptionId"`
}

type LookupSubscriptionResult struct {
	// The approximate amount of time (on a best-effort basis) Pub/Sub waits for the subscriber to acknowledge receipt before resending the message. In the interval after the message is delivered and before it is acknowledged, it is considered to be *outstanding*. During that time period, the message will not be redelivered (on a best-effort basis). For pull subscriptions, this value is used as the initial value for the ack deadline. To override this value for a given message, call `ModifyAckDeadline` with the corresponding `ack_id` if using non-streaming pull or send the `ack_id` in a `StreamingModifyAckDeadlineRequest` if using streaming pull. The minimum custom deadline you can specify is 10 seconds. The maximum custom deadline you can specify is 600 seconds (10 minutes). If this parameter is 0, a default value of 10 seconds is used. For push delivery, this value is also used to set the request timeout for the call to the push endpoint. If the subscriber never acknowledges the message, the Pub/Sub system will eventually redeliver the message.
	AckDeadlineSeconds int `pulumi:"ackDeadlineSeconds"`
	// A policy that specifies the conditions for dead lettering messages in this subscription. If dead_letter_policy is not set, dead lettering is disabled. The Cloud Pub/Sub service account associated with this subscriptions's parent project (i.e., service-{project_number}@gcp-sa-pubsub.iam.gserviceaccount.com) must have permission to Acknowledge() messages on this subscription.
	DeadLetterPolicy DeadLetterPolicyResponse `pulumi:"deadLetterPolicy"`
	// Indicates whether the subscription is detached from its topic. Detached subscriptions don't receive messages from their topic and don't retain any backlog. `Pull` and `StreamingPull` requests will return FAILED_PRECONDITION. If the subscription is a push subscription, pushes to the endpoint will not be made.
	Detached bool `pulumi:"detached"`
	// If true, Pub/Sub provides the following guarantees for the delivery of a message with a given value of `message_id` on this subscription: * The message sent to a subscriber is guaranteed not to be resent before the message's acknowledgement deadline expires. * An acknowledged message will not be resent to a subscriber. Note that subscribers may still receive multiple copies of a message when `enable_exactly_once_delivery` is true if the message was published multiple times by a publisher client. These copies are considered distinct by Pub/Sub and have distinct `message_id` values.
	EnableExactlyOnceDelivery bool `pulumi:"enableExactlyOnceDelivery"`
	// If true, messages published with the same `ordering_key` in `PubsubMessage` will be delivered to the subscribers in the order in which they are received by the Pub/Sub system. Otherwise, they may be delivered in any order.
	EnableMessageOrdering bool `pulumi:"enableMessageOrdering"`
	// A policy that specifies the conditions for this subscription's expiration. A subscription is considered active as long as any connected subscriber is successfully consuming messages from the subscription or is issuing operations on the subscription. If `expiration_policy` is not set, a *default policy* with `ttl` of 31 days will be used. The minimum allowed value for `expiration_policy.ttl` is 1 day. If `expiration_policy` is set, but `expiration_policy.ttl` is not set, the subscription never expires.
	ExpirationPolicy ExpirationPolicyResponse `pulumi:"expirationPolicy"`
	// An expression written in the Pub/Sub [filter language](https://cloud.google.com/pubsub/docs/filtering). If non-empty, then only `PubsubMessage`s whose `attributes` field matches the filter are delivered on this subscription. If empty, then no messages are filtered out.
	Filter string `pulumi:"filter"`
	// See Creating and managing labels.
	Labels map[string]string `pulumi:"labels"`
	// How long to retain unacknowledged messages in the subscription's backlog, from the moment a message is published. If `retain_acked_messages` is true, then this also configures the retention of acknowledged messages, and thus configures how far back in time a `Seek` can be done. Defaults to 7 days. Cannot be more than 7 days or less than 10 minutes.
	MessageRetentionDuration string `pulumi:"messageRetentionDuration"`
	// The name of the subscription. It must have the format `"projects/{project}/subscriptions/{subscription}"`. `{subscription}` must start with a letter, and contain only letters (`[A-Za-z]`), numbers (`[0-9]`), dashes (`-`), underscores (`_`), periods (`.`), tildes (`~`), plus (`+`) or percent signs (`%`). It must be between 3 and 255 characters in length, and it must not start with `"goog"`.
	Name string `pulumi:"name"`
	// If push delivery is used with this subscription, this field is used to configure it. At most one of `pushConfig` and `bigQueryConfig` can be set. If both are empty, then the subscriber will pull and ack messages using API methods.
	PushConfig PushConfigResponse `pulumi:"pushConfig"`
	// Indicates whether to retain acknowledged messages. If true, then messages are not expunged from the subscription's backlog, even if they are acknowledged, until they fall out of the `message_retention_duration` window. This must be true if you would like to [`Seek` to a timestamp] (https://cloud.google.com/pubsub/docs/replay-overview#seek_to_a_time) in the past to replay previously-acknowledged messages.
	RetainAckedMessages bool `pulumi:"retainAckedMessages"`
	// A policy that specifies how Pub/Sub retries message delivery for this subscription. If not set, the default retry policy is applied. This generally implies that messages will be retried as soon as possible for healthy subscribers. RetryPolicy will be triggered on NACKs or acknowledgement deadline exceeded events for a given message.
	RetryPolicy RetryPolicyResponse `pulumi:"retryPolicy"`
	// An output-only field indicating whether or not the subscription can receive messages.
	State string `pulumi:"state"`
	// The name of the topic from which this subscription is receiving messages. Format is `projects/{project}/topics/{topic}`. The value of this field will be `_deleted-topic_` if the topic has been deleted.
	Topic string `pulumi:"topic"`
	// Indicates the minimum duration for which a message is retained after it is published to the subscription's topic. If this field is set, messages published to the subscription's topic in the last `topic_message_retention_duration` are always available to subscribers. See the `message_retention_duration` field in `Topic`. This field is set only in responses from the server; it is ignored if it is set in any requests.
	TopicMessageRetentionDuration string `pulumi:"topicMessageRetentionDuration"`
}

func LookupSubscriptionOutput(ctx *pulumi.Context, args LookupSubscriptionOutputArgs, opts ...pulumi.InvokeOption) LookupSubscriptionResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupSubscriptionResult, error) {
			args := v.(LookupSubscriptionArgs)
			r, err := LookupSubscription(ctx, &args, opts...)
			return *r, err
		}).(LookupSubscriptionResultOutput)
}

type LookupSubscriptionOutputArgs struct {
	Project        pulumi.StringPtrInput `pulumi:"project"`
	SubscriptionId pulumi.StringInput    `pulumi:"subscriptionId"`
}

func (LookupSubscriptionOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupSubscriptionArgs)(nil)).Elem()
}

type LookupSubscriptionResultOutput struct{ *pulumi.OutputState }

func (LookupSubscriptionResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupSubscriptionResult)(nil)).Elem()
}

func (o LookupSubscriptionResultOutput) ToLookupSubscriptionResultOutput() LookupSubscriptionResultOutput {
	return o
}

func (o LookupSubscriptionResultOutput) ToLookupSubscriptionResultOutputWithContext(ctx context.Context) LookupSubscriptionResultOutput {
	return o
}

// The approximate amount of time (on a best-effort basis) Pub/Sub waits for the subscriber to acknowledge receipt before resending the message. In the interval after the message is delivered and before it is acknowledged, it is considered to be *outstanding*. During that time period, the message will not be redelivered (on a best-effort basis). For pull subscriptions, this value is used as the initial value for the ack deadline. To override this value for a given message, call `ModifyAckDeadline` with the corresponding `ack_id` if using non-streaming pull or send the `ack_id` in a `StreamingModifyAckDeadlineRequest` if using streaming pull. The minimum custom deadline you can specify is 10 seconds. The maximum custom deadline you can specify is 600 seconds (10 minutes). If this parameter is 0, a default value of 10 seconds is used. For push delivery, this value is also used to set the request timeout for the call to the push endpoint. If the subscriber never acknowledges the message, the Pub/Sub system will eventually redeliver the message.
func (o LookupSubscriptionResultOutput) AckDeadlineSeconds() pulumi.IntOutput {
	return o.ApplyT(func(v LookupSubscriptionResult) int { return v.AckDeadlineSeconds }).(pulumi.IntOutput)
}

// A policy that specifies the conditions for dead lettering messages in this subscription. If dead_letter_policy is not set, dead lettering is disabled. The Cloud Pub/Sub service account associated with this subscriptions's parent project (i.e., service-{project_number}@gcp-sa-pubsub.iam.gserviceaccount.com) must have permission to Acknowledge() messages on this subscription.
func (o LookupSubscriptionResultOutput) DeadLetterPolicy() DeadLetterPolicyResponseOutput {
	return o.ApplyT(func(v LookupSubscriptionResult) DeadLetterPolicyResponse { return v.DeadLetterPolicy }).(DeadLetterPolicyResponseOutput)
}

// Indicates whether the subscription is detached from its topic. Detached subscriptions don't receive messages from their topic and don't retain any backlog. `Pull` and `StreamingPull` requests will return FAILED_PRECONDITION. If the subscription is a push subscription, pushes to the endpoint will not be made.
func (o LookupSubscriptionResultOutput) Detached() pulumi.BoolOutput {
	return o.ApplyT(func(v LookupSubscriptionResult) bool { return v.Detached }).(pulumi.BoolOutput)
}

// If true, Pub/Sub provides the following guarantees for the delivery of a message with a given value of `message_id` on this subscription: * The message sent to a subscriber is guaranteed not to be resent before the message's acknowledgement deadline expires. * An acknowledged message will not be resent to a subscriber. Note that subscribers may still receive multiple copies of a message when `enable_exactly_once_delivery` is true if the message was published multiple times by a publisher client. These copies are considered distinct by Pub/Sub and have distinct `message_id` values.
func (o LookupSubscriptionResultOutput) EnableExactlyOnceDelivery() pulumi.BoolOutput {
	return o.ApplyT(func(v LookupSubscriptionResult) bool { return v.EnableExactlyOnceDelivery }).(pulumi.BoolOutput)
}

// If true, messages published with the same `ordering_key` in `PubsubMessage` will be delivered to the subscribers in the order in which they are received by the Pub/Sub system. Otherwise, they may be delivered in any order.
func (o LookupSubscriptionResultOutput) EnableMessageOrdering() pulumi.BoolOutput {
	return o.ApplyT(func(v LookupSubscriptionResult) bool { return v.EnableMessageOrdering }).(pulumi.BoolOutput)
}

// A policy that specifies the conditions for this subscription's expiration. A subscription is considered active as long as any connected subscriber is successfully consuming messages from the subscription or is issuing operations on the subscription. If `expiration_policy` is not set, a *default policy* with `ttl` of 31 days will be used. The minimum allowed value for `expiration_policy.ttl` is 1 day. If `expiration_policy` is set, but `expiration_policy.ttl` is not set, the subscription never expires.
func (o LookupSubscriptionResultOutput) ExpirationPolicy() ExpirationPolicyResponseOutput {
	return o.ApplyT(func(v LookupSubscriptionResult) ExpirationPolicyResponse { return v.ExpirationPolicy }).(ExpirationPolicyResponseOutput)
}

// An expression written in the Pub/Sub [filter language](https://cloud.google.com/pubsub/docs/filtering). If non-empty, then only `PubsubMessage`s whose `attributes` field matches the filter are delivered on this subscription. If empty, then no messages are filtered out.
func (o LookupSubscriptionResultOutput) Filter() pulumi.StringOutput {
	return o.ApplyT(func(v LookupSubscriptionResult) string { return v.Filter }).(pulumi.StringOutput)
}

// See Creating and managing labels.
func (o LookupSubscriptionResultOutput) Labels() pulumi.StringMapOutput {
	return o.ApplyT(func(v LookupSubscriptionResult) map[string]string { return v.Labels }).(pulumi.StringMapOutput)
}

// How long to retain unacknowledged messages in the subscription's backlog, from the moment a message is published. If `retain_acked_messages` is true, then this also configures the retention of acknowledged messages, and thus configures how far back in time a `Seek` can be done. Defaults to 7 days. Cannot be more than 7 days or less than 10 minutes.
func (o LookupSubscriptionResultOutput) MessageRetentionDuration() pulumi.StringOutput {
	return o.ApplyT(func(v LookupSubscriptionResult) string { return v.MessageRetentionDuration }).(pulumi.StringOutput)
}

// The name of the subscription. It must have the format `"projects/{project}/subscriptions/{subscription}"`. `{subscription}` must start with a letter, and contain only letters (`[A-Za-z]`), numbers (`[0-9]`), dashes (`-`), underscores (`_`), periods (`.`), tildes (`~`), plus (`+`) or percent signs (`%`). It must be between 3 and 255 characters in length, and it must not start with `"goog"`.
func (o LookupSubscriptionResultOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v LookupSubscriptionResult) string { return v.Name }).(pulumi.StringOutput)
}

// If push delivery is used with this subscription, this field is used to configure it. At most one of `pushConfig` and `bigQueryConfig` can be set. If both are empty, then the subscriber will pull and ack messages using API methods.
func (o LookupSubscriptionResultOutput) PushConfig() PushConfigResponseOutput {
	return o.ApplyT(func(v LookupSubscriptionResult) PushConfigResponse { return v.PushConfig }).(PushConfigResponseOutput)
}

// Indicates whether to retain acknowledged messages. If true, then messages are not expunged from the subscription's backlog, even if they are acknowledged, until they fall out of the `message_retention_duration` window. This must be true if you would like to [`Seek` to a timestamp] (https://cloud.google.com/pubsub/docs/replay-overview#seek_to_a_time) in the past to replay previously-acknowledged messages.
func (o LookupSubscriptionResultOutput) RetainAckedMessages() pulumi.BoolOutput {
	return o.ApplyT(func(v LookupSubscriptionResult) bool { return v.RetainAckedMessages }).(pulumi.BoolOutput)
}

// A policy that specifies how Pub/Sub retries message delivery for this subscription. If not set, the default retry policy is applied. This generally implies that messages will be retried as soon as possible for healthy subscribers. RetryPolicy will be triggered on NACKs or acknowledgement deadline exceeded events for a given message.
func (o LookupSubscriptionResultOutput) RetryPolicy() RetryPolicyResponseOutput {
	return o.ApplyT(func(v LookupSubscriptionResult) RetryPolicyResponse { return v.RetryPolicy }).(RetryPolicyResponseOutput)
}

// An output-only field indicating whether or not the subscription can receive messages.
func (o LookupSubscriptionResultOutput) State() pulumi.StringOutput {
	return o.ApplyT(func(v LookupSubscriptionResult) string { return v.State }).(pulumi.StringOutput)
}

// The name of the topic from which this subscription is receiving messages. Format is `projects/{project}/topics/{topic}`. The value of this field will be `_deleted-topic_` if the topic has been deleted.
func (o LookupSubscriptionResultOutput) Topic() pulumi.StringOutput {
	return o.ApplyT(func(v LookupSubscriptionResult) string { return v.Topic }).(pulumi.StringOutput)
}

// Indicates the minimum duration for which a message is retained after it is published to the subscription's topic. If this field is set, messages published to the subscription's topic in the last `topic_message_retention_duration` are always available to subscribers. See the `message_retention_duration` field in `Topic`. This field is set only in responses from the server; it is ignored if it is set in any requests.
func (o LookupSubscriptionResultOutput) TopicMessageRetentionDuration() pulumi.StringOutput {
	return o.ApplyT(func(v LookupSubscriptionResult) string { return v.TopicMessageRetentionDuration }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupSubscriptionResultOutput{})
}
