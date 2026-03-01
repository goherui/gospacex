package service

import (
	"gospaacex/bff/basic/config"
	"gospaacex/bff/handler/request"
	"gospaacex/bff/handler/response"
	__ "gospaacex/proto"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

func PosCreate(c *gin.Context) {
	var form request.PosCreate
	// 根据 Content-Type Header 推断使用哪个绑定器。
	if err := c.ShouldBind(&form); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
			"code":  http.StatusBadRequest,
			"msg":   "参数错误",
		})
		return
	}
	r, err := config.PosClient.PosCreate(c, &__.PosCreateReq{
		Title:        form.Title,
		Company:      form.Company,
		Salary:       form.Salary,
		Location:     form.Location,
		Description:  form.Description,
		Requirements: form.Requirements,
	})
	if err != nil {
		log.Fatalf("could not greet: %v", err)
	}
	c.JSON(http.StatusOK, response.PosCreate{
		Code: int(r.Code),
		Msg:  r.Msg,
	})
}
func PosDel(c *gin.Context) {
	var form request.PosDel
	// 根据 Content-Type Header 推断使用哪个绑定器。
	if err := c.ShouldBind(&form); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
			"code":  http.StatusBadRequest,
			"msg":   "参数错误",
		})
		return
	}
	r, err := config.PosClient.PosDel(c, &__.PosDelReq{
		Id: int64(form.Id),
	})
	if err != nil {
		log.Fatalf("could not greet: %v", err)
	}
	c.JSON(http.StatusOK, response.PosDel{
		Code: int(r.Code),
		Msg:  r.Msg,
	})
}
func PosUpdate(c *gin.Context) {
	var form request.PosUpdate
	// 根据 Content-Type Header 推断使用哪个绑定器。
	if err := c.ShouldBind(&form); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
			"code":  http.StatusBadRequest,
			"msg":   "参数错误",
		})
		return
	}
	r, err := config.PosClient.PosUpdate(c, &__.PosUpdateReq{
		Title:        form.Title,
		Company:      form.Company,
		Salary:       form.Salary,
		Location:     form.Location,
		Description:  form.Description,
		Requirements: form.Requirements,
		Id:           int64(form.Id),
	})
	if err != nil {
		log.Fatalf("could not greet: %v", err)
	}
	c.JSON(http.StatusOK, response.PosCUpdate{
		Code: int(r.Code),
		Msg:  r.Msg,
	})
}
func PosList(c *gin.Context) {
	var form request.PosList
	// 根据 Content-Type Header 推断使用哪个绑定器。
	if err := c.ShouldBind(&form); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
			"code":  http.StatusBadRequest,
			"msg":   "参数错误",
		})
		return
	}
	r, err := config.PosClient.PosList(c, &__.PosListReq{
		Page: int64(form.Page),
		Size: int64(form.Size),
	})
	if err != nil {
		log.Fatalf("could not greet: %v", err)
	}
	var list []response.Pos
	for _, position := range r.List {
		list = append(list, response.Pos{
			ID:           position.Id,
			Title:        position.Title,
			Company:      position.Company,
			Salary:       position.Salary,
			Location:     position.Location,
			Description:  position.Description,
			Requirements: position.Requirements,
		})
	}
	c.JSON(http.StatusOK, response.PosList{
		List: list,
		Code: int(r.Code),
		Msg:  r.Msg,
	})
}
