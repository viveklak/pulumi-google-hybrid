// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package v1beta2

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Creates a new Response Policy Rule.
type ResponsePolicyRule struct {
	pulumi.CustomResourceState

	// Answer this query with a behavior rather than DNS data.
	Behavior pulumi.StringOutput `pulumi:"behavior"`
	// The DNS name (wildcard or exact) to apply this rule to. Must be unique within the Response Policy Rule.
	DnsName pulumi.StringOutput `pulumi:"dnsName"`
	Kind    pulumi.StringOutput `pulumi:"kind"`
	// Answer this query directly with DNS data. These ResourceRecordSets override any other DNS behavior for the matched name; in particular they override private zones, the public internet, and GCP internal DNS. No SOA nor NS types are allowed.
	LocalData ResponsePolicyRuleLocalDataResponseOutput `pulumi:"localData"`
	// An identifier for this rule. Must be unique with the ResponsePolicy.
	RuleName pulumi.StringOutput `pulumi:"ruleName"`
}

// NewResponsePolicyRule registers a new resource with the given unique name, arguments, and options.
func NewResponsePolicyRule(ctx *pulumi.Context,
	name string, args *ResponsePolicyRuleArgs, opts ...pulumi.ResourceOption) (*ResponsePolicyRule, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.Project == nil {
		return nil, errors.New("invalid value for required argument 'Project'")
	}
	if args.ResponsePolicy == nil {
		return nil, errors.New("invalid value for required argument 'ResponsePolicy'")
	}
	var resource ResponsePolicyRule
	err := ctx.RegisterResource("google-native:dns/v1beta2:ResponsePolicyRule", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetResponsePolicyRule gets an existing ResponsePolicyRule resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetResponsePolicyRule(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *ResponsePolicyRuleState, opts ...pulumi.ResourceOption) (*ResponsePolicyRule, error) {
	var resource ResponsePolicyRule
	err := ctx.ReadResource("google-native:dns/v1beta2:ResponsePolicyRule", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering ResponsePolicyRule resources.
type responsePolicyRuleState struct {
}

type ResponsePolicyRuleState struct {
}

func (ResponsePolicyRuleState) ElementType() reflect.Type {
	return reflect.TypeOf((*responsePolicyRuleState)(nil)).Elem()
}

type responsePolicyRuleArgs struct {
	// Answer this query with a behavior rather than DNS data.
	Behavior          *string `pulumi:"behavior"`
	ClientOperationId *string `pulumi:"clientOperationId"`
	// The DNS name (wildcard or exact) to apply this rule to. Must be unique within the Response Policy Rule.
	DnsName *string `pulumi:"dnsName"`
	Kind    *string `pulumi:"kind"`
	// Answer this query directly with DNS data. These ResourceRecordSets override any other DNS behavior for the matched name; in particular they override private zones, the public internet, and GCP internal DNS. No SOA nor NS types are allowed.
	LocalData      *ResponsePolicyRuleLocalData `pulumi:"localData"`
	Project        string                       `pulumi:"project"`
	ResponsePolicy string                       `pulumi:"responsePolicy"`
	// An identifier for this rule. Must be unique with the ResponsePolicy.
	RuleName *string `pulumi:"ruleName"`
}

// The set of arguments for constructing a ResponsePolicyRule resource.
type ResponsePolicyRuleArgs struct {
	// Answer this query with a behavior rather than DNS data.
	Behavior          *ResponsePolicyRuleBehavior
	ClientOperationId pulumi.StringPtrInput
	// The DNS name (wildcard or exact) to apply this rule to. Must be unique within the Response Policy Rule.
	DnsName pulumi.StringPtrInput
	Kind    pulumi.StringPtrInput
	// Answer this query directly with DNS data. These ResourceRecordSets override any other DNS behavior for the matched name; in particular they override private zones, the public internet, and GCP internal DNS. No SOA nor NS types are allowed.
	LocalData      ResponsePolicyRuleLocalDataPtrInput
	Project        pulumi.StringInput
	ResponsePolicy pulumi.StringInput
	// An identifier for this rule. Must be unique with the ResponsePolicy.
	RuleName pulumi.StringPtrInput
}

func (ResponsePolicyRuleArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*responsePolicyRuleArgs)(nil)).Elem()
}

type ResponsePolicyRuleInput interface {
	pulumi.Input

	ToResponsePolicyRuleOutput() ResponsePolicyRuleOutput
	ToResponsePolicyRuleOutputWithContext(ctx context.Context) ResponsePolicyRuleOutput
}

func (*ResponsePolicyRule) ElementType() reflect.Type {
	return reflect.TypeOf((*ResponsePolicyRule)(nil))
}

func (i *ResponsePolicyRule) ToResponsePolicyRuleOutput() ResponsePolicyRuleOutput {
	return i.ToResponsePolicyRuleOutputWithContext(context.Background())
}

func (i *ResponsePolicyRule) ToResponsePolicyRuleOutputWithContext(ctx context.Context) ResponsePolicyRuleOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ResponsePolicyRuleOutput)
}

type ResponsePolicyRuleOutput struct {
	*pulumi.OutputState
}

func (ResponsePolicyRuleOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*ResponsePolicyRule)(nil))
}

func (o ResponsePolicyRuleOutput) ToResponsePolicyRuleOutput() ResponsePolicyRuleOutput {
	return o
}

func (o ResponsePolicyRuleOutput) ToResponsePolicyRuleOutputWithContext(ctx context.Context) ResponsePolicyRuleOutput {
	return o
}

func init() {
	pulumi.RegisterOutputType(ResponsePolicyRuleOutput{})
}
