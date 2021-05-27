// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import { input as inputs, output as outputs } from "../../types";
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
    public static readonly __pulumiType = 'google-native:websecurityscanner/v1beta:ScanConfig';

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
    public readonly authentication!: pulumi.Output<outputs.websecurityscanner.v1beta.AuthenticationResponse>;
    /**
     * The excluded URL patterns as described in https://cloud.google.com/security-command-center/docs/how-to-use-web-security-scanner#excluding_urls
     */
    public readonly blacklistPatterns!: pulumi.Output<string[]>;
    /**
     * Required. The user provided display name of the ScanConfig.
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
     * Latest ScanRun if available.
     */
    public readonly latestRun!: pulumi.Output<outputs.websecurityscanner.v1beta.ScanRunResponse>;
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
    public readonly schedule!: pulumi.Output<outputs.websecurityscanner.v1beta.ScheduleResponse>;
    /**
     * Required. The starting URLs from which the scanner finds site pages.
     */
    public readonly startingUrls!: pulumi.Output<string[]>;
    /**
     * Whether the scan configuration has enabled static IP address scan feature. If enabled, the scanner will access applications from static IP addresses.
     */
    public readonly staticIpScan!: pulumi.Output<boolean>;
    /**
     * Set of Google Cloud platforms targeted by the scan. If empty, APP_ENGINE will be used as a default.
     */
    public readonly targetPlatforms!: pulumi.Output<string[]>;
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
        let inputs: pulumi.Inputs = {};
        opts = opts || {};
        if (!opts.id) {
            if ((!args || args.project === undefined) && !opts.urn) {
                throw new Error("Missing required property 'project'");
            }
            inputs["authentication"] = args ? args.authentication : undefined;
            inputs["blacklistPatterns"] = args ? args.blacklistPatterns : undefined;
            inputs["displayName"] = args ? args.displayName : undefined;
            inputs["exportToSecurityCommandCenter"] = args ? args.exportToSecurityCommandCenter : undefined;
            inputs["ignoreHttpStatusErrors"] = args ? args.ignoreHttpStatusErrors : undefined;
            inputs["latestRun"] = args ? args.latestRun : undefined;
            inputs["managedScan"] = args ? args.managedScan : undefined;
            inputs["maxQps"] = args ? args.maxQps : undefined;
            inputs["name"] = args ? args.name : undefined;
            inputs["project"] = args ? args.project : undefined;
            inputs["riskLevel"] = args ? args.riskLevel : undefined;
            inputs["schedule"] = args ? args.schedule : undefined;
            inputs["startingUrls"] = args ? args.startingUrls : undefined;
            inputs["staticIpScan"] = args ? args.staticIpScan : undefined;
            inputs["targetPlatforms"] = args ? args.targetPlatforms : undefined;
            inputs["userAgent"] = args ? args.userAgent : undefined;
        } else {
            inputs["authentication"] = undefined /*out*/;
            inputs["blacklistPatterns"] = undefined /*out*/;
            inputs["displayName"] = undefined /*out*/;
            inputs["exportToSecurityCommandCenter"] = undefined /*out*/;
            inputs["ignoreHttpStatusErrors"] = undefined /*out*/;
            inputs["latestRun"] = undefined /*out*/;
            inputs["managedScan"] = undefined /*out*/;
            inputs["maxQps"] = undefined /*out*/;
            inputs["name"] = undefined /*out*/;
            inputs["riskLevel"] = undefined /*out*/;
            inputs["schedule"] = undefined /*out*/;
            inputs["startingUrls"] = undefined /*out*/;
            inputs["staticIpScan"] = undefined /*out*/;
            inputs["targetPlatforms"] = undefined /*out*/;
            inputs["userAgent"] = undefined /*out*/;
        }
        if (!opts.version) {
            opts = pulumi.mergeOptions(opts, { version: utilities.getVersion()});
        }
        super(ScanConfig.__pulumiType, name, inputs, opts);
    }
}

/**
 * The set of arguments for constructing a ScanConfig resource.
 */
export interface ScanConfigArgs {
    /**
     * The authentication configuration. If specified, service will use the authentication configuration during scanning.
     */
    readonly authentication?: pulumi.Input<inputs.websecurityscanner.v1beta.AuthenticationArgs>;
    /**
     * The excluded URL patterns as described in https://cloud.google.com/security-command-center/docs/how-to-use-web-security-scanner#excluding_urls
     */
    readonly blacklistPatterns?: pulumi.Input<pulumi.Input<string>[]>;
    /**
     * Required. The user provided display name of the ScanConfig.
     */
    readonly displayName?: pulumi.Input<string>;
    /**
     * Controls export of scan configurations and results to Security Command Center.
     */
    readonly exportToSecurityCommandCenter?: pulumi.Input<string>;
    /**
     * Whether to keep scanning even if most requests return HTTP error codes.
     */
    readonly ignoreHttpStatusErrors?: pulumi.Input<boolean>;
    /**
     * Latest ScanRun if available.
     */
    readonly latestRun?: pulumi.Input<inputs.websecurityscanner.v1beta.ScanRunArgs>;
    /**
     * Whether the scan config is managed by Web Security Scanner, output only.
     */
    readonly managedScan?: pulumi.Input<boolean>;
    /**
     * The maximum QPS during scanning. A valid value ranges from 5 to 20 inclusively. If the field is unspecified or its value is set 0, server will default to 15. Other values outside of [5, 20] range will be rejected with INVALID_ARGUMENT error.
     */
    readonly maxQps?: pulumi.Input<number>;
    /**
     * The resource name of the ScanConfig. The name follows the format of 'projects/{projectId}/scanConfigs/{scanConfigId}'. The ScanConfig IDs are generated by the system.
     */
    readonly name?: pulumi.Input<string>;
    readonly project: pulumi.Input<string>;
    /**
     * The risk level selected for the scan
     */
    readonly riskLevel?: pulumi.Input<string>;
    /**
     * The schedule of the ScanConfig.
     */
    readonly schedule?: pulumi.Input<inputs.websecurityscanner.v1beta.ScheduleArgs>;
    /**
     * Required. The starting URLs from which the scanner finds site pages.
     */
    readonly startingUrls?: pulumi.Input<pulumi.Input<string>[]>;
    /**
     * Whether the scan configuration has enabled static IP address scan feature. If enabled, the scanner will access applications from static IP addresses.
     */
    readonly staticIpScan?: pulumi.Input<boolean>;
    /**
     * Set of Google Cloud platforms targeted by the scan. If empty, APP_ENGINE will be used as a default.
     */
    readonly targetPlatforms?: pulumi.Input<pulumi.Input<string>[]>;
    /**
     * The user agent used during scanning.
     */
    readonly userAgent?: pulumi.Input<string>;
}
