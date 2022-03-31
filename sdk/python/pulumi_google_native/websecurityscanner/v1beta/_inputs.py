# coding=utf-8
# *** WARNING: this file was generated by the Pulumi SDK Generator. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import warnings
import pulumi
import pulumi.runtime
from typing import Any, Mapping, Optional, Sequence, Union, overload
from ... import _utilities
from ._enums import *

__all__ = [
    'AuthenticationArgs',
    'CustomAccountArgs',
    'GoogleAccountArgs',
    'IapCredentialArgs',
    'IapTestServiceAccountInfoArgs',
    'ScanConfigErrorArgs',
    'ScanRunErrorTraceArgs',
    'ScanRunWarningTraceArgs',
    'ScanRunArgs',
    'ScheduleArgs',
]

@pulumi.input_type
class AuthenticationArgs:
    def __init__(__self__, *,
                 custom_account: Optional[pulumi.Input['CustomAccountArgs']] = None,
                 google_account: Optional[pulumi.Input['GoogleAccountArgs']] = None,
                 iap_credential: Optional[pulumi.Input['IapCredentialArgs']] = None):
        """
        Scan authentication configuration.
        :param pulumi.Input['CustomAccountArgs'] custom_account: Authentication using a custom account.
        :param pulumi.Input['GoogleAccountArgs'] google_account: Authentication using a Google account.
        :param pulumi.Input['IapCredentialArgs'] iap_credential: Authentication using Identity-Aware-Proxy (IAP).
        """
        if custom_account is not None:
            pulumi.set(__self__, "custom_account", custom_account)
        if google_account is not None:
            pulumi.set(__self__, "google_account", google_account)
        if iap_credential is not None:
            pulumi.set(__self__, "iap_credential", iap_credential)

    @property
    @pulumi.getter(name="customAccount")
    def custom_account(self) -> Optional[pulumi.Input['CustomAccountArgs']]:
        """
        Authentication using a custom account.
        """
        return pulumi.get(self, "custom_account")

    @custom_account.setter
    def custom_account(self, value: Optional[pulumi.Input['CustomAccountArgs']]):
        pulumi.set(self, "custom_account", value)

    @property
    @pulumi.getter(name="googleAccount")
    def google_account(self) -> Optional[pulumi.Input['GoogleAccountArgs']]:
        """
        Authentication using a Google account.
        """
        return pulumi.get(self, "google_account")

    @google_account.setter
    def google_account(self, value: Optional[pulumi.Input['GoogleAccountArgs']]):
        pulumi.set(self, "google_account", value)

    @property
    @pulumi.getter(name="iapCredential")
    def iap_credential(self) -> Optional[pulumi.Input['IapCredentialArgs']]:
        """
        Authentication using Identity-Aware-Proxy (IAP).
        """
        return pulumi.get(self, "iap_credential")

    @iap_credential.setter
    def iap_credential(self, value: Optional[pulumi.Input['IapCredentialArgs']]):
        pulumi.set(self, "iap_credential", value)


@pulumi.input_type
class CustomAccountArgs:
    def __init__(__self__, *,
                 login_url: pulumi.Input[str],
                 password: pulumi.Input[str],
                 username: pulumi.Input[str]):
        """
        Describes authentication configuration that uses a custom account.
        :param pulumi.Input[str] login_url: The login form URL of the website.
        :param pulumi.Input[str] password: Input only. The password of the custom account. The credential is stored encrypted and not returned in any response nor included in audit logs.
        :param pulumi.Input[str] username: The user name of the custom account.
        """
        pulumi.set(__self__, "login_url", login_url)
        pulumi.set(__self__, "password", password)
        pulumi.set(__self__, "username", username)

    @property
    @pulumi.getter(name="loginUrl")
    def login_url(self) -> pulumi.Input[str]:
        """
        The login form URL of the website.
        """
        return pulumi.get(self, "login_url")

    @login_url.setter
    def login_url(self, value: pulumi.Input[str]):
        pulumi.set(self, "login_url", value)

    @property
    @pulumi.getter
    def password(self) -> pulumi.Input[str]:
        """
        Input only. The password of the custom account. The credential is stored encrypted and not returned in any response nor included in audit logs.
        """
        return pulumi.get(self, "password")

    @password.setter
    def password(self, value: pulumi.Input[str]):
        pulumi.set(self, "password", value)

    @property
    @pulumi.getter
    def username(self) -> pulumi.Input[str]:
        """
        The user name of the custom account.
        """
        return pulumi.get(self, "username")

    @username.setter
    def username(self, value: pulumi.Input[str]):
        pulumi.set(self, "username", value)


