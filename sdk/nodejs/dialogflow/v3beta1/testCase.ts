// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import { input as inputs, output as outputs, enums } from "../../types";
import * as utilities from "../../utilities";

/**
 * Creates a test case for the given agent.
 * Note - this resource's API doesn't support deletion. When deleted, the resource will persist
 * on Google Cloud even though it will be deleted from Pulumi state.
 */
export class TestCase extends pulumi.CustomResource {
    /**
     * Get an existing TestCase resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, opts?: pulumi.CustomResourceOptions): TestCase {
        return new TestCase(name, undefined as any, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'google-native:dialogflow/v3beta1:TestCase';

    /**
     * Returns true if the given object is an instance of TestCase.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is TestCase {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === TestCase.__pulumiType;
    }

    /**
     * When the test was created.
     */
    public /*out*/ readonly creationTime!: pulumi.Output<string>;
    /**
     * The human-readable name of the test case, unique within the agent. Limit of 200 characters.
     */
    public readonly displayName!: pulumi.Output<string>;
    /**
     * The latest test result.
     */
    public readonly lastTestResult!: pulumi.Output<outputs.dialogflow.v3beta1.GoogleCloudDialogflowCxV3beta1TestCaseResultResponse>;
    /**
     * The unique identifier of the test case. TestCases.CreateTestCase will populate the name automatically. Otherwise use format: `projects//locations//agents/ /testCases/`.
     */
    public readonly name!: pulumi.Output<string>;
    /**
     * Additional freeform notes about the test case. Limit of 400 characters.
     */
    public readonly notes!: pulumi.Output<string>;
    /**
     * Tags are short descriptions that users may apply to test cases for organizational and filtering purposes. Each tag should start with "#" and has a limit of 30 characters.
     */
    public readonly tags!: pulumi.Output<string[]>;
    /**
     * The conversation turns uttered when the test case was created, in chronological order. These include the canonical set of agent utterances that should occur when the agent is working properly.
     */
    public readonly testCaseConversationTurns!: pulumi.Output<outputs.dialogflow.v3beta1.GoogleCloudDialogflowCxV3beta1ConversationTurnResponse[]>;
    /**
     * Config for the test case.
     */
    public readonly testConfig!: pulumi.Output<outputs.dialogflow.v3beta1.GoogleCloudDialogflowCxV3beta1TestConfigResponse>;

    /**
     * Create a TestCase resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: TestCaseArgs, opts?: pulumi.CustomResourceOptions) {
        let resourceInputs: pulumi.Inputs = {};
        opts = opts || {};
        if (!opts.id) {
            if ((!args || args.agentId === undefined) && !opts.urn) {
                throw new Error("Missing required property 'agentId'");
            }
            if ((!args || args.displayName === undefined) && !opts.urn) {
                throw new Error("Missing required property 'displayName'");
            }
            resourceInputs["agentId"] = args ? args.agentId : undefined;
            resourceInputs["displayName"] = args ? args.displayName : undefined;
            resourceInputs["lastTestResult"] = args ? args.lastTestResult : undefined;
            resourceInputs["location"] = args ? args.location : undefined;
            resourceInputs["name"] = args ? args.name : undefined;
            resourceInputs["notes"] = args ? args.notes : undefined;
            resourceInputs["project"] = args ? args.project : undefined;
            resourceInputs["tags"] = args ? args.tags : undefined;
            resourceInputs["testCaseConversationTurns"] = args ? args.testCaseConversationTurns : undefined;
            resourceInputs["testConfig"] = args ? args.testConfig : undefined;
            resourceInputs["creationTime"] = undefined /*out*/;
        } else {
            resourceInputs["creationTime"] = undefined /*out*/;
            resourceInputs["displayName"] = undefined /*out*/;
            resourceInputs["lastTestResult"] = undefined /*out*/;
            resourceInputs["name"] = undefined /*out*/;
            resourceInputs["notes"] = undefined /*out*/;
            resourceInputs["tags"] = undefined /*out*/;
            resourceInputs["testCaseConversationTurns"] = undefined /*out*/;
            resourceInputs["testConfig"] = undefined /*out*/;
        }
        opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
        super(TestCase.__pulumiType, name, resourceInputs, opts);
    }
}

/**
 * The set of arguments for constructing a TestCase resource.
 */
export interface TestCaseArgs {
    agentId: pulumi.Input<string>;
    /**
     * The human-readable name of the test case, unique within the agent. Limit of 200 characters.
     */
    displayName: pulumi.Input<string>;
    /**
     * The latest test result.
     */
    lastTestResult?: pulumi.Input<inputs.dialogflow.v3beta1.GoogleCloudDialogflowCxV3beta1TestCaseResultArgs>;
    location?: pulumi.Input<string>;
    /**
     * The unique identifier of the test case. TestCases.CreateTestCase will populate the name automatically. Otherwise use format: `projects//locations//agents/ /testCases/`.
     */
    name?: pulumi.Input<string>;
    /**
     * Additional freeform notes about the test case. Limit of 400 characters.
     */
    notes?: pulumi.Input<string>;
    project?: pulumi.Input<string>;
    /**
     * Tags are short descriptions that users may apply to test cases for organizational and filtering purposes. Each tag should start with "#" and has a limit of 30 characters.
     */
    tags?: pulumi.Input<pulumi.Input<string>[]>;
    /**
     * The conversation turns uttered when the test case was created, in chronological order. These include the canonical set of agent utterances that should occur when the agent is working properly.
     */
    testCaseConversationTurns?: pulumi.Input<pulumi.Input<inputs.dialogflow.v3beta1.GoogleCloudDialogflowCxV3beta1ConversationTurnArgs>[]>;
    /**
     * Config for the test case.
     */
    testConfig?: pulumi.Input<inputs.dialogflow.v3beta1.GoogleCloudDialogflowCxV3beta1TestConfigArgs>;
}
