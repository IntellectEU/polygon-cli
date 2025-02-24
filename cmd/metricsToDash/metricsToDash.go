/*
Copyright © 2022 Polygon <engineering@polygon.technology>

This program is free software: you can redistribute it and/or modify
it under the terms of the GNU Lesser General Public License as published by
the Free Software Foundation, either version 3 of the License, or
(at your option) any later version.

This program is distributed in the hope that it will be useful,
but WITHOUT ANY WARRANTY; without even the implied warranty of
MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
GNU Lesser General Public License for more details.

You should have received a copy of the GNU Lesser General Public License
along with this program. If not, see <http://www.gnu.org/licenses/>.
*/
package metricsToDash

import (
	"fmt"

	"github.com/maticnetwork/polygon-cli/dashboard"
	"github.com/spf13/cobra"
)

var (
	inputMetricsToDashFile                *string
	inputMetricsToDashPrefix              *string
	inputMetricsToDashTitle               *string
	inputMetricsToDashDesc                *string
	inputMetricsToDashHeight              *int
	inputMetricsToDashWidth               *int
	inputMetricsToDashTemplateVars        *[]string
	inputMetricsToDashTemplateVarDefaults *[]string
	inputMetricsToDashStripPrefixes       *[]string
	inputMetricsToDashPretty              *bool
	inputMetricsToDashShowHelp            *bool
)

// MetricsToDashCmd represents the metricsToDash command
var MetricsToDashCmd = &cobra.Command{
	Use:     "metrics-to-dash",
	Aliases: []string{"metricstodash", "metricsToDash"},
	Short:   "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	RunE: func(cmd *cobra.Command, args []string) error {
		do := dashboard.DashboardOptions{
			File:                *inputMetricsToDashFile,
			Title:               *inputMetricsToDashTitle,
			Prefix:              *inputMetricsToDashPrefix,
			Description:         *inputMetricsToDashDesc,
			WidgetHeight:        *inputMetricsToDashHeight,
			WidgetWidth:         *inputMetricsToDashWidth,
			TemplateVars:        *inputMetricsToDashTemplateVars,
			TemplateVarDefaults: *inputMetricsToDashTemplateVarDefaults,
			StripPrefixes:       *inputMetricsToDashStripPrefixes,
			Pretty:              *inputMetricsToDashPretty,
			ShowHelp:            *inputMetricsToDashShowHelp,
		}
		data, err := dashboard.ConvertMetricsToDashboard(&do)
		if err != nil {
			return err
		}
		fmt.Println(string(data))
		return nil
	},
}

func init() {
	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	inputMetricsToDashFile = MetricsToDashCmd.PersistentFlags().StringP("input-file", "i", "", "the metrics file to be used")
	inputMetricsToDashPrefix = MetricsToDashCmd.PersistentFlags().StringP("prefix", "p", "", "prefix to use before all metrics")
	inputMetricsToDashTitle = MetricsToDashCmd.PersistentFlags().StringP("title", "t", "Polycli Dashboard", "title for the dashboard")
	inputMetricsToDashDesc = MetricsToDashCmd.PersistentFlags().StringP("desc", "d", "Polycli Dashboard", "description for the dashboard")
	inputMetricsToDashWidth = MetricsToDashCmd.PersistentFlags().IntP("width", "W", 4, "widget width")
	inputMetricsToDashHeight = MetricsToDashCmd.PersistentFlags().IntP("height", "H", 3, "widget height")

	inputMetricsToDashTemplateVars = MetricsToDashCmd.PersistentFlags().StringArrayP("template-vars", "T", []string{}, "The template variables to use for the dashboard")
	inputMetricsToDashTemplateVarDefaults = MetricsToDashCmd.PersistentFlags().StringArrayP("template-var-defaults", "D", []string{}, "The defaults to use for the template variables")

	inputMetricsToDashStripPrefixes = MetricsToDashCmd.PersistentFlags().StringArrayP("strip-prefix", "s", []string{}, "A prefix that can be removed from the metrics")
	inputMetricsToDashPretty = MetricsToDashCmd.PersistentFlags().BoolP("pretty-name", "P", true, "Should the metric names be prettified")

	inputMetricsToDashShowHelp = MetricsToDashCmd.PersistentFlags().BoolP("show-help", "S", false, "Should we show the help text for each metric")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// MetricsToDashCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
