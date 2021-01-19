package questionboxrepo

import (
	"justasking/GO/core/model/boxes/basebox"
	"justasking/GO/core/model/boxes/questionbox"
	"justasking/GO/core/model/questionboxentry"
	"justasking/GO/core/model/questionboxentryvote"
	"justasking/GO/core/repo/boxes/basebox"
	"justasking/GO/core/startup/flight"
	"time"

	uuid "github.com/satori/go.uuid"
)

// InsertQuestionBox creates a base box and questionbox
func InsertQuestionBox(questionBox questionboxmodel.QuestionBox) error {
	db := flight.Context(nil, nil).DB

	// Wrapping box creation in a transaction. We wouldn't want the QuestionBox to be created without the BaseBox
	tx := db.Begin()

	questionBox.BaseBox.ID = questionBox.BoxId
	baseBox := questionBox.BaseBox

	if err := baseboxrepo.InsertBaseBox(baseBox, tx); err != nil {
		tx.Rollback()
		return err
	}

	if err := tx.Create(&questionBox).Error; err != nil {
		tx.Rollback()
		return err
	}

	tx.Commit()
	return nil
}

// GetQuestionBoxByBoxId returns questionbox
func GetQuestionBoxByBoxId(guid uuid.UUID) (questionboxmodel.QuestionBox, error) {
	db := flight.Context(nil, nil).DB

	baseBox := baseboxmodel.BaseBox{}
	questionBox := questionboxmodel.QuestionBox{}

	//get basebox
	baseBox, err := baseboxrepo.GetBaseBoxById(guid)
	if err != nil {
		return questionBox, err
	}
	//get questionbox
	err = db.Where("box_id = ?", guid).Find(&questionBox).Error
	if err != nil {
		return questionBox, err
	}

	//assign questionBox's basebox
	questionBox.BaseBox = baseBox

	return questionBox, err
}

// GetQuestionBoxByBoxCode returns questionbox
func GetQuestionBoxByBoxCode(code string) (questionboxmodel.QuestionBox, error) {
	db := flight.Context(nil, nil).DB

	//baseBox := baseboxmodel.BaseBox{}
	questionBox := questionboxmodel.QuestionBox{}

	//get basebox
	baseBox, err := baseboxrepo.GetBaseBoxByCode(code)
	if err != nil {
		return questionBox, err
	}
	//get questionbox
	err = db.Where("box_id = ?", baseBox.ID).Find(&questionBox).Error
	if err != nil {
		return questionBox, err
	}

	//assign questionbox's basebox
	questionBox.BaseBox = baseBox

	return questionBox, err
}

// InsertQuestionBoxEntry adds an entry to the question box entries table
func InsertQuestionBoxEntry(entry questionboxentrymodel.QuestionBoxEntry) error {
	db := flight.Context(nil, nil).DB

	err := db.Exec(`INSERT INTO question_box_entries (entry_id, box_id, question, is_hidden, is_favorite, created_by) 
		VALUES (?, ?, ?, ?, ?, ?)`,
		entry.EntryId, entry.BoxId, entry.Question, entry.IsHidden, entry.IsFavorite, entry.CreatedBy).Error

	return err
}

// InsertQuestionBoxEntryVote adds an entry vote to the question box entries table
func InsertQuestionBoxEntryVote(entryVote questionboxentryvotemodel.QuestionBoxEntryVote) error {
	db := flight.Context(nil, nil).DB

	err := db.Create(&entryVote).Error

	return err
}

// GetQuestionBoxEntriesByBoxId gets all entries for a specific question box
func GetQuestionBoxEntriesByBoxId(boxId uuid.UUID) ([]questionboxentrymodel.QuestionBoxEntry, error) {
	db := flight.Context(nil, nil).DB

	entries := []questionboxentrymodel.QuestionBoxEntry{}
	err := db.Raw(`SELECT e.entry_id, e.box_id, e.question, e.is_hidden, e.is_favorite, e.created_by, e.created_at,
			SUM(IF(ev.vote_type = 'upvote', IFNULL(ev.vote_value, 0), 0)) as upvotes, 
			SUM(IF(ev.vote_type = 'downvote', IFNULL(ev.vote_value, 0), 0)) as downvotes
			FROM question_box_entries e LEFT JOIN question_box_entries_votes ev ON e.entry_id = ev.entry_id
			WHERE e.box_id = ?
           	GROUP BY e.entry_id
			Order by e.created_at asc`, boxId).Scan(&entries).Error

	return entries, err
}

// GetQuestionBoxEntriesByCode gets all entries for a specific question box
func GetQuestionBoxEntriesByCode(code string) ([]questionboxentrymodel.QuestionBoxEntry, error) {
	db := flight.Context(nil, nil).DB

	entries := []questionboxentrymodel.QuestionBoxEntry{}
	err := db.Raw(`SELECT e.entry_id, e.box_id, e.question, e.is_hidden, e.is_favorite, e.created_by, e.created_at,
			SUM(IF(ev.vote_type = 'upvote', IFNULL(ev.vote_value, 0), 0)) as upvotes, 
			SUM(IF(ev.vote_type = 'downvote', IFNULL(ev.vote_value, 0), 0)) as downvotes
			FROM question_box_entries e LEFT JOIN question_box_entries_votes ev ON e.entry_id = ev.entry_id
			WHERE e.box_id = (SELECT bb.id FROM base_box bb WHERE code = ?)
		   	GROUP BY e.entry_id
			Order by e.created_at asc`, code).Scan(&entries).Error

	return entries, err
}

