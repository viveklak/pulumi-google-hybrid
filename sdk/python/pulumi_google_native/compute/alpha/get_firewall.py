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
    'GetFirewallResult',
    'AwaitableGetFirewallResult',
    'get_firewall',
]

@pulumi.output_type
class GetFirewallResult:
    def __init__(__self__, allowed=None, creation_timestamp=None, denied=None, description=None, destination_ranges=None, direction=None, disabled=None, kind=None, log_config=None, name=None, network=None, priority=None, self_link=None, self_link_with_id=None, source_ranges=None, source_service_accounts=None, source_tags=None, target_service_accounts=None, target_tags=None):
        if allowed and not isinstance(allowed, list):
            raise TypeError("Expected argument 'allowed' to be a list")
        pulumi.set(__self__, "allowed", allowed)
        if creation_timestamp and not isinstance(creation_timestamp, str):
            raise TypeError("Expected argument 'creation_timestamp' to be a str")
        pulumi.set(__self__, "creation_timestamp", creation_timestamp)
        if denied and not isinstance(denied, list):
            raise TypeError("Expected argument 'denied' to be a list")
        pulumi.set(__self__, "denied", denied)
        if description and not isinstance(description, str):
            raise TypeError("Expected argument 'description' to be a str")
        pulumi.set(__self__, "description", description)
        if destination_ranges and not isinstance(destination_ranges, list):
            raise TypeError("Expected argument 'destination_ranges' to be a list")
        pulumi.set(__self__, "destination_ranges", destination_ranges)
        if direction and not isinstance(direction, str):
            raise TypeError("Expected argument 'direction' to be a str")
        pulumi.set(__self__, "direction", direction)
        if disabled and not isinstance(disabled, bool):
            raise TypeError("Expected argument 'disabled' to be a bool")
        pulumi.set(__self__, "disabled", disabled)
        if kind and not isinstance(kind, str):
            raise TypeError("Expected argument 'kind' to be a str")
        pulumi.set(__self__, "kind", kind)
        if log_config and not isinstance(log_config, dict):
            raise TypeError("Expected argument 'log_config' to be a dict")
        pulumi.set(__self__, "log_config", log_config)
        if name and not isinstance(name, str):
            raise TypeError("Expected argument 'name' to be a str")
        pulumi.set(__self__, "name", name)
        if network and not isinstance(network, str):
            raise TypeError("Expected argument 'network' to be a str")
        pulumi.set(__self__, "network", network)
        if priority and not isinstance(priority, int):
            raise TypeError("Expected argument 'priority' to be a int")
        pulumi.set(__self__, "priority", priority)
        if self_link and not isinstance(self_link, str):
            raise TypeError("Expected argument 'self_link' to be a str")
        pulumi.set(__self__, "self_link", self_link)
        if self_link_with_id and not isinstance(self_link_with_id, str):
            raise TypeError("Expected argument 'self_link_with_id' to be a str")
        pulumi.set(__self__, "self_link_with_id", self_link_with_id)
        if source_ranges and not isinstance(source_ranges, list):
            raise TypeError("Expected argument 'source_ranges' to be a list")
        pulumi.set(__self__, "source_ranges", source_ranges)
        if source_service_accounts and not isinstance(source_service_accounts, list):
            raise TypeError("Expected argument 'source_service_accounts' to be a list")
        pulumi.set(__self__, "source_service_accounts", source_service_accounts)
        if source_tags and not isinstance(source_tags, list):
            raise TypeError("Expected argument 'source_tags' to be a list")
        pulumi.set(__self__, "source_tags", source_tags)
        if target_service_accounts and not isinstance(target_service_accounts, list):
            raise TypeError("Expected argument 'target_service_accounts' to be a list")
        pulumi.set(__self__, "target_service_accounts", target_service_accounts)
        if target_tags and not isinstance(target_tags, list):
            raise TypeError("Expected argument 'target_tags' to be a list")
        pulumi.set(__self__, "target_tags", target_tags)

    @property
    @pulumi.getter
    def allowed(self) -> Sequence['outputs.FirewallAllowedItemResponse']:
        """
        The list of ALLOW rules specified by this firewall. Each rule specifies a protocol and port-range tuple that describes a permitted connection.
        """
        return pulumi.get(self, "allowed")

    @property
    @pulumi.getter(name="creationTimestamp")
    def creation_timestamp(self) -> str:
        """
        Creation timestamp in RFC3339 text format.
        """
        return pulumi.get(self, "creation_timestamp")

    @property
    @pulumi.getter
    def denied(self) -> Sequence['outputs.FirewallDeniedItemResponse']:
        """
        The list of DENY rules specified by this firewall. Each rule specifies a protocol and port-range tuple that describes a denied connection.
        """
        return pulumi.get(self, "denied")

    @property
    @pulumi.getter
    def description(self) -> str:
        """
        An optional description of this resource. Provide this field when you create the resource.
        """
        return pulumi.get(self, "description")

    @property
    @pulumi.getter(name="destinationRanges")
    def destination_ranges(self) -> Sequence[str]:
        """
        If destination ranges are specified, the firewall rule applies only to traffic that has destination IP address in these ranges. These ranges must be expressed in CIDR format. Only IPv4 is supported.
        """
        return pulumi.get(self, "destination_ranges")

    @property
    @pulumi.getter
    def direction(self) -> str:
        """
        Direction of traffic to which this firewall applies, either `INGRESS` or `EGRESS`. The default is `INGRESS`. For `INGRESS` traffic, you cannot specify the destinationRanges field, and for `EGRESS` traffic, you cannot specify the sourceRanges or sourceTags fields.
        """
        return pulumi.get(self, "direction")

    @property
    @pulumi.getter
    def disabled(self) -> bool:
        """
        Denotes whether the firewall rule is disabled. When set to true, the firewall rule is not enforced and the network behaves as if it did not exist. If this is unspecified, the firewall rule will be enabled.
        """
        return pulumi.get(self, "disabled")

    @property
    @pulumi.getter
    def kind(self) -> str:
        """
        Type of the resource. Always compute#firewall for firewall rules.
        """
        return pulumi.get(self, "kind")

    @property
    @pulumi.getter(name="logConfig")
    def log_config(self) -> 'outputs.FirewallLogConfigResponse':
        """
        This field denotes the logging options for a particular firewall rule. If logging is enabled, logs will be exported to Cloud Logging.
        """
        return pulumi.get(self, "log_config")

    @property
    @pulumi.getter
    def name(self) -> str:
        """
        Name of the resource; provided by the client when the resource is created. The name must be 1-63 characters long, and comply with RFC1035. Specifically, the name must be 1-63 characters long and match the regular expression `[a-z]([-a-z0-9]*[a-z0-9])?. The first character must be a lowercase letter, and all following characters (except for the last character) must be a dash, lowercase letter, or digit. The last character must be a lowercase letter or digit.
        """
        return pulumi.get(self, "name")

    @property
    @pulumi.getter
    def network(self) -> str:
        """
        URL of the network resource for this firewall rule. If not specified when creating a firewall rule, the default network is used:
        global/networks/default
        If you choose to specify this field, you can specify the network as a full or partial URL. For example, the following are all valid URLs:  
        - https://www.googleapis.com/compute/v1/projects/myproject/global/networks/my-network 
        - projects/myproject/global/networks/my-network 
        - global/networks/default
        """
        return pulumi.get(self, "network")

    @property
    @pulumi.getter
    def priority(self) -> int:
        """
        Priority for this rule. This is an integer between `0` and `65535`, both inclusive. The default value is `1000`. Relative priorities determine which rule takes effect if multiple rules apply. Lower values indicate higher priority. For example, a rule with priority `0` has higher precedence than a rule with priority `1`. DENY rules take precedence over ALLOW rules if they have equal priority. Note that VPC networks have implied rules with a priority of `65535`. To avoid conflicts with the implied rules, use a priority number less than `65535`.
        """
        return pulumi.get(self, "priority")

    @property
    @pulumi.getter(name="selfLink")
    def self_link(self) -> str:
        """
        Server-defined URL for the resource.
        """
        return pulumi.get(self, "self_link")

    @property
    @pulumi.getter(name="selfLinkWithId")
    def self_link_with_id(self) -> str:
        """
        Server-defined URL for this resource with the resource id.
        """
        return pulumi.get(self, "self_link_with_id")

    @property
    @pulumi.getter(name="sourceRanges")
    def source_ranges(self) -> Sequence[str]:
        """
        If source ranges are specified, the firewall rule applies only to traffic that has a source IP address in these ranges. These ranges must be expressed in CIDR format. One or both of sourceRanges and sourceTags may be set. If both fields are set, the rule applies to traffic that has a source IP address within sourceRanges OR a source IP from a resource with a matching tag listed in the sourceTags field. The connection does not need to match both fields for the rule to apply. Only IPv4 is supported.
        """
        return pulumi.get(self, "source_ranges")

    @property
    @pulumi.getter(name="sourceServiceAccounts")
    def source_service_accounts(self) -> Sequence[str]:
        """
        If source service accounts are specified, the firewall rules apply only to traffic originating from an instance with a service account in this list. Source service accounts cannot be used to control traffic to an instance's external IP address because service accounts are associated with an instance, not an IP address. sourceRanges can be set at the same time as sourceServiceAccounts. If both are set, the firewall applies to traffic that has a source IP address within the sourceRanges OR a source IP that belongs to an instance with service account listed in sourceServiceAccount. The connection does not need to match both fields for the firewall to apply. sourceServiceAccounts cannot be used at the same time as sourceTags or targetTags.
        """
        return pulumi.get(self, "source_service_accounts")

    @property
    @pulumi.getter(name="sourceTags")
    def source_tags(self) -> Sequence[str]:
        """
        If source tags are specified, the firewall rule applies only to traffic with source IPs that match the primary network interfaces of VM instances that have the tag and are in the same VPC network. Source tags cannot be used to control traffic to an instance's external IP address, it only applies to traffic between instances in the same virtual network. Because tags are associated with instances, not IP addresses. One or both of sourceRanges and sourceTags may be set. If both fields are set, the firewall applies to traffic that has a source IP address within sourceRanges OR a source IP from a resource with a matching tag listed in the sourceTags field. The connection does not need to match both fields for the firewall to apply.
        """
        return pulumi.get(self, "source_tags")

    @property
    @pulumi.getter(name="targetServiceAccounts")
    def target_service_accounts(self) -> Sequence[str]:
        """
        A list of service accounts indicating sets of instances located in the network that may make network connections as specified in allowed[]. targetServiceAccounts cannot be used at the same time as targetTags or sourceTags. If neither targetServiceAccounts nor targetTags are specified, the firewall rule applies to all instances on the specified network.
        """
        return pulumi.get(self, "target_service_accounts")

    @property
    @pulumi.getter(name="targetTags")
    def target_tags(self) -> Sequence[str]:
        """
        A list of tags that controls which instances the firewall rule applies to. If targetTags are specified, then the firewall rule applies only to instances in the VPC network that have one of those tags. If no targetTags are specified, the firewall rule applies to all instances on the specified network.
        """
        return pulumi.get(self, "target_tags")


