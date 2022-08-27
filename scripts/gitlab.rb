pages_external_url 'http://127.0.0.1:5051'
pages_nginx['redirect_http_to_https'] = false
pages_nginx['ssl_certificate'] = "/etc/gitlab/ssl/gitlab-registry.pem"
pages_nginx['ssl_certificate_key'] = "/etc/gitlab/ssl/gitlab-registry.key"

registry_external_url 'http://127.0.0.1:5050'
registry['enable']                    = true
registry_nginx['ssl_certificate']     = "/etc/gitlab/ssl/gitlab-registry.pem"
registry_nginx['ssl_certificate_key'] = "/etc/gitlab/ssl/gitlab-registry.key"

gitlab_rails['initial_shared_runners_registration_token'] = "ACCTEST1234567890123_RUNNER_REG_TOKEN"

# This setting is required to disable caching for application settings
# which is required to test different scenarios in the acceptance tests.
# see https://gitlab.com/gitlab-org/gitlab/-/issues/364812#note_986366898
# see https://github.com/gitlabhq/terraform-provider-gitlab/pull/1128
gitlab_rails['application_settings_cache_seconds'] = 0
gitlab_rails['env'] = {
    'IN_MEMORY_APPLICATION_SETTINGS' => 'false'
}

# Enable SAML authentication for GitLab (required for SAML group links).
# see https://docs.gitlab.com/ee/integration/saml.html
gitlab_rails['omniauth_allow_single_sign_on'] = ['saml']
gitlab_rails['omniauth_block_auto_created_users'] = false
gitlab_rails['omniauth_providers'] = [
  {
    name: "saml",
    label: "Dummy Test Provider",
    args: {
      assertion_consumer_service_url: "https://gitlab.example.com/users/auth/saml/callback",
      idp_cert_fingerprint: "aa:aa:aa:aa:aa:aa:aa:aa:aa:aa:aa:aa:aa:aa:aa:aa",
      idp_sso_target_url: "https://saml.provider.example.com/sso/saml",
      issuer: "https://gitlab.example.com",
      name_identifier_format: "urn:oasis:names:tc:SAML:2.0:nameid-format:persistent"
    }
  }
]