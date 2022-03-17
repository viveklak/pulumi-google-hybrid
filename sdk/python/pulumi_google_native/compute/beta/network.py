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

__all__ = ['NetworkArgs', 'Network']

@pulumi.input_type
class NetworkArgs:
    def __init__(__self__, *,
                 auto_create_subnetworks: Optional[pulumi.Input[bool]] = None,
                 description: Optional[pulumi.Input[str]] = None,
                 enable_ula_internal_ipv6: Optional[pulumi.Input[bool]] = None,
                 internal_ipv6_range: Optional[pulumi.Input[str]] = None,
                 ipv4_range: Optional[pulumi.Input[str]] = None,
                 mtu: Optional[pulumi.Input[int]] = None,
                 name: Optional[pulumi.Input[str]] = None,
                 network_firewall_policy_enforcement_order: Optional[pulumi.Input['NetworkNetworkFirewallPolicyEnforcementOrder']] = None,
                 project: Optional[pulumi.Input[str]] = None,
                 request_id: Optional[pulumi.Input[str]] = None,
                 routing_config: Optional[pulumi.Input['NetworkRoutingConfigArgs']] = None):
        """
        The set of arguments for constructing a Network resource.
        :param pulumi.Input[bool] auto_create_subnetworks: Must be set to create a VPC network. If not set, a legacy network is created. When set to true, the VPC network is created in auto mode. When set to false, the VPC network is created in custom mode. An auto mode VPC network starts with one subnet per region. Each subnet has a predetermined range as described in Auto mode VPC network IP ranges. For custom mode VPC networks, you can add subnets using the subnetworks insert method.
        :param pulumi.Input[str] description: An optional description of this resource. Provide this field when you create the resource.
        :param pulumi.Input[bool] enable_ula_internal_ipv6: Enable ULA internal ipv6 on this network. Enabling this feature will assign a /48 from google defined ULA prefix fd20::/20. .
        :param pulumi.Input[str] internal_ipv6_range: When enabling ula internal ipv6, caller optionally can specify the /48 range they want from the google defined ULA prefix fd20::/20. The input must be a valid /48 ULA IPv6 address and must be within the fd20::/20. Operation will fail if the speficied /48 is already in used by another resource. If the field is not speficied, then a /48 range will be randomly allocated from fd20::/20 and returned via this field. .
        :param pulumi.Input[str] ipv4_range: Deprecated in favor of subnet mode networks. The range of internal addresses that are legal on this network. This range is a CIDR specification, for example: 192.168.0.0/16. Provided by the client when the network is created.
        :param pulumi.Input[int] mtu: Maximum Transmission Unit in bytes. The minimum value for this field is 1460 and the maximum value is 1500 bytes. If unspecified, defaults to 1460.
        :param pulumi.Input[str] name: Name of the resource. Provided by the client when the resource is created. The name must be 1-63 characters long, and comply with RFC1035. Specifically, the name must be 1-63 characters long and match the regular expression `[a-z]([-a-z0-9]*[a-z0-9])?`. The first character must be a lowercase letter, and all following characters (except for the last character) must be a dash, lowercase letter, or digit. The last character must be a lowercase letter or digit.
        :param pulumi.Input['NetworkNetworkFirewallPolicyEnforcementOrder'] network_firewall_policy_enforcement_order: The network firewall policy enforcement order. Can be either AFTER_CLASSIC_FIREWALL or BEFORE_CLASSIC_FIREWALL. Defaults to AFTER_CLASSIC_FIREWALL if the field is not specified.
        :param pulumi.Input[str] request_id: An optional request ID to identify requests. Specify a unique request ID so that if you must retry your request, the server will know to ignore the request if it has already been completed. For example, consider a situation where you make an initial request and the request times out. If you make the request again with the same request ID, the server can check if original operation with the same request ID was received, and if so, will ignore the second request. This prevents clients from accidentally creating duplicate commitments. The request ID must be a valid UUID with the exception that zero UUID is not supported ( 00000000-0000-0000-0000-000000000000).
        :param pulumi.Input['NetworkRoutingConfigArgs'] routing_config: The network-level routing configuration for this network. Used by Cloud Router to determine what type of network-wide routing behavior to enforce.
        """
        if auto_create_subnetworks is not None:
            pulumi.set(__self__, "auto_create_subnetworks", auto_create_subnetworks)
        if description is not None:
            pulumi.set(__self__, "description", description)
        if enable_ula_internal_ipv6 is not None:
            pulumi.set(__self__, "enable_ula_internal_ipv6", enable_ula_internal_ipv6)
        if internal_ipv6_range is not None:
            pulumi.set(__self__, "internal_ipv6_range", internal_ipv6_range)
        if ipv4_range is not None:
            warnings.warn("""Deprecated in favor of subnet mode networks. The range of internal addresses that are legal on this network. This range is a CIDR specification, for example: 192.168.0.0/16. Provided by the client when the network is created.""", DeprecationWarning)
            pulumi.log.warn("""ipv4_range is deprecated: Deprecated in favor of subnet mode networks. The range of internal addresses that are legal on this network. This range is a CIDR specification, for example: 192.168.0.0/16. Provided by the client when the network is created.""")
        if ipv4_range is not None:
            pulumi.set(__self__, "ipv4_range", ipv4_range)
        if mtu is not None:
            pulumi.set(__self__, "mtu", mtu)
        if name is not None:
            pulumi.set(__self__, "name", name)
        if network_firewall_policy_enforcement_order is not None:
            pulumi.set(__self__, "network_firewall_policy_enforcement_order", network_firewall_policy_enforcement_order)
        if project is not None:
            pulumi.set(__self__, "project", project)
        if request_id is not None:
            pulumi.set(__self__, "request_id", request_id)
        if routing_config is not None:
            pulumi.set(__self__, "routing_config", routing_config)

    @property
    @pulumi.getter(name="autoCreateSubnetworks")
    def auto_create_subnetworks(self) -> Optional[pulumi.Input[bool]]:
        """
        Must be set to create a VPC network. If not set, a legacy network is created. When set to true, the VPC network is created in auto mode. When set to false, the VPC network is created in custom mode. An auto mode VPC network starts with one subnet per region. Each subnet has a predetermined range as described in Auto mode VPC network IP ranges. For custom mode VPC networks, you can add subnets using the subnetworks insert method.
        """
        return pulumi.get(self, "auto_create_subnetworks")

    @auto_create_subnetworks.setter
    def auto_create_subnetworks(self, value: Optional[pulumi.Input[bool]]):
        pulumi.set(self, "auto_create_subnetworks", value)

    @property
    @pulumi.getter
    def description(self) -> Optional[pulumi.Input[str]]:
        """
        An optional description of this resource. Provide this field when you create the resource.
        """
        return pulumi.get(self, "description")

    @description.setter
    def description(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "description", value)

    @property
    @pulumi.getter(name="enableUlaInternalIpv6")
    def enable_ula_internal_ipv6(self) -> Optional[pulumi.Input[bool]]:
        """
        Enable ULA internal ipv6 on this network. Enabling this feature will assign a /48 from google defined ULA prefix fd20::/20. .
        """
        return pulumi.get(self, "enable_ula_internal_ipv6")

    @enable_ula_internal_ipv6.setter
    def enable_ula_internal_ipv6(self, value: Optional[pulumi.Input[bool]]):
        pulumi.set(self, "enable_ula_internal_ipv6", value)

    @property
    @pulumi.getter(name="internalIpv6Range")
    def internal_ipv6_range(self) -> Optional[pulumi.Input[str]]:
        """
        When enabling ula internal ipv6, caller optionally can specify the /48 range they want from the google defined ULA prefix fd20::/20. The input must be a valid /48 ULA IPv6 address and must be within the fd20::/20. Operation will fail if the speficied /48 is already in used by another resource. If the field is not speficied, then a /48 range will be randomly allocated from fd20::/20 and returned via this field. .
        """
        return pulumi.get(self, "internal_ipv6_range")

    @internal_ipv6_range.setter
    def internal_ipv6_range(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "internal_ipv6_range", value)

    @property
    @pulumi.getter(name="ipv4Range")
    def ipv4_range(self) -> Optional[pulumi.Input[str]]:
        """
        Deprecated in favor of subnet mode networks. The range of internal addresses that are legal on this network. This range is a CIDR specification, for example: 192.168.0.0/16. Provided by the client when the network is created.
        """
        return pulumi.get(self, "ipv4_range")

    @ipv4_range.setter
    def ipv4_range(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "ipv4_range", value)

    @property
    @pulumi.getter
    def mtu(self) -> Optional[pulumi.Input[int]]:
        """
        Maximum Transmission Unit in bytes. The minimum value for this field is 1460 and the maximum value is 1500 bytes. If unspecified, defaults to 1460.
        """
        return pulumi.get(self, "mtu")

    @mtu.setter
    def mtu(self, value: Optional[pulumi.Input[int]]):
        pulumi.set(self, "mtu", value)

    @property
    @pulumi.getter
    def name(self) -> Optional[pulumi.Input[str]]:
        """
        Name of the resource. Provided by the client when the resource is created. The name must be 1-63 characters long, and comply with RFC1035. Specifically, the name must be 1-63 characters long and match the regular expression `[a-z]([-a-z0-9]*[a-z0-9])?`. The first character must be a lowercase letter, and all following characters (except for the last character) must be a dash, lowercase letter, or digit. The last character must be a lowercase letter or digit.
        """
        return pulumi.get(self, "name")

    @name.setter
    def name(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "name", value)

    @property
    @pulumi.getter(name="networkFirewallPolicyEnforcementOrder")
    def network_firewall_policy_enforcement_order(self) -> Optional[pulumi.Input['NetworkNetworkFirewallPolicyEnforcementOrder']]:
        """
        The network firewall policy enforcement order. Can be either AFTER_CLASSIC_FIREWALL or BEFORE_CLASSIC_FIREWALL. Defaults to AFTER_CLASSIC_FIREWALL if the field is not specified.
        """
        return pulumi.get(self, "network_firewall_policy_enforcement_order")

    @network_firewall_policy_enforcement_order.setter
    def network_firewall_policy_enforcement_order(self, value: Optional[pulumi.Input['NetworkNetworkFirewallPolicyEnforcementOrder']]):
        pulumi.set(self, "network_firewall_policy_enforcement_order", value)

    @property
    @pulumi.getter
    def project(self) -> Optional[pulumi.Input[str]]:
        return pulumi.get(self, "project")

    @project.setter
    def project(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "project", value)

    @property
    @pulumi.getter(name="requestId")
    def request_id(self) -> Optional[pulumi.Input[str]]:
        """
        An optional request ID to identify requests. Specify a unique request ID so that if you must retry your request, the server will know to ignore the request if it has already been completed. For example, consider a situation where you make an initial request and the request times out. If you make the request again with the same request ID, the server can check if original operation with the same request ID was received, and if so, will ignore the second request. This prevents clients from accidentally creating duplicate commitments. The request ID must be a valid UUID with the exception that zero UUID is not supported ( 00000000-0000-0000-0000-000000000000).
        """
        return pulumi.get(self, "request_id")

    @request_id.setter
    def request_id(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "request_id", value)

    @property
    @pulumi.getter(name="routingConfig")
    def routing_config(self) -> Optional[pulumi.Input['NetworkRoutingConfigArgs']]:
        """
        The network-level routing configuration for this network. Used by Cloud Router to determine what type of network-wide routing behavior to enforce.
        """
        return pulumi.get(self, "routing_config")

    @routing_config.setter
    def routing_config(self, value: Optional[pulumi.Input['NetworkRoutingConfigArgs']]):
        pulumi.set(self, "routing_config", value)


