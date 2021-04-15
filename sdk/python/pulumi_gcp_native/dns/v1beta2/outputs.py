# coding=utf-8
# *** WARNING: this file was generated by the Pulumi SDK Generator. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import warnings
import pulumi
import pulumi.runtime
from typing import Any, Mapping, Optional, Sequence, Union, overload
from ... import _utilities
from . import outputs

__all__ = [
    'DnsKeySpecResponse',
    'ManagedZoneDnsSecConfigResponse',
    'ManagedZoneForwardingConfigNameServerTargetResponse',
    'ManagedZoneForwardingConfigResponse',
    'ManagedZonePeeringConfigResponse',
    'ManagedZonePeeringConfigTargetNetworkResponse',
    'ManagedZonePrivateVisibilityConfigNetworkResponse',
    'ManagedZonePrivateVisibilityConfigResponse',
    'ManagedZoneReverseLookupConfigResponse',
    'ManagedZoneServiceDirectoryConfigNamespaceResponse',
    'ManagedZoneServiceDirectoryConfigResponse',
    'PolicyAlternativeNameServerConfigResponse',
    'PolicyAlternativeNameServerConfigTargetNameServerResponse',
    'PolicyNetworkResponse',
    'ResourceRecordSetResponse',
    'ResponsePolicyNetworkResponse',
    'ResponsePolicyRuleLocalDataResponse',
]

@pulumi.output_type
class DnsKeySpecResponse(dict):
    """
    Parameters for DnsKey key generation. Used for generating initial keys for a new ManagedZone and as default when adding a new DnsKey.
    """
    @staticmethod
    def __key_warning(key: str):
        suggest = None
        if key == "keyLength":
            suggest = "key_length"
        elif key == "keyType":
            suggest = "key_type"

        if suggest:
            pulumi.log.warn(f"Key '{key}' not found in DnsKeySpecResponse. Access the value via the '{suggest}' property getter instead.")

    def __getitem__(self, key: str) -> Any:
        DnsKeySpecResponse.__key_warning(key)
        return super().__getitem__(key)

    def get(self, key: str, default = None) -> Any:
        DnsKeySpecResponse.__key_warning(key)
        return super().get(key, default)

    def __init__(__self__, *,
                 algorithm: str,
                 key_length: int,
                 key_type: str,
                 kind: str):
        """
        Parameters for DnsKey key generation. Used for generating initial keys for a new ManagedZone and as default when adding a new DnsKey.
        :param str algorithm: String mnemonic specifying the DNSSEC algorithm of this key.
        :param int key_length: Length of the keys in bits.
        :param str key_type: Specifies whether this is a key signing key (KSK) or a zone signing key (ZSK). Key signing keys have the Secure Entry Point flag set and, when active, are only used to sign resource record sets of type DNSKEY. Zone signing keys do not have the Secure Entry Point flag set and are used to sign all other types of resource record sets.
        """
        pulumi.set(__self__, "algorithm", algorithm)
        pulumi.set(__self__, "key_length", key_length)
        pulumi.set(__self__, "key_type", key_type)
        pulumi.set(__self__, "kind", kind)

    @property
    @pulumi.getter
    def algorithm(self) -> str:
        """
        String mnemonic specifying the DNSSEC algorithm of this key.
        """
        return pulumi.get(self, "algorithm")

    @property
    @pulumi.getter(name="keyLength")
    def key_length(self) -> int:
        """
        Length of the keys in bits.
        """
        return pulumi.get(self, "key_length")

    @property
    @pulumi.getter(name="keyType")
    def key_type(self) -> str:
        """
        Specifies whether this is a key signing key (KSK) or a zone signing key (ZSK). Key signing keys have the Secure Entry Point flag set and, when active, are only used to sign resource record sets of type DNSKEY. Zone signing keys do not have the Secure Entry Point flag set and are used to sign all other types of resource record sets.
        """
        return pulumi.get(self, "key_type")

    @property
    @pulumi.getter
    def kind(self) -> str:
        return pulumi.get(self, "kind")


