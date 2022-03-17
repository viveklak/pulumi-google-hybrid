# coding=utf-8
# *** WARNING: this file was generated by the Pulumi SDK Generator. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import warnings
import pulumi
import pulumi.runtime
from typing import Any, Mapping, Optional, Sequence, Union, overload
from ... import _utilities
from . import outputs

__all__ = [
    'GetSubscriptionResult',
    'AwaitableGetSubscriptionResult',
    'get_subscription',
    'get_subscription_output',
]

@pulumi.output_type
class GetSubscriptionResult:
    def __init__(__self__, ack_deadline_seconds=None, name=None, push_config=None, topic=None):
        if ack_deadline_seconds and not isinstance(ack_deadline_seconds, int):
            raise TypeError("Expected argument 'ack_deadline_seconds' to be a int")
        pulumi.set(__self__, "ack_deadline_seconds", ack_deadline_seconds)
        if name and not isinstance(name, str):
            raise TypeError("Expected argument 'name' to be a str")
        pulumi.set(__self__, "name", name)
        if push_config and not isinstance(push_config, dict):
            raise TypeError("Expected argument 'push_config' to be a dict")
        pulumi.set(__self__, "push_config", push_config)
        if topic and not isinstance(topic, str):
            raise TypeError("Expected argument 'topic' to be a str")
        pulumi.set(__self__, "topic", topic)

    @property
    @pulumi.getter(name="ackDeadlineSeconds")
    def ack_deadline_seconds(self) -> int:
        """
        This value is the maximum time after a subscriber receives a message before the subscriber should acknowledge the message. After message delivery but before the ack deadline expires and before the message is acknowledged, it is an outstanding message and will not be delivered again during that time (on a best-effort basis). For pull subscriptions, this value is used as the initial value for the ack deadline. To override this value for a given message, call `ModifyAckDeadline` with the corresponding `ack_id` if using pull. The maximum custom deadline you can specify is 600 seconds (10 minutes). For push delivery, this value is also used to set the request timeout for the call to the push endpoint. If the subscriber never acknowledges the message, the Pub/Sub system will eventually redeliver the message. If this parameter is 0, a default value of 10 seconds is used.
        """
        return pulumi.get(self, "ack_deadline_seconds")

    @property
    @pulumi.getter
    def name(self) -> str:
        """
        The name of the subscription. It must have the format `"projects/{project}/subscriptions/{subscription}"`. `{subscription}` must start with a letter, and contain only letters (`[A-Za-z]`), numbers (`[0-9]`), dashes (`-`), underscores (`_`), periods (`.`), tildes (`~`), plus (`+`) or percent signs (`%`). It must be between 3 and 255 characters in length, and it must not start with `"goog"`.
        """
        return pulumi.get(self, "name")

    @property
    @pulumi.getter(name="pushConfig")
    def push_config(self) -> 'outputs.PushConfigResponse':
        """
        If push delivery is used with this subscription, this field is used to configure it. An empty `pushConfig` signifies that the subscriber will pull and ack messages using API methods.
        """
        return pulumi.get(self, "push_config")

    @property
    @pulumi.getter
    def topic(self) -> str:
        """
        The name of the topic from which this subscription is receiving messages. The value of this field will be `_deleted-topic_` if the topic has been deleted.
        """
        return pulumi.get(self, "topic")


class AwaitableGetSubscriptionResult(GetSubscriptionResult):
    # pylint: disable=using-constant-test
    def __await__(self):
        if False:
            yield self
        return GetSubscriptionResult(
            ack_deadline_seconds=self.ack_deadline_seconds,
            name=self.name,
            push_config=self.push_config,
            topic=self.topic)


def get_subscription(project: Optional[str] = None,
                     subscription_id: Optional[str] = None,
                     opts: Optional[pulumi.InvokeOptions] = None) -> AwaitableGetSubscriptionResult:
    """
    Gets the configuration details of a subscription.
    """
    __args__ = dict()
    __args__['project'] = project
    __args__['subscriptionId'] = subscription_id
    if opts is None:
        opts = pulumi.InvokeOptions()
    if opts.version is None:
        opts.version = _utilities.get_version()
    __ret__ = pulumi.runtime.invoke('google-native:pubsub/v1beta2:getSubscription', __args__, opts=opts, typ=GetSubscriptionResult).value

    return AwaitableGetSubscriptionResult(
        ack_deadline_seconds=__ret__.ack_deadline_seconds,
        name=__ret__.name,
        push_config=__ret__.push_config,
        topic=__ret__.topic)


@_utilities.lift_output_func(get_subscription)
def get_subscription_output(project: Optional[pulumi.Input[Optional[str]]] = None,
                            subscription_id: Optional[pulumi.Input[str]] = None,
                            opts: Optional[pulumi.InvokeOptions] = None) -> pulumi.Output[GetSubscriptionResult]:
    """
    Gets the configuration details of a subscription.
    """
    ...
