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

__all__ = ['RegionWorkflowTemplateArgs', 'RegionWorkflowTemplate']

@pulumi.input_type
class RegionWorkflowTemplateArgs:
    def __init__(__self__, *,
                 projects_id: pulumi.Input[str],
                 regions_id: pulumi.Input[str],
                 workflow_templates_id: pulumi.Input[str],
                 dag_timeout: Optional[pulumi.Input[str]] = None,
                 id: Optional[pulumi.Input[str]] = None,
                 jobs: Optional[pulumi.Input[Sequence[pulumi.Input['OrderedJobArgs']]]] = None,
                 labels: Optional[pulumi.Input[Mapping[str, pulumi.Input[str]]]] = None,
                 parameters: Optional[pulumi.Input[Sequence[pulumi.Input['TemplateParameterArgs']]]] = None,
                 placement: Optional[pulumi.Input['WorkflowTemplatePlacementArgs']] = None,
                 version: Optional[pulumi.Input[int]] = None):
        """
        The set of arguments for constructing a RegionWorkflowTemplate resource.
        :param pulumi.Input[str] dag_timeout: Optional. Timeout duration for the DAG of jobs, expressed in seconds (see JSON representation of duration (https://developers.google.com/protocol-buffers/docs/proto3#json)). The timeout duration must be from 10 minutes ("600s") to 24 hours ("86400s"). The timer begins when the first job is submitted. If the workflow is running at the end of the timeout period, any remaining jobs are cancelled, the workflow is ended, and if the workflow was running on a managed cluster, the cluster is deleted.
        :param pulumi.Input[str] id: Required. The template id.The id must contain only letters (a-z, A-Z), numbers (0-9), underscores (_), and hyphens (-). Cannot begin or end with underscore or hyphen. Must consist of between 3 and 50 characters..
        :param pulumi.Input[Sequence[pulumi.Input['OrderedJobArgs']]] jobs: Required. The Directed Acyclic Graph of Jobs to submit.
        :param pulumi.Input[Mapping[str, pulumi.Input[str]]] labels: Optional. The labels to associate with this template. These labels will be propagated to all jobs and clusters created by the workflow instance.Label keys must contain 1 to 63 characters, and must conform to RFC 1035 (https://www.ietf.org/rfc/rfc1035.txt).Label values may be empty, but, if present, must contain 1 to 63 characters, and must conform to RFC 1035 (https://www.ietf.org/rfc/rfc1035.txt).No more than 32 labels can be associated with a template.
        :param pulumi.Input[Sequence[pulumi.Input['TemplateParameterArgs']]] parameters: Optional. Template parameters whose values are substituted into the template. Values for parameters must be provided when the template is instantiated.
        :param pulumi.Input['WorkflowTemplatePlacementArgs'] placement: Required. WorkflowTemplate scheduling information.
        :param pulumi.Input[int] version: Optional. Used to perform a consistent read-modify-write.This field should be left blank for a CreateWorkflowTemplate request. It is required for an UpdateWorkflowTemplate request, and must match the current server version. A typical update template flow would fetch the current template with a GetWorkflowTemplate request, which will return the current template with the version field filled in with the current server version. The user updates other fields in the template, then returns it as part of the UpdateWorkflowTemplate request.
        """
        pulumi.set(__self__, "projects_id", projects_id)
        pulumi.set(__self__, "regions_id", regions_id)
        pulumi.set(__self__, "workflow_templates_id", workflow_templates_id)
        if dag_timeout is not None:
            pulumi.set(__self__, "dag_timeout", dag_timeout)
        if id is not None:
            pulumi.set(__self__, "id", id)
        if jobs is not None:
            pulumi.set(__self__, "jobs", jobs)
        if labels is not None:
            pulumi.set(__self__, "labels", labels)
        if parameters is not None:
            pulumi.set(__self__, "parameters", parameters)
        if placement is not None:
            pulumi.set(__self__, "placement", placement)
        if version is not None:
            pulumi.set(__self__, "version", version)

    @property
    @pulumi.getter(name="projectsId")
    def projects_id(self) -> pulumi.Input[str]:
        return pulumi.get(self, "projects_id")

    @projects_id.setter
    def projects_id(self, value: pulumi.Input[str]):
        pulumi.set(self, "projects_id", value)

    @property
    @pulumi.getter(name="regionsId")
    def regions_id(self) -> pulumi.Input[str]:
        return pulumi.get(self, "regions_id")

    @regions_id.setter
    def regions_id(self, value: pulumi.Input[str]):
        pulumi.set(self, "regions_id", value)

    @property
    @pulumi.getter(name="workflowTemplatesId")
    def workflow_templates_id(self) -> pulumi.Input[str]:
        return pulumi.get(self, "workflow_templates_id")

    @workflow_templates_id.setter
    def workflow_templates_id(self, value: pulumi.Input[str]):
        pulumi.set(self, "workflow_templates_id", value)

    @property
    @pulumi.getter(name="dagTimeout")
    def dag_timeout(self) -> Optional[pulumi.Input[str]]:
        """
        Optional. Timeout duration for the DAG of jobs, expressed in seconds (see JSON representation of duration (https://developers.google.com/protocol-buffers/docs/proto3#json)). The timeout duration must be from 10 minutes ("600s") to 24 hours ("86400s"). The timer begins when the first job is submitted. If the workflow is running at the end of the timeout period, any remaining jobs are cancelled, the workflow is ended, and if the workflow was running on a managed cluster, the cluster is deleted.
        """
        return pulumi.get(self, "dag_timeout")

    @dag_timeout.setter
    def dag_timeout(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "dag_timeout", value)

    @property
    @pulumi.getter
    def id(self) -> Optional[pulumi.Input[str]]:
        """
        Required. The template id.The id must contain only letters (a-z, A-Z), numbers (0-9), underscores (_), and hyphens (-). Cannot begin or end with underscore or hyphen. Must consist of between 3 and 50 characters..
        """
        return pulumi.get(self, "id")

    @id.setter
    def id(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "id", value)

    @property
    @pulumi.getter
    def jobs(self) -> Optional[pulumi.Input[Sequence[pulumi.Input['OrderedJobArgs']]]]:
        """
        Required. The Directed Acyclic Graph of Jobs to submit.
        """
        return pulumi.get(self, "jobs")

    @jobs.setter
    def jobs(self, value: Optional[pulumi.Input[Sequence[pulumi.Input['OrderedJobArgs']]]]):
        pulumi.set(self, "jobs", value)

    @property
    @pulumi.getter
    def labels(self) -> Optional[pulumi.Input[Mapping[str, pulumi.Input[str]]]]:
        """
        Optional. The labels to associate with this template. These labels will be propagated to all jobs and clusters created by the workflow instance.Label keys must contain 1 to 63 characters, and must conform to RFC 1035 (https://www.ietf.org/rfc/rfc1035.txt).Label values may be empty, but, if present, must contain 1 to 63 characters, and must conform to RFC 1035 (https://www.ietf.org/rfc/rfc1035.txt).No more than 32 labels can be associated with a template.
        """
        return pulumi.get(self, "labels")

    @labels.setter
    def labels(self, value: Optional[pulumi.Input[Mapping[str, pulumi.Input[str]]]]):
        pulumi.set(self, "labels", value)

    @property
    @pulumi.getter
    def parameters(self) -> Optional[pulumi.Input[Sequence[pulumi.Input['TemplateParameterArgs']]]]:
        """
        Optional. Template parameters whose values are substituted into the template. Values for parameters must be provided when the template is instantiated.
        """
        return pulumi.get(self, "parameters")

    @parameters.setter
    def parameters(self, value: Optional[pulumi.Input[Sequence[pulumi.Input['TemplateParameterArgs']]]]):
        pulumi.set(self, "parameters", value)

    @property
    @pulumi.getter
    def placement(self) -> Optional[pulumi.Input['WorkflowTemplatePlacementArgs']]:
        """
        Required. WorkflowTemplate scheduling information.
        """
        return pulumi.get(self, "placement")

    @placement.setter
    def placement(self, value: Optional[pulumi.Input['WorkflowTemplatePlacementArgs']]):
        pulumi.set(self, "placement", value)

    @property
    @pulumi.getter
    def version(self) -> Optional[pulumi.Input[int]]:
        """
        Optional. Used to perform a consistent read-modify-write.This field should be left blank for a CreateWorkflowTemplate request. It is required for an UpdateWorkflowTemplate request, and must match the current server version. A typical update template flow would fetch the current template with a GetWorkflowTemplate request, which will return the current template with the version field filled in with the current server version. The user updates other fields in the template, then returns it as part of the UpdateWorkflowTemplate request.
        """
        return pulumi.get(self, "version")

    @version.setter
    def version(self, value: Optional[pulumi.Input[int]]):
        pulumi.set(self, "version", value)


