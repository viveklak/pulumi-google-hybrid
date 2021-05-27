// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import { input as inputs, output as outputs } from "../../types";
import * as utilities from "../../utilities";

/**
 * Creates a new Annotation record. It is valid to create Annotation objects for the same source more than once since a unique ID is assigned to each record by this service.
 */
export class DatasetAnnotationStoreAnnotation extends pulumi.CustomResource {
    /**
     * Get an existing DatasetAnnotationStoreAnnotation resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, opts?: pulumi.CustomResourceOptions): DatasetAnnotationStoreAnnotation {
        return new DatasetAnnotationStoreAnnotation(name, undefined as any, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'google-native:healthcare/v1beta1:DatasetAnnotationStoreAnnotation';

    /**
     * Returns true if the given object is an instance of DatasetAnnotationStoreAnnotation.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is DatasetAnnotationStoreAnnotation {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === DatasetAnnotationStoreAnnotation.__pulumiType;
    }

    /**
     * Details of the source.
     */
    public readonly annotationSource!: pulumi.Output<outputs.healthcare.v1beta1.AnnotationSourceResponse>;
    /**
     * Additional information for this annotation record, such as annotator and verifier information or study campaign.
     */
    public readonly customData!: pulumi.Output<{[key: string]: string}>;
    /**
     * Annotations for images. For example, bounding polygons.
     */
    public readonly imageAnnotation!: pulumi.Output<outputs.healthcare.v1beta1.ImageAnnotationResponse>;
    /**
     * Resource name of the Annotation, of the form `projects/{project_id}/locations/{location_id}/datasets/{dataset_id}/annotationStores/{annotation_store_id}/annotations/{annotation_id}`.
     */
    public readonly name!: pulumi.Output<string>;
    /**
     * Annotations for resource. For example, classification tags.
     */
    public readonly resourceAnnotation!: pulumi.Output<outputs.healthcare.v1beta1.ResourceAnnotationResponse>;
    /**
     * Annotations for sensitive texts. For example, a range that describes the location of sensitive text.
     */
    public readonly textAnnotation!: pulumi.Output<outputs.healthcare.v1beta1.SensitiveTextAnnotationResponse>;

    /**
     * Create a DatasetAnnotationStoreAnnotation resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: DatasetAnnotationStoreAnnotationArgs, opts?: pulumi.CustomResourceOptions) {
        let inputs: pulumi.Inputs = {};
        opts = opts || {};
        if (!opts.id) {
            if ((!args || args.annotationStoreId === undefined) && !opts.urn) {
                throw new Error("Missing required property 'annotationStoreId'");
            }
            if ((!args || args.datasetId === undefined) && !opts.urn) {
                throw new Error("Missing required property 'datasetId'");
            }
            if ((!args || args.location === undefined) && !opts.urn) {
                throw new Error("Missing required property 'location'");
            }
            if ((!args || args.project === undefined) && !opts.urn) {
                throw new Error("Missing required property 'project'");
            }
            inputs["annotationSource"] = args ? args.annotationSource : undefined;
            inputs["annotationStoreId"] = args ? args.annotationStoreId : undefined;
            inputs["customData"] = args ? args.customData : undefined;
            inputs["datasetId"] = args ? args.datasetId : undefined;
            inputs["imageAnnotation"] = args ? args.imageAnnotation : undefined;
            inputs["location"] = args ? args.location : undefined;
            inputs["name"] = args ? args.name : undefined;
            inputs["project"] = args ? args.project : undefined;
            inputs["resourceAnnotation"] = args ? args.resourceAnnotation : undefined;
            inputs["textAnnotation"] = args ? args.textAnnotation : undefined;
        } else {
            inputs["annotationSource"] = undefined /*out*/;
            inputs["customData"] = undefined /*out*/;
            inputs["imageAnnotation"] = undefined /*out*/;
            inputs["name"] = undefined /*out*/;
            inputs["resourceAnnotation"] = undefined /*out*/;
            inputs["textAnnotation"] = undefined /*out*/;
        }
        if (!opts.version) {
            opts = pulumi.mergeOptions(opts, { version: utilities.getVersion()});
        }
        super(DatasetAnnotationStoreAnnotation.__pulumiType, name, inputs, opts);
    }
}

/**
 * The set of arguments for constructing a DatasetAnnotationStoreAnnotation resource.
 */
export interface DatasetAnnotationStoreAnnotationArgs {
    /**
     * Details of the source.
     */
    readonly annotationSource?: pulumi.Input<inputs.healthcare.v1beta1.AnnotationSourceArgs>;
    readonly annotationStoreId: pulumi.Input<string>;
    /**
     * Additional information for this annotation record, such as annotator and verifier information or study campaign.
     */
    readonly customData?: pulumi.Input<{[key: string]: pulumi.Input<string>}>;
    readonly datasetId: pulumi.Input<string>;
    /**
     * Annotations for images. For example, bounding polygons.
     */
    readonly imageAnnotation?: pulumi.Input<inputs.healthcare.v1beta1.ImageAnnotationArgs>;
    readonly location: pulumi.Input<string>;
    /**
     * Resource name of the Annotation, of the form `projects/{project_id}/locations/{location_id}/datasets/{dataset_id}/annotationStores/{annotation_store_id}/annotations/{annotation_id}`.
     */
    readonly name?: pulumi.Input<string>;
    readonly project: pulumi.Input<string>;
    /**
     * Annotations for resource. For example, classification tags.
     */
    readonly resourceAnnotation?: pulumi.Input<inputs.healthcare.v1beta1.ResourceAnnotationArgs>;
    /**
     * Annotations for sensitive texts. For example, a range that describes the location of sensitive text.
     */
    readonly textAnnotation?: pulumi.Input<inputs.healthcare.v1beta1.SensitiveTextAnnotationArgs>;
}
