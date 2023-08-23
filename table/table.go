package table

import (
	"fmt"
	"github.com/babafemi99/testerone/load"
	"github.com/olekukonko/tablewriter"
	"os"
)

func RenderTable(data load.ResponseData) {

	fmt.Printf("response average time :%.4fs\n", data.AverageResponseTime)
	fmt.Printf("response error rate: %.2f%%\n", data.ErrorRate)
	fmt.Printf("response success rate: %.2f%%\n", data.SuccessRate)
	fmt.Printf("response maximum time :%.4fs\n", data.MaximumTime)
	fmt.Printf("response minimum time :%.4fs\n", data.MinimumTime)

	arr := processDataArr(data.Responses)
	table := tablewriter.NewWriter(os.Stdout)
	table.SetHeader([]string{"Index", "Time Taken", "Throughput"})
	table.AppendBulk(arr)
	//table.SetFooter([]string{"", "Avg Time", fmt.Sprintf("%.2f", data.AverageResponseTime)})
	//table.SetFooter([]string{"", "Min Time", fmt.Sprintf("%.2f", data.MinimumTime)})
	//table.SetFooter([]string{"", "Max Time", fmt.Sprintf("%.2f", data.MaximumTime)})
	//table.SetFooter([]string{"", "Success Rate", fmt.Sprintf("%.2f %", data.SuccessRate)})
	//table.SetFooter([]string{"", "Error Rate", fmt.Sprintf("%.2f %", data.ErrorRate)})
	table.Render()
}
func processDataArr(data []load.ResponseTime) [][]string {
	var final [][]string

	for i, val := range data {
		row := []string{
			fmt.Sprintf("%d", i+1),
			fmt.Sprintf("%.4f", val.Time),
			fmt.Sprintf("%t", val.Success),
		}
		final = append(final, row)
	}

	return final

}
