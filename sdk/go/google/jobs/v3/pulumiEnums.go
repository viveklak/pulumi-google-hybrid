// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package v3

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Optional. The employer's company size.
type CompanySize pulumi.String

const (
	// Default value if the size is not specified.
	CompanySizeCompanySizeUnspecified = CompanySize("COMPANY_SIZE_UNSPECIFIED")
	// The company has less than 50 employees.
	CompanySizeMini = CompanySize("MINI")
	// The company has between 50 and 99 employees.
	CompanySizeSmall = CompanySize("SMALL")
	// The company has between 100 and 499 employees.
	CompanySizeSmedium = CompanySize("SMEDIUM")
	// The company has between 500 and 999 employees.
	CompanySizeMedium = CompanySize("MEDIUM")
	// The company has between 1,000 and 4,999 employees.
	CompanySizeBig = CompanySize("BIG")
	// The company has between 5,000 and 9,999 employees.
	CompanySizeBigger = CompanySize("BIGGER")
	// The company has 10,000 or more employees.
	CompanySizeGiant = CompanySize("GIANT")
)

func (CompanySize) ElementType() reflect.Type {
	return reflect.TypeOf((*pulumi.String)(nil)).Elem()
}

func (e CompanySize) ToStringOutput() pulumi.StringOutput {
	return pulumi.ToOutput(pulumi.String(e)).(pulumi.StringOutput)
}

func (e CompanySize) ToStringOutputWithContext(ctx context.Context) pulumi.StringOutput {
	return pulumi.ToOutputWithContext(ctx, pulumi.String(e)).(pulumi.StringOutput)
}

func (e CompanySize) ToStringPtrOutput() pulumi.StringPtrOutput {
	return pulumi.String(e).ToStringPtrOutputWithContext(context.Background())
}

func (e CompanySize) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return pulumi.String(e).ToStringOutputWithContext(ctx).ToStringPtrOutputWithContext(ctx)
}

// Optional. Compensation type. Default is CompensationUnit.COMPENSATION_TYPE_UNSPECIFIED.
type CompensationEntryType pulumi.String

const (
	// Default value.
	CompensationEntryTypeCompensationTypeUnspecified = CompensationEntryType("COMPENSATION_TYPE_UNSPECIFIED")
	// Base compensation: Refers to the fixed amount of money paid to an employee by an employer in return for work performed. Base compensation does not include benefits, bonuses or any other potential compensation from an employer.
	CompensationEntryTypeBase = CompensationEntryType("BASE")
	// Bonus.
	CompensationEntryTypeBonus = CompensationEntryType("BONUS")
	// Signing bonus.
	CompensationEntryTypeSigningBonus = CompensationEntryType("SIGNING_BONUS")
	// Equity.
	CompensationEntryTypeEquity = CompensationEntryType("EQUITY")
	// Profit sharing.
	CompensationEntryTypeProfitSharing = CompensationEntryType("PROFIT_SHARING")
	// Commission.
	CompensationEntryTypeCommissions = CompensationEntryType("COMMISSIONS")
	// Tips.
	CompensationEntryTypeTips = CompensationEntryType("TIPS")
	// Other compensation type.
	CompensationEntryTypeOtherCompensationType = CompensationEntryType("OTHER_COMPENSATION_TYPE")
)

func (CompensationEntryType) ElementType() reflect.Type {
	return reflect.TypeOf((*pulumi.String)(nil)).Elem()
}

func (e CompensationEntryType) ToStringOutput() pulumi.StringOutput {
	return pulumi.ToOutput(pulumi.String(e)).(pulumi.StringOutput)
}

func (e CompensationEntryType) ToStringOutputWithContext(ctx context.Context) pulumi.StringOutput {
	return pulumi.ToOutputWithContext(ctx, pulumi.String(e)).(pulumi.StringOutput)
}

func (e CompensationEntryType) ToStringPtrOutput() pulumi.StringPtrOutput {
	return pulumi.String(e).ToStringPtrOutputWithContext(context.Background())
}

func (e CompensationEntryType) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return pulumi.String(e).ToStringOutputWithContext(ctx).ToStringPtrOutputWithContext(ctx)
}

