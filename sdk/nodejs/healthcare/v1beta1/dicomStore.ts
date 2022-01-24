// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import { input as inputs, output as outputs, enums } from "../../types";
import * as utilities from "../../utilities";

/**
 * Creates a new DICOM store within the parent dataset.
 */
export class DicomStore extends pulumi.CustomResource {
    /**
     * Get an existing DicomStore resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, opts?: pulumi.CustomResourceOptions): DicomStore {
        return new DicomStore(name, undefined as any, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'google-native:healthcare/v1beta1:DicomStore';

    /**
     * Returns true if the given object is an instance of DicomStore.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is DicomStore {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === DicomStore.__pulumiType;
    }

    /**
     * User-supplied key-value pairs used to organize DICOM stores. Label keys must be between 1 and 63 characters long, have a UTF-8 encoding of maximum 128 bytes, and must conform to the following PCRE regular expression: \p{Ll}\p{Lo}{0,62} Label values are optional, must be between 1 and 63 characters long, have a UTF-8 encoding of maximum 128 bytes, and must conform to the following PCRE regular expression: [\p{Ll}\p{Lo}\p{N}_-]{0,63} No more than 64 labels can be associated with a given store.
     */
    public readonly labels!: pulumi.Output<{[key: string]: string}>;
    /**
     * Resource name of the DICOM store, of the form `projects/{project_id}/locations/{location_id}/datasets/{dataset_id}/dicomStores/{dicom_store_id}`.
     */
    public readonly name!: pulumi.Output<string>;
    /**
     * Notification destination for new DICOM instances. Supplied by the client.
     */
    public readonly notificationConfig!: pulumi.Output<outputs.healthcare.v1beta1.NotificationConfigResponse>;
    /**
     * A list of streaming configs used to configure the destination of streaming exports for every DICOM instance insertion in this DICOM store. After a new config is added to `stream_configs`, DICOM instance insertions are streamed to the new destination. When a config is removed from `stream_configs`, the server stops streaming to that destination. Each config must contain a unique destination.
     */
    public readonly streamConfigs!: pulumi.Output<outputs.healthcare.v1beta1.GoogleCloudHealthcareV1beta1DicomStreamConfigResponse[]>;

    /**
     * Create a DicomStore resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: DicomStoreArgs, opts?: pulumi.CustomResourceOptions) {
        let resourceInputs: pulumi.Inputs = {};
        opts = opts || {};
        if (!opts.id) {
            if ((!args || args.datasetId === undefined) && !opts.urn) {
                throw new Error("Missing required property 'datasetId'");
            }
            resourceInputs["datasetId"] = args ? args.datasetId : undefined;
            resourceInputs["dicomStoreId"] = args ? args.dicomStoreId : undefined;
            resourceInputs["labels"] = args ? args.labels : undefined;
            resourceInputs["location"] = args ? args.location : undefined;
            resourceInputs["name"] = args ? args.name : undefined;
            resourceInputs["notificationConfig"] = args ? args.notificationConfig : undefined;
            resourceInputs["project"] = args ? args.project : undefined;
            resourceInputs["streamConfigs"] = args ? args.streamConfigs : undefined;
        } else {
            resourceInputs["labels"] = undefined /*out*/;
            resourceInputs["name"] = undefined /*out*/;
            resourceInputs["notificationConfig"] = undefined /*out*/;
            resourceInputs["streamConfigs"] = undefined /*out*/;
        }
        opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
        super(DicomStore.__pulumiType, name, resourceInputs, opts);
    }
}

/**
 * The set of arguments for constructing a DicomStore resource.
 */
export interface DicomStoreArgs {
    datasetId: pulumi.Input<string>;
    dicomStoreId?: pulumi.Input<string>;
    /**
     * User-supplied key-value pairs used to organize DICOM stores. Label keys must be between 1 and 63 characters long, have a UTF-8 encoding of maximum 128 bytes, and must conform to the following PCRE regular expression: \p{Ll}\p{Lo}{0,62} Label values are optional, must be between 1 and 63 characters long, have a UTF-8 encoding of maximum 128 bytes, and must conform to the following PCRE regular expression: [\p{Ll}\p{Lo}\p{N}_-]{0,63} No more than 64 labels can be associated with a given store.
     */
    labels?: pulumi.Input<{[key: string]: pulumi.Input<string>}>;
    location?: pulumi.Input<string>;
    /**
     * Resource name of the DICOM store, of the form `projects/{project_id}/locations/{location_id}/datasets/{dataset_id}/dicomStores/{dicom_store_id}`.
     */
    name?: pulumi.Input<string>;
    /**
     * Notification destination for new DICOM instances. Supplied by the client.
     */
    notificationConfig?: pulumi.Input<inputs.healthcare.v1beta1.NotificationConfigArgs>;
    project?: pulumi.Input<string>;
    /**
     * A list of streaming configs used to configure the destination of streaming exports for every DICOM instance insertion in this DICOM store. After a new config is added to `stream_configs`, DICOM instance insertions are streamed to the new destination. When a config is removed from `stream_configs`, the server stops streaming to that destination. Each config must contain a unique destination.
     */
    streamConfigs?: pulumi.Input<pulumi.Input<inputs.healthcare.v1beta1.GoogleCloudHealthcareV1beta1DicomStreamConfigArgs>[]>;
}
