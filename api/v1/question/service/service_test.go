package service

import (
	"reflect"
	"testing"

	"github.com/Hidayathamir/opendiscuss/api/v1/question/repository"
	"github.com/Hidayathamir/opendiscuss/utils"
)

func TestNewQuestionService(t *testing.T) {
	type args struct {
		repo      repository.IQuestionRepository
		trManager utils.ITransactionManager
	}
	type testStruct struct {
		name string
		args args
		want IQuestionService
	}

	tests := []testStruct{
		{
			name: "NewQuestionService should return IQuestionService implementation",
			args: args{},
			want: &QuestionService{},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NewQuestionService(tt.args.repo, tt.args.trManager); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewQuestionService() = %v, want %v", got, tt.want)
			}
		})
	}
}