// Optional. Frequency of the specified amount. Default is CompensationUnit.COMPENSATION_UNIT_UNSPECIFIED.
type CompensationEntryUnit pulumi.String

const (
	// Default value.
	CompensationEntryUnitCompensationUnitUnspecified = CompensationEntryUnit("COMPENSATION_UNIT_UNSPECIFIED")
	// Hourly.
	CompensationEntryUnitHourly = CompensationEntryUnit("HOURLY")
	// Daily.
	CompensationEntryUnitDaily = CompensationEntryUnit("DAILY")
	// Weekly
	CompensationEntryUnitWeekly = CompensationEntryUnit("WEEKLY")
	// Monthly.
	CompensationEntryUnitMonthly = CompensationEntryUnit("MONTHLY")
	// Yearly.
	CompensationEntryUnitYearly = CompensationEntryUnit("YEARLY")
	// One time.
	CompensationEntryUnitOneTime = CompensationEntryUnit("ONE_TIME")
	// Other compensation units.
	CompensationEntryUnitOtherCompensationUnit = CompensationEntryUnit("OTHER_COMPENSATION_UNIT")
)

func (CompensationEntryUnit) ElementType() reflect.Type {
	return reflect.TypeOf((*pulumi.String)(nil)).Elem()
}

func (e CompensationEntryUnit) ToStringOutput() pulumi.StringOutput {
	return pulumi.ToOutput(pulumi.String(e)).(pulumi.StringOutput)
}

func (e CompensationEntryUnit) ToStringOutputWithContext(ctx context.Context) pulumi.StringOutput {
	return pulumi.ToOutputWithContext(ctx, pulumi.String(e)).(pulumi.StringOutput)
}

func (e CompensationEntryUnit) ToStringPtrOutput() pulumi.StringPtrOutput {
	return pulumi.String(e).ToStringPtrOutputWithContext(context.Background())
}

func (e CompensationEntryUnit) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return pulumi.String(e).ToStringOutputWithContext(ctx).ToStringPtrOutputWithContext(ctx)
}

type JobDegreeTypesItem pulumi.String

const (
	// Default value. Represents no degree, or early childhood education. Maps to ISCED code 0. Ex) Kindergarten
	JobDegreeTypesItemDegreeTypeUnspecified = JobDegreeTypesItem("DEGREE_TYPE_UNSPECIFIED")
	// Primary education which is typically the first stage of compulsory education. ISCED code 1. Ex) Elementary school
	JobDegreeTypesItemPrimaryEducation = JobDegreeTypesItem("PRIMARY_EDUCATION")
	// Lower secondary education; First stage of secondary education building on primary education, typically with a more subject-oriented curriculum. ISCED code 2. Ex) Middle school
	JobDegreeTypesItemLowerSecondaryEducation = JobDegreeTypesItem("LOWER_SECONDARY_EDUCATION")
	// Middle education; Second/final stage of secondary education preparing for tertiary education and/or providing skills relevant to employment. Usually with an increased range of subject options and streams. ISCED code 3. Ex) High school
	JobDegreeTypesItemUpperSecondaryEducation = JobDegreeTypesItem("UPPER_SECONDARY_EDUCATION")
	// Adult Remedial Education; Programmes providing learning experiences that build on secondary education and prepare for labour market entry and/or tertiary education. The content is broader than secondary but not as complex as tertiary education. ISCED code 4.
	JobDegreeTypesItemAdultRemedialEducation = JobDegreeTypesItem("ADULT_REMEDIAL_EDUCATION")
	// Associate's or equivalent; Short first tertiary programmes that are typically practically-based, occupationally-specific and prepare for labour market entry. These programmes may also provide a pathway to other tertiary programmes. ISCED code 5.
	JobDegreeTypesItemAssociatesOrEquivalent = JobDegreeTypesItem("ASSOCIATES_OR_EQUIVALENT")
	// Bachelor's or equivalent; Programmes designed to provide intermediate academic and/or professional knowledge, skills and competencies leading to a first tertiary degree or equivalent qualification. ISCED code 6.
	JobDegreeTypesItemBachelorsOrEquivalent = JobDegreeTypesItem("BACHELORS_OR_EQUIVALENT")
	// Master's or equivalent; Programmes designed to provide advanced academic and/or professional knowledge, skills and competencies leading to a second tertiary degree or equivalent qualification. ISCED code 7.
	JobDegreeTypesItemMastersOrEquivalent = JobDegreeTypesItem("MASTERS_OR_EQUIVALENT")
	// Doctoral or equivalent; Programmes designed primarily to lead to an advanced research qualification, usually concluding with the submission and defense of a substantive dissertation of publishable quality based on original research. ISCED code 8.
	JobDegreeTypesItemDoctoralOrEquivalent = JobDegreeTypesItem("DOCTORAL_OR_EQUIVALENT")
)

