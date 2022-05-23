package api

import (
	"encoding/json"
	"errors"
	"github.com/parnurzeal/gorequest"
	"os"
	"strings"
)

type ApiClient struct {
	ip     string
	port   string
	debug  bool
	url    string
	client *gorequest.SuperAgent
	baseName string
	basePasswd string
}

func NewApi(ip, port string, debug bool) *ApiClient {
	api := &ApiClient{
		ip:     ip,
		port:   port,
		debug:  debug,
		url:    "http://" + ip + ":" + port + "/",
		client: gorequest.New().SetDebug(debug),
	}

	api.SetBasicAuth("", "")
	return api
}

func NewApiWithserver(server string) *ApiClient {
	api := &ApiClient{
		url:    "http://" + server + "/",
		client: gorequest.New().SetDebug(false),
	}
	api.SetBasicAuth("", "")
	return api
}


func (api *ApiClient)SetBasicAuth(name,passwd string)  {
	if name == ""||passwd == "" {
		authStr := os.Getenv("AUTH")
		if authStr != "" {
			authStrlist := strings.Split(authStr, "-")
			api.baseName,api.basePasswd = authStrlist[0], authStrlist[1]
		} else {
			api.baseName,api.basePasswd = "dog", "admin"
		}
	}else {
		api.baseName,api.basePasswd = name,passwd
	}

	api.client.SetBasicAuth(api.baseName,api.basePasswd)
}

//在dog上设置一个命令
func (api *ApiClient) SetCommand(command SetCommand) (ResponseFront, error) {
	rsp := ResponseFront{}
	_, _, errs := api.client.Post(api.url + "set_command").SendStruct(command).EndStruct(&rsp)
	if len(errs) != 0 {
		return ResponseFront{}, errs[0]
	}
	return rsp, nil
}
//在dog上执行一个命令
func (api *ApiClient) ExecuteCommand(command Execute) (ResponseFront, error) {
	rsp := ResponseFront{}
	_, _, errs := api.client.Post(api.url + "execute").SendStruct(command).EndStruct(&rsp)
	if len(errs) != 0 {
		return ResponseFront{}, errs[0]
	}
	return rsp, nil
}

//在dog上停止一个命令
func (api *ApiClient) StopCommand(command Execute) (ResponseFront, error) {
	rsp := ResponseFront{}
	_, _, errs := api.client.Post(api.url + "stop").SendStruct(command).EndStruct(&rsp)
	if len(errs) != 0 {
		return ResponseFront{}, errs[0]
	}
	return rsp, nil
}

//在dog上获取命令列表
func (api *ApiClient) GetCommand() ([]CommandAndStatus, error) {
	rsp := ResponseFront{}
	_, _, errs := api.client.Get(api.url + "get_command").EndStruct(&rsp)
	if len(errs) != 0 {
		return []CommandAndStatus{}, errs[0]
	}
	if rsp.Code != "ok" {
		return []CommandAndStatus{}, errors.New(rsp.Msg)
	}
	if rsp.Data == nil {
		return []CommandAndStatus{}, nil
	}
	var commandList []CommandAndStatus
	data, err := json.Marshal(rsp.Data)
	if err != nil {
		return []CommandAndStatus{}, err
	}
	err = json.Unmarshal(data, &commandList)
	if err != nil {
		return []CommandAndStatus{}, err
	}
	return commandList, nil
}

//让dag下载某个命令的程序
func (api *ApiClient) Download(command SetCommand) (ResponseFront, error) {
	rsp := ResponseFront{}
	_, _, errs := api.client.Post(api.url + "download").SendStruct(command).EndStruct(&rsp)
	if len(errs) != 0 {
		return ResponseFront{}, errs[0]
	}
	return rsp, nil
}

//dog心跳检测
func (api *ApiClient) Ping() (ResponseFront, error) {
	rsp := ResponseFront{}
	_, _, errs := api.client.Get(api.url + "ping").EndStruct(&rsp)
	if len(errs) != 0 {
		return ResponseFront{}, errs[0]
	}
	return rsp, nil
}

