// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.ComponentModel;
using Pulumi;

namespace Pulumi.GoogleNative.Jobs.V4
{
    /// <summary>
    /// The employer's company size.
    /// </summary>
    [EnumType]
    public readonly struct CompanySize : IEquatable<CompanySize>
    {
        private readonly string _value;

        private CompanySize(string value)
        {
            _value = value ?? throw new ArgumentNullException(nameof(value));
        }

        /// <summary>
        /// Default value if the size isn't specified.
        /// </summary>
        public static CompanySize CompanySizeUnspecified { get; } = new CompanySize("COMPANY_SIZE_UNSPECIFIED");
        /// <summary>
        /// The company has less than 50 employees.
        /// </summary>
        public static CompanySize Mini { get; } = new CompanySize("MINI");
        /// <summary>
        /// The company has between 50 and 99 employees.
        /// </summary>
        public static CompanySize Small { get; } = new CompanySize("SMALL");
        /// <summary>
        /// The company has between 100 and 499 employees.
        /// </summary>
        public static CompanySize Smedium { get; } = new CompanySize("SMEDIUM");
        /// <summary>
        /// The company has between 500 and 999 employees.
        /// </summary>
        public static CompanySize Medium { get; } = new CompanySize("MEDIUM");
        /// <summary>
        /// The company has between 1,000 and 4,999 employees.
        /// </summary>
        public static CompanySize Big { get; } = new CompanySize("BIG");
        /// <summary>
        /// The company has between 5,000 and 9,999 employees.
        /// </summary>
        public static CompanySize Bigger { get; } = new CompanySize("BIGGER");
        /// <summary>
        /// The company has 10,000 or more employees.
        /// </summary>
        public static CompanySize Giant { get; } = new CompanySize("GIANT");

        public static bool operator ==(CompanySize left, CompanySize right) => left.Equals(right);
        public static bool operator !=(CompanySize left, CompanySize right) => !left.Equals(right);

        public static explicit operator string(CompanySize value) => value._value;

        [EditorBrowsable(EditorBrowsableState.Never)]
        public override bool Equals(object? obj) => obj is CompanySize other && Equals(other);
        public bool Equals(CompanySize other) => string.Equals(_value, other._value, StringComparison.Ordinal);

        [EditorBrowsable(EditorBrowsableState.Never)]
        public override int GetHashCode() => _value?.GetHashCode() ?? 0;

        public override string ToString() => _value;
    }

    /// <summary>
    /// Compensation type. Default is CompensationType.COMPENSATION_TYPE_UNSPECIFIED.
    /// </summary>
    [EnumType]
    public readonly struct CompensationEntryType : IEquatable<CompensationEntryType>
    {
        private readonly string _value;

        private CompensationEntryType(string value)
        {
            _value = value ?? throw new ArgumentNullException(nameof(value));
        }

        /// <summary>
        /// Default value.
        /// </summary>
        public static CompensationEntryType CompensationTypeUnspecified { get; } = new CompensationEntryType("COMPENSATION_TYPE_UNSPECIFIED");
        /// <summary>
        /// Base compensation: Refers to the fixed amount of money paid to an employee by an employer in return for work performed. Base compensation does not include benefits, bonuses or any other potential compensation from an employer.
        /// </summary>
        public static CompensationEntryType Base { get; } = new CompensationEntryType("BASE");
        /// <summary>
        /// Bonus.
        /// </summary>
        public static CompensationEntryType Bonus { get; } = new CompensationEntryType("BONUS");
        /// <summary>
        /// Signing bonus.
        /// </summary>
        public static CompensationEntryType SigningBonus { get; } = new CompensationEntryType("SIGNING_BONUS");
        /// <summary>
        /// Equity.
        /// </summary>
        public static CompensationEntryType Equity { get; } = new CompensationEntryType("EQUITY");
        /// <summary>
        /// Profit sharing.
        /// </summary>
        public static CompensationEntryType ProfitSharing { get; } = new CompensationEntryType("PROFIT_SHARING");
        /// <summary>
        /// Commission.
        /// </summary>
        public static CompensationEntryType Commissions { get; } = new CompensationEntryType("COMMISSIONS");
        /// <summary>
        /// Tips.
        /// </summary>
        public static CompensationEntryType Tips { get; } = new CompensationEntryType("TIPS");
        /// <summary>
        /// Other compensation type.
        /// </summary>
        public static CompensationEntryType OtherCompensationType { get; } = new CompensationEntryType("OTHER_COMPENSATION_TYPE");

        public static bool operator ==(CompensationEntryType left, CompensationEntryType right) => left.Equals(right);
        public static bool operator !=(CompensationEntryType left, CompensationEntryType right) => !left.Equals(right);

        public static explicit operator string(CompensationEntryType value) => value._value;

        [EditorBrowsable(EditorBrowsableState.Never)]
        public override bool Equals(object? obj) => obj is CompensationEntryType other && Equals(other);
        public bool Equals(CompensationEntryType other) => string.Equals(_value, other._value, StringComparison.Ordinal);

        [EditorBrowsable(EditorBrowsableState.Never)]
        public override int GetHashCode() => _value?.GetHashCode() ?? 0;

        public override string ToString() => _value;
    }

