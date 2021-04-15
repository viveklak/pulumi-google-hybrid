# coding=utf-8
# *** WARNING: this file was generated by the Pulumi SDK Generator. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import warnings
import pulumi
import pulumi.runtime
from typing import Any, Mapping, Optional, Sequence, Union, overload
from ... import _utilities

__all__ = ['KnowledgeBaseArgs', 'KnowledgeBase']

@pulumi.input_type
class KnowledgeBaseArgs:
    def __init__(__self__, *,
                 knowledge_bases_id: pulumi.Input[str],
                 locations_id: pulumi.Input[str],
                 projects_id: pulumi.Input[str],
                 display_name: Optional[pulumi.Input[str]] = None,
                 language_code: Optional[pulumi.Input[str]] = None,
                 name: Optional[pulumi.Input[str]] = None):
        """
        The set of arguments for constructing a KnowledgeBase resource.
        :param pulumi.Input[str] display_name: Required. The display name of the knowledge base. The name must be 1024 bytes or less; otherwise, the creation request fails.
        :param pulumi.Input[str] language_code: Language which represents the KnowledgeBase. When the KnowledgeBase is created/updated, this is populated for all non en-us languages. If not populated, the default language en-us applies.
        :param pulumi.Input[str] name: The knowledge base resource name. The name must be empty when creating a knowledge base. Format: `projects//locations//knowledgeBases/`.
        """
        pulumi.set(__self__, "knowledge_bases_id", knowledge_bases_id)
        pulumi.set(__self__, "locations_id", locations_id)
        pulumi.set(__self__, "projects_id", projects_id)
        if display_name is not None:
            pulumi.set(__self__, "display_name", display_name)
        if language_code is not None:
            pulumi.set(__self__, "language_code", language_code)
        if name is not None:
            pulumi.set(__self__, "name", name)

    @property
    @pulumi.getter(name="knowledgeBasesId")
    def knowledge_bases_id(self) -> pulumi.Input[str]:
        return pulumi.get(self, "knowledge_bases_id")

    @knowledge_bases_id.setter
    def knowledge_bases_id(self, value: pulumi.Input[str]):
        pulumi.set(self, "knowledge_bases_id", value)

    @property
    @pulumi.getter(name="locationsId")
    def locations_id(self) -> pulumi.Input[str]:
        return pulumi.get(self, "locations_id")

    @locations_id.setter
    def locations_id(self, value: pulumi.Input[str]):
        pulumi.set(self, "locations_id", value)

    @property
    @pulumi.getter(name="projectsId")
    def projects_id(self) -> pulumi.Input[str]:
        return pulumi.get(self, "projects_id")

    @projects_id.setter
    def projects_id(self, value: pulumi.Input[str]):
        pulumi.set(self, "projects_id", value)

    @property
    @pulumi.getter(name="displayName")
    def display_name(self) -> Optional[pulumi.Input[str]]:
        """
        Required. The display name of the knowledge base. The name must be 1024 bytes or less; otherwise, the creation request fails.
        """
        return pulumi.get(self, "display_name")

    @display_name.setter
    def display_name(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "display_name", value)

    @property
    @pulumi.getter(name="languageCode")
    def language_code(self) -> Optional[pulumi.Input[str]]:
        """
        Language which represents the KnowledgeBase. When the KnowledgeBase is created/updated, this is populated for all non en-us languages. If not populated, the default language en-us applies.
        """
        return pulumi.get(self, "language_code")

    @language_code.setter
    def language_code(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "language_code", value)

    @property
    @pulumi.getter
    def name(self) -> Optional[pulumi.Input[str]]:
        """
        The knowledge base resource name. The name must be empty when creating a knowledge base. Format: `projects//locations//knowledgeBases/`.
        """
        return pulumi.get(self, "name")

    @name.setter
    def name(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "name", value)


