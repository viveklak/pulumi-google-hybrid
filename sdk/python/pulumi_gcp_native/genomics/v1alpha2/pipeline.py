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

__all__ = ['PipelineArgs', 'Pipeline']

@pulumi.input_type
class PipelineArgs:
    def __init__(__self__, *,
                 pipeline_id: pulumi.Input[str],
                 description: Optional[pulumi.Input[str]] = None,
                 docker: Optional[pulumi.Input['DockerExecutorArgs']] = None,
                 input_parameters: Optional[pulumi.Input[Sequence[pulumi.Input['PipelineParameterArgs']]]] = None,
                 name: Optional[pulumi.Input[str]] = None,
                 output_parameters: Optional[pulumi.Input[Sequence[pulumi.Input['PipelineParameterArgs']]]] = None,
                 project_id: Optional[pulumi.Input[str]] = None,
                 resources: Optional[pulumi.Input['PipelineResourcesArgs']] = None):
        """
        The set of arguments for constructing a Pipeline resource.
        :param pulumi.Input[str] pipeline_id: Unique pipeline id that is generated by the service when CreatePipeline is called. Cannot be specified in the Pipeline used in the CreatePipelineRequest, and will be populated in the response to CreatePipeline and all subsequent Get and List calls. Indicates that the service has registered this pipeline.
        :param pulumi.Input[str] description: User-specified description.
        :param pulumi.Input['DockerExecutorArgs'] docker: Specifies the docker run information.
        :param pulumi.Input[Sequence[pulumi.Input['PipelineParameterArgs']]] input_parameters: Input parameters of the pipeline.
        :param pulumi.Input[str] name: Required. A user specified pipeline name that does not have to be unique. This name can be used for filtering Pipelines in ListPipelines.
        :param pulumi.Input[Sequence[pulumi.Input['PipelineParameterArgs']]] output_parameters: Output parameters of the pipeline.
        :param pulumi.Input[str] project_id: Required. The project in which to create the pipeline. The caller must have WRITE access.
        :param pulumi.Input['PipelineResourcesArgs'] resources: Required. Specifies resource requirements for the pipeline run. Required fields: * minimumCpuCores * minimumRamGb
        """
        pulumi.set(__self__, "pipeline_id", pipeline_id)
        if description is not None:
            pulumi.set(__self__, "description", description)
        if docker is not None:
            pulumi.set(__self__, "docker", docker)
        if input_parameters is not None:
            pulumi.set(__self__, "input_parameters", input_parameters)
        if name is not None:
            pulumi.set(__self__, "name", name)
        if output_parameters is not None:
            pulumi.set(__self__, "output_parameters", output_parameters)
        if project_id is not None:
            pulumi.set(__self__, "project_id", project_id)
        if resources is not None:
            pulumi.set(__self__, "resources", resources)

    @property
    @pulumi.getter(name="pipelineId")
    def pipeline_id(self) -> pulumi.Input[str]:
        """
        Unique pipeline id that is generated by the service when CreatePipeline is called. Cannot be specified in the Pipeline used in the CreatePipelineRequest, and will be populated in the response to CreatePipeline and all subsequent Get and List calls. Indicates that the service has registered this pipeline.
        """
        return pulumi.get(self, "pipeline_id")

    @pipeline_id.setter
    def pipeline_id(self, value: pulumi.Input[str]):
        pulumi.set(self, "pipeline_id", value)

    @property
    @pulumi.getter
    def description(self) -> Optional[pulumi.Input[str]]:
        """
        User-specified description.
        """
        return pulumi.get(self, "description")

    @description.setter
    def description(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "description", value)

    @property
    @pulumi.getter
    def docker(self) -> Optional[pulumi.Input['DockerExecutorArgs']]:
        """
        Specifies the docker run information.
        """
        return pulumi.get(self, "docker")

    @docker.setter
    def docker(self, value: Optional[pulumi.Input['DockerExecutorArgs']]):
        pulumi.set(self, "docker", value)

    @property
    @pulumi.getter(name="inputParameters")
    def input_parameters(self) -> Optional[pulumi.Input[Sequence[pulumi.Input['PipelineParameterArgs']]]]:
        """
        Input parameters of the pipeline.
        """
        return pulumi.get(self, "input_parameters")

    @input_parameters.setter
    def input_parameters(self, value: Optional[pulumi.Input[Sequence[pulumi.Input['PipelineParameterArgs']]]]):
        pulumi.set(self, "input_parameters", value)

    @property
    @pulumi.getter
    def name(self) -> Optional[pulumi.Input[str]]:
        """
        Required. A user specified pipeline name that does not have to be unique. This name can be used for filtering Pipelines in ListPipelines.
        """
        return pulumi.get(self, "name")

    @name.setter
    def name(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "name", value)

    @property
    @pulumi.getter(name="outputParameters")
    def output_parameters(self) -> Optional[pulumi.Input[Sequence[pulumi.Input['PipelineParameterArgs']]]]:
        """
        Output parameters of the pipeline.
        """
        return pulumi.get(self, "output_parameters")

    @output_parameters.setter
    def output_parameters(self, value: Optional[pulumi.Input[Sequence[pulumi.Input['PipelineParameterArgs']]]]):
        pulumi.set(self, "output_parameters", value)

    @property
    @pulumi.getter(name="projectId")
    def project_id(self) -> Optional[pulumi.Input[str]]:
        """
        Required. The project in which to create the pipeline. The caller must have WRITE access.
        """
        return pulumi.get(self, "project_id")

    @project_id.setter
    def project_id(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "project_id", value)

    @property
    @pulumi.getter
    def resources(self) -> Optional[pulumi.Input['PipelineResourcesArgs']]:
        """
        Required. Specifies resource requirements for the pipeline run. Required fields: * minimumCpuCores * minimumRamGb
        """
        return pulumi.get(self, "resources")

    @resources.setter
    def resources(self, value: Optional[pulumi.Input['PipelineResourcesArgs']]):
        pulumi.set(self, "resources", value)