class Network(pulumi.CustomResource):
    @overload
    def __init__(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 auto_create_subnetworks: Optional[pulumi.Input[bool]] = None,
                 description: Optional[pulumi.Input[str]] = None,
                 enable_ula_internal_ipv6: Optional[pulumi.Input[bool]] = None,
                 internal_ipv6_range: Optional[pulumi.Input[str]] = None,
                 ipv4_range: Optional[pulumi.Input[str]] = None,
                 mtu: Optional[pulumi.Input[int]] = None,
                 name: Optional[pulumi.Input[str]] = None,
                 network_firewall_policy_enforcement_order: Optional[pulumi.Input['NetworkNetworkFirewallPolicyEnforcementOrder']] = None,
                 project: Optional[pulumi.Input[str]] = None,
                 request_id: Optional[pulumi.Input[str]] = None,
                 routing_config: Optional[pulumi.Input[pulumi.InputType['NetworkRoutingConfigArgs']]] = None,
                 __props__=None):
        """
        Creates a network in the specified project using the data included in the request.

        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[bool] auto_create_subnetworks: Must be set to create a VPC network. If not set, a legacy network is created. When set to true, the VPC network is created in auto mode. When set to false, the VPC network is created in custom mode. An auto mode VPC network starts with one subnet per region. Each subnet has a predetermined range as described in Auto mode VPC network IP ranges. For custom mode VPC networks, you can add subnets using the subnetworks insert method.
        :param pulumi.Input[str] description: An optional description of this resource. Provide this field when you create the resource.
        :param pulumi.Input[bool] enable_ula_internal_ipv6: Enable ULA internal ipv6 on this network. Enabling this feature will assign a /48 from google defined ULA prefix fd20::/20. .
        :param pulumi.Input[str] internal_ipv6_range: When enabling ula internal ipv6, caller optionally can specify the /48 range they want from the google defined ULA prefix fd20::/20. The input must be a valid /48 ULA IPv6 address and must be within the fd20::/20. Operation will fail if the speficied /48 is already in used by another resource. If the field is not speficied, then a /48 range will be randomly allocated from fd20::/20 and returned via this field. .
        :param pulumi.Input[str] ipv4_range: Deprecated in favor of subnet mode networks. The range of internal addresses that are legal on this network. This range is a CIDR specification, for example: 192.168.0.0/16. Provided by the client when the network is created.
        :param pulumi.Input[int] mtu: Maximum Transmission Unit in bytes. The minimum value for this field is 1460 and the maximum value is 1500 bytes. If unspecified, defaults to 1460.
        :param pulumi.Input[str] name: Name of the resource. Provided by the client when the resource is created. The name must be 1-63 characters long, and comply with RFC1035. Specifically, the name must be 1-63 characters long and match the regular expression `[a-z]([-a-z0-9]*[a-z0-9])?`. The first character must be a lowercase letter, and all following characters (except for the last character) must be a dash, lowercase letter, or digit. The last character must be a lowercase letter or digit.
        :param pulumi.Input['NetworkNetworkFirewallPolicyEnforcementOrder'] network_firewall_policy_enforcement_order: The network firewall policy enforcement order. Can be either AFTER_CLASSIC_FIREWALL or BEFORE_CLASSIC_FIREWALL. Defaults to AFTER_CLASSIC_FIREWALL if the field is not specified.
        :param pulumi.Input[str] request_id: An optional request ID to identify requests. Specify a unique request ID so that if you must retry your request, the server will know to ignore the request if it has already been completed. For example, consider a situation where you make an initial request and the request times out. If you make the request again with the same request ID, the server can check if original operation with the same request ID was received, and if so, will ignore the second request. This prevents clients from accidentally creating duplicate commitments. The request ID must be a valid UUID with the exception that zero UUID is not supported ( 00000000-0000-0000-0000-000000000000).
        :param pulumi.Input[pulumi.InputType['NetworkRoutingConfigArgs']] routing_config: The network-level routing configuration for this network. Used by Cloud Router to determine what type of network-wide routing behavior to enforce.
        """
        ...
    @overload
    def __init__(__self__,
                 resource_name: str,
                 args: Optional[NetworkArgs] = None,
                 opts: Optional[pulumi.ResourceOptions] = None):
        """
        Creates a network in the specified project using the data included in the request.

        :param str resource_name: The name of the resource.
        :param NetworkArgs args: The arguments to use to populate this resource's properties.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        ...
    def __init__(__self__, resource_name: str, *args, **kwargs):
        resource_args, opts = _utilities.get_resource_args_opts(NetworkArgs, pulumi.ResourceOptions, *args, **kwargs)
        if resource_args is not None:
            __self__._internal_init(resource_name, opts, **resource_args.__dict__)
        else:
            __self__._internal_init(resource_name, *args, **kwargs)

    def _internal_init(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 auto_create_subnetworks: Optional[pulumi.Input[bool]] = None,
                 description: Optional[pulumi.Input[str]] = None,
                 enable_ula_internal_ipv6: Optional[pulumi.Input[bool]] = None,
                 internal_ipv6_range: Optional[pulumi.Input[str]] = None,
                 ipv4_range: Optional[pulumi.Input[str]] = None,
                 mtu: Optional[pulumi.Input[int]] = None,
                 name: Optional[pulumi.Input[str]] = None,
                 network_firewall_policy_enforcement_order: Optional[pulumi.Input['NetworkNetworkFirewallPolicyEnforcementOrder']] = None,
                 project: Optional[pulumi.Input[str]] = None,
                 request_id: Optional[pulumi.Input[str]] = None,
                 routing_config: Optional[pulumi.Input[pulumi.InputType['NetworkRoutingConfigArgs']]] = None,
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
            __props__ = NetworkArgs.__new__(NetworkArgs)

            __props__.__dict__["auto_create_subnetworks"] = auto_create_subnetworks
            __props__.__dict__["description"] = description
            __props__.__dict__["enable_ula_internal_ipv6"] = enable_ula_internal_ipv6
            __props__.__dict__["internal_ipv6_range"] = internal_ipv6_range
            if ipv4_range is not None and not opts.urn:
                warnings.warn("""Deprecated in favor of subnet mode networks. The range of internal addresses that are legal on this network. This range is a CIDR specification, for example: 192.168.0.0/16. Provided by the client when the network is created.""", DeprecationWarning)
                pulumi.log.warn("""ipv4_range is deprecated: Deprecated in favor of subnet mode networks. The range of internal addresses that are legal on this network. This range is a CIDR specification, for example: 192.168.0.0/16. Provided by the client when the network is created.""")
            __props__.__dict__["ipv4_range"] = ipv4_range
            __props__.__dict__["mtu"] = mtu
            __props__.__dict__["name"] = name
            __props__.__dict__["network_firewall_policy_enforcement_order"] = network_firewall_policy_enforcement_order
            __props__.__dict__["project"] = project
            __props__.__dict__["request_id"] = request_id
            __props__.__dict__["routing_config"] = routing_config
            __props__.__dict__["creation_timestamp"] = None
            __props__.__dict__["gateway_i_pv4"] = None
            __props__.__dict__["kind"] = None
            __props__.__dict__["peerings"] = None
            __props__.__dict__["self_link"] = None
            __props__.__dict__["self_link_with_id"] = None
            __props__.__dict__["subnetworks"] = None
        super(Network, __self__).__init__(
            'google-native:compute/beta:Network',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name: str,
            id: pulumi.Input[str],
            opts: Optional[pulumi.ResourceOptions] = None) -> 'Network':
        """
        Get an existing Network resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param pulumi.Input[str] id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = NetworkArgs.__new__(NetworkArgs)

        __props__.__dict__["auto_create_subnetworks"] = None
        __props__.__dict__["creation_timestamp"] = None
        __props__.__dict__["description"] = None
        __props__.__dict__["enable_ula_internal_ipv6"] = None
        __props__.__dict__["gateway_i_pv4"] = None
        __props__.__dict__["internal_ipv6_range"] = None
        __props__.__dict__["ipv4_range"] = None
        __props__.__dict__["kind"] = None
        __props__.__dict__["mtu"] = None
        __props__.__dict__["name"] = None
        __props__.__dict__["network_firewall_policy_enforcement_order"] = None
        __props__.__dict__["peerings"] = None
        __props__.__dict__["routing_config"] = None
        __props__.__dict__["self_link"] = None
        __props__.__dict__["self_link_with_id"] = None
        __props__.__dict__["subnetworks"] = None
        return Network(resource_name, opts=opts, __props__=__props__)

    @property
    @pulumi.getter(name="autoCreateSubnetworks")
    def auto_create_subnetworks(self) -> pulumi.Output[bool]:
        """
        Must be set to create a VPC network. If not set, a legacy network is created. When set to true, the VPC network is created in auto mode. When set to false, the VPC network is created in custom mode. An auto mode VPC network starts with one subnet per region. Each subnet has a predetermined range as described in Auto mode VPC network IP ranges. For custom mode VPC networks, you can add subnets using the subnetworks insert method.
        """
        return pulumi.get(self, "auto_create_subnetworks")

    @property
    @pulumi.getter(name="creationTimestamp")
    def creation_timestamp(self) -> pulumi.Output[str]:
        """
        Creation timestamp in RFC3339 text format.
        """
        return pulumi.get(self, "creation_timestamp")

    @property
    @pulumi.getter
    def description(self) -> pulumi.Output[str]:
        """
        An optional description of this resource. Provide this field when you create the resource.
        """
        return pulumi.get(self, "description")

    @property
    @pulumi.getter(name="enableUlaInternalIpv6")
    def enable_ula_internal_ipv6(self) -> pulumi.Output[bool]:
        """
        Enable ULA internal ipv6 on this network. Enabling this feature will assign a /48 from google defined ULA prefix fd20::/20. .
        """
        return pulumi.get(self, "enable_ula_internal_ipv6")

    @property
    @pulumi.getter(name="gatewayIPv4")
    def gateway_i_pv4(self) -> pulumi.Output[str]:
        """
        The gateway address for default routing out of the network, selected by GCP.
        """
        return pulumi.get(self, "gateway_i_pv4")

    @property
    @pulumi.getter(name="internalIpv6Range")
    def internal_ipv6_range(self) -> pulumi.Output[str]:
        """
        When enabling ula internal ipv6, caller optionally can specify the /48 range they want from the google defined ULA prefix fd20::/20. The input must be a valid /48 ULA IPv6 address and must be within the fd20::/20. Operation will fail if the speficied /48 is already in used by another resource. If the field is not speficied, then a /48 range will be randomly allocated from fd20::/20 and returned via this field. .
        """
        return pulumi.get(self, "internal_ipv6_range")

    @property
    @pulumi.getter(name="ipv4Range")
    def ipv4_range(self) -> pulumi.Output[str]:
        """
        Deprecated in favor of subnet mode networks. The range of internal addresses that are legal on this network. This range is a CIDR specification, for example: 192.168.0.0/16. Provided by the client when the network is created.
        """
        return pulumi.get(self, "ipv4_range")

    @property
    @pulumi.getter
    def kind(self) -> pulumi.Output[str]:
        """
        Type of the resource. Always compute#network for networks.
        """
        return pulumi.get(self, "kind")

    @property
    @pulumi.getter
    def mtu(self) -> pulumi.Output[int]:
        """
        Maximum Transmission Unit in bytes. The minimum value for this field is 1460 and the maximum value is 1500 bytes. If unspecified, defaults to 1460.
        """
        return pulumi.get(self, "mtu")

    @property
    @pulumi.getter
    def name(self) -> pulumi.Output[str]:
        """
        Name of the resource. Provided by the client when the resource is created. The name must be 1-63 characters long, and comply with RFC1035. Specifically, the name must be 1-63 characters long and match the regular expression `[a-z]([-a-z0-9]*[a-z0-9])?`. The first character must be a lowercase letter, and all following characters (except for the last character) must be a dash, lowercase letter, or digit. The last character must be a lowercase letter or digit.
        """
        return pulumi.get(self, "name")

    @property
    @pulumi.getter(name="networkFirewallPolicyEnforcementOrder")
    def network_firewall_policy_enforcement_order(self) -> pulumi.Output[str]:
        """
        The network firewall policy enforcement order. Can be either AFTER_CLASSIC_FIREWALL or BEFORE_CLASSIC_FIREWALL. Defaults to AFTER_CLASSIC_FIREWALL if the field is not specified.
        """
        return pulumi.get(self, "network_firewall_policy_enforcement_order")

    @property
    @pulumi.getter
    def peerings(self) -> pulumi.Output[Sequence['outputs.NetworkPeeringResponse']]:
        """
        A list of network peerings for the resource.
        """
        return pulumi.get(self, "peerings")

    @property
    @pulumi.getter(name="routingConfig")
    def routing_config(self) -> pulumi.Output['outputs.NetworkRoutingConfigResponse']:
        """
        The network-level routing configuration for this network. Used by Cloud Router to determine what type of network-wide routing behavior to enforce.
        """
        return pulumi.get(self, "routing_config")

    @property
    @pulumi.getter(name="selfLink")
    def self_link(self) -> pulumi.Output[str]:
        """
        Server-defined URL for the resource.
        """
        return pulumi.get(self, "self_link")

    @property
    @pulumi.getter(name="selfLinkWithId")
    def self_link_with_id(self) -> pulumi.Output[str]:
        """
        Server-defined URL for this resource with the resource id.
        """
        return pulumi.get(self, "self_link_with_id")

    @property
    @pulumi.getter
    def subnetworks(self) -> pulumi.Output[Sequence[str]]:
        """
        Server-defined fully-qualified URLs for all subnetworks in this VPC network.
        """
        return pulumi.get(self, "subnetworks")

