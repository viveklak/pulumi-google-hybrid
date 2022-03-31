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

__all__ = ['TlsRouteArgs', 'TlsRoute']

@pulumi.input_type
class TlsRouteArgs:
    def __init__(__self__, *,
                 rules: pulumi.Input[Sequence[pulumi.Input['TlsRouteRouteRuleArgs']]],
                 tls_route_id: pulumi.Input[str],
                 description: Optional[pulumi.Input[str]] = None,
                 gateways: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]] = None,
                 location: Optional[pulumi.Input[str]] = None,
                 meshes: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]] = None,
                 name: Optional[pulumi.Input[str]] = None,
                 project: Optional[pulumi.Input[str]] = None):
        """
        The set of arguments for constructing a TlsRoute resource.
        :param pulumi.Input[Sequence[pulumi.Input['TlsRouteRouteRuleArgs']]] rules: Rules that define how traffic is routed and handled. At least one RouteRule must be supplied. If there are multiple rules then the action taken will be the first rule to match.
        :param pulumi.Input[str] tls_route_id: Required. Short name of the TlsRoute resource to be created. E.g. TODO(Add an example).
        :param pulumi.Input[str] description: Optional. A free-text description of the resource. Max length 1024 characters.
        :param pulumi.Input[Sequence[pulumi.Input[str]]] gateways: Optional. Gateways defines a list of gateways this TlsRoute is attached to, as one of the routing rules to route the requests served by the gateway. Each gateway reference should match the pattern: `projects/*/locations/global/gateways/`
        :param pulumi.Input[Sequence[pulumi.Input[str]]] meshes: Optional. Meshes defines a list of meshes this TlsRoute is attached to, as one of the routing rules to route the requests served by the mesh. Each mesh reference should match the pattern: `projects/*/locations/global/meshes/` The attached Mesh should be of a type SIDECAR
        :param pulumi.Input[str] name: Name of the TlsRoute resource. It matches pattern `projects/*/locations/global/tlsRoutes/tls_route_name>`.
        """
        pulumi.set(__self__, "rules", rules)
        pulumi.set(__self__, "tls_route_id", tls_route_id)
        if description is not None:
            pulumi.set(__self__, "description", description)
        if gateways is not None:
            pulumi.set(__self__, "gateways", gateways)
        if location is not None:
            pulumi.set(__self__, "location", location)
        if meshes is not None:
            pulumi.set(__self__, "meshes", meshes)
        if name is not None:
            pulumi.set(__self__, "name", name)
        if project is not None:
            pulumi.set(__self__, "project", project)

    @property
    @pulumi.getter
    def rules(self) -> pulumi.Input[Sequence[pulumi.Input['TlsRouteRouteRuleArgs']]]:
        """
        Rules that define how traffic is routed and handled. At least one RouteRule must be supplied. If there are multiple rules then the action taken will be the first rule to match.
        """
        return pulumi.get(self, "rules")

    @rules.setter
    def rules(self, value: pulumi.Input[Sequence[pulumi.Input['TlsRouteRouteRuleArgs']]]):
        pulumi.set(self, "rules", value)

    @property
    @pulumi.getter(name="tlsRouteId")
    def tls_route_id(self) -> pulumi.Input[str]:
        """
        Required. Short name of the TlsRoute resource to be created. E.g. TODO(Add an example).
        """
        return pulumi.get(self, "tls_route_id")

    @tls_route_id.setter
    def tls_route_id(self, value: pulumi.Input[str]):
        pulumi.set(self, "tls_route_id", value)

    @property
    @pulumi.getter
    def description(self) -> Optional[pulumi.Input[str]]:
        """
        Optional. A free-text description of the resource. Max length 1024 characters.
        """
        return pulumi.get(self, "description")

    @description.setter
    def description(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "description", value)

    @property
    @pulumi.getter
    def gateways(self) -> Optional[pulumi.Input[Sequence[pulumi.Input[str]]]]:
        """
        Optional. Gateways defines a list of gateways this TlsRoute is attached to, as one of the routing rules to route the requests served by the gateway. Each gateway reference should match the pattern: `projects/*/locations/global/gateways/`
        """
        return pulumi.get(self, "gateways")

    @gateways.setter
    def gateways(self, value: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]]):
        pulumi.set(self, "gateways", value)

    @property
    @pulumi.getter
    def location(self) -> Optional[pulumi.Input[str]]:
        return pulumi.get(self, "location")

    @location.setter
    def location(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "location", value)

    @property
    @pulumi.getter
    def meshes(self) -> Optional[pulumi.Input[Sequence[pulumi.Input[str]]]]:
        """
        Optional. Meshes defines a list of meshes this TlsRoute is attached to, as one of the routing rules to route the requests served by the mesh. Each mesh reference should match the pattern: `projects/*/locations/global/meshes/` The attached Mesh should be of a type SIDECAR
        """
        return pulumi.get(self, "meshes")

    @meshes.setter
    def meshes(self, value: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]]):
        pulumi.set(self, "meshes", value)

    @property
    @pulumi.getter
    def name(self) -> Optional[pulumi.Input[str]]:
        """
        Name of the TlsRoute resource. It matches pattern `projects/*/locations/global/tlsRoutes/tls_route_name>`.
        """
        return pulumi.get(self, "name")

    @name.setter
    def name(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "name", value)

    @property
    @pulumi.getter
    def project(self) -> Optional[pulumi.Input[str]]:
        return pulumi.get(self, "project")

    @project.setter
    def project(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "project", value)


