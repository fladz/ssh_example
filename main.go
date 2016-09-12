package main

import (
	"bytes"
	"ext/ssh"
	"fmt"
	"flag"
	"io/ioutil"
	"os/user"
)

var server string

func init() {
	flag.StringVar(&server, "server", "", "Server to SSH into & run command on")
}

func main() {
	// Get server to connect to.
	flag.Parse()
	if server == "" {
		flag.PrintDefaults()
		return
	}

	// Parse RSA file so we can call ssh without password.
	// The RSA file should be in $HOME/.ssh/id_rsa.
	u, err := user.Current()
	if err != nil {
		fmt.Printf("user.Current - %s\n", err)
		return
	}
	rsaFile := u.HomeDir + "/.ssh/id_rsa"

	b, err := ioutil.ReadFile(rsaFile)
	if err != nil {
		fmt.Printf("ReadFile - %s\n", err)
		return
	}

	s, err := ssh.ParsePrivateKey(b)
	if err != nil {
		fmt.Printf("ParsePrivateKey - %s\n", err)
		return
	}

	// Create SSH configuration struct.
	config := &ssh.ClientConfig{
		User: u.Username,
		Auth: []ssh.AuthMethod{ssh.PublicKeys(s)},
	}

	// Initialize SSH connection.
	client, err := ssh.Dial("tcp", server + ":22", config)
	if err != nil {
		fmt.Printf("Dial error - %s\n", err)
		return
	}

	// Ok now run a command.
	// In this example, run "date" command on remote host and print out the output.
	//
	// Before running a command, let's prepare a session over the init'd ssh connection.
	session, err := client.NewSession()
	if err != nil {
		fmt.Printf("NewSession - %s\n", err)
		return
	}

	// Set output to the buffer for debugging.
	var buf bytes.Buffer
	session.Stdout = &buf

	if err = session.Run("date"); err != nil {
		fmt.Printf("Run - %s\n", err)
		return
	}

	fmt.Printf("DATE IS: %s", buf.String())
}
