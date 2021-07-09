# coding=utf-8
# *** WARNING: this file was generated by the Pulumi SDK Generator. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import warnings
import pulumi
import pulumi.runtime
from typing import Any, Mapping, Optional, Sequence, Union, overload
from ... import _utilities
from . import outputs
from ._inputs import *

__all__ = ['WebhookArgs', 'Webhook']

@pulumi.input_type
class WebhookArgs:
    def __init__(__self__, *,
                 agent_id: pulumi.Input[str],
                 display_name: pulumi.Input[str],
                 location: pulumi.Input[str],
                 project: pulumi.Input[str],
                 disabled: Optional[pulumi.Input[bool]] = None,
                 generic_web_service: Optional[pulumi.Input['GoogleCloudDialogflowCxV3WebhookGenericWebServiceArgs']] = None,
                 name: Optional[pulumi.Input[str]] = None,
                 timeout: Optional[pulumi.Input[str]] = None):
        """
        The set of arguments for constructing a Webhook resource.
        :param pulumi.Input[str] display_name: The human-readable name of the webhook, unique within the agent.
        :param pulumi.Input[bool] disabled: Indicates whether the webhook is disabled.
        :param pulumi.Input['GoogleCloudDialogflowCxV3WebhookGenericWebServiceArgs'] generic_web_service: Configuration for a generic web service.
        :param pulumi.Input[str] name: The unique identifier of the webhook. Required for the Webhooks.UpdateWebhook method. Webhooks.CreateWebhook populates the name automatically. Format: `projects//locations//agents//webhooks/`.
        :param pulumi.Input[str] timeout: Webhook execution timeout. Execution is considered failed if Dialogflow doesn't receive a response from webhook at the end of the timeout period. Defaults to 5 seconds, maximum allowed timeout is 30 seconds.
        """
        pulumi.set(__self__, "agent_id", agent_id)
        pulumi.set(__self__, "display_name", display_name)
        pulumi.set(__self__, "location", location)
        pulumi.set(__self__, "project", project)
        if disabled is not None:
            pulumi.set(__self__, "disabled", disabled)
        if generic_web_service is not None:
            pulumi.set(__self__, "generic_web_service", generic_web_service)
        if name is not None:
            pulumi.set(__self__, "name", name)
        if timeout is not None:
            pulumi.set(__self__, "timeout", timeout)

    @property
    @pulumi.getter(name="agentId")
    def agent_id(self) -> pulumi.Input[str]:
        return pulumi.get(self, "agent_id")

    @agent_id.setter
    def agent_id(self, value: pulumi.Input[str]):
        pulumi.set(self, "agent_id", value)

    @property
    @pulumi.getter(name="displayName")
    def display_name(self) -> pulumi.Input[str]:
        """
        The human-readable name of the webhook, unique within the agent.
        """
        return pulumi.get(self, "display_name")

    @display_name.setter
    def display_name(self, value: pulumi.Input[str]):
        pulumi.set(self, "display_name", value)

    @property
    @pulumi.getter
    def location(self) -> pulumi.Input[str]:
        return pulumi.get(self, "location")

    @location.setter
    def location(self, value: pulumi.Input[str]):
        pulumi.set(self, "location", value)

    @property
    @pulumi.getter
    def project(self) -> pulumi.Input[str]:
        return pulumi.get(self, "project")

    @project.setter
    def project(self, value: pulumi.Input[str]):
        pulumi.set(self, "project", value)

    @property
    @pulumi.getter
    def disabled(self) -> Optional[pulumi.Input[bool]]:
        """
        Indicates whether the webhook is disabled.
        """
        return pulumi.get(self, "disabled")

    @disabled.setter
    def disabled(self, value: Optional[pulumi.Input[bool]]):
        pulumi.set(self, "disabled", value)

    @property
    @pulumi.getter(name="genericWebService")
    def generic_web_service(self) -> Optional[pulumi.Input['GoogleCloudDialogflowCxV3WebhookGenericWebServiceArgs']]:
        """
        Configuration for a generic web service.
        """
        return pulumi.get(self, "generic_web_service")

    @generic_web_service.setter
    def generic_web_service(self, value: Optional[pulumi.Input['GoogleCloudDialogflowCxV3WebhookGenericWebServiceArgs']]):
        pulumi.set(self, "generic_web_service", value)

    @property
    @pulumi.getter
    def name(self) -> Optional[pulumi.Input[str]]:
        """
        The unique identifier of the webhook. Required for the Webhooks.UpdateWebhook method. Webhooks.CreateWebhook populates the name automatically. Format: `projects//locations//agents//webhooks/`.
        """
        return pulumi.get(self, "name")

    @name.setter
    def name(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "name", value)

    @property
    @pulumi.getter
    def timeout(self) -> Optional[pulumi.Input[str]]:
        """
        Webhook execution timeout. Execution is considered failed if Dialogflow doesn't receive a response from webhook at the end of the timeout period. Defaults to 5 seconds, maximum allowed timeout is 30 seconds.
        """
        return pulumi.get(self, "timeout")

    @timeout.setter
    def timeout(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "timeout", value)


