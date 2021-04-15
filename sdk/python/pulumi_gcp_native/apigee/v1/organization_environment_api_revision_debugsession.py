# coding=utf-8
# *** WARNING: this file was generated by the Pulumi SDK Generator. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import warnings
import pulumi
import pulumi.runtime
from typing import Any, Mapping, Optional, Sequence, Union, overload
from ... import _utilities

__all__ = ['OrganizationEnvironmentApiRevisionDebugsessionArgs', 'OrganizationEnvironmentApiRevisionDebugsession']

@pulumi.input_type
class OrganizationEnvironmentApiRevisionDebugsessionArgs:
    def __init__(__self__, *,
                 apis_id: pulumi.Input[str],
                 debugsessions_id: pulumi.Input[str],
                 environments_id: pulumi.Input[str],
                 organizations_id: pulumi.Input[str],
                 revisions_id: pulumi.Input[str],
                 count: Optional[pulumi.Input[int]] = None,
                 filter: Optional[pulumi.Input[str]] = None,
                 name: Optional[pulumi.Input[str]] = None,
                 timeout: Optional[pulumi.Input[str]] = None,
                 tracesize: Optional[pulumi.Input[int]] = None,
                 validity: Optional[pulumi.Input[int]] = None):
        """
        The set of arguments for constructing a OrganizationEnvironmentApiRevisionDebugsession resource.
        :param pulumi.Input[int] count: Optional. The number of request to be traced. Min = 1, Max = 15, Default = 10.
        :param pulumi.Input[str] filter: Optional. A conditional statement which is evaluated against the request message to determine if it should be traced. Syntax matches that of on API Proxy bundle flow Condition.
        :param pulumi.Input[str] name: A unique ID for this DebugSession.
        :param pulumi.Input[str] timeout: Optional. The time in seconds after which this DebugSession should end. This value will override the value in query param, if both are provided.
        :param pulumi.Input[int] tracesize: Optional. The maximum number of bytes captured from the response payload. Min = 0, Max = 5120, Default = 5120.
        :param pulumi.Input[int] validity: Optional. The length of time, in seconds, that this debug session is valid, starting from when it's received in the control plane. Min = 1, Max = 15, Default = 10.
        """
        pulumi.set(__self__, "apis_id", apis_id)
        pulumi.set(__self__, "debugsessions_id", debugsessions_id)
        pulumi.set(__self__, "environments_id", environments_id)
        pulumi.set(__self__, "organizations_id", organizations_id)
        pulumi.set(__self__, "revisions_id", revisions_id)
        if count is not None:
            pulumi.set(__self__, "count", count)
        if filter is not None:
            pulumi.set(__self__, "filter", filter)
        if name is not None:
            pulumi.set(__self__, "name", name)
        if timeout is not None:
            pulumi.set(__self__, "timeout", timeout)
        if tracesize is not None:
            pulumi.set(__self__, "tracesize", tracesize)
        if validity is not None:
            pulumi.set(__self__, "validity", validity)

    @property
    @pulumi.getter(name="apisId")
    def apis_id(self) -> pulumi.Input[str]:
        return pulumi.get(self, "apis_id")

    @apis_id.setter
    def apis_id(self, value: pulumi.Input[str]):
        pulumi.set(self, "apis_id", value)

    @property
    @pulumi.getter(name="debugsessionsId")
    def debugsessions_id(self) -> pulumi.Input[str]:
        return pulumi.get(self, "debugsessions_id")

    @debugsessions_id.setter
    def debugsessions_id(self, value: pulumi.Input[str]):
        pulumi.set(self, "debugsessions_id", value)

    @property
    @pulumi.getter(name="environmentsId")
    def environments_id(self) -> pulumi.Input[str]:
        return pulumi.get(self, "environments_id")

    @environments_id.setter
    def environments_id(self, value: pulumi.Input[str]):
        pulumi.set(self, "environments_id", value)

    @property
    @pulumi.getter(name="organizationsId")
    def organizations_id(self) -> pulumi.Input[str]:
        return pulumi.get(self, "organizations_id")

    @organizations_id.setter
    def organizations_id(self, value: pulumi.Input[str]):
        pulumi.set(self, "organizations_id", value)

    @property
    @pulumi.getter(name="revisionsId")
    def revisions_id(self) -> pulumi.Input[str]:
        return pulumi.get(self, "revisions_id")

    @revisions_id.setter
    def revisions_id(self, value: pulumi.Input[str]):
        pulumi.set(self, "revisions_id", value)

    @property
    @pulumi.getter
    def count(self) -> Optional[pulumi.Input[int]]:
        """
        Optional. The number of request to be traced. Min = 1, Max = 15, Default = 10.
        """
        return pulumi.get(self, "count")

    @count.setter
    def count(self, value: Optional[pulumi.Input[int]]):
        pulumi.set(self, "count", value)

    @property
    @pulumi.getter
    def filter(self) -> Optional[pulumi.Input[str]]:
        """
        Optional. A conditional statement which is evaluated against the request message to determine if it should be traced. Syntax matches that of on API Proxy bundle flow Condition.
        """
        return pulumi.get(self, "filter")

    @filter.setter
    def filter(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "filter", value)

    @property
    @pulumi.getter
    def name(self) -> Optional[pulumi.Input[str]]:
        """
        A unique ID for this DebugSession.
        """
        return pulumi.get(self, "name")

    @name.setter
    def name(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "name", value)

    @property
    @pulumi.getter
    def timeout(self) -> Optional[pulumi.Input[str]]:
        """
        Optional. The time in seconds after which this DebugSession should end. This value will override the value in query param, if both are provided.
        """
        return pulumi.get(self, "timeout")

    @timeout.setter
    def timeout(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "timeout", value)

    @property
    @pulumi.getter
    def tracesize(self) -> Optional[pulumi.Input[int]]:
        """
        Optional. The maximum number of bytes captured from the response payload. Min = 0, Max = 5120, Default = 5120.
        """
        return pulumi.get(self, "tracesize")

    @tracesize.setter
    def tracesize(self, value: Optional[pulumi.Input[int]]):
        pulumi.set(self, "tracesize", value)

    @property
    @pulumi.getter
    def validity(self) -> Optional[pulumi.Input[int]]:
        """
        Optional. The length of time, in seconds, that this debug session is valid, starting from when it's received in the control plane. Min = 1, Max = 15, Default = 10.
        """
        return pulumi.get(self, "validity")

    @validity.setter
    def validity(self, value: Optional[pulumi.Input[int]]):
        pulumi.set(self, "validity", value)


