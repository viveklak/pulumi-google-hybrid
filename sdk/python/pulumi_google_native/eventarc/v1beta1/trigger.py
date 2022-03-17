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

__all__ = ['TriggerArgs', 'Trigger']

@pulumi.input_type
class TriggerArgs:
    def __init__(__self__, *,
                 destination: pulumi.Input['DestinationArgs'],
                 matching_criteria: pulumi.Input[Sequence[pulumi.Input['MatchingCriteriaArgs']]],
                 trigger_id: pulumi.Input[str],
                 validate_only: pulumi.Input[str],
                 labels: Optional[pulumi.Input[Mapping[str, pulumi.Input[str]]]] = None,
                 location: Optional[pulumi.Input[str]] = None,
                 name: Optional[pulumi.Input[str]] = None,
                 project: Optional[pulumi.Input[str]] = None,
                 service_account: Optional[pulumi.Input[str]] = None):
        """
        The set of arguments for constructing a Trigger resource.
        :param pulumi.Input['DestinationArgs'] destination: Destination specifies where the events should be sent to.
        :param pulumi.Input[Sequence[pulumi.Input['MatchingCriteriaArgs']]] matching_criteria: null The criteria by which events are filtered. Only events that match with this criteria will be sent to the destination.
        :param pulumi.Input[str] trigger_id: Required. The user-provided ID to be assigned to the trigger.
        :param pulumi.Input[str] validate_only: Required. If set, validate the request and preview the review, but do not actually post it.
        :param pulumi.Input[Mapping[str, pulumi.Input[str]]] labels: Optional. User labels attached to the triggers that can be used to group resources.
        :param pulumi.Input[str] name: The resource name of the trigger. Must be unique within the location on the project and must in `projects/{project}/locations/{location}/triggers/{trigger}` format.
        :param pulumi.Input[str] service_account: Optional. The IAM service account email associated with the trigger. The service account represents the identity of the trigger. The principal who calls this API must have `iam.serviceAccounts.actAs` permission in the service account. See https://cloud.google.com/iam/docs/understanding-service-accounts?hl=en#sa_common for more information. For Cloud Run destinations, this service account is used to generate identity tokens when invoking the service. See https://cloud.google.com/run/docs/triggering/pubsub-push#create-service-account for information on how to invoke authenticated Cloud Run services. In order to create Audit Log triggers, the service account should also have 'eventarc.events.receiveAuditLogV1Written' permission.
        """
        pulumi.set(__self__, "destination", destination)
        pulumi.set(__self__, "matching_criteria", matching_criteria)
        pulumi.set(__self__, "trigger_id", trigger_id)
        pulumi.set(__self__, "validate_only", validate_only)
        if labels is not None:
            pulumi.set(__self__, "labels", labels)
        if location is not None:
            pulumi.set(__self__, "location", location)
        if name is not None:
            pulumi.set(__self__, "name", name)
        if project is not None:
            pulumi.set(__self__, "project", project)
        if service_account is not None:
            pulumi.set(__self__, "service_account", service_account)

    @property
    @pulumi.getter
    def destination(self) -> pulumi.Input['DestinationArgs']:
        """
        Destination specifies where the events should be sent to.
        """
        return pulumi.get(self, "destination")

    @destination.setter
    def destination(self, value: pulumi.Input['DestinationArgs']):
        pulumi.set(self, "destination", value)

    @property
    @pulumi.getter(name="matchingCriteria")
    def matching_criteria(self) -> pulumi.Input[Sequence[pulumi.Input['MatchingCriteriaArgs']]]:
        """
        null The criteria by which events are filtered. Only events that match with this criteria will be sent to the destination.
        """
        return pulumi.get(self, "matching_criteria")

    @matching_criteria.setter
    def matching_criteria(self, value: pulumi.Input[Sequence[pulumi.Input['MatchingCriteriaArgs']]]):
        pulumi.set(self, "matching_criteria", value)

    @property
    @pulumi.getter(name="triggerId")
    def trigger_id(self) -> pulumi.Input[str]:
        """
        Required. The user-provided ID to be assigned to the trigger.
        """
        return pulumi.get(self, "trigger_id")

    @trigger_id.setter
    def trigger_id(self, value: pulumi.Input[str]):
        pulumi.set(self, "trigger_id", value)

    @property
    @pulumi.getter(name="validateOnly")
    def validate_only(self) -> pulumi.Input[str]:
        """
        Required. If set, validate the request and preview the review, but do not actually post it.
        """
        return pulumi.get(self, "validate_only")

    @validate_only.setter
    def validate_only(self, value: pulumi.Input[str]):
        pulumi.set(self, "validate_only", value)

    @property
    @pulumi.getter
    def labels(self) -> Optional[pulumi.Input[Mapping[str, pulumi.Input[str]]]]:
        """
        Optional. User labels attached to the triggers that can be used to group resources.
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
    def name(self) -> Optional[pulumi.Input[str]]:
        """
        The resource name of the trigger. Must be unique within the location on the project and must in `projects/{project}/locations/{location}/triggers/{trigger}` format.
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
    @pulumi.getter(name="serviceAccount")
    def service_account(self) -> Optional[pulumi.Input[str]]:
        """
        Optional. The IAM service account email associated with the trigger. The service account represents the identity of the trigger. The principal who calls this API must have `iam.serviceAccounts.actAs` permission in the service account. See https://cloud.google.com/iam/docs/understanding-service-accounts?hl=en#sa_common for more information. For Cloud Run destinations, this service account is used to generate identity tokens when invoking the service. See https://cloud.google.com/run/docs/triggering/pubsub-push#create-service-account for information on how to invoke authenticated Cloud Run services. In order to create Audit Log triggers, the service account should also have 'eventarc.events.receiveAuditLogV1Written' permission.
        """
        return pulumi.get(self, "service_account")

    @service_account.setter
    def service_account(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "service_account", value)


