package pushaction

import (
	"io"

	"code.cloudfoundry.org/cli/actor/v2action"
)

//go:generate counterfeiter . V2Actor

type V2Actor interface {
	BindRouteToApplication(routeGUID string, appGUID string) (v2action.Warnings, error)
	CheckRoute(route v2action.Route) (bool, v2action.Warnings, error)
	CreateApplication(application v2action.Application) (v2action.Application, v2action.Warnings, error)
	CreateRoute(route v2action.Route, generatePort bool) (v2action.Route, v2action.Warnings, error)
	GatherResources(sourceDir string) ([]v2action.Resource, error)
	GetApplicationByNameAndSpace(name string, spaceGUID string) (v2action.Application, v2action.Warnings, error)
	GetApplicationRoutes(applicationGUID string) ([]v2action.Route, v2action.Warnings, error)
	GetOrganizationDomains(orgGUID string) ([]v2action.Domain, v2action.Warnings, error)
	GetRouteByHostAndDomain(host string, domainGUID string) (v2action.Route, v2action.Warnings, error)
	UpdateApplication(application v2action.Application) (v2action.Application, v2action.Warnings, error)
	UploadApplicationPackage(appGUID string, existingResources []v2action.Resource, newResources io.Reader, newResourcesLength int64) (v2action.Warnings, error)
	ZipResources(sourceDir string, filesToInclude []v2action.Resource) (string, error)
}
