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

__all__ = ['NamespaceServiceArgs', 'NamespaceService']

@pulumi.input_type
class NamespaceServiceArgs:
    def __init__(__self__, *,
                 namespaces_id: pulumi.Input[str],
                 services_id: pulumi.Input[str],
                 api_version: Optional[pulumi.Input[str]] = None,
                 kind: Optional[pulumi.Input[str]] = None,
                 metadata: Optional[pulumi.Input['ObjectMetaArgs']] = None,
                 spec: Optional[pulumi.Input['ServiceSpecArgs']] = None,
                 status: Optional[pulumi.Input['ServiceStatusArgs']] = None):
        """
        The set of arguments for constructing a NamespaceService resource.
        :param pulumi.Input[str] api_version: The API version for this call such as "serving.knative.dev/v1alpha1".
        :param pulumi.Input[str] kind: The kind of resource, in this case "Service".
        :param pulumi.Input['ObjectMetaArgs'] metadata: Metadata associated with this Service, including name, namespace, labels, and annotations.
        :param pulumi.Input['ServiceSpecArgs'] spec: Spec holds the desired state of the Service (from the client).
        :param pulumi.Input['ServiceStatusArgs'] status: Status communicates the observed state of the Service (from the controller).
        """
        pulumi.set(__self__, "namespaces_id", namespaces_id)
        pulumi.set(__self__, "services_id", services_id)
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
    @pulumi.getter(name="namespacesId")
    def namespaces_id(self) -> pulumi.Input[str]:
        return pulumi.get(self, "namespaces_id")

    @namespaces_id.setter
    def namespaces_id(self, value: pulumi.Input[str]):
        pulumi.set(self, "namespaces_id", value)

    @property
    @pulumi.getter(name="servicesId")
    def services_id(self) -> pulumi.Input[str]:
        return pulumi.get(self, "services_id")

    @services_id.setter
    def services_id(self, value: pulumi.Input[str]):
        pulumi.set(self, "services_id", value)

    @property
    @pulumi.getter(name="apiVersion")
    def api_version(self) -> Optional[pulumi.Input[str]]:
        """
        The API version for this call such as "serving.knative.dev/v1alpha1".
        """
        return pulumi.get(self, "api_version")

    @api_version.setter
    def api_version(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "api_version", value)

    @property
    @pulumi.getter
    def kind(self) -> Optional[pulumi.Input[str]]:
        """
        The kind of resource, in this case "Service".
        """
        return pulumi.get(self, "kind")

    @kind.setter
    def kind(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "kind", value)

    @property
    @pulumi.getter
    def metadata(self) -> Optional[pulumi.Input['ObjectMetaArgs']]:
        """
        Metadata associated with this Service, including name, namespace, labels, and annotations.
        """
        return pulumi.get(self, "metadata")

    @metadata.setter
    def metadata(self, value: Optional[pulumi.Input['ObjectMetaArgs']]):
        pulumi.set(self, "metadata", value)

    @property
    @pulumi.getter
    def spec(self) -> Optional[pulumi.Input['ServiceSpecArgs']]:
        """
        Spec holds the desired state of the Service (from the client).
        """
        return pulumi.get(self, "spec")

    @spec.setter
    def spec(self, value: Optional[pulumi.Input['ServiceSpecArgs']]):
        pulumi.set(self, "spec", value)

    @property
    @pulumi.getter
    def status(self) -> Optional[pulumi.Input['ServiceStatusArgs']]:
        """
        Status communicates the observed state of the Service (from the controller).
        """
        return pulumi.get(self, "status")

    @status.setter
    def status(self, value: Optional[pulumi.Input['ServiceStatusArgs']]):
        pulumi.set(self, "status", value)


