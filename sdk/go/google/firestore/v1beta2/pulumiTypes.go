// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package v1beta2

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// A field in an index. The field_path describes which field is indexed, the value_mode describes how the field value is indexed.
type GoogleFirestoreAdminV1beta2IndexField struct {
	// Indicates that this field supports operations on `array_value`s.
	ArrayConfig *GoogleFirestoreAdminV1beta2IndexFieldArrayConfig `pulumi:"arrayConfig"`
	// Can be __name__. For single field indexes, this must match the name of the field or may be omitted.
	FieldPath *string `pulumi:"fieldPath"`
	// Indicates that this field supports ordering by the specified order or comparing using =, <, <=, >, >=.
	Order *GoogleFirestoreAdminV1beta2IndexFieldOrder `pulumi:"order"`
}

// GoogleFirestoreAdminV1beta2IndexFieldInput is an input type that accepts GoogleFirestoreAdminV1beta2IndexFieldArgs and GoogleFirestoreAdminV1beta2IndexFieldOutput values.
// You can construct a concrete instance of `GoogleFirestoreAdminV1beta2IndexFieldInput` via:
//
//          GoogleFirestoreAdminV1beta2IndexFieldArgs{...}
type GoogleFirestoreAdminV1beta2IndexFieldInput interface {
	pulumi.Input

	ToGoogleFirestoreAdminV1beta2IndexFieldOutput() GoogleFirestoreAdminV1beta2IndexFieldOutput
	ToGoogleFirestoreAdminV1beta2IndexFieldOutputWithContext(context.Context) GoogleFirestoreAdminV1beta2IndexFieldOutput
}

// A field in an index. The field_path describes which field is indexed, the value_mode describes how the field value is indexed.
type GoogleFirestoreAdminV1beta2IndexFieldArgs struct {
	// Indicates that this field supports operations on `array_value`s.
	ArrayConfig GoogleFirestoreAdminV1beta2IndexFieldArrayConfigPtrInput `pulumi:"arrayConfig"`
	// Can be __name__. For single field indexes, this must match the name of the field or may be omitted.
	FieldPath pulumi.StringPtrInput `pulumi:"fieldPath"`
	// Indicates that this field supports ordering by the specified order or comparing using =, <, <=, >, >=.
	Order GoogleFirestoreAdminV1beta2IndexFieldOrderPtrInput `pulumi:"order"`
}

func (GoogleFirestoreAdminV1beta2IndexFieldArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*GoogleFirestoreAdminV1beta2IndexField)(nil)).Elem()
}

func (i GoogleFirestoreAdminV1beta2IndexFieldArgs) ToGoogleFirestoreAdminV1beta2IndexFieldOutput() GoogleFirestoreAdminV1beta2IndexFieldOutput {
	return i.ToGoogleFirestoreAdminV1beta2IndexFieldOutputWithContext(context.Background())
}

func (i GoogleFirestoreAdminV1beta2IndexFieldArgs) ToGoogleFirestoreAdminV1beta2IndexFieldOutputWithContext(ctx context.Context) GoogleFirestoreAdminV1beta2IndexFieldOutput {
	return pulumi.ToOutputWithContext(ctx, i).(GoogleFirestoreAdminV1beta2IndexFieldOutput)
}

// GoogleFirestoreAdminV1beta2IndexFieldArrayInput is an input type that accepts GoogleFirestoreAdminV1beta2IndexFieldArray and GoogleFirestoreAdminV1beta2IndexFieldArrayOutput values.
// You can construct a concrete instance of `GoogleFirestoreAdminV1beta2IndexFieldArrayInput` via:
//
//          GoogleFirestoreAdminV1beta2IndexFieldArray{ GoogleFirestoreAdminV1beta2IndexFieldArgs{...} }
type GoogleFirestoreAdminV1beta2IndexFieldArrayInput interface {
	pulumi.Input

	ToGoogleFirestoreAdminV1beta2IndexFieldArrayOutput() GoogleFirestoreAdminV1beta2IndexFieldArrayOutput
	ToGoogleFirestoreAdminV1beta2IndexFieldArrayOutputWithContext(context.Context) GoogleFirestoreAdminV1beta2IndexFieldArrayOutput
}

type GoogleFirestoreAdminV1beta2IndexFieldArray []GoogleFirestoreAdminV1beta2IndexFieldInput

func (GoogleFirestoreAdminV1beta2IndexFieldArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]GoogleFirestoreAdminV1beta2IndexField)(nil)).Elem()
}

