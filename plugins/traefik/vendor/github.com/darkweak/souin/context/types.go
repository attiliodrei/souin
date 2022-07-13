package context

import (
	"net/http"

	"github.com/darkweak/souin/configurationtypes"
)

type (
	ctxKey string

	ctx interface {
		SetupContext(c configurationtypes.AbstractConfigurationInterface)
		SetContext(req *http.Request) *http.Request
	}

	Context struct {
		CacheName ctx
		GraphQL   ctx
		Key       ctx
		Method    ctx
	}
)

func GetContext() *Context {
	return &Context{
		CacheName: &cacheContext{},
		GraphQL:   &graphQLContext{},
		Key:       &keyContext{},
		Method:    &methodContext{},
	}
}

func (c *Context) Init(co configurationtypes.AbstractConfigurationInterface) {
	c.CacheName.SetupContext(co)
	c.GraphQL.SetupContext(co)
	c.Key.SetupContext(co)
	c.Method.SetupContext(co)
}

func (c *Context) SetBaseContext(req *http.Request) *http.Request {
	return c.Method.SetContext(c.CacheName.SetContext(req))
}

func (c *Context) SetContext(req *http.Request) *http.Request {
	return c.Key.SetContext(c.GraphQL.SetContext(req))
}
