package services

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/shirou/gopsutil/cpu"
	"github.com/shirou/gopsutil/host"
	"github.com/shirou/gopsutil/load"
	"github.com/shirou/gopsutil/mem"
)

func GetSystemInfo(ctx *gin.Context) {

	sysInfo := make(map[string]interface{})

	memory, err := mem.VirtualMemory()
	if err != nil {
		sysInfo["memory"] = ""
	} else {
		sysInfo["memory"] = memory
	}
	cpu, err := cpu.Info()
	if err != nil {
		sysInfo["cpu"] = ""
	} else {
		sysInfo["cpu"] = cpu
	}
	host, err := host.Info()
	if err != nil {
		sysInfo["host"] = ""
	} else {
		sysInfo["host"] = host
	}
	load, err := load.Avg()
	if err != nil {
		sysInfo["load"] = ""
	} else {
		sysInfo["load"] = load
	}
	ctx.JSON(http.StatusOK, sysInfo)
}
