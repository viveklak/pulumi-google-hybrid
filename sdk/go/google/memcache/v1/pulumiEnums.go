// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package v1

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// The major version of Memcached software. If not provided, latest supported version will be used. Currently the latest supported major version is `MEMCACHE_1_5`. The minor version will be automatically determined by our system based on the latest supported minor version.
type InstanceMemcacheVersion string

const (
	InstanceMemcacheVersionMemcacheVersionUnspecified = InstanceMemcacheVersion("MEMCACHE_VERSION_UNSPECIFIED")
	// Memcached 1.5 version.
	InstanceMemcacheVersionMemcache15 = InstanceMemcacheVersion("MEMCACHE_1_5")
)

func (InstanceMemcacheVersion) ElementType() reflect.Type {
	return reflect.TypeOf((*InstanceMemcacheVersion)(nil)).Elem()
}

func (e InstanceMemcacheVersion) ToInstanceMemcacheVersionOutput() InstanceMemcacheVersionOutput {
	return pulumi.ToOutput(e).(InstanceMemcacheVersionOutput)
}

func (e InstanceMemcacheVersion) ToInstanceMemcacheVersionOutputWithContext(ctx context.Context) InstanceMemcacheVersionOutput {
	return pulumi.ToOutputWithContext(ctx, e).(InstanceMemcacheVersionOutput)
}

func (e InstanceMemcacheVersion) ToInstanceMemcacheVersionPtrOutput() InstanceMemcacheVersionPtrOutput {
	return e.ToInstanceMemcacheVersionPtrOutputWithContext(context.Background())
}

func (e InstanceMemcacheVersion) ToInstanceMemcacheVersionPtrOutputWithContext(ctx context.Context) InstanceMemcacheVersionPtrOutput {
	return InstanceMemcacheVersion(e).ToInstanceMemcacheVersionOutputWithContext(ctx).ToInstanceMemcacheVersionPtrOutputWithContext(ctx)
}

func (e InstanceMemcacheVersion) ToStringOutput() pulumi.StringOutput {
	return pulumi.ToOutput(pulumi.String(e)).(pulumi.StringOutput)
}

func (e InstanceMemcacheVersion) ToStringOutputWithContext(ctx context.Context) pulumi.StringOutput {
	return pulumi.ToOutputWithContext(ctx, pulumi.String(e)).(pulumi.StringOutput)
}

func (e InstanceMemcacheVersion) ToStringPtrOutput() pulumi.StringPtrOutput {
	return pulumi.String(e).ToStringPtrOutputWithContext(context.Background())
}

func (e InstanceMemcacheVersion) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return pulumi.String(e).ToStringOutputWithContext(ctx).ToStringPtrOutputWithContext(ctx)
}

type InstanceMemcacheVersionOutput struct{ *pulumi.OutputState }

func (InstanceMemcacheVersionOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*InstanceMemcacheVersion)(nil)).Elem()
}

func (o InstanceMemcacheVersionOutput) ToInstanceMemcacheVersionOutput() InstanceMemcacheVersionOutput {
	return o
}

func (o InstanceMemcacheVersionOutput) ToInstanceMemcacheVersionOutputWithContext(ctx context.Context) InstanceMemcacheVersionOutput {
	return o
}

func (o InstanceMemcacheVersionOutput) ToInstanceMemcacheVersionPtrOutput() InstanceMemcacheVersionPtrOutput {
	return o.ToInstanceMemcacheVersionPtrOutputWithContext(context.Background())
}

func (o InstanceMemcacheVersionOutput) ToInstanceMemcacheVersionPtrOutputWithContext(ctx context.Context) InstanceMemcacheVersionPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v InstanceMemcacheVersion) *InstanceMemcacheVersion {
		return &v
	}).(InstanceMemcacheVersionPtrOutput)
}

func (o InstanceMemcacheVersionOutput) ToStringOutput() pulumi.StringOutput {
	return o.ToStringOutputWithContext(context.Background())
}

