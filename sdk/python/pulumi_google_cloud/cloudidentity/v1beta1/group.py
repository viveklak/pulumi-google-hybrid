# coding=utf-8
# *** WARNING: this file was generated by the Pulumi SDK Generator. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import warnings
import pulumi
import pulumi.runtime
from typing import Any, Mapping, Optional, Sequence, Union
from ... import _utilities, _tables
from ._inputs import *

__all__ = ['Group']


class Group(pulumi.CustomResource):
    def __init__(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 additional_group_keys: Optional[pulumi.Input[Sequence[pulumi.Input[pulumi.InputType['EntityKeyArgs']]]]] = None,
                 create_time: Optional[pulumi.Input[str]] = None,
                 description: Optional[pulumi.Input[str]] = None,
                 display_name: Optional[pulumi.Input[str]] = None,
                 dynamic_group_metadata: Optional[pulumi.Input[pulumi.InputType['DynamicGroupMetadataArgs']]] = None,
                 group_key: Optional[pulumi.Input[pulumi.InputType['EntityKeyArgs']]] = None,
                 groups_id: Optional[pulumi.Input[str]] = None,
                 labels: Optional[pulumi.Input[Mapping[str, pulumi.Input[str]]]] = None,
                 name: Optional[pulumi.Input[str]] = None,
                 parent: Optional[pulumi.Input[str]] = None,
                 update_time: Optional[pulumi.Input[str]] = None,
                 __props__=None,
                 __name__=None,
                 __opts__=None):
        """
        Creates a `Group`.

        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[Sequence[pulumi.Input[pulumi.InputType['EntityKeyArgs']]]] additional_group_keys: Additional entity key aliases for a Group.
        :param pulumi.Input[str] create_time: Output only. The time when the `Group` was created.
        :param pulumi.Input[str] description: An extended description to help users determine the purpose of a `Group`. Must not be longer than 4,096 characters.
        :param pulumi.Input[str] display_name: The display name of the `Group`.
        :param pulumi.Input[pulumi.InputType['DynamicGroupMetadataArgs']] dynamic_group_metadata: Optional. Dynamic group metadata like queries and status.
        :param pulumi.Input[pulumi.InputType['EntityKeyArgs']] group_key: Required. Immutable. The `EntityKey` of the `Group`.
        :param pulumi.Input[Mapping[str, pulumi.Input[str]]] labels: Required. One or more label entries that apply to the Group. Currently supported labels contain a key with an empty value. Google Groups are the default type of group and have a label with a key of `cloudidentity.googleapis.com/groups.discussion_forum` and an empty value. Existing Google Groups can have an additional label with a key of `cloudidentity.googleapis.com/groups.security` and an empty value added to them. **This is an immutable change and the security label cannot be removed once added.** Dynamic groups have a label with a key of `cloudidentity.googleapis.com/groups.dynamic`. Identity-mapped groups for Cloud Search have a label with a key of `system/groups/external` and an empty value. Examples: {"cloudidentity.googleapis.com/groups.discussion_forum": ""} or {"system/groups/external": ""}.
        :param pulumi.Input[str] name: Output only. The [resource name](https://cloud.google.com/apis/design/resource_names) of the `Group`. Shall be of the form `groups/{group_id}`.
        :param pulumi.Input[str] parent: Required. Immutable. The resource name of the entity under which this `Group` resides in the Cloud Identity resource hierarchy. Must be of the form `identitysources/{identity_source_id}` for external- identity-mapped groups or `customers/{customer_id}` for Google Groups.
        :param pulumi.Input[str] update_time: Output only. The time when the `Group` was last updated.
        """
        if __name__ is not None:
            warnings.warn("explicit use of __name__ is deprecated", DeprecationWarning)
            resource_name = __name__
        if __opts__ is not None:
            warnings.warn("explicit use of __opts__ is deprecated, use 'opts' instead", DeprecationWarning)
            opts = __opts__
        if opts is None:
            opts = pulumi.ResourceOptions()
        if not isinstance(opts, pulumi.ResourceOptions):
            raise TypeError('Expected resource options to be a ResourceOptions instance')
        if opts.version is None:
            opts.version = _utilities.get_version()
        if opts.id is None:
            if __props__ is not None:
                raise TypeError('__props__ is only valid when passed in combination with a valid opts.id to get an existing resource')
            __props__ = dict()

            __props__['additional_group_keys'] = additional_group_keys
            __props__['create_time'] = create_time
            __props__['description'] = description
            __props__['display_name'] = display_name
            __props__['dynamic_group_metadata'] = dynamic_group_metadata
            __props__['group_key'] = group_key
            if groups_id is None and not opts.urn:
                raise TypeError("Missing required property 'groups_id'")
            __props__['groups_id'] = groups_id
            __props__['labels'] = labels
            __props__['name'] = name
            __props__['parent'] = parent
            __props__['update_time'] = update_time
        super(Group, __self__).__init__(
            'google-cloud:cloudidentity/v1beta1:Group',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name: str,
            id: pulumi.Input[str],
            opts: Optional[pulumi.ResourceOptions] = None) -> 'Group':
        """
        Get an existing Group resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param pulumi.Input[str] id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = dict()

        return Group(resource_name, opts=opts, __props__=__props__)

    def translate_output_property(self, prop):
        return _tables.CAMEL_TO_SNAKE_CASE_TABLE.get(prop) or prop

    def translate_input_property(self, prop):
        return _tables.SNAKE_TO_CAMEL_CASE_TABLE.get(prop) or prop

