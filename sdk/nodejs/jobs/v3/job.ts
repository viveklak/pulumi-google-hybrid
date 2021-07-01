// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import { input as inputs, output as outputs, enums } from "../../types";
import * as utilities from "../../utilities";

/**
 * Creates a new job. Typically, the job becomes searchable within 10 seconds, but it may take up to 5 minutes.
 */
export class Job extends pulumi.CustomResource {
    /**
     * Get an existing Job resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, opts?: pulumi.CustomResourceOptions): Job {
        return new Job(name, undefined as any, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'google-native:jobs/v3:Job';

    /**
     * Returns true if the given object is an instance of Job.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is Job {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === Job.__pulumiType;
    }

    /**
     * Optional but strongly recommended for the best service experience. Location(s) where the employer is looking to hire for this job posting. Specifying the full street address(es) of the hiring location enables better API results, especially job searches by commute time. At most 50 locations are allowed for best search performance. If a job has more locations, it is suggested to split it into multiple jobs with unique requisition_ids (e.g. 'ReqA' becomes 'ReqA-1', 'ReqA-2', etc.) as multiple jobs with the same company_name, language_code and requisition_id are not allowed. If the original requisition_id must be preserved, a custom field should be used for storage. It is also suggested to group the locations that close to each other in the same job for better search experience. Jobs with multiple addresses must have their addresses with the same LocationType to allow location filtering to work properly. (For example, a Job with addresses "1600 Amphitheatre Parkway, Mountain View, CA, USA" and "London, UK" may not have location filters applied correctly at search time since the first is a LocationType.STREET_ADDRESS and the second is a LocationType.LOCALITY.) If a job needs to have multiple addresses, it is suggested to split it into multiple jobs with same LocationTypes. The maximum number of allowed characters is 500.
     */
    public readonly addresses!: pulumi.Output<string[]>;
    /**
     * Required. At least one field within ApplicationInfo must be specified. Job application information.
     */
    public readonly applicationInfo!: pulumi.Output<outputs.jobs.v3.ApplicationInfoResponse>;
    /**
     * Display name of the company listing the job.
     */
    public /*out*/ readonly companyDisplayName!: pulumi.Output<string>;
    /**
     * Required. The resource name of the company listing the job, such as "projects/api-test-project/companies/foo".
     */
    public readonly companyName!: pulumi.Output<string>;
    /**
     * Optional. Job compensation information.
     */
    public readonly compensationInfo!: pulumi.Output<outputs.jobs.v3.CompensationInfoResponse>;
    /**
     * Optional. A map of fields to hold both filterable and non-filterable custom job attributes that are not covered by the provided structured fields. The keys of the map are strings up to 64 bytes and must match the pattern: a-zA-Z*. For example, key0LikeThis or KEY_1_LIKE_THIS. At most 100 filterable and at most 100 unfilterable keys are supported. For filterable `string_values`, across all keys at most 200 values are allowed, with each string no more than 255 characters. For unfilterable `string_values`, the maximum total size of `string_values` across all keys is 50KB.
     */
    public readonly customAttributes!: pulumi.Output<{[key: string]: string}>;
    /**
     * Optional. The desired education degrees for the job, such as Bachelors, Masters.
     */
    public readonly degreeTypes!: pulumi.Output<string[]>;
    /**
     * Optional. The department or functional area within the company with the open position. The maximum number of allowed characters is 255.
     */
    public readonly department!: pulumi.Output<string>;
    /**
     * Derived details about the job posting.
     */
    public /*out*/ readonly derivedInfo!: pulumi.Output<outputs.jobs.v3.JobDerivedInfoResponse>;
    /**
     * Required. The description of the job, which typically includes a multi-paragraph description of the company and related information. Separate fields are provided on the job object for responsibilities, qualifications, and other job characteristics. Use of these separate job fields is recommended. This field accepts and sanitizes HTML input, and also accepts bold, italic, ordered list, and unordered list markup tags. The maximum number of allowed characters is 100,000.
     */
    public readonly description!: pulumi.Output<string>;
    /**
     * Optional. The employment type(s) of a job, for example, full time or part time.
     */
    public readonly employmentTypes!: pulumi.Output<string[]>;
    /**
     * Optional. A description of bonus, commission, and other compensation incentives associated with the job not including salary or pay. The maximum number of allowed characters is 10,000.
     */
    public readonly incentives!: pulumi.Output<string>;
    /**
     * Optional. The benefits included with the job.
     */
    public readonly jobBenefits!: pulumi.Output<string[]>;
    /**
     * Optional. The end timestamp of the job. Typically this field is used for contracting engagements. Invalid timestamps are ignored.
     */
    public readonly jobEndTime!: pulumi.Output<string>;
    /**
     * Optional. The experience level associated with the job, such as "Entry Level".
     */
    public readonly jobLevel!: pulumi.Output<string>;
    /**
     * Optional. The start timestamp of the job in UTC time zone. Typically this field is used for contracting engagements. Invalid timestamps are ignored.
     */
    public readonly jobStartTime!: pulumi.Output<string>;
    /**
     * Optional. The language of the posting. This field is distinct from any requirements for fluency that are associated with the job. Language codes must be in BCP-47 format, such as "en-US" or "sr-Latn". For more information, see [Tags for Identifying Languages](https://tools.ietf.org/html/bcp47){: class="external" target="_blank" }. If this field is unspecified and Job.description is present, detected language code based on Job.description is assigned, otherwise defaults to 'en_US'.
     */
    public readonly languageCode!: pulumi.Output<string>;
    /**
     * Required during job update. The resource name for the job. This is generated by the service when a job is created. The format is "projects/{project_id}/jobs/{job_id}", for example, "projects/api-test-project/jobs/1234". Use of this field in job queries and API calls is preferred over the use of requisition_id since this value is unique.
     */
    public readonly name!: pulumi.Output<string>;
    /**
     * The timestamp when this job posting was created.
     */
    public /*out*/ readonly postingCreateTime!: pulumi.Output<string>;
    /**
     * Optional but strongly recommended for the best service experience. The expiration timestamp of the job. After this timestamp, the job is marked as expired, and it no longer appears in search results. The expired job can't be deleted or listed by the DeleteJob and ListJobs APIs, but it can be retrieved with the GetJob API or updated with the UpdateJob API. An expired job can be updated and opened again by using a future expiration timestamp. Updating an expired job fails if there is another existing open job with same company_name, language_code and requisition_id. The expired jobs are retained in our system for 90 days. However, the overall expired job count cannot exceed 3 times the maximum of open jobs count over the past week, otherwise jobs with earlier expire time are cleaned first. Expired jobs are no longer accessible after they are cleaned out. Invalid timestamps are ignored, and treated as expire time not provided. Timestamp before the instant request is made is considered valid, the job will be treated as expired immediately. If this value is not provided at the time of job creation or is invalid, the job posting expires after 30 days from the job's creation time. For example, if the job was created on 2017/01/01 13:00AM UTC with an unspecified expiration date, the job expires after 2017/01/31 13:00AM UTC. If this value is not provided on job update, it depends on the field masks set by UpdateJobRequest.update_mask. If the field masks include expiry_time, or the masks are empty meaning that every field is updated, the job posting expires after 30 days from the job's last update time. Otherwise the expiration date isn't updated.
     */
    public readonly postingExpireTime!: pulumi.Output<string>;
    /**
     * Optional. The timestamp this job posting was most recently published. The default value is the time the request arrives at the server. Invalid timestamps are ignored.
     */
    public readonly postingPublishTime!: pulumi.Output<string>;
    /**
     * Optional. The job PostingRegion (for example, state, country) throughout which the job is available. If this field is set, a LocationFilter in a search query within the job region finds this job posting if an exact location match isn't specified. If this field is set to PostingRegion.NATION or PostingRegion.ADMINISTRATIVE_AREA, setting job Job.addresses to the same location level as this field is strongly recommended.
     */
    public readonly postingRegion!: pulumi.Output<string>;
    /**
     * The timestamp when this job posting was last updated.
     */
    public /*out*/ readonly postingUpdateTime!: pulumi.Output<string>;
    /**
     * Optional. Options for job processing.
     */
    public readonly processingOptions!: pulumi.Output<outputs.jobs.v3.ProcessingOptionsResponse>;
    /**
     * Optional. A promotion value of the job, as determined by the client. The value determines the sort order of the jobs returned when searching for jobs using the featured jobs search call, with higher promotional values being returned first and ties being resolved by relevance sort. Only the jobs with a promotionValue >0 are returned in a FEATURED_JOB_SEARCH. Default value is 0, and negative values are treated as 0.
     */
    public readonly promotionValue!: pulumi.Output<number>;
    /**
     * Optional. A description of the qualifications required to perform the job. The use of this field is recommended as an alternative to using the more general description field. This field accepts and sanitizes HTML input, and also accepts bold, italic, ordered list, and unordered list markup tags. The maximum number of allowed characters is 10,000.
     */
    public readonly qualifications!: pulumi.Output<string>;
    /**
     * Required. The requisition ID, also referred to as the posting ID, assigned by the client to identify a job. This field is intended to be used by clients for client identification and tracking of postings. A job is not allowed to be created if there is another job with the same [company_name], language_code and requisition_id. The maximum number of allowed characters is 255.
     */
    public readonly requisitionId!: pulumi.Output<string>;
    /**
     * Optional. A description of job responsibilities. The use of this field is recommended as an alternative to using the more general description field. This field accepts and sanitizes HTML input, and also accepts bold, italic, ordered list, and unordered list markup tags. The maximum number of allowed characters is 10,000.
     */
    public readonly responsibilities!: pulumi.Output<string>;
    /**
     * Required. The title of the job, such as "Software Engineer" The maximum number of allowed characters is 500.
     */
    public readonly title!: pulumi.Output<string>;

