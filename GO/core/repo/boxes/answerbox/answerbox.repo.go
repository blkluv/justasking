package answerboxrepo

import (
	"justasking/GO/core/model/answerboxentry"
	"justasking/GO/core/model/answerboxquestion"
	"justasking/GO/core/model/boxes/answerbox"
	"justasking/GO/core/model/boxes/basebox"
	"justasking/GO/core/repo/boxes/basebox"
	"justasking/GO/core/startup/flight"
	"time"

	uuid "github.com/satori/go.uuid"
)

// InsertAnswerBox inserts a new answer box
func InsertAnswerBox(answerBox answerboxmodel.AnswerBox) error {
	db := flight.Context(nil, nil).DB

	// Wrapping box creation in a transaction. We wouldn't want the AnswerBox to be created without the BaseBox
	tx := db.Begin()

	answerBox.BaseBox.ID = answerBox.BoxId
	baseBox := answerBox.BaseBox

	if err := baseboxrepo.InsertBaseBox(baseBox, tx); err != nil {
		tx.Rollback()
		return err
	}

	if err := tx.Create(&answerBox).Error; err != nil {
		tx.Rollback()
		return err
	}

	//looping over every question for this answer box
	var err error
	for _, question := range answerBox.Questions {
		question.BoxId = answerBox.BaseBox.ID
		question.QuestionId, _ = uuid.NewV4()
		question.CreatedBy = baseBox.CreatedBy
		err = tx.Create(&question).Error
		if err != nil {
			break
		}
	}

	if err != nil {
		tx.Rollback()
		return err
	}

	tx.Commit()
	return nil
}

// GetAnswerBoxByBoxId gets an answer box by id
func GetAnswerBoxByBoxId(guid uuid.UUID) (answerboxmodel.AnswerBox, error) {
	db := flight.Context(nil, nil).DB

	baseBox := baseboxmodel.BaseBox{}
	answerBox := answerboxmodel.AnswerBox{}
	answerBoxQuestions := []answerboxquestionmodel.AnswerBoxQuestion{}

	//get basebox
	baseBox, err := baseboxrepo.GetBaseBoxById(guid)
	if err != nil {
		return answerBox, err
	}

	//get answerbox
	err = db.Where("box_id = ?", guid).Find(&answerBox).Error
	if err != nil {
		return answerBox, err
	}

	//get answerboxquestions
	err = db.Where("box_id = ?", guid).Order("sort_order").Find(&answerBoxQuestions).Error
	if err != nil {
		return answerBox, err
	}

	//assign answerBox's basebox and questions
	answerBox.BaseBox = baseBox
	answerBox.Questions = answerBoxQuestions

	return answerBox, err
}

// GetAnswerBoxByBoxCode returns an answerbox
func GetAnswerBoxByBoxCode(code string) (answerboxmodel.AnswerBox, error) {
	db := flight.Context(nil, nil).DB

	answerBox := answerboxmodel.AnswerBox{}
	answerBoxQuestions := []answerboxquestionmodel.AnswerBoxQuestion{}

	//get basebox
	baseBox, err := baseboxrepo.GetBaseBoxByCode(code)
	if err != nil {
		return answerBox, err
	}
	//get answerbox
	err = db.Where("box_id = ?", baseBox.ID).Find(&answerBox).Error
	if err != nil {
		return answerBox, err
	}

	//get answerboxquestions
	err = db.Where("box_id = ?", baseBox.ID).Order("sort_order").Find(&answerBoxQuestions).Error
	if err != nil {
		return answerBox, err
	}

	//assign answerBox's basebox and questions
	answerBox.BaseBox = baseBox
	answerBox.Questions = answerBoxQuestions

	return answerBox, err
}

