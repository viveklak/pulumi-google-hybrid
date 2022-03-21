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

__all__ = ['ConnectionProfileArgs', 'ConnectionProfile']

@pulumi.input_type
class ConnectionProfileArgs:
    def __init__(__self__, *,
                 connection_profile_id: pulumi.Input[str],
                 cloudsql: Optional[pulumi.Input['CloudSqlConnectionProfileArgs']] = None,
                 display_name: Optional[pulumi.Input[str]] = None,
                 labels: Optional[pulumi.Input[Mapping[str, pulumi.Input[str]]]] = None,
                 location: Optional[pulumi.Input[str]] = None,
                 mysql: Optional[pulumi.Input['MySqlConnectionProfileArgs']] = None,
                 name: Optional[pulumi.Input[str]] = None,
                 postgresql: Optional[pulumi.Input['PostgreSqlConnectionProfileArgs']] = None,
                 project: Optional[pulumi.Input[str]] = None,
                 provider: Optional[pulumi.Input['ConnectionProfileProvider']] = None,
                 request_id: Optional[pulumi.Input[str]] = None,
                 state: Optional[pulumi.Input['ConnectionProfileState']] = None):
        """
        The set of arguments for constructing a ConnectionProfile resource.
        :param pulumi.Input[str] connection_profile_id: Required. The connection profile identifier.
        :param pulumi.Input['CloudSqlConnectionProfileArgs'] cloudsql: A CloudSQL database connection profile.
        :param pulumi.Input[str] display_name: The connection profile display name.
        :param pulumi.Input[Mapping[str, pulumi.Input[str]]] labels: The resource labels for connection profile to use to annotate any related underlying resources such as Compute Engine VMs. An object containing a list of "key": "value" pairs. Example: `{ "name": "wrench", "mass": "1.3kg", "count": "3" }`.
        :param pulumi.Input['MySqlConnectionProfileArgs'] mysql: A MySQL database connection profile.
        :param pulumi.Input[str] name: The name of this connection profile resource in the form of projects/{project}/locations/{location}/connectionProfiles/{connectionProfile}.
        :param pulumi.Input['PostgreSqlConnectionProfileArgs'] postgresql: A PostgreSQL database connection profile.
        :param pulumi.Input['ConnectionProfileProvider'] provider: The database provider.
        :param pulumi.Input[str] request_id: A unique id used to identify the request. If the server receives two requests with the same id, then the second request will be ignored. It is recommended to always set this value to a UUID. The id must contain only letters (a-z, A-Z), numbers (0-9), underscores (_), and hyphens (-). The maximum length is 40 characters.
        :param pulumi.Input['ConnectionProfileState'] state: The current connection profile state (e.g. DRAFT, READY, or FAILED).
        """
        pulumi.set(__self__, "connection_profile_id", connection_profile_id)
        if cloudsql is not None:
            pulumi.set(__self__, "cloudsql", cloudsql)
        if display_name is not None:
            pulumi.set(__self__, "display_name", display_name)
        if labels is not None:
            pulumi.set(__self__, "labels", labels)
        if location is not None:
            pulumi.set(__self__, "location", location)
        if mysql is not None:
            pulumi.set(__self__, "mysql", mysql)
        if name is not None:
            pulumi.set(__self__, "name", name)
        if postgresql is not None:
            pulumi.set(__self__, "postgresql", postgresql)
        if project is not None:
            pulumi.set(__self__, "project", project)
        if provider is not None:
            pulumi.set(__self__, "provider", provider)
        if request_id is not None:
            pulumi.set(__self__, "request_id", request_id)
        if state is not None:
            pulumi.set(__self__, "state", state)

    @property
    @pulumi.getter(name="connectionProfileId")
    def connection_profile_id(self) -> pulumi.Input[str]:
        """
        Required. The connection profile identifier.
        """
        return pulumi.get(self, "connection_profile_id")

    @connection_profile_id.setter
    def connection_profile_id(self, value: pulumi.Input[str]):
        pulumi.set(self, "connection_profile_id", value)

    @property
    @pulumi.getter
    def cloudsql(self) -> Optional[pulumi.Input['CloudSqlConnectionProfileArgs']]:
        """
        A CloudSQL database connection profile.
        """
        return pulumi.get(self, "cloudsql")

    @cloudsql.setter
    def cloudsql(self, value: Optional[pulumi.Input['CloudSqlConnectionProfileArgs']]):
        pulumi.set(self, "cloudsql", value)

    @property
    @pulumi.getter(name="displayName")
    def display_name(self) -> Optional[pulumi.Input[str]]:
        """
        The connection profile display name.
        """
        return pulumi.get(self, "display_name")

    @display_name.setter
    def display_name(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "display_name", value)

    @property
    @pulumi.getter
    def labels(self) -> Optional[pulumi.Input[Mapping[str, pulumi.Input[str]]]]:
        """
        The resource labels for connection profile to use to annotate any related underlying resources such as Compute Engine VMs. An object containing a list of "key": "value" pairs. Example: `{ "name": "wrench", "mass": "1.3kg", "count": "3" }`.
        """
        return pulumi.get(self, "labels")

    @labels.setter
    def labels(self, value: Optional[pulumi.Input[Mapping[str, pulumi.Input[str]]]]):
        pulumi.set(self, "labels", value)

    @property
    @pulumi.getter
    def location(self) -> Optional[pulumi.Input[str]]:
        return pulumi.get(self, "location")

    @location.setter
    def location(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "location", value)

    @property
    @pulumi.getter
    def mysql(self) -> Optional[pulumi.Input['MySqlConnectionProfileArgs']]:
        """
        A MySQL database connection profile.
        """
        return pulumi.get(self, "mysql")

    @mysql.setter
    def mysql(self, value: Optional[pulumi.Input['MySqlConnectionProfileArgs']]):
        pulumi.set(self, "mysql", value)

    @property
    @pulumi.getter
    def name(self) -> Optional[pulumi.Input[str]]:
        """
        The name of this connection profile resource in the form of projects/{project}/locations/{location}/connectionProfiles/{connectionProfile}.
        """
        return pulumi.get(self, "name")

    @name.setter
    def name(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "name", value)

    @property
    @pulumi.getter
    def postgresql(self) -> Optional[pulumi.Input['PostgreSqlConnectionProfileArgs']]:
        """
        A PostgreSQL database connection profile.
        """
        return pulumi.get(self, "postgresql")

    @postgresql.setter
    def postgresql(self, value: Optional[pulumi.Input['PostgreSqlConnectionProfileArgs']]):
        pulumi.set(self, "postgresql", value)

    @property
    @pulumi.getter
    def project(self) -> Optional[pulumi.Input[str]]:
        return pulumi.get(self, "project")

    @project.setter
    def project(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "project", value)

    @property
    @pulumi.getter
    def provider(self) -> Optional[pulumi.Input['ConnectionProfileProvider']]:
        """
        The database provider.
        """
        return pulumi.get(self, "provider")

    @provider.setter
    def provider(self, value: Optional[pulumi.Input['ConnectionProfileProvider']]):
        pulumi.set(self, "provider", value)

    @property
    @pulumi.getter(name="requestId")
    def request_id(self) -> Optional[pulumi.Input[str]]:
        """
        A unique id used to identify the request. If the server receives two requests with the same id, then the second request will be ignored. It is recommended to always set this value to a UUID. The id must contain only letters (a-z, A-Z), numbers (0-9), underscores (_), and hyphens (-). The maximum length is 40 characters.
        """
        return pulumi.get(self, "request_id")

    @request_id.setter
    def request_id(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "request_id", value)

    @property
    @pulumi.getter
    def state(self) -> Optional[pulumi.Input['ConnectionProfileState']]:
        """
        The current connection profile state (e.g. DRAFT, READY, or FAILED).
        """
        return pulumi.get(self, "state")

    @state.setter
    def state(self, value: Optional[pulumi.Input['ConnectionProfileState']]):
        pulumi.set(self, "state", value)


