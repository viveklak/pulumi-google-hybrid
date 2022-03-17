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
    'MaintenancePolicyArgs',
    'PersistenceConfigArgs',
    'TimeOfDayArgs',
    'WeeklyMaintenanceWindowArgs',
]

@pulumi.input_type
class MaintenancePolicyArgs:
    def __init__(__self__, *,
                 description: Optional[pulumi.Input[str]] = None,
                 weekly_maintenance_window: Optional[pulumi.Input[Sequence[pulumi.Input['WeeklyMaintenanceWindowArgs']]]] = None):
        """
        Maintenance policy for an instance.
        :param pulumi.Input[str] description: Optional. Description of what this policy is for. Create/Update methods return INVALID_ARGUMENT if the length is greater than 512.
        :param pulumi.Input[Sequence[pulumi.Input['WeeklyMaintenanceWindowArgs']]] weekly_maintenance_window: Optional. Maintenance window that is applied to resources covered by this policy. Minimum 1. For the current version, the maximum number of weekly_window is expected to be one.
        """
        if description is not None:
            pulumi.set(__self__, "description", description)
        if weekly_maintenance_window is not None:
            pulumi.set(__self__, "weekly_maintenance_window", weekly_maintenance_window)

    @property
    @pulumi.getter
    def description(self) -> Optional[pulumi.Input[str]]:
        """
        Optional. Description of what this policy is for. Create/Update methods return INVALID_ARGUMENT if the length is greater than 512.
        """
        return pulumi.get(self, "description")

    @description.setter
    def description(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "description", value)

    @property
    @pulumi.getter(name="weeklyMaintenanceWindow")
    def weekly_maintenance_window(self) -> Optional[pulumi.Input[Sequence[pulumi.Input['WeeklyMaintenanceWindowArgs']]]]:
        """
        Optional. Maintenance window that is applied to resources covered by this policy. Minimum 1. For the current version, the maximum number of weekly_window is expected to be one.
        """
        return pulumi.get(self, "weekly_maintenance_window")

    @weekly_maintenance_window.setter
    def weekly_maintenance_window(self, value: Optional[pulumi.Input[Sequence[pulumi.Input['WeeklyMaintenanceWindowArgs']]]]):
        pulumi.set(self, "weekly_maintenance_window", value)


@pulumi.input_type
class PersistenceConfigArgs:
    def __init__(__self__, *,
                 persistence_mode: Optional[pulumi.Input['PersistenceConfigPersistenceMode']] = None,
                 rdb_snapshot_period: Optional[pulumi.Input['PersistenceConfigRdbSnapshotPeriod']] = None,
                 rdb_snapshot_start_time: Optional[pulumi.Input[str]] = None):
        """
        Configuration of the persistence functionality.
        :param pulumi.Input['PersistenceConfigPersistenceMode'] persistence_mode: Optional. Controls whether Persistence features are enabled. If not provided, the existing value will be used.
        :param pulumi.Input['PersistenceConfigRdbSnapshotPeriod'] rdb_snapshot_period: Optional. Period between RDB snapshots. Snapshots will be attempted every period starting from the provided snapshot start time. For example, a start time of 01/01/2033 06:45 and SIX_HOURS snapshot period will do nothing until 01/01/2033, and then trigger snapshots every day at 06:45, 12:45, 18:45, and 00:45 the next day, and so on. If not provided, TWENTY_FOUR_HOURS will be used as default.
        :param pulumi.Input[str] rdb_snapshot_start_time: Optional. Date and time that the first snapshot was/will be attempted, and to which future snapshots will be aligned. If not provided, the current time will be used.
        """
        if persistence_mode is not None:
            pulumi.set(__self__, "persistence_mode", persistence_mode)
        if rdb_snapshot_period is not None:
            pulumi.set(__self__, "rdb_snapshot_period", rdb_snapshot_period)
        if rdb_snapshot_start_time is not None:
            pulumi.set(__self__, "rdb_snapshot_start_time", rdb_snapshot_start_time)

    @property
    @pulumi.getter(name="persistenceMode")
    def persistence_mode(self) -> Optional[pulumi.Input['PersistenceConfigPersistenceMode']]:
        """
        Optional. Controls whether Persistence features are enabled. If not provided, the existing value will be used.
        """
        return pulumi.get(self, "persistence_mode")

    @persistence_mode.setter
    def persistence_mode(self, value: Optional[pulumi.Input['PersistenceConfigPersistenceMode']]):
        pulumi.set(self, "persistence_mode", value)

    @property
    @pulumi.getter(name="rdbSnapshotPeriod")
    def rdb_snapshot_period(self) -> Optional[pulumi.Input['PersistenceConfigRdbSnapshotPeriod']]:
        """
        Optional. Period between RDB snapshots. Snapshots will be attempted every period starting from the provided snapshot start time. For example, a start time of 01/01/2033 06:45 and SIX_HOURS snapshot period will do nothing until 01/01/2033, and then trigger snapshots every day at 06:45, 12:45, 18:45, and 00:45 the next day, and so on. If not provided, TWENTY_FOUR_HOURS will be used as default.
        """
        return pulumi.get(self, "rdb_snapshot_period")

    @rdb_snapshot_period.setter
    def rdb_snapshot_period(self, value: Optional[pulumi.Input['PersistenceConfigRdbSnapshotPeriod']]):
        pulumi.set(self, "rdb_snapshot_period", value)

    @property
    @pulumi.getter(name="rdbSnapshotStartTime")
    def rdb_snapshot_start_time(self) -> Optional[pulumi.Input[str]]:
        """
        Optional. Date and time that the first snapshot was/will be attempted, and to which future snapshots will be aligned. If not provided, the current time will be used.
        """
        return pulumi.get(self, "rdb_snapshot_start_time")

    @rdb_snapshot_start_time.setter
    def rdb_snapshot_start_time(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "rdb_snapshot_start_time", value)


