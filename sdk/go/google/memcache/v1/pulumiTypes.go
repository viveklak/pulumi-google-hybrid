// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package v1

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

type InstanceMessage struct {
	// A code that correspond to one type of user-facing message.
	Code *InstanceMessageCode `pulumi:"code"`
	// Message on memcached instance which will be exposed to users.
	Message *string `pulumi:"message"`
}

// InstanceMessageInput is an input type that accepts InstanceMessageArgs and InstanceMessageOutput values.
// You can construct a concrete instance of `InstanceMessageInput` via:
//
//          InstanceMessageArgs{...}
type InstanceMessageInput interface {
	pulumi.Input

	ToInstanceMessageOutput() InstanceMessageOutput
	ToInstanceMessageOutputWithContext(context.Context) InstanceMessageOutput
}

type InstanceMessageArgs struct {
	// A code that correspond to one type of user-facing message.
	Code InstanceMessageCodePtrInput `pulumi:"code"`
	// Message on memcached instance which will be exposed to users.
	Message pulumi.StringPtrInput `pulumi:"message"`
}

func (InstanceMessageArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*InstanceMessage)(nil)).Elem()
}

func (i InstanceMessageArgs) ToInstanceMessageOutput() InstanceMessageOutput {
	return i.ToInstanceMessageOutputWithContext(context.Background())
}

func (i InstanceMessageArgs) ToInstanceMessageOutputWithContext(ctx context.Context) InstanceMessageOutput {
	return pulumi.ToOutputWithContext(ctx, i).(InstanceMessageOutput)
}

// InstanceMessageArrayInput is an input type that accepts InstanceMessageArray and InstanceMessageArrayOutput values.
// You can construct a concrete instance of `InstanceMessageArrayInput` via:
//
//          InstanceMessageArray{ InstanceMessageArgs{...} }
type InstanceMessageArrayInput interface {
	pulumi.Input

	ToInstanceMessageArrayOutput() InstanceMessageArrayOutput
	ToInstanceMessageArrayOutputWithContext(context.Context) InstanceMessageArrayOutput
}

type InstanceMessageArray []InstanceMessageInput

func (InstanceMessageArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]InstanceMessage)(nil)).Elem()
}

func (i InstanceMessageArray) ToInstanceMessageArrayOutput() InstanceMessageArrayOutput {
	return i.ToInstanceMessageArrayOutputWithContext(context.Background())
}

func (i InstanceMessageArray) ToInstanceMessageArrayOutputWithContext(ctx context.Context) InstanceMessageArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(InstanceMessageArrayOutput)
}

type InstanceMessageOutput struct{ *pulumi.OutputState }

func (InstanceMessageOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*InstanceMessage)(nil)).Elem()
}

func (o InstanceMessageOutput) ToInstanceMessageOutput() InstanceMessageOutput {
	return o
}

func (o InstanceMessageOutput) ToInstanceMessageOutputWithContext(ctx context.Context) InstanceMessageOutput {
	return o
}

// A code that correspond to one type of user-facing message.
func (o InstanceMessageOutput) Code() InstanceMessageCodePtrOutput {
	return o.ApplyT(func(v InstanceMessage) *InstanceMessageCode { return v.Code }).(InstanceMessageCodePtrOutput)
}

// Message on memcached instance which will be exposed to users.
func (o InstanceMessageOutput) Message() pulumi.StringPtrOutput {
	return o.ApplyT(func(v InstanceMessage) *string { return v.Message }).(pulumi.StringPtrOutput)
}

type InstanceMessageArrayOutput struct{ *pulumi.OutputState }

func (InstanceMessageArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]InstanceMessage)(nil)).Elem()
}

func (o InstanceMessageArrayOutput) ToInstanceMessageArrayOutput() InstanceMessageArrayOutput {
	return o
}

func (o InstanceMessageArrayOutput) ToInstanceMessageArrayOutputWithContext(ctx context.Context) InstanceMessageArrayOutput {
	return o
}

func (o InstanceMessageArrayOutput) Index(i pulumi.IntInput) InstanceMessageOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) InstanceMessage {
		return vs[0].([]InstanceMessage)[vs[1].(int)]
	}).(InstanceMessageOutput)
}

