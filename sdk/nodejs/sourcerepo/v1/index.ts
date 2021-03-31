// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "../../utilities";

// Export members:
export * from "./repo";
export * from "./repoIamPolicy";

// Import resources to register:
import { Repo } from "./repo";
import { RepoIamPolicy } from "./repoIamPolicy";

const _module = {
    version: utilities.getVersion(),
    construct: (name: string, type: string, urn: string): pulumi.Resource => {
        switch (type) {
            case "google-cloud:sourcerepo/v1:Repo":
                return new Repo(name, <any>undefined, { urn })
            case "google-cloud:sourcerepo/v1:RepoIamPolicy":
                return new RepoIamPolicy(name, <any>undefined, { urn })
            default:
                throw new Error(`unknown resource type ${type}`);
        }
    },
};
pulumi.runtime.registerResourceModule("google-cloud", "sourcerepo/v1", _module)