@pulumi.output_type
class ManagedZoneDnsSecConfigResponse(dict):
    @staticmethod
    def __key_warning(key: str):
        suggest = None
        if key == "defaultKeySpecs":
            suggest = "default_key_specs"
        elif key == "nonExistence":
            suggest = "non_existence"

        if suggest:
            pulumi.log.warn(f"Key '{key}' not found in ManagedZoneDnsSecConfigResponse. Access the value via the '{suggest}' property getter instead.")

    def __getitem__(self, key: str) -> Any:
        ManagedZoneDnsSecConfigResponse.__key_warning(key)
        return super().__getitem__(key)

    def get(self, key: str, default = None) -> Any:
        ManagedZoneDnsSecConfigResponse.__key_warning(key)
        return super().get(key, default)

    def __init__(__self__, *,
                 default_key_specs: Sequence['outputs.DnsKeySpecResponse'],
                 kind: str,
                 non_existence: str,
                 state: str):
        """
        :param Sequence['DnsKeySpecResponseArgs'] default_key_specs: Specifies parameters for generating initial DnsKeys for this ManagedZone. Can only be changed while the state is OFF.
        :param str non_existence: Specifies the mechanism for authenticated denial-of-existence responses. Can only be changed while the state is OFF.
        :param str state: Specifies whether DNSSEC is enabled, and what mode it is in.
        """
        pulumi.set(__self__, "default_key_specs", default_key_specs)
        pulumi.set(__self__, "kind", kind)
        pulumi.set(__self__, "non_existence", non_existence)
        pulumi.set(__self__, "state", state)

    @property
    @pulumi.getter(name="defaultKeySpecs")
    def default_key_specs(self) -> Sequence['outputs.DnsKeySpecResponse']:
        """
        Specifies parameters for generating initial DnsKeys for this ManagedZone. Can only be changed while the state is OFF.
        """
        return pulumi.get(self, "default_key_specs")

    @property
    @pulumi.getter
    def kind(self) -> str:
        return pulumi.get(self, "kind")

    @property
    @pulumi.getter(name="nonExistence")
    def non_existence(self) -> str:
        """
        Specifies the mechanism for authenticated denial-of-existence responses. Can only be changed while the state is OFF.
        """
        return pulumi.get(self, "non_existence")

    @property
    @pulumi.getter
    def state(self) -> str:
        """
        Specifies whether DNSSEC is enabled, and what mode it is in.
        """
        return pulumi.get(self, "state")


@pulumi.output_type
class ManagedZoneForwardingConfigNameServerTargetResponse(dict):
    @staticmethod
    def __key_warning(key: str):
        suggest = None
        if key == "forwardingPath":
            suggest = "forwarding_path"
        elif key == "ipv4Address":
            suggest = "ipv4_address"
        elif key == "ipv6Address":
            suggest = "ipv6_address"

        if suggest:
            pulumi.log.warn(f"Key '{key}' not found in ManagedZoneForwardingConfigNameServerTargetResponse. Access the value via the '{suggest}' property getter instead.")

    def __getitem__(self, key: str) -> Any:
        ManagedZoneForwardingConfigNameServerTargetResponse.__key_warning(key)
        return super().__getitem__(key)

    def get(self, key: str, default = None) -> Any:
        ManagedZoneForwardingConfigNameServerTargetResponse.__key_warning(key)
        return super().get(key, default)

    def __init__(__self__, *,
                 forwarding_path: str,
                 ipv4_address: str,
                 ipv6_address: str,
                 kind: str):
        """
        :param str forwarding_path: Forwarding path for this NameServerTarget. If unset or set to DEFAULT, Cloud DNS makes forwarding decisions based on IP address ranges; that is, RFC1918 addresses go to the VPC network, non-RFC1918 addresses go to the internet. When set to PRIVATE, Cloud DNS always sends queries through the VPC network for this target.
        :param str ipv4_address: IPv4 address of a target name server.
        :param str ipv6_address: IPv6 address of a target name server. Does not accept both fields (ipv4 & ipv6) being populated.
        """
        pulumi.set(__self__, "forwarding_path", forwarding_path)
        pulumi.set(__self__, "ipv4_address", ipv4_address)
        pulumi.set(__self__, "ipv6_address", ipv6_address)
        pulumi.set(__self__, "kind", kind)

    @property
    @pulumi.getter(name="forwardingPath")
    def forwarding_path(self) -> str:
        """
        Forwarding path for this NameServerTarget. If unset or set to DEFAULT, Cloud DNS makes forwarding decisions based on IP address ranges; that is, RFC1918 addresses go to the VPC network, non-RFC1918 addresses go to the internet. When set to PRIVATE, Cloud DNS always sends queries through the VPC network for this target.
        """
        return pulumi.get(self, "forwarding_path")

    @property
    @pulumi.getter(name="ipv4Address")
    def ipv4_address(self) -> str:
        """
        IPv4 address of a target name server.
        """
        return pulumi.get(self, "ipv4_address")

    @property
    @pulumi.getter(name="ipv6Address")
    def ipv6_address(self) -> str:
        """
        IPv6 address of a target name server. Does not accept both fields (ipv4 & ipv6) being populated.
        """
        return pulumi.get(self, "ipv6_address")

    @property
    @pulumi.getter
    def kind(self) -> str:
        return pulumi.get(self, "kind")


