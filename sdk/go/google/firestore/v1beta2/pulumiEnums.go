// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package v1beta2

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Indicates that this field supports operations on `array_value`s.
type GoogleFirestoreAdminV1beta2IndexFieldArrayConfig string

const (
	// The index does not support additional array queries.
	GoogleFirestoreAdminV1beta2IndexFieldArrayConfigArrayConfigUnspecified = GoogleFirestoreAdminV1beta2IndexFieldArrayConfig("ARRAY_CONFIG_UNSPECIFIED")
	// The index supports array containment queries.
	GoogleFirestoreAdminV1beta2IndexFieldArrayConfigContains = GoogleFirestoreAdminV1beta2IndexFieldArrayConfig("CONTAINS")
)

func (GoogleFirestoreAdminV1beta2IndexFieldArrayConfig) ElementType() reflect.Type {
	return reflect.TypeOf((*GoogleFirestoreAdminV1beta2IndexFieldArrayConfig)(nil)).Elem()
}

func (e GoogleFirestoreAdminV1beta2IndexFieldArrayConfig) ToGoogleFirestoreAdminV1beta2IndexFieldArrayConfigOutput() GoogleFirestoreAdminV1beta2IndexFieldArrayConfigOutput {
	return pulumi.ToOutput(e).(GoogleFirestoreAdminV1beta2IndexFieldArrayConfigOutput)
}

func (e GoogleFirestoreAdminV1beta2IndexFieldArrayConfig) ToGoogleFirestoreAdminV1beta2IndexFieldArrayConfigOutputWithContext(ctx context.Context) GoogleFirestoreAdminV1beta2IndexFieldArrayConfigOutput {
	return pulumi.ToOutputWithContext(ctx, e).(GoogleFirestoreAdminV1beta2IndexFieldArrayConfigOutput)
}

func (e GoogleFirestoreAdminV1beta2IndexFieldArrayConfig) ToGoogleFirestoreAdminV1beta2IndexFieldArrayConfigPtrOutput() GoogleFirestoreAdminV1beta2IndexFieldArrayConfigPtrOutput {
	return e.ToGoogleFirestoreAdminV1beta2IndexFieldArrayConfigPtrOutputWithContext(context.Background())
}

func (e GoogleFirestoreAdminV1beta2IndexFieldArrayConfig) ToGoogleFirestoreAdminV1beta2IndexFieldArrayConfigPtrOutputWithContext(ctx context.Context) GoogleFirestoreAdminV1beta2IndexFieldArrayConfigPtrOutput {
	return GoogleFirestoreAdminV1beta2IndexFieldArrayConfig(e).ToGoogleFirestoreAdminV1beta2IndexFieldArrayConfigOutputWithContext(ctx).ToGoogleFirestoreAdminV1beta2IndexFieldArrayConfigPtrOutputWithContext(ctx)
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

type GoogleFirestoreAdminV1beta2IndexFieldArrayConfigOutput struct{ *pulumi.OutputState }

func (GoogleFirestoreAdminV1beta2IndexFieldArrayConfigOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*GoogleFirestoreAdminV1beta2IndexFieldArrayConfig)(nil)).Elem()
}

func (o GoogleFirestoreAdminV1beta2IndexFieldArrayConfigOutput) ToGoogleFirestoreAdminV1beta2IndexFieldArrayConfigOutput() GoogleFirestoreAdminV1beta2IndexFieldArrayConfigOutput {
	return o
}

func (o GoogleFirestoreAdminV1beta2IndexFieldArrayConfigOutput) ToGoogleFirestoreAdminV1beta2IndexFieldArrayConfigOutputWithContext(ctx context.Context) GoogleFirestoreAdminV1beta2IndexFieldArrayConfigOutput {
	return o
}

func (o GoogleFirestoreAdminV1beta2IndexFieldArrayConfigOutput) ToGoogleFirestoreAdminV1beta2IndexFieldArrayConfigPtrOutput() GoogleFirestoreAdminV1beta2IndexFieldArrayConfigPtrOutput {
	return o.ToGoogleFirestoreAdminV1beta2IndexFieldArrayConfigPtrOutputWithContext(context.Background())
}