func (JobDegreeTypesItem) ElementType() reflect.Type {
	return reflect.TypeOf((*pulumi.String)(nil)).Elem()
}

func (e JobDegreeTypesItem) ToStringOutput() pulumi.StringOutput {
	return pulumi.ToOutput(pulumi.String(e)).(pulumi.StringOutput)
}

func (e JobDegreeTypesItem) ToStringOutputWithContext(ctx context.Context) pulumi.StringOutput {
	return pulumi.ToOutputWithContext(ctx, pulumi.String(e)).(pulumi.StringOutput)
}

func (e JobDegreeTypesItem) ToStringPtrOutput() pulumi.StringPtrOutput {
	return pulumi.String(e).ToStringPtrOutputWithContext(context.Background())
}

func (e JobDegreeTypesItem) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return pulumi.String(e).ToStringOutputWithContext(ctx).ToStringPtrOutputWithContext(ctx)
}

// JobDegreeTypesItemArrayInput is an input type that accepts JobDegreeTypesItemArray and JobDegreeTypesItemArrayOutput values.
// You can construct a concrete instance of `JobDegreeTypesItemArrayInput` via:
//
//          JobDegreeTypesItemArray{ JobDegreeTypesItemArgs{...} }
type JobDegreeTypesItemArrayInput interface {
	pulumi.Input

	ToJobDegreeTypesItemArrayOutput() JobDegreeTypesItemArrayOutput
	ToJobDegreeTypesItemArrayOutputWithContext(context.Context) JobDegreeTypesItemArrayOutput
}

type JobDegreeTypesItemArray []JobDegreeTypesItem

func (JobDegreeTypesItemArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]JobDegreeTypesItem)(nil)).Elem()
}

func (i JobDegreeTypesItemArray) ToJobDegreeTypesItemArrayOutput() JobDegreeTypesItemArrayOutput {
	return i.ToJobDegreeTypesItemArrayOutputWithContext(context.Background())
}

func (i JobDegreeTypesItemArray) ToJobDegreeTypesItemArrayOutputWithContext(ctx context.Context) JobDegreeTypesItemArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(JobDegreeTypesItemArrayOutput)
}

type JobDegreeTypesItemArrayOutput struct{ *pulumi.OutputState }

func (JobDegreeTypesItemArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]JobDegreeTypesItem)(nil)).Elem()
}

func (o JobDegreeTypesItemArrayOutput) ToJobDegreeTypesItemArrayOutput() JobDegreeTypesItemArrayOutput {
	return o
}

func (o JobDegreeTypesItemArrayOutput) ToJobDegreeTypesItemArrayOutputWithContext(ctx context.Context) JobDegreeTypesItemArrayOutput {
	return o
}

func (o JobDegreeTypesItemArrayOutput) Index(i pulumi.IntInput) pulumi.StringOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) pulumi.StringOutput {
		return vs[0].([]JobDegreeTypesItem)[vs[1].(int)].ToStringOutput()
	}).(pulumi.StringOutput)
}

type JobEmploymentTypesItem pulumi.String

