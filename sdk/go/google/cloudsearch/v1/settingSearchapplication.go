// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package v1

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
)

// Creates a search application. **Note:** This API requires an admin account to execute.
type SettingSearchapplication struct {
	pulumi.CustomResourceState
}

// NewSettingSearchapplication registers a new resource with the given unique name, arguments, and options.
func NewSettingSearchapplication(ctx *pulumi.Context,
	name string, args *SettingSearchapplicationArgs, opts ...pulumi.ResourceOption) (*SettingSearchapplication, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.SearchapplicationsId == nil {
		return nil, errors.New("invalid value for required argument 'SearchapplicationsId'")
	}
	var resource SettingSearchapplication
	err := ctx.RegisterResource("google-cloud:cloudsearch/v1:SettingSearchapplication", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetSettingSearchapplication gets an existing SettingSearchapplication resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetSettingSearchapplication(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *SettingSearchapplicationState, opts ...pulumi.ResourceOption) (*SettingSearchapplication, error) {
	var resource SettingSearchapplication
	err := ctx.ReadResource("google-cloud:cloudsearch/v1:SettingSearchapplication", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering SettingSearchapplication resources.
type settingSearchapplicationState struct {
}

type SettingSearchapplicationState struct {
}

func (SettingSearchapplicationState) ElementType() reflect.Type {
	return reflect.TypeOf((*settingSearchapplicationState)(nil)).Elem()
}

type settingSearchapplicationArgs struct {
	// Retrictions applied to the configurations. The maximum number of elements is 10.
	DataSourceRestrictions []DataSourceRestriction `pulumi:"dataSourceRestrictions"`
	// The default fields for returning facet results. The sources specified here also have been included in data_source_restrictions above.
	DefaultFacetOptions []FacetOptions `pulumi:"defaultFacetOptions"`
	// The default options for sorting the search results
	DefaultSortOptions *SortOptions `pulumi:"defaultSortOptions"`
	// Display name of the Search Application. The maximum length is 300 characters.
	DisplayName *string `pulumi:"displayName"`
	// Name of the Search Application. Format: searchapplications/{application_id}.
	Name *string `pulumi:"name"`
	// Output only. IDs of the Long Running Operations (LROs) currently running for this schema. Output only field.
	OperationIds []string `pulumi:"operationIds"`
	// Configuration for ranking results.
	ScoringConfig        *ScoringConfig `pulumi:"scoringConfig"`
	SearchapplicationsId string         `pulumi:"searchapplicationsId"`
	// Configuration for a sources specified in data_source_restrictions.
	SourceConfig []SourceConfig `pulumi:"sourceConfig"`
}

// The set of arguments for constructing a SettingSearchapplication resource.
type SettingSearchapplicationArgs struct {
	// Retrictions applied to the configurations. The maximum number of elements is 10.
	DataSourceRestrictions DataSourceRestrictionArrayInput
	// The default fields for returning facet results. The sources specified here also have been included in data_source_restrictions above.
	DefaultFacetOptions FacetOptionsArrayInput
	// The default options for sorting the search results
	DefaultSortOptions SortOptionsPtrInput
	// Display name of the Search Application. The maximum length is 300 characters.
	DisplayName pulumi.StringPtrInput
	// Name of the Search Application. Format: searchapplications/{application_id}.
	Name pulumi.StringPtrInput
	// Output only. IDs of the Long Running Operations (LROs) currently running for this schema. Output only field.
	OperationIds pulumi.StringArrayInput
	// Configuration for ranking results.
	ScoringConfig        ScoringConfigPtrInput
	SearchapplicationsId pulumi.StringInput
	// Configuration for a sources specified in data_source_restrictions.
	SourceConfig SourceConfigArrayInput
}

func (SettingSearchapplicationArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*settingSearchapplicationArgs)(nil)).Elem()
}

type SettingSearchapplicationInput interface {
	pulumi.Input

	ToSettingSearchapplicationOutput() SettingSearchapplicationOutput
	ToSettingSearchapplicationOutputWithContext(ctx context.Context) SettingSearchapplicationOutput
}

func (*SettingSearchapplication) ElementType() reflect.Type {
	return reflect.TypeOf((*SettingSearchapplication)(nil))
}

func (i *SettingSearchapplication) ToSettingSearchapplicationOutput() SettingSearchapplicationOutput {
	return i.ToSettingSearchapplicationOutputWithContext(context.Background())
}

func (i *SettingSearchapplication) ToSettingSearchapplicationOutputWithContext(ctx context.Context) SettingSearchapplicationOutput {
	return pulumi.ToOutputWithContext(ctx, i).(SettingSearchapplicationOutput)
}

type SettingSearchapplicationOutput struct {
	*pulumi.OutputState
}

func (SettingSearchapplicationOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*SettingSearchapplication)(nil))
}

func (o SettingSearchapplicationOutput) ToSettingSearchapplicationOutput() SettingSearchapplicationOutput {
	return o
}

func (o SettingSearchapplicationOutput) ToSettingSearchapplicationOutputWithContext(ctx context.Context) SettingSearchapplicationOutput {
	return o
}

func init() {
	pulumi.RegisterOutputType(SettingSearchapplicationOutput{})
}