func (o GoogleFirestoreAdminV1beta2IndexFieldArrayConfigOutput) ToGoogleFirestoreAdminV1beta2IndexFieldArrayConfigPtrOutputWithContext(ctx context.Context) GoogleFirestoreAdminV1beta2IndexFieldArrayConfigPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v GoogleFirestoreAdminV1beta2IndexFieldArrayConfig) *GoogleFirestoreAdminV1beta2IndexFieldArrayConfig {
		return &v
	}).(GoogleFirestoreAdminV1beta2IndexFieldArrayConfigPtrOutput)
}

func (o GoogleFirestoreAdminV1beta2IndexFieldArrayConfigOutput) ToStringOutput() pulumi.StringOutput {
	return o.ToStringOutputWithContext(context.Background())
}

func (o GoogleFirestoreAdminV1beta2IndexFieldArrayConfigOutput) ToStringOutputWithContext(ctx context.Context) pulumi.StringOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e GoogleFirestoreAdminV1beta2IndexFieldArrayConfig) string {
		return string(e)
	}).(pulumi.StringOutput)
}

func (o GoogleFirestoreAdminV1beta2IndexFieldArrayConfigOutput) ToStringPtrOutput() pulumi.StringPtrOutput {
	return o.ToStringPtrOutputWithContext(context.Background())
}

func (o GoogleFirestoreAdminV1beta2IndexFieldArrayConfigOutput) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e GoogleFirestoreAdminV1beta2IndexFieldArrayConfig) *string {
		v := string(e)
		return &v
	}).(pulumi.StringPtrOutput)
}

type GoogleFirestoreAdminV1beta2IndexFieldArrayConfigPtrOutput struct{ *pulumi.OutputState }

func (GoogleFirestoreAdminV1beta2IndexFieldArrayConfigPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**GoogleFirestoreAdminV1beta2IndexFieldArrayConfig)(nil)).Elem()
}

func (o GoogleFirestoreAdminV1beta2IndexFieldArrayConfigPtrOutput) ToGoogleFirestoreAdminV1beta2IndexFieldArrayConfigPtrOutput() GoogleFirestoreAdminV1beta2IndexFieldArrayConfigPtrOutput {
	return o
}

func (o GoogleFirestoreAdminV1beta2IndexFieldArrayConfigPtrOutput) ToGoogleFirestoreAdminV1beta2IndexFieldArrayConfigPtrOutputWithContext(ctx context.Context) GoogleFirestoreAdminV1beta2IndexFieldArrayConfigPtrOutput {
	return o
}

func (o GoogleFirestoreAdminV1beta2IndexFieldArrayConfigPtrOutput) Elem() GoogleFirestoreAdminV1beta2IndexFieldArrayConfigOutput {
	return o.ApplyT(func(v *GoogleFirestoreAdminV1beta2IndexFieldArrayConfig) GoogleFirestoreAdminV1beta2IndexFieldArrayConfig {
		if v != nil {
			return *v
		}
		var ret GoogleFirestoreAdminV1beta2IndexFieldArrayConfig
		return ret
	}).(GoogleFirestoreAdminV1beta2IndexFieldArrayConfigOutput)
}

func (o GoogleFirestoreAdminV1beta2IndexFieldArrayConfigPtrOutput) ToStringPtrOutput() pulumi.StringPtrOutput {
	return o.ToStringPtrOutputWithContext(context.Background())
}

func (o GoogleFirestoreAdminV1beta2IndexFieldArrayConfigPtrOutput) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e *GoogleFirestoreAdminV1beta2IndexFieldArrayConfig) *string {
		if e == nil {
			return nil
		}
		v := string(*e)
		return &v
	}).(pulumi.StringPtrOutput)
}

// GoogleFirestoreAdminV1beta2IndexFieldArrayConfigInput is an input type that accepts GoogleFirestoreAdminV1beta2IndexFieldArrayConfigArgs and GoogleFirestoreAdminV1beta2IndexFieldArrayConfigOutput values.
// You can construct a concrete instance of `GoogleFirestoreAdminV1beta2IndexFieldArrayConfigInput` via:
//
//          GoogleFirestoreAdminV1beta2IndexFieldArrayConfigArgs{...}
type GoogleFirestoreAdminV1beta2IndexFieldArrayConfigInput interface {
	pulumi.Input

	ToGoogleFirestoreAdminV1beta2IndexFieldArrayConfigOutput() GoogleFirestoreAdminV1beta2IndexFieldArrayConfigOutput
	ToGoogleFirestoreAdminV1beta2IndexFieldArrayConfigOutputWithContext(context.Context) GoogleFirestoreAdminV1beta2IndexFieldArrayConfigOutput
}

