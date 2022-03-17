# coding=utf-8
# *** WARNING: this file was generated by the Pulumi SDK Generator. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import warnings
import pulumi
import pulumi.runtime
from typing import Any, Mapping, Optional, Sequence, Union, overload
from ... import _utilities
from . import outputs
from ._enums import *
from ._inputs import *

__all__ = ['ResponsePolicyRuleArgs', 'ResponsePolicyRule']

@pulumi.input_type
class ResponsePolicyRuleArgs:
    def __init__(__self__, *,
                 response_policy: pulumi.Input[str],
                 behavior: Optional[pulumi.Input['ResponsePolicyRuleBehavior']] = None,
                 client_operation_id: Optional[pulumi.Input[str]] = None,
                 dns_name: Optional[pulumi.Input[str]] = None,
                 kind: Optional[pulumi.Input[str]] = None,
                 local_data: Optional[pulumi.Input['ResponsePolicyRuleLocalDataArgs']] = None,
                 project: Optional[pulumi.Input[str]] = None,
                 rule_name: Optional[pulumi.Input[str]] = None):
        """
        The set of arguments for constructing a ResponsePolicyRule resource.
        :param pulumi.Input['ResponsePolicyRuleBehavior'] behavior: Answer this query with a behavior rather than DNS data.
        :param pulumi.Input[str] client_operation_id: For mutating operation requests only. An optional identifier specified by the client. Must be unique for operation resources in the Operations collection.
        :param pulumi.Input[str] dns_name: The DNS name (wildcard or exact) to apply this rule to. Must be unique within the Response Policy Rule.
        :param pulumi.Input['ResponsePolicyRuleLocalDataArgs'] local_data: Answer this query directly with DNS data. These ResourceRecordSets override any other DNS behavior for the matched name; in particular they override private zones, the public internet, and GCP internal DNS. No SOA nor NS types are allowed.
        :param pulumi.Input[str] rule_name: An identifier for this rule. Must be unique with the ResponsePolicy.
        """
        pulumi.set(__self__, "response_policy", response_policy)
        if behavior is not None:
            pulumi.set(__self__, "behavior", behavior)
        if client_operation_id is not None:
            pulumi.set(__self__, "client_operation_id", client_operation_id)
        if dns_name is not None:
            pulumi.set(__self__, "dns_name", dns_name)
        if kind is not None:
            pulumi.set(__self__, "kind", kind)
        if local_data is not None:
            pulumi.set(__self__, "local_data", local_data)
        if project is not None:
            pulumi.set(__self__, "project", project)
        if rule_name is not None:
            pulumi.set(__self__, "rule_name", rule_name)

    @property
    @pulumi.getter(name="responsePolicy")
    def response_policy(self) -> pulumi.Input[str]:
        return pulumi.get(self, "response_policy")

    @response_policy.setter
    def response_policy(self, value: pulumi.Input[str]):
        pulumi.set(self, "response_policy", value)

    @property
    @pulumi.getter
    def behavior(self) -> Optional[pulumi.Input['ResponsePolicyRuleBehavior']]:
        """
        Answer this query with a behavior rather than DNS data.
        """
        return pulumi.get(self, "behavior")

    @behavior.setter
    def behavior(self, value: Optional[pulumi.Input['ResponsePolicyRuleBehavior']]):
        pulumi.set(self, "behavior", value)

    @property
    @pulumi.getter(name="clientOperationId")
    def client_operation_id(self) -> Optional[pulumi.Input[str]]:
        """
        For mutating operation requests only. An optional identifier specified by the client. Must be unique for operation resources in the Operations collection.
        """
        return pulumi.get(self, "client_operation_id")

    @client_operation_id.setter
    def client_operation_id(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "client_operation_id", value)

    @property
    @pulumi.getter(name="dnsName")
    def dns_name(self) -> Optional[pulumi.Input[str]]:
        """
        The DNS name (wildcard or exact) to apply this rule to. Must be unique within the Response Policy Rule.
        """
        return pulumi.get(self, "dns_name")

    @dns_name.setter
    def dns_name(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "dns_name", value)

    @property
    @pulumi.getter
    def kind(self) -> Optional[pulumi.Input[str]]:
        return pulumi.get(self, "kind")

    @kind.setter
    def kind(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "kind", value)

    @property
    @pulumi.getter(name="localData")
    def local_data(self) -> Optional[pulumi.Input['ResponsePolicyRuleLocalDataArgs']]:
        """
        Answer this query directly with DNS data. These ResourceRecordSets override any other DNS behavior for the matched name; in particular they override private zones, the public internet, and GCP internal DNS. No SOA nor NS types are allowed.
        """
        return pulumi.get(self, "local_data")

    @local_data.setter
    def local_data(self, value: Optional[pulumi.Input['ResponsePolicyRuleLocalDataArgs']]):
        pulumi.set(self, "local_data", value)

    @property
    @pulumi.getter
    def project(self) -> Optional[pulumi.Input[str]]:
        return pulumi.get(self, "project")

    @project.setter
    def project(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "project", value)

    @property
    @pulumi.getter(name="ruleName")
    def rule_name(self) -> Optional[pulumi.Input[str]]:
        """
        An identifier for this rule. Must be unique with the ResponsePolicy.
        """
        return pulumi.get(self, "rule_name")

    @rule_name.setter
    def rule_name(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "rule_name", value)


