// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package v1

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
)

// Creates a composite index. This returns a google.longrunning.Operation which may be used to track the status of the creation. The metadata for the operation will be the type IndexOperationMetadata.
type DatabaseCollectionGroupIndex struct {
	pulumi.CustomResourceState
}

// NewDatabaseCollectionGroupIndex registers a new resource with the given unique name, arguments, and options.
func NewDatabaseCollectionGroupIndex(ctx *pulumi.Context,
	name string, args *DatabaseCollectionGroupIndexArgs, opts ...pulumi.ResourceOption) (*DatabaseCollectionGroupIndex, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.CollectionGroupsId == nil {
		return nil, errors.New("invalid value for required argument 'CollectionGroupsId'")
	}
	if args.DatabasesId == nil {
		return nil, errors.New("invalid value for required argument 'DatabasesId'")
	}
	if args.IndexesId == nil {
		return nil, errors.New("invalid value for required argument 'IndexesId'")
	}
	if args.ProjectsId == nil {
		return nil, errors.New("invalid value for required argument 'ProjectsId'")
	}
	var resource DatabaseCollectionGroupIndex
	err := ctx.RegisterResource("google-cloud:firestore/v1:DatabaseCollectionGroupIndex", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetDatabaseCollectionGroupIndex gets an existing DatabaseCollectionGroupIndex resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetDatabaseCollectionGroupIndex(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *DatabaseCollectionGroupIndexState, opts ...pulumi.ResourceOption) (*DatabaseCollectionGroupIndex, error) {
	var resource DatabaseCollectionGroupIndex
	err := ctx.ReadResource("google-cloud:firestore/v1:DatabaseCollectionGroupIndex", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering DatabaseCollectionGroupIndex resources.
type databaseCollectionGroupIndexState struct {
}

type DatabaseCollectionGroupIndexState struct {
}

func (DatabaseCollectionGroupIndexState) ElementType() reflect.Type {
	return reflect.TypeOf((*databaseCollectionGroupIndexState)(nil)).Elem()
}

type databaseCollectionGroupIndexArgs struct {
	CollectionGroupsId string `pulumi:"collectionGroupsId"`
	DatabasesId        string `pulumi:"databasesId"`
	// The fields supported by this index. For composite indexes, this is always 2 or more fields. The last field entry is always for the field path `__name__`. If, on creation, `__name__` was not specified as the last field, it will be added automatically with the same direction as that of the last field defined. If the final field in a composite index is not directional, the `__name__` will be ordered ASCENDING (unless explicitly specified). For single field indexes, this will always be exactly one entry with a field path equal to the field path of the associated field.
	Fields    []GoogleFirestoreAdminV1IndexField `pulumi:"fields"`
	IndexesId string                             `pulumi:"indexesId"`
	// Output only. A server defined name for this index. The form of this name for composite indexes will be: `projects/{project_id}/databases/{database_id}/collectionGroups/{collection_id}/indexes/{composite_index_id}` For single field indexes, this field will be empty.
	Name       *string `pulumi:"name"`
	ProjectsId string  `pulumi:"projectsId"`
	// Indexes with a collection query scope specified allow queries against a collection that is the child of a specific document, specified at query time, and that has the same collection id. Indexes with a collection group query scope specified allow queries against all collections descended from a specific document, specified at query time, and that have the same collection id as this index.
	QueryScope *string `pulumi:"queryScope"`
	// Output only. The serving state of the index.
	State *string `pulumi:"state"`
}

// The set of arguments for constructing a DatabaseCollectionGroupIndex resource.
type DatabaseCollectionGroupIndexArgs struct {
	CollectionGroupsId pulumi.StringInput
	DatabasesId        pulumi.StringInput
	// The fields supported by this index. For composite indexes, this is always 2 or more fields. The last field entry is always for the field path `__name__`. If, on creation, `__name__` was not specified as the last field, it will be added automatically with the same direction as that of the last field defined. If the final field in a composite index is not directional, the `__name__` will be ordered ASCENDING (unless explicitly specified). For single field indexes, this will always be exactly one entry with a field path equal to the field path of the associated field.
	Fields    GoogleFirestoreAdminV1IndexFieldArrayInput
	IndexesId pulumi.StringInput
	// Output only. A server defined name for this index. The form of this name for composite indexes will be: `projects/{project_id}/databases/{database_id}/collectionGroups/{collection_id}/indexes/{composite_index_id}` For single field indexes, this field will be empty.
	Name       pulumi.StringPtrInput
	ProjectsId pulumi.StringInput
	// Indexes with a collection query scope specified allow queries against a collection that is the child of a specific document, specified at query time, and that has the same collection id. Indexes with a collection group query scope specified allow queries against all collections descended from a specific document, specified at query time, and that have the same collection id as this index.
	QueryScope pulumi.StringPtrInput
	// Output only. The serving state of the index.
	State pulumi.StringPtrInput
}

func (DatabaseCollectionGroupIndexArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*databaseCollectionGroupIndexArgs)(nil)).Elem()
}

type DatabaseCollectionGroupIndexInput interface {
	pulumi.Input

	ToDatabaseCollectionGroupIndexOutput() DatabaseCollectionGroupIndexOutput
	ToDatabaseCollectionGroupIndexOutputWithContext(ctx context.Context) DatabaseCollectionGroupIndexOutput
}

func (*DatabaseCollectionGroupIndex) ElementType() reflect.Type {
	return reflect.TypeOf((*DatabaseCollectionGroupIndex)(nil))
}

func (i *DatabaseCollectionGroupIndex) ToDatabaseCollectionGroupIndexOutput() DatabaseCollectionGroupIndexOutput {
	return i.ToDatabaseCollectionGroupIndexOutputWithContext(context.Background())
}

func (i *DatabaseCollectionGroupIndex) ToDatabaseCollectionGroupIndexOutputWithContext(ctx context.Context) DatabaseCollectionGroupIndexOutput {
	return pulumi.ToOutputWithContext(ctx, i).(DatabaseCollectionGroupIndexOutput)
}

type DatabaseCollectionGroupIndexOutput struct {
	*pulumi.OutputState
}

func (DatabaseCollectionGroupIndexOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*DatabaseCollectionGroupIndex)(nil))
}

func (o DatabaseCollectionGroupIndexOutput) ToDatabaseCollectionGroupIndexOutput() DatabaseCollectionGroupIndexOutput {
	return o
}

func (o DatabaseCollectionGroupIndexOutput) ToDatabaseCollectionGroupIndexOutputWithContext(ctx context.Context) DatabaseCollectionGroupIndexOutput {
	return o
}

func init() {
	pulumi.RegisterOutputType(DatabaseCollectionGroupIndexOutput{})
}