func DataUnmarshal(d, s interface{}) error {
	data, err := json.Marshal(d)
	if err != nil {
		return err
	}
	return json.Unmarshal(data, s)
}

//dog运行程序的状态
func (api *ApiClient) Status(nickName string) ([]RunningState, error) {
	rsp := ResponseFront{}
	_, _, errs := api.client.Get(api.url + "status").Query("nick_name=" + nickName).EndStruct(&rsp)
	if len(errs) != 0 {
		return nil, errs[0]
	}
	if rsp.Code != "ok" {
		return nil, errors.New(rsp.Msg)
	}
	var runningState []RunningState
	err := DataUnmarshal(rsp.Data, &runningState)
	if err != nil {
		return nil, err
	}
	return runningState, nil
}
//运行dog的用户
func (api *ApiClient) UserCurrent() (string, error) {
	rsp := ResponseFront{}
	_, _, errs := api.client.Get(api.url + "user_current").EndStruct(&rsp)
	if len(errs) != 0 {
		return "", errs[0]
	}
	if rsp.Code != "ok" {
		return "", errors.New(rsp.Msg)
	}
	return rsp.Data.(string), nil
}

//运行dog的版本，git的提交信息
func (api *ApiClient) Version() (string, error) {
	rsp := ResponseFront{}
	_, _, errs := api.client.Get(api.url + "version").EndStruct(&rsp)
	if len(errs) != 0 {
		return "", errs[0]
	}
	if rsp.Code != "ok" {
		return "", errors.New(rsp.Msg)
	}
	return rsp.Data.(string), nil
}
//获取所有的webhook
func (api *ApiClient) GetAllWebhook() ([]WebhookRepresentation, error) {
	rsp := ResponseFront{}
	_, _, errs := api.client.Get(api.url + "get_all_webhook").EndStruct(&rsp)
	if len(errs) != 0 {
		return []WebhookRepresentation{}, errs[0]
	}
	a, _ := json.Marshal(rsp.Data)
	var webhooklist []WebhookRepresentation
	err := json.Unmarshal(a, &webhooklist)
	if err != nil {
		return webhooklist, err
	}
	return webhooklist, nil
}

//获取某一个命令的webhook
func (api *ApiClient) GetWebhook(nickName string) ([]WebhookRes, error) {
	rsp := ResponseFront{}
	_, _, errs := api.client.Get(api.url + "get_webhook/" + nickName).EndStruct(&rsp)
	if len(errs) != 0 {
		return []WebhookRes{}, errs[0]
	}

	a, _ := json.Marshal(rsp.Data)
	var webhooklist []WebhookRes
	err := json.Unmarshal(a, &webhooklist)
	if err != nil {
		return webhooklist, err
	}
	return webhooklist, nil
}

//设置的webhook
func (api *ApiClient) SetWebhook(command Webhook) (ResponseFront, error) {
	rsp := ResponseFront{}
	_, _, errs := api.client.Post(api.url + "set_webhook").SendStruct(command).EndStruct(&rsp)
	if len(errs) != 0 {
		return ResponseFront{}, errs[0]
	}
	return rsp, nil
}
//删除webhook
func (api *ApiClient) DelWebhook(command DelWebhook) (ResponseFront, error) {
	rsp := ResponseFront{}
	_, _, errs := api.client.Post(api.url + "del_webhook").SendStruct(command).EndStruct(&rsp)
	if len(errs) != 0 {
		return ResponseFront{}, errs[0]
	}
	return rsp, nil
}

//获取保活列表
func (api *ApiClient) GetKeepLive() ([]KeepLive, error) {
	rsp := ResponseFront{}
	_, _, errs := api.client.Get(api.url + "get_keeplive").EndStruct(&rsp)
	if len(errs) != 0 {
		return []KeepLive{}, errs[0]
	}

	a, _ := json.Marshal(rsp.Data)
	var KeepLivelist []KeepLive
	err := json.Unmarshal(a, &KeepLivelist)
	if err != nil {
		return KeepLivelist, err
	}
	return KeepLivelist, nil
}

