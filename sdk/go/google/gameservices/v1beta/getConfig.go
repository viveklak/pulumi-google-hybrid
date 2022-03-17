// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package v1beta

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Gets details of a single game server config.
func LookupConfig(ctx *pulumi.Context, args *LookupConfigArgs, opts ...pulumi.InvokeOption) (*LookupConfigResult, error) {
	var rv LookupConfigResult
	err := ctx.Invoke("google-native:gameservices/v1beta:getConfig", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupConfigArgs struct {
	ConfigId               string  `pulumi:"configId"`
	GameServerDeploymentId string  `pulumi:"gameServerDeploymentId"`
	Location               string  `pulumi:"location"`
	Project                *string `pulumi:"project"`
}

type LookupConfigResult struct {
	// The creation time.
	CreateTime string `pulumi:"createTime"`
	// The description of the game server config.
	Description string `pulumi:"description"`
	// FleetConfig contains a list of Agones fleet specs. Only one FleetConfig is allowed.
	FleetConfigs []FleetConfigResponse `pulumi:"fleetConfigs"`
	// The labels associated with this game server config. Each label is a key-value pair.
	Labels map[string]string `pulumi:"labels"`
	// The resource name of the game server config, in the following form: `projects/{project}/locations/{locationId}/gameServerDeployments/{deploymentId}/configs/{configId}`. For example, `projects/my-project/locations/global/gameServerDeployments/my-game/configs/my-config`.
	Name string `pulumi:"name"`
	// The autoscaling settings.
	ScalingConfigs []ScalingConfigResponse `pulumi:"scalingConfigs"`
	// The last-modified time.
	UpdateTime string `pulumi:"updateTime"`
}

func LookupConfigOutput(ctx *pulumi.Context, args LookupConfigOutputArgs, opts ...pulumi.InvokeOption) LookupConfigResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupConfigResult, error) {
			args := v.(LookupConfigArgs)
			r, err := LookupConfig(ctx, &args, opts...)
			return *r, err
		}).(LookupConfigResultOutput)
}

type LookupConfigOutputArgs struct {
	ConfigId               pulumi.StringInput    `pulumi:"configId"`
	GameServerDeploymentId pulumi.StringInput    `pulumi:"gameServerDeploymentId"`
	Location               pulumi.StringInput    `pulumi:"location"`
	Project                pulumi.StringPtrInput `pulumi:"project"`
}

func (LookupConfigOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupConfigArgs)(nil)).Elem()
}

type LookupConfigResultOutput struct{ *pulumi.OutputState }

func (LookupConfigResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupConfigResult)(nil)).Elem()
}

func (o LookupConfigResultOutput) ToLookupConfigResultOutput() LookupConfigResultOutput {
	return o
}

func (o LookupConfigResultOutput) ToLookupConfigResultOutputWithContext(ctx context.Context) LookupConfigResultOutput {
	return o
}

// The creation time.
func (o LookupConfigResultOutput) CreateTime() pulumi.StringOutput {
	return o.ApplyT(func(v LookupConfigResult) string { return v.CreateTime }).(pulumi.StringOutput)
}

// The description of the game server config.
func (o LookupConfigResultOutput) Description() pulumi.StringOutput {
	return o.ApplyT(func(v LookupConfigResult) string { return v.Description }).(pulumi.StringOutput)
}

// FleetConfig contains a list of Agones fleet specs. Only one FleetConfig is allowed.
func (o LookupConfigResultOutput) FleetConfigs() FleetConfigResponseArrayOutput {
	return o.ApplyT(func(v LookupConfigResult) []FleetConfigResponse { return v.FleetConfigs }).(FleetConfigResponseArrayOutput)
}

// The labels associated with this game server config. Each label is a key-value pair.
func (o LookupConfigResultOutput) Labels() pulumi.StringMapOutput {
	return o.ApplyT(func(v LookupConfigResult) map[string]string { return v.Labels }).(pulumi.StringMapOutput)
}

// The resource name of the game server config, in the following form: `projects/{project}/locations/{locationId}/gameServerDeployments/{deploymentId}/configs/{configId}`. For example, `projects/my-project/locations/global/gameServerDeployments/my-game/configs/my-config`.
func (o LookupConfigResultOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v LookupConfigResult) string { return v.Name }).(pulumi.StringOutput)
}

// The autoscaling settings.
func (o LookupConfigResultOutput) ScalingConfigs() ScalingConfigResponseArrayOutput {
	return o.ApplyT(func(v LookupConfigResult) []ScalingConfigResponse { return v.ScalingConfigs }).(ScalingConfigResponseArrayOutput)
}

// The last-modified time.
func (o LookupConfigResultOutput) UpdateTime() pulumi.StringOutput {
	return o.ApplyT(func(v LookupConfigResult) string { return v.UpdateTime }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupConfigResultOutput{})
}
