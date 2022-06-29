package routes

import (
	"layout/initial"
	"layout/internal/backstage"

	"github.com/abulo/ratel/v3/gin"
)

type Routes struct {
}

func (obj *Routes) Route(r *gin.Engine) {
	objBackstage := initial.Core.Container.Get("internal.backstage").(*backstage.Backstage)
	r.GET("/index", "index", objBackstage.Index)
}
