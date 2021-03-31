# coding=utf-8
# *** WARNING: this file was generated by the Pulumi SDK Generator. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import warnings
import pulumi
import pulumi.runtime
from typing import Any, Mapping, Optional, Sequence, Union
from ... import _utilities, _tables
from ._inputs import *

__all__ = ['ResponsePolicy']


class ResponsePolicy(pulumi.CustomResource):
    def __init__(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 description: Optional[pulumi.Input[str]] = None,
                 id: Optional[pulumi.Input[str]] = None,
                 kind: Optional[pulumi.Input[str]] = None,
                 networks: Optional[pulumi.Input[Sequence[pulumi.Input[pulumi.InputType['ResponsePolicyNetworkArgs']]]]] = None,
                 project: Optional[pulumi.Input[str]] = None,
                 response_policy: Optional[pulumi.Input[str]] = None,
                 response_policy_name: Optional[pulumi.Input[str]] = None,
                 __props__=None,
                 __name__=None,
                 __opts__=None):
        """
        Creates a new Response Policy

        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] description: User-provided description for this Response Policy.
        :param pulumi.Input[str] id: Unique identifier for the resource; defined by the server (output only).
        :param pulumi.Input[Sequence[pulumi.Input[pulumi.InputType['ResponsePolicyNetworkArgs']]]] networks: List of network names specifying networks to which this policy is applied.
        :param pulumi.Input[str] response_policy_name: User assigned name for this Response Policy.
        """
        if __name__ is not None:
            warnings.warn("explicit use of __name__ is deprecated", DeprecationWarning)
            resource_name = __name__
        if __opts__ is not None:
            warnings.warn("explicit use of __opts__ is deprecated, use 'opts' instead", DeprecationWarning)
            opts = __opts__
        if opts is None:
            opts = pulumi.ResourceOptions()
        if not isinstance(opts, pulumi.ResourceOptions):
            raise TypeError('Expected resource options to be a ResourceOptions instance')
        if opts.version is None:
            opts.version = _utilities.get_version()
        if opts.id is None:
            if __props__ is not None:
                raise TypeError('__props__ is only valid when passed in combination with a valid opts.id to get an existing resource')
            __props__ = dict()

            __props__['description'] = description
            __props__['id'] = id
            __props__['kind'] = kind
            __props__['networks'] = networks
            if project is None and not opts.urn:
                raise TypeError("Missing required property 'project'")
            __props__['project'] = project
            if response_policy is None and not opts.urn:
                raise TypeError("Missing required property 'response_policy'")
            __props__['response_policy'] = response_policy
            __props__['response_policy_name'] = response_policy_name
        super(ResponsePolicy, __self__).__init__(
            'google-cloud:dns/v1beta2:ResponsePolicy',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name: str,
            id: pulumi.Input[str],
            opts: Optional[pulumi.ResourceOptions] = None) -> 'ResponsePolicy':
        """
        Get an existing ResponsePolicy resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param pulumi.Input[str] id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = dict()

        return ResponsePolicy(resource_name, opts=opts, __props__=__props__)

    def translate_output_property(self, prop):
        return _tables.CAMEL_TO_SNAKE_CASE_TABLE.get(prop) or prop

    def translate_input_property(self, prop):
        return _tables.SNAKE_TO_CAMEL_CASE_TABLE.get(prop) or prop

