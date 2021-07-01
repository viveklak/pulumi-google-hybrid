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

__all__ = ['FhirStoreArgs', 'FhirStore']

@pulumi.input_type
class FhirStoreArgs:
    def __init__(__self__, *,
                 dataset_id: pulumi.Input[str],
                 location: pulumi.Input[str],
                 project: pulumi.Input[str],
                 default_search_handling_strict: Optional[pulumi.Input[bool]] = None,
                 disable_referential_integrity: Optional[pulumi.Input[bool]] = None,
                 disable_resource_versioning: Optional[pulumi.Input[bool]] = None,
                 enable_update_create: Optional[pulumi.Input[bool]] = None,
                 fhir_store_id: Optional[pulumi.Input[str]] = None,
                 labels: Optional[pulumi.Input[Mapping[str, str]]] = None,
                 notification_config: Optional[pulumi.Input['NotificationConfigArgs']] = None,
                 stream_configs: Optional[pulumi.Input[Sequence[pulumi.Input['StreamConfigArgs']]]] = None,
                 validation_config: Optional[pulumi.Input['ValidationConfigArgs']] = None,
                 version: Optional[pulumi.Input['FhirStoreVersion']] = None):
        """
        The set of arguments for constructing a FhirStore resource.
        :param pulumi.Input[bool] default_search_handling_strict: If true, overrides the default search behavior for this FHIR store to `handling=strict` which returns an error for unrecognized search parameters. If false, uses the FHIR specification default `handling=lenient` which ignores unrecognized search parameters. The handling can always be changed from the default on an individual API call by setting the HTTP header `Prefer: handling=strict` or `Prefer: handling=lenient`.
        :param pulumi.Input[bool] disable_referential_integrity: Immutable. Whether to disable referential integrity in this FHIR store. This field is immutable after FHIR store creation. The default value is false, meaning that the API enforces referential integrity and fails the requests that result in inconsistent state in the FHIR store. When this field is set to true, the API skips referential integrity checks. Consequently, operations that rely on references, such as GetPatientEverything, do not return all the results if broken references exist.
        :param pulumi.Input[bool] disable_resource_versioning: Immutable. Whether to disable resource versioning for this FHIR store. This field can not be changed after the creation of FHIR store. If set to false, which is the default behavior, all write operations cause historical versions to be recorded automatically. The historical versions can be fetched through the history APIs, but cannot be updated. If set to true, no historical versions are kept. The server sends errors for attempts to read the historical versions.
        :param pulumi.Input[bool] enable_update_create: Whether this FHIR store has the [updateCreate capability](https://www.hl7.org/fhir/capabilitystatement-definitions.html#CapabilityStatement.rest.resource.updateCreate). This determines if the client can use an Update operation to create a new resource with a client-specified ID. If false, all IDs are server-assigned through the Create operation and attempts to update a non-existent resource return errors. It is strongly advised not to include or encode any sensitive data such as patient identifiers in client-specified resource IDs. Those IDs are part of the FHIR resource path recorded in Cloud audit logs and Pub/Sub notifications. Those IDs can also be contained in reference fields within other resources.
        :param pulumi.Input[Mapping[str, str]] labels: User-supplied key-value pairs used to organize FHIR stores. Label keys must be between 1 and 63 characters long, have a UTF-8 encoding of maximum 128 bytes, and must conform to the following PCRE regular expression: \p{Ll}\p{Lo}{0,62} Label values are optional, must be between 1 and 63 characters long, have a UTF-8 encoding of maximum 128 bytes, and must conform to the following PCRE regular expression: [\p{Ll}\p{Lo}\p{N}_-]{0,63} No more than 64 labels can be associated with a given store.
        :param pulumi.Input['NotificationConfigArgs'] notification_config: If non-empty, publish all resource modifications of this FHIR store to this destination. The Pub/Sub message attributes contain a map with a string describing the action that has triggered the notification. For example, "action":"CreateResource".
        :param pulumi.Input[Sequence[pulumi.Input['StreamConfigArgs']]] stream_configs: A list of streaming configs that configure the destinations of streaming export for every resource mutation in this FHIR store. Each store is allowed to have up to 10 streaming configs. After a new config is added, the next resource mutation is streamed to the new location in addition to the existing ones. When a location is removed from the list, the server stops streaming to that location. Before adding a new config, you must add the required [`bigquery.dataEditor`](https://cloud.google.com/bigquery/docs/access-control#bigquery.dataEditor) role to your project's **Cloud Healthcare Service Agent** [service account](https://cloud.google.com/iam/docs/service-accounts). Some lag (typically on the order of dozens of seconds) is expected before the results show up in the streaming destination.
        :param pulumi.Input['ValidationConfigArgs'] validation_config: Configuration for how to validate incoming FHIR resources against configured profiles.
        :param pulumi.Input['FhirStoreVersion'] version: Immutable. The FHIR specification version that this FHIR store supports natively. This field is immutable after store creation. Requests are rejected if they contain FHIR resources of a different version. Version is required for every FHIR store.
        """
        pulumi.set(__self__, "dataset_id", dataset_id)
        pulumi.set(__self__, "location", location)
        pulumi.set(__self__, "project", project)
        if default_search_handling_strict is not None:
            pulumi.set(__self__, "default_search_handling_strict", default_search_handling_strict)
        if disable_referential_integrity is not None:
            pulumi.set(__self__, "disable_referential_integrity", disable_referential_integrity)
        if disable_resource_versioning is not None:
            pulumi.set(__self__, "disable_resource_versioning", disable_resource_versioning)
        if enable_update_create is not None:
            pulumi.set(__self__, "enable_update_create", enable_update_create)
        if fhir_store_id is not None:
            pulumi.set(__self__, "fhir_store_id", fhir_store_id)
        if labels is not None:
            pulumi.set(__self__, "labels", labels)
        if notification_config is not None:
            pulumi.set(__self__, "notification_config", notification_config)
        if stream_configs is not None:
            pulumi.set(__self__, "stream_configs", stream_configs)
        if validation_config is not None:
            pulumi.set(__self__, "validation_config", validation_config)
        if version is not None:
            pulumi.set(__self__, "version", version)

    @property
    @pulumi.getter(name="datasetId")
    def dataset_id(self) -> pulumi.Input[str]:
        return pulumi.get(self, "dataset_id")

    @dataset_id.setter
    def dataset_id(self, value: pulumi.Input[str]):
        pulumi.set(self, "dataset_id", value)

    @property
    @pulumi.getter
    def location(self) -> pulumi.Input[str]:
        return pulumi.get(self, "location")

    @location.setter
    def location(self, value: pulumi.Input[str]):
        pulumi.set(self, "location", value)

    @property
    @pulumi.getter
    def project(self) -> pulumi.Input[str]:
        return pulumi.get(self, "project")

    @project.setter
    def project(self, value: pulumi.Input[str]):
        pulumi.set(self, "project", value)

    @property
    @pulumi.getter(name="defaultSearchHandlingStrict")
    def default_search_handling_strict(self) -> Optional[pulumi.Input[bool]]:
        """
        If true, overrides the default search behavior for this FHIR store to `handling=strict` which returns an error for unrecognized search parameters. If false, uses the FHIR specification default `handling=lenient` which ignores unrecognized search parameters. The handling can always be changed from the default on an individual API call by setting the HTTP header `Prefer: handling=strict` or `Prefer: handling=lenient`.
        """
        return pulumi.get(self, "default_search_handling_strict")

    @default_search_handling_strict.setter
    def default_search_handling_strict(self, value: Optional[pulumi.Input[bool]]):
        pulumi.set(self, "default_search_handling_strict", value)

    @property
    @pulumi.getter(name="disableReferentialIntegrity")
    def disable_referential_integrity(self) -> Optional[pulumi.Input[bool]]:
        """
        Immutable. Whether to disable referential integrity in this FHIR store. This field is immutable after FHIR store creation. The default value is false, meaning that the API enforces referential integrity and fails the requests that result in inconsistent state in the FHIR store. When this field is set to true, the API skips referential integrity checks. Consequently, operations that rely on references, such as GetPatientEverything, do not return all the results if broken references exist.
        """
        return pulumi.get(self, "disable_referential_integrity")

    @disable_referential_integrity.setter
    def disable_referential_integrity(self, value: Optional[pulumi.Input[bool]]):
        pulumi.set(self, "disable_referential_integrity", value)

    @property
    @pulumi.getter(name="disableResourceVersioning")
    def disable_resource_versioning(self) -> Optional[pulumi.Input[bool]]:
        """
        Immutable. Whether to disable resource versioning for this FHIR store. This field can not be changed after the creation of FHIR store. If set to false, which is the default behavior, all write operations cause historical versions to be recorded automatically. The historical versions can be fetched through the history APIs, but cannot be updated. If set to true, no historical versions are kept. The server sends errors for attempts to read the historical versions.
        """
        return pulumi.get(self, "disable_resource_versioning")

    @disable_resource_versioning.setter
    def disable_resource_versioning(self, value: Optional[pulumi.Input[bool]]):
        pulumi.set(self, "disable_resource_versioning", value)

    @property
    @pulumi.getter(name="enableUpdateCreate")
    def enable_update_create(self) -> Optional[pulumi.Input[bool]]:
        """
        Whether this FHIR store has the [updateCreate capability](https://www.hl7.org/fhir/capabilitystatement-definitions.html#CapabilityStatement.rest.resource.updateCreate). This determines if the client can use an Update operation to create a new resource with a client-specified ID. If false, all IDs are server-assigned through the Create operation and attempts to update a non-existent resource return errors. It is strongly advised not to include or encode any sensitive data such as patient identifiers in client-specified resource IDs. Those IDs are part of the FHIR resource path recorded in Cloud audit logs and Pub/Sub notifications. Those IDs can also be contained in reference fields within other resources.
        """
        return pulumi.get(self, "enable_update_create")

    @enable_update_create.setter
    def enable_update_create(self, value: Optional[pulumi.Input[bool]]):
        pulumi.set(self, "enable_update_create", value)

    @property
    @pulumi.getter(name="fhirStoreId")
    def fhir_store_id(self) -> Optional[pulumi.Input[str]]:
        return pulumi.get(self, "fhir_store_id")

    @fhir_store_id.setter
    def fhir_store_id(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "fhir_store_id", value)

    @property
    @pulumi.getter
    def labels(self) -> Optional[pulumi.Input[Mapping[str, str]]]:
        """
        User-supplied key-value pairs used to organize FHIR stores. Label keys must be between 1 and 63 characters long, have a UTF-8 encoding of maximum 128 bytes, and must conform to the following PCRE regular expression: \p{Ll}\p{Lo}{0,62} Label values are optional, must be between 1 and 63 characters long, have a UTF-8 encoding of maximum 128 bytes, and must conform to the following PCRE regular expression: [\p{Ll}\p{Lo}\p{N}_-]{0,63} No more than 64 labels can be associated with a given store.
        """
        return pulumi.get(self, "labels")

    @labels.setter
    def labels(self, value: Optional[pulumi.Input[Mapping[str, str]]]):
        pulumi.set(self, "labels", value)

    @property
    @pulumi.getter(name="notificationConfig")
    def notification_config(self) -> Optional[pulumi.Input['NotificationConfigArgs']]:
        """
        If non-empty, publish all resource modifications of this FHIR store to this destination. The Pub/Sub message attributes contain a map with a string describing the action that has triggered the notification. For example, "action":"CreateResource".
        """
        return pulumi.get(self, "notification_config")

    @notification_config.setter
    def notification_config(self, value: Optional[pulumi.Input['NotificationConfigArgs']]):
        pulumi.set(self, "notification_config", value)

    @property
    @pulumi.getter(name="streamConfigs")
    def stream_configs(self) -> Optional[pulumi.Input[Sequence[pulumi.Input['StreamConfigArgs']]]]:
        """
        A list of streaming configs that configure the destinations of streaming export for every resource mutation in this FHIR store. Each store is allowed to have up to 10 streaming configs. After a new config is added, the next resource mutation is streamed to the new location in addition to the existing ones. When a location is removed from the list, the server stops streaming to that location. Before adding a new config, you must add the required [`bigquery.dataEditor`](https://cloud.google.com/bigquery/docs/access-control#bigquery.dataEditor) role to your project's **Cloud Healthcare Service Agent** [service account](https://cloud.google.com/iam/docs/service-accounts). Some lag (typically on the order of dozens of seconds) is expected before the results show up in the streaming destination.
        """
        return pulumi.get(self, "stream_configs")

    @stream_configs.setter
    def stream_configs(self, value: Optional[pulumi.Input[Sequence[pulumi.Input['StreamConfigArgs']]]]):
        pulumi.set(self, "stream_configs", value)

    @property
    @pulumi.getter(name="validationConfig")
    def validation_config(self) -> Optional[pulumi.Input['ValidationConfigArgs']]:
        """
        Configuration for how to validate incoming FHIR resources against configured profiles.
        """
        return pulumi.get(self, "validation_config")

    @validation_config.setter
    def validation_config(self, value: Optional[pulumi.Input['ValidationConfigArgs']]):
        pulumi.set(self, "validation_config", value)

    @property
    @pulumi.getter
    def version(self) -> Optional[pulumi.Input['FhirStoreVersion']]:
        """
        Immutable. The FHIR specification version that this FHIR store supports natively. This field is immutable after store creation. Requests are rejected if they contain FHIR resources of a different version. Version is required for every FHIR store.
        """
        return pulumi.get(self, "version")

    @version.setter
    def version(self, value: Optional[pulumi.Input['FhirStoreVersion']]):
        pulumi.set(self, "version", value)


