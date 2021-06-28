// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import { input as inputs, output as outputs, enums } from "../../types";
import * as utilities from "../../utilities";

/**
 * Creates a new `Note`.
 */
export class Note extends pulumi.CustomResource {
    /**
     * Get an existing Note resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, opts?: pulumi.CustomResourceOptions): Note {
        return new Note(name, undefined as any, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'google-native:containeranalysis/v1alpha1:Note';

    /**
     * Returns true if the given object is an instance of Note.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is Note {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === Note.__pulumiType;
    }

    /**
     * A note describing an attestation role.
     */
    public readonly attestationAuthority!: pulumi.Output<outputs.containeranalysis.v1alpha1.AttestationAuthorityResponse>;
    /**
     * A note describing a base image.
     */
    public readonly baseImage!: pulumi.Output<outputs.containeranalysis.v1alpha1.BasisResponse>;
    /**
     * Build provenance type for a verifiable build.
     */
    public readonly buildType!: pulumi.Output<outputs.containeranalysis.v1alpha1.BuildTypeResponse>;
    /**
     * The time this note was created. This field can be used as a filter in list requests.
     */
    public /*out*/ readonly createTime!: pulumi.Output<string>;
    /**
     * A note describing something that can be deployed.
     */
    public readonly deployable!: pulumi.Output<outputs.containeranalysis.v1alpha1.DeployableResponse>;
    /**
     * A note describing a provider/analysis type.
     */
    public readonly discovery!: pulumi.Output<outputs.containeranalysis.v1alpha1.DiscoveryResponse>;
    /**
     * Time of expiration for this note, null if note does not expire.
     */
    public readonly expirationTime!: pulumi.Output<string>;
    /**
     * This explicitly denotes which kind of note is specified. This field can be used as a filter in list requests.
     */
    public /*out*/ readonly kind!: pulumi.Output<string>;
    /**
     * A detailed description of this `Note`.
     */
    public readonly longDescription!: pulumi.Output<string>;
    /**
     * The name of the note in the form "projects/{provider_project_id}/notes/{NOTE_ID}"
     */
    public readonly name!: pulumi.Output<string>;
    /**
     * A note describing a package hosted by various package managers.
     */
    public readonly package!: pulumi.Output<outputs.containeranalysis.v1alpha1.PackageResponse>;
    /**
     * URLs associated with this note
     */
    public readonly relatedUrl!: pulumi.Output<outputs.containeranalysis.v1alpha1.RelatedUrlResponse[]>;
    /**
     * A one sentence description of this `Note`.
     */
    public readonly shortDescription!: pulumi.Output<string>;
    /**
     * The time this note was last updated. This field can be used as a filter in list requests.
     */
    public /*out*/ readonly updateTime!: pulumi.Output<string>;
    /**
     * A note describing an upgrade.
     */
    public readonly upgrade!: pulumi.Output<outputs.containeranalysis.v1alpha1.UpgradeNoteResponse>;
    /**
     * A package vulnerability type of note.
     */
    public readonly vulnerabilityType!: pulumi.Output<outputs.containeranalysis.v1alpha1.VulnerabilityTypeResponse>;