@pulumi.input_type
class GoogleAccountArgs:
    def __init__(__self__, *,
                 password: pulumi.Input[str],
                 username: pulumi.Input[str]):
        """
        Describes authentication configuration that uses a Google account.
        :param pulumi.Input[str] password: Input only. The password of the Google account. The credential is stored encrypted and not returned in any response nor included in audit logs.
        :param pulumi.Input[str] username: The user name of the Google account.
        """
        pulumi.set(__self__, "password", password)
        pulumi.set(__self__, "username", username)

    @property
    @pulumi.getter
    def password(self) -> pulumi.Input[str]:
        """
        Input only. The password of the Google account. The credential is stored encrypted and not returned in any response nor included in audit logs.
        """
        return pulumi.get(self, "password")

    @password.setter
    def password(self, value: pulumi.Input[str]):
        pulumi.set(self, "password", value)

    @property
    @pulumi.getter
    def username(self) -> pulumi.Input[str]:
        """
        The user name of the Google account.
        """
        return pulumi.get(self, "username")

    @username.setter
    def username(self, value: pulumi.Input[str]):
        pulumi.set(self, "username", value)


@pulumi.input_type
class IapCredentialArgs:
    def __init__(__self__, *,
                 iap_test_service_account_info: Optional[pulumi.Input['IapTestServiceAccountInfoArgs']] = None):
        """
        Describes authentication configuration for Identity-Aware-Proxy (IAP).
        :param pulumi.Input['IapTestServiceAccountInfoArgs'] iap_test_service_account_info: Authentication configuration when Web-Security-Scanner service account is added in Identity-Aware-Proxy (IAP) access policies.
        """
        if iap_test_service_account_info is not None:
            pulumi.set(__self__, "iap_test_service_account_info", iap_test_service_account_info)

    @property
    @pulumi.getter(name="iapTestServiceAccountInfo")
    def iap_test_service_account_info(self) -> Optional[pulumi.Input['IapTestServiceAccountInfoArgs']]:
        """
        Authentication configuration when Web-Security-Scanner service account is added in Identity-Aware-Proxy (IAP) access policies.
        """
        return pulumi.get(self, "iap_test_service_account_info")

    @iap_test_service_account_info.setter
    def iap_test_service_account_info(self, value: Optional[pulumi.Input['IapTestServiceAccountInfoArgs']]):
        pulumi.set(self, "iap_test_service_account_info", value)


@pulumi.input_type
class IapTestServiceAccountInfoArgs:
    def __init__(__self__, *,
                 target_audience_client_id: pulumi.Input[str]):
        """
        Describes authentication configuration when Web-Security-Scanner service account is added in Identity-Aware-Proxy (IAP) access policies.
        :param pulumi.Input[str] target_audience_client_id: Describes OAuth2 Client ID of resources protected by Identity-Aware-Proxy(IAP).
        """
        pulumi.set(__self__, "target_audience_client_id", target_audience_client_id)

    @property
    @pulumi.getter(name="targetAudienceClientId")
    def target_audience_client_id(self) -> pulumi.Input[str]:
        """
        Describes OAuth2 Client ID of resources protected by Identity-Aware-Proxy(IAP).
        """
        return pulumi.get(self, "target_audience_client_id")

    @target_audience_client_id.setter
    def target_audience_client_id(self, value: pulumi.Input[str]):
        pulumi.set(self, "target_audience_client_id", value)


