package api

import (
	"cluster-api-management-dashboard/gintemplrenderer"
	"cluster-api-management-dashboard/views"
	"github.com/gin-gonic/gin"
	"net/http"
)

func GinRoutes(router *gin.Engine) {
	ginHtmlRenderer := router.HTMLRender
	router.HTMLRender = &gintemplrenderer.HTMLTemplRenderer{FallbackHtmlRenderer: ginHtmlRenderer}

	// Disable trusted proxy warning.
	err := router.SetTrustedProxies(nil)
	if err != nil {
		return
	}

	router.GET("/", func(c *gin.Context) {
		c.HTML(http.StatusOK, "", views.Index())
	})

	router.GET("/with-ctx", func(c *gin.Context) {
		r := gintemplrenderer.New(c.Request.Context(), http.StatusOK, views.Index())
		c.Render(http.StatusOK, r)
	})

	router.GET("/with-fallback-renderer", func(c *gin.Context) {
		c.HTML(http.StatusOK, "home.html", gin.H{})
	})
}
