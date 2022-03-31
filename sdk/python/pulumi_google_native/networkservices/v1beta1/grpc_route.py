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

__all__ = ['GrpcRouteArgs', 'GrpcRoute']

@pulumi.input_type
class GrpcRouteArgs:
    def __init__(__self__, *,
                 grpc_route_id: pulumi.Input[str],
                 hostnames: pulumi.Input[Sequence[pulumi.Input[str]]],
                 rules: pulumi.Input[Sequence[pulumi.Input['GrpcRouteRouteRuleArgs']]],
                 description: Optional[pulumi.Input[str]] = None,
                 gateways: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]] = None,
                 labels: Optional[pulumi.Input[Mapping[str, pulumi.Input[str]]]] = None,
                 location: Optional[pulumi.Input[str]] = None,
                 meshes: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]] = None,
                 name: Optional[pulumi.Input[str]] = None,
                 project: Optional[pulumi.Input[str]] = None):
        """
        The set of arguments for constructing a GrpcRoute resource.
        :param pulumi.Input[str] grpc_route_id: Required. Short name of the GrpcRoute resource to be created.
        :param pulumi.Input[Sequence[pulumi.Input[str]]] hostnames: Service hostnames with an optional port for which this route describes traffic. Format: [:] Hostname is the fully qualified domain name of a network host. This matches the RFC 1123 definition of a hostname with 2 notable exceptions: - IPs are not allowed. - A hostname may be prefixed with a wildcard label (*.). The wildcard label must appear by itself as the first label. Hostname can be "precise" which is a domain name without the terminating dot of a network host (e.g. "foo.example.com") or "wildcard", which is a domain name prefixed with a single wildcard label (e.g. *.example.com). Note that as per RFC1035 and RFC1123, a label must consist of lower case alphanumeric characters or '-', and must start and end with an alphanumeric character. No other punctuation is allowed. The routes associated with a Mesh or Gateway must have unique hostnames. If you attempt to attach multiple routes with conflicting hostnames, the configuration will be rejected. For example, while it is acceptable for routes for the hostnames "*.foo.bar.com" and "*.bar.com" to be associated with the same route, it is not possible to associate two routes both with "*.bar.com" or both with "bar.com". If a port is specified, then gRPC clients must use the channel URI with the port to match this rule (i.e. "xds:///service:123"), otherwise they must supply the URI without a port (i.e. "xds:///service").
        :param pulumi.Input[Sequence[pulumi.Input['GrpcRouteRouteRuleArgs']]] rules: A list of detailed rules defining how to route traffic. Within a single GrpcRoute, the GrpcRoute.RouteAction associated with the first matching GrpcRoute.RouteRule will be executed. At least one rule must be supplied.
        :param pulumi.Input[str] description: Optional. A free-text description of the resource. Max length 1024 characters.
        :param pulumi.Input[Sequence[pulumi.Input[str]]] gateways: Optional. Gateways defines a list of gateways this GrpcRoute is attached to, as one of the routing rules to route the requests served by the gateway. Each gateway reference should match the pattern: `projects/*/locations/global/gateways/`
        :param pulumi.Input[Mapping[str, pulumi.Input[str]]] labels: Optional. Set of label tags associated with the GrpcRoute resource.
        :param pulumi.Input[Sequence[pulumi.Input[str]]] meshes: Optional. Meshes defines a list of meshes this GrpcRoute is attached to, as one of the routing rules to route the requests served by the mesh. Each mesh reference should match the pattern: `projects/*/locations/global/meshes/`
        :param pulumi.Input[str] name: Name of the GrpcRoute resource. It matches pattern `projects/*/locations/global/grpcRoutes/`
        """
        pulumi.set(__self__, "grpc_route_id", grpc_route_id)
        pulumi.set(__self__, "hostnames", hostnames)
        pulumi.set(__self__, "rules", rules)
        if description is not None:
            pulumi.set(__self__, "description", description)
        if gateways is not None:
            pulumi.set(__self__, "gateways", gateways)
        if labels is not None:
            pulumi.set(__self__, "labels", labels)
        if location is not None:
            pulumi.set(__self__, "location", location)
        if meshes is not None:
            pulumi.set(__self__, "meshes", meshes)
        if name is not None:
            pulumi.set(__self__, "name", name)
        if project is not None:
            pulumi.set(__self__, "project", project)

    @property
    @pulumi.getter(name="grpcRouteId")
    def grpc_route_id(self) -> pulumi.Input[str]:
        """
        Required. Short name of the GrpcRoute resource to be created.
        """
        return pulumi.get(self, "grpc_route_id")

    @grpc_route_id.setter
    def grpc_route_id(self, value: pulumi.Input[str]):
        pulumi.set(self, "grpc_route_id", value)

    @property
    @pulumi.getter
    def hostnames(self) -> pulumi.Input[Sequence[pulumi.Input[str]]]:
        """
        Service hostnames with an optional port for which this route describes traffic. Format: [:] Hostname is the fully qualified domain name of a network host. This matches the RFC 1123 definition of a hostname with 2 notable exceptions: - IPs are not allowed. - A hostname may be prefixed with a wildcard label (*.). The wildcard label must appear by itself as the first label. Hostname can be "precise" which is a domain name without the terminating dot of a network host (e.g. "foo.example.com") or "wildcard", which is a domain name prefixed with a single wildcard label (e.g. *.example.com). Note that as per RFC1035 and RFC1123, a label must consist of lower case alphanumeric characters or '-', and must start and end with an alphanumeric character. No other punctuation is allowed. The routes associated with a Mesh or Gateway must have unique hostnames. If you attempt to attach multiple routes with conflicting hostnames, the configuration will be rejected. For example, while it is acceptable for routes for the hostnames "*.foo.bar.com" and "*.bar.com" to be associated with the same route, it is not possible to associate two routes both with "*.bar.com" or both with "bar.com". If a port is specified, then gRPC clients must use the channel URI with the port to match this rule (i.e. "xds:///service:123"), otherwise they must supply the URI without a port (i.e. "xds:///service").
        """
        return pulumi.get(self, "hostnames")

    @hostnames.setter
    def hostnames(self, value: pulumi.Input[Sequence[pulumi.Input[str]]]):
        pulumi.set(self, "hostnames", value)

    @property
    @pulumi.getter
    def rules(self) -> pulumi.Input[Sequence[pulumi.Input['GrpcRouteRouteRuleArgs']]]:
        """
        A list of detailed rules defining how to route traffic. Within a single GrpcRoute, the GrpcRoute.RouteAction associated with the first matching GrpcRoute.RouteRule will be executed. At least one rule must be supplied.
        """
        return pulumi.get(self, "rules")

    @rules.setter
    def rules(self, value: pulumi.Input[Sequence[pulumi.Input['GrpcRouteRouteRuleArgs']]]):
        pulumi.set(self, "rules", value)

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
        Optional. Gateways defines a list of gateways this GrpcRoute is attached to, as one of the routing rules to route the requests served by the gateway. Each gateway reference should match the pattern: `projects/*/locations/global/gateways/`
        """
        return pulumi.get(self, "gateways")

    @gateways.setter
    def gateways(self, value: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]]):
        pulumi.set(self, "gateways", value)

    @property
    @pulumi.getter
    def labels(self) -> Optional[pulumi.Input[Mapping[str, pulumi.Input[str]]]]:
        """
        Optional. Set of label tags associated with the GrpcRoute resource.
        """
        return pulumi.get(self, "labels")

    @labels.setter
    def labels(self, value: Optional[pulumi.Input[Mapping[str, pulumi.Input[str]]]]):
        pulumi.set(self, "labels", value)

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
        Optional. Meshes defines a list of meshes this GrpcRoute is attached to, as one of the routing rules to route the requests served by the mesh. Each mesh reference should match the pattern: `projects/*/locations/global/meshes/`
        """
        return pulumi.get(self, "meshes")

    @meshes.setter
    def meshes(self, value: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]]):
        pulumi.set(self, "meshes", value)

    @property
    @pulumi.getter
    def name(self) -> Optional[pulumi.Input[str]]:
        """
        Name of the GrpcRoute resource. It matches pattern `projects/*/locations/global/grpcRoutes/`
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


class GrpcRoute(pulumi.CustomResource):
    @overload
    def __init__(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 description: Optional[pulumi.Input[str]] = None,
                 gateways: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]] = None,
                 grpc_route_id: Optional[pulumi.Input[str]] = None,
                 hostnames: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]] = None,
                 labels: Optional[pulumi.Input[Mapping[str, pulumi.Input[str]]]] = None,
                 location: Optional[pulumi.Input[str]] = None,
                 meshes: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]] = None,
                 name: Optional[pulumi.Input[str]] = None,
                 project: Optional[pulumi.Input[str]] = None,
                 rules: Optional[pulumi.Input[Sequence[pulumi.Input[pulumi.InputType['GrpcRouteRouteRuleArgs']]]]] = None,
                 __props__=None):
        """
        Creates a new GrpcRoute in a given project and location.

        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] description: Optional. A free-text description of the resource. Max length 1024 characters.
        :param pulumi.Input[Sequence[pulumi.Input[str]]] gateways: Optional. Gateways defines a list of gateways this GrpcRoute is attached to, as one of the routing rules to route the requests served by the gateway. Each gateway reference should match the pattern: `projects/*/locations/global/gateways/`
        :param pulumi.Input[str] grpc_route_id: Required. Short name of the GrpcRoute resource to be created.
        :param pulumi.Input[Sequence[pulumi.Input[str]]] hostnames: Service hostnames with an optional port for which this route describes traffic. Format: [:] Hostname is the fully qualified domain name of a network host. This matches the RFC 1123 definition of a hostname with 2 notable exceptions: - IPs are not allowed. - A hostname may be prefixed with a wildcard label (*.). The wildcard label must appear by itself as the first label. Hostname can be "precise" which is a domain name without the terminating dot of a network host (e.g. "foo.example.com") or "wildcard", which is a domain name prefixed with a single wildcard label (e.g. *.example.com). Note that as per RFC1035 and RFC1123, a label must consist of lower case alphanumeric characters or '-', and must start and end with an alphanumeric character. No other punctuation is allowed. The routes associated with a Mesh or Gateway must have unique hostnames. If you attempt to attach multiple routes with conflicting hostnames, the configuration will be rejected. For example, while it is acceptable for routes for the hostnames "*.foo.bar.com" and "*.bar.com" to be associated with the same route, it is not possible to associate two routes both with "*.bar.com" or both with "bar.com". If a port is specified, then gRPC clients must use the channel URI with the port to match this rule (i.e. "xds:///service:123"), otherwise they must supply the URI without a port (i.e. "xds:///service").
        :param pulumi.Input[Mapping[str, pulumi.Input[str]]] labels: Optional. Set of label tags associated with the GrpcRoute resource.
        :param pulumi.Input[Sequence[pulumi.Input[str]]] meshes: Optional. Meshes defines a list of meshes this GrpcRoute is attached to, as one of the routing rules to route the requests served by the mesh. Each mesh reference should match the pattern: `projects/*/locations/global/meshes/`
        :param pulumi.Input[str] name: Name of the GrpcRoute resource. It matches pattern `projects/*/locations/global/grpcRoutes/`
        :param pulumi.Input[Sequence[pulumi.Input[pulumi.InputType['GrpcRouteRouteRuleArgs']]]] rules: A list of detailed rules defining how to route traffic. Within a single GrpcRoute, the GrpcRoute.RouteAction associated with the first matching GrpcRoute.RouteRule will be executed. At least one rule must be supplied.
        """
        ...
    @overload
    def __init__(__self__,
                 resource_name: str,
                 args: GrpcRouteArgs,
                 opts: Optional[pulumi.ResourceOptions] = None):
        """
        Creates a new GrpcRoute in a given project and location.

        :param str resource_name: The name of the resource.
        :param GrpcRouteArgs args: The arguments to use to populate this resource's properties.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        ...
    def __init__(__self__, resource_name: str, *args, **kwargs):
        resource_args, opts = _utilities.get_resource_args_opts(GrpcRouteArgs, pulumi.ResourceOptions, *args, **kwargs)
        if resource_args is not None:
            __self__._internal_init(resource_name, opts, **resource_args.__dict__)
        else:
            __self__._internal_init(resource_name, *args, **kwargs)

    def _internal_init(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 description: Optional[pulumi.Input[str]] = None,
                 gateways: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]] = None,
                 grpc_route_id: Optional[pulumi.Input[str]] = None,
                 hostnames: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]] = None,
                 labels: Optional[pulumi.Input[Mapping[str, pulumi.Input[str]]]] = None,
                 location: Optional[pulumi.Input[str]] = None,
                 meshes: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]] = None,
                 name: Optional[pulumi.Input[str]] = None,
                 project: Optional[pulumi.Input[str]] = None,
                 rules: Optional[pulumi.Input[Sequence[pulumi.Input[pulumi.InputType['GrpcRouteRouteRuleArgs']]]]] = None,
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
            __props__ = GrpcRouteArgs.__new__(GrpcRouteArgs)

            __props__.__dict__["description"] = description
            __props__.__dict__["gateways"] = gateways
            if grpc_route_id is None and not opts.urn:
                raise TypeError("Missing required property 'grpc_route_id'")
            __props__.__dict__["grpc_route_id"] = grpc_route_id
            if hostnames is None and not opts.urn:
                raise TypeError("Missing required property 'hostnames'")
            __props__.__dict__["hostnames"] = hostnames
            __props__.__dict__["labels"] = labels
            __props__.__dict__["location"] = location
            __props__.__dict__["meshes"] = meshes
            __props__.__dict__["name"] = name
            __props__.__dict__["project"] = project
            if rules is None and not opts.urn:
                raise TypeError("Missing required property 'rules'")
            __props__.__dict__["rules"] = rules
            __props__.__dict__["create_time"] = None
            __props__.__dict__["self_link"] = None
            __props__.__dict__["update_time"] = None
        super(GrpcRoute, __self__).__init__(
            'google-native:networkservices/v1beta1:GrpcRoute',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name: str,
            id: pulumi.Input[str],
            opts: Optional[pulumi.ResourceOptions] = None) -> 'GrpcRoute':
        """
        Get an existing GrpcRoute resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param pulumi.Input[str] id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = GrpcRouteArgs.__new__(GrpcRouteArgs)

        __props__.__dict__["create_time"] = None
        __props__.__dict__["description"] = None
        __props__.__dict__["gateways"] = None
        __props__.__dict__["hostnames"] = None
        __props__.__dict__["labels"] = None
        __props__.__dict__["meshes"] = None
        __props__.__dict__["name"] = None
        __props__.__dict__["rules"] = None
        __props__.__dict__["self_link"] = None
        __props__.__dict__["update_time"] = None
        return GrpcRoute(resource_name, opts=opts, __props__=__props__)

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
        Optional. Gateways defines a list of gateways this GrpcRoute is attached to, as one of the routing rules to route the requests served by the gateway. Each gateway reference should match the pattern: `projects/*/locations/global/gateways/`
        """
        return pulumi.get(self, "gateways")

    @property
    @pulumi.getter
    def hostnames(self) -> pulumi.Output[Sequence[str]]:
        """
        Service hostnames with an optional port for which this route describes traffic. Format: [:] Hostname is the fully qualified domain name of a network host. This matches the RFC 1123 definition of a hostname with 2 notable exceptions: - IPs are not allowed. - A hostname may be prefixed with a wildcard label (*.). The wildcard label must appear by itself as the first label. Hostname can be "precise" which is a domain name without the terminating dot of a network host (e.g. "foo.example.com") or "wildcard", which is a domain name prefixed with a single wildcard label (e.g. *.example.com). Note that as per RFC1035 and RFC1123, a label must consist of lower case alphanumeric characters or '-', and must start and end with an alphanumeric character. No other punctuation is allowed. The routes associated with a Mesh or Gateway must have unique hostnames. If you attempt to attach multiple routes with conflicting hostnames, the configuration will be rejected. For example, while it is acceptable for routes for the hostnames "*.foo.bar.com" and "*.bar.com" to be associated with the same route, it is not possible to associate two routes both with "*.bar.com" or both with "bar.com". If a port is specified, then gRPC clients must use the channel URI with the port to match this rule (i.e. "xds:///service:123"), otherwise they must supply the URI without a port (i.e. "xds:///service").
        """
        return pulumi.get(self, "hostnames")

    @property
    @pulumi.getter
    def labels(self) -> pulumi.Output[Mapping[str, str]]:
        """
        Optional. Set of label tags associated with the GrpcRoute resource.
        """
        return pulumi.get(self, "labels")

    @property
    @pulumi.getter
    def meshes(self) -> pulumi.Output[Sequence[str]]:
        """
        Optional. Meshes defines a list of meshes this GrpcRoute is attached to, as one of the routing rules to route the requests served by the mesh. Each mesh reference should match the pattern: `projects/*/locations/global/meshes/`
        """
        return pulumi.get(self, "meshes")

    @property
    @pulumi.getter
    def name(self) -> pulumi.Output[str]:
        """
        Name of the GrpcRoute resource. It matches pattern `projects/*/locations/global/grpcRoutes/`
        """
        return pulumi.get(self, "name")

    @property
    @pulumi.getter
    def rules(self) -> pulumi.Output[Sequence['outputs.GrpcRouteRouteRuleResponse']]:
        """
        A list of detailed rules defining how to route traffic. Within a single GrpcRoute, the GrpcRoute.RouteAction associated with the first matching GrpcRoute.RouteRule will be executed. At least one rule must be supplied.
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

