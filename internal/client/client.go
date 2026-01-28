package client

import (
	"context"

	"github.com/drone/drone-go/drone"
	"golang.org/x/oauth2"
)

type Client interface {
	ListRepos() ([]*drone.Repo, error)
	ListBuilds(namespace, name string, page int) ([]*drone.Build, error)
	GetBuild(namespace, name string, number int) (*drone.Build, error)
	GetLogs(owner, name string, build, stage, step int) ([]*drone.Line, error)
}

type droneClient struct {
	inner drone.Client
}

func New(server, token string) Client {
	conf := new(oauth2.Config)
	auth := conf.Client(context.Background(), &oauth2.Token{AccessToken: token})
	return &droneClient{inner: drone.NewClient(server, auth)}
}

func (c *droneClient) ListRepos() ([]*drone.Repo, error) {
	return c.inner.RepoList()
}

func (c *droneClient) ListBuilds(namespace, name string, page int) ([]*drone.Build, error) {
	return c.inner.BuildList(namespace, name, drone.ListOptions{Page: page})
}

func (c *droneClient) GetBuild(namespace, name string, number int) (*drone.Build, error) {
	return c.inner.Build(namespace, name, number)
}

func (c *droneClient) GetLogs(owner, name string, build, stage, step int) ([]*drone.Line, error) {
	return c.inner.Logs(owner, name, build, stage, step)
}
