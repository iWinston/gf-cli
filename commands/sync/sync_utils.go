package sync

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"

	"github.com/alecthomas/template"
	"github.com/gogf/gf/os/gcmd"
	"github.com/gogf/gf/os/gfile"
	"github.com/iWinston/gf-cli/commands/sync/apifox"
	"github.com/iWinston/gf-cli/library/mlog"
	"github.com/iWinston/gf-cli/library/utils"
)

func load() *apifox.Apifox {
	url := "https://apifox-api.apipark.cn/api/v1/export-data?__xAuthorization=Bearer%20eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJpZCI6MzY3NjgzLCJ0cyI6IjkwZjdhZTliNzU3NGU5ODgiLCJpYXQiOjE2MTU3OTIxNzM3NzN9.sRlRlPt0A9qqcUZoo14L0LVKWwqXlaeIgRbAWLhPrCc&__xProjectId="
	projectId := gcmd.GetArg(3)
	JsonParse := NewJsonStruct()

	//下面使用的是相对路径，config.json文件和main.go文件处于同一目录下
	v := apifox.Apifox{}
	JsonParse.Load(url+projectId, &v)

	return &v
}

type JsonStruct struct {
}

func NewJsonStruct() *JsonStruct {
	return &JsonStruct{}
}

func (jst *JsonStruct) Load(url string, v interface{}) (err error) {
	if strings.HasPrefix(url, "http") {
		err = getJSONByHttp(url, v)
		if err != nil {
			return
		}
	} else {
		//ReadFile函数会读取文件的全部内容，并将结果以[]byte类型返回
		if data, err := ioutil.ReadFile(url); err != nil {
			return err
		} else {
			//读取的数据为json格式，需要进行解码
			return json.Unmarshal(data, v)
		}

	}
	return
}

func syncFileForce(folder string, fileName string, temp string, data interface{}) {
	fileName = utils.SnakeString(fileName)
	if err := gfile.Mkdir(folder); err != nil {
		mlog.Fatalf("mkdir for generating path '%s' failed: %v", folder, err)
	}
	path := gfile.Join(folder, fileName)
	t, err := template.New(fileName).Parse(utils.Header + temp)
	if err != nil {
		mlog.Fatalf("template to '%s' failed: %v", path, err)
	}
	f, _ := gfile.Create(path)
	parseParam := &parseParam{
		ProjectName: projectName,
		Backtick:    "`",
		Data:        data,
	}
	if err := t.Execute(f, parseParam); err != nil {
		mlog.Fatalf("writing content to '%s' failed: %v", path, err)
	} else {
		mlog.Print("generated:", path)
	}
}

func syncFile(folder string, fileName string, temp string, data interface{}) {
	fileName = utils.SnakeString(fileName)
	if err := gfile.Mkdir(folder); err != nil {
		mlog.Fatalf("mkdir for generating path '%s' failed: %v", folder, err)
	}
	path := gfile.Join(folder, fileName)
	if gcmd.ContainsOpt("f") || !gfile.Exists(path) {
		t, err := template.New(fileName).Parse(utils.OnlyOnceHeader + temp)
		if err != nil {
			mlog.Fatalf("template to '%s' failed: %v", path, err)
		}
		f, _ := gfile.Create(path)
		parseParam := &parseParam{
			ProjectName: projectName,
			Backtick:    "`",
			Data:        data,
		}
		if err := t.Execute(f, parseParam); err != nil {
			mlog.Fatalf("writing content to '%s' failed: %v", path, err)
		} else {
			mlog.Print("generated:", path)
		}
	}
}

type parseParam struct {
	ProjectName string
	Backtick    string
	Data        interface{}
}

func getRef(refId string) string {
	refInfo := refs[refId]
	if refInfo.Name == "time.Date" {
		return "time.Time"
	}
	return refInfo.Name
}

func getRefs(schemaCollection *[]apifox.SchemaItem) map[string]DefineInfo {
	defineInfoMap := make(map[string]DefineInfo)
	for _, schemaSystem := range *schemaCollection {
		getMapFromItem(&defineInfoMap, &schemaSystem, schemaSystem.Name)
	}
	return defineInfoMap
}

func getMapFromItem(defineInfoMap *map[string]DefineInfo, schema *apifox.SchemaItem, systemName string) {
	if len(schema.Items) == 0 {
		return
	}
	for _, item := range schema.Items {
		(*defineInfoMap)[item.Id] = getDefine(&item, systemName)
		getMapFromItem(defineInfoMap, &item, systemName)
	}
}

func joinNotEmpty(strArr []string, sep string) string {
	newArr := []string{}
	for _, str := range strArr {
		if str != "" {
			newArr = append(newArr, str)
		}
	}
	return strings.Join(newArr, sep)
}

// getJSON fetches the contents of the given URL
// and decodes it as JSON into the given result,
// which should be a pointer to the expected data.
func getJSONByHttp(url string, result interface{}) error {
	resp, err := http.Get(url)
	if err != nil {
		return fmt.Errorf("cannot fetch URL %q: %v", url, err)
	}
	defer resp.Body.Close()
	if resp.StatusCode != http.StatusOK {
		return fmt.Errorf("unexpected http GET status: %s", resp.Status)
	}
	// We could check the resulting content type
	// here if desired.
	err = json.NewDecoder(resp.Body).Decode(result)
	if err != nil {
		return fmt.Errorf("cannot decode JSON: %v", err)
	}
	return nil
}
