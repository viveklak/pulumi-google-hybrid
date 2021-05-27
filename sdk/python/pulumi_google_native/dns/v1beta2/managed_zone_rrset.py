# coding=utf-8
# *** WARNING: this file was generated by the Pulumi SDK Generator. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import warnings
import pulumi
import pulumi.runtime
from typing import Any, Mapping, Optional, Sequence, Union, overload
from ... import _utilities

__all__ = ['ManagedZoneRrsetArgs', 'ManagedZoneRrset']

@pulumi.input_type
class ManagedZoneRrsetArgs:
    def __init__(__self__, *,
                 managed_zone: pulumi.Input[str],
                 project: pulumi.Input[str],
                 client_operation_id: Optional[pulumi.Input[str]] = None,
                 kind: Optional[pulumi.Input[str]] = None,
                 name: Optional[pulumi.Input[str]] = None,
                 rrdatas: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]] = None,
                 signature_rrdatas: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]] = None,
                 ttl: Optional[pulumi.Input[int]] = None,
                 type: Optional[pulumi.Input[str]] = None):
        """
        The set of arguments for constructing a ManagedZoneRrset resource.
        :param pulumi.Input[str] name: For example, www.example.com.
        :param pulumi.Input[Sequence[pulumi.Input[str]]] rrdatas: As defined in RFC 1035 (section 5) and RFC 1034 (section 3.6.1) -- see examples.
        :param pulumi.Input[Sequence[pulumi.Input[str]]] signature_rrdatas: As defined in RFC 4034 (section 3.2).
        :param pulumi.Input[int] ttl: Number of seconds that this ResourceRecordSet can be cached by resolvers.
        :param pulumi.Input[str] type: The identifier of a supported record type. See the list of Supported DNS record types.
        """
        pulumi.set(__self__, "managed_zone", managed_zone)
        pulumi.set(__self__, "project", project)
        if client_operation_id is not None:
            pulumi.set(__self__, "client_operation_id", client_operation_id)
        if kind is not None:
            pulumi.set(__self__, "kind", kind)
        if name is not None:
            pulumi.set(__self__, "name", name)
        if rrdatas is not None:
            pulumi.set(__self__, "rrdatas", rrdatas)
        if signature_rrdatas is not None:
            pulumi.set(__self__, "signature_rrdatas", signature_rrdatas)
        if ttl is not None:
            pulumi.set(__self__, "ttl", ttl)
        if type is not None:
            pulumi.set(__self__, "type", type)

    @property
    @pulumi.getter(name="managedZone")
    def managed_zone(self) -> pulumi.Input[str]:
        return pulumi.get(self, "managed_zone")

    @managed_zone.setter
    def managed_zone(self, value: pulumi.Input[str]):
        pulumi.set(self, "managed_zone", value)

    @property
    @pulumi.getter
    def project(self) -> pulumi.Input[str]:
        return pulumi.get(self, "project")

    @project.setter
    def project(self, value: pulumi.Input[str]):
        pulumi.set(self, "project", value)

    @property
    @pulumi.getter(name="clientOperationId")
    def client_operation_id(self) -> Optional[pulumi.Input[str]]:
        return pulumi.get(self, "client_operation_id")

    @client_operation_id.setter
    def client_operation_id(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "client_operation_id", value)

    @property
    @pulumi.getter
    def kind(self) -> Optional[pulumi.Input[str]]:
        return pulumi.get(self, "kind")

    @kind.setter
    def kind(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "kind", value)

    @property
    @pulumi.getter
    def name(self) -> Optional[pulumi.Input[str]]:
        """
        For example, www.example.com.
        """
        return pulumi.get(self, "name")

    @name.setter
    def name(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "name", value)

    @property
    @pulumi.getter
    def rrdatas(self) -> Optional[pulumi.Input[Sequence[pulumi.Input[str]]]]:
        """
        As defined in RFC 1035 (section 5) and RFC 1034 (section 3.6.1) -- see examples.
        """
        return pulumi.get(self, "rrdatas")

    @rrdatas.setter
    def rrdatas(self, value: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]]):
        pulumi.set(self, "rrdatas", value)

    @property
    @pulumi.getter(name="signatureRrdatas")
    def signature_rrdatas(self) -> Optional[pulumi.Input[Sequence[pulumi.Input[str]]]]:
        """
        As defined in RFC 4034 (section 3.2).
        """
        return pulumi.get(self, "signature_rrdatas")

    @signature_rrdatas.setter
    def signature_rrdatas(self, value: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]]):
        pulumi.set(self, "signature_rrdatas", value)

    @property
    @pulumi.getter
    def ttl(self) -> Optional[pulumi.Input[int]]:
        """
        Number of seconds that this ResourceRecordSet can be cached by resolvers.
        """
        return pulumi.get(self, "ttl")

    @ttl.setter
    def ttl(self, value: Optional[pulumi.Input[int]]):
        pulumi.set(self, "ttl", value)

    @property
    @pulumi.getter
    def type(self) -> Optional[pulumi.Input[str]]:
        """
        The identifier of a supported record type. See the list of Supported DNS record types.
        """
        return pulumi.get(self, "type")

    @type.setter
    def type(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "type", value)


