# coding=utf-8
# *** WARNING: this file was generated by the Pulumi SDK Generator. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import warnings
import pulumi
import pulumi.runtime
from typing import Any, Mapping, Optional, Sequence, Union, overload
from ... import _utilities

__all__ = ['ConversationParticipantArgs', 'ConversationParticipant']

@pulumi.input_type
class ConversationParticipantArgs:
    def __init__(__self__, *,
                 conversations_id: pulumi.Input[str],
                 locations_id: pulumi.Input[str],
                 participants_id: pulumi.Input[str],
                 projects_id: pulumi.Input[str],
                 name: Optional[pulumi.Input[str]] = None,
                 obfuscated_external_user_id: Optional[pulumi.Input[str]] = None,
                 role: Optional[pulumi.Input[str]] = None):
        """
        The set of arguments for constructing a ConversationParticipant resource.
        :param pulumi.Input[str] name: Optional. The unique identifier of this participant. Format: `projects//locations//conversations//participants/`.
        :param pulumi.Input[str] obfuscated_external_user_id: Optional. Obfuscated user id that should be associated with the created participant. You can specify a user id as follows: 1. If you set this field in CreateParticipantRequest or UpdateParticipantRequest, Dialogflow adds the obfuscated user id with the participant. 2. If you set this field in AnalyzeContent or StreamingAnalyzeContent, Dialogflow will update Participant.obfuscated_external_user_id. Dialogflow uses this user id for following purposes: 1) Billing and measurement. If user with the same obfuscated_external_user_id is created in a later conversation, dialogflow will know it's the same user. 2) Agent assist suggestion personalization. For example, Dialogflow can use it to provide personalized smart reply suggestions for this user. Note: * Please never pass raw user ids to Dialogflow. Always obfuscate your user id first. * Dialogflow only accepts a UTF-8 encoded string, e.g., a hex digest of a hash function like SHA-512. * The length of the user id must be <= 256 characters.
        :param pulumi.Input[str] role: Immutable. The role this participant plays in the conversation. This field must be set during participant creation and is then immutable.
        """
        pulumi.set(__self__, "conversations_id", conversations_id)
        pulumi.set(__self__, "locations_id", locations_id)
        pulumi.set(__self__, "participants_id", participants_id)
        pulumi.set(__self__, "projects_id", projects_id)
        if name is not None:
            pulumi.set(__self__, "name", name)
        if obfuscated_external_user_id is not None:
            pulumi.set(__self__, "obfuscated_external_user_id", obfuscated_external_user_id)
        if role is not None:
            pulumi.set(__self__, "role", role)

    @property
    @pulumi.getter(name="conversationsId")
    def conversations_id(self) -> pulumi.Input[str]:
        return pulumi.get(self, "conversations_id")

    @conversations_id.setter
    def conversations_id(self, value: pulumi.Input[str]):
        pulumi.set(self, "conversations_id", value)

    @property
    @pulumi.getter(name="locationsId")
    def locations_id(self) -> pulumi.Input[str]:
        return pulumi.get(self, "locations_id")

    @locations_id.setter
    def locations_id(self, value: pulumi.Input[str]):
        pulumi.set(self, "locations_id", value)

    @property
    @pulumi.getter(name="participantsId")
    def participants_id(self) -> pulumi.Input[str]:
        return pulumi.get(self, "participants_id")

    @participants_id.setter
    def participants_id(self, value: pulumi.Input[str]):
        pulumi.set(self, "participants_id", value)

    @property
    @pulumi.getter(name="projectsId")
    def projects_id(self) -> pulumi.Input[str]:
        return pulumi.get(self, "projects_id")

    @projects_id.setter
    def projects_id(self, value: pulumi.Input[str]):
        pulumi.set(self, "projects_id", value)

    @property
    @pulumi.getter
    def name(self) -> Optional[pulumi.Input[str]]:
        """
        Optional. The unique identifier of this participant. Format: `projects//locations//conversations//participants/`.
        """
        return pulumi.get(self, "name")

    @name.setter
    def name(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "name", value)

    @property
    @pulumi.getter(name="obfuscatedExternalUserId")
    def obfuscated_external_user_id(self) -> Optional[pulumi.Input[str]]:
        """
        Optional. Obfuscated user id that should be associated with the created participant. You can specify a user id as follows: 1. If you set this field in CreateParticipantRequest or UpdateParticipantRequest, Dialogflow adds the obfuscated user id with the participant. 2. If you set this field in AnalyzeContent or StreamingAnalyzeContent, Dialogflow will update Participant.obfuscated_external_user_id. Dialogflow uses this user id for following purposes: 1) Billing and measurement. If user with the same obfuscated_external_user_id is created in a later conversation, dialogflow will know it's the same user. 2) Agent assist suggestion personalization. For example, Dialogflow can use it to provide personalized smart reply suggestions for this user. Note: * Please never pass raw user ids to Dialogflow. Always obfuscate your user id first. * Dialogflow only accepts a UTF-8 encoded string, e.g., a hex digest of a hash function like SHA-512. * The length of the user id must be <= 256 characters.
        """
        return pulumi.get(self, "obfuscated_external_user_id")

    @obfuscated_external_user_id.setter
    def obfuscated_external_user_id(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "obfuscated_external_user_id", value)

    @property
    @pulumi.getter
    def role(self) -> Optional[pulumi.Input[str]]:
        """
        Immutable. The role this participant plays in the conversation. This field must be set during participant creation and is then immutable.
        """
        return pulumi.get(self, "role")

    @role.setter
    def role(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "role", value)


