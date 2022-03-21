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

__all__ = ['PolicyArgs', 'Policy']

@pulumi.input_type
class PolicyArgs:
    def __init__(__self__, *,
                 alternative_name_server_config: Optional[pulumi.Input['PolicyAlternativeNameServerConfigArgs']] = None,
                 client_operation_id: Optional[pulumi.Input[str]] = None,
                 description: Optional[pulumi.Input[str]] = None,
                 enable_inbound_forwarding: Optional[pulumi.Input[bool]] = None,
                 enable_logging: Optional[pulumi.Input[bool]] = None,
                 id: Optional[pulumi.Input[str]] = None,
                 kind: Optional[pulumi.Input[str]] = None,
                 location: Optional[pulumi.Input[str]] = None,
                 name: Optional[pulumi.Input[str]] = None,
                 networks: Optional[pulumi.Input[Sequence[pulumi.Input['PolicyNetworkArgs']]]] = None,
                 project: Optional[pulumi.Input[str]] = None):
        """
        The set of arguments for constructing a Policy resource.
        :param pulumi.Input['PolicyAlternativeNameServerConfigArgs'] alternative_name_server_config: Sets an alternative name server for the associated networks. When specified, all DNS queries are forwarded to a name server that you choose. Names such as .internal are not available when an alternative name server is specified.
        :param pulumi.Input[str] client_operation_id: For mutating operation requests only. An optional identifier specified by the client. Must be unique for operation resources in the Operations collection.
        :param pulumi.Input[str] description: A mutable string of at most 1024 characters associated with this resource for the user's convenience. Has no effect on the policy's function.
        :param pulumi.Input[bool] enable_inbound_forwarding: Allows networks bound to this policy to receive DNS queries sent by VMs or applications over VPN connections. When enabled, a virtual IP address is allocated from each of the subnetworks that are bound to this policy.
        :param pulumi.Input[bool] enable_logging: Controls whether logging is enabled for the networks bound to this policy. Defaults to no logging if not set.
        :param pulumi.Input[str] id: Unique identifier for the resource; defined by the server (output only).
        :param pulumi.Input[str] name: User-assigned name for this policy.
        :param pulumi.Input[Sequence[pulumi.Input['PolicyNetworkArgs']]] networks: List of network names specifying networks to which this policy is applied.
        """
        if alternative_name_server_config is not None:
            pulumi.set(__self__, "alternative_name_server_config", alternative_name_server_config)
        if client_operation_id is not None:
            pulumi.set(__self__, "client_operation_id", client_operation_id)
        if description is not None:
            pulumi.set(__self__, "description", description)
        if enable_inbound_forwarding is not None:
            pulumi.set(__self__, "enable_inbound_forwarding", enable_inbound_forwarding)
        if enable_logging is not None:
            pulumi.set(__self__, "enable_logging", enable_logging)
        if id is not None:
            pulumi.set(__self__, "id", id)
        if kind is not None:
            pulumi.set(__self__, "kind", kind)
        if location is not None:
            pulumi.set(__self__, "location", location)
        if name is not None:
            pulumi.set(__self__, "name", name)
        if networks is not None:
            pulumi.set(__self__, "networks", networks)
        if project is not None:
            pulumi.set(__self__, "project", project)

    @property
    @pulumi.getter(name="alternativeNameServerConfig")
    def alternative_name_server_config(self) -> Optional[pulumi.Input['PolicyAlternativeNameServerConfigArgs']]:
        """
        Sets an alternative name server for the associated networks. When specified, all DNS queries are forwarded to a name server that you choose. Names such as .internal are not available when an alternative name server is specified.
        """
        return pulumi.get(self, "alternative_name_server_config")

    @alternative_name_server_config.setter
    def alternative_name_server_config(self, value: Optional[pulumi.Input['PolicyAlternativeNameServerConfigArgs']]):
        pulumi.set(self, "alternative_name_server_config", value)

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
    @pulumi.getter
    def description(self) -> Optional[pulumi.Input[str]]:
        """
        A mutable string of at most 1024 characters associated with this resource for the user's convenience. Has no effect on the policy's function.
        """
        return pulumi.get(self, "description")

    @description.setter
    def description(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "description", value)

    @property
    @pulumi.getter(name="enableInboundForwarding")
    def enable_inbound_forwarding(self) -> Optional[pulumi.Input[bool]]:
        """
        Allows networks bound to this policy to receive DNS queries sent by VMs or applications over VPN connections. When enabled, a virtual IP address is allocated from each of the subnetworks that are bound to this policy.
        """
        return pulumi.get(self, "enable_inbound_forwarding")

    @enable_inbound_forwarding.setter
    def enable_inbound_forwarding(self, value: Optional[pulumi.Input[bool]]):
        pulumi.set(self, "enable_inbound_forwarding", value)

    @property
    @pulumi.getter(name="enableLogging")
    def enable_logging(self) -> Optional[pulumi.Input[bool]]:
        """
        Controls whether logging is enabled for the networks bound to this policy. Defaults to no logging if not set.
        """
        return pulumi.get(self, "enable_logging")

    @enable_logging.setter
    def enable_logging(self, value: Optional[pulumi.Input[bool]]):
        pulumi.set(self, "enable_logging", value)

    @property
    @pulumi.getter
    def id(self) -> Optional[pulumi.Input[str]]:
        """
        Unique identifier for the resource; defined by the server (output only).
        """
        return pulumi.get(self, "id")

    @id.setter
    def id(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "id", value)

    @property
    @pulumi.getter
    def kind(self) -> Optional[pulumi.Input[str]]:
        return pulumi.get(self, "kind")

    @kind.setter
    def kind(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "kind", value)

    @property
    @pulumi.getter
    def location(self) -> Optional[pulumi.Input[str]]:
        return pulumi.get(self, "location")

    @location.setter
    def location(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "location", value)

    @property
    @pulumi.getter
    def name(self) -> Optional[pulumi.Input[str]]:
        """
        User-assigned name for this policy.
        """
        return pulumi.get(self, "name")

    @name.setter
    def name(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "name", value)

    @property
    @pulumi.getter
    def networks(self) -> Optional[pulumi.Input[Sequence[pulumi.Input['PolicyNetworkArgs']]]]:
        """
        List of network names specifying networks to which this policy is applied.
        """
        return pulumi.get(self, "networks")

    @networks.setter
    def networks(self, value: Optional[pulumi.Input[Sequence[pulumi.Input['PolicyNetworkArgs']]]]):
        pulumi.set(self, "networks", value)

    @property
    @pulumi.getter
    def project(self) -> Optional[pulumi.Input[str]]:
        return pulumi.get(self, "project")

    @project.setter
    def project(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "project", value)