@pulumi.output_type
class ManagedZoneForwardingConfigResponse(dict):
    @staticmethod
    def __key_warning(key: str):
        suggest = None
        if key == "targetNameServers":
            suggest = "target_name_servers"

        if suggest:
            pulumi.log.warn(f"Key '{key}' not found in ManagedZoneForwardingConfigResponse. Access the value via the '{suggest}' property getter instead.")

    def __getitem__(self, key: str) -> Any:
        ManagedZoneForwardingConfigResponse.__key_warning(key)
        return super().__getitem__(key)

    def get(self, key: str, default = None) -> Any:
        ManagedZoneForwardingConfigResponse.__key_warning(key)
        return super().get(key, default)

    def __init__(__self__, *,
                 kind: str,
                 target_name_servers: Sequence['outputs.ManagedZoneForwardingConfigNameServerTargetResponse']):
        """
        :param Sequence['ManagedZoneForwardingConfigNameServerTargetResponseArgs'] target_name_servers: List of target name servers to forward to. Cloud DNS selects the best available name server if more than one target is given.
        """
        pulumi.set(__self__, "kind", kind)
        pulumi.set(__self__, "target_name_servers", target_name_servers)

    @property
    @pulumi.getter
    def kind(self) -> str:
        return pulumi.get(self, "kind")

    @property
    @pulumi.getter(name="targetNameServers")
    def target_name_servers(self) -> Sequence['outputs.ManagedZoneForwardingConfigNameServerTargetResponse']:
        """
        List of target name servers to forward to. Cloud DNS selects the best available name server if more than one target is given.
        """
        return pulumi.get(self, "target_name_servers")


@pulumi.output_type
class ManagedZonePeeringConfigResponse(dict):
    @staticmethod
    def __key_warning(key: str):
        suggest = None
        if key == "targetNetwork":
            suggest = "target_network"

        if suggest:
            pulumi.log.warn(f"Key '{key}' not found in ManagedZonePeeringConfigResponse. Access the value via the '{suggest}' property getter instead.")

    def __getitem__(self, key: str) -> Any:
        ManagedZonePeeringConfigResponse.__key_warning(key)
        return super().__getitem__(key)

    def get(self, key: str, default = None) -> Any:
        ManagedZonePeeringConfigResponse.__key_warning(key)
        return super().get(key, default)

    def __init__(__self__, *,
                 kind: str,
                 target_network: 'outputs.ManagedZonePeeringConfigTargetNetworkResponse'):
        """
        :param 'ManagedZonePeeringConfigTargetNetworkResponseArgs' target_network: The network with which to peer.
        """
        pulumi.set(__self__, "kind", kind)
        pulumi.set(__self__, "target_network", target_network)

    @property
    @pulumi.getter
    def kind(self) -> str:
        return pulumi.get(self, "kind")

    @property
    @pulumi.getter(name="targetNetwork")
    def target_network(self) -> 'outputs.ManagedZonePeeringConfigTargetNetworkResponse':
        """
        The network with which to peer.
        """
        return pulumi.get(self, "target_network")


