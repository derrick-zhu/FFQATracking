package main

import (
	"fmt"
	"strings"
)

// CommandModelHelper ...
type CommandModelHelper struct {
	cmd CommandModel
}

func (c *CommandModelHelper) GetExecDir() string {
	return c.cmd.execDir
}

func (c *CommandModelHelper) GetWorkspace() string {
	if c.cmd.IsBuildWorkspace() {
		return " -workspace " + strings.Join([]string{c.cmd.execDir, c.cmd.xcworkspace}, "/")
	}
	return ""
}

func (c *CommandModelHelper) GetProject() string {
	if c.cmd.IsBuildProject() {
		return " -project " + strings.Join([]string{c.cmd.execDir, c.cmd.xcproject}, "/")
	}
	return ""
}

func (c *CommandModelHelper) GetScheme() string {
	if len(c.cmd.scheme) > 0 {
		return " -scheme " + c.cmd.scheme
	}
	return ""
}

func (c *CommandModelHelper) GetDerivedPath() string {
	if len(c.cmd.derivedPath) > 0 {
		return " -derivedDataPath " + strings.Join([]string{c.cmd.execDir, c.cmd.derivedPath}, "/")
	}
	return ""
}

func (c *CommandModelHelper) GetOutPath() string {
	return strings.Join([]string{c.cmd.execDir, c.cmd.outPath}, "/")
}

func (c *CommandModelHelper) GetConfig() string {
	if len(c.cmd.config) > 0 {
		return " -configuration " + c.cmd.config
	}
	return ""
}

func (c *CommandModelHelper) Description() {
	if len(c.cmd.execDir) != 0 {
		fmt.Println(c.GetExecDir())
	}

	if len(c.cmd.xcworkspace) != 0 {
		fmt.Println(c.GetWorkspace())
	}

	if len(c.cmd.xcproject) != 0 {
		fmt.Println(c.GetProject())
	}

	if len(c.cmd.derivedPath) != 0 {
		fmt.Println(c.GetDerivedPath())
	}

	if len(c.cmd.outPath) != 0 {
		fmt.Println(c.GetOutPath())
	}

	if len(c.cmd.config) != 0 {
		fmt.Println(c.GetConfig())
	}
}
