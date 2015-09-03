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
	"io"
	"io/ioutil"
	"net/http"
)

func system(sudo bool, cmd string, arg string) string {
	var out string
	if (sudo) {
		if out, err := exec.Command("/bin/sh", "-c", "sudo " + cmd + " " + arg).Output(); err != nil {
			fmt.Fprintln(os.Stderr, "There was an error running git rev-parse command: ", err)
			fmt.Println(cmd + " " + arg)
			fmt.Println(out)
			os.Exit(1)
		}
	} else {
		if out, err := exec.Command(cmd, arg).Output(); err != nil {
			fmt.Fprintln(os.Stderr, "There was an error running git rev-parse command: ", err)
			fmt.Println(cmd + " " + arg)
			fmt.Println(out)
			os.Exit(1)
		}
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

func copy(target string, dest string) {
	r, err := os.Open(target)
	if err != nil {
		panic(err)
	}
	defer r.Close()

     w, err := os.Create(dest)
     if err != nil {
         panic(err)
     }
     defer w.Close()

     // do the actual work
     _, err = io.Copy(w, r)
     if err != nil {
         panic(err)
     }
}

func setUp() {
	pwd := os.Getenv("PWD")
	os.Chdir(pwd)
	system(false, "go", "build")
	system(true, "cp", "-f " + pwd + "/ant-worker /usr/bin/ant-worker")
	system(true, "chmod", "777 /usr/bin/ant-worker")
	system(true, "ant-worker", "install")
	system(true, "ant-worker", "start")
}

func tearDown() {
	system(true, "ant-worker", "stop")
	system(true, "ant-worker", "uninstall")
}

func TestService(t *testing.T) {
	assert := assert.New(t)
	setUp()
	defer tearDown()
	response, err := http.Get("http://localhost:2345/test")
	assert.Equal(true, err == nil)
	if err == nil {
		defer response.Body.Close()
		contents, err := ioutil.ReadAll(response.Body)
		assert.Equal(true, err == nil)
		assert.Equal("{\"status\":\"hello world\"}\n", string(contents))
	}
}