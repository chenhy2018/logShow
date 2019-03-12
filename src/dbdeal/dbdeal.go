package dbdeal

import (
	// "log"
	// "os"
	"fmt"
	"time"

	. "mqttlog"
	. "github.com/qiniu/pandora-go-sdk/logdb"
	"github.com/qiniu/pandora-go-sdk/base/config"
	. "github.com/qiniu/pandora-go-sdk/base"
)

var (
	logdbCfg *config.Config
	logdbClient LogdbAPI
	endpoint = config.DefaultLogDBEndpoint
	akLink = "your ak"
	skLink = "your sk"
	repoName = "ipcstatusrepo"//"ipcuasrepo"//"mqttrepo"
	gLogger Logger
	//logger = log.New(os.Stdout, "[LDS]", log.Lshortfile | log.Ldate | log.Ltime)
)

func CreateDocument(m map[string]string)(err error) {
	//map[string:strsing]
	var mqttRepoSchema []RepoSchemaEntry

	for key := range m {
		var repoSchemaInfo RepoSchemaEntry
		repoSchemaInfo.Key = key
		repoSchemaInfo.ValueType = m[key]
		fmt.Println("Key:", key, "Value:", m[key])
		mqttRepoSchema = append(mqttRepoSchema, repoSchemaInfo)
	}

	createInput := &CreateRepoInput{
		RepoName : repoName,
		Region : "nb",
		Schema : mqttRepoSchema,
		Retention : "30d",
	}
	err = logdbClient.CreateRepo(createInput)
	if err != nil {
		fmt.Println(err)
		return err
	}
	return
}

func AddLog(logInfo *MqttDocument)(err error) {

	//log := make(Log)
	// log["timestamp"] = logInfo.Timestr
	// log["module"] = logInfo.Module 
	// log["data"] = logInfo.Data 
	fmt.Println(logInfo.Timestr, logInfo.Module, logInfo.Data)
	 loc, _ := time.LoadLocation("Local")
	 timeLayout := "2006-01-02 15:04:05"
	// timestamp, _ := time.ParseInLocation("2006-01-02 03:04:05", logInfo.Timestr, time.Local)
	timestamp, _ := time.ParseInLocation(timeLayout, logInfo.Timestr, loc)
	fmt.Println(timestamp)
	fmt.Println(time.Now().Format("2006-01-02 15:04:05"))
	sendLogInput := &SendLogInput{
		RepoName : repoName,
		OmitInvalidLog : false,
		Logs : Logs{
			Log{
				"timestamp": timestamp, //timestamp.UTC(),
				"id": logInfo.Id,
				"module" : logInfo.Module,
				"data" : logInfo.Data,
			},
		},
	}
	fmt.Println(sendLogInput)
	sendOutput, err := logdbClient.SendLog(sendLogInput)
	if err != nil {
		fmt.Println("SendLog failede.\n")
		fmt.Println(err)
		return err
	}
	fmt.Println(sendOutput)

	return
}

func DeleteRepo() error {
	err := logdbClient.DeleteRepo(&DeleteRepoInput{RepoName:repoName})
	if err != nil {
		fmt.Println(err)
		return err
	}
	return nil
}

func GetMqttBeat()(output *QueryLogOutput, err error) {
	queryInput := &QueryLogInput{
		RepoName: repoName,
		Query: "module:MQHB",
		Sort: "data:desc",
		From: 0,
		Size: 100,
	}
	queryOutput, err := logdbClient.QueryLog(queryInput)
	if err != nil {
		fmt.Println(err)
		return nil, err
	}
	fmt.Println(queryOutput)
	return queryOutput, nil
}

func GetMqttBeatHistogram()(output *QueryHistogramLogOutput, err error) {
	
	startTimeStr := "2018-09-09 18:32:45"
	endTimeStr := "2018-09-09 18:32:50"
	
	startTime, _ := time.Parse("2006-01-02 15:04:05", startTimeStr)
	endTime, _ := time.Parse("2006-01-02 15:04:05", endTimeStr)
	startT := startTime.Unix()
	endT := endTime.Unix()
	histogramInput := &QueryHistogramLogInput{
		RepoName: repoName,
		Query: "module:MQHB",
		From: startT,
		To: endT,
		Field: "data",
	}
	histogramOutput, err := logdbClient.QueryHistogramLog(histogramInput)
	if err != nil {
		fmt.Println(err)
		return nil, err
	}
	fmt.Println(histogramOutput)
	return histogramOutput, nil
}

func GetOnRate() {
	//search MQHB > 0
	//search MQHB <= 0
	//rate = > 0 / all
}

func GetRtmpStatus() {

}

func GetRtmpVideoRate() {

}

func DAInit() {
	var err error
	
	gLogger = NewDefaultLogger()
	logdbCfg = NewConfig().
	WithEndpoint(endpoint).
	WithAccessKeySecretKey(akLink, skLink).
	WithLogger(gLogger).
	WithLoggerLevel(LogDebug)

	logdbClient, err = New(logdbCfg)
	if err  != nil {
		gLogger.Error("New logdb client failed. err:%v", err)
	}
	
}