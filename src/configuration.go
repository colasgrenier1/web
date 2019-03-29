//
// C O N F I G U R A T I O N   M A N A G E R
//

package main

import (
	"bufio"
	"errors"
	"fmt"
	"os"
	"strconv"
	"strings"
	"unicode"
)

type Configuration struct {
	//server information
	Port int

	//session informatio
	SessionTimeout int //in seconds

	//loggin
	LogFile string

	//database
	DatabaseAddress string
	DatabasePort int
	DatabaseUsername string
	DatabasePassword string
	DatabaseName string
}

func ReadConfigurationFile(path string) (*Configuration, error) {
	var f *os.File
	var e error
	var s *bufio.Scanner
	var c *Configuration = &Configuration{}
	var t []rune
	var v,k string

	//Try to open the file
	f, e = os.Open(path)
	if e != nil {
		return nil, e
	}
	defer f.Close()

	//Build a scanner
	s = bufio.NewScanner(f)

	//We read line by line
	for s.Scan() {
		t = []rune(strings.TrimRightFunc(s.Text(), unicode.IsSpace))
		//we do not take into account empty lines
		if len(t) != 0 {
			//we do not take into account lines with characters in the first column
			if t[0] == rune(' ') {
				if len(t) < 25 {
					return nil, errors.New(fmt.Sprintf("%v", t))
				} else {
					k = strings.TrimSpace(string(t[1:25]))
					v = strings.TrimSpace(string(t[26:]))
					if k == "PORT" {
						c.Port, e = strconv.Atoi(v)
						if e != nil {
							return nil, errors.New("PORT NUMBER IS NOT AN INTEGER")
						}
					} else if k == "SESSIONTIMEOUT" {
						c.SessionTimeout, e = strconv.Atoi(v)
						if e != nil {
							return nil, errors.New("SESSION TIMEOUT IS NOT A NUMBER")
						}
					} else if k == "DATABASEPORT" {
						c.DatabasePort, e = strconv.Atoi(v)
						if e != nil {
							return nil, errors.New("DATABASE PORT IS NOT A NUMBER")
						}
					} else if k=="LOGFILE" {
						c.LogFile = v
					} else if k=="DATABASEADDRESS" {
						c.DatabaseAddress = v
					} else if k=="DATABASEUSERNAME" {
						c.DatabaseUsername = v
					} else if k=="DATABASEPASSWORD" {
						c.DatabasePassword = v
					} else if k=="DATABASENAME" {
						c.DatabaseName = v
					} else {
						return nil, errors.New(fmt.Sprintf("UNKNOWN FIELD: %s", k))
					}
				}
			}
		}
	}

	//We return
	return c, nil
}

func main() {
	m, e := ReadConfigurationFile("conf")
	fmt.Printf("%v %v\n", m, e)
}
