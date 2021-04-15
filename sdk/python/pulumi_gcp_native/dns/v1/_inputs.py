# coding=utf-8
# *** WARNING: this file was generated by the Pulumi SDK Generator. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import warnings
import pulumi
import pulumi.runtime
from typing import Any, Mapping, Optional, Sequence, Union, overload
from ... import _utilities

__all__ = [
    'DnsKeySpecArgs',
    'ManagedZoneDnsSecConfigArgs',
    'ManagedZoneForwardingConfigArgs',
    'ManagedZoneForwardingConfigNameServerTargetArgs',
    'ManagedZonePeeringConfigArgs',
    'ManagedZonePeeringConfigTargetNetworkArgs',
    'ManagedZonePrivateVisibilityConfigArgs',
    'ManagedZonePrivateVisibilityConfigNetworkArgs',
    'ManagedZoneReverseLookupConfigArgs',
    'ManagedZoneServiceDirectoryConfigArgs',
    'ManagedZoneServiceDirectoryConfigNamespaceArgs',
    'PolicyAlternativeNameServerConfigArgs',
    'PolicyAlternativeNameServerConfigTargetNameServerArgs',
    'PolicyNetworkArgs',
    'ResourceRecordSetArgs',
]

@pulumi.input_type
class DnsKeySpecArgs:
    def __init__(__self__, *,
                 algorithm: Optional[pulumi.Input[str]] = None,
                 key_length: Optional[pulumi.Input[int]] = None,
                 key_type: Optional[pulumi.Input[str]] = None,
                 kind: Optional[pulumi.Input[str]] = None):
        """
        Parameters for DnsKey key generation. Used for generating initial keys for a new ManagedZone and as default when adding a new DnsKey.
        :param pulumi.Input[str] algorithm: String mnemonic specifying the DNSSEC algorithm of this key.
        :param pulumi.Input[int] key_length: Length of the keys in bits.
        :param pulumi.Input[str] key_type: Specifies whether this is a key signing key (KSK) or a zone signing key (ZSK). Key signing keys have the Secure Entry Point flag set and, when active, are only used to sign resource record sets of type DNSKEY. Zone signing keys do not have the Secure Entry Point flag set and are used to sign all other types of resource record sets.
        """
        if algorithm is not None:
            pulumi.set(__self__, "algorithm", algorithm)
        if key_length is not None:
            pulumi.set(__self__, "key_length", key_length)
        if key_type is not None:
            pulumi.set(__self__, "key_type", key_type)
        if kind is not None:
            pulumi.set(__self__, "kind", kind)

    @property
    @pulumi.getter
    def algorithm(self) -> Optional[pulumi.Input[str]]:
        """
        String mnemonic specifying the DNSSEC algorithm of this key.
        """
        return pulumi.get(self, "algorithm")

    @algorithm.setter
    def algorithm(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "algorithm", value)

    @property
    @pulumi.getter(name="keyLength")
    def key_length(self) -> Optional[pulumi.Input[int]]:
        """
        Length of the keys in bits.
        """
        return pulumi.get(self, "key_length")

    @key_length.setter
    def key_length(self, value: Optional[pulumi.Input[int]]):
        pulumi.set(self, "key_length", value)

    @property
    @pulumi.getter(name="keyType")
    def key_type(self) -> Optional[pulumi.Input[str]]:
        """
        Specifies whether this is a key signing key (KSK) or a zone signing key (ZSK). Key signing keys have the Secure Entry Point flag set and, when active, are only used to sign resource record sets of type DNSKEY. Zone signing keys do not have the Secure Entry Point flag set and are used to sign all other types of resource record sets.
        """
        return pulumi.get(self, "key_type")

    @key_type.setter
    def key_type(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "key_type", value)

    @property
    @pulumi.getter
    def kind(self) -> Optional[pulumi.Input[str]]:
        return pulumi.get(self, "kind")

    @kind.setter
    def kind(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "kind", value)


