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

__all__ = ['InstructionArgs', 'Instruction']

@pulumi.input_type
class InstructionArgs:
    def __init__(__self__, *,
                 instructions_id: pulumi.Input[str],
                 projects_id: pulumi.Input[str],
                 blocking_resources: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]] = None,
                 create_time: Optional[pulumi.Input[str]] = None,
                 csv_instruction: Optional[pulumi.Input['GoogleCloudDatalabelingV1beta1CsvInstructionArgs']] = None,
                 data_type: Optional[pulumi.Input[str]] = None,
                 description: Optional[pulumi.Input[str]] = None,
                 display_name: Optional[pulumi.Input[str]] = None,
                 name: Optional[pulumi.Input[str]] = None,
                 pdf_instruction: Optional[pulumi.Input['GoogleCloudDatalabelingV1beta1PdfInstructionArgs']] = None,
                 update_time: Optional[pulumi.Input[str]] = None):
        """
        The set of arguments for constructing a Instruction resource.
        :param pulumi.Input[Sequence[pulumi.Input[str]]] blocking_resources: The names of any related resources that are blocking changes to the instruction.
        :param pulumi.Input[str] create_time: Creation time of instruction.
        :param pulumi.Input['GoogleCloudDatalabelingV1beta1CsvInstructionArgs'] csv_instruction: Deprecated: this instruction format is not supported any more. Instruction from a CSV file, such as for classification task. The CSV file should have exact two columns, in the following format: * The first column is labeled data, such as an image reference, text. * The second column is comma separated labels associated with data.
        :param pulumi.Input[str] data_type: Required. The data type of this instruction.
        :param pulumi.Input[str] description: Optional. User-provided description of the instruction. The description can be up to 10000 characters long.
        :param pulumi.Input[str] display_name: Required. The display name of the instruction. Maximum of 64 characters.
        :param pulumi.Input[str] name: Instruction resource name, format: projects/{project_id}/instructions/{instruction_id}
        :param pulumi.Input['GoogleCloudDatalabelingV1beta1PdfInstructionArgs'] pdf_instruction: Instruction from a PDF document. The PDF should be in a Cloud Storage bucket.
        :param pulumi.Input[str] update_time: Last update time of instruction.
        """
        pulumi.set(__self__, "instructions_id", instructions_id)
        pulumi.set(__self__, "projects_id", projects_id)
        if blocking_resources is not None:
            pulumi.set(__self__, "blocking_resources", blocking_resources)
        if create_time is not None:
            pulumi.set(__self__, "create_time", create_time)
        if csv_instruction is not None:
            pulumi.set(__self__, "csv_instruction", csv_instruction)
        if data_type is not None:
            pulumi.set(__self__, "data_type", data_type)
        if description is not None:
            pulumi.set(__self__, "description", description)
        if display_name is not None:
            pulumi.set(__self__, "display_name", display_name)
        if name is not None:
            pulumi.set(__self__, "name", name)
        if pdf_instruction is not None:
            pulumi.set(__self__, "pdf_instruction", pdf_instruction)
        if update_time is not None:
            pulumi.set(__self__, "update_time", update_time)

    @property
    @pulumi.getter(name="instructionsId")
    def instructions_id(self) -> pulumi.Input[str]:
        return pulumi.get(self, "instructions_id")

    @instructions_id.setter
    def instructions_id(self, value: pulumi.Input[str]):
        pulumi.set(self, "instructions_id", value)

    @property
    @pulumi.getter(name="projectsId")
    def projects_id(self) -> pulumi.Input[str]:
        return pulumi.get(self, "projects_id")

    @projects_id.setter
    def projects_id(self, value: pulumi.Input[str]):
        pulumi.set(self, "projects_id", value)

    @property
    @pulumi.getter(name="blockingResources")
    def blocking_resources(self) -> Optional[pulumi.Input[Sequence[pulumi.Input[str]]]]:
        """
        The names of any related resources that are blocking changes to the instruction.
        """
        return pulumi.get(self, "blocking_resources")

    @blocking_resources.setter
    def blocking_resources(self, value: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]]):
        pulumi.set(self, "blocking_resources", value)

    @property
    @pulumi.getter(name="createTime")
    def create_time(self) -> Optional[pulumi.Input[str]]:
        """
        Creation time of instruction.
        """
        return pulumi.get(self, "create_time")

    @create_time.setter
    def create_time(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "create_time", value)

    @property
    @pulumi.getter(name="csvInstruction")
    def csv_instruction(self) -> Optional[pulumi.Input['GoogleCloudDatalabelingV1beta1CsvInstructionArgs']]:
        """
        Deprecated: this instruction format is not supported any more. Instruction from a CSV file, such as for classification task. The CSV file should have exact two columns, in the following format: * The first column is labeled data, such as an image reference, text. * The second column is comma separated labels associated with data.
        """
        return pulumi.get(self, "csv_instruction")

    @csv_instruction.setter
    def csv_instruction(self, value: Optional[pulumi.Input['GoogleCloudDatalabelingV1beta1CsvInstructionArgs']]):
        pulumi.set(self, "csv_instruction", value)

    @property
    @pulumi.getter(name="dataType")
    def data_type(self) -> Optional[pulumi.Input[str]]:
        """
        Required. The data type of this instruction.
        """
        return pulumi.get(self, "data_type")

    @data_type.setter
    def data_type(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "data_type", value)

    @property
    @pulumi.getter
    def description(self) -> Optional[pulumi.Input[str]]:
        """
        Optional. User-provided description of the instruction. The description can be up to 10000 characters long.
        """
        return pulumi.get(self, "description")

    @description.setter
    def description(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "description", value)

    @property
    @pulumi.getter(name="displayName")
    def display_name(self) -> Optional[pulumi.Input[str]]:
        """
        Required. The display name of the instruction. Maximum of 64 characters.
        """
        return pulumi.get(self, "display_name")

    @display_name.setter
    def display_name(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "display_name", value)

    @property
    @pulumi.getter
    def name(self) -> Optional[pulumi.Input[str]]:
        """
        Instruction resource name, format: projects/{project_id}/instructions/{instruction_id}
        """
        return pulumi.get(self, "name")

    @name.setter
    def name(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "name", value)

    @property
    @pulumi.getter(name="pdfInstruction")
    def pdf_instruction(self) -> Optional[pulumi.Input['GoogleCloudDatalabelingV1beta1PdfInstructionArgs']]:
        """
        Instruction from a PDF document. The PDF should be in a Cloud Storage bucket.
        """
        return pulumi.get(self, "pdf_instruction")

    @pdf_instruction.setter
    def pdf_instruction(self, value: Optional[pulumi.Input['GoogleCloudDatalabelingV1beta1PdfInstructionArgs']]):
        pulumi.set(self, "pdf_instruction", value)

    @property
    @pulumi.getter(name="updateTime")
    def update_time(self) -> Optional[pulumi.Input[str]]:
        """
        Last update time of instruction.
        """
        return pulumi.get(self, "update_time")

    @update_time.setter
    def update_time(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "update_time", value)