func (o InstanceMemcacheVersionOutput) ToStringOutputWithContext(ctx context.Context) pulumi.StringOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e InstanceMemcacheVersion) string {
		return string(e)
	}).(pulumi.StringOutput)
}

func (o InstanceMemcacheVersionOutput) ToStringPtrOutput() pulumi.StringPtrOutput {
	return o.ToStringPtrOutputWithContext(context.Background())
}

func (o InstanceMemcacheVersionOutput) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e InstanceMemcacheVersion) *string {
		v := string(e)
		return &v
	}).(pulumi.StringPtrOutput)
}

type InstanceMemcacheVersionPtrOutput struct{ *pulumi.OutputState }

func (InstanceMemcacheVersionPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**InstanceMemcacheVersion)(nil)).Elem()
}

func (o InstanceMemcacheVersionPtrOutput) ToInstanceMemcacheVersionPtrOutput() InstanceMemcacheVersionPtrOutput {
	return o
}

func (o InstanceMemcacheVersionPtrOutput) ToInstanceMemcacheVersionPtrOutputWithContext(ctx context.Context) InstanceMemcacheVersionPtrOutput {
	return o
}

func (o InstanceMemcacheVersionPtrOutput) Elem() InstanceMemcacheVersionOutput {
	return o.ApplyT(func(v *InstanceMemcacheVersion) InstanceMemcacheVersion {
		if v != nil {
			return *v
		}
		var ret InstanceMemcacheVersion
		return ret
	}).(InstanceMemcacheVersionOutput)
}

func (o InstanceMemcacheVersionPtrOutput) ToStringPtrOutput() pulumi.StringPtrOutput {
	return o.ToStringPtrOutputWithContext(context.Background())
}

func (o InstanceMemcacheVersionPtrOutput) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e *InstanceMemcacheVersion) *string {
		if e == nil {
			return nil
		}
		v := string(*e)
		return &v
	}).(pulumi.StringPtrOutput)
}

// InstanceMemcacheVersionInput is an input type that accepts InstanceMemcacheVersionArgs and InstanceMemcacheVersionOutput values.
// You can construct a concrete instance of `InstanceMemcacheVersionInput` via:
//
//          InstanceMemcacheVersionArgs{...}
type InstanceMemcacheVersionInput interface {
	pulumi.Input

	ToInstanceMemcacheVersionOutput() InstanceMemcacheVersionOutput
	ToInstanceMemcacheVersionOutputWithContext(context.Context) InstanceMemcacheVersionOutput
}

var instanceMemcacheVersionPtrType = reflect.TypeOf((**InstanceMemcacheVersion)(nil)).Elem()

type InstanceMemcacheVersionPtrInput interface {
	pulumi.Input

	ToInstanceMemcacheVersionPtrOutput() InstanceMemcacheVersionPtrOutput
	ToInstanceMemcacheVersionPtrOutputWithContext(context.Context) InstanceMemcacheVersionPtrOutput
}

type instanceMemcacheVersionPtr string

func InstanceMemcacheVersionPtr(v string) InstanceMemcacheVersionPtrInput {
	return (*instanceMemcacheVersionPtr)(&v)
}

func (*instanceMemcacheVersionPtr) ElementType() reflect.Type {
	return instanceMemcacheVersionPtrType
}

func (in *instanceMemcacheVersionPtr) ToInstanceMemcacheVersionPtrOutput() InstanceMemcacheVersionPtrOutput {
	return pulumi.ToOutput(in).(InstanceMemcacheVersionPtrOutput)
}

func (in *instanceMemcacheVersionPtr) ToInstanceMemcacheVersionPtrOutputWithContext(ctx context.Context) InstanceMemcacheVersionPtrOutput {
	return pulumi.ToOutputWithContext(ctx, in).(InstanceMemcacheVersionPtrOutput)
}

// A code that correspond to one type of user-facing message.
type InstanceMessageCode string

