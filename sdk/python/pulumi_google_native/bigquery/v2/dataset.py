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

__all__ = ['DatasetArgs', 'Dataset']

@pulumi.input_type
class DatasetArgs:
    def __init__(__self__, *,
                 access: Optional[pulumi.Input[Sequence[pulumi.Input['DatasetAccessItemArgs']]]] = None,
                 dataset_reference: Optional[pulumi.Input['DatasetReferenceArgs']] = None,
                 default_encryption_configuration: Optional[pulumi.Input['EncryptionConfigurationArgs']] = None,
                 default_partition_expiration_ms: Optional[pulumi.Input[str]] = None,
                 default_table_expiration_ms: Optional[pulumi.Input[str]] = None,
                 description: Optional[pulumi.Input[str]] = None,
                 friendly_name: Optional[pulumi.Input[str]] = None,
                 is_case_insensitive: Optional[pulumi.Input[bool]] = None,
                 labels: Optional[pulumi.Input[Mapping[str, pulumi.Input[str]]]] = None,
                 location: Optional[pulumi.Input[str]] = None,
                 max_time_travel_hours: Optional[pulumi.Input[str]] = None,
                 project: Optional[pulumi.Input[str]] = None,
                 tags: Optional[pulumi.Input[Sequence[pulumi.Input['DatasetTagsItemArgs']]]] = None):
        """
        The set of arguments for constructing a Dataset resource.
        :param pulumi.Input[Sequence[pulumi.Input['DatasetAccessItemArgs']]] access: [Optional] An array of objects that define dataset access for one or more entities. You can set this property when inserting or updating a dataset in order to control who is allowed to access the data. If unspecified at dataset creation time, BigQuery adds default dataset access for the following entities: access.specialGroup: projectReaders; access.role: READER; access.specialGroup: projectWriters; access.role: WRITER; access.specialGroup: projectOwners; access.role: OWNER; access.userByEmail: [dataset creator email]; access.role: OWNER;
        :param pulumi.Input['DatasetReferenceArgs'] dataset_reference: [Required] A reference that identifies the dataset.
        :param pulumi.Input[str] default_partition_expiration_ms: [Optional] The default partition expiration for all partitioned tables in the dataset, in milliseconds. Once this property is set, all newly-created partitioned tables in the dataset will have an expirationMs property in the timePartitioning settings set to this value, and changing the value will only affect new tables, not existing ones. The storage in a partition will have an expiration time of its partition time plus this value. Setting this property overrides the use of defaultTableExpirationMs for partitioned tables: only one of defaultTableExpirationMs and defaultPartitionExpirationMs will be used for any new partitioned table. If you provide an explicit timePartitioning.expirationMs when creating or updating a partitioned table, that value takes precedence over the default partition expiration time indicated by this property.
        :param pulumi.Input[str] default_table_expiration_ms: [Optional] The default lifetime of all tables in the dataset, in milliseconds. The minimum value is 3600000 milliseconds (one hour). Once this property is set, all newly-created tables in the dataset will have an expirationTime property set to the creation time plus the value in this property, and changing the value will only affect new tables, not existing ones. When the expirationTime for a given table is reached, that table will be deleted automatically. If a table's expirationTime is modified or removed before the table expires, or if you provide an explicit expirationTime when creating a table, that value takes precedence over the default expiration time indicated by this property.
        :param pulumi.Input[str] description: [Optional] A user-friendly description of the dataset.
        :param pulumi.Input[str] friendly_name: [Optional] A descriptive name for the dataset.
        :param pulumi.Input[bool] is_case_insensitive: [Optional] Indicates if table names are case insensitive in the dataset.
        :param pulumi.Input[Mapping[str, pulumi.Input[str]]] labels: The labels associated with this dataset. You can use these to organize and group your datasets. You can set this property when inserting or updating a dataset. See Creating and Updating Dataset Labels for more information.
        :param pulumi.Input[str] location: The geographic location where the dataset should reside. The default value is US. See details at https://cloud.google.com/bigquery/docs/locations.
        :param pulumi.Input[str] max_time_travel_hours: [Optional] Number of hours for the max time travel for all tables in the dataset.
        :param pulumi.Input[Sequence[pulumi.Input['DatasetTagsItemArgs']]] tags: [Optional]The tags associated with this dataset. Tag keys are globally unique.
        """
        if access is not None:
            pulumi.set(__self__, "access", access)
        if dataset_reference is not None:
            pulumi.set(__self__, "dataset_reference", dataset_reference)
        if default_encryption_configuration is not None:
            pulumi.set(__self__, "default_encryption_configuration", default_encryption_configuration)
        if default_partition_expiration_ms is not None:
            pulumi.set(__self__, "default_partition_expiration_ms", default_partition_expiration_ms)
        if default_table_expiration_ms is not None:
            pulumi.set(__self__, "default_table_expiration_ms", default_table_expiration_ms)
        if description is not None:
            pulumi.set(__self__, "description", description)
        if friendly_name is not None:
            pulumi.set(__self__, "friendly_name", friendly_name)
        if is_case_insensitive is not None:
            pulumi.set(__self__, "is_case_insensitive", is_case_insensitive)
        if labels is not None:
            pulumi.set(__self__, "labels", labels)
        if location is not None:
            pulumi.set(__self__, "location", location)
        if max_time_travel_hours is not None:
            pulumi.set(__self__, "max_time_travel_hours", max_time_travel_hours)
        if project is not None:
            pulumi.set(__self__, "project", project)
        if tags is not None:
            pulumi.set(__self__, "tags", tags)

    @property
    @pulumi.getter
    def access(self) -> Optional[pulumi.Input[Sequence[pulumi.Input['DatasetAccessItemArgs']]]]:
        """
        [Optional] An array of objects that define dataset access for one or more entities. You can set this property when inserting or updating a dataset in order to control who is allowed to access the data. If unspecified at dataset creation time, BigQuery adds default dataset access for the following entities: access.specialGroup: projectReaders; access.role: READER; access.specialGroup: projectWriters; access.role: WRITER; access.specialGroup: projectOwners; access.role: OWNER; access.userByEmail: [dataset creator email]; access.role: OWNER;
        """
        return pulumi.get(self, "access")

    @access.setter
    def access(self, value: Optional[pulumi.Input[Sequence[pulumi.Input['DatasetAccessItemArgs']]]]):
        pulumi.set(self, "access", value)

    @property
    @pulumi.getter(name="datasetReference")
    def dataset_reference(self) -> Optional[pulumi.Input['DatasetReferenceArgs']]:
        """
        [Required] A reference that identifies the dataset.
        """
        return pulumi.get(self, "dataset_reference")

    @dataset_reference.setter
    def dataset_reference(self, value: Optional[pulumi.Input['DatasetReferenceArgs']]):
        pulumi.set(self, "dataset_reference", value)

    @property
    @pulumi.getter(name="defaultEncryptionConfiguration")
    def default_encryption_configuration(self) -> Optional[pulumi.Input['EncryptionConfigurationArgs']]:
        return pulumi.get(self, "default_encryption_configuration")

    @default_encryption_configuration.setter
    def default_encryption_configuration(self, value: Optional[pulumi.Input['EncryptionConfigurationArgs']]):
        pulumi.set(self, "default_encryption_configuration", value)

    @property
    @pulumi.getter(name="defaultPartitionExpirationMs")
    def default_partition_expiration_ms(self) -> Optional[pulumi.Input[str]]:
        """
        [Optional] The default partition expiration for all partitioned tables in the dataset, in milliseconds. Once this property is set, all newly-created partitioned tables in the dataset will have an expirationMs property in the timePartitioning settings set to this value, and changing the value will only affect new tables, not existing ones. The storage in a partition will have an expiration time of its partition time plus this value. Setting this property overrides the use of defaultTableExpirationMs for partitioned tables: only one of defaultTableExpirationMs and defaultPartitionExpirationMs will be used for any new partitioned table. If you provide an explicit timePartitioning.expirationMs when creating or updating a partitioned table, that value takes precedence over the default partition expiration time indicated by this property.
        """
        return pulumi.get(self, "default_partition_expiration_ms")

    @default_partition_expiration_ms.setter
    def default_partition_expiration_ms(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "default_partition_expiration_ms", value)

    @property
    @pulumi.getter(name="defaultTableExpirationMs")
    def default_table_expiration_ms(self) -> Optional[pulumi.Input[str]]:
        """
        [Optional] The default lifetime of all tables in the dataset, in milliseconds. The minimum value is 3600000 milliseconds (one hour). Once this property is set, all newly-created tables in the dataset will have an expirationTime property set to the creation time plus the value in this property, and changing the value will only affect new tables, not existing ones. When the expirationTime for a given table is reached, that table will be deleted automatically. If a table's expirationTime is modified or removed before the table expires, or if you provide an explicit expirationTime when creating a table, that value takes precedence over the default expiration time indicated by this property.
        """
        return pulumi.get(self, "default_table_expiration_ms")

    @default_table_expiration_ms.setter
    def default_table_expiration_ms(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "default_table_expiration_ms", value)

    @property
    @pulumi.getter
    def description(self) -> Optional[pulumi.Input[str]]:
        """
        [Optional] A user-friendly description of the dataset.
        """
        return pulumi.get(self, "description")

    @description.setter
    def description(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "description", value)

    @property
    @pulumi.getter(name="friendlyName")
    def friendly_name(self) -> Optional[pulumi.Input[str]]:
        """
        [Optional] A descriptive name for the dataset.
        """
        return pulumi.get(self, "friendly_name")

    @friendly_name.setter
    def friendly_name(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "friendly_name", value)

    @property
    @pulumi.getter(name="isCaseInsensitive")
    def is_case_insensitive(self) -> Optional[pulumi.Input[bool]]:
        """
        [Optional] Indicates if table names are case insensitive in the dataset.
        """
        return pulumi.get(self, "is_case_insensitive")

    @is_case_insensitive.setter
    def is_case_insensitive(self, value: Optional[pulumi.Input[bool]]):
        pulumi.set(self, "is_case_insensitive", value)

    @property
    @pulumi.getter
    def labels(self) -> Optional[pulumi.Input[Mapping[str, pulumi.Input[str]]]]:
        """
        The labels associated with this dataset. You can use these to organize and group your datasets. You can set this property when inserting or updating a dataset. See Creating and Updating Dataset Labels for more information.
        """
        return pulumi.get(self, "labels")

    @labels.setter
    def labels(self, value: Optional[pulumi.Input[Mapping[str, pulumi.Input[str]]]]):
        pulumi.set(self, "labels", value)

    @property
    @pulumi.getter
    def location(self) -> Optional[pulumi.Input[str]]:
        """
        The geographic location where the dataset should reside. The default value is US. See details at https://cloud.google.com/bigquery/docs/locations.
        """
        return pulumi.get(self, "location")

    @location.setter
    def location(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "location", value)

    @property
    @pulumi.getter(name="maxTimeTravelHours")
    def max_time_travel_hours(self) -> Optional[pulumi.Input[str]]:
        """
        [Optional] Number of hours for the max time travel for all tables in the dataset.
        """
        return pulumi.get(self, "max_time_travel_hours")

    @max_time_travel_hours.setter
    def max_time_travel_hours(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "max_time_travel_hours", value)

    @property
    @pulumi.getter
    def project(self) -> Optional[pulumi.Input[str]]:
        return pulumi.get(self, "project")

    @project.setter
    def project(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "project", value)

    @property
    @pulumi.getter
    def tags(self) -> Optional[pulumi.Input[Sequence[pulumi.Input['DatasetTagsItemArgs']]]]:
        """
        [Optional]The tags associated with this dataset. Tag keys are globally unique.
        """
        return pulumi.get(self, "tags")

    @tags.setter
    def tags(self, value: Optional[pulumi.Input[Sequence[pulumi.Input['DatasetTagsItemArgs']]]]):
        pulumi.set(self, "tags", value)