class NamespaceService(pulumi.CustomResource):
    @overload
    def __init__(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 api_version: Optional[pulumi.Input[str]] = None,
                 kind: Optional[pulumi.Input[str]] = None,
                 metadata: Optional[pulumi.Input[pulumi.InputType['ObjectMetaArgs']]] = None,
                 namespaces_id: Optional[pulumi.Input[str]] = None,
                 services_id: Optional[pulumi.Input[str]] = None,
                 spec: Optional[pulumi.Input[pulumi.InputType['ServiceSpecArgs']]] = None,
                 status: Optional[pulumi.Input[pulumi.InputType['ServiceStatusArgs']]] = None,
                 __props__=None):
        """
        Rpc to create a service.

        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] api_version: The API version for this call such as "serving.knative.dev/v1alpha1".
        :param pulumi.Input[str] kind: The kind of resource, in this case "Service".
        :param pulumi.Input[pulumi.InputType['ObjectMetaArgs']] metadata: Metadata associated with this Service, including name, namespace, labels, and annotations.
        :param pulumi.Input[pulumi.InputType['ServiceSpecArgs']] spec: Spec holds the desired state of the Service (from the client).
        :param pulumi.Input[pulumi.InputType['ServiceStatusArgs']] status: Status communicates the observed state of the Service (from the controller).
        """
        ...
    @overload
    def __init__(__self__,
                 resource_name: str,
                 args: NamespaceServiceArgs,
                 opts: Optional[pulumi.ResourceOptions] = None):
        """
        Rpc to create a service.

        :param str resource_name: The name of the resource.
        :param NamespaceServiceArgs args: The arguments to use to populate this resource's properties.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        ...
    def __init__(__self__, resource_name: str, *args, **kwargs):
        resource_args, opts = _utilities.get_resource_args_opts(NamespaceServiceArgs, pulumi.ResourceOptions, *args, **kwargs)
        if resource_args is not None:
            __self__._internal_init(resource_name, opts, **resource_args.__dict__)
        else:
            __self__._internal_init(resource_name, *args, **kwargs)

    def _internal_init(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 api_version: Optional[pulumi.Input[str]] = None,
                 kind: Optional[pulumi.Input[str]] = None,
                 metadata: Optional[pulumi.Input[pulumi.InputType['ObjectMetaArgs']]] = None,
                 namespaces_id: Optional[pulumi.Input[str]] = None,
                 services_id: Optional[pulumi.Input[str]] = None,
                 spec: Optional[pulumi.Input[pulumi.InputType['ServiceSpecArgs']]] = None,
                 status: Optional[pulumi.Input[pulumi.InputType['ServiceStatusArgs']]] = None,
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
            __props__ = NamespaceServiceArgs.__new__(NamespaceServiceArgs)

            __props__.__dict__["api_version"] = api_version
            __props__.__dict__["kind"] = kind
            __props__.__dict__["metadata"] = metadata
            if namespaces_id is None and not opts.urn:
                raise TypeError("Missing required property 'namespaces_id'")
            __props__.__dict__["namespaces_id"] = namespaces_id
            if services_id is None and not opts.urn:
                raise TypeError("Missing required property 'services_id'")
            __props__.__dict__["services_id"] = services_id
            __props__.__dict__["spec"] = spec
            __props__.__dict__["status"] = status
        super(NamespaceService, __self__).__init__(
            'gcp-native:run/v1alpha1:NamespaceService',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name: str,
            id: pulumi.Input[str],
            opts: Optional[pulumi.ResourceOptions] = None) -> 'NamespaceService':
        """
        Get an existing NamespaceService resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param pulumi.Input[str] id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = NamespaceServiceArgs.__new__(NamespaceServiceArgs)

        __props__.__dict__["api_version"] = None
        __props__.__dict__["kind"] = None
        __props__.__dict__["metadata"] = None
        __props__.__dict__["spec"] = None
        __props__.__dict__["status"] = None
        return NamespaceService(resource_name, opts=opts, __props__=__props__)

    @property
    @pulumi.getter(name="apiVersion")
    def api_version(self) -> pulumi.Output[str]:
        """
        The API version for this call such as "serving.knative.dev/v1alpha1".
        """
        return pulumi.get(self, "api_version")

    @property
    @pulumi.getter
    def kind(self) -> pulumi.Output[str]:
        """
        The kind of resource, in this case "Service".
        """
        return pulumi.get(self, "kind")

    @property
    @pulumi.getter
    def metadata(self) -> pulumi.Output['outputs.ObjectMetaResponse']:
        """
        Metadata associated with this Service, including name, namespace, labels, and annotations.
        """
        return pulumi.get(self, "metadata")

    @property
    @pulumi.getter
    def spec(self) -> pulumi.Output['outputs.ServiceSpecResponse']:
        """
        Spec holds the desired state of the Service (from the client).
        """
        return pulumi.get(self, "spec")

    @property
    @pulumi.getter
    def status(self) -> pulumi.Output['outputs.ServiceStatusResponse']:
        """
        Status communicates the observed state of the Service (from the controller).
        """
        return pulumi.get(self, "status")

