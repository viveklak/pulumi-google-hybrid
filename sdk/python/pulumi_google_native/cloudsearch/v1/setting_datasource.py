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

__all__ = ['SettingDatasourceArgs', 'SettingDatasource']

@pulumi.input_type
class SettingDatasourceArgs:
    def __init__(__self__, *,
                 disable_modifications: Optional[pulumi.Input[bool]] = None,
                 disable_serving: Optional[pulumi.Input[bool]] = None,
                 display_name: Optional[pulumi.Input[str]] = None,
                 indexing_service_accounts: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]] = None,
                 items_visibility: Optional[pulumi.Input[Sequence[pulumi.Input['GSuitePrincipalArgs']]]] = None,
                 name: Optional[pulumi.Input[str]] = None,
                 operation_ids: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]] = None,
                 short_name: Optional[pulumi.Input[str]] = None):
        """
        The set of arguments for constructing a SettingDatasource resource.
        :param pulumi.Input[bool] disable_modifications: If true, sets the datasource to read-only mode. In read-only mode, the Indexing API rejects any requests to index or delete items in this source. Enabling read-only mode does not stop the processing of previously accepted data.
        :param pulumi.Input[bool] disable_serving: Disable serving any search or assist results.
        :param pulumi.Input[str] display_name: Required. Display name of the datasource The maximum length is 300 characters.
        :param pulumi.Input[Sequence[pulumi.Input[str]]] indexing_service_accounts: List of service accounts that have indexing access.
        :param pulumi.Input[Sequence[pulumi.Input['GSuitePrincipalArgs']]] items_visibility: This field restricts visibility to items at the datasource level. Items within the datasource are restricted to the union of users and groups included in this field. Note that, this does not ensure access to a specific item, as users need to have ACL permissions on the contained items. This ensures a high level access on the entire datasource, and that the individual items are not shared outside this visibility.
        :param pulumi.Input[str] name: Name of the datasource resource. Format: datasources/{source_id}. The name is ignored when creating a datasource.
        :param pulumi.Input[Sequence[pulumi.Input[str]]] operation_ids: IDs of the Long Running Operations (LROs) currently running for this schema.
        :param pulumi.Input[str] short_name: A short name or alias for the source. This value will be used to match the 'source' operator. For example, if the short name is *<value>* then queries like *source:<value>* will only return results for this source. The value must be unique across all datasources. The value must only contain alphanumeric characters (a-zA-Z0-9). The value cannot start with 'google' and cannot be one of the following: mail, gmail, docs, drive, groups, sites, calendar, hangouts, gplus, keep, people, teams. Its maximum length is 32 characters.
        """
        if disable_modifications is not None:
            pulumi.set(__self__, "disable_modifications", disable_modifications)
        if disable_serving is not None:
            pulumi.set(__self__, "disable_serving", disable_serving)
        if display_name is not None:
            pulumi.set(__self__, "display_name", display_name)
        if indexing_service_accounts is not None:
            pulumi.set(__self__, "indexing_service_accounts", indexing_service_accounts)
        if items_visibility is not None:
            pulumi.set(__self__, "items_visibility", items_visibility)
        if name is not None:
            pulumi.set(__self__, "name", name)
        if operation_ids is not None:
            pulumi.set(__self__, "operation_ids", operation_ids)
        if short_name is not None:
            pulumi.set(__self__, "short_name", short_name)

    @property
    @pulumi.getter(name="disableModifications")
    def disable_modifications(self) -> Optional[pulumi.Input[bool]]:
        """
        If true, sets the datasource to read-only mode. In read-only mode, the Indexing API rejects any requests to index or delete items in this source. Enabling read-only mode does not stop the processing of previously accepted data.
        """
        return pulumi.get(self, "disable_modifications")

    @disable_modifications.setter
    def disable_modifications(self, value: Optional[pulumi.Input[bool]]):
        pulumi.set(self, "disable_modifications", value)

    @property
    @pulumi.getter(name="disableServing")
    def disable_serving(self) -> Optional[pulumi.Input[bool]]:
        """
        Disable serving any search or assist results.
        """
        return pulumi.get(self, "disable_serving")

    @disable_serving.setter
    def disable_serving(self, value: Optional[pulumi.Input[bool]]):
        pulumi.set(self, "disable_serving", value)

    @property
    @pulumi.getter(name="displayName")
    def display_name(self) -> Optional[pulumi.Input[str]]:
        """
        Required. Display name of the datasource The maximum length is 300 characters.
        """
        return pulumi.get(self, "display_name")

    @display_name.setter
    def display_name(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "display_name", value)

    @property
    @pulumi.getter(name="indexingServiceAccounts")
    def indexing_service_accounts(self) -> Optional[pulumi.Input[Sequence[pulumi.Input[str]]]]:
        """
        List of service accounts that have indexing access.
        """
        return pulumi.get(self, "indexing_service_accounts")

    @indexing_service_accounts.setter
    def indexing_service_accounts(self, value: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]]):
        pulumi.set(self, "indexing_service_accounts", value)

    @property
    @pulumi.getter(name="itemsVisibility")
    def items_visibility(self) -> Optional[pulumi.Input[Sequence[pulumi.Input['GSuitePrincipalArgs']]]]:
        """
        This field restricts visibility to items at the datasource level. Items within the datasource are restricted to the union of users and groups included in this field. Note that, this does not ensure access to a specific item, as users need to have ACL permissions on the contained items. This ensures a high level access on the entire datasource, and that the individual items are not shared outside this visibility.
        """
        return pulumi.get(self, "items_visibility")

    @items_visibility.setter
    def items_visibility(self, value: Optional[pulumi.Input[Sequence[pulumi.Input['GSuitePrincipalArgs']]]]):
        pulumi.set(self, "items_visibility", value)

    @property
    @pulumi.getter
    def name(self) -> Optional[pulumi.Input[str]]:
        """
        Name of the datasource resource. Format: datasources/{source_id}. The name is ignored when creating a datasource.
        """
        return pulumi.get(self, "name")

    @name.setter
    def name(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "name", value)

    @property
    @pulumi.getter(name="operationIds")
    def operation_ids(self) -> Optional[pulumi.Input[Sequence[pulumi.Input[str]]]]:
        """
        IDs of the Long Running Operations (LROs) currently running for this schema.
        """
        return pulumi.get(self, "operation_ids")

    @operation_ids.setter
    def operation_ids(self, value: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]]):
        pulumi.set(self, "operation_ids", value)

    @property
    @pulumi.getter(name="shortName")
    def short_name(self) -> Optional[pulumi.Input[str]]:
        """
        A short name or alias for the source. This value will be used to match the 'source' operator. For example, if the short name is *<value>* then queries like *source:<value>* will only return results for this source. The value must be unique across all datasources. The value must only contain alphanumeric characters (a-zA-Z0-9). The value cannot start with 'google' and cannot be one of the following: mail, gmail, docs, drive, groups, sites, calendar, hangouts, gplus, keep, people, teams. Its maximum length is 32 characters.
        """
        return pulumi.get(self, "short_name")

    @short_name.setter
    def short_name(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "short_name", value)


