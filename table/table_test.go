package table

import (
	"github.com/babafemi99/testerone/load"
	"github.com/olekukonko/tablewriter"
	"net/http"
	"os"
	"testing"
)

func TestRenderTable(t *testing.T) {
	req := load.Req{
		NumberOfRequests: 50,
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

func TestCustomReq(t *testing.T) {
	req, err := http.NewRequest("GET", "http://localhost:1010/ping", nil)
	if err != nil {
		return
	}

	req22 := load.Req{
		ReqType:          "custom",
		NumberOfRequests: 250,
		URL:              "",
		Interval:         0,
		Func:             req,
	}
	run, _ := req22.Run()
	RenderTable(run)
}