@pulumi.output_type
class ManagedZonePeeringConfigTargetNetworkResponse(dict):
    @staticmethod
    def __key_warning(key: str):
        suggest = None
        if key == "deactivateTime":
            suggest = "deactivate_time"
        elif key == "networkUrl":
            suggest = "network_url"

        if suggest:
            pulumi.log.warn(f"Key '{key}' not found in ManagedZonePeeringConfigTargetNetworkResponse. Access the value via the '{suggest}' property getter instead.")

    def __getitem__(self, key: str) -> Any:
        ManagedZonePeeringConfigTargetNetworkResponse.__key_warning(key)
        return super().__getitem__(key)

    def get(self, key: str, default = None) -> Any:
        ManagedZonePeeringConfigTargetNetworkResponse.__key_warning(key)
        return super().get(key, default)

    def __init__(__self__, *,
                 deactivate_time: str,
                 kind: str,
                 network_url: str):
        """
        :param str deactivate_time: The time at which the zone was deactivated, in RFC 3339 date-time format. An empty string indicates that the peering connection is active. The producer network can deactivate a zone. The zone is automatically deactivated if the producer network that the zone targeted is deleted. Output only.
        :param str network_url: The fully qualified URL of the VPC network to forward queries to. This should be formatted like https://www.googleapis.com/compute/v1/projects/{project}/global/networks/{network}
        """
        pulumi.set(__self__, "deactivate_time", deactivate_time)
        pulumi.set(__self__, "kind", kind)
        pulumi.set(__self__, "network_url", network_url)

    @property
    @pulumi.getter(name="deactivateTime")
    def deactivate_time(self) -> str:
        """
        The time at which the zone was deactivated, in RFC 3339 date-time format. An empty string indicates that the peering connection is active. The producer network can deactivate a zone. The zone is automatically deactivated if the producer network that the zone targeted is deleted. Output only.
        """
        return pulumi.get(self, "deactivate_time")

    @property
    @pulumi.getter
    def kind(self) -> str:
        return pulumi.get(self, "kind")

    @property
    @pulumi.getter(name="networkUrl")
    def network_url(self) -> str:
        """
        The fully qualified URL of the VPC network to forward queries to. This should be formatted like https://www.googleapis.com/compute/v1/projects/{project}/global/networks/{network}
        """
        return pulumi.get(self, "network_url")


@pulumi.output_type
class ManagedZonePrivateVisibilityConfigNetworkResponse(dict):
    @staticmethod
    def __key_warning(key: str):
        suggest = None
        if key == "networkUrl":
            suggest = "network_url"

        if suggest:
            pulumi.log.warn(f"Key '{key}' not found in ManagedZonePrivateVisibilityConfigNetworkResponse. Access the value via the '{suggest}' property getter instead.")

    def __getitem__(self, key: str) -> Any:
        ManagedZonePrivateVisibilityConfigNetworkResponse.__key_warning(key)
        return super().__getitem__(key)

    def get(self, key: str, default = None) -> Any:
        ManagedZonePrivateVisibilityConfigNetworkResponse.__key_warning(key)
        return super().get(key, default)

    def __init__(__self__, *,
                 kind: str,
                 network_url: str):
        """
        :param str network_url: The fully qualified URL of the VPC network to bind to. Format this URL like https://www.googleapis.com/compute/v1/projects/{project}/global/networks/{network}
        """
        pulumi.set(__self__, "kind", kind)
        pulumi.set(__self__, "network_url", network_url)

    @property
    @pulumi.getter
    def kind(self) -> str:
        return pulumi.get(self, "kind")

    @property
    @pulumi.getter(name="networkUrl")
    def network_url(self) -> str:
        """
        The fully qualified URL of the VPC network to bind to. Format this URL like https://www.googleapis.com/compute/v1/projects/{project}/global/networks/{network}
        """
        return pulumi.get(self, "network_url")


