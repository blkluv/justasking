import { Component, ChangeDetectorRef, OnInit, ViewChild, AfterViewChecked, AfterViewInit } from '@angular/core';
import { ActivatedRoute } from '@angular/router';
import { MatDialog } from '@angular/material';
import { MatTabGroup, MatTabChangeEvent } from '@angular/material';
import { Router } from '@angular/router';
import { UUID } from 'angular2-uuid';

import { Observable } from 'rxjs/Observable';
import 'rxjs/add/operator/catch';
import 'rxjs/add/observable/throw';

import { QuestionModel } from '../../core/models/question.model';
import { VotesBoxQuestionModel } from '../../core/models/votes-box-question.model';
import { VotesBoxQuestionAnswerModel } from '../../core/models/votes-box-question-answer.model';
import { ColorModel } from '../../core/models/color.model';
import { NewBoxModel } from '../../core/models/new-box.model';
import { QuestionBoxModel } from '../../core/models/question-box.model';
import { WordCloudModel } from '../../core/models/word-cloud.model';
import { AnswerBoxModel } from '../../core/models/answer-box.model';
import { VotesBoxModel } from '../../core/models/votes-box.model';
import { UserModel } from '../../core/models/user.model';
import { $WebSocket } from 'angular2-websocket/angular2-websocket';

import { ColorService } from '../../core/services/color.service';
import { BoxService } from '../../core/services/box.service';
import { RandomizerService } from '../../core/services/randomizer.service';
import { AnswerBoxService } from '../../core/services/answer-box.service';
import { WordCloudService } from '../../core/services/word-cloud.service';
import { QuestionBoxService } from '../../core/services/question-box.service';
import { VotesBoxService } from '../../core/services/votes-box.service';
import { ProcessingService } from '../../core/services/processing.service';
import { NotificationsService } from '../../core/services/notifications.service';
import { UserService } from '../../core/services/user.service';
import { AccountService } from '../../core/services/account.service';
import { WebSocketService } from '../../core/services/web-socket.service';

import { ConfirmBoxCreationDialogComponent } from './confirm-box-creation-dialog/confirm-box-creation-dialog.component';
import { NewUserWelcomeDialogComponent } from './new-user-welcome-dialog/new-user-welcome-dialog.component';

@Component({
  selector: 'app-new-box',
  templateUrl: './new-box.component.html',
  styleUrls: ['./new-box.component.scss']
})
export class NewBoxComponent implements OnInit {

  @ViewChild('dynamicVotesQuestionsTabGroup') dynamicVotesQuestionsTabGroup: MatTabGroup;

  newBoxForm: any;
  newBox: any;
  questionBoxPreview: QuestionBoxModel;
  votesBoxPreview: VotesBoxModel;
  wordCloudPreview: WordCloudModel;
  answerBoxPreview: AnswerBoxModel;
  boxTypes: { Name: string, Model: any }[];
  themes: ColorModel[];
  renderWordCloudPresentation: boolean;
  codeIsValid: boolean;
  validatingCode: boolean;
  codeTypingTimeout: any;
  votesBoxQuestionsTabGroup: any;
  customCodeEnabled: boolean;
  accountWebSocket: $WebSocket;

  constructor(
    private dialog: MatDialog,
    private route: ActivatedRoute,
    private changeDetectorRef: ChangeDetectorRef,
    private notificationsService: NotificationsService,
    private router: Router,
    private boxService: BoxService,
    private questionBoxService: QuestionBoxService,
    private colorService: ColorService,
    private answerBoxService: AnswerBoxService,
    private wordCloudService: WordCloudService,
    private votesBoxService: VotesBoxService,
    private randomizerService: RandomizerService,
    public processingService: ProcessingService,
    private userService: UserService,
    private accountService: AccountService,
    private webSocketService: WebSocketService
  ) {
    this.customCodeEnabled = false;
    this.codeIsValid = true;
    this.newBox = {
      BaseBox: {
        BoxType: 'question'
      }
    }
  }

