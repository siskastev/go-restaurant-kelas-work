package resto_test

import (
	"github.com/golang/mock/gomock"
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
	"go-restaurant-kelas-work/internal/model"
	"go-restaurant-kelas-work/internal/model/constant"
	mocksMenu "go-restaurant-kelas-work/internal/repository/menu/mocks"
	mocksOrder "go-restaurant-kelas-work/internal/repository/order/mocks"
	mocksUser "go-restaurant-kelas-work/internal/repository/user/mocks"
	"go-restaurant-kelas-work/internal/usecase/resto"
	"golang.org/x/net/context"
)

var _ = Describe("GinkgoResto", func() {
	var usecase resto.Usecase
	var menuRepoMock *mocksMenu.MockMenuRepository
	var orderRepoMock *mocksOrder.MockOrderRepository
	var userRepoMock *mocksUser.MockUserRepository
	var mockController *gomock.Controller

	BeforeEach(func() {
		mockController = gomock.NewController(GinkgoT())
		menuRepoMock = mocksMenu.NewMockMenuRepository(mockController)
		orderRepoMock = mocksOrder.NewMockOrderRepository(mockController)
		userRepoMock = mocksUser.NewMockUserRepository(mockController)
		usecase = resto.GetUsecase(menuRepoMock, orderRepoMock, userRepoMock)
	})

	Describe("request order info", func() {
		Context("it gave the correct inputs", func() {
			inputs := model.GetOrderInfoRequest{
				OrderID: "valid_order_id",
				UserID:  "valid_user_id",
			}

			When("the requested OrderID is not the user's", func() {
				BeforeEach(func() {
					orderRepoMock.EXPECT().GetOrderInfo(gomock.Any(), inputs.OrderID).
						Times(1).
						Return(model.Order{
							OrderID:       "valid_order_id",
							UserID:        "valid_user_id_2",
							Status:        constant.OrderStatusProcessed,
							ProductOrders: []model.ProductOrder{},
							ReferenceID:   "ref_id",
						}, nil)
				})

				It("returns unauthorized error", func() {
					res, err := usecase.GetOrderInfo(context.Background(), inputs)
					Expect(err).Should(HaveOccurred())
					Expect(err.Error()).To(BeEquivalentTo("unauthorized"))
					Expect(res).To(BeEquivalentTo(model.Order{}))
				})
			})

			When("the requested OrderID is not the user's", func() {
				BeforeEach(func() {
					orderRepoMock.EXPECT().GetOrderInfo(gomock.Any(), inputs.OrderID).
						Times(1).
						Return(model.Order{
							OrderID:       "valid_order_id",
							UserID:        "valid_user_id",
							Status:        constant.OrderStatusProcessed,
							ProductOrders: []model.ProductOrder{},
							ReferenceID:   "ref_id",
						}, nil)
				})
				It("returns unauthorized error", func() {
					res, err := usecase.GetOrderInfo(context.Background(), inputs)
					Expect(err).ShouldNot(HaveOccurred())
					Expect(res).To(BeEquivalentTo(model.Order{
						OrderID:       "valid_order_id",
						UserID:        "valid_user_id",
						Status:        constant.OrderStatusProcessed,
						ProductOrders: []model.ProductOrder{},
						ReferenceID:   "ref_id",
					}))
				})
			})
		})
	})
})
