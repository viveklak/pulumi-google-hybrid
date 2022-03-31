// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "../../utilities";

// Export members:
export * from "./backup";
export * from "./database";
export * from "./getBackup";
export * from "./getDatabase";
export * from "./getInstance";
export * from "./getInstanceBackupIamPolicy";
export * from "./getInstanceDatabaseIamPolicy";
export * from "./getInstanceIamPolicy";
export * from "./getSession";
export * from "./instance";
export * from "./instanceBackupIamPolicy";
export * from "./instanceDatabaseIamPolicy";
export * from "./instanceIamPolicy";
export * from "./session";

// Export enums:
export * from "../../types/enums/spanner/v1";

// Import resources to register:
import { Backup } from "./backup";
import { Database } from "./database";
import { Instance } from "./instance";
import { InstanceBackupIamPolicy } from "./instanceBackupIamPolicy";
import { InstanceDatabaseIamPolicy } from "./instanceDatabaseIamPolicy";
import { InstanceIamPolicy } from "./instanceIamPolicy";
import { Session } from "./session";

const _module = {
    version: utilities.getVersion(),
    construct: (name: string, type: string, urn: string): pulumi.Resource => {
        switch (type) {
            case "google-native:spanner/v1:Backup":
                return new Backup(name, <any>undefined, { urn })
            case "google-native:spanner/v1:Database":
                return new Database(name, <any>undefined, { urn })
            case "google-native:spanner/v1:Instance":
                return new Instance(name, <any>undefined, { urn })
            case "google-native:spanner/v1:InstanceBackupIamPolicy":
                return new InstanceBackupIamPolicy(name, <any>undefined, { urn })
            case "google-native:spanner/v1:InstanceDatabaseIamPolicy":
                return new InstanceDatabaseIamPolicy(name, <any>undefined, { urn })
            case "google-native:spanner/v1:InstanceIamPolicy":
                return new InstanceIamPolicy(name, <any>undefined, { urn })
            case "google-native:spanner/v1:Session":
                return new Session(name, <any>undefined, { urn })
            default:
                throw new Error(`unknown resource type ${type}`);
        }
    },
};
pulumi.runtime.registerResourceModule("google-native", "spanner/v1", _module)
