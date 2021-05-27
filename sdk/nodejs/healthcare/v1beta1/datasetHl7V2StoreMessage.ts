// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import { input as inputs, output as outputs } from "../../types";
import * as utilities from "../../utilities";

/**
 * Parses and stores an HL7v2 message. This method triggers an asynchronous notification to any Pub/Sub topic configured in Hl7V2Store.Hl7V2NotificationConfig, if the filtering matches the message. If an MLLP adapter is configured to listen to a Pub/Sub topic, the adapter transmits the message when a notification is received.
 */
export class DatasetHl7V2StoreMessage extends pulumi.CustomResource {
    /**
     * Get an existing DatasetHl7V2StoreMessage resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, opts?: pulumi.CustomResourceOptions): DatasetHl7V2StoreMessage {
        return new DatasetHl7V2StoreMessage(name, undefined as any, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'google-native:healthcare/v1beta1:DatasetHl7V2StoreMessage';

    /**
     * Returns true if the given object is an instance of DatasetHl7V2StoreMessage.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is DatasetHl7V2StoreMessage {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === DatasetHl7V2StoreMessage.__pulumiType;
    }

    /**
     * The datetime when the message was created. Set by the server.
     */
    public /*out*/ readonly createTime!: pulumi.Output<string>;
    /**
     * Raw message bytes.
     */
    public readonly data!: pulumi.Output<string>;
    /**
     * User-supplied key-value pairs used to organize HL7v2 stores. Label keys must be between 1 and 63 characters long, have a UTF-8 encoding of maximum 128 bytes, and must conform to the following PCRE regular expression: \p{Ll}\p{Lo}{0,62} Label values are optional, must be between 1 and 63 characters long, have a UTF-8 encoding of maximum 128 bytes, and must conform to the following PCRE regular expression: [\p{Ll}\p{Lo}\p{N}_-]{0,63} No more than 64 labels can be associated with a given store.
     */
    public readonly labels!: pulumi.Output<{[key: string]: string}>;
    /**
     * The message type for this message. MSH-9.1.
     */
    public readonly messageType!: pulumi.Output<string>;
    /**
     * Resource name of the Message, of the form `projects/{project_id}/datasets/{dataset_id}/hl7V2Stores/{hl7_v2_store_id}/messages/{message_id}`. Assigned by the server.
     */
    public readonly name!: pulumi.Output<string>;
    /**
     * The parsed version of the raw message data.
     */
    public /*out*/ readonly parsedData!: pulumi.Output<outputs.healthcare.v1beta1.ParsedDataResponse>;
    /**
     * All patient IDs listed in the PID-2, PID-3, and PID-4 segments of this message.
     */
    public readonly patientIds!: pulumi.Output<outputs.healthcare.v1beta1.PatientIdResponse[]>;
    /**
     * The parsed version of the raw message data schematized according to this store's schemas and type definitions.
     */
    public readonly schematizedData!: pulumi.Output<outputs.healthcare.v1beta1.SchematizedDataResponse>;
    /**
     * The hospital that this message came from. MSH-4.
     */
    public readonly sendFacility!: pulumi.Output<string>;
    /**
     * The datetime the sending application sent this message. MSH-7.
     */
    public readonly sendTime!: pulumi.Output<string>;

