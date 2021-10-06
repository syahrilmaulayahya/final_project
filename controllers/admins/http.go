package admins

import (
	"final_project/business/admins"
	"final_project/controllers"
	"final_project/controllers/admins/requests"
	"final_project/controllers/admins/respons"
	"net/http"

	"github.com/labstack/echo/v4"
)

type AdminController struct {
	AdminUseCase admins.UseCase
}

func NewAdminController(adminUseCase admins.UseCase) *AdminController {
	return &AdminController{
		AdminUseCase: adminUseCase,
	}
}

func (AdminController AdminController) Register(c echo.Context) error {
	adminRegister := requests.AdminRegister{}
	c.Bind(&adminRegister)
	register := adminRegister.ToDomain()
	ctx := c.Request().Context()
	admin, err := AdminController.AdminUseCase.Register(ctx, register)
	if err != nil {
		return controllers.NewErrorResponse(c, http.StatusBadRequest, err)
	}
	return controllers.NewSuccessResponse(c, respons.AdminNoTokenFromDomain(admin))
}

func (AdminController AdminController) Login(c echo.Context) error {
	adminLogin := requests.AdminLogin{}
	c.Bind(&adminLogin)
	ctx := c.Request().Context()
	admin, err := AdminController.AdminUseCase.Login(ctx, adminLogin.Email, adminLogin.Password)
	if err != nil {
		return controllers.NewErrorResponse(c, http.StatusForbidden, err)
	}
	return controllers.NewSuccessResponse(c, respons.AdminFromDomain(admin))
}
