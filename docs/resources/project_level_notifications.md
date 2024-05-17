---
# generated by https://github.com/hashicorp/terraform-plugin-docs
page_title: "gitlab_project_level_notifications Resource - terraform-provider-gitlab"
subcategory: ""
description: |-
  The gitlab_project_level_notifications resource allows to manage notifications for a project.
  ~> While the API supports both groups and projects, this resource only supports projects currently.
  Upstream API: GitLab REST API docs https://docs.gitlab.com/ee/api/notification_settings.html#group--project-level-notification-settings
---

# gitlab_project_level_notifications (Resource)

The `gitlab_project_level_notifications` resource allows to manage notifications for a project.

~> While the API supports both groups and projects, this resource only supports projects currently.
		
**Upstream API**: [GitLab REST API docs](https://docs.gitlab.com/ee/api/notification_settings.html#group--project-level-notification-settings)



<!-- schema generated by tfplugindocs -->
## Schema

### Required

- `project` (String) The ID or URL-encoded path of a project where notifications will be configured.

### Optional

- `close_issue` (Boolean) Enable notifications for closed issues. Can only be used when `level` is `custom`.
- `close_merge_request` (Boolean) Enable notifications for closed merge requests. Can only be used when `level` is `custom`.
- `failed_pipeline` (Boolean) Enable notifications for failed pipelines. Can only be used when `level` is `custom`.
- `fixed_pipeline` (Boolean) Enable notifications for fixed pipelines. Can only be used when `level` is `custom`.
- `issue_due` (Boolean) Enable notifications for due issues. Can only be used when `level` is `custom`.
- `level` (String) The level of the notification. Valid values are: `disabled`, `participating`, `watch`, `global`, `mention`, `custom`.
- `merge_merge_request` (Boolean) Enable notifications for merged merge requests. Can only be used when `level` is `custom`.
- `merge_when_pipeline_succeeds` (Boolean) Enable notifications for merged merge requests when the pipeline succeeds. Can only be used when `level` is `custom`.
- `moved_project` (Boolean) Enable notifications for moved projects. Can only be used when `level` is `custom`.
- `new_issue` (Boolean) Enable notifications for new issues. Can only be used when `level` is `custom`.
- `new_merge_request` (Boolean) Enable notifications for new merge requests. Can only be used when `level` is `custom`.
- `new_note` (Boolean) Enable notifications for new notes on merge requests. Can only be used when `level` is `custom`.
- `push_to_merge_request` (Boolean) Enable notifications for push to merge request branches. Can only be used when `level` is `custom`.
- `reassign_issue` (Boolean) Enable notifications for issue reassignments. Can only be used when `level` is `custom`.
- `reassign_merge_request` (Boolean) Enable notifications for merge request reassignments. Can only be used when `level` is `custom`.
- `reopen_issue` (Boolean) Enable notifications for reopened issues. Can only be used when `level` is `custom`.
- `reopen_merge_request` (Boolean) Enable notifications for reopened merge requests. Can only be used when `level` is `custom`.
- `success_pipeline` (Boolean) Enable notifications for successful pipelines. Can only be used when `level` is `custom`.

### Read-Only

- `id` (String) The ID of the resource. Matches the `project` value.