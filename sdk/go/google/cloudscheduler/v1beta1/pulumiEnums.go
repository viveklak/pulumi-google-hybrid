// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package v1beta1

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// The HTTP method to use for the request. PATCH and OPTIONS are not permitted.
type AppEngineHttpTargetHttpMethod string

const (
	// HTTP method unspecified. Defaults to POST.
	AppEngineHttpTargetHttpMethodHttpMethodUnspecified = AppEngineHttpTargetHttpMethod("HTTP_METHOD_UNSPECIFIED")
	// HTTP POST
	AppEngineHttpTargetHttpMethodPost = AppEngineHttpTargetHttpMethod("POST")
	// HTTP GET
	AppEngineHttpTargetHttpMethodGet = AppEngineHttpTargetHttpMethod("GET")
	// HTTP HEAD
	AppEngineHttpTargetHttpMethodHead = AppEngineHttpTargetHttpMethod("HEAD")
	// HTTP PUT
	AppEngineHttpTargetHttpMethodPut = AppEngineHttpTargetHttpMethod("PUT")
	// HTTP DELETE
	AppEngineHttpTargetHttpMethodDelete = AppEngineHttpTargetHttpMethod("DELETE")
	// HTTP PATCH
	AppEngineHttpTargetHttpMethodPatch = AppEngineHttpTargetHttpMethod("PATCH")
	// HTTP OPTIONS
	AppEngineHttpTargetHttpMethodOptions = AppEngineHttpTargetHttpMethod("OPTIONS")
)

func (AppEngineHttpTargetHttpMethod) ElementType() reflect.Type {
	return reflect.TypeOf((*AppEngineHttpTargetHttpMethod)(nil)).Elem()
}

func (e AppEngineHttpTargetHttpMethod) ToAppEngineHttpTargetHttpMethodOutput() AppEngineHttpTargetHttpMethodOutput {
	return pulumi.ToOutput(e).(AppEngineHttpTargetHttpMethodOutput)
}

func (e AppEngineHttpTargetHttpMethod) ToAppEngineHttpTargetHttpMethodOutputWithContext(ctx context.Context) AppEngineHttpTargetHttpMethodOutput {
	return pulumi.ToOutputWithContext(ctx, e).(AppEngineHttpTargetHttpMethodOutput)
}

func (e AppEngineHttpTargetHttpMethod) ToAppEngineHttpTargetHttpMethodPtrOutput() AppEngineHttpTargetHttpMethodPtrOutput {
	return e.ToAppEngineHttpTargetHttpMethodPtrOutputWithContext(context.Background())
}

func (e AppEngineHttpTargetHttpMethod) ToAppEngineHttpTargetHttpMethodPtrOutputWithContext(ctx context.Context) AppEngineHttpTargetHttpMethodPtrOutput {
	return AppEngineHttpTargetHttpMethod(e).ToAppEngineHttpTargetHttpMethodOutputWithContext(ctx).ToAppEngineHttpTargetHttpMethodPtrOutputWithContext(ctx)
}

func (e AppEngineHttpTargetHttpMethod) ToStringOutput() pulumi.StringOutput {
	return pulumi.ToOutput(pulumi.String(e)).(pulumi.StringOutput)
}

func (e AppEngineHttpTargetHttpMethod) ToStringOutputWithContext(ctx context.Context) pulumi.StringOutput {
	return pulumi.ToOutputWithContext(ctx, pulumi.String(e)).(pulumi.StringOutput)
}

func (e AppEngineHttpTargetHttpMethod) ToStringPtrOutput() pulumi.StringPtrOutput {
	return pulumi.String(e).ToStringPtrOutputWithContext(context.Background())
}

func (e AppEngineHttpTargetHttpMethod) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return pulumi.String(e).ToStringOutputWithContext(ctx).ToStringPtrOutputWithContext(ctx)
}

type AppEngineHttpTargetHttpMethodOutput struct{ *pulumi.OutputState }

func (AppEngineHttpTargetHttpMethodOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*AppEngineHttpTargetHttpMethod)(nil)).Elem()
}

func (o AppEngineHttpTargetHttpMethodOutput) ToAppEngineHttpTargetHttpMethodOutput() AppEngineHttpTargetHttpMethodOutput {
	return o
}

func (o AppEngineHttpTargetHttpMethodOutput) ToAppEngineHttpTargetHttpMethodOutputWithContext(ctx context.Context) AppEngineHttpTargetHttpMethodOutput {
	return o
}

func (o AppEngineHttpTargetHttpMethodOutput) ToAppEngineHttpTargetHttpMethodPtrOutput() AppEngineHttpTargetHttpMethodPtrOutput {
	return o.ToAppEngineHttpTargetHttpMethodPtrOutputWithContext(context.Background())
}

