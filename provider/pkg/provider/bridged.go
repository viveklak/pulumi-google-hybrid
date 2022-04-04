package provider

import (
	"fmt"
	google "github.com/hashicorp/terraform-provider-google-beta/google-beta"
	"github.com/pulumi/pulumi-google-hybrid/provider/pkg/version"
	"github.com/pulumi/pulumi-terraform-bridge/v3/pkg/tfbridge"
	shimv2 "github.com/pulumi/pulumi-terraform-bridge/v3/pkg/tfshim/sdk-v2"
	"github.com/pulumi/pulumi/sdk/v3/go/common/tokens"
	"path/filepath"
	"strings"
)

// all of the Google Cloud Platform token components used below.
const (
	// packages:
	gcpPackage = "google-hybrid"
	// modules; in general, we took naming inspiration from the Google Cloud SDK for Go:
	// https://github.com/GoogleCloudPlatform/google-cloud-go
	gcpAccessContextManager = "AccessContextManager" // Access Context Manager resources
	gcpActiveDirectory      = "ActiveDirectory"      // Active Directory resources
	// nolint:golint
	gcpApiGateway           = "ApiGateway"           // ApiGateway resources
	gcpApigee               = "Apigee"               // Apigee resources
	gcpAppEngine            = "AppEngine"            // AppEngine resources
	gcpArtifactRegistry     = "ArtifactRegistry"     // ArtifactRegistry resources
	gcpAssuredWorkloads     = "AssuredWorkloads"     // AssuredWorkloads resources
	gcpBigQuery             = "BigQuery"             // BigQuery resources
	gcpBigTable             = "BigTable"             // BitTable resources
	gcpBilling              = "Billing"              // Billing resources
	gcpBinaryAuthorization  = "BinaryAuthorization"  // Binary Authorization resources
	gcpCertificateAuthority = "CertificateAuthority" // CertificateAuthority resources
	gcpCloudAsset           = "CloudAsset"           // CloudAsset resources
	gcpCloudBuild           = "CloudBuild"           // CloudBuild resources
	gcpCloudFunctions       = "CloudFunctions"       // CloudFunction resources
	gcpCloudFunctionsV2     = "CloudFunctionsV2"     // CloudFunction (2nd Gen) resources
	gcpCloudIdentity        = "CloudIdentity"        // CloudIdentity resources
	gcpCloudRun             = "CloudRun"             // CloudRun resources
	gcpCloudScheduler       = "CloudScheduler"       // Cloud Scheduler resources
	gcpCloudTasks           = "CloudTasks"           // Cloud Tasks resources
	gcpComposer             = "Composer"             // Cloud Composer resources
	gcpCompute              = "Compute"              // Compute resources
	gcpContainerAnalysis    = "ContainerAnalysis"    // Container Analysis resources
	gcpDNS                  = "Dns"                  // DNS resources
	gcpDataCatalog          = "DataCatalog"          // Data Catalog resources
	gcpDataFlow             = "Dataflow"             // DataFlow resources
	gcpDataFusion           = "DataFusion"           // DataFusion resources
	gcpDataLoss             = "DataLoss"             // DataLoss resources
	gcpDataProc             = "Dataproc"             // DataProc resources
	gcpDatastore            = "Datastore"            // Datastore resources
	gcpDeploymentManager    = "DeploymentManager"    // DeploymentManager resources
	gcpDiagflow             = "Diagflow"             // Diagflow resources
	gcpEndPoints            = "Endpoints"            // End Point resources
	gcpEssentialContacts    = "EssentialContacts"    // Essential Contacts resources
	gcpEventarc             = "Eventarc"             // Eventarc
	gcpFilestore            = "Filestore"            // Filestore resources
	gcpFirebase             = "Firebase"             // Firebase resources
	gcpFirestore            = "Firestore"            // Firestore resources
	gcpFolder               = "Folder"               // Folder resources
	gcpGameServices         = "GameServices"         // Game Services resources
	gcpGkeHub               = "GkeHub"               // Gke Hub resources
	gcpHealthcare           = "Healthcare"           // Healthcare resources
	gcpIAM                  = "Iam"                  // IAM resources
	gcpIAP                  = "Iap"                  // IAP resources
	gcpIdentityPlatform     = "IdentityPlatform"     // IdentityPlatform resources
	gcpIot                  = "Iot"                  // Iot resources
	gcpKMS                  = "Kms"                  // KMS resources
	gcpKubernetes           = "Container"            // Kubernetes Engine resources
	gcpLogging              = "Logging"              // Logging resources
	gcpMachingLearning      = "ML"                   // Machine Learning
	gcpMemcache             = "Memcache"             // Memcache resources
	gcpMonitoring           = "Monitoring"           // Monitoring resources
	gcpNetworkConnectivity  = "NetworkConnectivity"  // Network Connectivity resources
	gcpNetworkManagement    = "NetworkManagement"    // Network Management resources
	gcpNetworkServices      = "NetworkServices"      // Network Services resources
	gcpNotebooks            = "Notebooks"            // Notebooks resources
	gcpOrganization         = "Organizations"        // Organization resources
	gcpOrgPolicy            = "OrgPolicy"            // Org Policy
	gcpOsConfig             = "OsConfig"             // OsConfig resources
	gcpOsLogin              = "OsLogin"              // OsLogin resources
	gcpProject              = "Projects"             // Project resources
	gcpPubSub               = "PubSub"               // PubSub resources
	gcpRecaptcha            = "Recaptcha"            //Recaptcha resources
	gcpRedis                = "Redis"                // Redis resources
	gcpResourceManager      = "ResourceManager"      // Resource Manager resources
	gcpRuntimeConfig        = "RuntimeConfig"        // Runtime Config resources
	gcpSecretManager        = "SecretManager"        // Secret Manager resources
	gcpServiceDirectory     = "ServiceDirectory"     // Service Directory resources
	gcpServiceNetworking    = "ServiceNetworking"    // Service Networking resources
	gcpSecurityCenter       = "SecurityCenter"       // Security Center
	gcpSQL                  = "Sql"                  // SQL resources
	gcpServiceAccount       = "ServiceAccount"       // Service Account resources
	gcpServiceUsage         = "ServiceUsage"         // Service Usage resources
	gcpSourceRepo           = "SourceRepo"           // Source Repo resources
	gcpSpanner              = "Spanner"              // Spanner Resources
	gcpStorage              = "Storage"              // Storage resources
	gcpTags                 = "Tags"                 // Tags
	gcpTPU                  = "Tpu"                  // Tensor Processing Units
	gcpVertex               = "Vertex"               // Vertex
	gcpVpcAccess            = "VpcAccess"            // VPC Access
	gcpWorkflows            = "Workflows"            // Workflows
)

var namespaceMap = map[string]string{
	"gcp": "Gcp",
}

// There is a bug that 'serviceAccount' is used instead of 'serviceaccount' in Node.js.
// We should eventually fix it and get rid of this map.
var specialNamesMap = map[string]string{
	"ServiceAccount": "serviceAccount",
}

// gcpMember manufactures a type token for the GCP package and the given module and type.  It automatically uses the GCP
// package and names the file by simply lower casing the resource's first character.
func gcpMember(moduleTitle string, mem string) tokens.ModuleMember {
	moduleName := strings.ToLower(moduleTitle)
	if value, exist := specialNamesMap[moduleTitle]; exist {
		moduleName = value
	}
	namespaceMap[moduleName] = moduleTitle
	// fn := string(unicode.ToLower(rune(mem[0]))) + mem[1:]
	token := moduleName + "/" + "classic"
	return tokens.ModuleMember(gcpPackage + ":" + token + ":" + mem)
}

// gcpType manufactures a type token for the GCP package and the given module and type.
func gcpType(mod string, typ string) tokens.Type {
	return tokens.Type(gcpMember(mod, typ))
}

// gcpDataSource manufactures a standard member given a module and resource name.
func gcpDataSource(mod string, res string) tokens.ModuleMember {
	return gcpMember(mod, res)
}

// googleHybridResource manufactures a standard resource token given a module and resource name.
func googleHybridResource(mod string, res string) tokens.Type {
	return gcpType(mod, res)
}

// managedByPulumi is a default used for some managed resources, in the absence of something more meaningful.
var managedByPulumi = &tfbridge.DefaultInfo{Value: "Managed by Pulumi"}

func boolRef(b bool) *bool {
	return &b
}

