package main

import (
	"flag"
	"strings"
)

// CommandModel ...
type CommandModel struct {
	execDir     string
	xcworkspace string
	xcproject   string
	scheme      string
	sdk         string
	derivedPath string
	outPath     string
	config      string
}

func (c *CommandModel) Parse(args []string) {
	xcworkspace := flag.String("xcworkspace", "", "Define a Xcode workspace file to build.")
	xcproj := flag.String("xcproject", "", "Define a Xcode project file to build.")
	scheme := flag.String("scheme", "", "Scheme settings for building.")
	config := flag.String("config", "Release", "Build config setting likes, Debug, Release")
	derivedPath := flag.String("d", "build", "Derived Dir for Xcode")
	outPath := flag.String("o", "out", "Output dir for binary")

	flag.Parse()

	c.xcworkspace = *xcworkspace
	c.xcproject = *xcproj
	c.scheme = *scheme
	c.config = *config
	c.derivedPath = *derivedPath
	c.outPath = *outPath
}

func (c *CommandModel) IsValid() bool {
	if c.IsBuildWorkspace() == false && c.IsBuildProject() == false {
		return false
	}
	if c.IsSchemeValid() == false {
		return false
	}
	return true
}

func (c *CommandModel) IsBuildWorkspace() bool {
	return len(c.xcworkspace) > 0
}

func (c *CommandModel) IsBuildProject() bool {
	return len(c.xcproject) > 0
}

func (c *CommandModel) IsSchemeValid() bool {
	return len(c.scheme) > 0
}

func (c *CommandModel) ExecDir() string {
	return c.execDir
}

func (c *CommandModel) XCworkspace() string {
	return strings.Join([]string{c.execDir, c.xcworkspace}, "/")
}

func (c *CommandModel) XCproject() string {
	return strings.Join([]string{c.execDir, c.xcproject}, "/")
}

func (c *CommandModel) Scheme() string {
	return c.scheme
}

func (c *CommandModel) DerivedPath() string {
	return strings.Join([]string{c.execDir, c.derivedPath}, "/")
}

func (c *CommandModel) OutPath() string {
	return strings.Join([]string{c.execDir, c.outPath}, "/")
}

func (c *CommandModel) Config() string {
	return c.config
}
