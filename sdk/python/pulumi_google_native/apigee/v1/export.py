# coding=utf-8
# *** WARNING: this file was generated by the Pulumi SDK Generator. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import warnings
import pulumi
import pulumi.runtime
from typing import Any, Mapping, Optional, Sequence, Union, overload
from ... import _utilities
from ._inputs import *

__all__ = ['ExportArgs', 'Export']

@pulumi.input_type
class ExportArgs:
    def __init__(__self__, *,
                 datastore_name: pulumi.Input[str],
                 date_range: pulumi.Input['GoogleCloudApigeeV1DateRangeArgs'],
                 environment_id: pulumi.Input[str],
                 name: pulumi.Input[str],
                 organization_id: pulumi.Input[str],
                 csv_delimiter: Optional[pulumi.Input[str]] = None,
                 description: Optional[pulumi.Input[str]] = None,
                 output_format: Optional[pulumi.Input[str]] = None):
        """
        The set of arguments for constructing a Export resource.
        :param pulumi.Input[str] datastore_name: Name of the preconfigured datastore.
        :param pulumi.Input['GoogleCloudApigeeV1DateRangeArgs'] date_range: Date range of the data to export.
        :param pulumi.Input[str] name: Display name of the export job.
        :param pulumi.Input[str] csv_delimiter: Optional. Delimiter used in the CSV file, if `outputFormat` is set to `csv`. Defaults to the `,` (comma) character. Supported delimiter characters include comma (`,`), pipe (`|`), and tab (`\t`).
        :param pulumi.Input[str] description: Optional. Description of the export job.
        :param pulumi.Input[str] output_format: Optional. Output format of the export. Valid values include: `csv` or `json`. Defaults to `json`. Note: Configure the delimiter for CSV output using the `csvDelimiter` property.
        """
        pulumi.set(__self__, "datastore_name", datastore_name)
        pulumi.set(__self__, "date_range", date_range)
        pulumi.set(__self__, "environment_id", environment_id)
        pulumi.set(__self__, "name", name)
        pulumi.set(__self__, "organization_id", organization_id)
        if csv_delimiter is not None:
            pulumi.set(__self__, "csv_delimiter", csv_delimiter)
        if description is not None:
            pulumi.set(__self__, "description", description)
        if output_format is not None:
            pulumi.set(__self__, "output_format", output_format)

    @property
    @pulumi.getter(name="datastoreName")
    def datastore_name(self) -> pulumi.Input[str]:
        """
        Name of the preconfigured datastore.
        """
        return pulumi.get(self, "datastore_name")

    @datastore_name.setter
    def datastore_name(self, value: pulumi.Input[str]):
        pulumi.set(self, "datastore_name", value)

    @property
    @pulumi.getter(name="dateRange")
    def date_range(self) -> pulumi.Input['GoogleCloudApigeeV1DateRangeArgs']:
        """
        Date range of the data to export.
        """
        return pulumi.get(self, "date_range")

    @date_range.setter
    def date_range(self, value: pulumi.Input['GoogleCloudApigeeV1DateRangeArgs']):
        pulumi.set(self, "date_range", value)

    @property
    @pulumi.getter(name="environmentId")
    def environment_id(self) -> pulumi.Input[str]:
        return pulumi.get(self, "environment_id")

    @environment_id.setter
    def environment_id(self, value: pulumi.Input[str]):
        pulumi.set(self, "environment_id", value)

    @property
    @pulumi.getter
    def name(self) -> pulumi.Input[str]:
        """
        Display name of the export job.
        """
        return pulumi.get(self, "name")

    @name.setter
    def name(self, value: pulumi.Input[str]):
        pulumi.set(self, "name", value)

    @property
    @pulumi.getter(name="organizationId")
    def organization_id(self) -> pulumi.Input[str]:
        return pulumi.get(self, "organization_id")

    @organization_id.setter
    def organization_id(self, value: pulumi.Input[str]):
        pulumi.set(self, "organization_id", value)

    @property
    @pulumi.getter(name="csvDelimiter")
    def csv_delimiter(self) -> Optional[pulumi.Input[str]]:
        """
        Optional. Delimiter used in the CSV file, if `outputFormat` is set to `csv`. Defaults to the `,` (comma) character. Supported delimiter characters include comma (`,`), pipe (`|`), and tab (`\t`).
        """
        return pulumi.get(self, "csv_delimiter")

    @csv_delimiter.setter
    def csv_delimiter(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "csv_delimiter", value)

    @property
    @pulumi.getter
    def description(self) -> Optional[pulumi.Input[str]]:
        """
        Optional. Description of the export job.
        """
        return pulumi.get(self, "description")

    @description.setter
    def description(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "description", value)

    @property
    @pulumi.getter(name="outputFormat")
    def output_format(self) -> Optional[pulumi.Input[str]]:
        """
        Optional. Output format of the export. Valid values include: `csv` or `json`. Defaults to `json`. Note: Configure the delimiter for CSV output using the `csvDelimiter` property.
        """
        return pulumi.get(self, "output_format")

    @output_format.setter
    def output_format(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "output_format", value)


