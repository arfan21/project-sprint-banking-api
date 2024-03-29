package transactionctrl

import (
	"github.com/arfan21/project-sprint-banking-api/internal/model"
	transactionsvc "github.com/arfan21/project-sprint-banking-api/internal/transaction/service"
	"github.com/arfan21/project-sprint-banking-api/pkg/constant"
	"github.com/arfan21/project-sprint-banking-api/pkg/exception"
	"github.com/arfan21/project-sprint-banking-api/pkg/logger"
	"github.com/arfan21/project-sprint-banking-api/pkg/pkgutil"
	"github.com/arfan21/project-sprint-banking-api/pkg/validation"
	"github.com/gofiber/fiber/v2"
)

type ControllerHTTP struct {
	svc transactionsvc.Service
}

func New(svc transactionsvc.Service) ControllerHTTP {
	return ControllerHTTP{svc: svc}
}

// @Summary Add Balance
// @Description Add balance to user account
// @Tags Balance
// @Accept json
// @Produce json
// @Param body body model.TransactionAddBalanceRequest true "Payload balance add request"
// @Success 200 {object} pkgutil.HTTPResponse
// @Failure 400 {object} pkgutil.HTTPResponse{data=[]pkgutil.ErrValidationResponse} "Error validation field"
// @Failure 500 {object} pkgutil.HTTPResponse
// @Router /v1/balance [post]
func (ctrl ControllerHTTP) AddBalance(c *fiber.Ctx) error {
	claims, ok := c.Locals(constant.JWTClaimsContextKey).(model.JWTClaims)
	if !ok {
		logger.Log(c.UserContext()).Error().Msg("cannot get claims from context")
		return c.Status(fiber.StatusForbidden).JSON(fiber.Map{
			"message": "invalid or expired token",
		})
	}

	var req model.TransactionAddBalanceRequest
	err := c.BodyParser(&req)
	exception.PanicIfNeeded(err)

	req.UserID = claims.Subject

	err = ctrl.svc.AddBalance(c.UserContext(), req)
	exception.PanicIfNeeded(err)

	return c.JSON(pkgutil.HTTPResponse{
		Message: "Balance added successfully",
	})
}

// @Summary Transfer Balance
// @Description Transfer balance to bank
// @Tags Transaction
// @Accept json
// @Produce json
// @Param body body model.TransactionTransferBalanceRequest true "Payload balance transfer request"
// @Success 200 {object} pkgutil.HTTPResponse
// @Failure 400 {object} pkgutil.HTTPResponse{data=[]pkgutil.ErrValidationResponse} "Error validation field"
// @Failure 500 {object} pkgutil.HTTPResponse
// @Router /v1/transaction [post]
func (ctrl ControllerHTTP) TransferBalance(c *fiber.Ctx) error {
	claims, ok := c.Locals(constant.JWTClaimsContextKey).(model.JWTClaims)
	if !ok {
		logger.Log(c.UserContext()).Error().Msg("cannot get claims from context")
		return c.Status(fiber.StatusForbidden).JSON(fiber.Map{
			"message": "invalid or expired token",
		})
	}

	var req model.TransactionTransferBalanceRequest
	err := c.BodyParser(&req)
	exception.PanicIfNeeded(err)

	req.UserID = claims.Subject

	err = ctrl.svc.TransferBalance(c.UserContext(), req)
	exception.PanicIfNeeded(err)

	return c.JSON(pkgutil.HTTPResponse{
		Message: "Balance transfered successfully",
	})
}

// @Summary Get Balance
// @Description Get balance from user account
// @Tags Balance
// @Accept json
// @Produce json
// @Success 200 {object} pkgutil.HTTPResponse{data=model.BalanceGetResponse}
// @Failure 500 {object} pkgutil.HTTPResponse
// @Router /v1/balance [get]
func (ctrl ControllerHTTP) GetBalance(c *fiber.Ctx) error {
	claims, ok := c.Locals(constant.JWTClaimsContextKey).(model.JWTClaims)
	if !ok {
		logger.Log(c.UserContext()).Error().Msg("cannot get claims from context")
		return c.Status(fiber.StatusForbidden).JSON(fiber.Map{
			"message": "invalid or expired token",
		})
	}

	balance, err := ctrl.svc.GetBalanceByUserID(c.UserContext(), claims.Subject)
	exception.PanicIfNeeded(err)

	return c.JSON(pkgutil.HTTPResponse{
		Message: "success",
		Data:    balance,
	})
}

// @Summary Get Transaction
// @Description Get transaction from user account
// @Tags Balance
// @Accept json
// @Produce json
// @Param page query int false "Page number"
// @Param limit query int false "Limit data"
// @Success 200 {object} pkgutil.HTTPResponse{data=[]model.TransactionGetResponse}
// @Failure 500 {object} pkgutil.HTTPResponse
// @Router /v1/balance/history [get]
func (ctrl ControllerHTTP) GetListTransaction(c *fiber.Ctx) error {
	claims, ok := c.Locals(constant.JWTClaimsContextKey).(model.JWTClaims)
	if !ok {
		logger.Log(c.UserContext()).Error().Msg("cannot get claims from context")
		return c.Status(fiber.StatusForbidden).JSON(fiber.Map{
			"message": "invalid or expired token",
		})
	}

	mapQuery := c.Queries()
	err := validation.ValidateQuery(mapQuery)
	exception.PanicIfNeeded(err)

	var req model.TransactionGetListRequest
	err = c.QueryParser(&req)
	exception.PanicIfNeeded(err)

	req.UserID = claims.Subject

	data, meta, err := ctrl.svc.GetListByUserID(c.UserContext(), req)
	exception.PanicIfNeeded(err)

	return c.JSON(pkgutil.HTTPResponse{
		Message: "success",
		Data:    data,
		Meta:    meta,
	})
}
