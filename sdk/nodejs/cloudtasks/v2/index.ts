// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "../../utilities";

// Export members:
export * from "./queue";
export * from "./queueIamPolicy";
export * from "./queueTask";

// Import resources to register:
import { Queue } from "./queue";
import { QueueIamPolicy } from "./queueIamPolicy";
import { QueueTask } from "./queueTask";

const _module = {
    version: utilities.getVersion(),
    construct: (name: string, type: string, urn: string): pulumi.Resource => {
        switch (type) {
            case "google-cloud:cloudtasks/v2:Queue":
                return new Queue(name, <any>undefined, { urn })
            case "google-cloud:cloudtasks/v2:QueueIamPolicy":
                return new QueueIamPolicy(name, <any>undefined, { urn })
            case "google-cloud:cloudtasks/v2:QueueTask":
                return new QueueTask(name, <any>undefined, { urn })
            default:
                throw new Error(`unknown resource type ${type}`);
        }
    },
};
pulumi.runtime.registerResourceModule("google-cloud", "cloudtasks/v2", _module)