class SettingDatasource(pulumi.CustomResource):
    @overload
    def __init__(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 disable_modifications: Optional[pulumi.Input[bool]] = None,
                 disable_serving: Optional[pulumi.Input[bool]] = None,
                 display_name: Optional[pulumi.Input[str]] = None,
                 indexing_service_accounts: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]] = None,
                 items_visibility: Optional[pulumi.Input[Sequence[pulumi.Input[pulumi.InputType['GSuitePrincipalArgs']]]]] = None,
                 name: Optional[pulumi.Input[str]] = None,
                 operation_ids: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]] = None,
                 short_name: Optional[pulumi.Input[str]] = None,
                 __props__=None):
        """
        Creates a datasource. **Note:** This API requires an admin account to execute.

        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[bool] disable_modifications: If true, sets the datasource to read-only mode. In read-only mode, the Indexing API rejects any requests to index or delete items in this source. Enabling read-only mode does not stop the processing of previously accepted data.
        :param pulumi.Input[bool] disable_serving: Disable serving any search or assist results.
        :param pulumi.Input[str] display_name: Required. Display name of the datasource The maximum length is 300 characters.
        :param pulumi.Input[Sequence[pulumi.Input[str]]] indexing_service_accounts: List of service accounts that have indexing access.
        :param pulumi.Input[Sequence[pulumi.Input[pulumi.InputType['GSuitePrincipalArgs']]]] items_visibility: This field restricts visibility to items at the datasource level. Items within the datasource are restricted to the union of users and groups included in this field. Note that, this does not ensure access to a specific item, as users need to have ACL permissions on the contained items. This ensures a high level access on the entire datasource, and that the individual items are not shared outside this visibility.
        :param pulumi.Input[str] name: Name of the datasource resource. Format: datasources/{source_id}. The name is ignored when creating a datasource.
        :param pulumi.Input[Sequence[pulumi.Input[str]]] operation_ids: IDs of the Long Running Operations (LROs) currently running for this schema.
        :param pulumi.Input[str] short_name: A short name or alias for the source. This value will be used to match the 'source' operator. For example, if the short name is *<value>* then queries like *source:<value>* will only return results for this source. The value must be unique across all datasources. The value must only contain alphanumeric characters (a-zA-Z0-9). The value cannot start with 'google' and cannot be one of the following: mail, gmail, docs, drive, groups, sites, calendar, hangouts, gplus, keep, people, teams. Its maximum length is 32 characters.
        """
        ...
    @overload
    def __init__(__self__,
                 resource_name: str,
                 args: Optional[SettingDatasourceArgs] = None,
                 opts: Optional[pulumi.ResourceOptions] = None):
        """
        Creates a datasource. **Note:** This API requires an admin account to execute.

        :param str resource_name: The name of the resource.
        :param SettingDatasourceArgs args: The arguments to use to populate this resource's properties.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        ...
    def __init__(__self__, resource_name: str, *args, **kwargs):
        resource_args, opts = _utilities.get_resource_args_opts(SettingDatasourceArgs, pulumi.ResourceOptions, *args, **kwargs)
        if resource_args is not None:
            __self__._internal_init(resource_name, opts, **resource_args.__dict__)
        else:
            __self__._internal_init(resource_name, *args, **kwargs)

    def _internal_init(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 disable_modifications: Optional[pulumi.Input[bool]] = None,
                 disable_serving: Optional[pulumi.Input[bool]] = None,
                 display_name: Optional[pulumi.Input[str]] = None,
                 indexing_service_accounts: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]] = None,
                 items_visibility: Optional[pulumi.Input[Sequence[pulumi.Input[pulumi.InputType['GSuitePrincipalArgs']]]]] = None,
                 name: Optional[pulumi.Input[str]] = None,
                 operation_ids: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]] = None,
                 short_name: Optional[pulumi.Input[str]] = None,
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
            __props__ = SettingDatasourceArgs.__new__(SettingDatasourceArgs)

            __props__.__dict__["disable_modifications"] = disable_modifications
            __props__.__dict__["disable_serving"] = disable_serving
            __props__.__dict__["display_name"] = display_name
            __props__.__dict__["indexing_service_accounts"] = indexing_service_accounts
            __props__.__dict__["items_visibility"] = items_visibility
            __props__.__dict__["name"] = name
            __props__.__dict__["operation_ids"] = operation_ids
            __props__.__dict__["short_name"] = short_name
        super(SettingDatasource, __self__).__init__(
            'google-native:cloudsearch/v1:SettingDatasource',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name: str,
            id: pulumi.Input[str],
            opts: Optional[pulumi.ResourceOptions] = None) -> 'SettingDatasource':
        """
        Get an existing SettingDatasource resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param pulumi.Input[str] id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = SettingDatasourceArgs.__new__(SettingDatasourceArgs)

        __props__.__dict__["disable_modifications"] = None
        __props__.__dict__["disable_serving"] = None
        __props__.__dict__["display_name"] = None
        __props__.__dict__["indexing_service_accounts"] = None
        __props__.__dict__["items_visibility"] = None
        __props__.__dict__["name"] = None
        __props__.__dict__["operation_ids"] = None
        __props__.__dict__["short_name"] = None
        return SettingDatasource(resource_name, opts=opts, __props__=__props__)

    @property
    @pulumi.getter(name="disableModifications")
    def disable_modifications(self) -> pulumi.Output[bool]:
        """
        If true, sets the datasource to read-only mode. In read-only mode, the Indexing API rejects any requests to index or delete items in this source. Enabling read-only mode does not stop the processing of previously accepted data.
        """
        return pulumi.get(self, "disable_modifications")

    @property
    @pulumi.getter(name="disableServing")
    def disable_serving(self) -> pulumi.Output[bool]:
        """
        Disable serving any search or assist results.
        """
        return pulumi.get(self, "disable_serving")

    @property
    @pulumi.getter(name="displayName")
    def display_name(self) -> pulumi.Output[str]:
        """
        Required. Display name of the datasource The maximum length is 300 characters.
        """
        return pulumi.get(self, "display_name")

    @property
    @pulumi.getter(name="indexingServiceAccounts")
    def indexing_service_accounts(self) -> pulumi.Output[Sequence[str]]:
        """
        List of service accounts that have indexing access.
        """
        return pulumi.get(self, "indexing_service_accounts")

    @property
    @pulumi.getter(name="itemsVisibility")
    def items_visibility(self) -> pulumi.Output[Sequence['outputs.GSuitePrincipalResponse']]:
        """
        This field restricts visibility to items at the datasource level. Items within the datasource are restricted to the union of users and groups included in this field. Note that, this does not ensure access to a specific item, as users need to have ACL permissions on the contained items. This ensures a high level access on the entire datasource, and that the individual items are not shared outside this visibility.
        """
        return pulumi.get(self, "items_visibility")

    @property
    @pulumi.getter
    def name(self) -> pulumi.Output[str]:
        """
        Name of the datasource resource. Format: datasources/{source_id}. The name is ignored when creating a datasource.
        """
        return pulumi.get(self, "name")

    @property
    @pulumi.getter(name="operationIds")
    def operation_ids(self) -> pulumi.Output[Sequence[str]]:
        """
        IDs of the Long Running Operations (LROs) currently running for this schema.
        """
        return pulumi.get(self, "operation_ids")

    @property
    @pulumi.getter(name="shortName")
    def short_name(self) -> pulumi.Output[str]:
        """
        A short name or alias for the source. This value will be used to match the 'source' operator. For example, if the short name is *<value>* then queries like *source:<value>* will only return results for this source. The value must be unique across all datasources. The value must only contain alphanumeric characters (a-zA-Z0-9). The value cannot start with 'google' and cannot be one of the following: mail, gmail, docs, drive, groups, sites, calendar, hangouts, gplus, keep, people, teams. Its maximum length is 32 characters.
        """
        return pulumi.get(self, "short_name")

