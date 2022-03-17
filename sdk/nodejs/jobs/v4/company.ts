// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import { input as inputs, output as outputs, enums } from "../../types";
import * as utilities from "../../utilities";

/**
 * Creates a new company entity.
 */
export class Company extends pulumi.CustomResource {
    /**
     * Get an existing Company resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, opts?: pulumi.CustomResourceOptions): Company {
        return new Company(name, undefined as any, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'google-native:jobs/v4:Company';

    /**
     * Returns true if the given object is an instance of Company.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is Company {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === Company.__pulumiType;
    }

    /**
     * The URI to employer's career site or careers page on the employer's web site, for example, "https://careers.google.com".
     */
    public readonly careerSiteUri!: pulumi.Output<string>;
    /**
     * Derived details about the company.
     */
    public /*out*/ readonly derivedInfo!: pulumi.Output<outputs.jobs.v4.CompanyDerivedInfoResponse>;
    /**
     * The display name of the company, for example, "Google LLC".
     */
    public readonly displayName!: pulumi.Output<string>;
    /**
     * Equal Employment Opportunity legal disclaimer text to be associated with all jobs, and typically to be displayed in all roles. The maximum number of allowed characters is 500.
     */
    public readonly eeoText!: pulumi.Output<string>;
    /**
     * Client side company identifier, used to uniquely identify the company. The maximum number of allowed characters is 255.
     */
    public readonly externalId!: pulumi.Output<string>;
    /**
     * The street address of the company's main headquarters, which may be different from the job location. The service attempts to geolocate the provided address, and populates a more specific location wherever possible in DerivedInfo.headquarters_location.
     */
    public readonly headquartersAddress!: pulumi.Output<string>;
    /**
     * Set to true if it is the hiring agency that post jobs for other employers. Defaults to false if not provided.
     */
    public readonly hiringAgency!: pulumi.Output<boolean>;
    /**
     * A URI that hosts the employer's company logo.
     */
    public readonly imageUri!: pulumi.Output<string>;
    /**
     * A list of keys of filterable Job.custom_attributes, whose corresponding `string_values` are used in keyword searches. Jobs with `string_values` under these specified field keys are returned if any of the values match the search keyword. Custom field values with parenthesis, brackets and special symbols are not searchable as-is, and those keyword queries must be surrounded by quotes.
     */
    public readonly keywordSearchableJobCustomAttributes!: pulumi.Output<string[]>;
    /**
     * Required during company update. The resource name for a company. This is generated by the service when a company is created. The format is "projects/{project_id}/tenants/{tenant_id}/companies/{company_id}", for example, "projects/foo/tenants/bar/companies/baz".
     */
    public readonly name!: pulumi.Output<string>;
    /**
     * The employer's company size.
     */
    public readonly size!: pulumi.Output<string>;
    /**
     * Indicates whether a company is flagged to be suspended from public availability by the service when job content appears suspicious, abusive, or spammy.
     */
    public /*out*/ readonly suspended!: pulumi.Output<boolean>;
    /**
     * The URI representing the company's primary web site or home page, for example, "https://www.google.com". The maximum number of allowed characters is 255.
     */
    public readonly websiteUri!: pulumi.Output<string>;

