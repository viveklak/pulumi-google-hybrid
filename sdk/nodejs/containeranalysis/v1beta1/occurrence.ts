// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import { input as inputs, output as outputs, enums } from "../../types";
import * as utilities from "../../utilities";

/**
 * Creates a new occurrence.
 * Auto-naming is currently not supported for this resource.
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
    public static readonly __pulumiType = 'google-native:containeranalysis/v1beta1:Occurrence';

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
     * Describes an attestation of an artifact.
     */
    public readonly attestation!: pulumi.Output<outputs.containeranalysis.v1beta1.DetailsResponse>;
    /**
     * Describes a verifiable build.
     */
    public readonly build!: pulumi.Output<outputs.containeranalysis.v1beta1.GrafeasV1beta1BuildDetailsResponse>;
    /**
     * The time this occurrence was created.
     */
    public /*out*/ readonly createTime!: pulumi.Output<string>;
    /**
     * Describes the deployment of an artifact on a runtime.
     */
    public readonly deployment!: pulumi.Output<outputs.containeranalysis.v1beta1.GrafeasV1beta1DeploymentDetailsResponse>;
    /**
     * Describes how this resource derives from the basis in the associated note.
     */
    public readonly derivedImage!: pulumi.Output<outputs.containeranalysis.v1beta1.GrafeasV1beta1ImageDetailsResponse>;
    /**
     * Describes when a resource was discovered.
     */
    public readonly discovered!: pulumi.Output<outputs.containeranalysis.v1beta1.GrafeasV1beta1DiscoveryDetailsResponse>;
    /**
     * Describes the installation of a package on the linked resource.
     */
    public readonly installation!: pulumi.Output<outputs.containeranalysis.v1beta1.GrafeasV1beta1PackageDetailsResponse>;
    /**
     * Describes a specific in-toto link.
     */
    public readonly intoto!: pulumi.Output<outputs.containeranalysis.v1beta1.GrafeasV1beta1IntotoDetailsResponse>;
    /**
     * This explicitly denotes which of the occurrence details are specified. This field can be used as a filter in list requests.
     */
    public /*out*/ readonly kind!: pulumi.Output<string>;
    /**
     * The name of the occurrence in the form of `projects/[PROJECT_ID]/occurrences/[OCCURRENCE_ID]`.
     */
    public /*out*/ readonly name!: pulumi.Output<string>;
    /**
     * Immutable. The analysis note associated with this occurrence, in the form of `projects/[PROVIDER_ID]/notes/[NOTE_ID]`. This field can be used as a filter in list requests.
     */
    public readonly noteName!: pulumi.Output<string>;
    /**
     * A description of actions that can be taken to remedy the note.
     */
    public readonly remediation!: pulumi.Output<string>;
    /**
     * Immutable. The resource for which the occurrence applies.
     */
    public readonly resource!: pulumi.Output<outputs.containeranalysis.v1beta1.ResourceResponse>;
    /**
     * Describes a specific software bill of materials document.
     */
    public readonly sbom!: pulumi.Output<outputs.containeranalysis.v1beta1.DocumentOccurrenceResponse>;
    /**
     * Describes a specific SPDX File.
     */
    public readonly spdxFile!: pulumi.Output<outputs.containeranalysis.v1beta1.FileOccurrenceResponse>;
    /**
     * Describes a specific SPDX Package.
     */
    public readonly spdxPackage!: pulumi.Output<outputs.containeranalysis.v1beta1.PackageInfoOccurrenceResponse>;
    /**
     * Describes a specific SPDX Relationship.
     */
    public readonly spdxRelationship!: pulumi.Output<outputs.containeranalysis.v1beta1.RelationshipOccurrenceResponse>;
    /**
     * The time this occurrence was last updated.
     */
    public /*out*/ readonly updateTime!: pulumi.Output<string>;
    /**
     * Describes a security vulnerability.
     */
    public readonly vulnerability!: pulumi.Output<outputs.containeranalysis.v1beta1.GrafeasV1beta1VulnerabilityDetailsResponse>;

    /**
     * Create a Occurrence resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: OccurrenceArgs, opts?: pulumi.CustomResourceOptions) {
        let resourceInputs: pulumi.Inputs = {};
        opts = opts || {};
        if (!opts.id) {
            if ((!args || args.noteName === undefined) && !opts.urn) {
                throw new Error("Missing required property 'noteName'");
            }
            if ((!args || args.resource === undefined) && !opts.urn) {
                throw new Error("Missing required property 'resource'");
            }
            resourceInputs["attestation"] = args ? args.attestation : undefined;
            resourceInputs["build"] = args ? args.build : undefined;
            resourceInputs["deployment"] = args ? args.deployment : undefined;
            resourceInputs["derivedImage"] = args ? args.derivedImage : undefined;
            resourceInputs["discovered"] = args ? args.discovered : undefined;
            resourceInputs["installation"] = args ? args.installation : undefined;
            resourceInputs["intoto"] = args ? args.intoto : undefined;
            resourceInputs["noteName"] = args ? args.noteName : undefined;
            resourceInputs["project"] = args ? args.project : undefined;
            resourceInputs["remediation"] = args ? args.remediation : undefined;
            resourceInputs["resource"] = args ? args.resource : undefined;
            resourceInputs["sbom"] = args ? args.sbom : undefined;
            resourceInputs["spdxFile"] = args ? args.spdxFile : undefined;
            resourceInputs["spdxPackage"] = args ? args.spdxPackage : undefined;
            resourceInputs["spdxRelationship"] = args ? args.spdxRelationship : undefined;
            resourceInputs["vulnerability"] = args ? args.vulnerability : undefined;
            resourceInputs["createTime"] = undefined /*out*/;
            resourceInputs["kind"] = undefined /*out*/;
            resourceInputs["name"] = undefined /*out*/;
            resourceInputs["updateTime"] = undefined /*out*/;
        } else {
            resourceInputs["attestation"] = undefined /*out*/;
            resourceInputs["build"] = undefined /*out*/;
            resourceInputs["createTime"] = undefined /*out*/;
            resourceInputs["deployment"] = undefined /*out*/;
            resourceInputs["derivedImage"] = undefined /*out*/;
            resourceInputs["discovered"] = undefined /*out*/;
            resourceInputs["installation"] = undefined /*out*/;
            resourceInputs["intoto"] = undefined /*out*/;
            resourceInputs["kind"] = undefined /*out*/;
            resourceInputs["name"] = undefined /*out*/;
            resourceInputs["noteName"] = undefined /*out*/;
            resourceInputs["remediation"] = undefined /*out*/;
            resourceInputs["resource"] = undefined /*out*/;
            resourceInputs["sbom"] = undefined /*out*/;
            resourceInputs["spdxFile"] = undefined /*out*/;
            resourceInputs["spdxPackage"] = undefined /*out*/;
            resourceInputs["spdxRelationship"] = undefined /*out*/;
            resourceInputs["updateTime"] = undefined /*out*/;
            resourceInputs["vulnerability"] = undefined /*out*/;
        }
        if (!opts.version) {
            opts = pulumi.mergeOptions(opts, { version: utilities.getVersion()});
        }
        super(Occurrence.__pulumiType, name, resourceInputs, opts);
    }
}