func (o AppEngineHttpTargetHttpMethodOutput) ToAppEngineHttpTargetHttpMethodPtrOutputWithContext(ctx context.Context) AppEngineHttpTargetHttpMethodPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v AppEngineHttpTargetHttpMethod) *AppEngineHttpTargetHttpMethod {
		return &v
	}).(AppEngineHttpTargetHttpMethodPtrOutput)
}

func (o AppEngineHttpTargetHttpMethodOutput) ToStringOutput() pulumi.StringOutput {
	return o.ToStringOutputWithContext(context.Background())
}

func (o AppEngineHttpTargetHttpMethodOutput) ToStringOutputWithContext(ctx context.Context) pulumi.StringOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e AppEngineHttpTargetHttpMethod) string {
		return string(e)
	}).(pulumi.StringOutput)
}

func (o AppEngineHttpTargetHttpMethodOutput) ToStringPtrOutput() pulumi.StringPtrOutput {
	return o.ToStringPtrOutputWithContext(context.Background())
}

func (o AppEngineHttpTargetHttpMethodOutput) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e AppEngineHttpTargetHttpMethod) *string {
		v := string(e)
		return &v
	}).(pulumi.StringPtrOutput)
}

type AppEngineHttpTargetHttpMethodPtrOutput struct{ *pulumi.OutputState }

func (AppEngineHttpTargetHttpMethodPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**AppEngineHttpTargetHttpMethod)(nil)).Elem()
}

func (o AppEngineHttpTargetHttpMethodPtrOutput) ToAppEngineHttpTargetHttpMethodPtrOutput() AppEngineHttpTargetHttpMethodPtrOutput {
	return o
}

func (o AppEngineHttpTargetHttpMethodPtrOutput) ToAppEngineHttpTargetHttpMethodPtrOutputWithContext(ctx context.Context) AppEngineHttpTargetHttpMethodPtrOutput {
	return o
}

func (o AppEngineHttpTargetHttpMethodPtrOutput) Elem() AppEngineHttpTargetHttpMethodOutput {
	return o.ApplyT(func(v *AppEngineHttpTargetHttpMethod) AppEngineHttpTargetHttpMethod {
		if v != nil {
			return *v
		}
		var ret AppEngineHttpTargetHttpMethod
		return ret
	}).(AppEngineHttpTargetHttpMethodOutput)
}

func (o AppEngineHttpTargetHttpMethodPtrOutput) ToStringPtrOutput() pulumi.StringPtrOutput {
	return o.ToStringPtrOutputWithContext(context.Background())
}

func (o AppEngineHttpTargetHttpMethodPtrOutput) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e *AppEngineHttpTargetHttpMethod) *string {
		if e == nil {
			return nil
		}
		v := string(*e)
		return &v
	}).(pulumi.StringPtrOutput)
}

// AppEngineHttpTargetHttpMethodInput is an input type that accepts AppEngineHttpTargetHttpMethodArgs and AppEngineHttpTargetHttpMethodOutput values.
// You can construct a concrete instance of `AppEngineHttpTargetHttpMethodInput` via:
//
//          AppEngineHttpTargetHttpMethodArgs{...}
type AppEngineHttpTargetHttpMethodInput interface {
	pulumi.Input

	ToAppEngineHttpTargetHttpMethodOutput() AppEngineHttpTargetHttpMethodOutput
	ToAppEngineHttpTargetHttpMethodOutputWithContext(context.Context) AppEngineHttpTargetHttpMethodOutput
}

var appEngineHttpTargetHttpMethodPtrType = reflect.TypeOf((**AppEngineHttpTargetHttpMethod)(nil)).Elem()

type AppEngineHttpTargetHttpMethodPtrInput interface {
	pulumi.Input

	ToAppEngineHttpTargetHttpMethodPtrOutput() AppEngineHttpTargetHttpMethodPtrOutput
	ToAppEngineHttpTargetHttpMethodPtrOutputWithContext(context.Context) AppEngineHttpTargetHttpMethodPtrOutput
}

type appEngineHttpTargetHttpMethodPtr string

func AppEngineHttpTargetHttpMethodPtr(v string) AppEngineHttpTargetHttpMethodPtrInput {
	return (*appEngineHttpTargetHttpMethodPtr)(&v)
}

func (*appEngineHttpTargetHttpMethodPtr) ElementType() reflect.Type {
	return appEngineHttpTargetHttpMethodPtrType
}

func (in *appEngineHttpTargetHttpMethodPtr) ToAppEngineHttpTargetHttpMethodPtrOutput() AppEngineHttpTargetHttpMethodPtrOutput {
	return pulumi.ToOutput(in).(AppEngineHttpTargetHttpMethodPtrOutput)
}

