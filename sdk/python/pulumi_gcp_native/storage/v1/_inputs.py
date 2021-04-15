# coding=utf-8
# *** WARNING: this file was generated by the Pulumi SDK Generator. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import warnings
import pulumi
import pulumi.runtime
from typing import Any, Mapping, Optional, Sequence, Union, overload
from ... import _utilities

__all__ = [
    'BucketAccessControlArgs',
    'ObjectAccessControlArgs',
]

@pulumi.input_type
class BucketAccessControlArgs:
    def __init__(__self__, *,
                 bucket: Optional[pulumi.Input[str]] = None,
                 domain: Optional[pulumi.Input[str]] = None,
                 email: Optional[pulumi.Input[str]] = None,
                 entity: Optional[pulumi.Input[str]] = None,
                 entity_id: Optional[pulumi.Input[str]] = None,
                 etag: Optional[pulumi.Input[str]] = None,
                 id: Optional[pulumi.Input[str]] = None,
                 kind: Optional[pulumi.Input[str]] = None,
                 project_team: Optional[pulumi.Input[Mapping[str, pulumi.Input[str]]]] = None,
                 role: Optional[pulumi.Input[str]] = None,
                 self_link: Optional[pulumi.Input[str]] = None):
        """
        An access-control entry.
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
        :param pulumi.Input[str] id: The ID of the access-control entry.
        :param pulumi.Input[str] kind: The kind of item this is. For bucket access control entries, this is always storage#bucketAccessControl.
        :param pulumi.Input[Mapping[str, pulumi.Input[str]]] project_team: The project team associated with the entity, if any.
        :param pulumi.Input[str] role: The access permission for the entity.
        :param pulumi.Input[str] self_link: The link to this access-control entry.
        """
        if bucket is not None:
            pulumi.set(__self__, "bucket", bucket)
        if domain is not None:
            pulumi.set(__self__, "domain", domain)
        if email is not None:
            pulumi.set(__self__, "email", email)
        if entity is not None:
            pulumi.set(__self__, "entity", entity)
        if entity_id is not None:
            pulumi.set(__self__, "entity_id", entity_id)
        if etag is not None:
            pulumi.set(__self__, "etag", etag)
        if id is not None:
            pulumi.set(__self__, "id", id)
        if kind is not None:
            pulumi.set(__self__, "kind", kind)
        if project_team is not None:
            pulumi.set(__self__, "project_team", project_team)
        if role is not None:
            pulumi.set(__self__, "role", role)
        if self_link is not None:
            pulumi.set(__self__, "self_link", self_link)

    @property
    @pulumi.getter
    def bucket(self) -> Optional[pulumi.Input[str]]:
        """
        The name of the bucket.
        """
        return pulumi.get(self, "bucket")

    @bucket.setter
    def bucket(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "bucket", value)

    @property
    @pulumi.getter
    def domain(self) -> Optional[pulumi.Input[str]]:
        """
        The domain associated with the entity, if any.
        """
        return pulumi.get(self, "domain")

    @domain.setter
    def domain(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "domain", value)

    @property
    @pulumi.getter
    def email(self) -> Optional[pulumi.Input[str]]:
        """
        The email address associated with the entity, if any.
        """
        return pulumi.get(self, "email")

    @email.setter
    def email(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "email", value)

    @property
    @pulumi.getter
    def entity(self) -> Optional[pulumi.Input[str]]:
        """
        The entity holding the permission, in one of the following forms: 
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
        """
        return pulumi.get(self, "entity")

    @entity.setter
    def entity(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "entity", value)

    @property
    @pulumi.getter(name="entityId")
    def entity_id(self) -> Optional[pulumi.Input[str]]:
        """
        The ID for the entity, if any.
        """
        return pulumi.get(self, "entity_id")

    @entity_id.setter
    def entity_id(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "entity_id", value)

    @property
    @pulumi.getter
    def etag(self) -> Optional[pulumi.Input[str]]:
        """
        HTTP 1.1 Entity tag for the access-control entry.
        """
        return pulumi.get(self, "etag")

    @etag.setter
    def etag(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "etag", value)

    @property
    @pulumi.getter
    def id(self) -> Optional[pulumi.Input[str]]:
        """
        The ID of the access-control entry.
        """
        return pulumi.get(self, "id")

    @id.setter
    def id(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "id", value)

    @property
    @pulumi.getter
    def kind(self) -> Optional[pulumi.Input[str]]:
        """
        The kind of item this is. For bucket access control entries, this is always storage#bucketAccessControl.
        """
        return pulumi.get(self, "kind")

    @kind.setter
    def kind(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "kind", value)

    @property
    @pulumi.getter(name="projectTeam")
    def project_team(self) -> Optional[pulumi.Input[Mapping[str, pulumi.Input[str]]]]:
        """
        The project team associated with the entity, if any.
        """
        return pulumi.get(self, "project_team")

    @project_team.setter
    def project_team(self, value: Optional[pulumi.Input[Mapping[str, pulumi.Input[str]]]]):
        pulumi.set(self, "project_team", value)

    @property
    @pulumi.getter
    def role(self) -> Optional[pulumi.Input[str]]:
        """
        The access permission for the entity.
        """
        return pulumi.get(self, "role")

    @role.setter
    def role(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "role", value)

    @property
    @pulumi.getter(name="selfLink")
    def self_link(self) -> Optional[pulumi.Input[str]]:
        """
        The link to this access-control entry.
        """
        return pulumi.get(self, "self_link")

    @self_link.setter
    def self_link(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "self_link", value)