  private findThemeId(theme: string) {
    for (var i = 0; i < this.themes.length; i++) {
      if (this.themes[i].Value == theme) {
        return this.themes[i].ID
      }
    }
  }

  confirmBoxCreation(newBox: NewBoxModel) {
    newBox.BaseBox.ThemeId = this.findThemeId(newBox.BaseBox.Theme);
    let dialogRef = this.dialog.open(ConfirmBoxCreationDialogComponent, {
      data: { newBox: newBox, accountWebSocket: this.accountWebSocket }
    });
  }

  invalidateCode($event: any) {
    var inp = String.fromCharCode($event.keyCode);
    //input was a letter, number, hyphen, underscore or space");
    if (/[a-zA-Z0-9]/.test(inp)) {
      this.validatingCode = true;
      this.codeIsValid = false;
    }
  }

  validateCode(code: string) {
    if (code) {
      this.validatingCode = true;
      this.codeIsValid = false;
      if (this.codeTypingTimeout) {
        clearTimeout(this.codeTypingTimeout);
      }
      this.codeTypingTimeout = setTimeout(() => {
        code = code && code.trim();
        this.validatingCode = true;
        this.boxService.codeIsTaken(code)
          .catch(error => {
            this.validatingCode = false;
            this.codeIsValid = false;
            return Observable.throw(error);
          })
          .subscribe(code => {
            this.codeIsValid = !code.Exists;
            this.validatingCode = false;
          })
      }, 400);
    }
  }

  disableWordCloudRendering() {
    this.renderWordCloudPresentation = false;
  }

  toggleRenderWordCloudPresentation(tabgroup: MatTabGroup) {
    this.renderWordCloudPresentation = tabgroup.selectedIndex == 1;
  }

  addQuestionToAnswerBox() {
    let newQuestion: QuestionModel = {
      QuestionId: UUID.UUID(),
      SortOrder: this.newBox.Questions.length + 1,
      Question: 'Ask a question here',
      Entries: this.randomizerService.getRandomAnswerBoxEntries(25),
      Entry: '',
      IsActive: true
    };
    this.newBox.Questions.push(newQuestion);
  }

  moveQuestionUpwards(question: QuestionModel) {
    if (question.SortOrder - 2 > -1) {
      var tempQuestion = this.newBox.Questions[question.SortOrder - 1];
      this.newBox.Questions[question.SortOrder - 1] = this.newBox.Questions[question.SortOrder - 2];
      this.newBox.Questions[question.SortOrder - 2] = tempQuestion;
    }
    this.newBox.Questions.forEach((question, index) => {
      question.SortOrder = index + 1;
    });
  }

  moveQuestionDownwards(question: QuestionModel) {
    if (question.SortOrder <= this.newBox.Questions.length) {
      var tempQuestion = this.newBox.Questions[question.SortOrder];
      this.newBox.Questions[question.SortOrder] = this.newBox.Questions[question.SortOrder - 1];
      this.newBox.Questions[question.SortOrder - 1] = tempQuestion;
    }
    this.newBox.Questions.forEach((question, index) => {
      question.SortOrder = index + 1;
    });
  }

  removeQuestionFromAnswerBox(questionToRemove: QuestionModel) {
    this.newBox.Questions.forEach((question, index) => {
      if (question.QuestionId == questionToRemove.QuestionId) {
        this.newBox.Questions.splice(index, 1);
      }
    });
    this.newBox.Questions.forEach((question, index) => {
      question.SortOrder = index + 1;
    });
  }

  addVotesBoxQuestion() {
    let votesBoxQuestion: VotesBoxQuestionModel = {
      QuestionId: UUID.UUID(),
      Header: "Ask a question here",
      SortOrder: this.newBox.Questions.length + 1,
      IsActive: true,
      Answers: [{
        AnswerId: UUID.UUID(),
        Answer: "Answer option",
        SortOrder: 1,
        Votes: this.randomizerService.getRandomVotesBoxEntries(50)
      }, {
        AnswerId: UUID.UUID(),
        Answer: "Answer option",
        SortOrder: 2,
        Votes: this.randomizerService.getRandomVotesBoxEntries(50)
      }]
    }
    this.newBox.Questions.push(votesBoxQuestion);
    this.newBox.Questions.forEach((question, index) => {
      question.SortOrder = index + 1;
    });
    this.dynamicVotesQuestionsTabGroup.selectedIndex = this.newBox.Questions.length - 1;
    this.changeDetectorRef.detectChanges();
  }

