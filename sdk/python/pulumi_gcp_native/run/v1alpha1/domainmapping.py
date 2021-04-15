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

__all__ = ['DomainmappingArgs', 'Domainmapping']

@pulumi.input_type
class DomainmappingArgs:
    def __init__(__self__, *,
                 domainmappings_id: pulumi.Input[str],
                 locations_id: pulumi.Input[str],
                 projects_id: pulumi.Input[str],
                 api_version: Optional[pulumi.Input[str]] = None,
                 kind: Optional[pulumi.Input[str]] = None,
                 metadata: Optional[pulumi.Input['ObjectMetaArgs']] = None,
                 spec: Optional[pulumi.Input['DomainMappingSpecArgs']] = None,
                 status: Optional[pulumi.Input['DomainMappingStatusArgs']] = None):
        """
        The set of arguments for constructing a Domainmapping resource.
        :param pulumi.Input[str] api_version: The API version for this call such as "domains.cloudrun.com/v1alpha1".
        :param pulumi.Input[str] kind: The kind of resource, in this case "DomainMapping".
        :param pulumi.Input['ObjectMetaArgs'] metadata: Metadata associated with this BuildTemplate.
        :param pulumi.Input['DomainMappingSpecArgs'] spec: The spec for this DomainMapping.
        :param pulumi.Input['DomainMappingStatusArgs'] status: The current status of the DomainMapping.
        """
        pulumi.set(__self__, "domainmappings_id", domainmappings_id)
        pulumi.set(__self__, "locations_id", locations_id)
        pulumi.set(__self__, "projects_id", projects_id)
        if api_version is not None:
            pulumi.set(__self__, "api_version", api_version)
        if kind is not None:
            pulumi.set(__self__, "kind", kind)
        if metadata is not None:
            pulumi.set(__self__, "metadata", metadata)
        if spec is not None:
            pulumi.set(__self__, "spec", spec)
        if status is not None:
            pulumi.set(__self__, "status", status)

    @property
    @pulumi.getter(name="domainmappingsId")
    def domainmappings_id(self) -> pulumi.Input[str]:
        return pulumi.get(self, "domainmappings_id")

    @domainmappings_id.setter
    def domainmappings_id(self, value: pulumi.Input[str]):
        pulumi.set(self, "domainmappings_id", value)

    @property
    @pulumi.getter(name="locationsId")
    def locations_id(self) -> pulumi.Input[str]:
        return pulumi.get(self, "locations_id")

    @locations_id.setter
    def locations_id(self, value: pulumi.Input[str]):
        pulumi.set(self, "locations_id", value)

    @property
    @pulumi.getter(name="projectsId")
    def projects_id(self) -> pulumi.Input[str]:
        return pulumi.get(self, "projects_id")

    @projects_id.setter
    def projects_id(self, value: pulumi.Input[str]):
        pulumi.set(self, "projects_id", value)

    @property
    @pulumi.getter(name="apiVersion")
    def api_version(self) -> Optional[pulumi.Input[str]]:
        """
        The API version for this call such as "domains.cloudrun.com/v1alpha1".
        """
        return pulumi.get(self, "api_version")

    @api_version.setter
    def api_version(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "api_version", value)

    @property
    @pulumi.getter
    def kind(self) -> Optional[pulumi.Input[str]]:
        """
        The kind of resource, in this case "DomainMapping".
        """
        return pulumi.get(self, "kind")

    @kind.setter
    def kind(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "kind", value)

    @property
    @pulumi.getter
    def metadata(self) -> Optional[pulumi.Input['ObjectMetaArgs']]:
        """
        Metadata associated with this BuildTemplate.
        """
        return pulumi.get(self, "metadata")

    @metadata.setter
    def metadata(self, value: Optional[pulumi.Input['ObjectMetaArgs']]):
        pulumi.set(self, "metadata", value)

    @property
    @pulumi.getter
    def spec(self) -> Optional[pulumi.Input['DomainMappingSpecArgs']]:
        """
        The spec for this DomainMapping.
        """
        return pulumi.get(self, "spec")

    @spec.setter
    def spec(self, value: Optional[pulumi.Input['DomainMappingSpecArgs']]):
        pulumi.set(self, "spec", value)

    @property
    @pulumi.getter
    def status(self) -> Optional[pulumi.Input['DomainMappingStatusArgs']]:
        """
        The current status of the DomainMapping.
        """
        return pulumi.get(self, "status")

    @status.setter
    def status(self, value: Optional[pulumi.Input['DomainMappingStatusArgs']]):
        pulumi.set(self, "status", value)


