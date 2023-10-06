package service

import (
	"context"
	"errors"
	"fmt"
	"testing"

	"github.com/Hidayathamir/opendiscuss/api/v1/user/dto"
	"github.com/Hidayathamir/opendiscuss/api/v1/user/repository"
	"github.com/Hidayathamir/opendiscuss/mocks"
	"github.com/stretchr/testify/mock"
)

func TestUserService_RegisterUser(t *testing.T) {
	type fields struct {
		repo repository.IUserRepository
	}
	type args struct {
		ctx context.Context
		req dto.ReqRegisterUser
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
				args:    args{req: dto.ReqRegisterUser{}},
				want:    0,
				wantErr: true,
			}
		}
		fmt.Printf("%T\n", Error_Validate)

		Error_GenerateHashPassword := func() testStruct {
			// long password will error generate password
			// as documentation said.
			//
			// "GenerateFromPassword does not accept passwords longer than 72 bytes, which
			// is the longest password bcrypt will operate on."
			longPassword := "1234567890"
			longPassword += "1234567890"
			longPassword += "1234567890"
			longPassword += "1234567890"
			longPassword += "1234567890"
			longPassword += "1234567890"
			longPassword += "1234567890"
			longPassword += "1234567890"

			req := dto.ReqRegisterUser{
				Username: "dummy username", Password: longPassword,
			}

			return testStruct{
				name:    "if generate hash password error, should return error",
				fields:  fields{},
				args:    args{req: req},
				want:    0,
				wantErr: true,
			}
		}
		fmt.Printf("%T\n", Error_GenerateHashPassword)

		Error_RegisterUser := func() testStruct {
			repo := mocks.NewIUserRepository(t)

			req := dto.ReqRegisterUser{
				Username: "dummy username", Password: "dummy password",
			}

			repo.
				On("RegisterUser", mock.Anything, mock.Anything).
				Return(0, errors.New("dummy error"))

			return testStruct{
				name:    "if register user error, should return error",
				fields:  fields{repo: repo},
				args:    args{req: req},
				want:    0,
				wantErr: true,
			}
		}
		fmt.Printf("%T\n", Error_RegisterUser)

		Success := func() testStruct {
			repo := mocks.NewIUserRepository(t)

			req := dto.ReqRegisterUser{
				Username: "dummy username", Password: "dummy password",
			}

			wantUserID := 232
			repo.
				On("RegisterUser", mock.Anything, mock.Anything).
				Return(wantUserID, nil)

			return testStruct{
				name:    "if register user success, should return success",
				fields:  fields{repo: repo},
				args:    args{req: req},
				want:    wantUserID,
				wantErr: false,
			}
		}
		fmt.Printf("%T\n", Success)

		tests = append(tests, Error_Validate())
		tests = append(tests, Error_GenerateHashPassword())
		tests = append(tests, Error_RegisterUser())
		tests = append(tests, Success())
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			us := &UserService{
				repo: tt.fields.repo,
			}
			got, err := us.RegisterUser(tt.args.ctx, tt.args.req)
			if (err != nil) != tt.wantErr {
				t.Errorf("UserService.RegisterUser() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("UserService.RegisterUser() = %v, want %v", got, tt.want)
			}
		})
	}
}
