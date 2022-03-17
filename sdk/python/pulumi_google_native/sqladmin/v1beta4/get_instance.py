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
    'GetInstanceResult',
    'AwaitableGetInstanceResult',
    'get_instance',
    'get_instance_output',
]

@pulumi.output_type
class GetInstanceResult:
    def __init__(__self__, available_maintenance_versions=None, backend_type=None, connection_name=None, create_time=None, current_disk_size=None, database_installed_version=None, database_version=None, disk_encryption_configuration=None, disk_encryption_status=None, etag=None, failover_replica=None, gce_zone=None, instance_type=None, ip_addresses=None, ipv6_address=None, kind=None, maintenance_version=None, master_instance_name=None, max_disk_size=None, name=None, on_premises_configuration=None, out_of_disk_report=None, project=None, region=None, replica_configuration=None, replica_names=None, root_password=None, satisfies_pzs=None, scheduled_maintenance=None, secondary_gce_zone=None, self_link=None, server_ca_cert=None, service_account_email_address=None, settings=None, state=None, suspension_reason=None):
        if available_maintenance_versions and not isinstance(available_maintenance_versions, list):
            raise TypeError("Expected argument 'available_maintenance_versions' to be a list")
        pulumi.set(__self__, "available_maintenance_versions", available_maintenance_versions)
        if backend_type and not isinstance(backend_type, str):
            raise TypeError("Expected argument 'backend_type' to be a str")
        pulumi.set(__self__, "backend_type", backend_type)
        if connection_name and not isinstance(connection_name, str):
            raise TypeError("Expected argument 'connection_name' to be a str")
        pulumi.set(__self__, "connection_name", connection_name)
        if create_time and not isinstance(create_time, str):
            raise TypeError("Expected argument 'create_time' to be a str")
        pulumi.set(__self__, "create_time", create_time)
        if current_disk_size and not isinstance(current_disk_size, str):
            raise TypeError("Expected argument 'current_disk_size' to be a str")
        pulumi.set(__self__, "current_disk_size", current_disk_size)
        if database_installed_version and not isinstance(database_installed_version, str):
            raise TypeError("Expected argument 'database_installed_version' to be a str")
        pulumi.set(__self__, "database_installed_version", database_installed_version)
        if database_version and not isinstance(database_version, str):
            raise TypeError("Expected argument 'database_version' to be a str")
        pulumi.set(__self__, "database_version", database_version)
        if disk_encryption_configuration and not isinstance(disk_encryption_configuration, dict):
            raise TypeError("Expected argument 'disk_encryption_configuration' to be a dict")
        pulumi.set(__self__, "disk_encryption_configuration", disk_encryption_configuration)
        if disk_encryption_status and not isinstance(disk_encryption_status, dict):
            raise TypeError("Expected argument 'disk_encryption_status' to be a dict")
        pulumi.set(__self__, "disk_encryption_status", disk_encryption_status)
        if etag and not isinstance(etag, str):
            raise TypeError("Expected argument 'etag' to be a str")
        if etag is not None:
            warnings.warn("""This field is deprecated and will be removed from a future version of the API. Use the `settings.settingsVersion` field instead.""", DeprecationWarning)
            pulumi.log.warn("""etag is deprecated: This field is deprecated and will be removed from a future version of the API. Use the `settings.settingsVersion` field instead.""")

        pulumi.set(__self__, "etag", etag)
        if failover_replica and not isinstance(failover_replica, dict):
            raise TypeError("Expected argument 'failover_replica' to be a dict")
        pulumi.set(__self__, "failover_replica", failover_replica)
        if gce_zone and not isinstance(gce_zone, str):
            raise TypeError("Expected argument 'gce_zone' to be a str")
        pulumi.set(__self__, "gce_zone", gce_zone)
        if instance_type and not isinstance(instance_type, str):
            raise TypeError("Expected argument 'instance_type' to be a str")
        pulumi.set(__self__, "instance_type", instance_type)
        if ip_addresses and not isinstance(ip_addresses, list):
            raise TypeError("Expected argument 'ip_addresses' to be a list")
        pulumi.set(__self__, "ip_addresses", ip_addresses)
        if ipv6_address and not isinstance(ipv6_address, str):
            raise TypeError("Expected argument 'ipv6_address' to be a str")
        if ipv6_address is not None:
            warnings.warn("""The IPv6 address assigned to the instance. (Deprecated) This property was applicable only to First Generation instances.""", DeprecationWarning)
            pulumi.log.warn("""ipv6_address is deprecated: The IPv6 address assigned to the instance. (Deprecated) This property was applicable only to First Generation instances.""")

        pulumi.set(__self__, "ipv6_address", ipv6_address)
        if kind and not isinstance(kind, str):
            raise TypeError("Expected argument 'kind' to be a str")
        pulumi.set(__self__, "kind", kind)
        if maintenance_version and not isinstance(maintenance_version, str):
            raise TypeError("Expected argument 'maintenance_version' to be a str")
        pulumi.set(__self__, "maintenance_version", maintenance_version)
        if master_instance_name and not isinstance(master_instance_name, str):
            raise TypeError("Expected argument 'master_instance_name' to be a str")
        pulumi.set(__self__, "master_instance_name", master_instance_name)
        if max_disk_size and not isinstance(max_disk_size, str):
            raise TypeError("Expected argument 'max_disk_size' to be a str")
        pulumi.set(__self__, "max_disk_size", max_disk_size)
        if name and not isinstance(name, str):
            raise TypeError("Expected argument 'name' to be a str")
        pulumi.set(__self__, "name", name)
        if on_premises_configuration and not isinstance(on_premises_configuration, dict):
            raise TypeError("Expected argument 'on_premises_configuration' to be a dict")
        pulumi.set(__self__, "on_premises_configuration", on_premises_configuration)
        if out_of_disk_report and not isinstance(out_of_disk_report, dict):
            raise TypeError("Expected argument 'out_of_disk_report' to be a dict")
        pulumi.set(__self__, "out_of_disk_report", out_of_disk_report)
        if project and not isinstance(project, str):
            raise TypeError("Expected argument 'project' to be a str")
        pulumi.set(__self__, "project", project)
        if region and not isinstance(region, str):
            raise TypeError("Expected argument 'region' to be a str")
        pulumi.set(__self__, "region", region)
        if replica_configuration and not isinstance(replica_configuration, dict):
            raise TypeError("Expected argument 'replica_configuration' to be a dict")
        pulumi.set(__self__, "replica_configuration", replica_configuration)
        if replica_names and not isinstance(replica_names, list):
            raise TypeError("Expected argument 'replica_names' to be a list")
        pulumi.set(__self__, "replica_names", replica_names)
        if root_password and not isinstance(root_password, str):
            raise TypeError("Expected argument 'root_password' to be a str")
        pulumi.set(__self__, "root_password", root_password)
        if satisfies_pzs and not isinstance(satisfies_pzs, bool):
            raise TypeError("Expected argument 'satisfies_pzs' to be a bool")
        pulumi.set(__self__, "satisfies_pzs", satisfies_pzs)
        if scheduled_maintenance and not isinstance(scheduled_maintenance, dict):
            raise TypeError("Expected argument 'scheduled_maintenance' to be a dict")
        pulumi.set(__self__, "scheduled_maintenance", scheduled_maintenance)
        if secondary_gce_zone and not isinstance(secondary_gce_zone, str):
            raise TypeError("Expected argument 'secondary_gce_zone' to be a str")
        pulumi.set(__self__, "secondary_gce_zone", secondary_gce_zone)
        if self_link and not isinstance(self_link, str):
            raise TypeError("Expected argument 'self_link' to be a str")
        pulumi.set(__self__, "self_link", self_link)
        if server_ca_cert and not isinstance(server_ca_cert, dict):
            raise TypeError("Expected argument 'server_ca_cert' to be a dict")
        pulumi.set(__self__, "server_ca_cert", server_ca_cert)
        if service_account_email_address and not isinstance(service_account_email_address, str):
            raise TypeError("Expected argument 'service_account_email_address' to be a str")
        pulumi.set(__self__, "service_account_email_address", service_account_email_address)
        if settings and not isinstance(settings, dict):
            raise TypeError("Expected argument 'settings' to be a dict")
        pulumi.set(__self__, "settings", settings)
        if state and not isinstance(state, str):
            raise TypeError("Expected argument 'state' to be a str")
        pulumi.set(__self__, "state", state)
        if suspension_reason and not isinstance(suspension_reason, list):
            raise TypeError("Expected argument 'suspension_reason' to be a list")
        pulumi.set(__self__, "suspension_reason", suspension_reason)

    @property
    @pulumi.getter(name="availableMaintenanceVersions")
    def available_maintenance_versions(self) -> Sequence[str]:
        """
        List all maintenance versions applicable on the instance
        """
        return pulumi.get(self, "available_maintenance_versions")

    @property
    @pulumi.getter(name="backendType")
    def backend_type(self) -> str:
        """
        The backend type. `SECOND_GEN`: Cloud SQL database instance. `EXTERNAL`: A database server that is not managed by Google. This property is read-only; use the `tier` property in the `settings` object to determine the database type.
        """
        return pulumi.get(self, "backend_type")

    @property
    @pulumi.getter(name="connectionName")
    def connection_name(self) -> str:
        """
        Connection name of the Cloud SQL instance used in connection strings.
        """
        return pulumi.get(self, "connection_name")

    @property
    @pulumi.getter(name="createTime")
    def create_time(self) -> str:
        """
        The time when the instance was created in [RFC 3339](https://tools.ietf.org/html/rfc3339) format, for example `2012-11-15T16:19:00.094Z`.
        """
        return pulumi.get(self, "create_time")

    @property
    @pulumi.getter(name="currentDiskSize")
    def current_disk_size(self) -> str:
        """
        The current disk usage of the instance in bytes. This property has been deprecated. Use the "cloudsql.googleapis.com/database/disk/bytes_used" metric in Cloud Monitoring API instead. Please see [this announcement](https://groups.google.com/d/msg/google-cloud-sql-announce/I_7-F9EBhT0/BtvFtdFeAgAJ) for details.
        """
        return pulumi.get(self, "current_disk_size")

    @property
    @pulumi.getter(name="databaseInstalledVersion")
    def database_installed_version(self) -> str:
        """
        Stores the current database version running on the instance including minor version such as `MYSQL_8_0_18`.
        """
        return pulumi.get(self, "database_installed_version")

    @property
    @pulumi.getter(name="databaseVersion")
    def database_version(self) -> str:
        """
        The database engine type and version. The `databaseVersion` field cannot be changed after instance creation.
        """
        return pulumi.get(self, "database_version")

    @property
    @pulumi.getter(name="diskEncryptionConfiguration")
    def disk_encryption_configuration(self) -> 'outputs.DiskEncryptionConfigurationResponse':
        """
        Disk encryption configuration specific to an instance.
        """
        return pulumi.get(self, "disk_encryption_configuration")

    @property
    @pulumi.getter(name="diskEncryptionStatus")
    def disk_encryption_status(self) -> 'outputs.DiskEncryptionStatusResponse':
        """
        Disk encryption status specific to an instance.
        """
        return pulumi.get(self, "disk_encryption_status")

    @property
    @pulumi.getter
    def etag(self) -> str:
        """
        This field is deprecated and will be removed from a future version of the API. Use the `settings.settingsVersion` field instead.
        """
        return pulumi.get(self, "etag")

    @property
    @pulumi.getter(name="failoverReplica")
    def failover_replica(self) -> 'outputs.InstanceFailoverReplicaResponse':
        """
        The name and status of the failover replica.
        """
        return pulumi.get(self, "failover_replica")

    @property
    @pulumi.getter(name="gceZone")
    def gce_zone(self) -> str:
        """
        The Compute Engine zone that the instance is currently serving from. This value could be different from the zone that was specified when the instance was created if the instance has failed over to its secondary zone. WARNING: Changing this might restart the instance.
        """
        return pulumi.get(self, "gce_zone")

    @property
    @pulumi.getter(name="instanceType")
    def instance_type(self) -> str:
        """
        The instance type.
        """
        return pulumi.get(self, "instance_type")

    @property
    @pulumi.getter(name="ipAddresses")
    def ip_addresses(self) -> Sequence['outputs.IpMappingResponse']:
        """
        The assigned IP addresses for the instance.
        """
        return pulumi.get(self, "ip_addresses")

    @property
    @pulumi.getter(name="ipv6Address")
    def ipv6_address(self) -> str:
        """
        The IPv6 address assigned to the instance. (Deprecated) This property was applicable only to First Generation instances.
        """
        return pulumi.get(self, "ipv6_address")

    @property
    @pulumi.getter
    def kind(self) -> str:
        """
        This is always `sql#instance`.
        """
        return pulumi.get(self, "kind")

    @property
    @pulumi.getter(name="maintenanceVersion")
    def maintenance_version(self) -> str:
        """
        The current software version on the instance.
        """
        return pulumi.get(self, "maintenance_version")

    @property
    @pulumi.getter(name="masterInstanceName")
    def master_instance_name(self) -> str:
        """
        The name of the instance which will act as primary in the replication setup.
        """
        return pulumi.get(self, "master_instance_name")

    @property
    @pulumi.getter(name="maxDiskSize")
    def max_disk_size(self) -> str:
        """
        The maximum disk size of the instance in bytes.
        """
        return pulumi.get(self, "max_disk_size")

    @property
    @pulumi.getter
    def name(self) -> str:
        """
        Name of the Cloud SQL instance. This does not include the project ID.
        """
        return pulumi.get(self, "name")

    @property
    @pulumi.getter(name="onPremisesConfiguration")
    def on_premises_configuration(self) -> 'outputs.OnPremisesConfigurationResponse':
        """
        Configuration specific to on-premises instances.
        """
        return pulumi.get(self, "on_premises_configuration")

    @property
    @pulumi.getter(name="outOfDiskReport")
    def out_of_disk_report(self) -> 'outputs.SqlOutOfDiskReportResponse':
        """
        This field represents the report generated by the proactive database wellness job for OutOfDisk issues. * Writers: * the proactive database wellness job for OOD. * Readers: * the proactive database wellness job
        """
        return pulumi.get(self, "out_of_disk_report")

    @property
    @pulumi.getter
    def project(self) -> str:
        """
        The project ID of the project containing the Cloud SQL instance. The Google apps domain is prefixed if applicable.
        """
        return pulumi.get(self, "project")

    @property
    @pulumi.getter
    def region(self) -> str:
        """
        The geographical region. Can be: * `us-central` (`FIRST_GEN` instances only) * `us-central1` (`SECOND_GEN` instances only) * `asia-east1` or `europe-west1`. Defaults to `us-central` or `us-central1` depending on the instance type. The region cannot be changed after instance creation.
        """
        return pulumi.get(self, "region")

    @property
    @pulumi.getter(name="replicaConfiguration")
    def replica_configuration(self) -> 'outputs.ReplicaConfigurationResponse':
        """
        Configuration specific to failover replicas and read replicas.
        """
        return pulumi.get(self, "replica_configuration")

    @property
    @pulumi.getter(name="replicaNames")
    def replica_names(self) -> Sequence[str]:
        """
        The replicas of the instance.
        """
        return pulumi.get(self, "replica_names")

    @property
    @pulumi.getter(name="rootPassword")
    def root_password(self) -> str:
        """
        Initial root password. Use only on creation.
        """
        return pulumi.get(self, "root_password")

    @property
    @pulumi.getter(name="satisfiesPzs")
    def satisfies_pzs(self) -> bool:
        """
        The status indicating if instance satisfiesPzs. Reserved for future use.
        """
        return pulumi.get(self, "satisfies_pzs")

    @property
    @pulumi.getter(name="scheduledMaintenance")
    def scheduled_maintenance(self) -> 'outputs.SqlScheduledMaintenanceResponse':
        """
        The start time of any upcoming scheduled maintenance for this instance.
        """
        return pulumi.get(self, "scheduled_maintenance")

    @property
    @pulumi.getter(name="secondaryGceZone")
    def secondary_gce_zone(self) -> str:
        """
        The Compute Engine zone that the failover instance is currently serving from for a regional instance. This value could be different from the zone that was specified when the instance was created if the instance has failed over to its secondary/failover zone.
        """
        return pulumi.get(self, "secondary_gce_zone")

    @property
    @pulumi.getter(name="selfLink")
    def self_link(self) -> str:
        """
        The URI of this resource.
        """
        return pulumi.get(self, "self_link")

    @property
    @pulumi.getter(name="serverCaCert")
    def server_ca_cert(self) -> 'outputs.SslCertResponse':
        """
        SSL configuration.
        """
        return pulumi.get(self, "server_ca_cert")

    @property
    @pulumi.getter(name="serviceAccountEmailAddress")
    def service_account_email_address(self) -> str:
        """
        The service account email address assigned to the instance. \This property is read-only.
        """
        return pulumi.get(self, "service_account_email_address")

    @property
    @pulumi.getter
    def settings(self) -> 'outputs.SettingsResponse':
        """
        The user settings.
        """
        return pulumi.get(self, "settings")

    @property
    @pulumi.getter
    def state(self) -> str:
        """
        The current serving state of the Cloud SQL instance.
        """
        return pulumi.get(self, "state")

    @property
    @pulumi.getter(name="suspensionReason")
    def suspension_reason(self) -> Sequence[str]:
        """
        If the instance state is SUSPENDED, the reason for the suspension.
        """
        return pulumi.get(self, "suspension_reason")