var googleFirestoreAdminV1beta2IndexFieldArrayConfigPtrType = reflect.TypeOf((**GoogleFirestoreAdminV1beta2IndexFieldArrayConfig)(nil)).Elem()

type GoogleFirestoreAdminV1beta2IndexFieldArrayConfigPtrInput interface {
	pulumi.Input

	ToGoogleFirestoreAdminV1beta2IndexFieldArrayConfigPtrOutput() GoogleFirestoreAdminV1beta2IndexFieldArrayConfigPtrOutput
	ToGoogleFirestoreAdminV1beta2IndexFieldArrayConfigPtrOutputWithContext(context.Context) GoogleFirestoreAdminV1beta2IndexFieldArrayConfigPtrOutput
}

type googleFirestoreAdminV1beta2IndexFieldArrayConfigPtr string

func GoogleFirestoreAdminV1beta2IndexFieldArrayConfigPtr(v string) GoogleFirestoreAdminV1beta2IndexFieldArrayConfigPtrInput {
	return (*googleFirestoreAdminV1beta2IndexFieldArrayConfigPtr)(&v)
}

func (*googleFirestoreAdminV1beta2IndexFieldArrayConfigPtr) ElementType() reflect.Type {
	return googleFirestoreAdminV1beta2IndexFieldArrayConfigPtrType
}

func (in *googleFirestoreAdminV1beta2IndexFieldArrayConfigPtr) ToGoogleFirestoreAdminV1beta2IndexFieldArrayConfigPtrOutput() GoogleFirestoreAdminV1beta2IndexFieldArrayConfigPtrOutput {
	return pulumi.ToOutput(in).(GoogleFirestoreAdminV1beta2IndexFieldArrayConfigPtrOutput)
}

func (in *googleFirestoreAdminV1beta2IndexFieldArrayConfigPtr) ToGoogleFirestoreAdminV1beta2IndexFieldArrayConfigPtrOutputWithContext(ctx context.Context) GoogleFirestoreAdminV1beta2IndexFieldArrayConfigPtrOutput {
	return pulumi.ToOutputWithContext(ctx, in).(GoogleFirestoreAdminV1beta2IndexFieldArrayConfigPtrOutput)
}

// Indicates that this field supports ordering by the specified order or comparing using =, <, <=, >, >=.
type GoogleFirestoreAdminV1beta2IndexFieldOrder string

const (
	// The ordering is unspecified. Not a valid option.
	GoogleFirestoreAdminV1beta2IndexFieldOrderOrderUnspecified = GoogleFirestoreAdminV1beta2IndexFieldOrder("ORDER_UNSPECIFIED")
	// The field is ordered by ascending field value.
	GoogleFirestoreAdminV1beta2IndexFieldOrderAscending = GoogleFirestoreAdminV1beta2IndexFieldOrder("ASCENDING")
	// The field is ordered by descending field value.
	GoogleFirestoreAdminV1beta2IndexFieldOrderDescending = GoogleFirestoreAdminV1beta2IndexFieldOrder("DESCENDING")
)

func (GoogleFirestoreAdminV1beta2IndexFieldOrder) ElementType() reflect.Type {
	return reflect.TypeOf((*GoogleFirestoreAdminV1beta2IndexFieldOrder)(nil)).Elem()
}

func (e GoogleFirestoreAdminV1beta2IndexFieldOrder) ToGoogleFirestoreAdminV1beta2IndexFieldOrderOutput() GoogleFirestoreAdminV1beta2IndexFieldOrderOutput {
	return pulumi.ToOutput(e).(GoogleFirestoreAdminV1beta2IndexFieldOrderOutput)
}

func (e GoogleFirestoreAdminV1beta2IndexFieldOrder) ToGoogleFirestoreAdminV1beta2IndexFieldOrderOutputWithContext(ctx context.Context) GoogleFirestoreAdminV1beta2IndexFieldOrderOutput {
	return pulumi.ToOutputWithContext(ctx, e).(GoogleFirestoreAdminV1beta2IndexFieldOrderOutput)
}

