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
    public static readonly __pulumiType = 'google-cloud:healthcare/v1beta1:DatasetAnnotationStoreAnnotation';

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
            if ((!args || args.annotationStoresId === undefined) && !opts.urn) {
                throw new Error("Missing required property 'annotationStoresId'");
            }
            if ((!args || args.annotationsId === undefined) && !opts.urn) {
                throw new Error("Missing required property 'annotationsId'");
            }
            if ((!args || args.datasetsId === undefined) && !opts.urn) {
                throw new Error("Missing required property 'datasetsId'");
            }
            if ((!args || args.locationsId === undefined) && !opts.urn) {
                throw new Error("Missing required property 'locationsId'");
            }
            if ((!args || args.projectsId === undefined) && !opts.urn) {
                throw new Error("Missing required property 'projectsId'");
            }
            inputs["annotationSource"] = args ? args.annotationSource : undefined;
            inputs["annotationStoresId"] = args ? args.annotationStoresId : undefined;
            inputs["annotationsId"] = args ? args.annotationsId : undefined;
            inputs["customData"] = args ? args.customData : undefined;
            inputs["datasetsId"] = args ? args.datasetsId : undefined;
            inputs["imageAnnotation"] = args ? args.imageAnnotation : undefined;
            inputs["locationsId"] = args ? args.locationsId : undefined;
            inputs["name"] = args ? args.name : undefined;
            inputs["projectsId"] = args ? args.projectsId : undefined;
            inputs["resourceAnnotation"] = args ? args.resourceAnnotation : undefined;
            inputs["textAnnotation"] = args ? args.textAnnotation : undefined;
        } else {
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
    readonly annotationSource?: pulumi.Input<inputs.healthcare.v1beta1.AnnotationSource>;
    readonly annotationStoresId: pulumi.Input<string>;
    readonly annotationsId: pulumi.Input<string>;
    /**
     * Additional information for this annotation record, such as annotator and verifier information or study campaign.
     */
    readonly customData?: pulumi.Input<{[key: string]: pulumi.Input<string>}>;
    readonly datasetsId: pulumi.Input<string>;
    /**
     * Annotations for images. For example, bounding polygons.
     */
    readonly imageAnnotation?: pulumi.Input<inputs.healthcare.v1beta1.ImageAnnotation>;
    readonly locationsId: pulumi.Input<string>;
    /**
     * Resource name of the Annotation, of the form `projects/{project_id}/locations/{location_id}/datasets/{dataset_id}/annotationStores/{annotation_store_id}/annotations/{annotation_id}`.
     */
    readonly name?: pulumi.Input<string>;
    readonly projectsId: pulumi.Input<string>;
    /**
     * Annotations for resource. For example, classification tags.
     */
    readonly resourceAnnotation?: pulumi.Input<inputs.healthcare.v1beta1.ResourceAnnotation>;
    /**
     * Annotations for sensitive texts. For example, a range that describes the location of sensitive text.
     */
    readonly textAnnotation?: pulumi.Input<inputs.healthcare.v1beta1.SensitiveTextAnnotation>;
}
