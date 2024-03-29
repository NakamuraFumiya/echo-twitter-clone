package detail

import (
	"echo-twitter-clone/core/usecase/post"
	"net/http"

	"github.com/labstack/echo/v4"
)

type Controller struct {
	PostUseCase post.PostUseCase
}

func NewController(useCase post.PostUseCase) Controller {
	return Controller{useCase}
}

func (c Controller) GetPost(ctx echo.Context) error {
	post, err := c.PostUseCase.GetPost(ctx)
	if err != nil {
		return echo.NewHTTPError(http.StatusNotFound, "post does not exist")
	}
	return ctx.JSON(http.StatusOK, post)
}
