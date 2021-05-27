# coding=utf-8
# *** WARNING: this file was generated by the Pulumi SDK Generator. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import warnings
import pulumi
import pulumi.runtime
from typing import Any, Mapping, Optional, Sequence, Union, overload
from ... import _utilities

__all__ = ['OrganizationRoleArgs', 'OrganizationRole']

@pulumi.input_type
class OrganizationRoleArgs:
    def __init__(__self__, *,
                 organization_id: pulumi.Input[str],
                 deleted: Optional[pulumi.Input[bool]] = None,
                 description: Optional[pulumi.Input[str]] = None,
                 etag: Optional[pulumi.Input[str]] = None,
                 included_permissions: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]] = None,
                 name: Optional[pulumi.Input[str]] = None,
                 role_id: Optional[pulumi.Input[str]] = None,
                 stage: Optional[pulumi.Input[str]] = None,
                 title: Optional[pulumi.Input[str]] = None):
        """
        The set of arguments for constructing a OrganizationRole resource.
        :param pulumi.Input[bool] deleted: The current deleted state of the role. This field is read only. It will be ignored in calls to CreateRole and UpdateRole.
        :param pulumi.Input[str] description: Optional. A human-readable description for the role.
        :param pulumi.Input[str] etag: Used to perform a consistent read-modify-write.
        :param pulumi.Input[Sequence[pulumi.Input[str]]] included_permissions: The names of the permissions this role grants when bound in an IAM policy.
        :param pulumi.Input[str] name: The name of the role. When Role is used in CreateRole, the role name must not be set. When Role is used in output and other input such as UpdateRole, the role name is the complete path, e.g., roles/logging.viewer for predefined roles and organizations/{ORGANIZATION_ID}/roles/logging.viewer for custom roles.
        :param pulumi.Input[str] role_id: The role ID to use for this role. A role ID may contain alphanumeric characters, underscores (`_`), and periods (`.`). It must contain a minimum of 3 characters and a maximum of 64 characters.
        :param pulumi.Input[str] stage: The current launch stage of the role. If the `ALPHA` launch stage has been selected for a role, the `stage` field will not be included in the returned definition for the role.
        :param pulumi.Input[str] title: Optional. A human-readable title for the role. Typically this is limited to 100 UTF-8 bytes.
        """
        pulumi.set(__self__, "organization_id", organization_id)
        if deleted is not None:
            pulumi.set(__self__, "deleted", deleted)
        if description is not None:
            pulumi.set(__self__, "description", description)
        if etag is not None:
            pulumi.set(__self__, "etag", etag)
        if included_permissions is not None:
            pulumi.set(__self__, "included_permissions", included_permissions)
        if name is not None:
            pulumi.set(__self__, "name", name)
        if role_id is not None:
            pulumi.set(__self__, "role_id", role_id)
        if stage is not None:
            pulumi.set(__self__, "stage", stage)
        if title is not None:
            pulumi.set(__self__, "title", title)

    @property
    @pulumi.getter(name="organizationId")
    def organization_id(self) -> pulumi.Input[str]:
        return pulumi.get(self, "organization_id")

    @organization_id.setter
    def organization_id(self, value: pulumi.Input[str]):
        pulumi.set(self, "organization_id", value)

    @property
    @pulumi.getter
    def deleted(self) -> Optional[pulumi.Input[bool]]:
        """
        The current deleted state of the role. This field is read only. It will be ignored in calls to CreateRole and UpdateRole.
        """
        return pulumi.get(self, "deleted")

    @deleted.setter
    def deleted(self, value: Optional[pulumi.Input[bool]]):
        pulumi.set(self, "deleted", value)

    @property
    @pulumi.getter
    def description(self) -> Optional[pulumi.Input[str]]:
        """
        Optional. A human-readable description for the role.
        """
        return pulumi.get(self, "description")

    @description.setter
    def description(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "description", value)

    @property
    @pulumi.getter
    def etag(self) -> Optional[pulumi.Input[str]]:
        """
        Used to perform a consistent read-modify-write.
        """
        return pulumi.get(self, "etag")

    @etag.setter
    def etag(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "etag", value)

    @property
    @pulumi.getter(name="includedPermissions")
    def included_permissions(self) -> Optional[pulumi.Input[Sequence[pulumi.Input[str]]]]:
        """
        The names of the permissions this role grants when bound in an IAM policy.
        """
        return pulumi.get(self, "included_permissions")

    @included_permissions.setter
    def included_permissions(self, value: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]]):
        pulumi.set(self, "included_permissions", value)

    @property
    @pulumi.getter
    def name(self) -> Optional[pulumi.Input[str]]:
        """
        The name of the role. When Role is used in CreateRole, the role name must not be set. When Role is used in output and other input such as UpdateRole, the role name is the complete path, e.g., roles/logging.viewer for predefined roles and organizations/{ORGANIZATION_ID}/roles/logging.viewer for custom roles.
        """
        return pulumi.get(self, "name")

    @name.setter
    def name(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "name", value)

    @property
    @pulumi.getter(name="roleId")
    def role_id(self) -> Optional[pulumi.Input[str]]:
        """
        The role ID to use for this role. A role ID may contain alphanumeric characters, underscores (`_`), and periods (`.`). It must contain a minimum of 3 characters and a maximum of 64 characters.
        """
        return pulumi.get(self, "role_id")

    @role_id.setter
    def role_id(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "role_id", value)

    @property
    @pulumi.getter
    def stage(self) -> Optional[pulumi.Input[str]]:
        """
        The current launch stage of the role. If the `ALPHA` launch stage has been selected for a role, the `stage` field will not be included in the returned definition for the role.
        """
        return pulumi.get(self, "stage")

    @stage.setter
    def stage(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "stage", value)

    @property
    @pulumi.getter
    def title(self) -> Optional[pulumi.Input[str]]:
        """
        Optional. A human-readable title for the role. Typically this is limited to 100 UTF-8 bytes.
        """
        return pulumi.get(self, "title")

    @title.setter
    def title(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "title", value)


