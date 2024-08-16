package load

import (
	"context"
	"encoding/json"
	"testing"
	"time"
)

func TestCustomReq_loadCustomTarget2(t *testing.T) {

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

	req22 := CustomReq{
		NumberOfRequests: 700,
		Interval:         1,
		Func2: []CustomFunction{
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
	req22.Run()
}

func TestCustomReq_RunWithContext(t *testing.T) {
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

	req22 := CustomReq{
		NumberOfRequests: 100,
		Interval:         1,
		Func2: []CustomFunction{
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
		RunAfterDuration: 5 * time.Nanosecond,
		RunDuration:      6,
	}
	ctx, cancelFunc := context.WithTimeout(context.Background(), time.Minute*10)
	defer cancelFunc()
	_, err = req22.RunAfterWithContext(ctx)
	if err != nil {
		t.Errorf("error: %v", err)
	}
}