class ManagedZoneRrset(pulumi.CustomResource):
    @overload
    def __init__(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 client_operation_id: Optional[pulumi.Input[str]] = None,
                 kind: Optional[pulumi.Input[str]] = None,
                 managed_zone: Optional[pulumi.Input[str]] = None,
                 name: Optional[pulumi.Input[str]] = None,
                 project: Optional[pulumi.Input[str]] = None,
                 rrdatas: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]] = None,
                 signature_rrdatas: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]] = None,
                 ttl: Optional[pulumi.Input[int]] = None,
                 type: Optional[pulumi.Input[str]] = None,
                 __props__=None):
        """
        Creates a new ResourceRecordSet.

        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] name: For example, www.example.com.
        :param pulumi.Input[Sequence[pulumi.Input[str]]] rrdatas: As defined in RFC 1035 (section 5) and RFC 1034 (section 3.6.1) -- see examples.
        :param pulumi.Input[Sequence[pulumi.Input[str]]] signature_rrdatas: As defined in RFC 4034 (section 3.2).
        :param pulumi.Input[int] ttl: Number of seconds that this ResourceRecordSet can be cached by resolvers.
        :param pulumi.Input[str] type: The identifier of a supported record type. See the list of Supported DNS record types.
        """
        ...
    @overload
    def __init__(__self__,
                 resource_name: str,
                 args: ManagedZoneRrsetArgs,
                 opts: Optional[pulumi.ResourceOptions] = None):
        """
        Creates a new ResourceRecordSet.

        :param str resource_name: The name of the resource.
        :param ManagedZoneRrsetArgs args: The arguments to use to populate this resource's properties.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        ...
    def __init__(__self__, resource_name: str, *args, **kwargs):
        resource_args, opts = _utilities.get_resource_args_opts(ManagedZoneRrsetArgs, pulumi.ResourceOptions, *args, **kwargs)
        if resource_args is not None:
            __self__._internal_init(resource_name, opts, **resource_args.__dict__)
        else:
            __self__._internal_init(resource_name, *args, **kwargs)

    def _internal_init(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 client_operation_id: Optional[pulumi.Input[str]] = None,
                 kind: Optional[pulumi.Input[str]] = None,
                 managed_zone: Optional[pulumi.Input[str]] = None,
                 name: Optional[pulumi.Input[str]] = None,
                 project: Optional[pulumi.Input[str]] = None,
                 rrdatas: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]] = None,
                 signature_rrdatas: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]] = None,
                 ttl: Optional[pulumi.Input[int]] = None,
                 type: Optional[pulumi.Input[str]] = None,
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
            __props__ = ManagedZoneRrsetArgs.__new__(ManagedZoneRrsetArgs)

            __props__.__dict__["client_operation_id"] = client_operation_id
            __props__.__dict__["kind"] = kind
            if managed_zone is None and not opts.urn:
                raise TypeError("Missing required property 'managed_zone'")
            __props__.__dict__["managed_zone"] = managed_zone
            __props__.__dict__["name"] = name
            if project is None and not opts.urn:
                raise TypeError("Missing required property 'project'")
            __props__.__dict__["project"] = project
            __props__.__dict__["rrdatas"] = rrdatas
            __props__.__dict__["signature_rrdatas"] = signature_rrdatas
            __props__.__dict__["ttl"] = ttl
            __props__.__dict__["type"] = type
        super(ManagedZoneRrset, __self__).__init__(
            'google-native:dns/v1beta2:ManagedZoneRrset',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name: str,
            id: pulumi.Input[str],
            opts: Optional[pulumi.ResourceOptions] = None) -> 'ManagedZoneRrset':
        """
        Get an existing ManagedZoneRrset resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param pulumi.Input[str] id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = ManagedZoneRrsetArgs.__new__(ManagedZoneRrsetArgs)

        __props__.__dict__["kind"] = None
        __props__.__dict__["name"] = None
        __props__.__dict__["rrdatas"] = None
        __props__.__dict__["signature_rrdatas"] = None
        __props__.__dict__["ttl"] = None
        __props__.__dict__["type"] = None
        return ManagedZoneRrset(resource_name, opts=opts, __props__=__props__)

    @property
    @pulumi.getter
    def kind(self) -> pulumi.Output[str]:
        return pulumi.get(self, "kind")

    @property
    @pulumi.getter
    def name(self) -> pulumi.Output[str]:
        """
        For example, www.example.com.
        """
        return pulumi.get(self, "name")

    @property
    @pulumi.getter
    def rrdatas(self) -> pulumi.Output[Sequence[str]]:
        """
        As defined in RFC 1035 (section 5) and RFC 1034 (section 3.6.1) -- see examples.
        """
        return pulumi.get(self, "rrdatas")

    @property
    @pulumi.getter(name="signatureRrdatas")
    def signature_rrdatas(self) -> pulumi.Output[Sequence[str]]:
        """
        As defined in RFC 4034 (section 3.2).
        """
        return pulumi.get(self, "signature_rrdatas")

    @property
    @pulumi.getter
    def ttl(self) -> pulumi.Output[int]:
        """
        Number of seconds that this ResourceRecordSet can be cached by resolvers.
        """
        return pulumi.get(self, "ttl")

    @property
    @pulumi.getter
    def type(self) -> pulumi.Output[str]:
        """
        The identifier of a supported record type. See the list of Supported DNS record types.
        """
        return pulumi.get(self, "type")

