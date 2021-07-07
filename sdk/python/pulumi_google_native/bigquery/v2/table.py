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

__all__ = ['TableArgs', 'Table']

@pulumi.input_type
class TableArgs:
    def __init__(__self__, *,
                 dataset_id: pulumi.Input[str],
                 project: pulumi.Input[str],
                 clustering: Optional[pulumi.Input['ClusteringArgs']] = None,
                 description: Optional[pulumi.Input[str]] = None,
                 encryption_configuration: Optional[pulumi.Input['EncryptionConfigurationArgs']] = None,
                 expiration_time: Optional[pulumi.Input[str]] = None,
                 external_data_configuration: Optional[pulumi.Input['ExternalDataConfigurationArgs']] = None,
                 friendly_name: Optional[pulumi.Input[str]] = None,
                 labels: Optional[pulumi.Input[Mapping[str, pulumi.Input[str]]]] = None,
                 materialized_view: Optional[pulumi.Input['MaterializedViewDefinitionArgs']] = None,
                 model: Optional[pulumi.Input['ModelDefinitionArgs']] = None,
                 range_partitioning: Optional[pulumi.Input['RangePartitioningArgs']] = None,
                 require_partition_filter: Optional[pulumi.Input[bool]] = None,
                 schema: Optional[pulumi.Input['TableSchemaArgs']] = None,
                 table_reference: Optional[pulumi.Input['TableReferenceArgs']] = None,
                 time_partitioning: Optional[pulumi.Input['TimePartitioningArgs']] = None,
                 view: Optional[pulumi.Input['ViewDefinitionArgs']] = None):
        """
        The set of arguments for constructing a Table resource.
        :param pulumi.Input['ClusteringArgs'] clustering: [Beta] Clustering specification for the table. Must be specified with partitioning, data in the table will be first partitioned and subsequently clustered.
        :param pulumi.Input[str] description: [Optional] A user-friendly description of this table.
        :param pulumi.Input['EncryptionConfigurationArgs'] encryption_configuration: Custom encryption configuration (e.g., Cloud KMS keys).
        :param pulumi.Input[str] expiration_time: [Optional] The time when this table expires, in milliseconds since the epoch. If not present, the table will persist indefinitely. Expired tables will be deleted and their storage reclaimed. The defaultTableExpirationMs property of the encapsulating dataset can be used to set a default expirationTime on newly created tables.
        :param pulumi.Input['ExternalDataConfigurationArgs'] external_data_configuration: [Optional] Describes the data format, location, and other properties of a table stored outside of BigQuery. By defining these properties, the data source can then be queried as if it were a standard BigQuery table.
        :param pulumi.Input[str] friendly_name: [Optional] A descriptive name for this table.
        :param pulumi.Input[Mapping[str, pulumi.Input[str]]] labels: The labels associated with this table. You can use these to organize and group your tables. Label keys and values can be no longer than 63 characters, can only contain lowercase letters, numeric characters, underscores and dashes. International characters are allowed. Label values are optional. Label keys must start with a letter and each label in the list must have a different key.
        :param pulumi.Input['MaterializedViewDefinitionArgs'] materialized_view: [Optional] Materialized view definition.
        :param pulumi.Input['ModelDefinitionArgs'] model: [Output-only, Beta] Present iff this table represents a ML model. Describes the training information for the model, and it is required to run 'PREDICT' queries.
        :param pulumi.Input['RangePartitioningArgs'] range_partitioning: [TrustedTester] Range partitioning specification for this table. Only one of timePartitioning and rangePartitioning should be specified.
        :param pulumi.Input[bool] require_partition_filter: [Optional] If set to true, queries over this table require a partition filter that can be used for partition elimination to be specified.
        :param pulumi.Input['TableSchemaArgs'] schema: [Optional] Describes the schema of this table.
        :param pulumi.Input['TableReferenceArgs'] table_reference: [Required] Reference describing the ID of this table.
        :param pulumi.Input['TimePartitioningArgs'] time_partitioning: Time-based partitioning specification for this table. Only one of timePartitioning and rangePartitioning should be specified.
        :param pulumi.Input['ViewDefinitionArgs'] view: [Optional] The view definition.
        """
        pulumi.set(__self__, "dataset_id", dataset_id)
        pulumi.set(__self__, "project", project)
        if clustering is not None:
            pulumi.set(__self__, "clustering", clustering)
        if description is not None:
            pulumi.set(__self__, "description", description)
        if encryption_configuration is not None:
            pulumi.set(__self__, "encryption_configuration", encryption_configuration)
        if expiration_time is not None:
            pulumi.set(__self__, "expiration_time", expiration_time)
        if external_data_configuration is not None:
            pulumi.set(__self__, "external_data_configuration", external_data_configuration)
        if friendly_name is not None:
            pulumi.set(__self__, "friendly_name", friendly_name)
        if labels is not None:
            pulumi.set(__self__, "labels", labels)
        if materialized_view is not None:
            pulumi.set(__self__, "materialized_view", materialized_view)
        if model is not None:
            pulumi.set(__self__, "model", model)
        if range_partitioning is not None:
            pulumi.set(__self__, "range_partitioning", range_partitioning)
        if require_partition_filter is not None:
            pulumi.set(__self__, "require_partition_filter", require_partition_filter)
        if schema is not None:
            pulumi.set(__self__, "schema", schema)
        if table_reference is not None:
            pulumi.set(__self__, "table_reference", table_reference)
        if time_partitioning is not None:
            pulumi.set(__self__, "time_partitioning", time_partitioning)
        if view is not None:
            pulumi.set(__self__, "view", view)

    @property
    @pulumi.getter(name="datasetId")
    def dataset_id(self) -> pulumi.Input[str]:
        return pulumi.get(self, "dataset_id")

    @dataset_id.setter
    def dataset_id(self, value: pulumi.Input[str]):
        pulumi.set(self, "dataset_id", value)

    @property
    @pulumi.getter
    def project(self) -> pulumi.Input[str]:
        return pulumi.get(self, "project")

    @project.setter
    def project(self, value: pulumi.Input[str]):
        pulumi.set(self, "project", value)

    @property
    @pulumi.getter
    def clustering(self) -> Optional[pulumi.Input['ClusteringArgs']]:
        """
        [Beta] Clustering specification for the table. Must be specified with partitioning, data in the table will be first partitioned and subsequently clustered.
        """
        return pulumi.get(self, "clustering")

    @clustering.setter
    def clustering(self, value: Optional[pulumi.Input['ClusteringArgs']]):
        pulumi.set(self, "clustering", value)

    @property
    @pulumi.getter
    def description(self) -> Optional[pulumi.Input[str]]:
        """
        [Optional] A user-friendly description of this table.
        """
        return pulumi.get(self, "description")

    @description.setter
    def description(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "description", value)

    @property
    @pulumi.getter(name="encryptionConfiguration")
    def encryption_configuration(self) -> Optional[pulumi.Input['EncryptionConfigurationArgs']]:
        """
        Custom encryption configuration (e.g., Cloud KMS keys).
        """
        return pulumi.get(self, "encryption_configuration")

    @encryption_configuration.setter
    def encryption_configuration(self, value: Optional[pulumi.Input['EncryptionConfigurationArgs']]):
        pulumi.set(self, "encryption_configuration", value)

    @property
    @pulumi.getter(name="expirationTime")
    def expiration_time(self) -> Optional[pulumi.Input[str]]:
        """
        [Optional] The time when this table expires, in milliseconds since the epoch. If not present, the table will persist indefinitely. Expired tables will be deleted and their storage reclaimed. The defaultTableExpirationMs property of the encapsulating dataset can be used to set a default expirationTime on newly created tables.
        """
        return pulumi.get(self, "expiration_time")

    @expiration_time.setter
    def expiration_time(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "expiration_time", value)

    @property
    @pulumi.getter(name="externalDataConfiguration")
    def external_data_configuration(self) -> Optional[pulumi.Input['ExternalDataConfigurationArgs']]:
        """
        [Optional] Describes the data format, location, and other properties of a table stored outside of BigQuery. By defining these properties, the data source can then be queried as if it were a standard BigQuery table.
        """
        return pulumi.get(self, "external_data_configuration")

    @external_data_configuration.setter
    def external_data_configuration(self, value: Optional[pulumi.Input['ExternalDataConfigurationArgs']]):
        pulumi.set(self, "external_data_configuration", value)

    @property
    @pulumi.getter(name="friendlyName")
    def friendly_name(self) -> Optional[pulumi.Input[str]]:
        """
        [Optional] A descriptive name for this table.
        """
        return pulumi.get(self, "friendly_name")

    @friendly_name.setter
    def friendly_name(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "friendly_name", value)

    @property
    @pulumi.getter
    def labels(self) -> Optional[pulumi.Input[Mapping[str, pulumi.Input[str]]]]:
        """
        The labels associated with this table. You can use these to organize and group your tables. Label keys and values can be no longer than 63 characters, can only contain lowercase letters, numeric characters, underscores and dashes. International characters are allowed. Label values are optional. Label keys must start with a letter and each label in the list must have a different key.
        """
        return pulumi.get(self, "labels")

    @labels.setter
    def labels(self, value: Optional[pulumi.Input[Mapping[str, pulumi.Input[str]]]]):
        pulumi.set(self, "labels", value)

    @property
    @pulumi.getter(name="materializedView")
    def materialized_view(self) -> Optional[pulumi.Input['MaterializedViewDefinitionArgs']]:
        """
        [Optional] Materialized view definition.
        """
        return pulumi.get(self, "materialized_view")

    @materialized_view.setter
    def materialized_view(self, value: Optional[pulumi.Input['MaterializedViewDefinitionArgs']]):
        pulumi.set(self, "materialized_view", value)

    @property
    @pulumi.getter
    def model(self) -> Optional[pulumi.Input['ModelDefinitionArgs']]:
        """
        [Output-only, Beta] Present iff this table represents a ML model. Describes the training information for the model, and it is required to run 'PREDICT' queries.
        """
        return pulumi.get(self, "model")

    @model.setter
    def model(self, value: Optional[pulumi.Input['ModelDefinitionArgs']]):
        pulumi.set(self, "model", value)

    @property
    @pulumi.getter(name="rangePartitioning")
    def range_partitioning(self) -> Optional[pulumi.Input['RangePartitioningArgs']]:
        """
        [TrustedTester] Range partitioning specification for this table. Only one of timePartitioning and rangePartitioning should be specified.
        """
        return pulumi.get(self, "range_partitioning")

    @range_partitioning.setter
    def range_partitioning(self, value: Optional[pulumi.Input['RangePartitioningArgs']]):
        pulumi.set(self, "range_partitioning", value)

    @property
    @pulumi.getter(name="requirePartitionFilter")
    def require_partition_filter(self) -> Optional[pulumi.Input[bool]]:
        """
        [Optional] If set to true, queries over this table require a partition filter that can be used for partition elimination to be specified.
        """
        return pulumi.get(self, "require_partition_filter")

    @require_partition_filter.setter
    def require_partition_filter(self, value: Optional[pulumi.Input[bool]]):
        pulumi.set(self, "require_partition_filter", value)

    @property
    @pulumi.getter
    def schema(self) -> Optional[pulumi.Input['TableSchemaArgs']]:
        """
        [Optional] Describes the schema of this table.
        """
        return pulumi.get(self, "schema")

    @schema.setter
    def schema(self, value: Optional[pulumi.Input['TableSchemaArgs']]):
        pulumi.set(self, "schema", value)

    @property
    @pulumi.getter(name="tableReference")
    def table_reference(self) -> Optional[pulumi.Input['TableReferenceArgs']]:
        """
        [Required] Reference describing the ID of this table.
        """
        return pulumi.get(self, "table_reference")

    @table_reference.setter
    def table_reference(self, value: Optional[pulumi.Input['TableReferenceArgs']]):
        pulumi.set(self, "table_reference", value)

    @property
    @pulumi.getter(name="timePartitioning")
    def time_partitioning(self) -> Optional[pulumi.Input['TimePartitioningArgs']]:
        """
        Time-based partitioning specification for this table. Only one of timePartitioning and rangePartitioning should be specified.
        """
        return pulumi.get(self, "time_partitioning")

    @time_partitioning.setter
    def time_partitioning(self, value: Optional[pulumi.Input['TimePartitioningArgs']]):
        pulumi.set(self, "time_partitioning", value)

    @property
    @pulumi.getter
    def view(self) -> Optional[pulumi.Input['ViewDefinitionArgs']]:
        """
        [Optional] The view definition.
        """
        return pulumi.get(self, "view")

    @view.setter
    def view(self, value: Optional[pulumi.Input['ViewDefinitionArgs']]):
        pulumi.set(self, "view", value)