func (e GoogleFirestoreAdminV1beta2IndexFieldOrder) ToGoogleFirestoreAdminV1beta2IndexFieldOrderPtrOutput() GoogleFirestoreAdminV1beta2IndexFieldOrderPtrOutput {
	return e.ToGoogleFirestoreAdminV1beta2IndexFieldOrderPtrOutputWithContext(context.Background())
}

func (e GoogleFirestoreAdminV1beta2IndexFieldOrder) ToGoogleFirestoreAdminV1beta2IndexFieldOrderPtrOutputWithContext(ctx context.Context) GoogleFirestoreAdminV1beta2IndexFieldOrderPtrOutput {
	return GoogleFirestoreAdminV1beta2IndexFieldOrder(e).ToGoogleFirestoreAdminV1beta2IndexFieldOrderOutputWithContext(ctx).ToGoogleFirestoreAdminV1beta2IndexFieldOrderPtrOutputWithContext(ctx)
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

type GoogleFirestoreAdminV1beta2IndexFieldOrderOutput struct{ *pulumi.OutputState }

func (GoogleFirestoreAdminV1beta2IndexFieldOrderOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*GoogleFirestoreAdminV1beta2IndexFieldOrder)(nil)).Elem()
}

func (o GoogleFirestoreAdminV1beta2IndexFieldOrderOutput) ToGoogleFirestoreAdminV1beta2IndexFieldOrderOutput() GoogleFirestoreAdminV1beta2IndexFieldOrderOutput {
	return o
}

func (o GoogleFirestoreAdminV1beta2IndexFieldOrderOutput) ToGoogleFirestoreAdminV1beta2IndexFieldOrderOutputWithContext(ctx context.Context) GoogleFirestoreAdminV1beta2IndexFieldOrderOutput {
	return o
}

func (o GoogleFirestoreAdminV1beta2IndexFieldOrderOutput) ToGoogleFirestoreAdminV1beta2IndexFieldOrderPtrOutput() GoogleFirestoreAdminV1beta2IndexFieldOrderPtrOutput {
	return o.ToGoogleFirestoreAdminV1beta2IndexFieldOrderPtrOutputWithContext(context.Background())
}

func (o GoogleFirestoreAdminV1beta2IndexFieldOrderOutput) ToGoogleFirestoreAdminV1beta2IndexFieldOrderPtrOutputWithContext(ctx context.Context) GoogleFirestoreAdminV1beta2IndexFieldOrderPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v GoogleFirestoreAdminV1beta2IndexFieldOrder) *GoogleFirestoreAdminV1beta2IndexFieldOrder {
		return &v
	}).(GoogleFirestoreAdminV1beta2IndexFieldOrderPtrOutput)
}

func (o GoogleFirestoreAdminV1beta2IndexFieldOrderOutput) ToStringOutput() pulumi.StringOutput {
	return o.ToStringOutputWithContext(context.Background())
}

func (o GoogleFirestoreAdminV1beta2IndexFieldOrderOutput) ToStringOutputWithContext(ctx context.Context) pulumi.StringOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e GoogleFirestoreAdminV1beta2IndexFieldOrder) string {
		return string(e)
	}).(pulumi.StringOutput)
}

func (o GoogleFirestoreAdminV1beta2IndexFieldOrderOutput) ToStringPtrOutput() pulumi.StringPtrOutput {
	return o.ToStringPtrOutputWithContext(context.Background())
}

func (o GoogleFirestoreAdminV1beta2IndexFieldOrderOutput) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e GoogleFirestoreAdminV1beta2IndexFieldOrder) *string {
		v := string(e)
		return &v
	}).(pulumi.StringPtrOutput)
}

type GoogleFirestoreAdminV1beta2IndexFieldOrderPtrOutput struct{ *pulumi.OutputState }

func (GoogleFirestoreAdminV1beta2IndexFieldOrderPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**GoogleFirestoreAdminV1beta2IndexFieldOrder)(nil)).Elem()
}

func (o GoogleFirestoreAdminV1beta2IndexFieldOrderPtrOutput) ToGoogleFirestoreAdminV1beta2IndexFieldOrderPtrOutput() GoogleFirestoreAdminV1beta2IndexFieldOrderPtrOutput {
	return o
}

