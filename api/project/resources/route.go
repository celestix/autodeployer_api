package resources

import (
	"github.com/gin-gonic/gin"
	"github.com/shirou/gopsutil/v4/cpu"
	"github.com/shirou/gopsutil/v4/mem"
)

type Response struct {
	Cpu float64 `json:"cpu"`
	Mem float64 `json:"mem"`
}

func resourcesCallback(ctx *gin.Context) {
	perc, err := cpu.Percent(0, false)
	if err != nil {
		perc = []float64{39.0} //placeholder
	}
	v, _ := mem.VirtualMemory()
	var r = Response{
		Cpu: perc[0],
		Mem: v.UsedPercent,
	}
	ctx.JSON(200, r)
}

func Load(rg *gin.RouterGroup) {
	rg.GET("/resources", resourcesCallback)
}
