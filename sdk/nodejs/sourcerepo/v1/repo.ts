// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import { input as inputs, output as outputs, enums } from "../../types";
import * as utilities from "../../utilities";

/**
 * Creates a repo in the given project with the given name. If the named repository already exists, `CreateRepo` returns `ALREADY_EXISTS`.
 */
export class Repo extends pulumi.CustomResource {
    /**
     * Get an existing Repo resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, opts?: pulumi.CustomResourceOptions): Repo {
        return new Repo(name, undefined as any, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'google-native:sourcerepo/v1:Repo';

    /**
     * Returns true if the given object is an instance of Repo.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is Repo {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === Repo.__pulumiType;
    }

    /**
     * How this repository mirrors a repository managed by another service. Read-only field.
     */
    public readonly mirrorConfig!: pulumi.Output<outputs.sourcerepo.v1.MirrorConfigResponse>;
    /**
     * Resource name of the repository, of the form `projects//repos/`. The repo name may contain slashes. eg, `projects/myproject/repos/name/with/slash`
     */
    public readonly name!: pulumi.Output<string>;
    /**
     * How this repository publishes a change in the repository through Cloud Pub/Sub. Keyed by the topic names.
     */
    public readonly pubsubConfigs!: pulumi.Output<{[key: string]: string}>;
    /**
     * The disk usage of the repo, in bytes. Read-only field. Size is only returned by GetRepo.
     */
    public readonly size!: pulumi.Output<string>;
    /**
     * URL to clone the repository from Google Cloud Source Repositories. Read-only field.
     */
    public readonly url!: pulumi.Output<string>;

    /**
     * Create a Repo resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args?: RepoArgs, opts?: pulumi.CustomResourceOptions) {
        let resourceInputs: pulumi.Inputs = {};
        opts = opts || {};
        if (!opts.id) {
            resourceInputs["mirrorConfig"] = args ? args.mirrorConfig : undefined;
            resourceInputs["name"] = args ? args.name : undefined;
            resourceInputs["project"] = args ? args.project : undefined;
            resourceInputs["pubsubConfigs"] = args ? args.pubsubConfigs : undefined;
            resourceInputs["size"] = args ? args.size : undefined;
            resourceInputs["url"] = args ? args.url : undefined;
        } else {
            resourceInputs["mirrorConfig"] = undefined /*out*/;
            resourceInputs["name"] = undefined /*out*/;
            resourceInputs["pubsubConfigs"] = undefined /*out*/;
            resourceInputs["size"] = undefined /*out*/;
            resourceInputs["url"] = undefined /*out*/;
        }
        opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
        super(Repo.__pulumiType, name, resourceInputs, opts);
    }
}

/**
 * The set of arguments for constructing a Repo resource.
 */
export interface RepoArgs {
    /**
     * How this repository mirrors a repository managed by another service. Read-only field.
     */
    mirrorConfig?: pulumi.Input<inputs.sourcerepo.v1.MirrorConfigArgs>;
    /**
     * Resource name of the repository, of the form `projects//repos/`. The repo name may contain slashes. eg, `projects/myproject/repos/name/with/slash`
     */
    name?: pulumi.Input<string>;
    project?: pulumi.Input<string>;
    /**
     * How this repository publishes a change in the repository through Cloud Pub/Sub. Keyed by the topic names.
     */
    pubsubConfigs?: pulumi.Input<{[key: string]: pulumi.Input<string>}>;
    /**
     * The disk usage of the repo, in bytes. Read-only field. Size is only returned by GetRepo.
     */
    size?: pulumi.Input<string>;
    /**
     * URL to clone the repository from Google Cloud Source Repositories. Read-only field.
     */
    url?: pulumi.Input<string>;
}
