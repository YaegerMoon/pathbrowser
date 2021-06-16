package slide

import (
	"fmt"
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
)

type SlideCtl struct {
	router   *gin.RouterGroup
	resource string
}

func New(router *gin.Engine, resource string) *SlideCtl {
	r := router.Group(resource)
	ctl := SlideCtl{r, resource}
	r.GET("/", ctl.findAll)
	r.GET("/:name", ctl.findOne)
	r.POST("/", ctl.create)
	r.DELETE("/:name", ctl.delete)
	r.PUT("/:name", ctl.update)
	return &ctl
}

func (ctl *SlideCtl) findAll(ctx *gin.Context) {
	entry, err := os.ReadDir("./")

	if err != nil {
		ctx.JSON(http.StatusInternalServerError, err)
		return
	}
	for _, val := range entry {
		fsInfo, _ := val.Info()
		fmt.Println(fsInfo.Mode().Type())
	}

	ctx.JSON(http.StatusOK, entry)

}

func (ctl *SlideCtl) findOne(ctx *gin.Context) {

}

func (ctl *SlideCtl) create(ctx *gin.Context) {

}

func (ctl *SlideCtl) delete(ctx *gin.Context) {

}

func (ctl *SlideCtl) update(ctx *gin.Context) {

}