@pulumi.input_type
class ObjectAccessControlArgs:
    def __init__(__self__, *,
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
                 self_link: Optional[pulumi.Input[str]] = None):
        """
        An access-control entry.
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
        if bucket is not None:
            pulumi.set(__self__, "bucket", bucket)
        if domain is not None:
            pulumi.set(__self__, "domain", domain)
        if email is not None:
            pulumi.set(__self__, "email", email)
        if entity is not None:
            pulumi.set(__self__, "entity", entity)
        if entity_id is not None:
            pulumi.set(__self__, "entity_id", entity_id)
        if etag is not None:
            pulumi.set(__self__, "etag", etag)
        if generation is not None:
            pulumi.set(__self__, "generation", generation)
        if id is not None:
            pulumi.set(__self__, "id", id)
        if kind is not None:
            pulumi.set(__self__, "kind", kind)
        if object is not None:
            pulumi.set(__self__, "object", object)
        if project_team is not None:
            pulumi.set(__self__, "project_team", project_team)
        if role is not None:
            pulumi.set(__self__, "role", role)
        if self_link is not None:
            pulumi.set(__self__, "self_link", self_link)

    @property
    @pulumi.getter
    def bucket(self) -> Optional[pulumi.Input[str]]:
        """
        The name of the bucket.
        """
        return pulumi.get(self, "bucket")

    @bucket.setter
    def bucket(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "bucket", value)

    @property
    @pulumi.getter
    def domain(self) -> Optional[pulumi.Input[str]]:
        """
        The domain associated with the entity, if any.
        """
        return pulumi.get(self, "domain")

    @domain.setter
    def domain(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "domain", value)

    @property
    @pulumi.getter
    def email(self) -> Optional[pulumi.Input[str]]:
        """
        The email address associated with the entity, if any.
        """
        return pulumi.get(self, "email")

    @email.setter
    def email(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "email", value)

    @property
    @pulumi.getter
    def entity(self) -> Optional[pulumi.Input[str]]:
        """
        The entity holding the permission, in one of the following forms: 
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
        """
        return pulumi.get(self, "entity")

    @entity.setter
    def entity(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "entity", value)

    @property
    @pulumi.getter(name="entityId")
    def entity_id(self) -> Optional[pulumi.Input[str]]:
        """
        The ID for the entity, if any.
        """
        return pulumi.get(self, "entity_id")

    @entity_id.setter
    def entity_id(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "entity_id", value)

    @property
    @pulumi.getter
    def etag(self) -> Optional[pulumi.Input[str]]:
        """
        HTTP 1.1 Entity tag for the access-control entry.
        """
        return pulumi.get(self, "etag")

    @etag.setter
    def etag(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "etag", value)

    @property
    @pulumi.getter
    def generation(self) -> Optional[pulumi.Input[str]]:
        """
        The content generation of the object, if applied to an object.
        """
        return pulumi.get(self, "generation")

    @generation.setter
    def generation(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "generation", value)

    @property
    @pulumi.getter
    def id(self) -> Optional[pulumi.Input[str]]:
        """
        The ID of the access-control entry.
        """
        return pulumi.get(self, "id")

    @id.setter
    def id(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "id", value)

    @property
    @pulumi.getter
    def kind(self) -> Optional[pulumi.Input[str]]:
        """
        The kind of item this is. For object access control entries, this is always storage#objectAccessControl.
        """
        return pulumi.get(self, "kind")

    @kind.setter
    def kind(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "kind", value)

    @property
    @pulumi.getter
    def object(self) -> Optional[pulumi.Input[str]]:
        """
        The name of the object, if applied to an object.
        """
        return pulumi.get(self, "object")

    @object.setter
    def object(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "object", value)

    @property
    @pulumi.getter(name="projectTeam")
    def project_team(self) -> Optional[pulumi.Input[Mapping[str, pulumi.Input[str]]]]:
        """
        The project team associated with the entity, if any.
        """
        return pulumi.get(self, "project_team")

    @project_team.setter
    def project_team(self, value: Optional[pulumi.Input[Mapping[str, pulumi.Input[str]]]]):
        pulumi.set(self, "project_team", value)

    @property
    @pulumi.getter
    def role(self) -> Optional[pulumi.Input[str]]:
        """
        The access permission for the entity.
        """
        return pulumi.get(self, "role")

    @role.setter
    def role(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "role", value)

    @property
    @pulumi.getter(name="selfLink")
    def self_link(self) -> Optional[pulumi.Input[str]]:
        """
        The link to this access-control entry.
        """
        return pulumi.get(self, "self_link")

    @self_link.setter
    def self_link(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "self_link", value)


