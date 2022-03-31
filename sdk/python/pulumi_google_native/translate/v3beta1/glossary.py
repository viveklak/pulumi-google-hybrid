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

__all__ = ['GlossaryArgs', 'Glossary']

@pulumi.input_type
class GlossaryArgs:
    def __init__(__self__, *,
                 input_config: pulumi.Input['GlossaryInputConfigArgs'],
                 language_codes_set: Optional[pulumi.Input['LanguageCodesSetArgs']] = None,
                 language_pair: Optional[pulumi.Input['LanguageCodePairArgs']] = None,
                 location: Optional[pulumi.Input[str]] = None,
                 name: Optional[pulumi.Input[str]] = None,
                 project: Optional[pulumi.Input[str]] = None):
        """
        The set of arguments for constructing a Glossary resource.
        :param pulumi.Input['GlossaryInputConfigArgs'] input_config: Provides examples to build the glossary from. Total glossary must not exceed 10M Unicode codepoints.
        :param pulumi.Input['LanguageCodesSetArgs'] language_codes_set: Used with equivalent term set glossaries.
        :param pulumi.Input['LanguageCodePairArgs'] language_pair: Used with unidirectional glossaries.
        :param pulumi.Input[str] name: The resource name of the glossary. Glossary names have the form `projects/{project-number-or-id}/locations/{location-id}/glossaries/{glossary-id}`.
        """
        pulumi.set(__self__, "input_config", input_config)
        if language_codes_set is not None:
            pulumi.set(__self__, "language_codes_set", language_codes_set)
        if language_pair is not None:
            pulumi.set(__self__, "language_pair", language_pair)
        if location is not None:
            pulumi.set(__self__, "location", location)
        if name is not None:
            pulumi.set(__self__, "name", name)
        if project is not None:
            pulumi.set(__self__, "project", project)

    @property
    @pulumi.getter(name="inputConfig")
    def input_config(self) -> pulumi.Input['GlossaryInputConfigArgs']:
        """
        Provides examples to build the glossary from. Total glossary must not exceed 10M Unicode codepoints.
        """
        return pulumi.get(self, "input_config")

    @input_config.setter
    def input_config(self, value: pulumi.Input['GlossaryInputConfigArgs']):
        pulumi.set(self, "input_config", value)

    @property
    @pulumi.getter(name="languageCodesSet")
    def language_codes_set(self) -> Optional[pulumi.Input['LanguageCodesSetArgs']]:
        """
        Used with equivalent term set glossaries.
        """
        return pulumi.get(self, "language_codes_set")

    @language_codes_set.setter
    def language_codes_set(self, value: Optional[pulumi.Input['LanguageCodesSetArgs']]):
        pulumi.set(self, "language_codes_set", value)

    @property
    @pulumi.getter(name="languagePair")
    def language_pair(self) -> Optional[pulumi.Input['LanguageCodePairArgs']]:
        """
        Used with unidirectional glossaries.
        """
        return pulumi.get(self, "language_pair")

    @language_pair.setter
    def language_pair(self, value: Optional[pulumi.Input['LanguageCodePairArgs']]):
        pulumi.set(self, "language_pair", value)

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
        The resource name of the glossary. Glossary names have the form `projects/{project-number-or-id}/locations/{location-id}/glossaries/{glossary-id}`.
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