func (o GoogleFirestoreAdminV1beta2IndexFieldOrderPtrOutput) ToGoogleFirestoreAdminV1beta2IndexFieldOrderPtrOutputWithContext(ctx context.Context) GoogleFirestoreAdminV1beta2IndexFieldOrderPtrOutput {
	return o
}

func (o GoogleFirestoreAdminV1beta2IndexFieldOrderPtrOutput) Elem() GoogleFirestoreAdminV1beta2IndexFieldOrderOutput {
	return o.ApplyT(func(v *GoogleFirestoreAdminV1beta2IndexFieldOrder) GoogleFirestoreAdminV1beta2IndexFieldOrder {
		if v != nil {
			return *v
		}
		var ret GoogleFirestoreAdminV1beta2IndexFieldOrder
		return ret
	}).(GoogleFirestoreAdminV1beta2IndexFieldOrderOutput)
}

func (o GoogleFirestoreAdminV1beta2IndexFieldOrderPtrOutput) ToStringPtrOutput() pulumi.StringPtrOutput {
	return o.ToStringPtrOutputWithContext(context.Background())
}

func (o GoogleFirestoreAdminV1beta2IndexFieldOrderPtrOutput) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e *GoogleFirestoreAdminV1beta2IndexFieldOrder) *string {
		if e == nil {
			return nil
		}
		v := string(*e)
		return &v
	}).(pulumi.StringPtrOutput)
}

// GoogleFirestoreAdminV1beta2IndexFieldOrderInput is an input type that accepts GoogleFirestoreAdminV1beta2IndexFieldOrderArgs and GoogleFirestoreAdminV1beta2IndexFieldOrderOutput values.
// You can construct a concrete instance of `GoogleFirestoreAdminV1beta2IndexFieldOrderInput` via:
//
//          GoogleFirestoreAdminV1beta2IndexFieldOrderArgs{...}
type GoogleFirestoreAdminV1beta2IndexFieldOrderInput interface {
	pulumi.Input

	ToGoogleFirestoreAdminV1beta2IndexFieldOrderOutput() GoogleFirestoreAdminV1beta2IndexFieldOrderOutput
	ToGoogleFirestoreAdminV1beta2IndexFieldOrderOutputWithContext(context.Context) GoogleFirestoreAdminV1beta2IndexFieldOrderOutput
}

var googleFirestoreAdminV1beta2IndexFieldOrderPtrType = reflect.TypeOf((**GoogleFirestoreAdminV1beta2IndexFieldOrder)(nil)).Elem()

type GoogleFirestoreAdminV1beta2IndexFieldOrderPtrInput interface {
	pulumi.Input

	ToGoogleFirestoreAdminV1beta2IndexFieldOrderPtrOutput() GoogleFirestoreAdminV1beta2IndexFieldOrderPtrOutput
	ToGoogleFirestoreAdminV1beta2IndexFieldOrderPtrOutputWithContext(context.Context) GoogleFirestoreAdminV1beta2IndexFieldOrderPtrOutput
}

type googleFirestoreAdminV1beta2IndexFieldOrderPtr string

func GoogleFirestoreAdminV1beta2IndexFieldOrderPtr(v string) GoogleFirestoreAdminV1beta2IndexFieldOrderPtrInput {
	return (*googleFirestoreAdminV1beta2IndexFieldOrderPtr)(&v)
}

func (*googleFirestoreAdminV1beta2IndexFieldOrderPtr) ElementType() reflect.Type {
	return googleFirestoreAdminV1beta2IndexFieldOrderPtrType
}

func (in *googleFirestoreAdminV1beta2IndexFieldOrderPtr) ToGoogleFirestoreAdminV1beta2IndexFieldOrderPtrOutput() GoogleFirestoreAdminV1beta2IndexFieldOrderPtrOutput {
	return pulumi.ToOutput(in).(GoogleFirestoreAdminV1beta2IndexFieldOrderPtrOutput)
}

func (in *googleFirestoreAdminV1beta2IndexFieldOrderPtr) ToGoogleFirestoreAdminV1beta2IndexFieldOrderPtrOutputWithContext(ctx context.Context) GoogleFirestoreAdminV1beta2IndexFieldOrderPtrOutput {
	return pulumi.ToOutputWithContext(ctx, in).(GoogleFirestoreAdminV1beta2IndexFieldOrderPtrOutput)
}

