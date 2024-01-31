package main

import(
"fmt"
"github.com/gin-gonic-gin"
)

type Device struct {
	ID string
	Name string
	Location string
	Status bool
}

var devices []Device

func main() {
r := gin.Default()

r.GET("/devices",getDevices)

r.GET("/devices/:id",getDeviceByID)

r.POST("/devices/:id/control", ControlDevice)

r.Run(":8080")
}

func getDevices(c *gin.Context) {
c.JSON(200,devices)
}

func getDeviceByID(c *gin.Context) {
id:= c.Params("id")
for _,d:=  range devices {
	if d.ID == id {
		c.JSON(200,d)
		return
	}
}
c.JSON(400,gin.H{"error":"device not found"})
}

func ControlDevice(c *gin.Context) {
id:= c.Params("Id")
for i,d:= range devices {
	if d.ID == id {
		devices[i].Status = !devices[i].Status
		c.JSON(200,devices[i])
		return
	}
}
c.JSON(400,gin.H{"error":"Device not found"})
}


func init() {
	devices = append(devices,Device{ID:"1",Name:"smart Bulb",Location:"Living Room",Status: false})
	devices = append(devices,Device{ID:"2",Name:"Smart Thermostat",Location:"Bedroom",Status: true})
	devices = append(devices,Device{ID:"3",Name:"smart Lock",Location:"Frontdoor",Status: false})
	
	fmt.Println("Smart Home Automation System Started.")

}
