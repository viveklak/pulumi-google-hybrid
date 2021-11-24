// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import { input as inputs, output as outputs, enums } from "../../types";
import * as utilities from "../../utilities";

/**
 * Creates a History. The returned History will have the id set. May return any of the following canonical error codes: - PERMISSION_DENIED - if the user is not authorized to write to project - INVALID_ARGUMENT - if the request is malformed - NOT_FOUND - if the containing project does not exist
 * Note - this resource's API doesn't support deletion. When deleted, the resource will persist
 * on Google Cloud even though it will be deleted from Pulumi state.
 */
export class History extends pulumi.CustomResource {
    /**
     * Get an existing History resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, opts?: pulumi.CustomResourceOptions): History {
        return new History(name, undefined as any, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'google-native:toolresults/v1beta3:History';

    /**
     * Returns true if the given object is an instance of History.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is History {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === History.__pulumiType;
    }

    /**
     * A short human-readable (plain text) name to display in the UI. Maximum of 100 characters. - In response: present if set during create. - In create request: optional
     */
    public readonly displayName!: pulumi.Output<string>;
    /**
     * A unique identifier within a project for this History. Returns INVALID_ARGUMENT if this field is set or overwritten by the caller. - In response always set - In create request: never set
     */
    public readonly historyId!: pulumi.Output<string>;
    /**
     * A name to uniquely identify a history within a project. Maximum of 200 characters. - In response always set - In create request: always set
     */
    public readonly name!: pulumi.Output<string>;
    /**
     * The platform of the test history. - In response: always set. Returns the platform of the last execution if unknown.
     */
    public readonly testPlatform!: pulumi.Output<string>;

    /**
     * Create a History resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args?: HistoryArgs, opts?: pulumi.CustomResourceOptions) {
        let resourceInputs: pulumi.Inputs = {};
        opts = opts || {};
        if (!opts.id) {
            resourceInputs["displayName"] = args ? args.displayName : undefined;
            resourceInputs["historyId"] = args ? args.historyId : undefined;
            resourceInputs["name"] = args ? args.name : undefined;
            resourceInputs["project"] = args ? args.project : undefined;
            resourceInputs["requestId"] = args ? args.requestId : undefined;
            resourceInputs["testPlatform"] = args ? args.testPlatform : undefined;
        } else {
            resourceInputs["displayName"] = undefined /*out*/;
            resourceInputs["historyId"] = undefined /*out*/;
            resourceInputs["name"] = undefined /*out*/;
            resourceInputs["testPlatform"] = undefined /*out*/;
        }
        if (!opts.version) {
            opts = pulumi.mergeOptions(opts, { version: utilities.getVersion()});
        }
        super(History.__pulumiType, name, resourceInputs, opts);
    }
}

/**
 * The set of arguments for constructing a History resource.
 */
export interface HistoryArgs {
    /**
     * A short human-readable (plain text) name to display in the UI. Maximum of 100 characters. - In response: present if set during create. - In create request: optional
     */
    displayName?: pulumi.Input<string>;
    /**
     * A unique identifier within a project for this History. Returns INVALID_ARGUMENT if this field is set or overwritten by the caller. - In response always set - In create request: never set
     */
    historyId?: pulumi.Input<string>;
    /**
     * A name to uniquely identify a history within a project. Maximum of 200 characters. - In response always set - In create request: always set
     */
    name?: pulumi.Input<string>;
    project?: pulumi.Input<string>;
    requestId?: pulumi.Input<string>;
    /**
     * The platform of the test history. - In response: always set. Returns the platform of the last execution if unknown.
     */
    testPlatform?: pulumi.Input<enums.toolresults.v1beta3.HistoryTestPlatform>;
}
