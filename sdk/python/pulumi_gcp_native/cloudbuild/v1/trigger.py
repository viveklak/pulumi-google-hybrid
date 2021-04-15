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

__all__ = ['TriggerArgs', 'Trigger']

@pulumi.input_type
class TriggerArgs:
    def __init__(__self__, *,
                 project_id: pulumi.Input[str],
                 trigger_id: pulumi.Input[str],
                 build: Optional[pulumi.Input['BuildArgs']] = None,
                 description: Optional[pulumi.Input[str]] = None,
                 disabled: Optional[pulumi.Input[bool]] = None,
                 filename: Optional[pulumi.Input[str]] = None,
                 github: Optional[pulumi.Input['GitHubEventsConfigArgs']] = None,
                 ignored_files: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]] = None,
                 included_files: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]] = None,
                 name: Optional[pulumi.Input[str]] = None,
                 substitutions: Optional[pulumi.Input[Mapping[str, pulumi.Input[str]]]] = None,
                 tags: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]] = None,
                 trigger_template: Optional[pulumi.Input['RepoSourceArgs']] = None):
        """
        The set of arguments for constructing a Trigger resource.
        :param pulumi.Input['BuildArgs'] build: Contents of the build template.
        :param pulumi.Input[str] description: Human-readable description of this trigger.
        :param pulumi.Input[bool] disabled: If true, the trigger will never automatically execute a build.
        :param pulumi.Input[str] filename: Path, from the source root, to the build configuration file (i.e. cloudbuild.yaml).
        :param pulumi.Input['GitHubEventsConfigArgs'] github: GitHubEventsConfig describes the configuration of a trigger that creates a build whenever a GitHub event is received. Mutually exclusive with `trigger_template`.
        :param pulumi.Input[Sequence[pulumi.Input[str]]] ignored_files: ignored_files and included_files are file glob matches using https://golang.org/pkg/path/filepath/#Match extended with support for "**". If ignored_files and changed files are both empty, then they are not used to determine whether or not to trigger a build. If ignored_files is not empty, then we ignore any files that match any of the ignored_file globs. If the change has no files that are outside of the ignored_files globs, then we do not trigger a build.
        :param pulumi.Input[Sequence[pulumi.Input[str]]] included_files: If any of the files altered in the commit pass the ignored_files filter and included_files is empty, then as far as this filter is concerned, we should trigger the build. If any of the files altered in the commit pass the ignored_files filter and included_files is not empty, then we make sure that at least one of those files matches a included_files glob. If not, then we do not trigger a build.
        :param pulumi.Input[str] name: User-assigned name of the trigger. Must be unique within the project. Trigger names must meet the following requirements: + They must contain only alphanumeric characters and dashes. + They can be 1-64 characters long. + They must begin and end with an alphanumeric character.
        :param pulumi.Input[Mapping[str, pulumi.Input[str]]] substitutions: Substitutions for Build resource. The keys must match the following regular expression: `^_[A-Z0-9_]+$`.
        :param pulumi.Input[Sequence[pulumi.Input[str]]] tags: Tags for annotation of a `BuildTrigger`
        :param pulumi.Input['RepoSourceArgs'] trigger_template: Template describing the types of source changes to trigger a build. Branch and tag names in trigger templates are interpreted as regular expressions. Any branch or tag change that matches that regular expression will trigger a build. Mutually exclusive with `github`.
        """
        pulumi.set(__self__, "project_id", project_id)
        pulumi.set(__self__, "trigger_id", trigger_id)
        if build is not None:
            pulumi.set(__self__, "build", build)
        if description is not None:
            pulumi.set(__self__, "description", description)
        if disabled is not None:
            pulumi.set(__self__, "disabled", disabled)
        if filename is not None:
            pulumi.set(__self__, "filename", filename)
        if github is not None:
            pulumi.set(__self__, "github", github)
        if ignored_files is not None:
            pulumi.set(__self__, "ignored_files", ignored_files)
        if included_files is not None:
            pulumi.set(__self__, "included_files", included_files)
        if name is not None:
            pulumi.set(__self__, "name", name)
        if substitutions is not None:
            pulumi.set(__self__, "substitutions", substitutions)
        if tags is not None:
            pulumi.set(__self__, "tags", tags)
        if trigger_template is not None:
            pulumi.set(__self__, "trigger_template", trigger_template)

    @property
    @pulumi.getter(name="projectId")
    def project_id(self) -> pulumi.Input[str]:
        return pulumi.get(self, "project_id")

    @project_id.setter
    def project_id(self, value: pulumi.Input[str]):
        pulumi.set(self, "project_id", value)

    @property
    @pulumi.getter(name="triggerId")
    def trigger_id(self) -> pulumi.Input[str]:
        return pulumi.get(self, "trigger_id")

    @trigger_id.setter
    def trigger_id(self, value: pulumi.Input[str]):
        pulumi.set(self, "trigger_id", value)

    @property
    @pulumi.getter
    def build(self) -> Optional[pulumi.Input['BuildArgs']]:
        """
        Contents of the build template.
        """
        return pulumi.get(self, "build")

    @build.setter
    def build(self, value: Optional[pulumi.Input['BuildArgs']]):
        pulumi.set(self, "build", value)

    @property
    @pulumi.getter
    def description(self) -> Optional[pulumi.Input[str]]:
        """
        Human-readable description of this trigger.
        """
        return pulumi.get(self, "description")

    @description.setter
    def description(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "description", value)

    @property
    @pulumi.getter
    def disabled(self) -> Optional[pulumi.Input[bool]]:
        """
        If true, the trigger will never automatically execute a build.
        """
        return pulumi.get(self, "disabled")

    @disabled.setter
    def disabled(self, value: Optional[pulumi.Input[bool]]):
        pulumi.set(self, "disabled", value)

    @property
    @pulumi.getter
    def filename(self) -> Optional[pulumi.Input[str]]:
        """
        Path, from the source root, to the build configuration file (i.e. cloudbuild.yaml).
        """
        return pulumi.get(self, "filename")

    @filename.setter
    def filename(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "filename", value)

    @property
    @pulumi.getter
    def github(self) -> Optional[pulumi.Input['GitHubEventsConfigArgs']]:
        """
        GitHubEventsConfig describes the configuration of a trigger that creates a build whenever a GitHub event is received. Mutually exclusive with `trigger_template`.
        """
        return pulumi.get(self, "github")

    @github.setter
    def github(self, value: Optional[pulumi.Input['GitHubEventsConfigArgs']]):
        pulumi.set(self, "github", value)

    @property
    @pulumi.getter(name="ignoredFiles")
    def ignored_files(self) -> Optional[pulumi.Input[Sequence[pulumi.Input[str]]]]:
        """
        ignored_files and included_files are file glob matches using https://golang.org/pkg/path/filepath/#Match extended with support for "**". If ignored_files and changed files are both empty, then they are not used to determine whether or not to trigger a build. If ignored_files is not empty, then we ignore any files that match any of the ignored_file globs. If the change has no files that are outside of the ignored_files globs, then we do not trigger a build.
        """
        return pulumi.get(self, "ignored_files")

    @ignored_files.setter
    def ignored_files(self, value: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]]):
        pulumi.set(self, "ignored_files", value)

    @property
    @pulumi.getter(name="includedFiles")
    def included_files(self) -> Optional[pulumi.Input[Sequence[pulumi.Input[str]]]]:
        """
        If any of the files altered in the commit pass the ignored_files filter and included_files is empty, then as far as this filter is concerned, we should trigger the build. If any of the files altered in the commit pass the ignored_files filter and included_files is not empty, then we make sure that at least one of those files matches a included_files glob. If not, then we do not trigger a build.
        """
        return pulumi.get(self, "included_files")

    @included_files.setter
    def included_files(self, value: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]]):
        pulumi.set(self, "included_files", value)

    @property
    @pulumi.getter
    def name(self) -> Optional[pulumi.Input[str]]:
        """
        User-assigned name of the trigger. Must be unique within the project. Trigger names must meet the following requirements: + They must contain only alphanumeric characters and dashes. + They can be 1-64 characters long. + They must begin and end with an alphanumeric character.
        """
        return pulumi.get(self, "name")

    @name.setter
    def name(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "name", value)

    @property
    @pulumi.getter
    def substitutions(self) -> Optional[pulumi.Input[Mapping[str, pulumi.Input[str]]]]:
        """
        Substitutions for Build resource. The keys must match the following regular expression: `^_[A-Z0-9_]+$`.
        """
        return pulumi.get(self, "substitutions")

    @substitutions.setter
    def substitutions(self, value: Optional[pulumi.Input[Mapping[str, pulumi.Input[str]]]]):
        pulumi.set(self, "substitutions", value)

    @property
    @pulumi.getter
    def tags(self) -> Optional[pulumi.Input[Sequence[pulumi.Input[str]]]]:
        """
        Tags for annotation of a `BuildTrigger`
        """
        return pulumi.get(self, "tags")

    @tags.setter
    def tags(self, value: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]]):
        pulumi.set(self, "tags", value)

    @property
    @pulumi.getter(name="triggerTemplate")
    def trigger_template(self) -> Optional[pulumi.Input['RepoSourceArgs']]:
        """
        Template describing the types of source changes to trigger a build. Branch and tag names in trigger templates are interpreted as regular expressions. Any branch or tag change that matches that regular expression will trigger a build. Mutually exclusive with `github`.
        """
        return pulumi.get(self, "trigger_template")

    @trigger_template.setter
    def trigger_template(self, value: Optional[pulumi.Input['RepoSourceArgs']]):
        pulumi.set(self, "trigger_template", value)


