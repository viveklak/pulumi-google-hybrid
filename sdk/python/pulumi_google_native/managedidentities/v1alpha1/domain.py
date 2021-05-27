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

__all__ = ['DomainArgs', 'Domain']

@pulumi.input_type
class DomainArgs:
    def __init__(__self__, *,
                 project: pulumi.Input[str],
                 audit_logs_enabled: Optional[pulumi.Input[bool]] = None,
                 authorized_networks: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]] = None,
                 create_time: Optional[pulumi.Input[str]] = None,
                 domain_name: Optional[pulumi.Input[str]] = None,
                 fqdn: Optional[pulumi.Input[str]] = None,
                 labels: Optional[pulumi.Input[Mapping[str, pulumi.Input[str]]]] = None,
                 locations: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]] = None,
                 managed_identities_admin_name: Optional[pulumi.Input[str]] = None,
                 name: Optional[pulumi.Input[str]] = None,
                 reserved_ip_range: Optional[pulumi.Input[str]] = None,
                 state: Optional[pulumi.Input[str]] = None,
                 status_message: Optional[pulumi.Input[str]] = None,
                 trusts: Optional[pulumi.Input[Sequence[pulumi.Input['TrustArgs']]]] = None,
                 update_time: Optional[pulumi.Input[str]] = None):
        """
        The set of arguments for constructing a Domain resource.
        :param pulumi.Input[bool] audit_logs_enabled: Optional. Configuration for audit logs. True if audit logs are enabled, else false. Default is audit logs disabled.
        :param pulumi.Input[Sequence[pulumi.Input[str]]] authorized_networks: Optional. The full names of the Google Compute Engine [networks](/compute/docs/networks-and-firewalls#networks) to which the instance is connected. Network can be added using UpdateDomain later. Domain is only available on network part of authorized_networks. Caller needs to make sure that CIDR subnets do not overlap between networks, else domain creation will fail.
        :param pulumi.Input[str] create_time: The time the instance was created. Synthetic field is populated automatically by CCFE. go/ccfe-synthetic-field-user-guide
        :param pulumi.Input[str] fqdn: Fully-qualified domain name of the exposed domain used by clients to connect to the service. Similar to what would be chosen for an Active Directory that is set up on an internal network.
        :param pulumi.Input[Mapping[str, pulumi.Input[str]]] labels: Optional. Resource labels to represent user provided metadata
        :param pulumi.Input[Sequence[pulumi.Input[str]]] locations: Required. Locations where domain needs to be provisioned. regions e.g. us-west1 or us-east4 Service supports up to 4 locations at once. Each location will use a /26 block.
        :param pulumi.Input[str] managed_identities_admin_name: Optional. Name of customer-visible admin used to perform Active Directory operations. If not specified `setupadmin` would be used.
        :param pulumi.Input[str] name: Unique name of the domain in this scope including projects and location using the form: `projects/{project_id}/locations/global/domains/{domain_name}`.
        :param pulumi.Input[str] reserved_ip_range: Required. The CIDR range of internal addresses that are reserved for this domain. Reserved networks must be /24 or larger. Ranges must be unique and non-overlapping with existing subnets in [Domain].[authorized_networks].
        :param pulumi.Input[str] state: The current state of this domain.
        :param pulumi.Input[str] status_message: Additional information about the current status of this domain, if available.
        :param pulumi.Input[Sequence[pulumi.Input['TrustArgs']]] trusts: The current trusts associated with the domain.
        :param pulumi.Input[str] update_time: Last update time. Synthetic field is populated automatically by CCFE.
        """
        pulumi.set(__self__, "project", project)
        if audit_logs_enabled is not None:
            pulumi.set(__self__, "audit_logs_enabled", audit_logs_enabled)
        if authorized_networks is not None:
            pulumi.set(__self__, "authorized_networks", authorized_networks)
        if create_time is not None:
            pulumi.set(__self__, "create_time", create_time)
        if domain_name is not None:
            pulumi.set(__self__, "domain_name", domain_name)
        if fqdn is not None:
            pulumi.set(__self__, "fqdn", fqdn)
        if labels is not None:
            pulumi.set(__self__, "labels", labels)
        if locations is not None:
            pulumi.set(__self__, "locations", locations)
        if managed_identities_admin_name is not None:
            pulumi.set(__self__, "managed_identities_admin_name", managed_identities_admin_name)
        if name is not None:
            pulumi.set(__self__, "name", name)
        if reserved_ip_range is not None:
            pulumi.set(__self__, "reserved_ip_range", reserved_ip_range)
        if state is not None:
            pulumi.set(__self__, "state", state)
        if status_message is not None:
            pulumi.set(__self__, "status_message", status_message)
        if trusts is not None:
            pulumi.set(__self__, "trusts", trusts)
        if update_time is not None:
            pulumi.set(__self__, "update_time", update_time)

    @property
    @pulumi.getter
    def project(self) -> pulumi.Input[str]:
        return pulumi.get(self, "project")

    @project.setter
    def project(self, value: pulumi.Input[str]):
        pulumi.set(self, "project", value)

    @property
    @pulumi.getter(name="auditLogsEnabled")
    def audit_logs_enabled(self) -> Optional[pulumi.Input[bool]]:
        """
        Optional. Configuration for audit logs. True if audit logs are enabled, else false. Default is audit logs disabled.
        """
        return pulumi.get(self, "audit_logs_enabled")

    @audit_logs_enabled.setter
    def audit_logs_enabled(self, value: Optional[pulumi.Input[bool]]):
        pulumi.set(self, "audit_logs_enabled", value)

    @property
    @pulumi.getter(name="authorizedNetworks")
    def authorized_networks(self) -> Optional[pulumi.Input[Sequence[pulumi.Input[str]]]]:
        """
        Optional. The full names of the Google Compute Engine [networks](/compute/docs/networks-and-firewalls#networks) to which the instance is connected. Network can be added using UpdateDomain later. Domain is only available on network part of authorized_networks. Caller needs to make sure that CIDR subnets do not overlap between networks, else domain creation will fail.
        """
        return pulumi.get(self, "authorized_networks")

    @authorized_networks.setter
    def authorized_networks(self, value: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]]):
        pulumi.set(self, "authorized_networks", value)

    @property
    @pulumi.getter(name="createTime")
    def create_time(self) -> Optional[pulumi.Input[str]]:
        """
        The time the instance was created. Synthetic field is populated automatically by CCFE. go/ccfe-synthetic-field-user-guide
        """
        return pulumi.get(self, "create_time")

    @create_time.setter
    def create_time(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "create_time", value)

    @property
    @pulumi.getter(name="domainName")
    def domain_name(self) -> Optional[pulumi.Input[str]]:
        return pulumi.get(self, "domain_name")

    @domain_name.setter
    def domain_name(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "domain_name", value)

    @property
    @pulumi.getter
    def fqdn(self) -> Optional[pulumi.Input[str]]:
        """
        Fully-qualified domain name of the exposed domain used by clients to connect to the service. Similar to what would be chosen for an Active Directory that is set up on an internal network.
        """
        return pulumi.get(self, "fqdn")

    @fqdn.setter
    def fqdn(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "fqdn", value)

    @property
    @pulumi.getter
    def labels(self) -> Optional[pulumi.Input[Mapping[str, pulumi.Input[str]]]]:
        """
        Optional. Resource labels to represent user provided metadata
        """
        return pulumi.get(self, "labels")

    @labels.setter
    def labels(self, value: Optional[pulumi.Input[Mapping[str, pulumi.Input[str]]]]):
        pulumi.set(self, "labels", value)

    @property
    @pulumi.getter
    def locations(self) -> Optional[pulumi.Input[Sequence[pulumi.Input[str]]]]:
        """
        Required. Locations where domain needs to be provisioned. regions e.g. us-west1 or us-east4 Service supports up to 4 locations at once. Each location will use a /26 block.
        """
        return pulumi.get(self, "locations")

    @locations.setter
    def locations(self, value: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]]):
        pulumi.set(self, "locations", value)

    @property
    @pulumi.getter(name="managedIdentitiesAdminName")
    def managed_identities_admin_name(self) -> Optional[pulumi.Input[str]]:
        """
        Optional. Name of customer-visible admin used to perform Active Directory operations. If not specified `setupadmin` would be used.
        """
        return pulumi.get(self, "managed_identities_admin_name")

    @managed_identities_admin_name.setter
    def managed_identities_admin_name(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "managed_identities_admin_name", value)

    @property
    @pulumi.getter
    def name(self) -> Optional[pulumi.Input[str]]:
        """
        Unique name of the domain in this scope including projects and location using the form: `projects/{project_id}/locations/global/domains/{domain_name}`.
        """
        return pulumi.get(self, "name")

    @name.setter
    def name(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "name", value)

    @property
    @pulumi.getter(name="reservedIpRange")
    def reserved_ip_range(self) -> Optional[pulumi.Input[str]]:
        """
        Required. The CIDR range of internal addresses that are reserved for this domain. Reserved networks must be /24 or larger. Ranges must be unique and non-overlapping with existing subnets in [Domain].[authorized_networks].
        """
        return pulumi.get(self, "reserved_ip_range")

    @reserved_ip_range.setter
    def reserved_ip_range(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "reserved_ip_range", value)

    @property
    @pulumi.getter
    def state(self) -> Optional[pulumi.Input[str]]:
        """
        The current state of this domain.
        """
        return pulumi.get(self, "state")

    @state.setter
    def state(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "state", value)

    @property
    @pulumi.getter(name="statusMessage")
    def status_message(self) -> Optional[pulumi.Input[str]]:
        """
        Additional information about the current status of this domain, if available.
        """
        return pulumi.get(self, "status_message")

    @status_message.setter
    def status_message(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "status_message", value)

    @property
    @pulumi.getter
    def trusts(self) -> Optional[pulumi.Input[Sequence[pulumi.Input['TrustArgs']]]]:
        """
        The current trusts associated with the domain.
        """
        return pulumi.get(self, "trusts")

    @trusts.setter
    def trusts(self, value: Optional[pulumi.Input[Sequence[pulumi.Input['TrustArgs']]]]):
        pulumi.set(self, "trusts", value)

    @property
    @pulumi.getter(name="updateTime")
    def update_time(self) -> Optional[pulumi.Input[str]]:
        """
        Last update time. Synthetic field is populated automatically by CCFE.
        """
        return pulumi.get(self, "update_time")

    @update_time.setter
    def update_time(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "update_time", value)


