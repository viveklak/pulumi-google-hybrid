// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.GoogleNative.OSConfig.V1Beta.Inputs
{

    /// <summary>
    /// A software recipe is a set of instructions for installing and configuring a piece of software. It consists of a set of artifacts that are downloaded, and a set of steps that install, configure, and/or update the software. Recipes support installing and updating software from artifacts in the following formats: Zip archive, Tar archive, Windows MSI, Debian package, and RPM package. Additionally, recipes support executing a script (either defined in a file or directly in this api) in bash, sh, cmd, and powershell. Updating a software recipe If a recipe is assigned to an instance and there is a recipe with the same name but a lower version already installed and the assigned state of the recipe is `UPDATED`, then the recipe is updated to the new version. Script Working Directories Each script or execution step is run in its own temporary directory which is deleted after completing the step.
    /// </summary>
    public sealed class SoftwareRecipeArgs : Pulumi.ResourceArgs
    {
        [Input("artifacts")]
        private InputList<Inputs.SoftwareRecipeArtifactArgs>? _artifacts;

        /// <summary>
        /// Resources available to be used in the steps in the recipe.
        /// </summary>
        public InputList<Inputs.SoftwareRecipeArtifactArgs> Artifacts
        {
            get => _artifacts ?? (_artifacts = new InputList<Inputs.SoftwareRecipeArtifactArgs>());
            set => _artifacts = value;
        }

        /// <summary>
        /// Default is INSTALLED. The desired state the agent should maintain for this recipe. INSTALLED: The software recipe is installed on the instance but won't be updated to new versions. UPDATED: The software recipe is installed on the instance. The recipe is updated to a higher version, if a higher version of the recipe is assigned to this instance. REMOVE: Remove is unsupported for software recipes and attempts to create or update a recipe to the REMOVE state is rejected.
        /// </summary>
        [Input("desiredState")]
        public Input<Pulumi.GoogleNative.OSConfig.V1Beta.SoftwareRecipeDesiredState>? DesiredState { get; set; }

        [Input("installSteps")]
        private InputList<Inputs.SoftwareRecipeStepArgs>? _installSteps;

        /// <summary>
        /// Actions to be taken for installing this recipe. On failure it stops executing steps and does not attempt another installation. Any steps taken (including partially completed steps) are not rolled back.
        /// </summary>
        public InputList<Inputs.SoftwareRecipeStepArgs> InstallSteps
        {
            get => _installSteps ?? (_installSteps = new InputList<Inputs.SoftwareRecipeStepArgs>());
            set => _installSteps = value;
        }

        /// <summary>
        /// Unique identifier for the recipe. Only one recipe with a given name is installed on an instance. Names are also used to identify resources which helps to determine whether guest policies have conflicts. This means that requests to create multiple recipes with the same name and version are rejected since they could potentially have conflicting assignments.
        /// </summary>
        [Input("name", required: true)]
        public Input<string> Name { get; set; } = null!;

        [Input("updateSteps")]
        private InputList<Inputs.SoftwareRecipeStepArgs>? _updateSteps;

        /// <summary>
        /// Actions to be taken for updating this recipe. On failure it stops executing steps and does not attempt another update for this recipe. Any steps taken (including partially completed steps) are not rolled back.
        /// </summary>
        public InputList<Inputs.SoftwareRecipeStepArgs> UpdateSteps
        {
            get => _updateSteps ?? (_updateSteps = new InputList<Inputs.SoftwareRecipeStepArgs>());
            set => _updateSteps = value;
        }

        /// <summary>
        /// The version of this software recipe. Version can be up to 4 period separated numbers (e.g. 12.34.56.78).
        /// </summary>
        [Input("version")]
        public Input<string>? Version { get; set; }

        public SoftwareRecipeArgs()
        {
        }
    }
}