class Trigger(pulumi.CustomResource):
    @overload
    def __init__(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 build: Optional[pulumi.Input[pulumi.InputType['BuildArgs']]] = None,
                 description: Optional[pulumi.Input[str]] = None,
                 disabled: Optional[pulumi.Input[bool]] = None,
                 filename: Optional[pulumi.Input[str]] = None,
                 github: Optional[pulumi.Input[pulumi.InputType['GitHubEventsConfigArgs']]] = None,
                 ignored_files: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]] = None,
                 included_files: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]] = None,
                 name: Optional[pulumi.Input[str]] = None,
                 project_id: Optional[pulumi.Input[str]] = None,
                 substitutions: Optional[pulumi.Input[Mapping[str, pulumi.Input[str]]]] = None,
                 tags: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]] = None,
                 trigger_id: Optional[pulumi.Input[str]] = None,
                 trigger_template: Optional[pulumi.Input[pulumi.InputType['RepoSourceArgs']]] = None,
                 __props__=None):
        """
        Creates a new `BuildTrigger`. This API is experimental.

        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[pulumi.InputType['BuildArgs']] build: Contents of the build template.
        :param pulumi.Input[str] description: Human-readable description of this trigger.
        :param pulumi.Input[bool] disabled: If true, the trigger will never automatically execute a build.
        :param pulumi.Input[str] filename: Path, from the source root, to the build configuration file (i.e. cloudbuild.yaml).
        :param pulumi.Input[pulumi.InputType['GitHubEventsConfigArgs']] github: GitHubEventsConfig describes the configuration of a trigger that creates a build whenever a GitHub event is received. Mutually exclusive with `trigger_template`.
        :param pulumi.Input[Sequence[pulumi.Input[str]]] ignored_files: ignored_files and included_files are file glob matches using https://golang.org/pkg/path/filepath/#Match extended with support for "**". If ignored_files and changed files are both empty, then they are not used to determine whether or not to trigger a build. If ignored_files is not empty, then we ignore any files that match any of the ignored_file globs. If the change has no files that are outside of the ignored_files globs, then we do not trigger a build.
        :param pulumi.Input[Sequence[pulumi.Input[str]]] included_files: If any of the files altered in the commit pass the ignored_files filter and included_files is empty, then as far as this filter is concerned, we should trigger the build. If any of the files altered in the commit pass the ignored_files filter and included_files is not empty, then we make sure that at least one of those files matches a included_files glob. If not, then we do not trigger a build.
        :param pulumi.Input[str] name: User-assigned name of the trigger. Must be unique within the project. Trigger names must meet the following requirements: + They must contain only alphanumeric characters and dashes. + They can be 1-64 characters long. + They must begin and end with an alphanumeric character.
        :param pulumi.Input[Mapping[str, pulumi.Input[str]]] substitutions: Substitutions for Build resource. The keys must match the following regular expression: `^_[A-Z0-9_]+$`.
        :param pulumi.Input[Sequence[pulumi.Input[str]]] tags: Tags for annotation of a `BuildTrigger`
        :param pulumi.Input[pulumi.InputType['RepoSourceArgs']] trigger_template: Template describing the types of source changes to trigger a build. Branch and tag names in trigger templates are interpreted as regular expressions. Any branch or tag change that matches that regular expression will trigger a build. Mutually exclusive with `github`.
        """
        ...
    @overload
    def __init__(__self__,
                 resource_name: str,
                 args: TriggerArgs,
                 opts: Optional[pulumi.ResourceOptions] = None):
        """
        Creates a new `BuildTrigger`. This API is experimental.

        :param str resource_name: The name of the resource.
        :param TriggerArgs args: The arguments to use to populate this resource's properties.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        ...
    def __init__(__self__, resource_name: str, *args, **kwargs):
        resource_args, opts = _utilities.get_resource_args_opts(TriggerArgs, pulumi.ResourceOptions, *args, **kwargs)
        if resource_args is not None:
            __self__._internal_init(resource_name, opts, **resource_args.__dict__)
        else:
            __self__._internal_init(resource_name, *args, **kwargs)

    def _internal_init(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 build: Optional[pulumi.Input[pulumi.InputType['BuildArgs']]] = None,
                 description: Optional[pulumi.Input[str]] = None,
                 disabled: Optional[pulumi.Input[bool]] = None,
                 filename: Optional[pulumi.Input[str]] = None,
                 github: Optional[pulumi.Input[pulumi.InputType['GitHubEventsConfigArgs']]] = None,
                 ignored_files: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]] = None,
                 included_files: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]] = None,
                 name: Optional[pulumi.Input[str]] = None,
                 project_id: Optional[pulumi.Input[str]] = None,
                 substitutions: Optional[pulumi.Input[Mapping[str, pulumi.Input[str]]]] = None,
                 tags: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]] = None,
                 trigger_id: Optional[pulumi.Input[str]] = None,
                 trigger_template: Optional[pulumi.Input[pulumi.InputType['RepoSourceArgs']]] = None,
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
            __props__ = TriggerArgs.__new__(TriggerArgs)

            __props__.__dict__["build"] = build
            __props__.__dict__["description"] = description
            __props__.__dict__["disabled"] = disabled
            __props__.__dict__["filename"] = filename
            __props__.__dict__["github"] = github
            __props__.__dict__["ignored_files"] = ignored_files
            __props__.__dict__["included_files"] = included_files
            __props__.__dict__["name"] = name
            if project_id is None and not opts.urn:
                raise TypeError("Missing required property 'project_id'")
            __props__.__dict__["project_id"] = project_id
            __props__.__dict__["substitutions"] = substitutions
            __props__.__dict__["tags"] = tags
            if trigger_id is None and not opts.urn:
                raise TypeError("Missing required property 'trigger_id'")
            __props__.__dict__["trigger_id"] = trigger_id
            __props__.__dict__["trigger_template"] = trigger_template
            __props__.__dict__["create_time"] = None
        super(Trigger, __self__).__init__(
            'gcp-native:cloudbuild/v1:Trigger',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name: str,
            id: pulumi.Input[str],
            opts: Optional[pulumi.ResourceOptions] = None) -> 'Trigger':
        """
        Get an existing Trigger resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param pulumi.Input[str] id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = TriggerArgs.__new__(TriggerArgs)

        __props__.__dict__["build"] = None
        __props__.__dict__["create_time"] = None
        __props__.__dict__["description"] = None
        __props__.__dict__["disabled"] = None
        __props__.__dict__["filename"] = None
        __props__.__dict__["github"] = None
        __props__.__dict__["ignored_files"] = None
        __props__.__dict__["included_files"] = None
        __props__.__dict__["name"] = None
        __props__.__dict__["substitutions"] = None
        __props__.__dict__["tags"] = None
        __props__.__dict__["trigger_template"] = None
        return Trigger(resource_name, opts=opts, __props__=__props__)

    @property
    @pulumi.getter
    def build(self) -> pulumi.Output['outputs.BuildResponse']:
        """
        Contents of the build template.
        """
        return pulumi.get(self, "build")

    @property
    @pulumi.getter(name="createTime")
    def create_time(self) -> pulumi.Output[str]:
        """
        Time when the trigger was created.
        """
        return pulumi.get(self, "create_time")

    @property
    @pulumi.getter
    def description(self) -> pulumi.Output[str]:
        """
        Human-readable description of this trigger.
        """
        return pulumi.get(self, "description")

    @property
    @pulumi.getter
    def disabled(self) -> pulumi.Output[bool]:
        """
        If true, the trigger will never automatically execute a build.
        """
        return pulumi.get(self, "disabled")

    @property
    @pulumi.getter
    def filename(self) -> pulumi.Output[str]:
        """
        Path, from the source root, to the build configuration file (i.e. cloudbuild.yaml).
        """
        return pulumi.get(self, "filename")

    @property
    @pulumi.getter
    def github(self) -> pulumi.Output['outputs.GitHubEventsConfigResponse']:
        """
        GitHubEventsConfig describes the configuration of a trigger that creates a build whenever a GitHub event is received. Mutually exclusive with `trigger_template`.
        """
        return pulumi.get(self, "github")

    @property
    @pulumi.getter(name="ignoredFiles")
    def ignored_files(self) -> pulumi.Output[Sequence[str]]:
        """
        ignored_files and included_files are file glob matches using https://golang.org/pkg/path/filepath/#Match extended with support for "**". If ignored_files and changed files are both empty, then they are not used to determine whether or not to trigger a build. If ignored_files is not empty, then we ignore any files that match any of the ignored_file globs. If the change has no files that are outside of the ignored_files globs, then we do not trigger a build.
        """
        return pulumi.get(self, "ignored_files")

    @property
    @pulumi.getter(name="includedFiles")
    def included_files(self) -> pulumi.Output[Sequence[str]]:
        """
        If any of the files altered in the commit pass the ignored_files filter and included_files is empty, then as far as this filter is concerned, we should trigger the build. If any of the files altered in the commit pass the ignored_files filter and included_files is not empty, then we make sure that at least one of those files matches a included_files glob. If not, then we do not trigger a build.
        """
        return pulumi.get(self, "included_files")

    @property
    @pulumi.getter
    def name(self) -> pulumi.Output[str]:
        """
        User-assigned name of the trigger. Must be unique within the project. Trigger names must meet the following requirements: + They must contain only alphanumeric characters and dashes. + They can be 1-64 characters long. + They must begin and end with an alphanumeric character.
        """
        return pulumi.get(self, "name")

    @property
    @pulumi.getter
    def substitutions(self) -> pulumi.Output[Mapping[str, str]]:
        """
        Substitutions for Build resource. The keys must match the following regular expression: `^_[A-Z0-9_]+$`.
        """
        return pulumi.get(self, "substitutions")

    @property
    @pulumi.getter
    def tags(self) -> pulumi.Output[Sequence[str]]:
        """
        Tags for annotation of a `BuildTrigger`
        """
        return pulumi.get(self, "tags")

    @property
    @pulumi.getter(name="triggerTemplate")
    def trigger_template(self) -> pulumi.Output['outputs.RepoSourceResponse']:
        """
        Template describing the types of source changes to trigger a build. Branch and tag names in trigger templates are interpreted as regular expressions. Any branch or tag change that matches that regular expression will trigger a build. Mutually exclusive with `github`.
        """
        return pulumi.get(self, "trigger_template")

