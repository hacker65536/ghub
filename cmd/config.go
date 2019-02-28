package cmd

import (
	"fmt"
	"os"
	"os/exec"
	"path/filepath"
	"syscall"
)

type GithubConfig struct {
	UserName  string
	UserEmail string
	Host      string
	Repo      string
	Local     bool
	Init      bool
	Secrets   bool
}

var git = "git"

var G GithubConfig

func (g *GithubConfig) set() {
	/*
		execRun("git", "init")
		execRun("git", "config", "--local", "user.name", g.UserName)
		execRun("git", "config", "--local", "user.email", g.UserEmail)
	*/

	err := exec.Command(git, "status").Run()
	// exit status 128
	if err != nil {
		if e2, ok := err.(*exec.ExitError); ok {
			if s, ok := e2.Sys().(syscall.WaitStatus); ok {
				if s.ExitStatus() == 128 {
					g.init()
				} else {
					panic(err)
				}
			} else {
				panic(err)
			}
		}
	}

	g.config()
	g.remote()
	g.secrets()
	g.list()

}

func (g *GithubConfig) init() {
	execRun(git, "init")
	g.Init = true
}

func (g *GithubConfig) config() {
	execRun(git, "config", "--local", "user.name", g.UserName)
	execRun(git, "config", "--local", "user.email", g.UserEmail)
}

func (g *GithubConfig) remote() {
	op := "set-url"

	if g.Init {
		op = "add"

	}

	pwd, _ := os.Getwd()
	repo := filepath.Base(pwd)

	if g.Repo != "" {
		repo = g.Repo
	}

	arg := []string{
		"remote",
		op,
		"origin",
		"git@" + g.Host + ":" + g.UserName + "/" + repo,
	}
	execRun("git", arg...)
}

func (g *GithubConfig) secrets() {
	if g.Secrets {
		execRun(git, "secrets", "--install")
		execRun(git, "secrets", "--register-aws")
	}
}

func (g *GithubConfig) list() {
	execRun(git, "config", "--list")
}

func execRun(name string, args ...string) {
	out, err := exec.Command(name, args...).Output()
	if err != nil {
		fmt.Println("ERROR: ", name, args, err)
	}

	if len(out) != 0 {
		fmt.Printf("%s", out)
	}

}
