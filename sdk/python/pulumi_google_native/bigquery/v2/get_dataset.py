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
    'GetDatasetResult',
    'AwaitableGetDatasetResult',
    'get_dataset',
    'get_dataset_output',
]

@pulumi.output_type
class GetDatasetResult:
    def __init__(__self__, access=None, creation_time=None, dataset_reference=None, default_collation=None, default_encryption_configuration=None, default_partition_expiration_ms=None, default_table_expiration_ms=None, description=None, etag=None, friendly_name=None, is_case_insensitive=None, kind=None, labels=None, last_modified_time=None, location=None, max_time_travel_hours=None, satisfies_pzs=None, self_link=None, tags=None):
        if access and not isinstance(access, list):
            raise TypeError("Expected argument 'access' to be a list")
        pulumi.set(__self__, "access", access)
        if creation_time and not isinstance(creation_time, str):
            raise TypeError("Expected argument 'creation_time' to be a str")
        pulumi.set(__self__, "creation_time", creation_time)
        if dataset_reference and not isinstance(dataset_reference, dict):
            raise TypeError("Expected argument 'dataset_reference' to be a dict")
        pulumi.set(__self__, "dataset_reference", dataset_reference)
        if default_collation and not isinstance(default_collation, str):
            raise TypeError("Expected argument 'default_collation' to be a str")
        pulumi.set(__self__, "default_collation", default_collation)
        if default_encryption_configuration and not isinstance(default_encryption_configuration, dict):
            raise TypeError("Expected argument 'default_encryption_configuration' to be a dict")
        pulumi.set(__self__, "default_encryption_configuration", default_encryption_configuration)
        if default_partition_expiration_ms and not isinstance(default_partition_expiration_ms, str):
            raise TypeError("Expected argument 'default_partition_expiration_ms' to be a str")
        pulumi.set(__self__, "default_partition_expiration_ms", default_partition_expiration_ms)
        if default_table_expiration_ms and not isinstance(default_table_expiration_ms, str):
            raise TypeError("Expected argument 'default_table_expiration_ms' to be a str")
        pulumi.set(__self__, "default_table_expiration_ms", default_table_expiration_ms)
        if description and not isinstance(description, str):
            raise TypeError("Expected argument 'description' to be a str")
        pulumi.set(__self__, "description", description)
        if etag and not isinstance(etag, str):
            raise TypeError("Expected argument 'etag' to be a str")
        pulumi.set(__self__, "etag", etag)
        if friendly_name and not isinstance(friendly_name, str):
            raise TypeError("Expected argument 'friendly_name' to be a str")
        pulumi.set(__self__, "friendly_name", friendly_name)
        if is_case_insensitive and not isinstance(is_case_insensitive, bool):
            raise TypeError("Expected argument 'is_case_insensitive' to be a bool")
        pulumi.set(__self__, "is_case_insensitive", is_case_insensitive)
        if kind and not isinstance(kind, str):
            raise TypeError("Expected argument 'kind' to be a str")
        pulumi.set(__self__, "kind", kind)
        if labels and not isinstance(labels, dict):
            raise TypeError("Expected argument 'labels' to be a dict")
        pulumi.set(__self__, "labels", labels)
        if last_modified_time and not isinstance(last_modified_time, str):
            raise TypeError("Expected argument 'last_modified_time' to be a str")
        pulumi.set(__self__, "last_modified_time", last_modified_time)
        if location and not isinstance(location, str):
            raise TypeError("Expected argument 'location' to be a str")
        pulumi.set(__self__, "location", location)
        if max_time_travel_hours and not isinstance(max_time_travel_hours, str):
            raise TypeError("Expected argument 'max_time_travel_hours' to be a str")
        pulumi.set(__self__, "max_time_travel_hours", max_time_travel_hours)
        if satisfies_pzs and not isinstance(satisfies_pzs, bool):
            raise TypeError("Expected argument 'satisfies_pzs' to be a bool")
        pulumi.set(__self__, "satisfies_pzs", satisfies_pzs)
        if self_link and not isinstance(self_link, str):
            raise TypeError("Expected argument 'self_link' to be a str")
        pulumi.set(__self__, "self_link", self_link)
        if tags and not isinstance(tags, list):
            raise TypeError("Expected argument 'tags' to be a list")
        pulumi.set(__self__, "tags", tags)

    @property
    @pulumi.getter
    def access(self) -> Sequence['outputs.DatasetAccessItemResponse']:
        """
        [Optional] An array of objects that define dataset access for one or more entities. You can set this property when inserting or updating a dataset in order to control who is allowed to access the data. If unspecified at dataset creation time, BigQuery adds default dataset access for the following entities: access.specialGroup: projectReaders; access.role: READER; access.specialGroup: projectWriters; access.role: WRITER; access.specialGroup: projectOwners; access.role: OWNER; access.userByEmail: [dataset creator email]; access.role: OWNER;
        """
        return pulumi.get(self, "access")

    @property
    @pulumi.getter(name="creationTime")
    def creation_time(self) -> str:
        """
        The time when this dataset was created, in milliseconds since the epoch.
        """
        return pulumi.get(self, "creation_time")

    @property
    @pulumi.getter(name="datasetReference")
    def dataset_reference(self) -> 'outputs.DatasetReferenceResponse':
        """
        [Required] A reference that identifies the dataset.
        """
        return pulumi.get(self, "dataset_reference")

    @property
    @pulumi.getter(name="defaultCollation")
    def default_collation(self) -> str:
        """
        The default collation of the dataset.
        """
        return pulumi.get(self, "default_collation")

    @property
    @pulumi.getter(name="defaultEncryptionConfiguration")
    def default_encryption_configuration(self) -> 'outputs.EncryptionConfigurationResponse':
        return pulumi.get(self, "default_encryption_configuration")

    @property
    @pulumi.getter(name="defaultPartitionExpirationMs")
    def default_partition_expiration_ms(self) -> str:
        """
        [Optional] The default partition expiration for all partitioned tables in the dataset, in milliseconds. Once this property is set, all newly-created partitioned tables in the dataset will have an expirationMs property in the timePartitioning settings set to this value, and changing the value will only affect new tables, not existing ones. The storage in a partition will have an expiration time of its partition time plus this value. Setting this property overrides the use of defaultTableExpirationMs for partitioned tables: only one of defaultTableExpirationMs and defaultPartitionExpirationMs will be used for any new partitioned table. If you provide an explicit timePartitioning.expirationMs when creating or updating a partitioned table, that value takes precedence over the default partition expiration time indicated by this property.
        """
        return pulumi.get(self, "default_partition_expiration_ms")

    @property
    @pulumi.getter(name="defaultTableExpirationMs")
    def default_table_expiration_ms(self) -> str:
        """
        [Optional] The default lifetime of all tables in the dataset, in milliseconds. The minimum value is 3600000 milliseconds (one hour). Once this property is set, all newly-created tables in the dataset will have an expirationTime property set to the creation time plus the value in this property, and changing the value will only affect new tables, not existing ones. When the expirationTime for a given table is reached, that table will be deleted automatically. If a table's expirationTime is modified or removed before the table expires, or if you provide an explicit expirationTime when creating a table, that value takes precedence over the default expiration time indicated by this property.
        """
        return pulumi.get(self, "default_table_expiration_ms")

    @property
    @pulumi.getter
    def description(self) -> str:
        """
        [Optional] A user-friendly description of the dataset.
        """
        return pulumi.get(self, "description")

    @property
    @pulumi.getter
    def etag(self) -> str:
        """
        A hash of the resource.
        """
        return pulumi.get(self, "etag")

    @property
    @pulumi.getter(name="friendlyName")
    def friendly_name(self) -> str:
        """
        [Optional] A descriptive name for the dataset.
        """
        return pulumi.get(self, "friendly_name")

    @property
    @pulumi.getter(name="isCaseInsensitive")
    def is_case_insensitive(self) -> bool:
        """
        [Optional] Indicates if table names are case insensitive in the dataset.
        """
        return pulumi.get(self, "is_case_insensitive")

    @property
    @pulumi.getter
    def kind(self) -> str:
        """
        The resource type.
        """
        return pulumi.get(self, "kind")

    @property
    @pulumi.getter
    def labels(self) -> Mapping[str, str]:
        """
        The labels associated with this dataset. You can use these to organize and group your datasets. You can set this property when inserting or updating a dataset. See Creating and Updating Dataset Labels for more information.
        """
        return pulumi.get(self, "labels")

    @property
    @pulumi.getter(name="lastModifiedTime")
    def last_modified_time(self) -> str:
        """
        The date when this dataset or any of its tables was last modified, in milliseconds since the epoch.
        """
        return pulumi.get(self, "last_modified_time")

    @property
    @pulumi.getter
    def location(self) -> str:
        """
        The geographic location where the dataset should reside. The default value is US. See details at https://cloud.google.com/bigquery/docs/locations.
        """
        return pulumi.get(self, "location")

    @property
    @pulumi.getter(name="maxTimeTravelHours")
    def max_time_travel_hours(self) -> str:
        """
        [Optional] Number of hours for the max time travel for all tables in the dataset.
        """
        return pulumi.get(self, "max_time_travel_hours")

    @property
    @pulumi.getter(name="satisfiesPZS")
    def satisfies_pzs(self) -> bool:
        """
        Reserved for future use.
        """
        return pulumi.get(self, "satisfies_pzs")

    @property
    @pulumi.getter(name="selfLink")
    def self_link(self) -> str:
        """
        A URL that can be used to access the resource again. You can use this URL in Get or Update requests to the resource.
        """
        return pulumi.get(self, "self_link")

    @property
    @pulumi.getter
    def tags(self) -> Sequence['outputs.DatasetTagsItemResponse']:
        """
        [Optional]The tags associated with this dataset. Tag keys are globally unique.
        """
        return pulumi.get(self, "tags")


