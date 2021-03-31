# coding=utf-8
# *** WARNING: this file was generated by the Pulumi SDK Generator. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import warnings
import pulumi
import pulumi.runtime
from typing import Any, Mapping, Optional, Sequence, Union
from ... import _utilities, _tables
from ._inputs import *

__all__ = ['DatasetConsentStoreConsent']


class DatasetConsentStoreConsent(pulumi.CustomResource):
    def __init__(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 consent_artifact: Optional[pulumi.Input[str]] = None,
                 consent_stores_id: Optional[pulumi.Input[str]] = None,
                 consents_id: Optional[pulumi.Input[str]] = None,
                 datasets_id: Optional[pulumi.Input[str]] = None,
                 expire_time: Optional[pulumi.Input[str]] = None,
                 locations_id: Optional[pulumi.Input[str]] = None,
                 metadata: Optional[pulumi.Input[Mapping[str, pulumi.Input[str]]]] = None,
                 name: Optional[pulumi.Input[str]] = None,
                 policies: Optional[pulumi.Input[Sequence[pulumi.Input[pulumi.InputType['GoogleCloudHealthcareV1ConsentPolicyArgs']]]]] = None,
                 projects_id: Optional[pulumi.Input[str]] = None,
                 revision_create_time: Optional[pulumi.Input[str]] = None,
                 revision_id: Optional[pulumi.Input[str]] = None,
                 state: Optional[pulumi.Input[str]] = None,
                 ttl: Optional[pulumi.Input[str]] = None,
                 user_id: Optional[pulumi.Input[str]] = None,
                 __props__=None,
                 __name__=None,
                 __opts__=None):
        """
        Creates a new Consent in the parent consent store.

        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] consent_artifact: Required. The resource name of the Consent artifact that contains proof of the end user's consent, of the form `projects/{project_id}/locations/{location_id}/datasets/{dataset_id}/consentStores/{consent_store_id}/consentArtifacts/{consent_artifact_id}`.
        :param pulumi.Input[str] expire_time: Timestamp in UTC of when this Consent is considered expired.
        :param pulumi.Input[Mapping[str, pulumi.Input[str]]] metadata: Optional. User-supplied key-value pairs used to organize Consent resources. Metadata keys must: - be between 1 and 63 characters long - have a UTF-8 encoding of maximum 128 bytes - begin with a letter - consist of up to 63 characters including lowercase letters, numeric characters, underscores, and dashes Metadata values must be: - be between 1 and 63 characters long - have a UTF-8 encoding of maximum 128 bytes - consist of up to 63 characters including lowercase letters, numeric characters, underscores, and dashes No more than 64 metadata entries can be associated with a given consent.
        :param pulumi.Input[str] name: Resource name of the Consent, of the form `projects/{project_id}/locations/{location_id}/datasets/{dataset_id}/consentStores/{consent_store_id}/consents/{consent_id}`. Cannot be changed after creation.
        :param pulumi.Input[Sequence[pulumi.Input[pulumi.InputType['GoogleCloudHealthcareV1ConsentPolicyArgs']]]] policies: Optional. Represents a user's consent in terms of the resources that can be accessed and under what conditions.
        :param pulumi.Input[str] revision_create_time: Output only. The timestamp that the revision was created.
        :param pulumi.Input[str] revision_id: Output only. The revision ID of the Consent. The format is an 8-character hexadecimal string. Refer to a specific revision of a Consent by appending `@{revision_id}` to the Consent's resource name.
        :param pulumi.Input[str] state: Required. Indicates the current state of this Consent.
        :param pulumi.Input[str] ttl: Input only. The time to live for this Consent from when it is created.
        :param pulumi.Input[str] user_id: Required. User's UUID provided by the client.
        """
        if __name__ is not None:
            warnings.warn("explicit use of __name__ is deprecated", DeprecationWarning)
            resource_name = __name__
        if __opts__ is not None:
            warnings.warn("explicit use of __opts__ is deprecated, use 'opts' instead", DeprecationWarning)
            opts = __opts__
        if opts is None:
            opts = pulumi.ResourceOptions()
        if not isinstance(opts, pulumi.ResourceOptions):
            raise TypeError('Expected resource options to be a ResourceOptions instance')
        if opts.version is None:
            opts.version = _utilities.get_version()
        if opts.id is None:
            if __props__ is not None:
                raise TypeError('__props__ is only valid when passed in combination with a valid opts.id to get an existing resource')
            __props__ = dict()

            __props__['consent_artifact'] = consent_artifact
            if consent_stores_id is None and not opts.urn:
                raise TypeError("Missing required property 'consent_stores_id'")
            __props__['consent_stores_id'] = consent_stores_id
            if consents_id is None and not opts.urn:
                raise TypeError("Missing required property 'consents_id'")
            __props__['consents_id'] = consents_id
            if datasets_id is None and not opts.urn:
                raise TypeError("Missing required property 'datasets_id'")
            __props__['datasets_id'] = datasets_id
            __props__['expire_time'] = expire_time
            if locations_id is None and not opts.urn:
                raise TypeError("Missing required property 'locations_id'")
            __props__['locations_id'] = locations_id
            __props__['metadata'] = metadata
            __props__['name'] = name
            __props__['policies'] = policies
            if projects_id is None and not opts.urn:
                raise TypeError("Missing required property 'projects_id'")
            __props__['projects_id'] = projects_id
            __props__['revision_create_time'] = revision_create_time
            __props__['revision_id'] = revision_id
            __props__['state'] = state
            __props__['ttl'] = ttl
            __props__['user_id'] = user_id
        super(DatasetConsentStoreConsent, __self__).__init__(
            'google-cloud:healthcare/v1:DatasetConsentStoreConsent',
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

        __props__ = dict()

        return DatasetConsentStoreConsent(resource_name, opts=opts, __props__=__props__)

    def translate_output_property(self, prop):
        return _tables.CAMEL_TO_SNAKE_CASE_TABLE.get(prop) or prop

    def translate_input_property(self, prop):
        return _tables.SNAKE_TO_CAMEL_CASE_TABLE.get(prop) or prop

