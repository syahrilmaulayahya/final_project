package users

import (
	"final_project/app/middleware"
	"final_project/business/users"
	"final_project/controllers"
	"final_project/controllers/users/requests"
	"final_project/controllers/users/respons"
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
)

type UserController struct {
	UserUseCase users.UseCase
}

func NewUserController(userUseCase users.UseCase) *UserController {
	return &UserController{
		UserUseCase: userUseCase,
	}
}
func (usercontroller UserController) Register(c echo.Context) error {

	userRegister := requests.UserRegister{}
	c.Bind(&userRegister)
	register := userRegister.ToDomain()
	ctx := c.Request().Context()
	user, err := usercontroller.UserUseCase.Register(ctx, register)

	if err != nil {
		return controllers.NewErrorResponse(c, http.StatusBadRequest, err)
	}
	return controllers.NewSuccessResponse(c, respons.NoTokenFromDomain(user))

}
func (usercontroller UserController) Login(c echo.Context) error {
	userLogin := requests.UserLogin{}
	c.Bind(&userLogin)

	ctx := c.Request().Context()

	user, err := usercontroller.UserUseCase.Login(ctx, userLogin.Email, userLogin.Password)

	if err != nil {
		return controllers.NewErrorResponse(c, http.StatusForbidden, err)
	}
	return controllers.NewSuccessResponse(c, respons.FromDomain(user))
}

func (UserController UserController) Details(c echo.Context) error {
	ctx := c.Request().Context()
	id, _ := strconv.Atoi(c.Param("id"))
	user, err := UserController.UserUseCase.Details(ctx, id)
	if err != nil {
		return controllers.NewErrorResponse(c, http.StatusNoContent, err)
	}
	return controllers.NewSuccessResponse(c, respons.NoTokenFromDomain(user))
}

func (UserController UserController) UploadReview(c echo.Context) error {
	userReview := requests.UserReview{}
	c.Bind(&userReview)
	newReview := userReview.ToDomain()
	ctx := c.Request().Context()
	review, err := UserController.UserUseCase.UploadReview(ctx, newReview, middleware.GetClaimsUserId(c))
	if err != nil {
		return controllers.NewErrorResponse(c, http.StatusBadRequest, err)
	}
	return controllers.NewSuccessResponse(c, respons.ReviewFromDomain(review))
}
