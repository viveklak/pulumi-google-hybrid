// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import { input as inputs, output as outputs } from "../../types";
import * as utilities from "../../utilities";

/**
 * Create a FeedbackMessage object.
 */
export class DatasetAnnotatedDatasetFeedbackThreadFeedbackMessage extends pulumi.CustomResource {
    /**
     * Get an existing DatasetAnnotatedDatasetFeedbackThreadFeedbackMessage resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, opts?: pulumi.CustomResourceOptions): DatasetAnnotatedDatasetFeedbackThreadFeedbackMessage {
        return new DatasetAnnotatedDatasetFeedbackThreadFeedbackMessage(name, undefined as any, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'google-cloud:datalabeling/v1beta1:DatasetAnnotatedDatasetFeedbackThreadFeedbackMessage';

    /**
     * Returns true if the given object is an instance of DatasetAnnotatedDatasetFeedbackThreadFeedbackMessage.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is DatasetAnnotatedDatasetFeedbackThreadFeedbackMessage {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === DatasetAnnotatedDatasetFeedbackThreadFeedbackMessage.__pulumiType;
    }


    /**
     * Create a DatasetAnnotatedDatasetFeedbackThreadFeedbackMessage resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: DatasetAnnotatedDatasetFeedbackThreadFeedbackMessageArgs, opts?: pulumi.CustomResourceOptions) {
        let inputs: pulumi.Inputs = {};
        opts = opts || {};
        if (!opts.id) {
            if ((!args || args.annotatedDatasetsId === undefined) && !opts.urn) {
                throw new Error("Missing required property 'annotatedDatasetsId'");
            }
            if ((!args || args.datasetsId === undefined) && !opts.urn) {
                throw new Error("Missing required property 'datasetsId'");
            }
            if ((!args || args.feedbackMessagesId === undefined) && !opts.urn) {
                throw new Error("Missing required property 'feedbackMessagesId'");
            }
            if ((!args || args.feedbackThreadsId === undefined) && !opts.urn) {
                throw new Error("Missing required property 'feedbackThreadsId'");
            }
            if ((!args || args.projectsId === undefined) && !opts.urn) {
                throw new Error("Missing required property 'projectsId'");
            }
            inputs["annotatedDatasetsId"] = args ? args.annotatedDatasetsId : undefined;
            inputs["body"] = args ? args.body : undefined;
            inputs["createTime"] = args ? args.createTime : undefined;
            inputs["datasetsId"] = args ? args.datasetsId : undefined;
            inputs["feedbackMessagesId"] = args ? args.feedbackMessagesId : undefined;
            inputs["feedbackThreadsId"] = args ? args.feedbackThreadsId : undefined;
            inputs["image"] = args ? args.image : undefined;
            inputs["name"] = args ? args.name : undefined;
            inputs["operatorFeedbackMetadata"] = args ? args.operatorFeedbackMetadata : undefined;
            inputs["projectsId"] = args ? args.projectsId : undefined;
            inputs["requesterFeedbackMetadata"] = args ? args.requesterFeedbackMetadata : undefined;
        } else {
        }
        if (!opts.version) {
            opts = pulumi.mergeOptions(opts, { version: utilities.getVersion()});
        }
        super(DatasetAnnotatedDatasetFeedbackThreadFeedbackMessage.__pulumiType, name, inputs, opts);
    }
}

/**
 * The set of arguments for constructing a DatasetAnnotatedDatasetFeedbackThreadFeedbackMessage resource.
 */
export interface DatasetAnnotatedDatasetFeedbackThreadFeedbackMessageArgs {
    readonly annotatedDatasetsId: pulumi.Input<string>;
    /**
     * String content of the feedback. Maximum of 10000 characters.
     */
    readonly body?: pulumi.Input<string>;
    /**
     * Create time.
     */
    readonly createTime?: pulumi.Input<string>;
    readonly datasetsId: pulumi.Input<string>;
    readonly feedbackMessagesId: pulumi.Input<string>;
    readonly feedbackThreadsId: pulumi.Input<string>;
    /**
     * The image storing this feedback if the feedback is an image representing operator's comments.
     */
    readonly image?: pulumi.Input<string>;
    /**
     * Name of the feedback message in a feedback thread. Format: 'project/{project_id}/datasets/{dataset_id}/annotatedDatasets/{annotated_dataset_id}/feedbackThreads/{feedback_thread_id}/feedbackMessage/{feedback_message_id}'
     */
    readonly name?: pulumi.Input<string>;
    readonly operatorFeedbackMetadata?: pulumi.Input<inputs.datalabeling.v1beta1.GoogleCloudDatalabelingV1beta1OperatorFeedbackMetadata>;
    readonly projectsId: pulumi.Input<string>;
    readonly requesterFeedbackMetadata?: pulumi.Input<inputs.datalabeling.v1beta1.GoogleCloudDatalabelingV1beta1RequesterFeedbackMetadata>;
}