class ResponsePolicyRule(pulumi.CustomResource):
    @overload
    def __init__(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 behavior: Optional[pulumi.Input['ResponsePolicyRuleBehavior']] = None,
                 client_operation_id: Optional[pulumi.Input[str]] = None,
                 dns_name: Optional[pulumi.Input[str]] = None,
                 kind: Optional[pulumi.Input[str]] = None,
                 local_data: Optional[pulumi.Input[pulumi.InputType['ResponsePolicyRuleLocalDataArgs']]] = None,
                 project: Optional[pulumi.Input[str]] = None,
                 response_policy: Optional[pulumi.Input[str]] = None,
                 rule_name: Optional[pulumi.Input[str]] = None,
                 __props__=None):
        """
        Creates a new Response Policy Rule.
        Auto-naming is currently not supported for this resource.

        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input['ResponsePolicyRuleBehavior'] behavior: Answer this query with a behavior rather than DNS data.
        :param pulumi.Input[str] client_operation_id: For mutating operation requests only. An optional identifier specified by the client. Must be unique for operation resources in the Operations collection.
        :param pulumi.Input[str] dns_name: The DNS name (wildcard or exact) to apply this rule to. Must be unique within the Response Policy Rule.
        :param pulumi.Input[pulumi.InputType['ResponsePolicyRuleLocalDataArgs']] local_data: Answer this query directly with DNS data. These ResourceRecordSets override any other DNS behavior for the matched name; in particular they override private zones, the public internet, and GCP internal DNS. No SOA nor NS types are allowed.
        :param pulumi.Input[str] rule_name: An identifier for this rule. Must be unique with the ResponsePolicy.
        """
        ...
    @overload
    def __init__(__self__,
                 resource_name: str,
                 args: ResponsePolicyRuleArgs,
                 opts: Optional[pulumi.ResourceOptions] = None):
        """
        Creates a new Response Policy Rule.
        Auto-naming is currently not supported for this resource.

        :param str resource_name: The name of the resource.
        :param ResponsePolicyRuleArgs args: The arguments to use to populate this resource's properties.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        ...
    def __init__(__self__, resource_name: str, *args, **kwargs):
        resource_args, opts = _utilities.get_resource_args_opts(ResponsePolicyRuleArgs, pulumi.ResourceOptions, *args, **kwargs)
        if resource_args is not None:
            __self__._internal_init(resource_name, opts, **resource_args.__dict__)
        else:
            __self__._internal_init(resource_name, *args, **kwargs)

    def _internal_init(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 behavior: Optional[pulumi.Input['ResponsePolicyRuleBehavior']] = None,
                 client_operation_id: Optional[pulumi.Input[str]] = None,
                 dns_name: Optional[pulumi.Input[str]] = None,
                 kind: Optional[pulumi.Input[str]] = None,
                 local_data: Optional[pulumi.Input[pulumi.InputType['ResponsePolicyRuleLocalDataArgs']]] = None,
                 project: Optional[pulumi.Input[str]] = None,
                 response_policy: Optional[pulumi.Input[str]] = None,
                 rule_name: Optional[pulumi.Input[str]] = None,
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
            __props__ = ResponsePolicyRuleArgs.__new__(ResponsePolicyRuleArgs)

            __props__.__dict__["behavior"] = behavior
            __props__.__dict__["client_operation_id"] = client_operation_id
            __props__.__dict__["dns_name"] = dns_name
            __props__.__dict__["kind"] = kind
            __props__.__dict__["local_data"] = local_data
            __props__.__dict__["project"] = project
            if response_policy is None and not opts.urn:
                raise TypeError("Missing required property 'response_policy'")
            __props__.__dict__["response_policy"] = response_policy
            __props__.__dict__["rule_name"] = rule_name
        super(ResponsePolicyRule, __self__).__init__(
            'google-native:dns/v1beta2:ResponsePolicyRule',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name: str,
            id: pulumi.Input[str],
            opts: Optional[pulumi.ResourceOptions] = None) -> 'ResponsePolicyRule':
        """
        Get an existing ResponsePolicyRule resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param pulumi.Input[str] id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = ResponsePolicyRuleArgs.__new__(ResponsePolicyRuleArgs)

        __props__.__dict__["behavior"] = None
        __props__.__dict__["dns_name"] = None
        __props__.__dict__["kind"] = None
        __props__.__dict__["local_data"] = None
        __props__.__dict__["rule_name"] = None
        return ResponsePolicyRule(resource_name, opts=opts, __props__=__props__)

    @property
    @pulumi.getter
    def behavior(self) -> pulumi.Output[str]:
        """
        Answer this query with a behavior rather than DNS data.
        """
        return pulumi.get(self, "behavior")

    @property
    @pulumi.getter(name="dnsName")
    def dns_name(self) -> pulumi.Output[str]:
        """
        The DNS name (wildcard or exact) to apply this rule to. Must be unique within the Response Policy Rule.
        """
        return pulumi.get(self, "dns_name")

    @property
    @pulumi.getter
    def kind(self) -> pulumi.Output[str]:
        return pulumi.get(self, "kind")

    @property
    @pulumi.getter(name="localData")
    def local_data(self) -> pulumi.Output['outputs.ResponsePolicyRuleLocalDataResponse']:
        """
        Answer this query directly with DNS data. These ResourceRecordSets override any other DNS behavior for the matched name; in particular they override private zones, the public internet, and GCP internal DNS. No SOA nor NS types are allowed.
        """
        return pulumi.get(self, "local_data")

    @property
    @pulumi.getter(name="ruleName")
    def rule_name(self) -> pulumi.Output[str]:
        """
        An identifier for this rule. Must be unique with the ResponsePolicy.
        """
        return pulumi.get(self, "rule_name")