/**
 * The set of arguments for constructing a Occurrence resource.
 */
export interface OccurrenceArgs {
    /**
     * Describes an attestation of an artifact.
     */
    attestation?: pulumi.Input<inputs.containeranalysis.v1beta1.DetailsArgs>;
    /**
     * Describes a verifiable build.
     */
    build?: pulumi.Input<inputs.containeranalysis.v1beta1.GrafeasV1beta1BuildDetailsArgs>;
    /**
     * Describes the deployment of an artifact on a runtime.
     */
    deployment?: pulumi.Input<inputs.containeranalysis.v1beta1.GrafeasV1beta1DeploymentDetailsArgs>;
    /**
     * Describes how this resource derives from the basis in the associated note.
     */
    derivedImage?: pulumi.Input<inputs.containeranalysis.v1beta1.GrafeasV1beta1ImageDetailsArgs>;
    /**
     * Describes when a resource was discovered.
     */
    discovered?: pulumi.Input<inputs.containeranalysis.v1beta1.GrafeasV1beta1DiscoveryDetailsArgs>;
    /**
     * Describes the installation of a package on the linked resource.
     */
    installation?: pulumi.Input<inputs.containeranalysis.v1beta1.GrafeasV1beta1PackageDetailsArgs>;
    /**
     * Describes a specific in-toto link.
     */
    intoto?: pulumi.Input<inputs.containeranalysis.v1beta1.GrafeasV1beta1IntotoDetailsArgs>;
    /**
     * Immutable. The analysis note associated with this occurrence, in the form of `projects/[PROVIDER_ID]/notes/[NOTE_ID]`. This field can be used as a filter in list requests.
     */
    noteName: pulumi.Input<string>;
    project?: pulumi.Input<string>;
    /**
     * A description of actions that can be taken to remedy the note.
     */
    remediation?: pulumi.Input<string>;
    /**
     * Immutable. The resource for which the occurrence applies.
     */
    resource: pulumi.Input<inputs.containeranalysis.v1beta1.ResourceArgs>;
    /**
     * Describes a specific software bill of materials document.
     */
    sbom?: pulumi.Input<inputs.containeranalysis.v1beta1.DocumentOccurrenceArgs>;
    /**
     * Describes a specific SPDX File.
     */
    spdxFile?: pulumi.Input<inputs.containeranalysis.v1beta1.FileOccurrenceArgs>;
    /**
     * Describes a specific SPDX Package.
     */
    spdxPackage?: pulumi.Input<inputs.containeranalysis.v1beta1.PackageInfoOccurrenceArgs>;
    /**
     * Describes a specific SPDX Relationship.
     */
    spdxRelationship?: pulumi.Input<inputs.containeranalysis.v1beta1.RelationshipOccurrenceArgs>;
    /**
     * Describes a security vulnerability.
     */
    vulnerability?: pulumi.Input<inputs.containeranalysis.v1beta1.GrafeasV1beta1VulnerabilityDetailsArgs>;
}