class Instruction(pulumi.CustomResource):
    @overload
    def __init__(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 blocking_resources: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]] = None,
                 create_time: Optional[pulumi.Input[str]] = None,
                 csv_instruction: Optional[pulumi.Input[pulumi.InputType['GoogleCloudDatalabelingV1beta1CsvInstructionArgs']]] = None,
                 data_type: Optional[pulumi.Input[str]] = None,
                 description: Optional[pulumi.Input[str]] = None,
                 display_name: Optional[pulumi.Input[str]] = None,
                 instructions_id: Optional[pulumi.Input[str]] = None,
                 name: Optional[pulumi.Input[str]] = None,
                 pdf_instruction: Optional[pulumi.Input[pulumi.InputType['GoogleCloudDatalabelingV1beta1PdfInstructionArgs']]] = None,
                 projects_id: Optional[pulumi.Input[str]] = None,
                 update_time: Optional[pulumi.Input[str]] = None,
                 __props__=None):
        """
        Creates an instruction for how data should be labeled.

        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[Sequence[pulumi.Input[str]]] blocking_resources: The names of any related resources that are blocking changes to the instruction.
        :param pulumi.Input[str] create_time: Creation time of instruction.
        :param pulumi.Input[pulumi.InputType['GoogleCloudDatalabelingV1beta1CsvInstructionArgs']] csv_instruction: Deprecated: this instruction format is not supported any more. Instruction from a CSV file, such as for classification task. The CSV file should have exact two columns, in the following format: * The first column is labeled data, such as an image reference, text. * The second column is comma separated labels associated with data.
        :param pulumi.Input[str] data_type: Required. The data type of this instruction.
        :param pulumi.Input[str] description: Optional. User-provided description of the instruction. The description can be up to 10000 characters long.
        :param pulumi.Input[str] display_name: Required. The display name of the instruction. Maximum of 64 characters.
        :param pulumi.Input[str] name: Instruction resource name, format: projects/{project_id}/instructions/{instruction_id}
        :param pulumi.Input[pulumi.InputType['GoogleCloudDatalabelingV1beta1PdfInstructionArgs']] pdf_instruction: Instruction from a PDF document. The PDF should be in a Cloud Storage bucket.
        :param pulumi.Input[str] update_time: Last update time of instruction.
        """
        ...
    @overload
    def __init__(__self__,
                 resource_name: str,
                 args: InstructionArgs,
                 opts: Optional[pulumi.ResourceOptions] = None):
        """
        Creates an instruction for how data should be labeled.

        :param str resource_name: The name of the resource.
        :param InstructionArgs args: The arguments to use to populate this resource's properties.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        ...
    def __init__(__self__, resource_name: str, *args, **kwargs):
        resource_args, opts = _utilities.get_resource_args_opts(InstructionArgs, pulumi.ResourceOptions, *args, **kwargs)
        if resource_args is not None:
            __self__._internal_init(resource_name, opts, **resource_args.__dict__)
        else:
            __self__._internal_init(resource_name, *args, **kwargs)

    def _internal_init(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 blocking_resources: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]] = None,
                 create_time: Optional[pulumi.Input[str]] = None,
                 csv_instruction: Optional[pulumi.Input[pulumi.InputType['GoogleCloudDatalabelingV1beta1CsvInstructionArgs']]] = None,
                 data_type: Optional[pulumi.Input[str]] = None,
                 description: Optional[pulumi.Input[str]] = None,
                 display_name: Optional[pulumi.Input[str]] = None,
                 instructions_id: Optional[pulumi.Input[str]] = None,
                 name: Optional[pulumi.Input[str]] = None,
                 pdf_instruction: Optional[pulumi.Input[pulumi.InputType['GoogleCloudDatalabelingV1beta1PdfInstructionArgs']]] = None,
                 projects_id: Optional[pulumi.Input[str]] = None,
                 update_time: Optional[pulumi.Input[str]] = None,
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
            __props__ = InstructionArgs.__new__(InstructionArgs)

            __props__.__dict__["blocking_resources"] = blocking_resources
            __props__.__dict__["create_time"] = create_time
            __props__.__dict__["csv_instruction"] = csv_instruction
            __props__.__dict__["data_type"] = data_type
            __props__.__dict__["description"] = description
            __props__.__dict__["display_name"] = display_name
            if instructions_id is None and not opts.urn:
                raise TypeError("Missing required property 'instructions_id'")
            __props__.__dict__["instructions_id"] = instructions_id
            __props__.__dict__["name"] = name
            __props__.__dict__["pdf_instruction"] = pdf_instruction
            if projects_id is None and not opts.urn:
                raise TypeError("Missing required property 'projects_id'")
            __props__.__dict__["projects_id"] = projects_id
            __props__.__dict__["update_time"] = update_time
        super(Instruction, __self__).__init__(
            'gcp-native:datalabeling/v1beta1:Instruction',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name: str,
            id: pulumi.Input[str],
            opts: Optional[pulumi.ResourceOptions] = None) -> 'Instruction':
        """
        Get an existing Instruction resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param pulumi.Input[str] id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = InstructionArgs.__new__(InstructionArgs)

        __props__.__dict__["blocking_resources"] = None
        __props__.__dict__["create_time"] = None
        __props__.__dict__["csv_instruction"] = None
        __props__.__dict__["data_type"] = None
        __props__.__dict__["description"] = None
        __props__.__dict__["display_name"] = None
        __props__.__dict__["name"] = None
        __props__.__dict__["pdf_instruction"] = None
        __props__.__dict__["update_time"] = None
        return Instruction(resource_name, opts=opts, __props__=__props__)

    @property
    @pulumi.getter(name="blockingResources")
    def blocking_resources(self) -> pulumi.Output[Sequence[str]]:
        """
        The names of any related resources that are blocking changes to the instruction.
        """
        return pulumi.get(self, "blocking_resources")

    @property
    @pulumi.getter(name="createTime")
    def create_time(self) -> pulumi.Output[str]:
        """
        Creation time of instruction.
        """
        return pulumi.get(self, "create_time")

    @property
    @pulumi.getter(name="csvInstruction")
    def csv_instruction(self) -> pulumi.Output['outputs.GoogleCloudDatalabelingV1beta1CsvInstructionResponse']:
        """
        Deprecated: this instruction format is not supported any more. Instruction from a CSV file, such as for classification task. The CSV file should have exact two columns, in the following format: * The first column is labeled data, such as an image reference, text. * The second column is comma separated labels associated with data.
        """
        return pulumi.get(self, "csv_instruction")

    @property
    @pulumi.getter(name="dataType")
    def data_type(self) -> pulumi.Output[str]:
        """
        Required. The data type of this instruction.
        """
        return pulumi.get(self, "data_type")

    @property
    @pulumi.getter
    def description(self) -> pulumi.Output[str]:
        """
        Optional. User-provided description of the instruction. The description can be up to 10000 characters long.
        """
        return pulumi.get(self, "description")

    @property
    @pulumi.getter(name="displayName")
    def display_name(self) -> pulumi.Output[str]:
        """
        Required. The display name of the instruction. Maximum of 64 characters.
        """
        return pulumi.get(self, "display_name")

    @property
    @pulumi.getter
    def name(self) -> pulumi.Output[str]:
        """
        Instruction resource name, format: projects/{project_id}/instructions/{instruction_id}
        """
        return pulumi.get(self, "name")

    @property
    @pulumi.getter(name="pdfInstruction")
    def pdf_instruction(self) -> pulumi.Output['outputs.GoogleCloudDatalabelingV1beta1PdfInstructionResponse']:
        """
        Instruction from a PDF document. The PDF should be in a Cloud Storage bucket.
        """
        return pulumi.get(self, "pdf_instruction")

    @property
    @pulumi.getter(name="updateTime")
    def update_time(self) -> pulumi.Output[str]:
        """
        Last update time of instruction.
        """
        return pulumi.get(self, "update_time")