class Policy(pulumi.CustomResource):
    @overload
    def __init__(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 alternative_name_server_config: Optional[pulumi.Input[pulumi.InputType['PolicyAlternativeNameServerConfigArgs']]] = None,
                 client_operation_id: Optional[pulumi.Input[str]] = None,
                 description: Optional[pulumi.Input[str]] = None,
                 enable_inbound_forwarding: Optional[pulumi.Input[bool]] = None,
                 enable_logging: Optional[pulumi.Input[bool]] = None,
                 id: Optional[pulumi.Input[str]] = None,
                 kind: Optional[pulumi.Input[str]] = None,
                 location: Optional[pulumi.Input[str]] = None,
                 name: Optional[pulumi.Input[str]] = None,
                 networks: Optional[pulumi.Input[Sequence[pulumi.Input[pulumi.InputType['PolicyNetworkArgs']]]]] = None,
                 project: Optional[pulumi.Input[str]] = None,
                 __props__=None):
        """
        Creates a new Policy.

        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[pulumi.InputType['PolicyAlternativeNameServerConfigArgs']] alternative_name_server_config: Sets an alternative name server for the associated networks. When specified, all DNS queries are forwarded to a name server that you choose. Names such as .internal are not available when an alternative name server is specified.
        :param pulumi.Input[str] client_operation_id: For mutating operation requests only. An optional identifier specified by the client. Must be unique for operation resources in the Operations collection.
        :param pulumi.Input[str] description: A mutable string of at most 1024 characters associated with this resource for the user's convenience. Has no effect on the policy's function.
        :param pulumi.Input[bool] enable_inbound_forwarding: Allows networks bound to this policy to receive DNS queries sent by VMs or applications over VPN connections. When enabled, a virtual IP address is allocated from each of the subnetworks that are bound to this policy.
        :param pulumi.Input[bool] enable_logging: Controls whether logging is enabled for the networks bound to this policy. Defaults to no logging if not set.
        :param pulumi.Input[str] id: Unique identifier for the resource; defined by the server (output only).
        :param pulumi.Input[str] name: User-assigned name for this policy.
        :param pulumi.Input[Sequence[pulumi.Input[pulumi.InputType['PolicyNetworkArgs']]]] networks: List of network names specifying networks to which this policy is applied.
        """
        ...
    @overload
    def __init__(__self__,
                 resource_name: str,
                 args: Optional[PolicyArgs] = None,
                 opts: Optional[pulumi.ResourceOptions] = None):
        """
        Creates a new Policy.

        :param str resource_name: The name of the resource.
        :param PolicyArgs args: The arguments to use to populate this resource's properties.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        ...
    def __init__(__self__, resource_name: str, *args, **kwargs):
        resource_args, opts = _utilities.get_resource_args_opts(PolicyArgs, pulumi.ResourceOptions, *args, **kwargs)
        if resource_args is not None:
            __self__._internal_init(resource_name, opts, **resource_args.__dict__)
        else:
            __self__._internal_init(resource_name, *args, **kwargs)

    def _internal_init(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 alternative_name_server_config: Optional[pulumi.Input[pulumi.InputType['PolicyAlternativeNameServerConfigArgs']]] = None,
                 client_operation_id: Optional[pulumi.Input[str]] = None,
                 description: Optional[pulumi.Input[str]] = None,
                 enable_inbound_forwarding: Optional[pulumi.Input[bool]] = None,
                 enable_logging: Optional[pulumi.Input[bool]] = None,
                 id: Optional[pulumi.Input[str]] = None,
                 kind: Optional[pulumi.Input[str]] = None,
                 location: Optional[pulumi.Input[str]] = None,
                 name: Optional[pulumi.Input[str]] = None,
                 networks: Optional[pulumi.Input[Sequence[pulumi.Input[pulumi.InputType['PolicyNetworkArgs']]]]] = None,
                 project: Optional[pulumi.Input[str]] = None,
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
            __props__ = PolicyArgs.__new__(PolicyArgs)

            __props__.__dict__["alternative_name_server_config"] = alternative_name_server_config
            __props__.__dict__["client_operation_id"] = client_operation_id
            __props__.__dict__["description"] = description
            __props__.__dict__["enable_inbound_forwarding"] = enable_inbound_forwarding
            __props__.__dict__["enable_logging"] = enable_logging
            __props__.__dict__["id"] = id
            __props__.__dict__["kind"] = kind
            __props__.__dict__["location"] = location
            __props__.__dict__["name"] = name
            __props__.__dict__["networks"] = networks
            __props__.__dict__["project"] = project
        super(Policy, __self__).__init__(
            'google-native:dns/v2:Policy',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name: str,
            id: pulumi.Input[str],
            opts: Optional[pulumi.ResourceOptions] = None) -> 'Policy':
        """
        Get an existing Policy resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param pulumi.Input[str] id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = PolicyArgs.__new__(PolicyArgs)

        __props__.__dict__["alternative_name_server_config"] = None
        __props__.__dict__["description"] = None
        __props__.__dict__["enable_inbound_forwarding"] = None
        __props__.__dict__["enable_logging"] = None
        __props__.__dict__["kind"] = None
        __props__.__dict__["name"] = None
        __props__.__dict__["networks"] = None
        return Policy(resource_name, opts=opts, __props__=__props__)

    @property
    @pulumi.getter(name="alternativeNameServerConfig")
    def alternative_name_server_config(self) -> pulumi.Output['outputs.PolicyAlternativeNameServerConfigResponse']:
        """
        Sets an alternative name server for the associated networks. When specified, all DNS queries are forwarded to a name server that you choose. Names such as .internal are not available when an alternative name server is specified.
        """
        return pulumi.get(self, "alternative_name_server_config")

    @property
    @pulumi.getter
    def description(self) -> pulumi.Output[str]:
        """
        A mutable string of at most 1024 characters associated with this resource for the user's convenience. Has no effect on the policy's function.
        """
        return pulumi.get(self, "description")

    @property
    @pulumi.getter(name="enableInboundForwarding")
    def enable_inbound_forwarding(self) -> pulumi.Output[bool]:
        """
        Allows networks bound to this policy to receive DNS queries sent by VMs or applications over VPN connections. When enabled, a virtual IP address is allocated from each of the subnetworks that are bound to this policy.
        """
        return pulumi.get(self, "enable_inbound_forwarding")

    @property
    @pulumi.getter(name="enableLogging")
    def enable_logging(self) -> pulumi.Output[bool]:
        """
        Controls whether logging is enabled for the networks bound to this policy. Defaults to no logging if not set.
        """
        return pulumi.get(self, "enable_logging")

    @property
    @pulumi.getter
    def kind(self) -> pulumi.Output[str]:
        return pulumi.get(self, "kind")

    @property
    @pulumi.getter
    def name(self) -> pulumi.Output[str]:
        """
        User-assigned name for this policy.
        """
        return pulumi.get(self, "name")

    @property
    @pulumi.getter
    def networks(self) -> pulumi.Output[Sequence['outputs.PolicyNetworkResponse']]:
        """
        List of network names specifying networks to which this policy is applied.
        """
        return pulumi.get(self, "networks")

