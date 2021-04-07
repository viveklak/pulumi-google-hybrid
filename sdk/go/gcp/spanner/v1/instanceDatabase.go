// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package v1

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
)

// Creates a new Cloud Spanner database and starts to prepare it for serving. The returned long-running operation will have a name of the format `/operations/` and can be used to track preparation of the database. The metadata field type is CreateDatabaseMetadata. The response field type is Database, if successful.
type InstanceDatabase struct {
	pulumi.CustomResourceState

	// If exists, the time at which the database creation started.
	CreateTime pulumi.StringOutput `pulumi:"createTime"`
	// Earliest timestamp at which older versions of the data can be read. This value is continuously updated by Cloud Spanner and becomes stale the moment it is queried. If you are using this value to recover data, make sure to account for the time from the moment when the value is queried to the moment when you initiate the recovery.
	EarliestVersionTime pulumi.StringOutput `pulumi:"earliestVersionTime"`
	// For databases that are using customer managed encryption, this field contains the encryption configuration for the database. For databases that are using Google default or other types of encryption, this field is empty.
	EncryptionConfig EncryptionConfigResponseOutput `pulumi:"encryptionConfig"`
	// For databases that are using customer managed encryption, this field contains the encryption information for the database, such as encryption state and the Cloud KMS key versions that are in use. For databases that are using Google default or other types of encryption, this field is empty. This field is propagated lazily from the backend. There might be a delay from when a key version is being used and when it appears in this field.
	EncryptionInfo EncryptionInfoResponseArrayOutput `pulumi:"encryptionInfo"`
	// Required. The name of the database. Values are of the form `projects//instances//databases/`, where `` is as specified in the `CREATE DATABASE` statement. This name can be passed to other API methods to identify the database.
	Name pulumi.StringOutput `pulumi:"name"`
	// Applicable only for restored databases. Contains information about the restore source.
	RestoreInfo RestoreInfoResponseOutput `pulumi:"restoreInfo"`
	// The current database state.
	State pulumi.StringOutput `pulumi:"state"`
	// The period in which Cloud Spanner retains all versions of data for the database. This is the same as the value of version_retention_period database option set using UpdateDatabaseDdl. Defaults to 1 hour, if not set.
	VersionRetentionPeriod pulumi.StringOutput `pulumi:"versionRetentionPeriod"`
}