const (
	// The default value if the employment type is not specified.
	JobEmploymentTypesItemEmploymentTypeUnspecified = JobEmploymentTypesItem("EMPLOYMENT_TYPE_UNSPECIFIED")
	// The job requires working a number of hours that constitute full time employment, typically 40 or more hours per week.
	JobEmploymentTypesItemFullTime = JobEmploymentTypesItem("FULL_TIME")
	// The job entails working fewer hours than a full time job, typically less than 40 hours a week.
	JobEmploymentTypesItemPartTime = JobEmploymentTypesItem("PART_TIME")
	// The job is offered as a contracted, as opposed to a salaried employee, position.
	JobEmploymentTypesItemContractor = JobEmploymentTypesItem("CONTRACTOR")
	// The job is offered as a contracted position with the understanding that it's converted into a full-time position at the end of the contract. Jobs of this type are also returned by a search for EmploymentType.CONTRACTOR jobs.
	JobEmploymentTypesItemContractToHire = JobEmploymentTypesItem("CONTRACT_TO_HIRE")
	// The job is offered as a temporary employment opportunity, usually a short-term engagement.
	JobEmploymentTypesItemTemporary = JobEmploymentTypesItem("TEMPORARY")
	// The job is a fixed-term opportunity for students or entry-level job seekers to obtain on-the-job training, typically offered as a summer position.
	JobEmploymentTypesItemIntern = JobEmploymentTypesItem("INTERN")
	// The is an opportunity for an individual to volunteer, where there's no expectation of compensation for the provided services.
	JobEmploymentTypesItemVolunteer = JobEmploymentTypesItem("VOLUNTEER")
	// The job requires an employee to work on an as-needed basis with a flexible schedule.
	JobEmploymentTypesItemPerDiem = JobEmploymentTypesItem("PER_DIEM")
	// The job involves employing people in remote areas and flying them temporarily to the work site instead of relocating employees and their families permanently.
	JobEmploymentTypesItemFlyInFlyOut = JobEmploymentTypesItem("FLY_IN_FLY_OUT")
	// The job does not fit any of the other listed types.
	JobEmploymentTypesItemOtherEmploymentType = JobEmploymentTypesItem("OTHER_EMPLOYMENT_TYPE")
)

func (JobEmploymentTypesItem) ElementType() reflect.Type {
	return reflect.TypeOf((*pulumi.String)(nil)).Elem()
}

func (e JobEmploymentTypesItem) ToStringOutput() pulumi.StringOutput {
	return pulumi.ToOutput(pulumi.String(e)).(pulumi.StringOutput)
}

func (e JobEmploymentTypesItem) ToStringOutputWithContext(ctx context.Context) pulumi.StringOutput {
	return pulumi.ToOutputWithContext(ctx, pulumi.String(e)).(pulumi.StringOutput)
}

func (e JobEmploymentTypesItem) ToStringPtrOutput() pulumi.StringPtrOutput {
	return pulumi.String(e).ToStringPtrOutputWithContext(context.Background())
}

func (e JobEmploymentTypesItem) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return pulumi.String(e).ToStringOutputWithContext(ctx).ToStringPtrOutputWithContext(ctx)
}

// JobEmploymentTypesItemArrayInput is an input type that accepts JobEmploymentTypesItemArray and JobEmploymentTypesItemArrayOutput values.
// You can construct a concrete instance of `JobEmploymentTypesItemArrayInput` via:
//
//          JobEmploymentTypesItemArray{ JobEmploymentTypesItemArgs{...} }
type JobEmploymentTypesItemArrayInput interface {
	pulumi.Input

	ToJobEmploymentTypesItemArrayOutput() JobEmploymentTypesItemArrayOutput
	ToJobEmploymentTypesItemArrayOutputWithContext(context.Context) JobEmploymentTypesItemArrayOutput
}

type JobEmploymentTypesItemArray []JobEmploymentTypesItem

func (JobEmploymentTypesItemArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]JobEmploymentTypesItem)(nil)).Elem()
}

func (i JobEmploymentTypesItemArray) ToJobEmploymentTypesItemArrayOutput() JobEmploymentTypesItemArrayOutput {
	return i.ToJobEmploymentTypesItemArrayOutputWithContext(context.Background())
}

func (i JobEmploymentTypesItemArray) ToJobEmploymentTypesItemArrayOutputWithContext(ctx context.Context) JobEmploymentTypesItemArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(JobEmploymentTypesItemArrayOutput)
}