class Domainmapping(pulumi.CustomResource):
    @overload
    def __init__(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 api_version: Optional[pulumi.Input[str]] = None,
                 domainmappings_id: Optional[pulumi.Input[str]] = None,
                 kind: Optional[pulumi.Input[str]] = None,
                 locations_id: Optional[pulumi.Input[str]] = None,
                 metadata: Optional[pulumi.Input[pulumi.InputType['ObjectMetaArgs']]] = None,
                 projects_id: Optional[pulumi.Input[str]] = None,
                 spec: Optional[pulumi.Input[pulumi.InputType['DomainMappingSpecArgs']]] = None,
                 status: Optional[pulumi.Input[pulumi.InputType['DomainMappingStatusArgs']]] = None,
                 __props__=None):
        """
        Creates a new domain mapping.

        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] api_version: The API version for this call such as "domains.cloudrun.com/v1alpha1".
        :param pulumi.Input[str] kind: The kind of resource, in this case "DomainMapping".
        :param pulumi.Input[pulumi.InputType['ObjectMetaArgs']] metadata: Metadata associated with this BuildTemplate.
        :param pulumi.Input[pulumi.InputType['DomainMappingSpecArgs']] spec: The spec for this DomainMapping.
        :param pulumi.Input[pulumi.InputType['DomainMappingStatusArgs']] status: The current status of the DomainMapping.
        """
        ...
    @overload
    def __init__(__self__,
                 resource_name: str,
                 args: DomainmappingArgs,
                 opts: Optional[pulumi.ResourceOptions] = None):
        """
        Creates a new domain mapping.

        :param str resource_name: The name of the resource.
        :param DomainmappingArgs args: The arguments to use to populate this resource's properties.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        ...
    def __init__(__self__, resource_name: str, *args, **kwargs):
        resource_args, opts = _utilities.get_resource_args_opts(DomainmappingArgs, pulumi.ResourceOptions, *args, **kwargs)
        if resource_args is not None:
            __self__._internal_init(resource_name, opts, **resource_args.__dict__)
        else:
            __self__._internal_init(resource_name, *args, **kwargs)

    def _internal_init(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 api_version: Optional[pulumi.Input[str]] = None,
                 domainmappings_id: Optional[pulumi.Input[str]] = None,
                 kind: Optional[pulumi.Input[str]] = None,
                 locations_id: Optional[pulumi.Input[str]] = None,
                 metadata: Optional[pulumi.Input[pulumi.InputType['ObjectMetaArgs']]] = None,
                 projects_id: Optional[pulumi.Input[str]] = None,
                 spec: Optional[pulumi.Input[pulumi.InputType['DomainMappingSpecArgs']]] = None,
                 status: Optional[pulumi.Input[pulumi.InputType['DomainMappingStatusArgs']]] = None,
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
            __props__ = DomainmappingArgs.__new__(DomainmappingArgs)

            __props__.__dict__["api_version"] = api_version
            if domainmappings_id is None and not opts.urn:
                raise TypeError("Missing required property 'domainmappings_id'")
            __props__.__dict__["domainmappings_id"] = domainmappings_id
            __props__.__dict__["kind"] = kind
            if locations_id is None and not opts.urn:
                raise TypeError("Missing required property 'locations_id'")
            __props__.__dict__["locations_id"] = locations_id
            __props__.__dict__["metadata"] = metadata
            if projects_id is None and not opts.urn:
                raise TypeError("Missing required property 'projects_id'")
            __props__.__dict__["projects_id"] = projects_id
            __props__.__dict__["spec"] = spec
            __props__.__dict__["status"] = status
        super(Domainmapping, __self__).__init__(
            'gcp-native:run/v1alpha1:Domainmapping',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name: str,
            id: pulumi.Input[str],
            opts: Optional[pulumi.ResourceOptions] = None) -> 'Domainmapping':
        """
        Get an existing Domainmapping resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param pulumi.Input[str] id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = DomainmappingArgs.__new__(DomainmappingArgs)

        __props__.__dict__["api_version"] = None
        __props__.__dict__["kind"] = None
        __props__.__dict__["metadata"] = None
        __props__.__dict__["spec"] = None
        __props__.__dict__["status"] = None
        return Domainmapping(resource_name, opts=opts, __props__=__props__)

    @property
    @pulumi.getter(name="apiVersion")
    def api_version(self) -> pulumi.Output[str]:
        """
        The API version for this call such as "domains.cloudrun.com/v1alpha1".
        """
        return pulumi.get(self, "api_version")

    @property
    @pulumi.getter
    def kind(self) -> pulumi.Output[str]:
        """
        The kind of resource, in this case "DomainMapping".
        """
        return pulumi.get(self, "kind")

    @property
    @pulumi.getter
    def metadata(self) -> pulumi.Output['outputs.ObjectMetaResponse']:
        """
        Metadata associated with this BuildTemplate.
        """
        return pulumi.get(self, "metadata")

    @property
    @pulumi.getter
    def spec(self) -> pulumi.Output['outputs.DomainMappingSpecResponse']:
        """
        The spec for this DomainMapping.
        """
        return pulumi.get(self, "spec")

    @property
    @pulumi.getter
    def status(self) -> pulumi.Output['outputs.DomainMappingStatusResponse']:
        """
        The current status of the DomainMapping.
        """
        return pulumi.get(self, "status")

