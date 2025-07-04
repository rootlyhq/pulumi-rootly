// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Rootly.Outputs
{

    [OutputType]
    public sealed class DashboardPanelParams
    {
        public readonly ImmutableArray<Outputs.DashboardPanelParamsDataset> Datasets;
        public readonly string Display;
        public readonly Outputs.DashboardPanelParamsLegend? Legend;

        [OutputConstructor]
        private DashboardPanelParams(
            ImmutableArray<Outputs.DashboardPanelParamsDataset> datasets,

            string display,

            Outputs.DashboardPanelParamsLegend? legend)
        {
            Datasets = datasets;
            Display = display;
            Legend = legend;
        }
    }
}
