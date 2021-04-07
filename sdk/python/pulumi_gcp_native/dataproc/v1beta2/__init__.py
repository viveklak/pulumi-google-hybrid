# coding=utf-8
# *** WARNING: this file was generated by the Pulumi SDK Generator. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

# Export this package's modules as members:
from .autoscaling_policy import *
from .autoscaling_policy_iam_policy import *
from .region_autoscaling_policy import *
from .region_autoscaling_policy_iam_policy import *
from .region_cluster import *
from .region_cluster_iam_policy import *
from .region_job_iam_policy import *
from .region_operation_iam_policy import *
from .region_workflow_template import *
from .region_workflow_template_iam_policy import *
from .workflow_template import *
from .workflow_template_iam_policy import *
from ._inputs import *
from . import outputs

def _register_module():
    import pulumi
    from ... import _utilities


    class Module(pulumi.runtime.ResourceModule):
        _version = _utilities.get_semver_version()

        def version(self):
            return Module._version

        def construct(self, name: str, typ: str, urn: str) -> pulumi.Resource:
            if typ == "gcp-native:dataproc/v1beta2:AutoscalingPolicy":
                return AutoscalingPolicy(name, pulumi.ResourceOptions(urn=urn))
            elif typ == "gcp-native:dataproc/v1beta2:AutoscalingPolicyIamPolicy":
                return AutoscalingPolicyIamPolicy(name, pulumi.ResourceOptions(urn=urn))
            elif typ == "gcp-native:dataproc/v1beta2:RegionAutoscalingPolicy":
                return RegionAutoscalingPolicy(name, pulumi.ResourceOptions(urn=urn))
            elif typ == "gcp-native:dataproc/v1beta2:RegionAutoscalingPolicyIamPolicy":
                return RegionAutoscalingPolicyIamPolicy(name, pulumi.ResourceOptions(urn=urn))
            elif typ == "gcp-native:dataproc/v1beta2:RegionCluster":
                return RegionCluster(name, pulumi.ResourceOptions(urn=urn))
            elif typ == "gcp-native:dataproc/v1beta2:RegionClusterIamPolicy":
                return RegionClusterIamPolicy(name, pulumi.ResourceOptions(urn=urn))
            elif typ == "gcp-native:dataproc/v1beta2:RegionJobIamPolicy":
                return RegionJobIamPolicy(name, pulumi.ResourceOptions(urn=urn))
            elif typ == "gcp-native:dataproc/v1beta2:RegionOperationIamPolicy":
                return RegionOperationIamPolicy(name, pulumi.ResourceOptions(urn=urn))
            elif typ == "gcp-native:dataproc/v1beta2:RegionWorkflowTemplate":
                return RegionWorkflowTemplate(name, pulumi.ResourceOptions(urn=urn))
            elif typ == "gcp-native:dataproc/v1beta2:RegionWorkflowTemplateIamPolicy":
                return RegionWorkflowTemplateIamPolicy(name, pulumi.ResourceOptions(urn=urn))
            elif typ == "gcp-native:dataproc/v1beta2:WorkflowTemplate":
                return WorkflowTemplate(name, pulumi.ResourceOptions(urn=urn))
            elif typ == "gcp-native:dataproc/v1beta2:WorkflowTemplateIamPolicy":
                return WorkflowTemplateIamPolicy(name, pulumi.ResourceOptions(urn=urn))
            else:
                raise Exception(f"unknown resource type {typ}")


    _module_instance = Module()
    pulumi.runtime.register_resource_module("gcp-native", "dataproc/v1beta2", _module_instance)

_register_module()