class ConnectionProfile(pulumi.CustomResource):
    @overload
    def __init__(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 cloudsql: Optional[pulumi.Input[pulumi.InputType['CloudSqlConnectionProfileArgs']]] = None,
                 connection_profile_id: Optional[pulumi.Input[str]] = None,
                 display_name: Optional[pulumi.Input[str]] = None,
                 labels: Optional[pulumi.Input[Mapping[str, pulumi.Input[str]]]] = None,
                 location: Optional[pulumi.Input[str]] = None,
                 mysql: Optional[pulumi.Input[pulumi.InputType['MySqlConnectionProfileArgs']]] = None,
                 name: Optional[pulumi.Input[str]] = None,
                 postgresql: Optional[pulumi.Input[pulumi.InputType['PostgreSqlConnectionProfileArgs']]] = None,
                 project: Optional[pulumi.Input[str]] = None,
                 provider: Optional[pulumi.Input['ConnectionProfileProvider']] = None,
                 request_id: Optional[pulumi.Input[str]] = None,
                 state: Optional[pulumi.Input['ConnectionProfileState']] = None,
                 __props__=None):
        """
        Creates a new connection profile in a given project and location.

        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[pulumi.InputType['CloudSqlConnectionProfileArgs']] cloudsql: A CloudSQL database connection profile.
        :param pulumi.Input[str] connection_profile_id: Required. The connection profile identifier.
        :param pulumi.Input[str] display_name: The connection profile display name.
        :param pulumi.Input[Mapping[str, pulumi.Input[str]]] labels: The resource labels for connection profile to use to annotate any related underlying resources such as Compute Engine VMs. An object containing a list of "key": "value" pairs. Example: `{ "name": "wrench", "mass": "1.3kg", "count": "3" }`.
        :param pulumi.Input[pulumi.InputType['MySqlConnectionProfileArgs']] mysql: A MySQL database connection profile.
        :param pulumi.Input[str] name: The name of this connection profile resource in the form of projects/{project}/locations/{location}/connectionProfiles/{connectionProfile}.
        :param pulumi.Input[pulumi.InputType['PostgreSqlConnectionProfileArgs']] postgresql: A PostgreSQL database connection profile.
        :param pulumi.Input['ConnectionProfileProvider'] provider: The database provider.
        :param pulumi.Input[str] request_id: A unique id used to identify the request. If the server receives two requests with the same id, then the second request will be ignored. It is recommended to always set this value to a UUID. The id must contain only letters (a-z, A-Z), numbers (0-9), underscores (_), and hyphens (-). The maximum length is 40 characters.
        :param pulumi.Input['ConnectionProfileState'] state: The current connection profile state (e.g. DRAFT, READY, or FAILED).
        """
        ...
    @overload
    def __init__(__self__,
                 resource_name: str,
                 args: ConnectionProfileArgs,
                 opts: Optional[pulumi.ResourceOptions] = None):
        """
        Creates a new connection profile in a given project and location.

        :param str resource_name: The name of the resource.
        :param ConnectionProfileArgs args: The arguments to use to populate this resource's properties.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        ...
    def __init__(__self__, resource_name: str, *args, **kwargs):
        resource_args, opts = _utilities.get_resource_args_opts(ConnectionProfileArgs, pulumi.ResourceOptions, *args, **kwargs)
        if resource_args is not None:
            __self__._internal_init(resource_name, opts, **resource_args.__dict__)
        else:
            __self__._internal_init(resource_name, *args, **kwargs)

    def _internal_init(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 cloudsql: Optional[pulumi.Input[pulumi.InputType['CloudSqlConnectionProfileArgs']]] = None,
                 connection_profile_id: Optional[pulumi.Input[str]] = None,
                 display_name: Optional[pulumi.Input[str]] = None,
                 labels: Optional[pulumi.Input[Mapping[str, pulumi.Input[str]]]] = None,
                 location: Optional[pulumi.Input[str]] = None,
                 mysql: Optional[pulumi.Input[pulumi.InputType['MySqlConnectionProfileArgs']]] = None,
                 name: Optional[pulumi.Input[str]] = None,
                 postgresql: Optional[pulumi.Input[pulumi.InputType['PostgreSqlConnectionProfileArgs']]] = None,
                 project: Optional[pulumi.Input[str]] = None,
                 provider: Optional[pulumi.Input['ConnectionProfileProvider']] = None,
                 request_id: Optional[pulumi.Input[str]] = None,
                 state: Optional[pulumi.Input['ConnectionProfileState']] = None,
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
            __props__ = ConnectionProfileArgs.__new__(ConnectionProfileArgs)

            __props__.__dict__["cloudsql"] = cloudsql
            if connection_profile_id is None and not opts.urn:
                raise TypeError("Missing required property 'connection_profile_id'")
            __props__.__dict__["connection_profile_id"] = connection_profile_id
            __props__.__dict__["display_name"] = display_name
            __props__.__dict__["labels"] = labels
            __props__.__dict__["location"] = location
            __props__.__dict__["mysql"] = mysql
            __props__.__dict__["name"] = name
            __props__.__dict__["postgresql"] = postgresql
            __props__.__dict__["project"] = project
            __props__.__dict__["provider"] = provider
            __props__.__dict__["request_id"] = request_id
            __props__.__dict__["state"] = state
            __props__.__dict__["create_time"] = None
            __props__.__dict__["error"] = None
            __props__.__dict__["update_time"] = None
        super(ConnectionProfile, __self__).__init__(
            'google-native:datamigration/v1:ConnectionProfile',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name: str,
            id: pulumi.Input[str],
            opts: Optional[pulumi.ResourceOptions] = None) -> 'ConnectionProfile':
        """
        Get an existing ConnectionProfile resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param pulumi.Input[str] id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = ConnectionProfileArgs.__new__(ConnectionProfileArgs)

        __props__.__dict__["cloudsql"] = None
        __props__.__dict__["create_time"] = None
        __props__.__dict__["display_name"] = None
        __props__.__dict__["error"] = None
        __props__.__dict__["labels"] = None
        __props__.__dict__["mysql"] = None
        __props__.__dict__["name"] = None
        __props__.__dict__["postgresql"] = None
        __props__.__dict__["provider"] = None
        __props__.__dict__["state"] = None
        __props__.__dict__["update_time"] = None
        return ConnectionProfile(resource_name, opts=opts, __props__=__props__)

    @property
    @pulumi.getter
    def cloudsql(self) -> pulumi.Output['outputs.CloudSqlConnectionProfileResponse']:
        """
        A CloudSQL database connection profile.
        """
        return pulumi.get(self, "cloudsql")

    @property
    @pulumi.getter(name="createTime")
    def create_time(self) -> pulumi.Output[str]:
        """
        The timestamp when the resource was created. A timestamp in RFC3339 UTC "Zulu" format, accurate to nanoseconds. Example: "2014-10-02T15:01:23.045123456Z".
        """
        return pulumi.get(self, "create_time")

    @property
    @pulumi.getter(name="displayName")
    def display_name(self) -> pulumi.Output[str]:
        """
        The connection profile display name.
        """
        return pulumi.get(self, "display_name")

    @property
    @pulumi.getter
    def error(self) -> pulumi.Output['outputs.StatusResponse']:
        """
        The error details in case of state FAILED.
        """
        return pulumi.get(self, "error")

    @property
    @pulumi.getter
    def labels(self) -> pulumi.Output[Mapping[str, str]]:
        """
        The resource labels for connection profile to use to annotate any related underlying resources such as Compute Engine VMs. An object containing a list of "key": "value" pairs. Example: `{ "name": "wrench", "mass": "1.3kg", "count": "3" }`.
        """
        return pulumi.get(self, "labels")

    @property
    @pulumi.getter
    def mysql(self) -> pulumi.Output['outputs.MySqlConnectionProfileResponse']:
        """
        A MySQL database connection profile.
        """
        return pulumi.get(self, "mysql")

    @property
    @pulumi.getter
    def name(self) -> pulumi.Output[str]:
        """
        The name of this connection profile resource in the form of projects/{project}/locations/{location}/connectionProfiles/{connectionProfile}.
        """
        return pulumi.get(self, "name")

    @property
    @pulumi.getter
    def postgresql(self) -> pulumi.Output['outputs.PostgreSqlConnectionProfileResponse']:
        """
        A PostgreSQL database connection profile.
        """
        return pulumi.get(self, "postgresql")

    @property
    @pulumi.getter
    def provider(self) -> pulumi.Output[str]:
        """
        The database provider.
        """
        return pulumi.get(self, "provider")

    @property
    @pulumi.getter
    def state(self) -> pulumi.Output[str]:
        """
        The current connection profile state (e.g. DRAFT, READY, or FAILED).
        """
        return pulumi.get(self, "state")

    @property
    @pulumi.getter(name="updateTime")
    def update_time(self) -> pulumi.Output[str]:
        """
        The timestamp when the resource was last updated. A timestamp in RFC3339 UTC "Zulu" format, accurate to nanoseconds. Example: "2014-10-02T15:01:23.045123456Z".
        """
        return pulumi.get(self, "update_time")

