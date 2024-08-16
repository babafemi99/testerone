package table

import (
	"context"
	"encoding/json"
	"github.com/babafemi99/testerone/load"
	"github.com/olekukonko/tablewriter"
	"log"
	"os"
	"testing"
	"time"
)

//func TestRenderTable(t *testing.T) {
//	req := load.Req{
//		NumberOfRequests: 1500,
//		//URL:              "http://localhost:1010/ping",
//		//URL:      "https://www.google.com/",
//		URL:      "http://localhost:2020",
//		Interval: 1,
//	}
//	//data, _ := req.()
//
//	RenderTable(data)
//	log.Println(data)
//}

func TestRenderTable1(t *testing.T) {

	data := [][]string{
		{"A", "The Good", "500"},
		{"B", "The Very very Bad Man", "288"},
		{"C", "The Ugly", "120"},
		{"D", "The Gopher", "800"},
	}

	table := tablewriter.NewWriter(os.Stdout)
	table.SetHeader([]string{"Name", "Sign", "Rating"})

	table.AppendBulk(data)
	table.Render()
}

func TestCustomReq2(t *testing.T) {
	b1 := struct {
		Name  string `json:"name"`
		Email string `json:"email"`
		Token string `json:"token"`
	}{
		Name:  "Termiii",
		Email: "tt@tikabodi.com",
		Token: "45yuhgdfrtyuiwop098uytghjko98w7yethjdiop098yutghjk",
	}

	jb1, err := json.Marshal(b1)
	if err != nil {
		return
	}

	b2 := struct {
		Title string `json:"title"`
		Body  string `json:"body"`
	}{
		Title: "Test me out",
		Body:  "Are you really doing this  ?",
	}

	jb2, err := json.Marshal(b2)
	if err != nil {
		return
	}

	b3 := struct {
		Age    int    `json:"age"`
		Gender string `json:"gender"`
	}{
		Age:    25,
		Gender: "Male",
	}

	jb3, err := json.Marshal(b3)
	if err != nil {
		return
	}

	req22 := load.CustomReq{
		NumberOfRequests: 1,
		Interval:         1,
		Func2: []load.CustomFunction{
			{
				Method: "POST",
				URL:    "http://localhost:1010/post1",
				Body:   jb1,
			},
			{
				Method: "POST",
				URL:    "http://localhost:1010/post2",
				Body:   jb2,
			},
			{
				Method: "POST",
				URL:    "http://localhost:1010/post3",
				Body:   jb3,
			},
		},
	}
	run, err := req22.Run()
	if err != nil {
		log.Println("error is ", err)
	}
	RenderTable(run)
}

func TestCustomReq3(t *testing.T) {
	req22 := load.CustomReq{
		NumberOfRequests: 1000,
		Interval:         100,
		Func2: []load.CustomFunction{
			{
				Method: "GET",
				URL:    "http://localhost:1010/ping",
				Body:   nil,
			},
			{
				Method: "GET",
				URL:    "http://localhost:1010/ping",
				Body:   nil,
			},
			{
				Method: "GET",
				URL:    "http://localhost:1010/ping",
				Body:   nil,
			},
		},
	}
	run, err := req22.Run()
	if err != nil {
		log.Println("error is ", err)
	}
	RenderTable(run)
}

func TestReq_RunX(t *testing.T) {
	req := load.Req{
		NumberOfRequests: 100,
		URL:              "http://localhost:1010/ping",
		Interval:         10,
		RunAfterDuration: 5 * time.Second,
		RunDuration:      10,
	}
	after, err := req.RunAfter()
	if err != nil {
		t.Errorf("error")
	}

	for i, a := range after {
		log.Println("Iteration: ", i+1)
		RenderTable(a)
	}
}
func TestReq_RunX_Custom(t *testing.T) {
	b1 := struct {
		Name  string `json:"name"`
		Email string `json:"email"`
		Token string `json:"token"`
	}{
		Name:  "Termiii",
		Email: "tt@tikabodi.com",
		Token: "45yuhgdfrtyuiwop098uytghjko98w7yethjdiop098yutghjk",
	}

	jb1, err := json.Marshal(b1)
	if err != nil {
		return
	}

	b2 := struct {
		Title string `json:"title"`
		Body  string `json:"body"`
	}{
		Title: "Test me out",
		Body:  "Are you really doing this  ?",
	}

	jb2, err := json.Marshal(b2)
	if err != nil {
		return
	}

	b3 := struct {
		Age    int    `json:"age"`
		Gender string `json:"gender"`
	}{
		Age:    25,
		Gender: "Male",
	}

	jb3, err := json.Marshal(b3)
	if err != nil {
		return
	}

	req22 := load.CustomReq{
		NumberOfRequests: 100,
		Interval:         10,
		Func2: []load.CustomFunction{
			{
				Method: "POST",
				URL:    "http://localhost:1010/post1",
				Body:   jb1,
			},
			{
				Method: "POST",
				URL:    "http://localhost:1010/post2",
				Body:   jb2,
			},
			{
				Method: "POST",
				URL:    "http://localhost:1010/post3",
				Body:   jb3,
			},
		},
		RunAfterDuration: 10 * time.Nanosecond,
		RunDuration:      10,
	}
	now := time.Now()
	timeout, cancelFunc := context.WithTimeout(context.Background(), 10*time.Minute)
	defer cancelFunc()
	run, err := req22.RunAfterWithContext(timeout)
	if err != nil {
		log.Println("error is ", err)
	}
	log.Println("total duration is", time.Since(now))

	for i, a := range run {
		log.Println("Iteration: ", i+1)
		RenderTable(a)
	}
}
