package main

import (
	"testing"

	steps "banking-system/features/steps"

	"github.com/cucumber/godog"
)

func TestBDDFeatures(t *testing.T) {
	opts := godog.Options{
		Format: "pretty", // pretty output
		Paths: []string{
			"./features/user.feature",
			"./features/deposit.feature",
			"./features/withdraw.feature",
			"./features/banking_lifecycle.feature",
		},
		Randomize: 0, // run scenarios in order
	}

	// Create TestSuite
	status := godog.TestSuite{
		Name:                "banking-system-bdd",
		ScenarioInitializer: steps.InitializeScenario, // your steps in steps.go
		Options:             &opts,
	}.Run()

	if status != 0 {
		t.Fatal("there are failed scenarios")
	}
}
