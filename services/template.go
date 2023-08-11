package services

import (
	"text/template"
)

func MainEntry() *template.Template {
	tmplStr := `package main

import (
	"{{ .Name }}/cmds"
)

func main() {
	cmds.Execute()
}
`
	tmpl := template.Must(template.New("codeTemplate").Parse(tmplStr))
	return tmpl
}

func ApisRouter() *template.Template {
	tmplStr := `package apis

import (
	"fmt"
	"log"

	"{{ .Name }}/configs"
	"{{ .Name }}/dbs"
	"{{ .Name }}/services/common/logs"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

// Run 执行服务
func Run() error {
	// 测试数据库连接
	err := dbs.DBPing()
	if err != nil {
		return err
	}
	// 读取服务器连接配置
	apiConfig, err := configs.APIConfig()
	if err != nil {
		return err
	}
	listenAddr := fmt.Sprintf("%s:%s", apiConfig["host"], apiConfig["port"])
	log.Println("Listen: " + listenAddr)
	// 配置日志
	logger, err := logs.GinLogger("access.log")
	if err != nil {
		return err
	}

	if configs.Debug {
		gin.SetMode(gin.DebugMode)
	} else {
		gin.SetMode(gin.ReleaseMode)
	}

	router := gin.New()
	router.SetTrustedProxies([]string{"127.0.0.1"})

	// 跨域
	config := cors.DefaultConfig()
	config.AllowOrigins = []string{"*"}
	config.AllowMethods = []string{"GET", "POST", "PUT", "DELETE", "OPTION"}
	config.AllowHeaders = []string{"Origin", "X-Requested-With", "content-Type", "Accept", "Authorization"}
	router.Use(cors.New(config))

	router.Use(gin.Recovery())
	router.Use(logger)

	v1 := router.Group("/api/v1")
	{
		v1.GET("/path/url", nil)
	}
	return router.Run(listenAddr)
}
`
	tmpl := template.Must(template.New("codeTemplate").Parse(tmplStr))
	return tmpl
}

func ApisCommonHandler() *template.Template {
	tmplStr := `package common

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// JSONHandler 返回 JSON 数据
func JSONHandler(c *gin.Context, status int, message string, data any) {
	c.JSON(http.StatusOK, gin.H{
		"status":  status,
		"message": message,
		"data":    data,
	})
}

// AuthHandler 返回认证数据
func AuthHandler(c *gin.Context, status int, message string) {
	c.JSON(http.StatusUnauthorized, gin.H{
		"status":  status,
		"message": message,
	})
}

// StringHandler 返回字符串数据
func StringHandler(c *gin.Context, status int, data string) {
	c.String(status, data)
}
`
	tmpl := template.Must(template.New("codeTemplate").Parse(tmplStr))
	return tmpl
}

func CmdsRoot() *template.Template {
	tmplStr := `package cmds

import (
	"fmt"
	"os"

	"{{ .Name }}/cmds/api"
	"{{ .Name }}/cmds/db"
	"{{ .Name }}/utils"

	"github.com/spf13/cobra"
)

var (
	rootCmd = &cobra.Command{
		Use:     "{{ .Name }}",
		Short:   "some desc",
		Version: utils.Version,
		Run: func(cmd *cobra.Command, args []string) {
			cmd.Help()
		},
	}
)

func Execute() {
	rootCmd.CompletionOptions.DisableDefaultCmd = true
	rootCmd.DisableFlagsInUseLine = true
	rootCmd.AddCommand(
		api.APICmd,
		db.DBCmd,
	)
	if err := rootCmd.Execute(); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}
`
	tmpl := template.Must(template.New("codeTemplate").Parse(tmplStr))
	return tmpl
}

func CmdsApiApi() *template.Template {
	tmplStr := `package api

import (
	"log"

	"{{ .Name }}/apis"

	"github.com/spf13/cobra"
)

var (
	APICmd = &cobra.Command{
		Use:   "api",
		Short: "Run Web API",
		Run: func(cmd *cobra.Command, args []string) {
			err := apis.Run()
			if err != nil {
				log.Fatal(err)
			}
		},
	}
)

func init() {
	APICmd.DisableFlagsInUseLine = true
}
`
	tmpl := template.Must(template.New("codeTemplate").Parse(tmplStr))
	return tmpl
}

func CmdsDbDb() *template.Template {
	tmplStr := `package db

import (
	"log"

	"{{ .Name }}/dbs"

	"github.com/spf13/cobra"
)

var (
	inviteCode string
	DBCmd      = &cobra.Command{
		Use:   "db",
		Short: "Database manager",
		Run: func(cmd *cobra.Command, args []string) {
			cmd.Help()
		},
	}
	initCmd = &cobra.Command{
		Use:   "init",
		Short: "Init tables",
		Run: func(cmd *cobra.Command, args []string) {
			err := dbs.DBCreateTable()
			if err != nil {
				log.Fatal(err)
			}
			err = dbs.DBComment()
			if err != nil {
				log.Fatal(err)
			}
		},
	}
)

func init() {
	DBCmd.AddCommand(initCmd)
	DBCmd.DisableFlagsInUseLine = true
}
`
	tmpl := template.Must(template.New("codeTemplate").Parse(tmplStr))
	return tmpl
}

