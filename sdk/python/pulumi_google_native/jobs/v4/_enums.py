# coding=utf-8
# *** WARNING: this file was generated by the Pulumi SDK Generator. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

from enum import Enum

__all__ = [
    'CompanySize',
    'CompensationEntryType',
    'CompensationEntryUnit',
    'JobDegreeTypesItem',
    'JobEmploymentTypesItem',
    'JobJobBenefitsItem',
    'JobJobLevel',
    'JobPostingRegion',
    'JobVisibility',
    'ProcessingOptionsHtmlSanitization',
]


class CompanySize(str, Enum):
    """
    The employer's company size.
    """
    COMPANY_SIZE_UNSPECIFIED = "COMPANY_SIZE_UNSPECIFIED"
    """
    Default value if the size isn't specified.
    """
    MINI = "MINI"
    """
    The company has less than 50 employees.
    """
    SMALL = "SMALL"
    """
    The company has between 50 and 99 employees.
    """
    SMEDIUM = "SMEDIUM"
    """
    The company has between 100 and 499 employees.
    """
    MEDIUM = "MEDIUM"
    """
    The company has between 500 and 999 employees.
    """
    BIG = "BIG"
    """
    The company has between 1,000 and 4,999 employees.
    """
    BIGGER = "BIGGER"
    """
    The company has between 5,000 and 9,999 employees.
    """
    GIANT = "GIANT"
    """
    The company has 10,000 or more employees.
    """


class CompensationEntryType(str, Enum):
    """
    Compensation type. Default is CompensationType.COMPENSATION_TYPE_UNSPECIFIED.
    """
    COMPENSATION_TYPE_UNSPECIFIED = "COMPENSATION_TYPE_UNSPECIFIED"
    """
    Default value.
    """
    BASE = "BASE"
    """
    Base compensation: Refers to the fixed amount of money paid to an employee by an employer in return for work performed. Base compensation does not include benefits, bonuses or any other potential compensation from an employer.
    """
    BONUS = "BONUS"
    """
    Bonus.
    """
    SIGNING_BONUS = "SIGNING_BONUS"
    """
    Signing bonus.
    """
    EQUITY = "EQUITY"
    """
    Equity.
    """
    PROFIT_SHARING = "PROFIT_SHARING"
    """
    Profit sharing.
    """
    COMMISSIONS = "COMMISSIONS"
    """
    Commission.
    """
    TIPS = "TIPS"
    """
    Tips.
    """
    OTHER_COMPENSATION_TYPE = "OTHER_COMPENSATION_TYPE"
    """
    Other compensation type.
    """


class CompensationEntryUnit(str, Enum):
    """
    Frequency of the specified amount. Default is CompensationUnit.COMPENSATION_UNIT_UNSPECIFIED.
    """
    COMPENSATION_UNIT_UNSPECIFIED = "COMPENSATION_UNIT_UNSPECIFIED"
    """
    Default value.
    """
    HOURLY = "HOURLY"
    """
    Hourly.
    """
    DAILY = "DAILY"
    """
    Daily.
    """
    WEEKLY = "WEEKLY"
    """
    Weekly
    """
    MONTHLY = "MONTHLY"
    """
    Monthly.
    """
    YEARLY = "YEARLY"
    """
    Yearly.
    """
    ONE_TIME = "ONE_TIME"
    """
    One time.
    """
    OTHER_COMPENSATION_UNIT = "OTHER_COMPENSATION_UNIT"
    """
    Other compensation units.
    """