class AwaitableGetFirewallResult(GetFirewallResult):
    # pylint: disable=using-constant-test
    def __await__(self):
        if False:
            yield self
        return GetFirewallResult(
            allowed=self.allowed,
            creation_timestamp=self.creation_timestamp,
            denied=self.denied,
            description=self.description,
            destination_ranges=self.destination_ranges,
            direction=self.direction,
            disabled=self.disabled,
            kind=self.kind,
            log_config=self.log_config,
            name=self.name,
            network=self.network,
            priority=self.priority,
            self_link=self.self_link,
            self_link_with_id=self.self_link_with_id,
            source_ranges=self.source_ranges,
            source_service_accounts=self.source_service_accounts,
            source_tags=self.source_tags,
            target_service_accounts=self.target_service_accounts,
            target_tags=self.target_tags)


def get_firewall(firewall: Optional[str] = None,
                 project: Optional[str] = None,
                 opts: Optional[pulumi.InvokeOptions] = None) -> AwaitableGetFirewallResult:
    """
    Returns the specified firewall.
    """
    __args__ = dict()
    __args__['firewall'] = firewall
    __args__['project'] = project
    if opts is None:
        opts = pulumi.InvokeOptions()
    if opts.version is None:
        opts.version = _utilities.get_version()
    __ret__ = pulumi.runtime.invoke('google-native:compute/alpha:getFirewall', __args__, opts=opts, typ=GetFirewallResult).value

    return AwaitableGetFirewallResult(
        allowed=__ret__.allowed,
        creation_timestamp=__ret__.creation_timestamp,
        denied=__ret__.denied,
        description=__ret__.description,
        destination_ranges=__ret__.destination_ranges,
        direction=__ret__.direction,
        disabled=__ret__.disabled,
        kind=__ret__.kind,
        log_config=__ret__.log_config,
        name=__ret__.name,
        network=__ret__.network,
        priority=__ret__.priority,
        self_link=__ret__.self_link,
        self_link_with_id=__ret__.self_link_with_id,
        source_ranges=__ret__.source_ranges,
        source_service_accounts=__ret__.source_service_accounts,
        source_tags=__ret__.source_tags,
        target_service_accounts=__ret__.target_service_accounts,
        target_tags=__ret__.target_tags)
