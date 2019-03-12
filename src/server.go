package main

import (
	 "os"
	 "log"
	 "io"
	 "fmt"
	 "bufio"
	// "strings"
	// "strconv"
	// "time"

	"handler"
	"mqttlog"
	"dbdeal"
	"gopkg.in/gin-gonic/gin.v1"
	// "net/http"
)

var (
	gLogFileName string = "/home/samba/iWork/logShow/config/rtmp_test_mqtt.log"
	gLogger = log.New(os.Stdout, "[LDS]", log.Lshortfile | log.Ldate | log.Ltime)
	gLogFile *os.File
	rd *bufio.Reader
)

func LogInit()(err error) {

	gLogFile, err = os.Open(gLogFileName)
	if err != nil {
		gLogger.Printf("Open % failed.\n", gLogFileName)
		return err
	}

	rd = bufio.NewReader(gLogFile)
	return
}

func LogExit() {
	gLogFile.Close()
}

func LogReadLine()(line string, err error) {
	// var logRow string
	for {
		logRow, err := rd.ReadString('\n')
		if err == io.EOF {
			fmt.Printf("===io.EOF\n", err)
			return "", err
		}
		if err != nil {
			gLogger.Println("ReadString failed.")
			return "", err
		}
		if logRow == "" {
			fmt.Println("line is nul")
		}else {
			fmt.Printf("ReadString:%s\n", logRow)
			return logRow, nil
		}
	}
}

func CreateMqttrepo() {
	var m map[string]string 
	m = make(map[string]string)
	m["timestamp"] = "date"
	m["id"] = "string"
	m["module"] = "string"
	m["data"] = "long"
	dbdeal.CreateDocument(m)
}

func UpdateMqttrepo() {

	lineCount := 0
	for {
		line, err := LogReadLine()
		if err != nil {
			fmt.Println("LogReadLine failed. lineCount:", lineCount)
			return
		}
		if line == "" {
			fmt.Println("lineCount:", lineCount)
			break
		}
		fmt.Printf("line: %s\n", line)
		lineCount += 1
		mqttInfo := mqttlog.NewMqttDocument()
		mqttInfo.GetDoc(line, " ")
		dbdeal.AddLog(mqttInfo)
		//  if lineCount >= 40 {
		// 	 break;
		//  }
		fmt.Printf("mqttInfo.timestr:%s  module:%s  data:%d\n", mqttInfo.Timestr, mqttInfo.Module, mqttInfo.Data)
	}
}

func DelMqttrepo() {
	 dbdeal.DeleteRepo()
}

func main() {

	args := os.Args 
	operate := args[1]

	LogInit()
	dbdeal.DAInit()
	fmt.Printf("=================================================================\n")
	switch operate {
		case "d": {
			fmt.Println("delete repo")
			DelMqttrepo()
		}
		case "c": {
			fmt.Println("create repo")
			CreateMqttrepo()
		}
		case "u": {
			fmt.Println("update repo")
			UpdateMqttrepo()
		}
		case "g": {
			fmt.Println("get repo data")
			dbdeal.GetMqttBeat()
		}
	}
	//fmt.Println("-----------------------------------------------------------------")
	//dbdeal.GetMqttBeatHistogram()
	fmt.Printf("=================================================================\n")
	//使用gin Default方法创建一个路由handler:router
	router := gin.Default()    
    //通过HTTP方法绑定路由规则和路由函数
	//router.GET("/ops/GetName", GetName)
	router.GET("/ops/GetLogCount", handler.GetLogCount)
	
	//启动路由，监听端口
	//router.Run(":1323")
}
