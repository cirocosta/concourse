local forkBranches = std.extVar('branches');

local resources = [{
  name: 'fork-' + branch,
  type: 'git',
  source: {
    uri: 'https://github.com/cirocosta/concourse',
    branch: branch,
  },
} for branch in forkBranches] + [{
  name: 'dev-image-' + branch,
  type: 'registry-image',
  source: {
    repository: 'cirocosta/concourse',
    tag: branch,
    username: '((docker-user))',
    password: '((docker-password))',
  },

} for branch in forkBranches];

local jobs = [
  {
    name: 'build-image-' + branch,
    plan: [
      {
        get: 'fork-' + branch,
        trigger: true,
      },
      {
        task: 'build',
        privileged: true,
        config: {
          platform: 'linux',
          image_resource: {
            type: 'registry-image',
            source: {
              repository: 'concourse/builder',
            },
          },
          container_limits: {},
          params: {
            REPOSITORY: 'cirocosta/concourse',
            TAG: branch,
          },
          run: {
            path: 'build',
          },
          inputs: [
            {
              name: 'fork-' + branch,
              path: './',
            },
          ],
          outputs: [
            {
              name: 'image',
            },
          ],
        },
      },
      {
        put: 'dev-image-' + branch,
        inputs: [
          'image',
        ],
        params: {
          image: 'image/image.tar',
        },
        get_params: {
          format: 'oci',
        },
      },
    ],
  }
  for branch in forkBranches
];

{
  groups: null,
  resources: resources,
  resource_types: null,
  jobs: jobs,
}