type JobEmploymentTypesItemArrayOutput struct{ *pulumi.OutputState }

func (JobEmploymentTypesItemArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]JobEmploymentTypesItem)(nil)).Elem()
}

func (o JobEmploymentTypesItemArrayOutput) ToJobEmploymentTypesItemArrayOutput() JobEmploymentTypesItemArrayOutput {
	return o
}

func (o JobEmploymentTypesItemArrayOutput) ToJobEmploymentTypesItemArrayOutputWithContext(ctx context.Context) JobEmploymentTypesItemArrayOutput {
	return o
}

func (o JobEmploymentTypesItemArrayOutput) Index(i pulumi.IntInput) pulumi.StringOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) pulumi.StringOutput {
		return vs[0].([]JobEmploymentTypesItem)[vs[1].(int)].ToStringOutput()
	}).(pulumi.StringOutput)
}

type JobJobBenefitsItem pulumi.String

const (
	// Default value if the type is not specified.
	JobJobBenefitsItemJobBenefitUnspecified = JobJobBenefitsItem("JOB_BENEFIT_UNSPECIFIED")
	// The job includes access to programs that support child care, such as daycare.
	JobJobBenefitsItemChildCare = JobJobBenefitsItem("CHILD_CARE")
	// The job includes dental services covered by a dental insurance plan.
	JobJobBenefitsItemDental = JobJobBenefitsItem("DENTAL")
	// The job offers specific benefits to domestic partners.
	JobJobBenefitsItemDomesticPartner = JobJobBenefitsItem("DOMESTIC_PARTNER")
	// The job allows for a flexible work schedule.
	JobJobBenefitsItemFlexibleHours = JobJobBenefitsItem("FLEXIBLE_HOURS")
	// The job includes health services covered by a medical insurance plan.
	JobJobBenefitsItemMedical = JobJobBenefitsItem("MEDICAL")
	// The job includes a life insurance plan provided by the employer or available for purchase by the employee.
	JobJobBenefitsItemLifeInsurance = JobJobBenefitsItem("LIFE_INSURANCE")
	// The job allows for a leave of absence to a parent to care for a newborn child.
	JobJobBenefitsItemParentalLeave = JobJobBenefitsItem("PARENTAL_LEAVE")
	// The job includes a workplace retirement plan provided by the employer or available for purchase by the employee.
	JobJobBenefitsItemRetirementPlan = JobJobBenefitsItem("RETIREMENT_PLAN")
	// The job allows for paid time off due to illness.
	JobJobBenefitsItemSickDays = JobJobBenefitsItem("SICK_DAYS")
	// The job includes paid time off for vacation.
	JobJobBenefitsItemVacation = JobJobBenefitsItem("VACATION")
	// The job includes vision services covered by a vision insurance plan.
	JobJobBenefitsItemVision = JobJobBenefitsItem("VISION")
)

func (JobJobBenefitsItem) ElementType() reflect.Type {
	return reflect.TypeOf((*pulumi.String)(nil)).Elem()
}

func (e JobJobBenefitsItem) ToStringOutput() pulumi.StringOutput {
	return pulumi.ToOutput(pulumi.String(e)).(pulumi.StringOutput)
}

func (e JobJobBenefitsItem) ToStringOutputWithContext(ctx context.Context) pulumi.StringOutput {
	return pulumi.ToOutputWithContext(ctx, pulumi.String(e)).(pulumi.StringOutput)
}

func (e JobJobBenefitsItem) ToStringPtrOutput() pulumi.StringPtrOutput {
	return pulumi.String(e).ToStringPtrOutputWithContext(context.Background())
}

func (e JobJobBenefitsItem) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return pulumi.String(e).ToStringOutputWithContext(ctx).ToStringPtrOutputWithContext(ctx)
}