class Dataset(pulumi.CustomResource):
    @overload
    def __init__(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 access: Optional[pulumi.Input[Sequence[pulumi.Input[pulumi.InputType['DatasetAccessItemArgs']]]]] = None,
                 dataset_reference: Optional[pulumi.Input[pulumi.InputType['DatasetReferenceArgs']]] = None,
                 default_encryption_configuration: Optional[pulumi.Input[pulumi.InputType['EncryptionConfigurationArgs']]] = None,
                 default_partition_expiration_ms: Optional[pulumi.Input[str]] = None,
                 default_table_expiration_ms: Optional[pulumi.Input[str]] = None,
                 description: Optional[pulumi.Input[str]] = None,
                 friendly_name: Optional[pulumi.Input[str]] = None,
                 is_case_insensitive: Optional[pulumi.Input[bool]] = None,
                 labels: Optional[pulumi.Input[Mapping[str, pulumi.Input[str]]]] = None,
                 location: Optional[pulumi.Input[str]] = None,
                 max_time_travel_hours: Optional[pulumi.Input[str]] = None,
                 project: Optional[pulumi.Input[str]] = None,
                 tags: Optional[pulumi.Input[Sequence[pulumi.Input[pulumi.InputType['DatasetTagsItemArgs']]]]] = None,
                 __props__=None):
        """
        Creates a new empty dataset.
        Auto-naming is currently not supported for this resource.

        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[Sequence[pulumi.Input[pulumi.InputType['DatasetAccessItemArgs']]]] access: [Optional] An array of objects that define dataset access for one or more entities. You can set this property when inserting or updating a dataset in order to control who is allowed to access the data. If unspecified at dataset creation time, BigQuery adds default dataset access for the following entities: access.specialGroup: projectReaders; access.role: READER; access.specialGroup: projectWriters; access.role: WRITER; access.specialGroup: projectOwners; access.role: OWNER; access.userByEmail: [dataset creator email]; access.role: OWNER;
        :param pulumi.Input[pulumi.InputType['DatasetReferenceArgs']] dataset_reference: [Required] A reference that identifies the dataset.
        :param pulumi.Input[str] default_partition_expiration_ms: [Optional] The default partition expiration for all partitioned tables in the dataset, in milliseconds. Once this property is set, all newly-created partitioned tables in the dataset will have an expirationMs property in the timePartitioning settings set to this value, and changing the value will only affect new tables, not existing ones. The storage in a partition will have an expiration time of its partition time plus this value. Setting this property overrides the use of defaultTableExpirationMs for partitioned tables: only one of defaultTableExpirationMs and defaultPartitionExpirationMs will be used for any new partitioned table. If you provide an explicit timePartitioning.expirationMs when creating or updating a partitioned table, that value takes precedence over the default partition expiration time indicated by this property.
        :param pulumi.Input[str] default_table_expiration_ms: [Optional] The default lifetime of all tables in the dataset, in milliseconds. The minimum value is 3600000 milliseconds (one hour). Once this property is set, all newly-created tables in the dataset will have an expirationTime property set to the creation time plus the value in this property, and changing the value will only affect new tables, not existing ones. When the expirationTime for a given table is reached, that table will be deleted automatically. If a table's expirationTime is modified or removed before the table expires, or if you provide an explicit expirationTime when creating a table, that value takes precedence over the default expiration time indicated by this property.
        :param pulumi.Input[str] description: [Optional] A user-friendly description of the dataset.
        :param pulumi.Input[str] friendly_name: [Optional] A descriptive name for the dataset.
        :param pulumi.Input[bool] is_case_insensitive: [Optional] Indicates if table names are case insensitive in the dataset.
        :param pulumi.Input[Mapping[str, pulumi.Input[str]]] labels: The labels associated with this dataset. You can use these to organize and group your datasets. You can set this property when inserting or updating a dataset. See Creating and Updating Dataset Labels for more information.
        :param pulumi.Input[str] location: The geographic location where the dataset should reside. The default value is US. See details at https://cloud.google.com/bigquery/docs/locations.
        :param pulumi.Input[str] max_time_travel_hours: [Optional] Number of hours for the max time travel for all tables in the dataset.
        :param pulumi.Input[Sequence[pulumi.Input[pulumi.InputType['DatasetTagsItemArgs']]]] tags: [Optional]The tags associated with this dataset. Tag keys are globally unique.
        """
        ...
    @overload
    def __init__(__self__,
                 resource_name: str,
                 args: Optional[DatasetArgs] = None,
                 opts: Optional[pulumi.ResourceOptions] = None):
        """
        Creates a new empty dataset.
        Auto-naming is currently not supported for this resource.

        :param str resource_name: The name of the resource.
        :param DatasetArgs args: The arguments to use to populate this resource's properties.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        ...
    def __init__(__self__, resource_name: str, *args, **kwargs):
        resource_args, opts = _utilities.get_resource_args_opts(DatasetArgs, pulumi.ResourceOptions, *args, **kwargs)
        if resource_args is not None:
            __self__._internal_init(resource_name, opts, **resource_args.__dict__)
        else:
            __self__._internal_init(resource_name, *args, **kwargs)

    def _internal_init(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 access: Optional[pulumi.Input[Sequence[pulumi.Input[pulumi.InputType['DatasetAccessItemArgs']]]]] = None,
                 dataset_reference: Optional[pulumi.Input[pulumi.InputType['DatasetReferenceArgs']]] = None,
                 default_encryption_configuration: Optional[pulumi.Input[pulumi.InputType['EncryptionConfigurationArgs']]] = None,
                 default_partition_expiration_ms: Optional[pulumi.Input[str]] = None,
                 default_table_expiration_ms: Optional[pulumi.Input[str]] = None,
                 description: Optional[pulumi.Input[str]] = None,
                 friendly_name: Optional[pulumi.Input[str]] = None,
                 is_case_insensitive: Optional[pulumi.Input[bool]] = None,
                 labels: Optional[pulumi.Input[Mapping[str, pulumi.Input[str]]]] = None,
                 location: Optional[pulumi.Input[str]] = None,
                 max_time_travel_hours: Optional[pulumi.Input[str]] = None,
                 project: Optional[pulumi.Input[str]] = None,
                 tags: Optional[pulumi.Input[Sequence[pulumi.Input[pulumi.InputType['DatasetTagsItemArgs']]]]] = None,
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
            __props__ = DatasetArgs.__new__(DatasetArgs)

            __props__.__dict__["access"] = access
            __props__.__dict__["dataset_reference"] = dataset_reference
            __props__.__dict__["default_encryption_configuration"] = default_encryption_configuration
            __props__.__dict__["default_partition_expiration_ms"] = default_partition_expiration_ms
            __props__.__dict__["default_table_expiration_ms"] = default_table_expiration_ms
            __props__.__dict__["description"] = description
            __props__.__dict__["friendly_name"] = friendly_name
            __props__.__dict__["is_case_insensitive"] = is_case_insensitive
            __props__.__dict__["labels"] = labels
            __props__.__dict__["location"] = location
            __props__.__dict__["max_time_travel_hours"] = max_time_travel_hours
            __props__.__dict__["project"] = project
            __props__.__dict__["tags"] = tags
            __props__.__dict__["creation_time"] = None
            __props__.__dict__["default_collation"] = None
            __props__.__dict__["etag"] = None
            __props__.__dict__["kind"] = None
            __props__.__dict__["last_modified_time"] = None
            __props__.__dict__["satisfies_pzs"] = None
            __props__.__dict__["self_link"] = None
        super(Dataset, __self__).__init__(
            'google-native:bigquery/v2:Dataset',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name: str,
            id: pulumi.Input[str],
            opts: Optional[pulumi.ResourceOptions] = None) -> 'Dataset':
        """
        Get an existing Dataset resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param pulumi.Input[str] id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = DatasetArgs.__new__(DatasetArgs)

        __props__.__dict__["access"] = None
        __props__.__dict__["creation_time"] = None
        __props__.__dict__["dataset_reference"] = None
        __props__.__dict__["default_collation"] = None
        __props__.__dict__["default_encryption_configuration"] = None
        __props__.__dict__["default_partition_expiration_ms"] = None
        __props__.__dict__["default_table_expiration_ms"] = None
        __props__.__dict__["description"] = None
        __props__.__dict__["etag"] = None
        __props__.__dict__["friendly_name"] = None
        __props__.__dict__["is_case_insensitive"] = None
        __props__.__dict__["kind"] = None
        __props__.__dict__["labels"] = None
        __props__.__dict__["last_modified_time"] = None
        __props__.__dict__["location"] = None
        __props__.__dict__["max_time_travel_hours"] = None
        __props__.__dict__["satisfies_pzs"] = None
        __props__.__dict__["self_link"] = None
        __props__.__dict__["tags"] = None
        return Dataset(resource_name, opts=opts, __props__=__props__)

    @property
    @pulumi.getter
    def access(self) -> pulumi.Output[Sequence['outputs.DatasetAccessItemResponse']]:
        """
        [Optional] An array of objects that define dataset access for one or more entities. You can set this property when inserting or updating a dataset in order to control who is allowed to access the data. If unspecified at dataset creation time, BigQuery adds default dataset access for the following entities: access.specialGroup: projectReaders; access.role: READER; access.specialGroup: projectWriters; access.role: WRITER; access.specialGroup: projectOwners; access.role: OWNER; access.userByEmail: [dataset creator email]; access.role: OWNER;
        """
        return pulumi.get(self, "access")

    @property
    @pulumi.getter(name="creationTime")
    def creation_time(self) -> pulumi.Output[str]:
        """
        The time when this dataset was created, in milliseconds since the epoch.
        """
        return pulumi.get(self, "creation_time")

    @property
    @pulumi.getter(name="datasetReference")
    def dataset_reference(self) -> pulumi.Output['outputs.DatasetReferenceResponse']:
        """
        [Required] A reference that identifies the dataset.
        """
        return pulumi.get(self, "dataset_reference")

    @property
    @pulumi.getter(name="defaultCollation")
    def default_collation(self) -> pulumi.Output[str]:
        """
        The default collation of the dataset.
        """
        return pulumi.get(self, "default_collation")

    @property
    @pulumi.getter(name="defaultEncryptionConfiguration")
    def default_encryption_configuration(self) -> pulumi.Output['outputs.EncryptionConfigurationResponse']:
        return pulumi.get(self, "default_encryption_configuration")

    @property
    @pulumi.getter(name="defaultPartitionExpirationMs")
    def default_partition_expiration_ms(self) -> pulumi.Output[str]:
        """
        [Optional] The default partition expiration for all partitioned tables in the dataset, in milliseconds. Once this property is set, all newly-created partitioned tables in the dataset will have an expirationMs property in the timePartitioning settings set to this value, and changing the value will only affect new tables, not existing ones. The storage in a partition will have an expiration time of its partition time plus this value. Setting this property overrides the use of defaultTableExpirationMs for partitioned tables: only one of defaultTableExpirationMs and defaultPartitionExpirationMs will be used for any new partitioned table. If you provide an explicit timePartitioning.expirationMs when creating or updating a partitioned table, that value takes precedence over the default partition expiration time indicated by this property.
        """
        return pulumi.get(self, "default_partition_expiration_ms")

    @property
    @pulumi.getter(name="defaultTableExpirationMs")
    def default_table_expiration_ms(self) -> pulumi.Output[str]:
        """
        [Optional] The default lifetime of all tables in the dataset, in milliseconds. The minimum value is 3600000 milliseconds (one hour). Once this property is set, all newly-created tables in the dataset will have an expirationTime property set to the creation time plus the value in this property, and changing the value will only affect new tables, not existing ones. When the expirationTime for a given table is reached, that table will be deleted automatically. If a table's expirationTime is modified or removed before the table expires, or if you provide an explicit expirationTime when creating a table, that value takes precedence over the default expiration time indicated by this property.
        """
        return pulumi.get(self, "default_table_expiration_ms")

    @property
    @pulumi.getter
    def description(self) -> pulumi.Output[str]:
        """
        [Optional] A user-friendly description of the dataset.
        """
        return pulumi.get(self, "description")

    @property
    @pulumi.getter
    def etag(self) -> pulumi.Output[str]:
        """
        A hash of the resource.
        """
        return pulumi.get(self, "etag")

    @property
    @pulumi.getter(name="friendlyName")
    def friendly_name(self) -> pulumi.Output[str]:
        """
        [Optional] A descriptive name for the dataset.
        """
        return pulumi.get(self, "friendly_name")

    @property
    @pulumi.getter(name="isCaseInsensitive")
    def is_case_insensitive(self) -> pulumi.Output[bool]:
        """
        [Optional] Indicates if table names are case insensitive in the dataset.
        """
        return pulumi.get(self, "is_case_insensitive")

    @property
    @pulumi.getter
    def kind(self) -> pulumi.Output[str]:
        """
        The resource type.
        """
        return pulumi.get(self, "kind")

    @property
    @pulumi.getter
    def labels(self) -> pulumi.Output[Mapping[str, str]]:
        """
        The labels associated with this dataset. You can use these to organize and group your datasets. You can set this property when inserting or updating a dataset. See Creating and Updating Dataset Labels for more information.
        """
        return pulumi.get(self, "labels")

    @property
    @pulumi.getter(name="lastModifiedTime")
    def last_modified_time(self) -> pulumi.Output[str]:
        """
        The date when this dataset or any of its tables was last modified, in milliseconds since the epoch.
        """
        return pulumi.get(self, "last_modified_time")

    @property
    @pulumi.getter
    def location(self) -> pulumi.Output[str]:
        """
        The geographic location where the dataset should reside. The default value is US. See details at https://cloud.google.com/bigquery/docs/locations.
        """
        return pulumi.get(self, "location")

    @property
    @pulumi.getter(name="maxTimeTravelHours")
    def max_time_travel_hours(self) -> pulumi.Output[str]:
        """
        [Optional] Number of hours for the max time travel for all tables in the dataset.
        """
        return pulumi.get(self, "max_time_travel_hours")

    @property
    @pulumi.getter(name="satisfiesPZS")
    def satisfies_pzs(self) -> pulumi.Output[bool]:
        """
        Reserved for future use.
        """
        return pulumi.get(self, "satisfies_pzs")

    @property
    @pulumi.getter(name="selfLink")
    def self_link(self) -> pulumi.Output[str]:
        """
        A URL that can be used to access the resource again. You can use this URL in Get or Update requests to the resource.
        """
        return pulumi.get(self, "self_link")

    @property
    @pulumi.getter
    def tags(self) -> pulumi.Output[Sequence['outputs.DatasetTagsItemResponse']]:
        """
        [Optional]The tags associated with this dataset. Tag keys are globally unique.
        """
        return pulumi.get(self, "tags")

