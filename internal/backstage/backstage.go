package backstage

import (
	"net/http"

	"github.com/abulo/ratel/v3/gin"
)

type Backstage struct {
}

func (obj *Backstage) Index(ctx *gin.Context) {
	ctx.HTML(http.StatusOK, "index/index.html", gin.H{})
}
