# coding=utf-8
# *** WARNING: this file was generated by the Pulumi SDK Generator. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import warnings
import pulumi
import pulumi.runtime
from typing import Any, Mapping, Optional, Sequence, Union, overload
from ... import _utilities
from . import outputs
from ._enums import *
from ._inputs import *

__all__ = ['ConversationModelArgs', 'ConversationModel']

@pulumi.input_type
class ConversationModelArgs:
    def __init__(__self__, *,
                 datasets: pulumi.Input[Sequence[pulumi.Input['GoogleCloudDialogflowV2InputDatasetArgs']]],
                 display_name: pulumi.Input[str],
                 article_suggestion_model_metadata: Optional[pulumi.Input['GoogleCloudDialogflowV2ArticleSuggestionModelMetadataArgs']] = None,
                 language_code: Optional[pulumi.Input[str]] = None,
                 location: Optional[pulumi.Input[str]] = None,
                 name: Optional[pulumi.Input[str]] = None,
                 project: Optional[pulumi.Input[str]] = None,
                 smart_reply_model_metadata: Optional[pulumi.Input['GoogleCloudDialogflowV2SmartReplyModelMetadataArgs']] = None):
        """
        The set of arguments for constructing a ConversationModel resource.
        :param pulumi.Input[Sequence[pulumi.Input['GoogleCloudDialogflowV2InputDatasetArgs']]] datasets: Datasets used to create model.
        :param pulumi.Input[str] display_name: The display name of the model. At most 64 bytes long.
        :param pulumi.Input['GoogleCloudDialogflowV2ArticleSuggestionModelMetadataArgs'] article_suggestion_model_metadata: Metadata for article suggestion models.
        :param pulumi.Input[str] language_code: Language code for the conversation model. If not specified, the language is en-US. Language at ConversationModel should be set for all non en-us languages. This should be a [BCP-47](https://www.rfc-editor.org/rfc/bcp/bcp47.txt) language tag. Example: "en-US".
        :param pulumi.Input[str] name: ConversationModel resource name. Format: `projects//conversationModels/`
        :param pulumi.Input['GoogleCloudDialogflowV2SmartReplyModelMetadataArgs'] smart_reply_model_metadata: Metadata for smart reply models.
        """
        pulumi.set(__self__, "datasets", datasets)
        pulumi.set(__self__, "display_name", display_name)
        if article_suggestion_model_metadata is not None:
            pulumi.set(__self__, "article_suggestion_model_metadata", article_suggestion_model_metadata)
        if language_code is not None:
            pulumi.set(__self__, "language_code", language_code)
        if location is not None:
            pulumi.set(__self__, "location", location)
        if name is not None:
            pulumi.set(__self__, "name", name)
        if project is not None:
            pulumi.set(__self__, "project", project)
        if smart_reply_model_metadata is not None:
            pulumi.set(__self__, "smart_reply_model_metadata", smart_reply_model_metadata)

    @property
    @pulumi.getter
    def datasets(self) -> pulumi.Input[Sequence[pulumi.Input['GoogleCloudDialogflowV2InputDatasetArgs']]]:
        """
        Datasets used to create model.
        """
        return pulumi.get(self, "datasets")

    @datasets.setter
    def datasets(self, value: pulumi.Input[Sequence[pulumi.Input['GoogleCloudDialogflowV2InputDatasetArgs']]]):
        pulumi.set(self, "datasets", value)

    @property
    @pulumi.getter(name="displayName")
    def display_name(self) -> pulumi.Input[str]:
        """
        The display name of the model. At most 64 bytes long.
        """
        return pulumi.get(self, "display_name")

    @display_name.setter
    def display_name(self, value: pulumi.Input[str]):
        pulumi.set(self, "display_name", value)

    @property
    @pulumi.getter(name="articleSuggestionModelMetadata")
    def article_suggestion_model_metadata(self) -> Optional[pulumi.Input['GoogleCloudDialogflowV2ArticleSuggestionModelMetadataArgs']]:
        """
        Metadata for article suggestion models.
        """
        return pulumi.get(self, "article_suggestion_model_metadata")

    @article_suggestion_model_metadata.setter
    def article_suggestion_model_metadata(self, value: Optional[pulumi.Input['GoogleCloudDialogflowV2ArticleSuggestionModelMetadataArgs']]):
        pulumi.set(self, "article_suggestion_model_metadata", value)

    @property
    @pulumi.getter(name="languageCode")
    def language_code(self) -> Optional[pulumi.Input[str]]:
        """
        Language code for the conversation model. If not specified, the language is en-US. Language at ConversationModel should be set for all non en-us languages. This should be a [BCP-47](https://www.rfc-editor.org/rfc/bcp/bcp47.txt) language tag. Example: "en-US".
        """
        return pulumi.get(self, "language_code")

    @language_code.setter
    def language_code(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "language_code", value)

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
        ConversationModel resource name. Format: `projects//conversationModels/`
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
    @pulumi.getter(name="smartReplyModelMetadata")
    def smart_reply_model_metadata(self) -> Optional[pulumi.Input['GoogleCloudDialogflowV2SmartReplyModelMetadataArgs']]:
        """
        Metadata for smart reply models.
        """
        return pulumi.get(self, "smart_reply_model_metadata")

    @smart_reply_model_metadata.setter
    def smart_reply_model_metadata(self, value: Optional[pulumi.Input['GoogleCloudDialogflowV2SmartReplyModelMetadataArgs']]):
        pulumi.set(self, "smart_reply_model_metadata", value)


