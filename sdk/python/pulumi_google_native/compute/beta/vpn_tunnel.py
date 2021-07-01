# coding=utf-8
# *** WARNING: this file was generated by the Pulumi SDK Generator. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import warnings
import pulumi
import pulumi.runtime
from typing import Any, Mapping, Optional, Sequence, Union, overload
from ... import _utilities

__all__ = ['VpnTunnelArgs', 'VpnTunnel']

@pulumi.input_type
class VpnTunnelArgs:
    def __init__(__self__, *,
                 project: pulumi.Input[str],
                 region: pulumi.Input[str],
                 description: Optional[pulumi.Input[str]] = None,
                 ike_version: Optional[pulumi.Input[int]] = None,
                 labels: Optional[pulumi.Input[Mapping[str, str]]] = None,
                 local_traffic_selector: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]] = None,
                 name: Optional[pulumi.Input[str]] = None,
                 peer_external_gateway: Optional[pulumi.Input[str]] = None,
                 peer_external_gateway_interface: Optional[pulumi.Input[int]] = None,
                 peer_gcp_gateway: Optional[pulumi.Input[str]] = None,
                 peer_ip: Optional[pulumi.Input[str]] = None,
                 remote_traffic_selector: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]] = None,
                 request_id: Optional[pulumi.Input[str]] = None,
                 router: Optional[pulumi.Input[str]] = None,
                 shared_secret: Optional[pulumi.Input[str]] = None,
                 shared_secret_hash: Optional[pulumi.Input[str]] = None,
                 target_vpn_gateway: Optional[pulumi.Input[str]] = None,
                 vpn_gateway: Optional[pulumi.Input[str]] = None,
                 vpn_gateway_interface: Optional[pulumi.Input[int]] = None):
        """
        The set of arguments for constructing a VpnTunnel resource.
        :param pulumi.Input[str] description: An optional description of this resource. Provide this property when you create the resource.
        :param pulumi.Input[int] ike_version: IKE protocol version to use when establishing the VPN tunnel with the peer VPN gateway. Acceptable IKE versions are 1 or 2. The default version is 2.
        :param pulumi.Input[Mapping[str, str]] labels: Labels for this resource. These can only be added or modified by the setLabels method. Each label key/value pair must comply with RFC1035. Label values may be empty.
        :param pulumi.Input[Sequence[pulumi.Input[str]]] local_traffic_selector: Local traffic selector to use when establishing the VPN tunnel with the peer VPN gateway. The value should be a CIDR formatted string, for example: 192.168.0.0/16. The ranges must be disjoint. Only IPv4 is supported.
        :param pulumi.Input[str] name: Name of the resource. Provided by the client when the resource is created. The name must be 1-63 characters long, and comply with RFC1035. Specifically, the name must be 1-63 characters long and match the regular expression `[a-z]([-a-z0-9]*[a-z0-9])?` which means the first character must be a lowercase letter, and all following characters must be a dash, lowercase letter, or digit, except the last character, which cannot be a dash.
        :param pulumi.Input[str] peer_external_gateway: URL of the peer side external VPN gateway to which this VPN tunnel is connected. Provided by the client when the VPN tunnel is created. This field is exclusive with the field peerGcpGateway.
        :param pulumi.Input[int] peer_external_gateway_interface: The interface ID of the external VPN gateway to which this VPN tunnel is connected. Provided by the client when the VPN tunnel is created.
        :param pulumi.Input[str] peer_gcp_gateway: URL of the peer side HA GCP VPN gateway to which this VPN tunnel is connected. Provided by the client when the VPN tunnel is created. This field can be used when creating highly available VPN from VPC network to VPC network, the field is exclusive with the field peerExternalGateway. If provided, the VPN tunnel will automatically use the same vpnGatewayInterface ID in the peer GCP VPN gateway.
        :param pulumi.Input[str] peer_ip: IP address of the peer VPN gateway. Only IPv4 is supported.
        :param pulumi.Input[Sequence[pulumi.Input[str]]] remote_traffic_selector: Remote traffic selectors to use when establishing the VPN tunnel with the peer VPN gateway. The value should be a CIDR formatted string, for example: 192.168.0.0/16. The ranges should be disjoint. Only IPv4 is supported.
        :param pulumi.Input[str] router: URL of the router resource to be used for dynamic routing.
        :param pulumi.Input[str] shared_secret: Shared secret used to set the secure session between the Cloud VPN gateway and the peer VPN gateway.
        :param pulumi.Input[str] shared_secret_hash: Hash of the shared secret.
        :param pulumi.Input[str] target_vpn_gateway: URL of the Target VPN gateway with which this VPN tunnel is associated. Provided by the client when the VPN tunnel is created.
        :param pulumi.Input[str] vpn_gateway: URL of the VPN gateway with which this VPN tunnel is associated. Provided by the client when the VPN tunnel is created. This must be used (instead of target_vpn_gateway) if a High Availability VPN gateway resource is created.
        :param pulumi.Input[int] vpn_gateway_interface: The interface ID of the VPN gateway with which this VPN tunnel is associated.
        """
        pulumi.set(__self__, "project", project)
        pulumi.set(__self__, "region", region)
        if description is not None:
            pulumi.set(__self__, "description", description)
        if ike_version is not None:
            pulumi.set(__self__, "ike_version", ike_version)
        if labels is not None:
            pulumi.set(__self__, "labels", labels)
        if local_traffic_selector is not None:
            pulumi.set(__self__, "local_traffic_selector", local_traffic_selector)
        if name is not None:
            pulumi.set(__self__, "name", name)
        if peer_external_gateway is not None:
            pulumi.set(__self__, "peer_external_gateway", peer_external_gateway)
        if peer_external_gateway_interface is not None:
            pulumi.set(__self__, "peer_external_gateway_interface", peer_external_gateway_interface)
        if peer_gcp_gateway is not None:
            pulumi.set(__self__, "peer_gcp_gateway", peer_gcp_gateway)
        if peer_ip is not None:
            pulumi.set(__self__, "peer_ip", peer_ip)
        if remote_traffic_selector is not None:
            pulumi.set(__self__, "remote_traffic_selector", remote_traffic_selector)
        if request_id is not None:
            pulumi.set(__self__, "request_id", request_id)
        if router is not None:
            pulumi.set(__self__, "router", router)
        if shared_secret is not None:
            pulumi.set(__self__, "shared_secret", shared_secret)
        if shared_secret_hash is not None:
            pulumi.set(__self__, "shared_secret_hash", shared_secret_hash)
        if target_vpn_gateway is not None:
            pulumi.set(__self__, "target_vpn_gateway", target_vpn_gateway)
        if vpn_gateway is not None:
            pulumi.set(__self__, "vpn_gateway", vpn_gateway)
        if vpn_gateway_interface is not None:
            pulumi.set(__self__, "vpn_gateway_interface", vpn_gateway_interface)

    @property
    @pulumi.getter
    def project(self) -> pulumi.Input[str]:
        return pulumi.get(self, "project")

    @project.setter
    def project(self, value: pulumi.Input[str]):
        pulumi.set(self, "project", value)

    @property
    @pulumi.getter
    def region(self) -> pulumi.Input[str]:
        return pulumi.get(self, "region")

    @region.setter
    def region(self, value: pulumi.Input[str]):
        pulumi.set(self, "region", value)

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
    @pulumi.getter(name="ikeVersion")
    def ike_version(self) -> Optional[pulumi.Input[int]]:
        """
        IKE protocol version to use when establishing the VPN tunnel with the peer VPN gateway. Acceptable IKE versions are 1 or 2. The default version is 2.
        """
        return pulumi.get(self, "ike_version")

    @ike_version.setter
    def ike_version(self, value: Optional[pulumi.Input[int]]):
        pulumi.set(self, "ike_version", value)

    @property
    @pulumi.getter
    def labels(self) -> Optional[pulumi.Input[Mapping[str, str]]]:
        """
        Labels for this resource. These can only be added or modified by the setLabels method. Each label key/value pair must comply with RFC1035. Label values may be empty.
        """
        return pulumi.get(self, "labels")

    @labels.setter
    def labels(self, value: Optional[pulumi.Input[Mapping[str, str]]]):
        pulumi.set(self, "labels", value)

    @property
    @pulumi.getter(name="localTrafficSelector")
    def local_traffic_selector(self) -> Optional[pulumi.Input[Sequence[pulumi.Input[str]]]]:
        """
        Local traffic selector to use when establishing the VPN tunnel with the peer VPN gateway. The value should be a CIDR formatted string, for example: 192.168.0.0/16. The ranges must be disjoint. Only IPv4 is supported.
        """
        return pulumi.get(self, "local_traffic_selector")

    @local_traffic_selector.setter
    def local_traffic_selector(self, value: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]]):
        pulumi.set(self, "local_traffic_selector", value)

    @property
    @pulumi.getter
    def name(self) -> Optional[pulumi.Input[str]]:
        """
        Name of the resource. Provided by the client when the resource is created. The name must be 1-63 characters long, and comply with RFC1035. Specifically, the name must be 1-63 characters long and match the regular expression `[a-z]([-a-z0-9]*[a-z0-9])?` which means the first character must be a lowercase letter, and all following characters must be a dash, lowercase letter, or digit, except the last character, which cannot be a dash.
        """
        return pulumi.get(self, "name")

    @name.setter
    def name(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "name", value)

    @property
    @pulumi.getter(name="peerExternalGateway")
    def peer_external_gateway(self) -> Optional[pulumi.Input[str]]:
        """
        URL of the peer side external VPN gateway to which this VPN tunnel is connected. Provided by the client when the VPN tunnel is created. This field is exclusive with the field peerGcpGateway.
        """
        return pulumi.get(self, "peer_external_gateway")

    @peer_external_gateway.setter
    def peer_external_gateway(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "peer_external_gateway", value)

    @property
    @pulumi.getter(name="peerExternalGatewayInterface")
    def peer_external_gateway_interface(self) -> Optional[pulumi.Input[int]]:
        """
        The interface ID of the external VPN gateway to which this VPN tunnel is connected. Provided by the client when the VPN tunnel is created.
        """
        return pulumi.get(self, "peer_external_gateway_interface")

    @peer_external_gateway_interface.setter
    def peer_external_gateway_interface(self, value: Optional[pulumi.Input[int]]):
        pulumi.set(self, "peer_external_gateway_interface", value)

    @property
    @pulumi.getter(name="peerGcpGateway")
    def peer_gcp_gateway(self) -> Optional[pulumi.Input[str]]:
        """
        URL of the peer side HA GCP VPN gateway to which this VPN tunnel is connected. Provided by the client when the VPN tunnel is created. This field can be used when creating highly available VPN from VPC network to VPC network, the field is exclusive with the field peerExternalGateway. If provided, the VPN tunnel will automatically use the same vpnGatewayInterface ID in the peer GCP VPN gateway.
        """
        return pulumi.get(self, "peer_gcp_gateway")

    @peer_gcp_gateway.setter
    def peer_gcp_gateway(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "peer_gcp_gateway", value)

    @property
    @pulumi.getter(name="peerIp")
    def peer_ip(self) -> Optional[pulumi.Input[str]]:
        """
        IP address of the peer VPN gateway. Only IPv4 is supported.
        """
        return pulumi.get(self, "peer_ip")

    @peer_ip.setter
    def peer_ip(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "peer_ip", value)

    @property
    @pulumi.getter(name="remoteTrafficSelector")
    def remote_traffic_selector(self) -> Optional[pulumi.Input[Sequence[pulumi.Input[str]]]]:
        """
        Remote traffic selectors to use when establishing the VPN tunnel with the peer VPN gateway. The value should be a CIDR formatted string, for example: 192.168.0.0/16. The ranges should be disjoint. Only IPv4 is supported.
        """
        return pulumi.get(self, "remote_traffic_selector")

    @remote_traffic_selector.setter
    def remote_traffic_selector(self, value: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]]):
        pulumi.set(self, "remote_traffic_selector", value)

    @property
    @pulumi.getter(name="requestId")
    def request_id(self) -> Optional[pulumi.Input[str]]:
        return pulumi.get(self, "request_id")

    @request_id.setter
    def request_id(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "request_id", value)

    @property
    @pulumi.getter
    def router(self) -> Optional[pulumi.Input[str]]:
        """
        URL of the router resource to be used for dynamic routing.
        """
        return pulumi.get(self, "router")

    @router.setter
    def router(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "router", value)

    @property
    @pulumi.getter(name="sharedSecret")
    def shared_secret(self) -> Optional[pulumi.Input[str]]:
        """
        Shared secret used to set the secure session between the Cloud VPN gateway and the peer VPN gateway.
        """
        return pulumi.get(self, "shared_secret")

    @shared_secret.setter
    def shared_secret(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "shared_secret", value)

    @property
    @pulumi.getter(name="sharedSecretHash")
    def shared_secret_hash(self) -> Optional[pulumi.Input[str]]:
        """
        Hash of the shared secret.
        """
        return pulumi.get(self, "shared_secret_hash")

    @shared_secret_hash.setter
    def shared_secret_hash(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "shared_secret_hash", value)

    @property
    @pulumi.getter(name="targetVpnGateway")
    def target_vpn_gateway(self) -> Optional[pulumi.Input[str]]:
        """
        URL of the Target VPN gateway with which this VPN tunnel is associated. Provided by the client when the VPN tunnel is created.
        """
        return pulumi.get(self, "target_vpn_gateway")

    @target_vpn_gateway.setter
    def target_vpn_gateway(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "target_vpn_gateway", value)

    @property
    @pulumi.getter(name="vpnGateway")
    def vpn_gateway(self) -> Optional[pulumi.Input[str]]:
        """
        URL of the VPN gateway with which this VPN tunnel is associated. Provided by the client when the VPN tunnel is created. This must be used (instead of target_vpn_gateway) if a High Availability VPN gateway resource is created.
        """
        return pulumi.get(self, "vpn_gateway")

    @vpn_gateway.setter
    def vpn_gateway(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "vpn_gateway", value)

    @property
    @pulumi.getter(name="vpnGatewayInterface")
    def vpn_gateway_interface(self) -> Optional[pulumi.Input[int]]:
        """
        The interface ID of the VPN gateway with which this VPN tunnel is associated.
        """
        return pulumi.get(self, "vpn_gateway_interface")

    @vpn_gateway_interface.setter
    def vpn_gateway_interface(self, value: Optional[pulumi.Input[int]]):
        pulumi.set(self, "vpn_gateway_interface", value)