    /**
     * Create a Job resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: JobArgs, opts?: pulumi.CustomResourceOptions) {
        let inputs: pulumi.Inputs = {};
        opts = opts || {};
        if (!opts.id) {
            if ((!args || args.project === undefined) && !opts.urn) {
                throw new Error("Missing required property 'project'");
            }
            inputs["addresses"] = args ? args.addresses : undefined;
            inputs["applicationInfo"] = args ? args.applicationInfo : undefined;
            inputs["companyName"] = args ? args.companyName : undefined;
            inputs["compensationInfo"] = args ? args.compensationInfo : undefined;
            inputs["customAttributes"] = args ? args.customAttributes : undefined;
            inputs["degreeTypes"] = args ? args.degreeTypes : undefined;
            inputs["department"] = args ? args.department : undefined;
            inputs["description"] = args ? args.description : undefined;
            inputs["employmentTypes"] = args ? args.employmentTypes : undefined;
            inputs["incentives"] = args ? args.incentives : undefined;
            inputs["jobBenefits"] = args ? args.jobBenefits : undefined;
            inputs["jobEndTime"] = args ? args.jobEndTime : undefined;
            inputs["jobLevel"] = args ? args.jobLevel : undefined;
            inputs["jobStartTime"] = args ? args.jobStartTime : undefined;
            inputs["languageCode"] = args ? args.languageCode : undefined;
            inputs["name"] = args ? args.name : undefined;
            inputs["postingExpireTime"] = args ? args.postingExpireTime : undefined;
            inputs["postingPublishTime"] = args ? args.postingPublishTime : undefined;
            inputs["postingRegion"] = args ? args.postingRegion : undefined;
            inputs["processingOptions"] = args ? args.processingOptions : undefined;
            inputs["project"] = args ? args.project : undefined;
            inputs["promotionValue"] = args ? args.promotionValue : undefined;
            inputs["qualifications"] = args ? args.qualifications : undefined;
            inputs["requisitionId"] = args ? args.requisitionId : undefined;
            inputs["responsibilities"] = args ? args.responsibilities : undefined;
            inputs["title"] = args ? args.title : undefined;
            inputs["companyDisplayName"] = undefined /*out*/;
            inputs["derivedInfo"] = undefined /*out*/;
            inputs["postingCreateTime"] = undefined /*out*/;
            inputs["postingUpdateTime"] = undefined /*out*/;
        } else {
            inputs["addresses"] = undefined /*out*/;
            inputs["applicationInfo"] = undefined /*out*/;
            inputs["companyDisplayName"] = undefined /*out*/;
            inputs["companyName"] = undefined /*out*/;
            inputs["compensationInfo"] = undefined /*out*/;
            inputs["customAttributes"] = undefined /*out*/;
            inputs["degreeTypes"] = undefined /*out*/;
            inputs["department"] = undefined /*out*/;
            inputs["derivedInfo"] = undefined /*out*/;
            inputs["description"] = undefined /*out*/;
            inputs["employmentTypes"] = undefined /*out*/;
            inputs["incentives"] = undefined /*out*/;
            inputs["jobBenefits"] = undefined /*out*/;
            inputs["jobEndTime"] = undefined /*out*/;
            inputs["jobLevel"] = undefined /*out*/;
            inputs["jobStartTime"] = undefined /*out*/;
            inputs["languageCode"] = undefined /*out*/;
            inputs["name"] = undefined /*out*/;
            inputs["postingCreateTime"] = undefined /*out*/;
            inputs["postingExpireTime"] = undefined /*out*/;
            inputs["postingPublishTime"] = undefined /*out*/;
            inputs["postingRegion"] = undefined /*out*/;
            inputs["postingUpdateTime"] = undefined /*out*/;
            inputs["processingOptions"] = undefined /*out*/;
            inputs["promotionValue"] = undefined /*out*/;
            inputs["qualifications"] = undefined /*out*/;
            inputs["requisitionId"] = undefined /*out*/;
            inputs["responsibilities"] = undefined /*out*/;
            inputs["title"] = undefined /*out*/;
        }
        if (!opts.version) {
            opts = pulumi.mergeOptions(opts, { version: utilities.getVersion()});
        }
        super(Job.__pulumiType, name, inputs, opts);
    }
}