class OrganizationEnvironmentApiRevisionDebugsession(pulumi.CustomResource):
    @overload
    def __init__(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 apis_id: Optional[pulumi.Input[str]] = None,
                 count: Optional[pulumi.Input[int]] = None,
                 debugsessions_id: Optional[pulumi.Input[str]] = None,
                 environments_id: Optional[pulumi.Input[str]] = None,
                 filter: Optional[pulumi.Input[str]] = None,
                 name: Optional[pulumi.Input[str]] = None,
                 organizations_id: Optional[pulumi.Input[str]] = None,
                 revisions_id: Optional[pulumi.Input[str]] = None,
                 timeout: Optional[pulumi.Input[str]] = None,
                 tracesize: Optional[pulumi.Input[int]] = None,
                 validity: Optional[pulumi.Input[int]] = None,
                 __props__=None):
        """
        Creates a debug session for a deployed API Proxy revision.

        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[int] count: Optional. The number of request to be traced. Min = 1, Max = 15, Default = 10.
        :param pulumi.Input[str] filter: Optional. A conditional statement which is evaluated against the request message to determine if it should be traced. Syntax matches that of on API Proxy bundle flow Condition.
        :param pulumi.Input[str] name: A unique ID for this DebugSession.
        :param pulumi.Input[str] timeout: Optional. The time in seconds after which this DebugSession should end. This value will override the value in query param, if both are provided.
        :param pulumi.Input[int] tracesize: Optional. The maximum number of bytes captured from the response payload. Min = 0, Max = 5120, Default = 5120.
        :param pulumi.Input[int] validity: Optional. The length of time, in seconds, that this debug session is valid, starting from when it's received in the control plane. Min = 1, Max = 15, Default = 10.
        """
        ...
    @overload
    def __init__(__self__,
                 resource_name: str,
                 args: OrganizationEnvironmentApiRevisionDebugsessionArgs,
                 opts: Optional[pulumi.ResourceOptions] = None):
        """
        Creates a debug session for a deployed API Proxy revision.

        :param str resource_name: The name of the resource.
        :param OrganizationEnvironmentApiRevisionDebugsessionArgs args: The arguments to use to populate this resource's properties.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        ...
    def __init__(__self__, resource_name: str, *args, **kwargs):
        resource_args, opts = _utilities.get_resource_args_opts(OrganizationEnvironmentApiRevisionDebugsessionArgs, pulumi.ResourceOptions, *args, **kwargs)
        if resource_args is not None:
            __self__._internal_init(resource_name, opts, **resource_args.__dict__)
        else:
            __self__._internal_init(resource_name, *args, **kwargs)

    def _internal_init(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 apis_id: Optional[pulumi.Input[str]] = None,
                 count: Optional[pulumi.Input[int]] = None,
                 debugsessions_id: Optional[pulumi.Input[str]] = None,
                 environments_id: Optional[pulumi.Input[str]] = None,
                 filter: Optional[pulumi.Input[str]] = None,
                 name: Optional[pulumi.Input[str]] = None,
                 organizations_id: Optional[pulumi.Input[str]] = None,
                 revisions_id: Optional[pulumi.Input[str]] = None,
                 timeout: Optional[pulumi.Input[str]] = None,
                 tracesize: Optional[pulumi.Input[int]] = None,
                 validity: Optional[pulumi.Input[int]] = None,
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
            __props__ = OrganizationEnvironmentApiRevisionDebugsessionArgs.__new__(OrganizationEnvironmentApiRevisionDebugsessionArgs)

            if apis_id is None and not opts.urn:
                raise TypeError("Missing required property 'apis_id'")
            __props__.__dict__["apis_id"] = apis_id
            __props__.__dict__["count"] = count
            if debugsessions_id is None and not opts.urn:
                raise TypeError("Missing required property 'debugsessions_id'")
            __props__.__dict__["debugsessions_id"] = debugsessions_id
            if environments_id is None and not opts.urn:
                raise TypeError("Missing required property 'environments_id'")
            __props__.__dict__["environments_id"] = environments_id
            __props__.__dict__["filter"] = filter
            __props__.__dict__["name"] = name
            if organizations_id is None and not opts.urn:
                raise TypeError("Missing required property 'organizations_id'")
            __props__.__dict__["organizations_id"] = organizations_id
            if revisions_id is None and not opts.urn:
                raise TypeError("Missing required property 'revisions_id'")
            __props__.__dict__["revisions_id"] = revisions_id
            __props__.__dict__["timeout"] = timeout
            __props__.__dict__["tracesize"] = tracesize
            __props__.__dict__["validity"] = validity
        super(OrganizationEnvironmentApiRevisionDebugsession, __self__).__init__(
            'gcp-native:apigee/v1:OrganizationEnvironmentApiRevisionDebugsession',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name: str,
            id: pulumi.Input[str],
            opts: Optional[pulumi.ResourceOptions] = None) -> 'OrganizationEnvironmentApiRevisionDebugsession':
        """
        Get an existing OrganizationEnvironmentApiRevisionDebugsession resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param pulumi.Input[str] id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = OrganizationEnvironmentApiRevisionDebugsessionArgs.__new__(OrganizationEnvironmentApiRevisionDebugsessionArgs)

        __props__.__dict__["count"] = None
        __props__.__dict__["filter"] = None
        __props__.__dict__["name"] = None
        __props__.__dict__["timeout"] = None
        __props__.__dict__["tracesize"] = None
        __props__.__dict__["validity"] = None
        return OrganizationEnvironmentApiRevisionDebugsession(resource_name, opts=opts, __props__=__props__)

    @property
    @pulumi.getter
    def count(self) -> pulumi.Output[int]:
        """
        Optional. The number of request to be traced. Min = 1, Max = 15, Default = 10.
        """
        return pulumi.get(self, "count")

    @property
    @pulumi.getter
    def filter(self) -> pulumi.Output[str]:
        """
        Optional. A conditional statement which is evaluated against the request message to determine if it should be traced. Syntax matches that of on API Proxy bundle flow Condition.
        """
        return pulumi.get(self, "filter")

    @property
    @pulumi.getter
    def name(self) -> pulumi.Output[str]:
        """
        A unique ID for this DebugSession.
        """
        return pulumi.get(self, "name")

    @property
    @pulumi.getter
    def timeout(self) -> pulumi.Output[str]:
        """
        Optional. The time in seconds after which this DebugSession should end. This value will override the value in query param, if both are provided.
        """
        return pulumi.get(self, "timeout")

    @property
    @pulumi.getter
    def tracesize(self) -> pulumi.Output[int]:
        """
        Optional. The maximum number of bytes captured from the response payload. Min = 0, Max = 5120, Default = 5120.
        """
        return pulumi.get(self, "tracesize")

    @property
    @pulumi.getter
    def validity(self) -> pulumi.Output[int]:
        """
        Optional. The length of time, in seconds, that this debug session is valid, starting from when it's received in the control plane. Min = 1, Max = 15, Default = 10.
        """
        return pulumi.get(self, "validity")