class VpnTunnel(pulumi.CustomResource):
    @overload
    def __init__(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 description: Optional[pulumi.Input[str]] = None,
                 ike_version: Optional[pulumi.Input[int]] = None,
                 labels: Optional[pulumi.Input[Mapping[str, str]]] = None,
                 local_traffic_selector: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]] = None,
                 name: Optional[pulumi.Input[str]] = None,
                 peer_external_gateway: Optional[pulumi.Input[str]] = None,
                 peer_external_gateway_interface: Optional[pulumi.Input[int]] = None,
                 peer_gcp_gateway: Optional[pulumi.Input[str]] = None,
                 peer_ip: Optional[pulumi.Input[str]] = None,
                 project: Optional[pulumi.Input[str]] = None,
                 region: Optional[pulumi.Input[str]] = None,
                 remote_traffic_selector: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]] = None,
                 request_id: Optional[pulumi.Input[str]] = None,
                 router: Optional[pulumi.Input[str]] = None,
                 shared_secret: Optional[pulumi.Input[str]] = None,
                 shared_secret_hash: Optional[pulumi.Input[str]] = None,
                 target_vpn_gateway: Optional[pulumi.Input[str]] = None,
                 vpn_gateway: Optional[pulumi.Input[str]] = None,
                 vpn_gateway_interface: Optional[pulumi.Input[int]] = None,
                 __props__=None):
        """
        Creates a VpnTunnel resource in the specified project and region using the data included in the request.

        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] description: An optional description of this resource. Provide this property when you create the resource.
        :param pulumi.Input[int] ike_version: IKE protocol version to use when establishing the VPN tunnel with the peer VPN gateway. Acceptable IKE versions are 1 or 2. The default version is 2.
        :param pulumi.Input[Mapping[str, str]] labels: Labels for this resource. These can only be added or modified by the setLabels method. Each label key/value pair must comply with RFC1035. Label values may be empty.
        :param pulumi.Input[Sequence[pulumi.Input[str]]] local_traffic_selector: Local traffic selector to use when establishing the VPN tunnel with the peer VPN gateway. The value should be a CIDR formatted string, for example: 192.168.0.0/16. The ranges must be disjoint. Only IPv4 is supported.
        :param pulumi.Input[str] name: Name of the resource. Provided by the client when the resource is created. The name must be 1-63 characters long, and comply with RFC1035. Specifically, the name must be 1-63 characters long and match the regular expression `[a-z]([-a-z0-9]*[a-z0-9])?` which means the first character must be a lowercase letter, and all following characters must be a dash, lowercase letter, or digit, except the last character, which cannot be a dash.
        :param pulumi.Input[str] peer_external_gateway: URL of the peer side external VPN gateway to which this VPN tunnel is connected. Provided by the client when the VPN tunnel is created. This field is exclusive with the field peerGcpGateway.
        :param pulumi.Input[int] peer_external_gateway_interface: The interface ID of the external VPN gateway to which this VPN tunnel is connected. Provided by the client when the VPN tunnel is created.
        :param pulumi.Input[str] peer_gcp_gateway: URL of the peer side HA GCP VPN gateway to which this VPN tunnel is connected. Provided by the client when the VPN tunnel is created. This field can be used when creating highly available VPN from VPC network to VPC network, the field is exclusive with the field peerExternalGateway. If provided, the VPN tunnel will automatically use the same vpnGatewayInterface ID in the peer GCP VPN gateway.
        :param pulumi.Input[str] peer_ip: IP address of the peer VPN gateway. Only IPv4 is supported.
        :param pulumi.Input[Sequence[pulumi.Input[str]]] remote_traffic_selector: Remote traffic selectors to use when establishing the VPN tunnel with the peer VPN gateway. The value should be a CIDR formatted string, for example: 192.168.0.0/16. The ranges should be disjoint. Only IPv4 is supported.
        :param pulumi.Input[str] router: URL of the router resource to be used for dynamic routing.
        :param pulumi.Input[str] shared_secret: Shared secret used to set the secure session between the Cloud VPN gateway and the peer VPN gateway.
        :param pulumi.Input[str] shared_secret_hash: Hash of the shared secret.
        :param pulumi.Input[str] target_vpn_gateway: URL of the Target VPN gateway with which this VPN tunnel is associated. Provided by the client when the VPN tunnel is created.
        :param pulumi.Input[str] vpn_gateway: URL of the VPN gateway with which this VPN tunnel is associated. Provided by the client when the VPN tunnel is created. This must be used (instead of target_vpn_gateway) if a High Availability VPN gateway resource is created.
        :param pulumi.Input[int] vpn_gateway_interface: The interface ID of the VPN gateway with which this VPN tunnel is associated.
        """
        ...
    @overload
    def __init__(__self__,
                 resource_name: str,
                 args: VpnTunnelArgs,
                 opts: Optional[pulumi.ResourceOptions] = None):
        """
        Creates a VpnTunnel resource in the specified project and region using the data included in the request.

        :param str resource_name: The name of the resource.
        :param VpnTunnelArgs args: The arguments to use to populate this resource's properties.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        ...
    def __init__(__self__, resource_name: str, *args, **kwargs):
        resource_args, opts = _utilities.get_resource_args_opts(VpnTunnelArgs, pulumi.ResourceOptions, *args, **kwargs)
        if resource_args is not None:
            __self__._internal_init(resource_name, opts, **resource_args.__dict__)
        else:
            __self__._internal_init(resource_name, *args, **kwargs)

    def _internal_init(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 description: Optional[pulumi.Input[str]] = None,
                 ike_version: Optional[pulumi.Input[int]] = None,
                 labels: Optional[pulumi.Input[Mapping[str, str]]] = None,
                 local_traffic_selector: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]] = None,
                 name: Optional[pulumi.Input[str]] = None,
                 peer_external_gateway: Optional[pulumi.Input[str]] = None,
                 peer_external_gateway_interface: Optional[pulumi.Input[int]] = None,
                 peer_gcp_gateway: Optional[pulumi.Input[str]] = None,
                 peer_ip: Optional[pulumi.Input[str]] = None,
                 project: Optional[pulumi.Input[str]] = None,
                 region: Optional[pulumi.Input[str]] = None,
                 remote_traffic_selector: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]] = None,
                 request_id: Optional[pulumi.Input[str]] = None,
                 router: Optional[pulumi.Input[str]] = None,
                 shared_secret: Optional[pulumi.Input[str]] = None,
                 shared_secret_hash: Optional[pulumi.Input[str]] = None,
                 target_vpn_gateway: Optional[pulumi.Input[str]] = None,
                 vpn_gateway: Optional[pulumi.Input[str]] = None,
                 vpn_gateway_interface: Optional[pulumi.Input[int]] = None,
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
            __props__ = VpnTunnelArgs.__new__(VpnTunnelArgs)

            __props__.__dict__["description"] = description
            __props__.__dict__["ike_version"] = ike_version
            __props__.__dict__["labels"] = labels
            __props__.__dict__["local_traffic_selector"] = local_traffic_selector
            __props__.__dict__["name"] = name
            __props__.__dict__["peer_external_gateway"] = peer_external_gateway
            __props__.__dict__["peer_external_gateway_interface"] = peer_external_gateway_interface
            __props__.__dict__["peer_gcp_gateway"] = peer_gcp_gateway
            __props__.__dict__["peer_ip"] = peer_ip
            if project is None and not opts.urn:
                raise TypeError("Missing required property 'project'")
            __props__.__dict__["project"] = project
            if region is None and not opts.urn:
                raise TypeError("Missing required property 'region'")
            __props__.__dict__["region"] = region
            __props__.__dict__["remote_traffic_selector"] = remote_traffic_selector
            __props__.__dict__["request_id"] = request_id
            __props__.__dict__["router"] = router
            __props__.__dict__["shared_secret"] = shared_secret
            __props__.__dict__["shared_secret_hash"] = shared_secret_hash
            __props__.__dict__["target_vpn_gateway"] = target_vpn_gateway
            __props__.__dict__["vpn_gateway"] = vpn_gateway
            __props__.__dict__["vpn_gateway_interface"] = vpn_gateway_interface
            __props__.__dict__["creation_timestamp"] = None
            __props__.__dict__["detailed_status"] = None
            __props__.__dict__["kind"] = None
            __props__.__dict__["label_fingerprint"] = None
            __props__.__dict__["self_link"] = None
            __props__.__dict__["status"] = None
        super(VpnTunnel, __self__).__init__(
            'google-native:compute/beta:VpnTunnel',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name: str,
            id: pulumi.Input[str],
            opts: Optional[pulumi.ResourceOptions] = None) -> 'VpnTunnel':
        """
        Get an existing VpnTunnel resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param pulumi.Input[str] id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = VpnTunnelArgs.__new__(VpnTunnelArgs)

        __props__.__dict__["creation_timestamp"] = None
        __props__.__dict__["description"] = None
        __props__.__dict__["detailed_status"] = None
        __props__.__dict__["ike_version"] = None
        __props__.__dict__["kind"] = None
        __props__.__dict__["label_fingerprint"] = None
        __props__.__dict__["labels"] = None
        __props__.__dict__["local_traffic_selector"] = None
        __props__.__dict__["name"] = None
        __props__.__dict__["peer_external_gateway"] = None
        __props__.__dict__["peer_external_gateway_interface"] = None
        __props__.__dict__["peer_gcp_gateway"] = None
        __props__.__dict__["peer_ip"] = None
        __props__.__dict__["region"] = None
        __props__.__dict__["remote_traffic_selector"] = None
        __props__.__dict__["router"] = None
        __props__.__dict__["self_link"] = None
        __props__.__dict__["shared_secret"] = None
        __props__.__dict__["shared_secret_hash"] = None
        __props__.__dict__["status"] = None
        __props__.__dict__["target_vpn_gateway"] = None
        __props__.__dict__["vpn_gateway"] = None
        __props__.__dict__["vpn_gateway_interface"] = None
        return VpnTunnel(resource_name, opts=opts, __props__=__props__)

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
        An optional description of this resource. Provide this property when you create the resource.
        """
        return pulumi.get(self, "description")

    @property
    @pulumi.getter(name="detailedStatus")
    def detailed_status(self) -> pulumi.Output[str]:
        """
        Detailed status message for the VPN tunnel.
        """
        return pulumi.get(self, "detailed_status")

    @property
    @pulumi.getter(name="ikeVersion")
    def ike_version(self) -> pulumi.Output[int]:
        """
        IKE protocol version to use when establishing the VPN tunnel with the peer VPN gateway. Acceptable IKE versions are 1 or 2. The default version is 2.
        """
        return pulumi.get(self, "ike_version")

    @property
    @pulumi.getter
    def kind(self) -> pulumi.Output[str]:
        """
        Type of resource. Always compute#vpnTunnel for VPN tunnels.
        """
        return pulumi.get(self, "kind")

    @property
    @pulumi.getter(name="labelFingerprint")
    def label_fingerprint(self) -> pulumi.Output[str]:
        """
        A fingerprint for the labels being applied to this VpnTunnel, which is essentially a hash of the labels set used for optimistic locking. The fingerprint is initially generated by Compute Engine and changes after every request to modify or update labels. You must always provide an up-to-date fingerprint hash in order to update or change labels, otherwise the request will fail with error 412 conditionNotMet.

        To see the latest fingerprint, make a get() request to retrieve a VpnTunnel.
        """
        return pulumi.get(self, "label_fingerprint")

    @property
    @pulumi.getter
    def labels(self) -> pulumi.Output[Mapping[str, str]]:
        """
        Labels for this resource. These can only be added or modified by the setLabels method. Each label key/value pair must comply with RFC1035. Label values may be empty.
        """
        return pulumi.get(self, "labels")

    @property
    @pulumi.getter(name="localTrafficSelector")
    def local_traffic_selector(self) -> pulumi.Output[Sequence[str]]:
        """
        Local traffic selector to use when establishing the VPN tunnel with the peer VPN gateway. The value should be a CIDR formatted string, for example: 192.168.0.0/16. The ranges must be disjoint. Only IPv4 is supported.
        """
        return pulumi.get(self, "local_traffic_selector")

    @property
    @pulumi.getter
    def name(self) -> pulumi.Output[str]:
        """
        Name of the resource. Provided by the client when the resource is created. The name must be 1-63 characters long, and comply with RFC1035. Specifically, the name must be 1-63 characters long and match the regular expression `[a-z]([-a-z0-9]*[a-z0-9])?` which means the first character must be a lowercase letter, and all following characters must be a dash, lowercase letter, or digit, except the last character, which cannot be a dash.
        """
        return pulumi.get(self, "name")

    @property
    @pulumi.getter(name="peerExternalGateway")
    def peer_external_gateway(self) -> pulumi.Output[str]:
        """
        URL of the peer side external VPN gateway to which this VPN tunnel is connected. Provided by the client when the VPN tunnel is created. This field is exclusive with the field peerGcpGateway.
        """
        return pulumi.get(self, "peer_external_gateway")

    @property
    @pulumi.getter(name="peerExternalGatewayInterface")
    def peer_external_gateway_interface(self) -> pulumi.Output[int]:
        """
        The interface ID of the external VPN gateway to which this VPN tunnel is connected. Provided by the client when the VPN tunnel is created.
        """
        return pulumi.get(self, "peer_external_gateway_interface")

    @property
    @pulumi.getter(name="peerGcpGateway")
    def peer_gcp_gateway(self) -> pulumi.Output[str]:
        """
        URL of the peer side HA GCP VPN gateway to which this VPN tunnel is connected. Provided by the client when the VPN tunnel is created. This field can be used when creating highly available VPN from VPC network to VPC network, the field is exclusive with the field peerExternalGateway. If provided, the VPN tunnel will automatically use the same vpnGatewayInterface ID in the peer GCP VPN gateway.
        """
        return pulumi.get(self, "peer_gcp_gateway")

    @property
    @pulumi.getter(name="peerIp")
    def peer_ip(self) -> pulumi.Output[str]:
        """
        IP address of the peer VPN gateway. Only IPv4 is supported.
        """
        return pulumi.get(self, "peer_ip")

    @property
    @pulumi.getter
    def region(self) -> pulumi.Output[str]:
        """
        URL of the region where the VPN tunnel resides. You must specify this field as part of the HTTP request URL. It is not settable as a field in the request body.
        """
        return pulumi.get(self, "region")

    @property
    @pulumi.getter(name="remoteTrafficSelector")
    def remote_traffic_selector(self) -> pulumi.Output[Sequence[str]]:
        """
        Remote traffic selectors to use when establishing the VPN tunnel with the peer VPN gateway. The value should be a CIDR formatted string, for example: 192.168.0.0/16. The ranges should be disjoint. Only IPv4 is supported.
        """
        return pulumi.get(self, "remote_traffic_selector")

    @property
    @pulumi.getter
    def router(self) -> pulumi.Output[str]:
        """
        URL of the router resource to be used for dynamic routing.
        """
        return pulumi.get(self, "router")

    @property
    @pulumi.getter(name="selfLink")
    def self_link(self) -> pulumi.Output[str]:
        """
        Server-defined URL for the resource.
        """
        return pulumi.get(self, "self_link")

    @property
    @pulumi.getter(name="sharedSecret")
    def shared_secret(self) -> pulumi.Output[str]:
        """
        Shared secret used to set the secure session between the Cloud VPN gateway and the peer VPN gateway.
        """
        return pulumi.get(self, "shared_secret")

    @property
    @pulumi.getter(name="sharedSecretHash")
    def shared_secret_hash(self) -> pulumi.Output[str]:
        """
        Hash of the shared secret.
        """
        return pulumi.get(self, "shared_secret_hash")

    @property
    @pulumi.getter
    def status(self) -> pulumi.Output[str]:
        """
        The status of the VPN tunnel, which can be one of the following: 
        - PROVISIONING: Resource is being allocated for the VPN tunnel. 
        - WAITING_FOR_FULL_CONFIG: Waiting to receive all VPN-related configs from the user. Network, TargetVpnGateway, VpnTunnel, ForwardingRule, and Route resources are needed to setup the VPN tunnel. 
        - FIRST_HANDSHAKE: Successful first handshake with the peer VPN. 
        - ESTABLISHED: Secure session is successfully established with the peer VPN. 
        - NETWORK_ERROR: Deprecated, replaced by NO_INCOMING_PACKETS 
        - AUTHORIZATION_ERROR: Auth error (for example, bad shared secret). 
        - NEGOTIATION_FAILURE: Handshake failed. 
        - DEPROVISIONING: Resources are being deallocated for the VPN tunnel. 
        - FAILED: Tunnel creation has failed and the tunnel is not ready to be used. 
        - NO_INCOMING_PACKETS: No incoming packets from peer. 
        - REJECTED: Tunnel configuration was rejected, can be result of being denied access. 
        - ALLOCATING_RESOURCES: Cloud VPN is in the process of allocating all required resources. 
        - STOPPED: Tunnel is stopped due to its Forwarding Rules being deleted for Classic VPN tunnels or the project is in frozen state. 
        - PEER_IDENTITY_MISMATCH: Peer identity does not match peer IP, probably behind NAT. 
        - TS_NARROWING_NOT_ALLOWED: Traffic selector narrowing not allowed for an HA-VPN tunnel.
        """
        return pulumi.get(self, "status")

    @property
    @pulumi.getter(name="targetVpnGateway")
    def target_vpn_gateway(self) -> pulumi.Output[str]:
        """
        URL of the Target VPN gateway with which this VPN tunnel is associated. Provided by the client when the VPN tunnel is created.
        """
        return pulumi.get(self, "target_vpn_gateway")

    @property
    @pulumi.getter(name="vpnGateway")
    def vpn_gateway(self) -> pulumi.Output[str]:
        """
        URL of the VPN gateway with which this VPN tunnel is associated. Provided by the client when the VPN tunnel is created. This must be used (instead of target_vpn_gateway) if a High Availability VPN gateway resource is created.
        """
        return pulumi.get(self, "vpn_gateway")

    @property
    @pulumi.getter(name="vpnGatewayInterface")
    def vpn_gateway_interface(self) -> pulumi.Output[int]:
        """
        The interface ID of the VPN gateway with which this VPN tunnel is associated.
        """
        return pulumi.get(self, "vpn_gateway_interface")

