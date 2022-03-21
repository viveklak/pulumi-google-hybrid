# coding=utf-8
# *** WARNING: this file was generated by the Pulumi SDK Generator. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import warnings
import pulumi
import pulumi.runtime
from typing import Any, Mapping, Optional, Sequence, Union, overload
from ... import _utilities
from . import outputs
from ._enums import *
from ._inputs import *

__all__ = ['IndexArgs', 'Index']

@pulumi.input_type
class IndexArgs:
    def __init__(__self__, *,
                 collection_group_id: pulumi.Input[str],
                 database_id: pulumi.Input[str],
                 fields: Optional[pulumi.Input[Sequence[pulumi.Input['GoogleFirestoreAdminV1beta2IndexFieldArgs']]]] = None,
                 project: Optional[pulumi.Input[str]] = None,
                 query_scope: Optional[pulumi.Input['IndexQueryScope']] = None):
        """
        The set of arguments for constructing a Index resource.
        :param pulumi.Input[Sequence[pulumi.Input['GoogleFirestoreAdminV1beta2IndexFieldArgs']]] fields: The fields supported by this index. For composite indexes, this is always 2 or more fields. The last field entry is always for the field path `__name__`. If, on creation, `__name__` was not specified as the last field, it will be added automatically with the same direction as that of the last field defined. If the final field in a composite index is not directional, the `__name__` will be ordered ASCENDING (unless explicitly specified). For single field indexes, this will always be exactly one entry with a field path equal to the field path of the associated field.
        :param pulumi.Input['IndexQueryScope'] query_scope: Indexes with a collection query scope specified allow queries against a collection that is the child of a specific document, specified at query time, and that has the same collection id. Indexes with a collection group query scope specified allow queries against all collections descended from a specific document, specified at query time, and that have the same collection id as this index.
        """
        pulumi.set(__self__, "collection_group_id", collection_group_id)
        pulumi.set(__self__, "database_id", database_id)
        if fields is not None:
            pulumi.set(__self__, "fields", fields)
        if project is not None:
            pulumi.set(__self__, "project", project)
        if query_scope is not None:
            pulumi.set(__self__, "query_scope", query_scope)

    @property
    @pulumi.getter(name="collectionGroupId")
    def collection_group_id(self) -> pulumi.Input[str]:
        return pulumi.get(self, "collection_group_id")

    @collection_group_id.setter
    def collection_group_id(self, value: pulumi.Input[str]):
        pulumi.set(self, "collection_group_id", value)

    @property
    @pulumi.getter(name="databaseId")
    def database_id(self) -> pulumi.Input[str]:
        return pulumi.get(self, "database_id")

    @database_id.setter
    def database_id(self, value: pulumi.Input[str]):
        pulumi.set(self, "database_id", value)

    @property
    @pulumi.getter
    def fields(self) -> Optional[pulumi.Input[Sequence[pulumi.Input['GoogleFirestoreAdminV1beta2IndexFieldArgs']]]]:
        """
        The fields supported by this index. For composite indexes, this is always 2 or more fields. The last field entry is always for the field path `__name__`. If, on creation, `__name__` was not specified as the last field, it will be added automatically with the same direction as that of the last field defined. If the final field in a composite index is not directional, the `__name__` will be ordered ASCENDING (unless explicitly specified). For single field indexes, this will always be exactly one entry with a field path equal to the field path of the associated field.
        """
        return pulumi.get(self, "fields")

    @fields.setter
    def fields(self, value: Optional[pulumi.Input[Sequence[pulumi.Input['GoogleFirestoreAdminV1beta2IndexFieldArgs']]]]):
        pulumi.set(self, "fields", value)

    @property
    @pulumi.getter
    def project(self) -> Optional[pulumi.Input[str]]:
        return pulumi.get(self, "project")

    @project.setter
    def project(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "project", value)

    @property
    @pulumi.getter(name="queryScope")
    def query_scope(self) -> Optional[pulumi.Input['IndexQueryScope']]:
        """
        Indexes with a collection query scope specified allow queries against a collection that is the child of a specific document, specified at query time, and that has the same collection id. Indexes with a collection group query scope specified allow queries against all collections descended from a specific document, specified at query time, and that have the same collection id as this index.
        """
        return pulumi.get(self, "query_scope")

    @query_scope.setter
    def query_scope(self, value: Optional[pulumi.Input['IndexQueryScope']]):
        pulumi.set(self, "query_scope", value)