@pulumi.output_type
class ManagedZonePrivateVisibilityConfigResponse(dict):
    def __init__(__self__, *,
                 kind: str,
                 networks: Sequence['outputs.ManagedZonePrivateVisibilityConfigNetworkResponse']):
        """
        :param Sequence['ManagedZonePrivateVisibilityConfigNetworkResponseArgs'] networks: The list of VPC networks that can see this zone.
        """
        pulumi.set(__self__, "kind", kind)
        pulumi.set(__self__, "networks", networks)

    @property
    @pulumi.getter
    def kind(self) -> str:
        return pulumi.get(self, "kind")

    @property
    @pulumi.getter
    def networks(self) -> Sequence['outputs.ManagedZonePrivateVisibilityConfigNetworkResponse']:
        """
        The list of VPC networks that can see this zone.
        """
        return pulumi.get(self, "networks")


@pulumi.output_type
class ManagedZoneReverseLookupConfigResponse(dict):
    def __init__(__self__, *,
                 kind: str):
        pulumi.set(__self__, "kind", kind)

    @property
    @pulumi.getter
    def kind(self) -> str:
        return pulumi.get(self, "kind")


@pulumi.output_type
class ManagedZoneServiceDirectoryConfigNamespaceResponse(dict):
    @staticmethod
    def __key_warning(key: str):
        suggest = None
        if key == "deletionTime":
            suggest = "deletion_time"
        elif key == "namespaceUrl":
            suggest = "namespace_url"

        if suggest:
            pulumi.log.warn(f"Key '{key}' not found in ManagedZoneServiceDirectoryConfigNamespaceResponse. Access the value via the '{suggest}' property getter instead.")

    def __getitem__(self, key: str) -> Any:
        ManagedZoneServiceDirectoryConfigNamespaceResponse.__key_warning(key)
        return super().__getitem__(key)

    def get(self, key: str, default = None) -> Any:
        ManagedZoneServiceDirectoryConfigNamespaceResponse.__key_warning(key)
        return super().get(key, default)

    def __init__(__self__, *,
                 deletion_time: str,
                 kind: str,
                 namespace_url: str):
        """
        :param str deletion_time: The time that the namespace backing this zone was deleted; an empty string if it still exists. This is in RFC3339 text format. Output only.
        :param str namespace_url: The fully qualified URL of the namespace associated with the zone. Format must be https://servicedirectory.googleapis.com/v1/projects/{project}/locations/{location}/namespaces/{namespace}
        """
        pulumi.set(__self__, "deletion_time", deletion_time)
        pulumi.set(__self__, "kind", kind)
        pulumi.set(__self__, "namespace_url", namespace_url)

    @property
    @pulumi.getter(name="deletionTime")
    def deletion_time(self) -> str:
        """
        The time that the namespace backing this zone was deleted; an empty string if it still exists. This is in RFC3339 text format. Output only.
        """
        return pulumi.get(self, "deletion_time")

    @property
    @pulumi.getter
    def kind(self) -> str:
        return pulumi.get(self, "kind")

    @property
    @pulumi.getter(name="namespaceUrl")
    def namespace_url(self) -> str:
        """
        The fully qualified URL of the namespace associated with the zone. Format must be https://servicedirectory.googleapis.com/v1/projects/{project}/locations/{location}/namespaces/{namespace}
        """
        return pulumi.get(self, "namespace_url")


@pulumi.output_type
class ManagedZoneServiceDirectoryConfigResponse(dict):
    """
    Contains information about Service Directory-backed zones.
    """
    def __init__(__self__, *,
                 kind: str,
                 namespace: 'outputs.ManagedZoneServiceDirectoryConfigNamespaceResponse'):
        """
        Contains information about Service Directory-backed zones.
        :param 'ManagedZoneServiceDirectoryConfigNamespaceResponseArgs' namespace: Contains information about the namespace associated with the zone.
        """
        pulumi.set(__self__, "kind", kind)
        pulumi.set(__self__, "namespace", namespace)

    @property
    @pulumi.getter
    def kind(self) -> str:
        return pulumi.get(self, "kind")

    @property
    @pulumi.getter
    def namespace(self) -> 'outputs.ManagedZoneServiceDirectoryConfigNamespaceResponse':
        """
        Contains information about the namespace associated with the zone.
        """
        return pulumi.get(self, "namespace")


