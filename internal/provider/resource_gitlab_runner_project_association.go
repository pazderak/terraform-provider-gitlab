package provider

import "github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"

var _ = registerResource("gitlab_runner_project_association", func() *schema.Resource {
	return &schema.Resource{
		Description: "The" + "`gitlab_runner_project_association`" + `allows to enable or disable a particular runner for project

A runner must be shared instance runner, i.e. registered at the instnance level

**Upstream API**:

	- [Gitlab REST API Docs](https://docs.gitlab.com/ee/api/runners.html#enable-a-runner-in-project)
	- [Gitlab REST API Docs](https://docs.gitlab.com/ee/api/runners.html#disable-a-runner-from-project)`,
	}
})