class Pipeline(pulumi.CustomResource):
    @overload
    def __init__(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 description: Optional[pulumi.Input[str]] = None,
                 docker: Optional[pulumi.Input[pulumi.InputType['DockerExecutorArgs']]] = None,
                 input_parameters: Optional[pulumi.Input[Sequence[pulumi.Input[pulumi.InputType['PipelineParameterArgs']]]]] = None,
                 name: Optional[pulumi.Input[str]] = None,
                 output_parameters: Optional[pulumi.Input[Sequence[pulumi.Input[pulumi.InputType['PipelineParameterArgs']]]]] = None,
                 pipeline_id: Optional[pulumi.Input[str]] = None,
                 project_id: Optional[pulumi.Input[str]] = None,
                 resources: Optional[pulumi.Input[pulumi.InputType['PipelineResourcesArgs']]] = None,
                 __props__=None):
        """
        Creates a pipeline that can be run later. Create takes a Pipeline that has all fields other than `pipelineId` populated, and then returns the same pipeline with `pipelineId` populated. This id can be used to run the pipeline. Caller must have WRITE permission to the project.

        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] description: User-specified description.
        :param pulumi.Input[pulumi.InputType['DockerExecutorArgs']] docker: Specifies the docker run information.
        :param pulumi.Input[Sequence[pulumi.Input[pulumi.InputType['PipelineParameterArgs']]]] input_parameters: Input parameters of the pipeline.
        :param pulumi.Input[str] name: Required. A user specified pipeline name that does not have to be unique. This name can be used for filtering Pipelines in ListPipelines.
        :param pulumi.Input[Sequence[pulumi.Input[pulumi.InputType['PipelineParameterArgs']]]] output_parameters: Output parameters of the pipeline.
        :param pulumi.Input[str] pipeline_id: Unique pipeline id that is generated by the service when CreatePipeline is called. Cannot be specified in the Pipeline used in the CreatePipelineRequest, and will be populated in the response to CreatePipeline and all subsequent Get and List calls. Indicates that the service has registered this pipeline.
        :param pulumi.Input[str] project_id: Required. The project in which to create the pipeline. The caller must have WRITE access.
        :param pulumi.Input[pulumi.InputType['PipelineResourcesArgs']] resources: Required. Specifies resource requirements for the pipeline run. Required fields: * minimumCpuCores * minimumRamGb
        """
        ...
    @overload
    def __init__(__self__,
                 resource_name: str,
                 args: PipelineArgs,
                 opts: Optional[pulumi.ResourceOptions] = None):
        """
        Creates a pipeline that can be run later. Create takes a Pipeline that has all fields other than `pipelineId` populated, and then returns the same pipeline with `pipelineId` populated. This id can be used to run the pipeline. Caller must have WRITE permission to the project.

        :param str resource_name: The name of the resource.
        :param PipelineArgs args: The arguments to use to populate this resource's properties.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        ...
    def __init__(__self__, resource_name: str, *args, **kwargs):
        resource_args, opts = _utilities.get_resource_args_opts(PipelineArgs, pulumi.ResourceOptions, *args, **kwargs)
        if resource_args is not None:
            __self__._internal_init(resource_name, opts, **resource_args.__dict__)
        else:
            __self__._internal_init(resource_name, *args, **kwargs)

    def _internal_init(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 description: Optional[pulumi.Input[str]] = None,
                 docker: Optional[pulumi.Input[pulumi.InputType['DockerExecutorArgs']]] = None,
                 input_parameters: Optional[pulumi.Input[Sequence[pulumi.Input[pulumi.InputType['PipelineParameterArgs']]]]] = None,
                 name: Optional[pulumi.Input[str]] = None,
                 output_parameters: Optional[pulumi.Input[Sequence[pulumi.Input[pulumi.InputType['PipelineParameterArgs']]]]] = None,
                 pipeline_id: Optional[pulumi.Input[str]] = None,
                 project_id: Optional[pulumi.Input[str]] = None,
                 resources: Optional[pulumi.Input[pulumi.InputType['PipelineResourcesArgs']]] = None,
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
            __props__ = PipelineArgs.__new__(PipelineArgs)

            __props__.__dict__["description"] = description
            __props__.__dict__["docker"] = docker
            __props__.__dict__["input_parameters"] = input_parameters
            __props__.__dict__["name"] = name
            __props__.__dict__["output_parameters"] = output_parameters
            if pipeline_id is None and not opts.urn:
                raise TypeError("Missing required property 'pipeline_id'")
            __props__.__dict__["pipeline_id"] = pipeline_id
            __props__.__dict__["project_id"] = project_id
            __props__.__dict__["resources"] = resources
        super(Pipeline, __self__).__init__(
            'gcp-native:genomics/v1alpha2:Pipeline',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name: str,
            id: pulumi.Input[str],
            opts: Optional[pulumi.ResourceOptions] = None) -> 'Pipeline':
        """
        Get an existing Pipeline resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param pulumi.Input[str] id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = PipelineArgs.__new__(PipelineArgs)

        __props__.__dict__["description"] = None
        __props__.__dict__["docker"] = None
        __props__.__dict__["input_parameters"] = None
        __props__.__dict__["name"] = None
        __props__.__dict__["output_parameters"] = None
        __props__.__dict__["pipeline_id"] = None
        __props__.__dict__["project_id"] = None
        __props__.__dict__["resources"] = None
        return Pipeline(resource_name, opts=opts, __props__=__props__)

    @property
    @pulumi.getter
    def description(self) -> pulumi.Output[str]:
        """
        User-specified description.
        """
        return pulumi.get(self, "description")

    @property
    @pulumi.getter
    def docker(self) -> pulumi.Output['outputs.DockerExecutorResponse']:
        """
        Specifies the docker run information.
        """
        return pulumi.get(self, "docker")

    @property
    @pulumi.getter(name="inputParameters")
    def input_parameters(self) -> pulumi.Output[Sequence['outputs.PipelineParameterResponse']]:
        """
        Input parameters of the pipeline.
        """
        return pulumi.get(self, "input_parameters")

    @property
    @pulumi.getter
    def name(self) -> pulumi.Output[str]:
        """
        Required. A user specified pipeline name that does not have to be unique. This name can be used for filtering Pipelines in ListPipelines.
        """
        return pulumi.get(self, "name")

    @property
    @pulumi.getter(name="outputParameters")
    def output_parameters(self) -> pulumi.Output[Sequence['outputs.PipelineParameterResponse']]:
        """
        Output parameters of the pipeline.
        """
        return pulumi.get(self, "output_parameters")

    @property
    @pulumi.getter(name="pipelineId")
    def pipeline_id(self) -> pulumi.Output[str]:
        """
        Unique pipeline id that is generated by the service when CreatePipeline is called. Cannot be specified in the Pipeline used in the CreatePipelineRequest, and will be populated in the response to CreatePipeline and all subsequent Get and List calls. Indicates that the service has registered this pipeline.
        """
        return pulumi.get(self, "pipeline_id")

    @property
    @pulumi.getter(name="projectId")
    def project_id(self) -> pulumi.Output[str]:
        """
        Required. The project in which to create the pipeline. The caller must have WRITE access.
        """
        return pulumi.get(self, "project_id")

    @property
    @pulumi.getter
    def resources(self) -> pulumi.Output['outputs.PipelineResourcesResponse']:
        """
        Required. Specifies resource requirements for the pipeline run. Required fields: * minimumCpuCores * minimumRamGb
        """
        return pulumi.get(self, "resources")