// JobJobBenefitsItemArrayInput is an input type that accepts JobJobBenefitsItemArray and JobJobBenefitsItemArrayOutput values.
// You can construct a concrete instance of `JobJobBenefitsItemArrayInput` via:
//
//          JobJobBenefitsItemArray{ JobJobBenefitsItemArgs{...} }
type JobJobBenefitsItemArrayInput interface {
	pulumi.Input

	ToJobJobBenefitsItemArrayOutput() JobJobBenefitsItemArrayOutput
	ToJobJobBenefitsItemArrayOutputWithContext(context.Context) JobJobBenefitsItemArrayOutput
}

type JobJobBenefitsItemArray []JobJobBenefitsItem

func (JobJobBenefitsItemArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]JobJobBenefitsItem)(nil)).Elem()
}

func (i JobJobBenefitsItemArray) ToJobJobBenefitsItemArrayOutput() JobJobBenefitsItemArrayOutput {
	return i.ToJobJobBenefitsItemArrayOutputWithContext(context.Background())
}

func (i JobJobBenefitsItemArray) ToJobJobBenefitsItemArrayOutputWithContext(ctx context.Context) JobJobBenefitsItemArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(JobJobBenefitsItemArrayOutput)
}

type JobJobBenefitsItemArrayOutput struct{ *pulumi.OutputState }

func (JobJobBenefitsItemArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]JobJobBenefitsItem)(nil)).Elem()
}

func (o JobJobBenefitsItemArrayOutput) ToJobJobBenefitsItemArrayOutput() JobJobBenefitsItemArrayOutput {
	return o
}

func (o JobJobBenefitsItemArrayOutput) ToJobJobBenefitsItemArrayOutputWithContext(ctx context.Context) JobJobBenefitsItemArrayOutput {
	return o
}

func (o JobJobBenefitsItemArrayOutput) Index(i pulumi.IntInput) pulumi.StringOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) pulumi.StringOutput {
		return vs[0].([]JobJobBenefitsItem)[vs[1].(int)].ToStringOutput()
	}).(pulumi.StringOutput)
}

// Optional. The experience level associated with the job, such as "Entry Level".
type JobJobLevel pulumi.String

const (
	// The default value if the level is not specified.
	JobJobLevelJobLevelUnspecified = JobJobLevel("JOB_LEVEL_UNSPECIFIED")
	// Entry-level individual contributors, typically with less than 2 years of experience in a similar role. Includes interns.
	JobJobLevelEntryLevel = JobJobLevel("ENTRY_LEVEL")
	// Experienced individual contributors, typically with 2+ years of experience in a similar role.
	JobJobLevelExperienced = JobJobLevel("EXPERIENCED")
	// Entry- to mid-level managers responsible for managing a team of people.
	JobJobLevelManager = JobJobLevel("MANAGER")
	// Senior-level managers responsible for managing teams of managers.
	JobJobLevelDirector = JobJobLevel("DIRECTOR")
	// Executive-level managers and above, including C-level positions.
	JobJobLevelExecutive = JobJobLevel("EXECUTIVE")
)

func (JobJobLevel) ElementType() reflect.Type {
	return reflect.TypeOf((*pulumi.String)(nil)).Elem()
}

func (e JobJobLevel) ToStringOutput() pulumi.StringOutput {
	return pulumi.ToOutput(pulumi.String(e)).(pulumi.StringOutput)
}

func (e JobJobLevel) ToStringOutputWithContext(ctx context.Context) pulumi.StringOutput {
	return pulumi.ToOutputWithContext(ctx, pulumi.String(e)).(pulumi.StringOutput)
}

func (e JobJobLevel) ToStringPtrOutput() pulumi.StringPtrOutput {
	return pulumi.String(e).ToStringPtrOutputWithContext(context.Background())
}

func (e JobJobLevel) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return pulumi.String(e).ToStringOutputWithContext(ctx).ToStringPtrOutputWithContext(ctx)
}

// Optional. The job PostingRegion (for example, state, country) throughout which the job is available. If this field is set, a LocationFilter in a search query within the job region finds this job posting if an exact location match isn't specified. If this field is set to PostingRegion.NATION or PostingRegion.ADMINISTRATIVE_AREA, setting job Job.addresses to the same location level as this field is strongly recommended.
type JobPostingRegion pulumi.String