@pulumi.input_type
class ManagedZoneDnsSecConfigArgs:
    def __init__(__self__, *,
                 default_key_specs: Optional[pulumi.Input[Sequence[pulumi.Input['DnsKeySpecArgs']]]] = None,
                 kind: Optional[pulumi.Input[str]] = None,
                 non_existence: Optional[pulumi.Input[str]] = None,
                 state: Optional[pulumi.Input[str]] = None):
        """
        :param pulumi.Input[Sequence[pulumi.Input['DnsKeySpecArgs']]] default_key_specs: Specifies parameters for generating initial DnsKeys for this ManagedZone. Can only be changed while the state is OFF.
        :param pulumi.Input[str] non_existence: Specifies the mechanism for authenticated denial-of-existence responses. Can only be changed while the state is OFF.
        :param pulumi.Input[str] state: Specifies whether DNSSEC is enabled, and what mode it is in.
        """
        if default_key_specs is not None:
            pulumi.set(__self__, "default_key_specs", default_key_specs)
        if kind is not None:
            pulumi.set(__self__, "kind", kind)
        if non_existence is not None:
            pulumi.set(__self__, "non_existence", non_existence)
        if state is not None:
            pulumi.set(__self__, "state", state)

    @property
    @pulumi.getter(name="defaultKeySpecs")
    def default_key_specs(self) -> Optional[pulumi.Input[Sequence[pulumi.Input['DnsKeySpecArgs']]]]:
        """
        Specifies parameters for generating initial DnsKeys for this ManagedZone. Can only be changed while the state is OFF.
        """
        return pulumi.get(self, "default_key_specs")

    @default_key_specs.setter
    def default_key_specs(self, value: Optional[pulumi.Input[Sequence[pulumi.Input['DnsKeySpecArgs']]]]):
        pulumi.set(self, "default_key_specs", value)

    @property
    @pulumi.getter
    def kind(self) -> Optional[pulumi.Input[str]]:
        return pulumi.get(self, "kind")

    @kind.setter
    def kind(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "kind", value)

    @property
    @pulumi.getter(name="nonExistence")
    def non_existence(self) -> Optional[pulumi.Input[str]]:
        """
        Specifies the mechanism for authenticated denial-of-existence responses. Can only be changed while the state is OFF.
        """
        return pulumi.get(self, "non_existence")

    @non_existence.setter
    def non_existence(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "non_existence", value)

    @property
    @pulumi.getter
    def state(self) -> Optional[pulumi.Input[str]]:
        """
        Specifies whether DNSSEC is enabled, and what mode it is in.
        """
        return pulumi.get(self, "state")

    @state.setter
    def state(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "state", value)


@pulumi.input_type
class ManagedZoneForwardingConfigArgs:
    def __init__(__self__, *,
                 kind: Optional[pulumi.Input[str]] = None,
                 target_name_servers: Optional[pulumi.Input[Sequence[pulumi.Input['ManagedZoneForwardingConfigNameServerTargetArgs']]]] = None):
        """
        :param pulumi.Input[Sequence[pulumi.Input['ManagedZoneForwardingConfigNameServerTargetArgs']]] target_name_servers: List of target name servers to forward to. Cloud DNS selects the best available name server if more than one target is given.
        """
        if kind is not None:
            pulumi.set(__self__, "kind", kind)
        if target_name_servers is not None:
            pulumi.set(__self__, "target_name_servers", target_name_servers)

    @property
    @pulumi.getter
    def kind(self) -> Optional[pulumi.Input[str]]:
        return pulumi.get(self, "kind")

    @kind.setter
    def kind(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "kind", value)

    @property
    @pulumi.getter(name="targetNameServers")
    def target_name_servers(self) -> Optional[pulumi.Input[Sequence[pulumi.Input['ManagedZoneForwardingConfigNameServerTargetArgs']]]]:
        """
        List of target name servers to forward to. Cloud DNS selects the best available name server if more than one target is given.
        """
        return pulumi.get(self, "target_name_servers")

    @target_name_servers.setter
    def target_name_servers(self, value: Optional[pulumi.Input[Sequence[pulumi.Input['ManagedZoneForwardingConfigNameServerTargetArgs']]]]):
        pulumi.set(self, "target_name_servers", value)