@pulumi.input_type
class ScanConfigErrorArgs:
    def __init__(__self__, *,
                 code: Optional[pulumi.Input['ScanConfigErrorCode']] = None,
                 field_name: Optional[pulumi.Input[str]] = None):
        """
        Defines a custom error message used by CreateScanConfig and UpdateScanConfig APIs when scan configuration validation fails. It is also reported as part of a ScanRunErrorTrace message if scan validation fails due to a scan configuration error.
        :param pulumi.Input['ScanConfigErrorCode'] code: Indicates the reason code for a configuration failure.
        :param pulumi.Input[str] field_name: Indicates the full name of the ScanConfig field that triggers this error, for example "scan_config.max_qps". This field is provided for troubleshooting purposes only and its actual value can change in the future.
        """
        if code is not None:
            pulumi.set(__self__, "code", code)
        if field_name is not None:
            pulumi.set(__self__, "field_name", field_name)

    @property
    @pulumi.getter
    def code(self) -> Optional[pulumi.Input['ScanConfigErrorCode']]:
        """
        Indicates the reason code for a configuration failure.
        """
        return pulumi.get(self, "code")

    @code.setter
    def code(self, value: Optional[pulumi.Input['ScanConfigErrorCode']]):
        pulumi.set(self, "code", value)

    @property
    @pulumi.getter(name="fieldName")
    def field_name(self) -> Optional[pulumi.Input[str]]:
        """
        Indicates the full name of the ScanConfig field that triggers this error, for example "scan_config.max_qps". This field is provided for troubleshooting purposes only and its actual value can change in the future.
        """
        return pulumi.get(self, "field_name")

    @field_name.setter
    def field_name(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "field_name", value)


@pulumi.input_type
class ScanRunErrorTraceArgs:
    def __init__(__self__, *,
                 code: Optional[pulumi.Input['ScanRunErrorTraceCode']] = None,
                 most_common_http_error_code: Optional[pulumi.Input[int]] = None,
                 scan_config_error: Optional[pulumi.Input['ScanConfigErrorArgs']] = None):
        """
        Output only. Defines an error trace message for a ScanRun.
        :param pulumi.Input['ScanRunErrorTraceCode'] code: Indicates the error reason code.
        :param pulumi.Input[int] most_common_http_error_code: If the scan encounters TOO_MANY_HTTP_ERRORS, this field indicates the most common HTTP error code, if such is available. For example, if this code is 404, the scan has encountered too many NOT_FOUND responses.
        :param pulumi.Input['ScanConfigErrorArgs'] scan_config_error: If the scan encounters SCAN_CONFIG_ISSUE error, this field has the error message encountered during scan configuration validation that is performed before each scan run.
        """
        if code is not None:
            pulumi.set(__self__, "code", code)
        if most_common_http_error_code is not None:
            pulumi.set(__self__, "most_common_http_error_code", most_common_http_error_code)
        if scan_config_error is not None:
            pulumi.set(__self__, "scan_config_error", scan_config_error)

    @property
    @pulumi.getter
    def code(self) -> Optional[pulumi.Input['ScanRunErrorTraceCode']]:
        """
        Indicates the error reason code.
        """
        return pulumi.get(self, "code")

    @code.setter
    def code(self, value: Optional[pulumi.Input['ScanRunErrorTraceCode']]):
        pulumi.set(self, "code", value)

    @property
    @pulumi.getter(name="mostCommonHttpErrorCode")
    def most_common_http_error_code(self) -> Optional[pulumi.Input[int]]:
        """
        If the scan encounters TOO_MANY_HTTP_ERRORS, this field indicates the most common HTTP error code, if such is available. For example, if this code is 404, the scan has encountered too many NOT_FOUND responses.
        """
        return pulumi.get(self, "most_common_http_error_code")

    @most_common_http_error_code.setter
    def most_common_http_error_code(self, value: Optional[pulumi.Input[int]]):
        pulumi.set(self, "most_common_http_error_code", value)

    @property
    @pulumi.getter(name="scanConfigError")
    def scan_config_error(self) -> Optional[pulumi.Input['ScanConfigErrorArgs']]:
        """
        If the scan encounters SCAN_CONFIG_ISSUE error, this field has the error message encountered during scan configuration validation that is performed before each scan run.
        """
        return pulumi.get(self, "scan_config_error")

    @scan_config_error.setter
    def scan_config_error(self, value: Optional[pulumi.Input['ScanConfigErrorArgs']]):
        pulumi.set(self, "scan_config_error", value)