type InstanceMessageResponse struct {
	// A code that correspond to one type of user-facing message.
	Code string `pulumi:"code"`
	// Message on memcached instance which will be exposed to users.
	Message string `pulumi:"message"`
}

type InstanceMessageResponseOutput struct{ *pulumi.OutputState }

func (InstanceMessageResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*InstanceMessageResponse)(nil)).Elem()
}

func (o InstanceMessageResponseOutput) ToInstanceMessageResponseOutput() InstanceMessageResponseOutput {
	return o
}

func (o InstanceMessageResponseOutput) ToInstanceMessageResponseOutputWithContext(ctx context.Context) InstanceMessageResponseOutput {
	return o
}

// A code that correspond to one type of user-facing message.
func (o InstanceMessageResponseOutput) Code() pulumi.StringOutput {
	return o.ApplyT(func(v InstanceMessageResponse) string { return v.Code }).(pulumi.StringOutput)
}

// Message on memcached instance which will be exposed to users.
func (o InstanceMessageResponseOutput) Message() pulumi.StringOutput {
	return o.ApplyT(func(v InstanceMessageResponse) string { return v.Message }).(pulumi.StringOutput)
}

type InstanceMessageResponseArrayOutput struct{ *pulumi.OutputState }

func (InstanceMessageResponseArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]InstanceMessageResponse)(nil)).Elem()
}

func (o InstanceMessageResponseArrayOutput) ToInstanceMessageResponseArrayOutput() InstanceMessageResponseArrayOutput {
	return o
}

func (o InstanceMessageResponseArrayOutput) ToInstanceMessageResponseArrayOutputWithContext(ctx context.Context) InstanceMessageResponseArrayOutput {
	return o
}

func (o InstanceMessageResponseArrayOutput) Index(i pulumi.IntInput) InstanceMessageResponseOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) InstanceMessageResponse {
		return vs[0].([]InstanceMessageResponse)[vs[1].(int)]
	}).(InstanceMessageResponseOutput)
}

type MemcacheParameters struct {
	// User defined set of parameters to use in the memcached process.
	Params map[string]string `pulumi:"params"`
}

// MemcacheParametersInput is an input type that accepts MemcacheParametersArgs and MemcacheParametersOutput values.
// You can construct a concrete instance of `MemcacheParametersInput` via:
//
//          MemcacheParametersArgs{...}
type MemcacheParametersInput interface {
	pulumi.Input

	ToMemcacheParametersOutput() MemcacheParametersOutput
	ToMemcacheParametersOutputWithContext(context.Context) MemcacheParametersOutput
}

type MemcacheParametersArgs struct {
	// User defined set of parameters to use in the memcached process.
	Params pulumi.StringMapInput `pulumi:"params"`
}

func (MemcacheParametersArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*MemcacheParameters)(nil)).Elem()
}

func (i MemcacheParametersArgs) ToMemcacheParametersOutput() MemcacheParametersOutput {
	return i.ToMemcacheParametersOutputWithContext(context.Background())
}

func (i MemcacheParametersArgs) ToMemcacheParametersOutputWithContext(ctx context.Context) MemcacheParametersOutput {
	return pulumi.ToOutputWithContext(ctx, i).(MemcacheParametersOutput)
}

func (i MemcacheParametersArgs) ToMemcacheParametersPtrOutput() MemcacheParametersPtrOutput {
	return i.ToMemcacheParametersPtrOutputWithContext(context.Background())
}

func (i MemcacheParametersArgs) ToMemcacheParametersPtrOutputWithContext(ctx context.Context) MemcacheParametersPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(MemcacheParametersOutput).ToMemcacheParametersPtrOutputWithContext(ctx)
}

// MemcacheParametersPtrInput is an input type that accepts MemcacheParametersArgs, MemcacheParametersPtr and MemcacheParametersPtrOutput values.
// You can construct a concrete instance of `MemcacheParametersPtrInput` via:
//
//          MemcacheParametersArgs{...}
//
//  or:
//
//          nil
type MemcacheParametersPtrInput interface {
	pulumi.Input

	ToMemcacheParametersPtrOutput() MemcacheParametersPtrOutput
	ToMemcacheParametersPtrOutputWithContext(context.Context) MemcacheParametersPtrOutput
}

