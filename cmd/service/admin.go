package service

import (
	"html/template"
	"layout/initial"
	"layout/initial/view"
	"layout/routes"
	"os"
	"path/filepath"
	"strings"

	"github.com/abulo/ratel/v3/gin"
	"github.com/abulo/ratel/v3/gin/render"
	"github.com/abulo/ratel/v3/logger"
	"github.com/abulo/ratel/v3/server/xgin"
	"github.com/abulo/ratel/v3/util"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func FormatDateTime(str interface{}) string {

	if util.Empty(str) {
		return ""
	}
	return util.Date("Y-m-d H:i:s", str.(primitive.DateTime).Time())
}

//加载模板文件
func loadTemplate(funcMap template.FuncMap, r render.Delims) (*template.Template, error) {
	t := template.New("").Delims(r.Left, r.Right).Funcs(funcMap)

	for _, name := range view.AssetNames() {
		if !strings.HasSuffix(name, ".html") {
			continue
		}
		asset, err := view.Asset(name)
		if err != nil {
			continue
		}
		name := strings.Replace(name, "view/", "", 1)
		t, err = t.New(name).Parse(string(asset))
		if err != nil {
			return nil, err
		}
	}
	return t, nil
}

func loadGlobTemplate(dir string) ([]string, error) {
	fileList := []string{}
	err := filepath.Walk(dir, func(path string, f os.FileInfo, err error) error {
		if !f.IsDir() {
			fileList = append(fileList, filepath.FromSlash(path))
		}
		return nil
	})
	return fileList, err
}

func (eng *Engine) AdminServer() error {

	configAdmin := initial.Core.Config.Get("http.admin")
	cfg := configAdmin.(map[string]interface{})
	//先获取这个服务是否是需要开启
	if disable := util.ToBool(cfg["Disable"]); disable {
		logger.Logger.Error("http.admin is disabled")
		return nil
	}
	client := xgin.New()
	client.Host = util.ToString(cfg["Host"])
	client.Port = util.ToInt(cfg["Port"])
	client.Deployment = util.ToString(cfg["Deployment"])
	client.DisableMetric = util.ToBool(cfg["DisableMetric"])
	client.DisableTrace = util.ToBool(cfg["DisableTrace"])
	client.DisableSlowQuery = util.ToBool(cfg["DisableSlowQuery"])
	client.ServiceAddress = util.ToString(cfg["ServiceAddress"])
	client.SlowQueryThresholdInMilli = util.ToInt64(cfg["SlowQueryThresholdInMilli"])
	if !initial.Core.Config.Bool("DisableDebug", true) {
		client.Mode = gin.DebugMode
	} else {
		client.Mode = gin.ReleaseMode
	}

	server := client.Build()
	server.SetTrustedProxies([]string{"0.0.0.0/0"})
	//辅助函数
	server.InitFuncMap()
	server.AddFuncMap("config", initial.Core.Config.String)
	server.AddFuncMap("marshalHtml", util.MarshalHTML)
	server.AddFuncMap("marshalJs", util.MarshalJS)
	server.AddFuncMap("static", util.Static)
	server.AddFuncMap("js", util.JS)
	server.AddFuncMap("formatDate", util.FormatDate)
	server.AddFuncMap("formatDateTime", util.FormatDateTime)
	server.AddFuncMap("unixTimeFormatDate", FormatDateTime)
	server.AddFuncMap("inArray", util.InArray)
	server.AddFuncMap("multiArray", util.MultiArray)
	server.AddFuncMap("inMultiArray", util.InMultiArray)
	server.AddFuncMap("empty", util.Empty)
	server.AddFuncMap("divide", util.Divide)
	server.AddFuncMap("add", util.Add)
	server.AddFuncMap("strReplace", util.StrReplace)
	// 开发模式
	if !initial.Core.Config.Bool("DisableDebug", true) {
		//模板
		// server.LoadHTMLGlob(initial.Core.Path + "/view/**/*.html")
		t, err := loadGlobTemplate(initial.Core.Path + "/view")
		if err != nil {
			panic(err)
		}
		server.LoadHTMLFiles(t...)
		server.Use(gin.Logger())
	} else {
		//加载模板文件
		t, err := loadTemplate(server.FuncMap, server.GetDelims())
		if err != nil {
			panic(err)
		}
		server.SetHTMLTemplate(t)
	}
	//静态文件目录
	server.Static("/static", initial.Core.Path+"/static")

	objRoute := initial.Core.Container.Get("routes").(*routes.Routes)
	objRoute.Route(server.Engine)

	if gin.IsDebugging() {
		gin.App.Table.Render()
	}
	return eng.Serve(server)
}
