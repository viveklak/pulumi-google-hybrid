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

__all__ = ['ServiceAttachmentArgs', 'ServiceAttachment']

@pulumi.input_type
class ServiceAttachmentArgs:
    def __init__(__self__, *,
                 project: pulumi.Input[str],
                 region: pulumi.Input[str],
                 service_attachment: pulumi.Input[str],
                 connection_preference: Optional[pulumi.Input[str]] = None,
                 consumer_forwarding_rules: Optional[pulumi.Input[Sequence[pulumi.Input['ServiceAttachmentConsumerForwardingRuleArgs']]]] = None,
                 creation_timestamp: Optional[pulumi.Input[str]] = None,
                 description: Optional[pulumi.Input[str]] = None,
                 enable_proxy_protocol: Optional[pulumi.Input[bool]] = None,
                 id: Optional[pulumi.Input[str]] = None,
                 kind: Optional[pulumi.Input[str]] = None,
                 name: Optional[pulumi.Input[str]] = None,
                 nat_subnets: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]] = None,
                 producer_forwarding_rule: Optional[pulumi.Input[str]] = None,
                 self_link: Optional[pulumi.Input[str]] = None):
        """
        The set of arguments for constructing a ServiceAttachment resource.
        :param pulumi.Input[str] region: [Output Only] URL of the region where the service attachment resides. This field applies only to the region resource. You must specify this field as part of the HTTP request URL. It is not settable as a field in the request body.
        :param pulumi.Input[str] connection_preference: The connection preference of service attachment. The value can be set to ACCEPT_AUTOMATIC. An ACCEPT_AUTOMATIC service attachment is one that always accepts the connection from consumer forwarding rules.
        :param pulumi.Input[Sequence[pulumi.Input['ServiceAttachmentConsumerForwardingRuleArgs']]] consumer_forwarding_rules: [Output Only] An array of forwarding rules for all the consumers connected to this service attachment.
        :param pulumi.Input[str] creation_timestamp: [Output Only] Creation timestamp in RFC3339 text format.
        :param pulumi.Input[str] description: An optional description of this resource. Provide this property when you create the resource.
        :param pulumi.Input[bool] enable_proxy_protocol: If true, enable the proxy protocol which is for supplying client TCP/IP address data in TCP connections that traverse proxies on their way to destination servers.
        :param pulumi.Input[str] id: [Output Only] The unique identifier for the resource type. The server generates this identifier.
        :param pulumi.Input[str] kind: [Output Only] Type of the resource. Always compute#serviceAttachment for service attachments.
        :param pulumi.Input[str] name: Name of the resource. Provided by the client when the resource is created. The name must be 1-63 characters long, and comply with RFC1035. Specifically, the name must be 1-63 characters long and match the regular expression `[a-z]([-a-z0-9]*[a-z0-9])?` which means the first character must be a lowercase letter, and all following characters must be a dash, lowercase letter, or digit, except the last character, which cannot be a dash.
        :param pulumi.Input[Sequence[pulumi.Input[str]]] nat_subnets: An array of URLs where each entry is the URL of a subnet provided by the service producer to use for NAT in this service attachment.
        :param pulumi.Input[str] producer_forwarding_rule: The URL of a forwarding rule with loadBalancingScheme INTERNAL* that is serving the endpoint identified by this service attachment.
        :param pulumi.Input[str] self_link: [Output Only] Server-defined URL for the resource.
        """
        pulumi.set(__self__, "project", project)
        pulumi.set(__self__, "region", region)
        pulumi.set(__self__, "service_attachment", service_attachment)
        if connection_preference is not None:
            pulumi.set(__self__, "connection_preference", connection_preference)
        if consumer_forwarding_rules is not None:
            pulumi.set(__self__, "consumer_forwarding_rules", consumer_forwarding_rules)
        if creation_timestamp is not None:
            pulumi.set(__self__, "creation_timestamp", creation_timestamp)
        if description is not None:
            pulumi.set(__self__, "description", description)
        if enable_proxy_protocol is not None:
            pulumi.set(__self__, "enable_proxy_protocol", enable_proxy_protocol)
        if id is not None:
            pulumi.set(__self__, "id", id)
        if kind is not None:
            pulumi.set(__self__, "kind", kind)
        if name is not None:
            pulumi.set(__self__, "name", name)
        if nat_subnets is not None:
            pulumi.set(__self__, "nat_subnets", nat_subnets)
        if producer_forwarding_rule is not None:
            pulumi.set(__self__, "producer_forwarding_rule", producer_forwarding_rule)
        if self_link is not None:
            pulumi.set(__self__, "self_link", self_link)

    @property
    @pulumi.getter
    def project(self) -> pulumi.Input[str]:
        return pulumi.get(self, "project")

    @project.setter
    def project(self, value: pulumi.Input[str]):
        pulumi.set(self, "project", value)

    @property
    @pulumi.getter
    def region(self) -> pulumi.Input[str]:
        """
        [Output Only] URL of the region where the service attachment resides. This field applies only to the region resource. You must specify this field as part of the HTTP request URL. It is not settable as a field in the request body.
        """
        return pulumi.get(self, "region")

    @region.setter
    def region(self, value: pulumi.Input[str]):
        pulumi.set(self, "region", value)

    @property
    @pulumi.getter(name="serviceAttachment")
    def service_attachment(self) -> pulumi.Input[str]:
        return pulumi.get(self, "service_attachment")

    @service_attachment.setter
    def service_attachment(self, value: pulumi.Input[str]):
        pulumi.set(self, "service_attachment", value)

    @property
    @pulumi.getter(name="connectionPreference")
    def connection_preference(self) -> Optional[pulumi.Input[str]]:
        """
        The connection preference of service attachment. The value can be set to ACCEPT_AUTOMATIC. An ACCEPT_AUTOMATIC service attachment is one that always accepts the connection from consumer forwarding rules.
        """
        return pulumi.get(self, "connection_preference")

    @connection_preference.setter
    def connection_preference(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "connection_preference", value)

    @property
    @pulumi.getter(name="consumerForwardingRules")
    def consumer_forwarding_rules(self) -> Optional[pulumi.Input[Sequence[pulumi.Input['ServiceAttachmentConsumerForwardingRuleArgs']]]]:
        """
        [Output Only] An array of forwarding rules for all the consumers connected to this service attachment.
        """
        return pulumi.get(self, "consumer_forwarding_rules")

    @consumer_forwarding_rules.setter
    def consumer_forwarding_rules(self, value: Optional[pulumi.Input[Sequence[pulumi.Input['ServiceAttachmentConsumerForwardingRuleArgs']]]]):
        pulumi.set(self, "consumer_forwarding_rules", value)

    @property
    @pulumi.getter(name="creationTimestamp")
    def creation_timestamp(self) -> Optional[pulumi.Input[str]]:
        """
        [Output Only] Creation timestamp in RFC3339 text format.
        """
        return pulumi.get(self, "creation_timestamp")

    @creation_timestamp.setter
    def creation_timestamp(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "creation_timestamp", value)

    @property
    @pulumi.getter
    def description(self) -> Optional[pulumi.Input[str]]:
        """
        An optional description of this resource. Provide this property when you create the resource.
        """
        return pulumi.get(self, "description")

    @description.setter
    def description(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "description", value)

    @property
    @pulumi.getter(name="enableProxyProtocol")
    def enable_proxy_protocol(self) -> Optional[pulumi.Input[bool]]:
        """
        If true, enable the proxy protocol which is for supplying client TCP/IP address data in TCP connections that traverse proxies on their way to destination servers.
        """
        return pulumi.get(self, "enable_proxy_protocol")

    @enable_proxy_protocol.setter
    def enable_proxy_protocol(self, value: Optional[pulumi.Input[bool]]):
        pulumi.set(self, "enable_proxy_protocol", value)

    @property
    @pulumi.getter
    def id(self) -> Optional[pulumi.Input[str]]:
        """
        [Output Only] The unique identifier for the resource type. The server generates this identifier.
        """
        return pulumi.get(self, "id")

    @id.setter
    def id(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "id", value)

    @property
    @pulumi.getter
    def kind(self) -> Optional[pulumi.Input[str]]:
        """
        [Output Only] Type of the resource. Always compute#serviceAttachment for service attachments.
        """
        return pulumi.get(self, "kind")

    @kind.setter
    def kind(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "kind", value)

    @property
    @pulumi.getter
    def name(self) -> Optional[pulumi.Input[str]]:
        """
        Name of the resource. Provided by the client when the resource is created. The name must be 1-63 characters long, and comply with RFC1035. Specifically, the name must be 1-63 characters long and match the regular expression `[a-z]([-a-z0-9]*[a-z0-9])?` which means the first character must be a lowercase letter, and all following characters must be a dash, lowercase letter, or digit, except the last character, which cannot be a dash.
        """
        return pulumi.get(self, "name")

    @name.setter
    def name(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "name", value)

    @property
    @pulumi.getter(name="natSubnets")
    def nat_subnets(self) -> Optional[pulumi.Input[Sequence[pulumi.Input[str]]]]:
        """
        An array of URLs where each entry is the URL of a subnet provided by the service producer to use for NAT in this service attachment.
        """
        return pulumi.get(self, "nat_subnets")

    @nat_subnets.setter
    def nat_subnets(self, value: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]]):
        pulumi.set(self, "nat_subnets", value)

    @property
    @pulumi.getter(name="producerForwardingRule")
    def producer_forwarding_rule(self) -> Optional[pulumi.Input[str]]:
        """
        The URL of a forwarding rule with loadBalancingScheme INTERNAL* that is serving the endpoint identified by this service attachment.
        """
        return pulumi.get(self, "producer_forwarding_rule")

    @producer_forwarding_rule.setter
    def producer_forwarding_rule(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "producer_forwarding_rule", value)

    @property
    @pulumi.getter(name="selfLink")
    def self_link(self) -> Optional[pulumi.Input[str]]:
        """
        [Output Only] Server-defined URL for the resource.
        """
        return pulumi.get(self, "self_link")

    @self_link.setter
    def self_link(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "self_link", value)


