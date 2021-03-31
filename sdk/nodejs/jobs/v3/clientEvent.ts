// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import { input as inputs, output as outputs } from "../../types";
import * as utilities from "../../utilities";

/**
 * Report events issued when end user interacts with customer's application that uses Cloud Talent Solution. You may inspect the created events in [self service tools](https://console.cloud.google.com/talent-solution/overview). [Learn more](https://cloud.google.com/talent-solution/docs/management-tools) about self service tools.
 */
export class ClientEvent extends pulumi.CustomResource {
    /**
     * Get an existing ClientEvent resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, opts?: pulumi.CustomResourceOptions): ClientEvent {
        return new ClientEvent(name, undefined as any, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'google-cloud:jobs/v3:ClientEvent';

    /**
     * Returns true if the given object is an instance of ClientEvent.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is ClientEvent {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === ClientEvent.__pulumiType;
    }


    /**
     * Create a ClientEvent resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: ClientEventArgs, opts?: pulumi.CustomResourceOptions) {
        let inputs: pulumi.Inputs = {};
        opts = opts || {};
        if (!opts.id) {
            if ((!args || args.projectsId === undefined) && !opts.urn) {
                throw new Error("Missing required property 'projectsId'");
            }
            inputs["clientEvent"] = args ? args.clientEvent : undefined;
            inputs["projectsId"] = args ? args.projectsId : undefined;
        } else {
        }
        if (!opts.version) {
            opts = pulumi.mergeOptions(opts, { version: utilities.getVersion()});
        }
        super(ClientEvent.__pulumiType, name, inputs, opts);
    }
}

/**
 * The set of arguments for constructing a ClientEvent resource.
 */
export interface ClientEventArgs {
    /**
     * Required. Events issued when end user interacts with customer's application that uses Cloud Talent Solution.
     */
    readonly clientEvent?: pulumi.Input<inputs.jobs.v3.ClientEvent>;
    readonly projectsId: pulumi.Input<string>;
}
