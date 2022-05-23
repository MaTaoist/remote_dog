package api
// SetCommand 设置命令的入参
type SetCommand struct {
	Url         string   `json:"url"`          // 下载地址
	Params      []string `json:"params"`       // 执行参数
	Env         []string `json:"env"`          // 程序的环境变量
	NickName    string   `json:"nick_name"`    // 程序昵称，别名
	ProgramName string   `json:"program_name"` // 程序名，启动用的名称
	Md5         string   `json:"md5"`          //程序的MD5
	Delete      bool     `json:"delete"`       //是否为删除
}

type Execute struct {
	NickName    string `json:"nick_name"`    // 程序昵称，别名
	ProgramName string `json:"program_name"` // 程序名，启动用的名称
}

// CommandAndStatus 查询命令详情的返回
type CommandAndStatus struct {
	NickName    string   `json:"nick_name"`    // 程序昵称，别名
	ProgramName string   `json:"program_name"` // 程序名，启动用的名称
	Params      []string `json:"params"`       // 执行参数
	Env         []string `json:"env"`          // 程序的环境变量
	Status      int      `json:"status"`       //程序状态
	Md5         string   `json:"md5"`          //程序的MD5 用户填写的用于验证的
	RealMd5     string   `json:"real_md5"`     //实际算出来的
	Url         string   `json:"url"`          // 下载地址
	KeepLive    int      `json:"keep_live"`
}

type RunningState struct {
	Status      int    `json:"status"`   //1开2关
	LogPath     string `json:"log_path"` //程序log目录
	Md5         string `json:"md5"`      //程序的MD5
	NickName    string `json:"nick_name"`
	ProgramName string `json:"program_name"`
}


type WebhookRepresentation struct {
	Types         int      `json:"types"` // 1 钉钉 2 其他自定义（推送为json信息）
	Url           string   `json:"url"`   //若为钉钉，此处填写token
	Secret        string   `json:"secret"`
	NickName      []string `json:"nick_name"`
	EventTypeList []int    `json:"event_type_list"` // 上线，下线，设置新命令，更新命令，删除命令，update下载了文件
	List          []Representation
}
type Representation struct {
	Wenhook int   //wenhook类型
	Key     int64 //wenhook key
}


type WebhookRes struct {
	Types         int      `json:"types"` // 1 钉钉 2 其他自定义（推送为json信息）
	Url           string   `json:"url"`   //若为钉钉，此处填写token
	Secret        string   `json:"secret"`
	NickName      []string `json:"nick_name"`
	EventTypeList []int    `json:"event_type_list"` // 上线，下线，设置新命令，更新命令，删除命令，update下载了文件
	Id            int64    `json:"id"`
}

type Webhook struct {
	Types         int      `json:"types"` // 1 钉钉 2 其他自定义（推送为json信息）
	Url           string   `json:"url"`   //若为钉钉，此处填写token
	Secret        string   `json:"secret"`
	NickName      []string `json:"nick_name"`
	EventTypeList []int    `json:"event_type_list"` // 上线，下线，设置新命令，更新命令，删除命令，update下载了文件
}


type DelWebhook struct {
	Key int64 `json:"key"`
}

type KeepLive struct {
	NickName string `json:"nick_name"`
	Keep     int    `json:"keep"` // 1 保活 2 不保活
}

type Upload struct {
	NickName string `json:"nick_name"` // 程序昵称，别名
}

type IpStruct2 struct {
	Ip       string `json:"ip"`
	NickName string `json:"nick_name"`
	Types    int    `json:"types"`
}
type IpList2 struct {
	Ips   []IpStruct2 `json:"good_ips"`
	Token string      `json:"token"`
}

type Plugin struct {
	Name string `json:"name"`
	Port int64  `json:"port"`
}

type IpStruct struct {
	Ip       string `json:"ip"`
	NickName string `json:"nick_name"`
}

type Relay struct {
	Ip          string `json:"ip"`           //转发的机器ip
	Execute     int    `json:"execute"`      //1 开始 2 结束
	NickName    string `json:"nick_name"`    // 程序昵称，别名
	ProgramName string `json:"program_name"` // 程序名，启动用的名称
}

