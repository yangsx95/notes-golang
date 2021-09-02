package main

import (
	"bufio"
	"fmt"
	"github.com/magiconair/properties"
	"io/ioutil"
	"log"
	"os"
	"path/filepath"
	"strings"
	"time"
)

var (
	// 当前执行程序的所在目录，也是项目所在目录
	execDir string
	// 行读取
	input *bufio.Scanner
)

const m = `@CopyRight 苏州农村商业银行
@author yangsx

项目路径： %s

选择操作：
	1. 修改别名
	2. 注释证书校验
	3. 注释登录密码校验
	4. 注释登录验证码校验
	5. 开启https登录
	6. 手机联调(包括2，3)
`

func init() {
	_execDir, err := filepath.Abs(os.Args[0])
	execDir = _execDir
	execDir = "/Users/yangsx/Desktop/framework-master"
	if err != nil {
		log.Fatalln("获取当前路径出错")
	}

	input = bufio.NewScanner(os.Stdin) // 按行扫描代码扫描器
}

// 用于初始化项目的命令行小工具
func main() {
	fmt.Printf(`@CopyRight 苏州农村商业银行
@author yangsx

项目路径： %s
`, execDir)
	for readLine(`选择操作：
	1. 修改别名
	2. 注释证书校验
	3. 注释登录密码校验
	4. 注释登录验证码校验
	5. 开启https登录
	6. 手机联调(包括2，3，4)
`) { // 读入下一行，并移除行末的换行符，如果到达没有后面一行，函数返回false
		switch input.Text() {
		case "1":
			jsfAliasModify()
		case "2":
			certRemove()
		case "3":
			pwdRemove()
		case "4":
			vcodeRemove()
		case "5":
			httpsOpen()
		case "6":
			withMob()
		default:
			fmt.Println("输入错误")
		}
		time.Sleep(1 * time.Second)
		fmt.Println()
	}
}

func vcodeRemove() {

}

func readLine(format string, a ...interface{}) bool {
	if format != "" {
		fmt.Printf(format, a...)
		fmt.Println()
	}
	fmt.Print("> ")
	return input.Scan()
}

func jsfAliasModify() {
	suffix := ""

	if readLine("请输入别名：") {
		suffix = input.Text()
	}
	if suffix == "" {
		return
	}

	pf := filepath.Join(execDir, "framework-common", "src", "main", "resources", "config", "mode", "dev", "jsf-alias.properties")
	p, err := properties.LoadFile(pf, properties.UTF8)
	if err != nil {
		log.Println("读取文件出错，路径：", pf)
		return
	}

	var versions = [10]string{"version.mca", "version.mcm", "version.auth",
		"version.transfer", "version.limit", "version.product", "version.rt",
		"version.ms-user", "version.batch", "version.points"}

	for _, version := range versions {
		p.Set(version, p.GetString(version, "")+suffix)
	}

	w, _ := os.OpenFile(pf, os.O_RDWR|os.O_TRUNC, os.ModeSetuid)
	defer w.Close()

	n, err := p.Write(w, properties.UTF8)
	if err != nil || n == 0 {
		log.Println("写入properties文件失败", err)
	}

	fmt.Println("------操作成功-------")
}

func withMob() {
	certRemove()
	//pwdRemove()
	//vcodeRemove()

	p := filepath.Join(execDir, "pe-fix", "src", "main", "java", "com", "csii", "pe", "fix", "channel", "http", "servlet", "DecryptJsonView.java")
	data := make(map[int]string, 2)
	data[158] = ""
	data[159] = ""
	modifyLineByNumber(p, data)

	p = filepath.Join(execDir, "framework-web", "eweb", "eweb-core", "src", "main", "java", "com", "csii", "ibs", "common", "policy", "BindDeviceControlPolicy.java")
	data = make(map[int]string, 16)
	for i := 20; i < 36; i++ {
		data[i] = ""
	}
	modifyLineByNumber(p, data)

	p = filepath.Join(execDir, "pe-fix", "src", "main", "java", "com", "csii", "pe", "transform", "DecryptPackageParse.java")
	data = make(map[int]string, 7)
	for i := 16; i < 23; i++ {
		data[i] = ""
		if i == 22 {
			data[i] = "\t\treturn obj;"
		}
	}
	modifyLineByNumber(p, data)

}

