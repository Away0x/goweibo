package sdhandler

import (
  "fmt"
  "github.com/labstack/echo/v4"
  "github.com/shirou/gopsutil/cpu"
  "github.com/shirou/gopsutil/disk"
  "github.com/shirou/gopsutil/load"
  "github.com/shirou/gopsutil/mem"
  "net/http"
)

const (
  B  = 1
  KB = 1024 * B
  MB = 1024 * KB
  GB = 1024 * MB
)

// HealthCheck shows `OK` as the ping-pong result.
func HealthCheck(c echo.Context) error {
  return c.String(http.StatusOK, "\nOK")
}

// DiskCheck checks the disk usage.
func DiskCheck(c echo.Context) error {
  u, _ := disk.Usage("/")

  usedMB := int(u.Used) / MB
  usedGB := int(u.Used) / GB
  totalMB := int(u.Total) / MB
  totalGB := int(u.Total) / GB
  usedPercent := int(u.UsedPercent)

  status := http.StatusOK
  text := "OK"

  if usedPercent >= 95 {
    status = http.StatusOK
    text = "CRITICAL"
  } else if usedPercent >= 90 {
    status = http.StatusTooManyRequests
    text = "WARNING"
  }

  message := fmt.Sprintf("%s - Free space: %dMB (%dGB) / %dMB (%dGB) | Used: %d%%",
    text, usedMB, usedGB, totalMB, totalGB, usedPercent)

  return c.String(status, "\n"+message)
}

// CPUCheck checks the cpu usage.
func CPUCheck(c echo.Context) error {
  cores, _ := cpu.Counts(false)

  a, _ := load.Avg()
  l1 := a.Load1
  l5 := a.Load5
  l15 := a.Load15

  status := http.StatusOK
  text := "OK"

  if l5 >= float64(cores-1) {
    status = http.StatusInternalServerError
    text = "CRITICAL"
  } else if l5 >= float64(cores-2) {
    status = http.StatusTooManyRequests
    text = "WARNING"
  }

  message := fmt.Sprintf("%s - Load average: %.2f, %.2f, %.2f | Cores: %d",
    text, l1, l5, l15, cores)

  return c.String(status, "\n"+message)
}

// RAMCheck checks the disk usage.
func RAMCheck(c echo.Context) error {
  u, _ := mem.VirtualMemory()

  usedMB := int(u.Used) / MB
  usedGB := int(u.Used) / GB
  totalMB := int(u.Total) / MB
  totalGB := int(u.Total) / GB
  usedPercent := int(u.UsedPercent)

  status := http.StatusOK
  text := "OK"

  if usedPercent >= 95 {
    status = http.StatusInternalServerError
    text = "CRITICAL"
  } else if usedPercent >= 90 {
    status = http.StatusTooManyRequests
    text = "WARNING"
  }

  message := fmt.Sprintf("%s - Free space: %dMB (%dGB) / %dMB (%dGB) | Used: %d%%",
    text, usedMB, usedGB, totalMB, totalGB, usedPercent)

  return c.String(status, "\n"+message)
}

// RegisterSDHandler 注册 sd handlers
func RegisterSDHandlers(e *echo.Echo, prefix string) {
  // 服务器健康自检
  sdRouter := e.Group(prefix)
  {
    sdRouter.GET("/health", HealthCheck).Name = "__sd.health"
    sdRouter.GET("/disk", DiskCheck).Name = "__sd.disk"
    sdRouter.GET("/cpu", CPUCheck).Name = "__sd.cpu"
    sdRouter.GET("/ram", HealthCheck).Name = "__sd.ram"
  }
}
