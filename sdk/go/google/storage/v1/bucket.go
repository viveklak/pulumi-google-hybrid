// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package v1

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Creates a new bucket.
type Bucket struct {
	pulumi.CustomResourceState

	// Access controls on the bucket.
	Acl BucketAccessControlResponseArrayOutput `pulumi:"acl"`
	// The bucket's billing configuration.
	Billing BucketBillingResponseOutput `pulumi:"billing"`
	// The bucket's Cross-Origin Resource Sharing (CORS) configuration.
	Cors BucketCorsItemResponseArrayOutput `pulumi:"cors"`
	// The default value for event-based hold on newly created objects in this bucket. Event-based hold is a way to retain objects indefinitely until an event occurs, signified by the hold's release. After being released, such objects will be subject to bucket-level retention (if any). One sample use case of this flag is for banks to hold loan documents for at least 3 years after loan is paid in full. Here, bucket-level retention is 3 years and the event is loan being paid in full. In this example, these objects will be held intact for any number of years until the event has occurred (event-based hold on the object is released) and then 3 more years after that. That means retention duration of the objects begins from the moment event-based hold transitioned from true to false. Objects under event-based hold cannot be deleted, overwritten or archived until the hold is removed.
	DefaultEventBasedHold pulumi.BoolOutput `pulumi:"defaultEventBasedHold"`
	// Default access controls to apply to new objects when no ACL is provided.
	DefaultObjectAcl ObjectAccessControlResponseArrayOutput `pulumi:"defaultObjectAcl"`
	// Encryption configuration for a bucket.
	Encryption BucketEncryptionResponseOutput `pulumi:"encryption"`
	// HTTP 1.1 Entity tag for the bucket.
	Etag pulumi.StringOutput `pulumi:"etag"`
	// The bucket's IAM configuration.
	IamConfiguration BucketIamConfigurationResponseOutput `pulumi:"iamConfiguration"`
	// The kind of item this is. For buckets, this is always storage#bucket.
	Kind pulumi.StringOutput `pulumi:"kind"`
	// User-provided labels, in key/value pairs.
	Labels pulumi.StringMapOutput `pulumi:"labels"`
	// The bucket's lifecycle configuration. See lifecycle management for more information.
	Lifecycle BucketLifecycleResponseOutput `pulumi:"lifecycle"`
	// The location of the bucket. Object data for objects in the bucket resides in physical storage within this region. Defaults to US. See the developer's guide for the authoritative list.
	Location pulumi.StringOutput `pulumi:"location"`
	// The type of the bucket location.
	LocationType pulumi.StringOutput `pulumi:"locationType"`
	// The bucket's logging configuration, which defines the destination bucket and optional name prefix for the current bucket's logs.
	Logging BucketLoggingResponseOutput `pulumi:"logging"`
	// The metadata generation of this bucket.
	Metageneration pulumi.StringOutput `pulumi:"metageneration"`
	// The name of the bucket.
	Name pulumi.StringOutput `pulumi:"name"`
	// The owner of the bucket. This is always the project team's owner group.
	Owner BucketOwnerResponseOutput `pulumi:"owner"`
	// The project number of the project the bucket belongs to.
	ProjectNumber pulumi.StringOutput `pulumi:"projectNumber"`
	// The bucket's retention policy. The retention policy enforces a minimum retention time for all objects contained in the bucket, based on their creation time. Any attempt to overwrite or delete objects younger than the retention period will result in a PERMISSION_DENIED error. An unlocked retention policy can be modified or removed from the bucket via a storage.buckets.update operation. A locked retention policy cannot be removed or shortened in duration for the lifetime of the bucket. Attempting to remove or decrease period of a locked retention policy will result in a PERMISSION_DENIED error.
	RetentionPolicy BucketRetentionPolicyResponseOutput `pulumi:"retentionPolicy"`
	// Reserved for future use.
	SatisfiesPZS pulumi.BoolOutput `pulumi:"satisfiesPZS"`
	// The URI of this bucket.
	SelfLink pulumi.StringOutput `pulumi:"selfLink"`
	// The bucket's default storage class, used whenever no storageClass is specified for a newly-created object. This defines how objects in the bucket are stored and determines the SLA and the cost of storage. Values include MULTI_REGIONAL, REGIONAL, STANDARD, NEARLINE, COLDLINE, ARCHIVE, and DURABLE_REDUCED_AVAILABILITY. If this value is not specified when the bucket is created, it will default to STANDARD. For more information, see storage classes.
	StorageClass pulumi.StringOutput `pulumi:"storageClass"`
	// The creation time of the bucket in RFC 3339 format.
	TimeCreated pulumi.StringOutput `pulumi:"timeCreated"`
	// The modification time of the bucket in RFC 3339 format.
	Updated pulumi.StringOutput `pulumi:"updated"`
	// The bucket's versioning configuration.
	Versioning BucketVersioningResponseOutput `pulumi:"versioning"`
	// The bucket's website configuration, controlling how the service behaves when accessing bucket contents as a web site. See the Static Website Examples for more information.
	Website BucketWebsiteResponseOutput `pulumi:"website"`
	// The zone or zones from which the bucket is intended to use zonal quota. Requests for data from outside the specified affinities are still allowed but won't be able to use zonal quota. The zone or zones need to be within the bucket location otherwise the requests will fail with a 400 Bad Request response.
	ZoneAffinity pulumi.StringArrayOutput `pulumi:"zoneAffinity"`
}

