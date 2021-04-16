// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package v2

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Creates a bucket that can be used to store log entries. Once a bucket has been created, the region cannot be changed.
type OrganizationBucket struct {
	pulumi.CustomResourceState

	// The creation timestamp of the bucket. This is not set for any of the default buckets.
	CreateTime pulumi.StringOutput `pulumi:"createTime"`
	// Describes this bucket.
	Description pulumi.StringOutput `pulumi:"description"`
	// The bucket lifecycle state.
	LifecycleState pulumi.StringOutput `pulumi:"lifecycleState"`
	// Whether the bucket has been locked. The retention period on a locked bucket may not be changed. Locked buckets may only be deleted if they are empty.
	Locked pulumi.BoolOutput `pulumi:"locked"`
	// The resource name of the bucket. For example: "projects/my-project-id/locations/my-location/buckets/my-bucket-id" The supported locations are: global, us-central1, us-east1, us-west1, asia-east1, europe-west1.For the location of global it is unspecified where logs are actually stored. Once a bucket has been created, the location can not be changed.
	Name pulumi.StringOutput `pulumi:"name"`
	// Logs will be retained by default for this amount of time, after which they will automatically be deleted. The minimum retention period is 1 day. If this value is set to zero at bucket creation time, the default time of 30 days will be used.
	RetentionDays pulumi.IntOutput `pulumi:"retentionDays"`
	// The last update timestamp of the bucket.
	UpdateTime pulumi.StringOutput `pulumi:"updateTime"`
}

