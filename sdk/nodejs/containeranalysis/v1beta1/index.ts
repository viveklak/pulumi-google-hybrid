// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "../../utilities";

// Export members:
export * from "./note";
export * from "./noteIamPolicy";
export * from "./occurrence";
export * from "./occurrenceIamPolicy";

// Import resources to register:
import { Note } from "./note";
import { NoteIamPolicy } from "./noteIamPolicy";
import { Occurrence } from "./occurrence";
import { OccurrenceIamPolicy } from "./occurrenceIamPolicy";

const _module = {
    version: utilities.getVersion(),
    construct: (name: string, type: string, urn: string): pulumi.Resource => {
        switch (type) {
            case "google-cloud:containeranalysis/v1beta1:Note":
                return new Note(name, <any>undefined, { urn })
            case "google-cloud:containeranalysis/v1beta1:NoteIamPolicy":
                return new NoteIamPolicy(name, <any>undefined, { urn })
            case "google-cloud:containeranalysis/v1beta1:Occurrence":
                return new Occurrence(name, <any>undefined, { urn })
            case "google-cloud:containeranalysis/v1beta1:OccurrenceIamPolicy":
                return new OccurrenceIamPolicy(name, <any>undefined, { urn })
            default:
                throw new Error(`unknown resource type ${type}`);
        }
    },
};
pulumi.runtime.registerResourceModule("google-cloud", "containeranalysis/v1beta1", _module)