// NewBucket registers a new resource with the given unique name, arguments, and options.
func NewBucket(ctx *pulumi.Context,
	name string, args *BucketArgs, opts ...pulumi.ResourceOption) (*Bucket, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.Project == nil {
		return nil, errors.New("invalid value for required argument 'Project'")
	}
	var resource Bucket
	err := ctx.RegisterResource("google-native:storage/v1:Bucket", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetBucket gets an existing Bucket resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetBucket(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *BucketState, opts ...pulumi.ResourceOption) (*Bucket, error) {
	var resource Bucket
	err := ctx.ReadResource("google-native:storage/v1:Bucket", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering Bucket resources.
type bucketState struct {
}

type BucketState struct {
}

func (BucketState) ElementType() reflect.Type {
	return reflect.TypeOf((*bucketState)(nil)).Elem()
}

type bucketArgs struct {
	// Access controls on the bucket.
	Acl []BucketAccessControlType `pulumi:"acl"`
	// The bucket's billing configuration.
	Billing *BucketBilling `pulumi:"billing"`
	// The bucket's Cross-Origin Resource Sharing (CORS) configuration.
	Cors []BucketCorsItem `pulumi:"cors"`
	// The default value for event-based hold on newly created objects in this bucket. Event-based hold is a way to retain objects indefinitely until an event occurs, signified by the hold's release. After being released, such objects will be subject to bucket-level retention (if any). One sample use case of this flag is for banks to hold loan documents for at least 3 years after loan is paid in full. Here, bucket-level retention is 3 years and the event is loan being paid in full. In this example, these objects will be held intact for any number of years until the event has occurred (event-based hold on the object is released) and then 3 more years after that. That means retention duration of the objects begins from the moment event-based hold transitioned from true to false. Objects under event-based hold cannot be deleted, overwritten or archived until the hold is removed.
	DefaultEventBasedHold *bool `pulumi:"defaultEventBasedHold"`
	// Default access controls to apply to new objects when no ACL is provided.
	DefaultObjectAcl []ObjectAccessControlType `pulumi:"defaultObjectAcl"`
	// Encryption configuration for a bucket.
	Encryption *BucketEncryption `pulumi:"encryption"`
	// HTTP 1.1 Entity tag for the bucket.
	Etag *string `pulumi:"etag"`
	// The bucket's IAM configuration.
	IamConfiguration *BucketIamConfiguration `pulumi:"iamConfiguration"`
	// The ID of the bucket. For buckets, the id and name properties are the same.
	Id *string `pulumi:"id"`
	// The kind of item this is. For buckets, this is always storage#bucket.
	Kind *string `pulumi:"kind"`
	// User-provided labels, in key/value pairs.
	Labels map[string]string `pulumi:"labels"`
	// The bucket's lifecycle configuration. See lifecycle management for more information.
	Lifecycle *BucketLifecycle `pulumi:"lifecycle"`
	// The location of the bucket. Object data for objects in the bucket resides in physical storage within this region. Defaults to US. See the developer's guide for the authoritative list.
	Location *string `pulumi:"location"`
	// The type of the bucket location.
	LocationType *string `pulumi:"locationType"`
	// The bucket's logging configuration, which defines the destination bucket and optional name prefix for the current bucket's logs.
	Logging *BucketLogging `pulumi:"logging"`
	// The metadata generation of this bucket.
	Metageneration *string `pulumi:"metageneration"`
	// The name of the bucket.
	Name *string `pulumi:"name"`
	// The owner of the bucket. This is always the project team's owner group.
	Owner                      *BucketOwner `pulumi:"owner"`
	PredefinedAcl              *string      `pulumi:"predefinedAcl"`
	PredefinedDefaultObjectAcl *string      `pulumi:"predefinedDefaultObjectAcl"`
	Project                    string       `pulumi:"project"`
	// The project number of the project the bucket belongs to.
	ProjectNumber          *string `pulumi:"projectNumber"`
	Projection             *string `pulumi:"projection"`
	ProvisionalUserProject *string `pulumi:"provisionalUserProject"`
	// The bucket's retention policy. The retention policy enforces a minimum retention time for all objects contained in the bucket, based on their creation time. Any attempt to overwrite or delete objects younger than the retention period will result in a PERMISSION_DENIED error. An unlocked retention policy can be modified or removed from the bucket via a storage.buckets.update operation. A locked retention policy cannot be removed or shortened in duration for the lifetime of the bucket. Attempting to remove or decrease period of a locked retention policy will result in a PERMISSION_DENIED error.
	RetentionPolicy *BucketRetentionPolicy `pulumi:"retentionPolicy"`
	// Reserved for future use.
	SatisfiesPZS *bool `pulumi:"satisfiesPZS"`
	// The URI of this bucket.
	SelfLink *string `pulumi:"selfLink"`
	// The bucket's default storage class, used whenever no storageClass is specified for a newly-created object. This defines how objects in the bucket are stored and determines the SLA and the cost of storage. Values include MULTI_REGIONAL, REGIONAL, STANDARD, NEARLINE, COLDLINE, ARCHIVE, and DURABLE_REDUCED_AVAILABILITY. If this value is not specified when the bucket is created, it will default to STANDARD. For more information, see storage classes.
	StorageClass *string `pulumi:"storageClass"`
	// The creation time of the bucket in RFC 3339 format.
	TimeCreated *string `pulumi:"timeCreated"`
	// The modification time of the bucket in RFC 3339 format.
	Updated     *string `pulumi:"updated"`
	UserProject *string `pulumi:"userProject"`
	// The bucket's versioning configuration.
	Versioning *BucketVersioning `pulumi:"versioning"`
	// The bucket's website configuration, controlling how the service behaves when accessing bucket contents as a web site. See the Static Website Examples for more information.
	Website *BucketWebsite `pulumi:"website"`
	// The zone or zones from which the bucket is intended to use zonal quota. Requests for data from outside the specified affinities are still allowed but won't be able to use zonal quota. The zone or zones need to be within the bucket location otherwise the requests will fail with a 400 Bad Request response.
	ZoneAffinity []string `pulumi:"zoneAffinity"`
}

// The set of arguments for constructing a Bucket resource.
type BucketArgs struct {
	// Access controls on the bucket.
	Acl BucketAccessControlTypeArrayInput
	// The bucket's billing configuration.
	Billing BucketBillingPtrInput
	// The bucket's Cross-Origin Resource Sharing (CORS) configuration.
	Cors BucketCorsItemArrayInput
	// The default value for event-based hold on newly created objects in this bucket. Event-based hold is a way to retain objects indefinitely until an event occurs, signified by the hold's release. After being released, such objects will be subject to bucket-level retention (if any). One sample use case of this flag is for banks to hold loan documents for at least 3 years after loan is paid in full. Here, bucket-level retention is 3 years and the event is loan being paid in full. In this example, these objects will be held intact for any number of years until the event has occurred (event-based hold on the object is released) and then 3 more years after that. That means retention duration of the objects begins from the moment event-based hold transitioned from true to false. Objects under event-based hold cannot be deleted, overwritten or archived until the hold is removed.
	DefaultEventBasedHold pulumi.BoolPtrInput
	// Default access controls to apply to new objects when no ACL is provided.
	DefaultObjectAcl ObjectAccessControlTypeArrayInput
	// Encryption configuration for a bucket.
	Encryption BucketEncryptionPtrInput
	// HTTP 1.1 Entity tag for the bucket.
	Etag pulumi.StringPtrInput
	// The bucket's IAM configuration.
	IamConfiguration BucketIamConfigurationPtrInput
	// The ID of the bucket. For buckets, the id and name properties are the same.
	Id pulumi.StringPtrInput
	// The kind of item this is. For buckets, this is always storage#bucket.
	Kind pulumi.StringPtrInput
	// User-provided labels, in key/value pairs.
	Labels pulumi.StringMapInput
	// The bucket's lifecycle configuration. See lifecycle management for more information.
	Lifecycle BucketLifecyclePtrInput
	// The location of the bucket. Object data for objects in the bucket resides in physical storage within this region. Defaults to US. See the developer's guide for the authoritative list.
	Location pulumi.StringPtrInput
	// The type of the bucket location.
	LocationType pulumi.StringPtrInput
	// The bucket's logging configuration, which defines the destination bucket and optional name prefix for the current bucket's logs.
	Logging BucketLoggingPtrInput
	// The metadata generation of this bucket.
	Metageneration pulumi.StringPtrInput
	// The name of the bucket.
	Name pulumi.StringPtrInput
	// The owner of the bucket. This is always the project team's owner group.
	Owner                      BucketOwnerPtrInput
	PredefinedAcl              pulumi.StringPtrInput
	PredefinedDefaultObjectAcl pulumi.StringPtrInput
	Project                    pulumi.StringInput
	// The project number of the project the bucket belongs to.
	ProjectNumber          pulumi.StringPtrInput
	Projection             pulumi.StringPtrInput
	ProvisionalUserProject pulumi.StringPtrInput
	// The bucket's retention policy. The retention policy enforces a minimum retention time for all objects contained in the bucket, based on their creation time. Any attempt to overwrite or delete objects younger than the retention period will result in a PERMISSION_DENIED error. An unlocked retention policy can be modified or removed from the bucket via a storage.buckets.update operation. A locked retention policy cannot be removed or shortened in duration for the lifetime of the bucket. Attempting to remove or decrease period of a locked retention policy will result in a PERMISSION_DENIED error.
	RetentionPolicy BucketRetentionPolicyPtrInput
	// Reserved for future use.
	SatisfiesPZS pulumi.BoolPtrInput
	// The URI of this bucket.
	SelfLink pulumi.StringPtrInput
	// The bucket's default storage class, used whenever no storageClass is specified for a newly-created object. This defines how objects in the bucket are stored and determines the SLA and the cost of storage. Values include MULTI_REGIONAL, REGIONAL, STANDARD, NEARLINE, COLDLINE, ARCHIVE, and DURABLE_REDUCED_AVAILABILITY. If this value is not specified when the bucket is created, it will default to STANDARD. For more information, see storage classes.
	StorageClass pulumi.StringPtrInput
	// The creation time of the bucket in RFC 3339 format.
	TimeCreated pulumi.StringPtrInput
	// The modification time of the bucket in RFC 3339 format.
	Updated     pulumi.StringPtrInput
	UserProject pulumi.StringPtrInput
	// The bucket's versioning configuration.
	Versioning BucketVersioningPtrInput
	// The bucket's website configuration, controlling how the service behaves when accessing bucket contents as a web site. See the Static Website Examples for more information.
	Website BucketWebsitePtrInput
	// The zone or zones from which the bucket is intended to use zonal quota. Requests for data from outside the specified affinities are still allowed but won't be able to use zonal quota. The zone or zones need to be within the bucket location otherwise the requests will fail with a 400 Bad Request response.
	ZoneAffinity pulumi.StringArrayInput
}

func (BucketArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*bucketArgs)(nil)).Elem()
}

type BucketInput interface {
	pulumi.Input

	ToBucketOutput() BucketOutput
	ToBucketOutputWithContext(ctx context.Context) BucketOutput
}

func (*Bucket) ElementType() reflect.Type {
	return reflect.TypeOf((*Bucket)(nil))
}

func (i *Bucket) ToBucketOutput() BucketOutput {
	return i.ToBucketOutputWithContext(context.Background())
}

func (i *Bucket) ToBucketOutputWithContext(ctx context.Context) BucketOutput {
	return pulumi.ToOutputWithContext(ctx, i).(BucketOutput)
}

type BucketOutput struct {
	*pulumi.OutputState
}

func (BucketOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*Bucket)(nil))
}

func (o BucketOutput) ToBucketOutput() BucketOutput {
	return o
}

func (o BucketOutput) ToBucketOutputWithContext(ctx context.Context) BucketOutput {
	return o
}

func init() {
	pulumi.RegisterOutputType(BucketOutput{})
}