@pulumi.input_type
class ScanRunWarningTraceArgs:
    def __init__(__self__, *,
                 code: Optional[pulumi.Input['ScanRunWarningTraceCode']] = None):
        """
        Output only. Defines a warning trace message for ScanRun. Warning traces provide customers with useful information that helps make the scanning process more effective.
        :param pulumi.Input['ScanRunWarningTraceCode'] code: Indicates the warning code.
        """
        if code is not None:
            pulumi.set(__self__, "code", code)

    @property
    @pulumi.getter
    def code(self) -> Optional[pulumi.Input['ScanRunWarningTraceCode']]:
        """
        Indicates the warning code.
        """
        return pulumi.get(self, "code")

    @code.setter
    def code(self, value: Optional[pulumi.Input['ScanRunWarningTraceCode']]):
        pulumi.set(self, "code", value)


@pulumi.input_type
class ScanRunArgs:
    def __init__(__self__, *,
                 end_time: Optional[pulumi.Input[str]] = None,
                 error_trace: Optional[pulumi.Input['ScanRunErrorTraceArgs']] = None,
                 execution_state: Optional[pulumi.Input['ScanRunExecutionState']] = None,
                 has_vulnerabilities: Optional[pulumi.Input[bool]] = None,
                 name: Optional[pulumi.Input[str]] = None,
                 progress_percent: Optional[pulumi.Input[int]] = None,
                 result_state: Optional[pulumi.Input['ScanRunResultState']] = None,
                 start_time: Optional[pulumi.Input[str]] = None,
                 urls_crawled_count: Optional[pulumi.Input[str]] = None,
                 urls_tested_count: Optional[pulumi.Input[str]] = None,
                 warning_traces: Optional[pulumi.Input[Sequence[pulumi.Input['ScanRunWarningTraceArgs']]]] = None):
        """
        A ScanRun is a output-only resource representing an actual run of the scan. Next id: 12
        :param pulumi.Input[str] end_time: The time at which the ScanRun reached termination state - that the ScanRun is either finished or stopped by user.
        :param pulumi.Input['ScanRunErrorTraceArgs'] error_trace: If result_state is an ERROR, this field provides the primary reason for scan's termination and more details, if such are available.
        :param pulumi.Input['ScanRunExecutionState'] execution_state: The execution state of the ScanRun.
        :param pulumi.Input[bool] has_vulnerabilities: Whether the scan run has found any vulnerabilities.
        :param pulumi.Input[str] name: The resource name of the ScanRun. The name follows the format of 'projects/{projectId}/scanConfigs/{scanConfigId}/scanRuns/{scanRunId}'. The ScanRun IDs are generated by the system.
        :param pulumi.Input[int] progress_percent: The percentage of total completion ranging from 0 to 100. If the scan is in queue, the value is 0. If the scan is running, the value ranges from 0 to 100. If the scan is finished, the value is 100.
        :param pulumi.Input['ScanRunResultState'] result_state: The result state of the ScanRun. This field is only available after the execution state reaches "FINISHED".
        :param pulumi.Input[str] start_time: The time at which the ScanRun started.
        :param pulumi.Input[str] urls_crawled_count: The number of URLs crawled during this ScanRun. If the scan is in progress, the value represents the number of URLs crawled up to now.
        :param pulumi.Input[str] urls_tested_count: The number of URLs tested during this ScanRun. If the scan is in progress, the value represents the number of URLs tested up to now. The number of URLs tested is usually larger than the number URLS crawled because typically a crawled URL is tested with multiple test payloads.
        :param pulumi.Input[Sequence[pulumi.Input['ScanRunWarningTraceArgs']]] warning_traces: A list of warnings, if such are encountered during this scan run.
        """
        if end_time is not None:
            pulumi.set(__self__, "end_time", end_time)
        if error_trace is not None:
            pulumi.set(__self__, "error_trace", error_trace)
        if execution_state is not None:
            pulumi.set(__self__, "execution_state", execution_state)
        if has_vulnerabilities is not None:
            pulumi.set(__self__, "has_vulnerabilities", has_vulnerabilities)
        if name is not None:
            pulumi.set(__self__, "name", name)
        if progress_percent is not None:
            pulumi.set(__self__, "progress_percent", progress_percent)
        if result_state is not None:
            pulumi.set(__self__, "result_state", result_state)
        if start_time is not None:
            pulumi.set(__self__, "start_time", start_time)
        if urls_crawled_count is not None:
            pulumi.set(__self__, "urls_crawled_count", urls_crawled_count)
        if urls_tested_count is not None:
            pulumi.set(__self__, "urls_tested_count", urls_tested_count)
        if warning_traces is not None:
            pulumi.set(__self__, "warning_traces", warning_traces)

    @property
    @pulumi.getter(name="endTime")
    def end_time(self) -> Optional[pulumi.Input[str]]:
        """
        The time at which the ScanRun reached termination state - that the ScanRun is either finished or stopped by user.
        """
        return pulumi.get(self, "end_time")

    @end_time.setter
    def end_time(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "end_time", value)

    @property
    @pulumi.getter(name="errorTrace")
    def error_trace(self) -> Optional[pulumi.Input['ScanRunErrorTraceArgs']]:
        """
        If result_state is an ERROR, this field provides the primary reason for scan's termination and more details, if such are available.
        """
        return pulumi.get(self, "error_trace")

    @error_trace.setter
    def error_trace(self, value: Optional[pulumi.Input['ScanRunErrorTraceArgs']]):
        pulumi.set(self, "error_trace", value)

    @property
    @pulumi.getter(name="executionState")
    def execution_state(self) -> Optional[pulumi.Input['ScanRunExecutionState']]:
        """
        The execution state of the ScanRun.
        """
        return pulumi.get(self, "execution_state")

    @execution_state.setter
    def execution_state(self, value: Optional[pulumi.Input['ScanRunExecutionState']]):
        pulumi.set(self, "execution_state", value)

    @property
    @pulumi.getter(name="hasVulnerabilities")
    def has_vulnerabilities(self) -> Optional[pulumi.Input[bool]]:
        """
        Whether the scan run has found any vulnerabilities.
        """
        return pulumi.get(self, "has_vulnerabilities")

    @has_vulnerabilities.setter
    def has_vulnerabilities(self, value: Optional[pulumi.Input[bool]]):
        pulumi.set(self, "has_vulnerabilities", value)

    @property
    @pulumi.getter
    def name(self) -> Optional[pulumi.Input[str]]:
        """
        The resource name of the ScanRun. The name follows the format of 'projects/{projectId}/scanConfigs/{scanConfigId}/scanRuns/{scanRunId}'. The ScanRun IDs are generated by the system.
        """
        return pulumi.get(self, "name")

    @name.setter
    def name(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "name", value)

    @property
    @pulumi.getter(name="progressPercent")
    def progress_percent(self) -> Optional[pulumi.Input[int]]:
        """
        The percentage of total completion ranging from 0 to 100. If the scan is in queue, the value is 0. If the scan is running, the value ranges from 0 to 100. If the scan is finished, the value is 100.
        """
        return pulumi.get(self, "progress_percent")

    @progress_percent.setter
    def progress_percent(self, value: Optional[pulumi.Input[int]]):
        pulumi.set(self, "progress_percent", value)

    @property
    @pulumi.getter(name="resultState")
    def result_state(self) -> Optional[pulumi.Input['ScanRunResultState']]:
        """
        The result state of the ScanRun. This field is only available after the execution state reaches "FINISHED".
        """
        return pulumi.get(self, "result_state")

    @result_state.setter
    def result_state(self, value: Optional[pulumi.Input['ScanRunResultState']]):
        pulumi.set(self, "result_state", value)

    @property
    @pulumi.getter(name="startTime")
    def start_time(self) -> Optional[pulumi.Input[str]]:
        """
        The time at which the ScanRun started.
        """
        return pulumi.get(self, "start_time")

    @start_time.setter
    def start_time(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "start_time", value)

    @property
    @pulumi.getter(name="urlsCrawledCount")
    def urls_crawled_count(self) -> Optional[pulumi.Input[str]]:
        """
        The number of URLs crawled during this ScanRun. If the scan is in progress, the value represents the number of URLs crawled up to now.
        """
        return pulumi.get(self, "urls_crawled_count")

    @urls_crawled_count.setter
    def urls_crawled_count(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "urls_crawled_count", value)

    @property
    @pulumi.getter(name="urlsTestedCount")
    def urls_tested_count(self) -> Optional[pulumi.Input[str]]:
        """
        The number of URLs tested during this ScanRun. If the scan is in progress, the value represents the number of URLs tested up to now. The number of URLs tested is usually larger than the number URLS crawled because typically a crawled URL is tested with multiple test payloads.
        """
        return pulumi.get(self, "urls_tested_count")

    @urls_tested_count.setter
    def urls_tested_count(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "urls_tested_count", value)

    @property
    @pulumi.getter(name="warningTraces")
    def warning_traces(self) -> Optional[pulumi.Input[Sequence[pulumi.Input['ScanRunWarningTraceArgs']]]]:
        """
        A list of warnings, if such are encountered during this scan run.
        """
        return pulumi.get(self, "warning_traces")

    @warning_traces.setter
    def warning_traces(self, value: Optional[pulumi.Input[Sequence[pulumi.Input['ScanRunWarningTraceArgs']]]]):
        pulumi.set(self, "warning_traces", value)