func ConfigsConfig() *template.Template {
	tmplStr := `package configs

import (
	"path/filepath"

	"{{ .Name }}/utils"
)

const configPath = "configs"

func readMapData(cfgName string) (map[string]any, error) {
	mainRoot, err := utils.FileSuite.RootPath(configPath)
	if err != nil {
		return nil, err
	}

	cfgPath := filepath.Join(mainRoot, cfgName)
	err = utils.FileSuite.IsExist(cfgPath)
	if err != nil {
		return nil, err
	}
	rawData, err := utils.FileSuite.ReadFile(cfgPath)
	if err != nil {
		return nil, err
	}
	config, err := utils.DataSuite.RawMap2Map(rawData)
	return config, nil
}

// DatabaseConfig 数据库连接配置
func DatabaseConfig() (map[string]any, error) {
	return readMapData("database.json")
}

// APIConfig 接口连接配置
func APIConfig() (map[string]any, error) {
	return readMapData("config.json")
}
`
	tmpl := template.Must(template.New("codeTemplate").Parse(tmplStr))
	return tmpl
}

func ConfigsEnv() *template.Template {
	tmplStr := `package configs

const Debug bool = false
`
	tmpl := template.Must(template.New("codeTemplate").Parse(tmplStr))
	return tmpl
}

func ConfigsDatabaseJson() *template.Template {
	tmplStr := `{
	"host": "127.0.0.1",
	"port": 5432,
	"username": "",
	"password": "",
	"dbname": ""
}
`
	tmpl := template.Must(template.New("codeTemplate").Parse(tmplStr))
	return tmpl
}

func ConfigsConfigJSON() *template.Template {
	tmplStr := `{
	"host": "127.0.0.1",
	"port": "8373"
}
`
	tmpl := template.Must(template.New("codeTemplate").Parse(tmplStr))
	return tmpl
}

func DbDbsDbs() *template.Template {
	tmplStr := `package dbs

import (
	"log"

	"{{ .Name }}/configs"

	"github.com/qmaru/qdb/postgresql"
)

const (
	SomeTable        string = "some_table"
)

var Psql *postgresql.PostgreSQL

func init() {
	cfg, err := configs.DatabaseConfig()
	if err != nil {
		log.Fatal(err)
	}

	host := cfg["host"].(string)
	port := int(cfg["port"].(float64))
	username := cfg["username"].(string)
	password := cfg["password"].(string)
	dbname := cfg["dbname"].(string)

	Psql = postgresql.New(host, port, username, password, dbname)
}

// DBPing 测试数据库连接
func DBPing() error {
	return Psql.Ping()
}

// DBCreateTable 创建数据表
func DBCreateTable() error {
	tables := []any{}
	return Psql.CreateTable(tables)
}

// DBComment 注释数据表的字段
func DBComment() error {
	tables := []any{}
	return Psql.Comment(tables)
}

// DBCreateIndex 给指定表的字段创建索引
func DBCreateIndex() error {
	tables := []any{}
	return Psql.CreateIndex(tables)
}
`
	tmpl := template.Must(template.New("codeTemplate").Parse(tmplStr))
	return tmpl
}

func DbDbsModelsModel() *template.Template {
	tmplStr := `package models

import (
	"time"
)

// CommonModel 公用结构
type CommonModel struct {
	ID        uint64    ` + "`json:\"id\" db:\"serial;PRIMARY KEY\" comment:\"ID\"`" + `
	CreatedAt time.Time ` + "`json:\"created_at\" db:\"timestamp;DEFAULT NULL\" comment:\"创建时间\"`" + `
	UpdatedAt time.Time ` + "`json:\"updated_at\" db:\"timestamp;DEFAULT NULL\" comment:\"更新时间\"`" + `
	DeletedAt time.Time ` + "`json:\"deleted_at\" db:\"timestamp;DEFAULT NULL\" comment:\"删除时间\"`" + `
	State     bool      ` + "`json:\"state\" db:\"boolean;DEFAULT true\" comment:\"状态\"`" + `
	Remark    string    ` + "`json:\"remark\" db:\"varchar;DEFAULT ''\" comment:\"备注\"`" + `
}
`
	tmpl := template.Must(template.New("codeTemplate").Parse(tmplStr))
	return tmpl
}

