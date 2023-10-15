package main_test

import (
	main "home_sensor_lambda/cmd/home_sensor_lambda"
	"home_sensor_lambda/test/assertions"
	"home_sensor_lambda/test/mocks"
	"testing"
	"time"
)

func TestProcess(t *testing.T) {
	mockSnsPublisher := mocks.SnsManagerSuccess()
	request := main.Request{
		Headers: map[string]string{
			"X-Timestamp": time.Now().String(),
		},
	}
	statusCode, body, err := main.Process(request, mockSnsPublisher)

	if statusCode != 200 {
		t.Fatalf("expected 200, got %d", statusCode)
	}

	if body != "OK" {
		t.Fatalf("expected OK, got %s", body)
	}

	assertions.AssertEquals(t, statusCode, 200)
	assertions.AssertEquals(t, body, "OK")
	assertions.AssertEquals(t, err, nil)
}