@pulumi.output_type
class PolicyAlternativeNameServerConfigResponse(dict):
    @staticmethod
    def __key_warning(key: str):
        suggest = None
        if key == "targetNameServers":
            suggest = "target_name_servers"

        if suggest:
            pulumi.log.warn(f"Key '{key}' not found in PolicyAlternativeNameServerConfigResponse. Access the value via the '{suggest}' property getter instead.")

    def __getitem__(self, key: str) -> Any:
        PolicyAlternativeNameServerConfigResponse.__key_warning(key)
        return super().__getitem__(key)

    def get(self, key: str, default = None) -> Any:
        PolicyAlternativeNameServerConfigResponse.__key_warning(key)
        return super().get(key, default)

    def __init__(__self__, *,
                 kind: str,
                 target_name_servers: Sequence['outputs.PolicyAlternativeNameServerConfigTargetNameServerResponse']):
        """
        :param Sequence['PolicyAlternativeNameServerConfigTargetNameServerResponseArgs'] target_name_servers: Sets an alternative name server for the associated networks. When specified, all DNS queries are forwarded to a name server that you choose. Names such as .internal are not available when an alternative name server is specified.
        """
        pulumi.set(__self__, "kind", kind)
        pulumi.set(__self__, "target_name_servers", target_name_servers)

    @property
    @pulumi.getter
    def kind(self) -> str:
        return pulumi.get(self, "kind")

    @property
    @pulumi.getter(name="targetNameServers")
    def target_name_servers(self) -> Sequence['outputs.PolicyAlternativeNameServerConfigTargetNameServerResponse']:
        """
        Sets an alternative name server for the associated networks. When specified, all DNS queries are forwarded to a name server that you choose. Names such as .internal are not available when an alternative name server is specified.
        """
        return pulumi.get(self, "target_name_servers")


@pulumi.output_type
class PolicyAlternativeNameServerConfigTargetNameServerResponse(dict):
    @staticmethod
    def __key_warning(key: str):
        suggest = None
        if key == "forwardingPath":
            suggest = "forwarding_path"
        elif key == "ipv4Address":
            suggest = "ipv4_address"
        elif key == "ipv6Address":
            suggest = "ipv6_address"

        if suggest:
            pulumi.log.warn(f"Key '{key}' not found in PolicyAlternativeNameServerConfigTargetNameServerResponse. Access the value via the '{suggest}' property getter instead.")

    def __getitem__(self, key: str) -> Any:
        PolicyAlternativeNameServerConfigTargetNameServerResponse.__key_warning(key)
        return super().__getitem__(key)

    def get(self, key: str, default = None) -> Any:
        PolicyAlternativeNameServerConfigTargetNameServerResponse.__key_warning(key)
        return super().get(key, default)

    def __init__(__self__, *,
                 forwarding_path: str,
                 ipv4_address: str,
                 ipv6_address: str,
                 kind: str):
        """
        :param str forwarding_path: Forwarding path for this TargetNameServer. If unset or set to DEFAULT, Cloud DNS makes forwarding decisions based on address ranges; that is, RFC1918 addresses go to the VPC network, non-RFC1918 addresses go to the internet. When set to PRIVATE, Cloud DNS always sends queries through the VPC network for this target.
        :param str ipv4_address: IPv4 address to forward to.
        :param str ipv6_address: IPv6 address to forward to. Does not accept both fields (ipv4 & ipv6) being populated.
        """
        pulumi.set(__self__, "forwarding_path", forwarding_path)
        pulumi.set(__self__, "ipv4_address", ipv4_address)
        pulumi.set(__self__, "ipv6_address", ipv6_address)
        pulumi.set(__self__, "kind", kind)

    @property
    @pulumi.getter(name="forwardingPath")
    def forwarding_path(self) -> str:
        """
        Forwarding path for this TargetNameServer. If unset or set to DEFAULT, Cloud DNS makes forwarding decisions based on address ranges; that is, RFC1918 addresses go to the VPC network, non-RFC1918 addresses go to the internet. When set to PRIVATE, Cloud DNS always sends queries through the VPC network for this target.
        """
        return pulumi.get(self, "forwarding_path")

    @property
    @pulumi.getter(name="ipv4Address")
    def ipv4_address(self) -> str:
        """
        IPv4 address to forward to.
        """
        return pulumi.get(self, "ipv4_address")

    @property
    @pulumi.getter(name="ipv6Address")
    def ipv6_address(self) -> str:
        """
        IPv6 address to forward to. Does not accept both fields (ipv4 & ipv6) being populated.
        """
        return pulumi.get(self, "ipv6_address")

    @property
    @pulumi.getter
    def kind(self) -> str:
        return pulumi.get(self, "kind")


