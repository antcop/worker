/**
 * Ant Worker Project
 *
 * Copyright (c) 2015 Epinion Online Research Team
 *
 * --------------------------------------------------------------------
 *
 * This program is free software: you can redistribute it and/or
 * modify it under the terms of the GNU Affero General Public License
 * as published by the Free Software Foundation, either version 3
 * of the License, or (at your option) any later version.
 *
 * This program is distributed in the hope that it will be useful,
 * but WITHOUT ANY WARRANTY; without even the implied warranty of
 * MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE. See the
 * GNU Affero General Public License for more details.
 *
 * You should have received a copy of the GNU Affero General Public
 * License along with this program.
 * If not, see <http://www.gnu.org/licenses/>.
 *
 * --------------------------------------------------------------------
 *
 * Author:
 *     Jerry Pham       <jerry@andjerry.com>
 *     Loi Nguyen       <loint@penlook.com>
 */

package main

import (
	"fmt"
	"testing"
	"github.com/stretchr/testify/assert"
	"path/filepath"
	"os"
	"log"
	"os/exec"
)

func system(cmd string, arg string) string {
	var out string
	if out, err := exec.Command(cmd, arg).Output(); err != nil {
		fmt.Fprintln(os.Stderr, "There was an error running git rev-parse command: ", err)
		fmt.Println(out)
		os.Exit(1)
	}
	return string(out)
}

func cwd() string {
	dir, err := filepath.Abs(filepath.Dir(os.Args[0]))
	if err != nil {
		log.Fatal(err)
	}
	return string(dir)
}

func TestService(t *testing.T) {
	assert := assert.New(t)
	assert.Equal("hello", "hello")
	
	// Build and start service
	//cur := cwd()
	pwd := os.Getenv("PWD")
	os.Chdir(pwd)
	system("go", "build")
	os.Link(pwd + "/ant-worker", "/usr/bin/ant-worker")
	//system("ant-worker", "install")
	//system("./ant-worker", "start")
	/*
	assert.Equal("hello", "hello")
	
	// Stop and cleanup
	system("./ant-worker", "stop")
	system("./ant-worker", "uninstall")
	os.Chdir(cur)
	*/
}