@pulumi.input_type
class TimeOfDayArgs:
    def __init__(__self__, *,
                 hours: Optional[pulumi.Input[int]] = None,
                 minutes: Optional[pulumi.Input[int]] = None,
                 nanos: Optional[pulumi.Input[int]] = None,
                 seconds: Optional[pulumi.Input[int]] = None):
        """
        Represents a time of day. The date and time zone are either not significant or are specified elsewhere. An API may choose to allow leap seconds. Related types are google.type.Date and `google.protobuf.Timestamp`.
        :param pulumi.Input[int] hours: Hours of day in 24 hour format. Should be from 0 to 23. An API may choose to allow the value "24:00:00" for scenarios like business closing time.
        :param pulumi.Input[int] minutes: Minutes of hour of day. Must be from 0 to 59.
        :param pulumi.Input[int] nanos: Fractions of seconds in nanoseconds. Must be from 0 to 999,999,999.
        :param pulumi.Input[int] seconds: Seconds of minutes of the time. Must normally be from 0 to 59. An API may allow the value 60 if it allows leap-seconds.
        """
        if hours is not None:
            pulumi.set(__self__, "hours", hours)
        if minutes is not None:
            pulumi.set(__self__, "minutes", minutes)
        if nanos is not None:
            pulumi.set(__self__, "nanos", nanos)
        if seconds is not None:
            pulumi.set(__self__, "seconds", seconds)

    @property
    @pulumi.getter
    def hours(self) -> Optional[pulumi.Input[int]]:
        """
        Hours of day in 24 hour format. Should be from 0 to 23. An API may choose to allow the value "24:00:00" for scenarios like business closing time.
        """
        return pulumi.get(self, "hours")

    @hours.setter
    def hours(self, value: Optional[pulumi.Input[int]]):
        pulumi.set(self, "hours", value)

    @property
    @pulumi.getter
    def minutes(self) -> Optional[pulumi.Input[int]]:
        """
        Minutes of hour of day. Must be from 0 to 59.
        """
        return pulumi.get(self, "minutes")

    @minutes.setter
    def minutes(self, value: Optional[pulumi.Input[int]]):
        pulumi.set(self, "minutes", value)

    @property
    @pulumi.getter
    def nanos(self) -> Optional[pulumi.Input[int]]:
        """
        Fractions of seconds in nanoseconds. Must be from 0 to 999,999,999.
        """
        return pulumi.get(self, "nanos")

    @nanos.setter
    def nanos(self, value: Optional[pulumi.Input[int]]):
        pulumi.set(self, "nanos", value)

    @property
    @pulumi.getter
    def seconds(self) -> Optional[pulumi.Input[int]]:
        """
        Seconds of minutes of the time. Must normally be from 0 to 59. An API may allow the value 60 if it allows leap-seconds.
        """
        return pulumi.get(self, "seconds")

    @seconds.setter
    def seconds(self, value: Optional[pulumi.Input[int]]):
        pulumi.set(self, "seconds", value)


@pulumi.input_type
class WeeklyMaintenanceWindowArgs:
    def __init__(__self__, *,
                 day: pulumi.Input['WeeklyMaintenanceWindowDay'],
                 start_time: pulumi.Input['TimeOfDayArgs']):
        """
        Time window in which disruptive maintenance updates occur. Non-disruptive updates can occur inside or outside this window.
        :param pulumi.Input['WeeklyMaintenanceWindowDay'] day: The day of week that maintenance updates occur.
        :param pulumi.Input['TimeOfDayArgs'] start_time: Start time of the window in UTC time.
        """
        pulumi.set(__self__, "day", day)
        pulumi.set(__self__, "start_time", start_time)

    @property
    @pulumi.getter
    def day(self) -> pulumi.Input['WeeklyMaintenanceWindowDay']:
        """
        The day of week that maintenance updates occur.
        """
        return pulumi.get(self, "day")

    @day.setter
    def day(self, value: pulumi.Input['WeeklyMaintenanceWindowDay']):
        pulumi.set(self, "day", value)

    @property
    @pulumi.getter(name="startTime")
    def start_time(self) -> pulumi.Input['TimeOfDayArgs']:
        """
        Start time of the window in UTC time.
        """
        return pulumi.get(self, "start_time")

    @start_time.setter
    def start_time(self, value: pulumi.Input['TimeOfDayArgs']):
        pulumi.set(self, "start_time", value)


