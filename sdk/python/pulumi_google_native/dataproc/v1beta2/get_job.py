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
    'GetJobResult',
    'AwaitableGetJobResult',
    'get_job',
    'get_job_output',
]

@pulumi.output_type
class GetJobResult:
    def __init__(__self__, done=None, driver_control_files_uri=None, driver_output_resource_uri=None, hadoop_job=None, hive_job=None, job_uuid=None, labels=None, pig_job=None, placement=None, presto_job=None, pyspark_job=None, reference=None, scheduling=None, spark_job=None, spark_r_job=None, spark_sql_job=None, status=None, status_history=None, submitted_by=None, yarn_applications=None):
        if done and not isinstance(done, bool):
            raise TypeError("Expected argument 'done' to be a bool")
        pulumi.set(__self__, "done", done)
        if driver_control_files_uri and not isinstance(driver_control_files_uri, str):
            raise TypeError("Expected argument 'driver_control_files_uri' to be a str")
        pulumi.set(__self__, "driver_control_files_uri", driver_control_files_uri)
        if driver_output_resource_uri and not isinstance(driver_output_resource_uri, str):
            raise TypeError("Expected argument 'driver_output_resource_uri' to be a str")
        pulumi.set(__self__, "driver_output_resource_uri", driver_output_resource_uri)
        if hadoop_job and not isinstance(hadoop_job, dict):
            raise TypeError("Expected argument 'hadoop_job' to be a dict")
        pulumi.set(__self__, "hadoop_job", hadoop_job)
        if hive_job and not isinstance(hive_job, dict):
            raise TypeError("Expected argument 'hive_job' to be a dict")
        pulumi.set(__self__, "hive_job", hive_job)
        if job_uuid and not isinstance(job_uuid, str):
            raise TypeError("Expected argument 'job_uuid' to be a str")
        pulumi.set(__self__, "job_uuid", job_uuid)
        if labels and not isinstance(labels, dict):
            raise TypeError("Expected argument 'labels' to be a dict")
        pulumi.set(__self__, "labels", labels)
        if pig_job and not isinstance(pig_job, dict):
            raise TypeError("Expected argument 'pig_job' to be a dict")
        pulumi.set(__self__, "pig_job", pig_job)
        if placement and not isinstance(placement, dict):
            raise TypeError("Expected argument 'placement' to be a dict")
        pulumi.set(__self__, "placement", placement)
        if presto_job and not isinstance(presto_job, dict):
            raise TypeError("Expected argument 'presto_job' to be a dict")
        pulumi.set(__self__, "presto_job", presto_job)
        if pyspark_job and not isinstance(pyspark_job, dict):
            raise TypeError("Expected argument 'pyspark_job' to be a dict")
        pulumi.set(__self__, "pyspark_job", pyspark_job)
        if reference and not isinstance(reference, dict):
            raise TypeError("Expected argument 'reference' to be a dict")
        pulumi.set(__self__, "reference", reference)
        if scheduling and not isinstance(scheduling, dict):
            raise TypeError("Expected argument 'scheduling' to be a dict")
        pulumi.set(__self__, "scheduling", scheduling)
        if spark_job and not isinstance(spark_job, dict):
            raise TypeError("Expected argument 'spark_job' to be a dict")
        pulumi.set(__self__, "spark_job", spark_job)
        if spark_r_job and not isinstance(spark_r_job, dict):
            raise TypeError("Expected argument 'spark_r_job' to be a dict")
        pulumi.set(__self__, "spark_r_job", spark_r_job)
        if spark_sql_job and not isinstance(spark_sql_job, dict):
            raise TypeError("Expected argument 'spark_sql_job' to be a dict")
        pulumi.set(__self__, "spark_sql_job", spark_sql_job)
        if status and not isinstance(status, dict):
            raise TypeError("Expected argument 'status' to be a dict")
        pulumi.set(__self__, "status", status)
        if status_history and not isinstance(status_history, list):
            raise TypeError("Expected argument 'status_history' to be a list")
        pulumi.set(__self__, "status_history", status_history)
        if submitted_by and not isinstance(submitted_by, str):
            raise TypeError("Expected argument 'submitted_by' to be a str")
        pulumi.set(__self__, "submitted_by", submitted_by)
        if yarn_applications and not isinstance(yarn_applications, list):
            raise TypeError("Expected argument 'yarn_applications' to be a list")
        pulumi.set(__self__, "yarn_applications", yarn_applications)

    @property
    @pulumi.getter
    def done(self) -> bool:
        """
        Indicates whether the job is completed. If the value is false, the job is still in progress. If true, the job is completed, and status.state field will indicate if it was successful, failed, or cancelled.
        """
        return pulumi.get(self, "done")

    @property
    @pulumi.getter(name="driverControlFilesUri")
    def driver_control_files_uri(self) -> str:
        """
        If present, the location of miscellaneous control files which may be used as part of job setup and handling. If not present, control files may be placed in the same location as driver_output_uri.
        """
        return pulumi.get(self, "driver_control_files_uri")

    @property
    @pulumi.getter(name="driverOutputResourceUri")
    def driver_output_resource_uri(self) -> str:
        """
        A URI pointing to the location of the stdout of the job's driver program.
        """
        return pulumi.get(self, "driver_output_resource_uri")

    @property
    @pulumi.getter(name="hadoopJob")
    def hadoop_job(self) -> 'outputs.HadoopJobResponse':
        """
        Optional. Job is a Hadoop job.
        """
        return pulumi.get(self, "hadoop_job")

    @property
    @pulumi.getter(name="hiveJob")
    def hive_job(self) -> 'outputs.HiveJobResponse':
        """
        Optional. Job is a Hive job.
        """
        return pulumi.get(self, "hive_job")

    @property
    @pulumi.getter(name="jobUuid")
    def job_uuid(self) -> str:
        """
        A UUID that uniquely identifies a job within the project over time. This is in contrast to a user-settable reference.job_id that may be reused over time.
        """
        return pulumi.get(self, "job_uuid")

    @property
    @pulumi.getter
    def labels(self) -> Mapping[str, str]:
        """
        Optional. The labels to associate with this job. Label keys must contain 1 to 63 characters, and must conform to RFC 1035 (https://www.ietf.org/rfc/rfc1035.txt). Label values may be empty, but, if present, must contain 1 to 63 characters, and must conform to RFC 1035 (https://www.ietf.org/rfc/rfc1035.txt). No more than 32 labels can be associated with a job.
        """
        return pulumi.get(self, "labels")

    @property
    @pulumi.getter(name="pigJob")
    def pig_job(self) -> 'outputs.PigJobResponse':
        """
        Optional. Job is a Pig job.
        """
        return pulumi.get(self, "pig_job")

    @property
    @pulumi.getter
    def placement(self) -> 'outputs.JobPlacementResponse':
        """
        Job information, including how, when, and where to run the job.
        """
        return pulumi.get(self, "placement")

    @property
    @pulumi.getter(name="prestoJob")
    def presto_job(self) -> 'outputs.PrestoJobResponse':
        """
        Optional. Job is a Presto job.
        """
        return pulumi.get(self, "presto_job")

    @property
    @pulumi.getter(name="pysparkJob")
    def pyspark_job(self) -> 'outputs.PySparkJobResponse':
        """
        Optional. Job is a PySpark job.
        """
        return pulumi.get(self, "pyspark_job")

    @property
    @pulumi.getter
    def reference(self) -> 'outputs.JobReferenceResponse':
        """
        Optional. The fully qualified reference to the job, which can be used to obtain the equivalent REST path of the job resource. If this property is not specified when a job is created, the server generates a job_id.
        """
        return pulumi.get(self, "reference")

    @property
    @pulumi.getter
    def scheduling(self) -> 'outputs.JobSchedulingResponse':
        """
        Optional. Job scheduling configuration.
        """
        return pulumi.get(self, "scheduling")

    @property
    @pulumi.getter(name="sparkJob")
    def spark_job(self) -> 'outputs.SparkJobResponse':
        """
        Optional. Job is a Spark job.
        """
        return pulumi.get(self, "spark_job")

    @property
    @pulumi.getter(name="sparkRJob")
    def spark_r_job(self) -> 'outputs.SparkRJobResponse':
        """
        Optional. Job is a SparkR job.
        """
        return pulumi.get(self, "spark_r_job")

    @property
    @pulumi.getter(name="sparkSqlJob")
    def spark_sql_job(self) -> 'outputs.SparkSqlJobResponse':
        """
        Optional. Job is a SparkSql job.
        """
        return pulumi.get(self, "spark_sql_job")

    @property
    @pulumi.getter
    def status(self) -> 'outputs.JobStatusResponse':
        """
        The job status. Additional application-specific status information may be contained in the type_job and yarn_applications fields.
        """
        return pulumi.get(self, "status")

    @property
    @pulumi.getter(name="statusHistory")
    def status_history(self) -> Sequence['outputs.JobStatusResponse']:
        """
        The previous job status.
        """
        return pulumi.get(self, "status_history")

    @property
    @pulumi.getter(name="submittedBy")
    def submitted_by(self) -> str:
        """
        The email address of the user submitting the job. For jobs submitted on the cluster, the address is username@hostname.
        """
        return pulumi.get(self, "submitted_by")

    @property
    @pulumi.getter(name="yarnApplications")
    def yarn_applications(self) -> Sequence['outputs.YarnApplicationResponse']:
        """
        The collection of YARN applications spun up by this job.Beta Feature: This report is available for testing purposes only. It may be changed before final release.
        """
        return pulumi.get(self, "yarn_applications")


