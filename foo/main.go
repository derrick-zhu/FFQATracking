package main

import (
	"bytes"
	"fmt"
	"log"
	"os"
	"os/exec"
	"path/filepath"
)

var (
	// TargetList ...
	TargetList = [2]string{"iphoneos", "iphonesimulator"}
	// ArchList ...
	ArchList = [...]string{"i386", "x86_64", "armv7", "armv7s", "arm64"}
)

const (
	cmdXcodebuild = "xcodebuild"
	cmdLibtool    = "libtool"
	cmdLipo       = "lipo"
)

func funnyFunc() {
	fmt.Printf(`
======================================================================
  FFSharing SDK 
  Copyright 2018 By Farfetch

  前方ハイエナ-ジの!前方ハイエナ-ジの!前方ハイエナ-ジの!前方ハイエナ-ジの!
======================================================================

`)
}

func showHelp() {

}

func main() {
	funnyFunc()

	cmd := CommandModel{
		execDir:     getExecPath(),
		derivedPath: "build",
		outPath:     "output",
		config:      "Release",
	}
	cmd.Parse(os.Args)

	cmdHelper := CommandModelHelper{cmd: cmd}
	cmdHelper.Description()

	for _, eachTarget := range TargetList {
		// var allFileList []string

		for _, eachArch := range ArchList {
			var buildArgs string

			buildArgs = buildArgs + cmdHelper.GetWorkspace()
			buildArgs = buildArgs + cmdHelper.GetProject()
			buildArgs = buildArgs + cmdHelper.GetScheme()

			buildArgs = buildArgs + cmdHelper.GetConfig()
			buildArgs = buildArgs + " -arch " + eachArch

			buildArgs = buildArgs + " -sdk " + eachTarget + "${BASH_REMATCH[1]}"
			buildArgs = buildArgs + cmdHelper.GetDerivedPath()
			// buildArgs = buildArgs + " " + cmdHelper.GetOutPath()

			buildCommand := cmdXcodebuild + buildArgs

			fmt.Println(buildCommand)
		}
	}
}

// execShell 阻塞式的执行外部shell命令的函数,等待执行完毕并返回标准输出
func execShell(s string) (string, error) {
	//函数返回一个*Cmd，用于使用给出的参数执行name指定的程序
	cmd := exec.Command("/bin/bash", "", s)
	//读取io.Writer类型的cmd.Stdout，再通过bytes.Buffer(缓冲byte类型的缓冲器)将byte类型转化为string类型(out.String():这是bytes类型提供的接口)
	var out bytes.Buffer
	cmd.Stdout = &out //Run执行c包含的命令，并阻塞直到完成。这里stdout被取出，cmd.Wait()无法正确获取stdin,stdout,stderr，则阻塞在那了
	cmd.Stderr = &out
	cmd.Stdin = os.Stdin
	err := cmd.Run()
	checkErr(err)
	return out.String(), err
}

func checkErr(err error) {
	if err != nil {
		fmt.Println(err)
	}
}

func getExecPath() string {
	dir, err := filepath.Abs(filepath.Dir(os.Args[0]))
	if err != nil {
		log.Fatal(err)
	}
	return dir
}
