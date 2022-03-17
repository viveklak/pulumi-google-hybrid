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

__all__ = ['ScanConfigArgs', 'ScanConfig']

@pulumi.input_type
class ScanConfigArgs:
    def __init__(__self__, *,
                 display_name: pulumi.Input[str],
                 starting_urls: pulumi.Input[Sequence[pulumi.Input[str]]],
                 authentication: Optional[pulumi.Input['AuthenticationArgs']] = None,
                 blacklist_patterns: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]] = None,
                 export_to_security_command_center: Optional[pulumi.Input['ScanConfigExportToSecurityCommandCenter']] = None,
                 ignore_http_status_errors: Optional[pulumi.Input[bool]] = None,
                 managed_scan: Optional[pulumi.Input[bool]] = None,
                 max_qps: Optional[pulumi.Input[int]] = None,
                 name: Optional[pulumi.Input[str]] = None,
                 project: Optional[pulumi.Input[str]] = None,
                 risk_level: Optional[pulumi.Input['ScanConfigRiskLevel']] = None,
                 schedule: Optional[pulumi.Input['ScheduleArgs']] = None,
                 static_ip_scan: Optional[pulumi.Input[bool]] = None,
                 user_agent: Optional[pulumi.Input['ScanConfigUserAgent']] = None):
        """
        The set of arguments for constructing a ScanConfig resource.
        :param pulumi.Input[str] display_name: The user provided display name of the ScanConfig.
        :param pulumi.Input[Sequence[pulumi.Input[str]]] starting_urls: The starting URLs from which the scanner finds site pages.
        :param pulumi.Input['AuthenticationArgs'] authentication: The authentication configuration. If specified, service will use the authentication configuration during scanning.
        :param pulumi.Input[Sequence[pulumi.Input[str]]] blacklist_patterns: The excluded URL patterns as described in https://cloud.google.com/security-command-center/docs/how-to-use-web-security-scanner#excluding_urls
        :param pulumi.Input['ScanConfigExportToSecurityCommandCenter'] export_to_security_command_center: Controls export of scan configurations and results to Security Command Center.
        :param pulumi.Input[bool] ignore_http_status_errors: Whether to keep scanning even if most requests return HTTP error codes.
        :param pulumi.Input[bool] managed_scan: Whether the scan config is managed by Web Security Scanner, output only.
        :param pulumi.Input[int] max_qps: The maximum QPS during scanning. A valid value ranges from 5 to 20 inclusively. If the field is unspecified or its value is set 0, server will default to 15. Other values outside of [5, 20] range will be rejected with INVALID_ARGUMENT error.
        :param pulumi.Input[str] name: The resource name of the ScanConfig. The name follows the format of 'projects/{projectId}/scanConfigs/{scanConfigId}'. The ScanConfig IDs are generated by the system.
        :param pulumi.Input['ScanConfigRiskLevel'] risk_level: The risk level selected for the scan
        :param pulumi.Input['ScheduleArgs'] schedule: The schedule of the ScanConfig.
        :param pulumi.Input[bool] static_ip_scan: Whether the scan configuration has enabled static IP address scan feature. If enabled, the scanner will access applications from static IP addresses.
        :param pulumi.Input['ScanConfigUserAgent'] user_agent: The user agent used during scanning.
        """
        pulumi.set(__self__, "display_name", display_name)
        pulumi.set(__self__, "starting_urls", starting_urls)
        if authentication is not None:
            pulumi.set(__self__, "authentication", authentication)
        if blacklist_patterns is not None:
            pulumi.set(__self__, "blacklist_patterns", blacklist_patterns)
        if export_to_security_command_center is not None:
            pulumi.set(__self__, "export_to_security_command_center", export_to_security_command_center)
        if ignore_http_status_errors is not None:
            pulumi.set(__self__, "ignore_http_status_errors", ignore_http_status_errors)
        if managed_scan is not None:
            pulumi.set(__self__, "managed_scan", managed_scan)
        if max_qps is not None:
            pulumi.set(__self__, "max_qps", max_qps)
        if name is not None:
            pulumi.set(__self__, "name", name)
        if project is not None:
            pulumi.set(__self__, "project", project)
        if risk_level is not None:
            pulumi.set(__self__, "risk_level", risk_level)
        if schedule is not None:
            pulumi.set(__self__, "schedule", schedule)
        if static_ip_scan is not None:
            pulumi.set(__self__, "static_ip_scan", static_ip_scan)
        if user_agent is not None:
            pulumi.set(__self__, "user_agent", user_agent)

    @property
    @pulumi.getter(name="displayName")
    def display_name(self) -> pulumi.Input[str]:
        """
        The user provided display name of the ScanConfig.
        """
        return pulumi.get(self, "display_name")

    @display_name.setter
    def display_name(self, value: pulumi.Input[str]):
        pulumi.set(self, "display_name", value)

    @property
    @pulumi.getter(name="startingUrls")
    def starting_urls(self) -> pulumi.Input[Sequence[pulumi.Input[str]]]:
        """
        The starting URLs from which the scanner finds site pages.
        """
        return pulumi.get(self, "starting_urls")

    @starting_urls.setter
    def starting_urls(self, value: pulumi.Input[Sequence[pulumi.Input[str]]]):
        pulumi.set(self, "starting_urls", value)

    @property
    @pulumi.getter
    def authentication(self) -> Optional[pulumi.Input['AuthenticationArgs']]:
        """
        The authentication configuration. If specified, service will use the authentication configuration during scanning.
        """
        return pulumi.get(self, "authentication")

    @authentication.setter
    def authentication(self, value: Optional[pulumi.Input['AuthenticationArgs']]):
        pulumi.set(self, "authentication", value)

    @property
    @pulumi.getter(name="blacklistPatterns")
    def blacklist_patterns(self) -> Optional[pulumi.Input[Sequence[pulumi.Input[str]]]]:
        """
        The excluded URL patterns as described in https://cloud.google.com/security-command-center/docs/how-to-use-web-security-scanner#excluding_urls
        """
        return pulumi.get(self, "blacklist_patterns")

    @blacklist_patterns.setter
    def blacklist_patterns(self, value: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]]):
        pulumi.set(self, "blacklist_patterns", value)

    @property
    @pulumi.getter(name="exportToSecurityCommandCenter")
    def export_to_security_command_center(self) -> Optional[pulumi.Input['ScanConfigExportToSecurityCommandCenter']]:
        """
        Controls export of scan configurations and results to Security Command Center.
        """
        return pulumi.get(self, "export_to_security_command_center")

    @export_to_security_command_center.setter
    def export_to_security_command_center(self, value: Optional[pulumi.Input['ScanConfigExportToSecurityCommandCenter']]):
        pulumi.set(self, "export_to_security_command_center", value)

    @property
    @pulumi.getter(name="ignoreHttpStatusErrors")
    def ignore_http_status_errors(self) -> Optional[pulumi.Input[bool]]:
        """
        Whether to keep scanning even if most requests return HTTP error codes.
        """
        return pulumi.get(self, "ignore_http_status_errors")

    @ignore_http_status_errors.setter
    def ignore_http_status_errors(self, value: Optional[pulumi.Input[bool]]):
        pulumi.set(self, "ignore_http_status_errors", value)

    @property
    @pulumi.getter(name="managedScan")
    def managed_scan(self) -> Optional[pulumi.Input[bool]]:
        """
        Whether the scan config is managed by Web Security Scanner, output only.
        """
        return pulumi.get(self, "managed_scan")

    @managed_scan.setter
    def managed_scan(self, value: Optional[pulumi.Input[bool]]):
        pulumi.set(self, "managed_scan", value)

    @property
    @pulumi.getter(name="maxQps")
    def max_qps(self) -> Optional[pulumi.Input[int]]:
        """
        The maximum QPS during scanning. A valid value ranges from 5 to 20 inclusively. If the field is unspecified or its value is set 0, server will default to 15. Other values outside of [5, 20] range will be rejected with INVALID_ARGUMENT error.
        """
        return pulumi.get(self, "max_qps")

    @max_qps.setter
    def max_qps(self, value: Optional[pulumi.Input[int]]):
        pulumi.set(self, "max_qps", value)

    @property
    @pulumi.getter
    def name(self) -> Optional[pulumi.Input[str]]:
        """
        The resource name of the ScanConfig. The name follows the format of 'projects/{projectId}/scanConfigs/{scanConfigId}'. The ScanConfig IDs are generated by the system.
        """
        return pulumi.get(self, "name")

    @name.setter
    def name(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "name", value)

    @property
    @pulumi.getter
    def project(self) -> Optional[pulumi.Input[str]]:
        return pulumi.get(self, "project")

    @project.setter
    def project(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "project", value)

    @property
    @pulumi.getter(name="riskLevel")
    def risk_level(self) -> Optional[pulumi.Input['ScanConfigRiskLevel']]:
        """
        The risk level selected for the scan
        """
        return pulumi.get(self, "risk_level")

    @risk_level.setter
    def risk_level(self, value: Optional[pulumi.Input['ScanConfigRiskLevel']]):
        pulumi.set(self, "risk_level", value)

    @property
    @pulumi.getter
    def schedule(self) -> Optional[pulumi.Input['ScheduleArgs']]:
        """
        The schedule of the ScanConfig.
        """
        return pulumi.get(self, "schedule")

    @schedule.setter
    def schedule(self, value: Optional[pulumi.Input['ScheduleArgs']]):
        pulumi.set(self, "schedule", value)

    @property
    @pulumi.getter(name="staticIpScan")
    def static_ip_scan(self) -> Optional[pulumi.Input[bool]]:
        """
        Whether the scan configuration has enabled static IP address scan feature. If enabled, the scanner will access applications from static IP addresses.
        """
        return pulumi.get(self, "static_ip_scan")

    @static_ip_scan.setter
    def static_ip_scan(self, value: Optional[pulumi.Input[bool]]):
        pulumi.set(self, "static_ip_scan", value)

    @property
    @pulumi.getter(name="userAgent")
    def user_agent(self) -> Optional[pulumi.Input['ScanConfigUserAgent']]:
        """
        The user agent used during scanning.
        """
        return pulumi.get(self, "user_agent")

    @user_agent.setter
    def user_agent(self, value: Optional[pulumi.Input['ScanConfigUserAgent']]):
        pulumi.set(self, "user_agent", value)


