package votesboxrepo

import (
	"justasking/GO/core/model/boxes/basebox"
	"justasking/GO/core/model/boxes/votesbox"
	"justasking/GO/core/model/votesboxquestion"
	"justasking/GO/core/model/votesboxquestionanswer"
	"justasking/GO/core/model/votesboxquestionanswervote"
	"justasking/GO/core/repo/boxes/basebox"
	"justasking/GO/core/startup/flight"
	"time"

	uuid "github.com/satori/go.uuid"
)

// InsertVotesBox inserts a new votes box
func InsertVotesBox(votesBox votesboxmodel.VotesBox) error {
	db := flight.Context(nil, nil).DB

	// Wrapping box creation in a transaction. We wouldn't want the VotesBox to be created without the BaseBox
	tx := db.Begin()

	votesBox.BaseBox.ID = votesBox.BoxId
	baseBox := votesBox.BaseBox

	if err := baseboxrepo.InsertBaseBox(baseBox, tx); err != nil {
		tx.Rollback()
		return err
	}

	if err := tx.Create(&votesBox).Error; err != nil {
		tx.Rollback()
		return err
	}

	//looping over every question and answer set for this votes box
	var err error
	for _, question := range votesBox.Questions {
		question.BoxId = votesBox.BaseBox.ID
		question.QuestionId, _ = uuid.NewV4()
		err = tx.Create(&question).Error
		if err != nil {
			break
		}
		for _, answer := range question.Answers {
			answer.AnswerId, _ = uuid.NewV4()
			answer.QuestionId = question.QuestionId
			err = tx.Create(&answer).Error
			if err != nil {
				break
			}
		}
	}

	if err != nil {
		tx.Rollback()
		return err
	}

	tx.Commit()
	return nil
}

// GetVotesBoxByBoxId gets a votes box by id
func GetVotesBoxByBoxId(guid uuid.UUID) (votesboxmodel.VotesBox, error) {
	db := flight.Context(nil, nil).DB

	baseBox := baseboxmodel.BaseBox{}
	votesBox := votesboxmodel.VotesBox{}
	votesBoxQuestions := []votesboxquestionmodel.VotesBoxQuestion{}
	votesBoxQuestionAnswers := []votesboxquestionanswermodel.VotesBoxQuestionAnswer{}

	//get basebox
	baseBox, err := baseboxrepo.GetBaseBoxById(guid)
	if err != nil {
		return votesBox, err
	}

	//get votesbox
	err = db.Where("box_id = ?", guid).Find(&votesBox).Error
	if err != nil {
		return votesBox, err
	}

	//get votesboxquestions
	err = db.Where("box_id = ?", guid).Order("sort_order").Find(&votesBoxQuestions).Error
	if err != nil {
		return votesBox, err
	}

	//get votesboxquestionanswers
	for _, question := range votesBoxQuestions {
		err = db.Raw(`SELECT a.answer_id, a.question_id, answer, sort_order, COUNT(v.answer_id) as votes, IFNULL(v.created_at, current_timestamp) as created_at, IFNULL(v.created_by, 0) as created_by, v.updated_at, v.updated_by, v.deleted_at
		FROM votes_box_question_answers a LEFT JOIN votes_box_question_answers_votes v ON a.answer_id = v.answer_id
		WHERE a.question_id = ?
		GROUP BY a.answer_id
		ORDER BY a.sort_order ASC;`, question.QuestionId).Scan(&votesBoxQuestionAnswers).Error
		if err != nil {
			return votesBox, err
		}
		question.Answers = votesBoxQuestionAnswers
		votesBox.Questions = append(votesBox.Questions, question)
	}

	//assign votesBox's basebox and questions
	votesBox.BaseBox = baseBox

	return votesBox, err
}

// GetVotesBoxByBoxCode returns a votes box
func GetVotesBoxByBoxCode(code string) (votesboxmodel.VotesBox, error) {
	db := flight.Context(nil, nil).DB

	votesBox := votesboxmodel.VotesBox{}
	votesBoxQuestions := []votesboxquestionmodel.VotesBoxQuestion{}
	votesBoxQuestionAnswers := []votesboxquestionanswermodel.VotesBoxQuestionAnswer{}

	//get basebox
	baseBox, err := baseboxrepo.GetBaseBoxByCode(code)
	if err != nil {
		return votesBox, err
	}

	//get votesbox
	err = db.Where("box_id = ?", baseBox.ID).Find(&votesBox).Error
	if err != nil {
		return votesBox, err
	}

	//get votesboxquestions
	err = db.Where("box_id = ?", baseBox.ID).Order("sort_order").Find(&votesBoxQuestions).Error
	if err != nil {
		return votesBox, err
	}

	//get votesboxquestionanswers
	for _, question := range votesBoxQuestions {
		err = db.Raw(`SELECT a.answer_id, a.question_id, answer, sort_order, COUNT(v.answer_id) as votes, IFNULL(v.created_at, current_timestamp) as created_at, IFNULL(v.created_by, 0) as created_by, v.updated_at, v.updated_by, v.deleted_at
		FROM votes_box_question_answers a LEFT JOIN votes_box_question_answers_votes v ON a.answer_id = v.answer_id
		WHERE a.question_id = ?
		GROUP BY a.answer_id
		ORDER BY a.sort_order ASC;`, question.QuestionId).Scan(&votesBoxQuestionAnswers).Error
		if err != nil {
			return votesBox, err
		}
		question.Answers = votesBoxQuestionAnswers
		votesBox.Questions = append(votesBox.Questions, question)
	}
	//assign votesBox's basebox and questions
	votesBox.BaseBox = baseBox

	return votesBox, err
}

