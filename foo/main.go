package main

import (
	"bytes"
	"fmt"
	"io"
	"log"
	"os"
	"os/exec"
	"path/filepath"
)

var (
	// TargetIphoneos ...
	TargetIphoneos = "iphoneos"
	// TargetSimulator ...
	TargetSimulator = "iphonesimulator"
	// TargetList ...
	TargetList = [2]string{TargetIphoneos, TargetSimulator}
	// ArchList ...
	ArchList = [...]string{"x86_64", "arm64", "arm64e", "armv7", "armv7s"}
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

	if result, _ := exists(cmd.DerivedPath()); result != true {
		os.Mkdir(cmd.DerivedPath(), os.ModePerm)
	}

	if result, _ := exists(cmd.OutPath()); result != true {
		os.Mkdir(cmd.OutPath(), os.ModePerm)
	}

	for _, eachTarget := range TargetList {
		// var allFileList []string

		// for _, eachArch := range ArchList {

		// if strings.HasPrefix(eachArch, "arm") {
		// 	if eachTarget != TargetIphoneos {
		// 		continue
		// 	}
		// } else {
		// 	if eachTarget != TargetSimulator {
		// 		continue
		// 	}
		// }

		var buildArgs []string

		buildArgs = append(buildArgs, cmdHelper.GetWorkspace()...)
		buildArgs = append(buildArgs, cmdHelper.GetProject()...)
		buildArgs = append(buildArgs, cmdHelper.GetScheme()...)

		buildArgs = append(buildArgs, cmdHelper.GetConfig()...)
		buildArgs = append(buildArgs, "-jobs", "4")

		for _, eachArch := range ArchList {
			buildArgs = append(buildArgs, "-arch", eachArch)
		}

		buildArgs = append(buildArgs, "-sdk", eachTarget)
		buildArgs = append(buildArgs, cmdHelper.GetDerivedPath()...)

		execShell(cmdXcodebuild, buildArgs)
		// }
	}
}

func copyAndCapture(w io.Writer, r io.Reader) ([]byte, error) {
	var out []byte
	buf := make([]byte, 1024, 1024)
	for {
		n, err := r.Read(buf[:])
		if n > 0 {
			d := buf[:n]
			out = append(out, d...)
			_, err := w.Write(d)
			if err != nil {
				return out, err
			}
		}
		if err != nil {
			// Read returns io.EOF at the end of file, which is not an error for us
			if err == io.EOF {
				err = nil
			}
			return out, err
		}
	}
	// never reached
	panic(true)
	return nil, nil
}

// exists returns whether the given file or directory exists
func exists(path string) (bool, error) {
	_, err := os.Stat(path)
	if err == nil {
		return true, nil
	}
	if os.IsNotExist(err) {
		return false, nil
	}
	return true, err
}

// execShell 阻塞式的执行外部shell命令的函数,等待执行完毕并返回标准输出
func execShell(s string, param []string) {
	//函数返回一个*Cmd，用于使用给出的参数执行name指定的程序
	cmd := exec.Command(s, param...)

	// var stdout, stderr []byte
	// var errStdout, errStderr error
	// stdoutIn, _ := cmd.StdoutPipe()
	// stderrIn, _ := cmd.StderrPipe()
	// cmd.Start()

	// go func() {
	// 	stdout, errStdout = copyAndCapture(os.Stdout, stdoutIn)
	// }()

	// go func() {
	// 	stderr, errStderr = copyAndCapture(os.Stderr, stderrIn)
	// }()

	// err := cmd.Wait()
	// if err != nil {
	// 	log.Fatalf("cmd.Run() failed with %s\n", err)
	// }
	// if errStdout != nil || errStderr != nil {
	// 	log.Fatalf("failed to capture stdout or stderr\n")
	// }
	// outStr, errStr := string(stdout), string(stderr)
	// fmt.Printf("\nout:\n%s\nerr:\n%s\n", outStr, errStr)

	//读取io.Writer类型的cmd.Stdout，再通过bytes.Buffer(缓冲byte类型的缓冲器)将byte类型转化为string类型(out.String():这是bytes类型提供的接口)
	var out bytes.Buffer
	cmd.Stdout = &out //Run执行c包含的命令，并阻塞直到完成。这里stdout被取出，cmd.Wait()无法正确获取stdin,stdout,stderr，则阻塞在那了
	cmd.Stderr = &out
	cmd.Stdin = os.Stdin
	err := cmd.Run()
	checkErr(err)
	fmt.Printf(out.String())
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
