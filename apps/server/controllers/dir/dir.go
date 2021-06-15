package dir

import (
	"fmt"
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
)

type DirCtl struct {
	router   *gin.RouterGroup
	resource string
}

func New(router *gin.Engine, resource string) *DirCtl {
	r := router.Group(resource)
	ctl := DirCtl{r, resource}
	r.GET("/", ctl.findAll)
	r.GET("/:name", ctl.findOne)
	r.POST("/", ctl.create)
	r.DELETE("/:name", ctl.delete)
	r.PUT("/:name", ctl.update)
	return &ctl
}

func (ctl *DirCtl) findAll(ctx *gin.Context) {
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

func (ctl *DirCtl) findOne(ctx *gin.Context) {

}

func (ctl *DirCtl) create(ctx *gin.Context) {

}

func (ctl *DirCtl) delete(ctx *gin.Context) {

}

func (ctl *DirCtl) update(ctx *gin.Context) {

}
