package main

import (
	"fmt"
	"remote_dog/api"
)

func main()  {
	//默认调用
	testApi:=api.NewApi("127.0.0.1","1568",false)
	version,err:=testApi.Version()
	if err !=nil {
		fmt.Println(err)
	}
	fmt.Println(version)

	//设置基础认证,默认dog，admin
	testApi.SetBasicAuth("dog","admin")
	user,err:=testApi.UserCurrent()
	if err !=nil {
		fmt.Println(err)
	}
	fmt.Println(user)
	testApi.SetCommand(api.SetCommand{
		Url:         "",
		Params:      nil,
		Env:         nil,
		NickName:    "",
		ProgramName: "",
		Md5:         "",
		Delete:      false,
	})

}
