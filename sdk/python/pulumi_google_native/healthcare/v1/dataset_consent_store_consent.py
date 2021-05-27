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

__all__ = ['DatasetConsentStoreConsentArgs', 'DatasetConsentStoreConsent']

@pulumi.input_type
class DatasetConsentStoreConsentArgs:
    def __init__(__self__, *,
                 consent_store_id: pulumi.Input[str],
                 dataset_id: pulumi.Input[str],
                 location: pulumi.Input[str],
                 project: pulumi.Input[str],
                 consent_artifact: Optional[pulumi.Input[str]] = None,
                 expire_time: Optional[pulumi.Input[str]] = None,
                 metadata: Optional[pulumi.Input[Mapping[str, pulumi.Input[str]]]] = None,
                 name: Optional[pulumi.Input[str]] = None,
                 policies: Optional[pulumi.Input[Sequence[pulumi.Input['GoogleCloudHealthcareV1ConsentPolicyArgs']]]] = None,
                 state: Optional[pulumi.Input[str]] = None,
                 ttl: Optional[pulumi.Input[str]] = None,
                 user_id: Optional[pulumi.Input[str]] = None):
        """
        The set of arguments for constructing a DatasetConsentStoreConsent resource.
        :param pulumi.Input[str] consent_artifact: Required. The resource name of the Consent artifact that contains proof of the end user's consent, of the form `projects/{project_id}/locations/{location_id}/datasets/{dataset_id}/consentStores/{consent_store_id}/consentArtifacts/{consent_artifact_id}`.
        :param pulumi.Input[str] expire_time: Timestamp in UTC of when this Consent is considered expired.
        :param pulumi.Input[Mapping[str, pulumi.Input[str]]] metadata: Optional. User-supplied key-value pairs used to organize Consent resources. Metadata keys must: - be between 1 and 63 characters long - have a UTF-8 encoding of maximum 128 bytes - begin with a letter - consist of up to 63 characters including lowercase letters, numeric characters, underscores, and dashes Metadata values must be: - be between 1 and 63 characters long - have a UTF-8 encoding of maximum 128 bytes - consist of up to 63 characters including lowercase letters, numeric characters, underscores, and dashes No more than 64 metadata entries can be associated with a given consent.
        :param pulumi.Input[str] name: Resource name of the Consent, of the form `projects/{project_id}/locations/{location_id}/datasets/{dataset_id}/consentStores/{consent_store_id}/consents/{consent_id}`. Cannot be changed after creation.
        :param pulumi.Input[Sequence[pulumi.Input['GoogleCloudHealthcareV1ConsentPolicyArgs']]] policies: Optional. Represents a user's consent in terms of the resources that can be accessed and under what conditions.
        :param pulumi.Input[str] state: Required. Indicates the current state of this Consent.
        :param pulumi.Input[str] ttl: Input only. The time to live for this Consent from when it is created.
        :param pulumi.Input[str] user_id: Required. User's UUID provided by the client.
        """
        pulumi.set(__self__, "consent_store_id", consent_store_id)
        pulumi.set(__self__, "dataset_id", dataset_id)
        pulumi.set(__self__, "location", location)
        pulumi.set(__self__, "project", project)
        if consent_artifact is not None:
            pulumi.set(__self__, "consent_artifact", consent_artifact)
        if expire_time is not None:
            pulumi.set(__self__, "expire_time", expire_time)
        if metadata is not None:
            pulumi.set(__self__, "metadata", metadata)
        if name is not None:
            pulumi.set(__self__, "name", name)
        if policies is not None:
            pulumi.set(__self__, "policies", policies)
        if state is not None:
            pulumi.set(__self__, "state", state)
        if ttl is not None:
            pulumi.set(__self__, "ttl", ttl)
        if user_id is not None:
            pulumi.set(__self__, "user_id", user_id)

    @property
    @pulumi.getter(name="consentStoreId")
    def consent_store_id(self) -> pulumi.Input[str]:
        return pulumi.get(self, "consent_store_id")

    @consent_store_id.setter
    def consent_store_id(self, value: pulumi.Input[str]):
        pulumi.set(self, "consent_store_id", value)

    @property
    @pulumi.getter(name="datasetId")
    def dataset_id(self) -> pulumi.Input[str]:
        return pulumi.get(self, "dataset_id")

    @dataset_id.setter
    def dataset_id(self, value: pulumi.Input[str]):
        pulumi.set(self, "dataset_id", value)

    @property
    @pulumi.getter
    def location(self) -> pulumi.Input[str]:
        return pulumi.get(self, "location")

    @location.setter
    def location(self, value: pulumi.Input[str]):
        pulumi.set(self, "location", value)

    @property
    @pulumi.getter
    def project(self) -> pulumi.Input[str]:
        return pulumi.get(self, "project")

    @project.setter
    def project(self, value: pulumi.Input[str]):
        pulumi.set(self, "project", value)

    @property
    @pulumi.getter(name="consentArtifact")
    def consent_artifact(self) -> Optional[pulumi.Input[str]]:
        """
        Required. The resource name of the Consent artifact that contains proof of the end user's consent, of the form `projects/{project_id}/locations/{location_id}/datasets/{dataset_id}/consentStores/{consent_store_id}/consentArtifacts/{consent_artifact_id}`.
        """
        return pulumi.get(self, "consent_artifact")

    @consent_artifact.setter
    def consent_artifact(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "consent_artifact", value)

    @property
    @pulumi.getter(name="expireTime")
    def expire_time(self) -> Optional[pulumi.Input[str]]:
        """
        Timestamp in UTC of when this Consent is considered expired.
        """
        return pulumi.get(self, "expire_time")

    @expire_time.setter
    def expire_time(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "expire_time", value)

    @property
    @pulumi.getter
    def metadata(self) -> Optional[pulumi.Input[Mapping[str, pulumi.Input[str]]]]:
        """
        Optional. User-supplied key-value pairs used to organize Consent resources. Metadata keys must: - be between 1 and 63 characters long - have a UTF-8 encoding of maximum 128 bytes - begin with a letter - consist of up to 63 characters including lowercase letters, numeric characters, underscores, and dashes Metadata values must be: - be between 1 and 63 characters long - have a UTF-8 encoding of maximum 128 bytes - consist of up to 63 characters including lowercase letters, numeric characters, underscores, and dashes No more than 64 metadata entries can be associated with a given consent.
        """
        return pulumi.get(self, "metadata")

    @metadata.setter
    def metadata(self, value: Optional[pulumi.Input[Mapping[str, pulumi.Input[str]]]]):
        pulumi.set(self, "metadata", value)

    @property
    @pulumi.getter
    def name(self) -> Optional[pulumi.Input[str]]:
        """
        Resource name of the Consent, of the form `projects/{project_id}/locations/{location_id}/datasets/{dataset_id}/consentStores/{consent_store_id}/consents/{consent_id}`. Cannot be changed after creation.
        """
        return pulumi.get(self, "name")

    @name.setter
    def name(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "name", value)

    @property
    @pulumi.getter
    def policies(self) -> Optional[pulumi.Input[Sequence[pulumi.Input['GoogleCloudHealthcareV1ConsentPolicyArgs']]]]:
        """
        Optional. Represents a user's consent in terms of the resources that can be accessed and under what conditions.
        """
        return pulumi.get(self, "policies")

    @policies.setter
    def policies(self, value: Optional[pulumi.Input[Sequence[pulumi.Input['GoogleCloudHealthcareV1ConsentPolicyArgs']]]]):
        pulumi.set(self, "policies", value)

    @property
    @pulumi.getter
    def state(self) -> Optional[pulumi.Input[str]]:
        """
        Required. Indicates the current state of this Consent.
        """
        return pulumi.get(self, "state")

    @state.setter
    def state(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "state", value)

    @property
    @pulumi.getter
    def ttl(self) -> Optional[pulumi.Input[str]]:
        """
        Input only. The time to live for this Consent from when it is created.
        """
        return pulumi.get(self, "ttl")

    @ttl.setter
    def ttl(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "ttl", value)

    @property
    @pulumi.getter(name="userId")
    def user_id(self) -> Optional[pulumi.Input[str]]:
        """
        Required. User's UUID provided by the client.
        """
        return pulumi.get(self, "user_id")

    @user_id.setter
    def user_id(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "user_id", value)


