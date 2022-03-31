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

__all__ = ['NetworkEndpointGroupArgs', 'NetworkEndpointGroup']

@pulumi.input_type
class NetworkEndpointGroupArgs:
    def __init__(__self__, *,
                 annotations: Optional[pulumi.Input[Mapping[str, pulumi.Input[str]]]] = None,
                 app_engine: Optional[pulumi.Input['NetworkEndpointGroupAppEngineArgs']] = None,
                 cloud_function: Optional[pulumi.Input['NetworkEndpointGroupCloudFunctionArgs']] = None,
                 cloud_run: Optional[pulumi.Input['NetworkEndpointGroupCloudRunArgs']] = None,
                 default_port: Optional[pulumi.Input[int]] = None,
                 description: Optional[pulumi.Input[str]] = None,
                 name: Optional[pulumi.Input[str]] = None,
                 network: Optional[pulumi.Input[str]] = None,
                 network_endpoint_type: Optional[pulumi.Input['NetworkEndpointGroupNetworkEndpointType']] = None,
                 project: Optional[pulumi.Input[str]] = None,
                 psc_target_service: Optional[pulumi.Input[str]] = None,
                 request_id: Optional[pulumi.Input[str]] = None,
                 subnetwork: Optional[pulumi.Input[str]] = None,
                 zone: Optional[pulumi.Input[str]] = None):
        """
        The set of arguments for constructing a NetworkEndpointGroup resource.
        :param pulumi.Input[Mapping[str, pulumi.Input[str]]] annotations: Metadata defined as annotations on the network endpoint group.
        :param pulumi.Input['NetworkEndpointGroupAppEngineArgs'] app_engine: Only valid when networkEndpointType is "SERVERLESS". Only one of cloudRun, appEngine or cloudFunction may be set.
        :param pulumi.Input['NetworkEndpointGroupCloudFunctionArgs'] cloud_function: Only valid when networkEndpointType is "SERVERLESS". Only one of cloudRun, appEngine or cloudFunction may be set.
        :param pulumi.Input['NetworkEndpointGroupCloudRunArgs'] cloud_run: Only valid when networkEndpointType is "SERVERLESS". Only one of cloudRun, appEngine or cloudFunction may be set.
        :param pulumi.Input[int] default_port: The default port used if the port number is not specified in the network endpoint.
        :param pulumi.Input[str] description: An optional description of this resource. Provide this property when you create the resource.
        :param pulumi.Input[str] name: Name of the resource; provided by the client when the resource is created. The name must be 1-63 characters long, and comply with RFC1035. Specifically, the name must be 1-63 characters long and match the regular expression `[a-z]([-a-z0-9]*[a-z0-9])?` which means the first character must be a lowercase letter, and all following characters must be a dash, lowercase letter, or digit, except the last character, which cannot be a dash.
        :param pulumi.Input[str] network: The URL of the network to which all network endpoints in the NEG belong. Uses "default" project network if unspecified.
        :param pulumi.Input['NetworkEndpointGroupNetworkEndpointType'] network_endpoint_type: Type of network endpoints in this network endpoint group. Can be one of GCE_VM_IP, GCE_VM_IP_PORT, NON_GCP_PRIVATE_IP_PORT, INTERNET_FQDN_PORT, INTERNET_IP_PORT, SERVERLESS, PRIVATE_SERVICE_CONNECT.
        :param pulumi.Input[str] psc_target_service: The target service url used to set up private service connection to a Google API. An example value is: "asia-northeast3-cloudkms.googleapis.com"
        :param pulumi.Input[str] request_id: An optional request ID to identify requests. Specify a unique request ID so that if you must retry your request, the server will know to ignore the request if it has already been completed. For example, consider a situation where you make an initial request and the request times out. If you make the request again with the same request ID, the server can check if original operation with the same request ID was received, and if so, will ignore the second request. This prevents clients from accidentally creating duplicate commitments. The request ID must be a valid UUID with the exception that zero UUID is not supported ( 00000000-0000-0000-0000-000000000000).
        :param pulumi.Input[str] subnetwork: Optional URL of the subnetwork to which all network endpoints in the NEG belong.
        """
        if annotations is not None:
            pulumi.set(__self__, "annotations", annotations)
        if app_engine is not None:
            pulumi.set(__self__, "app_engine", app_engine)
        if cloud_function is not None:
            pulumi.set(__self__, "cloud_function", cloud_function)
        if cloud_run is not None:
            pulumi.set(__self__, "cloud_run", cloud_run)
        if default_port is not None:
            pulumi.set(__self__, "default_port", default_port)
        if description is not None:
            pulumi.set(__self__, "description", description)
        if name is not None:
            pulumi.set(__self__, "name", name)
        if network is not None:
            pulumi.set(__self__, "network", network)
        if network_endpoint_type is not None:
            pulumi.set(__self__, "network_endpoint_type", network_endpoint_type)
        if project is not None:
            pulumi.set(__self__, "project", project)
        if psc_target_service is not None:
            pulumi.set(__self__, "psc_target_service", psc_target_service)
        if request_id is not None:
            pulumi.set(__self__, "request_id", request_id)
        if subnetwork is not None:
            pulumi.set(__self__, "subnetwork", subnetwork)
        if zone is not None:
            pulumi.set(__self__, "zone", zone)

    @property
    @pulumi.getter
    def annotations(self) -> Optional[pulumi.Input[Mapping[str, pulumi.Input[str]]]]:
        """
        Metadata defined as annotations on the network endpoint group.
        """
        return pulumi.get(self, "annotations")

    @annotations.setter
    def annotations(self, value: Optional[pulumi.Input[Mapping[str, pulumi.Input[str]]]]):
        pulumi.set(self, "annotations", value)

    @property
    @pulumi.getter(name="appEngine")
    def app_engine(self) -> Optional[pulumi.Input['NetworkEndpointGroupAppEngineArgs']]:
        """
        Only valid when networkEndpointType is "SERVERLESS". Only one of cloudRun, appEngine or cloudFunction may be set.
        """
        return pulumi.get(self, "app_engine")

    @app_engine.setter
    def app_engine(self, value: Optional[pulumi.Input['NetworkEndpointGroupAppEngineArgs']]):
        pulumi.set(self, "app_engine", value)

    @property
    @pulumi.getter(name="cloudFunction")
    def cloud_function(self) -> Optional[pulumi.Input['NetworkEndpointGroupCloudFunctionArgs']]:
        """
        Only valid when networkEndpointType is "SERVERLESS". Only one of cloudRun, appEngine or cloudFunction may be set.
        """
        return pulumi.get(self, "cloud_function")

    @cloud_function.setter
    def cloud_function(self, value: Optional[pulumi.Input['NetworkEndpointGroupCloudFunctionArgs']]):
        pulumi.set(self, "cloud_function", value)

    @property
    @pulumi.getter(name="cloudRun")
    def cloud_run(self) -> Optional[pulumi.Input['NetworkEndpointGroupCloudRunArgs']]:
        """
        Only valid when networkEndpointType is "SERVERLESS". Only one of cloudRun, appEngine or cloudFunction may be set.
        """
        return pulumi.get(self, "cloud_run")

    @cloud_run.setter
    def cloud_run(self, value: Optional[pulumi.Input['NetworkEndpointGroupCloudRunArgs']]):
        pulumi.set(self, "cloud_run", value)

    @property
    @pulumi.getter(name="defaultPort")
    def default_port(self) -> Optional[pulumi.Input[int]]:
        """
        The default port used if the port number is not specified in the network endpoint.
        """
        return pulumi.get(self, "default_port")

    @default_port.setter
    def default_port(self, value: Optional[pulumi.Input[int]]):
        pulumi.set(self, "default_port", value)

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
    @pulumi.getter
    def name(self) -> Optional[pulumi.Input[str]]:
        """
        Name of the resource; provided by the client when the resource is created. The name must be 1-63 characters long, and comply with RFC1035. Specifically, the name must be 1-63 characters long and match the regular expression `[a-z]([-a-z0-9]*[a-z0-9])?` which means the first character must be a lowercase letter, and all following characters must be a dash, lowercase letter, or digit, except the last character, which cannot be a dash.
        """
        return pulumi.get(self, "name")

    @name.setter
    def name(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "name", value)

    @property
    @pulumi.getter
    def network(self) -> Optional[pulumi.Input[str]]:
        """
        The URL of the network to which all network endpoints in the NEG belong. Uses "default" project network if unspecified.
        """
        return pulumi.get(self, "network")

    @network.setter
    def network(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "network", value)

    @property
    @pulumi.getter(name="networkEndpointType")
    def network_endpoint_type(self) -> Optional[pulumi.Input['NetworkEndpointGroupNetworkEndpointType']]:
        """
        Type of network endpoints in this network endpoint group. Can be one of GCE_VM_IP, GCE_VM_IP_PORT, NON_GCP_PRIVATE_IP_PORT, INTERNET_FQDN_PORT, INTERNET_IP_PORT, SERVERLESS, PRIVATE_SERVICE_CONNECT.
        """
        return pulumi.get(self, "network_endpoint_type")

    @network_endpoint_type.setter
    def network_endpoint_type(self, value: Optional[pulumi.Input['NetworkEndpointGroupNetworkEndpointType']]):
        pulumi.set(self, "network_endpoint_type", value)

    @property
    @pulumi.getter
    def project(self) -> Optional[pulumi.Input[str]]:
        return pulumi.get(self, "project")

    @project.setter
    def project(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "project", value)

    @property
    @pulumi.getter(name="pscTargetService")
    def psc_target_service(self) -> Optional[pulumi.Input[str]]:
        """
        The target service url used to set up private service connection to a Google API. An example value is: "asia-northeast3-cloudkms.googleapis.com"
        """
        return pulumi.get(self, "psc_target_service")

    @psc_target_service.setter
    def psc_target_service(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "psc_target_service", value)

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
    @pulumi.getter
    def subnetwork(self) -> Optional[pulumi.Input[str]]:
        """
        Optional URL of the subnetwork to which all network endpoints in the NEG belong.
        """
        return pulumi.get(self, "subnetwork")

    @subnetwork.setter
    def subnetwork(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "subnetwork", value)

    @property
    @pulumi.getter
    def zone(self) -> Optional[pulumi.Input[str]]:
        return pulumi.get(self, "zone")

    @zone.setter
    def zone(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "zone", value)