type memcacheParametersPtrType MemcacheParametersArgs

func MemcacheParametersPtr(v *MemcacheParametersArgs) MemcacheParametersPtrInput {
	return (*memcacheParametersPtrType)(v)
}

func (*memcacheParametersPtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**MemcacheParameters)(nil)).Elem()
}

func (i *memcacheParametersPtrType) ToMemcacheParametersPtrOutput() MemcacheParametersPtrOutput {
	return i.ToMemcacheParametersPtrOutputWithContext(context.Background())
}

func (i *memcacheParametersPtrType) ToMemcacheParametersPtrOutputWithContext(ctx context.Context) MemcacheParametersPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(MemcacheParametersPtrOutput)
}

type MemcacheParametersOutput struct{ *pulumi.OutputState }

func (MemcacheParametersOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*MemcacheParameters)(nil)).Elem()
}

func (o MemcacheParametersOutput) ToMemcacheParametersOutput() MemcacheParametersOutput {
	return o
}

func (o MemcacheParametersOutput) ToMemcacheParametersOutputWithContext(ctx context.Context) MemcacheParametersOutput {
	return o
}

func (o MemcacheParametersOutput) ToMemcacheParametersPtrOutput() MemcacheParametersPtrOutput {
	return o.ToMemcacheParametersPtrOutputWithContext(context.Background())
}

func (o MemcacheParametersOutput) ToMemcacheParametersPtrOutputWithContext(ctx context.Context) MemcacheParametersPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v MemcacheParameters) *MemcacheParameters {
		return &v
	}).(MemcacheParametersPtrOutput)
}

// User defined set of parameters to use in the memcached process.
func (o MemcacheParametersOutput) Params() pulumi.StringMapOutput {
	return o.ApplyT(func(v MemcacheParameters) map[string]string { return v.Params }).(pulumi.StringMapOutput)
}

type MemcacheParametersPtrOutput struct{ *pulumi.OutputState }

func (MemcacheParametersPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**MemcacheParameters)(nil)).Elem()
}

func (o MemcacheParametersPtrOutput) ToMemcacheParametersPtrOutput() MemcacheParametersPtrOutput {
	return o
}

func (o MemcacheParametersPtrOutput) ToMemcacheParametersPtrOutputWithContext(ctx context.Context) MemcacheParametersPtrOutput {
	return o
}

func (o MemcacheParametersPtrOutput) Elem() MemcacheParametersOutput {
	return o.ApplyT(func(v *MemcacheParameters) MemcacheParameters {
		if v != nil {
			return *v
		}
		var ret MemcacheParameters
		return ret
	}).(MemcacheParametersOutput)
}

// User defined set of parameters to use in the memcached process.
func (o MemcacheParametersPtrOutput) Params() pulumi.StringMapOutput {
	return o.ApplyT(func(v *MemcacheParameters) map[string]string {
		if v == nil {
			return nil
		}
		return v.Params
	}).(pulumi.StringMapOutput)
}

type MemcacheParametersResponse struct {
	// User defined set of parameters to use in the memcached process.
	Params map[string]string `pulumi:"params"`
}

type MemcacheParametersResponseOutput struct{ *pulumi.OutputState }

func (MemcacheParametersResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*MemcacheParametersResponse)(nil)).Elem()
}

func (o MemcacheParametersResponseOutput) ToMemcacheParametersResponseOutput() MemcacheParametersResponseOutput {
	return o
}

func (o MemcacheParametersResponseOutput) ToMemcacheParametersResponseOutputWithContext(ctx context.Context) MemcacheParametersResponseOutput {
	return o
}

// User defined set of parameters to use in the memcached process.
func (o MemcacheParametersResponseOutput) Params() pulumi.StringMapOutput {
	return o.ApplyT(func(v MemcacheParametersResponse) map[string]string { return v.Params }).(pulumi.StringMapOutput)
}

