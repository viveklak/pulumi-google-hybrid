// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import { input as inputs, output as outputs, enums } from "../../types";
import * as utilities from "../../utilities";

/**
 * Creates a new ScanConfig.
 */
export class ScanConfig extends pulumi.CustomResource {
    /**
     * Get an existing ScanConfig resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, opts?: pulumi.CustomResourceOptions): ScanConfig {
        return new ScanConfig(name, undefined as any, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'google-native:websecurityscanner/v1:ScanConfig';

    /**
     * Returns true if the given object is an instance of ScanConfig.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is ScanConfig {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === ScanConfig.__pulumiType;
    }

    /**
     * The authentication configuration. If specified, service will use the authentication configuration during scanning.
     */
    public readonly authentication!: pulumi.Output<outputs.websecurityscanner.v1.AuthenticationResponse>;
    /**
     * The excluded URL patterns as described in https://cloud.google.com/security-command-center/docs/how-to-use-web-security-scanner#excluding_urls
     */
    public readonly blacklistPatterns!: pulumi.Output<string[]>;
    /**
     * The user provided display name of the ScanConfig.
     */
    public readonly displayName!: pulumi.Output<string>;
    /**
     * Controls export of scan configurations and results to Security Command Center.
     */
    public readonly exportToSecurityCommandCenter!: pulumi.Output<string>;
    /**
     * Whether to keep scanning even if most requests return HTTP error codes.
     */
    public readonly ignoreHttpStatusErrors!: pulumi.Output<boolean>;
    /**
     * Whether the scan config is managed by Web Security Scanner, output only.
     */
    public readonly managedScan!: pulumi.Output<boolean>;
    /**
     * The maximum QPS during scanning. A valid value ranges from 5 to 20 inclusively. If the field is unspecified or its value is set 0, server will default to 15. Other values outside of [5, 20] range will be rejected with INVALID_ARGUMENT error.
     */
    public readonly maxQps!: pulumi.Output<number>;
    /**
     * The resource name of the ScanConfig. The name follows the format of 'projects/{projectId}/scanConfigs/{scanConfigId}'. The ScanConfig IDs are generated by the system.
     */
    public readonly name!: pulumi.Output<string>;
    /**
     * The risk level selected for the scan
     */
    public readonly riskLevel!: pulumi.Output<string>;
    /**
     * The schedule of the ScanConfig.
     */
    public readonly schedule!: pulumi.Output<outputs.websecurityscanner.v1.ScheduleResponse>;
    /**
     * The starting URLs from which the scanner finds site pages.
     */
    public readonly startingUrls!: pulumi.Output<string[]>;
    /**
     * Whether the scan configuration has enabled static IP address scan feature. If enabled, the scanner will access applications from static IP addresses.
     */
    public readonly staticIpScan!: pulumi.Output<boolean>;
    /**
     * The user agent used during scanning.
     */
    public readonly userAgent!: pulumi.Output<string>;

