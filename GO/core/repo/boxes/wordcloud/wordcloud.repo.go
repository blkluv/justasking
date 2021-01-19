package wordcloudrepo

import (
	"justasking/GO/core/model/boxes/basebox"
	"justasking/GO/core/model/boxes/wordcloud"
	"justasking/GO/core/model/wordcloudresponse"
	"justasking/GO/core/repo/boxes/basebox"
	"justasking/GO/core/startup/flight"

	"time"

	"github.com/satori/go.uuid"
)

// GetWordCloudByBoxId returns wordcloud
func GetWordCloudByBoxId(guid uuid.UUID) (wordcloudmodel.WordCloud, error) {
	db := flight.Context(nil, nil).DB

	baseBox := baseboxmodel.BaseBox{}
	wordCloud := wordcloudmodel.WordCloud{}

	//get basebox
	baseBox, err := baseboxrepo.GetBaseBoxById(guid)
	if err != nil {
		return wordCloud, err
	}
	//get wordcloud
	err = db.Where("box_id = ?", guid).Find(&wordCloud).Error
	if err != nil {
		return wordCloud, err
	}

	//assign wordcloud's basebox
	wordCloud.BaseBox = baseBox

	return wordCloud, err
}

// GetWordCloudByBoxCode returns wordcloud
func GetWordCloudByCode(code string) (wordcloudmodel.WordCloud, error) {
	db := flight.Context(nil, nil).DB

	baseBox := baseboxmodel.BaseBox{}
	wordCloud := wordcloudmodel.WordCloud{}

	//get basebox
	baseBox, err := baseboxrepo.GetBaseBoxByCode(code)
	if err != nil {
		return wordCloud, err
	}
	//get wordcloud
	err = db.Where("box_id = ?", baseBox.ID).Find(&wordCloud).Error
	if err != nil {
		return wordCloud, err
	}

	//assign wordcloud's basebox
	wordCloud.BaseBox = baseBox

	return wordCloud, err
}

// InsertWordCloud creates a base box and wordcloud
func InsertWordCloud(wordCloud wordcloudmodel.WordCloud) error {
	db := flight.Context(nil, nil).DB

	// Wrapping box creation in a transaction. We wouldn't want the WordCloud to be created without the BaseBox
	tx := db.Begin()

	wordCloud.BaseBox.ID = wordCloud.BoxId
	baseBox := wordCloud.BaseBox

	if err := baseboxrepo.InsertBaseBox(baseBox, tx); err != nil {
		tx.Rollback()
		return err
	}

	if err := tx.Create(&wordCloud).Error; err != nil {
		tx.Rollback()
		return err
	}

	tx.Commit()
	return nil
}

// InsertWordCloudResponse inserts a response to a word cloud
func InsertWordCloudResponse(response wordcloudresponsemodel.WordCloudResponse) (wordcloudresponsemodel.WordCloudResponse, error) {
	db := flight.Context(nil, nil).DB

	err := db.Create(&response).Error

	return response, err
}

// GetWordCloudResponsesByBoxId gets all answers for a specific box
func GetWordCloudResponsesByBoxId(guid uuid.UUID) ([]wordcloudresponsemodel.WordCloudResponse, error) {
	db := flight.Context(nil, nil).DB

	answers := []wordcloudresponsemodel.WordCloudResponse{}
	err := db.Raw(`SELECT box_id,response,is_hidden,created_at,created_by,updated_at,updated_by,deleted_at FROM word_cloud_box_responses 
		WHERE box_id = ?
		Order by created_at asc`, guid).Scan(&answers).Error

	return answers, err
}

// GetWordCloudResponsesByCode gets all answers for a specific box
func GetWordCloudResponsesByCode(code string) ([]wordcloudresponsemodel.WordCloudResponse, error) {
	db := flight.Context(nil, nil).DB

	answers := []wordcloudresponsemodel.WordCloudResponse{}
	err := db.Raw(`SELECT box_id,response,is_hidden,created_at,created_by,updated_at,updated_by,deleted_at FROM word_cloud_box_responses 
		WHERE box_id = (SELECT id FROM base_box WHERE code = ?)
		Order by created_at asc`, code).Scan(&answers).Error

	return answers, err
}

// HideAnswer hides the given answer for a given box
func HideAnswer(guid uuid.UUID, answerToHide string, updatedBy string) error {
	db := flight.Context(nil, nil).DB

	err := db.Exec("UPDATE word_cloud_box_responses SET is_hidden = 1, updated_at = ?, updated_by = ? WHERE box_id = ? AND response = ?", time.Now(), updatedBy, guid, answerToHide).Error

	return err
}

// UnhideAnswer hides the given answer for a given box
func UnhideAnswer(guid uuid.UUID, answerToHide string, updatedBy string) error {
	db := flight.Context(nil, nil).DB

	err := db.Exec("UPDATE word_cloud_box_responses SET is_hidden = 0, updated_at = ?, updated_by = ? WHERE box_id = ? AND response = ?", time.Now(), updatedBy, guid, answerToHide).Error

	return err
}

// HideAllAnswers hides all answers for a given wordcloud
func HideAllAnswers(guid uuid.UUID, updatedBy string) error {
	db := flight.Context(nil, nil).DB

	err := db.Exec("UPDATE word_cloud_box_responses SET is_hidden = 1, updated_at = ?, updated_by = ? WHERE box_id = ?", time.Now(), updatedBy, guid).Error

	return err
}

// UnhideAllAnswers hides all answers for a given wordcloud
func UnhideAllAnswers(guid uuid.UUID, updatedBy string) error {
	db := flight.Context(nil, nil).DB

	err := db.Exec("UPDATE word_cloud_box_responses SET is_hidden = 0, updated_at = ?, updated_by = ? WHERE box_id = ?", time.Now(), updatedBy, guid).Error

	return err
}