class AwaitableGetJobResult(GetJobResult):
    # pylint: disable=using-constant-test
    def __await__(self):
        if False:
            yield self
        return GetJobResult(
            done=self.done,
            driver_control_files_uri=self.driver_control_files_uri,
            driver_output_resource_uri=self.driver_output_resource_uri,
            hadoop_job=self.hadoop_job,
            hive_job=self.hive_job,
            job_uuid=self.job_uuid,
            labels=self.labels,
            pig_job=self.pig_job,
            placement=self.placement,
            presto_job=self.presto_job,
            pyspark_job=self.pyspark_job,
            reference=self.reference,
            scheduling=self.scheduling,
            spark_job=self.spark_job,
            spark_r_job=self.spark_r_job,
            spark_sql_job=self.spark_sql_job,
            status=self.status,
            status_history=self.status_history,
            submitted_by=self.submitted_by,
            yarn_applications=self.yarn_applications)


def get_job(job_id: Optional[str] = None,
            project: Optional[str] = None,
            region: Optional[str] = None,
            opts: Optional[pulumi.InvokeOptions] = None) -> AwaitableGetJobResult:
    """
    Gets the resource representation for a job in a project.
    """
    __args__ = dict()
    __args__['jobId'] = job_id
    __args__['project'] = project
    __args__['region'] = region
    if opts is None:
        opts = pulumi.InvokeOptions()
    if opts.version is None:
        opts.version = _utilities.get_version()
    __ret__ = pulumi.runtime.invoke('google-native:dataproc/v1beta2:getJob', __args__, opts=opts, typ=GetJobResult).value

    return AwaitableGetJobResult(
        done=__ret__.done,
        driver_control_files_uri=__ret__.driver_control_files_uri,
        driver_output_resource_uri=__ret__.driver_output_resource_uri,
        hadoop_job=__ret__.hadoop_job,
        hive_job=__ret__.hive_job,
        job_uuid=__ret__.job_uuid,
        labels=__ret__.labels,
        pig_job=__ret__.pig_job,
        placement=__ret__.placement,
        presto_job=__ret__.presto_job,
        pyspark_job=__ret__.pyspark_job,
        reference=__ret__.reference,
        scheduling=__ret__.scheduling,
        spark_job=__ret__.spark_job,
        spark_r_job=__ret__.spark_r_job,
        spark_sql_job=__ret__.spark_sql_job,
        status=__ret__.status,
        status_history=__ret__.status_history,
        submitted_by=__ret__.submitted_by,
        yarn_applications=__ret__.yarn_applications)


@_utilities.lift_output_func(get_job)
def get_job_output(job_id: Optional[pulumi.Input[str]] = None,
                   project: Optional[pulumi.Input[Optional[str]]] = None,
                   region: Optional[pulumi.Input[str]] = None,
                   opts: Optional[pulumi.InvokeOptions] = None) -> pulumi.Output[GetJobResult]:
    """
    Gets the resource representation for a job in a project.
    """
    ...