class ConversationParticipant(pulumi.CustomResource):
    @overload
    def __init__(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 conversations_id: Optional[pulumi.Input[str]] = None,
                 locations_id: Optional[pulumi.Input[str]] = None,
                 name: Optional[pulumi.Input[str]] = None,
                 obfuscated_external_user_id: Optional[pulumi.Input[str]] = None,
                 participants_id: Optional[pulumi.Input[str]] = None,
                 projects_id: Optional[pulumi.Input[str]] = None,
                 role: Optional[pulumi.Input[str]] = None,
                 __props__=None):
        """
        Creates a new participant in a conversation.

        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] name: Optional. The unique identifier of this participant. Format: `projects//locations//conversations//participants/`.
        :param pulumi.Input[str] obfuscated_external_user_id: Optional. Obfuscated user id that should be associated with the created participant. You can specify a user id as follows: 1. If you set this field in CreateParticipantRequest or UpdateParticipantRequest, Dialogflow adds the obfuscated user id with the participant. 2. If you set this field in AnalyzeContent or StreamingAnalyzeContent, Dialogflow will update Participant.obfuscated_external_user_id. Dialogflow uses this user id for following purposes: 1) Billing and measurement. If user with the same obfuscated_external_user_id is created in a later conversation, dialogflow will know it's the same user. 2) Agent assist suggestion personalization. For example, Dialogflow can use it to provide personalized smart reply suggestions for this user. Note: * Please never pass raw user ids to Dialogflow. Always obfuscate your user id first. * Dialogflow only accepts a UTF-8 encoded string, e.g., a hex digest of a hash function like SHA-512. * The length of the user id must be <= 256 characters.
        :param pulumi.Input[str] role: Immutable. The role this participant plays in the conversation. This field must be set during participant creation and is then immutable.
        """
        ...
    @overload
    def __init__(__self__,
                 resource_name: str,
                 args: ConversationParticipantArgs,
                 opts: Optional[pulumi.ResourceOptions] = None):
        """
        Creates a new participant in a conversation.

        :param str resource_name: The name of the resource.
        :param ConversationParticipantArgs args: The arguments to use to populate this resource's properties.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        ...
    def __init__(__self__, resource_name: str, *args, **kwargs):
        resource_args, opts = _utilities.get_resource_args_opts(ConversationParticipantArgs, pulumi.ResourceOptions, *args, **kwargs)
        if resource_args is not None:
            __self__._internal_init(resource_name, opts, **resource_args.__dict__)
        else:
            __self__._internal_init(resource_name, *args, **kwargs)

    def _internal_init(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 conversations_id: Optional[pulumi.Input[str]] = None,
                 locations_id: Optional[pulumi.Input[str]] = None,
                 name: Optional[pulumi.Input[str]] = None,
                 obfuscated_external_user_id: Optional[pulumi.Input[str]] = None,
                 participants_id: Optional[pulumi.Input[str]] = None,
                 projects_id: Optional[pulumi.Input[str]] = None,
                 role: Optional[pulumi.Input[str]] = None,
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
            __props__ = ConversationParticipantArgs.__new__(ConversationParticipantArgs)

            if conversations_id is None and not opts.urn:
                raise TypeError("Missing required property 'conversations_id'")
            __props__.__dict__["conversations_id"] = conversations_id
            if locations_id is None and not opts.urn:
                raise TypeError("Missing required property 'locations_id'")
            __props__.__dict__["locations_id"] = locations_id
            __props__.__dict__["name"] = name
            __props__.__dict__["obfuscated_external_user_id"] = obfuscated_external_user_id
            if participants_id is None and not opts.urn:
                raise TypeError("Missing required property 'participants_id'")
            __props__.__dict__["participants_id"] = participants_id
            if projects_id is None and not opts.urn:
                raise TypeError("Missing required property 'projects_id'")
            __props__.__dict__["projects_id"] = projects_id
            __props__.__dict__["role"] = role
        super(ConversationParticipant, __self__).__init__(
            'gcp-native:dialogflow/v2beta1:ConversationParticipant',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name: str,
            id: pulumi.Input[str],
            opts: Optional[pulumi.ResourceOptions] = None) -> 'ConversationParticipant':
        """
        Get an existing ConversationParticipant resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param pulumi.Input[str] id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = ConversationParticipantArgs.__new__(ConversationParticipantArgs)

        __props__.__dict__["name"] = None
        __props__.__dict__["obfuscated_external_user_id"] = None
        __props__.__dict__["role"] = None
        return ConversationParticipant(resource_name, opts=opts, __props__=__props__)

    @property
    @pulumi.getter
    def name(self) -> pulumi.Output[str]:
        """
        Optional. The unique identifier of this participant. Format: `projects//locations//conversations//participants/`.
        """
        return pulumi.get(self, "name")

    @property
    @pulumi.getter(name="obfuscatedExternalUserId")
    def obfuscated_external_user_id(self) -> pulumi.Output[str]:
        """
        Optional. Obfuscated user id that should be associated with the created participant. You can specify a user id as follows: 1. If you set this field in CreateParticipantRequest or UpdateParticipantRequest, Dialogflow adds the obfuscated user id with the participant. 2. If you set this field in AnalyzeContent or StreamingAnalyzeContent, Dialogflow will update Participant.obfuscated_external_user_id. Dialogflow uses this user id for following purposes: 1) Billing and measurement. If user with the same obfuscated_external_user_id is created in a later conversation, dialogflow will know it's the same user. 2) Agent assist suggestion personalization. For example, Dialogflow can use it to provide personalized smart reply suggestions for this user. Note: * Please never pass raw user ids to Dialogflow. Always obfuscate your user id first. * Dialogflow only accepts a UTF-8 encoded string, e.g., a hex digest of a hash function like SHA-512. * The length of the user id must be <= 256 characters.
        """
        return pulumi.get(self, "obfuscated_external_user_id")

    @property
    @pulumi.getter
    def role(self) -> pulumi.Output[str]:
        """
        Immutable. The role this participant plays in the conversation. This field must be set during participant creation and is then immutable.
        """
        return pulumi.get(self, "role")