class TlsRoute(pulumi.CustomResource):
    @overload
    def __init__(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 description: Optional[pulumi.Input[str]] = None,
                 gateways: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]] = None,
                 location: Optional[pulumi.Input[str]] = None,
                 meshes: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]] = None,
                 name: Optional[pulumi.Input[str]] = None,
                 project: Optional[pulumi.Input[str]] = None,
                 rules: Optional[pulumi.Input[Sequence[pulumi.Input[pulumi.InputType['TlsRouteRouteRuleArgs']]]]] = None,
                 tls_route_id: Optional[pulumi.Input[str]] = None,
                 __props__=None):
        """
        Creates a new TlsRoute in a given project and location.

        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] description: Optional. A free-text description of the resource. Max length 1024 characters.
        :param pulumi.Input[Sequence[pulumi.Input[str]]] gateways: Optional. Gateways defines a list of gateways this TlsRoute is attached to, as one of the routing rules to route the requests served by the gateway. Each gateway reference should match the pattern: `projects/*/locations/global/gateways/`
        :param pulumi.Input[Sequence[pulumi.Input[str]]] meshes: Optional. Meshes defines a list of meshes this TlsRoute is attached to, as one of the routing rules to route the requests served by the mesh. Each mesh reference should match the pattern: `projects/*/locations/global/meshes/` The attached Mesh should be of a type SIDECAR
        :param pulumi.Input[str] name: Name of the TlsRoute resource. It matches pattern `projects/*/locations/global/tlsRoutes/tls_route_name>`.
        :param pulumi.Input[Sequence[pulumi.Input[pulumi.InputType['TlsRouteRouteRuleArgs']]]] rules: Rules that define how traffic is routed and handled. At least one RouteRule must be supplied. If there are multiple rules then the action taken will be the first rule to match.
        :param pulumi.Input[str] tls_route_id: Required. Short name of the TlsRoute resource to be created. E.g. TODO(Add an example).
        """
        ...
    @overload
    def __init__(__self__,
                 resource_name: str,
                 args: TlsRouteArgs,
                 opts: Optional[pulumi.ResourceOptions] = None):
        """
        Creates a new TlsRoute in a given project and location.

        :param str resource_name: The name of the resource.
        :param TlsRouteArgs args: The arguments to use to populate this resource's properties.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        ...
    def __init__(__self__, resource_name: str, *args, **kwargs):
        resource_args, opts = _utilities.get_resource_args_opts(TlsRouteArgs, pulumi.ResourceOptions, *args, **kwargs)
        if resource_args is not None:
            __self__._internal_init(resource_name, opts, **resource_args.__dict__)
        else:
            __self__._internal_init(resource_name, *args, **kwargs)

    def _internal_init(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 description: Optional[pulumi.Input[str]] = None,
                 gateways: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]] = None,
                 location: Optional[pulumi.Input[str]] = None,
                 meshes: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]] = None,
                 name: Optional[pulumi.Input[str]] = None,
                 project: Optional[pulumi.Input[str]] = None,
                 rules: Optional[pulumi.Input[Sequence[pulumi.Input[pulumi.InputType['TlsRouteRouteRuleArgs']]]]] = None,
                 tls_route_id: Optional[pulumi.Input[str]] = None,
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
            __props__ = TlsRouteArgs.__new__(TlsRouteArgs)

            __props__.__dict__["description"] = description
            __props__.__dict__["gateways"] = gateways
            __props__.__dict__["location"] = location
            __props__.__dict__["meshes"] = meshes
            __props__.__dict__["name"] = name
            __props__.__dict__["project"] = project
            if rules is None and not opts.urn:
                raise TypeError("Missing required property 'rules'")
            __props__.__dict__["rules"] = rules
            if tls_route_id is None and not opts.urn:
                raise TypeError("Missing required property 'tls_route_id'")
            __props__.__dict__["tls_route_id"] = tls_route_id
            __props__.__dict__["create_time"] = None
            __props__.__dict__["self_link"] = None
            __props__.__dict__["update_time"] = None
        super(TlsRoute, __self__).__init__(
            'google-native:networkservices/v1beta1:TlsRoute',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name: str,
            id: pulumi.Input[str],
            opts: Optional[pulumi.ResourceOptions] = None) -> 'TlsRoute':
        """
        Get an existing TlsRoute resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param pulumi.Input[str] id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = TlsRouteArgs.__new__(TlsRouteArgs)

        __props__.__dict__["create_time"] = None
        __props__.__dict__["description"] = None
        __props__.__dict__["gateways"] = None
        __props__.__dict__["meshes"] = None
        __props__.__dict__["name"] = None
        __props__.__dict__["rules"] = None
        __props__.__dict__["self_link"] = None
        __props__.__dict__["update_time"] = None
        return TlsRoute(resource_name, opts=opts, __props__=__props__)

    @property
    @pulumi.getter(name="createTime")
    def create_time(self) -> pulumi.Output[str]:
        """
        The timestamp when the resource was created.
        """
        return pulumi.get(self, "create_time")

    @property
    @pulumi.getter
    def description(self) -> pulumi.Output[str]:
        """
        Optional. A free-text description of the resource. Max length 1024 characters.
        """
        return pulumi.get(self, "description")

    @property
    @pulumi.getter
    def gateways(self) -> pulumi.Output[Sequence[str]]:
        """
        Optional. Gateways defines a list of gateways this TlsRoute is attached to, as one of the routing rules to route the requests served by the gateway. Each gateway reference should match the pattern: `projects/*/locations/global/gateways/`
        """
        return pulumi.get(self, "gateways")

    @property
    @pulumi.getter
    def meshes(self) -> pulumi.Output[Sequence[str]]:
        """
        Optional. Meshes defines a list of meshes this TlsRoute is attached to, as one of the routing rules to route the requests served by the mesh. Each mesh reference should match the pattern: `projects/*/locations/global/meshes/` The attached Mesh should be of a type SIDECAR
        """
        return pulumi.get(self, "meshes")

    @property
    @pulumi.getter
    def name(self) -> pulumi.Output[str]:
        """
        Name of the TlsRoute resource. It matches pattern `projects/*/locations/global/tlsRoutes/tls_route_name>`.
        """
        return pulumi.get(self, "name")

    @property
    @pulumi.getter
    def rules(self) -> pulumi.Output[Sequence['outputs.TlsRouteRouteRuleResponse']]:
        """
        Rules that define how traffic is routed and handled. At least one RouteRule must be supplied. If there are multiple rules then the action taken will be the first rule to match.
        """
        return pulumi.get(self, "rules")

    @property
    @pulumi.getter(name="selfLink")
    def self_link(self) -> pulumi.Output[str]:
        """
        Server-defined URL of this resource
        """
        return pulumi.get(self, "self_link")

    @property
    @pulumi.getter(name="updateTime")
    def update_time(self) -> pulumi.Output[str]:
        """
        The timestamp when the resource was updated.
        """
        return pulumi.get(self, "update_time")

