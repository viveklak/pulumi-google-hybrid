// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import { input as inputs, output as outputs, enums } from "../../types";
import * as utilities from "../../utilities";

/**
 * Gets a model resource.
 */
export function getModel(args: GetModelArgs, opts?: pulumi.InvokeOptions): Promise<GetModelResult> {
    if (!opts) {
        opts = {}
    }

    opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
    return pulumi.runtime.invoke("google-native:firebaseml/v1beta2:getModel", {
        "modelId": args.modelId,
        "project": args.project,
    }, opts);
}

export interface GetModelArgs {
    modelId: string;
    project?: string;
}

export interface GetModelResult {
    /**
     * Lists operation ids associated with this model whose status is NOT done.
     */
    readonly activeOperations: outputs.firebaseml.v1beta2.OperationResponse[];
    /**
     * Timestamp when this model was created in Firebase ML.
     */
    readonly createTime: string;
    /**
     * The name of the model to create. The name can be up to 32 characters long and can consist only of ASCII Latin letters A-Z and a-z, underscores(_) and ASCII digits 0-9. It must start with a letter.
     */
    readonly displayName: string;
    /**
     * See RFC7232 https://tools.ietf.org/html/rfc7232#section-2.3
     */
    readonly etag: string;
    /**
     * The model_hash will change if a new file is available for download.
     */
    readonly modelHash: string;
    /**
     * The resource name of the Model. Model names have the form `projects/{project_id}/models/{model_id}` The name is ignored when creating a model.
     */
    readonly name: string;
    /**
     * State common to all model types. Includes publishing and validation information.
     */
    readonly state: outputs.firebaseml.v1beta2.ModelStateResponse;
    /**
     * User defined tags which can be used to group/filter models during listing
     */
    readonly tags: string[];
    /**
     * A TFLite Model
     */
    readonly tfliteModel: outputs.firebaseml.v1beta2.TfLiteModelResponse;
    /**
     * Timestamp when this model was updated in Firebase ML.
     */
    readonly updateTime: string;
}

export function getModelOutput(args: GetModelOutputArgs, opts?: pulumi.InvokeOptions): pulumi.Output<GetModelResult> {
    return pulumi.output(args).apply(a => getModel(a, opts))
}

export interface GetModelOutputArgs {
    modelId: pulumi.Input<string>;
    project?: pulumi.Input<string>;
}