class Index(pulumi.CustomResource):
    @overload
    def __init__(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 collection_group_id: Optional[pulumi.Input[str]] = None,
                 database_id: Optional[pulumi.Input[str]] = None,
                 fields: Optional[pulumi.Input[Sequence[pulumi.Input[pulumi.InputType['GoogleFirestoreAdminV1beta2IndexFieldArgs']]]]] = None,
                 project: Optional[pulumi.Input[str]] = None,
                 query_scope: Optional[pulumi.Input['IndexQueryScope']] = None,
                 __props__=None):
        """
        Creates a composite index. This returns a google.longrunning.Operation which may be used to track the status of the creation. The metadata for the operation will be the type IndexOperationMetadata.
        Auto-naming is currently not supported for this resource.

        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[Sequence[pulumi.Input[pulumi.InputType['GoogleFirestoreAdminV1beta2IndexFieldArgs']]]] fields: The fields supported by this index. For composite indexes, this is always 2 or more fields. The last field entry is always for the field path `__name__`. If, on creation, `__name__` was not specified as the last field, it will be added automatically with the same direction as that of the last field defined. If the final field in a composite index is not directional, the `__name__` will be ordered ASCENDING (unless explicitly specified). For single field indexes, this will always be exactly one entry with a field path equal to the field path of the associated field.
        :param pulumi.Input['IndexQueryScope'] query_scope: Indexes with a collection query scope specified allow queries against a collection that is the child of a specific document, specified at query time, and that has the same collection id. Indexes with a collection group query scope specified allow queries against all collections descended from a specific document, specified at query time, and that have the same collection id as this index.
        """
        ...
    @overload
    def __init__(__self__,
                 resource_name: str,
                 args: IndexArgs,
                 opts: Optional[pulumi.ResourceOptions] = None):
        """
        Creates a composite index. This returns a google.longrunning.Operation which may be used to track the status of the creation. The metadata for the operation will be the type IndexOperationMetadata.
        Auto-naming is currently not supported for this resource.

        :param str resource_name: The name of the resource.
        :param IndexArgs args: The arguments to use to populate this resource's properties.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        ...
    def __init__(__self__, resource_name: str, *args, **kwargs):
        resource_args, opts = _utilities.get_resource_args_opts(IndexArgs, pulumi.ResourceOptions, *args, **kwargs)
        if resource_args is not None:
            __self__._internal_init(resource_name, opts, **resource_args.__dict__)
        else:
            __self__._internal_init(resource_name, *args, **kwargs)

    def _internal_init(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 collection_group_id: Optional[pulumi.Input[str]] = None,
                 database_id: Optional[pulumi.Input[str]] = None,
                 fields: Optional[pulumi.Input[Sequence[pulumi.Input[pulumi.InputType['GoogleFirestoreAdminV1beta2IndexFieldArgs']]]]] = None,
                 project: Optional[pulumi.Input[str]] = None,
                 query_scope: Optional[pulumi.Input['IndexQueryScope']] = None,
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
            __props__ = IndexArgs.__new__(IndexArgs)

            if collection_group_id is None and not opts.urn:
                raise TypeError("Missing required property 'collection_group_id'")
            __props__.__dict__["collection_group_id"] = collection_group_id
            if database_id is None and not opts.urn:
                raise TypeError("Missing required property 'database_id'")
            __props__.__dict__["database_id"] = database_id
            __props__.__dict__["fields"] = fields
            __props__.__dict__["project"] = project
            __props__.__dict__["query_scope"] = query_scope
            __props__.__dict__["name"] = None
            __props__.__dict__["state"] = None
        super(Index, __self__).__init__(
            'google-native:firestore/v1beta2:Index',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name: str,
            id: pulumi.Input[str],
            opts: Optional[pulumi.ResourceOptions] = None) -> 'Index':
        """
        Get an existing Index resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param pulumi.Input[str] id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = IndexArgs.__new__(IndexArgs)

        __props__.__dict__["fields"] = None
        __props__.__dict__["name"] = None
        __props__.__dict__["query_scope"] = None
        __props__.__dict__["state"] = None
        return Index(resource_name, opts=opts, __props__=__props__)

    @property
    @pulumi.getter
    def fields(self) -> pulumi.Output[Sequence['outputs.GoogleFirestoreAdminV1beta2IndexFieldResponse']]:
        """
        The fields supported by this index. For composite indexes, this is always 2 or more fields. The last field entry is always for the field path `__name__`. If, on creation, `__name__` was not specified as the last field, it will be added automatically with the same direction as that of the last field defined. If the final field in a composite index is not directional, the `__name__` will be ordered ASCENDING (unless explicitly specified). For single field indexes, this will always be exactly one entry with a field path equal to the field path of the associated field.
        """
        return pulumi.get(self, "fields")

    @property
    @pulumi.getter
    def name(self) -> pulumi.Output[str]:
        """
        A server defined name for this index. The form of this name for composite indexes will be: `projects/{project_id}/databases/{database_id}/collectionGroups/{collection_id}/indexes/{composite_index_id}` For single field indexes, this field will be empty.
        """
        return pulumi.get(self, "name")

    @property
    @pulumi.getter(name="queryScope")
    def query_scope(self) -> pulumi.Output[str]:
        """
        Indexes with a collection query scope specified allow queries against a collection that is the child of a specific document, specified at query time, and that has the same collection id. Indexes with a collection group query scope specified allow queries against all collections descended from a specific document, specified at query time, and that have the same collection id as this index.
        """
        return pulumi.get(self, "query_scope")

    @property
    @pulumi.getter
    def state(self) -> pulumi.Output[str]:
        """
        The serving state of the index.
        """
        return pulumi.get(self, "state")