class JobDegreeTypesItem(str, Enum):
    DEGREE_TYPE_UNSPECIFIED = "DEGREE_TYPE_UNSPECIFIED"
    """
    Default value. Represents no degree, or early childhood education. Maps to ISCED code 0. Ex) Kindergarten
    """
    PRIMARY_EDUCATION = "PRIMARY_EDUCATION"
    """
    Primary education which is typically the first stage of compulsory education. ISCED code 1. Ex) Elementary school
    """
    LOWER_SECONDARY_EDUCATION = "LOWER_SECONDARY_EDUCATION"
    """
    Lower secondary education; First stage of secondary education building on primary education, typically with a more subject-oriented curriculum. ISCED code 2. Ex) Middle school
    """
    UPPER_SECONDARY_EDUCATION = "UPPER_SECONDARY_EDUCATION"
    """
    Middle education; Second/final stage of secondary education preparing for tertiary education and/or providing skills relevant to employment. Usually with an increased range of subject options and streams. ISCED code 3. Ex) High school
    """
    ADULT_REMEDIAL_EDUCATION = "ADULT_REMEDIAL_EDUCATION"
    """
    Adult Remedial Education; Programmes providing learning experiences that build on secondary education and prepare for labour market entry and/or tertiary education. The content is broader than secondary but not as complex as tertiary education. ISCED code 4.
    """
    ASSOCIATES_OR_EQUIVALENT = "ASSOCIATES_OR_EQUIVALENT"
    """
    Associate's or equivalent; Short first tertiary programmes that are typically practically-based, occupationally-specific and prepare for labour market entry. These programmes may also provide a pathway to other tertiary programmes. ISCED code 5.
    """
    BACHELORS_OR_EQUIVALENT = "BACHELORS_OR_EQUIVALENT"
    """
    Bachelor's or equivalent; Programmes designed to provide intermediate academic and/or professional knowledge, skills and competencies leading to a first tertiary degree or equivalent qualification. ISCED code 6.
    """
    MASTERS_OR_EQUIVALENT = "MASTERS_OR_EQUIVALENT"
    """
    Master's or equivalent; Programmes designed to provide advanced academic and/or professional knowledge, skills and competencies leading to a second tertiary degree or equivalent qualification. ISCED code 7.
    """
    DOCTORAL_OR_EQUIVALENT = "DOCTORAL_OR_EQUIVALENT"
    """
    Doctoral or equivalent; Programmes designed primarily to lead to an advanced research qualification, usually concluding with the submission and defense of a substantive dissertation of publishable quality based on original research. ISCED code 8.
    """


class JobEmploymentTypesItem(str, Enum):
    EMPLOYMENT_TYPE_UNSPECIFIED = "EMPLOYMENT_TYPE_UNSPECIFIED"
    """
    The default value if the employment type isn't specified.
    """
    FULL_TIME = "FULL_TIME"
    """
    The job requires working a number of hours that constitute full time employment, typically 40 or more hours per week.
    """
    PART_TIME = "PART_TIME"
    """
    The job entails working fewer hours than a full time job, typically less than 40 hours a week.
    """
    CONTRACTOR = "CONTRACTOR"
    """
    The job is offered as a contracted, as opposed to a salaried employee, position.
    """
    CONTRACT_TO_HIRE = "CONTRACT_TO_HIRE"
    """
    The job is offered as a contracted position with the understanding that it's converted into a full-time position at the end of the contract. Jobs of this type are also returned by a search for EmploymentType.CONTRACTOR jobs.
    """
    TEMPORARY = "TEMPORARY"
    """
    The job is offered as a temporary employment opportunity, usually a short-term engagement.
    """
    INTERN = "INTERN"
    """
    The job is a fixed-term opportunity for students or entry-level job seekers to obtain on-the-job training, typically offered as a summer position.
    """
    VOLUNTEER = "VOLUNTEER"
    """
    The is an opportunity for an individual to volunteer, where there's no expectation of compensation for the provided services.
    """
    PER_DIEM = "PER_DIEM"
    """
    The job requires an employee to work on an as-needed basis with a flexible schedule.
    """
    FLY_IN_FLY_OUT = "FLY_IN_FLY_OUT"
    """
    The job involves employing people in remote areas and flying them temporarily to the work site instead of relocating employees and their families permanently.
    """
    OTHER_EMPLOYMENT_TYPE = "OTHER_EMPLOYMENT_TYPE"
    """
    The job does not fit any of the other listed types.
    """


class JobJobBenefitsItem(str, Enum):
    JOB_BENEFIT_UNSPECIFIED = "JOB_BENEFIT_UNSPECIFIED"
    """
    Default value if the type isn't specified.
    """
    CHILD_CARE = "CHILD_CARE"
    """
    The job includes access to programs that support child care, such as daycare.
    """
    DENTAL = "DENTAL"
    """
    The job includes dental services covered by a dental insurance plan.
    """
    DOMESTIC_PARTNER = "DOMESTIC_PARTNER"
    """
    The job offers specific benefits to domestic partners.
    """
    FLEXIBLE_HOURS = "FLEXIBLE_HOURS"
    """
    The job allows for a flexible work schedule.
    """
    MEDICAL = "MEDICAL"
    """
    The job includes health services covered by a medical insurance plan.
    """
    LIFE_INSURANCE = "LIFE_INSURANCE"
    """
    The job includes a life insurance plan provided by the employer or available for purchase by the employee.
    """
    PARENTAL_LEAVE = "PARENTAL_LEAVE"
    """
    The job allows for a leave of absence to a parent to care for a newborn child.
    """
    RETIREMENT_PLAN = "RETIREMENT_PLAN"
    """
    The job includes a workplace retirement plan provided by the employer or available for purchase by the employee.
    """
    SICK_DAYS = "SICK_DAYS"
    """
    The job allows for paid time off due to illness.
    """
    VACATION = "VACATION"
    """
    The job includes paid time off for vacation.
    """
    VISION = "VISION"
    """
    The job includes vision services covered by a vision insurance plan.
    """