// Configuration for a Memcached Node.
type NodeConfig struct {
	// Number of cpus per Memcached node.
	CpuCount int `pulumi:"cpuCount"`
	// Memory size in MiB for each Memcached node.
	MemorySizeMb int `pulumi:"memorySizeMb"`
}

// NodeConfigInput is an input type that accepts NodeConfigArgs and NodeConfigOutput values.
// You can construct a concrete instance of `NodeConfigInput` via:
//
//          NodeConfigArgs{...}
type NodeConfigInput interface {
	pulumi.Input

	ToNodeConfigOutput() NodeConfigOutput
	ToNodeConfigOutputWithContext(context.Context) NodeConfigOutput
}

// Configuration for a Memcached Node.
type NodeConfigArgs struct {
	// Number of cpus per Memcached node.
	CpuCount pulumi.IntInput `pulumi:"cpuCount"`
	// Memory size in MiB for each Memcached node.
	MemorySizeMb pulumi.IntInput `pulumi:"memorySizeMb"`
}

func (NodeConfigArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*NodeConfig)(nil)).Elem()
}

func (i NodeConfigArgs) ToNodeConfigOutput() NodeConfigOutput {
	return i.ToNodeConfigOutputWithContext(context.Background())
}

func (i NodeConfigArgs) ToNodeConfigOutputWithContext(ctx context.Context) NodeConfigOutput {
	return pulumi.ToOutputWithContext(ctx, i).(NodeConfigOutput)
}

// Configuration for a Memcached Node.
type NodeConfigOutput struct{ *pulumi.OutputState }

func (NodeConfigOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*NodeConfig)(nil)).Elem()
}

func (o NodeConfigOutput) ToNodeConfigOutput() NodeConfigOutput {
	return o
}

func (o NodeConfigOutput) ToNodeConfigOutputWithContext(ctx context.Context) NodeConfigOutput {
	return o
}

// Number of cpus per Memcached node.
func (o NodeConfigOutput) CpuCount() pulumi.IntOutput {
	return o.ApplyT(func(v NodeConfig) int { return v.CpuCount }).(pulumi.IntOutput)
}

// Memory size in MiB for each Memcached node.
func (o NodeConfigOutput) MemorySizeMb() pulumi.IntOutput {
	return o.ApplyT(func(v NodeConfig) int { return v.MemorySizeMb }).(pulumi.IntOutput)
}

// Configuration for a Memcached Node.
type NodeConfigResponse struct {
	// Number of cpus per Memcached node.
	CpuCount int `pulumi:"cpuCount"`
	// Memory size in MiB for each Memcached node.
	MemorySizeMb int `pulumi:"memorySizeMb"`
}

// Configuration for a Memcached Node.
type NodeConfigResponseOutput struct{ *pulumi.OutputState }

func (NodeConfigResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*NodeConfigResponse)(nil)).Elem()
}

func (o NodeConfigResponseOutput) ToNodeConfigResponseOutput() NodeConfigResponseOutput {
	return o
}

func (o NodeConfigResponseOutput) ToNodeConfigResponseOutputWithContext(ctx context.Context) NodeConfigResponseOutput {
	return o
}

// Number of cpus per Memcached node.
func (o NodeConfigResponseOutput) CpuCount() pulumi.IntOutput {
	return o.ApplyT(func(v NodeConfigResponse) int { return v.CpuCount }).(pulumi.IntOutput)
}

// Memory size in MiB for each Memcached node.
func (o NodeConfigResponseOutput) MemorySizeMb() pulumi.IntOutput {
	return o.ApplyT(func(v NodeConfigResponse) int { return v.MemorySizeMb }).(pulumi.IntOutput)
}

type NodeResponse struct {
	// Hostname or IP address of the Memcached node used by the clients to connect to the Memcached server on this node.
	Host string `pulumi:"host"`
	// Identifier of the Memcached node. The node id does not include project or location like the Memcached instance name.
	NodeId string `pulumi:"nodeId"`
	// User defined parameters currently applied to the node.
	Parameters MemcacheParametersResponse `pulumi:"parameters"`
	// The port number of the Memcached server on this node.
	Port int `pulumi:"port"`
	// Current state of the Memcached node.
	State string `pulumi:"state"`
	// Location (GCP Zone) for the Memcached node.
	Zone string `pulumi:"zone"`
}

