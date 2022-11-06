package describe

import (
	"fmt"
	"os"
	"strings"
	"text/tabwriter"

	models "github.com/avocadohooman/logistic-regression-magic-sorting-hat/pkg/models/stats"
)

func PrintData(stats models.StatsArray) {
	headers := prepareRow("")
	count := prepareRow("Count")
	mean := prepareRow("Mean")
	std := prepareRow("Std")
	min := prepareRow("Min")
	max := prepareRow("Max")
	twentyFive := prepareRow("25%")
	fifty := prepareRow("50%")
	seventyFive := prepareRow("75%")

	for _, stat := range stats {
		headers = append(headers, strings.ToUpper(stat.FeatureName[0:3]))
		count = append(count, fmt.Sprintf("%d", stat.Count))
		mean = append(mean, fmt.Sprintf("%.2f", stat.Mean))
		std = append(std, fmt.Sprintf("%.2f", stat.Std))
		min = append(min, fmt.Sprintf("%.2f", stat.Min))
		max = append(max, fmt.Sprintf("%.2f", stat.Max))
		twentyFive = append(twentyFive, fmt.Sprintf("%.2f", stat.TwentyFive))
		fifty = append(fifty, fmt.Sprintf("%.2f", stat.Fifty))
		seventyFive = append(seventyFive, fmt.Sprintf("%.2f", stat.SeventyFive))
	}

	tabFeatures := strings.Join(headers, "\t")
	tabMean := strings.Join(mean, "\t")
	tabCount := strings.Join(count, "\t")
	tabStd := strings.Join(std, "\t")
	tabMin := strings.Join(min, "\t")
	tabMax := strings.Join(max, "\t")
	tabTwentFive := strings.Join(twentyFive, "\t")
	tabFifty := strings.Join(fifty, "\t")
	tabSeventyFive := strings.Join(seventyFive, "\t")

	const padding = 2
	tabWriter := tabwriter.NewWriter(os.Stdout, 0, 0, padding, ' ', tabwriter.Debug)
	
	fmt.Fprintln(tabWriter, tabFeatures)
	fmt.Fprintln(tabWriter, tabCount)
	fmt.Fprintln(tabWriter, tabMean)
	fmt.Fprintln(tabWriter, tabStd)
	fmt.Fprintln(tabWriter, tabMin)
	fmt.Fprintln(tabWriter, tabMax)
	fmt.Fprintln(tabWriter, tabTwentFive)
	fmt.Fprintln(tabWriter, tabFifty)
	fmt.Fprintln(tabWriter, tabSeventyFive)

	tabWriter.Flush()
}

func prepareRow(describeLabel string) []string {
	var describeRow []string
	describeRow = append(describeRow, describeLabel)
	return describeRow
}