class Domain(pulumi.CustomResource):
    @overload
    def __init__(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 audit_logs_enabled: Optional[pulumi.Input[bool]] = None,
                 authorized_networks: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]] = None,
                 create_time: Optional[pulumi.Input[str]] = None,
                 domain_name: Optional[pulumi.Input[str]] = None,
                 fqdn: Optional[pulumi.Input[str]] = None,
                 labels: Optional[pulumi.Input[Mapping[str, pulumi.Input[str]]]] = None,
                 locations: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]] = None,
                 managed_identities_admin_name: Optional[pulumi.Input[str]] = None,
                 name: Optional[pulumi.Input[str]] = None,
                 project: Optional[pulumi.Input[str]] = None,
                 reserved_ip_range: Optional[pulumi.Input[str]] = None,
                 state: Optional[pulumi.Input[str]] = None,
                 status_message: Optional[pulumi.Input[str]] = None,
                 trusts: Optional[pulumi.Input[Sequence[pulumi.Input[pulumi.InputType['TrustArgs']]]]] = None,
                 update_time: Optional[pulumi.Input[str]] = None,
                 __props__=None):
        """
        Creates a Microsoft AD Domain in a given project. Operation

        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[bool] audit_logs_enabled: Optional. Configuration for audit logs. True if audit logs are enabled, else false. Default is audit logs disabled.
        :param pulumi.Input[Sequence[pulumi.Input[str]]] authorized_networks: Optional. The full names of the Google Compute Engine [networks](/compute/docs/networks-and-firewalls#networks) to which the instance is connected. Network can be added using UpdateDomain later. Domain is only available on network part of authorized_networks. Caller needs to make sure that CIDR subnets do not overlap between networks, else domain creation will fail.
        :param pulumi.Input[str] create_time: The time the instance was created. Synthetic field is populated automatically by CCFE. go/ccfe-synthetic-field-user-guide
        :param pulumi.Input[str] fqdn: Fully-qualified domain name of the exposed domain used by clients to connect to the service. Similar to what would be chosen for an Active Directory that is set up on an internal network.
        :param pulumi.Input[Mapping[str, pulumi.Input[str]]] labels: Optional. Resource labels to represent user provided metadata
        :param pulumi.Input[Sequence[pulumi.Input[str]]] locations: Required. Locations where domain needs to be provisioned. regions e.g. us-west1 or us-east4 Service supports up to 4 locations at once. Each location will use a /26 block.
        :param pulumi.Input[str] managed_identities_admin_name: Optional. Name of customer-visible admin used to perform Active Directory operations. If not specified `setupadmin` would be used.
        :param pulumi.Input[str] name: Unique name of the domain in this scope including projects and location using the form: `projects/{project_id}/locations/global/domains/{domain_name}`.
        :param pulumi.Input[str] reserved_ip_range: Required. The CIDR range of internal addresses that are reserved for this domain. Reserved networks must be /24 or larger. Ranges must be unique and non-overlapping with existing subnets in [Domain].[authorized_networks].
        :param pulumi.Input[str] state: The current state of this domain.
        :param pulumi.Input[str] status_message: Additional information about the current status of this domain, if available.
        :param pulumi.Input[Sequence[pulumi.Input[pulumi.InputType['TrustArgs']]]] trusts: The current trusts associated with the domain.
        :param pulumi.Input[str] update_time: Last update time. Synthetic field is populated automatically by CCFE.
        """
        ...
    @overload
    def __init__(__self__,
                 resource_name: str,
                 args: DomainArgs,
                 opts: Optional[pulumi.ResourceOptions] = None):
        """
        Creates a Microsoft AD Domain in a given project. Operation

        :param str resource_name: The name of the resource.
        :param DomainArgs args: The arguments to use to populate this resource's properties.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        ...
    def __init__(__self__, resource_name: str, *args, **kwargs):
        resource_args, opts = _utilities.get_resource_args_opts(DomainArgs, pulumi.ResourceOptions, *args, **kwargs)
        if resource_args is not None:
            __self__._internal_init(resource_name, opts, **resource_args.__dict__)
        else:
            __self__._internal_init(resource_name, *args, **kwargs)

    def _internal_init(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 audit_logs_enabled: Optional[pulumi.Input[bool]] = None,
                 authorized_networks: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]] = None,
                 create_time: Optional[pulumi.Input[str]] = None,
                 domain_name: Optional[pulumi.Input[str]] = None,
                 fqdn: Optional[pulumi.Input[str]] = None,
                 labels: Optional[pulumi.Input[Mapping[str, pulumi.Input[str]]]] = None,
                 locations: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]] = None,
                 managed_identities_admin_name: Optional[pulumi.Input[str]] = None,
                 name: Optional[pulumi.Input[str]] = None,
                 project: Optional[pulumi.Input[str]] = None,
                 reserved_ip_range: Optional[pulumi.Input[str]] = None,
                 state: Optional[pulumi.Input[str]] = None,
                 status_message: Optional[pulumi.Input[str]] = None,
                 trusts: Optional[pulumi.Input[Sequence[pulumi.Input[pulumi.InputType['TrustArgs']]]]] = None,
                 update_time: Optional[pulumi.Input[str]] = None,
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
            __props__ = DomainArgs.__new__(DomainArgs)

            __props__.__dict__["audit_logs_enabled"] = audit_logs_enabled
            __props__.__dict__["authorized_networks"] = authorized_networks
            __props__.__dict__["create_time"] = create_time
            __props__.__dict__["domain_name"] = domain_name
            __props__.__dict__["fqdn"] = fqdn
            __props__.__dict__["labels"] = labels
            __props__.__dict__["locations"] = locations
            __props__.__dict__["managed_identities_admin_name"] = managed_identities_admin_name
            __props__.__dict__["name"] = name
            if project is None and not opts.urn:
                raise TypeError("Missing required property 'project'")
            __props__.__dict__["project"] = project
            __props__.__dict__["reserved_ip_range"] = reserved_ip_range
            __props__.__dict__["state"] = state
            __props__.__dict__["status_message"] = status_message
            __props__.__dict__["trusts"] = trusts
            __props__.__dict__["update_time"] = update_time
        super(Domain, __self__).__init__(
            'google-native:managedidentities/v1alpha1:Domain',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name: str,
            id: pulumi.Input[str],
            opts: Optional[pulumi.ResourceOptions] = None) -> 'Domain':
        """
        Get an existing Domain resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param pulumi.Input[str] id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = DomainArgs.__new__(DomainArgs)

        __props__.__dict__["audit_logs_enabled"] = None
        __props__.__dict__["authorized_networks"] = None
        __props__.__dict__["create_time"] = None
        __props__.__dict__["fqdn"] = None
        __props__.__dict__["labels"] = None
        __props__.__dict__["locations"] = None
        __props__.__dict__["managed_identities_admin_name"] = None
        __props__.__dict__["name"] = None
        __props__.__dict__["reserved_ip_range"] = None
        __props__.__dict__["state"] = None
        __props__.__dict__["status_message"] = None
        __props__.__dict__["trusts"] = None
        __props__.__dict__["update_time"] = None
        return Domain(resource_name, opts=opts, __props__=__props__)

    @property
    @pulumi.getter(name="auditLogsEnabled")
    def audit_logs_enabled(self) -> pulumi.Output[bool]:
        """
        Optional. Configuration for audit logs. True if audit logs are enabled, else false. Default is audit logs disabled.
        """
        return pulumi.get(self, "audit_logs_enabled")

    @property
    @pulumi.getter(name="authorizedNetworks")
    def authorized_networks(self) -> pulumi.Output[Sequence[str]]:
        """
        Optional. The full names of the Google Compute Engine [networks](/compute/docs/networks-and-firewalls#networks) to which the instance is connected. Network can be added using UpdateDomain later. Domain is only available on network part of authorized_networks. Caller needs to make sure that CIDR subnets do not overlap between networks, else domain creation will fail.
        """
        return pulumi.get(self, "authorized_networks")

    @property
    @pulumi.getter(name="createTime")
    def create_time(self) -> pulumi.Output[str]:
        """
        The time the instance was created. Synthetic field is populated automatically by CCFE. go/ccfe-synthetic-field-user-guide
        """
        return pulumi.get(self, "create_time")

    @property
    @pulumi.getter
    def fqdn(self) -> pulumi.Output[str]:
        """
        Fully-qualified domain name of the exposed domain used by clients to connect to the service. Similar to what would be chosen for an Active Directory that is set up on an internal network.
        """
        return pulumi.get(self, "fqdn")

    @property
    @pulumi.getter
    def labels(self) -> pulumi.Output[Mapping[str, str]]:
        """
        Optional. Resource labels to represent user provided metadata
        """
        return pulumi.get(self, "labels")

    @property
    @pulumi.getter
    def locations(self) -> pulumi.Output[Sequence[str]]:
        """
        Required. Locations where domain needs to be provisioned. regions e.g. us-west1 or us-east4 Service supports up to 4 locations at once. Each location will use a /26 block.
        """
        return pulumi.get(self, "locations")

    @property
    @pulumi.getter(name="managedIdentitiesAdminName")
    def managed_identities_admin_name(self) -> pulumi.Output[str]:
        """
        Optional. Name of customer-visible admin used to perform Active Directory operations. If not specified `setupadmin` would be used.
        """
        return pulumi.get(self, "managed_identities_admin_name")

    @property
    @pulumi.getter
    def name(self) -> pulumi.Output[str]:
        """
        Unique name of the domain in this scope including projects and location using the form: `projects/{project_id}/locations/global/domains/{domain_name}`.
        """
        return pulumi.get(self, "name")

    @property
    @pulumi.getter(name="reservedIpRange")
    def reserved_ip_range(self) -> pulumi.Output[str]:
        """
        Required. The CIDR range of internal addresses that are reserved for this domain. Reserved networks must be /24 or larger. Ranges must be unique and non-overlapping with existing subnets in [Domain].[authorized_networks].
        """
        return pulumi.get(self, "reserved_ip_range")

    @property
    @pulumi.getter
    def state(self) -> pulumi.Output[str]:
        """
        The current state of this domain.
        """
        return pulumi.get(self, "state")

    @property
    @pulumi.getter(name="statusMessage")
    def status_message(self) -> pulumi.Output[str]:
        """
        Additional information about the current status of this domain, if available.
        """
        return pulumi.get(self, "status_message")

    @property
    @pulumi.getter
    def trusts(self) -> pulumi.Output[Sequence['outputs.TrustResponse']]:
        """
        The current trusts associated with the domain.
        """
        return pulumi.get(self, "trusts")

    @property
    @pulumi.getter(name="updateTime")
    def update_time(self) -> pulumi.Output[str]:
        """
        Last update time. Synthetic field is populated automatically by CCFE.
        """
        return pulumi.get(self, "update_time")