func (in *appEngineHttpTargetHttpMethodPtr) ToAppEngineHttpTargetHttpMethodPtrOutputWithContext(ctx context.Context) AppEngineHttpTargetHttpMethodPtrOutput {
	return pulumi.ToOutputWithContext(ctx, in).(AppEngineHttpTargetHttpMethodPtrOutput)
}

// Which HTTP method to use for the request.
type HttpTargetHttpMethod string

const (
	// HTTP method unspecified. Defaults to POST.
	HttpTargetHttpMethodHttpMethodUnspecified = HttpTargetHttpMethod("HTTP_METHOD_UNSPECIFIED")
	// HTTP POST
	HttpTargetHttpMethodPost = HttpTargetHttpMethod("POST")
	// HTTP GET
	HttpTargetHttpMethodGet = HttpTargetHttpMethod("GET")
	// HTTP HEAD
	HttpTargetHttpMethodHead = HttpTargetHttpMethod("HEAD")
	// HTTP PUT
	HttpTargetHttpMethodPut = HttpTargetHttpMethod("PUT")
	// HTTP DELETE
	HttpTargetHttpMethodDelete = HttpTargetHttpMethod("DELETE")
	// HTTP PATCH
	HttpTargetHttpMethodPatch = HttpTargetHttpMethod("PATCH")
	// HTTP OPTIONS
	HttpTargetHttpMethodOptions = HttpTargetHttpMethod("OPTIONS")
)

func (HttpTargetHttpMethod) ElementType() reflect.Type {
	return reflect.TypeOf((*HttpTargetHttpMethod)(nil)).Elem()
}

func (e HttpTargetHttpMethod) ToHttpTargetHttpMethodOutput() HttpTargetHttpMethodOutput {
	return pulumi.ToOutput(e).(HttpTargetHttpMethodOutput)
}

func (e HttpTargetHttpMethod) ToHttpTargetHttpMethodOutputWithContext(ctx context.Context) HttpTargetHttpMethodOutput {
	return pulumi.ToOutputWithContext(ctx, e).(HttpTargetHttpMethodOutput)
}

func (e HttpTargetHttpMethod) ToHttpTargetHttpMethodPtrOutput() HttpTargetHttpMethodPtrOutput {
	return e.ToHttpTargetHttpMethodPtrOutputWithContext(context.Background())
}

func (e HttpTargetHttpMethod) ToHttpTargetHttpMethodPtrOutputWithContext(ctx context.Context) HttpTargetHttpMethodPtrOutput {
	return HttpTargetHttpMethod(e).ToHttpTargetHttpMethodOutputWithContext(ctx).ToHttpTargetHttpMethodPtrOutputWithContext(ctx)
}

func (e HttpTargetHttpMethod) ToStringOutput() pulumi.StringOutput {
	return pulumi.ToOutput(pulumi.String(e)).(pulumi.StringOutput)
}

func (e HttpTargetHttpMethod) ToStringOutputWithContext(ctx context.Context) pulumi.StringOutput {
	return pulumi.ToOutputWithContext(ctx, pulumi.String(e)).(pulumi.StringOutput)
}

func (e HttpTargetHttpMethod) ToStringPtrOutput() pulumi.StringPtrOutput {
	return pulumi.String(e).ToStringPtrOutputWithContext(context.Background())
}

func (e HttpTargetHttpMethod) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return pulumi.String(e).ToStringOutputWithContext(ctx).ToStringPtrOutputWithContext(ctx)
}

type HttpTargetHttpMethodOutput struct{ *pulumi.OutputState }

func (HttpTargetHttpMethodOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*HttpTargetHttpMethod)(nil)).Elem()
}

func (o HttpTargetHttpMethodOutput) ToHttpTargetHttpMethodOutput() HttpTargetHttpMethodOutput {
	return o
}

func (o HttpTargetHttpMethodOutput) ToHttpTargetHttpMethodOutputWithContext(ctx context.Context) HttpTargetHttpMethodOutput {
	return o
}

func (o HttpTargetHttpMethodOutput) ToHttpTargetHttpMethodPtrOutput() HttpTargetHttpMethodPtrOutput {
	return o.ToHttpTargetHttpMethodPtrOutputWithContext(context.Background())
}

func (o HttpTargetHttpMethodOutput) ToHttpTargetHttpMethodPtrOutputWithContext(ctx context.Context) HttpTargetHttpMethodPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v HttpTargetHttpMethod) *HttpTargetHttpMethod {
		return &v
	}).(HttpTargetHttpMethodPtrOutput)
}

func (o HttpTargetHttpMethodOutput) ToStringOutput() pulumi.StringOutput {
	return o.ToStringOutputWithContext(context.Background())
}

func (o HttpTargetHttpMethodOutput) ToStringOutputWithContext(ctx context.Context) pulumi.StringOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e HttpTargetHttpMethod) string {
		return string(e)
	}).(pulumi.StringOutput)
}

