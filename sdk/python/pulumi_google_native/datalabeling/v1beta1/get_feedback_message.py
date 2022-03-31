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
    'GetFeedbackMessageResult',
    'AwaitableGetFeedbackMessageResult',
    'get_feedback_message',
    'get_feedback_message_output',
]

@pulumi.output_type
class GetFeedbackMessageResult:
    def __init__(__self__, body=None, create_time=None, image=None, name=None, operator_feedback_metadata=None, requester_feedback_metadata=None):
        if body and not isinstance(body, str):
            raise TypeError("Expected argument 'body' to be a str")
        pulumi.set(__self__, "body", body)
        if create_time and not isinstance(create_time, str):
            raise TypeError("Expected argument 'create_time' to be a str")
        pulumi.set(__self__, "create_time", create_time)
        if image and not isinstance(image, str):
            raise TypeError("Expected argument 'image' to be a str")
        pulumi.set(__self__, "image", image)
        if name and not isinstance(name, str):
            raise TypeError("Expected argument 'name' to be a str")
        pulumi.set(__self__, "name", name)
        if operator_feedback_metadata and not isinstance(operator_feedback_metadata, dict):
            raise TypeError("Expected argument 'operator_feedback_metadata' to be a dict")
        pulumi.set(__self__, "operator_feedback_metadata", operator_feedback_metadata)
        if requester_feedback_metadata and not isinstance(requester_feedback_metadata, dict):
            raise TypeError("Expected argument 'requester_feedback_metadata' to be a dict")
        pulumi.set(__self__, "requester_feedback_metadata", requester_feedback_metadata)

    @property
    @pulumi.getter
    def body(self) -> str:
        """
        String content of the feedback. Maximum of 10000 characters.
        """
        return pulumi.get(self, "body")

    @property
    @pulumi.getter(name="createTime")
    def create_time(self) -> str:
        """
        Create time.
        """
        return pulumi.get(self, "create_time")

    @property
    @pulumi.getter
    def image(self) -> str:
        """
        The image storing this feedback if the feedback is an image representing operator's comments.
        """
        return pulumi.get(self, "image")

    @property
    @pulumi.getter
    def name(self) -> str:
        """
        Name of the feedback message in a feedback thread. Format: 'project/{project_id}/datasets/{dataset_id}/annotatedDatasets/{annotated_dataset_id}/feedbackThreads/{feedback_thread_id}/feedbackMessage/{feedback_message_id}'
        """
        return pulumi.get(self, "name")

    @property
    @pulumi.getter(name="operatorFeedbackMetadata")
    def operator_feedback_metadata(self) -> 'outputs.GoogleCloudDatalabelingV1beta1OperatorFeedbackMetadataResponse':
        return pulumi.get(self, "operator_feedback_metadata")

    @property
    @pulumi.getter(name="requesterFeedbackMetadata")
    def requester_feedback_metadata(self) -> 'outputs.GoogleCloudDatalabelingV1beta1RequesterFeedbackMetadataResponse':
        return pulumi.get(self, "requester_feedback_metadata")


class AwaitableGetFeedbackMessageResult(GetFeedbackMessageResult):
    # pylint: disable=using-constant-test
    def __await__(self):
        if False:
            yield self
        return GetFeedbackMessageResult(
            body=self.body,
            create_time=self.create_time,
            image=self.image,
            name=self.name,
            operator_feedback_metadata=self.operator_feedback_metadata,
            requester_feedback_metadata=self.requester_feedback_metadata)


def get_feedback_message(annotated_dataset_id: Optional[str] = None,
                         dataset_id: Optional[str] = None,
                         feedback_message_id: Optional[str] = None,
                         feedback_thread_id: Optional[str] = None,
                         project: Optional[str] = None,
                         opts: Optional[pulumi.InvokeOptions] = None) -> AwaitableGetFeedbackMessageResult:
    """
    Get a FeedbackMessage object.
    """
    __args__ = dict()
    __args__['annotatedDatasetId'] = annotated_dataset_id
    __args__['datasetId'] = dataset_id
    __args__['feedbackMessageId'] = feedback_message_id
    __args__['feedbackThreadId'] = feedback_thread_id
    __args__['project'] = project
    if opts is None:
        opts = pulumi.InvokeOptions()
    if opts.version is None:
        opts.version = _utilities.get_version()
    __ret__ = pulumi.runtime.invoke('google-native:datalabeling/v1beta1:getFeedbackMessage', __args__, opts=opts, typ=GetFeedbackMessageResult).value

    return AwaitableGetFeedbackMessageResult(
        body=__ret__.body,
        create_time=__ret__.create_time,
        image=__ret__.image,
        name=__ret__.name,
        operator_feedback_metadata=__ret__.operator_feedback_metadata,
        requester_feedback_metadata=__ret__.requester_feedback_metadata)


@_utilities.lift_output_func(get_feedback_message)
def get_feedback_message_output(annotated_dataset_id: Optional[pulumi.Input[str]] = None,
                                dataset_id: Optional[pulumi.Input[str]] = None,
                                feedback_message_id: Optional[pulumi.Input[str]] = None,
                                feedback_thread_id: Optional[pulumi.Input[str]] = None,
                                project: Optional[pulumi.Input[Optional[str]]] = None,
                                opts: Optional[pulumi.InvokeOptions] = None) -> pulumi.Output[GetFeedbackMessageResult]:
    """
    Get a FeedbackMessage object.
    """
    ...