class NetworkEndpointGroup(pulumi.CustomResource):
    @overload
    def __init__(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 annotations: Optional[pulumi.Input[Mapping[str, pulumi.Input[str]]]] = None,
                 app_engine: Optional[pulumi.Input[pulumi.InputType['NetworkEndpointGroupAppEngineArgs']]] = None,
                 cloud_function: Optional[pulumi.Input[pulumi.InputType['NetworkEndpointGroupCloudFunctionArgs']]] = None,
                 cloud_run: Optional[pulumi.Input[pulumi.InputType['NetworkEndpointGroupCloudRunArgs']]] = None,
                 default_port: Optional[pulumi.Input[int]] = None,
                 description: Optional[pulumi.Input[str]] = None,
                 name: Optional[pulumi.Input[str]] = None,
                 network: Optional[pulumi.Input[str]] = None,
                 network_endpoint_type: Optional[pulumi.Input['NetworkEndpointGroupNetworkEndpointType']] = None,
                 project: Optional[pulumi.Input[str]] = None,
                 psc_target_service: Optional[pulumi.Input[str]] = None,
                 request_id: Optional[pulumi.Input[str]] = None,
                 subnetwork: Optional[pulumi.Input[str]] = None,
                 zone: Optional[pulumi.Input[str]] = None,
                 __props__=None):
        """
        Creates a network endpoint group in the specified project using the parameters that are included in the request.

        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[Mapping[str, pulumi.Input[str]]] annotations: Metadata defined as annotations on the network endpoint group.
        :param pulumi.Input[pulumi.InputType['NetworkEndpointGroupAppEngineArgs']] app_engine: Only valid when networkEndpointType is "SERVERLESS". Only one of cloudRun, appEngine or cloudFunction may be set.
        :param pulumi.Input[pulumi.InputType['NetworkEndpointGroupCloudFunctionArgs']] cloud_function: Only valid when networkEndpointType is "SERVERLESS". Only one of cloudRun, appEngine or cloudFunction may be set.
        :param pulumi.Input[pulumi.InputType['NetworkEndpointGroupCloudRunArgs']] cloud_run: Only valid when networkEndpointType is "SERVERLESS". Only one of cloudRun, appEngine or cloudFunction may be set.
        :param pulumi.Input[int] default_port: The default port used if the port number is not specified in the network endpoint.
        :param pulumi.Input[str] description: An optional description of this resource. Provide this property when you create the resource.
        :param pulumi.Input[str] name: Name of the resource; provided by the client when the resource is created. The name must be 1-63 characters long, and comply with RFC1035. Specifically, the name must be 1-63 characters long and match the regular expression `[a-z]([-a-z0-9]*[a-z0-9])?` which means the first character must be a lowercase letter, and all following characters must be a dash, lowercase letter, or digit, except the last character, which cannot be a dash.
        :param pulumi.Input[str] network: The URL of the network to which all network endpoints in the NEG belong. Uses "default" project network if unspecified.
        :param pulumi.Input['NetworkEndpointGroupNetworkEndpointType'] network_endpoint_type: Type of network endpoints in this network endpoint group. Can be one of GCE_VM_IP, GCE_VM_IP_PORT, NON_GCP_PRIVATE_IP_PORT, INTERNET_FQDN_PORT, INTERNET_IP_PORT, SERVERLESS, PRIVATE_SERVICE_CONNECT.
        :param pulumi.Input[str] psc_target_service: The target service url used to set up private service connection to a Google API. An example value is: "asia-northeast3-cloudkms.googleapis.com"
        :param pulumi.Input[str] request_id: An optional request ID to identify requests. Specify a unique request ID so that if you must retry your request, the server will know to ignore the request if it has already been completed. For example, consider a situation where you make an initial request and the request times out. If you make the request again with the same request ID, the server can check if original operation with the same request ID was received, and if so, will ignore the second request. This prevents clients from accidentally creating duplicate commitments. The request ID must be a valid UUID with the exception that zero UUID is not supported ( 00000000-0000-0000-0000-000000000000).
        :param pulumi.Input[str] subnetwork: Optional URL of the subnetwork to which all network endpoints in the NEG belong.
        """
        ...
    @overload
    def __init__(__self__,
                 resource_name: str,
                 args: Optional[NetworkEndpointGroupArgs] = None,
                 opts: Optional[pulumi.ResourceOptions] = None):
        """
        Creates a network endpoint group in the specified project using the parameters that are included in the request.

        :param str resource_name: The name of the resource.
        :param NetworkEndpointGroupArgs args: The arguments to use to populate this resource's properties.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        ...
    def __init__(__self__, resource_name: str, *args, **kwargs):
        resource_args, opts = _utilities.get_resource_args_opts(NetworkEndpointGroupArgs, pulumi.ResourceOptions, *args, **kwargs)
        if resource_args is not None:
            __self__._internal_init(resource_name, opts, **resource_args.__dict__)
        else:
            __self__._internal_init(resource_name, *args, **kwargs)

    def _internal_init(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 annotations: Optional[pulumi.Input[Mapping[str, pulumi.Input[str]]]] = None,
                 app_engine: Optional[pulumi.Input[pulumi.InputType['NetworkEndpointGroupAppEngineArgs']]] = None,
                 cloud_function: Optional[pulumi.Input[pulumi.InputType['NetworkEndpointGroupCloudFunctionArgs']]] = None,
                 cloud_run: Optional[pulumi.Input[pulumi.InputType['NetworkEndpointGroupCloudRunArgs']]] = None,
                 default_port: Optional[pulumi.Input[int]] = None,
                 description: Optional[pulumi.Input[str]] = None,
                 name: Optional[pulumi.Input[str]] = None,
                 network: Optional[pulumi.Input[str]] = None,
                 network_endpoint_type: Optional[pulumi.Input['NetworkEndpointGroupNetworkEndpointType']] = None,
                 project: Optional[pulumi.Input[str]] = None,
                 psc_target_service: Optional[pulumi.Input[str]] = None,
                 request_id: Optional[pulumi.Input[str]] = None,
                 subnetwork: Optional[pulumi.Input[str]] = None,
                 zone: Optional[pulumi.Input[str]] = None,
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
            __props__ = NetworkEndpointGroupArgs.__new__(NetworkEndpointGroupArgs)

            __props__.__dict__["annotations"] = annotations
            __props__.__dict__["app_engine"] = app_engine
            __props__.__dict__["cloud_function"] = cloud_function
            __props__.__dict__["cloud_run"] = cloud_run
            __props__.__dict__["default_port"] = default_port
            __props__.__dict__["description"] = description
            __props__.__dict__["name"] = name
            __props__.__dict__["network"] = network
            __props__.__dict__["network_endpoint_type"] = network_endpoint_type
            __props__.__dict__["project"] = project
            __props__.__dict__["psc_target_service"] = psc_target_service
            __props__.__dict__["request_id"] = request_id
            __props__.__dict__["subnetwork"] = subnetwork
            __props__.__dict__["zone"] = zone
            __props__.__dict__["creation_timestamp"] = None
            __props__.__dict__["kind"] = None
            __props__.__dict__["region"] = None
            __props__.__dict__["self_link"] = None
            __props__.__dict__["size"] = None
        super(NetworkEndpointGroup, __self__).__init__(
            'google-native:compute/v1:NetworkEndpointGroup',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name: str,
            id: pulumi.Input[str],
            opts: Optional[pulumi.ResourceOptions] = None) -> 'NetworkEndpointGroup':
        """
        Get an existing NetworkEndpointGroup resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param pulumi.Input[str] id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = NetworkEndpointGroupArgs.__new__(NetworkEndpointGroupArgs)

        __props__.__dict__["annotations"] = None
        __props__.__dict__["app_engine"] = None
        __props__.__dict__["cloud_function"] = None
        __props__.__dict__["cloud_run"] = None
        __props__.__dict__["creation_timestamp"] = None
        __props__.__dict__["default_port"] = None
        __props__.__dict__["description"] = None
        __props__.__dict__["kind"] = None
        __props__.__dict__["name"] = None
        __props__.__dict__["network"] = None
        __props__.__dict__["network_endpoint_type"] = None
        __props__.__dict__["psc_target_service"] = None
        __props__.__dict__["region"] = None
        __props__.__dict__["self_link"] = None
        __props__.__dict__["size"] = None
        __props__.__dict__["subnetwork"] = None
        __props__.__dict__["zone"] = None
        return NetworkEndpointGroup(resource_name, opts=opts, __props__=__props__)

    @property
    @pulumi.getter
    def annotations(self) -> pulumi.Output[Mapping[str, str]]:
        """
        Metadata defined as annotations on the network endpoint group.
        """
        return pulumi.get(self, "annotations")

    @property
    @pulumi.getter(name="appEngine")
    def app_engine(self) -> pulumi.Output['outputs.NetworkEndpointGroupAppEngineResponse']:
        """
        Only valid when networkEndpointType is "SERVERLESS". Only one of cloudRun, appEngine or cloudFunction may be set.
        """
        return pulumi.get(self, "app_engine")

    @property
    @pulumi.getter(name="cloudFunction")
    def cloud_function(self) -> pulumi.Output['outputs.NetworkEndpointGroupCloudFunctionResponse']:
        """
        Only valid when networkEndpointType is "SERVERLESS". Only one of cloudRun, appEngine or cloudFunction may be set.
        """
        return pulumi.get(self, "cloud_function")

    @property
    @pulumi.getter(name="cloudRun")
    def cloud_run(self) -> pulumi.Output['outputs.NetworkEndpointGroupCloudRunResponse']:
        """
        Only valid when networkEndpointType is "SERVERLESS". Only one of cloudRun, appEngine or cloudFunction may be set.
        """
        return pulumi.get(self, "cloud_run")

    @property
    @pulumi.getter(name="creationTimestamp")
    def creation_timestamp(self) -> pulumi.Output[str]:
        """
        Creation timestamp in RFC3339 text format.
        """
        return pulumi.get(self, "creation_timestamp")

    @property
    @pulumi.getter(name="defaultPort")
    def default_port(self) -> pulumi.Output[int]:
        """
        The default port used if the port number is not specified in the network endpoint.
        """
        return pulumi.get(self, "default_port")

    @property
    @pulumi.getter
    def description(self) -> pulumi.Output[str]:
        """
        An optional description of this resource. Provide this property when you create the resource.
        """
        return pulumi.get(self, "description")

    @property
    @pulumi.getter
    def kind(self) -> pulumi.Output[str]:
        """
        Type of the resource. Always compute#networkEndpointGroup for network endpoint group.
        """
        return pulumi.get(self, "kind")

    @property
    @pulumi.getter
    def name(self) -> pulumi.Output[str]:
        """
        Name of the resource; provided by the client when the resource is created. The name must be 1-63 characters long, and comply with RFC1035. Specifically, the name must be 1-63 characters long and match the regular expression `[a-z]([-a-z0-9]*[a-z0-9])?` which means the first character must be a lowercase letter, and all following characters must be a dash, lowercase letter, or digit, except the last character, which cannot be a dash.
        """
        return pulumi.get(self, "name")

    @property
    @pulumi.getter
    def network(self) -> pulumi.Output[str]:
        """
        The URL of the network to which all network endpoints in the NEG belong. Uses "default" project network if unspecified.
        """
        return pulumi.get(self, "network")

    @property
    @pulumi.getter(name="networkEndpointType")
    def network_endpoint_type(self) -> pulumi.Output[str]:
        """
        Type of network endpoints in this network endpoint group. Can be one of GCE_VM_IP, GCE_VM_IP_PORT, NON_GCP_PRIVATE_IP_PORT, INTERNET_FQDN_PORT, INTERNET_IP_PORT, SERVERLESS, PRIVATE_SERVICE_CONNECT.
        """
        return pulumi.get(self, "network_endpoint_type")

    @property
    @pulumi.getter(name="pscTargetService")
    def psc_target_service(self) -> pulumi.Output[str]:
        """
        The target service url used to set up private service connection to a Google API. An example value is: "asia-northeast3-cloudkms.googleapis.com"
        """
        return pulumi.get(self, "psc_target_service")

    @property
    @pulumi.getter
    def region(self) -> pulumi.Output[str]:
        """
        The URL of the region where the network endpoint group is located.
        """
        return pulumi.get(self, "region")

    @property
    @pulumi.getter(name="selfLink")
    def self_link(self) -> pulumi.Output[str]:
        """
        Server-defined URL for the resource.
        """
        return pulumi.get(self, "self_link")

    @property
    @pulumi.getter
    def size(self) -> pulumi.Output[int]:
        """
        [Output only] Number of network endpoints in the network endpoint group.
        """
        return pulumi.get(self, "size")

    @property
    @pulumi.getter
    def subnetwork(self) -> pulumi.Output[str]:
        """
        Optional URL of the subnetwork to which all network endpoints in the NEG belong.
        """
        return pulumi.get(self, "subnetwork")

    @property
    @pulumi.getter
    def zone(self) -> pulumi.Output[str]:
        """
        The URL of the zone where the network endpoint group is located.
        """
        return pulumi.get(self, "zone")