/**
 * The set of arguments for constructing a Job resource.
 */
export interface JobArgs {
    /**
     * Optional but strongly recommended for the best service experience. Location(s) where the employer is looking to hire for this job posting. Specifying the full street address(es) of the hiring location enables better API results, especially job searches by commute time. At most 50 locations are allowed for best search performance. If a job has more locations, it is suggested to split it into multiple jobs with unique requisition_ids (e.g. 'ReqA' becomes 'ReqA-1', 'ReqA-2', etc.) as multiple jobs with the same company_name, language_code and requisition_id are not allowed. If the original requisition_id must be preserved, a custom field should be used for storage. It is also suggested to group the locations that close to each other in the same job for better search experience. Jobs with multiple addresses must have their addresses with the same LocationType to allow location filtering to work properly. (For example, a Job with addresses "1600 Amphitheatre Parkway, Mountain View, CA, USA" and "London, UK" may not have location filters applied correctly at search time since the first is a LocationType.STREET_ADDRESS and the second is a LocationType.LOCALITY.) If a job needs to have multiple addresses, it is suggested to split it into multiple jobs with same LocationTypes. The maximum number of allowed characters is 500.
     */
    addresses?: pulumi.Input<pulumi.Input<string>[]>;
    /**
     * Required. At least one field within ApplicationInfo must be specified. Job application information.
     */
    applicationInfo?: pulumi.Input<inputs.jobs.v3.ApplicationInfoArgs>;
    /**
     * Required. The resource name of the company listing the job, such as "projects/api-test-project/companies/foo".
     */
    companyName?: pulumi.Input<string>;
    /**
     * Optional. Job compensation information.
     */
    compensationInfo?: pulumi.Input<inputs.jobs.v3.CompensationInfoArgs>;
    /**
     * Optional. A map of fields to hold both filterable and non-filterable custom job attributes that are not covered by the provided structured fields. The keys of the map are strings up to 64 bytes and must match the pattern: a-zA-Z*. For example, key0LikeThis or KEY_1_LIKE_THIS. At most 100 filterable and at most 100 unfilterable keys are supported. For filterable `string_values`, across all keys at most 200 values are allowed, with each string no more than 255 characters. For unfilterable `string_values`, the maximum total size of `string_values` across all keys is 50KB.
     */
    customAttributes?: pulumi.Input<{[key: string]: string}>;
    /**
     * Optional. The desired education degrees for the job, such as Bachelors, Masters.
     */
    degreeTypes?: pulumi.Input<pulumi.Input<enums.jobs.v3.JobDegreeTypesItem>[]>;
    /**
     * Optional. The department or functional area within the company with the open position. The maximum number of allowed characters is 255.
     */
    department?: pulumi.Input<string>;
    /**
     * Required. The description of the job, which typically includes a multi-paragraph description of the company and related information. Separate fields are provided on the job object for responsibilities, qualifications, and other job characteristics. Use of these separate job fields is recommended. This field accepts and sanitizes HTML input, and also accepts bold, italic, ordered list, and unordered list markup tags. The maximum number of allowed characters is 100,000.
     */
    description?: pulumi.Input<string>;
    /**
     * Optional. The employment type(s) of a job, for example, full time or part time.
     */
    employmentTypes?: pulumi.Input<pulumi.Input<enums.jobs.v3.JobEmploymentTypesItem>[]>;
    /**
     * Optional. A description of bonus, commission, and other compensation incentives associated with the job not including salary or pay. The maximum number of allowed characters is 10,000.
     */
    incentives?: pulumi.Input<string>;
    /**
     * Optional. The benefits included with the job.
     */
    jobBenefits?: pulumi.Input<pulumi.Input<enums.jobs.v3.JobJobBenefitsItem>[]>;
    /**
     * Optional. The end timestamp of the job. Typically this field is used for contracting engagements. Invalid timestamps are ignored.
     */
    jobEndTime?: pulumi.Input<string>;
    /**
     * Optional. The experience level associated with the job, such as "Entry Level".
     */
    jobLevel?: pulumi.Input<enums.jobs.v3.JobJobLevel>;
    /**
     * Optional. The start timestamp of the job in UTC time zone. Typically this field is used for contracting engagements. Invalid timestamps are ignored.
     */
    jobStartTime?: pulumi.Input<string>;
    /**
     * Optional. The language of the posting. This field is distinct from any requirements for fluency that are associated with the job. Language codes must be in BCP-47 format, such as "en-US" or "sr-Latn". For more information, see [Tags for Identifying Languages](https://tools.ietf.org/html/bcp47){: class="external" target="_blank" }. If this field is unspecified and Job.description is present, detected language code based on Job.description is assigned, otherwise defaults to 'en_US'.
     */
    languageCode?: pulumi.Input<string>;
    /**
     * Required during job update. The resource name for the job. This is generated by the service when a job is created. The format is "projects/{project_id}/jobs/{job_id}", for example, "projects/api-test-project/jobs/1234". Use of this field in job queries and API calls is preferred over the use of requisition_id since this value is unique.
     */
    name?: pulumi.Input<string>;
    /**
     * Optional but strongly recommended for the best service experience. The expiration timestamp of the job. After this timestamp, the job is marked as expired, and it no longer appears in search results. The expired job can't be deleted or listed by the DeleteJob and ListJobs APIs, but it can be retrieved with the GetJob API or updated with the UpdateJob API. An expired job can be updated and opened again by using a future expiration timestamp. Updating an expired job fails if there is another existing open job with same company_name, language_code and requisition_id. The expired jobs are retained in our system for 90 days. However, the overall expired job count cannot exceed 3 times the maximum of open jobs count over the past week, otherwise jobs with earlier expire time are cleaned first. Expired jobs are no longer accessible after they are cleaned out. Invalid timestamps are ignored, and treated as expire time not provided. Timestamp before the instant request is made is considered valid, the job will be treated as expired immediately. If this value is not provided at the time of job creation or is invalid, the job posting expires after 30 days from the job's creation time. For example, if the job was created on 2017/01/01 13:00AM UTC with an unspecified expiration date, the job expires after 2017/01/31 13:00AM UTC. If this value is not provided on job update, it depends on the field masks set by UpdateJobRequest.update_mask. If the field masks include expiry_time, or the masks are empty meaning that every field is updated, the job posting expires after 30 days from the job's last update time. Otherwise the expiration date isn't updated.
     */
    postingExpireTime?: pulumi.Input<string>;
    /**
     * Optional. The timestamp this job posting was most recently published. The default value is the time the request arrives at the server. Invalid timestamps are ignored.
     */
    postingPublishTime?: pulumi.Input<string>;
    /**
     * Optional. The job PostingRegion (for example, state, country) throughout which the job is available. If this field is set, a LocationFilter in a search query within the job region finds this job posting if an exact location match isn't specified. If this field is set to PostingRegion.NATION or PostingRegion.ADMINISTRATIVE_AREA, setting job Job.addresses to the same location level as this field is strongly recommended.
     */
    postingRegion?: pulumi.Input<enums.jobs.v3.JobPostingRegion>;
    /**
     * Optional. Options for job processing.
     */
    processingOptions?: pulumi.Input<inputs.jobs.v3.ProcessingOptionsArgs>;
    project: pulumi.Input<string>;
    /**
     * Optional. A promotion value of the job, as determined by the client. The value determines the sort order of the jobs returned when searching for jobs using the featured jobs search call, with higher promotional values being returned first and ties being resolved by relevance sort. Only the jobs with a promotionValue >0 are returned in a FEATURED_JOB_SEARCH. Default value is 0, and negative values are treated as 0.
     */
    promotionValue?: pulumi.Input<number>;
    /**
     * Optional. A description of the qualifications required to perform the job. The use of this field is recommended as an alternative to using the more general description field. This field accepts and sanitizes HTML input, and also accepts bold, italic, ordered list, and unordered list markup tags. The maximum number of allowed characters is 10,000.
     */
    qualifications?: pulumi.Input<string>;
    /**
     * Required. The requisition ID, also referred to as the posting ID, assigned by the client to identify a job. This field is intended to be used by clients for client identification and tracking of postings. A job is not allowed to be created if there is another job with the same [company_name], language_code and requisition_id. The maximum number of allowed characters is 255.
     */
    requisitionId?: pulumi.Input<string>;
    /**
     * Optional. A description of job responsibilities. The use of this field is recommended as an alternative to using the more general description field. This field accepts and sanitizes HTML input, and also accepts bold, italic, ordered list, and unordered list markup tags. The maximum number of allowed characters is 10,000.
     */
    responsibilities?: pulumi.Input<string>;
    /**
     * Required. The title of the job, such as "Software Engineer" The maximum number of allowed characters is 500.
     */
    title?: pulumi.Input<string>;
}