// InsertVoteNamed inserts a vote for a given answer
func InsertVoteNamed(votesBoxQuestionAnswer votesboxquestionanswermodel.VotesBoxQuestionAnswer) (votesboxquestionanswervotemodel.VotesBoxQuestionAnswerVote, error) {
	db := flight.Context(nil, nil).DB

	inserted := votesboxquestionanswervotemodel.VotesBoxQuestionAnswerVote{}

	err := db.Raw(`CALL InsertVotesBoxVoteNamed(?, ?, ?) `,
		votesBoxQuestionAnswer.AnswerId, votesBoxQuestionAnswer.QuestionId, votesBoxQuestionAnswer.CreatedBy).Scan(&inserted).Error

	return inserted, err
}

// InsertVoteAnonymous inserts a vote with a default createdBy of 0. for use when voting from frontend.
func InsertVoteAnonymous(votesBoxQuestionAnswer votesboxquestionanswermodel.VotesBoxQuestionAnswer) error {
	db := flight.Context(nil, nil).DB

	err := db.Exec(`INSERT INTO votes_box_question_answers_votes (answer_id, question_id, created_by) 
					VALUES (?, ?, ?)`,
		votesBoxQuestionAnswer.AnswerId, votesBoxQuestionAnswer.QuestionId, "0").Error

	return err
}

// GetActiveQuestions gets active questions for a votes box. For use with sms.
func GetActiveQuestions(boxId uuid.UUID) ([]votesboxquestionmodel.VotesBoxQuestion, error) {
	db := flight.Context(nil, nil).DB

	votesboxquestionsResult := []votesboxquestionmodel.VotesBoxQuestion{}
	votesboxQuestions := []votesboxquestionmodel.VotesBoxQuestion{}
	votesBoxQuestionAnswers := []votesboxquestionanswermodel.VotesBoxQuestionAnswer{}
	var err error

	err = db.Raw(`SELECT question_id, box_id, header, is_active, sort_order, created_at, created_by, updated_at, updated_by, deleted_at 
				FROM justasking.votes_box_questions WHERE is_active = 1 AND box_id = ?`, boxId).Scan(&votesboxQuestions).Error

	//get votesboxquestionanswers
	for _, question := range votesboxQuestions {
		err = db.Raw(`SELECT answer_id, question_id, answer, sort_order, created_at, created_by, updated_at, updated_by, deleted_at
		FROM votes_box_question_answers a
		WHERE a.question_id = ? 
		ORDER BY a.sort_order ASC;`, question.QuestionId).Scan(&votesBoxQuestionAnswers).Error
		if err != nil {
			return votesboxQuestions, err
		}
		question.Answers = votesBoxQuestionAnswers
		votesboxquestionsResult = append(votesboxquestionsResult, question)
	}

	return votesboxquestionsResult, err
}

// ActivateQuestion activates the given question for a votes box
func ActivateQuestion(votesboxQuestion votesboxquestionmodel.VotesBoxQuestion, updatedBy string) error {
	db := flight.Context(nil, nil).DB

	err := db.Exec("UPDATE votes_box_questions SET is_active = 1, updated_at = ?, updated_by = ? WHERE question_id = ?", time.Now(), updatedBy, votesboxQuestion.QuestionId).Error

	return err
}

// DeactivateQuestion activates the given question for a votes box
func DeactivateQuestion(votesboxQuestion votesboxquestionmodel.VotesBoxQuestion, updatedBy string) error {
	db := flight.Context(nil, nil).DB

	err := db.Exec("UPDATE votes_box_questions SET is_active = 0, updated_at = ?, updated_by = ? WHERE question_id = ?", time.Now(), updatedBy, votesboxQuestion.QuestionId).Error

	return err
}

// GetNumberOfVotesBoxEntries gets the number of votes box entries for a particular box
func GetNumberOfVotesBoxEntriesByQuestionId(questionId uuid.UUID) (int, error) {
	db := flight.Context(nil, nil).DB

	box := []votesboxquestionanswervotemodel.VotesBoxQuestionAnswerVote{}

	err := db.Raw(`SELECT vbav.answer_id, vbav.question_id
		FROM votes_box_question_answers_votes vbav
		JOIN votes_box_questions vbq ON vbav.question_id = vbq.question_id
		JOIN base_box bb ON bb.id = vbq.box_id
		WHERE vbq.box_id = (SELECT box_id FROM votes_box_question_answers WHERE question_id = ? LIMIT 1);`, questionId).Scan(&box).Error
	count := len(box)

	return count, err
}