@pulumi.input_type
class ManagedZoneForwardingConfigNameServerTargetArgs:
    def __init__(__self__, *,
                 forwarding_path: Optional[pulumi.Input[str]] = None,
                 ipv4_address: Optional[pulumi.Input[str]] = None,
                 kind: Optional[pulumi.Input[str]] = None):
        """
        :param pulumi.Input[str] forwarding_path: Forwarding path for this NameServerTarget. If unset or set to DEFAULT, Cloud DNS makes forwarding decisions based on IP address ranges; that is, RFC1918 addresses go to the VPC network, non-RFC1918 addresses go to the internet. When set to PRIVATE, Cloud DNS always sends queries through the VPC network for this target.
        :param pulumi.Input[str] ipv4_address: IPv4 address of a target name server.
        """
        if forwarding_path is not None:
            pulumi.set(__self__, "forwarding_path", forwarding_path)
        if ipv4_address is not None:
            pulumi.set(__self__, "ipv4_address", ipv4_address)
        if kind is not None:
            pulumi.set(__self__, "kind", kind)

    @property
    @pulumi.getter(name="forwardingPath")
    def forwarding_path(self) -> Optional[pulumi.Input[str]]:
        """
        Forwarding path for this NameServerTarget. If unset or set to DEFAULT, Cloud DNS makes forwarding decisions based on IP address ranges; that is, RFC1918 addresses go to the VPC network, non-RFC1918 addresses go to the internet. When set to PRIVATE, Cloud DNS always sends queries through the VPC network for this target.
        """
        return pulumi.get(self, "forwarding_path")

    @forwarding_path.setter
    def forwarding_path(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "forwarding_path", value)

    @property
    @pulumi.getter(name="ipv4Address")
    def ipv4_address(self) -> Optional[pulumi.Input[str]]:
        """
        IPv4 address of a target name server.
        """
        return pulumi.get(self, "ipv4_address")

    @ipv4_address.setter
    def ipv4_address(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "ipv4_address", value)

    @property
    @pulumi.getter
    def kind(self) -> Optional[pulumi.Input[str]]:
        return pulumi.get(self, "kind")

    @kind.setter
    def kind(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "kind", value)


@pulumi.input_type
class ManagedZonePeeringConfigArgs:
    def __init__(__self__, *,
                 kind: Optional[pulumi.Input[str]] = None,
                 target_network: Optional[pulumi.Input['ManagedZonePeeringConfigTargetNetworkArgs']] = None):
        """
        :param pulumi.Input['ManagedZonePeeringConfigTargetNetworkArgs'] target_network: The network with which to peer.
        """
        if kind is not None:
            pulumi.set(__self__, "kind", kind)
        if target_network is not None:
            pulumi.set(__self__, "target_network", target_network)

    @property
    @pulumi.getter
    def kind(self) -> Optional[pulumi.Input[str]]:
        return pulumi.get(self, "kind")

    @kind.setter
    def kind(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "kind", value)

    @property
    @pulumi.getter(name="targetNetwork")
    def target_network(self) -> Optional[pulumi.Input['ManagedZonePeeringConfigTargetNetworkArgs']]:
        """
        The network with which to peer.
        """
        return pulumi.get(self, "target_network")

    @target_network.setter
    def target_network(self, value: Optional[pulumi.Input['ManagedZonePeeringConfigTargetNetworkArgs']]):
        pulumi.set(self, "target_network", value)


@pulumi.input_type
class ManagedZonePeeringConfigTargetNetworkArgs:
    def __init__(__self__, *,
                 deactivate_time: Optional[pulumi.Input[str]] = None,
                 kind: Optional[pulumi.Input[str]] = None,
                 network_url: Optional[pulumi.Input[str]] = None):
        """
        :param pulumi.Input[str] deactivate_time: The time at which the zone was deactivated, in RFC 3339 date-time format. An empty string indicates that the peering connection is active. The producer network can deactivate a zone. The zone is automatically deactivated if the producer network that the zone targeted is deleted. Output only.
        :param pulumi.Input[str] network_url: The fully qualified URL of the VPC network to forward queries to. This should be formatted like https://www.googleapis.com/compute/v1/projects/{project}/global/networks/{network}
        """
        if deactivate_time is not None:
            pulumi.set(__self__, "deactivate_time", deactivate_time)
        if kind is not None:
            pulumi.set(__self__, "kind", kind)
        if network_url is not None:
            pulumi.set(__self__, "network_url", network_url)

    @property
    @pulumi.getter(name="deactivateTime")
    def deactivate_time(self) -> Optional[pulumi.Input[str]]:
        """
        The time at which the zone was deactivated, in RFC 3339 date-time format. An empty string indicates that the peering connection is active. The producer network can deactivate a zone. The zone is automatically deactivated if the producer network that the zone targeted is deleted. Output only.
        """
        return pulumi.get(self, "deactivate_time")

    @deactivate_time.setter
    def deactivate_time(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "deactivate_time", value)

    @property
    @pulumi.getter
    def kind(self) -> Optional[pulumi.Input[str]]:
        return pulumi.get(self, "kind")

    @kind.setter
    def kind(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "kind", value)

    @property
    @pulumi.getter(name="networkUrl")
    def network_url(self) -> Optional[pulumi.Input[str]]:
        """
        The fully qualified URL of the VPC network to forward queries to. This should be formatted like https://www.googleapis.com/compute/v1/projects/{project}/global/networks/{network}
        """
        return pulumi.get(self, "network_url")

    @network_url.setter
    def network_url(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "network_url", value)


