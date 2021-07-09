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
    'GetWebhookResult',
    'AwaitableGetWebhookResult',
    'get_webhook',
]

@pulumi.output_type
class GetWebhookResult:
    def __init__(__self__, disabled=None, display_name=None, generic_web_service=None, name=None, timeout=None):
        if disabled and not isinstance(disabled, bool):
            raise TypeError("Expected argument 'disabled' to be a bool")
        pulumi.set(__self__, "disabled", disabled)
        if display_name and not isinstance(display_name, str):
            raise TypeError("Expected argument 'display_name' to be a str")
        pulumi.set(__self__, "display_name", display_name)
        if generic_web_service and not isinstance(generic_web_service, dict):
            raise TypeError("Expected argument 'generic_web_service' to be a dict")
        pulumi.set(__self__, "generic_web_service", generic_web_service)
        if name and not isinstance(name, str):
            raise TypeError("Expected argument 'name' to be a str")
        pulumi.set(__self__, "name", name)
        if timeout and not isinstance(timeout, str):
            raise TypeError("Expected argument 'timeout' to be a str")
        pulumi.set(__self__, "timeout", timeout)

    @property
    @pulumi.getter
    def disabled(self) -> bool:
        """
        Indicates whether the webhook is disabled.
        """
        return pulumi.get(self, "disabled")

    @property
    @pulumi.getter(name="displayName")
    def display_name(self) -> str:
        """
        The human-readable name of the webhook, unique within the agent.
        """
        return pulumi.get(self, "display_name")

    @property
    @pulumi.getter(name="genericWebService")
    def generic_web_service(self) -> 'outputs.GoogleCloudDialogflowCxV3beta1WebhookGenericWebServiceResponse':
        """
        Configuration for a generic web service.
        """
        return pulumi.get(self, "generic_web_service")

    @property
    @pulumi.getter
    def name(self) -> str:
        """
        The unique identifier of the webhook. Required for the Webhooks.UpdateWebhook method. Webhooks.CreateWebhook populates the name automatically. Format: `projects//locations//agents//webhooks/`.
        """
        return pulumi.get(self, "name")

    @property
    @pulumi.getter
    def timeout(self) -> str:
        """
        Webhook execution timeout. Execution is considered failed if Dialogflow doesn't receive a response from webhook at the end of the timeout period. Defaults to 5 seconds, maximum allowed timeout is 30 seconds.
        """
        return pulumi.get(self, "timeout")


class AwaitableGetWebhookResult(GetWebhookResult):
    # pylint: disable=using-constant-test
    def __await__(self):
        if False:
            yield self
        return GetWebhookResult(
            disabled=self.disabled,
            display_name=self.display_name,
            generic_web_service=self.generic_web_service,
            name=self.name,
            timeout=self.timeout)


def get_webhook(agent_id: Optional[str] = None,
                location: Optional[str] = None,
                project: Optional[str] = None,
                webhook_id: Optional[str] = None,
                opts: Optional[pulumi.InvokeOptions] = None) -> AwaitableGetWebhookResult:
    """
    Retrieves the specified webhook.
    """
    __args__ = dict()
    __args__['agentId'] = agent_id
    __args__['location'] = location
    __args__['project'] = project
    __args__['webhookId'] = webhook_id
    if opts is None:
        opts = pulumi.InvokeOptions()
    if opts.version is None:
        opts.version = _utilities.get_version()
    __ret__ = pulumi.runtime.invoke('google-native:dialogflow/v3beta1:getWebhook', __args__, opts=opts, typ=GetWebhookResult).value

    return AwaitableGetWebhookResult(
        disabled=__ret__.disabled,
        display_name=__ret__.display_name,
        generic_web_service=__ret__.generic_web_service,
        name=__ret__.name,
        timeout=__ret__.timeout)
