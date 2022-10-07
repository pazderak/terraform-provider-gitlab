//go:build acceptance
// +build acceptance

package provider

import (
	"fmt"
	"strconv"
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/v2/terraform"
	"github.com/xanzy/go-gitlab"
)

func TestAccGitlabRunnerProjectAssociation_basic(t *testing.T) {
	// This pulls from the gitlab.rb file, and is set on instance start-up
	token := "ACCTEST1234567890123_RUNNER_REG_TOKEN"

	project := testAccCreateProject(t)

	project, _, err := testGitlabClient.Projects.GetProject(project.ID, &gitlab.GetProjectOptions{})
	if err != nil {
		t.Fail()
	}

	resource.ParallelTest(t, resource.TestCase{
		ProviderFactories: providerFactories,
		CheckDestroy:      testAccCheckRunnerProjectAssocDestroy,
		Steps: []resource.TestStep{
			{
				Config: fmt.Sprintf(`
				resource "gitlab_runner" "this" {
					registration_token = "%s"
					description = "Lorem Ipsum"
				}

				resource "gitlab_runner_project_association" "this" {
					runner_id = gitlab_runner.this.id
					project = "%d"
				}
				`, token, project.ID),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttrSet("gitlab_runner.this", "registration_token"),
					resource.TestCheckResourceAttrSet("gitlab_runner.this", "authentication_token"),
				),
			}},
	})
}

func testAccCheckRunnerProjectAssocDestroy(state *terraform.State) error {
	for _, rs := range state.RootModule().Resources {
		if rs.Type != "gitlab_runner" && rs.Type != "gitlab_runner_project_association" {
			continue
		}

		id, _ := strconv.Atoi(rs.Primary.ID)

		runner, _, err := testGitlabClient.Runners.GetRunnerDetails(id)
		if err == nil {
			if runner != nil {
				return fmt.Errorf("runner still exists")
			}
		}

		if !is404(err) {
			return err
		}
		return nil
	}
	return nil
}