func (o HttpTargetHttpMethodOutput) ToStringPtrOutput() pulumi.StringPtrOutput {
	return o.ToStringPtrOutputWithContext(context.Background())
}

func (o HttpTargetHttpMethodOutput) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e HttpTargetHttpMethod) *string {
		v := string(e)
		return &v
	}).(pulumi.StringPtrOutput)
}

type HttpTargetHttpMethodPtrOutput struct{ *pulumi.OutputState }

func (HttpTargetHttpMethodPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**HttpTargetHttpMethod)(nil)).Elem()
}

func (o HttpTargetHttpMethodPtrOutput) ToHttpTargetHttpMethodPtrOutput() HttpTargetHttpMethodPtrOutput {
	return o
}

func (o HttpTargetHttpMethodPtrOutput) ToHttpTargetHttpMethodPtrOutputWithContext(ctx context.Context) HttpTargetHttpMethodPtrOutput {
	return o
}

func (o HttpTargetHttpMethodPtrOutput) Elem() HttpTargetHttpMethodOutput {
	return o.ApplyT(func(v *HttpTargetHttpMethod) HttpTargetHttpMethod {
		if v != nil {
			return *v
		}
		var ret HttpTargetHttpMethod
		return ret
	}).(HttpTargetHttpMethodOutput)
}

func (o HttpTargetHttpMethodPtrOutput) ToStringPtrOutput() pulumi.StringPtrOutput {
	return o.ToStringPtrOutputWithContext(context.Background())
}

func (o HttpTargetHttpMethodPtrOutput) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e *HttpTargetHttpMethod) *string {
		if e == nil {
			return nil
		}
		v := string(*e)
		return &v
	}).(pulumi.StringPtrOutput)
}

// HttpTargetHttpMethodInput is an input type that accepts HttpTargetHttpMethodArgs and HttpTargetHttpMethodOutput values.
// You can construct a concrete instance of `HttpTargetHttpMethodInput` via:
//
//          HttpTargetHttpMethodArgs{...}
type HttpTargetHttpMethodInput interface {
	pulumi.Input

	ToHttpTargetHttpMethodOutput() HttpTargetHttpMethodOutput
	ToHttpTargetHttpMethodOutputWithContext(context.Context) HttpTargetHttpMethodOutput
}

var httpTargetHttpMethodPtrType = reflect.TypeOf((**HttpTargetHttpMethod)(nil)).Elem()

type HttpTargetHttpMethodPtrInput interface {
	pulumi.Input

	ToHttpTargetHttpMethodPtrOutput() HttpTargetHttpMethodPtrOutput
	ToHttpTargetHttpMethodPtrOutputWithContext(context.Context) HttpTargetHttpMethodPtrOutput
}

type httpTargetHttpMethodPtr string

func HttpTargetHttpMethodPtr(v string) HttpTargetHttpMethodPtrInput {
	return (*httpTargetHttpMethodPtr)(&v)
}

func (*httpTargetHttpMethodPtr) ElementType() reflect.Type {
	return httpTargetHttpMethodPtrType
}

func (in *httpTargetHttpMethodPtr) ToHttpTargetHttpMethodPtrOutput() HttpTargetHttpMethodPtrOutput {
	return pulumi.ToOutput(in).(HttpTargetHttpMethodPtrOutput)
}

func (in *httpTargetHttpMethodPtr) ToHttpTargetHttpMethodPtrOutputWithContext(ctx context.Context) HttpTargetHttpMethodPtrOutput {
	return pulumi.ToOutputWithContext(ctx, in).(HttpTargetHttpMethodPtrOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*AppEngineHttpTargetHttpMethodInput)(nil)).Elem(), AppEngineHttpTargetHttpMethod("HTTP_METHOD_UNSPECIFIED"))
	pulumi.RegisterInputType(reflect.TypeOf((*AppEngineHttpTargetHttpMethodPtrInput)(nil)).Elem(), AppEngineHttpTargetHttpMethod("HTTP_METHOD_UNSPECIFIED"))
	pulumi.RegisterInputType(reflect.TypeOf((*HttpTargetHttpMethodInput)(nil)).Elem(), HttpTargetHttpMethod("HTTP_METHOD_UNSPECIFIED"))
	pulumi.RegisterInputType(reflect.TypeOf((*HttpTargetHttpMethodPtrInput)(nil)).Elem(), HttpTargetHttpMethod("HTTP_METHOD_UNSPECIFIED"))
	pulumi.RegisterOutputType(AppEngineHttpTargetHttpMethodOutput{})
	pulumi.RegisterOutputType(AppEngineHttpTargetHttpMethodPtrOutput{})
	pulumi.RegisterOutputType(HttpTargetHttpMethodOutput{})
	pulumi.RegisterOutputType(HttpTargetHttpMethodPtrOutput{})
}