// GetVisibleQuestionBoxEntriesByCode gets all entries for a specific question box
func GetVisibleQuestionBoxEntriesByCode(code string) ([]questionboxentrymodel.QuestionBoxEntry, error) {
	db := flight.Context(nil, nil).DB

	entries := []questionboxentrymodel.QuestionBoxEntry{}
	err := db.Raw(`SELECT e.entry_id, e.box_id, e.question, e.is_hidden, e.is_favorite, e.created_by, e.created_at,
			SUM(IF(ev.vote_type = 'upvote', IFNULL(ev.vote_value, 0), 0)) as upvotes, 
			SUM(IF(ev.vote_type = 'downvote', IFNULL(ev.vote_value, 0), 0)) as downvotes
			FROM question_box_entries e LEFT JOIN question_box_entries_votes ev ON e.entry_id = ev.entry_id
			WHERE e.box_id = (SELECT bb.id FROM base_box bb WHERE code = ?)
			AND e.is_hidden = 0
		   	GROUP BY e.entry_id
			Order by e.created_at asc`, code).Scan(&entries).Error

	return entries, err
}

// HideEntry hides the given answer for a given box
func HideEntry(questionBoxEntry questionboxentrymodel.QuestionBoxEntry, updatedBy string) error {
	db := flight.Context(nil, nil).DB

	err := db.Exec("UPDATE question_box_entries SET is_hidden = 1, updated_at = ?, updated_by = ? WHERE box_id = ? AND entry_id = ?", time.Now(), updatedBy, questionBoxEntry.BoxId, questionBoxEntry.EntryId).Error

	return err
}

// UnhideEntry hides the given answer for a given box
func UnhideEntry(questionBoxEntry questionboxentrymodel.QuestionBoxEntry, updatedBy string) error {
	db := flight.Context(nil, nil).DB

	err := db.Exec("UPDATE question_box_entries SET is_hidden = 0, updated_at = ?, updated_by = ? WHERE box_id = ? AND entry_id = ?", time.Now(), updatedBy, questionBoxEntry.BoxId, questionBoxEntry.EntryId).Error

	return err
}

// UpvoteFromDownvote cancels out previous downvote by giving it a -1, and adds an upvote
func UpvoteFromDownvote(entryVote questionboxentryvotemodel.QuestionBoxEntryVote) error {
	db := flight.Context(nil, nil).DB

	// Wrapping vote in a transaction. We wouldn't want one of the votes to be created without the other.
	tx := db.Begin()

	if err := db.Exec("INSERT INTO question_box_entries_votes VALUES (?, 'upvote', 1, ?, ?, ?, ?, null)", entryVote.EntryId, time.Now(), entryVote.CreatedBy, time.Now(), entryVote.CreatedBy).Error; err != nil {
		tx.Rollback()
		return err
	}

	if err := db.Exec("INSERT INTO question_box_entries_votes VALUES (?, 'downvote', -1, ?, ?, ?, ?, null)", entryVote.EntryId, time.Now(), entryVote.CreatedBy, time.Now(), entryVote.CreatedBy).Error; err != nil {
		tx.Rollback()
		return err
	}

	tx.Commit()
	return nil
}

// DownvoteFromUpvote cancels out previous upvote by giving it a -1, and adds a downvote
func DownvoteFromUpvote(entryVote questionboxentryvotemodel.QuestionBoxEntryVote) error {
	db := flight.Context(nil, nil).DB

	// Wrapping vote in a transaction. We wouldn't want one of the votes to be created without the other.
	tx := db.Begin()

	if err := db.Exec("INSERT INTO question_box_entries_votes VALUES (?, 'upvote', -1, ?, ?, ?, ?, null)", entryVote.EntryId, time.Now(), entryVote.CreatedBy, time.Now(), entryVote.CreatedBy).Error; err != nil {
		tx.Rollback()
		return err
	}

	if err := db.Exec("INSERT INTO question_box_entries_votes VALUES (?, 'downvote', 1, ?, ?, ?, ?, null)", entryVote.EntryId, time.Now(), entryVote.CreatedBy, time.Now(), entryVote.CreatedBy).Error; err != nil {
		tx.Rollback()
		return err
	}

	tx.Commit()
	return nil
}

// UndoUpvote adds a -1 to the upvotes for a question
func UndoUpvote(entryVote questionboxentryvotemodel.QuestionBoxEntryVote) error {
	db := flight.Context(nil, nil).DB

	err := db.Exec("INSERT INTO question_box_entries_votes VALUES (?, 'upvote', -1, CURRENT_TIMESTAMP, ?, null, null, null)", entryVote.EntryId, entryVote.CreatedBy).Error

	return err
}

// UndoDownvote adds a -1 to the downvotes for a question
func UndoDownvote(entryVote questionboxentryvotemodel.QuestionBoxEntryVote) error {
	db := flight.Context(nil, nil).DB

	err := db.Exec("INSERT INTO question_box_entries_votes VALUES (?, 'downvote', -1, CURRENT_TIMESTAMP, ?, null, null, null)", entryVote.EntryId, entryVote.CreatedBy).Error

	return err
}

// HideAllEntries hides the given answer for a given box
func HideAllEntries(questionBox questionboxmodel.QuestionBox, updatedBy string) error {
	db := flight.Context(nil, nil).DB

	err := db.Exec("UPDATE question_box_entries SET is_hidden = 1, updated_at = ?, updated_by = ? WHERE box_id = ?", time.Now(), updatedBy, questionBox.BoxId).Error

	return err
}

// UnhideAllEntries hides the given answer for a given box
func UnhideAllEntries(questionBox questionboxmodel.QuestionBox, updatedBy string) error {
	db := flight.Context(nil, nil).DB

	err := db.Exec("UPDATE question_box_entries SET is_hidden = 0, updated_at = ?, updated_by = ? WHERE box_id = ?", time.Now(), updatedBy, questionBox.BoxId).Error

	return err
}
