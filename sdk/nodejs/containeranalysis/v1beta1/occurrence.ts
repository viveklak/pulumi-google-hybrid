// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import { input as inputs, output as outputs } from "../../types";
import * as utilities from "../../utilities";

/**
 * Creates a new occurrence.
 */
export class Occurrence extends pulumi.CustomResource {
    /**
     * Get an existing Occurrence resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, opts?: pulumi.CustomResourceOptions): Occurrence {
        return new Occurrence(name, undefined as any, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'google-cloud:containeranalysis/v1beta1:Occurrence';

    /**
     * Returns true if the given object is an instance of Occurrence.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is Occurrence {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === Occurrence.__pulumiType;
    }


    /**
     * Create a Occurrence resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: OccurrenceArgs, opts?: pulumi.CustomResourceOptions) {
        let inputs: pulumi.Inputs = {};
        opts = opts || {};
        if (!opts.id) {
            if ((!args || args.occurrencesId === undefined) && !opts.urn) {
                throw new Error("Missing required property 'occurrencesId'");
            }
            if ((!args || args.projectsId === undefined) && !opts.urn) {
                throw new Error("Missing required property 'projectsId'");
            }
            inputs["attestation"] = args ? args.attestation : undefined;
            inputs["build"] = args ? args.build : undefined;
            inputs["createTime"] = args ? args.createTime : undefined;
            inputs["deployment"] = args ? args.deployment : undefined;
            inputs["derivedImage"] = args ? args.derivedImage : undefined;
            inputs["discovered"] = args ? args.discovered : undefined;
            inputs["installation"] = args ? args.installation : undefined;
            inputs["intoto"] = args ? args.intoto : undefined;
            inputs["kind"] = args ? args.kind : undefined;
            inputs["name"] = args ? args.name : undefined;
            inputs["noteName"] = args ? args.noteName : undefined;
            inputs["occurrencesId"] = args ? args.occurrencesId : undefined;
            inputs["projectsId"] = args ? args.projectsId : undefined;
            inputs["remediation"] = args ? args.remediation : undefined;
            inputs["resource"] = args ? args.resource : undefined;
            inputs["updateTime"] = args ? args.updateTime : undefined;
            inputs["vulnerability"] = args ? args.vulnerability : undefined;
        } else {
        }
        if (!opts.version) {
            opts = pulumi.mergeOptions(opts, { version: utilities.getVersion()});
        }
        super(Occurrence.__pulumiType, name, inputs, opts);
    }
}

/**
 * The set of arguments for constructing a Occurrence resource.
 */
export interface OccurrenceArgs {
    /**
     * Describes an attestation of an artifact.
     */
    readonly attestation?: pulumi.Input<inputs.containeranalysis.v1beta1.Details>;
    /**
     * Describes a verifiable build.
     */
    readonly build?: pulumi.Input<inputs.containeranalysis.v1beta1.GrafeasV1beta1BuildDetails>;
    /**
     * Output only. The time this occurrence was created.
     */
    readonly createTime?: pulumi.Input<string>;
    /**
     * Describes the deployment of an artifact on a runtime.
     */
    readonly deployment?: pulumi.Input<inputs.containeranalysis.v1beta1.GrafeasV1beta1DeploymentDetails>;
    /**
     * Describes how this resource derives from the basis in the associated note.
     */
    readonly derivedImage?: pulumi.Input<inputs.containeranalysis.v1beta1.GrafeasV1beta1ImageDetails>;
    /**
     * Describes when a resource was discovered.
     */
    readonly discovered?: pulumi.Input<inputs.containeranalysis.v1beta1.GrafeasV1beta1DiscoveryDetails>;
    /**
     * Describes the installation of a package on the linked resource.
     */
    readonly installation?: pulumi.Input<inputs.containeranalysis.v1beta1.GrafeasV1beta1PackageDetails>;
    /**
     * Describes a specific in-toto link.
     */
    readonly intoto?: pulumi.Input<inputs.containeranalysis.v1beta1.GrafeasV1beta1IntotoDetails>;
    /**
     * Output only. This explicitly denotes which of the occurrence details are specified. This field can be used as a filter in list requests.
     */
    readonly kind?: pulumi.Input<string>;
    /**
     * Output only. The name of the occurrence in the form of `projects/[PROJECT_ID]/occurrences/[OCCURRENCE_ID]`.
     */
    readonly name?: pulumi.Input<string>;
    /**
     * Required. Immutable. The analysis note associated with this occurrence, in the form of `projects/[PROVIDER_ID]/notes/[NOTE_ID]`. This field can be used as a filter in list requests.
     */
    readonly noteName?: pulumi.Input<string>;
    readonly occurrencesId: pulumi.Input<string>;
    readonly projectsId: pulumi.Input<string>;
    /**
     * A description of actions that can be taken to remedy the note.
     */
    readonly remediation?: pulumi.Input<string>;
    /**
     * Required. Immutable. The resource for which the occurrence applies.
     */
    readonly resource?: pulumi.Input<inputs.containeranalysis.v1beta1.Resource>;
    /**
     * Output only. The time this occurrence was last updated.
     */
    readonly updateTime?: pulumi.Input<string>;
    /**
     * Describes a security vulnerability.
     */
    readonly vulnerability?: pulumi.Input<inputs.containeranalysis.v1beta1.GrafeasV1beta1VulnerabilityDetails>;
}