const (
	// If the region is unspecified, the job is only returned if it matches the LocationFilter.
	JobPostingRegionPostingRegionUnspecified = JobPostingRegion("POSTING_REGION_UNSPECIFIED")
	// In addition to exact location matching, job posting is returned when the LocationFilter in the search query is in the same administrative area as the returned job posting. For example, if a `ADMINISTRATIVE_AREA` job is posted in "CA, USA", it's returned if LocationFilter has "Mountain View". Administrative area refers to top-level administrative subdivision of this country. For example, US state, IT region, UK constituent nation and JP prefecture.
	JobPostingRegionAdministrativeArea = JobPostingRegion("ADMINISTRATIVE_AREA")
	// In addition to exact location matching, job is returned when LocationFilter in search query is in the same country as this job. For example, if a `NATION_WIDE` job is posted in "USA", it's returned if LocationFilter has 'Mountain View'.
	JobPostingRegionNation = JobPostingRegion("NATION")
	// Job allows employees to work remotely (telecommute). If locations are provided with this value, the job is considered as having a location, but telecommuting is allowed.
	JobPostingRegionTelecommute = JobPostingRegion("TELECOMMUTE")
)

func (JobPostingRegion) ElementType() reflect.Type {
	return reflect.TypeOf((*pulumi.String)(nil)).Elem()
}

func (e JobPostingRegion) ToStringOutput() pulumi.StringOutput {
	return pulumi.ToOutput(pulumi.String(e)).(pulumi.StringOutput)
}

func (e JobPostingRegion) ToStringOutputWithContext(ctx context.Context) pulumi.StringOutput {
	return pulumi.ToOutputWithContext(ctx, pulumi.String(e)).(pulumi.StringOutput)
}

func (e JobPostingRegion) ToStringPtrOutput() pulumi.StringPtrOutput {
	return pulumi.String(e).ToStringPtrOutputWithContext(context.Background())
}

func (e JobPostingRegion) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return pulumi.String(e).ToStringOutputWithContext(ctx).ToStringPtrOutputWithContext(ctx)
}

// Optional. Option for job HTML content sanitization. Applied fields are: * description * applicationInfo.instruction * incentives * qualifications * responsibilities HTML tags in these fields may be stripped if sanitiazation is not disabled. Defaults to HtmlSanitization.SIMPLE_FORMATTING_ONLY.
type ProcessingOptionsHtmlSanitization pulumi.String

const (
	// Default value.
	ProcessingOptionsHtmlSanitizationHtmlSanitizationUnspecified = ProcessingOptionsHtmlSanitization("HTML_SANITIZATION_UNSPECIFIED")
	// Disables sanitization on HTML input.
	ProcessingOptionsHtmlSanitizationHtmlSanitizationDisabled = ProcessingOptionsHtmlSanitization("HTML_SANITIZATION_DISABLED")
	// Sanitizes HTML input, only accepts bold, italic, ordered list, and unordered list markup tags.
	ProcessingOptionsHtmlSanitizationSimpleFormattingOnly = ProcessingOptionsHtmlSanitization("SIMPLE_FORMATTING_ONLY")
)

func (ProcessingOptionsHtmlSanitization) ElementType() reflect.Type {
	return reflect.TypeOf((*pulumi.String)(nil)).Elem()
}

func (e ProcessingOptionsHtmlSanitization) ToStringOutput() pulumi.StringOutput {
	return pulumi.ToOutput(pulumi.String(e)).(pulumi.StringOutput)
}

func (e ProcessingOptionsHtmlSanitization) ToStringOutputWithContext(ctx context.Context) pulumi.StringOutput {
	return pulumi.ToOutputWithContext(ctx, pulumi.String(e)).(pulumi.StringOutput)
}

func (e ProcessingOptionsHtmlSanitization) ToStringPtrOutput() pulumi.StringPtrOutput {
	return pulumi.String(e).ToStringPtrOutputWithContext(context.Background())
}

func (e ProcessingOptionsHtmlSanitization) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return pulumi.String(e).ToStringOutputWithContext(ctx).ToStringPtrOutputWithContext(ctx)
}
