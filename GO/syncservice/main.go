package main

import (
	"fmt"
	"justasking/GO/core/domain/appconfigs"
	"log"
	"strconv"
	"time"

	"justasking/GO/core/domain/applogs"
	"justasking/GO/core/domain/sync"
	"justasking/GO/core/startup/boot"
	"justasking/GO/core/startup/env"
)

var serviceName = "SyncService"

func main() {
	// Load the configuration file
	config, err := env.LoadConfig("config.json")
	if err != nil {
		log.Fatalln(err)
	}

	// Register the services
	boot.RegisterServices(config)

	interval, err := strconv.Atoi(config.Settings["IntervalInMinutes"])
	if err == nil {
		ticker := time.NewTicker(time.Duration(interval) * time.Minute)
		quit := make(chan struct{})
		for {
			select {
			case <-ticker.C:
				StartSync()
			case <-quit:
				ticker.Stop()
				return
			}
		}
	}
}

// StartSync is the main entry point for the sync
func StartSync() {
	functionName := "StartSync"

	applogsdomain.LogInfo(serviceName, functionName, "Starting SyncService Iteration")

	configs, configsResponse := appconfigsdomain.GetAppConfigs("syncservice")
	if configsResponse.IsSuccess() {
		runPlanExpirationSync, _ := strconv.ParseBool(configs["RunPlanExpirationSync"])
		runPhoneNumbersSync, _ := strconv.ParseBool(configs["RunPhoneNumbersSync"])

		if runPlanExpirationSync {
			//1. Sync Canceled plans
			syncPlansResponse := syncdomain.CancelExpiredPlans()
			if syncPlansResponse.IsSuccess() {
				applogsdomain.LogInfo(serviceName, functionName, "CancelExpiredPlans ran successfully.")
			}
		} else {
			applogsdomain.LogInfo(serviceName, functionName, "Skipping plan expiration sync.")
		}

		if runPhoneNumbersSync {
			//2. Sync Phone Numbers
			phoneNumbersThreshold, _ := strconv.Atoi(configs["PhoneNumbersThreshold"])
			syncPhoneNumbersResponse := syncdomain.SyncPhoneNumbers(phoneNumbersThreshold)
			if syncPhoneNumbersResponse.IsSuccess() {
				applogsdomain.LogInfo(serviceName, functionName, "SyncPhoneNumbers ran successfully.")
			}
		} else {
			applogsdomain.LogInfo(serviceName, functionName, "Skipping phone numbers sync.")
		}

	} else {
		applogsdomain.LogError(serviceName, functionName, fmt.Sprintf("Could not retrieve configs for SyncService. Message: [%v]", configsResponse.Message), true)
	}

	applogsdomain.LogInfo(serviceName, functionName, "Ending SyncService Iteration")
}
