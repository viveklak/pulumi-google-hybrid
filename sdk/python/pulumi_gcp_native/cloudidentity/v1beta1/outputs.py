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
    'AndroidAttributesResponse',
    'DynamicGroupMetadataResponse',
    'DynamicGroupQueryResponse',
    'DynamicGroupStatusResponse',
    'EntityKeyResponse',
    'ExpiryDetailResponse',
    'MembershipRoleResponse',
]

@pulumi.output_type
class AndroidAttributesResponse(dict):
    """
    Resource representing the Android specific attributes of a Device.
    """
    @staticmethod
    def __key_warning(key: str):
        suggest = None
        if key == "enabledUnknownSources":
            suggest = "enabled_unknown_sources"
        elif key == "ownerProfileAccount":
            suggest = "owner_profile_account"
        elif key == "ownershipPrivilege":
            suggest = "ownership_privilege"
        elif key == "supportsWorkProfile":
            suggest = "supports_work_profile"

        if suggest:
            pulumi.log.warn(f"Key '{key}' not found in AndroidAttributesResponse. Access the value via the '{suggest}' property getter instead.")

    def __getitem__(self, key: str) -> Any:
        AndroidAttributesResponse.__key_warning(key)
        return super().__getitem__(key)

    def get(self, key: str, default = None) -> Any:
        AndroidAttributesResponse.__key_warning(key)
        return super().get(key, default)

    def __init__(__self__, *,
                 enabled_unknown_sources: bool,
                 owner_profile_account: bool,
                 ownership_privilege: str,
                 supports_work_profile: bool):
        """
        Resource representing the Android specific attributes of a Device.
        :param bool enabled_unknown_sources: Whether applications from unknown sources can be installed on device.
        :param bool owner_profile_account: Whether this account is on an owner/primary profile. For phones, only true for owner profiles. Android 4+ devices can have secondary or restricted user profiles.
        :param str ownership_privilege: Ownership privileges on device.
        :param bool supports_work_profile: Whether device supports Android work profiles. If false, this service will not block access to corp data even if an administrator turns on the "Enforce Work Profile" policy.
        """
        pulumi.set(__self__, "enabled_unknown_sources", enabled_unknown_sources)
        pulumi.set(__self__, "owner_profile_account", owner_profile_account)
        pulumi.set(__self__, "ownership_privilege", ownership_privilege)
        pulumi.set(__self__, "supports_work_profile", supports_work_profile)

    @property
    @pulumi.getter(name="enabledUnknownSources")
    def enabled_unknown_sources(self) -> bool:
        """
        Whether applications from unknown sources can be installed on device.
        """
        return pulumi.get(self, "enabled_unknown_sources")

    @property
    @pulumi.getter(name="ownerProfileAccount")
    def owner_profile_account(self) -> bool:
        """
        Whether this account is on an owner/primary profile. For phones, only true for owner profiles. Android 4+ devices can have secondary or restricted user profiles.
        """
        return pulumi.get(self, "owner_profile_account")

    @property
    @pulumi.getter(name="ownershipPrivilege")
    def ownership_privilege(self) -> str:
        """
        Ownership privileges on device.
        """
        return pulumi.get(self, "ownership_privilege")

    @property
    @pulumi.getter(name="supportsWorkProfile")
    def supports_work_profile(self) -> bool:
        """
        Whether device supports Android work profiles. If false, this service will not block access to corp data even if an administrator turns on the "Enforce Work Profile" policy.
        """
        return pulumi.get(self, "supports_work_profile")


@pulumi.output_type
class DynamicGroupMetadataResponse(dict):
    """
    Dynamic group metadata like queries and status.
    """
    def __init__(__self__, *,
                 queries: Sequence['outputs.DynamicGroupQueryResponse'],
                 status: 'outputs.DynamicGroupStatusResponse'):
        """
        Dynamic group metadata like queries and status.
        :param Sequence['DynamicGroupQueryResponseArgs'] queries: Memberships will be the union of all queries. Only one entry with USER resource is currently supported. Customers can create up to 100 dynamic groups.
        :param 'DynamicGroupStatusResponseArgs' status: Status of the dynamic group.
        """
        pulumi.set(__self__, "queries", queries)
        pulumi.set(__self__, "status", status)

    @property
    @pulumi.getter
    def queries(self) -> Sequence['outputs.DynamicGroupQueryResponse']:
        """
        Memberships will be the union of all queries. Only one entry with USER resource is currently supported. Customers can create up to 100 dynamic groups.
        """
        return pulumi.get(self, "queries")

    @property
    @pulumi.getter
    def status(self) -> 'outputs.DynamicGroupStatusResponse':
        """
        Status of the dynamic group.
        """
        return pulumi.get(self, "status")