    /// <summary>
    /// Frequency of the specified amount. Default is CompensationUnit.COMPENSATION_UNIT_UNSPECIFIED.
    /// </summary>
    [EnumType]
    public readonly struct CompensationEntryUnit : IEquatable<CompensationEntryUnit>
    {
        private readonly string _value;

        private CompensationEntryUnit(string value)
        {
            _value = value ?? throw new ArgumentNullException(nameof(value));
        }

        /// <summary>
        /// Default value.
        /// </summary>
        public static CompensationEntryUnit CompensationUnitUnspecified { get; } = new CompensationEntryUnit("COMPENSATION_UNIT_UNSPECIFIED");
        /// <summary>
        /// Hourly.
        /// </summary>
        public static CompensationEntryUnit Hourly { get; } = new CompensationEntryUnit("HOURLY");
        /// <summary>
        /// Daily.
        /// </summary>
        public static CompensationEntryUnit Daily { get; } = new CompensationEntryUnit("DAILY");
        /// <summary>
        /// Weekly
        /// </summary>
        public static CompensationEntryUnit Weekly { get; } = new CompensationEntryUnit("WEEKLY");
        /// <summary>
        /// Monthly.
        /// </summary>
        public static CompensationEntryUnit Monthly { get; } = new CompensationEntryUnit("MONTHLY");
        /// <summary>
        /// Yearly.
        /// </summary>
        public static CompensationEntryUnit Yearly { get; } = new CompensationEntryUnit("YEARLY");
        /// <summary>
        /// One time.
        /// </summary>
        public static CompensationEntryUnit OneTime { get; } = new CompensationEntryUnit("ONE_TIME");
        /// <summary>
        /// Other compensation units.
        /// </summary>
        public static CompensationEntryUnit OtherCompensationUnit { get; } = new CompensationEntryUnit("OTHER_COMPENSATION_UNIT");

        public static bool operator ==(CompensationEntryUnit left, CompensationEntryUnit right) => left.Equals(right);
        public static bool operator !=(CompensationEntryUnit left, CompensationEntryUnit right) => !left.Equals(right);

        public static explicit operator string(CompensationEntryUnit value) => value._value;

        [EditorBrowsable(EditorBrowsableState.Never)]
        public override bool Equals(object? obj) => obj is CompensationEntryUnit other && Equals(other);
        public bool Equals(CompensationEntryUnit other) => string.Equals(_value, other._value, StringComparison.Ordinal);

        [EditorBrowsable(EditorBrowsableState.Never)]
        public override int GetHashCode() => _value?.GetHashCode() ?? 0;

        public override string ToString() => _value;
    }