func httpsOpen() {
	var (
		// ssl配置文件路径
		sslp = []string{filepath.Join(execDir, "framework-web", "eweb", "eweb-starter", "src", "main", "resources", "application-dev.properties")}
		sslm = map[string]string{
			"server.ssl.key-store":            "classpath:server.jks",
			"server.ssl.key-store-password":   "123456",
			"server.ssl.key-password":         "123456",
			"server.ssl.trust-store":          "classpath:server.jks",
			"server.ssl.trust-store-password": "123456",
			"server.ssl.client-auth":          "need",
		}
		// 证书服务配置文件路径
		sps = []string{
			filepath.Join(execDir, "framework-micro-services", "user", "user-core", "src", "main", "resources", "pe", "config.properties"),
			filepath.Join(execDir, "framework-router", "router-services", "src", "main", "java", "config.properties"),
			filepath.Join(execDir, "framework-router", "router-services", "src", "main", "java", "config_en_US.properties"),
			filepath.Join(execDir, "framework-router", "router-services", "src", "main", "java", "config_zh_CN.properties"),
		}
		spm = map[string]string{
			"hosts": "192.168.169.191",
		}
		// 验签服务配置文件路径
		nps = []string{
			filepath.Join(execDir, "framework-router", "router-services", "src", "main", "java", "netsignagent.properties"),
			filepath.Join(execDir, "framework-web", "eweb", "eweb-core", "src", "main", "java", "netsignagent.properties"),
		}
		npm = map[string]string{
			"ServerIP":   "192.168.169.185",
			"ServerPort": "8888",
			"Timeout":    "30000",
		}
	)
	_ = modifyProperties(sslp, sslm)
	_ = modifyProperties(sps, spm)
	_ = modifyProperties(nps, npm)

	fmt.Println("操作成功")
}

func modifyProperties(paths []string, kv map[string]string) error {
	for _, path := range paths {
		p, err := properties.LoadFile(path, properties.UTF8)
		if err != nil {
			return err
		}
		for k, v := range kv {
			p.Set(k, v)
		}
		file, err := os.OpenFile(path, os.O_TRUNC|os.O_WRONLY, 0644)
		if err != nil {
			return err
		}
		p.Write(file, properties.UTF8)
		fmt.Println("处理文件", path, "成功")
	}
	return nil
}

func pwdRemove() {

}

func certRemove() {
	p := filepath.Join(execDir, "framework-web", "eweb", "eweb-core", "src", "main", "java", "com", "csii", "framework", "common", "login", "MyLoginAction.java")
	// 读取源代码
	err := cmLineByKeyword(p, `CsiiUtils.getClientCertificate(context, null);`, `user.setDigitalSignEnabled(true)`)
	showError(err)
	fmt.Println("操作成功")
}

func cmLineByKeyword(path string, startLine string, endLine string) error {
	fs, err := ioutil.ReadFile(path)
	if err != nil {
		return err
	}
	lines := strings.Split(string(fs), "\n")

	// 修改源代码
	flag := false
	for i, line := range lines {
		if strings.Contains(line, startLine) {
			flag = true
		} else if strings.Contains(line, endLine) {
			flag = false
		}
		if flag {
			lines[i] = "//" + line
		}
	}
	code := strings.Join(lines, "\n")

	// 写入源代码
	return ioutil.WriteFile(path, []byte(code), 0644)
}

// 按行修改
func modifyLineByNumber(path string, data map[int]string) error {
	fs, err := ioutil.ReadFile(path)
	if err != nil {
		return err
	}
	lines := strings.Split(string(fs), "\n")

	// 修改源代码
	for i, _ := range lines {
		d, ok := data[i]
		if ok {
			lines[i] = d
		}
	}
	code := strings.Join(lines, "\n")

	// 写入源代码
	return ioutil.WriteFile(path, []byte(code), 0644)
}

func showError(err error) {
	if err != nil {
		log.Fatalln(err)
	}
}