@pulumi.output_type
class DynamicGroupQueryResponse(dict):
    """
    Defines a query on a resource.
    """
    @staticmethod
    def __key_warning(key: str):
        suggest = None
        if key == "resourceType":
            suggest = "resource_type"

        if suggest:
            pulumi.log.warn(f"Key '{key}' not found in DynamicGroupQueryResponse. Access the value via the '{suggest}' property getter instead.")

    def __getitem__(self, key: str) -> Any:
        DynamicGroupQueryResponse.__key_warning(key)
        return super().__getitem__(key)

    def get(self, key: str, default = None) -> Any:
        DynamicGroupQueryResponse.__key_warning(key)
        return super().get(key, default)

    def __init__(__self__, *,
                 query: str,
                 resource_type: str):
        """
        Defines a query on a resource.
        :param str query: Query that determines the memberships of the dynamic group. Examples: All users with at least one `organizations.department` of engineering. `user.organizations.exists(org, org.department=='engineering')` All users with at least one location that has `area` of `foo` and `building_id` of `bar`. `user.locations.exists(loc, loc.area=='foo' && loc.building_id=='bar')`
        """
        pulumi.set(__self__, "query", query)
        pulumi.set(__self__, "resource_type", resource_type)

    @property
    @pulumi.getter
    def query(self) -> str:
        """
        Query that determines the memberships of the dynamic group. Examples: All users with at least one `organizations.department` of engineering. `user.organizations.exists(org, org.department=='engineering')` All users with at least one location that has `area` of `foo` and `building_id` of `bar`. `user.locations.exists(loc, loc.area=='foo' && loc.building_id=='bar')`
        """
        return pulumi.get(self, "query")

    @property
    @pulumi.getter(name="resourceType")
    def resource_type(self) -> str:
        return pulumi.get(self, "resource_type")


@pulumi.output_type
class DynamicGroupStatusResponse(dict):
    """
    The current status of a dynamic group along with timestamp.
    """
    @staticmethod
    def __key_warning(key: str):
        suggest = None
        if key == "statusTime":
            suggest = "status_time"

        if suggest:
            pulumi.log.warn(f"Key '{key}' not found in DynamicGroupStatusResponse. Access the value via the '{suggest}' property getter instead.")

    def __getitem__(self, key: str) -> Any:
        DynamicGroupStatusResponse.__key_warning(key)
        return super().__getitem__(key)

    def get(self, key: str, default = None) -> Any:
        DynamicGroupStatusResponse.__key_warning(key)
        return super().get(key, default)

    def __init__(__self__, *,
                 status: str,
                 status_time: str):
        """
        The current status of a dynamic group along with timestamp.
        :param str status: Status of the dynamic group.
        :param str status_time: The latest time at which the dynamic group is guaranteed to be in the given status. If status is `UP_TO_DATE`, the latest time at which the dynamic group was confirmed to be up-to-date. If status is `UPDATING_MEMBERSHIPS`, the time at which dynamic group was created.
        """
        pulumi.set(__self__, "status", status)
        pulumi.set(__self__, "status_time", status_time)

    @property
    @pulumi.getter
    def status(self) -> str:
        """
        Status of the dynamic group.
        """
        return pulumi.get(self, "status")

    @property
    @pulumi.getter(name="statusTime")
    def status_time(self) -> str:
        """
        The latest time at which the dynamic group is guaranteed to be in the given status. If status is `UP_TO_DATE`, the latest time at which the dynamic group was confirmed to be up-to-date. If status is `UPDATING_MEMBERSHIPS`, the time at which dynamic group was created.
        """
        return pulumi.get(self, "status_time")