@pulumi.input_type
class ManagedZonePrivateVisibilityConfigArgs:
    def __init__(__self__, *,
                 kind: Optional[pulumi.Input[str]] = None,
                 networks: Optional[pulumi.Input[Sequence[pulumi.Input['ManagedZonePrivateVisibilityConfigNetworkArgs']]]] = None):
        """
        :param pulumi.Input[Sequence[pulumi.Input['ManagedZonePrivateVisibilityConfigNetworkArgs']]] networks: The list of VPC networks that can see this zone.
        """
        if kind is not None:
            pulumi.set(__self__, "kind", kind)
        if networks is not None:
            pulumi.set(__self__, "networks", networks)

    @property
    @pulumi.getter
    def kind(self) -> Optional[pulumi.Input[str]]:
        return pulumi.get(self, "kind")

    @kind.setter
    def kind(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "kind", value)

    @property
    @pulumi.getter
    def networks(self) -> Optional[pulumi.Input[Sequence[pulumi.Input['ManagedZonePrivateVisibilityConfigNetworkArgs']]]]:
        """
        The list of VPC networks that can see this zone.
        """
        return pulumi.get(self, "networks")

    @networks.setter
    def networks(self, value: Optional[pulumi.Input[Sequence[pulumi.Input['ManagedZonePrivateVisibilityConfigNetworkArgs']]]]):
        pulumi.set(self, "networks", value)


@pulumi.input_type
class ManagedZonePrivateVisibilityConfigNetworkArgs:
    def __init__(__self__, *,
                 kind: Optional[pulumi.Input[str]] = None,
                 network_url: Optional[pulumi.Input[str]] = None):
        """
        :param pulumi.Input[str] network_url: The fully qualified URL of the VPC network to bind to. Format this URL like https://www.googleapis.com/compute/v1/projects/{project}/global/networks/{network}
        """
        if kind is not None:
            pulumi.set(__self__, "kind", kind)
        if network_url is not None:
            pulumi.set(__self__, "network_url", network_url)

    @property
    @pulumi.getter
    def kind(self) -> Optional[pulumi.Input[str]]:
        return pulumi.get(self, "kind")

    @kind.setter
    def kind(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "kind", value)

    @property
    @pulumi.getter(name="networkUrl")
    def network_url(self) -> Optional[pulumi.Input[str]]:
        """
        The fully qualified URL of the VPC network to bind to. Format this URL like https://www.googleapis.com/compute/v1/projects/{project}/global/networks/{network}
        """
        return pulumi.get(self, "network_url")

    @network_url.setter
    def network_url(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "network_url", value)


@pulumi.input_type
class ManagedZoneReverseLookupConfigArgs:
    def __init__(__self__, *,
                 kind: Optional[pulumi.Input[str]] = None):
        if kind is not None:
            pulumi.set(__self__, "kind", kind)

    @property
    @pulumi.getter
    def kind(self) -> Optional[pulumi.Input[str]]:
        return pulumi.get(self, "kind")

    @kind.setter
    def kind(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "kind", value)


@pulumi.input_type
class ManagedZoneServiceDirectoryConfigArgs:
    def __init__(__self__, *,
                 kind: Optional[pulumi.Input[str]] = None,
                 namespace: Optional[pulumi.Input['ManagedZoneServiceDirectoryConfigNamespaceArgs']] = None):
        """
        Contains information about Service Directory-backed zones.
        :param pulumi.Input['ManagedZoneServiceDirectoryConfigNamespaceArgs'] namespace: Contains information about the namespace associated with the zone.
        """
        if kind is not None:
            pulumi.set(__self__, "kind", kind)
        if namespace is not None:
            pulumi.set(__self__, "namespace", namespace)

    @property
    @pulumi.getter
    def kind(self) -> Optional[pulumi.Input[str]]:
        return pulumi.get(self, "kind")

    @kind.setter
    def kind(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "kind", value)

    @property
    @pulumi.getter
    def namespace(self) -> Optional[pulumi.Input['ManagedZoneServiceDirectoryConfigNamespaceArgs']]:
        """
        Contains information about the namespace associated with the zone.
        """
        return pulumi.get(self, "namespace")

    @namespace.setter
    def namespace(self, value: Optional[pulumi.Input['ManagedZoneServiceDirectoryConfigNamespaceArgs']]):
        pulumi.set(self, "namespace", value)