@pulumi.output_type
class PolicyNetworkResponse(dict):
    @staticmethod
    def __key_warning(key: str):
        suggest = None
        if key == "networkUrl":
            suggest = "network_url"

        if suggest:
            pulumi.log.warn(f"Key '{key}' not found in PolicyNetworkResponse. Access the value via the '{suggest}' property getter instead.")

    def __getitem__(self, key: str) -> Any:
        PolicyNetworkResponse.__key_warning(key)
        return super().__getitem__(key)

    def get(self, key: str, default = None) -> Any:
        PolicyNetworkResponse.__key_warning(key)
        return super().get(key, default)

    def __init__(__self__, *,
                 kind: str,
                 network_url: str):
        """
        :param str network_url: The fully qualified URL of the VPC network to bind to. This should be formatted like https://www.googleapis.com/compute/v1/projects/{project}/global/networks/{network}
        """
        pulumi.set(__self__, "kind", kind)
        pulumi.set(__self__, "network_url", network_url)

    @property
    @pulumi.getter
    def kind(self) -> str:
        return pulumi.get(self, "kind")

    @property
    @pulumi.getter(name="networkUrl")
    def network_url(self) -> str:
        """
        The fully qualified URL of the VPC network to bind to. This should be formatted like https://www.googleapis.com/compute/v1/projects/{project}/global/networks/{network}
        """
        return pulumi.get(self, "network_url")


@pulumi.output_type
class ResourceRecordSetResponse(dict):
    """
    A unit of data that is returned by the DNS servers.
    """
    @staticmethod
    def __key_warning(key: str):
        suggest = None
        if key == "signatureRrdatas":
            suggest = "signature_rrdatas"

        if suggest:
            pulumi.log.warn(f"Key '{key}' not found in ResourceRecordSetResponse. Access the value via the '{suggest}' property getter instead.")

    def __getitem__(self, key: str) -> Any:
        ResourceRecordSetResponse.__key_warning(key)
        return super().__getitem__(key)

    def get(self, key: str, default = None) -> Any:
        ResourceRecordSetResponse.__key_warning(key)
        return super().get(key, default)

    def __init__(__self__, *,
                 kind: str,
                 name: str,
                 rrdatas: Sequence[str],
                 signature_rrdatas: Sequence[str],
                 ttl: int,
                 type: str):
        """
        A unit of data that is returned by the DNS servers.
        :param str name: For example, www.example.com.
        :param Sequence[str] rrdatas: As defined in RFC 1035 (section 5) and RFC 1034 (section 3.6.1) -- see examples.
        :param Sequence[str] signature_rrdatas: As defined in RFC 4034 (section 3.2).
        :param int ttl: Number of seconds that this ResourceRecordSet can be cached by resolvers.
        :param str type: The identifier of a supported record type. See the list of Supported DNS record types.
        """
        pulumi.set(__self__, "kind", kind)
        pulumi.set(__self__, "name", name)
        pulumi.set(__self__, "rrdatas", rrdatas)
        pulumi.set(__self__, "signature_rrdatas", signature_rrdatas)
        pulumi.set(__self__, "ttl", ttl)
        pulumi.set(__self__, "type", type)

    @property
    @pulumi.getter
    def kind(self) -> str:
        return pulumi.get(self, "kind")

    @property
    @pulumi.getter
    def name(self) -> str:
        """
        For example, www.example.com.
        """
        return pulumi.get(self, "name")

    @property
    @pulumi.getter
    def rrdatas(self) -> Sequence[str]:
        """
        As defined in RFC 1035 (section 5) and RFC 1034 (section 3.6.1) -- see examples.
        """
        return pulumi.get(self, "rrdatas")

    @property
    @pulumi.getter(name="signatureRrdatas")
    def signature_rrdatas(self) -> Sequence[str]:
        """
        As defined in RFC 4034 (section 3.2).
        """
        return pulumi.get(self, "signature_rrdatas")

    @property
    @pulumi.getter
    def ttl(self) -> int:
        """
        Number of seconds that this ResourceRecordSet can be cached by resolvers.
        """
        return pulumi.get(self, "ttl")

    @property
    @pulumi.getter
    def type(self) -> str:
        """
        The identifier of a supported record type. See the list of Supported DNS record types.
        """
        return pulumi.get(self, "type")