class Table(pulumi.CustomResource):
    @overload
    def __init__(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 clustering: Optional[pulumi.Input[pulumi.InputType['ClusteringArgs']]] = None,
                 dataset_id: Optional[pulumi.Input[str]] = None,
                 description: Optional[pulumi.Input[str]] = None,
                 encryption_configuration: Optional[pulumi.Input[pulumi.InputType['EncryptionConfigurationArgs']]] = None,
                 expiration_time: Optional[pulumi.Input[str]] = None,
                 external_data_configuration: Optional[pulumi.Input[pulumi.InputType['ExternalDataConfigurationArgs']]] = None,
                 friendly_name: Optional[pulumi.Input[str]] = None,
                 labels: Optional[pulumi.Input[Mapping[str, pulumi.Input[str]]]] = None,
                 materialized_view: Optional[pulumi.Input[pulumi.InputType['MaterializedViewDefinitionArgs']]] = None,
                 model: Optional[pulumi.Input[pulumi.InputType['ModelDefinitionArgs']]] = None,
                 project: Optional[pulumi.Input[str]] = None,
                 range_partitioning: Optional[pulumi.Input[pulumi.InputType['RangePartitioningArgs']]] = None,
                 require_partition_filter: Optional[pulumi.Input[bool]] = None,
                 schema: Optional[pulumi.Input[pulumi.InputType['TableSchemaArgs']]] = None,
                 table_reference: Optional[pulumi.Input[pulumi.InputType['TableReferenceArgs']]] = None,
                 time_partitioning: Optional[pulumi.Input[pulumi.InputType['TimePartitioningArgs']]] = None,
                 view: Optional[pulumi.Input[pulumi.InputType['ViewDefinitionArgs']]] = None,
                 __props__=None):
        """
        Creates a new, empty table in the dataset.

        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[pulumi.InputType['ClusteringArgs']] clustering: [Beta] Clustering specification for the table. Must be specified with partitioning, data in the table will be first partitioned and subsequently clustered.
        :param pulumi.Input[str] description: [Optional] A user-friendly description of this table.
        :param pulumi.Input[pulumi.InputType['EncryptionConfigurationArgs']] encryption_configuration: Custom encryption configuration (e.g., Cloud KMS keys).
        :param pulumi.Input[str] expiration_time: [Optional] The time when this table expires, in milliseconds since the epoch. If not present, the table will persist indefinitely. Expired tables will be deleted and their storage reclaimed. The defaultTableExpirationMs property of the encapsulating dataset can be used to set a default expirationTime on newly created tables.
        :param pulumi.Input[pulumi.InputType['ExternalDataConfigurationArgs']] external_data_configuration: [Optional] Describes the data format, location, and other properties of a table stored outside of BigQuery. By defining these properties, the data source can then be queried as if it were a standard BigQuery table.
        :param pulumi.Input[str] friendly_name: [Optional] A descriptive name for this table.
        :param pulumi.Input[Mapping[str, pulumi.Input[str]]] labels: The labels associated with this table. You can use these to organize and group your tables. Label keys and values can be no longer than 63 characters, can only contain lowercase letters, numeric characters, underscores and dashes. International characters are allowed. Label values are optional. Label keys must start with a letter and each label in the list must have a different key.
        :param pulumi.Input[pulumi.InputType['MaterializedViewDefinitionArgs']] materialized_view: [Optional] Materialized view definition.
        :param pulumi.Input[pulumi.InputType['ModelDefinitionArgs']] model: [Output-only, Beta] Present iff this table represents a ML model. Describes the training information for the model, and it is required to run 'PREDICT' queries.
        :param pulumi.Input[pulumi.InputType['RangePartitioningArgs']] range_partitioning: [TrustedTester] Range partitioning specification for this table. Only one of timePartitioning and rangePartitioning should be specified.
        :param pulumi.Input[bool] require_partition_filter: [Optional] If set to true, queries over this table require a partition filter that can be used for partition elimination to be specified.
        :param pulumi.Input[pulumi.InputType['TableSchemaArgs']] schema: [Optional] Describes the schema of this table.
        :param pulumi.Input[pulumi.InputType['TableReferenceArgs']] table_reference: [Required] Reference describing the ID of this table.
        :param pulumi.Input[pulumi.InputType['TimePartitioningArgs']] time_partitioning: Time-based partitioning specification for this table. Only one of timePartitioning and rangePartitioning should be specified.
        :param pulumi.Input[pulumi.InputType['ViewDefinitionArgs']] view: [Optional] The view definition.
        """
        ...
    @overload
    def __init__(__self__,
                 resource_name: str,
                 args: TableArgs,
                 opts: Optional[pulumi.ResourceOptions] = None):
        """
        Creates a new, empty table in the dataset.

        :param str resource_name: The name of the resource.
        :param TableArgs args: The arguments to use to populate this resource's properties.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        ...
    def __init__(__self__, resource_name: str, *args, **kwargs):
        resource_args, opts = _utilities.get_resource_args_opts(TableArgs, pulumi.ResourceOptions, *args, **kwargs)
        if resource_args is not None:
            __self__._internal_init(resource_name, opts, **resource_args.__dict__)
        else:
            __self__._internal_init(resource_name, *args, **kwargs)

    def _internal_init(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 clustering: Optional[pulumi.Input[pulumi.InputType['ClusteringArgs']]] = None,
                 dataset_id: Optional[pulumi.Input[str]] = None,
                 description: Optional[pulumi.Input[str]] = None,
                 encryption_configuration: Optional[pulumi.Input[pulumi.InputType['EncryptionConfigurationArgs']]] = None,
                 expiration_time: Optional[pulumi.Input[str]] = None,
                 external_data_configuration: Optional[pulumi.Input[pulumi.InputType['ExternalDataConfigurationArgs']]] = None,
                 friendly_name: Optional[pulumi.Input[str]] = None,
                 labels: Optional[pulumi.Input[Mapping[str, pulumi.Input[str]]]] = None,
                 materialized_view: Optional[pulumi.Input[pulumi.InputType['MaterializedViewDefinitionArgs']]] = None,
                 model: Optional[pulumi.Input[pulumi.InputType['ModelDefinitionArgs']]] = None,
                 project: Optional[pulumi.Input[str]] = None,
                 range_partitioning: Optional[pulumi.Input[pulumi.InputType['RangePartitioningArgs']]] = None,
                 require_partition_filter: Optional[pulumi.Input[bool]] = None,
                 schema: Optional[pulumi.Input[pulumi.InputType['TableSchemaArgs']]] = None,
                 table_reference: Optional[pulumi.Input[pulumi.InputType['TableReferenceArgs']]] = None,
                 time_partitioning: Optional[pulumi.Input[pulumi.InputType['TimePartitioningArgs']]] = None,
                 view: Optional[pulumi.Input[pulumi.InputType['ViewDefinitionArgs']]] = None,
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
            __props__ = TableArgs.__new__(TableArgs)

            __props__.__dict__["clustering"] = clustering
            if dataset_id is None and not opts.urn:
                raise TypeError("Missing required property 'dataset_id'")
            __props__.__dict__["dataset_id"] = dataset_id
            __props__.__dict__["description"] = description
            __props__.__dict__["encryption_configuration"] = encryption_configuration
            __props__.__dict__["expiration_time"] = expiration_time
            __props__.__dict__["external_data_configuration"] = external_data_configuration
            __props__.__dict__["friendly_name"] = friendly_name
            __props__.__dict__["labels"] = labels
            __props__.__dict__["materialized_view"] = materialized_view
            __props__.__dict__["model"] = model
            if project is None and not opts.urn:
                raise TypeError("Missing required property 'project'")
            __props__.__dict__["project"] = project
            __props__.__dict__["range_partitioning"] = range_partitioning
            __props__.__dict__["require_partition_filter"] = require_partition_filter
            __props__.__dict__["schema"] = schema
            __props__.__dict__["table_reference"] = table_reference
            __props__.__dict__["time_partitioning"] = time_partitioning
            __props__.__dict__["view"] = view
            __props__.__dict__["creation_time"] = None
            __props__.__dict__["etag"] = None
            __props__.__dict__["kind"] = None
            __props__.__dict__["last_modified_time"] = None
            __props__.__dict__["location"] = None
            __props__.__dict__["num_bytes"] = None
            __props__.__dict__["num_long_term_bytes"] = None
            __props__.__dict__["num_physical_bytes"] = None
            __props__.__dict__["num_rows"] = None
            __props__.__dict__["self_link"] = None
            __props__.__dict__["snapshot_definition"] = None
            __props__.__dict__["streaming_buffer"] = None
            __props__.__dict__["type"] = None
        super(Table, __self__).__init__(
            'google-native:bigquery/v2:Table',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name: str,
            id: pulumi.Input[str],
            opts: Optional[pulumi.ResourceOptions] = None) -> 'Table':
        """
        Get an existing Table resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param pulumi.Input[str] id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = TableArgs.__new__(TableArgs)

        __props__.__dict__["clustering"] = None
        __props__.__dict__["creation_time"] = None
        __props__.__dict__["description"] = None
        __props__.__dict__["encryption_configuration"] = None
        __props__.__dict__["etag"] = None
        __props__.__dict__["expiration_time"] = None
        __props__.__dict__["external_data_configuration"] = None
        __props__.__dict__["friendly_name"] = None
        __props__.__dict__["kind"] = None
        __props__.__dict__["labels"] = None
        __props__.__dict__["last_modified_time"] = None
        __props__.__dict__["location"] = None
        __props__.__dict__["materialized_view"] = None
        __props__.__dict__["model"] = None
        __props__.__dict__["num_bytes"] = None
        __props__.__dict__["num_long_term_bytes"] = None
        __props__.__dict__["num_physical_bytes"] = None
        __props__.__dict__["num_rows"] = None
        __props__.__dict__["range_partitioning"] = None
        __props__.__dict__["require_partition_filter"] = None
        __props__.__dict__["schema"] = None
        __props__.__dict__["self_link"] = None
        __props__.__dict__["snapshot_definition"] = None
        __props__.__dict__["streaming_buffer"] = None
        __props__.__dict__["table_reference"] = None
        __props__.__dict__["time_partitioning"] = None
        __props__.__dict__["type"] = None
        __props__.__dict__["view"] = None
        return Table(resource_name, opts=opts, __props__=__props__)

    @property
    @pulumi.getter
    def clustering(self) -> pulumi.Output['outputs.ClusteringResponse']:
        """
        [Beta] Clustering specification for the table. Must be specified with partitioning, data in the table will be first partitioned and subsequently clustered.
        """
        return pulumi.get(self, "clustering")

    @property
    @pulumi.getter(name="creationTime")
    def creation_time(self) -> pulumi.Output[str]:
        """
        The time when this table was created, in milliseconds since the epoch.
        """
        return pulumi.get(self, "creation_time")

    @property
    @pulumi.getter
    def description(self) -> pulumi.Output[str]:
        """
        [Optional] A user-friendly description of this table.
        """
        return pulumi.get(self, "description")

    @property
    @pulumi.getter(name="encryptionConfiguration")
    def encryption_configuration(self) -> pulumi.Output['outputs.EncryptionConfigurationResponse']:
        """
        Custom encryption configuration (e.g., Cloud KMS keys).
        """
        return pulumi.get(self, "encryption_configuration")

    @property
    @pulumi.getter
    def etag(self) -> pulumi.Output[str]:
        """
        A hash of the table metadata. Used to ensure there were no concurrent modifications to the resource when attempting an update. Not guaranteed to change when the table contents or the fields numRows, numBytes, numLongTermBytes or lastModifiedTime change.
        """
        return pulumi.get(self, "etag")

    @property
    @pulumi.getter(name="expirationTime")
    def expiration_time(self) -> pulumi.Output[str]:
        """
        [Optional] The time when this table expires, in milliseconds since the epoch. If not present, the table will persist indefinitely. Expired tables will be deleted and their storage reclaimed. The defaultTableExpirationMs property of the encapsulating dataset can be used to set a default expirationTime on newly created tables.
        """
        return pulumi.get(self, "expiration_time")

    @property
    @pulumi.getter(name="externalDataConfiguration")
    def external_data_configuration(self) -> pulumi.Output['outputs.ExternalDataConfigurationResponse']:
        """
        [Optional] Describes the data format, location, and other properties of a table stored outside of BigQuery. By defining these properties, the data source can then be queried as if it were a standard BigQuery table.
        """
        return pulumi.get(self, "external_data_configuration")

    @property
    @pulumi.getter(name="friendlyName")
    def friendly_name(self) -> pulumi.Output[str]:
        """
        [Optional] A descriptive name for this table.
        """
        return pulumi.get(self, "friendly_name")

    @property
    @pulumi.getter
    def kind(self) -> pulumi.Output[str]:
        """
        The type of the resource.
        """
        return pulumi.get(self, "kind")

    @property
    @pulumi.getter
    def labels(self) -> pulumi.Output[Mapping[str, str]]:
        """
        The labels associated with this table. You can use these to organize and group your tables. Label keys and values can be no longer than 63 characters, can only contain lowercase letters, numeric characters, underscores and dashes. International characters are allowed. Label values are optional. Label keys must start with a letter and each label in the list must have a different key.
        """
        return pulumi.get(self, "labels")

    @property
    @pulumi.getter(name="lastModifiedTime")
    def last_modified_time(self) -> pulumi.Output[str]:
        """
        The time when this table was last modified, in milliseconds since the epoch.
        """
        return pulumi.get(self, "last_modified_time")

    @property
    @pulumi.getter
    def location(self) -> pulumi.Output[str]:
        """
        The geographic location where the table resides. This value is inherited from the dataset.
        """
        return pulumi.get(self, "location")

    @property
    @pulumi.getter(name="materializedView")
    def materialized_view(self) -> pulumi.Output['outputs.MaterializedViewDefinitionResponse']:
        """
        [Optional] Materialized view definition.
        """
        return pulumi.get(self, "materialized_view")

    @property
    @pulumi.getter
    def model(self) -> pulumi.Output['outputs.ModelDefinitionResponse']:
        """
        [Output-only, Beta] Present iff this table represents a ML model. Describes the training information for the model, and it is required to run 'PREDICT' queries.
        """
        return pulumi.get(self, "model")

    @property
    @pulumi.getter(name="numBytes")
    def num_bytes(self) -> pulumi.Output[str]:
        """
        The size of this table in bytes, excluding any data in the streaming buffer.
        """
        return pulumi.get(self, "num_bytes")

    @property
    @pulumi.getter(name="numLongTermBytes")
    def num_long_term_bytes(self) -> pulumi.Output[str]:
        """
        The number of bytes in the table that are considered "long-term storage".
        """
        return pulumi.get(self, "num_long_term_bytes")

    @property
    @pulumi.getter(name="numPhysicalBytes")
    def num_physical_bytes(self) -> pulumi.Output[str]:
        """
        [TrustedTester] The physical size of this table in bytes, excluding any data in the streaming buffer. This includes compression and storage used for time travel.
        """
        return pulumi.get(self, "num_physical_bytes")

    @property
    @pulumi.getter(name="numRows")
    def num_rows(self) -> pulumi.Output[str]:
        """
        The number of rows of data in this table, excluding any data in the streaming buffer.
        """
        return pulumi.get(self, "num_rows")

    @property
    @pulumi.getter(name="rangePartitioning")
    def range_partitioning(self) -> pulumi.Output['outputs.RangePartitioningResponse']:
        """
        [TrustedTester] Range partitioning specification for this table. Only one of timePartitioning and rangePartitioning should be specified.
        """
        return pulumi.get(self, "range_partitioning")

    @property
    @pulumi.getter(name="requirePartitionFilter")
    def require_partition_filter(self) -> pulumi.Output[bool]:
        """
        [Optional] If set to true, queries over this table require a partition filter that can be used for partition elimination to be specified.
        """
        return pulumi.get(self, "require_partition_filter")

    @property
    @pulumi.getter
    def schema(self) -> pulumi.Output['outputs.TableSchemaResponse']:
        """
        [Optional] Describes the schema of this table.
        """
        return pulumi.get(self, "schema")

    @property
    @pulumi.getter(name="selfLink")
    def self_link(self) -> pulumi.Output[str]:
        """
        A URL that can be used to access this resource again.
        """
        return pulumi.get(self, "self_link")

    @property
    @pulumi.getter(name="snapshotDefinition")
    def snapshot_definition(self) -> pulumi.Output['outputs.SnapshotDefinitionResponse']:
        """
        Snapshot definition.
        """
        return pulumi.get(self, "snapshot_definition")

    @property
    @pulumi.getter(name="streamingBuffer")
    def streaming_buffer(self) -> pulumi.Output['outputs.StreamingbufferResponse']:
        """
        Contains information regarding this table's streaming buffer, if one is present. This field will be absent if the table is not being streamed to or if there is no data in the streaming buffer.
        """
        return pulumi.get(self, "streaming_buffer")

    @property
    @pulumi.getter(name="tableReference")
    def table_reference(self) -> pulumi.Output['outputs.TableReferenceResponse']:
        """
        [Required] Reference describing the ID of this table.
        """
        return pulumi.get(self, "table_reference")

    @property
    @pulumi.getter(name="timePartitioning")
    def time_partitioning(self) -> pulumi.Output['outputs.TimePartitioningResponse']:
        """
        Time-based partitioning specification for this table. Only one of timePartitioning and rangePartitioning should be specified.
        """
        return pulumi.get(self, "time_partitioning")

    @property
    @pulumi.getter
    def type(self) -> pulumi.Output[str]:
        """
        Describes the table type. The following values are supported: TABLE: A normal BigQuery table. VIEW: A virtual table defined by a SQL query. SNAPSHOT: An immutable, read-only table that is a copy of another table. [TrustedTester] MATERIALIZED_VIEW: SQL query whose result is persisted. EXTERNAL: A table that references data stored in an external storage system, such as Google Cloud Storage. The default value is TABLE.
        """
        return pulumi.get(self, "type")

    @property
    @pulumi.getter
    def view(self) -> pulumi.Output['outputs.ViewDefinitionResponse']:
        """
        [Optional] The view definition.
        """
        return pulumi.get(self, "view")

