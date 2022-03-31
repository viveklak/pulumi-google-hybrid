# coding=utf-8
# *** WARNING: this file was generated by the Pulumi SDK Generator. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import warnings
import pulumi
import pulumi.runtime
from typing import Any, Mapping, Optional, Sequence, Union, overload
from ... import _utilities
from ._enums import *

__all__ = ['InstanceArgs', 'Instance']

@pulumi.input_type
class InstanceArgs:
    def __init__(__self__, *,
                 organization_id: pulumi.Input[str],
                 consumer_accept_list: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]] = None,
                 description: Optional[pulumi.Input[str]] = None,
                 disk_encryption_key_name: Optional[pulumi.Input[str]] = None,
                 display_name: Optional[pulumi.Input[str]] = None,
                 ip_range: Optional[pulumi.Input[str]] = None,
                 location: Optional[pulumi.Input[str]] = None,
                 name: Optional[pulumi.Input[str]] = None,
                 peering_cidr_range: Optional[pulumi.Input['InstancePeeringCidrRange']] = None):
        """
        The set of arguments for constructing a Instance resource.
        :param pulumi.Input[Sequence[pulumi.Input[str]]] consumer_accept_list: Optional. Customer accept list represents the list of projects (id/number) on customer side that can privately connect to the service attachment. It is an optional field which the customers can provide during the instance creation. By default, the customer project associated with the Apigee organization will be included to the list.
        :param pulumi.Input[str] description: Optional. Description of the instance.
        :param pulumi.Input[str] disk_encryption_key_name: Customer Managed Encryption Key (CMEK) used for disk and volume encryption. Required for Apigee paid subscriptions only. Use the following format: `projects/([^/]+)/locations/([^/]+)/keyRings/([^/]+)/cryptoKeys/([^/]+)`
        :param pulumi.Input[str] display_name: Optional. Display name for the instance.
        :param pulumi.Input[str] ip_range: Optional. IP range represents the customer-provided CIDR block of length 22 that will be used for the Apigee instance creation. This optional range, if provided, should be freely available as part of larger named range the customer has allocated to the Service Networking peering. If this is not provided, Apigee will automatically request for any available /22 CIDR block from Service Networking. The customer should use this CIDR block for configuring their firewall needs to allow traffic from Apigee. Input format: "a.b.c.d/22", Output format: a.b.c.d/22, e.f.g.h/28"
        :param pulumi.Input[str] location: Compute Engine location where the instance resides.
        :param pulumi.Input[str] name: Resource ID of the instance. Values must match the regular expression `^a-z{0,30}[a-z\d]$`.
        :param pulumi.Input['InstancePeeringCidrRange'] peering_cidr_range: Optional. Size of the CIDR block range that will be reserved by the instance. PAID organizations support `SLASH_16` to `SLASH_20` and defaults to `SLASH_16`. Evaluation organizations support only `SLASH_23`.
        """
        pulumi.set(__self__, "organization_id", organization_id)
        if consumer_accept_list is not None:
            pulumi.set(__self__, "consumer_accept_list", consumer_accept_list)
        if description is not None:
            pulumi.set(__self__, "description", description)
        if disk_encryption_key_name is not None:
            pulumi.set(__self__, "disk_encryption_key_name", disk_encryption_key_name)
        if display_name is not None:
            pulumi.set(__self__, "display_name", display_name)
        if ip_range is not None:
            pulumi.set(__self__, "ip_range", ip_range)
        if location is not None:
            pulumi.set(__self__, "location", location)
        if name is not None:
            pulumi.set(__self__, "name", name)
        if peering_cidr_range is not None:
            pulumi.set(__self__, "peering_cidr_range", peering_cidr_range)

    @property
    @pulumi.getter(name="organizationId")
    def organization_id(self) -> pulumi.Input[str]:
        return pulumi.get(self, "organization_id")

    @organization_id.setter
    def organization_id(self, value: pulumi.Input[str]):
        pulumi.set(self, "organization_id", value)

    @property
    @pulumi.getter(name="consumerAcceptList")
    def consumer_accept_list(self) -> Optional[pulumi.Input[Sequence[pulumi.Input[str]]]]:
        """
        Optional. Customer accept list represents the list of projects (id/number) on customer side that can privately connect to the service attachment. It is an optional field which the customers can provide during the instance creation. By default, the customer project associated with the Apigee organization will be included to the list.
        """
        return pulumi.get(self, "consumer_accept_list")

    @consumer_accept_list.setter
    def consumer_accept_list(self, value: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]]):
        pulumi.set(self, "consumer_accept_list", value)

    @property
    @pulumi.getter
    def description(self) -> Optional[pulumi.Input[str]]:
        """
        Optional. Description of the instance.
        """
        return pulumi.get(self, "description")

    @description.setter
    def description(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "description", value)

    @property
    @pulumi.getter(name="diskEncryptionKeyName")
    def disk_encryption_key_name(self) -> Optional[pulumi.Input[str]]:
        """
        Customer Managed Encryption Key (CMEK) used for disk and volume encryption. Required for Apigee paid subscriptions only. Use the following format: `projects/([^/]+)/locations/([^/]+)/keyRings/([^/]+)/cryptoKeys/([^/]+)`
        """
        return pulumi.get(self, "disk_encryption_key_name")

    @disk_encryption_key_name.setter
    def disk_encryption_key_name(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "disk_encryption_key_name", value)

    @property
    @pulumi.getter(name="displayName")
    def display_name(self) -> Optional[pulumi.Input[str]]:
        """
        Optional. Display name for the instance.
        """
        return pulumi.get(self, "display_name")

    @display_name.setter
    def display_name(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "display_name", value)

    @property
    @pulumi.getter(name="ipRange")
    def ip_range(self) -> Optional[pulumi.Input[str]]:
        """
        Optional. IP range represents the customer-provided CIDR block of length 22 that will be used for the Apigee instance creation. This optional range, if provided, should be freely available as part of larger named range the customer has allocated to the Service Networking peering. If this is not provided, Apigee will automatically request for any available /22 CIDR block from Service Networking. The customer should use this CIDR block for configuring their firewall needs to allow traffic from Apigee. Input format: "a.b.c.d/22", Output format: a.b.c.d/22, e.f.g.h/28"
        """
        return pulumi.get(self, "ip_range")

    @ip_range.setter
    def ip_range(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "ip_range", value)

    @property
    @pulumi.getter
    def location(self) -> Optional[pulumi.Input[str]]:
        """
        Compute Engine location where the instance resides.
        """
        return pulumi.get(self, "location")

    @location.setter
    def location(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "location", value)

    @property
    @pulumi.getter
    def name(self) -> Optional[pulumi.Input[str]]:
        """
        Resource ID of the instance. Values must match the regular expression `^a-z{0,30}[a-z\d]$`.
        """
        return pulumi.get(self, "name")

    @name.setter
    def name(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "name", value)

    @property
    @pulumi.getter(name="peeringCidrRange")
    def peering_cidr_range(self) -> Optional[pulumi.Input['InstancePeeringCidrRange']]:
        """
        Optional. Size of the CIDR block range that will be reserved by the instance. PAID organizations support `SLASH_16` to `SLASH_20` and defaults to `SLASH_16`. Evaluation organizations support only `SLASH_23`.
        """
        return pulumi.get(self, "peering_cidr_range")

    @peering_cidr_range.setter
    def peering_cidr_range(self, value: Optional[pulumi.Input['InstancePeeringCidrRange']]):
        pulumi.set(self, "peering_cidr_range", value)