// Provider returns additional overlaid schema and metadata associated with the gcp package.
func Provider() tfbridge.ProviderInfo {
	p := shimv2.NewProvider(google.Provider())
	prov := tfbridge.ProviderInfo{
		P:              p,
		Name:           "google-beta",
		ResourcePrefix: "google",
		GitHubOrg:      "hashicorp",
		Description:    "A Pulumi package for creating and managing Google Cloud Platform resources.",
		Keywords:       []string{"pulumi", "gcp"},
		License:        "Apache-2.0",
		Homepage:       "https://pulumi.io",
		Repository:     "https://github.com/pulumi/pulumi-gcp",
		Config: map[string]*tfbridge.SchemaInfo{
			"project": {
				Default: &tfbridge.DefaultInfo{
					EnvVars: []string{
						"GOOGLE_PROJECT",
						"GOOGLE_CLOUD_PROJECT",
						"GCLOUD_PROJECT",
						"CLOUDSDK_CORE_PROJECT",
					},
				},
			},
			"region": {
				Default: &tfbridge.DefaultInfo{
					EnvVars: []string{
						"GOOGLE_REGION",
						"GCLOUD_REGION",
						"CLOUDSDK_COMPUTE_REGION",
					},
				},
			},
			"zone": {
				Default: &tfbridge.DefaultInfo{
					EnvVars: []string{
						"GOOGLE_ZONE",
						"GCLOUD_ZONE",
						"CLOUDSDK_COMPUTE_ZONE",
					},
				},
			},
		},
		Resources: map[string]*tfbridge.ResourceInfo{
			//// Access Context Manager
			//"google_access_context_manager_access_level": {
			//	Tok: googleHybridResource(gcpAccessContextManager, "AccessLevel"),
			//},
			//"google_access_context_manager_access_policy": {
			//	Tok: googleHybridResource(gcpAccessContextManager, "AccessPolicy"),
			//},
			//"google_access_context_manager_service_perimeter": {
			//	Tok: googleHybridResource(gcpAccessContextManager, "ServicePerimeter"),
			//},
			//"google_access_context_manager_service_perimeter_resource": {
			//	Tok: googleHybridResource(gcpAccessContextManager, "ServicePerimeterResource"),
			//},
			//"google_access_context_manager_service_perimeters": {
			//	Tok: googleHybridResource(gcpAccessContextManager, "ServicePerimeters"),
			//	Fields: map[string]*tfbridge.SchemaInfo{
			//		"service_perimeters": {
			//			CSharpName: "ServicePerimeterDetails",
			//		},
			//	},
			//},
			//"google_access_context_manager_access_levels": {
			//	Tok: googleHybridResource(gcpAccessContextManager, "AccessLevels"),
			//	Fields: map[string]*tfbridge.SchemaInfo{
			//		"access_levels": {
			//			CSharpName: "AccessLevelDetails",
			//		},
			//	},
			//},
			//"google_access_context_manager_access_level_condition": {
			//	Tok: googleHybridResource(gcpAccessContextManager, "AccessLevelCondition"),
			//},
			//"google_access_context_manager_gcp_user_access_binding": {
			//	Tok: googleHybridResource(gcpAccessContextManager, "GcpUserAccessBinding"),
			//},
			//
			//// AppEngine
			//"google_app_engine_application": {Tok: googleHybridResource(gcpAppEngine, "Application")},
			//"google_app_engine_firewall_rule": {
			//	Tok: googleHybridResource(gcpAppEngine, "FirewallRule"),
			//	Docs: &tfbridge.DocInfo{
			//		Source: "appengine_firewall_rule.html.markdown",
			//	},
			//},
			//"google_app_engine_standard_app_version": {Tok: googleHybridResource(gcpAppEngine, "StandardAppVersion")},
			//"google_app_engine_domain_mapping":       {Tok: googleHybridResource(gcpAppEngine, "DomainMapping")},
			//"google_app_engine_application_url_dispatch_rules": {
			//	Tok: googleHybridResource(gcpAppEngine, "ApplicationUrlDispatchRules"),
			//},
			//"google_app_engine_service_split_traffic":    {Tok: googleHybridResource(gcpAppEngine, "EngineSplitTraffic")},
			//"google_app_engine_flexible_app_version":     {Tok: googleHybridResource(gcpAppEngine, "FlexibleAppVersion")},
			//"google_app_engine_service_network_settings": {Tok: googleHybridResource(gcpAppEngine, "ServiceNetworkSettings")},
			//
			//// AssuredWorkloads
			//"google_assured_workloads_workload": {Tok: googleHybridResource(gcpAssuredWorkloads, "Workload")},
			//
			//// BigQuery
			//"google_bigquery_dataset":              {Tok: googleHybridResource(gcpBigQuery, "Dataset")},
			//"google_bigquery_table":                {Tok: googleHybridResource(gcpBigQuery, "Table")},
			//"google_bigquery_data_transfer_config": {Tok: googleHybridResource(gcpBigQuery, "DataTransferConfig")},
			//"google_bigquery_reservation":          {Tok: googleHybridResource(gcpBigQuery, "Reservation")},
			//"google_bigtable_app_profile":          {Tok: googleHybridResource(gcpBigQuery, "AppProfile")},
			//"google_bigquery_dataset_access": {
			//	Tok: googleHybridResource(gcpBigQuery, "DatasetAccess"),
			//	// The upstream provider has nested attributes, both called "dataset", which causes a panic in the
			//	// bridge due to the duplicated names. In order to resolve the panic (and also to clarify the meaning of
			//	// the field), we use the name "authorizedDataset", which is derived from the title of the example code
			//	// in the upstream docs:
			//	// https://registry.terraform.io/providers/hashicorp/google-beta/latest/docs/resources/bigquery_dataset_access#example-usage---bigquery-dataset-access-authorized-dataset
			//	Fields: map[string]*tfbridge.SchemaInfo{
			//		"dataset": {
			//			Name: "authorizedDataset",
			//		},
			//	},
			//},
			//"google_bigquery_job":        {Tok: googleHybridResource(gcpBigQuery, "Job")},
			//"google_bigquery_connection": {Tok: googleHybridResource(gcpBigQuery, "Connection")},
			//"google_bigquery_dataset_iam_binding": {
			//	Tok: googleHybridResource(gcpBigQuery, "DatasetIamBinding"),
			//	Docs: &tfbridge.DocInfo{
			//		Source: "bigquery_dataset_iam.html.markdown",
			//	},
			//},
			//"google_bigquery_dataset_iam_member": {
			//	Tok: googleHybridResource(gcpBigQuery, "DatasetIamMember"),
			//	Docs: &tfbridge.DocInfo{
			//		Source: "bigquery_dataset_iam.html.markdown",
			//	},
			//},
			//"google_bigquery_dataset_iam_policy": {
			//	Tok: googleHybridResource(gcpBigQuery, "DatasetIamPolicy"),
			//	Docs: &tfbridge.DocInfo{
			//		Source: "bigquery_dataset_iam.html.markdown",
			//	},
			//},
			//"google_bigquery_table_iam_policy": {
			//	Tok: googleHybridResource(gcpBigQuery, "IamPolicy"),
			//	Docs: &tfbridge.DocInfo{
			//		Source: "bigquery_table_iam.html.markdown",
			//	},
			//},
			//"google_bigquery_table_iam_binding": {
			//	Tok: googleHybridResource(gcpBigQuery, "IamBinding"),
			//	Docs: &tfbridge.DocInfo{
			//		Source: "bigquery_table_iam.html.markdown",
			//	},
			//},
			//"google_bigquery_table_iam_member": {
			//	Tok: googleHybridResource(gcpBigQuery, "IamMember"),
			//	Docs: &tfbridge.DocInfo{
			//		Source: "bigquery_table_iam.html.markdown",
			//	},
			//},
			//"google_bigquery_routine":                {Tok: googleHybridResource(gcpBigQuery, "Routine")},
			//"google_bigquery_reservation_assignment": {Tok: googleHybridResource(gcpBigQuery, "ReservationAssignment")},
			//
			//// BigTable
			//"google_bigtable_instance": {Tok: googleHybridResource(gcpBigTable, "Instance")},
			//"google_bigtable_table":    {Tok: googleHybridResource(gcpBigTable, "Table")},
			//"google_bigtable_instance_iam_binding": {
			//	Tok: googleHybridResource(gcpBigTable, "InstanceIamBinding"),
			//	Docs: &tfbridge.DocInfo{
			//		Source: "bigtable_instance_iam.html.markdown",
			//	},
			//},
			//"google_bigtable_instance_iam_member": {
			//	Tok: googleHybridResource(gcpBigTable, "InstanceIamMember"),
			//	Docs: &tfbridge.DocInfo{
			//		Source: "bigtable_instance_iam.html.markdown",
			//	},
			//},
			//"google_bigtable_instance_iam_policy": {
			//	Tok: googleHybridResource(gcpBigTable, "InstanceIamPolicy"),
			//	Docs: &tfbridge.DocInfo{
			//		Source: "bigtable_instance_iam.html.markdown",
			//	},
			//},
			//"google_bigtable_gc_policy": {
			//	Tok: googleHybridResource(gcpBigTable, "GCPolicy"),
			//	Docs: &tfbridge.DocInfo{
			//		Source: "bigtable_gc_policy.html.markdown",
			//	},
			//},
			//"google_bigtable_table_iam_member": {
			//	Tok: googleHybridResource(gcpBigTable, "TableIamMember"),
			//	Docs: &tfbridge.DocInfo{
			//		Source: "bigtable_table_iam.html.markdown",
			//	},
			//},
			//"google_bigtable_table_iam_binding": {
			//	Tok: googleHybridResource(gcpBigTable, "TableIamBinding"),
			//	Docs: &tfbridge.DocInfo{
			//		Source: "bigtable_table_iam.html.markdown",
			//	},
			//},
			//"google_bigtable_table_iam_policy": {
			//	Tok: googleHybridResource(gcpBigTable, "TableIamPolicy"),
			//	Docs: &tfbridge.DocInfo{
			//		Source: "bigtable_table_iam.html.markdown",
			//	},
			//},
			//
			//// Billing
			//"google_billing_account_iam_binding": {
			//	Tok: googleHybridResource(gcpBilling, "AccountIamBinding"),
			//	Docs: &tfbridge.DocInfo{
			//		Source: "billing_account_iam.html.markdown",
			//	},
			//},
			//"google_billing_account_iam_member": {
			//	Tok: googleHybridResource(gcpBilling, "AccountIamMember"),
			//	Docs: &tfbridge.DocInfo{
			//		Source: "billing_account_iam.html.markdown",
			//	},
			//},
			//"google_billing_account_iam_policy": {
			//	Tok: googleHybridResource(gcpBilling, "AccountIamPolicy"),
			//	Docs: &tfbridge.DocInfo{
			//		Source: "billing_account_iam.html.markdown",
			//	},
			//},
			//"google_billing_budget":     {Tok: googleHybridResource(gcpBilling, "Budget")},
			//"google_billing_subaccount": {Tok: googleHybridResource(gcpBilling, "SubAccount")},
			//
			//// Binary Authorization
			//"google_binary_authorization_attestor": {
			//	Tok: googleHybridResource(gcpBinaryAuthorization, "Attestor"),
			//	Docs: &tfbridge.DocInfo{
			//		Source: "binaryauthorization_attestor.html.markdown",
			//	},
			//},
			//"google_binary_authorization_policy": {
			//	Tok: googleHybridResource(gcpBinaryAuthorization, "Policy"),
			//	Docs: &tfbridge.DocInfo{
			//		Source: "binaryauthorization_policy.html.markdown",
			//	},
			//},
			//"google_binary_authorization_attestor_iam_binding": {
			//	Tok: googleHybridResource(gcpBinaryAuthorization, "AttestorIamBinding"),
			//	Docs: &tfbridge.DocInfo{
			//		Source: "binary_authorization_attestor_iam.html.markdown",
			//	},
			//},
			//"google_binary_authorization_attestor_iam_member": {
			//	Tok: googleHybridResource(gcpBinaryAuthorization, "AttestorIamMember"),
			//	Docs: &tfbridge.DocInfo{
			//		Source: "binary_authorization_attestor_iam.html.markdown",
			//	},
			//},
			//"google_binary_authorization_attestor_iam_policy": {
			//	Tok: googleHybridResource(gcpBinaryAuthorization, "AttestorIamPolicy"),
			//	Docs: &tfbridge.DocInfo{
			//		Source: "binary_authorization_attestor_iam.html.markdown",
			//	},
			//},
			//
			//// Cloud Build
			//"google_cloudbuild_trigger": {
			//	Tok: googleHybridResource(gcpCloudBuild, "Trigger"),
			//	Docs: &tfbridge.DocInfo{
			//		Source: "cloud_build_trigger.html.markdown",
			//	},
			//},
			//"google_cloudbuild_worker_pool": {Tok: googleHybridResource(gcpCloudBuild, "WorkerPool")},
			//
			//// Cloud Functions
			//"google_cloudfunctions_function": {
			//	Tok: googleHybridResource(gcpCloudFunctions, "Function"),
			//	Fields: map[string]*tfbridge.SchemaInfo{
			//		// Name must start with a letter followed by up to 62 letters, numbers, or
			//		// hyphens, and cannot end with a hyphen
			//		"name": tfbridge.AutoName("name", 63, "-"),
			//	},
			//},
			//"google_cloudfunctions_function_iam_binding": {
			//	Tok: googleHybridResource(gcpCloudFunctions, "FunctionIamBinding"),
			//	Docs: &tfbridge.DocInfo{
			//		Source: "cloudfunctions_function_iam.html.markdown",
			//	},
			//},
			//"google_cloudfunctions_function_iam_member": {
			//	Tok: googleHybridResource(gcpCloudFunctions, "FunctionIamMember"),
			//	Docs: &tfbridge.DocInfo{
			//		Source: "cloudfunctions_function_iam.html.markdown",
			//	},
			//},
			//"google_cloudfunctions_function_iam_policy": {
			//	Tok: googleHybridResource(gcpCloudFunctions, "FunctionIamPolicy"),
			//	Docs: &tfbridge.DocInfo{
			//		Source: "cloudfunctions_function_iam.html.markdown",
			//	},
			//},
			//
			//// Cloud Functions (2nd gen)
			//"google_cloudfunctions2_function": {Tok: googleHybridResource(gcpCloudFunctionsV2, "Function")},
			//
			//// Cloud Scheduler
			//"google_cloud_scheduler_job": {Tok: googleHybridResource(gcpCloudScheduler, "Job")},
			//
			//// Composer
			//"google_composer_environment": {Tok: googleHybridResource(gcpComposer, "Environment")},
			//
			//// Core functions
			//"google_folder": {
			//	Tok: googleHybridResource(gcpOrganization, "Folder"),
			//	Docs: &tfbridge.DocInfo{
			//		Source: "google_folder.html.markdown",
			//	},
			//},
			//"google_folder_iam_binding": {
			//	Tok: googleHybridResource(gcpFolder, "IAMBinding"),
			//	Docs: &tfbridge.DocInfo{
			//		Source: "google_folder_iam.html.markdown",
			//	},
			//},
			//"google_folder_iam_member": {
			//	Tok: googleHybridResource(gcpFolder, "IAMMember"),
			//	Docs: &tfbridge.DocInfo{
			//		Source: "google_folder_iam.html.markdown",
			//	},
			//},
			//"google_folder_iam_policy": {
			//	Tok: googleHybridResource(gcpFolder, "IAMPolicy"),
			//	Docs: &tfbridge.DocInfo{
			//		Source: "google_folder_iam.html.markdown",
			//	},
			//},
			//"google_folder_organization_policy": {
			//	Tok: googleHybridResource(gcpFolder, "OrganizationPolicy"),
			//	Docs: &tfbridge.DocInfo{
			//		Source: "google_folder_organization_policy.html.markdown",
			//	},
			//},
			//"google_folder_iam_audit_config": {
			//	Tok: googleHybridResource(gcpFolder, "IamAuditConfig"),
			//	Docs: &tfbridge.DocInfo{
			//		Source: "google_folder_iam.html.markdown",
			//	},
			//},
			//"google_folder_access_approval_settings": {Tok: googleHybridResource(gcpFolder, "AccessApprovalSettings")},
			//
			//"google_organization_policy": {
			//	Tok: googleHybridResource(gcpOrganization, "Policy"),
			//	Docs: &tfbridge.DocInfo{
			//		Source: "google_organization_policy.html.markdown",
			//	},
			//},
			//"google_organization_iam_binding": {
			//	Tok: googleHybridResource(gcpOrganization, "IAMBinding"),
			//	Docs: &tfbridge.DocInfo{
			//		Source: "google_organization_iam.html.markdown",
			//	},
			//},
			//"google_organization_iam_custom_role": {
			//	Tok: googleHybridResource(gcpOrganization, "IAMCustomRole"),
			//	Docs: &tfbridge.DocInfo{
			//		Source: "google_organization_iam_custom_role.html.markdown",
			//	},
			//},
			//"google_organization_iam_member": {
			//	Tok: googleHybridResource(gcpOrganization, "IAMMember"),
			//	Docs: &tfbridge.DocInfo{
			//		Source: "google_organization_iam.html.markdown",
			//	},
			//},
			//"google_organization_iam_policy": {
			//	Tok: googleHybridResource(gcpOrganization, "IAMPolicy"),
			//	Docs: &tfbridge.DocInfo{
			//		Source: "google_organization_iam.html.markdown",
			//	},
			//},
			//"google_organization_iam_audit_config": {
			//	Tok: googleHybridResource(gcpOrganization, "IamAuditConfig"),
			//	Docs: &tfbridge.DocInfo{
			//		Source: "google_organization_iam.html.markdown",
			//	},
			//},
			//"google_organization_access_approval_settings": {
			//	Tok: googleHybridResource(gcpOrganization, "AccessApprovalSettings"),
			//},
			"google_project": {
				Tok: googleHybridResource(gcpOrganization, "Project"),
				Docs: &tfbridge.DocInfo{
					Source: "google_project.html.markdown",
				},
			},

			"google_project_iam_audit_config": {
				Tok: googleHybridResource(gcpProject, "IAMAuditConfig"),
				Docs: &tfbridge.DocInfo{
					Source: "google_project_iam.html.markdown",
				},
			},
			"google_project_iam_binding": {
				Tok: googleHybridResource(gcpProject, "IAMBinding"),
				Docs: &tfbridge.DocInfo{
					Source: "google_project_iam.html.markdown",
				},
			},
			"google_project_iam_custom_role": {
				Tok: googleHybridResource(gcpProject, "IAMCustomRole"),
				Docs: &tfbridge.DocInfo{
					Source: "google_project_iam_custom_role.html.markdown",
				},
			},
			"google_project_iam_member": {
				Tok: googleHybridResource(gcpProject, "IAMMember"),
				Docs: &tfbridge.DocInfo{
					Source: "google_project_iam.html.markdown",
				},
			},
			"google_project_iam_policy": {
				Tok: googleHybridResource(gcpProject, "IAMPolicy"),
				Docs: &tfbridge.DocInfo{
					Source: "google_project_iam.html.markdown",
				},
			},
			"google_project_organization_policy": {
				Tok: googleHybridResource(gcpProject, "OrganizationPolicy"),
				Docs: &tfbridge.DocInfo{
					Source: "google_project_organization_policy.html.markdown",
				},
			},
			"google_project_service": {
				Tok: googleHybridResource(gcpProject, "Service"),
				Docs: &tfbridge.DocInfo{
					Source: "google_project_service.html.markdown",
				},
				Fields: map[string]*tfbridge.SchemaInfo{
					"service": {
						CSharpName: "ServiceName",
					},
				},
			},
			"google_project_service_identity": {
				Tok: googleHybridResource(gcpProject, "ServiceIdentity"),
			},
			"google_project_default_service_accounts": {Tok: googleHybridResource(gcpProject, "DefaultServiceAccounts")},
			"google_project_usage_export_bucket": {
				Tok: googleHybridResource(gcpProject, "UsageExportBucket"),
				Docs: &tfbridge.DocInfo{
					Source: "google_project.html.markdown",
				},
			},
			"google_project_access_approval_settings": {
				Tok: googleHybridResource(gcpProject, "AccessApprovalSettings"),
			},
			//// This resource is in the root namespace in the TF provider at the time of writing. The GCP SDK does not
			//// give an obvious namespace choice either. Since an API key authenticates an application, we put it under
			//// the gcpProject module:
			//"google_apikeys_key": {Tok: googleHybridResource(gcpProject, "ApiKey")},
			//
			//"google_service_account": {
			//	Tok: googleHybridResource(gcpServiceAccount, "Account"),
			//	Docs: &tfbridge.DocInfo{
			//		Source: "google_service_account.html.markdown",
			//	},
			//},
			//"google_service_account_iam_binding": {
			//	Tok: googleHybridResource(gcpServiceAccount, "IAMBinding"),
			//	Docs: &tfbridge.DocInfo{
			//		Source: "google_service_account_iam.html.markdown",
			//	},
			//},
			//"google_service_account_iam_member": {
			//	Tok: googleHybridResource(gcpServiceAccount, "IAMMember"),
			//	Docs: &tfbridge.DocInfo{
			//		Source: "google_service_account_iam.html.markdown",
			//	},
			//},
			//"google_service_account_iam_policy": {
			//	Tok: googleHybridResource(gcpServiceAccount, "IAMPolicy"),
			//	Docs: &tfbridge.DocInfo{
			//		Source: "google_service_account_iam.html.markdown",
			//	},
			//},
			//"google_service_account_key": {
			//	Tok: googleHybridResource(gcpServiceAccount, "Key"),
			//	Docs: &tfbridge.DocInfo{
			//		Source: "google_service_account_key.html.markdown",
			//	},
			//},
			//
			//// Service Usage
			//"google_service_usage_consumer_quota_override": {
			//	Tok: googleHybridResource(gcpServiceUsage, "ConsumerQuotaOverride"),
			//},
			//
			//// Compute
			//"google_compute_address": {
			//	Tok: googleHybridResource(gcpCompute, "Address"),
			//	Docs: &tfbridge.DocInfo{
			//		Source: "compute_address.html.markdown",
			//	},
			//	Fields: map[string]*tfbridge.SchemaInfo{
			//		"address": {
			//			CSharpName: "IPAddress",
			//		},
			//	},
			//},
			//"google_compute_attached_disk": {
			//	Tok: googleHybridResource(gcpCompute, "AttachedDisk"),
			//	Docs: &tfbridge.DocInfo{
			//		Source: "compute_attached_disk.html.markdown",
			//	},
			//},
			//"google_compute_autoscaler": {
			//	Tok: googleHybridResource(gcpCompute, "Autoscalar"),
			//	Docs: &tfbridge.DocInfo{
			//		Source: "compute_autoscaler.html.markdown",
			//	},
			//},
			//"google_compute_backend_bucket": {
			//	Tok: googleHybridResource(gcpCompute, "BackendBucket"),
			//	Docs: &tfbridge.DocInfo{
			//		Source: "compute_backend_bucket.html.markdown",
			//	},
			//},
			//"google_compute_backend_bucket_signed_url_key": {
			//	Tok: googleHybridResource(gcpCompute, "BackendBucketSignedUrlKey"),
			//	Docs: &tfbridge.DocInfo{
			//		Source: "compute_backend_bucket_signed_url_key.html.markdown",
			//	},
			//},
			//"google_compute_backend_service": {
			//	Tok: googleHybridResource(gcpCompute, "BackendService"),
			//	Docs: &tfbridge.DocInfo{
			//		Source: "compute_backend_service.html.markdown",
			//	},
			//},
			//"google_compute_backend_service_signed_url_key": {
			//	Tok: googleHybridResource(gcpCompute, "BackendServiceSignedUrlKey"),
			//	Docs: &tfbridge.DocInfo{
			//		Source: "compute_backend_service_signed_url_key.html.markdown",
			//	},
			//},
			//"google_compute_disk":                          {Tok: googleHybridResource(gcpCompute, "Disk")},
			//"google_compute_firewall":                      {Tok: googleHybridResource(gcpCompute, "Firewall")},
			//"google_compute_firewall_policy":               {Tok: googleHybridResource(gcpCompute, "FirewallPolicy")},
			//"google_compute_firewall_policy_rule":          {Tok: googleHybridResource(gcpCompute, "FirewallPolicyRule")},
			//"google_compute_firewall_policy_association":   {Tok: googleHybridResource(gcpCompute, "FirewallPolicyAssociation")},
			//"google_compute_forwarding_rule":               {Tok: googleHybridResource(gcpCompute, "ForwardingRule")},
			//"google_compute_external_vpn_gateway":          {Tok: googleHybridResource(gcpCompute, "ExternalVpnGateway")},
			//"google_compute_global_address":                {Tok: googleHybridResource(gcpCompute, "GlobalAddress")},
			//"google_compute_global_forwarding_rule":        {Tok: googleHybridResource(gcpCompute, "GlobalForwardingRule")},
			//"google_compute_global_network_endpoint":       {Tok: googleHybridResource(gcpCompute, "GlobalNetworkEndpoint")},
			//"google_compute_global_network_endpoint_group": {Tok: googleHybridResource(gcpCompute, "GlobalNetworkEndpointGroup")},
			//"google_compute_ha_vpn_gateway":                {Tok: googleHybridResource(gcpCompute, "HaVpnGateway")},
			//"google_compute_health_check":                  {Tok: googleHybridResource(gcpCompute, "HealthCheck")},
			//"google_compute_http_health_check":             {Tok: googleHybridResource(gcpCompute, "HttpHealthCheck")},
			//"google_compute_https_health_check":            {Tok: googleHybridResource(gcpCompute, "HttpsHealthCheck")},
			//"google_compute_image":                         {Tok: googleHybridResource(gcpCompute, "Image")},
			//"google_compute_instance":                      {Tok: googleHybridResource(gcpCompute, "Instance")},
			//"google_compute_instance_from_template":        {Tok: googleHybridResource(gcpCompute, "InstanceFromTemplate")},
			//"google_compute_instance_group":                {Tok: googleHybridResource(gcpCompute, "InstanceGroup")},
			//"google_compute_instance_group_manager":        {Tok: googleHybridResource(gcpCompute, "InstanceGroupManager")},
			//"google_compute_instance_from_machine_image":   {Tok: googleHybridResource(gcpCompute, "InstanceFromMachineImage")},
			//"google_compute_instance_iam_binding": {
			//	Tok: googleHybridResource(gcpCompute, "InstanceIAMBinding"),
			//	Docs: &tfbridge.DocInfo{
			//		Source: "compute_instance_iam.html.markdown",
			//	},
			//},
			//"google_compute_instance_iam_member": {
			//	Tok: googleHybridResource(gcpCompute, "InstanceIAMMember"),
			//	Docs: &tfbridge.DocInfo{
			//		Source: "compute_instance_iam.html.markdown",
			//	},
			//},
			//"google_compute_instance_iam_policy": {
			//	Tok: googleHybridResource(gcpCompute, "InstanceIAMPolicy"),
			//	Docs: &tfbridge.DocInfo{
			//		Source: "compute_instance_iam.html.markdown",
			//	},
			//},
			//"google_compute_instance_template":       {Tok: googleHybridResource(gcpCompute, "InstanceTemplate")},
			//"google_compute_interconnect_attachment": {Tok: googleHybridResource(gcpCompute, "InterconnectAttachment")},
			//"google_compute_machine_image_iam_binding": {
			//	Tok: googleHybridResource(gcpCompute, "MachineImageIamBinding"),
			//	Docs: &tfbridge.DocInfo{
			//		Source: "compute_machine_image_iam.html.markdown",
			//	},
			//},
			//"google_compute_machine_image_iam_member": {
			//	Tok: googleHybridResource(gcpCompute, "MachineImageIamMember"),
			//	Docs: &tfbridge.DocInfo{
			//		Source: "compute_machine_image_iam.html.markdown",
			//	},
			//},
			//"google_compute_machine_image_iam_policy": {
			//	Tok: googleHybridResource(gcpCompute, "MachineImageIamPolicy"),
			//	Docs: &tfbridge.DocInfo{
			//		Source: "compute_machine_image_iam.html.markdown",
			//	},
			//},
			//"google_compute_node_group":       {Tok: googleHybridResource(gcpCompute, "NodeGroup")},
			//"google_compute_node_template":    {Tok: googleHybridResource(gcpCompute, "NodeTemplate")},
			//"google_compute_network_endpoint": {Tok: googleHybridResource(gcpCompute, "NetworkEndpoint")},
			//"google_compute_network_endpoint_group": {
			//	Tok: googleHybridResource(gcpCompute, "NetworkEndpointGroup"),
			//	Docs: &tfbridge.DocInfo{
			//		Source: "compute_network_endpoint_group.html.markdown",
			//	},
			//},
			//"google_compute_network_peering":               {Tok: googleHybridResource(gcpCompute, "NetworkPeering")},
			//"google_compute_network_peering_routes_config": {Tok: googleHybridResource(gcpCompute, "NetworkPeeringRoutesConfig")},
			//"google_compute_network":                       {Tok: googleHybridResource(gcpCompute, "Network")},
			//"google_compute_project_default_network_tier":  {Tok: googleHybridResource(gcpCompute, "ProjectDefaultNetworkTier")},
			//"google_compute_project_metadata":              {Tok: googleHybridResource(gcpCompute, "ProjectMetadata")},
			//"google_compute_project_metadata_item":         {Tok: googleHybridResource(gcpCompute, "ProjectMetadataItem")},
			//"google_compute_region_autoscaler":             {Tok: googleHybridResource(gcpCompute, "RegionAutoscaler")},
			//"google_compute_region_backend_service":        {Tok: googleHybridResource(gcpCompute, "RegionBackendService")},
			//"google_compute_region_disk":                   {Tok: googleHybridResource(gcpCompute, "RegionDisk")},
			//"google_compute_region_instance_group_manager": {Tok: googleHybridResource(gcpCompute, "RegionInstanceGroupManager")},
			//"google_compute_region_ssl_certificate":        {Tok: googleHybridResource(gcpCompute, "RegionSslCertificate")},
			//"google_compute_region_target_http_proxy":      {Tok: googleHybridResource(gcpCompute, "RegionTargetHttpProxy")},
			//"google_compute_region_target_https_proxy":     {Tok: googleHybridResource(gcpCompute, "RegionTargetHttpsProxy")},
			//"google_compute_region_network_endpoint_group": {Tok: googleHybridResource(gcpCompute, "RegionNetworkEndpointGroup")},
			//"google_compute_resource_policy":               {Tok: googleHybridResource(gcpCompute, "ResourcePolicy")},
			//"google_compute_route":                         {Tok: googleHybridResource(gcpCompute, "Route")},
			//"google_compute_router":                        {Tok: googleHybridResource(gcpCompute, "Router")},
			//"google_compute_router_interface":              {Tok: googleHybridResource(gcpCompute, "RouterInterface")},
			//"google_compute_router_nat":                    {Tok: googleHybridResource(gcpCompute, "RouterNat")},
			//"google_compute_router_peer": {
			//	Tok: googleHybridResource(gcpCompute, "RouterPeer"),
			//	Docs: &tfbridge.DocInfo{
			//		Source: "compute_router_bgp_peer.html.markdown",
			//	},
			//},
			//"google_compute_security_policy":            {Tok: googleHybridResource(gcpCompute, "SecurityPolicy")},
			//"google_security_scanner_scan_config":       {Tok: googleHybridResource(gcpCompute, "SecurityScanConfig")},
			//"google_compute_shared_vpc_host_project":    {Tok: googleHybridResource(gcpCompute, "SharedVPCHostProject")},
			//"google_compute_shared_vpc_service_project": {Tok: googleHybridResource(gcpCompute, "SharedVPCServiceProject")},
			//"google_compute_snapshot":                   {Tok: googleHybridResource(gcpCompute, "Snapshot")},
			//"google_compute_ssl_certificate":            {Tok: googleHybridResource(gcpCompute, "SSLCertificate")},
			//"google_compute_ssl_policy":                 {Tok: googleHybridResource(gcpCompute, "SSLPolicy")},
			//"google_compute_subnetwork":                 {Tok: googleHybridResource(gcpCompute, "Subnetwork")},
			//"google_compute_subnetwork_iam_binding": {
			//	Tok: googleHybridResource(gcpCompute, "SubnetworkIAMBinding"),
			//	Docs: &tfbridge.DocInfo{
			//		Source: "compute_subnetwork_iam.html.markdown",
			//	},
			//},
			//"google_compute_subnetwork_iam_member": {
			//	Tok: googleHybridResource(gcpCompute, "SubnetworkIAMMember"),
			//	Docs: &tfbridge.DocInfo{
			//		Source: "compute_subnetwork_iam.html.markdown",
			//	},
			//},
			//"google_compute_subnetwork_iam_policy": {
			//	Tok: googleHybridResource(gcpCompute, "SubnetworkIAMPolicy"),
			//	Docs: &tfbridge.DocInfo{
			//		Source: "compute_subnetwork_iam.html.markdown",
			//	},
			//},
			//"google_compute_target_http_proxy":   {Tok: googleHybridResource(gcpCompute, "TargetHttpProxy")},
			//"google_compute_target_https_proxy":  {Tok: googleHybridResource(gcpCompute, "TargetHttpsProxy")},
			//"google_compute_target_instance":     {Tok: googleHybridResource(gcpCompute, "TargetInstance")},
			//"google_compute_target_ssl_proxy":    {Tok: googleHybridResource(gcpCompute, "TargetSSLProxy")},
			//"google_compute_target_tcp_proxy":    {Tok: googleHybridResource(gcpCompute, "TargetTCPProxy")},
			//"google_compute_target_pool":         {Tok: googleHybridResource(gcpCompute, "TargetPool")},
			//"google_compute_target_grpc_proxy":   {Tok: googleHybridResource(gcpCompute, "TargetGrpcProxy")},
			//"google_compute_url_map":             {Tok: googleHybridResource(gcpCompute, "URLMap")},
			//"google_compute_vpn_gateway":         {Tok: googleHybridResource(gcpCompute, "VPNGateway")},
			//"google_compute_vpn_tunnel":          {Tok: googleHybridResource(gcpCompute, "VPNTunnel")},
			//"google_compute_reservation":         {Tok: googleHybridResource(gcpCompute, "Reservation")},
			//"google_compute_region_health_check": {Tok: googleHybridResource(gcpCompute, "RegionHealthCheck")},
			//"google_compute_region_url_map":      {Tok: googleHybridResource(gcpCompute, "RegionUrlMap")},
			//"google_compute_region_disk_resource_policy_attachment": {
			//	Tok: googleHybridResource(gcpCompute, "RegionDiskResourcePolicyAttachment"),
			//},
			//"google_compute_disk_resource_policy_attachment": {
			//	Tok: googleHybridResource(gcpCompute, "DiskResourcePolicyAttachment"),
			//},
			//"google_compute_packet_mirroring":           {Tok: googleHybridResource(gcpCompute, "PacketMirroring")},
			//"google_compute_instance_group_named_port":  {Tok: googleHybridResource(gcpCompute, "InstanceGroupNamedPort")},
			//"google_compute_per_instance_config":        {Tok: googleHybridResource(gcpCompute, "PerInstanceConfig")},
			//"google_compute_region_per_instance_config": {Tok: googleHybridResource(gcpCompute, "RegionPerInstanceConfig")},
			//"google_compute_machine_image":              {Tok: googleHybridResource(gcpCompute, "MachineImage")},
			//"google_compute_image_iam_binding": {
			//	Tok: googleHybridResource(gcpCompute, "ImageIamBinding"),
			//	Docs: &tfbridge.DocInfo{
			//		Source: "compute_image_iam.html.markdown",
			//	},
			//},
			//"google_compute_image_iam_member": {
			//	Tok: googleHybridResource(gcpCompute, "ImageIamMember"),
			//	Docs: &tfbridge.DocInfo{
			//		Source: "compute_image_iam.html.markdown",
			//	},
			//},
			//"google_compute_image_iam_policy": {
			//	Tok: googleHybridResource(gcpCompute, "ImageIamPolicy"),
			//	Docs: &tfbridge.DocInfo{
			//		Source: "compute_image_iam.html.markdown",
			//	},
			//},
			//"google_compute_disk_iam_binding": {
			//	Tok: googleHybridResource(gcpCompute, "DiskIamBinding"),
			//	Docs: &tfbridge.DocInfo{
			//		Source: "compute_disk_iam.html.markdown",
			//	},
			//},
			//"google_compute_disk_iam_member": {
			//	Tok: googleHybridResource(gcpCompute, "DiskIamMember"),
			//	Docs: &tfbridge.DocInfo{
			//		Source: "compute_disk_iam.html.markdown",
			//	},
			//},
			//"google_compute_disk_iam_policy": {
			//	Tok: googleHybridResource(gcpCompute, "DiskIamPolicy"),
			//	Docs: &tfbridge.DocInfo{
			//		Source: "compute_disk_iam.html.markdown",
			//	},
			//},
			//"google_compute_region_disk_iam_binding": {
			//	Tok: googleHybridResource(gcpCompute, "RegionDiskIamBinding"),
			//	Docs: &tfbridge.DocInfo{
			//		Source: "compute_disk_iam.html.markdown",
			//	},
			//},
			//"google_compute_region_disk_iam_member": {
			//	Tok: googleHybridResource(gcpCompute, "RegionDiskIamMember"),
			//	Docs: &tfbridge.DocInfo{
			//		Source: "compute_disk_iam.html.markdown",
			//	},
			//},
			//"google_compute_region_disk_iam_policy": {
			//	Tok: googleHybridResource(gcpCompute, "RegionDiskIamPolicy"),
			//	Docs: &tfbridge.DocInfo{
			//		Source: "compute_disk_iam.html.markdown",
			//	},
			//},
			//"google_compute_organization_security_policy_rule": {
			//	Tok: googleHybridResource(gcpCompute, "OrganizationSecurityPolicyRule"),
			//},
			//"google_compute_organization_security_policy_association": {
			//	Tok: googleHybridResource(gcpCompute, "OrganizationSecurityPolicyAssociation"),
			//},
			//"google_compute_organization_security_policy": {
			//	Tok: googleHybridResource(gcpCompute, "OrganizationSecurityPolicy"),
			//},
			//"google_compute_service_attachment": {Tok: googleHybridResource(gcpCompute, "ServiceAttachment")},
			//"google_compute_backend_service_iam_binding": {
			//	Tok: googleHybridResource(gcpCompute, "BackendServiceIamBinding"),
			//	Docs: &tfbridge.DocInfo{
			//		Source: "compute_backend_service_iam.html.markdown",
			//	},
			//},
			//"google_compute_backend_service_iam_member": {
			//	Tok: googleHybridResource(gcpCompute, "BackendServiceIamMember"),
			//	Docs: &tfbridge.DocInfo{
			//		Source: "compute_backend_service_iam.html.markdown",
			//	},
			//},
			//"google_compute_backend_service_iam_policy": {
			//	Tok: googleHybridResource(gcpCompute, "BackendServiceIamPolicy"),
			//	Docs: &tfbridge.DocInfo{
			//		Source: "compute_backend_service_iam.html.markdown",
			//	},
			//},
			//"google_compute_region_backend_service_iam_binding": {
			//	Tok: googleHybridResource(gcpCompute, "RegionBackendServiceIamBinding"),
			//	Docs: &tfbridge.DocInfo{
			//		Source: "compute_region_backend_service_iam.html.markdown",
			//	},
			//},
			//"google_compute_region_backend_service_iam_member": {
			//	Tok: googleHybridResource(gcpCompute, "RegionBackendServiceIamMember"),
			//	Docs: &tfbridge.DocInfo{
			//		Source: "compute_region_backend_service_iam.html.markdown",
			//	},
			//},
			//"google_compute_region_backend_service_iam_policy": {
			//	Tok: googleHybridResource(gcpCompute, "RegionBackendServiceIamPolicy"),
			//	Docs: &tfbridge.DocInfo{
			//		Source: "compute_region_backend_service_iam.html.markdown",
			//	},
			//},
			//
			//// Container Analysis resources
			//"google_container_analysis_note": {
			//	Tok: googleHybridResource(gcpContainerAnalysis, "Note"),
			//	Docs: &tfbridge.DocInfo{
			//		Source: "containeranalysis_note.html.markdown",
			//	},
			//},
			//"google_container_analysis_occurrence": {
			//	// nolint:misspell
			//	Tok: googleHybridResource(gcpContainerAnalysis, "Occurence"),
			//},

			// Container/Kubernetes resources
			"google_container_cluster":         {Tok: googleHybridResource(gcpKubernetes, "Cluster")},
			"google_container_node_pool":       {Tok: googleHybridResource(gcpKubernetes, "NodePool")},
			"google_container_registry":        {Tok: googleHybridResource(gcpKubernetes, "Registry")},
			"google_container_aws_cluster":     {Tok: googleHybridResource(gcpKubernetes, "AwsCluster")},
			"google_container_aws_node_pool":   {Tok: googleHybridResource(gcpKubernetes, "AwsNodePool")},
			"google_container_azure_client":    {Tok: googleHybridResource(gcpKubernetes, "AzureClient")},
			"google_container_azure_cluster":   {Tok: googleHybridResource(gcpKubernetes, "AzureCluster")},
			"google_container_azure_node_pool": {Tok: googleHybridResource(gcpKubernetes, "AzureNodePool")},

			//// Data Flow resources
			//"google_dataflow_job":               {Tok: googleHybridResource(gcpDataFlow, "Job")},
			//"google_dataflow_flex_template_job": {Tok: googleHybridResource(gcpDataFlow, "FlexTemplateJob")},
			//
			//// Data Proc resources
			//"google_dataproc_cluster": {Tok: googleHybridResource(gcpDataProc, "Cluster")},
			//"google_dataproc_cluster_iam_binding": {
			//	Tok: googleHybridResource(gcpDataProc, "ClusterIAMBinding"),
			//	Docs: &tfbridge.DocInfo{
			//		Source: "dataproc_cluster_iam.html.markdown",
			//	},
			//},
			//"google_dataproc_cluster_iam_member": {
			//	Tok: googleHybridResource(gcpDataProc, "ClusterIAMMember"),
			//	Docs: &tfbridge.DocInfo{
			//		Source: "dataproc_cluster_iam.html.markdown",
			//	},
			//},
			//"google_dataproc_cluster_iam_policy": {
			//	Tok: googleHybridResource(gcpDataProc, "ClusterIAMPolicy"),
			//	Docs: &tfbridge.DocInfo{
			//		Source: "dataproc_cluster_iam.html.markdown",
			//	},
			//},
			//"google_dataproc_job": {Tok: googleHybridResource(gcpDataProc, "Job")},
			//"google_dataproc_job_iam_binding": {
			//	Tok: googleHybridResource(gcpDataProc, "JobIAMBinding"),
			//	Docs: &tfbridge.DocInfo{
			//		Source: "dataproc_job_iam.html.markdown",
			//	},
			//},
			//"google_dataproc_job_iam_member": {
			//	Tok: googleHybridResource(gcpDataProc, "JobIAMMember"),
			//	Docs: &tfbridge.DocInfo{
			//		Source: "dataproc_job_iam.html.markdown",
			//	},
			//},
			//"google_dataproc_job_iam_policy": {
			//	Tok: googleHybridResource(gcpDataProc, "JobIAMPolicy"),
			//	Docs: &tfbridge.DocInfo{
			//		Source: "dataproc_job_iam.html.markdown",
			//	},
			//},
			//"google_dataproc_autoscaling_policy": {Tok: googleHybridResource(gcpDataProc, "AutoscalingPolicy")},
			//"google_dataproc_metastore_service":  {Tok: googleHybridResource(gcpDataProc, "MetastoreService")},
			//"google_dataproc_workflow_template":  {Tok: googleHybridResource(gcpDataProc, "WorkflowTemplate")},
			//
			//// DataStore resources
			//"google_datastore_index": {Tok: googleHybridResource(gcpDatastore, "DataStoreIndex")},
			//
			//// DNS resources
			//"google_dns_managed_zone": {
			//	Tok: googleHybridResource(gcpDNS, "ManagedZone"),
			//	Fields: map[string]*tfbridge.SchemaInfo{
			//		"description": {
			//			Default: managedByPulumi,
			//		},
			//	},
			//},
			//"google_dns_policy": {Tok: googleHybridResource(gcpDNS, "Policy")},
			//"google_dns_record_set": {
			//	Tok: googleHybridResource(gcpDNS, "RecordSet"),
			//	Fields: map[string]*tfbridge.SchemaInfo{
			//		// Do not autoname RecordSet records, as the "name" of these is actually the true
			//		// domain name of the DNS record.
			//		"name": {Name: "name"},
			//	},
			//},
			//"google_dns_response_policy":      {Tok: googleHybridResource(gcpDNS, "ResponsePolicy")},
			//"google_dns_response_policy_rule": {Tok: googleHybridResource(gcpDNS, "ResponsePolicyRule")},
			//
			//// EndPoints resources
			//"google_endpoints_service": {Tok: googleHybridResource(gcpEndPoints, "Service")},
			//"google_endpoints_service_iam_binding": {
			//	Tok: googleHybridResource(gcpEndPoints, "ServiceIamBinding"),
			//	Docs: &tfbridge.DocInfo{
			//		Source: "endpoints_service_iam.html.markdown",
			//	},
			//},
			//"google_endpoints_service_iam_member": {
			//	Tok: googleHybridResource(gcpEndPoints, "ServiceIamMember"),
			//	Docs: &tfbridge.DocInfo{
			//		Source: "endpoints_service_iam.html.markdown",
			//	},
			//},
			//"google_endpoints_service_iam_policy": {
			//	Tok: googleHybridResource(gcpEndPoints, "ServiceIamPolicy"),
			//	Docs: &tfbridge.DocInfo{
			//		Source: "endpoints_service_iam.html.markdown",
			//	},
			//},
			//
			//// Filestore resources
			//"google_filestore_instance": {Tok: googleHybridResource(gcpFilestore, "Instance")},
			//
			//// Firebase
			//"google_firebase_project": {
			//	Tok: googleHybridResource(gcpFirebase, "Project"),
			//	Fields: map[string]*tfbridge.SchemaInfo{
			//		"project": {
			//			CSharpName: "ProjectID",
			//		},
			//	},
			//},
			//"google_firebase_project_location": {Tok: googleHybridResource(gcpFirebase, "ProjectLocation")},
			//"google_firebase_web_app":          {Tok: googleHybridResource(gcpFirebase, "WebApp")},
			//
			//// Firestore resources
			//"google_firestore_index":    {Tok: googleHybridResource(gcpFirestore, "Index")},
			//"google_firestore_document": {Tok: googleHybridResource(gcpFirestore, "Document")},
			//
			//// Monitoring resources
			//"google_monitoring_alert_policy":         {Tok: googleHybridResource(gcpMonitoring, "AlertPolicy")},
			//"google_monitoring_group":                {Tok: googleHybridResource(gcpMonitoring, "Group")},
			//"google_monitoring_notification_channel": {Tok: googleHybridResource(gcpMonitoring, "NotificationChannel")},
			//"google_monitoring_uptime_check_config":  {Tok: googleHybridResource(gcpMonitoring, "UptimeCheckConfig")},
			//"google_monitoring_custom_service": {
			//	Tok: googleHybridResource(gcpMonitoring, "CustomService"),
			//	Docs: &tfbridge.DocInfo{
			//		Source: "monitoring_service.html.markdown",
			//	},
			//},
			//"google_monitoring_slo":               {Tok: googleHybridResource(gcpMonitoring, "Slo")},
			//"google_monitoring_dashboard":         {Tok: googleHybridResource(gcpMonitoring, "Dashboard")},
			//"google_monitoring_metric_descriptor": {Tok: googleHybridResource(gcpMonitoring, "MetricDescriptor")},
			//"google_monitoring_monitored_project": {Tok: googleHybridResource(gcpMonitoring, "MonitoredProject")},
			//
			//// PubSub resources
			//"google_pubsub_topic": {Tok: googleHybridResource(gcpPubSub, "Topic")},
			//"google_pubsub_topic_iam_binding": {
			//	Tok: googleHybridResource(gcpPubSub, "TopicIAMBinding"),
			//	Docs: &tfbridge.DocInfo{
			//		Source: "pubsub_topic_iam.html.markdown",
			//	},
			//},
			//"google_pubsub_topic_iam_member": {
			//	Tok: googleHybridResource(gcpPubSub, "TopicIAMMember"),
			//	Docs: &tfbridge.DocInfo{
			//		Source: "pubsub_topic_iam.html.markdown",
			//	},
			//},
			//"google_pubsub_topic_iam_policy": {
			//	Tok: googleHybridResource(gcpPubSub, "TopicIAMPolicy"),
			//	Docs: &tfbridge.DocInfo{
			//		Source: "pubsub_topic_iam.html.markdown",
			//	},
			//},
			//"google_pubsub_subscription":     {Tok: googleHybridResource(gcpPubSub, "Subscription")},
			//"google_pubsub_lite_reservation": {Tok: googleHybridResource(gcpPubSub, "LiteReservation")},
			//"google_pubsub_subscription_iam_binding": {
			//	Tok: googleHybridResource(gcpPubSub, "SubscriptionIAMBinding"),
			//	Docs: &tfbridge.DocInfo{
			//		Source: "pubsub_subscription_iam.html.markdown",
			//	},
			//},
			//"google_pubsub_subscription_iam_member": {
			//	Tok: googleHybridResource(gcpPubSub, "SubscriptionIAMMember"),
			//	Docs: &tfbridge.DocInfo{
			//		Source: "pubsub_subscription_iam.html.markdown",
			//	},
			//},
			//"google_pubsub_subscription_iam_policy": {
			//	Tok: googleHybridResource(gcpPubSub, "SubscriptionIAMPolicy"),
			//	Docs: &tfbridge.DocInfo{
			//		Source: "pubsub_subscription_iam.html.markdown",
			//	},
			//},
			//"google_pubsub_lite_subscription": {Tok: googleHybridResource(gcpPubSub, "LiteSubscription")},
			//"google_pubsub_lite_topic":        {Tok: googleHybridResource(gcpPubSub, "LiteTopic")},
			//"google_pubsub_schema":            {Tok: googleHybridResource(gcpPubSub, "Schema")},
			//
			//// Redis resources
			//"google_redis_instance": {Tok: googleHybridResource(gcpRedis, "Instance")},
			//
			//// Resource Manager resources
			//"google_resource_manager_lien": {Tok: googleHybridResource(gcpResourceManager, "Lien"),
			//	Docs: &tfbridge.DocInfo{
			//		Source: "resourcemanager_lien.html.markdown",
			//	},
			//},
			//
			//// Runtime Config resources
			//"google_runtimeconfig_config":   {Tok: googleHybridResource(gcpRuntimeConfig, "Config")},
			//"google_runtimeconfig_variable": {Tok: googleHybridResource(gcpRuntimeConfig, "Variable")},
			//"google_runtimeconfig_config_iam_binding": {
			//	Tok: googleHybridResource(gcpRuntimeConfig, "ConfigIamBinding"),
			//	Docs: &tfbridge.DocInfo{
			//		Source: "runtimeconfig_config_iam.html.markdown",
			//	},
			//},
			//"google_runtimeconfig_config_iam_member": {
			//	Tok: googleHybridResource(gcpRuntimeConfig, "ConfigIamMember"),
			//	Docs: &tfbridge.DocInfo{
			//		Source: "runtimeconfig_config_iam.html.markdown",
			//	},
			//},
			//"google_runtimeconfig_config_iam_policy": {
			//	Tok: googleHybridResource(gcpRuntimeConfig, "ConfigIamPolicy"),
			//	Docs: &tfbridge.DocInfo{
			//		Source: "runtimeconfig_config_iam.html.markdown",
			//	},
			//},
			//
			//// Service Networking resources
			//"google_service_networking_connection":        {Tok: googleHybridResource(gcpServiceNetworking, "Connection")},
			//"google_service_networking_peered_dns_domain": {Tok: googleHybridResource(gcpServiceNetworking, "PeeredDnsDomain")},
			//
			//// Source Repository resources
			//"google_sourcerepo_repository": {
			//	Tok: googleHybridResource(gcpSourceRepo, "Repository"),
			//	Docs: &tfbridge.DocInfo{
			//		Source: "source_repo_repository.html.markdown",
			//	},
			//},
			//"google_sourcerepo_repository_iam_binding": {
			//	Tok: googleHybridResource(gcpSourceRepo, "RepositoryIamBinding"),
			//	Docs: &tfbridge.DocInfo{
			//		Source: "pubsub_topic_iam.html.markdown",
			//	},
			//},
			//"google_sourcerepo_repository_iam_member": {
			//	Tok: googleHybridResource(gcpSourceRepo, "RepositoryIamMember"),
			//	Docs: &tfbridge.DocInfo{
			//		Source: "pubsub_topic_iam.html.markdown",
			//	},
			//},
			//"google_sourcerepo_repository_iam_policy": {
			//	Tok: googleHybridResource(gcpSourceRepo, "RepositoryIamPolicy"),
			//	Docs: &tfbridge.DocInfo{
			//		Source: "pubsub_topic_iam.html.markdown",
			//	},
			//},
			//
			//// Spanner resources
			//"google_spanner_database": {Tok: googleHybridResource(gcpSpanner, "Database")},
			//"google_spanner_database_iam_binding": {
			//	Tok: googleHybridResource(gcpSpanner, "DatabaseIAMBinding"),
			//	Docs: &tfbridge.DocInfo{
			//		Source: "spanner_database_iam.html.markdown",
			//	},
			//},
			//"google_spanner_database_iam_member": {
			//	Tok: googleHybridResource(gcpSpanner, "DatabaseIAMMember"),
			//	Docs: &tfbridge.DocInfo{
			//		Source: "spanner_database_iam.html.markdown",
			//	},
			//},
			//"google_spanner_database_iam_policy": {
			//	Tok: googleHybridResource(gcpSpanner, "DatabaseIAMPolicy"),
			//	Docs: &tfbridge.DocInfo{
			//		Source: "spanner_database_iam.html.markdown",
			//	},
			//},
			//"google_spanner_instance": {Tok: googleHybridResource(gcpSpanner, "Instance")},
			//"google_spanner_instance_iam_binding": {
			//	Tok: googleHybridResource(gcpSpanner, "InstanceIAMBinding"),
			//	Docs: &tfbridge.DocInfo{
			//		Source: "spanner_instance_iam.html.markdown",
			//	},
			//},
			//"google_spanner_instance_iam_member": {
			//	Tok: googleHybridResource(gcpSpanner, "InstanceIAMMember"),
			//	Docs: &tfbridge.DocInfo{
			//		Source: "spanner_instance_iam.html.markdown",
			//	},
			//},
			//"google_spanner_instance_iam_policy": {
			//	Tok: googleHybridResource(gcpSpanner, "InstanceIAMPolicy"),
			//	Docs: &tfbridge.DocInfo{
			//		Source: "spanner_instance_iam.html.markdown",
			//	},
			//},
			//
			//// SQL resources
			//"google_sql_database":                       {Tok: googleHybridResource(gcpSQL, "Database")},
			//"google_sql_database_instance":              {Tok: googleHybridResource(gcpSQL, "DatabaseInstance")},
			//"google_sql_user":                           {Tok: googleHybridResource(gcpSQL, "User")},
			//"google_sql_ssl_cert":                       {Tok: googleHybridResource(gcpSQL, "SslCert")},
			//"google_sql_source_representation_instance": {Tok: googleHybridResource(gcpSQL, "SourceRepresentationInstance")},
			//
			//// Stackdriver Logging resources
			//"google_logging_billing_account_exclusion":     {Tok: googleHybridResource(gcpLogging, "BillingAccountExclusion")},
			//"google_logging_billing_account_sink":          {Tok: googleHybridResource(gcpLogging, "BillingAccountSink")},
			//"google_logging_folder_exclusion":              {Tok: googleHybridResource(gcpLogging, "FolderExclusion")},
			//"google_logging_folder_sink":                   {Tok: googleHybridResource(gcpLogging, "FolderSink")},
			//"google_logging_metric":                        {Tok: googleHybridResource(gcpLogging, "Metric")},
			//"google_logging_organization_exclusion":        {Tok: googleHybridResource(gcpLogging, "OrganizationExclusion")},
			//"google_logging_organization_sink":             {Tok: googleHybridResource(gcpLogging, "OrganizationSink")},
			//"google_logging_project_exclusion":             {Tok: googleHybridResource(gcpLogging, "ProjectExclusion")},
			//"google_logging_project_sink":                  {Tok: googleHybridResource(gcpLogging, "ProjectSink")},
			//"google_logging_billing_account_bucket_config": {Tok: googleHybridResource(gcpLogging, "BillingAccountBucketConfig")},
			//"google_logging_folder_bucket_config":          {Tok: googleHybridResource(gcpLogging, "FolderBucketConfig")},
			//"google_logging_organization_bucket_config":    {Tok: googleHybridResource(gcpLogging, "OrganizationBucketConfig")},
			//"google_logging_project_bucket_config":         {Tok: googleHybridResource(gcpLogging, "ProjectBucketConfig")},
			//"google_logging_log_view":                      {Tok: googleHybridResource(gcpLogging, "LogView")},
			//
			//// Storage resources
			//"google_storage_bucket": {
			//	Tok: googleHybridResource(gcpStorage, "Bucket"),
			//	Fields: map[string]*tfbridge.SchemaInfo{
			//		// https://cloud.google.com/storage/docs/naming
			//		// Bucket names must contain 3 to 63 characters. Names containing dots can
			//		// contain up to 222 characters, but each dot-separated component can be no
			//		// longer than 63 characters.
			//		"name": tfbridge.AutoName("name", 222, "-"),
			//	},
			//},
			//"google_storage_bucket_acl": {Tok: googleHybridResource(gcpStorage, "BucketACL")},
			//"google_storage_bucket_iam_binding": {
			//	Tok: googleHybridResource(gcpStorage, "BucketIAMBinding"),
			//	Docs: &tfbridge.DocInfo{
			//		Source: "storage_bucket_iam.html.markdown",
			//	},
			//},
			//"google_storage_bucket_iam_member": {
			//	Tok: googleHybridResource(gcpStorage, "BucketIAMMember"),
			//	Docs: &tfbridge.DocInfo{
			//		Source: "storage_bucket_iam.html.markdown",
			//	},
			//},
			//"google_storage_bucket_iam_policy": {
			//	Tok: googleHybridResource(gcpStorage, "BucketIAMPolicy"),
			//	Docs: &tfbridge.DocInfo{
			//		Source: "storage_bucket_iam.html.markdown",
			//	},
			//},
			//"google_storage_bucket_object": {
			//	Tok: googleHybridResource(gcpStorage, "BucketObject"),
			//	Fields: map[string]*tfbridge.SchemaInfo{
			//		"source": {
			//			Asset: &tfbridge.AssetTranslation{
			//				Kind:   tfbridge.FileAsset,
			//				Format: resource.ZIPArchive,
			//			},
			//		},
			//	},
			//},
			//"google_storage_default_object_access_control": {Tok: googleHybridResource(gcpStorage, "DefaultObjectAccessControl")},
			//"google_storage_default_object_acl":            {Tok: googleHybridResource(gcpStorage, "DefaultObjectACL")},
			//"google_storage_notification":                  {Tok: googleHybridResource(gcpStorage, "Notification")},
			//"google_storage_object_access_control":         {Tok: googleHybridResource(gcpStorage, "ObjectAccessControl")},
			//"google_storage_object_acl":                    {Tok: googleHybridResource(gcpStorage, "ObjectACL")},
			//"google_storage_transfer_job":                  {Tok: googleHybridResource(gcpStorage, "TransferJob")},
			//"google_storage_bucket_access_control":         {Tok: googleHybridResource(gcpStorage, "BucketAccessControl")},
			//"google_storage_hmac_key":                      {Tok: googleHybridResource(gcpStorage, "HmacKey")},
			//
			//// TPU resources
			//"google_tpu_node": {Tok: googleHybridResource(gcpTPU, "Node")},
			//
			//// Vertex
			//"google_vertex_ai_dataset":                 {Tok: googleHybridResource(gcpVertex, "AiDataset")},
			//"google_vertex_ai_featurestore_entitytype": {Tok: googleHybridResource(gcpVertex, "AiFeatureStoreEntityType")},
			//"google_vertex_ai_featurestore":            {Tok: googleHybridResource(gcpVertex, "AiFeatureStore")},
			//"google_vertex_ai_metadata_store":          {Tok: googleHybridResource(gcpVertex, "AiMetadataStore")},
			//
			//// Key Management Service resources
			//"google_kms_key_ring": {
			//	Tok: googleHybridResource(gcpKMS, "KeyRing"),
			//	Docs: &tfbridge.DocInfo{
			//		Source: "google_kms_key_ring.html.markdown",
			//	},
			//},
			//"google_kms_key_ring_iam_binding": {
			//	Tok: googleHybridResource(gcpKMS, "KeyRingIAMBinding"),
			//	Docs: &tfbridge.DocInfo{
			//		Source: "google_kms_key_ring_iam.html.markdown",
			//	},
			//},
			//"google_kms_key_ring_iam_member": {
			//	Tok: googleHybridResource(gcpKMS, "KeyRingIAMMember"),
			//	Docs: &tfbridge.DocInfo{
			//		Source: "google_kms_key_ring_iam.html.markdown",
			//	},
			//},
			//"google_kms_key_ring_iam_policy": {
			//	Tok: googleHybridResource(gcpKMS, "KeyRingIAMPolicy"),
			//	Docs: &tfbridge.DocInfo{
			//		Source: "google_kms_key_ring_iam.html.markdown",
			//	},
			//},
			//"google_kms_crypto_key": {
			//	Tok: googleHybridResource(gcpKMS, "CryptoKey"),
			//	Docs: &tfbridge.DocInfo{
			//		Source: "google_kms_crypto_key.html.markdown",
			//	},
			//},
			//"google_kms_crypto_key_iam_binding": {
			//	Tok: googleHybridResource(gcpKMS, "CryptoKeyIAMBinding"),
			//	Docs: &tfbridge.DocInfo{
			//		Source: "google_kms_crypto_key_iam.html.markdown",
			//	},
			//},
			//"google_kms_crypto_key_iam_member": {
			//	Tok: googleHybridResource(gcpKMS, "CryptoKeyIAMMember"),
			//	Docs: &tfbridge.DocInfo{
			//		Source: "google_kms_crypto_key_iam.html.markdown",
			//	},
			//},
			//"google_kms_crypto_key_iam_policy": {
			//	Tok: googleHybridResource(gcpKMS, "CryptoKeyIAMPolicy"),
			//	Docs: &tfbridge.DocInfo{
			//		Source: "google_kms_crypto_key_iam.html.markdown",
			//	},
			//},
			//"google_kms_secret_ciphertext":   {Tok: googleHybridResource(gcpKMS, "SecretCiphertext")},
			//"google_kms_key_ring_import_job": {Tok: googleHybridResource(gcpKMS, "KeyRingImportJob")},
			//
			//// Cloud IAP Resources
			//"google_iap_tunnel_instance_iam_binding": {
			//	Tok: googleHybridResource(gcpIAP, "TunnelInstanceIAMBinding"),
			//	Docs: &tfbridge.DocInfo{
			//		Source: "iap_tunnel_instance_iam.html.markdown",
			//	},
			//},
			//"google_iap_tunnel_instance_iam_member": {
			//	Tok: googleHybridResource(gcpIAP, "TunnelInstanceIAMMember"),
			//	Docs: &tfbridge.DocInfo{
			//		Source: "iap_tunnel_instance_iam.html.markdown",
			//	},
			//},
			//"google_iap_tunnel_instance_iam_policy": {
			//	Tok: googleHybridResource(gcpIAP, "TunnelInstanceIAMPolicy"),
			//	Docs: &tfbridge.DocInfo{
			//		Source: "iap_tunnel_instance_iam.html.markdown",
			//	},
			//},
			//"google_iap_tunnel_iam_binding": {
			//	Tok: googleHybridResource(gcpIAP, "TunnelIamBinding"),
			//	Docs: &tfbridge.DocInfo{
			//		Source: "iap_tunnel_iam.html.markdown",
			//	},
			//},
			//"google_iap_tunnel_iam_member": {
			//	Tok: googleHybridResource(gcpIAP, "TunnelIamMember"),
			//	Docs: &tfbridge.DocInfo{
			//		Source: "iap_tunnel_iam.html.markdown",
			//	},
			//},
			//"google_iap_tunnel_iam_policy": {
			//	Tok: googleHybridResource(gcpIAP, "TunnelIamPolicy"),
			//	Docs: &tfbridge.DocInfo{
			//		Source: "iap_tunnel_iam.html.markdown",
			//	},
			//},
			//"google_iap_web_backend_service_iam_binding": {
			//	Tok: googleHybridResource(gcpIAP, "WebBackendServiceIamBinding"),
			//	Docs: &tfbridge.DocInfo{
			//		Source: "iap_web_backend_service_iam.html.markdown",
			//	},
			//},
			//"google_iap_web_backend_service_iam_member": {
			//	Tok: googleHybridResource(gcpIAP, "WebBackendServiceIamMember"),
			//	Docs: &tfbridge.DocInfo{
			//		Source: "iap_web_backend_service_iam.html.markdown",
			//	},
			//},
			//"google_iap_web_backend_service_iam_policy": {
			//	Tok: googleHybridResource(gcpIAP, "WebBackendServiceIamPolicy"),
			//	Docs: &tfbridge.DocInfo{
			//		Source: "iap_web_backend_service_iam.html.markdown",
			//	},
			//},
			//"google_iap_web_iam_binding": {
			//	Tok: googleHybridResource(gcpIAP, "WebIamBinding"),
			//	Docs: &tfbridge.DocInfo{
			//		Source: "iap_web_iam.html.markdown",
			//	},
			//},
			//"google_iap_web_iam_member": {
			//	Tok: googleHybridResource(gcpIAP, "WebIamMember"),
			//	Docs: &tfbridge.DocInfo{
			//		Source: "iap_web_iam.html.markdown",
			//	},
			//},
			//"google_iap_web_iam_policy": {
			//	Tok: googleHybridResource(gcpIAP, "WebIamPolicy"),
			//	Docs: &tfbridge.DocInfo{
			//		Source: "iap_web_iam.html.markdown",
			//	},
			//},
			//"google_iap_web_type_app_engine_iam_binding": {
			//	Tok: googleHybridResource(gcpIAP, "WebTypeAppEngingIamBinding"),
			//	Docs: &tfbridge.DocInfo{
			//		Source: "iap_web_type_app_engine_iam.html.markdown",
			//	},
			//},
			//"google_iap_web_type_app_engine_iam_member": {
			//	Tok: googleHybridResource(gcpIAP, "WebTypeAppEngingIamMember"),
			//	Docs: &tfbridge.DocInfo{
			//		Source: "iap_web_type_app_engine_iam.html.markdown",
			//	},
			//},
			//"google_iap_web_type_app_engine_iam_policy": {
			//	Tok: googleHybridResource(gcpIAP, "WebTypeAppEngingIamPolicy"),
			//	Docs: &tfbridge.DocInfo{
			//		Source: "iap_web_type_app_engine_iam.html.markdown",
			//	},
			//},
			//"google_iap_web_type_compute_iam_binding": {
			//	Tok: googleHybridResource(gcpIAP, "WebTypeComputeIamBinding"),
			//	Docs: &tfbridge.DocInfo{
			//		Source: "iap_web_type_compute_iam.html.markdown",
			//	},
			//},
			//"google_iap_web_type_compute_iam_member": {
			//	Tok: googleHybridResource(gcpIAP, "WebTypeComputeIamMember"),
			//	Docs: &tfbridge.DocInfo{
			//		Source: "iap_web_type_compute_iam.html.markdown",
			//	},
			//},
			//"google_iap_web_type_compute_iam_policy": {
			//	Tok: googleHybridResource(gcpIAP, "WebTypeComputeIamPolicy"),
			//	Docs: &tfbridge.DocInfo{
			//		Source: "iap_web_type_compute_iam.html.markdown",
			//	},
			//},
			//"google_iap_app_engine_service_iam_binding": {
			//	Tok: googleHybridResource(gcpIAP, "AppEngineServiceIamBinding"),
			//	Docs: &tfbridge.DocInfo{
			//		Source: "iap_app_engine_service_iam.html.markdown",
			//	},
			//},
			//"google_iap_app_engine_service_iam_member": {
			//	Tok: googleHybridResource(gcpIAP, "AppEngineServiceIamMember"),
			//	Docs: &tfbridge.DocInfo{
			//		Source: "iap_app_engine_service_iam.html.markdown",
			//	},
			//},
			//"google_iap_app_engine_service_iam_policy": {
			//	Tok: googleHybridResource(gcpIAP, "AppEngineServiceIamPolicy"),
			//	Docs: &tfbridge.DocInfo{
			//		Source: "iap_app_engine_service_iam.html.markdown",
			//	},
			//},
			//"google_iap_app_engine_version_iam_binding": {
			//	Tok: googleHybridResource(gcpIAP, "AppEngineVersionIamBinding"),
			//	Docs: &tfbridge.DocInfo{
			//		Source: "iap_app_engine_version_iam.html.markdown",
			//	},
			//},
			//"google_iap_app_engine_version_iam_member": {
			//	Tok: googleHybridResource(gcpIAP, "AppEngineVersionIamMember"),
			//	Docs: &tfbridge.DocInfo{
			//		Source: "iap_app_engine_version_iam.html.markdown",
			//	},
			//},
			//"google_iap_app_engine_version_iam_policy": {
			//	Tok: googleHybridResource(gcpIAP, "AppEngineVersionIamPolicy"),
			//	Docs: &tfbridge.DocInfo{
			//		Source: "iap_app_engine_version_iam.html.markdown",
			//	},
			//},
			//"google_iap_brand":  {Tok: googleHybridResource(gcpIAP, "Brand")},
			//"google_iap_client": {Tok: googleHybridResource(gcpIAP, "Client")},
			//
			//// Game Services Resources
			//"google_game_services_game_server_cluster":    {Tok: googleHybridResource(gcpGameServices, "GameServerCluster")},
			//"google_game_services_game_server_config":     {Tok: googleHybridResource(gcpGameServices, "GameServerConfig")},
			//"google_game_services_game_server_deployment": {Tok: googleHybridResource(gcpGameServices, "GameServerDeployment")},
			//"google_game_services_realm":                  {Tok: googleHybridResource(gcpGameServices, "Realm")},
			//"google_game_services_game_server_deployment_rollout": {
			//	Tok: googleHybridResource(gcpGameServices, "GameServerDeploymentRollout"),
			//},
			//
			//// Healthcare resources
			//"google_healthcare_dataset": {
			//	Tok: googleHybridResource(gcpHealthcare, "Dataset"),
			//	Docs: &tfbridge.DocInfo{
			//		Source: "healthcare_dataset.html.markdown",
			//	},
			//},
			//"google_healthcare_dataset_iam_binding": {
			//	Tok: googleHybridResource(gcpHealthcare, "DatasetIamBinding"),
			//	Docs: &tfbridge.DocInfo{
			//		Source: "healthcare_dataset_iam.html.markdown",
			//	},
			//},
			//"google_healthcare_dataset_iam_member": {
			//	Tok: googleHybridResource(gcpHealthcare, "DatasetIamMember"),
			//	Docs: &tfbridge.DocInfo{
			//		Source: "healthcare_dataset_iam.html.markdown",
			//	},
			//},
			//"google_healthcare_dataset_iam_policy": {
			//	Tok: googleHybridResource(gcpHealthcare, "DatasetIamPolicy"),
			//	Docs: &tfbridge.DocInfo{
			//		Source: "healthcare_dataset_iam.html.markdown",
			//	},
			//},
			//"google_healthcare_dicom_store": {
			//	Tok: googleHybridResource(gcpHealthcare, "DicomStore"),
			//	Docs: &tfbridge.DocInfo{
			//		Source: "healthcare_dicom_store.html.markdown",
			//	},
			//},
			//"google_healthcare_dicom_store_iam_binding": {
			//	Tok: googleHybridResource(gcpHealthcare, "DicomStoreIamBinding"),
			//	Docs: &tfbridge.DocInfo{
			//		Source: "healthcare_dicom_store_iam.html.markdown",
			//	},
			//},
			//"google_healthcare_dicom_store_iam_member": {
			//	Tok: googleHybridResource(gcpHealthcare, "DicomStoreIamMember"),
			//	Docs: &tfbridge.DocInfo{
			//		Source: "healthcare_dicom_store_iam.html.markdown",
			//	},
			//},
			//"google_healthcare_dicom_store_iam_policy": {
			//	Tok: googleHybridResource(gcpHealthcare, "DicomStoreIamPolicy"),
			//	Docs: &tfbridge.DocInfo{
			//		Source: "healthcare_dicom_store_iam.html.markdown",
			//	},
			//},
			//"google_healthcare_fhir_store": {
			//	Tok: googleHybridResource(gcpHealthcare, "FhirStore"),
			//	Docs: &tfbridge.DocInfo{
			//		Source: "healthcare_fhir_store.html.markdown",
			//	},
			//},
			//"google_healthcare_fhir_store_iam_binding": {
			//	Tok: googleHybridResource(gcpHealthcare, "FhirStoreIamBinding"),
			//	Docs: &tfbridge.DocInfo{
			//		Source: "healthcare_fhir_store_iam.html.markdown",
			//	},
			//},
			//"google_healthcare_fhir_store_iam_member": {
			//	Tok: googleHybridResource(gcpHealthcare, "FhirStoreIamMember"),
			//	Docs: &tfbridge.DocInfo{
			//		Source: "healthcare_fhir_store_iam.html.markdown",
			//	},
			//},
			//"google_healthcare_fhir_store_iam_policy": {
			//	Tok: googleHybridResource(gcpHealthcare, "FhirStoreIamPolicy"),
			//	Docs: &tfbridge.DocInfo{
			//		Source: "healthcare_fhir_store_iam.html.markdown",
			//	},
			//},
			//"google_healthcare_hl7_v2_store": {
			//	Tok: googleHybridResource(gcpHealthcare, "Hl7Store"),
			//	Fields: map[string]*tfbridge.SchemaInfo{
			//		// Prevent unwanted singularization of the property type to `Hl7StoreNotificationConfig` as that
			//		// clashes with an existing deprecated property
			//		"notification_configs": {
			//			Elem: &tfbridge.SchemaInfo{
			//				NestedType: "Hl7StoreNotificationConfigs",
			//			},
			//		},
			//	},
			//	Docs: &tfbridge.DocInfo{
			//		Source: "healthcare_hl7_v2_store.html.markdown",
			//	},
			//},
			//"google_healthcare_hl7_v2_store_iam_binding": {
			//	Tok: googleHybridResource(gcpHealthcare, "Hl7StoreIamBinding"),
			//	Docs: &tfbridge.DocInfo{
			//		Source: "healthcare_hl7_v2_store_iam.html.markdown",
			//	},
			//},
			//"google_healthcare_hl7_v2_store_iam_member": {
			//	Tok: googleHybridResource(gcpHealthcare, "Hl7StoreIamMember"),
			//	Docs: &tfbridge.DocInfo{
			//		Source: "healthcare_hl7_v2_store_iam.html.markdown",
			//	},
			//},
			//"google_healthcare_hl7_v2_store_iam_policy": {
			//	Tok: googleHybridResource(gcpHealthcare, "Hl7StoreIamPolicy"),
			//	Docs: &tfbridge.DocInfo{
			//		Source: "healthcare_hl7_v2_store_iam.html.markdown",
			//	},
			//},
			//"google_healthcare_consent_store": {Tok: googleHybridResource(gcpHealthcare, "ConsentStore")},
			//"google_healthcare_consent_store_iam_binding": {
			//	Tok: googleHybridResource(gcpHealthcare, "ConsentStoreIamBinding"),
			//	Docs: &tfbridge.DocInfo{
			//		Source: "healthcare_consent_store_iam.html.markdown",
			//	},
			//},
			//"google_healthcare_consent_store_iam_member": {
			//	Tok: googleHybridResource(gcpHealthcare, "ConsentStoreIamMember"),
			//	Docs: &tfbridge.DocInfo{
			//		Source: "healthcare_consent_store_iam.html.markdown",
			//	},
			//},
			//"google_healthcare_consent_store_iam_policy": {
			//	Tok: googleHybridResource(gcpHealthcare, "ConsentStoreIamPolicy"),
			//	Docs: &tfbridge.DocInfo{
			//		Source: "healthcare_consent_store_iam.html.markdown",
			//	},
			//},
			//
			////CloudRun Resources
			//"google_cloud_run_domain_mapping": {Tok: googleHybridResource(gcpCloudRun, "DomainMapping")},
			//"google_cloud_run_service":        {Tok: googleHybridResource(gcpCloudRun, "Service")},
			//"google_cloud_run_service_iam_binding": {
			//	Tok:  googleHybridResource(gcpCloudRun, "IamBinding"),
			//	Docs: &tfbridge.DocInfo{Source: "cloud_run_service_iam.html.markdown"},
			//},
			//"google_cloud_run_service_iam_member": {
			//	Tok:  googleHybridResource(gcpCloudRun, "IamMember"),
			//	Docs: &tfbridge.DocInfo{Source: "cloud_run_service_iam.html.markdown"},
			//},
			//"google_cloud_run_service_iam_policy": {
			//	Tok:  googleHybridResource(gcpCloudRun, "IamPolicy"),
			//	Docs: &tfbridge.DocInfo{Source: "cloud_run_service_iam.html.markdown"},
			//},
			//
			//// Machine Learning
			//"google_ml_engine_model": {Tok: googleHybridResource(gcpMachingLearning, "EngineModel")},
			//
			//// Security Center
			//"google_scc_source":              {Tok: googleHybridResource(gcpSecurityCenter, "Source")},
			//"google_scc_notification_config": {Tok: googleHybridResource(gcpSecurityCenter, "NotificationConfig")},
			//
			//// VPC Access
			//"google_vpc_access_connector": {Tok: googleHybridResource(gcpVpcAccess, "Connector")},
			//
			//// DataFusion
			//"google_data_fusion_instance": {Tok: googleHybridResource(gcpDataFusion, "Instance")},
			//
			//// Cloudtasks
			//"google_cloud_tasks_queue": {Tok: googleHybridResource(gcpCloudTasks, "Queue")},
			//
			//// Deployment Manager
			//"google_deployment_manager_deployment": {Tok: googleHybridResource(gcpDeploymentManager, "Deployment")},
			//
			//// Identity Platform
			//"google_identity_platform_default_supported_idp_config": {
			//	Tok: googleHybridResource(gcpIdentityPlatform, "DefaultSupportedIdpConfig"),
			//},
			//"google_identity_platform_inbound_saml_config": {Tok: googleHybridResource(gcpIdentityPlatform, "InboundSamlConfig")},
			//"google_identity_platform_oauth_idp_config":    {Tok: googleHybridResource(gcpIdentityPlatform, "OauthIdpConfig")},
			//"google_identity_platform_tenant_default_supported_idp_config": {
			//	Tok: googleHybridResource(gcpIdentityPlatform, "TenantDefaultSupportedIdpConfig"),
			//},
			//"google_identity_platform_tenant_inbound_saml_config": {
			//	Tok: googleHybridResource(gcpIdentityPlatform, "TenantInboundSamlConfig"),
			//},
			//"google_identity_platform_tenant_oauth_idp_config": {
			//	Tok: googleHybridResource(gcpIdentityPlatform, "TenantOauthIdpConfig"),
			//},
			//"google_identity_platform_tenant": {Tok: googleHybridResource(gcpIdentityPlatform, "Tenant")},
			//
			//// Diagflow
			//"google_dialogflow_agent":          {Tok: googleHybridResource(gcpDiagflow, "Agent")},
			//"google_dialogflow_intent":         {Tok: googleHybridResource(gcpDiagflow, "Intent")},
			//"google_dialogflow_entity_type":    {Tok: googleHybridResource(gcpDiagflow, "EntityType")},
			//"google_dialogflow_fulfillment":    {Tok: googleHybridResource(gcpDiagflow, "Fulfillment")},
			//"google_dialogflow_cx_agent":       {Tok: googleHybridResource(gcpDiagflow, "CxAgent")},
			//"google_dialogflow_cx_flow":        {Tok: googleHybridResource(gcpDiagflow, "CxFlow")},
			//"google_dialogflow_cx_intent":      {Tok: googleHybridResource(gcpDiagflow, "CxIntent")},
			//"google_dialogflow_cx_version":     {Tok: googleHybridResource(gcpDiagflow, "CxVersion")},
			//"google_dialogflow_cx_entity_type": {Tok: googleHybridResource(gcpDiagflow, "CxEntityType")},
			//"google_dialogflow_cx_page":        {Tok: googleHybridResource(gcpDiagflow, "CxPage")},
			//"google_dialogflow_cx_environment": {Tok: googleHybridResource(gcpDiagflow, "CxEnvironment")},
			//
			//// Secret Manager
			//"google_secret_manager_secret": {Tok: googleHybridResource(gcpSecretManager, "Secret")},
			//"google_secret_manager_secret_iam_binding": {
			//	Tok: googleHybridResource(gcpSecretManager, "SecretIamBinding"),
			//	Docs: &tfbridge.DocInfo{
			//		Source: "secret_manager_secret_iam.html.markdown",
			//	},
			//},
			//"google_secret_manager_secret_iam_member": {
			//	Tok: googleHybridResource(gcpSecretManager, "SecretIamMember"),
			//	Docs: &tfbridge.DocInfo{
			//		Source: "secret_manager_secret_iam.html.markdown",
			//	},
			//},
			//"google_secret_manager_secret_iam_policy": {
			//	Tok: googleHybridResource(gcpSecretManager, "SecretIamPolicy"),
			//	Docs: &tfbridge.DocInfo{
			//		Source: "secret_manager_secret_iam.html.markdown",
			//	},
			//},
			//"google_secret_manager_secret_version": {Tok: googleHybridResource(gcpSecretManager, "SecretVersion")},
			//
			//// OS Login
			//"google_os_login_ssh_public_key": {Tok: googleHybridResource(gcpOsLogin, "SshPublicKey")},
			//
			//// Org Policy
			//"google_org_policy_policy": {Tok: googleHybridResource(gcpOrgPolicy, "Policy")},
			//
			//// OS Config
			//"google_os_config_patch_deployment":     {Tok: googleHybridResource(gcpOsConfig, "PatchDeployment")},
			//"google_os_config_guest_policies":       {Tok: googleHybridResource(gcpOsConfig, "GuestPolicies")},
			//"google_os_config_os_policy_assignment": {Tok: googleHybridResource(gcpOsConfig, "OsPolicyAssignment")},
			//
			//// Recaptcha
			//"google_recaptcha_enterprise_key": {Tok: googleHybridResource(gcpRecaptcha, "EnterpriseKey")},
			//
			//// Service Directory
			//"google_service_directory_endpoint":  {Tok: googleHybridResource(gcpServiceDirectory, "Endpoint")},
			//"google_service_directory_namespace": {Tok: googleHybridResource(gcpServiceDirectory, "Namespace")},
			//"google_service_directory_namespace_iam_binding": {
			//	Tok: googleHybridResource(gcpServiceDirectory, "NamespaceIamBinding"),
			//	Docs: &tfbridge.DocInfo{
			//		Source: "service_directory_namespace_iam.html.markdown",
			//	},
			//},
			//"google_service_directory_namespace_iam_member": {
			//	Tok: googleHybridResource(gcpServiceDirectory, "NamespaceIamMember"),
			//	Docs: &tfbridge.DocInfo{
			//		Source: "service_directory_namespace_iam.html.markdown",
			//	},
			//},
			//"google_service_directory_namespace_iam_policy": {
			//	Tok: googleHybridResource(gcpServiceDirectory, "NamespaceIamPolicy"),
			//	Docs: &tfbridge.DocInfo{
			//		Source: "service_directory_namespace_iam.html.markdown",
			//	},
			//},
			//"google_service_directory_service": {Tok: googleHybridResource(gcpServiceDirectory, "Service")},
			//"google_service_directory_service_iam_binding": {
			//	Tok: googleHybridResource(gcpServiceDirectory, "ServiceIamBinding"),
			//	Docs: &tfbridge.DocInfo{
			//		Source: "service_directory_service_iam.html.markdown",
			//	},
			//},
			//"google_service_directory_service_iam_member": {
			//	Tok: googleHybridResource(gcpServiceDirectory, "ServiceIamMember"),
			//	Docs: &tfbridge.DocInfo{
			//		Source: "service_directory_service_iam.html.markdown",
			//	},
			//},
			//"google_service_directory_service_iam_policy": {
			//	Tok: googleHybridResource(gcpServiceDirectory, "ServiceIamPolicy"),
			//	Docs: &tfbridge.DocInfo{
			//		Source: "service_directory_service_iam.html.markdown",
			//	},
			//},
			//
			//// ArtifactRegistry
			//"google_artifact_registry_repository": {Tok: googleHybridResource(gcpArtifactRegistry, "Repository")},
			//"google_artifact_registry_repository_iam_policy": {
			//	Tok: googleHybridResource(gcpArtifactRegistry, "RepositoryIamPolicy"),
			//	Docs: &tfbridge.DocInfo{
			//		Source: "artifact_registry_repository_iam.html.markdown",
			//	},
			//},
			//"google_artifact_registry_repository_iam_binding": {
			//	Tok: googleHybridResource(gcpArtifactRegistry, "RepositoryIamBinding"),
			//	Docs: &tfbridge.DocInfo{
			//		Source: "artifact_registry_repository_iam.html.markdown",
			//	},
			//},
			//"google_artifact_registry_repository_iam_member": {
			//	Tok: googleHybridResource(gcpArtifactRegistry, "RepositoryIamMember"),
			//	Docs: &tfbridge.DocInfo{
			//		Source: "artifact_registry_repository_iam.html.markdown",
			//	},
			//},
			//
			//// Data Catalog
			//"google_data_catalog_entry":        {Tok: googleHybridResource(gcpDataCatalog, "Entry")},
			//"google_data_catalog_tag":          {Tok: googleHybridResource(gcpDataCatalog, "Tag")},
			//"google_data_catalog_entry_group":  {Tok: googleHybridResource(gcpDataCatalog, "EntryGroup")},
			//"google_data_catalog_tag_template": {Tok: googleHybridResource(gcpDataCatalog, "TagTemplate")},
			//"google_data_catalog_taxonomy": {
			//	Tok: googleHybridResource(gcpDataCatalog, "Taxonomy"),
			//	Docs: &tfbridge.DocInfo{
			//		Source: "data_catalog_taxonomy.html.markdown",
			//	},
			//},
			//"google_data_catalog_policy_tag": {
			//	Tok: googleHybridResource(gcpDataCatalog, "PolicyTag"),
			//	Docs: &tfbridge.DocInfo{
			//		Source: "data_catalog_policy_tag.html.markdown ",
			//	},
			//},
			//"google_data_catalog_taxonomy_iam_binding": {
			//	Tok: googleHybridResource(gcpDataCatalog, "TaxonomyIamBinding"),
			//	Docs: &tfbridge.DocInfo{
			//		Source: "data_catalog_taxonomy_iam.html.markdown",
			//	},
			//},
			//"google_data_catalog_taxonomy_iam_member": {
			//	Tok: googleHybridResource(gcpDataCatalog, "TaxonomyIamMember"),
			//	Docs: &tfbridge.DocInfo{
			//		Source: "data_catalog_taxonomy_iam.html.markdown",
			//	},
			//},
			//"google_data_catalog_taxonomy_iam_policy": {
			//	Tok: googleHybridResource(gcpDataCatalog, "TaxonomyIamPolicy"),
			//	Docs: &tfbridge.DocInfo{
			//		Source: "data_catalog_taxonomy_iam.html.markdown",
			//	},
			//},
			//"google_data_catalog_policy_tag_iam_binding": {
			//	Tok: googleHybridResource(gcpDataCatalog, "PolicyTagIamBinding"),
			//	Docs: &tfbridge.DocInfo{
			//		Source: "data_catalog_policy_tag_iam.html.markdown",
			//	},
			//},
			//"google_data_catalog_policy_tag_iam_member": {
			//	Tok: googleHybridResource(gcpDataCatalog, "PolicyTagIamMember"),
			//	Docs: &tfbridge.DocInfo{
			//		Source: "data_catalog_policy_tag_iam.html.markdown",
			//	},
			//},
			//"google_data_catalog_policy_tag_iam_policy": {
			//	Tok: googleHybridResource(gcpDataCatalog, "PolicyTagIamPolicy"),
			//	Docs: &tfbridge.DocInfo{
			//		Source: "data_catalog_policy_tag_iam.html.markdown",
			//	},
			//},
			//"google_data_catalog_entry_group_iam_binding": {
			//	Tok: googleHybridResource(gcpDataCatalog, "EntryGroupIamBinding"),
			//	Docs: &tfbridge.DocInfo{
			//		Source: "data_catalog_entry_group_iam.html.markdown",
			//	},
			//},
			//"google_data_catalog_entry_group_iam_member": {
			//	Tok: googleHybridResource(gcpDataCatalog, "EntryGroupIamMember"),
			//	Docs: &tfbridge.DocInfo{
			//		Source: "data_catalog_entry_group_iam.html.markdown",
			//	},
			//},
			//"google_data_catalog_entry_group_iam_policy": {
			//	Tok: googleHybridResource(gcpDataCatalog, "EntryGroupIamPolicy"),
			//	Docs: &tfbridge.DocInfo{
			//		Source: "data_catalog_entry_group_iam.html.markdown",
			//	},
			//},
			//"google_data_catalog_tag_template_iam_binding": {
			//	Tok: googleHybridResource(gcpDataCatalog, "TagTemplateIamBinding"),
			//	Docs: &tfbridge.DocInfo{
			//		Source: "data_catalog_tag_template_iam.html.markdown",
			//	},
			//},
			//"google_data_catalog_tag_template_iam_member": {
			//	Tok: googleHybridResource(gcpDataCatalog, "TagTemplateIamMember"),
			//	Docs: &tfbridge.DocInfo{
			//		Source: "data_catalog_tag_template_iam.html.markdown",
			//	},
			//},
			//"google_data_catalog_tag_template_iam_policy": {
			//	Tok: googleHybridResource(gcpDataCatalog, "TagTemplateIamPolicy"),
			//	Docs: &tfbridge.DocInfo{
			//		Source: "data_catalog_tag_template_iam.html.markdown",
			//	},
			//},
			//
			//// Memcache
			//"google_memcache_instance": {Tok: googleHybridResource(gcpMemcache, "Instance")},
			//
			//// Network Connectivity
			//"google_network_connectivity_hub":   {Tok: googleHybridResource(gcpNetworkConnectivity, "Hub")},
			//"google_network_connectivity_spoke": {Tok: googleHybridResource(gcpNetworkConnectivity, "Spoke")},

			//// Network Management
			//"google_network_management_connectivity_test": {
			//	Tok: googleHybridResource(gcpNetworkManagement, "ConnectivityTest"),
			//	Docs: &tfbridge.DocInfo{
			//		Source: "network_management_connectivity_test_resource.html.markdown",
			//	},
			//},
			//
			//// Network Services
			//"google_network_services_edge_cache_keyset":  {Tok: googleHybridResource(gcpNetworkServices, "EdgeCacheKeyset")},
			//"google_network_services_edge_cache_origin":  {Tok: googleHybridResource(gcpNetworkServices, "EdgeCacheOrigin")},
			//"google_network_services_edge_cache_service": {Tok: googleHybridResource(gcpNetworkServices, "EdgeCacheService")},
			//
			//// Notebook
			//"google_notebooks_environment": {Tok: googleHybridResource(gcpNotebooks, "Environment")},
			//"google_notebooks_instance":    {Tok: googleHybridResource(gcpNotebooks, "Instance")},
			//"google_notebooks_location":    {Tok: googleHybridResource(gcpNotebooks, "Location")},
			//"google_notebooks_instance_iam_binding": {
			//	Tok: googleHybridResource(gcpNotebooks, "InstanceIamBinding"),
			//	Docs: &tfbridge.DocInfo{
			//		Source: "notebooks_instance_iam.html.markdown",
			//	},
			//},
			//"google_notebooks_instance_iam_member": {
			//	Tok: googleHybridResource(gcpNotebooks, "InstanceIamMember"),
			//	Docs: &tfbridge.DocInfo{
			//		Source: "notebooks_instance_iam.html.markdown",
			//	},
			//},
			//"google_notebooks_instance_iam_policy": {
			//	Tok: googleHybridResource(gcpNotebooks, "InstanceIamPolicy"),
			//	Docs: &tfbridge.DocInfo{
			//		Source: "notebooks_instance_iam.html.markdown",
			//	},
			//},
			//"google_notebooks_runtime": {Tok: googleHybridResource(gcpNotebooks, "Runtime")},
			//"google_notebooks_runtime_iam_binding": {
			//	Tok: googleHybridResource(gcpNotebooks, "RuntimeIamBinding"),
			//	Docs: &tfbridge.DocInfo{
			//		Source: "notebooks_runtime_iam.html.markdown",
			//	},
			//},
			//"google_notebooks_runtime_iam_member": {
			//	Tok: googleHybridResource(gcpNotebooks, "RuntimeIamMember"),
			//	Docs: &tfbridge.DocInfo{
			//		Source: "notebooks_runtime_iam.html.markdown",
			//	},
			//},
			//"google_notebooks_runtime_iam_policy": {
			//	Tok: googleHybridResource(gcpNotebooks, "RuntimeIamPolicy"),
			//	Docs: &tfbridge.DocInfo{
			//		Source: "notebooks_runtime_iam.html.markdown",
			//	},
			//},
			//
			//// CloudIdentity
			//"google_cloud_identity_group_membership": {Tok: googleHybridResource(gcpCloudIdentity, "GroupMembership")},
			//"google_cloud_identity_group":            {Tok: googleHybridResource(gcpCloudIdentity, "Group")},
			//
			//// CloudIOT
			//"google_cloudiot_device": {Tok: googleHybridResource(gcpIot, "Device")},
			//
			//// CloudAsset
			//"google_cloud_asset_folder_feed":       {Tok: googleHybridResource(gcpCloudAsset, "FolderFeed")},
			//"google_cloud_asset_organization_feed": {Tok: googleHybridResource(gcpCloudAsset, "OrganizationFeed")},
			//"google_cloud_asset_project_feed":      {Tok: googleHybridResource(gcpCloudAsset, "ProjectFeed")},
			//
			//// ActiveDirectory
			//"google_active_directory_domain":       {Tok: googleHybridResource(gcpActiveDirectory, "Domain")},
			//"google_active_directory_domain_trust": {Tok: googleHybridResource(gcpActiveDirectory, "DomainTrust")},
			//
			//// DataLoss
			//"google_data_loss_prevention_inspect_template": {
			//	Tok: googleHybridResource(gcpDataLoss, "PreventionInspectTemplate"),
			//},
			//"google_data_loss_prevention_job_trigger": {
			//	Tok: googleHybridResource(gcpDataLoss, "PreventionJobTrigger"),
			//},
			//"google_data_loss_prevention_stored_info_type": {
			//	Tok: googleHybridResource(gcpDataLoss, "PreventionStoredInfoType"),
			//},
			//"google_data_loss_prevention_deidentify_template": {
			//	Tok: googleHybridResource(gcpDataLoss, "PreventionDeidentifyTemplate"),
			//},
			//
			//// IAM
			//"google_iam_workload_identity_pool":          {Tok: googleHybridResource(gcpIAM, "WorkloadIdentityPool")},
			//"google_iam_workload_identity_pool_provider": {Tok: googleHybridResource(gcpIAM, "WorkloadIdentityPoolProvider")},
			//
			//// apigee
			//"google_apigee_organization":        {Tok: googleHybridResource(gcpApigee, "Organization")},
			//"google_apigee_instance":            {Tok: googleHybridResource(gcpApigee, "Instance")},
			//"google_apigee_envgroup":            {Tok: googleHybridResource(gcpApigee, "EnvGroup")},
			//"google_apigee_environment":         {Tok: googleHybridResource(gcpApigee, "Environment")},
			//"google_apigee_instance_attachment": {Tok: googleHybridResource(gcpApigee, "InstanceAttachment")},
			//"google_apigee_envgroup_attachment": {Tok: googleHybridResource(gcpApigee, "EnvGroupAttachment")},
			//"google_apigee_environment_iam_binding": {
			//	Tok: googleHybridResource(gcpApigee, "EnvironmentIamBinding"),
			//	Docs: &tfbridge.DocInfo{
			//		Source: "apigee_environment_iam.html.markdown",
			//	},
			//},
			//"google_apigee_environment_iam_member": {
			//	Tok: googleHybridResource(gcpApigee, "EnvironmentIamMember"),
			//	Docs: &tfbridge.DocInfo{
			//		Source: "apigee_environment_iam.html.markdown",
			//	},
			//},
			//"google_apigee_environment_iam_policy": {
			//	Tok: googleHybridResource(gcpApigee, "EnvironmentIamPolicy"),
			//	Docs: &tfbridge.DocInfo{
			//		Source: "apigee_environment_iam.html.markdown",
			//	},
			//},
			//"google_apigee_endpoint_attachment": {Tok: googleHybridResource(gcpApigee, "EndpointAttachment")},
			//
			//// API Gateway
			//"google_api_gateway_api_config_iam_binding": {
			//	Tok: googleHybridResource(gcpApiGateway, "ApiConfigIamBinding"),
			//	Docs: &tfbridge.DocInfo{
			//		Source: "api_gateway_api_config_iam.html.markdown",
			//	},
			//},
			//"google_api_gateway_api_config_iam_member": {
			//	Tok: googleHybridResource(gcpApiGateway, "ApiConfigIamMember"),
			//	Docs: &tfbridge.DocInfo{
			//		Source: "api_gateway_api_config_iam.html.markdown",
			//	},
			//},
			//"google_api_gateway_api_config_iam_policy": {
			//	Tok: googleHybridResource(gcpApiGateway, "ApiConfigIamPolicy"),
			//	Docs: &tfbridge.DocInfo{
			//		Source: "api_gateway_api_config_iam.html.markdown",
			//	},
			//},
			//"google_api_gateway_api_config": {Tok: googleHybridResource(gcpApiGateway, "ApiConfig")},
			//"google_api_gateway_api_iam_binding": {
			//	Tok: googleHybridResource(gcpApiGateway, "ApiIamBinding"),
			//	Docs: &tfbridge.DocInfo{
			//		Source: "api_gateway_api_iam.html.markdown",
			//	},
			//},
			//"google_api_gateway_api_iam_member": {
			//	Tok: googleHybridResource(gcpApiGateway, "ApiIamMember"),
			//	Docs: &tfbridge.DocInfo{
			//		Source: "api_gateway_api_iam.html.markdown",
			//	},
			//},
			//"google_api_gateway_api_iam_policy": {
			//	Tok: googleHybridResource(gcpApiGateway, "ApiIamPolicy"),
			//	Docs: &tfbridge.DocInfo{
			//		Source: "api_gateway_api_iam.html.markdown",
			//	},
			//},
			//"google_api_gateway_api": {Tok: googleHybridResource(gcpApiGateway, "Api")},
			//"google_api_gateway_gateway_iam_binding": {
			//	Tok: googleHybridResource(gcpApiGateway, "GatewayIamBinding"),
			//	Docs: &tfbridge.DocInfo{
			//		Source: "api_gateway_gateway_iam.html.markdown",
			//	},
			//},
			//"google_api_gateway_gateway_iam_member": {
			//	Tok: googleHybridResource(gcpApiGateway, "GatewayIamMember"),
			//	Docs: &tfbridge.DocInfo{
			//		Source: "api_gateway_gateway_iam.html.markdown",
			//	},
			//},
			//"google_api_gateway_gateway_iam_policy": {
			//	Tok: googleHybridResource(gcpApiGateway, "GatewayIamPolicy"),
			//	Docs: &tfbridge.DocInfo{
			//		Source: "api_gateway_gateway_iam.html.markdown",
			//	},
			//},
			//"google_api_gateway_gateway": {Tok: googleHybridResource(gcpApiGateway, "Gateway")},
			//
			//// Certificate Authority
			//"google_privateca_certificate_authority": {Tok: googleHybridResource(gcpCertificateAuthority, "Authority")},
			//"google_privateca_certificate":           {Tok: googleHybridResource(gcpCertificateAuthority, "Certificate")},
			//"google_privateca_ca_pool":               {Tok: googleHybridResource(gcpCertificateAuthority, "CaPool")},
			//"google_privateca_certificate_template":  {Tok: googleHybridResource(gcpCertificateAuthority, "CertificateTemplate")},
			//"google_privateca_ca_pool_iam_binding": {
			//	Tok: googleHybridResource(gcpCertificateAuthority, "CaPoolIamBinding"),
			//	Docs: &tfbridge.DocInfo{
			//		Source: "privateca_ca_pool_iam.html.markdown",
			//	},
			//},
			//"google_privateca_ca_pool_iam_member": {
			//	Tok: googleHybridResource(gcpCertificateAuthority, "CaPoolIamMember"),
			//	Docs: &tfbridge.DocInfo{
			//		Source: "privateca_ca_pool_iam.html.markdown",
			//	},
			//},
			//"google_privateca_ca_pool_iam_policy": {
			//	Tok: googleHybridResource(gcpCertificateAuthority, "CaPoolIamPolicy"),
			//	Docs: &tfbridge.DocInfo{
			//		Source: "privateca_ca_pool_iam.html.markdown",
			//	},
			//},
			//
			//// essential contacts
			//"google_essential_contacts_contact": {Tok: googleHybridResource(gcpEssentialContacts, "Contact")},
			//
			//// workflows
			//"google_workflows_workflow": {Tok: googleHybridResource(gcpWorkflows, "Workflow")},
			//
			////eventarc
			//"google_eventarc_trigger": {
			//	Tok: googleHybridResource(gcpEventarc, "Trigger"),
			//	Fields: map[string]*tfbridge.SchemaInfo{
			//		"transport": {
			//			Name:        "transports",
			//			MaxItemsOne: boolRef(false),
			//			Elem: &tfbridge.SchemaInfo{
			//				Fields: map[string]*tfbridge.SchemaInfo{
			//					"pubsub": {
			//						Name:        "pubsubs",
			//						MaxItemsOne: boolRef(false),
			//					},
			//				},
			//			},
			//		},
			//	},
			//},
			//
			//// gke hub
			//"google_gke_hub_membership":         {Tok: googleHybridResource(gcpGkeHub, "Membership")},
			//"google_gke_hub_feature":            {Tok: googleHybridResource(gcpGkeHub, "Feature")},
			//"google_gke_hub_feature_membership": {Tok: googleHybridResource(gcpGkeHub, "FeatureMembership")},
			//
			//// tags
			//"google_tags_tag_key":     {Tok: googleHybridResource(gcpTags, "TagKey")},
			//"google_tags_tag_value":   {Tok: googleHybridResource(gcpTags, "TagValue")},
			//"google_tags_tag_binding": {Tok: googleHybridResource(gcpTags, "TagBinding")},
			//"google_tags_tag_key_iam_binding": {
			//	Tok: googleHybridResource(gcpTags, "TagKeyIamBinding"),
			//	Docs: &tfbridge.DocInfo{
			//		Source: "tags_tag_key_iam.html.markdown",
			//	},
			//},
			//"google_tags_tag_key_iam_member": {
			//	Tok: googleHybridResource(gcpTags, "TagKeyIamMember"),
			//	Docs: &tfbridge.DocInfo{
			//		Source: "tags_tag_key_iam.html.markdown",
			//	},
			//},
			//"google_tags_tag_key_iam_policy": {
			//	Tok: googleHybridResource(gcpTags, "TagKeyIamPolicy"),
			//	Docs: &tfbridge.DocInfo{
			//		Source: "tags_tag_key_iam.html.markdown",
			//	},
			//},
			//"google_tags_tag_value_iam_binding": {
			//	Tok: googleHybridResource(gcpTags, "TagValueIamBinding"),
			//	Docs: &tfbridge.DocInfo{
			//		Source: "tags_tag_value_iam.html.markdown",
			//	},
			//},
			//"google_tags_tag_value_iam_member": {
			//	Tok: googleHybridResource(gcpTags, "TagValueIamMember"),
			//	Docs: &tfbridge.DocInfo{
			//		Source: "tags_tag_value_iam.html.markdown",
			//	},
			//},
			//"google_tags_tag_value_iam_policy": {
			//	Tok: googleHybridResource(gcpTags, "TagValueIamPolicy"),
			//	Docs: &tfbridge.DocInfo{
			//		Source: "tags_tag_value_iam.html.markdown",
			//	},
			//},
		},
		DataSources: map[string]*tfbridge.DataSourceInfo{
			//"google_billing_account": {
			//	Tok: gcpDataSource(gcpOrganization, "getBillingAccount"),
			//	Docs: &tfbridge.DocInfo{
			//		Source: "google_billing_account.html.markdown",
			//	},
			//},
			//"google_client_config": {
			//	Tok: gcpDataSource(gcpOrganization, "getClientConfig"),
			//	Docs: &tfbridge.DocInfo{
			//		Source: "datasource_client_config.html.markdown",
			//	},
			//},
			//"google_client_openid_userinfo": {
			//	Tok: gcpDataSource(gcpOrganization, "getClientOpenIdUserInfo"),
			//	Docs: &tfbridge.DocInfo{
			//		Source: "datasource_google_client_openid_userinfo.html.markdown",
			//	},
			//},
			//"google_folders": {Tok: gcpDataSource(gcpOrganization, "getFolders")},
			//
			//"google_cloudfunctions_function": {
			//	Tok: gcpDataSource(gcpCloudFunctions, "getFunction"),
			//	Docs: &tfbridge.DocInfo{
			//		Source: "datasource_cloudfunctions_function.html.markdown",
			//	},
			//},
			//"google_compute_address": {
			//	Tok: gcpDataSource(gcpCompute, "getAddress"),
			//	Docs: &tfbridge.DocInfo{
			//		Source: "datasource_compute_address.html.markdown",
			//	},
			//},
			//"google_compute_default_service_account": {
			//	Tok: gcpDataSource(gcpCompute, "getDefaultServiceAccount"),
			//	Docs: &tfbridge.DocInfo{
			//		Source: "google_compute_default_service_account.html.markdown",
			//	},
			//},
			//"google_compute_backend_service": {
			//	Tok: gcpDataSource(gcpCompute, "getBackendService"),
			//	Fields: map[string]*tfbridge.SchemaInfo{
			//		"consistent_hash": {
			//			Name: "consistentHash",
			//			Elem: &tfbridge.SchemaInfo{
			//				Fields: map[string]*tfbridge.SchemaInfo{
			//					"http_cookie": {
			//						Name: "httpCookies",
			//					},
			//				},
			//			},
			//		},
			//	},
			//	Docs: &tfbridge.DocInfo{
			//		Source: "datasource_google_compute_backend_service.html.markdown",
			//	},
			//},
			//"google_compute_backend_bucket": {
			//	Tok: gcpDataSource(gcpCompute, "getBackendBucket"),
			//	Docs: &tfbridge.DocInfo{
			//		Source: "datasource_google_compute_backend_bucket.html.markdown",
			//	},
			//},
			//"google_compute_instance_template": {Tok: gcpDataSource(gcpCompute, "getInstanceTemplate")},
			//"google_compute_image": {
			//	Tok: gcpDataSource(gcpCompute, "getImage"),
			//	Docs: &tfbridge.DocInfo{
			//		Source: "datasource_compute_image.html.markdown",
			//	},
			//},
			//"google_compute_instance": {
			//	Tok: gcpDataSource(gcpCompute, "getInstance"),
			//	Docs: &tfbridge.DocInfo{
			//		Source: "datasource_compute_instance.html.markdown",
			//	},
			//},
			//"google_compute_forwarding_rule": {
			//	Tok: gcpDataSource(gcpCompute, "getForwardingRule"),
			//	Docs: &tfbridge.DocInfo{
			//		Source: "datasource_compute_forwarding_rule.html.markdown",
			//	},
			//},
			//"google_compute_global_address": {
			//	Tok: gcpDataSource(gcpCompute, "getGlobalAddress"),
			//	Docs: &tfbridge.DocInfo{
			//		Source: "datasource_compute_global_address.html.markdown",
			//	},
			//},
			//"google_compute_global_forwarding_rule": {
			//	Tok: gcpDataSource(gcpCompute, "getGlobalForwardingRule"),
			//},
			//"google_compute_network": {
			//	Tok: gcpDataSource(gcpCompute, "getNetwork"),
			//	Docs: &tfbridge.DocInfo{
			//		Source: "datasource_compute_network.html.markdown",
			//	},
			//},
			//"google_compute_network_endpoint_group": {
			//	Tok: gcpDataSource(gcpCompute, "getNetworkEndpointGroup"),
			//	Docs: &tfbridge.DocInfo{
			//		Source: "datasource_google_compute_network_endpoint_group.html.markdown",
			//	},
			//},
			//"google_compute_router_status": {Tok: gcpDataSource(gcpCompute, "RouterStatus")},
			//
			//"google_composer_image_versions": {
			//	Tok: gcpDataSource(gcpComposer, "getImageVersions"),
			//	Docs: &tfbridge.DocInfo{
			//		Source: "datasource_google_composer_image_versions.html.markdown",
			//	},
			//},
			//"google_composer_environment": {Tok: gcpDataSource(gcpComposer, "getEnvironment")},
			//"google_iam_role": {
			//	Tok: gcpDataSource(gcpIAM, "getRule"),
			//	Docs: &tfbridge.DocInfo{
			//		Source: "datasource_google_iam_role.html.markdown",
			//	},
			//},
			//"google_iam_testable_permissions": {
			//	Tok: gcpDataSource(gcpIAM, "getTestablePermissions"),
			//	Docs: &tfbridge.DocInfo{
			//		Source: "datasource_iam_testable_permissions.html.markdown",
			//	},
			//},
			//"google_iam_workload_identity_pool":          {Tok: gcpDataSource(gcpIAM, "getWorkloadIdentityPool")},
			//"google_iam_workload_identity_pool_provider": {Tok: gcpDataSource(gcpIAM, "getWorkloadIdentityPoolProvider")},
			//"google_netblock_ip_ranges": {
			//	Tok: gcpDataSource(gcpCompute, "getNetblockIPRanges"),
			//	Docs: &tfbridge.DocInfo{
			//		Source: "datasource_google_netblock_ip_ranges.html.markdown",
			//	},
			//},
			"google_project": {
				Tok: gcpDataSource(gcpOrganization, "getProject"),
				Docs: &tfbridge.DocInfo{
					Source: "google_project.html.markdown",
				},
			},
			//"google_compute_node_types": {
			//	Tok: gcpDataSource(gcpCompute, "getNodeTypes"),
			//},
			//"google_compute_subnetwork": {
			//	Tok: gcpDataSource(gcpCompute, "getSubnetwork"),
			//	Docs: &tfbridge.DocInfo{
			//		Source: "datasource_compute_subnetwork.html.markdown",
			//	},
			//},
			//"google_compute_vpn_gateway": {
			//	Tok: gcpDataSource(gcpCompute, "getVPNGateway"),
			//	Docs: &tfbridge.DocInfo{
			//		Source: "datasource_compute_vpn_gateway.html.markdown",
			//	},
			//},
			//"google_compute_zones": {
			//	Tok: gcpDataSource(gcpCompute, "getZones"),
			//	Docs: &tfbridge.DocInfo{
			//		Source: "google_compute_zones.html.markdown",
			//	},
			//},
			//"google_compute_regions": {
			//	Tok: gcpDataSource(gcpCompute, "getRegions"),
			//	Docs: &tfbridge.DocInfo{
			//		Source: "google_compute_regions.html.markdown",
			//	},
			//},
			//"google_compute_ssl_certificate": {
			//	Tok: gcpDataSource(gcpCompute, "getCertificate"),
			//	Docs: &tfbridge.DocInfo{
			//		Source: "datasource_compute_ssl_certificate.html.markdown",
			//	},
			//},
			//"google_compute_ssl_policy": {
			//	Tok: gcpDataSource(gcpCompute, "getSSLPolicy"),
			//	Docs: &tfbridge.DocInfo{
			//		Source: "datasource_compute_ssl_policy.html.markdown",
			//	},
			//},
			//"google_compute_region_instance_group": {
			//	Tok: gcpDataSource(gcpCompute, "getRegionInstanceGroup"),
			//	Docs: &tfbridge.DocInfo{
			//		Source: "datasource_compute_region_instance_group.html.markdown",
			//	},
			//},
			//"google_compute_ha_vpn_gateway": {Tok: gcpDataSource(gcpCompute, "getHcVpnGateway")},
			//"google_compute_region_ssl_certificate": {
			//	Tok: gcpDataSource(gcpCompute, "getRegionSslCertificate"),
			//},
			//"google_compute_instance_group": {
			//	Tok: gcpDataSource(gcpCompute, "getInstanceGroup"),
			//	Docs: &tfbridge.DocInfo{
			//		Source: "google_compute_instance_group.html.markdown",
			//	},
			//},
			//"google_compute_lb_ip_ranges": {
			//	Tok: gcpDataSource(gcpCompute, "getLBIPRanges"),
			//	Docs: &tfbridge.DocInfo{
			//		Source: "datasource_compute_lb_ip_ranges.html.markdown",
			//	},
			//},
			//"google_compute_resource_policy": {
			//	Tok: gcpDataSource(gcpCompute, "getResourcePolicy"),
			//},
			//"google_compute_health_check": {Tok: gcpDataSource(gcpCompute, "getHealthCheck")},
			//
			//"google_container_cluster": {
			//	Tok: gcpDataSource(gcpKubernetes, "getCluster"),
			//	Docs: &tfbridge.DocInfo{
			//		Source: "google_container_cluster.html.markdown",
			//	},
			//},
			//"google_container_engine_versions": {
			//	Tok: gcpDataSource(gcpKubernetes, "getEngineVersions"),
			//	Docs: &tfbridge.DocInfo{
			//		Source: "google_container_engine_versions.html.markdown",
			//	},
			//},
			//"google_container_registry_repository": {
			//	Tok: gcpDataSource(gcpKubernetes, "getRegistryRepository"),
			//	Docs: &tfbridge.DocInfo{
			//		Source: "google_container_registry_repository.html.markdown",
			//	},
			//},
			//"google_container_registry_image": {
			//	Tok: gcpDataSource(gcpKubernetes, "getRegistryImage"),
			//	Docs: &tfbridge.DocInfo{
			//		Source: "google_container_registry_repository.html.markdown",
			//	},
			//},
			//"google_container_aws_versions": {
			//	Tok: gcpDataSource(gcpKubernetes, "getAwsVersions"),
			//	Docs: &tfbridge.DocInfo{
			//		Source: "container_aws_versions.html.markdown",
			//	},
			//},
			//"google_container_azure_versions": {
			//	Tok: gcpDataSource(gcpKubernetes, "getAzureVersions"),
			//	Docs: &tfbridge.DocInfo{
			//		Source: "container_azure_versions.html.markdown",
			//	},
			//},
			//
			//"google_dns_managed_zone": {
			//	Tok: gcpDataSource(gcpDNS, "getManagedZone"),
			//	Docs: &tfbridge.DocInfo{
			//		Source: "dns_managed_zone.html.markdown",
			//	},
			//},
			//"google_dns_keys": {
			//	Tok: gcpDataSource(gcpDNS, "getKeys"),
			//	Docs: &tfbridge.DocInfo{
			//		Source: "datasource_dns_keys.html.markdown",
			//	},
			//},
			//"google_dns_record_set": {Tok: gcpDataSource(gcpDNS, "getRecordSet")},
			//
			//"google_active_folder": {
			//	Tok: gcpDataSource(gcpOrganization, "getActiveFolder"),
			//	Docs: &tfbridge.DocInfo{
			//		Source: "google_active_folder.html.markdown",
			//	},
			//},
			//"google_folder": {
			//	Tok: gcpDataSource(gcpOrganization, "getFolder"),
			//	Docs: &tfbridge.DocInfo{
			//		Source: "google_folder.html.markdown",
			//	},
			//},
			//"google_folder_organization_policy": {
			//	Tok: gcpDataSource(gcpFolder, "getOrganizationPolicy"),
			//	Docs: &tfbridge.DocInfo{
			//		Source: "datasource_google_folder_organization_policy.html.markdown",
			//	},
			//},
			//"google_iam_policy": {
			//	Tok: gcpDataSource(gcpOrganization, "getIAMPolicy"),
			//	Docs: &tfbridge.DocInfo{
			//		Source: "google_iam_policy.html.markdown",
			//	},
			//},
			//"google_kms_crypto_key": {
			//	Tok: gcpDataSource(gcpKMS, "getKMSCryptoKey"),
			//	Docs: &tfbridge.DocInfo{
			//		Source: "google_kms_crypto_key.html.markdown",
			//	},
			//},
			//"google_kms_key_ring": {
			//	Tok: gcpDataSource(gcpKMS, "getKMSKeyRing"),
			//	Docs: &tfbridge.DocInfo{
			//		Source: "google_kms_key_ring.html.markdown",
			//	},
			//},
			//"google_kms_secret": {
			//	Tok: gcpDataSource(gcpKMS, "getKMSSecret"),
			//	Docs: &tfbridge.DocInfo{
			//		Source: "google_kms_secret.html.markdown",
			//	},
			//},
			//"google_kms_crypto_key_version": {Tok: gcpDataSource(gcpKMS, "getKMSCryptoKeyVersion")},
			//"google_kms_secret_asymmetric":  {Tok: gcpDataSource(gcpKMS, "getKMSSecretAsymmetric")},
			//"google_kms_secret_ciphertext":  {Tok: gcpDataSource(gcpKMS, "getKMSSecretCiphertext")},
			//"google_organization": {
			//	Tok: gcpDataSource(gcpOrganization, "getOrganization"),
			//	Docs: &tfbridge.DocInfo{
			//		Source: "google_organization.html.markdown",
			//	},
			//},
			//
			//"google_privateca_certificate_authority": {Tok: gcpDataSource(gcpCertificateAuthority, "getAuthority")},
			//
			//"google_projects": {
			//	Tok: gcpDataSource(gcpProject, "getProject"),
			//	Docs: &tfbridge.DocInfo{
			//		Source: "google_projects.html.markdown",
			//	},
			//},
			//"google_project_organization_policy": {
			//	Tok: gcpDataSource(gcpProject, "getOrganizationPolicy"),
			//	Docs: &tfbridge.DocInfo{
			//		Source: "datasource_google_project_organization_policy.html.markdown",
			//	},
			//},
			//"google_storage_bucket_object": {
			//	Tok: gcpDataSource(gcpStorage, "getBucketObject"),
			//	Docs: &tfbridge.DocInfo{
			//		Source: "bucket_object.html.markdown",
			//	},
			//},
			//"google_storage_object_signed_url": {
			//	Tok: gcpDataSource(gcpStorage, "getObjectSignedUrl"),
			//	Docs: &tfbridge.DocInfo{
			//		Source: "signed_url.html.markdown",
			//	},
			//},
			//"google_storage_project_service_account": {
			//	Tok: gcpDataSource(gcpStorage, "getProjectServiceAccount"),
			//	Docs: &tfbridge.DocInfo{
			//		Source: "google_storage_project_service_account.html.markdown",
			//	},
			//},
			//"google_storage_transfer_project_service_account": {
			//	Tok: gcpDataSource(gcpStorage, "getTransferProjectServieAccount"),
			//	Docs: &tfbridge.DocInfo{
			//		Source: "google_storage_transfer_project_service_account.html.markdown",
			//	},
			//},
			//"google_storage_bucket_object_content": {
			//	Tok: gcpDataSource(gcpStorage, "getBucketObjectContent"),
			//},
			//"google_storage_bucket": {Tok: gcpDataSource(gcpStorage, "getBucket")},
			//"google_service_account": {
			//	Tok: gcpDataSource(gcpServiceAccount, "getAccount"),
			//	Docs: &tfbridge.DocInfo{
			//		Source: "datasource_google_service_account.html.markdown",
			//	},
			//},
			//"google_service_account_id_token": {
			//	Tok: gcpDataSource(gcpServiceAccount, "getAccountIdToken"),
			//	Docs: &tfbridge.DocInfo{
			//		Source: "datasource_google_service_account_id_token.html.markdown",
			//	},
			//},
			//"google_service_account_access_token": {
			//	Tok: gcpDataSource(gcpServiceAccount, "getAccountAccessToken"),
			//	Docs: &tfbridge.DocInfo{
			//		Source: "datasource_google_service_account_access_token.html.markdown",
			//	},
			//},
			//"google_service_account_key": {
			//	Tok: gcpDataSource(gcpServiceAccount, "getAccountKey"),
			//	Docs: &tfbridge.DocInfo{
			//		Source: "datasource_google_service_account_key.html.markdown",
			//	},
			//},
			//"google_tpu_tensorflow_versions": {
			//	Tok: gcpDataSource(gcpTPU, "getTensorflowVersions"),
			//	Docs: &tfbridge.DocInfo{
			//		Source: "datasource_tpu_tensorflow_versions.html.markdown",
			//	},
			//},
			//"google_compute_router": {
			//	Tok: gcpDataSource(gcpCompute, "getRouter"),
			//	Docs: &tfbridge.DocInfo{
			//		Source: "datasource_compute_router.html.markdown",
			//	},
			//},
			//"google_compute_instance_serial_port": {
			//	Tok: gcpDataSource(gcpCompute, "getInstanceSerialPort"),
			//	Docs: &tfbridge.DocInfo{
			//		Source: "datasource_compute_instance_serial_port.html.markdown",
			//	},
			//},
			//"google_bigquery_default_service_account": {
			//	Tok:  gcpDataSource(gcpBigQuery, "getDefaultServiceAccount"),
			//	Docs: &tfbridge.DocInfo{Source: "google_bigquery_default_service_account.html"},
			//},
			//"google_sql_ca_certs": {
			//	Tok: gcpDataSource(gcpSQL, "getCaCerts"),
			//	Docs: &tfbridge.DocInfo{
			//		Source: "datasource_google_sql_ca_certs.html.markdown",
			//	},
			//},
			//"google_sql_database_instance": {
			//	Tok: gcpDataSource(gcpSQL, "getDatabaseInstance"),
			//},
			//"google_service_networking_peered_dns_domain": {
			//	Tok: gcpDataSource(gcpServiceNetworking, "getPeeredDnsDomain"),
			//	// At the time of writing this data source does not have any upstream docs at all, so we override
			//	// with blank contents.
			//	Docs: &tfbridge.DocInfo{
			//		Markdown: []byte(" "),
			//	},
			//},
			//"google_sql_backup_run": {Tok: gcpDataSource(gcpSQL, "getBackupRun")},
			//"google_monitoring_notification_channel": {
			//	Tok: gcpDataSource(gcpMonitoring, "getNotificationChannel"),
			//	Docs: &tfbridge.DocInfo{
			//		Source: "datasource_monitoring_notification_channel.html.markdown",
			//	},
			//},
			//"google_monitoring_uptime_check_ips": {
			//	Tok: gcpDataSource(gcpMonitoring, "getUptimeCheckIPs"),
			//	Docs: &tfbridge.DocInfo{
			//		Source: "datasource_google_monitoring_uptime_check_ips.html.markdown",
			//	},
			//},
			//"google_monitoring_app_engine_service": {
			//	Tok: gcpDataSource(gcpMonitoring, "getAppEngineService"),
			//	Docs: &tfbridge.DocInfo{
			//		Source: "datasource_monitoring_app_engine_service.html.markdown",
			//	},
			//},
			//"google_monitoring_cluster_istio_service":   {Tok: gcpDataSource(gcpMonitoring, "getClusterIstioService")},
			//"google_monitoring_mesh_istio_service":      {Tok: gcpDataSource(gcpMonitoring, "getMeshIstioService")},
			//"google_monitoring_istio_canonical_service": {Tok: gcpDataSource(gcpMonitoring, "getIstioCanonicalService")},
			//
			//// Firebase
			//"google_firebase_web_app_config": {
			//	Tok: gcpDataSource(gcpFirebase, "getWebAppConfig"),
			//	Docs: &tfbridge.DocInfo{
			//		Source: "datasource_firebase_web_app_config.html.markdown",
			//	},
			//},
			//"google_firebase_web_app": {
			//	Tok: gcpDataSource(gcpFirebase, "getWebApp"),
			//	Docs: &tfbridge.DocInfo{
			//		Source: "datasource_firebase_web_app.html.markdown",
			//	},
			//},
			//
			//// Redis resources
			//"google_redis_instance": {
			//	Tok: gcpDataSource(gcpRedis, "getInstance"),
			//	Docs: &tfbridge.DocInfo{
			//		Source: "datasource_google_redis_instance.html.markdown",
			//	},
			//},
			//
			//// CloudIdentity
			//"google_cloud_identity_group_memberships": {
			//	Tok: gcpDataSource(gcpCloudIdentity, "getGroupMemberships"),
			//	Docs: &tfbridge.DocInfo{
			//		Source: "cloud_identity_group_membership.html.markdown",
			//	},
			//},
			//"google_cloud_identity_groups": {Tok: gcpDataSource(gcpCloudIdentity, "getGroups")},
			//
			//// GameServices
			//"google_game_services_game_server_deployment_rollout": {
			//	Tok: gcpDataSource(gcpGameServices, "getGameServerDeploymentRollout"),
			//},
			//
			//// Pubsub
			//"google_pubsub_topic": {Tok: gcpDataSource(gcpPubSub, "getTopic")},
			//
			////Cloud Run
			//"google_cloud_run_service":   {Tok: gcpDataSource(gcpCloudRun, "getService")},
			//"google_cloud_run_locations": {Tok: gcpDataSource(gcpCloudRun, "getLocations")},
			//
			////Appengine
			//"google_app_engine_default_service_account": {Tok: gcpDataSource(gcpAppEngine, "getDefaultServiceAccount")},
			//
			//// Source repo
			//"google_sourcerepo_repository": {
			//	Tok: gcpDataSource(gcpSourceRepo, "getRepository"),
			//	Docs: &tfbridge.DocInfo{
			//		Source: "data_source_sourcerepo_repository.html.markdown",
			//	},
			//},
			//
			//// Spanner
			//"google_spanner_instance": {Tok: gcpDataSource(gcpSpanner, "getInstance")},
			//
			////runtime config
			//"google_runtimeconfig_config":   {Tok: gcpDataSource(gcpRuntimeConfig, "getConfig")},
			//"google_runtimeconfig_variable": {Tok: gcpDataSource(gcpRuntimeConfig, "getVariable")},
			//
			//// IAP
			//"google_iap_client": {Tok: gcpDataSource(gcpIAP, "getClient")},
			//
			//// Secret Manager
			//"google_secret_manager_secret": {Tok: gcpDataSource(gcpSecretManager, "getSecret")},
		},
		JavaScript: &tfbridge.JavaScriptInfo{
			Dependencies: map[string]string{
				"@pulumi/pulumi":    "^3.0.0",
				"read-package-json": "^2.0.13",
				"@types/express":    "^4.16.0",
			},
			DevDependencies: map[string]string{
				"@types/node": "^10.0.0", // so we can access strongly typed node definitions.
			},
			Overlay: &tfbridge.OverlayInfo{
				DestFiles: []string{
					"utils.ts", // Helpers,
				},
				Modules: map[string]*tfbridge.OverlayInfo{
					"cloudfunctions": {
						DestFiles: []string{
							// named with 'z' to come after all relevant files in package since the
							// generator sorts these by name.
							"zMixins.ts",
						},
					},
					"pubsub": {
						DestFiles: []string{
							// named with 'z' to come after all relevant files in package since the
							// generator sorts these by name.
							"zMixins.ts",
						},
					},
					"storage": {
						DestFiles: []string{
							// named with 'z' to come after all relevant files in package since the
							// generator sorts these by name.
							"zMixins.ts",
						},
					},
				},
			},
		},
		Golang: &tfbridge.GolangInfo{
			ImportBasePath: filepath.Join(
				fmt.Sprintf("github.com/pulumi/pulumi-%[1]s/sdk/", gcpPackage),
				tfbridge.GetModuleMajorVersion(version.Version),
				"go",
				gcpPackage,
			),
			GenerateResourceContainerTypes: true,
		},
		Python: &tfbridge.PythonInfo{
			Requires: map[string]string{
				"pulumi": ">=3.0.0,<4.0.0",
			},
		},
		CSharp: &tfbridge.CSharpInfo{
			PackageReferences: map[string]string{
				"Pulumi": "3.*",
			},
			Namespaces: namespaceMap,
		},
	}

	prov.RenameResourceWithAlias("google_compute_autoscaler", googleHybridResource(gcpCompute,
		"Autoscalar"), googleHybridResource(gcpCompute, "Autoscaler"), gcpCompute, gcpCompute, &tfbridge.ResourceInfo{
		Docs: &tfbridge.DocInfo{
			Source: "compute_autoscaler.html.markdown",
		},
	})

	prov.RenameResourceWithAlias("google_compute_managed_ssl_certificate", googleHybridResource(gcpCompute,
		"MangedSslCertificate"), googleHybridResource(gcpCompute, "ManagedSslCertificate"), gcpCompute, gcpCompute,
		&tfbridge.ResourceInfo{
			Docs: &tfbridge.DocInfo{
				Source: "compute_managed_ssl_certificate.html.markdown",
			},
		})

	prov.RenameDataSource("google_secret_manager_secret_version", gcpDataSource(gcpMonitoring, "getSecretVersion"),
		gcpDataSource(gcpSecretManager, "getSecretVersion"), gcpMonitoring, gcpSecretManager,
		&tfbridge.DataSourceInfo{
			Docs: &tfbridge.DocInfo{
				Source: "secret_manager_secret_version.html.markdown",
			},
		})

	prov.RenameResourceWithAlias("google_cloudiot_registry", googleHybridResource(gcpKMS,
		"Registry"), googleHybridResource(gcpIot, "Registry"), gcpKMS, gcpIot, &tfbridge.ResourceInfo{
		Fields: map[string]*tfbridge.SchemaInfo{
			// This property's nested type name conflicts with the nested type of the existing (now deprecated)
			// `event_notification_config` property (singular, a TypeMap). A conflict occurs because the new
			// `event_notification_configs` property (plural, a TypeList) is a TypeList, which we singularize.
			// To avoid the conflict, we override the nested type name for the new property, appending an "Item"
			// suffix.
			"event_notification_configs": {
				Elem: &tfbridge.SchemaInfo{
					NestedType: "RegistryEventNotificationConfigItem",
				},
			},
		},
		Docs: &tfbridge.DocInfo{
			Source: "cloudiot_registry.html.markdown",
		},
	})

	prov.SetAutonaming(255, "-")

	return prov
}
