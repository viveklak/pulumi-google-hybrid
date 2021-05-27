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

__all__ = ['SslPolicyArgs', 'SslPolicy']

@pulumi.input_type
class SslPolicyArgs:
    def __init__(__self__, *,
                 project: pulumi.Input[str],
                 creation_timestamp: Optional[pulumi.Input[str]] = None,
                 custom_features: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]] = None,
                 description: Optional[pulumi.Input[str]] = None,
                 enabled_features: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]] = None,
                 fingerprint: Optional[pulumi.Input[str]] = None,
                 id: Optional[pulumi.Input[str]] = None,
                 kind: Optional[pulumi.Input[str]] = None,
                 min_tls_version: Optional[pulumi.Input[str]] = None,
                 name: Optional[pulumi.Input[str]] = None,
                 profile: Optional[pulumi.Input[str]] = None,
                 request_id: Optional[pulumi.Input[str]] = None,
                 self_link: Optional[pulumi.Input[str]] = None,
                 warnings: Optional[pulumi.Input[Sequence[pulumi.Input['SslPolicyWarningsItemArgs']]]] = None):
        """
        The set of arguments for constructing a SslPolicy resource.
        :param pulumi.Input[str] creation_timestamp: [Output Only] Creation timestamp in RFC3339 text format.
        :param pulumi.Input[Sequence[pulumi.Input[str]]] custom_features: A list of features enabled when the selected profile is CUSTOM. The
               - method returns the set of features that can be specified in this list. This field must be empty if the profile is not CUSTOM.
        :param pulumi.Input[str] description: An optional description of this resource. Provide this property when you create the resource.
        :param pulumi.Input[Sequence[pulumi.Input[str]]] enabled_features: [Output Only] The list of features enabled in the SSL policy.
        :param pulumi.Input[str] fingerprint: Fingerprint of this resource. A hash of the contents stored in this object. This field is used in optimistic locking. This field will be ignored when inserting a SslPolicy. An up-to-date fingerprint must be provided in order to update the SslPolicy, otherwise the request will fail with error 412 conditionNotMet.
               
               To see the latest fingerprint, make a get() request to retrieve an SslPolicy.
        :param pulumi.Input[str] id: [Output Only] The unique identifier for the resource. This identifier is defined by the server.
        :param pulumi.Input[str] kind: [Output only] Type of the resource. Always compute#sslPolicyfor SSL policies.
        :param pulumi.Input[str] min_tls_version: The minimum version of SSL protocol that can be used by the clients to establish a connection with the load balancer. This can be one of TLS_1_0, TLS_1_1, TLS_1_2.
        :param pulumi.Input[str] name: Name of the resource. The name must be 1-63 characters long, and comply with RFC1035. Specifically, the name must be 1-63 characters long and match the regular expression `[a-z]([-a-z0-9]*[a-z0-9])?` which means the first character must be a lowercase letter, and all following characters must be a dash, lowercase letter, or digit, except the last character, which cannot be a dash.
        :param pulumi.Input[str] profile: Profile specifies the set of SSL features that can be used by the load balancer when negotiating SSL with clients. This can be one of COMPATIBLE, MODERN, RESTRICTED, or CUSTOM. If using CUSTOM, the set of SSL features to enable must be specified in the customFeatures field.
        :param pulumi.Input[str] self_link: [Output Only] Server-defined URL for the resource.
        :param pulumi.Input[Sequence[pulumi.Input['SslPolicyWarningsItemArgs']]] warnings: [Output Only] If potential misconfigurations are detected for this SSL policy, this field will be populated with warning messages.
        """
        pulumi.set(__self__, "project", project)
        if creation_timestamp is not None:
            pulumi.set(__self__, "creation_timestamp", creation_timestamp)
        if custom_features is not None:
            pulumi.set(__self__, "custom_features", custom_features)
        if description is not None:
            pulumi.set(__self__, "description", description)
        if enabled_features is not None:
            pulumi.set(__self__, "enabled_features", enabled_features)
        if fingerprint is not None:
            pulumi.set(__self__, "fingerprint", fingerprint)
        if id is not None:
            pulumi.set(__self__, "id", id)
        if kind is not None:
            pulumi.set(__self__, "kind", kind)
        if min_tls_version is not None:
            pulumi.set(__self__, "min_tls_version", min_tls_version)
        if name is not None:
            pulumi.set(__self__, "name", name)
        if profile is not None:
            pulumi.set(__self__, "profile", profile)
        if request_id is not None:
            pulumi.set(__self__, "request_id", request_id)
        if self_link is not None:
            pulumi.set(__self__, "self_link", self_link)
        if warnings is not None:
            pulumi.set(__self__, "warnings", warnings)

    @property
    @pulumi.getter
    def project(self) -> pulumi.Input[str]:
        return pulumi.get(self, "project")

    @project.setter
    def project(self, value: pulumi.Input[str]):
        pulumi.set(self, "project", value)

    @property
    @pulumi.getter(name="creationTimestamp")
    def creation_timestamp(self) -> Optional[pulumi.Input[str]]:
        """
        [Output Only] Creation timestamp in RFC3339 text format.
        """
        return pulumi.get(self, "creation_timestamp")

    @creation_timestamp.setter
    def creation_timestamp(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "creation_timestamp", value)

    @property
    @pulumi.getter(name="customFeatures")
    def custom_features(self) -> Optional[pulumi.Input[Sequence[pulumi.Input[str]]]]:
        """
        A list of features enabled when the selected profile is CUSTOM. The
        - method returns the set of features that can be specified in this list. This field must be empty if the profile is not CUSTOM.
        """
        return pulumi.get(self, "custom_features")

    @custom_features.setter
    def custom_features(self, value: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]]):
        pulumi.set(self, "custom_features", value)

    @property
    @pulumi.getter
    def description(self) -> Optional[pulumi.Input[str]]:
        """
        An optional description of this resource. Provide this property when you create the resource.
        """
        return pulumi.get(self, "description")

    @description.setter
    def description(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "description", value)

    @property
    @pulumi.getter(name="enabledFeatures")
    def enabled_features(self) -> Optional[pulumi.Input[Sequence[pulumi.Input[str]]]]:
        """
        [Output Only] The list of features enabled in the SSL policy.
        """
        return pulumi.get(self, "enabled_features")

    @enabled_features.setter
    def enabled_features(self, value: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]]):
        pulumi.set(self, "enabled_features", value)

    @property
    @pulumi.getter
    def fingerprint(self) -> Optional[pulumi.Input[str]]:
        """
        Fingerprint of this resource. A hash of the contents stored in this object. This field is used in optimistic locking. This field will be ignored when inserting a SslPolicy. An up-to-date fingerprint must be provided in order to update the SslPolicy, otherwise the request will fail with error 412 conditionNotMet.

        To see the latest fingerprint, make a get() request to retrieve an SslPolicy.
        """
        return pulumi.get(self, "fingerprint")

    @fingerprint.setter
    def fingerprint(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "fingerprint", value)

    @property
    @pulumi.getter
    def id(self) -> Optional[pulumi.Input[str]]:
        """
        [Output Only] The unique identifier for the resource. This identifier is defined by the server.
        """
        return pulumi.get(self, "id")

    @id.setter
    def id(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "id", value)

    @property
    @pulumi.getter
    def kind(self) -> Optional[pulumi.Input[str]]:
        """
        [Output only] Type of the resource. Always compute#sslPolicyfor SSL policies.
        """
        return pulumi.get(self, "kind")

    @kind.setter
    def kind(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "kind", value)

    @property
    @pulumi.getter(name="minTlsVersion")
    def min_tls_version(self) -> Optional[pulumi.Input[str]]:
        """
        The minimum version of SSL protocol that can be used by the clients to establish a connection with the load balancer. This can be one of TLS_1_0, TLS_1_1, TLS_1_2.
        """
        return pulumi.get(self, "min_tls_version")

    @min_tls_version.setter
    def min_tls_version(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "min_tls_version", value)

    @property
    @pulumi.getter
    def name(self) -> Optional[pulumi.Input[str]]:
        """
        Name of the resource. The name must be 1-63 characters long, and comply with RFC1035. Specifically, the name must be 1-63 characters long and match the regular expression `[a-z]([-a-z0-9]*[a-z0-9])?` which means the first character must be a lowercase letter, and all following characters must be a dash, lowercase letter, or digit, except the last character, which cannot be a dash.
        """
        return pulumi.get(self, "name")

    @name.setter
    def name(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "name", value)

    @property
    @pulumi.getter
    def profile(self) -> Optional[pulumi.Input[str]]:
        """
        Profile specifies the set of SSL features that can be used by the load balancer when negotiating SSL with clients. This can be one of COMPATIBLE, MODERN, RESTRICTED, or CUSTOM. If using CUSTOM, the set of SSL features to enable must be specified in the customFeatures field.
        """
        return pulumi.get(self, "profile")

    @profile.setter
    def profile(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "profile", value)

    @property
    @pulumi.getter(name="requestId")
    def request_id(self) -> Optional[pulumi.Input[str]]:
        return pulumi.get(self, "request_id")

    @request_id.setter
    def request_id(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "request_id", value)

    @property
    @pulumi.getter(name="selfLink")
    def self_link(self) -> Optional[pulumi.Input[str]]:
        """
        [Output Only] Server-defined URL for the resource.
        """
        return pulumi.get(self, "self_link")

    @self_link.setter
    def self_link(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "self_link", value)

    @property
    @pulumi.getter
    def warnings(self) -> Optional[pulumi.Input[Sequence[pulumi.Input['SslPolicyWarningsItemArgs']]]]:
        """
        [Output Only] If potential misconfigurations are detected for this SSL policy, this field will be populated with warning messages.
        """
        return pulumi.get(self, "warnings")

    @warnings.setter
    def warnings(self, value: Optional[pulumi.Input[Sequence[pulumi.Input['SslPolicyWarningsItemArgs']]]]):
        pulumi.set(self, "warnings", value)


