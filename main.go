package main

import (
	"encoding/json"
	"fmt"
	gin "github.com/gin-gonic/gin"
	"io/ioutil"
	"net/http"
	"strconv"
	"time"
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

func addDevice() {

	request, err := http.NewRequest("GET", "http://192.168.1.230:8181/onos/mao/MaoIntegration/netconf/biKnownLinks", nil)
	request.Header.Add("Authorization", "Basic a2FyYWY6a2FyYWY=")
	if err != nil {
		fmt.Println(err)
		return
	}

	//resp, err := http.Get("http://192.168.1.230:8181/onos/mao/MaoIntegration/netconf/biKnownLinks")
	client := http.Client{}
	resp, err := client.Do(request)
	if err != nil {
		fmt.Println(err)
		return
	}
	defer resp.Body.Close()

	body,err := ioutil.ReadAll(resp.Body)
	//str := string(body)
	//fmt.Println(str)


	links := make([]map[string]interface{}, 0)
	err = json.Unmarshal(body, &links)
	if err != nil {
		fmt.Println("err")
		fmt.Println(err)
		return
	}
	//fmt.Println("*************************")
	//for index, link := range links {
	//	fmt.Println("============= " + strconv.Itoa(index) + " =============")
	//	fmt.Println(link["localDeviceName"])
	//	fmt.Println(link["localPortName"])
	//	fmt.Println(link["remoteDeviceName"])
	//	fmt.Println(link["remotePortName"])
	//}

	//fmt.Println(links)
}

func startRestful() {
	fmt.Printf("qingdao\n")
	//gin.SetMode(gin.ReleaseMode)
	restful := gin.Default()
	restful.LoadHTMLFiles("index.html")
	restful.GET("/", registerPage)
	restful.POST("/register", register)
	//restful.GET("/test", addDevice)
	restful.Run()
}

func main() {
	go startRestful()

	count := 1
	for {
		if count % 1000 == 0 {
			fmt.Println(count)
			fmt.Println(time.Now())
		}
		count++
		addDevice()
		//time.Sleep(3 * time.Second)
	}
}
