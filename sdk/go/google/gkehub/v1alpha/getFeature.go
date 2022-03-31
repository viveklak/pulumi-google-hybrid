// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package v1alpha

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Gets details of a single Feature.
func LookupFeature(ctx *pulumi.Context, args *LookupFeatureArgs, opts ...pulumi.InvokeOption) (*LookupFeatureResult, error) {
	var rv LookupFeatureResult
	err := ctx.Invoke("google-native:gkehub/v1alpha:getFeature", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupFeatureArgs struct {
	FeatureId string  `pulumi:"featureId"`
	Location  string  `pulumi:"location"`
	Project   *string `pulumi:"project"`
}

type LookupFeatureResult struct {
	// When the Feature resource was created.
	CreateTime string `pulumi:"createTime"`
	// When the Feature resource was deleted.
	DeleteTime string `pulumi:"deleteTime"`
	// GCP labels for this Feature.
	Labels map[string]string `pulumi:"labels"`
	// Optional. Membership-specific configuration for this Feature. If this Feature does not support any per-Membership configuration, this field may be unused. The keys indicate which Membership the configuration is for, in the form: `projects/{p}/locations/{l}/memberships/{m}` Where {p} is the project, {l} is a valid location and {m} is a valid Membership in this project at that location. {p} WILL match the Feature's project. {p} will always be returned as the project number, but the project ID is also accepted during input. If the same Membership is specified in the map twice (using the project ID form, and the project number form), exactly ONE of the entries will be saved, with no guarantees as to which. For this reason, it is recommended the same format be used for all entries when mutating a Feature.
	MembershipSpecs map[string]string `pulumi:"membershipSpecs"`
	// Membership-specific Feature status. If this Feature does report any per-Membership status, this field may be unused. The keys indicate which Membership the state is for, in the form: `projects/{p}/locations/{l}/memberships/{m}` Where {p} is the project number, {l} is a valid location and {m} is a valid Membership in this project at that location. {p} MUST match the Feature's project number.
	MembershipStates map[string]string `pulumi:"membershipStates"`
	// The full, unique name of this Feature resource in the format `projects/*/locations/*/features/*`.
	Name string `pulumi:"name"`
	// State of the Feature resource itself.
	ResourceState FeatureResourceStateResponse `pulumi:"resourceState"`
	// Optional. Hub-wide Feature configuration. If this Feature does not support any Hub-wide configuration, this field may be unused.
	Spec CommonFeatureSpecResponse `pulumi:"spec"`
	// The Hub-wide Feature state.
	State CommonFeatureStateResponse `pulumi:"state"`
	// When the Feature resource was last updated.
	UpdateTime string `pulumi:"updateTime"`
}

func LookupFeatureOutput(ctx *pulumi.Context, args LookupFeatureOutputArgs, opts ...pulumi.InvokeOption) LookupFeatureResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupFeatureResult, error) {
			args := v.(LookupFeatureArgs)
			r, err := LookupFeature(ctx, &args, opts...)
			return *r, err
		}).(LookupFeatureResultOutput)
}

type LookupFeatureOutputArgs struct {
	FeatureId pulumi.StringInput    `pulumi:"featureId"`
	Location  pulumi.StringInput    `pulumi:"location"`
	Project   pulumi.StringPtrInput `pulumi:"project"`
}

func (LookupFeatureOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupFeatureArgs)(nil)).Elem()
}

type LookupFeatureResultOutput struct{ *pulumi.OutputState }

func (LookupFeatureResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupFeatureResult)(nil)).Elem()
}

func (o LookupFeatureResultOutput) ToLookupFeatureResultOutput() LookupFeatureResultOutput {
	return o
}

func (o LookupFeatureResultOutput) ToLookupFeatureResultOutputWithContext(ctx context.Context) LookupFeatureResultOutput {
	return o
}

// When the Feature resource was created.
func (o LookupFeatureResultOutput) CreateTime() pulumi.StringOutput {
	return o.ApplyT(func(v LookupFeatureResult) string { return v.CreateTime }).(pulumi.StringOutput)
}

// When the Feature resource was deleted.
func (o LookupFeatureResultOutput) DeleteTime() pulumi.StringOutput {
	return o.ApplyT(func(v LookupFeatureResult) string { return v.DeleteTime }).(pulumi.StringOutput)
}

// GCP labels for this Feature.
func (o LookupFeatureResultOutput) Labels() pulumi.StringMapOutput {
	return o.ApplyT(func(v LookupFeatureResult) map[string]string { return v.Labels }).(pulumi.StringMapOutput)
}

// Optional. Membership-specific configuration for this Feature. If this Feature does not support any per-Membership configuration, this field may be unused. The keys indicate which Membership the configuration is for, in the form: `projects/{p}/locations/{l}/memberships/{m}` Where {p} is the project, {l} is a valid location and {m} is a valid Membership in this project at that location. {p} WILL match the Feature's project. {p} will always be returned as the project number, but the project ID is also accepted during input. If the same Membership is specified in the map twice (using the project ID form, and the project number form), exactly ONE of the entries will be saved, with no guarantees as to which. For this reason, it is recommended the same format be used for all entries when mutating a Feature.
func (o LookupFeatureResultOutput) MembershipSpecs() pulumi.StringMapOutput {
	return o.ApplyT(func(v LookupFeatureResult) map[string]string { return v.MembershipSpecs }).(pulumi.StringMapOutput)
}

// Membership-specific Feature status. If this Feature does report any per-Membership status, this field may be unused. The keys indicate which Membership the state is for, in the form: `projects/{p}/locations/{l}/memberships/{m}` Where {p} is the project number, {l} is a valid location and {m} is a valid Membership in this project at that location. {p} MUST match the Feature's project number.
func (o LookupFeatureResultOutput) MembershipStates() pulumi.StringMapOutput {
	return o.ApplyT(func(v LookupFeatureResult) map[string]string { return v.MembershipStates }).(pulumi.StringMapOutput)
}

// The full, unique name of this Feature resource in the format `projects/*/locations/*/features/*`.
func (o LookupFeatureResultOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v LookupFeatureResult) string { return v.Name }).(pulumi.StringOutput)
}

// State of the Feature resource itself.
func (o LookupFeatureResultOutput) ResourceState() FeatureResourceStateResponseOutput {
	return o.ApplyT(func(v LookupFeatureResult) FeatureResourceStateResponse { return v.ResourceState }).(FeatureResourceStateResponseOutput)
}

// Optional. Hub-wide Feature configuration. If this Feature does not support any Hub-wide configuration, this field may be unused.
func (o LookupFeatureResultOutput) Spec() CommonFeatureSpecResponseOutput {
	return o.ApplyT(func(v LookupFeatureResult) CommonFeatureSpecResponse { return v.Spec }).(CommonFeatureSpecResponseOutput)
}

// The Hub-wide Feature state.
func (o LookupFeatureResultOutput) State() CommonFeatureStateResponseOutput {
	return o.ApplyT(func(v LookupFeatureResult) CommonFeatureStateResponse { return v.State }).(CommonFeatureStateResponseOutput)
}

// When the Feature resource was last updated.
func (o LookupFeatureResultOutput) UpdateTime() pulumi.StringOutput {
	return o.ApplyT(func(v LookupFeatureResult) string { return v.UpdateTime }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupFeatureResultOutput{})
}
