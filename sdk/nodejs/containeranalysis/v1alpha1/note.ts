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
     * A note describing a compliance check.
     */
    public readonly compliance!: pulumi.Output<outputs.containeranalysis.v1alpha1.ComplianceNoteResponse>;
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
     * A note describing a dsse attestation note.
     */
    public readonly dsseAttestation!: pulumi.Output<outputs.containeranalysis.v1alpha1.DSSEAttestationNoteResponse>;
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
     * A note describing a software bill of materials.
     */
    public readonly sbom!: pulumi.Output<outputs.containeranalysis.v1alpha1.DocumentNoteResponse>;
    /**
     * A one sentence description of this `Note`.
     */
    public readonly shortDescription!: pulumi.Output<string>;
    /**
     * A note describing an SPDX File.
     */
    public readonly spdxFile!: pulumi.Output<outputs.containeranalysis.v1alpha1.FileNoteResponse>;
    /**
     * A note describing an SPDX Package.
     */
    public readonly spdxPackage!: pulumi.Output<outputs.containeranalysis.v1alpha1.PackageInfoNoteResponse>;
    /**
     * A note describing a relationship between SPDX elements.
     */
    public readonly spdxRelationship!: pulumi.Output<outputs.containeranalysis.v1alpha1.RelationshipNoteResponse>;
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
    constructor(name: string, args?: NoteArgs, opts?: pulumi.CustomResourceOptions) {
        let resourceInputs: pulumi.Inputs = {};
        opts = opts || {};
        if (!opts.id) {
            resourceInputs["attestationAuthority"] = args ? args.attestationAuthority : undefined;
            resourceInputs["baseImage"] = args ? args.baseImage : undefined;
            resourceInputs["buildType"] = args ? args.buildType : undefined;
            resourceInputs["compliance"] = args ? args.compliance : undefined;
            resourceInputs["deployable"] = args ? args.deployable : undefined;
            resourceInputs["discovery"] = args ? args.discovery : undefined;
            resourceInputs["dsseAttestation"] = args ? args.dsseAttestation : undefined;
            resourceInputs["expirationTime"] = args ? args.expirationTime : undefined;
            resourceInputs["longDescription"] = args ? args.longDescription : undefined;
            resourceInputs["name"] = args ? args.name : undefined;
            resourceInputs["noteId"] = args ? args.noteId : undefined;
            resourceInputs["package"] = args ? args.package : undefined;
            resourceInputs["project"] = args ? args.project : undefined;
            resourceInputs["relatedUrl"] = args ? args.relatedUrl : undefined;
            resourceInputs["sbom"] = args ? args.sbom : undefined;
            resourceInputs["shortDescription"] = args ? args.shortDescription : undefined;
            resourceInputs["spdxFile"] = args ? args.spdxFile : undefined;
            resourceInputs["spdxPackage"] = args ? args.spdxPackage : undefined;
            resourceInputs["spdxRelationship"] = args ? args.spdxRelationship : undefined;
            resourceInputs["upgrade"] = args ? args.upgrade : undefined;
            resourceInputs["vulnerabilityType"] = args ? args.vulnerabilityType : undefined;
            resourceInputs["createTime"] = undefined /*out*/;
            resourceInputs["kind"] = undefined /*out*/;
            resourceInputs["updateTime"] = undefined /*out*/;
        } else {
            resourceInputs["attestationAuthority"] = undefined /*out*/;
            resourceInputs["baseImage"] = undefined /*out*/;
            resourceInputs["buildType"] = undefined /*out*/;
            resourceInputs["compliance"] = undefined /*out*/;
            resourceInputs["createTime"] = undefined /*out*/;
            resourceInputs["deployable"] = undefined /*out*/;
            resourceInputs["discovery"] = undefined /*out*/;
            resourceInputs["dsseAttestation"] = undefined /*out*/;
            resourceInputs["expirationTime"] = undefined /*out*/;
            resourceInputs["kind"] = undefined /*out*/;
            resourceInputs["longDescription"] = undefined /*out*/;
            resourceInputs["name"] = undefined /*out*/;
            resourceInputs["package"] = undefined /*out*/;
            resourceInputs["relatedUrl"] = undefined /*out*/;
            resourceInputs["sbom"] = undefined /*out*/;
            resourceInputs["shortDescription"] = undefined /*out*/;
            resourceInputs["spdxFile"] = undefined /*out*/;
            resourceInputs["spdxPackage"] = undefined /*out*/;
            resourceInputs["spdxRelationship"] = undefined /*out*/;
            resourceInputs["updateTime"] = undefined /*out*/;
            resourceInputs["upgrade"] = undefined /*out*/;
            resourceInputs["vulnerabilityType"] = undefined /*out*/;
        }
        if (!opts.version) {
            opts = pulumi.mergeOptions(opts, { version: utilities.getVersion()});
        }
        super(Note.__pulumiType, name, resourceInputs, opts);
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
     * A note describing a compliance check.
     */
    compliance?: pulumi.Input<inputs.containeranalysis.v1alpha1.ComplianceNoteArgs>;
    /**
     * A note describing something that can be deployed.
     */
    deployable?: pulumi.Input<inputs.containeranalysis.v1alpha1.DeployableArgs>;
    /**
     * A note describing a provider/analysis type.
     */
    discovery?: pulumi.Input<inputs.containeranalysis.v1alpha1.DiscoveryArgs>;
    /**
     * A note describing a dsse attestation note.
     */
    dsseAttestation?: pulumi.Input<inputs.containeranalysis.v1alpha1.DSSEAttestationNoteArgs>;
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
    project?: pulumi.Input<string>;
    /**
     * URLs associated with this note
     */
    relatedUrl?: pulumi.Input<pulumi.Input<inputs.containeranalysis.v1alpha1.RelatedUrlArgs>[]>;
    /**
     * A note describing a software bill of materials.
     */
    sbom?: pulumi.Input<inputs.containeranalysis.v1alpha1.DocumentNoteArgs>;
    /**
     * A one sentence description of this `Note`.
     */
    shortDescription?: pulumi.Input<string>;
    /**
     * A note describing an SPDX File.
     */
    spdxFile?: pulumi.Input<inputs.containeranalysis.v1alpha1.FileNoteArgs>;
    /**
     * A note describing an SPDX Package.
     */
    spdxPackage?: pulumi.Input<inputs.containeranalysis.v1alpha1.PackageInfoNoteArgs>;
    /**
     * A note describing a relationship between SPDX elements.
     */
    spdxRelationship?: pulumi.Input<inputs.containeranalysis.v1alpha1.RelationshipNoteArgs>;
    /**
     * A note describing an upgrade.
     */
    upgrade?: pulumi.Input<inputs.containeranalysis.v1alpha1.UpgradeNoteArgs>;
    /**
     * A package vulnerability type of note.
     */
    vulnerabilityType?: pulumi.Input<inputs.containeranalysis.v1alpha1.VulnerabilityTypeArgs>;
}