class Export(pulumi.CustomResource):
    @overload
    def __init__(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 csv_delimiter: Optional[pulumi.Input[str]] = None,
                 datastore_name: Optional[pulumi.Input[str]] = None,
                 date_range: Optional[pulumi.Input[pulumi.InputType['GoogleCloudApigeeV1DateRangeArgs']]] = None,
                 description: Optional[pulumi.Input[str]] = None,
                 environment_id: Optional[pulumi.Input[str]] = None,
                 name: Optional[pulumi.Input[str]] = None,
                 organization_id: Optional[pulumi.Input[str]] = None,
                 output_format: Optional[pulumi.Input[str]] = None,
                 __props__=None):
        """
        Submit a data export job to be processed in the background. If the request is successful, the API returns a 201 status, a URI that can be used to retrieve the status of the export job, and the `state` value of "enqueued".

        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] csv_delimiter: Optional. Delimiter used in the CSV file, if `outputFormat` is set to `csv`. Defaults to the `,` (comma) character. Supported delimiter characters include comma (`,`), pipe (`|`), and tab (`\t`).
        :param pulumi.Input[str] datastore_name: Name of the preconfigured datastore.
        :param pulumi.Input[pulumi.InputType['GoogleCloudApigeeV1DateRangeArgs']] date_range: Date range of the data to export.
        :param pulumi.Input[str] description: Optional. Description of the export job.
        :param pulumi.Input[str] name: Display name of the export job.
        :param pulumi.Input[str] output_format: Optional. Output format of the export. Valid values include: `csv` or `json`. Defaults to `json`. Note: Configure the delimiter for CSV output using the `csvDelimiter` property.
        """
        ...
    @overload
    def __init__(__self__,
                 resource_name: str,
                 args: ExportArgs,
                 opts: Optional[pulumi.ResourceOptions] = None):
        """
        Submit a data export job to be processed in the background. If the request is successful, the API returns a 201 status, a URI that can be used to retrieve the status of the export job, and the `state` value of "enqueued".

        :param str resource_name: The name of the resource.
        :param ExportArgs args: The arguments to use to populate this resource's properties.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        ...
    def __init__(__self__, resource_name: str, *args, **kwargs):
        resource_args, opts = _utilities.get_resource_args_opts(ExportArgs, pulumi.ResourceOptions, *args, **kwargs)
        if resource_args is not None:
            __self__._internal_init(resource_name, opts, **resource_args.__dict__)
        else:
            __self__._internal_init(resource_name, *args, **kwargs)

    def _internal_init(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 csv_delimiter: Optional[pulumi.Input[str]] = None,
                 datastore_name: Optional[pulumi.Input[str]] = None,
                 date_range: Optional[pulumi.Input[pulumi.InputType['GoogleCloudApigeeV1DateRangeArgs']]] = None,
                 description: Optional[pulumi.Input[str]] = None,
                 environment_id: Optional[pulumi.Input[str]] = None,
                 name: Optional[pulumi.Input[str]] = None,
                 organization_id: Optional[pulumi.Input[str]] = None,
                 output_format: Optional[pulumi.Input[str]] = None,
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
            __props__ = ExportArgs.__new__(ExportArgs)

            __props__.__dict__["csv_delimiter"] = csv_delimiter
            if datastore_name is None and not opts.urn:
                raise TypeError("Missing required property 'datastore_name'")
            __props__.__dict__["datastore_name"] = datastore_name
            if date_range is None and not opts.urn:
                raise TypeError("Missing required property 'date_range'")
            __props__.__dict__["date_range"] = date_range
            __props__.__dict__["description"] = description
            if environment_id is None and not opts.urn:
                raise TypeError("Missing required property 'environment_id'")
            __props__.__dict__["environment_id"] = environment_id
            if name is None and not opts.urn:
                raise TypeError("Missing required property 'name'")
            __props__.__dict__["name"] = name
            if organization_id is None and not opts.urn:
                raise TypeError("Missing required property 'organization_id'")
            __props__.__dict__["organization_id"] = organization_id
            __props__.__dict__["output_format"] = output_format
            __props__.__dict__["created"] = None
            __props__.__dict__["error"] = None
            __props__.__dict__["execution_time"] = None
            __props__.__dict__["self"] = None
            __props__.__dict__["state"] = None
            __props__.__dict__["updated"] = None
        super(Export, __self__).__init__(
            'google-native:apigee/v1:Export',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name: str,
            id: pulumi.Input[str],
            opts: Optional[pulumi.ResourceOptions] = None) -> 'Export':
        """
        Get an existing Export resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param pulumi.Input[str] id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = ExportArgs.__new__(ExportArgs)

        __props__.__dict__["created"] = None
        __props__.__dict__["datastore_name"] = None
        __props__.__dict__["description"] = None
        __props__.__dict__["error"] = None
        __props__.__dict__["execution_time"] = None
        __props__.__dict__["name"] = None
        __props__.__dict__["self"] = None
        __props__.__dict__["state"] = None
        __props__.__dict__["updated"] = None
        return Export(resource_name, opts=opts, __props__=__props__)

    @property
    @pulumi.getter
    def created(self) -> pulumi.Output[str]:
        """
        Time the export job was created.
        """
        return pulumi.get(self, "created")

    @property
    @pulumi.getter(name="datastoreName")
    def datastore_name(self) -> pulumi.Output[str]:
        """
        Name of the datastore that is the destination of the export job [datastore]
        """
        return pulumi.get(self, "datastore_name")

    @property
    @pulumi.getter
    def description(self) -> pulumi.Output[str]:
        """
        Description of the export job.
        """
        return pulumi.get(self, "description")

    @property
    @pulumi.getter
    def error(self) -> pulumi.Output[str]:
        """
        Error is set when export fails
        """
        return pulumi.get(self, "error")

    @property
    @pulumi.getter(name="executionTime")
    def execution_time(self) -> pulumi.Output[str]:
        """
        Execution time for this export job. If the job is still in progress, it will be set to the amount of time that has elapsed since`created`, in seconds. Else, it will set to (`updated` - `created`), in seconds.
        """
        return pulumi.get(self, "execution_time")

    @property
    @pulumi.getter
    def name(self) -> pulumi.Output[str]:
        """
        Display name of the export job.
        """
        return pulumi.get(self, "name")

    @property
    @pulumi.getter
    def self(self) -> pulumi.Output[str]:
        """
        Self link of the export job. A URI that can be used to retrieve the status of an export job. Example: `/organizations/myorg/environments/myenv/analytics/exports/9cfc0d85-0f30-46d6-ae6f-318d0cb961bd`
        """
        return pulumi.get(self, "self")

    @property
    @pulumi.getter
    def state(self) -> pulumi.Output[str]:
        """
        Status of the export job. Valid values include `enqueued`, `running`, `completed`, and `failed`.
        """
        return pulumi.get(self, "state")

    @property
    @pulumi.getter
    def updated(self) -> pulumi.Output[str]:
        """
        Time the export job was last updated.
        """
        return pulumi.get(self, "updated")

