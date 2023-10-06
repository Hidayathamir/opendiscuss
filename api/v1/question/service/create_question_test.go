package service

import (
	"context"
	"errors"
	"fmt"
	"testing"

	"github.com/Hidayathamir/opendiscuss/api/v1/question/dto"
	"github.com/Hidayathamir/opendiscuss/api/v1/question/repository"
	"github.com/Hidayathamir/opendiscuss/constant"
	"github.com/Hidayathamir/opendiscuss/mocks"
	"github.com/Hidayathamir/opendiscuss/model"
	"github.com/Hidayathamir/opendiscuss/utils"
	"github.com/stretchr/testify/mock"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

func TestQuestionService_CreateQuestion(t *testing.T) {
	type fields struct {
		repo      repository.IQuestionRepository
		trManager utils.ITransactionManager
	}
	type args struct {
		ctx context.Context
		req dto.ReqCreateQuestion
	}

	type testStruct struct {
		name    string
		fields  fields
		args    args
		want    int
		wantErr bool
	}
	tests := []testStruct{}
	{
		Error_Validate := func() testStruct {
			return testStruct{
				name:    "if validate req error, should return error",
				fields:  fields{},
				args:    args{req: dto.ReqCreateQuestion{}},
				want:    0,
				wantErr: true,
			}
		}
		fmt.Printf("%T\n", Error_Validate)

		Error_CreateQuestion_User_Not_Found := func() testStruct {
			db, err := gorm.Open(sqlite.Open(constant.SQLITE_IN_MEMORY), &gorm.Config{DryRun: true})
			if err != nil {
				t.Fatalf(err.Error())
			}

			repo := mocks.NewIQuestionRepository(t)
			trManager := utils.NewTransactionManager(db)

			ctx := context.Background()
			req := dto.ReqCreateQuestion{UserID: 1, Question: "dummy question"}

			repo.
				On("CreateQuestion", mock.Anything, req.ToModelQuestion()).
				Return(0, errors.New("dummy error REFERENCES `users` (`id`)"))

			return testStruct{
				name:    "if create question error user not found, should fail transaction and return error",
				fields:  fields{repo: repo, trManager: trManager},
				args:    args{req: req, ctx: ctx},
				want:    0,
				wantErr: true,
			}
		}
		fmt.Printf("%T\n", Error_CreateQuestion_User_Not_Found)

		Error_CreateQuestion := func() testStruct {
			db, err := gorm.Open(sqlite.Open(constant.SQLITE_IN_MEMORY), &gorm.Config{DryRun: true})
			if err != nil {
				t.Fatalf(err.Error())
			}

			repo := mocks.NewIQuestionRepository(t)
			trManager := utils.NewTransactionManager(db)

			ctx := context.Background()
			req := dto.ReqCreateQuestion{UserID: 1, Question: "dummy question"}

			repo.
				On("CreateQuestion", mock.Anything, req.ToModelQuestion()).
				Return(0, errors.New("dummy error"))

			return testStruct{
				name:    "if create question error, should fail transaction and return error",
				fields:  fields{repo: repo, trManager: trManager},
				args:    args{req: req, ctx: ctx},
				want:    0,
				wantErr: true,
			}
		}
		fmt.Printf("%T\n", Error_CreateQuestion)

		Error_CreateQuestionStatistic := func() testStruct {
			db, err := gorm.Open(sqlite.Open(constant.SQLITE_IN_MEMORY), &gorm.Config{DryRun: true})
			if err != nil {
				t.Fatalf(err.Error())
			}

			repo := mocks.NewIQuestionRepository(t)
			trManager := utils.NewTransactionManager(db)

			ctx := context.Background()
			req := dto.ReqCreateQuestion{UserID: 1, Question: "dummy question"}

			questionID := 1233
			repo.
				On("CreateQuestion", mock.Anything, req.ToModelQuestion()).
				Return(questionID, nil)

			questionStatistic := model.QuestionStatistic{
				QuestionID: questionID,
			}
			repo.
				On("CreateQuestionStatistic", mock.Anything, questionStatistic).
				Return(0, errors.New("dummy error"))

			return testStruct{
				name:    "if create question statistic error, should fail transaction and return error",
				fields:  fields{repo: repo, trManager: trManager},
				args:    args{req: req, ctx: ctx},
				want:    0,
				wantErr: true,
			}
		}
		fmt.Printf("%T\n", Error_CreateQuestionStatistic)

		Success := func() testStruct {
			db, err := gorm.Open(sqlite.Open(constant.SQLITE_IN_MEMORY), &gorm.Config{DryRun: true})
			if err != nil {
				t.Fatalf(err.Error())
			}

			repo := mocks.NewIQuestionRepository(t)
			trManager := utils.NewTransactionManager(db)

			ctx := context.Background()
			req := dto.ReqCreateQuestion{UserID: 1, Question: "dummy question"}

			questionIDWanted := 123
			repo.
				On("CreateQuestion", mock.Anything, req.ToModelQuestion()).
				Return(questionIDWanted, nil)

			questionStatistic := model.QuestionStatistic{
				QuestionID: questionIDWanted,
			}
			repo.
				On("CreateQuestionStatistic", mock.Anything, questionStatistic).
				Return(999, nil)

			return testStruct{
				name:    "if transaction create question and create question statistic success, should success transaction and return success",
				fields:  fields{repo: repo, trManager: trManager},
				args:    args{req: req, ctx: ctx},
				want:    questionIDWanted,
				wantErr: false,
			}
		}
		fmt.Printf("%T\n", Success)

		tests = append(tests, Error_Validate())
		tests = append(tests, Error_CreateQuestion_User_Not_Found())
		tests = append(tests, Error_CreateQuestion())
		tests = append(tests, Error_CreateQuestionStatistic())
		tests = append(tests, Success())
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			qs := &QuestionService{
				repo:      tt.fields.repo,
				trManager: tt.fields.trManager,
			}
			got, err := qs.CreateQuestion(tt.args.ctx, tt.args.req)
			if (err != nil) != tt.wantErr {
				t.Errorf("QuestionService.CreateQuestion() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("QuestionService.CreateQuestion() = %v, want %v", got, tt.want)
			}
		})
	}
}