    [EnumType]
    public readonly struct JobDegreeTypesItem : IEquatable<JobDegreeTypesItem>
    {
        private readonly string _value;

        private JobDegreeTypesItem(string value)
        {
            _value = value ?? throw new ArgumentNullException(nameof(value));
        }

        /// <summary>
        /// Default value. Represents no degree, or early childhood education. Maps to ISCED code 0. Ex) Kindergarten
        /// </summary>
        public static JobDegreeTypesItem DegreeTypeUnspecified { get; } = new JobDegreeTypesItem("DEGREE_TYPE_UNSPECIFIED");
        /// <summary>
        /// Primary education which is typically the first stage of compulsory education. ISCED code 1. Ex) Elementary school
        /// </summary>
        public static JobDegreeTypesItem PrimaryEducation { get; } = new JobDegreeTypesItem("PRIMARY_EDUCATION");
        /// <summary>
        /// Lower secondary education; First stage of secondary education building on primary education, typically with a more subject-oriented curriculum. ISCED code 2. Ex) Middle school
        /// </summary>
        public static JobDegreeTypesItem LowerSecondaryEducation { get; } = new JobDegreeTypesItem("LOWER_SECONDARY_EDUCATION");
        /// <summary>
        /// Middle education; Second/final stage of secondary education preparing for tertiary education and/or providing skills relevant to employment. Usually with an increased range of subject options and streams. ISCED code 3. Ex) High school
        /// </summary>
        public static JobDegreeTypesItem UpperSecondaryEducation { get; } = new JobDegreeTypesItem("UPPER_SECONDARY_EDUCATION");
        /// <summary>
        /// Adult Remedial Education; Programmes providing learning experiences that build on secondary education and prepare for labour market entry and/or tertiary education. The content is broader than secondary but not as complex as tertiary education. ISCED code 4.
        /// </summary>
        public static JobDegreeTypesItem AdultRemedialEducation { get; } = new JobDegreeTypesItem("ADULT_REMEDIAL_EDUCATION");
        /// <summary>
        /// Associate's or equivalent; Short first tertiary programmes that are typically practically-based, occupationally-specific and prepare for labour market entry. These programmes may also provide a pathway to other tertiary programmes. ISCED code 5.
        /// </summary>
        public static JobDegreeTypesItem AssociatesOrEquivalent { get; } = new JobDegreeTypesItem("ASSOCIATES_OR_EQUIVALENT");
        /// <summary>
        /// Bachelor's or equivalent; Programmes designed to provide intermediate academic and/or professional knowledge, skills and competencies leading to a first tertiary degree or equivalent qualification. ISCED code 6.
        /// </summary>
        public static JobDegreeTypesItem BachelorsOrEquivalent { get; } = new JobDegreeTypesItem("BACHELORS_OR_EQUIVALENT");
        /// <summary>
        /// Master's or equivalent; Programmes designed to provide advanced academic and/or professional knowledge, skills and competencies leading to a second tertiary degree or equivalent qualification. ISCED code 7.
        /// </summary>
        public static JobDegreeTypesItem MastersOrEquivalent { get; } = new JobDegreeTypesItem("MASTERS_OR_EQUIVALENT");
        /// <summary>
        /// Doctoral or equivalent; Programmes designed primarily to lead to an advanced research qualification, usually concluding with the submission and defense of a substantive dissertation of publishable quality based on original research. ISCED code 8.
        /// </summary>
        public static JobDegreeTypesItem DoctoralOrEquivalent { get; } = new JobDegreeTypesItem("DOCTORAL_OR_EQUIVALENT");

        public static bool operator ==(JobDegreeTypesItem left, JobDegreeTypesItem right) => left.Equals(right);
        public static bool operator !=(JobDegreeTypesItem left, JobDegreeTypesItem right) => !left.Equals(right);

        public static explicit operator string(JobDegreeTypesItem value) => value._value;

        [EditorBrowsable(EditorBrowsableState.Never)]
        public override bool Equals(object? obj) => obj is JobDegreeTypesItem other && Equals(other);
        public bool Equals(JobDegreeTypesItem other) => string.Equals(_value, other._value, StringComparison.Ordinal);

        [EditorBrowsable(EditorBrowsableState.Never)]
        public override int GetHashCode() => _value?.GetHashCode() ?? 0;

        public override string ToString() => _value;
    }