  moveVoteBoxQuestionUpwards(question: VotesBoxQuestionModel) {
    let questions = this.newBox.Questions;
    if (question.SortOrder - 2 > -1) {
      var tempQuestion = questions[question.SortOrder - 1];
      questions[question.SortOrder - 1] = questions[question.SortOrder - 2];
      questions[question.SortOrder - 2] = tempQuestion;
    }
    questions.forEach((question, index) => {
      question.SortOrder = index + 1;
    });
  }

  moveVoteBoxQuestionDownwards(question: VotesBoxQuestionModel) {
    let questions = this.newBox.Questions;
    if (question.SortOrder <= questions.length) {
      var tempQuestion = questions[question.SortOrder];
      questions[question.SortOrder] = questions[question.SortOrder - 1];
      questions[question.SortOrder - 1] = tempQuestion;
    }
    questions.forEach((question, index) => {
      question.SortOrder = index + 1;
    });
  }

  removeVoteBoxQuestion(questionToRemove: VotesBoxQuestionModel) {
    let questions = this.newBox.Questions;
    questions.forEach((question, index) => {
      if (question.SortOrder == questionToRemove.SortOrder) {
        questions.splice(index, 1);
      }
    });
    questions.forEach((question, index) => {
      question.SortOrder = index + 1;
    });
  }

  addVotesBoxQuestionAnswer(questionId: string) {
    this.newBox.Questions.forEach((question: VotesBoxQuestionModel) => {
      if (question.QuestionId == questionId) {
        let votesAnswer: VotesBoxQuestionAnswerModel = {
          QuestionId: question.QuestionId,
          AnswerId: UUID.UUID(),
          SortOrder: question.Answers.length + 1,
          Answer: "Answer option",
          Votes: this.randomizerService.getRandomVotesBoxEntries(50),
        }
        question.Answers.push(votesAnswer);
        question.Answers.forEach((answer, index) => {
          answer.SortOrder = index + 1;
        });
        question.Answers = question.Answers.slice();
      }
    });
  }

  moveVoteBoxQuestionAnswerUpwards(question: VotesBoxQuestionModel, answer: VotesBoxQuestionAnswerModel) {
    let answers = question.Answers;
    if (answer.SortOrder - 2 > -1) {
      var tempAnswer = answers[answer.SortOrder - 1];
      answers[answer.SortOrder - 1] = answers[answer.SortOrder - 2];
      answers[answer.SortOrder - 2] = tempAnswer;
    }
    answers.forEach((answer, index) => {
      answer.SortOrder = index + 1;
    });
  }

  moveVoteBoxQuestionAnswerDownwards(question: VotesBoxQuestionModel, answer: VotesBoxQuestionAnswerModel) {
    let answers = question.Answers;
    if (answer.SortOrder <= answers.length) {
      var tempAnswer = answers[answer.SortOrder];
      answers[answer.SortOrder] = answers[answer.SortOrder - 1];
      answers[answer.SortOrder - 1] = tempAnswer;
    }
    answers.forEach((answer, index) => {
      answer.SortOrder = index + 1;
    });
  }

  removeVoteBoxQuestionAnswer(question: VotesBoxQuestionModel, answerToRemove: VotesBoxQuestionAnswerModel) {
    let answers = question.Answers;
    answers.forEach((answer, index) => {
      if (answer.SortOrder == answerToRemove.SortOrder) {
        answers.splice(index, 1);
      }
    });
    answers.forEach((answer, index) => {
      answer.SortOrder = index + 1;
    }); 
    question.Answers = question.Answers.slice();
  }

  ngAfterViewChecked() {
    this.changeDetectorRef.detectChanges();
  }

