// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "../../utilities";

// Export members:
export * from "./getQueue";
export * from "./getQueueIamPolicy";
export * from "./getTask";
export * from "./queue";
export * from "./queueIamPolicy";
export * from "./task";

// Export enums:
export * from "../../types/enums/cloudtasks/v2beta3";

// Import resources to register:
import { Queue } from "./queue";
import { QueueIamPolicy } from "./queueIamPolicy";
import { Task } from "./task";

const _module = {
    version: utilities.getVersion(),
    construct: (name: string, type: string, urn: string): pulumi.Resource => {
        switch (type) {
            case "google-native:cloudtasks/v2beta3:Queue":
                return new Queue(name, <any>undefined, { urn })
            case "google-native:cloudtasks/v2beta3:QueueIamPolicy":
                return new QueueIamPolicy(name, <any>undefined, { urn })
            case "google-native:cloudtasks/v2beta3:Task":
                return new Task(name, <any>undefined, { urn })
            default:
                throw new Error(`unknown resource type ${type}`);
        }
    },
};
pulumi.runtime.registerResourceModule("google-native", "cloudtasks/v2beta3", _module)