    /**
     * Create a Company resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: CompanyArgs, opts?: pulumi.CustomResourceOptions) {
        let resourceInputs: pulumi.Inputs = {};
        opts = opts || {};
        if (!opts.id) {
            if ((!args || args.displayName === undefined) && !opts.urn) {
                throw new Error("Missing required property 'displayName'");
            }
            if ((!args || args.externalId === undefined) && !opts.urn) {
                throw new Error("Missing required property 'externalId'");
            }
            if ((!args || args.tenantId === undefined) && !opts.urn) {
                throw new Error("Missing required property 'tenantId'");
            }
            resourceInputs["careerSiteUri"] = args ? args.careerSiteUri : undefined;
            resourceInputs["displayName"] = args ? args.displayName : undefined;
            resourceInputs["eeoText"] = args ? args.eeoText : undefined;
            resourceInputs["externalId"] = args ? args.externalId : undefined;
            resourceInputs["headquartersAddress"] = args ? args.headquartersAddress : undefined;
            resourceInputs["hiringAgency"] = args ? args.hiringAgency : undefined;
            resourceInputs["imageUri"] = args ? args.imageUri : undefined;
            resourceInputs["keywordSearchableJobCustomAttributes"] = args ? args.keywordSearchableJobCustomAttributes : undefined;
            resourceInputs["name"] = args ? args.name : undefined;
            resourceInputs["project"] = args ? args.project : undefined;
            resourceInputs["size"] = args ? args.size : undefined;
            resourceInputs["tenantId"] = args ? args.tenantId : undefined;
            resourceInputs["websiteUri"] = args ? args.websiteUri : undefined;
            resourceInputs["derivedInfo"] = undefined /*out*/;
            resourceInputs["suspended"] = undefined /*out*/;
        } else {
            resourceInputs["careerSiteUri"] = undefined /*out*/;
            resourceInputs["derivedInfo"] = undefined /*out*/;
            resourceInputs["displayName"] = undefined /*out*/;
            resourceInputs["eeoText"] = undefined /*out*/;
            resourceInputs["externalId"] = undefined /*out*/;
            resourceInputs["headquartersAddress"] = undefined /*out*/;
            resourceInputs["hiringAgency"] = undefined /*out*/;
            resourceInputs["imageUri"] = undefined /*out*/;
            resourceInputs["keywordSearchableJobCustomAttributes"] = undefined /*out*/;
            resourceInputs["name"] = undefined /*out*/;
            resourceInputs["size"] = undefined /*out*/;
            resourceInputs["suspended"] = undefined /*out*/;
            resourceInputs["websiteUri"] = undefined /*out*/;
        }
        opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
        super(Company.__pulumiType, name, resourceInputs, opts);
    }
}

/**
 * The set of arguments for constructing a Company resource.
 */
export interface CompanyArgs {
    /**
     * The URI to employer's career site or careers page on the employer's web site, for example, "https://careers.google.com".
     */
    careerSiteUri?: pulumi.Input<string>;
    /**
     * The display name of the company, for example, "Google LLC".
     */
    displayName: pulumi.Input<string>;
    /**
     * Equal Employment Opportunity legal disclaimer text to be associated with all jobs, and typically to be displayed in all roles. The maximum number of allowed characters is 500.
     */
    eeoText?: pulumi.Input<string>;
    /**
     * Client side company identifier, used to uniquely identify the company. The maximum number of allowed characters is 255.
     */
    externalId: pulumi.Input<string>;
    /**
     * The street address of the company's main headquarters, which may be different from the job location. The service attempts to geolocate the provided address, and populates a more specific location wherever possible in DerivedInfo.headquarters_location.
     */
    headquartersAddress?: pulumi.Input<string>;
    /**
     * Set to true if it is the hiring agency that post jobs for other employers. Defaults to false if not provided.
     */
    hiringAgency?: pulumi.Input<boolean>;
    /**
     * A URI that hosts the employer's company logo.
     */
    imageUri?: pulumi.Input<string>;
    /**
     * A list of keys of filterable Job.custom_attributes, whose corresponding `string_values` are used in keyword searches. Jobs with `string_values` under these specified field keys are returned if any of the values match the search keyword. Custom field values with parenthesis, brackets and special symbols are not searchable as-is, and those keyword queries must be surrounded by quotes.
     */
    keywordSearchableJobCustomAttributes?: pulumi.Input<pulumi.Input<string>[]>;
    /**
     * Required during company update. The resource name for a company. This is generated by the service when a company is created. The format is "projects/{project_id}/tenants/{tenant_id}/companies/{company_id}", for example, "projects/foo/tenants/bar/companies/baz".
     */
    name?: pulumi.Input<string>;
    project?: pulumi.Input<string>;
    /**
     * The employer's company size.
     */
    size?: pulumi.Input<enums.jobs.v4.CompanySize>;
    tenantId: pulumi.Input<string>;
    /**
     * The URI representing the company's primary web site or home page, for example, "https://www.google.com". The maximum number of allowed characters is 255.
     */
    websiteUri?: pulumi.Input<string>;
}
