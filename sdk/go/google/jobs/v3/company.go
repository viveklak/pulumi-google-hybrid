// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package v3

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Creates a new company entity.
type Company struct {
	pulumi.CustomResourceState

	// Optional. The URI to employer's career site or careers page on the employer's web site, for example, "https://careers.google.com".
	CareerSiteUri pulumi.StringOutput `pulumi:"careerSiteUri"`
	// Derived details about the company.
	DerivedInfo CompanyDerivedInfoResponseOutput `pulumi:"derivedInfo"`
	// Required. The display name of the company, for example, "Google LLC".
	DisplayName pulumi.StringOutput `pulumi:"displayName"`
	// Optional. Equal Employment Opportunity legal disclaimer text to be associated with all jobs, and typically to be displayed in all roles. The maximum number of allowed characters is 500.
	EeoText pulumi.StringOutput `pulumi:"eeoText"`
	// Required. Client side company identifier, used to uniquely identify the company. The maximum number of allowed characters is 255.
	ExternalId pulumi.StringOutput `pulumi:"externalId"`
	// Optional. The street address of the company's main headquarters, which may be different from the job location. The service attempts to geolocate the provided address, and populates a more specific location wherever possible in DerivedInfo.headquarters_location.
	HeadquartersAddress pulumi.StringOutput `pulumi:"headquartersAddress"`
	// Optional. Set to true if it is the hiring agency that post jobs for other employers. Defaults to false if not provided.
	HiringAgency pulumi.BoolOutput `pulumi:"hiringAgency"`
	// Optional. A URI that hosts the employer's company logo.
	ImageUri pulumi.StringOutput `pulumi:"imageUri"`
	// Optional. A list of keys of filterable Job.custom_attributes, whose corresponding `string_values` are used in keyword search. Jobs with `string_values` under these specified field keys are returned if any of the values matches the search keyword. Custom field values with parenthesis, brackets and special symbols won't be properly searchable, and those keyword queries need to be surrounded by quotes.
	KeywordSearchableJobCustomAttributes pulumi.StringArrayOutput `pulumi:"keywordSearchableJobCustomAttributes"`
	// Required during company update. The resource name for a company. This is generated by the service when a company is created. The format is "projects/{project_id}/companies/{company_id}", for example, "projects/api-test-project/companies/foo".
	Name pulumi.StringOutput `pulumi:"name"`
	// Optional. The employer's company size.
	Size pulumi.StringOutput `pulumi:"size"`
	// Indicates whether a company is flagged to be suspended from public availability by the service when job content appears suspicious, abusive, or spammy.
	Suspended pulumi.BoolOutput `pulumi:"suspended"`
	// Optional. The URI representing the company's primary web site or home page, for example, "https://www.google.com". The maximum number of allowed characters is 255.
	WebsiteUri pulumi.StringOutput `pulumi:"websiteUri"`
}

