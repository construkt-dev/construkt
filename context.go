package construkt

import (
	"context"
	"flag"
	"fmt"
	"k8s.io/client-go/rest"
	"k8s.io/client-go/tools/clientcmd"
	"k8s.io/client-go/util/homedir"
	"os"
	"os/signal"
	"path/filepath"
	"syscall"
	"time"
)

type Context struct {
	config       *rest.Config
	done         chan struct{}
	routinesDone map[string]chan struct{}
}

func (c *Context) Wait() {
	ctx, stop := signal.NotifyContext(context.Background(), os.Interrupt, syscall.SIGTERM)
	defer stop()

waitForSignal:
	for {
		select {
		case <-c.Done():
			break waitForSignal
		case <-ctx.Done():
			c.Shutdown()
		}
	}

	fmt.Printf("Waiting for routines to finish (timeout = 5 seconds)...\n")
	timeout := time.NewTimer(5 * time.Second)
waitForShutdown:
	for id, done := range c.routinesDone {
		fmt.Printf("Waiting for %s to finish\n", id)
		select {
		case <-timeout.C:
			fmt.Printf("Timeout reached, shutting down\n")
			break waitForShutdown
		case <-done:
			continue waitForShutdown
		}
	}
}

func (c *Context) Config() *rest.Config {
	return c.config
}

func (c *Context) Done() <-chan struct{} {
	return c.done
}

func (c *Context) Shutdown() {
	select {
	case <-c.done:
		return
	default:
		close(c.done)
	}
}

func NewContext() *Context {
	return &Context{
		config:       config(),
		done:         make(chan struct{}),
		routinesDone: make(map[string]chan struct{}),
	}
}

func (c *Context) RegisterComponent(id string) func() {
	done := make(chan struct{})
	c.routinesDone[id] = done
	return func() {
		close(done)
		delete(c.routinesDone, id)
	}
}

func config() *rest.Config {
	var kubeConfig *string
	if home := homedir.HomeDir(); home != "" {
		kubeConfig = flag.String("kubeconfig", filepath.Join(home, ".kube", "config"), "(optional) absolute path to the kubeconfig file")
	} else {
		kubeConfig = flag.String("kubeconfig", "", "absolute path to the kubeconfig file")
	}
	flag.Parse()

	config, err := clientcmd.BuildConfigFromFlags("", *kubeConfig)
	if err != nil {
		panic(err.Error())
	}

	return config
}