// Indexes with a collection query scope specified allow queries against a collection that is the child of a specific document, specified at query time, and that has the same collection id. Indexes with a collection group query scope specified allow queries against all collections descended from a specific document, specified at query time, and that have the same collection id as this index.
type IndexQueryScope string

const (
	// The query scope is unspecified. Not a valid option.
	IndexQueryScopeQueryScopeUnspecified = IndexQueryScope("QUERY_SCOPE_UNSPECIFIED")
	// Indexes with a collection query scope specified allow queries against a collection that is the child of a specific document, specified at query time, and that has the collection id specified by the index.
	IndexQueryScopeCollection = IndexQueryScope("COLLECTION")
	// Indexes with a collection group query scope specified allow queries against all collections that has the collection id specified by the index.
	IndexQueryScopeCollectionGroup = IndexQueryScope("COLLECTION_GROUP")
)

func (IndexQueryScope) ElementType() reflect.Type {
	return reflect.TypeOf((*IndexQueryScope)(nil)).Elem()
}

func (e IndexQueryScope) ToIndexQueryScopeOutput() IndexQueryScopeOutput {
	return pulumi.ToOutput(e).(IndexQueryScopeOutput)
}

func (e IndexQueryScope) ToIndexQueryScopeOutputWithContext(ctx context.Context) IndexQueryScopeOutput {
	return pulumi.ToOutputWithContext(ctx, e).(IndexQueryScopeOutput)
}

func (e IndexQueryScope) ToIndexQueryScopePtrOutput() IndexQueryScopePtrOutput {
	return e.ToIndexQueryScopePtrOutputWithContext(context.Background())
}

func (e IndexQueryScope) ToIndexQueryScopePtrOutputWithContext(ctx context.Context) IndexQueryScopePtrOutput {
	return IndexQueryScope(e).ToIndexQueryScopeOutputWithContext(ctx).ToIndexQueryScopePtrOutputWithContext(ctx)
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

type IndexQueryScopeOutput struct{ *pulumi.OutputState }

func (IndexQueryScopeOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*IndexQueryScope)(nil)).Elem()
}

func (o IndexQueryScopeOutput) ToIndexQueryScopeOutput() IndexQueryScopeOutput {
	return o
}

func (o IndexQueryScopeOutput) ToIndexQueryScopeOutputWithContext(ctx context.Context) IndexQueryScopeOutput {
	return o
}

func (o IndexQueryScopeOutput) ToIndexQueryScopePtrOutput() IndexQueryScopePtrOutput {
	return o.ToIndexQueryScopePtrOutputWithContext(context.Background())
}

func (o IndexQueryScopeOutput) ToIndexQueryScopePtrOutputWithContext(ctx context.Context) IndexQueryScopePtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v IndexQueryScope) *IndexQueryScope {
		return &v
	}).(IndexQueryScopePtrOutput)
}

func (o IndexQueryScopeOutput) ToStringOutput() pulumi.StringOutput {
	return o.ToStringOutputWithContext(context.Background())
}

func (o IndexQueryScopeOutput) ToStringOutputWithContext(ctx context.Context) pulumi.StringOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e IndexQueryScope) string {
		return string(e)
	}).(pulumi.StringOutput)
}

func (o IndexQueryScopeOutput) ToStringPtrOutput() pulumi.StringPtrOutput {
	return o.ToStringPtrOutputWithContext(context.Background())
}

func (o IndexQueryScopeOutput) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e IndexQueryScope) *string {
		v := string(e)
		return &v
	}).(pulumi.StringPtrOutput)
}

type IndexQueryScopePtrOutput struct{ *pulumi.OutputState }

func (IndexQueryScopePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**IndexQueryScope)(nil)).Elem()
}

func (o IndexQueryScopePtrOutput) ToIndexQueryScopePtrOutput() IndexQueryScopePtrOutput {
	return o
}

func (o IndexQueryScopePtrOutput) ToIndexQueryScopePtrOutputWithContext(ctx context.Context) IndexQueryScopePtrOutput {
	return o
}

func (o IndexQueryScopePtrOutput) Elem() IndexQueryScopeOutput {
	return o.ApplyT(func(v *IndexQueryScope) IndexQueryScope {
		if v != nil {
			return *v
		}
		var ret IndexQueryScope
		return ret
	}).(IndexQueryScopeOutput)
}

