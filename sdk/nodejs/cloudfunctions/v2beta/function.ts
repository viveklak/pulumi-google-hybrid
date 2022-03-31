// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import { input as inputs, output as outputs, enums } from "../../types";
import * as utilities from "../../utilities";

/**
 * Creates a new function. If a function with the given name already exists in the specified project, the long running operation will return `ALREADY_EXISTS` error.
 */
export class Function extends pulumi.CustomResource {
    /**
     * Get an existing Function resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, opts?: pulumi.CustomResourceOptions): Function {
        return new Function(name, undefined as any, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'google-native:cloudfunctions/v2beta:Function';

    /**
     * Returns true if the given object is an instance of Function.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is Function {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === Function.__pulumiType;
    }

    /**
     * Describes the Build step of the function that builds a container from the given source.
     */
    public readonly buildConfig!: pulumi.Output<outputs.cloudfunctions.v2beta.BuildConfigResponse>;
    /**
     * User-provided description of a function.
     */
    public readonly description!: pulumi.Output<string>;
    /**
     * Describe whether the function is gen1 or gen2.
     */
    public readonly environment!: pulumi.Output<string>;
    /**
     * An Eventarc trigger managed by Google Cloud Functions that fires events in response to a condition in another service.
     */
    public readonly eventTrigger!: pulumi.Output<outputs.cloudfunctions.v2beta.EventTriggerResponse>;
    /**
     * Labels associated with this Cloud Function.
     */
    public readonly labels!: pulumi.Output<{[key: string]: string}>;
    /**
     * A user-defined name of the function. Function names must be unique globally and match pattern `projects/*&#47;locations/*&#47;functions/*`
     */
    public readonly name!: pulumi.Output<string>;
    /**
     * Describes the Service being deployed. Currently deploys services to Cloud Run (fully managed).
     */
    public readonly serviceConfig!: pulumi.Output<outputs.cloudfunctions.v2beta.ServiceConfigResponse>;
    /**
     * State of the function.
     */
    public /*out*/ readonly state!: pulumi.Output<string>;
    /**
     * State Messages for this Cloud Function.
     */
    public /*out*/ readonly stateMessages!: pulumi.Output<outputs.cloudfunctions.v2beta.GoogleCloudFunctionsV2betaStateMessageResponse[]>;
    /**
     * The last update timestamp of a Cloud Function.
     */
    public /*out*/ readonly updateTime!: pulumi.Output<string>;

    /**
     * Create a Function resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args?: FunctionArgs, opts?: pulumi.CustomResourceOptions) {
        let resourceInputs: pulumi.Inputs = {};
        opts = opts || {};
        if (!opts.id) {
            resourceInputs["buildConfig"] = args ? args.buildConfig : undefined;
            resourceInputs["description"] = args ? args.description : undefined;
            resourceInputs["environment"] = args ? args.environment : undefined;
            resourceInputs["eventTrigger"] = args ? args.eventTrigger : undefined;
            resourceInputs["functionId"] = args ? args.functionId : undefined;
            resourceInputs["labels"] = args ? args.labels : undefined;
            resourceInputs["location"] = args ? args.location : undefined;
            resourceInputs["name"] = args ? args.name : undefined;
            resourceInputs["project"] = args ? args.project : undefined;
            resourceInputs["serviceConfig"] = args ? args.serviceConfig : undefined;
            resourceInputs["state"] = undefined /*out*/;
            resourceInputs["stateMessages"] = undefined /*out*/;
            resourceInputs["updateTime"] = undefined /*out*/;
        } else {
            resourceInputs["buildConfig"] = undefined /*out*/;
            resourceInputs["description"] = undefined /*out*/;
            resourceInputs["environment"] = undefined /*out*/;
            resourceInputs["eventTrigger"] = undefined /*out*/;
            resourceInputs["labels"] = undefined /*out*/;
            resourceInputs["name"] = undefined /*out*/;
            resourceInputs["serviceConfig"] = undefined /*out*/;
            resourceInputs["state"] = undefined /*out*/;
            resourceInputs["stateMessages"] = undefined /*out*/;
            resourceInputs["updateTime"] = undefined /*out*/;
        }
        opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
        super(Function.__pulumiType, name, resourceInputs, opts);
    }
}

/**
 * The set of arguments for constructing a Function resource.
 */
export interface FunctionArgs {
    /**
     * Describes the Build step of the function that builds a container from the given source.
     */
    buildConfig?: pulumi.Input<inputs.cloudfunctions.v2beta.BuildConfigArgs>;
    /**
     * User-provided description of a function.
     */
    description?: pulumi.Input<string>;
    /**
     * Describe whether the function is gen1 or gen2.
     */
    environment?: pulumi.Input<enums.cloudfunctions.v2beta.FunctionEnvironment>;
    /**
     * An Eventarc trigger managed by Google Cloud Functions that fires events in response to a condition in another service.
     */
    eventTrigger?: pulumi.Input<inputs.cloudfunctions.v2beta.EventTriggerArgs>;
    /**
     * The ID to use for the function, which will become the final component of the function's resource name. This value should be 4-63 characters, and valid characters are /a-z-/.
     */
    functionId?: pulumi.Input<string>;
    /**
     * Labels associated with this Cloud Function.
     */
    labels?: pulumi.Input<{[key: string]: pulumi.Input<string>}>;
    location?: pulumi.Input<string>;
    /**
     * A user-defined name of the function. Function names must be unique globally and match pattern `projects/*&#47;locations/*&#47;functions/*`
     */
    name?: pulumi.Input<string>;
    project?: pulumi.Input<string>;
    /**
     * Describes the Service being deployed. Currently deploys services to Cloud Run (fully managed).
     */
    serviceConfig?: pulumi.Input<inputs.cloudfunctions.v2beta.ServiceConfigArgs>;
}