class RegionWorkflowTemplate(pulumi.CustomResource):
    @overload
    def __init__(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 dag_timeout: Optional[pulumi.Input[str]] = None,
                 id: Optional[pulumi.Input[str]] = None,
                 jobs: Optional[pulumi.Input[Sequence[pulumi.Input[pulumi.InputType['OrderedJobArgs']]]]] = None,
                 labels: Optional[pulumi.Input[Mapping[str, pulumi.Input[str]]]] = None,
                 parameters: Optional[pulumi.Input[Sequence[pulumi.Input[pulumi.InputType['TemplateParameterArgs']]]]] = None,
                 placement: Optional[pulumi.Input[pulumi.InputType['WorkflowTemplatePlacementArgs']]] = None,
                 projects_id: Optional[pulumi.Input[str]] = None,
                 regions_id: Optional[pulumi.Input[str]] = None,
                 version: Optional[pulumi.Input[int]] = None,
                 workflow_templates_id: Optional[pulumi.Input[str]] = None,
                 __props__=None):
        """
        Creates new workflow template.

        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] dag_timeout: Optional. Timeout duration for the DAG of jobs, expressed in seconds (see JSON representation of duration (https://developers.google.com/protocol-buffers/docs/proto3#json)). The timeout duration must be from 10 minutes ("600s") to 24 hours ("86400s"). The timer begins when the first job is submitted. If the workflow is running at the end of the timeout period, any remaining jobs are cancelled, the workflow is ended, and if the workflow was running on a managed cluster, the cluster is deleted.
        :param pulumi.Input[str] id: Required. The template id.The id must contain only letters (a-z, A-Z), numbers (0-9), underscores (_), and hyphens (-). Cannot begin or end with underscore or hyphen. Must consist of between 3 and 50 characters..
        :param pulumi.Input[Sequence[pulumi.Input[pulumi.InputType['OrderedJobArgs']]]] jobs: Required. The Directed Acyclic Graph of Jobs to submit.
        :param pulumi.Input[Mapping[str, pulumi.Input[str]]] labels: Optional. The labels to associate with this template. These labels will be propagated to all jobs and clusters created by the workflow instance.Label keys must contain 1 to 63 characters, and must conform to RFC 1035 (https://www.ietf.org/rfc/rfc1035.txt).Label values may be empty, but, if present, must contain 1 to 63 characters, and must conform to RFC 1035 (https://www.ietf.org/rfc/rfc1035.txt).No more than 32 labels can be associated with a template.
        :param pulumi.Input[Sequence[pulumi.Input[pulumi.InputType['TemplateParameterArgs']]]] parameters: Optional. Template parameters whose values are substituted into the template. Values for parameters must be provided when the template is instantiated.
        :param pulumi.Input[pulumi.InputType['WorkflowTemplatePlacementArgs']] placement: Required. WorkflowTemplate scheduling information.
        :param pulumi.Input[int] version: Optional. Used to perform a consistent read-modify-write.This field should be left blank for a CreateWorkflowTemplate request. It is required for an UpdateWorkflowTemplate request, and must match the current server version. A typical update template flow would fetch the current template with a GetWorkflowTemplate request, which will return the current template with the version field filled in with the current server version. The user updates other fields in the template, then returns it as part of the UpdateWorkflowTemplate request.
        """
        ...
    @overload
    def __init__(__self__,
                 resource_name: str,
                 args: RegionWorkflowTemplateArgs,
                 opts: Optional[pulumi.ResourceOptions] = None):
        """
        Creates new workflow template.

        :param str resource_name: The name of the resource.
        :param RegionWorkflowTemplateArgs args: The arguments to use to populate this resource's properties.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        ...
    def __init__(__self__, resource_name: str, *args, **kwargs):
        resource_args, opts = _utilities.get_resource_args_opts(RegionWorkflowTemplateArgs, pulumi.ResourceOptions, *args, **kwargs)
        if resource_args is not None:
            __self__._internal_init(resource_name, opts, **resource_args.__dict__)
        else:
            __self__._internal_init(resource_name, *args, **kwargs)

    def _internal_init(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 dag_timeout: Optional[pulumi.Input[str]] = None,
                 id: Optional[pulumi.Input[str]] = None,
                 jobs: Optional[pulumi.Input[Sequence[pulumi.Input[pulumi.InputType['OrderedJobArgs']]]]] = None,
                 labels: Optional[pulumi.Input[Mapping[str, pulumi.Input[str]]]] = None,
                 parameters: Optional[pulumi.Input[Sequence[pulumi.Input[pulumi.InputType['TemplateParameterArgs']]]]] = None,
                 placement: Optional[pulumi.Input[pulumi.InputType['WorkflowTemplatePlacementArgs']]] = None,
                 projects_id: Optional[pulumi.Input[str]] = None,
                 regions_id: Optional[pulumi.Input[str]] = None,
                 version: Optional[pulumi.Input[int]] = None,
                 workflow_templates_id: Optional[pulumi.Input[str]] = None,
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
            __props__ = RegionWorkflowTemplateArgs.__new__(RegionWorkflowTemplateArgs)

            __props__.__dict__["dag_timeout"] = dag_timeout
            __props__.__dict__["id"] = id
            __props__.__dict__["jobs"] = jobs
            __props__.__dict__["labels"] = labels
            __props__.__dict__["parameters"] = parameters
            __props__.__dict__["placement"] = placement
            if projects_id is None and not opts.urn:
                raise TypeError("Missing required property 'projects_id'")
            __props__.__dict__["projects_id"] = projects_id
            if regions_id is None and not opts.urn:
                raise TypeError("Missing required property 'regions_id'")
            __props__.__dict__["regions_id"] = regions_id
            __props__.__dict__["version"] = version
            if workflow_templates_id is None and not opts.urn:
                raise TypeError("Missing required property 'workflow_templates_id'")
            __props__.__dict__["workflow_templates_id"] = workflow_templates_id
            __props__.__dict__["create_time"] = None
            __props__.__dict__["name"] = None
            __props__.__dict__["update_time"] = None
        super(RegionWorkflowTemplate, __self__).__init__(
            'gcp-native:dataproc/v1beta2:RegionWorkflowTemplate',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name: str,
            id: pulumi.Input[str],
            opts: Optional[pulumi.ResourceOptions] = None) -> 'RegionWorkflowTemplate':
        """
        Get an existing RegionWorkflowTemplate resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param pulumi.Input[str] id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = RegionWorkflowTemplateArgs.__new__(RegionWorkflowTemplateArgs)

        __props__.__dict__["create_time"] = None
        __props__.__dict__["dag_timeout"] = None
        __props__.__dict__["jobs"] = None
        __props__.__dict__["labels"] = None
        __props__.__dict__["name"] = None
        __props__.__dict__["parameters"] = None
        __props__.__dict__["placement"] = None
        __props__.__dict__["update_time"] = None
        __props__.__dict__["version"] = None
        return RegionWorkflowTemplate(resource_name, opts=opts, __props__=__props__)

    @property
    @pulumi.getter(name="createTime")
    def create_time(self) -> pulumi.Output[str]:
        """
        The time template was created.
        """
        return pulumi.get(self, "create_time")

    @property
    @pulumi.getter(name="dagTimeout")
    def dag_timeout(self) -> pulumi.Output[str]:
        """
        Optional. Timeout duration for the DAG of jobs, expressed in seconds (see JSON representation of duration (https://developers.google.com/protocol-buffers/docs/proto3#json)). The timeout duration must be from 10 minutes ("600s") to 24 hours ("86400s"). The timer begins when the first job is submitted. If the workflow is running at the end of the timeout period, any remaining jobs are cancelled, the workflow is ended, and if the workflow was running on a managed cluster, the cluster is deleted.
        """
        return pulumi.get(self, "dag_timeout")

    @property
    @pulumi.getter
    def jobs(self) -> pulumi.Output[Sequence['outputs.OrderedJobResponse']]:
        """
        Required. The Directed Acyclic Graph of Jobs to submit.
        """
        return pulumi.get(self, "jobs")

    @property
    @pulumi.getter
    def labels(self) -> pulumi.Output[Mapping[str, str]]:
        """
        Optional. The labels to associate with this template. These labels will be propagated to all jobs and clusters created by the workflow instance.Label keys must contain 1 to 63 characters, and must conform to RFC 1035 (https://www.ietf.org/rfc/rfc1035.txt).Label values may be empty, but, if present, must contain 1 to 63 characters, and must conform to RFC 1035 (https://www.ietf.org/rfc/rfc1035.txt).No more than 32 labels can be associated with a template.
        """
        return pulumi.get(self, "labels")

    @property
    @pulumi.getter
    def name(self) -> pulumi.Output[str]:
        """
        The resource name of the workflow template, as described in https://cloud.google.com/apis/design/resource_names. For projects.regions.workflowTemplates, the resource name of the template has the following format: projects/{project_id}/regions/{region}/workflowTemplates/{template_id} For projects.locations.workflowTemplates, the resource name of the template has the following format: projects/{project_id}/locations/{location}/workflowTemplates/{template_id}
        """
        return pulumi.get(self, "name")

    @property
    @pulumi.getter
    def parameters(self) -> pulumi.Output[Sequence['outputs.TemplateParameterResponse']]:
        """
        Optional. Template parameters whose values are substituted into the template. Values for parameters must be provided when the template is instantiated.
        """
        return pulumi.get(self, "parameters")

    @property
    @pulumi.getter
    def placement(self) -> pulumi.Output['outputs.WorkflowTemplatePlacementResponse']:
        """
        Required. WorkflowTemplate scheduling information.
        """
        return pulumi.get(self, "placement")

    @property
    @pulumi.getter(name="updateTime")
    def update_time(self) -> pulumi.Output[str]:
        """
        The time template was last updated.
        """
        return pulumi.get(self, "update_time")

    @property
    @pulumi.getter
    def version(self) -> pulumi.Output[int]:
        """
        Optional. Used to perform a consistent read-modify-write.This field should be left blank for a CreateWorkflowTemplate request. It is required for an UpdateWorkflowTemplate request, and must match the current server version. A typical update template flow would fetch the current template with a GetWorkflowTemplate request, which will return the current template with the version field filled in with the current server version. The user updates other fields in the template, then returns it as part of the UpdateWorkflowTemplate request.
        """
        return pulumi.get(self, "version")