class SslPolicy(pulumi.CustomResource):
    @overload
    def __init__(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 creation_timestamp: Optional[pulumi.Input[str]] = None,
                 custom_features: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]] = None,
                 description: Optional[pulumi.Input[str]] = None,
                 enabled_features: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]] = None,
                 fingerprint: Optional[pulumi.Input[str]] = None,
                 id: Optional[pulumi.Input[str]] = None,
                 kind: Optional[pulumi.Input[str]] = None,
                 min_tls_version: Optional[pulumi.Input[str]] = None,
                 name: Optional[pulumi.Input[str]] = None,
                 profile: Optional[pulumi.Input[str]] = None,
                 project: Optional[pulumi.Input[str]] = None,
                 request_id: Optional[pulumi.Input[str]] = None,
                 self_link: Optional[pulumi.Input[str]] = None,
                 warnings: Optional[pulumi.Input[Sequence[pulumi.Input[pulumi.InputType['SslPolicyWarningsItemArgs']]]]] = None,
                 __props__=None):
        """
        Returns the specified SSL policy resource. Gets a list of available SSL policies by making a list() request.

        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] creation_timestamp: [Output Only] Creation timestamp in RFC3339 text format.
        :param pulumi.Input[Sequence[pulumi.Input[str]]] custom_features: A list of features enabled when the selected profile is CUSTOM. The
               - method returns the set of features that can be specified in this list. This field must be empty if the profile is not CUSTOM.
        :param pulumi.Input[str] description: An optional description of this resource. Provide this property when you create the resource.
        :param pulumi.Input[Sequence[pulumi.Input[str]]] enabled_features: [Output Only] The list of features enabled in the SSL policy.
        :param pulumi.Input[str] fingerprint: Fingerprint of this resource. A hash of the contents stored in this object. This field is used in optimistic locking. This field will be ignored when inserting a SslPolicy. An up-to-date fingerprint must be provided in order to update the SslPolicy, otherwise the request will fail with error 412 conditionNotMet.
               
               To see the latest fingerprint, make a get() request to retrieve an SslPolicy.
        :param pulumi.Input[str] id: [Output Only] The unique identifier for the resource. This identifier is defined by the server.
        :param pulumi.Input[str] kind: [Output only] Type of the resource. Always compute#sslPolicyfor SSL policies.
        :param pulumi.Input[str] min_tls_version: The minimum version of SSL protocol that can be used by the clients to establish a connection with the load balancer. This can be one of TLS_1_0, TLS_1_1, TLS_1_2.
        :param pulumi.Input[str] name: Name of the resource. The name must be 1-63 characters long, and comply with RFC1035. Specifically, the name must be 1-63 characters long and match the regular expression `[a-z]([-a-z0-9]*[a-z0-9])?` which means the first character must be a lowercase letter, and all following characters must be a dash, lowercase letter, or digit, except the last character, which cannot be a dash.
        :param pulumi.Input[str] profile: Profile specifies the set of SSL features that can be used by the load balancer when negotiating SSL with clients. This can be one of COMPATIBLE, MODERN, RESTRICTED, or CUSTOM. If using CUSTOM, the set of SSL features to enable must be specified in the customFeatures field.
        :param pulumi.Input[str] self_link: [Output Only] Server-defined URL for the resource.
        :param pulumi.Input[Sequence[pulumi.Input[pulumi.InputType['SslPolicyWarningsItemArgs']]]] warnings: [Output Only] If potential misconfigurations are detected for this SSL policy, this field will be populated with warning messages.
        """
        ...
    @overload
    def __init__(__self__,
                 resource_name: str,
                 args: SslPolicyArgs,
                 opts: Optional[pulumi.ResourceOptions] = None):
        """
        Returns the specified SSL policy resource. Gets a list of available SSL policies by making a list() request.

        :param str resource_name: The name of the resource.
        :param SslPolicyArgs args: The arguments to use to populate this resource's properties.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        ...
    def __init__(__self__, resource_name: str, *args, **kwargs):
        resource_args, opts = _utilities.get_resource_args_opts(SslPolicyArgs, pulumi.ResourceOptions, *args, **kwargs)
        if resource_args is not None:
            __self__._internal_init(resource_name, opts, **resource_args.__dict__)
        else:
            __self__._internal_init(resource_name, *args, **kwargs)

    def _internal_init(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 creation_timestamp: Optional[pulumi.Input[str]] = None,
                 custom_features: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]] = None,
                 description: Optional[pulumi.Input[str]] = None,
                 enabled_features: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]] = None,
                 fingerprint: Optional[pulumi.Input[str]] = None,
                 id: Optional[pulumi.Input[str]] = None,
                 kind: Optional[pulumi.Input[str]] = None,
                 min_tls_version: Optional[pulumi.Input[str]] = None,
                 name: Optional[pulumi.Input[str]] = None,
                 profile: Optional[pulumi.Input[str]] = None,
                 project: Optional[pulumi.Input[str]] = None,
                 request_id: Optional[pulumi.Input[str]] = None,
                 self_link: Optional[pulumi.Input[str]] = None,
                 warnings: Optional[pulumi.Input[Sequence[pulumi.Input[pulumi.InputType['SslPolicyWarningsItemArgs']]]]] = None,
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
            __props__ = SslPolicyArgs.__new__(SslPolicyArgs)

            __props__.__dict__["creation_timestamp"] = creation_timestamp
            __props__.__dict__["custom_features"] = custom_features
            __props__.__dict__["description"] = description
            __props__.__dict__["enabled_features"] = enabled_features
            __props__.__dict__["fingerprint"] = fingerprint
            __props__.__dict__["id"] = id
            __props__.__dict__["kind"] = kind
            __props__.__dict__["min_tls_version"] = min_tls_version
            __props__.__dict__["name"] = name
            __props__.__dict__["profile"] = profile
            if project is None and not opts.urn:
                raise TypeError("Missing required property 'project'")
            __props__.__dict__["project"] = project
            __props__.__dict__["request_id"] = request_id
            __props__.__dict__["self_link"] = self_link
            __props__.__dict__["warnings"] = warnings
        super(SslPolicy, __self__).__init__(
            'google-native:compute/beta:SslPolicy',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name: str,
            id: pulumi.Input[str],
            opts: Optional[pulumi.ResourceOptions] = None) -> 'SslPolicy':
        """
        Get an existing SslPolicy resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param pulumi.Input[str] id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = SslPolicyArgs.__new__(SslPolicyArgs)

        __props__.__dict__["creation_timestamp"] = None
        __props__.__dict__["custom_features"] = None
        __props__.__dict__["description"] = None
        __props__.__dict__["enabled_features"] = None
        __props__.__dict__["fingerprint"] = None
        __props__.__dict__["kind"] = None
        __props__.__dict__["min_tls_version"] = None
        __props__.__dict__["name"] = None
        __props__.__dict__["profile"] = None
        __props__.__dict__["self_link"] = None
        __props__.__dict__["warnings"] = None
        return SslPolicy(resource_name, opts=opts, __props__=__props__)

    @property
    @pulumi.getter(name="creationTimestamp")
    def creation_timestamp(self) -> pulumi.Output[str]:
        """
        [Output Only] Creation timestamp in RFC3339 text format.
        """
        return pulumi.get(self, "creation_timestamp")

    @property
    @pulumi.getter(name="customFeatures")
    def custom_features(self) -> pulumi.Output[Sequence[str]]:
        """
        A list of features enabled when the selected profile is CUSTOM. The
        - method returns the set of features that can be specified in this list. This field must be empty if the profile is not CUSTOM.
        """
        return pulumi.get(self, "custom_features")

    @property
    @pulumi.getter
    def description(self) -> pulumi.Output[str]:
        """
        An optional description of this resource. Provide this property when you create the resource.
        """
        return pulumi.get(self, "description")

    @property
    @pulumi.getter(name="enabledFeatures")
    def enabled_features(self) -> pulumi.Output[Sequence[str]]:
        """
        [Output Only] The list of features enabled in the SSL policy.
        """
        return pulumi.get(self, "enabled_features")

    @property
    @pulumi.getter
    def fingerprint(self) -> pulumi.Output[str]:
        """
        Fingerprint of this resource. A hash of the contents stored in this object. This field is used in optimistic locking. This field will be ignored when inserting a SslPolicy. An up-to-date fingerprint must be provided in order to update the SslPolicy, otherwise the request will fail with error 412 conditionNotMet.

        To see the latest fingerprint, make a get() request to retrieve an SslPolicy.
        """
        return pulumi.get(self, "fingerprint")

    @property
    @pulumi.getter
    def kind(self) -> pulumi.Output[str]:
        """
        [Output only] Type of the resource. Always compute#sslPolicyfor SSL policies.
        """
        return pulumi.get(self, "kind")

    @property
    @pulumi.getter(name="minTlsVersion")
    def min_tls_version(self) -> pulumi.Output[str]:
        """
        The minimum version of SSL protocol that can be used by the clients to establish a connection with the load balancer. This can be one of TLS_1_0, TLS_1_1, TLS_1_2.
        """
        return pulumi.get(self, "min_tls_version")

    @property
    @pulumi.getter
    def name(self) -> pulumi.Output[str]:
        """
        Name of the resource. The name must be 1-63 characters long, and comply with RFC1035. Specifically, the name must be 1-63 characters long and match the regular expression `[a-z]([-a-z0-9]*[a-z0-9])?` which means the first character must be a lowercase letter, and all following characters must be a dash, lowercase letter, or digit, except the last character, which cannot be a dash.
        """
        return pulumi.get(self, "name")

    @property
    @pulumi.getter
    def profile(self) -> pulumi.Output[str]:
        """
        Profile specifies the set of SSL features that can be used by the load balancer when negotiating SSL with clients. This can be one of COMPATIBLE, MODERN, RESTRICTED, or CUSTOM. If using CUSTOM, the set of SSL features to enable must be specified in the customFeatures field.
        """
        return pulumi.get(self, "profile")

    @property
    @pulumi.getter(name="selfLink")
    def self_link(self) -> pulumi.Output[str]:
        """
        [Output Only] Server-defined URL for the resource.
        """
        return pulumi.get(self, "self_link")

    @property
    @pulumi.getter
    def warnings(self) -> pulumi.Output[Sequence['outputs.SslPolicyWarningsItemResponse']]:
        """
        [Output Only] If potential misconfigurations are detected for this SSL policy, this field will be populated with warning messages.
        """
        return pulumi.get(self, "warnings")

