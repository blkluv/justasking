package tests

import (
	"fmt"
	"log"
	"math/rand"
	"testing"
	"time"

	"github.com/chande/justasking/core/startup/boot"
	"github.com/chande/justasking/core/startup/env"
)

func init() {
	// Load the configuration file
	config, err := env.LoadConfig("config.json")
	if err != nil {
		log.Fatalln(err)
	}

	// Register the services
	boot.RegisterServices(config)
}

func XTestBaseBoxInsert(t *testing.T) {
	var baseBox baseboxmodel.BaseBox
	var wordCloud wordcloudmodel.WordCloud

	s1 := rand.NewSource(time.Now().UnixNano())
	r1 := rand.New(s1)

	baseBox.Code = fmt.Sprintf("TestBox%v", r1.Intn(10000))
	baseBox.OwnerId = 1
	baseBox.BoxType = 1
	baseBox.CreatedBy = "1"
	baseBox.UpdatedBy = "1"
	wordCloud.BaseBox = baseBox

	err := wordcloudrepo.InsertWordCloud(wordCloud)

	if err != nil {
		t.Error("error is ", err)
	}
}

func TestWordCloudInsert(t *testing.T) {
	var wordCloud wordcloudmodel.WordCloud

	s1 := rand.NewSource(time.Now().UnixNano())
	r1 := rand.New(s1)

	wordCloud.BaseBox.Code = fmt.Sprintf("TestBox%v", r1.Intn(10000))
	wordCloud.BaseBox.OwnerId = 1
	wordCloud.BaseBox.BoxType = 1
	wordCloud.BaseBox.CreatedBy = "1"
	wordCloud.BaseBox.UpdatedBy = "1"
	wordCloud.Header = "Is This a test?"
	wordCloud.ThemeId = 1

	err := wordcloudrepo.InsertWordCloud(wordCloud)

	if err != nil {
		t.Error("error is ", err)
	}
}