@pulumi.input_type
class ManagedZoneServiceDirectoryConfigNamespaceArgs:
    def __init__(__self__, *,
                 deletion_time: Optional[pulumi.Input[str]] = None,
                 kind: Optional[pulumi.Input[str]] = None,
                 namespace_url: Optional[pulumi.Input[str]] = None):
        """
        :param pulumi.Input[str] deletion_time: The time that the namespace backing this zone was deleted; an empty string if it still exists. This is in RFC3339 text format. Output only.
        :param pulumi.Input[str] namespace_url: The fully qualified URL of the namespace associated with the zone. Format must be https://servicedirectory.googleapis.com/v1/projects/{project}/locations/{location}/namespaces/{namespace}
        """
        if deletion_time is not None:
            pulumi.set(__self__, "deletion_time", deletion_time)
        if kind is not None:
            pulumi.set(__self__, "kind", kind)
        if namespace_url is not None:
            pulumi.set(__self__, "namespace_url", namespace_url)

    @property
    @pulumi.getter(name="deletionTime")
    def deletion_time(self) -> Optional[pulumi.Input[str]]:
        """
        The time that the namespace backing this zone was deleted; an empty string if it still exists. This is in RFC3339 text format. Output only.
        """
        return pulumi.get(self, "deletion_time")

    @deletion_time.setter
    def deletion_time(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "deletion_time", value)

    @property
    @pulumi.getter
    def kind(self) -> Optional[pulumi.Input[str]]:
        return pulumi.get(self, "kind")

    @kind.setter
    def kind(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "kind", value)

    @property
    @pulumi.getter(name="namespaceUrl")
    def namespace_url(self) -> Optional[pulumi.Input[str]]:
        """
        The fully qualified URL of the namespace associated with the zone. Format must be https://servicedirectory.googleapis.com/v1/projects/{project}/locations/{location}/namespaces/{namespace}
        """
        return pulumi.get(self, "namespace_url")

    @namespace_url.setter
    def namespace_url(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "namespace_url", value)


@pulumi.input_type
class PolicyAlternativeNameServerConfigArgs:
    def __init__(__self__, *,
                 kind: Optional[pulumi.Input[str]] = None,
                 target_name_servers: Optional[pulumi.Input[Sequence[pulumi.Input['PolicyAlternativeNameServerConfigTargetNameServerArgs']]]] = None):
        """
        :param pulumi.Input[Sequence[pulumi.Input['PolicyAlternativeNameServerConfigTargetNameServerArgs']]] target_name_servers: Sets an alternative name server for the associated networks. When specified, all DNS queries are forwarded to a name server that you choose. Names such as .internal are not available when an alternative name server is specified.
        """
        if kind is not None:
            pulumi.set(__self__, "kind", kind)
        if target_name_servers is not None:
            pulumi.set(__self__, "target_name_servers", target_name_servers)

    @property
    @pulumi.getter
    def kind(self) -> Optional[pulumi.Input[str]]:
        return pulumi.get(self, "kind")

    @kind.setter
    def kind(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "kind", value)

    @property
    @pulumi.getter(name="targetNameServers")
    def target_name_servers(self) -> Optional[pulumi.Input[Sequence[pulumi.Input['PolicyAlternativeNameServerConfigTargetNameServerArgs']]]]:
        """
        Sets an alternative name server for the associated networks. When specified, all DNS queries are forwarded to a name server that you choose. Names such as .internal are not available when an alternative name server is specified.
        """
        return pulumi.get(self, "target_name_servers")

    @target_name_servers.setter
    def target_name_servers(self, value: Optional[pulumi.Input[Sequence[pulumi.Input['PolicyAlternativeNameServerConfigTargetNameServerArgs']]]]):
        pulumi.set(self, "target_name_servers", value)