func ServiceServicesCommonLogsCommon() *template.Template {
	tmplStr := `package logs

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"
	"time"

	"{{ .Name }}/utils"
	"{{ .Name }}/configs"

	"github.com/sirupsen/logrus"
)

type myFormat struct{}

func (f *myFormat) Format(entry *logrus.Entry) ([]byte, error) {
	timestamp := time.Now().Local().Format("2006/01/02 15:04:05")
	msg := fmt.Sprintf("%s [%s] %s\n", timestamp, strings.ToUpper(entry.Level.String()), entry.Message)
	return []byte(msg), nil
}

// Logger 定义日志格式
func Logger(logName string) (*logrus.Logger, error) {
	logger := logrus.New()
	// 输出到文件
	if configs.Debug {
		logger.Out = os.Stdout
		logger.SetLevel(logrus.DebugLevel)
	} else {
		logPath, err := utils.FileSuite.RootPath("logs")
		if err != nil {
			return nil, err
		}
		logpath, err := utils.FileSuite.Mkdir(logPath)
		if err != nil {
			return nil, err
		}
		accessPath := filepath.Join(logpath, logName)
		accessFile, err := os.OpenFile(accessPath, os.O_WRONLY|os.O_APPEND|os.O_CREATE, 0644)
		if err != nil {
			return nil, err
		}
		logger.Out = accessFile
		logger.SetLevel(logrus.InfoLevel)
	}
	myformat := new(myFormat)
	logger.SetFormatter(myformat)
	return logger, nil
}
`
	tmpl := template.Must(template.New("codeTemplate").Parse(tmplStr))
	return tmpl
}

func ServiceServicesCommonLogsLogs() *template.Template {
	tmplStr := `package logs

import (
	"time"

	"github.com/gin-gonic/gin"
)

func GinLogger(logName string) (gin.HandlerFunc, error) {
	logger, err := Logger(logName)
	if err != nil {
		return nil, err
	}
	return func(c *gin.Context) {
		// 开始时间
		startTime := time.Now()
		// 处理请求
		c.Next()
		// 结束时间
		endTime := time.Now()
		// 执行时间
		latencyTime := endTime.Sub(startTime)
		// 请求方式
		reqMethod := c.Request.Method
		// 请求路由
		reqURI := c.Request.RequestURI
		// 状态码
		statusCode := c.Writer.Status()
		// 请求IP
		clientIP := c.ClientIP()
		// 日志格式
		logger.Infof("- %s %d %s %s %s",
			reqMethod,
			statusCode,
			clientIP,
			reqURI,
			latencyTime,
		)
	}, nil
}
`
	tmpl := template.Must(template.New("codeTemplate").Parse(tmplStr))
	return tmpl
}

func UtilsUtilsMinireq() *template.Template {
	tmplStr := `package utils

import (
	minireq "github.com/qmaru/minireq/v2"
)

// UserAgent 全局 UA
var UserAgent = "{{ .Name }}/" + DateVer

var Minireq *minireq.HttpClient

// MiniHeaders Headers
type MiniHeaders = minireq.Headers

// MiniParams Params
type MiniParams = minireq.Params

// MiniJSONData application/json
type MiniJSONData = minireq.JSONData

// MiniFormData multipart/form-data
type MiniFormData = minireq.FormData

// MiniFormKV application/x-www-from-urlencoded
type MiniFormKV = minireq.FormKV

// MiniAuth HTTP Basic Auth
type MiniAuth = minireq.Auth

// MiniResponse
type MiniResponse = minireq.MiniResponse

func init() {
	Minireq = minireq.NewClient()
}
`
	tmpl := template.Must(template.New("codeTemplate").Parse(tmplStr))
	return tmpl
}

func UtilsUtilsTools() *template.Template {
	tmplStr := `package utils

import (
	"github.com/qmaru/minitools"
)

// DataSuite 初始化
var DataSuite *minitools.DataSuiteBasic

// FileSuite 初始化
var FileSuite *minitools.FileSuiteBasic

// TimeSuite 初始化
var TimeSuite *minitools.TimeSuiteBasic

func init() {
	DataSuite = minitools.DataSuite()
	FileSuite = minitools.FileSuite()
	TimeSuite = minitools.TimeSuite()
}
`
	tmpl := template.Must(template.New("codeTemplate").Parse(tmplStr))
	return tmpl
}

func UtilsUtilsVersion() *template.Template {
	tmplStr := `package utils

import (
	"fmt"
)

const (
	Major     string = "2"
	Minor     string = "0"
	DateVer   string = "COMMIT_DATE"
	CommitVer string = "COMMIT_VERSION"
	GoVer     string = "COMMIT_GOVER"
)

var Version string = fmt.Sprintf("%s.%s-%s (git-%s) (%s)", Major, Minor, DateVer, CommitVer, GoVer)
`
	tmpl := template.Must(template.New("codeTemplate").Parse(tmplStr))
	return tmpl
}
