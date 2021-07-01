// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package v1

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Create a `Ruleset` from `Source`. The `Ruleset` is given a unique generated name which is returned to the caller. `Source` containing syntactic or semantics errors will result in an error response indicating the first error encountered. For a detailed view of `Source` issues, use TestRuleset.
type Ruleset struct {
	pulumi.CustomResourceState

	// Time the `Ruleset` was created. Output only.
	CreateTime pulumi.StringOutput `pulumi:"createTime"`
	// The metadata for this ruleset. Output only.
	Metadata MetadataResponseOutput `pulumi:"metadata"`
	// Name of the `Ruleset`. The ruleset_id is auto generated by the service. Format: `projects/{project_id}/rulesets/{ruleset_id}` Output only.
	Name pulumi.StringOutput `pulumi:"name"`
	// `Source` for the `Ruleset`.
	Source SourceResponseOutput `pulumi:"source"`
}

// NewRuleset registers a new resource with the given unique name, arguments, and options.
func NewRuleset(ctx *pulumi.Context,
	name string, args *RulesetArgs, opts ...pulumi.ResourceOption) (*Ruleset, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.Project == nil {
		return nil, errors.New("invalid value for required argument 'Project'")
	}
	var resource Ruleset
	err := ctx.RegisterResource("google-native:firebaserules/v1:Ruleset", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetRuleset gets an existing Ruleset resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetRuleset(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *RulesetState, opts ...pulumi.ResourceOption) (*Ruleset, error) {
	var resource Ruleset
	err := ctx.ReadResource("google-native:firebaserules/v1:Ruleset", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering Ruleset resources.
type rulesetState struct {
}

type RulesetState struct {
}

func (RulesetState) ElementType() reflect.Type {
	return reflect.TypeOf((*rulesetState)(nil)).Elem()
}

type rulesetArgs struct {
	// Time the `Ruleset` was created. Output only.
	CreateTime *string `pulumi:"createTime"`
	// The metadata for this ruleset. Output only.
	Metadata *Metadata `pulumi:"metadata"`
	// Name of the `Ruleset`. The ruleset_id is auto generated by the service. Format: `projects/{project_id}/rulesets/{ruleset_id}` Output only.
	Name    *string `pulumi:"name"`
	Project string  `pulumi:"project"`
	// `Source` for the `Ruleset`.
	Source *Source `pulumi:"source"`
}

// The set of arguments for constructing a Ruleset resource.
type RulesetArgs struct {
	// Time the `Ruleset` was created. Output only.
	CreateTime pulumi.StringPtrInput
	// The metadata for this ruleset. Output only.
	Metadata MetadataPtrInput
	// Name of the `Ruleset`. The ruleset_id is auto generated by the service. Format: `projects/{project_id}/rulesets/{ruleset_id}` Output only.
	Name    pulumi.StringPtrInput
	Project pulumi.StringInput
	// `Source` for the `Ruleset`.
	Source SourcePtrInput
}

func (RulesetArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*rulesetArgs)(nil)).Elem()
}

type RulesetInput interface {
	pulumi.Input

	ToRulesetOutput() RulesetOutput
	ToRulesetOutputWithContext(ctx context.Context) RulesetOutput
}

func (*Ruleset) ElementType() reflect.Type {
	return reflect.TypeOf((*Ruleset)(nil))
}

func (i *Ruleset) ToRulesetOutput() RulesetOutput {
	return i.ToRulesetOutputWithContext(context.Background())
}

func (i *Ruleset) ToRulesetOutputWithContext(ctx context.Context) RulesetOutput {
	return pulumi.ToOutputWithContext(ctx, i).(RulesetOutput)
}

type RulesetOutput struct {
	*pulumi.OutputState
}

func (RulesetOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*Ruleset)(nil))
}

func (o RulesetOutput) ToRulesetOutput() RulesetOutput {
	return o
}

func (o RulesetOutput) ToRulesetOutputWithContext(ctx context.Context) RulesetOutput {
	return o
}

func init() {
	pulumi.RegisterOutputType(RulesetOutput{})
}
