// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "../../utilities";

/**
 * Creates a new game server deployment in a given project and location.
 */
export class GameServerDeployment extends pulumi.CustomResource {
    /**
     * Get an existing GameServerDeployment resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, opts?: pulumi.CustomResourceOptions): GameServerDeployment {
        return new GameServerDeployment(name, undefined as any, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'google-cloud:gameservices/v1:GameServerDeployment';

    /**
     * Returns true if the given object is an instance of GameServerDeployment.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is GameServerDeployment {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === GameServerDeployment.__pulumiType;
    }


    /**
     * Create a GameServerDeployment resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: GameServerDeploymentArgs, opts?: pulumi.CustomResourceOptions) {
        let inputs: pulumi.Inputs = {};
        opts = opts || {};
        if (!opts.id) {
            if ((!args || args.gameServerDeploymentsId === undefined) && !opts.urn) {
                throw new Error("Missing required property 'gameServerDeploymentsId'");
            }
            if ((!args || args.locationsId === undefined) && !opts.urn) {
                throw new Error("Missing required property 'locationsId'");
            }
            if ((!args || args.projectsId === undefined) && !opts.urn) {
                throw new Error("Missing required property 'projectsId'");
            }
            inputs["createTime"] = args ? args.createTime : undefined;
            inputs["description"] = args ? args.description : undefined;
            inputs["etag"] = args ? args.etag : undefined;
            inputs["gameServerDeploymentsId"] = args ? args.gameServerDeploymentsId : undefined;
            inputs["labels"] = args ? args.labels : undefined;
            inputs["locationsId"] = args ? args.locationsId : undefined;
            inputs["name"] = args ? args.name : undefined;
            inputs["projectsId"] = args ? args.projectsId : undefined;
            inputs["updateTime"] = args ? args.updateTime : undefined;
        } else {
        }
        if (!opts.version) {
            opts = pulumi.mergeOptions(opts, { version: utilities.getVersion()});
        }
        super(GameServerDeployment.__pulumiType, name, inputs, opts);
    }
}

/**
 * The set of arguments for constructing a GameServerDeployment resource.
 */
export interface GameServerDeploymentArgs {
    /**
     * Output only. The creation time.
     */
    readonly createTime?: pulumi.Input<string>;
    /**
     * Human readable description of the game server delpoyment.
     */
    readonly description?: pulumi.Input<string>;
    /**
     * ETag of the resource.
     */
    readonly etag?: pulumi.Input<string>;
    readonly gameServerDeploymentsId: pulumi.Input<string>;
    /**
     * The labels associated with this game server deployment. Each label is a key-value pair.
     */
    readonly labels?: pulumi.Input<{[key: string]: pulumi.Input<string>}>;
    readonly locationsId: pulumi.Input<string>;
    /**
     * The resource name of the game server deployment, in the following form: `projects/{project}/locations/{location}/gameServerDeployments/{deployment}`. For example, `projects/my-project/locations/global/gameServerDeployments/my-deployment`.
     */
    readonly name?: pulumi.Input<string>;
    readonly projectsId: pulumi.Input<string>;
    /**
     * Output only. The last-modified time.
     */
    readonly updateTime?: pulumi.Input<string>;
}
