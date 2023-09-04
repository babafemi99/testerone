package yaml

import (
	"testing"
	"time"
)

func ConvertByteSliceToString(byteSlice []byte) string {
	var convertedString string
	for _, b := range byteSlice {
		convertedString += string(b)
	}
	return convertedString
}

func TestLoadYAMLFile(t *testing.T) {
	filePath := "config.yaml"

	t.Run("Test LoadYAMLFile", func(t *testing.T) {
		custom, basic, err := LoadYAMLFile(filePath)
		if err != nil {
			t.Errorf("LoadYAMLFile() error = %v", err)
			return
		}
		if custom != nil {
			if custom.NumberOfRequests != 200 {
				t.Errorf("Expected NumberOfRequests 200, but got %d", custom.NumberOfRequests)
			}

			if custom.Interval != 10 {
				t.Errorf("Expected Interval 10, but got %d", custom.Interval)
			}
			if custom.RunAfterDuration != 10*time.Nanosecond {
				t.Errorf("Expected RunAfterDuration 10ns, but got %v", custom.RunAfterDuration)
			}
			if custom.RunDuration != 500 {
				t.Errorf("Expected RunDuration 500, but got %d", custom.RunDuration)
			}

			// Assertions for the first CustomFunction in Func2
			if len(custom.Func2) != 3 {
				t.Errorf("Expected 3 elements in Func2, but got %d", len(custom.Func2))
			}
			cf1 := custom.Func2[0]
			if cf1.Method != "POST" {
				t.Errorf("Expected Method 'POST', but got '%s'", cf1.Method)
			}
			if cf1.URL != "http://localhost:1010/post1" {
				t.Errorf("Expected URL 'http://localhost:1010/post1', but got '%s'", cf1.URL)
			}
			// Assertions for cf1.Body

			if (string(cf1.Body)) != `{"email":"tt@tikabodi.com","name":"Termiii","token":"45yuhgdfrtyuiwop098uytghjko98w7yethjdiop098yutghjk"}` {
				t.Errorf("Unexpected body for cf1: %s", string(cf1.Body))
			}

			// Assertions for the second CustomFunction in Func2
			cf2 := custom.Func2[1]
			if cf2.Method != "POST" {
				t.Errorf("Expected Method 'POST', but got '%s'", cf2.Method)
			}
			if cf2.URL != "http://localhost:1010/post2" {
				t.Errorf("Expected URL 'http://localhost:1010/post2', but got '%s'", cf2.URL)
			}
			// Assertions for cf2.Body
			if (string(cf2.Body)) != `{"body":"Are you really doing this  ?","title":"Test me out"}` {
				t.Errorf("Unexpected body for cf2: %s", cf2.Body)
			}

			// Assertions for the third CustomFunction in Func2
			cf3 := custom.Func2[2]
			if cf3.Method != "POST" {
				t.Errorf("Expected Method 'POST', but got '%s'", cf3.Method)
			}
			if cf3.URL != "http://localhost:1010/post3" {
				t.Errorf("Expected URL 'http://localhost:1010/post3', but got '%s'", cf3.URL)
			}
			// Assertions for cf3.Body
			if (string(cf3.Body)) != `{"age":25,"gender":"Male"}` {
				t.Errorf("Unexpected body for cf3: %s", cf3.Body)
			}
		}
		if basic != nil {
			if basic.NumberOfRequests != 200 {
				t.Errorf("Expected NumberOfRequests 200, but got %d", basic.NumberOfRequests)
			}

			if basic.Interval != 10 {
				t.Errorf("Expected Interval 10, but got %d", basic.Interval)
			}

			if basic.URL != "https://google.com" {
				t.Errorf("Expected URL 'https://google.com', but got '%s'", basic.URL)
			}
		}
	})
}
