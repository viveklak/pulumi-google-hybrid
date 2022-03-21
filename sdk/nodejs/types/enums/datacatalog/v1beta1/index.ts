// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***


export const EntryType = {
    /**
     * Default unknown type.
     */
    EntryTypeUnspecified: "ENTRY_TYPE_UNSPECIFIED",
    /**
     * Output only. The type of entry that has a GoogleSQL schema, including logical views.
     */
    Table: "TABLE",
    /**
     * Output only. The type of models. https://cloud.google.com/bigquery-ml/docs/bigqueryml-intro
     */
    Model: "MODEL",
    /**
     * Output only. An entry type which is used for streaming entries. Example: Pub/Sub topic.
     */
    DataStream: "DATA_STREAM",
    /**
     * An entry type which is a set of files or objects. Example: Cloud Storage fileset.
     */
    Fileset: "FILESET",
} as const;

/**
 * The type of the entry. Only used for Entries with types in the EntryType enum.
 */
export type EntryType = (typeof EntryType)[keyof typeof EntryType];

export const TaxonomyActivatedPolicyTypesItem = {
    /**
     * Unspecified policy type.
     */
    PolicyTypeUnspecified: "POLICY_TYPE_UNSPECIFIED",
    /**
     * Fine grained access control policy, which enables access control on tagged resources.
     */
    FineGrainedAccessControl: "FINE_GRAINED_ACCESS_CONTROL",
} as const;

export type TaxonomyActivatedPolicyTypesItem = (typeof TaxonomyActivatedPolicyTypesItem)[keyof typeof TaxonomyActivatedPolicyTypesItem];