class FhirStore(pulumi.CustomResource):
    @overload
    def __init__(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 dataset_id: Optional[pulumi.Input[str]] = None,
                 default_search_handling_strict: Optional[pulumi.Input[bool]] = None,
                 disable_referential_integrity: Optional[pulumi.Input[bool]] = None,
                 disable_resource_versioning: Optional[pulumi.Input[bool]] = None,
                 enable_update_create: Optional[pulumi.Input[bool]] = None,
                 fhir_store_id: Optional[pulumi.Input[str]] = None,
                 labels: Optional[pulumi.Input[Mapping[str, str]]] = None,
                 location: Optional[pulumi.Input[str]] = None,
                 notification_config: Optional[pulumi.Input[pulumi.InputType['NotificationConfigArgs']]] = None,
                 project: Optional[pulumi.Input[str]] = None,
                 stream_configs: Optional[pulumi.Input[Sequence[pulumi.Input[pulumi.InputType['StreamConfigArgs']]]]] = None,
                 validation_config: Optional[pulumi.Input[pulumi.InputType['ValidationConfigArgs']]] = None,
                 version: Optional[pulumi.Input['FhirStoreVersion']] = None,
                 __props__=None):
        """
        Creates a new FHIR store within the parent dataset.

        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[bool] default_search_handling_strict: If true, overrides the default search behavior for this FHIR store to `handling=strict` which returns an error for unrecognized search parameters. If false, uses the FHIR specification default `handling=lenient` which ignores unrecognized search parameters. The handling can always be changed from the default on an individual API call by setting the HTTP header `Prefer: handling=strict` or `Prefer: handling=lenient`.
        :param pulumi.Input[bool] disable_referential_integrity: Immutable. Whether to disable referential integrity in this FHIR store. This field is immutable after FHIR store creation. The default value is false, meaning that the API enforces referential integrity and fails the requests that result in inconsistent state in the FHIR store. When this field is set to true, the API skips referential integrity checks. Consequently, operations that rely on references, such as GetPatientEverything, do not return all the results if broken references exist.
        :param pulumi.Input[bool] disable_resource_versioning: Immutable. Whether to disable resource versioning for this FHIR store. This field can not be changed after the creation of FHIR store. If set to false, which is the default behavior, all write operations cause historical versions to be recorded automatically. The historical versions can be fetched through the history APIs, but cannot be updated. If set to true, no historical versions are kept. The server sends errors for attempts to read the historical versions.
        :param pulumi.Input[bool] enable_update_create: Whether this FHIR store has the [updateCreate capability](https://www.hl7.org/fhir/capabilitystatement-definitions.html#CapabilityStatement.rest.resource.updateCreate). This determines if the client can use an Update operation to create a new resource with a client-specified ID. If false, all IDs are server-assigned through the Create operation and attempts to update a non-existent resource return errors. It is strongly advised not to include or encode any sensitive data such as patient identifiers in client-specified resource IDs. Those IDs are part of the FHIR resource path recorded in Cloud audit logs and Pub/Sub notifications. Those IDs can also be contained in reference fields within other resources.
        :param pulumi.Input[Mapping[str, str]] labels: User-supplied key-value pairs used to organize FHIR stores. Label keys must be between 1 and 63 characters long, have a UTF-8 encoding of maximum 128 bytes, and must conform to the following PCRE regular expression: \p{Ll}\p{Lo}{0,62} Label values are optional, must be between 1 and 63 characters long, have a UTF-8 encoding of maximum 128 bytes, and must conform to the following PCRE regular expression: [\p{Ll}\p{Lo}\p{N}_-]{0,63} No more than 64 labels can be associated with a given store.
        :param pulumi.Input[pulumi.InputType['NotificationConfigArgs']] notification_config: If non-empty, publish all resource modifications of this FHIR store to this destination. The Pub/Sub message attributes contain a map with a string describing the action that has triggered the notification. For example, "action":"CreateResource".
        :param pulumi.Input[Sequence[pulumi.Input[pulumi.InputType['StreamConfigArgs']]]] stream_configs: A list of streaming configs that configure the destinations of streaming export for every resource mutation in this FHIR store. Each store is allowed to have up to 10 streaming configs. After a new config is added, the next resource mutation is streamed to the new location in addition to the existing ones. When a location is removed from the list, the server stops streaming to that location. Before adding a new config, you must add the required [`bigquery.dataEditor`](https://cloud.google.com/bigquery/docs/access-control#bigquery.dataEditor) role to your project's **Cloud Healthcare Service Agent** [service account](https://cloud.google.com/iam/docs/service-accounts). Some lag (typically on the order of dozens of seconds) is expected before the results show up in the streaming destination.
        :param pulumi.Input[pulumi.InputType['ValidationConfigArgs']] validation_config: Configuration for how to validate incoming FHIR resources against configured profiles.
        :param pulumi.Input['FhirStoreVersion'] version: Immutable. The FHIR specification version that this FHIR store supports natively. This field is immutable after store creation. Requests are rejected if they contain FHIR resources of a different version. Version is required for every FHIR store.
        """
        ...
    @overload
    def __init__(__self__,
                 resource_name: str,
                 args: FhirStoreArgs,
                 opts: Optional[pulumi.ResourceOptions] = None):
        """
        Creates a new FHIR store within the parent dataset.

        :param str resource_name: The name of the resource.
        :param FhirStoreArgs args: The arguments to use to populate this resource's properties.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        ...
    def __init__(__self__, resource_name: str, *args, **kwargs):
        resource_args, opts = _utilities.get_resource_args_opts(FhirStoreArgs, pulumi.ResourceOptions, *args, **kwargs)
        if resource_args is not None:
            __self__._internal_init(resource_name, opts, **resource_args.__dict__)
        else:
            __self__._internal_init(resource_name, *args, **kwargs)

    def _internal_init(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 dataset_id: Optional[pulumi.Input[str]] = None,
                 default_search_handling_strict: Optional[pulumi.Input[bool]] = None,
                 disable_referential_integrity: Optional[pulumi.Input[bool]] = None,
                 disable_resource_versioning: Optional[pulumi.Input[bool]] = None,
                 enable_update_create: Optional[pulumi.Input[bool]] = None,
                 fhir_store_id: Optional[pulumi.Input[str]] = None,
                 labels: Optional[pulumi.Input[Mapping[str, str]]] = None,
                 location: Optional[pulumi.Input[str]] = None,
                 notification_config: Optional[pulumi.Input[pulumi.InputType['NotificationConfigArgs']]] = None,
                 project: Optional[pulumi.Input[str]] = None,
                 stream_configs: Optional[pulumi.Input[Sequence[pulumi.Input[pulumi.InputType['StreamConfigArgs']]]]] = None,
                 validation_config: Optional[pulumi.Input[pulumi.InputType['ValidationConfigArgs']]] = None,
                 version: Optional[pulumi.Input['FhirStoreVersion']] = None,
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
            __props__ = FhirStoreArgs.__new__(FhirStoreArgs)

            if dataset_id is None and not opts.urn:
                raise TypeError("Missing required property 'dataset_id'")
            __props__.__dict__["dataset_id"] = dataset_id
            __props__.__dict__["default_search_handling_strict"] = default_search_handling_strict
            __props__.__dict__["disable_referential_integrity"] = disable_referential_integrity
            __props__.__dict__["disable_resource_versioning"] = disable_resource_versioning
            __props__.__dict__["enable_update_create"] = enable_update_create
            __props__.__dict__["fhir_store_id"] = fhir_store_id
            __props__.__dict__["labels"] = labels
            if location is None and not opts.urn:
                raise TypeError("Missing required property 'location'")
            __props__.__dict__["location"] = location
            __props__.__dict__["notification_config"] = notification_config
            if project is None and not opts.urn:
                raise TypeError("Missing required property 'project'")
            __props__.__dict__["project"] = project
            __props__.__dict__["stream_configs"] = stream_configs
            __props__.__dict__["validation_config"] = validation_config
            __props__.__dict__["version"] = version
            __props__.__dict__["name"] = None
        super(FhirStore, __self__).__init__(
            'google-native:healthcare/v1beta1:FhirStore',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name: str,
            id: pulumi.Input[str],
            opts: Optional[pulumi.ResourceOptions] = None) -> 'FhirStore':
        """
        Get an existing FhirStore resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param pulumi.Input[str] id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = FhirStoreArgs.__new__(FhirStoreArgs)

        __props__.__dict__["default_search_handling_strict"] = None
        __props__.__dict__["disable_referential_integrity"] = None
        __props__.__dict__["disable_resource_versioning"] = None
        __props__.__dict__["enable_update_create"] = None
        __props__.__dict__["labels"] = None
        __props__.__dict__["name"] = None
        __props__.__dict__["notification_config"] = None
        __props__.__dict__["stream_configs"] = None
        __props__.__dict__["validation_config"] = None
        __props__.__dict__["version"] = None
        return FhirStore(resource_name, opts=opts, __props__=__props__)

    @property
    @pulumi.getter(name="defaultSearchHandlingStrict")
    def default_search_handling_strict(self) -> pulumi.Output[bool]:
        """
        If true, overrides the default search behavior for this FHIR store to `handling=strict` which returns an error for unrecognized search parameters. If false, uses the FHIR specification default `handling=lenient` which ignores unrecognized search parameters. The handling can always be changed from the default on an individual API call by setting the HTTP header `Prefer: handling=strict` or `Prefer: handling=lenient`.
        """
        return pulumi.get(self, "default_search_handling_strict")

    @property
    @pulumi.getter(name="disableReferentialIntegrity")
    def disable_referential_integrity(self) -> pulumi.Output[bool]:
        """
        Immutable. Whether to disable referential integrity in this FHIR store. This field is immutable after FHIR store creation. The default value is false, meaning that the API enforces referential integrity and fails the requests that result in inconsistent state in the FHIR store. When this field is set to true, the API skips referential integrity checks. Consequently, operations that rely on references, such as GetPatientEverything, do not return all the results if broken references exist.
        """
        return pulumi.get(self, "disable_referential_integrity")

    @property
    @pulumi.getter(name="disableResourceVersioning")
    def disable_resource_versioning(self) -> pulumi.Output[bool]:
        """
        Immutable. Whether to disable resource versioning for this FHIR store. This field can not be changed after the creation of FHIR store. If set to false, which is the default behavior, all write operations cause historical versions to be recorded automatically. The historical versions can be fetched through the history APIs, but cannot be updated. If set to true, no historical versions are kept. The server sends errors for attempts to read the historical versions.
        """
        return pulumi.get(self, "disable_resource_versioning")

    @property
    @pulumi.getter(name="enableUpdateCreate")
    def enable_update_create(self) -> pulumi.Output[bool]:
        """
        Whether this FHIR store has the [updateCreate capability](https://www.hl7.org/fhir/capabilitystatement-definitions.html#CapabilityStatement.rest.resource.updateCreate). This determines if the client can use an Update operation to create a new resource with a client-specified ID. If false, all IDs are server-assigned through the Create operation and attempts to update a non-existent resource return errors. It is strongly advised not to include or encode any sensitive data such as patient identifiers in client-specified resource IDs. Those IDs are part of the FHIR resource path recorded in Cloud audit logs and Pub/Sub notifications. Those IDs can also be contained in reference fields within other resources.
        """
        return pulumi.get(self, "enable_update_create")

    @property
    @pulumi.getter
    def labels(self) -> pulumi.Output[Mapping[str, str]]:
        """
        User-supplied key-value pairs used to organize FHIR stores. Label keys must be between 1 and 63 characters long, have a UTF-8 encoding of maximum 128 bytes, and must conform to the following PCRE regular expression: \p{Ll}\p{Lo}{0,62} Label values are optional, must be between 1 and 63 characters long, have a UTF-8 encoding of maximum 128 bytes, and must conform to the following PCRE regular expression: [\p{Ll}\p{Lo}\p{N}_-]{0,63} No more than 64 labels can be associated with a given store.
        """
        return pulumi.get(self, "labels")

    @property
    @pulumi.getter
    def name(self) -> pulumi.Output[str]:
        """
        Resource name of the FHIR store, of the form `projects/{project_id}/datasets/{dataset_id}/fhirStores/{fhir_store_id}`.
        """
        return pulumi.get(self, "name")

    @property
    @pulumi.getter(name="notificationConfig")
    def notification_config(self) -> pulumi.Output['outputs.NotificationConfigResponse']:
        """
        If non-empty, publish all resource modifications of this FHIR store to this destination. The Pub/Sub message attributes contain a map with a string describing the action that has triggered the notification. For example, "action":"CreateResource".
        """
        return pulumi.get(self, "notification_config")

    @property
    @pulumi.getter(name="streamConfigs")
    def stream_configs(self) -> pulumi.Output[Sequence['outputs.StreamConfigResponse']]:
        """
        A list of streaming configs that configure the destinations of streaming export for every resource mutation in this FHIR store. Each store is allowed to have up to 10 streaming configs. After a new config is added, the next resource mutation is streamed to the new location in addition to the existing ones. When a location is removed from the list, the server stops streaming to that location. Before adding a new config, you must add the required [`bigquery.dataEditor`](https://cloud.google.com/bigquery/docs/access-control#bigquery.dataEditor) role to your project's **Cloud Healthcare Service Agent** [service account](https://cloud.google.com/iam/docs/service-accounts). Some lag (typically on the order of dozens of seconds) is expected before the results show up in the streaming destination.
        """
        return pulumi.get(self, "stream_configs")

    @property
    @pulumi.getter(name="validationConfig")
    def validation_config(self) -> pulumi.Output['outputs.ValidationConfigResponse']:
        """
        Configuration for how to validate incoming FHIR resources against configured profiles.
        """
        return pulumi.get(self, "validation_config")

    @property
    @pulumi.getter
    def version(self) -> pulumi.Output[str]:
        """
        Immutable. The FHIR specification version that this FHIR store supports natively. This field is immutable after store creation. Requests are rejected if they contain FHIR resources of a different version. Version is required for every FHIR store.
        """
        return pulumi.get(self, "version")