class ConversationModel(pulumi.CustomResource):
    @overload
    def __init__(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 article_suggestion_model_metadata: Optional[pulumi.Input[pulumi.InputType['GoogleCloudDialogflowV2ArticleSuggestionModelMetadataArgs']]] = None,
                 datasets: Optional[pulumi.Input[Sequence[pulumi.Input[pulumi.InputType['GoogleCloudDialogflowV2InputDatasetArgs']]]]] = None,
                 display_name: Optional[pulumi.Input[str]] = None,
                 language_code: Optional[pulumi.Input[str]] = None,
                 location: Optional[pulumi.Input[str]] = None,
                 name: Optional[pulumi.Input[str]] = None,
                 project: Optional[pulumi.Input[str]] = None,
                 smart_reply_model_metadata: Optional[pulumi.Input[pulumi.InputType['GoogleCloudDialogflowV2SmartReplyModelMetadataArgs']]] = None,
                 __props__=None):
        """
        Creates a model. This method is a [long-running operation](https://cloud.google.com/dialogflow/es/docs/how/long-running-operations). The returned `Operation` type has the following method-specific fields: - `metadata`: CreateConversationModelOperationMetadata - `response`: ConversationModel

        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[pulumi.InputType['GoogleCloudDialogflowV2ArticleSuggestionModelMetadataArgs']] article_suggestion_model_metadata: Metadata for article suggestion models.
        :param pulumi.Input[Sequence[pulumi.Input[pulumi.InputType['GoogleCloudDialogflowV2InputDatasetArgs']]]] datasets: Datasets used to create model.
        :param pulumi.Input[str] display_name: The display name of the model. At most 64 bytes long.
        :param pulumi.Input[str] language_code: Language code for the conversation model. If not specified, the language is en-US. Language at ConversationModel should be set for all non en-us languages. This should be a [BCP-47](https://www.rfc-editor.org/rfc/bcp/bcp47.txt) language tag. Example: "en-US".
        :param pulumi.Input[str] name: ConversationModel resource name. Format: `projects//conversationModels/`
        :param pulumi.Input[pulumi.InputType['GoogleCloudDialogflowV2SmartReplyModelMetadataArgs']] smart_reply_model_metadata: Metadata for smart reply models.
        """
        ...
    @overload
    def __init__(__self__,
                 resource_name: str,
                 args: ConversationModelArgs,
                 opts: Optional[pulumi.ResourceOptions] = None):
        """
        Creates a model. This method is a [long-running operation](https://cloud.google.com/dialogflow/es/docs/how/long-running-operations). The returned `Operation` type has the following method-specific fields: - `metadata`: CreateConversationModelOperationMetadata - `response`: ConversationModel

        :param str resource_name: The name of the resource.
        :param ConversationModelArgs args: The arguments to use to populate this resource's properties.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        ...
    def __init__(__self__, resource_name: str, *args, **kwargs):
        resource_args, opts = _utilities.get_resource_args_opts(ConversationModelArgs, pulumi.ResourceOptions, *args, **kwargs)
        if resource_args is not None:
            __self__._internal_init(resource_name, opts, **resource_args.__dict__)
        else:
            __self__._internal_init(resource_name, *args, **kwargs)

    def _internal_init(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 article_suggestion_model_metadata: Optional[pulumi.Input[pulumi.InputType['GoogleCloudDialogflowV2ArticleSuggestionModelMetadataArgs']]] = None,
                 datasets: Optional[pulumi.Input[Sequence[pulumi.Input[pulumi.InputType['GoogleCloudDialogflowV2InputDatasetArgs']]]]] = None,
                 display_name: Optional[pulumi.Input[str]] = None,
                 language_code: Optional[pulumi.Input[str]] = None,
                 location: Optional[pulumi.Input[str]] = None,
                 name: Optional[pulumi.Input[str]] = None,
                 project: Optional[pulumi.Input[str]] = None,
                 smart_reply_model_metadata: Optional[pulumi.Input[pulumi.InputType['GoogleCloudDialogflowV2SmartReplyModelMetadataArgs']]] = None,
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
            __props__ = ConversationModelArgs.__new__(ConversationModelArgs)

            __props__.__dict__["article_suggestion_model_metadata"] = article_suggestion_model_metadata
            if datasets is None and not opts.urn:
                raise TypeError("Missing required property 'datasets'")
            __props__.__dict__["datasets"] = datasets
            if display_name is None and not opts.urn:
                raise TypeError("Missing required property 'display_name'")
            __props__.__dict__["display_name"] = display_name
            __props__.__dict__["language_code"] = language_code
            __props__.__dict__["location"] = location
            __props__.__dict__["name"] = name
            __props__.__dict__["project"] = project
            __props__.__dict__["smart_reply_model_metadata"] = smart_reply_model_metadata
            __props__.__dict__["create_time"] = None
            __props__.__dict__["state"] = None
        super(ConversationModel, __self__).__init__(
            'google-native:dialogflow/v2:ConversationModel',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name: str,
            id: pulumi.Input[str],
            opts: Optional[pulumi.ResourceOptions] = None) -> 'ConversationModel':
        """
        Get an existing ConversationModel resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param pulumi.Input[str] id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = ConversationModelArgs.__new__(ConversationModelArgs)

        __props__.__dict__["article_suggestion_model_metadata"] = None
        __props__.__dict__["create_time"] = None
        __props__.__dict__["datasets"] = None
        __props__.__dict__["display_name"] = None
        __props__.__dict__["language_code"] = None
        __props__.__dict__["name"] = None
        __props__.__dict__["smart_reply_model_metadata"] = None
        __props__.__dict__["state"] = None
        return ConversationModel(resource_name, opts=opts, __props__=__props__)

    @property
    @pulumi.getter(name="articleSuggestionModelMetadata")
    def article_suggestion_model_metadata(self) -> pulumi.Output['outputs.GoogleCloudDialogflowV2ArticleSuggestionModelMetadataResponse']:
        """
        Metadata for article suggestion models.
        """
        return pulumi.get(self, "article_suggestion_model_metadata")

    @property
    @pulumi.getter(name="createTime")
    def create_time(self) -> pulumi.Output[str]:
        """
        Creation time of this model.
        """
        return pulumi.get(self, "create_time")

    @property
    @pulumi.getter
    def datasets(self) -> pulumi.Output[Sequence['outputs.GoogleCloudDialogflowV2InputDatasetResponse']]:
        """
        Datasets used to create model.
        """
        return pulumi.get(self, "datasets")

    @property
    @pulumi.getter(name="displayName")
    def display_name(self) -> pulumi.Output[str]:
        """
        The display name of the model. At most 64 bytes long.
        """
        return pulumi.get(self, "display_name")

    @property
    @pulumi.getter(name="languageCode")
    def language_code(self) -> pulumi.Output[str]:
        """
        Language code for the conversation model. If not specified, the language is en-US. Language at ConversationModel should be set for all non en-us languages. This should be a [BCP-47](https://www.rfc-editor.org/rfc/bcp/bcp47.txt) language tag. Example: "en-US".
        """
        return pulumi.get(self, "language_code")

    @property
    @pulumi.getter
    def name(self) -> pulumi.Output[str]:
        """
        ConversationModel resource name. Format: `projects//conversationModels/`
        """
        return pulumi.get(self, "name")

    @property
    @pulumi.getter(name="smartReplyModelMetadata")
    def smart_reply_model_metadata(self) -> pulumi.Output['outputs.GoogleCloudDialogflowV2SmartReplyModelMetadataResponse']:
        """
        Metadata for smart reply models.
        """
        return pulumi.get(self, "smart_reply_model_metadata")

    @property
    @pulumi.getter
    def state(self) -> pulumi.Output[str]:
        """
        State of the model. A model can only serve prediction requests after it gets deployed.
        """
        return pulumi.get(self, "state")