// NewCompany registers a new resource with the given unique name, arguments, and options.
func NewCompany(ctx *pulumi.Context,
	name string, args *CompanyArgs, opts ...pulumi.ResourceOption) (*Company, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.Project == nil {
		return nil, errors.New("invalid value for required argument 'Project'")
	}
	var resource Company
	err := ctx.RegisterResource("google-native:jobs/v3:Company", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetCompany gets an existing Company resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetCompany(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *CompanyState, opts ...pulumi.ResourceOption) (*Company, error) {
	var resource Company
	err := ctx.ReadResource("google-native:jobs/v3:Company", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering Company resources.
type companyState struct {
}

type CompanyState struct {
}

func (CompanyState) ElementType() reflect.Type {
	return reflect.TypeOf((*companyState)(nil)).Elem()
}

type companyArgs struct {
	// Optional. The URI to employer's career site or careers page on the employer's web site, for example, "https://careers.google.com".
	CareerSiteUri *string `pulumi:"careerSiteUri"`
	// Required. The display name of the company, for example, "Google LLC".
	DisplayName *string `pulumi:"displayName"`
	// Optional. Equal Employment Opportunity legal disclaimer text to be associated with all jobs, and typically to be displayed in all roles. The maximum number of allowed characters is 500.
	EeoText *string `pulumi:"eeoText"`
	// Required. Client side company identifier, used to uniquely identify the company. The maximum number of allowed characters is 255.
	ExternalId *string `pulumi:"externalId"`
	// Optional. The street address of the company's main headquarters, which may be different from the job location. The service attempts to geolocate the provided address, and populates a more specific location wherever possible in DerivedInfo.headquarters_location.
	HeadquartersAddress *string `pulumi:"headquartersAddress"`
	// Optional. Set to true if it is the hiring agency that post jobs for other employers. Defaults to false if not provided.
	HiringAgency *bool `pulumi:"hiringAgency"`
	// Optional. A URI that hosts the employer's company logo.
	ImageUri *string `pulumi:"imageUri"`
	// Optional. A list of keys of filterable Job.custom_attributes, whose corresponding `string_values` are used in keyword search. Jobs with `string_values` under these specified field keys are returned if any of the values matches the search keyword. Custom field values with parenthesis, brackets and special symbols won't be properly searchable, and those keyword queries need to be surrounded by quotes.
	KeywordSearchableJobCustomAttributes []string `pulumi:"keywordSearchableJobCustomAttributes"`
	// Required during company update. The resource name for a company. This is generated by the service when a company is created. The format is "projects/{project_id}/companies/{company_id}", for example, "projects/api-test-project/companies/foo".
	Name    *string `pulumi:"name"`
	Project string  `pulumi:"project"`
	// Optional. The employer's company size.
	Size *string `pulumi:"size"`
	// Optional. The URI representing the company's primary web site or home page, for example, "https://www.google.com". The maximum number of allowed characters is 255.
	WebsiteUri *string `pulumi:"websiteUri"`
}

// The set of arguments for constructing a Company resource.
type CompanyArgs struct {
	// Optional. The URI to employer's career site or careers page on the employer's web site, for example, "https://careers.google.com".
	CareerSiteUri pulumi.StringPtrInput
	// Required. The display name of the company, for example, "Google LLC".
	DisplayName pulumi.StringPtrInput
	// Optional. Equal Employment Opportunity legal disclaimer text to be associated with all jobs, and typically to be displayed in all roles. The maximum number of allowed characters is 500.
	EeoText pulumi.StringPtrInput
	// Required. Client side company identifier, used to uniquely identify the company. The maximum number of allowed characters is 255.
	ExternalId pulumi.StringPtrInput
	// Optional. The street address of the company's main headquarters, which may be different from the job location. The service attempts to geolocate the provided address, and populates a more specific location wherever possible in DerivedInfo.headquarters_location.
	HeadquartersAddress pulumi.StringPtrInput
	// Optional. Set to true if it is the hiring agency that post jobs for other employers. Defaults to false if not provided.
	HiringAgency pulumi.BoolPtrInput
	// Optional. A URI that hosts the employer's company logo.
	ImageUri pulumi.StringPtrInput
	// Optional. A list of keys of filterable Job.custom_attributes, whose corresponding `string_values` are used in keyword search. Jobs with `string_values` under these specified field keys are returned if any of the values matches the search keyword. Custom field values with parenthesis, brackets and special symbols won't be properly searchable, and those keyword queries need to be surrounded by quotes.
	KeywordSearchableJobCustomAttributes pulumi.StringArrayInput
	// Required during company update. The resource name for a company. This is generated by the service when a company is created. The format is "projects/{project_id}/companies/{company_id}", for example, "projects/api-test-project/companies/foo".
	Name    pulumi.StringPtrInput
	Project pulumi.StringInput
	// Optional. The employer's company size.
	Size *CompanySize
	// Optional. The URI representing the company's primary web site or home page, for example, "https://www.google.com". The maximum number of allowed characters is 255.
	WebsiteUri pulumi.StringPtrInput
}

func (CompanyArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*companyArgs)(nil)).Elem()
}

type CompanyInput interface {
	pulumi.Input

	ToCompanyOutput() CompanyOutput
	ToCompanyOutputWithContext(ctx context.Context) CompanyOutput
}

func (*Company) ElementType() reflect.Type {
	return reflect.TypeOf((*Company)(nil))
}

func (i *Company) ToCompanyOutput() CompanyOutput {
	return i.ToCompanyOutputWithContext(context.Background())
}

func (i *Company) ToCompanyOutputWithContext(ctx context.Context) CompanyOutput {
	return pulumi.ToOutputWithContext(ctx, i).(CompanyOutput)
}

type CompanyOutput struct {
	*pulumi.OutputState
}

func (CompanyOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*Company)(nil))
}

func (o CompanyOutput) ToCompanyOutput() CompanyOutput {
	return o
}

func (o CompanyOutput) ToCompanyOutputWithContext(ctx context.Context) CompanyOutput {
	return o
}

func init() {
	pulumi.RegisterOutputType(CompanyOutput{})
}
