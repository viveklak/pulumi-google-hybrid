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
    'GetScanConfigResult',
    'AwaitableGetScanConfigResult',
    'get_scan_config',
    'get_scan_config_output',
]

@pulumi.output_type
class GetScanConfigResult:
    def __init__(__self__, authentication=None, blacklist_patterns=None, display_name=None, export_to_security_command_center=None, ignore_http_status_errors=None, managed_scan=None, max_qps=None, name=None, risk_level=None, schedule=None, starting_urls=None, static_ip_scan=None, user_agent=None):
        if authentication and not isinstance(authentication, dict):
            raise TypeError("Expected argument 'authentication' to be a dict")
        pulumi.set(__self__, "authentication", authentication)
        if blacklist_patterns and not isinstance(blacklist_patterns, list):
            raise TypeError("Expected argument 'blacklist_patterns' to be a list")
        pulumi.set(__self__, "blacklist_patterns", blacklist_patterns)
        if display_name and not isinstance(display_name, str):
            raise TypeError("Expected argument 'display_name' to be a str")
        pulumi.set(__self__, "display_name", display_name)
        if export_to_security_command_center and not isinstance(export_to_security_command_center, str):
            raise TypeError("Expected argument 'export_to_security_command_center' to be a str")
        pulumi.set(__self__, "export_to_security_command_center", export_to_security_command_center)
        if ignore_http_status_errors and not isinstance(ignore_http_status_errors, bool):
            raise TypeError("Expected argument 'ignore_http_status_errors' to be a bool")
        pulumi.set(__self__, "ignore_http_status_errors", ignore_http_status_errors)
        if managed_scan and not isinstance(managed_scan, bool):
            raise TypeError("Expected argument 'managed_scan' to be a bool")
        pulumi.set(__self__, "managed_scan", managed_scan)
        if max_qps and not isinstance(max_qps, int):
            raise TypeError("Expected argument 'max_qps' to be a int")
        pulumi.set(__self__, "max_qps", max_qps)
        if name and not isinstance(name, str):
            raise TypeError("Expected argument 'name' to be a str")
        pulumi.set(__self__, "name", name)
        if risk_level and not isinstance(risk_level, str):
            raise TypeError("Expected argument 'risk_level' to be a str")
        pulumi.set(__self__, "risk_level", risk_level)
        if schedule and not isinstance(schedule, dict):
            raise TypeError("Expected argument 'schedule' to be a dict")
        pulumi.set(__self__, "schedule", schedule)
        if starting_urls and not isinstance(starting_urls, list):
            raise TypeError("Expected argument 'starting_urls' to be a list")
        pulumi.set(__self__, "starting_urls", starting_urls)
        if static_ip_scan and not isinstance(static_ip_scan, bool):
            raise TypeError("Expected argument 'static_ip_scan' to be a bool")
        pulumi.set(__self__, "static_ip_scan", static_ip_scan)
        if user_agent and not isinstance(user_agent, str):
            raise TypeError("Expected argument 'user_agent' to be a str")
        pulumi.set(__self__, "user_agent", user_agent)

    @property
    @pulumi.getter
    def authentication(self) -> 'outputs.AuthenticationResponse':
        """
        The authentication configuration. If specified, service will use the authentication configuration during scanning.
        """
        return pulumi.get(self, "authentication")

    @property
    @pulumi.getter(name="blacklistPatterns")
    def blacklist_patterns(self) -> Sequence[str]:
        """
        The excluded URL patterns as described in https://cloud.google.com/security-command-center/docs/how-to-use-web-security-scanner#excluding_urls
        """
        return pulumi.get(self, "blacklist_patterns")

    @property
    @pulumi.getter(name="displayName")
    def display_name(self) -> str:
        """
        The user provided display name of the ScanConfig.
        """
        return pulumi.get(self, "display_name")

    @property
    @pulumi.getter(name="exportToSecurityCommandCenter")
    def export_to_security_command_center(self) -> str:
        """
        Controls export of scan configurations and results to Security Command Center.
        """
        return pulumi.get(self, "export_to_security_command_center")

    @property
    @pulumi.getter(name="ignoreHttpStatusErrors")
    def ignore_http_status_errors(self) -> bool:
        """
        Whether to keep scanning even if most requests return HTTP error codes.
        """
        return pulumi.get(self, "ignore_http_status_errors")

    @property
    @pulumi.getter(name="managedScan")
    def managed_scan(self) -> bool:
        """
        Whether the scan config is managed by Web Security Scanner, output only.
        """
        return pulumi.get(self, "managed_scan")

    @property
    @pulumi.getter(name="maxQps")
    def max_qps(self) -> int:
        """
        The maximum QPS during scanning. A valid value ranges from 5 to 20 inclusively. If the field is unspecified or its value is set 0, server will default to 15. Other values outside of [5, 20] range will be rejected with INVALID_ARGUMENT error.
        """
        return pulumi.get(self, "max_qps")

    @property
    @pulumi.getter
    def name(self) -> str:
        """
        The resource name of the ScanConfig. The name follows the format of 'projects/{projectId}/scanConfigs/{scanConfigId}'. The ScanConfig IDs are generated by the system.
        """
        return pulumi.get(self, "name")

    @property
    @pulumi.getter(name="riskLevel")
    def risk_level(self) -> str:
        """
        The risk level selected for the scan
        """
        return pulumi.get(self, "risk_level")

    @property
    @pulumi.getter
    def schedule(self) -> 'outputs.ScheduleResponse':
        """
        The schedule of the ScanConfig.
        """
        return pulumi.get(self, "schedule")

    @property
    @pulumi.getter(name="startingUrls")
    def starting_urls(self) -> Sequence[str]:
        """
        The starting URLs from which the scanner finds site pages.
        """
        return pulumi.get(self, "starting_urls")

    @property
    @pulumi.getter(name="staticIpScan")
    def static_ip_scan(self) -> bool:
        """
        Whether the scan configuration has enabled static IP address scan feature. If enabled, the scanner will access applications from static IP addresses.
        """
        return pulumi.get(self, "static_ip_scan")

    @property
    @pulumi.getter(name="userAgent")
    def user_agent(self) -> str:
        """
        The user agent used during scanning.
        """
        return pulumi.get(self, "user_agent")