// InsertAnswerBoxEntry adds an entry to the question box entries table
func InsertAnswerBoxEntry(entry answerboxentrymodel.AnswerBoxEntry) error {
	db := flight.Context(nil, nil).DB

	err := db.Exec(`INSERT INTO answer_box_entries (entry_id, question_id, entry, is_hidden, created_by) 
		VALUES (?, ?, ?, ?, ?)`,
		entry.EntryId, entry.QuestionId, entry.Entry, entry.IsHidden, entry.CreatedBy).Error

	return err
}

// GetAnswerBoxEntriesByBoxId gets all entries for a specific answer box
func GetAnswerBoxEntriesByBoxId(guid uuid.UUID) ([]answerboxentrymodel.AnswerBoxEntry, error) {
	db := flight.Context(nil, nil).DB

	entries := []answerboxentrymodel.AnswerBoxEntry{}

	err := db.Raw(`SELECT e.entry_id, e.question_id, e.entry, e.is_hidden, e.created_by, e.created_at
				FROM answer_box_entries e JOIN answer_box_questions q ON e.question_id = q.question_id JOIN base_box b ON b.id = q.box_id
				WHERE b.id = ? 
				Order by created_at asc`, guid).Scan(&entries).Error

	return entries, err
}

// GetAnswerBoxEntriesByCode gets all entries for a specific answer box
func GetAnswerBoxEntriesByCode(code string) ([]answerboxentrymodel.AnswerBoxEntry, error) {
	db := flight.Context(nil, nil).DB

	entries := []answerboxentrymodel.AnswerBoxEntry{}

	err := db.Raw(`SELECT e.entry_id, e.question_id, e.entry, e.is_hidden, e.created_by, e.created_at
				FROM answer_box_entries e JOIN answer_box_questions q ON e.question_id = q.question_id JOIN base_box b ON b.id = q.box_id
				WHERE b.code = ? 
				Order by created_at asc`, code).Scan(&entries).Error

	return entries, err
}

// GetVisibleAnswerBoxEntriesByCode gets all entries for a specific question box
func GetVisibleAnswerBoxEntriesByCode(code string) ([]answerboxentrymodel.AnswerBoxEntry, error) {
	db := flight.Context(nil, nil).DB

	entries := []answerboxentrymodel.AnswerBoxEntry{}

	err := db.Raw(`SELECT e.entry_id, e.question_id, e.entry, e.is_hidden, e.created_by, e.created_at
				FROM answer_box_entries e JOIN answer_box_questions q ON e.question_id = q.question_id JOIN base_box b ON b.id = q.box_id
				WHERE b.code = ? AND e.is_hidden = 0
				Order by created_at asc`, code).Scan(&entries).Error

	return entries, err
}

// ActivateQuestion activates the given question for an answer box
func ActivateQuestion(answerBoxQuestion answerboxquestionmodel.AnswerBoxQuestion, updatedBy string) error {
	db := flight.Context(nil, nil).DB

	err := db.Exec("UPDATE answer_box_questions SET is_active = 1, updated_at = ?, updated_by = ? WHERE question_id = ?", time.Now(), updatedBy, answerBoxQuestion.QuestionId).Error

	return err
}

// DeactivateQuestion activates the given question for an answer box
func DeactivateQuestion(answerBoxQuestion answerboxquestionmodel.AnswerBoxQuestion, updatedBy string) error {
	db := flight.Context(nil, nil).DB

	err := db.Exec("UPDATE answer_box_questions SET is_active = 0, updated_at = ?, updated_by = ? WHERE question_id = ?", time.Now(), updatedBy, answerBoxQuestion.QuestionId).Error

	return err
}

// HideEntry hides the given answer for a given box
func HideEntry(answerBoxEntry answerboxentrymodel.AnswerBoxEntry, updatedBy string) error {
	db := flight.Context(nil, nil).DB

	err := db.Exec("UPDATE answer_box_entries SET is_hidden = 1, updated_at = ?, updated_by = ? WHERE question_id = ? AND entry_id = ?", time.Now(), updatedBy, answerBoxEntry.QuestionId, answerBoxEntry.EntryId).Error

	return err
}

