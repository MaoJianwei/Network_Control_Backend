package main

import (
	"fmt"
	gin "github.com/gin-gonic/gin"
	"io/ioutil"
	"net/http"
	"strconv"
)

func maoGet(c *gin.Context) {
	c.JSON(200, gin.H{
		"beijing": 118.5,
	})
}

func registerPage(c *gin.Context) {
	c.HTML(200, "index.html", nil)
}

type DeviceConnector struct {
	ip string `form:"ipv4"`
	port uint16 `form:"port"`
}

func register(c *gin.Context) {
	//data := DeviceConnector{}
	//err := c.ShouldBind(&data)
	//if err == nil {
	//	c.JSON(200, gin.H{"ipv4": data.ip, "port": data.port})
	//} else {
	//	c.JSON(509, gin.H{"err": err})
	//}

	data := DeviceConnector{}
	data.ip = c.PostForm("ipv4")
	tmpPort,_ := strconv.ParseUint(c.PostForm("port"),0, 16)
	data.port = uint16(tmpPort)

	c.JSON(200, gin.H{"ipv4": data.ip, "port": data.port})
}

func addDevice(c *gin.Context) {
	resp, err := http.Get("http://127.0.0.1:8080")
	if err != nil {

	}

	defer resp.Body.Close()

	body,err := ioutil.ReadAll(resp.Body)
	str := string(body)

	fmt.Println(str)
	c.JSON(200, "MAO-OK")
}

func main() {
	fmt.Printf("qingdao\n")
	//gin.SetMode(gin.ReleaseMode)
	restful := gin.Default()
	restful.LoadHTMLFiles("index.html")
	restful.GET("/", registerPage)
	restful.POST("/register", register)
	restful.GET("/test", addDevice)
	restful.Run()
}
