# coding=utf-8
# *** WARNING: this file was generated by the Pulumi SDK Generator. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import warnings
import pulumi
import pulumi.runtime
from typing import Any, Mapping, Optional, Sequence, Union, overload
from ... import _utilities

__all__ = ['TenantArgs', 'Tenant']

@pulumi.input_type
class TenantArgs:
    def __init__(__self__, *,
                 external_id: pulumi.Input[str],
                 project: pulumi.Input[str],
                 name: Optional[pulumi.Input[str]] = None):
        """
        The set of arguments for constructing a Tenant resource.
        :param pulumi.Input[str] external_id: Client side tenant identifier, used to uniquely identify the tenant. The maximum number of allowed characters is 255.
        :param pulumi.Input[str] name: Required during tenant update. The resource name for a tenant. This is generated by the service when a tenant is created. The format is "projects/{project_id}/tenants/{tenant_id}", for example, "projects/foo/tenants/bar".
        """
        pulumi.set(__self__, "external_id", external_id)
        pulumi.set(__self__, "project", project)
        if name is not None:
            pulumi.set(__self__, "name", name)

    @property
    @pulumi.getter(name="externalId")
    def external_id(self) -> pulumi.Input[str]:
        """
        Client side tenant identifier, used to uniquely identify the tenant. The maximum number of allowed characters is 255.
        """
        return pulumi.get(self, "external_id")

    @external_id.setter
    def external_id(self, value: pulumi.Input[str]):
        pulumi.set(self, "external_id", value)

    @property
    @pulumi.getter
    def project(self) -> pulumi.Input[str]:
        return pulumi.get(self, "project")

    @project.setter
    def project(self, value: pulumi.Input[str]):
        pulumi.set(self, "project", value)

    @property
    @pulumi.getter
    def name(self) -> Optional[pulumi.Input[str]]:
        """
        Required during tenant update. The resource name for a tenant. This is generated by the service when a tenant is created. The format is "projects/{project_id}/tenants/{tenant_id}", for example, "projects/foo/tenants/bar".
        """
        return pulumi.get(self, "name")

    @name.setter
    def name(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "name", value)


class Tenant(pulumi.CustomResource):
    @overload
    def __init__(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 external_id: Optional[pulumi.Input[str]] = None,
                 name: Optional[pulumi.Input[str]] = None,
                 project: Optional[pulumi.Input[str]] = None,
                 __props__=None):
        """
        Creates a new tenant entity.

        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] external_id: Client side tenant identifier, used to uniquely identify the tenant. The maximum number of allowed characters is 255.
        :param pulumi.Input[str] name: Required during tenant update. The resource name for a tenant. This is generated by the service when a tenant is created. The format is "projects/{project_id}/tenants/{tenant_id}", for example, "projects/foo/tenants/bar".
        """
        ...
    @overload
    def __init__(__self__,
                 resource_name: str,
                 args: TenantArgs,
                 opts: Optional[pulumi.ResourceOptions] = None):
        """
        Creates a new tenant entity.

        :param str resource_name: The name of the resource.
        :param TenantArgs args: The arguments to use to populate this resource's properties.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        ...
    def __init__(__self__, resource_name: str, *args, **kwargs):
        resource_args, opts = _utilities.get_resource_args_opts(TenantArgs, pulumi.ResourceOptions, *args, **kwargs)
        if resource_args is not None:
            __self__._internal_init(resource_name, opts, **resource_args.__dict__)
        else:
            __self__._internal_init(resource_name, *args, **kwargs)

    def _internal_init(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 external_id: Optional[pulumi.Input[str]] = None,
                 name: Optional[pulumi.Input[str]] = None,
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
            __props__ = TenantArgs.__new__(TenantArgs)

            if external_id is None and not opts.urn:
                raise TypeError("Missing required property 'external_id'")
            __props__.__dict__["external_id"] = external_id
            __props__.__dict__["name"] = name
            if project is None and not opts.urn:
                raise TypeError("Missing required property 'project'")
            __props__.__dict__["project"] = project
        super(Tenant, __self__).__init__(
            'google-native:jobs/v4:Tenant',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name: str,
            id: pulumi.Input[str],
            opts: Optional[pulumi.ResourceOptions] = None) -> 'Tenant':
        """
        Get an existing Tenant resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param pulumi.Input[str] id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = TenantArgs.__new__(TenantArgs)

        __props__.__dict__["external_id"] = None
        __props__.__dict__["name"] = None
        return Tenant(resource_name, opts=opts, __props__=__props__)

    @property
    @pulumi.getter(name="externalId")
    def external_id(self) -> pulumi.Output[str]:
        """
        Client side tenant identifier, used to uniquely identify the tenant. The maximum number of allowed characters is 255.
        """
        return pulumi.get(self, "external_id")

    @property
    @pulumi.getter
    def name(self) -> pulumi.Output[str]:
        """
        Required during tenant update. The resource name for a tenant. This is generated by the service when a tenant is created. The format is "projects/{project_id}/tenants/{tenant_id}", for example, "projects/foo/tenants/bar".
        """
        return pulumi.get(self, "name")