// UnhideEntry hides the given answer for a given box
func UnhideEntry(answerBoxEntry answerboxentrymodel.AnswerBoxEntry, updatedBy string) error {
	db := flight.Context(nil, nil).DB

	err := db.Exec("UPDATE answer_box_entries SET is_hidden = 0, updated_at = ?, updated_by = ? WHERE question_id = ? AND entry_id = ?", time.Now(), updatedBy, answerBoxEntry.QuestionId, answerBoxEntry.EntryId).Error

	return err
}

// GetActiveQuestions gets active questions for an answer box. For use with sms.
func GetActiveQuestions(boxId uuid.UUID) ([]answerboxquestionmodel.AnswerBoxQuestion, error) {
	db := flight.Context(nil, nil).DB

	answerBoxQuestion := []answerboxquestionmodel.AnswerBoxQuestion{}
	var err error

	err = db.Raw(`SELECT question_id, box_id, question, is_active, sort_order, created_at, created_by, updated_at, updated_by, deleted_at 
				FROM justasking.answer_box_questions 
				WHERE is_active = 1 AND box_id = ?`, boxId).Scan(&answerBoxQuestion).Error

	return answerBoxQuestion, err
}

// HideEntriesForQuestion hides all entries for a given question
func HideEntriesForQuestion(answerBoxQuestion answerboxquestionmodel.AnswerBoxQuestion, updatedBy string) error {
	db := flight.Context(nil, nil).DB

	err := db.Exec("UPDATE answer_box_entries SET is_hidden = 1, updated_at = ?, updated_by = ? WHERE question_id = ?", time.Now(), updatedBy, answerBoxQuestion.QuestionId).Error

	return err
}

// UnhideEntriesForQuestion unhides all entries for a given question
func UnhideEntriesForQuestion(answerBoxQuestion answerboxquestionmodel.AnswerBoxQuestion, updatedBy string) error {
	db := flight.Context(nil, nil).DB

	err := db.Exec("UPDATE answer_box_entries SET is_hidden = 0, updated_at = ?, updated_by = ? WHERE question_id = ?", time.Now(), updatedBy, answerBoxQuestion.QuestionId).Error

	return err
}

// HideAllEntries hides all answers for a given box
func HideAllEntries(answerBox answerboxmodel.AnswerBox, updatedBy string) error {
	db := flight.Context(nil, nil).DB

	err := db.Exec("UPDATE answer_box_entries e JOIN justasking.answer_box_questions q ON e.question_id = q.question_id SET is_hidden = 1, e.updated_at = ?, e.updated_by = ?  WHERE q.box_id = ?;", time.Now(), updatedBy, answerBox.BoxId).Error

	return err
}

// UnhideAllEntries hides the given answer for a given box
func UnhideAllEntries(answerBox answerboxmodel.AnswerBox, updatedBy string) error {
	db := flight.Context(nil, nil).DB

	err := db.Exec("UPDATE answer_box_entries e JOIN justasking.answer_box_questions q ON e.question_id = q.question_id SET is_hidden = 0, e.updated_at = ?, e.updated_by = ?  WHERE q.box_id = ?;", time.Now(), updatedBy, answerBox.BoxId).Error

	return err
}

// GetNumberOfAnswerBoxEntries gets the number of answer box entries for a particular box
func GetNumberOfAnswerBoxEntriesByQuestionId(questionId uuid.UUID) (int, error) {
	db := flight.Context(nil, nil).DB

	box := []answerboxentrymodel.AnswerBoxEntry{}

	err := db.Raw(`SELECT abe.entry_id, abe.question_id, abe.entry, abe.is_hidden
		FROM answer_box_entries abe 
        JOIN answer_box_questions abq ON abe.question_id = abq.question_id
        JOIN base_box bb ON bb.id = abq.box_id
		WHERE abq.box_id = (SELECT box_id FROM answer_box_questions WHERE question_id = ? LIMIT 1);`, questionId).Scan(&box).Error
	count := len(box)

	return count, err
}