    [EnumType]
    public readonly struct JobEmploymentTypesItem : IEquatable<JobEmploymentTypesItem>
    {
        private readonly string _value;

        private JobEmploymentTypesItem(string value)
        {
            _value = value ?? throw new ArgumentNullException(nameof(value));
        }

        /// <summary>
        /// The default value if the employment type isn't specified.
        /// </summary>
        public static JobEmploymentTypesItem EmploymentTypeUnspecified { get; } = new JobEmploymentTypesItem("EMPLOYMENT_TYPE_UNSPECIFIED");
        /// <summary>
        /// The job requires working a number of hours that constitute full time employment, typically 40 or more hours per week.
        /// </summary>
        public static JobEmploymentTypesItem FullTime { get; } = new JobEmploymentTypesItem("FULL_TIME");
        /// <summary>
        /// The job entails working fewer hours than a full time job, typically less than 40 hours a week.
        /// </summary>
        public static JobEmploymentTypesItem PartTime { get; } = new JobEmploymentTypesItem("PART_TIME");
        /// <summary>
        /// The job is offered as a contracted, as opposed to a salaried employee, position.
        /// </summary>
        public static JobEmploymentTypesItem Contractor { get; } = new JobEmploymentTypesItem("CONTRACTOR");
        /// <summary>
        /// The job is offered as a contracted position with the understanding that it's converted into a full-time position at the end of the contract. Jobs of this type are also returned by a search for EmploymentType.CONTRACTOR jobs.
        /// </summary>
        public static JobEmploymentTypesItem ContractToHire { get; } = new JobEmploymentTypesItem("CONTRACT_TO_HIRE");
        /// <summary>
        /// The job is offered as a temporary employment opportunity, usually a short-term engagement.
        /// </summary>
        public static JobEmploymentTypesItem Temporary { get; } = new JobEmploymentTypesItem("TEMPORARY");
        /// <summary>
        /// The job is a fixed-term opportunity for students or entry-level job seekers to obtain on-the-job training, typically offered as a summer position.
        /// </summary>
        public static JobEmploymentTypesItem Intern { get; } = new JobEmploymentTypesItem("INTERN");
        /// <summary>
        /// The is an opportunity for an individual to volunteer, where there's no expectation of compensation for the provided services.
        /// </summary>
        public static JobEmploymentTypesItem Volunteer { get; } = new JobEmploymentTypesItem("VOLUNTEER");
        /// <summary>
        /// The job requires an employee to work on an as-needed basis with a flexible schedule.
        /// </summary>
        public static JobEmploymentTypesItem PerDiem { get; } = new JobEmploymentTypesItem("PER_DIEM");
        /// <summary>
        /// The job involves employing people in remote areas and flying them temporarily to the work site instead of relocating employees and their families permanently.
        /// </summary>
        public static JobEmploymentTypesItem FlyInFlyOut { get; } = new JobEmploymentTypesItem("FLY_IN_FLY_OUT");
        /// <summary>
        /// The job does not fit any of the other listed types.
        /// </summary>
        public static JobEmploymentTypesItem OtherEmploymentType { get; } = new JobEmploymentTypesItem("OTHER_EMPLOYMENT_TYPE");

        public static bool operator ==(JobEmploymentTypesItem left, JobEmploymentTypesItem right) => left.Equals(right);
        public static bool operator !=(JobEmploymentTypesItem left, JobEmploymentTypesItem right) => !left.Equals(right);

        public static explicit operator string(JobEmploymentTypesItem value) => value._value;

        [EditorBrowsable(EditorBrowsableState.Never)]
        public override bool Equals(object? obj) => obj is JobEmploymentTypesItem other && Equals(other);
        public bool Equals(JobEmploymentTypesItem other) => string.Equals(_value, other._value, StringComparison.Ordinal);

        [EditorBrowsable(EditorBrowsableState.Never)]
        public override int GetHashCode() => _value?.GetHashCode() ?? 0;

        public override string ToString() => _value;
    }

