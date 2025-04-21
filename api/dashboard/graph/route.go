package graph

import (
	"fmt"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/shirou/gopsutil/v4/cpu"
)

var n = 2
var cpuPercent = make([]float64, 3)

type GraphResp struct {
	Cpu  float64 `json:"cpu"`
	Time string  `json:"time"`
}

func graphCallback(ctx *gin.Context) {
	per, err := cpu.Percent(0, false)
	if err != nil {
		fmt.Println("Error getting CPU percent:", err)
	}
	cpuPercent = append(cpuPercent, per[0])[1:]
	resp := make([]GraphResp, 3)
	tn := time.Now()
	hr := tn.Hour()
	mn := tn.Minute()
	for i, x := range cpuPercent {
		currMin := mn + i - 2
		if currMin < 0 {
			currMin = 60 + currMin
		}
		currMinStr := strconv.Itoa(currMin)
		if len(currMinStr) == 1 {
			currMinStr = "0" + currMinStr
		}
		resp[i] = GraphResp{
			Cpu:  x,
			Time: fmt.Sprintf("%d:%s", hr, currMinStr),
		}
		n++
	}
	ctx.JSON(200, resp)
}

func Load(rg *gin.RouterGroup) {
	rg.GET("/graph", graphCallback)
}
