package render

import (
	"github.com/a-h/templ"
	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/render"
)

func Render(c *gin.Context, status int, template templ.Component) error {
	c.Status(status)
	return template.Render(c.Request.Context(), c.Writer)
}

func Redirect(c *gin.Context, url string, code int) {
	c.Render(-1, render.Redirect{
		Code:     code,
		Location: url,
		Request:  c.Request,
	})
}
