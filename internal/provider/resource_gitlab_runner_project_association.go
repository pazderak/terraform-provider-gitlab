package provider

import (
	"context"
	"errors"
	"log"
	"strconv"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/xanzy/go-gitlab"
)

var _ = registerResource("gitlab_runner_project_association", func() *schema.Resource {
	return &schema.Resource{
		Description: "The " + "`gitlab_runner_project_association`" + ` allows to enable or disable a particular runner for project

A runner must be shared instance runner, i.e. registered at the instnance level

**Upstream API**:

	- [Gitlab REST API Docs](https://docs.gitlab.com/ee/api/runners.html#enable-a-runner-in-project)
	- [Gitlab REST API Docs](https://docs.gitlab.com/ee/api/runners.html#disable-a-runner-from-project)`,

		CreateContext: resourceGitlabRunnerProjectAssociationCreate,
		ReadContext:   resourceGitlabRunnerProjectAssociationRead,
		DeleteContext: resourceGitlabRunnerProjectAssociationDelete,

		Schema: map[string]*schema.Schema{
			"runner_id": {
				Description: "The ID of the runner",
				Type:        schema.TypeInt,
				Required:    true,
				ForceNew:    true,
			},
			"project": {
				Description: "The name or ID of the project.",
				Type:        schema.TypeString,
				Required:    true,
				ForceNew:    true,
			},
		},
	}
})

func resourceGitlabRunnerProjectAssociationCreate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(*gitlab.Client)
	project := d.Get("project").(string)
	runner_id := d.Get("runner_id").(int)

	options := &gitlab.EnableProjectRunnerOptions{
		RunnerID: *gitlab.Int(runner_id),
	}

	log.Printf("[DEBUG] enable gitlab runner %d for project %s", runner_id, project)

	runner, _, err := client.Runners.EnableProjectRunner(project, options, gitlab.WithContext(ctx))
	if err != nil {
		return diag.FromErr(err)
	}
	idstr := strconv.Itoa(runner.ID)
	if err != nil {
		return diag.FromErr(err)
	}

	d.SetId(buildTwoPartID(&project, &idstr))
	return resourceGitlabRunnerProjectAssociationRead(ctx, d, meta)
}

func resourceGitlabRunnerProjectAssociationRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	return diag.FromErr(errors.New("not implemented"))
}

func resourceGitlabRunnerProjectAssociationDelete(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	return diag.FromErr(errors.New("not implemented"))
}
