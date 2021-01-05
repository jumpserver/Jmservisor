package main

import (
	"encoding/base64"
	"encoding/json"
	"flag"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"os/exec"
	"path"
	"runtime"
)

var AppletsDir = path.Join("/opt", "JumpServerRemoteApplets")

func osIsWindows() bool {
	return runtime.GOOS == "windows"
}

func init() {
	if runtime.GOOS == "windows" {
		AppletsDir = path.Join("c:", "JumpServerRemoteApplets")
	}
}

// 新建一个 applet
func NewApplet() {

}

// 列出目录中的 applet
func ListApplet() (applets []string) {
	fmt.Println("List applets")
	files, err := ioutil.ReadDir(AppletsDir)
	if err != nil {
		log.Println(err)
		return applets
	}
	for _, file := range files {
		if file.IsDir() {
			applets = append(applets, file.Name())
			fmt.Println(file.Name())
		}
	}
	fmt.Println("")
	return applets
}

func GetAppletProcPath(applet string) string {
	procName := "main"
	if osIsWindows() {
		procName = "main.exe"
	}
	theAppletDir := path.Join(AppletsDir, applet)
	theAppletMainPath := path.Join(theAppletDir, procName)
	return theAppletMainPath
}

func HasApplet(applet string) bool {
	theAppletMainPath := GetAppletProcPath(applet)
	if _, err := os.Stat(theAppletMainPath); os.IsNotExist(err) {
		fmt.Println("Applet Proc not exist: ", theAppletMainPath)
		return false
	}
	return true
}

func decodeAppletArgs(args string, argsInter *map[string]interface{}) (err error) {
	if args == "" {
		return nil
	}
	argsJson, err := base64.StdEncoding.DecodeString(args)
	if err != nil {
		err = fmt.Errorf("decode args error: %s", args)
		return
	}
	err = json.Unmarshal(argsJson, argsInter)
	if err != nil {
		err = fmt.Errorf("json decode args error: %s", argsJson)
		return err
	}
	return nil
}

func RunApplet(applet string, args string) error {
	var err error
	if !HasApplet(applet) {
		err = fmt.Errorf("applet not exists: %s", applet)
		fmt.Println(err)
		return err
	}

	var argsInter map[string]interface{}
	fmt.Println("Args is: ", args)
	err = decodeAppletArgs(args, &argsInter)
	if err != nil {
		err = fmt.Errorf("decode applet args error: %s", err)
		return err
	}
	var commands []string
	if osIsWindows() {
		commands = []string{"cmd.exe", "/C", "start"}
	}

	theAppletMainPath := GetAppletProcPath(applet)
	commands = append(commands, theAppletMainPath)
	for k, v := range argsInter {
		vs := fmt.Sprintf("%v", v)
		commands = append(commands, fmt.Sprintf("-%s", k))
		commands = append(commands, vs)
	}
	cmd := exec.Command(commands[0], commands[1:]...)
	stdoutStderr, err := cmd.CombinedOutput()
	if err != nil {
		log.Println("Error:", err)
	}
	fmt.Printf("%s\n", stdoutStderr)
	return err
}

func main() {
	action := flag.String("s", "start", "start | list applet")
	applet := flag.String("app", "chrome", "Remote applet name")
	args := flag.String("arg", "", "Applet args base64 code")
	flag.Parse()

	fmt.Printf("Action: %s, Applet: %s, Args: %s\n\n", *action, *applet, *args)
	ListApplet()

	if HasApplet(*applet) {
		fmt.Println("Has applet: ", *applet)
	} else {
		fmt.Println("No applet: ", *applet)
	}

	err := RunApplet(*applet, *args)
	if err != nil {
		fmt.Println("Run applet error: ", err)
	}
	var input string
	_, _ = fmt.Scanln(&input)
}
