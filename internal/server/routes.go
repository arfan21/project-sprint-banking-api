package server

import (
	fileuploaderctrl "github.com/arfan21/project-sprint-banking-api/internal/fileuploader/controller"
	fileuploadersvc "github.com/arfan21/project-sprint-banking-api/internal/fileuploader/service"
	transactionctrl "github.com/arfan21/project-sprint-banking-api/internal/transaction/controller"
	transactionrepo "github.com/arfan21/project-sprint-banking-api/internal/transaction/repository"
	transactionsvc "github.com/arfan21/project-sprint-banking-api/internal/transaction/service"
	userctrl "github.com/arfan21/project-sprint-banking-api/internal/user/controller"
	userrepo "github.com/arfan21/project-sprint-banking-api/internal/user/repository"
	usersvc "github.com/arfan21/project-sprint-banking-api/internal/user/service"
	"github.com/arfan21/project-sprint-banking-api/pkg/middleware"
	"github.com/gofiber/fiber/v2"
)

func (s *Server) Routes() {

	api := s.app.Group("")
	api.Get("/healthz", func(c *fiber.Ctx) error {
		err := s.db.Ping()
		if err != nil {
			return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
				"message": "database is not connected",
			})
		}

		return c.JSON(fiber.Map{
			"message": "OK, test update",
		})
	})

	userRepo := userrepo.NewStdLib(s.db)
	userSvc := usersvc.New(userRepo)
	userCtrl := userctrl.New(userSvc)

	fileUploaderSvc := fileuploadersvc.New()
	fileUploaderCtrl := fileuploaderctrl.New(fileUploaderSvc)

	transactionRepo := transactionrepo.NewStdLib(s.db)
	transactionSvc := transactionsvc.New(transactionRepo)
	transactionCtrl := transactionctrl.New(transactionSvc)

	s.RoutesCustomer(api, userCtrl)
	s.RoutesFileUploader(api, fileUploaderCtrl)
	s.RoutesTransaction(api, transactionCtrl)
}

func (s Server) RoutesCustomer(route fiber.Router, ctrl *userctrl.ControllerHTTP) {
	v1 := route.Group("/v1")
	usersV1 := v1.Group("/user")
	usersV1.Post("/register", ctrl.Register)
	usersV1.Post("/login", ctrl.Login)
}

func (s Server) RoutesFileUploader(route fiber.Router, ctrl *fileuploaderctrl.ControllerHTTP) {
	v1 := route.Group("/v1")
	fileUploaderV1 := v1.Group("/image", middleware.JWTAuth)
	fileUploaderV1.Post("", ctrl.UploadImage)
}

func (s Server) RoutesTransaction(route fiber.Router, ctrl transactionctrl.ControllerHTTP) {
	v1 := route.Group("/v1", middleware.JWTAuth)
	// balance routes
	balanceV1 := v1.Group("/balance")
	balanceV1.Post("", ctrl.AddBalance)
	balanceV1.Get("", ctrl.GetBalance)
	balanceV1.Get("/history", ctrl.GetListTransaction)
	balanceV1.Get("/history/:id", ctrl.GetListTransaction)

	// transaction routes
	transactionV1 := v1.Group("/transaction")
	transactionV1.Post("", ctrl.TransferBalance)

}
