package main
import (
//	"time"
	"log"
	. "github.com/qiniu/pandora-go-sdk/base"
	"github.com/qiniu/pandora-go-sdk/base/config"
	. "github.com/qiniu/pandora-go-sdk/logdb"
)

var (
	cfg *config.Config
	client LogdbAPI
	region = "<Region>"
	endpoint = config.DefaultLogDBEndpoint
	ak = "your ak"
	sk = "your sk"
	logger Logger
	defaultRepoSchema []RepoSchemaEntry
)

func demo_init() {
    var err error

    logger = NewDefaultLogger()
    cfg = NewConfig().
	WithEndpoint(endpoint).
	WithAccessKeySecretKey(ak, sk).
	WithLogger(logger).
	WithLoggerLevel(LogDebug)

	client, err = New(cfg)
	if err != nil {
		logger.Error("new logdb client failed, err:%v", err)
	}
	defaultRepoSchema = []RepoSchemaEntry{
		RepoSchemaEntry{
			Key: "f1",
			ValueType: "date",
		},
		RepoSchemaEntry{
			Key: "f2",
			ValueType: "string",
		},
		RepoSchemaEntry{
			Key: "f3",
			ValueType: "string",
		},
	}
}

func AddLog(client LogdbAPI, repoName string)(err error) {
	return 
}

func DeleteRepo(client LogdbAPI, repoName string)(err error) {
	
	err = client.DeleteRepo(&DeleteRepoInput{RepoName: repoName})
	if err != nil {
		logger.Error(err)
		return
	}
	return
}

func main() {
	demo_init()

	var repoName = "mqttrepo"
/*
	var mqttRepoSchema = []RepoSchemaEntry{
		RepoSchemaEntry{
			Key: "timestamp",
			ValueType: "date",
		},
		RepoSchemaEntry{
			Key: "module",
			ValueType: "string",
		},
		RepoSchemaEntry{
			Key: "data",
			ValueType: "long",
		},
	}
	createInput := &CreateRepoInput{
		RepoName: repoName,
		Region: "nb",
		Schema: mqttRepoSchema,
		Retention: "2d",
	}
	err := client.CreateRepo(createInput)
	if err != nil {
		logger.Error(err)
		return
	}
*/
	getRepoOutput, err := client.GetRepo(&GetRepoInput{RepoName: repoName})
	if err != nil {
		logger.Error(err)
		return
	}
	logger.Info(getRepoOutput)
	//add log
/*
	sendLogInput := &SendLogInput{
		RepoName: repoName,
		OmitInvalidLog: false,
		Logs: Logs{
			Log{
			"timestamp": time.Now().UTC().Format(time.RFC3339),
			"module": "RTMPHB",
			"data": 200,
			},
			},
	}

	sendOutput, err := client.SendLog(sendLogInput)
	if err != nil {
		logger.Error(err)
		return
	}
	logger.Info(sendOutput)
*/
	logger.Errorf("==========")
	//search a log
	queryInput :=  &QueryLogInput{
		RepoName:  "mqttrepo",
		Query: "module:RTMPHB",
		Sort: "data:desc",
		From: 0,
		Size: 10,
	}
	queryOutput, err := client.QueryLog(queryInput)
	if err != nil {
		logger.Error(err)
		return
	}
	if queryOutput.Total != 0 {
		log.Printf("Total:%d", queryOutput.Total)
	}
	//logger.Info(queryOutput)
	logger.Errorf("==========")
}




