    [EnumType]
    public readonly struct JobJobBenefitsItem : IEquatable<JobJobBenefitsItem>
    {
        private readonly string _value;

        private JobJobBenefitsItem(string value)
        {
            _value = value ?? throw new ArgumentNullException(nameof(value));
        }

        /// <summary>
        /// Default value if the type isn't specified.
        /// </summary>
        public static JobJobBenefitsItem JobBenefitUnspecified { get; } = new JobJobBenefitsItem("JOB_BENEFIT_UNSPECIFIED");
        /// <summary>
        /// The job includes access to programs that support child care, such as daycare.
        /// </summary>
        public static JobJobBenefitsItem ChildCare { get; } = new JobJobBenefitsItem("CHILD_CARE");
        /// <summary>
        /// The job includes dental services covered by a dental insurance plan.
        /// </summary>
        public static JobJobBenefitsItem Dental { get; } = new JobJobBenefitsItem("DENTAL");
        /// <summary>
        /// The job offers specific benefits to domestic partners.
        /// </summary>
        public static JobJobBenefitsItem DomesticPartner { get; } = new JobJobBenefitsItem("DOMESTIC_PARTNER");
        /// <summary>
        /// The job allows for a flexible work schedule.
        /// </summary>
        public static JobJobBenefitsItem FlexibleHours { get; } = new JobJobBenefitsItem("FLEXIBLE_HOURS");
        /// <summary>
        /// The job includes health services covered by a medical insurance plan.
        /// </summary>
        public static JobJobBenefitsItem Medical { get; } = new JobJobBenefitsItem("MEDICAL");
        /// <summary>
        /// The job includes a life insurance plan provided by the employer or available for purchase by the employee.
        /// </summary>
        public static JobJobBenefitsItem LifeInsurance { get; } = new JobJobBenefitsItem("LIFE_INSURANCE");
        /// <summary>
        /// The job allows for a leave of absence to a parent to care for a newborn child.
        /// </summary>
        public static JobJobBenefitsItem ParentalLeave { get; } = new JobJobBenefitsItem("PARENTAL_LEAVE");
        /// <summary>
        /// The job includes a workplace retirement plan provided by the employer or available for purchase by the employee.
        /// </summary>
        public static JobJobBenefitsItem RetirementPlan { get; } = new JobJobBenefitsItem("RETIREMENT_PLAN");
        /// <summary>
        /// The job allows for paid time off due to illness.
        /// </summary>
        public static JobJobBenefitsItem SickDays { get; } = new JobJobBenefitsItem("SICK_DAYS");
        /// <summary>
        /// The job includes paid time off for vacation.
        /// </summary>
        public static JobJobBenefitsItem Vacation { get; } = new JobJobBenefitsItem("VACATION");
        /// <summary>
        /// The job includes vision services covered by a vision insurance plan.
        /// </summary>
        public static JobJobBenefitsItem Vision { get; } = new JobJobBenefitsItem("VISION");

        public static bool operator ==(JobJobBenefitsItem left, JobJobBenefitsItem right) => left.Equals(right);
        public static bool operator !=(JobJobBenefitsItem left, JobJobBenefitsItem right) => !left.Equals(right);

        public static explicit operator string(JobJobBenefitsItem value) => value._value;

        [EditorBrowsable(EditorBrowsableState.Never)]
        public override bool Equals(object? obj) => obj is JobJobBenefitsItem other && Equals(other);
        public bool Equals(JobJobBenefitsItem other) => string.Equals(_value, other._value, StringComparison.Ordinal);

        [EditorBrowsable(EditorBrowsableState.Never)]
        public override int GetHashCode() => _value?.GetHashCode() ?? 0;

        public override string ToString() => _value;
    }