class Trigger(pulumi.CustomResource):
    @overload
    def __init__(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 destination: Optional[pulumi.Input[pulumi.InputType['DestinationArgs']]] = None,
                 labels: Optional[pulumi.Input[Mapping[str, pulumi.Input[str]]]] = None,
                 location: Optional[pulumi.Input[str]] = None,
                 matching_criteria: Optional[pulumi.Input[Sequence[pulumi.Input[pulumi.InputType['MatchingCriteriaArgs']]]]] = None,
                 name: Optional[pulumi.Input[str]] = None,
                 project: Optional[pulumi.Input[str]] = None,
                 service_account: Optional[pulumi.Input[str]] = None,
                 trigger_id: Optional[pulumi.Input[str]] = None,
                 validate_only: Optional[pulumi.Input[str]] = None,
                 __props__=None):
        """
        Create a new trigger in a particular project and location.

        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[pulumi.InputType['DestinationArgs']] destination: Destination specifies where the events should be sent to.
        :param pulumi.Input[Mapping[str, pulumi.Input[str]]] labels: Optional. User labels attached to the triggers that can be used to group resources.
        :param pulumi.Input[Sequence[pulumi.Input[pulumi.InputType['MatchingCriteriaArgs']]]] matching_criteria: null The criteria by which events are filtered. Only events that match with this criteria will be sent to the destination.
        :param pulumi.Input[str] name: The resource name of the trigger. Must be unique within the location on the project and must in `projects/{project}/locations/{location}/triggers/{trigger}` format.
        :param pulumi.Input[str] service_account: Optional. The IAM service account email associated with the trigger. The service account represents the identity of the trigger. The principal who calls this API must have `iam.serviceAccounts.actAs` permission in the service account. See https://cloud.google.com/iam/docs/understanding-service-accounts?hl=en#sa_common for more information. For Cloud Run destinations, this service account is used to generate identity tokens when invoking the service. See https://cloud.google.com/run/docs/triggering/pubsub-push#create-service-account for information on how to invoke authenticated Cloud Run services. In order to create Audit Log triggers, the service account should also have 'eventarc.events.receiveAuditLogV1Written' permission.
        :param pulumi.Input[str] trigger_id: Required. The user-provided ID to be assigned to the trigger.
        :param pulumi.Input[str] validate_only: Required. If set, validate the request and preview the review, but do not actually post it.
        """
        ...
    @overload
    def __init__(__self__,
                 resource_name: str,
                 args: TriggerArgs,
                 opts: Optional[pulumi.ResourceOptions] = None):
        """
        Create a new trigger in a particular project and location.

        :param str resource_name: The name of the resource.
        :param TriggerArgs args: The arguments to use to populate this resource's properties.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        ...
    def __init__(__self__, resource_name: str, *args, **kwargs):
        resource_args, opts = _utilities.get_resource_args_opts(TriggerArgs, pulumi.ResourceOptions, *args, **kwargs)
        if resource_args is not None:
            __self__._internal_init(resource_name, opts, **resource_args.__dict__)
        else:
            __self__._internal_init(resource_name, *args, **kwargs)

    def _internal_init(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 destination: Optional[pulumi.Input[pulumi.InputType['DestinationArgs']]] = None,
                 labels: Optional[pulumi.Input[Mapping[str, pulumi.Input[str]]]] = None,
                 location: Optional[pulumi.Input[str]] = None,
                 matching_criteria: Optional[pulumi.Input[Sequence[pulumi.Input[pulumi.InputType['MatchingCriteriaArgs']]]]] = None,
                 name: Optional[pulumi.Input[str]] = None,
                 project: Optional[pulumi.Input[str]] = None,
                 service_account: Optional[pulumi.Input[str]] = None,
                 trigger_id: Optional[pulumi.Input[str]] = None,
                 validate_only: Optional[pulumi.Input[str]] = None,
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
            __props__ = TriggerArgs.__new__(TriggerArgs)

            if destination is None and not opts.urn:
                raise TypeError("Missing required property 'destination'")
            __props__.__dict__["destination"] = destination
            __props__.__dict__["labels"] = labels
            __props__.__dict__["location"] = location
            if matching_criteria is None and not opts.urn:
                raise TypeError("Missing required property 'matching_criteria'")
            __props__.__dict__["matching_criteria"] = matching_criteria
            __props__.__dict__["name"] = name
            __props__.__dict__["project"] = project
            __props__.__dict__["service_account"] = service_account
            if trigger_id is None and not opts.urn:
                raise TypeError("Missing required property 'trigger_id'")
            __props__.__dict__["trigger_id"] = trigger_id
            if validate_only is None and not opts.urn:
                raise TypeError("Missing required property 'validate_only'")
            __props__.__dict__["validate_only"] = validate_only
            __props__.__dict__["create_time"] = None
            __props__.__dict__["etag"] = None
            __props__.__dict__["transport"] = None
            __props__.__dict__["update_time"] = None
        super(Trigger, __self__).__init__(
            'google-native:eventarc/v1beta1:Trigger',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name: str,
            id: pulumi.Input[str],
            opts: Optional[pulumi.ResourceOptions] = None) -> 'Trigger':
        """
        Get an existing Trigger resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param pulumi.Input[str] id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = TriggerArgs.__new__(TriggerArgs)

        __props__.__dict__["create_time"] = None
        __props__.__dict__["destination"] = None
        __props__.__dict__["etag"] = None
        __props__.__dict__["labels"] = None
        __props__.__dict__["matching_criteria"] = None
        __props__.__dict__["name"] = None
        __props__.__dict__["service_account"] = None
        __props__.__dict__["transport"] = None
        __props__.__dict__["update_time"] = None
        return Trigger(resource_name, opts=opts, __props__=__props__)

    @property
    @pulumi.getter(name="createTime")
    def create_time(self) -> pulumi.Output[str]:
        """
        The creation time.
        """
        return pulumi.get(self, "create_time")

    @property
    @pulumi.getter
    def destination(self) -> pulumi.Output['outputs.DestinationResponse']:
        """
        Destination specifies where the events should be sent to.
        """
        return pulumi.get(self, "destination")

    @property
    @pulumi.getter
    def etag(self) -> pulumi.Output[str]:
        """
        This checksum is computed by the server based on the value of other fields, and may be sent only on create requests to ensure the client has an up-to-date value before proceeding.
        """
        return pulumi.get(self, "etag")

    @property
    @pulumi.getter
    def labels(self) -> pulumi.Output[Mapping[str, str]]:
        """
        Optional. User labels attached to the triggers that can be used to group resources.
        """
        return pulumi.get(self, "labels")

    @property
    @pulumi.getter(name="matchingCriteria")
    def matching_criteria(self) -> pulumi.Output[Sequence['outputs.MatchingCriteriaResponse']]:
        """
        null The criteria by which events are filtered. Only events that match with this criteria will be sent to the destination.
        """
        return pulumi.get(self, "matching_criteria")

    @property
    @pulumi.getter
    def name(self) -> pulumi.Output[str]:
        """
        The resource name of the trigger. Must be unique within the location on the project and must in `projects/{project}/locations/{location}/triggers/{trigger}` format.
        """
        return pulumi.get(self, "name")

    @property
    @pulumi.getter(name="serviceAccount")
    def service_account(self) -> pulumi.Output[str]:
        """
        Optional. The IAM service account email associated with the trigger. The service account represents the identity of the trigger. The principal who calls this API must have `iam.serviceAccounts.actAs` permission in the service account. See https://cloud.google.com/iam/docs/understanding-service-accounts?hl=en#sa_common for more information. For Cloud Run destinations, this service account is used to generate identity tokens when invoking the service. See https://cloud.google.com/run/docs/triggering/pubsub-push#create-service-account for information on how to invoke authenticated Cloud Run services. In order to create Audit Log triggers, the service account should also have 'eventarc.events.receiveAuditLogV1Written' permission.
        """
        return pulumi.get(self, "service_account")

    @property
    @pulumi.getter
    def transport(self) -> pulumi.Output['outputs.TransportResponse']:
        """
        In order to deliver messages, Eventarc may use other GCP products as transport intermediary. This field contains a reference to that transport intermediary. This information can be used for debugging purposes.
        """
        return pulumi.get(self, "transport")

    @property
    @pulumi.getter(name="updateTime")
    def update_time(self) -> pulumi.Output[str]:
        """
        The last-modified time.
        """
        return pulumi.get(self, "update_time")