// NewInstanceDatabase registers a new resource with the given unique name, arguments, and options.
func NewInstanceDatabase(ctx *pulumi.Context,
	name string, args *InstanceDatabaseArgs, opts ...pulumi.ResourceOption) (*InstanceDatabase, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.DatabasesId == nil {
		return nil, errors.New("invalid value for required argument 'DatabasesId'")
	}
	if args.InstancesId == nil {
		return nil, errors.New("invalid value for required argument 'InstancesId'")
	}
	if args.ProjectsId == nil {
		return nil, errors.New("invalid value for required argument 'ProjectsId'")
	}
	var resource InstanceDatabase
	err := ctx.RegisterResource("gcp-native:spanner/v1:InstanceDatabase", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetInstanceDatabase gets an existing InstanceDatabase resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetInstanceDatabase(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *InstanceDatabaseState, opts ...pulumi.ResourceOption) (*InstanceDatabase, error) {
	var resource InstanceDatabase
	err := ctx.ReadResource("gcp-native:spanner/v1:InstanceDatabase", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering InstanceDatabase resources.
type instanceDatabaseState struct {
	// If exists, the time at which the database creation started.
	CreateTime *string `pulumi:"createTime"`
	// Earliest timestamp at which older versions of the data can be read. This value is continuously updated by Cloud Spanner and becomes stale the moment it is queried. If you are using this value to recover data, make sure to account for the time from the moment when the value is queried to the moment when you initiate the recovery.
	EarliestVersionTime *string `pulumi:"earliestVersionTime"`
	// For databases that are using customer managed encryption, this field contains the encryption configuration for the database. For databases that are using Google default or other types of encryption, this field is empty.
	EncryptionConfig *EncryptionConfigResponse `pulumi:"encryptionConfig"`
	// For databases that are using customer managed encryption, this field contains the encryption information for the database, such as encryption state and the Cloud KMS key versions that are in use. For databases that are using Google default or other types of encryption, this field is empty. This field is propagated lazily from the backend. There might be a delay from when a key version is being used and when it appears in this field.
	EncryptionInfo []EncryptionInfoResponse `pulumi:"encryptionInfo"`
	// Required. The name of the database. Values are of the form `projects//instances//databases/`, where `` is as specified in the `CREATE DATABASE` statement. This name can be passed to other API methods to identify the database.
	Name *string `pulumi:"name"`
	// Applicable only for restored databases. Contains information about the restore source.
	RestoreInfo *RestoreInfoResponse `pulumi:"restoreInfo"`
	// The current database state.
	State *string `pulumi:"state"`
	// The period in which Cloud Spanner retains all versions of data for the database. This is the same as the value of version_retention_period database option set using UpdateDatabaseDdl. Defaults to 1 hour, if not set.
	VersionRetentionPeriod *string `pulumi:"versionRetentionPeriod"`
}

type InstanceDatabaseState struct {
	// If exists, the time at which the database creation started.
	CreateTime pulumi.StringPtrInput
	// Earliest timestamp at which older versions of the data can be read. This value is continuously updated by Cloud Spanner and becomes stale the moment it is queried. If you are using this value to recover data, make sure to account for the time from the moment when the value is queried to the moment when you initiate the recovery.
	EarliestVersionTime pulumi.StringPtrInput
	// For databases that are using customer managed encryption, this field contains the encryption configuration for the database. For databases that are using Google default or other types of encryption, this field is empty.
	EncryptionConfig EncryptionConfigResponsePtrInput
	// For databases that are using customer managed encryption, this field contains the encryption information for the database, such as encryption state and the Cloud KMS key versions that are in use. For databases that are using Google default or other types of encryption, this field is empty. This field is propagated lazily from the backend. There might be a delay from when a key version is being used and when it appears in this field.
	EncryptionInfo EncryptionInfoResponseArrayInput
	// Required. The name of the database. Values are of the form `projects//instances//databases/`, where `` is as specified in the `CREATE DATABASE` statement. This name can be passed to other API methods to identify the database.
	Name pulumi.StringPtrInput
	// Applicable only for restored databases. Contains information about the restore source.
	RestoreInfo RestoreInfoResponsePtrInput
	// The current database state.
	State pulumi.StringPtrInput
	// The period in which Cloud Spanner retains all versions of data for the database. This is the same as the value of version_retention_period database option set using UpdateDatabaseDdl. Defaults to 1 hour, if not set.
	VersionRetentionPeriod pulumi.StringPtrInput
}

func (InstanceDatabaseState) ElementType() reflect.Type {
	return reflect.TypeOf((*instanceDatabaseState)(nil)).Elem()
}

type instanceDatabaseArgs struct {
	// Required. A `CREATE DATABASE` statement, which specifies the ID of the new database. The database ID must conform to the regular expression `a-z*[a-z0-9]` and be between 2 and 30 characters in length. If the database ID is a reserved word or if it contains a hyphen, the database ID must be enclosed in backticks (`` ` ``).
	CreateStatement *string `pulumi:"createStatement"`
	DatabasesId     string  `pulumi:"databasesId"`
	// Optional. The encryption configuration for the database. If this field is not specified, Cloud Spanner will encrypt/decrypt all data at rest using Google default encryption.
	EncryptionConfig *EncryptionConfig `pulumi:"encryptionConfig"`
	// Optional. A list of DDL statements to run inside the newly created database. Statements can create tables, indexes, etc. These statements execute atomically with the creation of the database: if there is an error in any statement, the database is not created.
	ExtraStatements []string `pulumi:"extraStatements"`
	InstancesId     string   `pulumi:"instancesId"`
	ProjectsId      string   `pulumi:"projectsId"`
}

// The set of arguments for constructing a InstanceDatabase resource.
type InstanceDatabaseArgs struct {
	// Required. A `CREATE DATABASE` statement, which specifies the ID of the new database. The database ID must conform to the regular expression `a-z*[a-z0-9]` and be between 2 and 30 characters in length. If the database ID is a reserved word or if it contains a hyphen, the database ID must be enclosed in backticks (`` ` ``).
	CreateStatement pulumi.StringPtrInput
	DatabasesId     pulumi.StringInput
	// Optional. The encryption configuration for the database. If this field is not specified, Cloud Spanner will encrypt/decrypt all data at rest using Google default encryption.
	EncryptionConfig EncryptionConfigPtrInput
	// Optional. A list of DDL statements to run inside the newly created database. Statements can create tables, indexes, etc. These statements execute atomically with the creation of the database: if there is an error in any statement, the database is not created.
	ExtraStatements pulumi.StringArrayInput
	InstancesId     pulumi.StringInput
	ProjectsId      pulumi.StringInput
}

func (InstanceDatabaseArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*instanceDatabaseArgs)(nil)).Elem()
}

type InstanceDatabaseInput interface {
	pulumi.Input

	ToInstanceDatabaseOutput() InstanceDatabaseOutput
	ToInstanceDatabaseOutputWithContext(ctx context.Context) InstanceDatabaseOutput
}

func (*InstanceDatabase) ElementType() reflect.Type {
	return reflect.TypeOf((*InstanceDatabase)(nil))
}

func (i *InstanceDatabase) ToInstanceDatabaseOutput() InstanceDatabaseOutput {
	return i.ToInstanceDatabaseOutputWithContext(context.Background())
}

func (i *InstanceDatabase) ToInstanceDatabaseOutputWithContext(ctx context.Context) InstanceDatabaseOutput {
	return pulumi.ToOutputWithContext(ctx, i).(InstanceDatabaseOutput)
}

type InstanceDatabaseOutput struct {
	*pulumi.OutputState
}

func (InstanceDatabaseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*InstanceDatabase)(nil))
}

func (o InstanceDatabaseOutput) ToInstanceDatabaseOutput() InstanceDatabaseOutput {
	return o
}

func (o InstanceDatabaseOutput) ToInstanceDatabaseOutputWithContext(ctx context.Context) InstanceDatabaseOutput {
	return o
}

func init() {
	pulumi.RegisterOutputType(InstanceDatabaseOutput{})
}