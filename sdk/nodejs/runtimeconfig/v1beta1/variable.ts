// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "../../utilities";

/**
 * Creates a variable within the given configuration. You cannot create a variable with a name that is a prefix of an existing variable name, or a name that has an existing variable name as a prefix. To learn more about creating a variable, read the [Setting and Getting Data](/deployment-manager/runtime-configurator/set-and-get-variables) documentation.
 */
export class Variable extends pulumi.CustomResource {
    /**
     * Get an existing Variable resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, opts?: pulumi.CustomResourceOptions): Variable {
        return new Variable(name, undefined as any, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'google-native:runtimeconfig/v1beta1:Variable';

    /**
     * Returns true if the given object is an instance of Variable.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is Variable {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === Variable.__pulumiType;
    }

    /**
     * The name of the variable resource, in the format: projects/[PROJECT_ID]/configs/[CONFIG_NAME]/variables/[VARIABLE_NAME] The `[PROJECT_ID]` must be a valid project ID, `[CONFIG_NAME]` must be a valid RuntimeConfig resource and `[VARIABLE_NAME]` follows Unix file system file path naming. The `[VARIABLE_NAME]` can contain ASCII letters, numbers, slashes and dashes. Slashes are used as path element separators and are not part of the `[VARIABLE_NAME]` itself, so `[VARIABLE_NAME]` must contain at least one non-slash character. Multiple slashes are coalesced into single slash character. Each path segment should match [0-9A-Za-z](?:[_.A-Za-z0-9-]{0,62}[_.A-Za-z0-9])? regular expression. The length of a `[VARIABLE_NAME]` must be less than 256 characters. Once you create a variable, you cannot change the variable name.
     */
    public readonly name!: pulumi.Output<string>;
    /**
     * The current state of the variable. The variable state indicates the outcome of the `variables().watch` call and is visible through the `get` and `list` calls.
     */
    public /*out*/ readonly state!: pulumi.Output<string>;
    /**
     * The string value of the variable. The length of the value must be less than 4096 bytes. Empty values are also accepted. For example, `text: "my text value"`. The string must be valid UTF-8.
     */
    public readonly text!: pulumi.Output<string>;
    /**
     * The time of the last variable update. Timestamp will be UTC timestamp.
     */
    public /*out*/ readonly updateTime!: pulumi.Output<string>;
    /**
     * The binary value of the variable. The length of the value must be less than 4096 bytes. Empty values are also accepted. The value must be base64 encoded, and must comply with IETF RFC4648 (https://www.ietf.org/rfc/rfc4648.txt). Only one of `value` or `text` can be set.
     */
    public readonly value!: pulumi.Output<string>;

    /**
     * Create a Variable resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: VariableArgs, opts?: pulumi.CustomResourceOptions) {
        let resourceInputs: pulumi.Inputs = {};
        opts = opts || {};
        if (!opts.id) {
            if ((!args || args.configId === undefined) && !opts.urn) {
                throw new Error("Missing required property 'configId'");
            }
            resourceInputs["configId"] = args ? args.configId : undefined;
            resourceInputs["name"] = args ? args.name : undefined;
            resourceInputs["project"] = args ? args.project : undefined;
            resourceInputs["requestId"] = args ? args.requestId : undefined;
            resourceInputs["text"] = args ? args.text : undefined;
            resourceInputs["value"] = args ? args.value : undefined;
            resourceInputs["state"] = undefined /*out*/;
            resourceInputs["updateTime"] = undefined /*out*/;
        } else {
            resourceInputs["name"] = undefined /*out*/;
            resourceInputs["state"] = undefined /*out*/;
            resourceInputs["text"] = undefined /*out*/;
            resourceInputs["updateTime"] = undefined /*out*/;
            resourceInputs["value"] = undefined /*out*/;
        }
        opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
        super(Variable.__pulumiType, name, resourceInputs, opts);
    }
}

/**
 * The set of arguments for constructing a Variable resource.
 */
export interface VariableArgs {
    configId: pulumi.Input<string>;
    /**
     * The name of the variable resource, in the format: projects/[PROJECT_ID]/configs/[CONFIG_NAME]/variables/[VARIABLE_NAME] The `[PROJECT_ID]` must be a valid project ID, `[CONFIG_NAME]` must be a valid RuntimeConfig resource and `[VARIABLE_NAME]` follows Unix file system file path naming. The `[VARIABLE_NAME]` can contain ASCII letters, numbers, slashes and dashes. Slashes are used as path element separators and are not part of the `[VARIABLE_NAME]` itself, so `[VARIABLE_NAME]` must contain at least one non-slash character. Multiple slashes are coalesced into single slash character. Each path segment should match [0-9A-Za-z](?:[_.A-Za-z0-9-]{0,62}[_.A-Za-z0-9])? regular expression. The length of a `[VARIABLE_NAME]` must be less than 256 characters. Once you create a variable, you cannot change the variable name.
     */
    name?: pulumi.Input<string>;
    project?: pulumi.Input<string>;
    /**
     * An optional but recommended unique `request_id`. If the server receives two `create()` requests with the same `request_id`, then the second request will be ignored and the first resource created and stored in the backend is returned. Empty `request_id` fields are ignored. It is responsibility of the client to ensure uniqueness of the `request_id` strings. `request_id` strings are limited to 64 characters.
     */
    requestId?: pulumi.Input<string>;
    /**
     * The string value of the variable. The length of the value must be less than 4096 bytes. Empty values are also accepted. For example, `text: "my text value"`. The string must be valid UTF-8.
     */
    text?: pulumi.Input<string>;
    /**
     * The binary value of the variable. The length of the value must be less than 4096 bytes. Empty values are also accepted. The value must be base64 encoded, and must comply with IETF RFC4648 (https://www.ietf.org/rfc/rfc4648.txt). Only one of `value` or `text` can be set.
     */
    value?: pulumi.Input<string>;
}