    /**
     * Create a DatasetHl7V2StoreMessage resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: DatasetHl7V2StoreMessageArgs, opts?: pulumi.CustomResourceOptions) {
        let inputs: pulumi.Inputs = {};
        opts = opts || {};
        if (!opts.id) {
            if ((!args || args.datasetId === undefined) && !opts.urn) {
                throw new Error("Missing required property 'datasetId'");
            }
            if ((!args || args.hl7V2StoreId === undefined) && !opts.urn) {
                throw new Error("Missing required property 'hl7V2StoreId'");
            }
            if ((!args || args.location === undefined) && !opts.urn) {
                throw new Error("Missing required property 'location'");
            }
            if ((!args || args.project === undefined) && !opts.urn) {
                throw new Error("Missing required property 'project'");
            }
            inputs["data"] = args ? args.data : undefined;
            inputs["datasetId"] = args ? args.datasetId : undefined;
            inputs["hl7V2StoreId"] = args ? args.hl7V2StoreId : undefined;
            inputs["labels"] = args ? args.labels : undefined;
            inputs["location"] = args ? args.location : undefined;
            inputs["messageType"] = args ? args.messageType : undefined;
            inputs["name"] = args ? args.name : undefined;
            inputs["patientIds"] = args ? args.patientIds : undefined;
            inputs["project"] = args ? args.project : undefined;
            inputs["schematizedData"] = args ? args.schematizedData : undefined;
            inputs["sendFacility"] = args ? args.sendFacility : undefined;
            inputs["sendTime"] = args ? args.sendTime : undefined;
            inputs["createTime"] = undefined /*out*/;
            inputs["parsedData"] = undefined /*out*/;
        } else {
            inputs["createTime"] = undefined /*out*/;
            inputs["data"] = undefined /*out*/;
            inputs["labels"] = undefined /*out*/;
            inputs["messageType"] = undefined /*out*/;
            inputs["name"] = undefined /*out*/;
            inputs["parsedData"] = undefined /*out*/;
            inputs["patientIds"] = undefined /*out*/;
            inputs["schematizedData"] = undefined /*out*/;
            inputs["sendFacility"] = undefined /*out*/;
            inputs["sendTime"] = undefined /*out*/;
        }
        if (!opts.version) {
            opts = pulumi.mergeOptions(opts, { version: utilities.getVersion()});
        }
        super(DatasetHl7V2StoreMessage.__pulumiType, name, inputs, opts);
    }
}

/**
 * The set of arguments for constructing a DatasetHl7V2StoreMessage resource.
 */
export interface DatasetHl7V2StoreMessageArgs {
    /**
     * Raw message bytes.
     */
    readonly data?: pulumi.Input<string>;
    readonly datasetId: pulumi.Input<string>;
    readonly hl7V2StoreId: pulumi.Input<string>;
    /**
     * User-supplied key-value pairs used to organize HL7v2 stores. Label keys must be between 1 and 63 characters long, have a UTF-8 encoding of maximum 128 bytes, and must conform to the following PCRE regular expression: \p{Ll}\p{Lo}{0,62} Label values are optional, must be between 1 and 63 characters long, have a UTF-8 encoding of maximum 128 bytes, and must conform to the following PCRE regular expression: [\p{Ll}\p{Lo}\p{N}_-]{0,63} No more than 64 labels can be associated with a given store.
     */
    readonly labels?: pulumi.Input<{[key: string]: pulumi.Input<string>}>;
    readonly location: pulumi.Input<string>;
    /**
     * The message type for this message. MSH-9.1.
     */
    readonly messageType?: pulumi.Input<string>;
    /**
     * Resource name of the Message, of the form `projects/{project_id}/datasets/{dataset_id}/hl7V2Stores/{hl7_v2_store_id}/messages/{message_id}`. Assigned by the server.
     */
    readonly name?: pulumi.Input<string>;
    /**
     * All patient IDs listed in the PID-2, PID-3, and PID-4 segments of this message.
     */
    readonly patientIds?: pulumi.Input<pulumi.Input<inputs.healthcare.v1beta1.PatientIdArgs>[]>;
    readonly project: pulumi.Input<string>;
    /**
     * The parsed version of the raw message data schematized according to this store's schemas and type definitions.
     */
    readonly schematizedData?: pulumi.Input<inputs.healthcare.v1beta1.SchematizedDataArgs>;
    /**
     * The hospital that this message came from. MSH-4.
     */
    readonly sendFacility?: pulumi.Input<string>;
    /**
     * The datetime the sending application sent this message. MSH-7.
     */
    readonly sendTime?: pulumi.Input<string>;
}
