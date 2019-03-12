package mqttlog

import (
	"os"
	"strings"
	"strconv"
	"log"
	//"io"
	//"bufio"
)

const (
	enDate = 0
	enTime = 1
	enId = 2
	enModule = 3
	enData = 4
)

var (
	gLogger = log.New(os.Stdout, "[LDS]", log.Lshortfile | log.Ldate | log.Ltime)
	gLogFileName string = "/home/samba/iWork/logShow/config/rtmp_test_mqtt.log"
	gLogFile *os.File
)

type MqttDocument struct {
	Timestr string
	Id string
	Module string
	Data int64
}

// type RtmpDocument struct {
// 	Timestr string
// 	Module string
// 	Func string
// 	Data int
// }

type MqttType struct {
	doc []MqttDocument
	size int
}

func NewMqttDocument() *MqttDocument {
	return &MqttDocument{
		Timestr: "",
		Module: "",
		Data: -1,
	}
}

func (d *MqttDocument)GetDoc(s, sep string)(err error) {
	var field []string = make([]string, 0)

	for _,v := range strings.Split(s, sep) {
		field = append(field, v)
	}
	
	var n = len(field)
	for i := 0; i < n; i++ {
		switch i {
			case enDate:
				d.Timestr = field[i]
			case enTime:
				d.Timestr +=  " " + field[i]
			case enId:
				d.Id = field[i]
			case enModule:
				d.Module = field[i]
			case enData:
				idx := len(field[i])
				data := field[i][:idx -1]
				d.Data, _ = strconv.ParseInt(data, 10, 64)
			default:
		}
	}
	return 
}