class OrganizationRole(pulumi.CustomResource):
    @overload
    def __init__(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 deleted: Optional[pulumi.Input[bool]] = None,
                 description: Optional[pulumi.Input[str]] = None,
                 etag: Optional[pulumi.Input[str]] = None,
                 included_permissions: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]] = None,
                 name: Optional[pulumi.Input[str]] = None,
                 organization_id: Optional[pulumi.Input[str]] = None,
                 role_id: Optional[pulumi.Input[str]] = None,
                 stage: Optional[pulumi.Input[str]] = None,
                 title: Optional[pulumi.Input[str]] = None,
                 __props__=None):
        """
        Creates a new custom Role.

        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[bool] deleted: The current deleted state of the role. This field is read only. It will be ignored in calls to CreateRole and UpdateRole.
        :param pulumi.Input[str] description: Optional. A human-readable description for the role.
        :param pulumi.Input[str] etag: Used to perform a consistent read-modify-write.
        :param pulumi.Input[Sequence[pulumi.Input[str]]] included_permissions: The names of the permissions this role grants when bound in an IAM policy.
        :param pulumi.Input[str] name: The name of the role. When Role is used in CreateRole, the role name must not be set. When Role is used in output and other input such as UpdateRole, the role name is the complete path, e.g., roles/logging.viewer for predefined roles and organizations/{ORGANIZATION_ID}/roles/logging.viewer for custom roles.
        :param pulumi.Input[str] role_id: The role ID to use for this role. A role ID may contain alphanumeric characters, underscores (`_`), and periods (`.`). It must contain a minimum of 3 characters and a maximum of 64 characters.
        :param pulumi.Input[str] stage: The current launch stage of the role. If the `ALPHA` launch stage has been selected for a role, the `stage` field will not be included in the returned definition for the role.
        :param pulumi.Input[str] title: Optional. A human-readable title for the role. Typically this is limited to 100 UTF-8 bytes.
        """
        ...
    @overload
    def __init__(__self__,
                 resource_name: str,
                 args: OrganizationRoleArgs,
                 opts: Optional[pulumi.ResourceOptions] = None):
        """
        Creates a new custom Role.

        :param str resource_name: The name of the resource.
        :param OrganizationRoleArgs args: The arguments to use to populate this resource's properties.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        ...
    def __init__(__self__, resource_name: str, *args, **kwargs):
        resource_args, opts = _utilities.get_resource_args_opts(OrganizationRoleArgs, pulumi.ResourceOptions, *args, **kwargs)
        if resource_args is not None:
            __self__._internal_init(resource_name, opts, **resource_args.__dict__)
        else:
            __self__._internal_init(resource_name, *args, **kwargs)

    def _internal_init(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 deleted: Optional[pulumi.Input[bool]] = None,
                 description: Optional[pulumi.Input[str]] = None,
                 etag: Optional[pulumi.Input[str]] = None,
                 included_permissions: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]] = None,
                 name: Optional[pulumi.Input[str]] = None,
                 organization_id: Optional[pulumi.Input[str]] = None,
                 role_id: Optional[pulumi.Input[str]] = None,
                 stage: Optional[pulumi.Input[str]] = None,
                 title: Optional[pulumi.Input[str]] = None,
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
            __props__ = OrganizationRoleArgs.__new__(OrganizationRoleArgs)

            __props__.__dict__["deleted"] = deleted
            __props__.__dict__["description"] = description
            __props__.__dict__["etag"] = etag
            __props__.__dict__["included_permissions"] = included_permissions
            __props__.__dict__["name"] = name
            if organization_id is None and not opts.urn:
                raise TypeError("Missing required property 'organization_id'")
            __props__.__dict__["organization_id"] = organization_id
            __props__.__dict__["role_id"] = role_id
            __props__.__dict__["stage"] = stage
            __props__.__dict__["title"] = title
        super(OrganizationRole, __self__).__init__(
            'google-native:iam/v1:OrganizationRole',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name: str,
            id: pulumi.Input[str],
            opts: Optional[pulumi.ResourceOptions] = None) -> 'OrganizationRole':
        """
        Get an existing OrganizationRole resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param pulumi.Input[str] id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = OrganizationRoleArgs.__new__(OrganizationRoleArgs)

        __props__.__dict__["deleted"] = None
        __props__.__dict__["description"] = None
        __props__.__dict__["etag"] = None
        __props__.__dict__["included_permissions"] = None
        __props__.__dict__["name"] = None
        __props__.__dict__["stage"] = None
        __props__.__dict__["title"] = None
        return OrganizationRole(resource_name, opts=opts, __props__=__props__)

    @property
    @pulumi.getter
    def deleted(self) -> pulumi.Output[bool]:
        """
        The current deleted state of the role. This field is read only. It will be ignored in calls to CreateRole and UpdateRole.
        """
        return pulumi.get(self, "deleted")

    @property
    @pulumi.getter
    def description(self) -> pulumi.Output[str]:
        """
        Optional. A human-readable description for the role.
        """
        return pulumi.get(self, "description")

    @property
    @pulumi.getter
    def etag(self) -> pulumi.Output[str]:
        """
        Used to perform a consistent read-modify-write.
        """
        return pulumi.get(self, "etag")

    @property
    @pulumi.getter(name="includedPermissions")
    def included_permissions(self) -> pulumi.Output[Sequence[str]]:
        """
        The names of the permissions this role grants when bound in an IAM policy.
        """
        return pulumi.get(self, "included_permissions")

    @property
    @pulumi.getter
    def name(self) -> pulumi.Output[str]:
        """
        The name of the role. When Role is used in CreateRole, the role name must not be set. When Role is used in output and other input such as UpdateRole, the role name is the complete path, e.g., roles/logging.viewer for predefined roles and organizations/{ORGANIZATION_ID}/roles/logging.viewer for custom roles.
        """
        return pulumi.get(self, "name")

    @property
    @pulumi.getter
    def stage(self) -> pulumi.Output[str]:
        """
        The current launch stage of the role. If the `ALPHA` launch stage has been selected for a role, the `stage` field will not be included in the returned definition for the role.
        """
        return pulumi.get(self, "stage")

    @property
    @pulumi.getter
    def title(self) -> pulumi.Output[str]:
        """
        Optional. A human-readable title for the role. Typically this is limited to 100 UTF-8 bytes.
        """
        return pulumi.get(self, "title")