    /// <summary>
    /// The experience level associated with the job, such as "Entry Level".
    /// </summary>
    [EnumType]
    public readonly struct JobJobLevel : IEquatable<JobJobLevel>
    {
        private readonly string _value;

        private JobJobLevel(string value)
        {
            _value = value ?? throw new ArgumentNullException(nameof(value));
        }

        /// <summary>
        /// The default value if the level isn't specified.
        /// </summary>
        public static JobJobLevel JobLevelUnspecified { get; } = new JobJobLevel("JOB_LEVEL_UNSPECIFIED");
        /// <summary>
        /// Entry-level individual contributors, typically with less than 2 years of experience in a similar role. Includes interns.
        /// </summary>
        public static JobJobLevel EntryLevel { get; } = new JobJobLevel("ENTRY_LEVEL");
        /// <summary>
        /// Experienced individual contributors, typically with 2+ years of experience in a similar role.
        /// </summary>
        public static JobJobLevel Experienced { get; } = new JobJobLevel("EXPERIENCED");
        /// <summary>
        /// Entry- to mid-level managers responsible for managing a team of people.
        /// </summary>
        public static JobJobLevel Manager { get; } = new JobJobLevel("MANAGER");
        /// <summary>
        /// Senior-level managers responsible for managing teams of managers.
        /// </summary>
        public static JobJobLevel Director { get; } = new JobJobLevel("DIRECTOR");
        /// <summary>
        /// Executive-level managers and above, including C-level positions.
        /// </summary>
        public static JobJobLevel Executive { get; } = new JobJobLevel("EXECUTIVE");

        public static bool operator ==(JobJobLevel left, JobJobLevel right) => left.Equals(right);
        public static bool operator !=(JobJobLevel left, JobJobLevel right) => !left.Equals(right);

        public static explicit operator string(JobJobLevel value) => value._value;

        [EditorBrowsable(EditorBrowsableState.Never)]
        public override bool Equals(object? obj) => obj is JobJobLevel other && Equals(other);
        public bool Equals(JobJobLevel other) => string.Equals(_value, other._value, StringComparison.Ordinal);

        [EditorBrowsable(EditorBrowsableState.Never)]
        public override int GetHashCode() => _value?.GetHashCode() ?? 0;

        public override string ToString() => _value;
    }

    /// <summary>
    /// The job PostingRegion (for example, state, country) throughout which the job is available. If this field is set, a LocationFilter in a search query within the job region finds this job posting if an exact location match isn't specified. If this field is set to PostingRegion.NATION or PostingRegion.ADMINISTRATIVE_AREA, setting job Job.addresses to the same location level as this field is strongly recommended.
    /// </summary>
    [EnumType]
    public readonly struct JobPostingRegion : IEquatable<JobPostingRegion>
    {
        private readonly string _value;

        private JobPostingRegion(string value)
        {
            _value = value ?? throw new ArgumentNullException(nameof(value));
        }

        /// <summary>
        /// If the region is unspecified, the job is only returned if it matches the LocationFilter.
        /// </summary>
        public static JobPostingRegion PostingRegionUnspecified { get; } = new JobPostingRegion("POSTING_REGION_UNSPECIFIED");
        /// <summary>
        /// In addition to exact location matching, job posting is returned when the LocationFilter in the search query is in the same administrative area as the returned job posting. For example, if a `ADMINISTRATIVE_AREA` job is posted in "CA, USA", it's returned if LocationFilter has "Mountain View". Administrative area refers to top-level administrative subdivision of this country. For example, US state, IT region, UK constituent nation and JP prefecture.
        /// </summary>
        public static JobPostingRegion AdministrativeArea { get; } = new JobPostingRegion("ADMINISTRATIVE_AREA");
        /// <summary>
        /// In addition to exact location matching, job is returned when LocationFilter in search query is in the same country as this job. For example, if a `NATION_WIDE` job is posted in "USA", it's returned if LocationFilter has 'Mountain View'.
        /// </summary>
        public static JobPostingRegion Nation { get; } = new JobPostingRegion("NATION");
        /// <summary>
        /// Job allows employees to work remotely (telecommute). If locations are provided with this value, the job is considered as having a location, but telecommuting is allowed.
        /// </summary>
        public static JobPostingRegion Telecommute { get; } = new JobPostingRegion("TELECOMMUTE");

        public static bool operator ==(JobPostingRegion left, JobPostingRegion right) => left.Equals(right);
        public static bool operator !=(JobPostingRegion left, JobPostingRegion right) => !left.Equals(right);

        public static explicit operator string(JobPostingRegion value) => value._value;

        [EditorBrowsable(EditorBrowsableState.Never)]
        public override bool Equals(object? obj) => obj is JobPostingRegion other && Equals(other);
        public bool Equals(JobPostingRegion other) => string.Equals(_value, other._value, StringComparison.Ordinal);

        [EditorBrowsable(EditorBrowsableState.Never)]
        public override int GetHashCode() => _value?.GetHashCode() ?? 0;

        public override string ToString() => _value;
    }