func (o IndexQueryScopePtrOutput) ToStringPtrOutput() pulumi.StringPtrOutput {
	return o.ToStringPtrOutputWithContext(context.Background())
}

func (o IndexQueryScopePtrOutput) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e *IndexQueryScope) *string {
		if e == nil {
			return nil
		}
		v := string(*e)
		return &v
	}).(pulumi.StringPtrOutput)
}

// IndexQueryScopeInput is an input type that accepts IndexQueryScopeArgs and IndexQueryScopeOutput values.
// You can construct a concrete instance of `IndexQueryScopeInput` via:
//
//          IndexQueryScopeArgs{...}
type IndexQueryScopeInput interface {
	pulumi.Input

	ToIndexQueryScopeOutput() IndexQueryScopeOutput
	ToIndexQueryScopeOutputWithContext(context.Context) IndexQueryScopeOutput
}

var indexQueryScopePtrType = reflect.TypeOf((**IndexQueryScope)(nil)).Elem()

type IndexQueryScopePtrInput interface {
	pulumi.Input

	ToIndexQueryScopePtrOutput() IndexQueryScopePtrOutput
	ToIndexQueryScopePtrOutputWithContext(context.Context) IndexQueryScopePtrOutput
}

type indexQueryScopePtr string

func IndexQueryScopePtr(v string) IndexQueryScopePtrInput {
	return (*indexQueryScopePtr)(&v)
}

func (*indexQueryScopePtr) ElementType() reflect.Type {
	return indexQueryScopePtrType
}

func (in *indexQueryScopePtr) ToIndexQueryScopePtrOutput() IndexQueryScopePtrOutput {
	return pulumi.ToOutput(in).(IndexQueryScopePtrOutput)
}

func (in *indexQueryScopePtr) ToIndexQueryScopePtrOutputWithContext(ctx context.Context) IndexQueryScopePtrOutput {
	return pulumi.ToOutputWithContext(ctx, in).(IndexQueryScopePtrOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*GoogleFirestoreAdminV1beta2IndexFieldArrayConfigInput)(nil)).Elem(), GoogleFirestoreAdminV1beta2IndexFieldArrayConfig("ARRAY_CONFIG_UNSPECIFIED"))
	pulumi.RegisterInputType(reflect.TypeOf((*GoogleFirestoreAdminV1beta2IndexFieldArrayConfigPtrInput)(nil)).Elem(), GoogleFirestoreAdminV1beta2IndexFieldArrayConfig("ARRAY_CONFIG_UNSPECIFIED"))
	pulumi.RegisterInputType(reflect.TypeOf((*GoogleFirestoreAdminV1beta2IndexFieldOrderInput)(nil)).Elem(), GoogleFirestoreAdminV1beta2IndexFieldOrder("ORDER_UNSPECIFIED"))
	pulumi.RegisterInputType(reflect.TypeOf((*GoogleFirestoreAdminV1beta2IndexFieldOrderPtrInput)(nil)).Elem(), GoogleFirestoreAdminV1beta2IndexFieldOrder("ORDER_UNSPECIFIED"))
	pulumi.RegisterInputType(reflect.TypeOf((*IndexQueryScopeInput)(nil)).Elem(), IndexQueryScope("QUERY_SCOPE_UNSPECIFIED"))
	pulumi.RegisterInputType(reflect.TypeOf((*IndexQueryScopePtrInput)(nil)).Elem(), IndexQueryScope("QUERY_SCOPE_UNSPECIFIED"))
	pulumi.RegisterOutputType(GoogleFirestoreAdminV1beta2IndexFieldArrayConfigOutput{})
	pulumi.RegisterOutputType(GoogleFirestoreAdminV1beta2IndexFieldArrayConfigPtrOutput{})
	pulumi.RegisterOutputType(GoogleFirestoreAdminV1beta2IndexFieldOrderOutput{})
	pulumi.RegisterOutputType(GoogleFirestoreAdminV1beta2IndexFieldOrderPtrOutput{})
	pulumi.RegisterOutputType(IndexQueryScopeOutput{})
	pulumi.RegisterOutputType(IndexQueryScopePtrOutput{})
}
