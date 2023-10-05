package service

import (
	"context"
	"errors"
	"testing"

	"github.com/Hidayathamir/opendiscuss/api/v1/question/dto"
	"github.com/Hidayathamir/opendiscuss/api/v1/question/repository"
	"github.com/Hidayathamir/opendiscuss/mocks"
	"github.com/stretchr/testify/mock"
)

func TestQuestionService_CreateQuestion(t *testing.T) {
	type fields struct {
		repo repository.IQuestionRepository
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

		Error_CreateQuestion_User_Not_Found := func() testStruct {
			repo := mocks.NewIQuestionRepository(t)
			req := dto.ReqCreateQuestion{UserID: 1, Question: "dummy question"}

			repo.
				On("CreateQuestion", mock.Anything, req.ToModelQuestion()).
				Return(0, errors.New("dummy error REFERENCES `users` (`id`)"))

			return testStruct{
				name:    "if create question error user not found, should return error",
				fields:  fields{repo: repo},
				args:    args{req: req},
				want:    0,
				wantErr: true,
			}
		}

		Error_CreateQuestion := func() testStruct {
			repo := mocks.NewIQuestionRepository(t)
			req := dto.ReqCreateQuestion{UserID: 1, Question: "dummy question"}

			repo.
				On("CreateQuestion", mock.Anything, req.ToModelQuestion()).
				Return(0, errors.New("dummy error"))

			return testStruct{
				name:    "if create question error, should return error",
				fields:  fields{repo: repo},
				args:    args{req: req},
				want:    0,
				wantErr: true,
			}
		}

		Success := func() testStruct {
			repo := mocks.NewIQuestionRepository(t)
			req := dto.ReqCreateQuestion{UserID: 1, Question: "dummy question"}

			want := 123
			repo.
				On("CreateQuestion", mock.Anything, req.ToModelQuestion()).
				Return(want, nil)

			return testStruct{
				name:    "if create question success, should return success",
				fields:  fields{repo: repo},
				args:    args{req: req},
				want:    want,
				wantErr: false,
			}
		}

		tests = append(tests, Error_Validate())
		tests = append(tests, Error_CreateQuestion_User_Not_Found())
		tests = append(tests, Error_CreateQuestion())
		tests = append(tests, Success())
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			qs := &QuestionService{
				repo: tt.fields.repo,
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