func (i GoogleFirestoreAdminV1beta2IndexFieldArray) ToGoogleFirestoreAdminV1beta2IndexFieldArrayOutput() GoogleFirestoreAdminV1beta2IndexFieldArrayOutput {
	return i.ToGoogleFirestoreAdminV1beta2IndexFieldArrayOutputWithContext(context.Background())
}

func (i GoogleFirestoreAdminV1beta2IndexFieldArray) ToGoogleFirestoreAdminV1beta2IndexFieldArrayOutputWithContext(ctx context.Context) GoogleFirestoreAdminV1beta2IndexFieldArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(GoogleFirestoreAdminV1beta2IndexFieldArrayOutput)
}

// A field in an index. The field_path describes which field is indexed, the value_mode describes how the field value is indexed.
type GoogleFirestoreAdminV1beta2IndexFieldOutput struct{ *pulumi.OutputState }

func (GoogleFirestoreAdminV1beta2IndexFieldOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*GoogleFirestoreAdminV1beta2IndexField)(nil)).Elem()
}

func (o GoogleFirestoreAdminV1beta2IndexFieldOutput) ToGoogleFirestoreAdminV1beta2IndexFieldOutput() GoogleFirestoreAdminV1beta2IndexFieldOutput {
	return o
}

func (o GoogleFirestoreAdminV1beta2IndexFieldOutput) ToGoogleFirestoreAdminV1beta2IndexFieldOutputWithContext(ctx context.Context) GoogleFirestoreAdminV1beta2IndexFieldOutput {
	return o
}

// Indicates that this field supports operations on `array_value`s.
func (o GoogleFirestoreAdminV1beta2IndexFieldOutput) ArrayConfig() GoogleFirestoreAdminV1beta2IndexFieldArrayConfigPtrOutput {
	return o.ApplyT(func(v GoogleFirestoreAdminV1beta2IndexField) *GoogleFirestoreAdminV1beta2IndexFieldArrayConfig {
		return v.ArrayConfig
	}).(GoogleFirestoreAdminV1beta2IndexFieldArrayConfigPtrOutput)
}

// Can be __name__. For single field indexes, this must match the name of the field or may be omitted.
func (o GoogleFirestoreAdminV1beta2IndexFieldOutput) FieldPath() pulumi.StringPtrOutput {
	return o.ApplyT(func(v GoogleFirestoreAdminV1beta2IndexField) *string { return v.FieldPath }).(pulumi.StringPtrOutput)
}

// Indicates that this field supports ordering by the specified order or comparing using =, <, <=, >, >=.
func (o GoogleFirestoreAdminV1beta2IndexFieldOutput) Order() GoogleFirestoreAdminV1beta2IndexFieldOrderPtrOutput {
	return o.ApplyT(func(v GoogleFirestoreAdminV1beta2IndexField) *GoogleFirestoreAdminV1beta2IndexFieldOrder {
		return v.Order
	}).(GoogleFirestoreAdminV1beta2IndexFieldOrderPtrOutput)
}

type GoogleFirestoreAdminV1beta2IndexFieldArrayOutput struct{ *pulumi.OutputState }

func (GoogleFirestoreAdminV1beta2IndexFieldArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]GoogleFirestoreAdminV1beta2IndexField)(nil)).Elem()
}

func (o GoogleFirestoreAdminV1beta2IndexFieldArrayOutput) ToGoogleFirestoreAdminV1beta2IndexFieldArrayOutput() GoogleFirestoreAdminV1beta2IndexFieldArrayOutput {
	return o
}

func (o GoogleFirestoreAdminV1beta2IndexFieldArrayOutput) ToGoogleFirestoreAdminV1beta2IndexFieldArrayOutputWithContext(ctx context.Context) GoogleFirestoreAdminV1beta2IndexFieldArrayOutput {
	return o
}

func (o GoogleFirestoreAdminV1beta2IndexFieldArrayOutput) Index(i pulumi.IntInput) GoogleFirestoreAdminV1beta2IndexFieldOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) GoogleFirestoreAdminV1beta2IndexField {
		return vs[0].([]GoogleFirestoreAdminV1beta2IndexField)[vs[1].(int)]
	}).(GoogleFirestoreAdminV1beta2IndexFieldOutput)
}

// A field in an index. The field_path describes which field is indexed, the value_mode describes how the field value is indexed.
type GoogleFirestoreAdminV1beta2IndexFieldResponse struct {
	// Indicates that this field supports operations on `array_value`s.
	ArrayConfig string `pulumi:"arrayConfig"`
	// Can be __name__. For single field indexes, this must match the name of the field or may be omitted.
	FieldPath string `pulumi:"fieldPath"`
	// Indicates that this field supports ordering by the specified order or comparing using =, <, <=, >, >=.
	Order string `pulumi:"order"`
}