const (
	// Message Code not set.
	InstanceMessageCodeCodeUnspecified = InstanceMessageCode("CODE_UNSPECIFIED")
	// Memcached nodes are distributed unevenly.
	InstanceMessageCodeZoneDistributionUnbalanced = InstanceMessageCode("ZONE_DISTRIBUTION_UNBALANCED")
)

func (InstanceMessageCode) ElementType() reflect.Type {
	return reflect.TypeOf((*InstanceMessageCode)(nil)).Elem()
}

func (e InstanceMessageCode) ToInstanceMessageCodeOutput() InstanceMessageCodeOutput {
	return pulumi.ToOutput(e).(InstanceMessageCodeOutput)
}

func (e InstanceMessageCode) ToInstanceMessageCodeOutputWithContext(ctx context.Context) InstanceMessageCodeOutput {
	return pulumi.ToOutputWithContext(ctx, e).(InstanceMessageCodeOutput)
}

func (e InstanceMessageCode) ToInstanceMessageCodePtrOutput() InstanceMessageCodePtrOutput {
	return e.ToInstanceMessageCodePtrOutputWithContext(context.Background())
}

func (e InstanceMessageCode) ToInstanceMessageCodePtrOutputWithContext(ctx context.Context) InstanceMessageCodePtrOutput {
	return InstanceMessageCode(e).ToInstanceMessageCodeOutputWithContext(ctx).ToInstanceMessageCodePtrOutputWithContext(ctx)
}

func (e InstanceMessageCode) ToStringOutput() pulumi.StringOutput {
	return pulumi.ToOutput(pulumi.String(e)).(pulumi.StringOutput)
}

func (e InstanceMessageCode) ToStringOutputWithContext(ctx context.Context) pulumi.StringOutput {
	return pulumi.ToOutputWithContext(ctx, pulumi.String(e)).(pulumi.StringOutput)
}

func (e InstanceMessageCode) ToStringPtrOutput() pulumi.StringPtrOutput {
	return pulumi.String(e).ToStringPtrOutputWithContext(context.Background())
}

func (e InstanceMessageCode) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return pulumi.String(e).ToStringOutputWithContext(ctx).ToStringPtrOutputWithContext(ctx)
}

type InstanceMessageCodeOutput struct{ *pulumi.OutputState }

func (InstanceMessageCodeOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*InstanceMessageCode)(nil)).Elem()
}

func (o InstanceMessageCodeOutput) ToInstanceMessageCodeOutput() InstanceMessageCodeOutput {
	return o
}

func (o InstanceMessageCodeOutput) ToInstanceMessageCodeOutputWithContext(ctx context.Context) InstanceMessageCodeOutput {
	return o
}

func (o InstanceMessageCodeOutput) ToInstanceMessageCodePtrOutput() InstanceMessageCodePtrOutput {
	return o.ToInstanceMessageCodePtrOutputWithContext(context.Background())
}

func (o InstanceMessageCodeOutput) ToInstanceMessageCodePtrOutputWithContext(ctx context.Context) InstanceMessageCodePtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v InstanceMessageCode) *InstanceMessageCode {
		return &v
	}).(InstanceMessageCodePtrOutput)
}

func (o InstanceMessageCodeOutput) ToStringOutput() pulumi.StringOutput {
	return o.ToStringOutputWithContext(context.Background())
}

func (o InstanceMessageCodeOutput) ToStringOutputWithContext(ctx context.Context) pulumi.StringOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e InstanceMessageCode) string {
		return string(e)
	}).(pulumi.StringOutput)
}

func (o InstanceMessageCodeOutput) ToStringPtrOutput() pulumi.StringPtrOutput {
	return o.ToStringPtrOutputWithContext(context.Background())
}

func (o InstanceMessageCodeOutput) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e InstanceMessageCode) *string {
		v := string(e)
		return &v
	}).(pulumi.StringPtrOutput)
}

type InstanceMessageCodePtrOutput struct{ *pulumi.OutputState }

