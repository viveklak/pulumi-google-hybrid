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
    'GetAutoscalingPolicyResult',
    'AwaitableGetAutoscalingPolicyResult',
    'get_autoscaling_policy',
    'get_autoscaling_policy_output',
]

@pulumi.output_type
class GetAutoscalingPolicyResult:
    def __init__(__self__, basic_algorithm=None, labels=None, name=None, secondary_worker_config=None, worker_config=None):
        if basic_algorithm and not isinstance(basic_algorithm, dict):
            raise TypeError("Expected argument 'basic_algorithm' to be a dict")
        pulumi.set(__self__, "basic_algorithm", basic_algorithm)
        if labels and not isinstance(labels, dict):
            raise TypeError("Expected argument 'labels' to be a dict")
        pulumi.set(__self__, "labels", labels)
        if name and not isinstance(name, str):
            raise TypeError("Expected argument 'name' to be a str")
        pulumi.set(__self__, "name", name)
        if secondary_worker_config and not isinstance(secondary_worker_config, dict):
            raise TypeError("Expected argument 'secondary_worker_config' to be a dict")
        pulumi.set(__self__, "secondary_worker_config", secondary_worker_config)
        if worker_config and not isinstance(worker_config, dict):
            raise TypeError("Expected argument 'worker_config' to be a dict")
        pulumi.set(__self__, "worker_config", worker_config)

    @property
    @pulumi.getter(name="basicAlgorithm")
    def basic_algorithm(self) -> 'outputs.BasicAutoscalingAlgorithmResponse':
        return pulumi.get(self, "basic_algorithm")

    @property
    @pulumi.getter
    def labels(self) -> Mapping[str, str]:
        """
        Optional. The labels to associate with this autoscaling policy. Label keys must contain 1 to 63 characters, and must conform to RFC 1035 (https://www.ietf.org/rfc/rfc1035.txt). Label values may be empty, but, if present, must contain 1 to 63 characters, and must conform to RFC 1035 (https://www.ietf.org/rfc/rfc1035.txt). No more than 32 labels can be associated with an autoscaling policy.
        """
        return pulumi.get(self, "labels")

    @property
    @pulumi.getter
    def name(self) -> str:
        """
        The "resource name" of the autoscaling policy, as described in https://cloud.google.com/apis/design/resource_names. For projects.regions.autoscalingPolicies, the resource name of the policy has the following format: projects/{project_id}/regions/{region}/autoscalingPolicies/{policy_id} For projects.locations.autoscalingPolicies, the resource name of the policy has the following format: projects/{project_id}/locations/{location}/autoscalingPolicies/{policy_id}
        """
        return pulumi.get(self, "name")

    @property
    @pulumi.getter(name="secondaryWorkerConfig")
    def secondary_worker_config(self) -> 'outputs.InstanceGroupAutoscalingPolicyConfigResponse':
        """
        Optional. Describes how the autoscaler will operate for secondary workers.
        """
        return pulumi.get(self, "secondary_worker_config")

    @property
    @pulumi.getter(name="workerConfig")
    def worker_config(self) -> 'outputs.InstanceGroupAutoscalingPolicyConfigResponse':
        """
        Describes how the autoscaler will operate for primary workers.
        """
        return pulumi.get(self, "worker_config")


class AwaitableGetAutoscalingPolicyResult(GetAutoscalingPolicyResult):
    # pylint: disable=using-constant-test
    def __await__(self):
        if False:
            yield self
        return GetAutoscalingPolicyResult(
            basic_algorithm=self.basic_algorithm,
            labels=self.labels,
            name=self.name,
            secondary_worker_config=self.secondary_worker_config,
            worker_config=self.worker_config)


def get_autoscaling_policy(autoscaling_policy_id: Optional[str] = None,
                           location: Optional[str] = None,
                           project: Optional[str] = None,
                           opts: Optional[pulumi.InvokeOptions] = None) -> AwaitableGetAutoscalingPolicyResult:
    """
    Retrieves autoscaling policy.
    """
    __args__ = dict()
    __args__['autoscalingPolicyId'] = autoscaling_policy_id
    __args__['location'] = location
    __args__['project'] = project
    if opts is None:
        opts = pulumi.InvokeOptions()
    if opts.version is None:
        opts.version = _utilities.get_version()
    __ret__ = pulumi.runtime.invoke('google-native:dataproc/v1:getAutoscalingPolicy', __args__, opts=opts, typ=GetAutoscalingPolicyResult).value

    return AwaitableGetAutoscalingPolicyResult(
        basic_algorithm=__ret__.basic_algorithm,
        labels=__ret__.labels,
        name=__ret__.name,
        secondary_worker_config=__ret__.secondary_worker_config,
        worker_config=__ret__.worker_config)


@_utilities.lift_output_func(get_autoscaling_policy)
def get_autoscaling_policy_output(autoscaling_policy_id: Optional[pulumi.Input[str]] = None,
                                  location: Optional[pulumi.Input[str]] = None,
                                  project: Optional[pulumi.Input[Optional[str]]] = None,
                                  opts: Optional[pulumi.InvokeOptions] = None) -> pulumi.Output[GetAutoscalingPolicyResult]:
    """
    Retrieves autoscaling policy.
    """
    ...