class Instance(pulumi.CustomResource):
    @overload
    def __init__(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 consumer_accept_list: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]] = None,
                 description: Optional[pulumi.Input[str]] = None,
                 disk_encryption_key_name: Optional[pulumi.Input[str]] = None,
                 display_name: Optional[pulumi.Input[str]] = None,
                 ip_range: Optional[pulumi.Input[str]] = None,
                 location: Optional[pulumi.Input[str]] = None,
                 name: Optional[pulumi.Input[str]] = None,
                 organization_id: Optional[pulumi.Input[str]] = None,
                 peering_cidr_range: Optional[pulumi.Input['InstancePeeringCidrRange']] = None,
                 __props__=None):
        """
        Creates an Apigee runtime instance. The instance is accessible from the authorized network configured on the organization. **Note:** Not supported for Apigee hybrid.

        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[Sequence[pulumi.Input[str]]] consumer_accept_list: Optional. Customer accept list represents the list of projects (id/number) on customer side that can privately connect to the service attachment. It is an optional field which the customers can provide during the instance creation. By default, the customer project associated with the Apigee organization will be included to the list.
        :param pulumi.Input[str] description: Optional. Description of the instance.
        :param pulumi.Input[str] disk_encryption_key_name: Customer Managed Encryption Key (CMEK) used for disk and volume encryption. Required for Apigee paid subscriptions only. Use the following format: `projects/([^/]+)/locations/([^/]+)/keyRings/([^/]+)/cryptoKeys/([^/]+)`
        :param pulumi.Input[str] display_name: Optional. Display name for the instance.
        :param pulumi.Input[str] ip_range: Optional. IP range represents the customer-provided CIDR block of length 22 that will be used for the Apigee instance creation. This optional range, if provided, should be freely available as part of larger named range the customer has allocated to the Service Networking peering. If this is not provided, Apigee will automatically request for any available /22 CIDR block from Service Networking. The customer should use this CIDR block for configuring their firewall needs to allow traffic from Apigee. Input format: "a.b.c.d/22", Output format: a.b.c.d/22, e.f.g.h/28"
        :param pulumi.Input[str] location: Compute Engine location where the instance resides.
        :param pulumi.Input[str] name: Resource ID of the instance. Values must match the regular expression `^a-z{0,30}[a-z\d]$`.
        :param pulumi.Input['InstancePeeringCidrRange'] peering_cidr_range: Optional. Size of the CIDR block range that will be reserved by the instance. PAID organizations support `SLASH_16` to `SLASH_20` and defaults to `SLASH_16`. Evaluation organizations support only `SLASH_23`.
        """
        ...
    @overload
    def __init__(__self__,
                 resource_name: str,
                 args: InstanceArgs,
                 opts: Optional[pulumi.ResourceOptions] = None):
        """
        Creates an Apigee runtime instance. The instance is accessible from the authorized network configured on the organization. **Note:** Not supported for Apigee hybrid.

        :param str resource_name: The name of the resource.
        :param InstanceArgs args: The arguments to use to populate this resource's properties.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        ...
    def __init__(__self__, resource_name: str, *args, **kwargs):
        resource_args, opts = _utilities.get_resource_args_opts(InstanceArgs, pulumi.ResourceOptions, *args, **kwargs)
        if resource_args is not None:
            __self__._internal_init(resource_name, opts, **resource_args.__dict__)
        else:
            __self__._internal_init(resource_name, *args, **kwargs)

    def _internal_init(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 consumer_accept_list: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]] = None,
                 description: Optional[pulumi.Input[str]] = None,
                 disk_encryption_key_name: Optional[pulumi.Input[str]] = None,
                 display_name: Optional[pulumi.Input[str]] = None,
                 ip_range: Optional[pulumi.Input[str]] = None,
                 location: Optional[pulumi.Input[str]] = None,
                 name: Optional[pulumi.Input[str]] = None,
                 organization_id: Optional[pulumi.Input[str]] = None,
                 peering_cidr_range: Optional[pulumi.Input['InstancePeeringCidrRange']] = None,
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
            __props__ = InstanceArgs.__new__(InstanceArgs)

            __props__.__dict__["consumer_accept_list"] = consumer_accept_list
            __props__.__dict__["description"] = description
            __props__.__dict__["disk_encryption_key_name"] = disk_encryption_key_name
            __props__.__dict__["display_name"] = display_name
            __props__.__dict__["ip_range"] = ip_range
            __props__.__dict__["location"] = location
            __props__.__dict__["name"] = name
            if organization_id is None and not opts.urn:
                raise TypeError("Missing required property 'organization_id'")
            __props__.__dict__["organization_id"] = organization_id
            __props__.__dict__["peering_cidr_range"] = peering_cidr_range
            __props__.__dict__["created_at"] = None
            __props__.__dict__["host"] = None
            __props__.__dict__["last_modified_at"] = None
            __props__.__dict__["port"] = None
            __props__.__dict__["runtime_version"] = None
            __props__.__dict__["service_attachment"] = None
            __props__.__dict__["state"] = None
        super(Instance, __self__).__init__(
            'google-native:apigee/v1:Instance',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name: str,
            id: pulumi.Input[str],
            opts: Optional[pulumi.ResourceOptions] = None) -> 'Instance':
        """
        Get an existing Instance resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param pulumi.Input[str] id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = InstanceArgs.__new__(InstanceArgs)

        __props__.__dict__["consumer_accept_list"] = None
        __props__.__dict__["created_at"] = None
        __props__.__dict__["description"] = None
        __props__.__dict__["disk_encryption_key_name"] = None
        __props__.__dict__["display_name"] = None
        __props__.__dict__["host"] = None
        __props__.__dict__["ip_range"] = None
        __props__.__dict__["last_modified_at"] = None
        __props__.__dict__["location"] = None
        __props__.__dict__["name"] = None
        __props__.__dict__["peering_cidr_range"] = None
        __props__.__dict__["port"] = None
        __props__.__dict__["runtime_version"] = None
        __props__.__dict__["service_attachment"] = None
        __props__.__dict__["state"] = None
        return Instance(resource_name, opts=opts, __props__=__props__)

    @property
    @pulumi.getter(name="consumerAcceptList")
    def consumer_accept_list(self) -> pulumi.Output[Sequence[str]]:
        """
        Optional. Customer accept list represents the list of projects (id/number) on customer side that can privately connect to the service attachment. It is an optional field which the customers can provide during the instance creation. By default, the customer project associated with the Apigee organization will be included to the list.
        """
        return pulumi.get(self, "consumer_accept_list")

    @property
    @pulumi.getter(name="createdAt")
    def created_at(self) -> pulumi.Output[str]:
        """
        Time the instance was created in milliseconds since epoch.
        """
        return pulumi.get(self, "created_at")

    @property
    @pulumi.getter
    def description(self) -> pulumi.Output[str]:
        """
        Optional. Description of the instance.
        """
        return pulumi.get(self, "description")

    @property
    @pulumi.getter(name="diskEncryptionKeyName")
    def disk_encryption_key_name(self) -> pulumi.Output[str]:
        """
        Customer Managed Encryption Key (CMEK) used for disk and volume encryption. Required for Apigee paid subscriptions only. Use the following format: `projects/([^/]+)/locations/([^/]+)/keyRings/([^/]+)/cryptoKeys/([^/]+)`
        """
        return pulumi.get(self, "disk_encryption_key_name")

    @property
    @pulumi.getter(name="displayName")
    def display_name(self) -> pulumi.Output[str]:
        """
        Optional. Display name for the instance.
        """
        return pulumi.get(self, "display_name")

    @property
    @pulumi.getter
    def host(self) -> pulumi.Output[str]:
        """
        Internal hostname or IP address of the Apigee endpoint used by clients to connect to the service.
        """
        return pulumi.get(self, "host")

    @property
    @pulumi.getter(name="ipRange")
    def ip_range(self) -> pulumi.Output[str]:
        """
        Optional. IP range represents the customer-provided CIDR block of length 22 that will be used for the Apigee instance creation. This optional range, if provided, should be freely available as part of larger named range the customer has allocated to the Service Networking peering. If this is not provided, Apigee will automatically request for any available /22 CIDR block from Service Networking. The customer should use this CIDR block for configuring their firewall needs to allow traffic from Apigee. Input format: "a.b.c.d/22", Output format: a.b.c.d/22, e.f.g.h/28"
        """
        return pulumi.get(self, "ip_range")

    @property
    @pulumi.getter(name="lastModifiedAt")
    def last_modified_at(self) -> pulumi.Output[str]:
        """
        Time the instance was last modified in milliseconds since epoch.
        """
        return pulumi.get(self, "last_modified_at")

    @property
    @pulumi.getter
    def location(self) -> pulumi.Output[str]:
        """
        Compute Engine location where the instance resides.
        """
        return pulumi.get(self, "location")

    @property
    @pulumi.getter
    def name(self) -> pulumi.Output[str]:
        """
        Resource ID of the instance. Values must match the regular expression `^a-z{0,30}[a-z\d]$`.
        """
        return pulumi.get(self, "name")

    @property
    @pulumi.getter(name="peeringCidrRange")
    def peering_cidr_range(self) -> pulumi.Output[str]:
        """
        Optional. Size of the CIDR block range that will be reserved by the instance. PAID organizations support `SLASH_16` to `SLASH_20` and defaults to `SLASH_16`. Evaluation organizations support only `SLASH_23`.
        """
        return pulumi.get(self, "peering_cidr_range")

    @property
    @pulumi.getter
    def port(self) -> pulumi.Output[str]:
        """
        Port number of the exposed Apigee endpoint.
        """
        return pulumi.get(self, "port")

    @property
    @pulumi.getter(name="runtimeVersion")
    def runtime_version(self) -> pulumi.Output[str]:
        """
        Version of the runtime system running in the instance. The runtime system is the set of components that serve the API Proxy traffic in your Environments.
        """
        return pulumi.get(self, "runtime_version")

    @property
    @pulumi.getter(name="serviceAttachment")
    def service_attachment(self) -> pulumi.Output[str]:
        """
        Resource name of the service attachment created for the instance in the format: `projects/*/regions/*/serviceAttachments/*` Apigee customers can privately forward traffic to this service attachment using the PSC endpoints.
        """
        return pulumi.get(self, "service_attachment")

    @property
    @pulumi.getter
    def state(self) -> pulumi.Output[str]:
        """
        State of the instance. Values other than `ACTIVE` means the resource is not ready to use.
        """
        return pulumi.get(self, "state")