// A field in an index. The field_path describes which field is indexed, the value_mode describes how the field value is indexed.
type GoogleFirestoreAdminV1beta2IndexFieldResponseOutput struct{ *pulumi.OutputState }

func (GoogleFirestoreAdminV1beta2IndexFieldResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*GoogleFirestoreAdminV1beta2IndexFieldResponse)(nil)).Elem()
}

func (o GoogleFirestoreAdminV1beta2IndexFieldResponseOutput) ToGoogleFirestoreAdminV1beta2IndexFieldResponseOutput() GoogleFirestoreAdminV1beta2IndexFieldResponseOutput {
	return o
}

func (o GoogleFirestoreAdminV1beta2IndexFieldResponseOutput) ToGoogleFirestoreAdminV1beta2IndexFieldResponseOutputWithContext(ctx context.Context) GoogleFirestoreAdminV1beta2IndexFieldResponseOutput {
	return o
}

// Indicates that this field supports operations on `array_value`s.
func (o GoogleFirestoreAdminV1beta2IndexFieldResponseOutput) ArrayConfig() pulumi.StringOutput {
	return o.ApplyT(func(v GoogleFirestoreAdminV1beta2IndexFieldResponse) string { return v.ArrayConfig }).(pulumi.StringOutput)
}

// Can be __name__. For single field indexes, this must match the name of the field or may be omitted.
func (o GoogleFirestoreAdminV1beta2IndexFieldResponseOutput) FieldPath() pulumi.StringOutput {
	return o.ApplyT(func(v GoogleFirestoreAdminV1beta2IndexFieldResponse) string { return v.FieldPath }).(pulumi.StringOutput)
}

// Indicates that this field supports ordering by the specified order or comparing using =, <, <=, >, >=.
func (o GoogleFirestoreAdminV1beta2IndexFieldResponseOutput) Order() pulumi.StringOutput {
	return o.ApplyT(func(v GoogleFirestoreAdminV1beta2IndexFieldResponse) string { return v.Order }).(pulumi.StringOutput)
}

type GoogleFirestoreAdminV1beta2IndexFieldResponseArrayOutput struct{ *pulumi.OutputState }

func (GoogleFirestoreAdminV1beta2IndexFieldResponseArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]GoogleFirestoreAdminV1beta2IndexFieldResponse)(nil)).Elem()
}

func (o GoogleFirestoreAdminV1beta2IndexFieldResponseArrayOutput) ToGoogleFirestoreAdminV1beta2IndexFieldResponseArrayOutput() GoogleFirestoreAdminV1beta2IndexFieldResponseArrayOutput {
	return o
}

func (o GoogleFirestoreAdminV1beta2IndexFieldResponseArrayOutput) ToGoogleFirestoreAdminV1beta2IndexFieldResponseArrayOutputWithContext(ctx context.Context) GoogleFirestoreAdminV1beta2IndexFieldResponseArrayOutput {
	return o
}

func (o GoogleFirestoreAdminV1beta2IndexFieldResponseArrayOutput) Index(i pulumi.IntInput) GoogleFirestoreAdminV1beta2IndexFieldResponseOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) GoogleFirestoreAdminV1beta2IndexFieldResponse {
		return vs[0].([]GoogleFirestoreAdminV1beta2IndexFieldResponse)[vs[1].(int)]
	}).(GoogleFirestoreAdminV1beta2IndexFieldResponseOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*GoogleFirestoreAdminV1beta2IndexFieldInput)(nil)).Elem(), GoogleFirestoreAdminV1beta2IndexFieldArgs{})
	pulumi.RegisterInputType(reflect.TypeOf((*GoogleFirestoreAdminV1beta2IndexFieldArrayInput)(nil)).Elem(), GoogleFirestoreAdminV1beta2IndexFieldArray{})
	pulumi.RegisterOutputType(GoogleFirestoreAdminV1beta2IndexFieldOutput{})
	pulumi.RegisterOutputType(GoogleFirestoreAdminV1beta2IndexFieldArrayOutput{})
	pulumi.RegisterOutputType(GoogleFirestoreAdminV1beta2IndexFieldResponseOutput{})
	pulumi.RegisterOutputType(GoogleFirestoreAdminV1beta2IndexFieldResponseArrayOutput{})
}