class AwaitableGetDatasetResult(GetDatasetResult):
    # pylint: disable=using-constant-test
    def __await__(self):
        if False:
            yield self
        return GetDatasetResult(
            access=self.access,
            creation_time=self.creation_time,
            dataset_reference=self.dataset_reference,
            default_collation=self.default_collation,
            default_encryption_configuration=self.default_encryption_configuration,
            default_partition_expiration_ms=self.default_partition_expiration_ms,
            default_table_expiration_ms=self.default_table_expiration_ms,
            description=self.description,
            etag=self.etag,
            friendly_name=self.friendly_name,
            is_case_insensitive=self.is_case_insensitive,
            kind=self.kind,
            labels=self.labels,
            last_modified_time=self.last_modified_time,
            location=self.location,
            max_time_travel_hours=self.max_time_travel_hours,
            satisfies_pzs=self.satisfies_pzs,
            self_link=self.self_link,
            tags=self.tags)


def get_dataset(dataset_id: Optional[str] = None,
                project: Optional[str] = None,
                opts: Optional[pulumi.InvokeOptions] = None) -> AwaitableGetDatasetResult:
    """
    Returns the dataset specified by datasetID.
    """
    __args__ = dict()
    __args__['datasetId'] = dataset_id
    __args__['project'] = project
    if opts is None:
        opts = pulumi.InvokeOptions()
    if opts.version is None:
        opts.version = _utilities.get_version()
    __ret__ = pulumi.runtime.invoke('google-native:bigquery/v2:getDataset', __args__, opts=opts, typ=GetDatasetResult).value

    return AwaitableGetDatasetResult(
        access=__ret__.access,
        creation_time=__ret__.creation_time,
        dataset_reference=__ret__.dataset_reference,
        default_collation=__ret__.default_collation,
        default_encryption_configuration=__ret__.default_encryption_configuration,
        default_partition_expiration_ms=__ret__.default_partition_expiration_ms,
        default_table_expiration_ms=__ret__.default_table_expiration_ms,
        description=__ret__.description,
        etag=__ret__.etag,
        friendly_name=__ret__.friendly_name,
        is_case_insensitive=__ret__.is_case_insensitive,
        kind=__ret__.kind,
        labels=__ret__.labels,
        last_modified_time=__ret__.last_modified_time,
        location=__ret__.location,
        max_time_travel_hours=__ret__.max_time_travel_hours,
        satisfies_pzs=__ret__.satisfies_pzs,
        self_link=__ret__.self_link,
        tags=__ret__.tags)


@_utilities.lift_output_func(get_dataset)
def get_dataset_output(dataset_id: Optional[pulumi.Input[str]] = None,
                       project: Optional[pulumi.Input[Optional[str]]] = None,
                       opts: Optional[pulumi.InvokeOptions] = None) -> pulumi.Output[GetDatasetResult]:
    """
    Returns the dataset specified by datasetID.
    """
    ...
