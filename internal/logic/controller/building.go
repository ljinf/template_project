package controller

import (
	"errors"
	"github.com/gin-gonic/gin"
	"github.com/ljinf/template_project/internal/app"
	"github.com/ljinf/template_project/pkg/config"
	errcode2 "github.com/ljinf/template_project/pkg/errcode"
	"github.com/ljinf/template_project/pkg/logger"
	"net/http"
)

// 存放一些项目搭建过程中验证效果用的接口Handler, 之前搭建过程中在main包中写的测试接口也挪到了这里

func TestPing(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"message": "pong",
	})
	return
}

func TestConfigRead(c *gin.Context) {
	database := config.Database
	c.JSON(http.StatusOK, gin.H{
		"type":     database.Type,
		"max_life": database.Master.MaxLifeTime,
	})
	return
}

func TestLogger(c *gin.Context) {
	logger.Info(c, "logger test", "key", "keyName", "val", 2)
	c.JSON(http.StatusOK, gin.H{
		"status": "ok",
	})
	return
}

func TestAccessLog(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"status": "ok",
	})
	return
}

func TestPanicLog(c *gin.Context) {
	var a map[string]string
	a["k"] = "v"
	c.JSON(http.StatusOK, gin.H{
		"status": "ok",
		"data":   a,
	})
	return
}

func TestAppError(c *gin.Context) {

	// 使用 Wrap 包装原因error 生成 项目error
	err := errors.New("a dao error")
	appErr := errcode2.Wrap("包装错误", err)
	bAppErr := errcode2.Wrap("再包装错误", appErr)
	logger.Error(c, "记录错误", "err", bAppErr)

	// 预定义的ErrServer, 给其追加错误原因的error
	err = errors.New("a domain error")
	apiErr := errcode2.ErrServer.WithCause(err)
	logger.Error(c, "API执行中出现错误", "err", apiErr)

	c.JSON(apiErr.HttpStatusCode(), gin.H{
		"code": apiErr.Code(),
		"msg":  apiErr.Msg(),
	})
	return
}

func TestResponseObj(ctx *gin.Context) {
	data := map[string]int{
		"a": 1,
		"b": 2,
	}
	app.Success(ctx, data)
	return
}

func TestResponseList(ctx *gin.Context) {

	pagination := app.NewPagination(ctx)
	// Mock fetch list data from db
	data := []struct {
		Name string `json:"name"`
		Age  int    `json:"age"`
	}{
		{
			Name: "Lily",
			Age:  26,
		},
		{
			Name: "Violet",
			Age:  25,
		},
	}
	pagination.SetTotalRows(2)
	app.SuccessWithPagination(ctx, data, pagination)
	return
}

func TestResponseError(ctx *gin.Context) {
	baseErr := errors.New("a dao error")
	// 这一步正式开发时写在service层
	err := errcode2.Wrap("encountered an error when xxx service did xxx", baseErr)
	app.Error(ctx, errcode2.ErrServer.WithCause(err))
	return
}
