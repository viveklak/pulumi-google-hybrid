// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package v1beta2

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Indicates that this field supports operations on `array_value`s.
type GoogleFirestoreAdminV1beta2IndexFieldArrayConfig pulumi.String

const (
	// The index does not support additional array queries.
	GoogleFirestoreAdminV1beta2IndexFieldArrayConfigArrayConfigUnspecified = GoogleFirestoreAdminV1beta2IndexFieldArrayConfig("ARRAY_CONFIG_UNSPECIFIED")
	// The index supports array containment queries.
	GoogleFirestoreAdminV1beta2IndexFieldArrayConfigContains = GoogleFirestoreAdminV1beta2IndexFieldArrayConfig("CONTAINS")
)

func (GoogleFirestoreAdminV1beta2IndexFieldArrayConfig) ElementType() reflect.Type {
	return reflect.TypeOf((*pulumi.String)(nil)).Elem()
}

func (e GoogleFirestoreAdminV1beta2IndexFieldArrayConfig) ToStringOutput() pulumi.StringOutput {
	return pulumi.ToOutput(pulumi.String(e)).(pulumi.StringOutput)
}

func (e GoogleFirestoreAdminV1beta2IndexFieldArrayConfig) ToStringOutputWithContext(ctx context.Context) pulumi.StringOutput {
	return pulumi.ToOutputWithContext(ctx, pulumi.String(e)).(pulumi.StringOutput)
}

func (e GoogleFirestoreAdminV1beta2IndexFieldArrayConfig) ToStringPtrOutput() pulumi.StringPtrOutput {
	return pulumi.String(e).ToStringPtrOutputWithContext(context.Background())
}

func (e GoogleFirestoreAdminV1beta2IndexFieldArrayConfig) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return pulumi.String(e).ToStringOutputWithContext(ctx).ToStringPtrOutputWithContext(ctx)
}

// Indicates that this field supports ordering by the specified order or comparing using =, <, <=, >, >=.
type GoogleFirestoreAdminV1beta2IndexFieldOrder pulumi.String

const (
	// The ordering is unspecified. Not a valid option.
	GoogleFirestoreAdminV1beta2IndexFieldOrderOrderUnspecified = GoogleFirestoreAdminV1beta2IndexFieldOrder("ORDER_UNSPECIFIED")
	// The field is ordered by ascending field value.
	GoogleFirestoreAdminV1beta2IndexFieldOrderAscending = GoogleFirestoreAdminV1beta2IndexFieldOrder("ASCENDING")
	// The field is ordered by descending field value.
	GoogleFirestoreAdminV1beta2IndexFieldOrderDescending = GoogleFirestoreAdminV1beta2IndexFieldOrder("DESCENDING")
)

func (GoogleFirestoreAdminV1beta2IndexFieldOrder) ElementType() reflect.Type {
	return reflect.TypeOf((*pulumi.String)(nil)).Elem()
}

func (e GoogleFirestoreAdminV1beta2IndexFieldOrder) ToStringOutput() pulumi.StringOutput {
	return pulumi.ToOutput(pulumi.String(e)).(pulumi.StringOutput)
}

func (e GoogleFirestoreAdminV1beta2IndexFieldOrder) ToStringOutputWithContext(ctx context.Context) pulumi.StringOutput {
	return pulumi.ToOutputWithContext(ctx, pulumi.String(e)).(pulumi.StringOutput)
}

func (e GoogleFirestoreAdminV1beta2IndexFieldOrder) ToStringPtrOutput() pulumi.StringPtrOutput {
	return pulumi.String(e).ToStringPtrOutputWithContext(context.Background())
}

func (e GoogleFirestoreAdminV1beta2IndexFieldOrder) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return pulumi.String(e).ToStringOutputWithContext(ctx).ToStringPtrOutputWithContext(ctx)
}

// Indexes with a collection query scope specified allow queries against a collection that is the child of a specific document, specified at query time, and that has the same collection id. Indexes with a collection group query scope specified allow queries against all collections descended from a specific document, specified at query time, and that have the same collection id as this index.
type IndexQueryScope pulumi.String

const (
	// The query scope is unspecified. Not a valid option.
	IndexQueryScopeQueryScopeUnspecified = IndexQueryScope("QUERY_SCOPE_UNSPECIFIED")
	// Indexes with a collection query scope specified allow queries against a collection that is the child of a specific document, specified at query time, and that has the collection id specified by the index.
	IndexQueryScopeCollection = IndexQueryScope("COLLECTION")
	// Indexes with a collection group query scope specified allow queries against all collections that has the collection id specified by the index.
	IndexQueryScopeCollectionGroup = IndexQueryScope("COLLECTION_GROUP")
)

func (IndexQueryScope) ElementType() reflect.Type {
	return reflect.TypeOf((*pulumi.String)(nil)).Elem()
}

func (e IndexQueryScope) ToStringOutput() pulumi.StringOutput {
	return pulumi.ToOutput(pulumi.String(e)).(pulumi.StringOutput)
}

func (e IndexQueryScope) ToStringOutputWithContext(ctx context.Context) pulumi.StringOutput {
	return pulumi.ToOutputWithContext(ctx, pulumi.String(e)).(pulumi.StringOutput)
}

func (e IndexQueryScope) ToStringPtrOutput() pulumi.StringPtrOutput {
	return pulumi.String(e).ToStringPtrOutputWithContext(context.Background())
}

func (e IndexQueryScope) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return pulumi.String(e).ToStringOutputWithContext(ctx).ToStringPtrOutputWithContext(ctx)
}