class JobJobLevel(str, Enum):
    """
    The experience level associated with the job, such as "Entry Level".
    """
    JOB_LEVEL_UNSPECIFIED = "JOB_LEVEL_UNSPECIFIED"
    """
    The default value if the level isn't specified.
    """
    ENTRY_LEVEL = "ENTRY_LEVEL"
    """
    Entry-level individual contributors, typically with less than 2 years of experience in a similar role. Includes interns.
    """
    EXPERIENCED = "EXPERIENCED"
    """
    Experienced individual contributors, typically with 2+ years of experience in a similar role.
    """
    MANAGER = "MANAGER"
    """
    Entry- to mid-level managers responsible for managing a team of people.
    """
    DIRECTOR = "DIRECTOR"
    """
    Senior-level managers responsible for managing teams of managers.
    """
    EXECUTIVE = "EXECUTIVE"
    """
    Executive-level managers and above, including C-level positions.
    """


class JobPostingRegion(str, Enum):
    """
    The job PostingRegion (for example, state, country) throughout which the job is available. If this field is set, a LocationFilter in a search query within the job region finds this job posting if an exact location match isn't specified. If this field is set to PostingRegion.NATION or PostingRegion.ADMINISTRATIVE_AREA, setting job Job.addresses to the same location level as this field is strongly recommended.
    """
    POSTING_REGION_UNSPECIFIED = "POSTING_REGION_UNSPECIFIED"
    """
    If the region is unspecified, the job is only returned if it matches the LocationFilter.
    """
    ADMINISTRATIVE_AREA = "ADMINISTRATIVE_AREA"
    """
    In addition to exact location matching, job posting is returned when the LocationFilter in the search query is in the same administrative area as the returned job posting. For example, if a `ADMINISTRATIVE_AREA` job is posted in "CA, USA", it's returned if LocationFilter has "Mountain View". Administrative area refers to top-level administrative subdivision of this country. For example, US state, IT region, UK constituent nation and JP prefecture.
    """
    NATION = "NATION"
    """
    In addition to exact location matching, job is returned when LocationFilter in search query is in the same country as this job. For example, if a `NATION_WIDE` job is posted in "USA", it's returned if LocationFilter has 'Mountain View'.
    """
    TELECOMMUTE = "TELECOMMUTE"
    """
    Job allows employees to work remotely (telecommute). If locations are provided with this value, the job is considered as having a location, but telecommuting is allowed.
    """


class JobVisibility(str, Enum):
    """
    Deprecated. The job is only visible to the owner. The visibility of the job. Defaults to Visibility.ACCOUNT_ONLY if not specified.
    """
    VISIBILITY_UNSPECIFIED = "VISIBILITY_UNSPECIFIED"
    """
    Default value.
    """
    ACCOUNT_ONLY = "ACCOUNT_ONLY"
    """
    The resource is only visible to the GCP account who owns it.
    """
    SHARED_WITH_GOOGLE = "SHARED_WITH_GOOGLE"
    """
    The resource is visible to the owner and may be visible to other applications and processes at Google.
    """
    SHARED_WITH_PUBLIC = "SHARED_WITH_PUBLIC"
    """
    The resource is visible to the owner and may be visible to all other API clients.
    """


class ProcessingOptionsHtmlSanitization(str, Enum):
    """
    Option for job HTML content sanitization. Applied fields are: * description * applicationInfo.instruction * incentives * qualifications * responsibilities HTML tags in these fields may be stripped if sanitiazation isn't disabled. Defaults to HtmlSanitization.SIMPLE_FORMATTING_ONLY.
    """
    HTML_SANITIZATION_UNSPECIFIED = "HTML_SANITIZATION_UNSPECIFIED"
    """
    Default value.
    """
    HTML_SANITIZATION_DISABLED = "HTML_SANITIZATION_DISABLED"
    """
    Disables sanitization on HTML input.
    """
    SIMPLE_FORMATTING_ONLY = "SIMPLE_FORMATTING_ONLY"
    """
    Sanitizes HTML input, only accepts bold, italic, ordered list, and unordered list markup tags.
    """
