// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package v2beta1

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Creates a knowledge base. Note: The `projects.agent.knowledgeBases` resource is deprecated; only use `projects.knowledgeBases`.
type KnowledgeBase struct {
	pulumi.CustomResourceState

	// Required. The display name of the knowledge base. The name must be 1024 bytes or less; otherwise, the creation request fails.
	DisplayName pulumi.StringOutput `pulumi:"displayName"`
	// Language which represents the KnowledgeBase. When the KnowledgeBase is created/updated, this is populated for all non en-us languages. If not populated, the default language en-us applies.
	LanguageCode pulumi.StringOutput `pulumi:"languageCode"`
	// The knowledge base resource name. The name must be empty when creating a knowledge base. Format: `projects//locations//knowledgeBases/`.
	Name pulumi.StringOutput `pulumi:"name"`
}

// NewKnowledgeBase registers a new resource with the given unique name, arguments, and options.
func NewKnowledgeBase(ctx *pulumi.Context,
	name string, args *KnowledgeBaseArgs, opts ...pulumi.ResourceOption) (*KnowledgeBase, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.Location == nil {
		return nil, errors.New("invalid value for required argument 'Location'")
	}
	if args.Project == nil {
		return nil, errors.New("invalid value for required argument 'Project'")
	}
	var resource KnowledgeBase
	err := ctx.RegisterResource("google-native:dialogflow/v2beta1:KnowledgeBase", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetKnowledgeBase gets an existing KnowledgeBase resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetKnowledgeBase(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *KnowledgeBaseState, opts ...pulumi.ResourceOption) (*KnowledgeBase, error) {
	var resource KnowledgeBase
	err := ctx.ReadResource("google-native:dialogflow/v2beta1:KnowledgeBase", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering KnowledgeBase resources.
type knowledgeBaseState struct {
}

type KnowledgeBaseState struct {
}

func (KnowledgeBaseState) ElementType() reflect.Type {
	return reflect.TypeOf((*knowledgeBaseState)(nil)).Elem()
}

type knowledgeBaseArgs struct {
	// Required. The display name of the knowledge base. The name must be 1024 bytes or less; otherwise, the creation request fails.
	DisplayName *string `pulumi:"displayName"`
	// Language which represents the KnowledgeBase. When the KnowledgeBase is created/updated, this is populated for all non en-us languages. If not populated, the default language en-us applies.
	LanguageCode *string `pulumi:"languageCode"`
	Location     string  `pulumi:"location"`
	// The knowledge base resource name. The name must be empty when creating a knowledge base. Format: `projects//locations//knowledgeBases/`.
	Name    *string `pulumi:"name"`
	Project string  `pulumi:"project"`
}

// The set of arguments for constructing a KnowledgeBase resource.
type KnowledgeBaseArgs struct {
	// Required. The display name of the knowledge base. The name must be 1024 bytes or less; otherwise, the creation request fails.
	DisplayName pulumi.StringPtrInput
	// Language which represents the KnowledgeBase. When the KnowledgeBase is created/updated, this is populated for all non en-us languages. If not populated, the default language en-us applies.
	LanguageCode pulumi.StringPtrInput
	Location     pulumi.StringInput
	// The knowledge base resource name. The name must be empty when creating a knowledge base. Format: `projects//locations//knowledgeBases/`.
	Name    pulumi.StringPtrInput
	Project pulumi.StringInput
}

func (KnowledgeBaseArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*knowledgeBaseArgs)(nil)).Elem()
}

type KnowledgeBaseInput interface {
	pulumi.Input

	ToKnowledgeBaseOutput() KnowledgeBaseOutput
	ToKnowledgeBaseOutputWithContext(ctx context.Context) KnowledgeBaseOutput
}

func (*KnowledgeBase) ElementType() reflect.Type {
	return reflect.TypeOf((*KnowledgeBase)(nil))
}

func (i *KnowledgeBase) ToKnowledgeBaseOutput() KnowledgeBaseOutput {
	return i.ToKnowledgeBaseOutputWithContext(context.Background())
}

func (i *KnowledgeBase) ToKnowledgeBaseOutputWithContext(ctx context.Context) KnowledgeBaseOutput {
	return pulumi.ToOutputWithContext(ctx, i).(KnowledgeBaseOutput)
}

type KnowledgeBaseOutput struct {
	*pulumi.OutputState
}

func (KnowledgeBaseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*KnowledgeBase)(nil))
}

func (o KnowledgeBaseOutput) ToKnowledgeBaseOutput() KnowledgeBaseOutput {
	return o
}

func (o KnowledgeBaseOutput) ToKnowledgeBaseOutputWithContext(ctx context.Context) KnowledgeBaseOutput {
	return o
}

func init() {
	pulumi.RegisterOutputType(KnowledgeBaseOutput{})
}