func (InstanceMessageCodePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**InstanceMessageCode)(nil)).Elem()
}

func (o InstanceMessageCodePtrOutput) ToInstanceMessageCodePtrOutput() InstanceMessageCodePtrOutput {
	return o
}

func (o InstanceMessageCodePtrOutput) ToInstanceMessageCodePtrOutputWithContext(ctx context.Context) InstanceMessageCodePtrOutput {
	return o
}

func (o InstanceMessageCodePtrOutput) Elem() InstanceMessageCodeOutput {
	return o.ApplyT(func(v *InstanceMessageCode) InstanceMessageCode {
		if v != nil {
			return *v
		}
		var ret InstanceMessageCode
		return ret
	}).(InstanceMessageCodeOutput)
}

func (o InstanceMessageCodePtrOutput) ToStringPtrOutput() pulumi.StringPtrOutput {
	return o.ToStringPtrOutputWithContext(context.Background())
}

func (o InstanceMessageCodePtrOutput) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e *InstanceMessageCode) *string {
		if e == nil {
			return nil
		}
		v := string(*e)
		return &v
	}).(pulumi.StringPtrOutput)
}

// InstanceMessageCodeInput is an input type that accepts InstanceMessageCodeArgs and InstanceMessageCodeOutput values.
// You can construct a concrete instance of `InstanceMessageCodeInput` via:
//
//          InstanceMessageCodeArgs{...}
type InstanceMessageCodeInput interface {
	pulumi.Input

	ToInstanceMessageCodeOutput() InstanceMessageCodeOutput
	ToInstanceMessageCodeOutputWithContext(context.Context) InstanceMessageCodeOutput
}

var instanceMessageCodePtrType = reflect.TypeOf((**InstanceMessageCode)(nil)).Elem()

type InstanceMessageCodePtrInput interface {
	pulumi.Input

	ToInstanceMessageCodePtrOutput() InstanceMessageCodePtrOutput
	ToInstanceMessageCodePtrOutputWithContext(context.Context) InstanceMessageCodePtrOutput
}

type instanceMessageCodePtr string

func InstanceMessageCodePtr(v string) InstanceMessageCodePtrInput {
	return (*instanceMessageCodePtr)(&v)
}

func (*instanceMessageCodePtr) ElementType() reflect.Type {
	return instanceMessageCodePtrType
}

func (in *instanceMessageCodePtr) ToInstanceMessageCodePtrOutput() InstanceMessageCodePtrOutput {
	return pulumi.ToOutput(in).(InstanceMessageCodePtrOutput)
}

func (in *instanceMessageCodePtr) ToInstanceMessageCodePtrOutputWithContext(ctx context.Context) InstanceMessageCodePtrOutput {
	return pulumi.ToOutputWithContext(ctx, in).(InstanceMessageCodePtrOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*InstanceMemcacheVersionInput)(nil)).Elem(), InstanceMemcacheVersion("MEMCACHE_VERSION_UNSPECIFIED"))
	pulumi.RegisterInputType(reflect.TypeOf((*InstanceMemcacheVersionPtrInput)(nil)).Elem(), InstanceMemcacheVersion("MEMCACHE_VERSION_UNSPECIFIED"))
	pulumi.RegisterInputType(reflect.TypeOf((*InstanceMessageCodeInput)(nil)).Elem(), InstanceMessageCode("CODE_UNSPECIFIED"))
	pulumi.RegisterInputType(reflect.TypeOf((*InstanceMessageCodePtrInput)(nil)).Elem(), InstanceMessageCode("CODE_UNSPECIFIED"))
	pulumi.RegisterOutputType(InstanceMemcacheVersionOutput{})
	pulumi.RegisterOutputType(InstanceMemcacheVersionPtrOutput{})
	pulumi.RegisterOutputType(InstanceMessageCodeOutput{})
	pulumi.RegisterOutputType(InstanceMessageCodePtrOutput{})
}
