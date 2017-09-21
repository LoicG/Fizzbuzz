package main

import (
	"client"
	"fmt"
	"net/http"
	"os"
	"os/exec"
	"path/filepath"
	"testing"

	. "gopkg.in/check.v1"
)

type TestSuite struct{}

func Test(t *testing.T) { TestingT(t) }

var (
	_ = Suite(&TestSuite{})
	// Default test port
	TestPort = uint32(8085)
)

// Process wrapper
type ServerProcess struct {
	Port uint32
	cmd  *exec.Cmd
}

func (s *ServerProcess) Stop() {
	s.cmd.Process.Kill()
}

func (s *ServerProcess) Start() error {
	return s.cmd.Start()
}

func (s *ServerProcess) Wait() {
	s.cmd.Process.Wait()
}

// startServerProcess starts fizzbuzz.exe process from the $GOBIN directory
func startServerProcess(router string) (*ServerProcess, error) {
	gobin := os.Getenv("GOBIN")
	server := &ServerProcess{
		Port: TestPort,
		cmd: exec.Command(filepath.Join(gobin, "fizzbuzz.exe"),
			[]string{"-port", fmt.Sprintf("%d", TestPort), "-router", router}...),
	}
	err := server.Start()
	if err != nil {
		return nil, err
	}
	return server, nil
}

// startServer returns a running process server and a http client
func startServer(router string) (*ServerProcess, *client.Client, error) {
	server, err := startServerProcess(router)
	if err != nil {
		return nil, nil, err
	}
	return server, client.NewClient(server.Port, http.DefaultMaxIdleConnsPerHost), nil
}

func checkError(c *C, client *client.Client, int1, int2, limit int, string1,
	string2 string) {

	_, err := client.FizzBuzz(int1, int2, limit, string1, string2)
	c.Assert(err, NotNil)
	c.Assert(err.Error(), Equals,
		fmt.Sprintf("server request failed, error code: %d",
			http.StatusBadRequest))
}

func testInvalidParameters(c *C, router string) {
	server, client, err := startServer(router)
	c.Assert(err, IsNil)
	c.Assert(server, NotNil)
	c.Assert(client, NotNil)
	defer server.Stop()

	// invalid int1 parameter
	checkError(c, client, 0, 0, 0, "", "")

	// invalid int2 parameter
	checkError(c, client, 3, 0, 0, "", "")

	// invalid string1 parameter
	checkError(c, client, 3, 5, 0, "", "")

	// invalid string2 parameter
	checkError(c, client, 3, 5, 0, "fizz", "")
}

func (s *TestSuite) TestGorillaInvalidParameters(c *C) {
	testInvalidParameters(c, "gorilla")
}

func (s *TestSuite) TestGojiInvalidParameters(c *C) {
	testInvalidParameters(c, "goji")
}

func (s *TestSuite) TestEmickleiInvalidParameters(c *C) {
	testInvalidParameters(c, "emicklei")
}

func testFizzBuzz(c *C, router string) {
	server, client, err := startServer(router)
	c.Assert(err, IsNil)
	c.Assert(server, NotNil)
	c.Assert(client, NotNil)
	defer server.Stop()

	// limit: -8 return empty list
	limit := -8
	result, err := client.FizzBuzz(3, 5, limit, "fizz", "buzz")
	c.Assert(err, IsNil)
	c.Assert(result, DeepEquals, []string{})

	// limit: 0 return empty list
	limit = 0
	result, err = client.FizzBuzz(3, 5, limit, "fizz", "buzz")
	c.Assert(err, IsNil)
	c.Assert(result, HasLen, limit)
	c.Assert(result, DeepEquals, []string{})

	// limit: 1 return "1"
	limit = 1
	result, err = client.FizzBuzz(3, 5, limit, "fizz", "buzz")
	c.Assert(err, IsNil)
	c.Assert(result, HasLen, limit)
	c.Assert(result, DeepEquals, []string{"1"})

	// fizz-buzz case
	limit = 16
	result, err = client.FizzBuzz(3, 5, limit, "fizz", "buzz")
	c.Assert(err, IsNil)
	c.Assert(result, HasLen, limit)
	c.Assert(result, DeepEquals, []string{"1", "2", "fizz", "4", "buzz",
		"fizz", "7", "8", "fizz", "buzz", "11", "fizz", "13", "14", "fizzbuzz", "16"})

	// fizz-fizz case
	result, err = client.FizzBuzz(2, 7, limit, "fizz", "fizz")
	c.Assert(err, IsNil)
	c.Assert(result, HasLen, limit)
	c.Assert(result, DeepEquals, []string{"1", "fizz", "3", "fizz", "5",
		"fizz", "fizz", "fizz", "9", "fizz", "11", "fizz", "13", "fizzfizz", "15", "fizz"})

	// pair case
	limit = 6
	result, err = client.FizzBuzz(2, 2, limit, "fizz", "buzz")
	c.Assert(err, IsNil)
	c.Assert(result, HasLen, limit)
	c.Assert(result, DeepEquals, []string{"1", "fizzbuzz", "3", "fizzbuzz",
		"5", "fizzbuzz"})

	// unchanged case
	result, err = client.FizzBuzz(8, 8, limit, "fizz", "buzz")
	c.Assert(err, IsNil)
	c.Assert(result, HasLen, limit)
	c.Assert(result, DeepEquals, []string{"1", "2", "3", "4", "5", "6"})

	// 1 case
	result, err = client.FizzBuzz(1, 1, limit, "fizz", "buzz")
	c.Assert(err, IsNil)
	c.Assert(result, HasLen, limit)
	c.Assert(result, DeepEquals, []string{"fizzbuzz", "fizzbuzz", "fizzbuzz",
		"fizzbuzz", "fizzbuzz", "fizzbuzz"})
}

func (s *TestSuite) TestGorillaFizzBuzz(c *C) {
	testFizzBuzz(c, "gorilla")
}

func (s *TestSuite) TestGojiFizzBuzzs(c *C) {
	testFizzBuzz(c, "goji")
}

func (s *TestSuite) TestEmickleiFizzBuzz(c *C) {
	testFizzBuzz(c, "emicklei")
}