class AwaitableGetInstanceResult(GetInstanceResult):
    # pylint: disable=using-constant-test
    def __await__(self):
        if False:
            yield self
        return GetInstanceResult(
            available_maintenance_versions=self.available_maintenance_versions,
            backend_type=self.backend_type,
            connection_name=self.connection_name,
            create_time=self.create_time,
            current_disk_size=self.current_disk_size,
            database_installed_version=self.database_installed_version,
            database_version=self.database_version,
            disk_encryption_configuration=self.disk_encryption_configuration,
            disk_encryption_status=self.disk_encryption_status,
            etag=self.etag,
            failover_replica=self.failover_replica,
            gce_zone=self.gce_zone,
            instance_type=self.instance_type,
            ip_addresses=self.ip_addresses,
            ipv6_address=self.ipv6_address,
            kind=self.kind,
            maintenance_version=self.maintenance_version,
            master_instance_name=self.master_instance_name,
            max_disk_size=self.max_disk_size,
            name=self.name,
            on_premises_configuration=self.on_premises_configuration,
            out_of_disk_report=self.out_of_disk_report,
            project=self.project,
            region=self.region,
            replica_configuration=self.replica_configuration,
            replica_names=self.replica_names,
            root_password=self.root_password,
            satisfies_pzs=self.satisfies_pzs,
            scheduled_maintenance=self.scheduled_maintenance,
            secondary_gce_zone=self.secondary_gce_zone,
            self_link=self.self_link,
            server_ca_cert=self.server_ca_cert,
            service_account_email_address=self.service_account_email_address,
            settings=self.settings,
            state=self.state,
            suspension_reason=self.suspension_reason)


