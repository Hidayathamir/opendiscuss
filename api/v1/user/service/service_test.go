package service

import (
	"reflect"
	"testing"

	"github.com/Hidayathamir/opendiscuss/api/v1/user/repository"
)

func TestNewUserService(t *testing.T) {
	type args struct {
		repo repository.IUserRepository
	}
	tests := []struct {
		name string
		args args
		want IUserService
	}{
		{
			name: "NewUserService should return IUserService implementation",
			args: args{},
			want: &UserService{},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NewUserService(tt.args.repo); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewUserService() = %v, want %v", got, tt.want)
			}
		})
	}
}
