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
    'GetTriggerResult',
    'AwaitableGetTriggerResult',
    'get_trigger',
]

@pulumi.output_type
class GetTriggerResult:
    def __init__(__self__, create_time=None, destination=None, etag=None, labels=None, matching_criteria=None, name=None, service_account=None, transport=None, update_time=None):
        if create_time and not isinstance(create_time, str):
            raise TypeError("Expected argument 'create_time' to be a str")
        pulumi.set(__self__, "create_time", create_time)
        if destination and not isinstance(destination, dict):
            raise TypeError("Expected argument 'destination' to be a dict")
        pulumi.set(__self__, "destination", destination)
        if etag and not isinstance(etag, str):
            raise TypeError("Expected argument 'etag' to be a str")
        pulumi.set(__self__, "etag", etag)
        if labels and not isinstance(labels, dict):
            raise TypeError("Expected argument 'labels' to be a dict")
        pulumi.set(__self__, "labels", labels)
        if matching_criteria and not isinstance(matching_criteria, list):
            raise TypeError("Expected argument 'matching_criteria' to be a list")
        pulumi.set(__self__, "matching_criteria", matching_criteria)
        if name and not isinstance(name, str):
            raise TypeError("Expected argument 'name' to be a str")
        pulumi.set(__self__, "name", name)
        if service_account and not isinstance(service_account, str):
            raise TypeError("Expected argument 'service_account' to be a str")
        pulumi.set(__self__, "service_account", service_account)
        if transport and not isinstance(transport, dict):
            raise TypeError("Expected argument 'transport' to be a dict")
        pulumi.set(__self__, "transport", transport)
        if update_time and not isinstance(update_time, str):
            raise TypeError("Expected argument 'update_time' to be a str")
        pulumi.set(__self__, "update_time", update_time)

    @property
    @pulumi.getter(name="createTime")
    def create_time(self) -> str:
        """
        The creation time.
        """
        return pulumi.get(self, "create_time")

    @property
    @pulumi.getter
    def destination(self) -> 'outputs.DestinationResponse':
        """
        Destination specifies where the events should be sent to.
        """
        return pulumi.get(self, "destination")

    @property
    @pulumi.getter
    def etag(self) -> str:
        """
        This checksum is computed by the server based on the value of other fields, and may be sent only on create requests to ensure the client has an up-to-date value before proceeding.
        """
        return pulumi.get(self, "etag")

    @property
    @pulumi.getter
    def labels(self) -> Mapping[str, str]:
        """
        Optional. User labels attached to the triggers that can be used to group resources.
        """
        return pulumi.get(self, "labels")

    @property
    @pulumi.getter(name="matchingCriteria")
    def matching_criteria(self) -> Sequence['outputs.MatchingCriteriaResponse']:
        """
        null The criteria by which events are filtered. Only events that match with this criteria will be sent to the destination.
        """
        return pulumi.get(self, "matching_criteria")

    @property
    @pulumi.getter
    def name(self) -> str:
        """
        The resource name of the trigger. Must be unique within the location on the project and must in `projects/{project}/locations/{location}/triggers/{trigger}` format.
        """
        return pulumi.get(self, "name")

    @property
    @pulumi.getter(name="serviceAccount")
    def service_account(self) -> str:
        """
        Optional. The IAM service account email associated with the trigger. The service account represents the identity of the trigger. The principal who calls this API must have `iam.serviceAccounts.actAs` permission in the service account. See https://cloud.google.com/iam/docs/understanding-service-accounts?hl=en#sa_common for more information. For Cloud Run destinations, this service account is used to generate identity tokens when invoking the service. See https://cloud.google.com/run/docs/triggering/pubsub-push#create-service-account for information on how to invoke authenticated Cloud Run services. In order to create Audit Log triggers, the service account should also have 'eventarc.events.receiveAuditLogV1Written' permission.
        """
        return pulumi.get(self, "service_account")

    @property
    @pulumi.getter
    def transport(self) -> 'outputs.TransportResponse':
        """
        In order to deliver messages, Eventarc may use other GCP products as transport intermediary. This field contains a reference to that transport intermediary. This information can be used for debugging purposes.
        """
        return pulumi.get(self, "transport")

    @property
    @pulumi.getter(name="updateTime")
    def update_time(self) -> str:
        """
        The last-modified time.
        """
        return pulumi.get(self, "update_time")


class AwaitableGetTriggerResult(GetTriggerResult):
    # pylint: disable=using-constant-test
    def __await__(self):
        if False:
            yield self
        return GetTriggerResult(
            create_time=self.create_time,
            destination=self.destination,
            etag=self.etag,
            labels=self.labels,
            matching_criteria=self.matching_criteria,
            name=self.name,
            service_account=self.service_account,
            transport=self.transport,
            update_time=self.update_time)


def get_trigger(location: Optional[str] = None,
                project: Optional[str] = None,
                trigger_id: Optional[str] = None,
                opts: Optional[pulumi.InvokeOptions] = None) -> AwaitableGetTriggerResult:
    """
    Get a single trigger.
    """
    __args__ = dict()
    __args__['location'] = location
    __args__['project'] = project
    __args__['triggerId'] = trigger_id
    if opts is None:
        opts = pulumi.InvokeOptions()
    if opts.version is None:
        opts.version = _utilities.get_version()
    __ret__ = pulumi.runtime.invoke('google-native:eventarc/v1beta1:getTrigger', __args__, opts=opts, typ=GetTriggerResult).value

    return AwaitableGetTriggerResult(
        create_time=__ret__.create_time,
        destination=__ret__.destination,
        etag=__ret__.etag,
        labels=__ret__.labels,
        matching_criteria=__ret__.matching_criteria,
        name=__ret__.name,
        service_account=__ret__.service_account,
        transport=__ret__.transport,
        update_time=__ret__.update_time)