def get_instance(instance: Optional[str] = None,
                 project: Optional[str] = None,
                 opts: Optional[pulumi.InvokeOptions] = None) -> AwaitableGetInstanceResult:
    """
    Retrieves a resource containing information about a Cloud SQL instance.
    """
    __args__ = dict()
    __args__['instance'] = instance
    __args__['project'] = project
    if opts is None:
        opts = pulumi.InvokeOptions()
    if opts.version is None:
        opts.version = _utilities.get_version()
    __ret__ = pulumi.runtime.invoke('google-native:sqladmin/v1beta4:getInstance', __args__, opts=opts, typ=GetInstanceResult).value

    return AwaitableGetInstanceResult(
        available_maintenance_versions=__ret__.available_maintenance_versions,
        backend_type=__ret__.backend_type,
        connection_name=__ret__.connection_name,
        create_time=__ret__.create_time,
        current_disk_size=__ret__.current_disk_size,
        database_installed_version=__ret__.database_installed_version,
        database_version=__ret__.database_version,
        disk_encryption_configuration=__ret__.disk_encryption_configuration,
        disk_encryption_status=__ret__.disk_encryption_status,
        etag=__ret__.etag,
        failover_replica=__ret__.failover_replica,
        gce_zone=__ret__.gce_zone,
        instance_type=__ret__.instance_type,
        ip_addresses=__ret__.ip_addresses,
        ipv6_address=__ret__.ipv6_address,
        kind=__ret__.kind,
        maintenance_version=__ret__.maintenance_version,
        master_instance_name=__ret__.master_instance_name,
        max_disk_size=__ret__.max_disk_size,
        name=__ret__.name,
        on_premises_configuration=__ret__.on_premises_configuration,
        out_of_disk_report=__ret__.out_of_disk_report,
        project=__ret__.project,
        region=__ret__.region,
        replica_configuration=__ret__.replica_configuration,
        replica_names=__ret__.replica_names,
        root_password=__ret__.root_password,
        satisfies_pzs=__ret__.satisfies_pzs,
        scheduled_maintenance=__ret__.scheduled_maintenance,
        secondary_gce_zone=__ret__.secondary_gce_zone,
        self_link=__ret__.self_link,
        server_ca_cert=__ret__.server_ca_cert,
        service_account_email_address=__ret__.service_account_email_address,
        settings=__ret__.settings,
        state=__ret__.state,
        suspension_reason=__ret__.suspension_reason)


@_utilities.lift_output_func(get_instance)
def get_instance_output(instance: Optional[pulumi.Input[str]] = None,
                        project: Optional[pulumi.Input[Optional[str]]] = None,
                        opts: Optional[pulumi.InvokeOptions] = None) -> pulumi.Output[GetInstanceResult]:
    """
    Retrieves a resource containing information about a Cloud SQL instance.
    """
    ...