// NewOrganizationBucket registers a new resource with the given unique name, arguments, and options.
func NewOrganizationBucket(ctx *pulumi.Context,
	name string, args *OrganizationBucketArgs, opts ...pulumi.ResourceOption) (*OrganizationBucket, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.BucketsId == nil {
		return nil, errors.New("invalid value for required argument 'BucketsId'")
	}
	if args.LocationsId == nil {
		return nil, errors.New("invalid value for required argument 'LocationsId'")
	}
	if args.OrganizationsId == nil {
		return nil, errors.New("invalid value for required argument 'OrganizationsId'")
	}
	var resource OrganizationBucket
	err := ctx.RegisterResource("google-native:logging/v2:OrganizationBucket", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetOrganizationBucket gets an existing OrganizationBucket resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetOrganizationBucket(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *OrganizationBucketState, opts ...pulumi.ResourceOption) (*OrganizationBucket, error) {
	var resource OrganizationBucket
	err := ctx.ReadResource("google-native:logging/v2:OrganizationBucket", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering OrganizationBucket resources.
type organizationBucketState struct {
	// The creation timestamp of the bucket. This is not set for any of the default buckets.
	CreateTime *string `pulumi:"createTime"`
	// Describes this bucket.
	Description *string `pulumi:"description"`
	// The bucket lifecycle state.
	LifecycleState *string `pulumi:"lifecycleState"`
	// Whether the bucket has been locked. The retention period on a locked bucket may not be changed. Locked buckets may only be deleted if they are empty.
	Locked *bool `pulumi:"locked"`
	// The resource name of the bucket. For example: "projects/my-project-id/locations/my-location/buckets/my-bucket-id" The supported locations are: global, us-central1, us-east1, us-west1, asia-east1, europe-west1.For the location of global it is unspecified where logs are actually stored. Once a bucket has been created, the location can not be changed.
	Name *string `pulumi:"name"`
	// Logs will be retained by default for this amount of time, after which they will automatically be deleted. The minimum retention period is 1 day. If this value is set to zero at bucket creation time, the default time of 30 days will be used.
	RetentionDays *int `pulumi:"retentionDays"`
	// The last update timestamp of the bucket.
	UpdateTime *string `pulumi:"updateTime"`
}

type OrganizationBucketState struct {
	// The creation timestamp of the bucket. This is not set for any of the default buckets.
	CreateTime pulumi.StringPtrInput
	// Describes this bucket.
	Description pulumi.StringPtrInput
	// The bucket lifecycle state.
	LifecycleState pulumi.StringPtrInput
	// Whether the bucket has been locked. The retention period on a locked bucket may not be changed. Locked buckets may only be deleted if they are empty.
	Locked pulumi.BoolPtrInput
	// The resource name of the bucket. For example: "projects/my-project-id/locations/my-location/buckets/my-bucket-id" The supported locations are: global, us-central1, us-east1, us-west1, asia-east1, europe-west1.For the location of global it is unspecified where logs are actually stored. Once a bucket has been created, the location can not be changed.
	Name pulumi.StringPtrInput
	// Logs will be retained by default for this amount of time, after which they will automatically be deleted. The minimum retention period is 1 day. If this value is set to zero at bucket creation time, the default time of 30 days will be used.
	RetentionDays pulumi.IntPtrInput
	// The last update timestamp of the bucket.
	UpdateTime pulumi.StringPtrInput
}

func (OrganizationBucketState) ElementType() reflect.Type {
	return reflect.TypeOf((*organizationBucketState)(nil)).Elem()
}

type organizationBucketArgs struct {
	BucketsId string `pulumi:"bucketsId"`
	// Describes this bucket.
	Description *string `pulumi:"description"`
	LocationsId string  `pulumi:"locationsId"`
	// Whether the bucket has been locked. The retention period on a locked bucket may not be changed. Locked buckets may only be deleted if they are empty.
	Locked          *bool  `pulumi:"locked"`
	OrganizationsId string `pulumi:"organizationsId"`
	// Logs will be retained by default for this amount of time, after which they will automatically be deleted. The minimum retention period is 1 day. If this value is set to zero at bucket creation time, the default time of 30 days will be used.
	RetentionDays *int `pulumi:"retentionDays"`
}

// The set of arguments for constructing a OrganizationBucket resource.
type OrganizationBucketArgs struct {
	BucketsId pulumi.StringInput
	// Describes this bucket.
	Description pulumi.StringPtrInput
	LocationsId pulumi.StringInput
	// Whether the bucket has been locked. The retention period on a locked bucket may not be changed. Locked buckets may only be deleted if they are empty.
	Locked          pulumi.BoolPtrInput
	OrganizationsId pulumi.StringInput
	// Logs will be retained by default for this amount of time, after which they will automatically be deleted. The minimum retention period is 1 day. If this value is set to zero at bucket creation time, the default time of 30 days will be used.
	RetentionDays pulumi.IntPtrInput
}

func (OrganizationBucketArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*organizationBucketArgs)(nil)).Elem()
}

type OrganizationBucketInput interface {
	pulumi.Input

	ToOrganizationBucketOutput() OrganizationBucketOutput
	ToOrganizationBucketOutputWithContext(ctx context.Context) OrganizationBucketOutput
}

func (*OrganizationBucket) ElementType() reflect.Type {
	return reflect.TypeOf((*OrganizationBucket)(nil))
}

func (i *OrganizationBucket) ToOrganizationBucketOutput() OrganizationBucketOutput {
	return i.ToOrganizationBucketOutputWithContext(context.Background())
}

func (i *OrganizationBucket) ToOrganizationBucketOutputWithContext(ctx context.Context) OrganizationBucketOutput {
	return pulumi.ToOutputWithContext(ctx, i).(OrganizationBucketOutput)
}

type OrganizationBucketOutput struct {
	*pulumi.OutputState
}

func (OrganizationBucketOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*OrganizationBucket)(nil))
}

func (o OrganizationBucketOutput) ToOrganizationBucketOutput() OrganizationBucketOutput {
	return o
}

func (o OrganizationBucketOutput) ToOrganizationBucketOutputWithContext(ctx context.Context) OrganizationBucketOutput {
	return o
}

func init() {
	pulumi.RegisterOutputType(OrganizationBucketOutput{})
}