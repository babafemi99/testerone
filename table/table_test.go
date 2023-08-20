package table

import (
	"encoding/base64"
	"encoding/json"
	"fmt"
	"github.com/babafemi99/testerone/load"
	"log"
	"testing"
)

func TestRenderTable(t *testing.T) {
	req := load.Req{
		NumberOfRequests: 1500,
		//URL:              "http://localhost:1010/ping",
		//URL:      "https://www.google.com/",
		URL:      "http://localhost:2020",
		Interval: 1,
	}
	data, _ := req.Run()

	RenderTable(data)

}

func TestRenderTable1(t *testing.T) {
	type Data struct {
		BinaryData []byte `json:"binaryData"`
	}
	binaryData := []byte{1, 2, 3, 4, 5}

	encodedData := base64.StdEncoding.EncodeToString(binaryData)

	data := Data{
		BinaryData: []byte(encodedData),
	}

	jsonBytes, err := json.Marshal(data)
	if err != nil {
		fmt.Println("Error marshaling struct:", err)
		return
	}

	fmt.Println(string(jsonBytes))

	//data := [][]string{
	//	[]string{"A", "The Good", "500"},
	//	[]string{"B", "The Very very Bad Man", "288"},
	//	[]string{"C", "The Ugly", "120"},
	//	[]string{"D", "The Gopher", "800"},
	//}
	//
	//table := tablewriter.NewWriter(os.Stdout)
	//table.SetHeader([]string{"Name", "Sign", "Rating"})
	//
	//table.AppendBulk(data)
	//table.Render()
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
		ReqType:          "custom",
		NumberOfRequests: 1,
		URL:              "",
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
		ReqType:          "custom",
		NumberOfRequests: 1000,
		URL:              "",
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