    /**
     * Create a ScanConfig resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: ScanConfigArgs, opts?: pulumi.CustomResourceOptions) {
        let resourceInputs: pulumi.Inputs = {};
        opts = opts || {};
        if (!opts.id) {
            if ((!args || args.displayName === undefined) && !opts.urn) {
                throw new Error("Missing required property 'displayName'");
            }
            if ((!args || args.startingUrls === undefined) && !opts.urn) {
                throw new Error("Missing required property 'startingUrls'");
            }
            resourceInputs["authentication"] = args ? args.authentication : undefined;
            resourceInputs["blacklistPatterns"] = args ? args.blacklistPatterns : undefined;
            resourceInputs["displayName"] = args ? args.displayName : undefined;
            resourceInputs["exportToSecurityCommandCenter"] = args ? args.exportToSecurityCommandCenter : undefined;
            resourceInputs["ignoreHttpStatusErrors"] = args ? args.ignoreHttpStatusErrors : undefined;
            resourceInputs["managedScan"] = args ? args.managedScan : undefined;
            resourceInputs["maxQps"] = args ? args.maxQps : undefined;
            resourceInputs["name"] = args ? args.name : undefined;
            resourceInputs["project"] = args ? args.project : undefined;
            resourceInputs["riskLevel"] = args ? args.riskLevel : undefined;
            resourceInputs["schedule"] = args ? args.schedule : undefined;
            resourceInputs["startingUrls"] = args ? args.startingUrls : undefined;
            resourceInputs["staticIpScan"] = args ? args.staticIpScan : undefined;
            resourceInputs["userAgent"] = args ? args.userAgent : undefined;
        } else {
            resourceInputs["authentication"] = undefined /*out*/;
            resourceInputs["blacklistPatterns"] = undefined /*out*/;
            resourceInputs["displayName"] = undefined /*out*/;
            resourceInputs["exportToSecurityCommandCenter"] = undefined /*out*/;
            resourceInputs["ignoreHttpStatusErrors"] = undefined /*out*/;
            resourceInputs["managedScan"] = undefined /*out*/;
            resourceInputs["maxQps"] = undefined /*out*/;
            resourceInputs["name"] = undefined /*out*/;
            resourceInputs["riskLevel"] = undefined /*out*/;
            resourceInputs["schedule"] = undefined /*out*/;
            resourceInputs["startingUrls"] = undefined /*out*/;
            resourceInputs["staticIpScan"] = undefined /*out*/;
            resourceInputs["userAgent"] = undefined /*out*/;
        }
        if (!opts.version) {
            opts = pulumi.mergeOptions(opts, { version: utilities.getVersion()});
        }
        super(ScanConfig.__pulumiType, name, resourceInputs, opts);
    }
}

/**
 * The set of arguments for constructing a ScanConfig resource.
 */
export interface ScanConfigArgs {
    /**
     * The authentication configuration. If specified, service will use the authentication configuration during scanning.
     */
    authentication?: pulumi.Input<inputs.websecurityscanner.v1.AuthenticationArgs>;
    /**
     * The excluded URL patterns as described in https://cloud.google.com/security-command-center/docs/how-to-use-web-security-scanner#excluding_urls
     */
    blacklistPatterns?: pulumi.Input<pulumi.Input<string>[]>;
    /**
     * The user provided display name of the ScanConfig.
     */
    displayName: pulumi.Input<string>;
    /**
     * Controls export of scan configurations and results to Security Command Center.
     */
    exportToSecurityCommandCenter?: pulumi.Input<enums.websecurityscanner.v1.ScanConfigExportToSecurityCommandCenter>;
    /**
     * Whether to keep scanning even if most requests return HTTP error codes.
     */
    ignoreHttpStatusErrors?: pulumi.Input<boolean>;
    /**
     * Whether the scan config is managed by Web Security Scanner, output only.
     */
    managedScan?: pulumi.Input<boolean>;
    /**
     * The maximum QPS during scanning. A valid value ranges from 5 to 20 inclusively. If the field is unspecified or its value is set 0, server will default to 15. Other values outside of [5, 20] range will be rejected with INVALID_ARGUMENT error.
     */
    maxQps?: pulumi.Input<number>;
    /**
     * The resource name of the ScanConfig. The name follows the format of 'projects/{projectId}/scanConfigs/{scanConfigId}'. The ScanConfig IDs are generated by the system.
     */
    name?: pulumi.Input<string>;
    project?: pulumi.Input<string>;
    /**
     * The risk level selected for the scan
     */
    riskLevel?: pulumi.Input<enums.websecurityscanner.v1.ScanConfigRiskLevel>;
    /**
     * The schedule of the ScanConfig.
     */
    schedule?: pulumi.Input<inputs.websecurityscanner.v1.ScheduleArgs>;
    /**
     * The starting URLs from which the scanner finds site pages.
     */
    startingUrls: pulumi.Input<pulumi.Input<string>[]>;
    /**
     * Whether the scan configuration has enabled static IP address scan feature. If enabled, the scanner will access applications from static IP addresses.
     */
    staticIpScan?: pulumi.Input<boolean>;
    /**
     * The user agent used during scanning.
     */
    userAgent?: pulumi.Input<enums.websecurityscanner.v1.ScanConfigUserAgent>;
}