class KnowledgeBase(pulumi.CustomResource):
    @overload
    def __init__(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 display_name: Optional[pulumi.Input[str]] = None,
                 knowledge_bases_id: Optional[pulumi.Input[str]] = None,
                 language_code: Optional[pulumi.Input[str]] = None,
                 locations_id: Optional[pulumi.Input[str]] = None,
                 name: Optional[pulumi.Input[str]] = None,
                 projects_id: Optional[pulumi.Input[str]] = None,
                 __props__=None):
        """
        Creates a knowledge base. Note: The `projects.agent.knowledgeBases` resource is deprecated; only use `projects.knowledgeBases`.

        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] display_name: Required. The display name of the knowledge base. The name must be 1024 bytes or less; otherwise, the creation request fails.
        :param pulumi.Input[str] language_code: Language which represents the KnowledgeBase. When the KnowledgeBase is created/updated, this is populated for all non en-us languages. If not populated, the default language en-us applies.
        :param pulumi.Input[str] name: The knowledge base resource name. The name must be empty when creating a knowledge base. Format: `projects//locations//knowledgeBases/`.
        """
        ...
    @overload
    def __init__(__self__,
                 resource_name: str,
                 args: KnowledgeBaseArgs,
                 opts: Optional[pulumi.ResourceOptions] = None):
        """
        Creates a knowledge base. Note: The `projects.agent.knowledgeBases` resource is deprecated; only use `projects.knowledgeBases`.

        :param str resource_name: The name of the resource.
        :param KnowledgeBaseArgs args: The arguments to use to populate this resource's properties.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        ...
    def __init__(__self__, resource_name: str, *args, **kwargs):
        resource_args, opts = _utilities.get_resource_args_opts(KnowledgeBaseArgs, pulumi.ResourceOptions, *args, **kwargs)
        if resource_args is not None:
            __self__._internal_init(resource_name, opts, **resource_args.__dict__)
        else:
            __self__._internal_init(resource_name, *args, **kwargs)

    def _internal_init(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 display_name: Optional[pulumi.Input[str]] = None,
                 knowledge_bases_id: Optional[pulumi.Input[str]] = None,
                 language_code: Optional[pulumi.Input[str]] = None,
                 locations_id: Optional[pulumi.Input[str]] = None,
                 name: Optional[pulumi.Input[str]] = None,
                 projects_id: Optional[pulumi.Input[str]] = None,
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
            __props__ = KnowledgeBaseArgs.__new__(KnowledgeBaseArgs)

            __props__.__dict__["display_name"] = display_name
            if knowledge_bases_id is None and not opts.urn:
                raise TypeError("Missing required property 'knowledge_bases_id'")
            __props__.__dict__["knowledge_bases_id"] = knowledge_bases_id
            __props__.__dict__["language_code"] = language_code
            if locations_id is None and not opts.urn:
                raise TypeError("Missing required property 'locations_id'")
            __props__.__dict__["locations_id"] = locations_id
            __props__.__dict__["name"] = name
            if projects_id is None and not opts.urn:
                raise TypeError("Missing required property 'projects_id'")
            __props__.__dict__["projects_id"] = projects_id
        super(KnowledgeBase, __self__).__init__(
            'gcp-native:dialogflow/v2beta1:KnowledgeBase',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name: str,
            id: pulumi.Input[str],
            opts: Optional[pulumi.ResourceOptions] = None) -> 'KnowledgeBase':
        """
        Get an existing KnowledgeBase resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param pulumi.Input[str] id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = KnowledgeBaseArgs.__new__(KnowledgeBaseArgs)

        __props__.__dict__["display_name"] = None
        __props__.__dict__["language_code"] = None
        __props__.__dict__["name"] = None
        return KnowledgeBase(resource_name, opts=opts, __props__=__props__)

    @property
    @pulumi.getter(name="displayName")
    def display_name(self) -> pulumi.Output[str]:
        """
        Required. The display name of the knowledge base. The name must be 1024 bytes or less; otherwise, the creation request fails.
        """
        return pulumi.get(self, "display_name")

    @property
    @pulumi.getter(name="languageCode")
    def language_code(self) -> pulumi.Output[str]:
        """
        Language which represents the KnowledgeBase. When the KnowledgeBase is created/updated, this is populated for all non en-us languages. If not populated, the default language en-us applies.
        """
        return pulumi.get(self, "language_code")

    @property
    @pulumi.getter
    def name(self) -> pulumi.Output[str]:
        """
        The knowledge base resource name. The name must be empty when creating a knowledge base. Format: `projects//locations//knowledgeBases/`.
        """
        return pulumi.get(self, "name")

