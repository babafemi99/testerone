package table

import (
	"github.com/babafemi99/testerone/load"
	"github.com/olekukonko/tablewriter"
	"os"
	"testing"
)

func TestRenderTable(t *testing.T) {
	req := load.Req{
		NumberOfRequests: 250,
		URL:              "http://localhost:1010/ping",
		//URL:      "https://www.google.com/",
		//URL:      "http://localhost:2020/process",
		Interval: 1,
	}
	data, _ := req.Run()

	RenderTable(data)

}

func TestRenderTable1(t *testing.T) {
	data := [][]string{
		[]string{"A", "The Good", "500"},
		[]string{"B", "The Very very Bad Man", "288"},
		[]string{"C", "The Ugly", "120"},
		[]string{"D", "The Gopher", "800"},
	}

	table := tablewriter.NewWriter(os.Stdout)
	table.SetHeader([]string{"Name", "Sign", "Rating"})

	table.AppendBulk(data)
	table.Render()
}