    /// <summary>
    /// Deprecated. The job is only visible to the owner. The visibility of the job. Defaults to Visibility.ACCOUNT_ONLY if not specified.
    /// </summary>
    [EnumType]
    public readonly struct JobVisibility : IEquatable<JobVisibility>
    {
        private readonly string _value;

        private JobVisibility(string value)
        {
            _value = value ?? throw new ArgumentNullException(nameof(value));
        }

        /// <summary>
        /// Default value.
        /// </summary>
        public static JobVisibility VisibilityUnspecified { get; } = new JobVisibility("VISIBILITY_UNSPECIFIED");
        /// <summary>
        /// The resource is only visible to the GCP account who owns it.
        /// </summary>
        public static JobVisibility AccountOnly { get; } = new JobVisibility("ACCOUNT_ONLY");
        /// <summary>
        /// The resource is visible to the owner and may be visible to other applications and processes at Google.
        /// </summary>
        public static JobVisibility SharedWithGoogle { get; } = new JobVisibility("SHARED_WITH_GOOGLE");
        /// <summary>
        /// The resource is visible to the owner and may be visible to all other API clients.
        /// </summary>
        public static JobVisibility SharedWithPublic { get; } = new JobVisibility("SHARED_WITH_PUBLIC");

        public static bool operator ==(JobVisibility left, JobVisibility right) => left.Equals(right);
        public static bool operator !=(JobVisibility left, JobVisibility right) => !left.Equals(right);

        public static explicit operator string(JobVisibility value) => value._value;

        [EditorBrowsable(EditorBrowsableState.Never)]
        public override bool Equals(object? obj) => obj is JobVisibility other && Equals(other);
        public bool Equals(JobVisibility other) => string.Equals(_value, other._value, StringComparison.Ordinal);

        [EditorBrowsable(EditorBrowsableState.Never)]
        public override int GetHashCode() => _value?.GetHashCode() ?? 0;

        public override string ToString() => _value;
    }

    /// <summary>
    /// Option for job HTML content sanitization. Applied fields are: * description * applicationInfo.instruction * incentives * qualifications * responsibilities HTML tags in these fields may be stripped if sanitiazation isn't disabled. Defaults to HtmlSanitization.SIMPLE_FORMATTING_ONLY.
    /// </summary>
    [EnumType]
    public readonly struct ProcessingOptionsHtmlSanitization : IEquatable<ProcessingOptionsHtmlSanitization>
    {
        private readonly string _value;

        private ProcessingOptionsHtmlSanitization(string value)
        {
            _value = value ?? throw new ArgumentNullException(nameof(value));
        }

        /// <summary>
        /// Default value.
        /// </summary>
        public static ProcessingOptionsHtmlSanitization HtmlSanitizationUnspecified { get; } = new ProcessingOptionsHtmlSanitization("HTML_SANITIZATION_UNSPECIFIED");
        /// <summary>
        /// Disables sanitization on HTML input.
        /// </summary>
        public static ProcessingOptionsHtmlSanitization HtmlSanitizationDisabled { get; } = new ProcessingOptionsHtmlSanitization("HTML_SANITIZATION_DISABLED");
        /// <summary>
        /// Sanitizes HTML input, only accepts bold, italic, ordered list, and unordered list markup tags.
        /// </summary>
        public static ProcessingOptionsHtmlSanitization SimpleFormattingOnly { get; } = new ProcessingOptionsHtmlSanitization("SIMPLE_FORMATTING_ONLY");

        public static bool operator ==(ProcessingOptionsHtmlSanitization left, ProcessingOptionsHtmlSanitization right) => left.Equals(right);
        public static bool operator !=(ProcessingOptionsHtmlSanitization left, ProcessingOptionsHtmlSanitization right) => !left.Equals(right);

        public static explicit operator string(ProcessingOptionsHtmlSanitization value) => value._value;

        [EditorBrowsable(EditorBrowsableState.Never)]
        public override bool Equals(object? obj) => obj is ProcessingOptionsHtmlSanitization other && Equals(other);
        public bool Equals(ProcessingOptionsHtmlSanitization other) => string.Equals(_value, other._value, StringComparison.Ordinal);

        [EditorBrowsable(EditorBrowsableState.Never)]
        public override int GetHashCode() => _value?.GetHashCode() ?? 0;

        public override string ToString() => _value;
    }
}