class Glossary(pulumi.CustomResource):
    @overload
    def __init__(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 input_config: Optional[pulumi.Input[pulumi.InputType['GlossaryInputConfigArgs']]] = None,
                 language_codes_set: Optional[pulumi.Input[pulumi.InputType['LanguageCodesSetArgs']]] = None,
                 language_pair: Optional[pulumi.Input[pulumi.InputType['LanguageCodePairArgs']]] = None,
                 location: Optional[pulumi.Input[str]] = None,
                 name: Optional[pulumi.Input[str]] = None,
                 project: Optional[pulumi.Input[str]] = None,
                 __props__=None):
        """
        Creates a glossary and returns the long-running operation. Returns NOT_FOUND, if the project doesn't exist.

        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[pulumi.InputType['GlossaryInputConfigArgs']] input_config: Provides examples to build the glossary from. Total glossary must not exceed 10M Unicode codepoints.
        :param pulumi.Input[pulumi.InputType['LanguageCodesSetArgs']] language_codes_set: Used with equivalent term set glossaries.
        :param pulumi.Input[pulumi.InputType['LanguageCodePairArgs']] language_pair: Used with unidirectional glossaries.
        :param pulumi.Input[str] name: The resource name of the glossary. Glossary names have the form `projects/{project-number-or-id}/locations/{location-id}/glossaries/{glossary-id}`.
        """
        ...
    @overload
    def __init__(__self__,
                 resource_name: str,
                 args: GlossaryArgs,
                 opts: Optional[pulumi.ResourceOptions] = None):
        """
        Creates a glossary and returns the long-running operation. Returns NOT_FOUND, if the project doesn't exist.

        :param str resource_name: The name of the resource.
        :param GlossaryArgs args: The arguments to use to populate this resource's properties.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        ...
    def __init__(__self__, resource_name: str, *args, **kwargs):
        resource_args, opts = _utilities.get_resource_args_opts(GlossaryArgs, pulumi.ResourceOptions, *args, **kwargs)
        if resource_args is not None:
            __self__._internal_init(resource_name, opts, **resource_args.__dict__)
        else:
            __self__._internal_init(resource_name, *args, **kwargs)

    def _internal_init(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 input_config: Optional[pulumi.Input[pulumi.InputType['GlossaryInputConfigArgs']]] = None,
                 language_codes_set: Optional[pulumi.Input[pulumi.InputType['LanguageCodesSetArgs']]] = None,
                 language_pair: Optional[pulumi.Input[pulumi.InputType['LanguageCodePairArgs']]] = None,
                 location: Optional[pulumi.Input[str]] = None,
                 name: Optional[pulumi.Input[str]] = None,
                 project: Optional[pulumi.Input[str]] = None,
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
            __props__ = GlossaryArgs.__new__(GlossaryArgs)

            if input_config is None and not opts.urn:
                raise TypeError("Missing required property 'input_config'")
            __props__.__dict__["input_config"] = input_config
            __props__.__dict__["language_codes_set"] = language_codes_set
            __props__.__dict__["language_pair"] = language_pair
            __props__.__dict__["location"] = location
            __props__.__dict__["name"] = name
            __props__.__dict__["project"] = project
            __props__.__dict__["end_time"] = None
            __props__.__dict__["entry_count"] = None
            __props__.__dict__["submit_time"] = None
        super(Glossary, __self__).__init__(
            'google-native:translate/v3beta1:Glossary',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name: str,
            id: pulumi.Input[str],
            opts: Optional[pulumi.ResourceOptions] = None) -> 'Glossary':
        """
        Get an existing Glossary resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param pulumi.Input[str] id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = GlossaryArgs.__new__(GlossaryArgs)

        __props__.__dict__["end_time"] = None
        __props__.__dict__["entry_count"] = None
        __props__.__dict__["input_config"] = None
        __props__.__dict__["language_codes_set"] = None
        __props__.__dict__["language_pair"] = None
        __props__.__dict__["name"] = None
        __props__.__dict__["submit_time"] = None
        return Glossary(resource_name, opts=opts, __props__=__props__)

    @property
    @pulumi.getter(name="endTime")
    def end_time(self) -> pulumi.Output[str]:
        """
        When the glossary creation was finished.
        """
        return pulumi.get(self, "end_time")

    @property
    @pulumi.getter(name="entryCount")
    def entry_count(self) -> pulumi.Output[int]:
        """
        The number of entries defined in the glossary.
        """
        return pulumi.get(self, "entry_count")

    @property
    @pulumi.getter(name="inputConfig")
    def input_config(self) -> pulumi.Output['outputs.GlossaryInputConfigResponse']:
        """
        Provides examples to build the glossary from. Total glossary must not exceed 10M Unicode codepoints.
        """
        return pulumi.get(self, "input_config")

    @property
    @pulumi.getter(name="languageCodesSet")
    def language_codes_set(self) -> pulumi.Output['outputs.LanguageCodesSetResponse']:
        """
        Used with equivalent term set glossaries.
        """
        return pulumi.get(self, "language_codes_set")

    @property
    @pulumi.getter(name="languagePair")
    def language_pair(self) -> pulumi.Output['outputs.LanguageCodePairResponse']:
        """
        Used with unidirectional glossaries.
        """
        return pulumi.get(self, "language_pair")

    @property
    @pulumi.getter
    def name(self) -> pulumi.Output[str]:
        """
        The resource name of the glossary. Glossary names have the form `projects/{project-number-or-id}/locations/{location-id}/glossaries/{glossary-id}`.
        """
        return pulumi.get(self, "name")

    @property
    @pulumi.getter(name="submitTime")
    def submit_time(self) -> pulumi.Output[str]:
        """
        When CreateGlossary was called.
        """
        return pulumi.get(self, "submit_time")

