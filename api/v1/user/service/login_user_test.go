package service

import (
	"context"
	"errors"
	"fmt"
	"testing"

	"github.com/Hidayathamir/opendiscuss/api/v1/user/dto"
	"github.com/Hidayathamir/opendiscuss/api/v1/user/repository"
	"github.com/Hidayathamir/opendiscuss/mocks"
	"github.com/Hidayathamir/opendiscuss/model"
	"github.com/Hidayathamir/opendiscuss/utils"
	"github.com/stretchr/testify/mock"
)

func TestUserService_LoginUser(t *testing.T) {
	type fields struct {
		repo repository.IUserRepository
	}
	type args struct {
		ctx context.Context
		req dto.ReqLoginUser
	}
	type testStruct struct {
		name    string
		fields  fields
		args    args
		want    string
		wantErr bool
	}
	tests := []testStruct{}
	{
		Error_Validate := func() testStruct {
			return testStruct{
				name:    "if validate req error, should return error",
				fields:  fields{},
				args:    args{},
				want:    "",
				wantErr: true,
			}
		}
		fmt.Printf("%T\n", Error_Validate)

		Error_GetUserByUsername := func() testStruct {
			repo := mocks.NewIUserRepository(t)

			req := dto.ReqLoginUser{
				Username: "dummy username",
				Password: "dummy password",
			}

			repo.
				On("GetUserByUsername", mock.Anything, req.Username).
				Return(model.User{}, errors.New("dummy error"))

			return testStruct{
				name:    "if get user by username error, should return error",
				fields:  fields{repo: repo},
				args:    args{req: req},
				want:    "",
				wantErr: true,
			}
		}
		fmt.Printf("%T\n", Error_GetUserByUsername)

		Error_CompareHashAndPassword := func() testStruct {
			repo := mocks.NewIUserRepository(t)

			req := dto.ReqLoginUser{
				Username: "dummy username",
				Password: "dummy password",
			}

			repo.
				On("GetUserByUsername", mock.Anything, req.Username).
				Return(model.User{}, nil)

			return testStruct{
				name:    "if compare hash and password error, should return error",
				fields:  fields{repo: repo},
				args:    args{req: req},
				want:    "",
				wantErr: true,
			}
		}
		fmt.Printf("%T\n", Error_CompareHashAndPassword)

		Success := func() testStruct {
			repo := mocks.NewIUserRepository(t)

			req := dto.ReqLoginUser{
				Username: "dummy username",
				Password: "dummy password",
			}

			id := 234
			hashedPassword, _ := utils.GenerateHashPassword(req.Password)

			repo.
				On("GetUserByUsername", mock.Anything, req.Username).
				Return(model.User{ID: id, Password: hashedPassword}, nil)

			wantTokenString, _ := utils.GenerateUserJWTToken(id)

			return testStruct{
				name:    "if login user success, should return success",
				fields:  fields{repo: repo},
				args:    args{req: req},
				want:    wantTokenString,
				wantErr: false,
			}
		}
		fmt.Printf("%T\n", Success)

		tests = append(tests, Error_Validate())
		tests = append(tests, Error_GetUserByUsername())
		tests = append(tests, Error_CompareHashAndPassword())
		tests = append(tests, Success())
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			us := &UserService{
				repo: tt.fields.repo,
			}
			got, err := us.LoginUser(tt.args.ctx, tt.args.req)
			if (err != nil) != tt.wantErr {
				t.Errorf("UserService.LoginUser() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("UserService.LoginUser() = %v, want %v", got, tt.want)
			}
		})
	}
}