class ScanConfig(pulumi.CustomResource):
    @overload
    def __init__(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 authentication: Optional[pulumi.Input[pulumi.InputType['AuthenticationArgs']]] = None,
                 blacklist_patterns: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]] = None,
                 display_name: Optional[pulumi.Input[str]] = None,
                 export_to_security_command_center: Optional[pulumi.Input['ScanConfigExportToSecurityCommandCenter']] = None,
                 ignore_http_status_errors: Optional[pulumi.Input[bool]] = None,
                 managed_scan: Optional[pulumi.Input[bool]] = None,
                 max_qps: Optional[pulumi.Input[int]] = None,
                 name: Optional[pulumi.Input[str]] = None,
                 project: Optional[pulumi.Input[str]] = None,
                 risk_level: Optional[pulumi.Input['ScanConfigRiskLevel']] = None,
                 schedule: Optional[pulumi.Input[pulumi.InputType['ScheduleArgs']]] = None,
                 starting_urls: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]] = None,
                 static_ip_scan: Optional[pulumi.Input[bool]] = None,
                 user_agent: Optional[pulumi.Input['ScanConfigUserAgent']] = None,
                 __props__=None):
        """
        Creates a new ScanConfig.

        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[pulumi.InputType['AuthenticationArgs']] authentication: The authentication configuration. If specified, service will use the authentication configuration during scanning.
        :param pulumi.Input[Sequence[pulumi.Input[str]]] blacklist_patterns: The excluded URL patterns as described in https://cloud.google.com/security-command-center/docs/how-to-use-web-security-scanner#excluding_urls
        :param pulumi.Input[str] display_name: The user provided display name of the ScanConfig.
        :param pulumi.Input['ScanConfigExportToSecurityCommandCenter'] export_to_security_command_center: Controls export of scan configurations and results to Security Command Center.
        :param pulumi.Input[bool] ignore_http_status_errors: Whether to keep scanning even if most requests return HTTP error codes.
        :param pulumi.Input[bool] managed_scan: Whether the scan config is managed by Web Security Scanner, output only.
        :param pulumi.Input[int] max_qps: The maximum QPS during scanning. A valid value ranges from 5 to 20 inclusively. If the field is unspecified or its value is set 0, server will default to 15. Other values outside of [5, 20] range will be rejected with INVALID_ARGUMENT error.
        :param pulumi.Input[str] name: The resource name of the ScanConfig. The name follows the format of 'projects/{projectId}/scanConfigs/{scanConfigId}'. The ScanConfig IDs are generated by the system.
        :param pulumi.Input['ScanConfigRiskLevel'] risk_level: The risk level selected for the scan
        :param pulumi.Input[pulumi.InputType['ScheduleArgs']] schedule: The schedule of the ScanConfig.
        :param pulumi.Input[Sequence[pulumi.Input[str]]] starting_urls: The starting URLs from which the scanner finds site pages.
        :param pulumi.Input[bool] static_ip_scan: Whether the scan configuration has enabled static IP address scan feature. If enabled, the scanner will access applications from static IP addresses.
        :param pulumi.Input['ScanConfigUserAgent'] user_agent: The user agent used during scanning.
        """
        ...
    @overload
    def __init__(__self__,
                 resource_name: str,
                 args: ScanConfigArgs,
                 opts: Optional[pulumi.ResourceOptions] = None):
        """
        Creates a new ScanConfig.

        :param str resource_name: The name of the resource.
        :param ScanConfigArgs args: The arguments to use to populate this resource's properties.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        ...
    def __init__(__self__, resource_name: str, *args, **kwargs):
        resource_args, opts = _utilities.get_resource_args_opts(ScanConfigArgs, pulumi.ResourceOptions, *args, **kwargs)
        if resource_args is not None:
            __self__._internal_init(resource_name, opts, **resource_args.__dict__)
        else:
            __self__._internal_init(resource_name, *args, **kwargs)

    def _internal_init(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 authentication: Optional[pulumi.Input[pulumi.InputType['AuthenticationArgs']]] = None,
                 blacklist_patterns: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]] = None,
                 display_name: Optional[pulumi.Input[str]] = None,
                 export_to_security_command_center: Optional[pulumi.Input['ScanConfigExportToSecurityCommandCenter']] = None,
                 ignore_http_status_errors: Optional[pulumi.Input[bool]] = None,
                 managed_scan: Optional[pulumi.Input[bool]] = None,
                 max_qps: Optional[pulumi.Input[int]] = None,
                 name: Optional[pulumi.Input[str]] = None,
                 project: Optional[pulumi.Input[str]] = None,
                 risk_level: Optional[pulumi.Input['ScanConfigRiskLevel']] = None,
                 schedule: Optional[pulumi.Input[pulumi.InputType['ScheduleArgs']]] = None,
                 starting_urls: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]] = None,
                 static_ip_scan: Optional[pulumi.Input[bool]] = None,
                 user_agent: Optional[pulumi.Input['ScanConfigUserAgent']] = None,
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
            __props__ = ScanConfigArgs.__new__(ScanConfigArgs)

            __props__.__dict__["authentication"] = authentication
            __props__.__dict__["blacklist_patterns"] = blacklist_patterns
            if display_name is None and not opts.urn:
                raise TypeError("Missing required property 'display_name'")
            __props__.__dict__["display_name"] = display_name
            __props__.__dict__["export_to_security_command_center"] = export_to_security_command_center
            __props__.__dict__["ignore_http_status_errors"] = ignore_http_status_errors
            __props__.__dict__["managed_scan"] = managed_scan
            __props__.__dict__["max_qps"] = max_qps
            __props__.__dict__["name"] = name
            __props__.__dict__["project"] = project
            __props__.__dict__["risk_level"] = risk_level
            __props__.__dict__["schedule"] = schedule
            if starting_urls is None and not opts.urn:
                raise TypeError("Missing required property 'starting_urls'")
            __props__.__dict__["starting_urls"] = starting_urls
            __props__.__dict__["static_ip_scan"] = static_ip_scan
            __props__.__dict__["user_agent"] = user_agent
        super(ScanConfig, __self__).__init__(
            'google-native:websecurityscanner/v1:ScanConfig',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name: str,
            id: pulumi.Input[str],
            opts: Optional[pulumi.ResourceOptions] = None) -> 'ScanConfig':
        """
        Get an existing ScanConfig resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param pulumi.Input[str] id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = ScanConfigArgs.__new__(ScanConfigArgs)

        __props__.__dict__["authentication"] = None
        __props__.__dict__["blacklist_patterns"] = None
        __props__.__dict__["display_name"] = None
        __props__.__dict__["export_to_security_command_center"] = None
        __props__.__dict__["ignore_http_status_errors"] = None
        __props__.__dict__["managed_scan"] = None
        __props__.__dict__["max_qps"] = None
        __props__.__dict__["name"] = None
        __props__.__dict__["risk_level"] = None
        __props__.__dict__["schedule"] = None
        __props__.__dict__["starting_urls"] = None
        __props__.__dict__["static_ip_scan"] = None
        __props__.__dict__["user_agent"] = None
        return ScanConfig(resource_name, opts=opts, __props__=__props__)

    @property
    @pulumi.getter
    def authentication(self) -> pulumi.Output['outputs.AuthenticationResponse']:
        """
        The authentication configuration. If specified, service will use the authentication configuration during scanning.
        """
        return pulumi.get(self, "authentication")

    @property
    @pulumi.getter(name="blacklistPatterns")
    def blacklist_patterns(self) -> pulumi.Output[Sequence[str]]:
        """
        The excluded URL patterns as described in https://cloud.google.com/security-command-center/docs/how-to-use-web-security-scanner#excluding_urls
        """
        return pulumi.get(self, "blacklist_patterns")

    @property
    @pulumi.getter(name="displayName")
    def display_name(self) -> pulumi.Output[str]:
        """
        The user provided display name of the ScanConfig.
        """
        return pulumi.get(self, "display_name")

    @property
    @pulumi.getter(name="exportToSecurityCommandCenter")
    def export_to_security_command_center(self) -> pulumi.Output[str]:
        """
        Controls export of scan configurations and results to Security Command Center.
        """
        return pulumi.get(self, "export_to_security_command_center")

    @property
    @pulumi.getter(name="ignoreHttpStatusErrors")
    def ignore_http_status_errors(self) -> pulumi.Output[bool]:
        """
        Whether to keep scanning even if most requests return HTTP error codes.
        """
        return pulumi.get(self, "ignore_http_status_errors")

    @property
    @pulumi.getter(name="managedScan")
    def managed_scan(self) -> pulumi.Output[bool]:
        """
        Whether the scan config is managed by Web Security Scanner, output only.
        """
        return pulumi.get(self, "managed_scan")

    @property
    @pulumi.getter(name="maxQps")
    def max_qps(self) -> pulumi.Output[int]:
        """
        The maximum QPS during scanning. A valid value ranges from 5 to 20 inclusively. If the field is unspecified or its value is set 0, server will default to 15. Other values outside of [5, 20] range will be rejected with INVALID_ARGUMENT error.
        """
        return pulumi.get(self, "max_qps")

    @property
    @pulumi.getter
    def name(self) -> pulumi.Output[str]:
        """
        The resource name of the ScanConfig. The name follows the format of 'projects/{projectId}/scanConfigs/{scanConfigId}'. The ScanConfig IDs are generated by the system.
        """
        return pulumi.get(self, "name")

    @property
    @pulumi.getter(name="riskLevel")
    def risk_level(self) -> pulumi.Output[str]:
        """
        The risk level selected for the scan
        """
        return pulumi.get(self, "risk_level")

    @property
    @pulumi.getter
    def schedule(self) -> pulumi.Output['outputs.ScheduleResponse']:
        """
        The schedule of the ScanConfig.
        """
        return pulumi.get(self, "schedule")

    @property
    @pulumi.getter(name="startingUrls")
    def starting_urls(self) -> pulumi.Output[Sequence[str]]:
        """
        The starting URLs from which the scanner finds site pages.
        """
        return pulumi.get(self, "starting_urls")

    @property
    @pulumi.getter(name="staticIpScan")
    def static_ip_scan(self) -> pulumi.Output[bool]:
        """
        Whether the scan configuration has enabled static IP address scan feature. If enabled, the scanner will access applications from static IP addresses.
        """
        return pulumi.get(self, "static_ip_scan")

    @property
    @pulumi.getter(name="userAgent")
    def user_agent(self) -> pulumi.Output[str]:
        """
        The user agent used during scanning.
        """
        return pulumi.get(self, "user_agent")

