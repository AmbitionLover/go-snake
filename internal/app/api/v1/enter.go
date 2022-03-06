package v1

import (
	"github.com/AmbitionLover/go-snake/internal/app/api/v1/apipost"
	"github.com/AmbitionLover/go-snake/internal/app/api/v1/portal"
	"github.com/AmbitionLover/go-snake/internal/app/api/v1/tools"
)

type ApiGroup struct {
	ApiPostGroup apipost.ApiGroup
	ToolsGroup   tools.ApiGroup
	PortalGroup  portal.ApiGroup
}

var ApiGroupApp = new(ApiGroup)