@pulumi.input_type
class ScheduleArgs:
    def __init__(__self__, *,
                 interval_duration_days: pulumi.Input[int],
                 schedule_time: Optional[pulumi.Input[str]] = None):
        """
        Scan schedule configuration.
        :param pulumi.Input[int] interval_duration_days: The duration of time between executions in days.
        :param pulumi.Input[str] schedule_time: A timestamp indicates when the next run will be scheduled. The value is refreshed by the server after each run. If unspecified, it will default to current server time, which means the scan will be scheduled to start immediately.
        """
        pulumi.set(__self__, "interval_duration_days", interval_duration_days)
        if schedule_time is not None:
            pulumi.set(__self__, "schedule_time", schedule_time)

    @property
    @pulumi.getter(name="intervalDurationDays")
    def interval_duration_days(self) -> pulumi.Input[int]:
        """
        The duration of time between executions in days.
        """
        return pulumi.get(self, "interval_duration_days")

    @interval_duration_days.setter
    def interval_duration_days(self, value: pulumi.Input[int]):
        pulumi.set(self, "interval_duration_days", value)

    @property
    @pulumi.getter(name="scheduleTime")
    def schedule_time(self) -> Optional[pulumi.Input[str]]:
        """
        A timestamp indicates when the next run will be scheduled. The value is refreshed by the server after each run. If unspecified, it will default to current server time, which means the scan will be scheduled to start immediately.
        """
        return pulumi.get(self, "schedule_time")

    @schedule_time.setter
    def schedule_time(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "schedule_time", value)


