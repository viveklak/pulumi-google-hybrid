# coding=utf-8
# *** WARNING: this file was generated by the Pulumi SDK Generator. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import warnings
import pulumi
import pulumi.runtime
from typing import Any, Mapping, Optional, Sequence, Union
from ... import _utilities, _tables
from ._inputs import *

__all__ = ['BucketObject']


class BucketObject(pulumi.CustomResource):
    def __init__(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 acl: Optional[pulumi.Input[Sequence[pulumi.Input[pulumi.InputType['ObjectAccessControlArgs']]]]] = None,
                 bucket: Optional[pulumi.Input[str]] = None,
                 cache_control: Optional[pulumi.Input[str]] = None,
                 component_count: Optional[pulumi.Input[int]] = None,
                 content_disposition: Optional[pulumi.Input[str]] = None,
                 content_encoding: Optional[pulumi.Input[str]] = None,
                 content_language: Optional[pulumi.Input[str]] = None,
                 content_type: Optional[pulumi.Input[str]] = None,
                 crc32c: Optional[pulumi.Input[str]] = None,
                 custom_time: Optional[pulumi.Input[str]] = None,
                 customer_encryption: Optional[pulumi.Input[Mapping[str, pulumi.Input[str]]]] = None,
                 etag: Optional[pulumi.Input[str]] = None,
                 event_based_hold: Optional[pulumi.Input[bool]] = None,
                 generation: Optional[pulumi.Input[str]] = None,
                 id: Optional[pulumi.Input[str]] = None,
                 kind: Optional[pulumi.Input[str]] = None,
                 kms_key_name: Optional[pulumi.Input[str]] = None,
                 md5_hash: Optional[pulumi.Input[str]] = None,
                 media_link: Optional[pulumi.Input[str]] = None,
                 metadata: Optional[pulumi.Input[Mapping[str, pulumi.Input[str]]]] = None,
                 metageneration: Optional[pulumi.Input[str]] = None,
                 name: Optional[pulumi.Input[str]] = None,
                 object: Optional[pulumi.Input[str]] = None,
                 owner: Optional[pulumi.Input[Mapping[str, pulumi.Input[str]]]] = None,
                 retention_expiration_time: Optional[pulumi.Input[str]] = None,
                 self_link: Optional[pulumi.Input[str]] = None,
                 size: Optional[pulumi.Input[str]] = None,
                 source: Optional[pulumi.Input[Union[pulumi.Asset, pulumi.Archive]]] = None,
                 storage_class: Optional[pulumi.Input[str]] = None,
                 temporary_hold: Optional[pulumi.Input[bool]] = None,
                 time_created: Optional[pulumi.Input[str]] = None,
                 time_deleted: Optional[pulumi.Input[str]] = None,
                 time_storage_class_updated: Optional[pulumi.Input[str]] = None,
                 updated: Optional[pulumi.Input[str]] = None,
                 __props__=None,
                 __name__=None,
                 __opts__=None):
        """
        Stores a new object and metadata.

        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[Sequence[pulumi.Input[pulumi.InputType['ObjectAccessControlArgs']]]] acl: Access controls on the object.
        :param pulumi.Input[str] bucket: The name of the bucket containing this object.
        :param pulumi.Input[str] cache_control: Cache-Control directive for the object data. If omitted, and the object is accessible to all anonymous users, the default will be public, max-age=3600.
        :param pulumi.Input[int] component_count: Number of underlying components that make up this object. Components are accumulated by compose operations.
        :param pulumi.Input[str] content_disposition: Content-Disposition of the object data.
        :param pulumi.Input[str] content_encoding: Content-Encoding of the object data.
        :param pulumi.Input[str] content_language: Content-Language of the object data.
        :param pulumi.Input[str] content_type: Content-Type of the object data. If an object is stored without a Content-Type, it is served as application/octet-stream.
        :param pulumi.Input[str] crc32c: CRC32c checksum, as described in RFC 4960, Appendix B; encoded using base64 in big-endian byte order. For more information about using the CRC32c checksum, see Hashes and ETags: Best Practices.
        :param pulumi.Input[str] custom_time: A timestamp in RFC 3339 format specified by the user for an object.
        :param pulumi.Input[Mapping[str, pulumi.Input[str]]] customer_encryption: Metadata of customer-supplied encryption key, if the object is encrypted by such a key.
        :param pulumi.Input[str] etag: HTTP 1.1 Entity tag for the object.
        :param pulumi.Input[bool] event_based_hold: Whether an object is under event-based hold. Event-based hold is a way to retain objects until an event occurs, which is signified by the hold's release (i.e. this value is set to false). After being released (set to false), such objects will be subject to bucket-level retention (if any). One sample use case of this flag is for banks to hold loan documents for at least 3 years after loan is paid in full. Here, bucket-level retention is 3 years and the event is the loan being paid in full. In this example, these objects will be held intact for any number of years until the event has occurred (event-based hold on the object is released) and then 3 more years after that. That means retention duration of the objects begins from the moment event-based hold transitioned from true to false.
        :param pulumi.Input[str] generation: The content generation of this object. Used for object versioning.
        :param pulumi.Input[str] id: The ID of the object, including the bucket name, object name, and generation number.
        :param pulumi.Input[str] kind: The kind of item this is. For objects, this is always storage#object.
        :param pulumi.Input[str] kms_key_name: Not currently supported. Specifying the parameter causes the request to fail with status code 400 - Bad Request.
        :param pulumi.Input[str] md5_hash: MD5 hash of the data; encoded using base64. For more information about using the MD5 hash, see Hashes and ETags: Best Practices.
        :param pulumi.Input[str] media_link: Media download link.
        :param pulumi.Input[Mapping[str, pulumi.Input[str]]] metadata: User-provided metadata, in key/value pairs.
        :param pulumi.Input[str] metageneration: The version of the metadata for this object at this generation. Used for preconditions and for detecting changes in metadata. A metageneration number is only meaningful in the context of a particular generation of a particular object.
        :param pulumi.Input[str] name: The name of the object. Required if not specified by URL parameter.
        :param pulumi.Input[Mapping[str, pulumi.Input[str]]] owner: The owner of the object. This will always be the uploader of the object.
        :param pulumi.Input[str] retention_expiration_time: A server-determined value that specifies the earliest time that the object's retention period expires. This value is in RFC 3339 format. Note 1: This field is not provided for objects with an active event-based hold, since retention expiration is unknown until the hold is removed. Note 2: This value can be provided even when temporary hold is set (so that the user can reason about policy without having to first unset the temporary hold).
        :param pulumi.Input[str] self_link: The link to this object.
        :param pulumi.Input[str] size: Content-Length of the data in bytes.
        :param pulumi.Input[str] storage_class: Storage class of the object.
        :param pulumi.Input[bool] temporary_hold: Whether an object is under temporary hold. While this flag is set to true, the object is protected against deletion and overwrites. A common use case of this flag is regulatory investigations where objects need to be retained while the investigation is ongoing. Note that unlike event-based hold, temporary hold does not impact retention expiration time of an object.
        :param pulumi.Input[str] time_created: The creation time of the object in RFC 3339 format.
        :param pulumi.Input[str] time_deleted: The deletion time of the object in RFC 3339 format. Will be returned if and only if this version of the object has been deleted.
        :param pulumi.Input[str] time_storage_class_updated: The time at which the object's storage class was last changed. When the object is initially created, it will be set to timeCreated.
        :param pulumi.Input[str] updated: The modification time of the object metadata in RFC 3339 format.
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

            __props__['acl'] = acl
            if bucket is None and not opts.urn:
                raise TypeError("Missing required property 'bucket'")
            __props__['bucket'] = bucket
            __props__['cache_control'] = cache_control
            __props__['component_count'] = component_count
            __props__['content_disposition'] = content_disposition
            __props__['content_encoding'] = content_encoding
            __props__['content_language'] = content_language
            __props__['content_type'] = content_type
            __props__['crc32c'] = crc32c
            __props__['custom_time'] = custom_time
            __props__['customer_encryption'] = customer_encryption
            __props__['etag'] = etag
            __props__['event_based_hold'] = event_based_hold
            __props__['generation'] = generation
            __props__['id'] = id
            __props__['kind'] = kind
            __props__['kms_key_name'] = kms_key_name
            __props__['md5_hash'] = md5_hash
            __props__['media_link'] = media_link
            __props__['metadata'] = metadata
            __props__['metageneration'] = metageneration
            __props__['name'] = name
            if object is None and not opts.urn:
                raise TypeError("Missing required property 'object'")
            __props__['object'] = object
            __props__['owner'] = owner
            __props__['retention_expiration_time'] = retention_expiration_time
            __props__['self_link'] = self_link
            __props__['size'] = size
            __props__['source'] = source
            __props__['storage_class'] = storage_class
            __props__['temporary_hold'] = temporary_hold
            __props__['time_created'] = time_created
            __props__['time_deleted'] = time_deleted
            __props__['time_storage_class_updated'] = time_storage_class_updated
            __props__['updated'] = updated
        super(BucketObject, __self__).__init__(
            'google-cloud:storage/v1:BucketObject',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name: str,
            id: pulumi.Input[str],
            opts: Optional[pulumi.ResourceOptions] = None) -> 'BucketObject':
        """
        Get an existing BucketObject resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param pulumi.Input[str] id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = dict()

        return BucketObject(resource_name, opts=opts, __props__=__props__)

    def translate_output_property(self, prop):
        return _tables.CAMEL_TO_SNAKE_CASE_TABLE.get(prop) or prop

    def translate_input_property(self, prop):
        return _tables.SNAKE_TO_CAMEL_CASE_TABLE.get(prop) or prop

