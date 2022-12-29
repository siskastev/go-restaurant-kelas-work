package resto

import (
	"context"
	"errors"
	"github.com/golang/mock/gomock"
	"go-restaurant-kelas-work/internal/model"
	"go-restaurant-kelas-work/internal/model/constant"
	"go-restaurant-kelas-work/internal/repository/menu"
	"go-restaurant-kelas-work/internal/repository/menu/mocks"
	"go-restaurant-kelas-work/internal/repository/order"
	"go-restaurant-kelas-work/internal/repository/user"
	"reflect"
	"testing"
)

func Test_restoUsecase_GetMenuList(t *testing.T) {
	type fields struct {
		menuRepo  menu.Repository
		orderRepo order.Repository
		userRepo  user.Repository
	}
	type args struct {
		ctx      context.Context
		menuType string
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    []model.MenuItem
		wantErr bool
	}{
		// TODO: Add test cases.
		{
			name: "success get menu list",
			fields: fields{
				menuRepo: func() menu.Repository {
					controller := gomock.NewController(t)
					mock := mocks.NewMockMenuRepository(controller)
					mock.EXPECT().GetMenuList(gomock.Any(), string(constant.TypeFood)).
						Times(1).
						Return([]model.MenuItem{
							{
								"bakmie",
								"Bakmie",
								35000,
								constant.TypeFood,
							},
						}, nil)
					return mock
				}(),
			},
			args: args{
				context.Background(),
				string(constant.TypeFood),
			},
			want: []model.MenuItem{
				{
					"bakmie",
					"Bakmie",
					35000,
					constant.TypeFood,
				},
			},
			wantErr: false,
		},
		{
			name: "fail get menu list",
			fields: fields{
				menuRepo: func() menu.Repository {
					controller := gomock.NewController(t)
					mock := mocks.NewMockMenuRepository(controller)
					mock.EXPECT().GetMenuList(gomock.Any(), string(constant.TypeFood)).
						Times(1).
						Return(nil, errors.New("mock error"))
					return mock
				}(),
			},
			args: args{
				context.Background(),
				string(constant.TypeFood),
			},
			want:    nil,
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := &restoUsecase{
				menuRepo:  tt.fields.menuRepo,
				orderRepo: tt.fields.orderRepo,
				userRepo:  tt.fields.userRepo,
			}
			got, err := r.GetMenuList(tt.args.ctx, tt.args.menuType)
			if (err != nil) != tt.wantErr {
				t.Errorf("GetMenuList() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("GetMenuList() got = %v, want %v", got, tt.want)
			}
		})
	}
}