//设置保活
func (api *ApiClient) KeepLive(command KeepLive) (ResponseFront, error) {
	rsp := ResponseFront{}
	_, _, errs := api.client.Post(api.url + "keep_live").SendStruct(command).EndStruct(&rsp)
	if len(errs) != 0 {
		return ResponseFront{}, errs[0]
	}
	return rsp, nil
}

//上传文件到dog
func (api *ApiClient) Upload(command Upload, filePath string) (ResponseFront, error) {
	rsp := ResponseFront{}
	_, _, errs := api.client.Post(api.url+"upload").Type("multipart").Send(command).SendFile(filePath, "", "command_file").EndStruct(&rsp)
	if len(errs) != 0 {
		return ResponseFront{}, errs[0]
	}
	return rsp, nil
}
//dog的iplist，节点发现下获取的iplist
func (api *ApiClient) IpList() (IpList2, error) {
	rsp := ResponseFront{}
	_, _, errs := api.client.Get(api.url + "ip_list").EndStruct(&rsp)
	if len(errs) != 0 {
		return IpList2{}, errs[0]
	}
	a, _ := json.Marshal(rsp.Data)
	var ipList IpList2
	err := json.Unmarshal(a, &ipList)
	if err != nil {
		return ipList, err
	}
	return ipList, nil
}

//操作日志
func (api *ApiClient) OperationLog(nickName string) (string, error) {
	rsp := ResponseFront{}
	_, _, errs := api.client.Get(api.url + "operation_log/" + nickName).EndStruct(&rsp)
	if len(errs) != 0 {
		return "", errs[0]
	}

	return rsp.Data.(string), nil
}

//dog上运行的插件及其端口
func (api *ApiClient) Plugin_list() ([]Plugin, error) {
	rsp := ResponseFront{}
	_, _, errs := api.client.Get(api.url + "plugin_list").EndStruct(&rsp)
	if len(errs) != 0 {
		return []Plugin{}, errs[0]

	}
	a, _ := json.Marshal(rsp.Data)
	var List []Plugin
	err := json.Unmarshal(a, &List)
	if err != nil {
		return List, err
	}
	return List, nil

}

//设置当前dog上发现的节点的ip昵称
func (api *ApiClient) SetIpNickname(ipStruct IpStruct) (ResponseFront, error) {
	rsp := ResponseFront{}
	_, _, errs := api.client.Post(api.url + "set_ip_nickname").SendStruct(ipStruct).EndStruct(&rsp)
	if len(errs) != 0 {
		return ResponseFront{}, errs[0]
	}
	return rsp, nil
}

//删除节点发现的节点列表中某一个不在线的节点
func (api *ApiClient) DeleteServer(ipStruct IpStruct) (ResponseFront, error) {
	rsp := ResponseFront{}
	_, _, errs := api.client.Post(api.url + "delete_server").SendStruct(ipStruct).EndStruct(&rsp)
	if len(errs) != 0 {
		return ResponseFront{}, errs[0]
	}
	return rsp, nil
}

//让同集群的某一个dog执行开始或者结束
func (api *ApiClient) Relay(relay Relay) (ResponseFront, error) {
	rsp := ResponseFront{}
	_, _, errs := api.client.Post(api.url + "relay").SendStruct(relay).EndStruct(&rsp)
	if len(errs) != 0 {
		return ResponseFront{}, errs[0]
	}
	return rsp, nil
}
//通过文件导入命令
func (api *ApiClient) Import(filePath string) (ResponseFront, error) {
	rsp := ResponseFront{}
	_, _, errs := api.client.Post(api.url+"import").Type("multipart").SendFile(filePath, "", "export_file").EndStruct(&rsp)
	if len(errs) != 0 {
		return ResponseFront{}, errs[0]
	}
	return rsp, nil
}
//导出dog上的命令
func (api *ApiClient) Export() (string, error) {
	var list string
	_, _, errs := api.client.Get(api.url + "export").EndBytes(func(response gorequest.Response, body []byte, errs []error) {
		list = string(body)
	})
	if len(errs) != 0 {
		return list, errs[0]
	}
	return list, nil
}