@pulumi.output_type
class EntityKeyResponse(dict):
    """
    A unique identifier for an entity in the Cloud Identity Groups API. An entity can represent either a group with an optional `namespace` or a user without a `namespace`. The combination of `id` and `namespace` must be unique; however, the same `id` can be used with different `namespace`s.
    """
    def __init__(__self__, *,
                 namespace: str):
        """
        A unique identifier for an entity in the Cloud Identity Groups API. An entity can represent either a group with an optional `namespace` or a user without a `namespace`. The combination of `id` and `namespace` must be unique; however, the same `id` can be used with different `namespace`s.
        :param str namespace: The namespace in which the entity exists. If not specified, the `EntityKey` represents a Google-managed entity such as a Google user or a Google Group. If specified, the `EntityKey` represents an external-identity-mapped group. The namespace must correspond to an identity source created in Admin Console and must be in the form of `identitysources/{identity_source_id}.
        """
        pulumi.set(__self__, "namespace", namespace)

    @property
    @pulumi.getter
    def namespace(self) -> str:
        """
        The namespace in which the entity exists. If not specified, the `EntityKey` represents a Google-managed entity such as a Google user or a Google Group. If specified, the `EntityKey` represents an external-identity-mapped group. The namespace must correspond to an identity source created in Admin Console and must be in the form of `identitysources/{identity_source_id}.
        """
        return pulumi.get(self, "namespace")


@pulumi.output_type
class ExpiryDetailResponse(dict):
    """
    The `MembershipRole` expiry details.
    """
    @staticmethod
    def __key_warning(key: str):
        suggest = None
        if key == "expireTime":
            suggest = "expire_time"

        if suggest:
            pulumi.log.warn(f"Key '{key}' not found in ExpiryDetailResponse. Access the value via the '{suggest}' property getter instead.")

    def __getitem__(self, key: str) -> Any:
        ExpiryDetailResponse.__key_warning(key)
        return super().__getitem__(key)

    def get(self, key: str, default = None) -> Any:
        ExpiryDetailResponse.__key_warning(key)
        return super().get(key, default)

    def __init__(__self__, *,
                 expire_time: str):
        """
        The `MembershipRole` expiry details.
        :param str expire_time: The time at which the `MembershipRole` will expire.
        """
        pulumi.set(__self__, "expire_time", expire_time)

    @property
    @pulumi.getter(name="expireTime")
    def expire_time(self) -> str:
        """
        The time at which the `MembershipRole` will expire.
        """
        return pulumi.get(self, "expire_time")


@pulumi.output_type
class MembershipRoleResponse(dict):
    """
    A membership role within the Cloud Identity Groups API. A `MembershipRole` defines the privileges granted to a `Membership`.
    """
    @staticmethod
    def __key_warning(key: str):
        suggest = None
        if key == "expiryDetail":
            suggest = "expiry_detail"

        if suggest:
            pulumi.log.warn(f"Key '{key}' not found in MembershipRoleResponse. Access the value via the '{suggest}' property getter instead.")

    def __getitem__(self, key: str) -> Any:
        MembershipRoleResponse.__key_warning(key)
        return super().__getitem__(key)

    def get(self, key: str, default = None) -> Any:
        MembershipRoleResponse.__key_warning(key)
        return super().get(key, default)

    def __init__(__self__, *,
                 expiry_detail: 'outputs.ExpiryDetailResponse',
                 name: str):
        """
        A membership role within the Cloud Identity Groups API. A `MembershipRole` defines the privileges granted to a `Membership`.
        :param 'ExpiryDetailResponseArgs' expiry_detail: The expiry details of the `MembershipRole`. Expiry details are only supported for `MEMBER` `MembershipRoles`. May be set if `name` is `MEMBER`. Must not be set if `name` is any other value.
        :param str name: The name of the `MembershipRole`. Must be one of `OWNER`, `MANAGER`, `MEMBER`.
        """
        pulumi.set(__self__, "expiry_detail", expiry_detail)
        pulumi.set(__self__, "name", name)

    @property
    @pulumi.getter(name="expiryDetail")
    def expiry_detail(self) -> 'outputs.ExpiryDetailResponse':
        """
        The expiry details of the `MembershipRole`. Expiry details are only supported for `MEMBER` `MembershipRoles`. May be set if `name` is `MEMBER`. Must not be set if `name` is any other value.
        """
        return pulumi.get(self, "expiry_detail")

    @property
    @pulumi.getter
    def name(self) -> str:
        """
        The name of the `MembershipRole`. Must be one of `OWNER`, `MANAGER`, `MEMBER`.
        """
        return pulumi.get(self, "name")