class DatasetConsentStoreConsent(pulumi.CustomResource):
    @overload
    def __init__(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 consent_artifact: Optional[pulumi.Input[str]] = None,
                 consent_store_id: Optional[pulumi.Input[str]] = None,
                 dataset_id: Optional[pulumi.Input[str]] = None,
                 expire_time: Optional[pulumi.Input[str]] = None,
                 location: Optional[pulumi.Input[str]] = None,
                 metadata: Optional[pulumi.Input[Mapping[str, pulumi.Input[str]]]] = None,
                 name: Optional[pulumi.Input[str]] = None,
                 policies: Optional[pulumi.Input[Sequence[pulumi.Input[pulumi.InputType['GoogleCloudHealthcareV1ConsentPolicyArgs']]]]] = None,
                 project: Optional[pulumi.Input[str]] = None,
                 state: Optional[pulumi.Input[str]] = None,
                 ttl: Optional[pulumi.Input[str]] = None,
                 user_id: Optional[pulumi.Input[str]] = None,
                 __props__=None):
        """
        Creates a new Consent in the parent consent store.

        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] consent_artifact: Required. The resource name of the Consent artifact that contains proof of the end user's consent, of the form `projects/{project_id}/locations/{location_id}/datasets/{dataset_id}/consentStores/{consent_store_id}/consentArtifacts/{consent_artifact_id}`.
        :param pulumi.Input[str] expire_time: Timestamp in UTC of when this Consent is considered expired.
        :param pulumi.Input[Mapping[str, pulumi.Input[str]]] metadata: Optional. User-supplied key-value pairs used to organize Consent resources. Metadata keys must: - be between 1 and 63 characters long - have a UTF-8 encoding of maximum 128 bytes - begin with a letter - consist of up to 63 characters including lowercase letters, numeric characters, underscores, and dashes Metadata values must be: - be between 1 and 63 characters long - have a UTF-8 encoding of maximum 128 bytes - consist of up to 63 characters including lowercase letters, numeric characters, underscores, and dashes No more than 64 metadata entries can be associated with a given consent.
        :param pulumi.Input[str] name: Resource name of the Consent, of the form `projects/{project_id}/locations/{location_id}/datasets/{dataset_id}/consentStores/{consent_store_id}/consents/{consent_id}`. Cannot be changed after creation.
        :param pulumi.Input[Sequence[pulumi.Input[pulumi.InputType['GoogleCloudHealthcareV1ConsentPolicyArgs']]]] policies: Optional. Represents a user's consent in terms of the resources that can be accessed and under what conditions.
        :param pulumi.Input[str] state: Required. Indicates the current state of this Consent.
        :param pulumi.Input[str] ttl: Input only. The time to live for this Consent from when it is created.
        :param pulumi.Input[str] user_id: Required. User's UUID provided by the client.
        """
        ...
    @overload
    def __init__(__self__,
                 resource_name: str,
                 args: DatasetConsentStoreConsentArgs,
                 opts: Optional[pulumi.ResourceOptions] = None):
        """
        Creates a new Consent in the parent consent store.

        :param str resource_name: The name of the resource.
        :param DatasetConsentStoreConsentArgs args: The arguments to use to populate this resource's properties.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        ...
    def __init__(__self__, resource_name: str, *args, **kwargs):
        resource_args, opts = _utilities.get_resource_args_opts(DatasetConsentStoreConsentArgs, pulumi.ResourceOptions, *args, **kwargs)
        if resource_args is not None:
            __self__._internal_init(resource_name, opts, **resource_args.__dict__)
        else:
            __self__._internal_init(resource_name, *args, **kwargs)

    def _internal_init(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 consent_artifact: Optional[pulumi.Input[str]] = None,
                 consent_store_id: Optional[pulumi.Input[str]] = None,
                 dataset_id: Optional[pulumi.Input[str]] = None,
                 expire_time: Optional[pulumi.Input[str]] = None,
                 location: Optional[pulumi.Input[str]] = None,
                 metadata: Optional[pulumi.Input[Mapping[str, pulumi.Input[str]]]] = None,
                 name: Optional[pulumi.Input[str]] = None,
                 policies: Optional[pulumi.Input[Sequence[pulumi.Input[pulumi.InputType['GoogleCloudHealthcareV1ConsentPolicyArgs']]]]] = None,
                 project: Optional[pulumi.Input[str]] = None,
                 state: Optional[pulumi.Input[str]] = None,
                 ttl: Optional[pulumi.Input[str]] = None,
                 user_id: Optional[pulumi.Input[str]] = None,
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
            __props__ = DatasetConsentStoreConsentArgs.__new__(DatasetConsentStoreConsentArgs)

            __props__.__dict__["consent_artifact"] = consent_artifact
            if consent_store_id is None and not opts.urn:
                raise TypeError("Missing required property 'consent_store_id'")
            __props__.__dict__["consent_store_id"] = consent_store_id
            if dataset_id is None and not opts.urn:
                raise TypeError("Missing required property 'dataset_id'")
            __props__.__dict__["dataset_id"] = dataset_id
            __props__.__dict__["expire_time"] = expire_time
            if location is None and not opts.urn:
                raise TypeError("Missing required property 'location'")
            __props__.__dict__["location"] = location
            __props__.__dict__["metadata"] = metadata
            __props__.__dict__["name"] = name
            __props__.__dict__["policies"] = policies
            if project is None and not opts.urn:
                raise TypeError("Missing required property 'project'")
            __props__.__dict__["project"] = project
            __props__.__dict__["state"] = state
            __props__.__dict__["ttl"] = ttl
            __props__.__dict__["user_id"] = user_id
            __props__.__dict__["revision_create_time"] = None
            __props__.__dict__["revision_id"] = None
        super(DatasetConsentStoreConsent, __self__).__init__(
            'google-native:healthcare/v1:DatasetConsentStoreConsent',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name: str,
            id: pulumi.Input[str],
            opts: Optional[pulumi.ResourceOptions] = None) -> 'DatasetConsentStoreConsent':
        """
        Get an existing DatasetConsentStoreConsent resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param pulumi.Input[str] id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = DatasetConsentStoreConsentArgs.__new__(DatasetConsentStoreConsentArgs)

        __props__.__dict__["consent_artifact"] = None
        __props__.__dict__["expire_time"] = None
        __props__.__dict__["metadata"] = None
        __props__.__dict__["name"] = None
        __props__.__dict__["policies"] = None
        __props__.__dict__["revision_create_time"] = None
        __props__.__dict__["revision_id"] = None
        __props__.__dict__["state"] = None
        __props__.__dict__["ttl"] = None
        __props__.__dict__["user_id"] = None
        return DatasetConsentStoreConsent(resource_name, opts=opts, __props__=__props__)

    @property
    @pulumi.getter(name="consentArtifact")
    def consent_artifact(self) -> pulumi.Output[str]:
        """
        Required. The resource name of the Consent artifact that contains proof of the end user's consent, of the form `projects/{project_id}/locations/{location_id}/datasets/{dataset_id}/consentStores/{consent_store_id}/consentArtifacts/{consent_artifact_id}`.
        """
        return pulumi.get(self, "consent_artifact")

    @property
    @pulumi.getter(name="expireTime")
    def expire_time(self) -> pulumi.Output[str]:
        """
        Timestamp in UTC of when this Consent is considered expired.
        """
        return pulumi.get(self, "expire_time")

    @property
    @pulumi.getter
    def metadata(self) -> pulumi.Output[Mapping[str, str]]:
        """
        Optional. User-supplied key-value pairs used to organize Consent resources. Metadata keys must: - be between 1 and 63 characters long - have a UTF-8 encoding of maximum 128 bytes - begin with a letter - consist of up to 63 characters including lowercase letters, numeric characters, underscores, and dashes Metadata values must be: - be between 1 and 63 characters long - have a UTF-8 encoding of maximum 128 bytes - consist of up to 63 characters including lowercase letters, numeric characters, underscores, and dashes No more than 64 metadata entries can be associated with a given consent.
        """
        return pulumi.get(self, "metadata")

    @property
    @pulumi.getter
    def name(self) -> pulumi.Output[str]:
        """
        Resource name of the Consent, of the form `projects/{project_id}/locations/{location_id}/datasets/{dataset_id}/consentStores/{consent_store_id}/consents/{consent_id}`. Cannot be changed after creation.
        """
        return pulumi.get(self, "name")

    @property
    @pulumi.getter
    def policies(self) -> pulumi.Output[Sequence['outputs.GoogleCloudHealthcareV1ConsentPolicyResponse']]:
        """
        Optional. Represents a user's consent in terms of the resources that can be accessed and under what conditions.
        """
        return pulumi.get(self, "policies")

    @property
    @pulumi.getter(name="revisionCreateTime")
    def revision_create_time(self) -> pulumi.Output[str]:
        """
        The timestamp that the revision was created.
        """
        return pulumi.get(self, "revision_create_time")

    @property
    @pulumi.getter(name="revisionId")
    def revision_id(self) -> pulumi.Output[str]:
        """
        The revision ID of the Consent. The format is an 8-character hexadecimal string. Refer to a specific revision of a Consent by appending `@{revision_id}` to the Consent's resource name.
        """
        return pulumi.get(self, "revision_id")

    @property
    @pulumi.getter
    def state(self) -> pulumi.Output[str]:
        """
        Required. Indicates the current state of this Consent.
        """
        return pulumi.get(self, "state")

    @property
    @pulumi.getter
    def ttl(self) -> pulumi.Output[str]:
        """
        Input only. The time to live for this Consent from when it is created.
        """
        return pulumi.get(self, "ttl")

    @property
    @pulumi.getter(name="userId")
    def user_id(self) -> pulumi.Output[str]:
        """
        Required. User's UUID provided by the client.
        """
        return pulumi.get(self, "user_id")