class AwaitableGetScanConfigResult(GetScanConfigResult):
    # pylint: disable=using-constant-test
    def __await__(self):
        if False:
            yield self
        return GetScanConfigResult(
            authentication=self.authentication,
            blacklist_patterns=self.blacklist_patterns,
            display_name=self.display_name,
            export_to_security_command_center=self.export_to_security_command_center,
            ignore_http_status_errors=self.ignore_http_status_errors,
            managed_scan=self.managed_scan,
            max_qps=self.max_qps,
            name=self.name,
            risk_level=self.risk_level,
            schedule=self.schedule,
            starting_urls=self.starting_urls,
            static_ip_scan=self.static_ip_scan,
            user_agent=self.user_agent)


def get_scan_config(project: Optional[str] = None,
                    scan_config_id: Optional[str] = None,
                    opts: Optional[pulumi.InvokeOptions] = None) -> AwaitableGetScanConfigResult:
    """
    Gets a ScanConfig.
    """
    __args__ = dict()
    __args__['project'] = project
    __args__['scanConfigId'] = scan_config_id
    if opts is None:
        opts = pulumi.InvokeOptions()
    if opts.version is None:
        opts.version = _utilities.get_version()
    __ret__ = pulumi.runtime.invoke('google-native:websecurityscanner/v1:getScanConfig', __args__, opts=opts, typ=GetScanConfigResult).value

    return AwaitableGetScanConfigResult(
        authentication=__ret__.authentication,
        blacklist_patterns=__ret__.blacklist_patterns,
        display_name=__ret__.display_name,
        export_to_security_command_center=__ret__.export_to_security_command_center,
        ignore_http_status_errors=__ret__.ignore_http_status_errors,
        managed_scan=__ret__.managed_scan,
        max_qps=__ret__.max_qps,
        name=__ret__.name,
        risk_level=__ret__.risk_level,
        schedule=__ret__.schedule,
        starting_urls=__ret__.starting_urls,
        static_ip_scan=__ret__.static_ip_scan,
        user_agent=__ret__.user_agent)


@_utilities.lift_output_func(get_scan_config)
def get_scan_config_output(project: Optional[pulumi.Input[Optional[str]]] = None,
                           scan_config_id: Optional[pulumi.Input[str]] = None,
                           opts: Optional[pulumi.InvokeOptions] = None) -> pulumi.Output[GetScanConfigResult]:
    """
    Gets a ScanConfig.
    """
    ...