class ServiceAttachment(pulumi.CustomResource):
    @overload
    def __init__(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 connection_preference: Optional[pulumi.Input[str]] = None,
                 consumer_forwarding_rules: Optional[pulumi.Input[Sequence[pulumi.Input[pulumi.InputType['ServiceAttachmentConsumerForwardingRuleArgs']]]]] = None,
                 creation_timestamp: Optional[pulumi.Input[str]] = None,
                 description: Optional[pulumi.Input[str]] = None,
                 enable_proxy_protocol: Optional[pulumi.Input[bool]] = None,
                 id: Optional[pulumi.Input[str]] = None,
                 kind: Optional[pulumi.Input[str]] = None,
                 name: Optional[pulumi.Input[str]] = None,
                 nat_subnets: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]] = None,
                 producer_forwarding_rule: Optional[pulumi.Input[str]] = None,
                 project: Optional[pulumi.Input[str]] = None,
                 region: Optional[pulumi.Input[str]] = None,
                 self_link: Optional[pulumi.Input[str]] = None,
                 service_attachment: Optional[pulumi.Input[str]] = None,
                 __props__=None):
        """
        Creates a ServiceAttachment in the specified project in the given scope using the parameters that are included in the request.

        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] connection_preference: The connection preference of service attachment. The value can be set to ACCEPT_AUTOMATIC. An ACCEPT_AUTOMATIC service attachment is one that always accepts the connection from consumer forwarding rules.
        :param pulumi.Input[Sequence[pulumi.Input[pulumi.InputType['ServiceAttachmentConsumerForwardingRuleArgs']]]] consumer_forwarding_rules: [Output Only] An array of forwarding rules for all the consumers connected to this service attachment.
        :param pulumi.Input[str] creation_timestamp: [Output Only] Creation timestamp in RFC3339 text format.
        :param pulumi.Input[str] description: An optional description of this resource. Provide this property when you create the resource.
        :param pulumi.Input[bool] enable_proxy_protocol: If true, enable the proxy protocol which is for supplying client TCP/IP address data in TCP connections that traverse proxies on their way to destination servers.
        :param pulumi.Input[str] id: [Output Only] The unique identifier for the resource type. The server generates this identifier.
        :param pulumi.Input[str] kind: [Output Only] Type of the resource. Always compute#serviceAttachment for service attachments.
        :param pulumi.Input[str] name: Name of the resource. Provided by the client when the resource is created. The name must be 1-63 characters long, and comply with RFC1035. Specifically, the name must be 1-63 characters long and match the regular expression `[a-z]([-a-z0-9]*[a-z0-9])?` which means the first character must be a lowercase letter, and all following characters must be a dash, lowercase letter, or digit, except the last character, which cannot be a dash.
        :param pulumi.Input[Sequence[pulumi.Input[str]]] nat_subnets: An array of URLs where each entry is the URL of a subnet provided by the service producer to use for NAT in this service attachment.
        :param pulumi.Input[str] producer_forwarding_rule: The URL of a forwarding rule with loadBalancingScheme INTERNAL* that is serving the endpoint identified by this service attachment.
        :param pulumi.Input[str] region: [Output Only] URL of the region where the service attachment resides. This field applies only to the region resource. You must specify this field as part of the HTTP request URL. It is not settable as a field in the request body.
        :param pulumi.Input[str] self_link: [Output Only] Server-defined URL for the resource.
        """
        ...
    @overload
    def __init__(__self__,
                 resource_name: str,
                 args: ServiceAttachmentArgs,
                 opts: Optional[pulumi.ResourceOptions] = None):
        """
        Creates a ServiceAttachment in the specified project in the given scope using the parameters that are included in the request.

        :param str resource_name: The name of the resource.
        :param ServiceAttachmentArgs args: The arguments to use to populate this resource's properties.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        ...
    def __init__(__self__, resource_name: str, *args, **kwargs):
        resource_args, opts = _utilities.get_resource_args_opts(ServiceAttachmentArgs, pulumi.ResourceOptions, *args, **kwargs)
        if resource_args is not None:
            __self__._internal_init(resource_name, opts, **resource_args.__dict__)
        else:
            __self__._internal_init(resource_name, *args, **kwargs)

    def _internal_init(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 connection_preference: Optional[pulumi.Input[str]] = None,
                 consumer_forwarding_rules: Optional[pulumi.Input[Sequence[pulumi.Input[pulumi.InputType['ServiceAttachmentConsumerForwardingRuleArgs']]]]] = None,
                 creation_timestamp: Optional[pulumi.Input[str]] = None,
                 description: Optional[pulumi.Input[str]] = None,
                 enable_proxy_protocol: Optional[pulumi.Input[bool]] = None,
                 id: Optional[pulumi.Input[str]] = None,
                 kind: Optional[pulumi.Input[str]] = None,
                 name: Optional[pulumi.Input[str]] = None,
                 nat_subnets: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]] = None,
                 producer_forwarding_rule: Optional[pulumi.Input[str]] = None,
                 project: Optional[pulumi.Input[str]] = None,
                 region: Optional[pulumi.Input[str]] = None,
                 self_link: Optional[pulumi.Input[str]] = None,
                 service_attachment: Optional[pulumi.Input[str]] = None,
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
            __props__ = ServiceAttachmentArgs.__new__(ServiceAttachmentArgs)

            __props__.__dict__["connection_preference"] = connection_preference
            __props__.__dict__["consumer_forwarding_rules"] = consumer_forwarding_rules
            __props__.__dict__["creation_timestamp"] = creation_timestamp
            __props__.__dict__["description"] = description
            __props__.__dict__["enable_proxy_protocol"] = enable_proxy_protocol
            __props__.__dict__["id"] = id
            __props__.__dict__["kind"] = kind
            __props__.__dict__["name"] = name
            __props__.__dict__["nat_subnets"] = nat_subnets
            __props__.__dict__["producer_forwarding_rule"] = producer_forwarding_rule
            if project is None and not opts.urn:
                raise TypeError("Missing required property 'project'")
            __props__.__dict__["project"] = project
            if region is None and not opts.urn:
                raise TypeError("Missing required property 'region'")
            __props__.__dict__["region"] = region
            __props__.__dict__["self_link"] = self_link
            if service_attachment is None and not opts.urn:
                raise TypeError("Missing required property 'service_attachment'")
            __props__.__dict__["service_attachment"] = service_attachment
        super(ServiceAttachment, __self__).__init__(
            'gcp-native:compute/beta:ServiceAttachment',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name: str,
            id: pulumi.Input[str],
            opts: Optional[pulumi.ResourceOptions] = None) -> 'ServiceAttachment':
        """
        Get an existing ServiceAttachment resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param pulumi.Input[str] id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = ServiceAttachmentArgs.__new__(ServiceAttachmentArgs)

        __props__.__dict__["connection_preference"] = None
        __props__.__dict__["consumer_forwarding_rules"] = None
        __props__.__dict__["creation_timestamp"] = None
        __props__.__dict__["description"] = None
        __props__.__dict__["enable_proxy_protocol"] = None
        __props__.__dict__["kind"] = None
        __props__.__dict__["name"] = None
        __props__.__dict__["nat_subnets"] = None
        __props__.__dict__["producer_forwarding_rule"] = None
        __props__.__dict__["region"] = None
        __props__.__dict__["self_link"] = None
        return ServiceAttachment(resource_name, opts=opts, __props__=__props__)

    @property
    @pulumi.getter(name="connectionPreference")
    def connection_preference(self) -> pulumi.Output[str]:
        """
        The connection preference of service attachment. The value can be set to ACCEPT_AUTOMATIC. An ACCEPT_AUTOMATIC service attachment is one that always accepts the connection from consumer forwarding rules.
        """
        return pulumi.get(self, "connection_preference")

    @property
    @pulumi.getter(name="consumerForwardingRules")
    def consumer_forwarding_rules(self) -> pulumi.Output[Sequence['outputs.ServiceAttachmentConsumerForwardingRuleResponse']]:
        """
        [Output Only] An array of forwarding rules for all the consumers connected to this service attachment.
        """
        return pulumi.get(self, "consumer_forwarding_rules")

    @property
    @pulumi.getter(name="creationTimestamp")
    def creation_timestamp(self) -> pulumi.Output[str]:
        """
        [Output Only] Creation timestamp in RFC3339 text format.
        """
        return pulumi.get(self, "creation_timestamp")

    @property
    @pulumi.getter
    def description(self) -> pulumi.Output[str]:
        """
        An optional description of this resource. Provide this property when you create the resource.
        """
        return pulumi.get(self, "description")

    @property
    @pulumi.getter(name="enableProxyProtocol")
    def enable_proxy_protocol(self) -> pulumi.Output[bool]:
        """
        If true, enable the proxy protocol which is for supplying client TCP/IP address data in TCP connections that traverse proxies on their way to destination servers.
        """
        return pulumi.get(self, "enable_proxy_protocol")

    @property
    @pulumi.getter
    def kind(self) -> pulumi.Output[str]:
        """
        [Output Only] Type of the resource. Always compute#serviceAttachment for service attachments.
        """
        return pulumi.get(self, "kind")

    @property
    @pulumi.getter
    def name(self) -> pulumi.Output[str]:
        """
        Name of the resource. Provided by the client when the resource is created. The name must be 1-63 characters long, and comply with RFC1035. Specifically, the name must be 1-63 characters long and match the regular expression `[a-z]([-a-z0-9]*[a-z0-9])?` which means the first character must be a lowercase letter, and all following characters must be a dash, lowercase letter, or digit, except the last character, which cannot be a dash.
        """
        return pulumi.get(self, "name")

    @property
    @pulumi.getter(name="natSubnets")
    def nat_subnets(self) -> pulumi.Output[Sequence[str]]:
        """
        An array of URLs where each entry is the URL of a subnet provided by the service producer to use for NAT in this service attachment.
        """
        return pulumi.get(self, "nat_subnets")

    @property
    @pulumi.getter(name="producerForwardingRule")
    def producer_forwarding_rule(self) -> pulumi.Output[str]:
        """
        The URL of a forwarding rule with loadBalancingScheme INTERNAL* that is serving the endpoint identified by this service attachment.
        """
        return pulumi.get(self, "producer_forwarding_rule")

    @property
    @pulumi.getter
    def region(self) -> pulumi.Output[str]:
        """
        [Output Only] URL of the region where the service attachment resides. This field applies only to the region resource. You must specify this field as part of the HTTP request URL. It is not settable as a field in the request body.
        """
        return pulumi.get(self, "region")

    @property
    @pulumi.getter(name="selfLink")
    def self_link(self) -> pulumi.Output[str]:
        """
        [Output Only] Server-defined URL for the resource.
        """
        return pulumi.get(self, "self_link")