type NodeResponseOutput struct{ *pulumi.OutputState }

func (NodeResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*NodeResponse)(nil)).Elem()
}

func (o NodeResponseOutput) ToNodeResponseOutput() NodeResponseOutput {
	return o
}

func (o NodeResponseOutput) ToNodeResponseOutputWithContext(ctx context.Context) NodeResponseOutput {
	return o
}

// Hostname or IP address of the Memcached node used by the clients to connect to the Memcached server on this node.
func (o NodeResponseOutput) Host() pulumi.StringOutput {
	return o.ApplyT(func(v NodeResponse) string { return v.Host }).(pulumi.StringOutput)
}

// Identifier of the Memcached node. The node id does not include project or location like the Memcached instance name.
func (o NodeResponseOutput) NodeId() pulumi.StringOutput {
	return o.ApplyT(func(v NodeResponse) string { return v.NodeId }).(pulumi.StringOutput)
}

// User defined parameters currently applied to the node.
func (o NodeResponseOutput) Parameters() MemcacheParametersResponseOutput {
	return o.ApplyT(func(v NodeResponse) MemcacheParametersResponse { return v.Parameters }).(MemcacheParametersResponseOutput)
}

// The port number of the Memcached server on this node.
func (o NodeResponseOutput) Port() pulumi.IntOutput {
	return o.ApplyT(func(v NodeResponse) int { return v.Port }).(pulumi.IntOutput)
}

// Current state of the Memcached node.
func (o NodeResponseOutput) State() pulumi.StringOutput {
	return o.ApplyT(func(v NodeResponse) string { return v.State }).(pulumi.StringOutput)
}

// Location (GCP Zone) for the Memcached node.
func (o NodeResponseOutput) Zone() pulumi.StringOutput {
	return o.ApplyT(func(v NodeResponse) string { return v.Zone }).(pulumi.StringOutput)
}

type NodeResponseArrayOutput struct{ *pulumi.OutputState }

func (NodeResponseArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]NodeResponse)(nil)).Elem()
}

func (o NodeResponseArrayOutput) ToNodeResponseArrayOutput() NodeResponseArrayOutput {
	return o
}

func (o NodeResponseArrayOutput) ToNodeResponseArrayOutputWithContext(ctx context.Context) NodeResponseArrayOutput {
	return o
}

func (o NodeResponseArrayOutput) Index(i pulumi.IntInput) NodeResponseOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) NodeResponse {
		return vs[0].([]NodeResponse)[vs[1].(int)]
	}).(NodeResponseOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*InstanceMessageInput)(nil)).Elem(), InstanceMessageArgs{})
	pulumi.RegisterInputType(reflect.TypeOf((*InstanceMessageArrayInput)(nil)).Elem(), InstanceMessageArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*MemcacheParametersInput)(nil)).Elem(), MemcacheParametersArgs{})
	pulumi.RegisterInputType(reflect.TypeOf((*MemcacheParametersPtrInput)(nil)).Elem(), MemcacheParametersArgs{})
	pulumi.RegisterInputType(reflect.TypeOf((*NodeConfigInput)(nil)).Elem(), NodeConfigArgs{})
	pulumi.RegisterOutputType(InstanceMessageOutput{})
	pulumi.RegisterOutputType(InstanceMessageArrayOutput{})
	pulumi.RegisterOutputType(InstanceMessageResponseOutput{})
	pulumi.RegisterOutputType(InstanceMessageResponseArrayOutput{})
	pulumi.RegisterOutputType(MemcacheParametersOutput{})
	pulumi.RegisterOutputType(MemcacheParametersPtrOutput{})
	pulumi.RegisterOutputType(MemcacheParametersResponseOutput{})
	pulumi.RegisterOutputType(NodeConfigOutput{})
	pulumi.RegisterOutputType(NodeConfigResponseOutput{})
	pulumi.RegisterOutputType(NodeResponseOutput{})
	pulumi.RegisterOutputType(NodeResponseArrayOutput{})
}