@pulumi.output_type
class ResponsePolicyNetworkResponse(dict):
    @staticmethod
    def __key_warning(key: str):
        suggest = None
        if key == "networkUrl":
            suggest = "network_url"

        if suggest:
            pulumi.log.warn(f"Key '{key}' not found in ResponsePolicyNetworkResponse. Access the value via the '{suggest}' property getter instead.")

    def __getitem__(self, key: str) -> Any:
        ResponsePolicyNetworkResponse.__key_warning(key)
        return super().__getitem__(key)

    def get(self, key: str, default = None) -> Any:
        ResponsePolicyNetworkResponse.__key_warning(key)
        return super().get(key, default)

    def __init__(__self__, *,
                 kind: str,
                 network_url: str):
        """
        :param str network_url: The fully qualified URL of the VPC network to bind to. This should be formatted like https://www.googleapis.com/compute/v1/projects/{project}/global/networks/{network}
        """
        pulumi.set(__self__, "kind", kind)
        pulumi.set(__self__, "network_url", network_url)

    @property
    @pulumi.getter
    def kind(self) -> str:
        return pulumi.get(self, "kind")

    @property
    @pulumi.getter(name="networkUrl")
    def network_url(self) -> str:
        """
        The fully qualified URL of the VPC network to bind to. This should be formatted like https://www.googleapis.com/compute/v1/projects/{project}/global/networks/{network}
        """
        return pulumi.get(self, "network_url")


@pulumi.output_type
class ResponsePolicyRuleLocalDataResponse(dict):
    @staticmethod
    def __key_warning(key: str):
        suggest = None
        if key == "localDatas":
            suggest = "local_datas"

        if suggest:
            pulumi.log.warn(f"Key '{key}' not found in ResponsePolicyRuleLocalDataResponse. Access the value via the '{suggest}' property getter instead.")

    def __getitem__(self, key: str) -> Any:
        ResponsePolicyRuleLocalDataResponse.__key_warning(key)
        return super().__getitem__(key)

    def get(self, key: str, default = None) -> Any:
        ResponsePolicyRuleLocalDataResponse.__key_warning(key)
        return super().get(key, default)

    def __init__(__self__, *,
                 local_datas: Sequence['outputs.ResourceRecordSetResponse']):
        """
        :param Sequence['ResourceRecordSetResponseArgs'] local_datas: All resource record sets for this selector, one per resource record type. The name must match the dns_name.
        """
        pulumi.set(__self__, "local_datas", local_datas)

    @property
    @pulumi.getter(name="localDatas")
    def local_datas(self) -> Sequence['outputs.ResourceRecordSetResponse']:
        """
        All resource record sets for this selector, one per resource record type. The name must match the dns_name.
        """
        return pulumi.get(self, "local_datas")