    /**
     * Create a Note resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: NoteArgs, opts?: pulumi.CustomResourceOptions) {
        let inputs: pulumi.Inputs = {};
        opts = opts || {};
        if (!opts.id) {
            if ((!args || args.project === undefined) && !opts.urn) {
                throw new Error("Missing required property 'project'");
            }
            inputs["attestationAuthority"] = args ? args.attestationAuthority : undefined;
            inputs["baseImage"] = args ? args.baseImage : undefined;
            inputs["buildType"] = args ? args.buildType : undefined;
            inputs["deployable"] = args ? args.deployable : undefined;
            inputs["discovery"] = args ? args.discovery : undefined;
            inputs["expirationTime"] = args ? args.expirationTime : undefined;
            inputs["longDescription"] = args ? args.longDescription : undefined;
            inputs["name"] = args ? args.name : undefined;
            inputs["noteId"] = args ? args.noteId : undefined;
            inputs["package"] = args ? args.package : undefined;
            inputs["project"] = args ? args.project : undefined;
            inputs["relatedUrl"] = args ? args.relatedUrl : undefined;
            inputs["shortDescription"] = args ? args.shortDescription : undefined;
            inputs["upgrade"] = args ? args.upgrade : undefined;
            inputs["vulnerabilityType"] = args ? args.vulnerabilityType : undefined;
            inputs["createTime"] = undefined /*out*/;
            inputs["kind"] = undefined /*out*/;
            inputs["updateTime"] = undefined /*out*/;
        } else {
            inputs["attestationAuthority"] = undefined /*out*/;
            inputs["baseImage"] = undefined /*out*/;
            inputs["buildType"] = undefined /*out*/;
            inputs["createTime"] = undefined /*out*/;
            inputs["deployable"] = undefined /*out*/;
            inputs["discovery"] = undefined /*out*/;
            inputs["expirationTime"] = undefined /*out*/;
            inputs["kind"] = undefined /*out*/;
            inputs["longDescription"] = undefined /*out*/;
            inputs["name"] = undefined /*out*/;
            inputs["package"] = undefined /*out*/;
            inputs["relatedUrl"] = undefined /*out*/;
            inputs["shortDescription"] = undefined /*out*/;
            inputs["updateTime"] = undefined /*out*/;
            inputs["upgrade"] = undefined /*out*/;
            inputs["vulnerabilityType"] = undefined /*out*/;
        }
        if (!opts.version) {
            opts = pulumi.mergeOptions(opts, { version: utilities.getVersion()});
        }
        super(Note.__pulumiType, name, inputs, opts);
    }
}

/**
 * The set of arguments for constructing a Note resource.
 */
export interface NoteArgs {
    /**
     * A note describing an attestation role.
     */
    attestationAuthority?: pulumi.Input<inputs.containeranalysis.v1alpha1.AttestationAuthorityArgs>;
    /**
     * A note describing a base image.
     */
    baseImage?: pulumi.Input<inputs.containeranalysis.v1alpha1.BasisArgs>;
    /**
     * Build provenance type for a verifiable build.
     */
    buildType?: pulumi.Input<inputs.containeranalysis.v1alpha1.BuildTypeArgs>;
    /**
     * A note describing something that can be deployed.
     */
    deployable?: pulumi.Input<inputs.containeranalysis.v1alpha1.DeployableArgs>;
    /**
     * A note describing a provider/analysis type.
     */
    discovery?: pulumi.Input<inputs.containeranalysis.v1alpha1.DiscoveryArgs>;
    /**
     * Time of expiration for this note, null if note does not expire.
     */
    expirationTime?: pulumi.Input<string>;
    /**
     * A detailed description of this `Note`.
     */
    longDescription?: pulumi.Input<string>;
    /**
     * The name of the note in the form "projects/{provider_project_id}/notes/{NOTE_ID}"
     */
    name?: pulumi.Input<string>;
    noteId?: pulumi.Input<string>;
    /**
     * A note describing a package hosted by various package managers.
     */
    package?: pulumi.Input<inputs.containeranalysis.v1alpha1.PackageArgs>;
    project: pulumi.Input<string>;
    /**
     * URLs associated with this note
     */
    relatedUrl?: pulumi.Input<pulumi.Input<inputs.containeranalysis.v1alpha1.RelatedUrlArgs>[]>;
    /**
     * A one sentence description of this `Note`.
     */
    shortDescription?: pulumi.Input<string>;
    /**
     * A note describing an upgrade.
     */
    upgrade?: pulumi.Input<inputs.containeranalysis.v1alpha1.UpgradeNoteArgs>;
    /**
     * A package vulnerability type of note.
     */
    vulnerabilityType?: pulumi.Input<inputs.containeranalysis.v1alpha1.VulnerabilityTypeArgs>;
}
