package user

import (
	"cloud/api/common"
	"cloud/code"
	"cloud/internal/response"
	"context"
	"strings"

	"github.com/abulo/ratel/v3/util"
	"github.com/cloudwego/hertz/pkg/app"
	"github.com/cloudwego/hertz/pkg/common/utils"
	"github.com/cloudwego/hertz/pkg/protocol/consts"
)

func SysUserCode(ctx context.Context, newCtx *app.RequestContext) {
	tenantId := newCtx.GetInt64("tenantId") // 租户 Id
	sysUserId := newCtx.GetInt64("userId")  // 用户 Id
	list, currentMenuIds, err := common.SysUserTree(ctx, sysUserId, tenantId)
	if err != nil {
		response.JSON(newCtx, consts.StatusOK, utils.H{
			"code": code.TokenError,
			"msg":  err.Error(),
		})
		return
	}
	var codeList []string
	for _, item := range list {
		if util.InArray(item.Id, currentMenuIds) {
			if !util.Empty(item.Code) {
				codeList = append(codeList, item.Code)
			}
		}
	}
	resList := make(map[string][]string)
	for _, v := range codeList {
		if !strings.Contains(v, ".") {
			resList[v] = append(resList[v], v)
		} else {
			tmp := strings.Split(v, ".")
			resList[tmp[0]] = append(resList[tmp[0]], v)
		}
	}
	response.JSON(newCtx, consts.StatusOK, utils.H{
		"code": code.Success,
		"msg":  code.StatusText(code.Success),
		"data": resList,
	})
}