@pulumi.input_type
class PolicyAlternativeNameServerConfigTargetNameServerArgs:
    def __init__(__self__, *,
                 forwarding_path: Optional[pulumi.Input[str]] = None,
                 ipv4_address: Optional[pulumi.Input[str]] = None,
                 kind: Optional[pulumi.Input[str]] = None):
        """
        :param pulumi.Input[str] forwarding_path: Forwarding path for this TargetNameServer. If unset or set to DEFAULT, Cloud DNS makes forwarding decisions based on address ranges; that is, RFC1918 addresses go to the VPC network, non-RFC1918 addresses go to the internet. When set to PRIVATE, Cloud DNS always sends queries through the VPC network for this target.
        :param pulumi.Input[str] ipv4_address: IPv4 address to forward to.
        """
        if forwarding_path is not None:
            pulumi.set(__self__, "forwarding_path", forwarding_path)
        if ipv4_address is not None:
            pulumi.set(__self__, "ipv4_address", ipv4_address)
        if kind is not None:
            pulumi.set(__self__, "kind", kind)

    @property
    @pulumi.getter(name="forwardingPath")
    def forwarding_path(self) -> Optional[pulumi.Input[str]]:
        """
        Forwarding path for this TargetNameServer. If unset or set to DEFAULT, Cloud DNS makes forwarding decisions based on address ranges; that is, RFC1918 addresses go to the VPC network, non-RFC1918 addresses go to the internet. When set to PRIVATE, Cloud DNS always sends queries through the VPC network for this target.
        """
        return pulumi.get(self, "forwarding_path")

    @forwarding_path.setter
    def forwarding_path(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "forwarding_path", value)

    @property
    @pulumi.getter(name="ipv4Address")
    def ipv4_address(self) -> Optional[pulumi.Input[str]]:
        """
        IPv4 address to forward to.
        """
        return pulumi.get(self, "ipv4_address")

    @ipv4_address.setter
    def ipv4_address(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "ipv4_address", value)

    @property
    @pulumi.getter
    def kind(self) -> Optional[pulumi.Input[str]]:
        return pulumi.get(self, "kind")

    @kind.setter
    def kind(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "kind", value)


@pulumi.input_type
class PolicyNetworkArgs:
    def __init__(__self__, *,
                 kind: Optional[pulumi.Input[str]] = None,
                 network_url: Optional[pulumi.Input[str]] = None):
        """
        :param pulumi.Input[str] network_url: The fully qualified URL of the VPC network to bind to. This should be formatted like https://www.googleapis.com/compute/v1/projects/{project}/global/networks/{network}
        """
        if kind is not None:
            pulumi.set(__self__, "kind", kind)
        if network_url is not None:
            pulumi.set(__self__, "network_url", network_url)

    @property
    @pulumi.getter
    def kind(self) -> Optional[pulumi.Input[str]]:
        return pulumi.get(self, "kind")

    @kind.setter
    def kind(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "kind", value)

    @property
    @pulumi.getter(name="networkUrl")
    def network_url(self) -> Optional[pulumi.Input[str]]:
        """
        The fully qualified URL of the VPC network to bind to. This should be formatted like https://www.googleapis.com/compute/v1/projects/{project}/global/networks/{network}
        """
        return pulumi.get(self, "network_url")

    @network_url.setter
    def network_url(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "network_url", value)


@pulumi.input_type
class ResourceRecordSetArgs:
    def __init__(__self__, *,
                 kind: Optional[pulumi.Input[str]] = None,
                 name: Optional[pulumi.Input[str]] = None,
                 rrdatas: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]] = None,
                 signature_rrdatas: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]] = None,
                 ttl: Optional[pulumi.Input[int]] = None,
                 type: Optional[pulumi.Input[str]] = None):
        """
        A unit of data that is returned by the DNS servers.
        :param pulumi.Input[str] name: For example, www.example.com.
        :param pulumi.Input[Sequence[pulumi.Input[str]]] rrdatas: As defined in RFC 1035 (section 5) and RFC 1034 (section 3.6.1) -- see examples.
        :param pulumi.Input[Sequence[pulumi.Input[str]]] signature_rrdatas: As defined in RFC 4034 (section 3.2).
        :param pulumi.Input[int] ttl: Number of seconds that this ResourceRecordSet can be cached by resolvers.
        :param pulumi.Input[str] type: The identifier of a supported record type. See the list of Supported DNS record types.
        """
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


