package service

import (
	"context"
	"errors"
	"fmt"
	"reflect"
	"testing"

	"github.com/Hidayathamir/opendiscuss/api/v1/question/dto"
	"github.com/Hidayathamir/opendiscuss/api/v1/question/repository"
	"github.com/Hidayathamir/opendiscuss/mocks"
	"github.com/Hidayathamir/opendiscuss/utils"
)

func TestQuestionService_GetQuestionList(t *testing.T) {
	type fields struct {
		repo      repository.IQuestionRepository
		trManager utils.ITransactionManager
	}
	type args struct {
		ctx context.Context
	}
	type testStruct struct {
		name    string
		fields  fields
		args    args
		want    []dto.QuestionHighlight
		wantErr bool
	}
	tests := []testStruct{}
	{
		Error_GetQuestionList := func() testStruct {
			repo := mocks.NewIQuestionRepository(t)

			ctx := context.Background()

			repo.
				On("GetQuestionList", ctx).
				Return(nil, errors.New("dummy error"))

			return testStruct{
				name:    "if get question list error, should return error",
				fields:  fields{repo: repo},
				args:    args{ctx: ctx},
				want:    nil,
				wantErr: true,
			}
		}
		fmt.Printf("%T\n", Error_GetQuestionList)

		Success := func() testStruct {
			repo := mocks.NewIQuestionRepository(t)

			ctx := context.Background()

			want := []dto.QuestionHighlight{}
			repo.
				On("GetQuestionList", ctx).
				Return(want, nil)

			return testStruct{
				name:    "if get question list success, should return success",
				fields:  fields{repo: repo},
				args:    args{ctx: ctx},
				want:    want,
				wantErr: false,
			}
		}
		fmt.Printf("%T\n", Success)

		tests = append(tests, Error_GetQuestionList())
		tests = append(tests, Success())
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			qs := &QuestionService{
				repo:      tt.fields.repo,
				trManager: tt.fields.trManager,
			}
			got, err := qs.GetQuestionList(tt.args.ctx)
			if (err != nil) != tt.wantErr {
				t.Errorf("QuestionService.GetQuestionList() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("QuestionService.GetQuestionList() = %v, want %v", got, tt.want)
			}
		})
	}
}
