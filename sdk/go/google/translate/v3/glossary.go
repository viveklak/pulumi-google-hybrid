// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package v3

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Creates a glossary and returns the long-running operation. Returns NOT_FOUND, if the project doesn't exist.
type Glossary struct {
	pulumi.CustomResourceState

	// When the glossary creation was finished.
	EndTime pulumi.StringOutput `pulumi:"endTime"`
	// The number of entries defined in the glossary.
	EntryCount pulumi.IntOutput `pulumi:"entryCount"`
	// Required. Provides examples to build the glossary from. Total glossary must not exceed 10M Unicode codepoints.
	InputConfig GlossaryInputConfigResponseOutput `pulumi:"inputConfig"`
	// Used with equivalent term set glossaries.
	LanguageCodesSet LanguageCodesSetResponseOutput `pulumi:"languageCodesSet"`
	// Used with unidirectional glossaries.
	LanguagePair LanguageCodePairResponseOutput `pulumi:"languagePair"`
	// Required. The resource name of the glossary. Glossary names have the form `projects/{project-number-or-id}/locations/{location-id}/glossaries/{glossary-id}`.
	Name pulumi.StringOutput `pulumi:"name"`
	// When CreateGlossary was called.
	SubmitTime pulumi.StringOutput `pulumi:"submitTime"`
}

// NewGlossary registers a new resource with the given unique name, arguments, and options.
func NewGlossary(ctx *pulumi.Context,
	name string, args *GlossaryArgs, opts ...pulumi.ResourceOption) (*Glossary, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.Location == nil {
		return nil, errors.New("invalid value for required argument 'Location'")
	}
	if args.Project == nil {
		return nil, errors.New("invalid value for required argument 'Project'")
	}
	var resource Glossary
	err := ctx.RegisterResource("google-native:translate/v3:Glossary", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetGlossary gets an existing Glossary resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetGlossary(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *GlossaryState, opts ...pulumi.ResourceOption) (*Glossary, error) {
	var resource Glossary
	err := ctx.ReadResource("google-native:translate/v3:Glossary", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering Glossary resources.
type glossaryState struct {
	// When the glossary creation was finished.
	EndTime *string `pulumi:"endTime"`
	// The number of entries defined in the glossary.
	EntryCount *int `pulumi:"entryCount"`
	// Required. Provides examples to build the glossary from. Total glossary must not exceed 10M Unicode codepoints.
	InputConfig *GlossaryInputConfigResponse `pulumi:"inputConfig"`
	// Used with equivalent term set glossaries.
	LanguageCodesSet *LanguageCodesSetResponse `pulumi:"languageCodesSet"`
	// Used with unidirectional glossaries.
	LanguagePair *LanguageCodePairResponse `pulumi:"languagePair"`
	// Required. The resource name of the glossary. Glossary names have the form `projects/{project-number-or-id}/locations/{location-id}/glossaries/{glossary-id}`.
	Name *string `pulumi:"name"`
	// When CreateGlossary was called.
	SubmitTime *string `pulumi:"submitTime"`
}

type GlossaryState struct {
	// When the glossary creation was finished.
	EndTime pulumi.StringPtrInput
	// The number of entries defined in the glossary.
	EntryCount pulumi.IntPtrInput
	// Required. Provides examples to build the glossary from. Total glossary must not exceed 10M Unicode codepoints.
	InputConfig GlossaryInputConfigResponsePtrInput
	// Used with equivalent term set glossaries.
	LanguageCodesSet LanguageCodesSetResponsePtrInput
	// Used with unidirectional glossaries.
	LanguagePair LanguageCodePairResponsePtrInput
	// Required. The resource name of the glossary. Glossary names have the form `projects/{project-number-or-id}/locations/{location-id}/glossaries/{glossary-id}`.
	Name pulumi.StringPtrInput
	// When CreateGlossary was called.
	SubmitTime pulumi.StringPtrInput
}

func (GlossaryState) ElementType() reflect.Type {
	return reflect.TypeOf((*glossaryState)(nil)).Elem()
}

type glossaryArgs struct {
	// Required. Provides examples to build the glossary from. Total glossary must not exceed 10M Unicode codepoints.
	InputConfig *GlossaryInputConfig `pulumi:"inputConfig"`
	// Used with equivalent term set glossaries.
	LanguageCodesSet *LanguageCodesSet `pulumi:"languageCodesSet"`
	// Used with unidirectional glossaries.
	LanguagePair *LanguageCodePair `pulumi:"languagePair"`
	Location     string            `pulumi:"location"`
	// Required. The resource name of the glossary. Glossary names have the form `projects/{project-number-or-id}/locations/{location-id}/glossaries/{glossary-id}`.
	Name    *string `pulumi:"name"`
	Project string  `pulumi:"project"`
}

// The set of arguments for constructing a Glossary resource.
type GlossaryArgs struct {
	// Required. Provides examples to build the glossary from. Total glossary must not exceed 10M Unicode codepoints.
	InputConfig GlossaryInputConfigPtrInput
	// Used with equivalent term set glossaries.
	LanguageCodesSet LanguageCodesSetPtrInput
	// Used with unidirectional glossaries.
	LanguagePair LanguageCodePairPtrInput
	Location     pulumi.StringInput
	// Required. The resource name of the glossary. Glossary names have the form `projects/{project-number-or-id}/locations/{location-id}/glossaries/{glossary-id}`.
	Name    pulumi.StringPtrInput
	Project pulumi.StringInput
}

func (GlossaryArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*glossaryArgs)(nil)).Elem()
}

type GlossaryInput interface {
	pulumi.Input

	ToGlossaryOutput() GlossaryOutput
	ToGlossaryOutputWithContext(ctx context.Context) GlossaryOutput
}

func (*Glossary) ElementType() reflect.Type {
	return reflect.TypeOf((*Glossary)(nil))
}

func (i *Glossary) ToGlossaryOutput() GlossaryOutput {
	return i.ToGlossaryOutputWithContext(context.Background())
}

func (i *Glossary) ToGlossaryOutputWithContext(ctx context.Context) GlossaryOutput {
	return pulumi.ToOutputWithContext(ctx, i).(GlossaryOutput)
}

type GlossaryOutput struct {
	*pulumi.OutputState
}

func (GlossaryOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*Glossary)(nil))
}

func (o GlossaryOutput) ToGlossaryOutput() GlossaryOutput {
	return o
}

func (o GlossaryOutput) ToGlossaryOutputWithContext(ctx context.Context) GlossaryOutput {
	return o
}

func init() {
	pulumi.RegisterOutputType(GlossaryOutput{})
}
