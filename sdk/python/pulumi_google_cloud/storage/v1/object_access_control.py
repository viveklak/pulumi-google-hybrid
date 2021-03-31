# coding=utf-8
# *** WARNING: this file was generated by the Pulumi SDK Generator. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import warnings
import pulumi
import pulumi.runtime
from typing import Any, Mapping, Optional, Sequence, Union
from ... import _utilities, _tables

__all__ = ['ObjectAccessControl']


class ObjectAccessControl(pulumi.CustomResource):
    def __init__(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 bucket: Optional[pulumi.Input[str]] = None,
                 domain: Optional[pulumi.Input[str]] = None,
                 email: Optional[pulumi.Input[str]] = None,
                 entity: Optional[pulumi.Input[str]] = None,
                 entity_id: Optional[pulumi.Input[str]] = None,
                 etag: Optional[pulumi.Input[str]] = None,
                 generation: Optional[pulumi.Input[str]] = None,
                 id: Optional[pulumi.Input[str]] = None,
                 kind: Optional[pulumi.Input[str]] = None,
                 object: Optional[pulumi.Input[str]] = None,
                 project_team: Optional[pulumi.Input[Mapping[str, pulumi.Input[str]]]] = None,
                 role: Optional[pulumi.Input[str]] = None,
                 self_link: Optional[pulumi.Input[str]] = None,
                 __props__=None,
                 __name__=None,
                 __opts__=None):
        """
        Creates a new ACL entry on the specified object.

        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] bucket: The name of the bucket.
        :param pulumi.Input[str] domain: The domain associated with the entity, if any.
        :param pulumi.Input[str] email: The email address associated with the entity, if any.
        :param pulumi.Input[str] entity: The entity holding the permission, in one of the following forms: 
               - user-userId 
               - user-email 
               - group-groupId 
               - group-email 
               - domain-domain 
               - project-team-projectId 
               - allUsers 
               - allAuthenticatedUsers Examples: 
               - The user liz@example.com would be user-liz@example.com. 
               - The group example@googlegroups.com would be group-example@googlegroups.com. 
               - To refer to all members of the Google Apps for Business domain example.com, the entity would be domain-example.com.
        :param pulumi.Input[str] entity_id: The ID for the entity, if any.
        :param pulumi.Input[str] etag: HTTP 1.1 Entity tag for the access-control entry.
        :param pulumi.Input[str] generation: The content generation of the object, if applied to an object.
        :param pulumi.Input[str] id: The ID of the access-control entry.
        :param pulumi.Input[str] kind: The kind of item this is. For object access control entries, this is always storage#objectAccessControl.
        :param pulumi.Input[str] object: The name of the object, if applied to an object.
        :param pulumi.Input[Mapping[str, pulumi.Input[str]]] project_team: The project team associated with the entity, if any.
        :param pulumi.Input[str] role: The access permission for the entity.
        :param pulumi.Input[str] self_link: The link to this access-control entry.
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

            if bucket is None and not opts.urn:
                raise TypeError("Missing required property 'bucket'")
            __props__['bucket'] = bucket
            __props__['domain'] = domain
            __props__['email'] = email
            if entity is None and not opts.urn:
                raise TypeError("Missing required property 'entity'")
            __props__['entity'] = entity
            __props__['entity_id'] = entity_id
            __props__['etag'] = etag
            __props__['generation'] = generation
            __props__['id'] = id
            __props__['kind'] = kind
            if object is None and not opts.urn:
                raise TypeError("Missing required property 'object'")
            __props__['object'] = object
            __props__['project_team'] = project_team
            __props__['role'] = role
            __props__['self_link'] = self_link
        super(ObjectAccessControl, __self__).__init__(
            'google-cloud:storage/v1:ObjectAccessControl',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name: str,
            id: pulumi.Input[str],
            opts: Optional[pulumi.ResourceOptions] = None) -> 'ObjectAccessControl':
        """
        Get an existing ObjectAccessControl resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param pulumi.Input[str] id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = dict()

        return ObjectAccessControl(resource_name, opts=opts, __props__=__props__)

    def translate_output_property(self, prop):
        return _tables.CAMEL_TO_SNAKE_CASE_TABLE.get(prop) or prop

    def translate_input_property(self, prop):
        return _tables.SNAKE_TO_CAMEL_CASE_TABLE.get(prop) or prop