class Webhook(pulumi.CustomResource):
    @overload
    def __init__(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 agent_id: Optional[pulumi.Input[str]] = None,
                 disabled: Optional[pulumi.Input[bool]] = None,
                 display_name: Optional[pulumi.Input[str]] = None,
                 generic_web_service: Optional[pulumi.Input[pulumi.InputType['GoogleCloudDialogflowCxV3WebhookGenericWebServiceArgs']]] = None,
                 location: Optional[pulumi.Input[str]] = None,
                 name: Optional[pulumi.Input[str]] = None,
                 project: Optional[pulumi.Input[str]] = None,
                 timeout: Optional[pulumi.Input[str]] = None,
                 __props__=None):
        """
        Creates a webhook in the specified agent.

        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[bool] disabled: Indicates whether the webhook is disabled.
        :param pulumi.Input[str] display_name: The human-readable name of the webhook, unique within the agent.
        :param pulumi.Input[pulumi.InputType['GoogleCloudDialogflowCxV3WebhookGenericWebServiceArgs']] generic_web_service: Configuration for a generic web service.
        :param pulumi.Input[str] name: The unique identifier of the webhook. Required for the Webhooks.UpdateWebhook method. Webhooks.CreateWebhook populates the name automatically. Format: `projects//locations//agents//webhooks/`.
        :param pulumi.Input[str] timeout: Webhook execution timeout. Execution is considered failed if Dialogflow doesn't receive a response from webhook at the end of the timeout period. Defaults to 5 seconds, maximum allowed timeout is 30 seconds.
        """
        ...
    @overload
    def __init__(__self__,
                 resource_name: str,
                 args: WebhookArgs,
                 opts: Optional[pulumi.ResourceOptions] = None):
        """
        Creates a webhook in the specified agent.

        :param str resource_name: The name of the resource.
        :param WebhookArgs args: The arguments to use to populate this resource's properties.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        ...
    def __init__(__self__, resource_name: str, *args, **kwargs):
        resource_args, opts = _utilities.get_resource_args_opts(WebhookArgs, pulumi.ResourceOptions, *args, **kwargs)
        if resource_args is not None:
            __self__._internal_init(resource_name, opts, **resource_args.__dict__)
        else:
            __self__._internal_init(resource_name, *args, **kwargs)

    def _internal_init(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 agent_id: Optional[pulumi.Input[str]] = None,
                 disabled: Optional[pulumi.Input[bool]] = None,
                 display_name: Optional[pulumi.Input[str]] = None,
                 generic_web_service: Optional[pulumi.Input[pulumi.InputType['GoogleCloudDialogflowCxV3WebhookGenericWebServiceArgs']]] = None,
                 location: Optional[pulumi.Input[str]] = None,
                 name: Optional[pulumi.Input[str]] = None,
                 project: Optional[pulumi.Input[str]] = None,
                 timeout: Optional[pulumi.Input[str]] = None,
                 __props__=None):
        if opts is None:
            opts = pulumi.ResourceOptions()
        if not isinstance(opts, pulumi.ResourceOptions):
            raise TypeError('Expected resource options to be a ResourceOptions instance')
        if opts.version is None:
            opts.version = _utilities.get_version()
        if opts.id is None:
            if __props__ is not None:
                raise TypeError('__props__ is only valid when passed in combination with a valid opts.id to get an existing resource')
            __props__ = WebhookArgs.__new__(WebhookArgs)

            if agent_id is None and not opts.urn:
                raise TypeError("Missing required property 'agent_id'")
            __props__.__dict__["agent_id"] = agent_id
            __props__.__dict__["disabled"] = disabled
            if display_name is None and not opts.urn:
                raise TypeError("Missing required property 'display_name'")
            __props__.__dict__["display_name"] = display_name
            __props__.__dict__["generic_web_service"] = generic_web_service
            if location is None and not opts.urn:
                raise TypeError("Missing required property 'location'")
            __props__.__dict__["location"] = location
            __props__.__dict__["name"] = name
            if project is None and not opts.urn:
                raise TypeError("Missing required property 'project'")
            __props__.__dict__["project"] = project
            __props__.__dict__["timeout"] = timeout
        super(Webhook, __self__).__init__(
            'google-native:dialogflow/v3:Webhook',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name: str,
            id: pulumi.Input[str],
            opts: Optional[pulumi.ResourceOptions] = None) -> 'Webhook':
        """
        Get an existing Webhook resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param pulumi.Input[str] id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = WebhookArgs.__new__(WebhookArgs)

        __props__.__dict__["disabled"] = None
        __props__.__dict__["display_name"] = None
        __props__.__dict__["generic_web_service"] = None
        __props__.__dict__["name"] = None
        __props__.__dict__["timeout"] = None
        return Webhook(resource_name, opts=opts, __props__=__props__)

    @property
    @pulumi.getter
    def disabled(self) -> pulumi.Output[bool]:
        """
        Indicates whether the webhook is disabled.
        """
        return pulumi.get(self, "disabled")

    @property
    @pulumi.getter(name="displayName")
    def display_name(self) -> pulumi.Output[str]:
        """
        The human-readable name of the webhook, unique within the agent.
        """
        return pulumi.get(self, "display_name")

    @property
    @pulumi.getter(name="genericWebService")
    def generic_web_service(self) -> pulumi.Output['outputs.GoogleCloudDialogflowCxV3WebhookGenericWebServiceResponse']:
        """
        Configuration for a generic web service.
        """
        return pulumi.get(self, "generic_web_service")

    @property
    @pulumi.getter
    def name(self) -> pulumi.Output[str]:
        """
        The unique identifier of the webhook. Required for the Webhooks.UpdateWebhook method. Webhooks.CreateWebhook populates the name automatically. Format: `projects//locations//agents//webhooks/`.
        """
        return pulumi.get(self, "name")

    @property
    @pulumi.getter
    def timeout(self) -> pulumi.Output[str]:
        """
        Webhook execution timeout. Execution is considered failed if Dialogflow doesn't receive a response from webhook at the end of the timeout period. Defaults to 5 seconds, maximum allowed timeout is 30 seconds.
        """
        return pulumi.get(self, "timeout")