  ngAfterViewInit() {
    this.changeDetectorRef.detectChanges();
  }

  ngOnInit() {
    let newUserWelcome = this.route.snapshot.queryParams['welcome'] == 'true';

    if (newUserWelcome) {
      setTimeout(() => {
        let dialogRef = this.dialog.open(NewUserWelcomeDialogComponent);
        dialogRef.afterClosed().subscribe(result => {
        });
      })
    }

    this.colorService.getColors()
      .catch(error => {
        console.error("Could not retrieve colors for populating themes dropdown.", error);
        return Observable.throw(error);
      })
      .subscribe(response => {
        this.themes = response;
      });

    this.userService.getUser()
      .catch(error => {
        return Observable.throw(error);
      })
      .subscribe(() => {
        this.accountWebSocket = this.webSocketService.connectToAccountChannel(this.userService.user.Account.Id);
        this.customCodeEnabled = this.userService.user && this.userService.user.MembershipDetails && this.userService.user.MembershipDetails.CustomCode;
        if (!this.customCodeEnabled) {
          this.boxService.getRandomCode()
            .catch(error => {
              return Observable.throw(error);
            })
            .subscribe((randomCode: string) => {
              this.questionBoxPreview.BaseBox.Code = randomCode;
              this.votesBoxPreview.BaseBox.Code = randomCode;
              this.wordCloudPreview.BaseBox.Code = randomCode;
              this.answerBoxPreview.BaseBox.Code = randomCode;
            });
        }
      });

    this.wordCloudPreview = {
      Header: 'Name your favorite baseball team',
      DefaultWord: "Marlins",
      IsPreview: true,
      BaseBox: {
        IsLive: true,
        BoxTypeId: 1,
        BoxType: 'Word Cloud',
        Theme: 'core-deep-orange',
        PhoneNumber: "+1234567890"
      }
    };

    this.questionBoxPreview = {
      Header: `I'm taking questions now, ask me anything!`,
      Action: 'Ask',
      Question: '',
      IsPreview: true,
      BaseBox: {
        IsLive: true,
        BoxTypeId: 2,
        BoxType: 'Question Poll',
        Theme: 'core-green',
        PhoneNumber: "+1234567890"
      }
    };

    this.answerBoxPreview = {
      Questions: [
        {
          QuestionId: UUID.UUID(),
          SortOrder: 1,
          Question: 'Where should we go for lunch today?',
          Entries: this.randomizerService.getRandomAnswerBoxEntries(25),
          Entry: '',
          IsActive: true
        }
      ],
      IsPreview: true,
      BaseBox: {
        IsLive: true,
        BoxTypeId: 3,
        BoxType: 'Answer Poll',
        Theme: 'core-cyan',
        PhoneNumber: "+1234567890"
      }
    };

    this.votesBoxPreview = {
      Questions: [{
        QuestionId: UUID.UUID(),
        Header: "Which is your favorite type of pizza?",
        Answers: [{
          AnswerId: UUID.UUID(),
          Answer: "Cheese",
          SortOrder: 1,
        }, {
          AnswerId: UUID.UUID(),
          Answer: "Pepperoni",
          SortOrder: 2,
        }, {
          AnswerId: UUID.UUID(),
          Answer: "Hawaiian",
          SortOrder: 3,
        }],
        SortOrder: 1,
        IsActive: true,
      }],
      IsPreview: true,
      BaseBox: {
        IsLive: true,
        BoxTypeId: 4,
        BoxType: 'Votes Poll',
        Theme: "core-blue",
        PhoneNumber: "+1234567890"
      }
    };

    this.boxTypes = [
      {
        Name: "Votes Poll",
        Model: this.votesBoxPreview
      },
      {
        Name: "Word Cloud",
        Model: this.wordCloudPreview
      },
      {
        Name: "Question Poll",
        Model: this.questionBoxPreview
      },
      {
        Name: "Answer Poll",
        Model: this.answerBoxPreview
      }
    ];

    this.newBox = this.votesBoxPreview;
  }
}
