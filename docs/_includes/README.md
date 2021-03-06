
# MY GO PACKAGES

### jeffshapes

_jeffshapes package makes it easy to compute area and circumference
of a circle._

[jeffshapes.go](https://github.com/JeffDeCola/my-go-packages/blob/master/jeffshapes/jeffshapes.go)
contains,

* type **Circle**
  * func (c Circle) **circleArea**() float64
  * func (c Circle) **circleCircumference**() float64

To use,

```bash
go get -u -v github.com/JeffDeCola/my-go-packages/jeffshapes
import github.com/JeffDeCola/my-go-packages/jeffshapes
```

Refer to
[test-jeffshapes](https://github.com/JeffDeCola/my-go-examples/tree/master/packages/test-jeffshapes)
in my repo `my-go-examples` for an example of its use.

## UPDATE GITHUB WEBPAGE USING CONCOURSE (OPTIONAL)

For fun, I use concourse to update
[my-go-packages GitHub Webpage](https://jeffdecola.github.io/my-go-packages/)
and alert me of the changes via repo status and slack.

A pipeline file [pipeline.yml](https://github.com/JeffDeCola/my-go-packages/tree/master/ci/pipeline.yml)
shows the entire ci flow. Visually, it looks like,

![IMAGE - my-go-packages concourse ci pipeline - IMAGE](pics/my-go-packages-pipeline.jpg)

The `jobs` and `tasks` are,

* `job-readme-github-pages` runs task
  [readme-github-pages.sh](https://github.com/JeffDeCola/my-go-packages/tree/master/ci/scripts/readme-github-pages.sh).

The concourse `resources types` are,

* `my-go-packages` uses a resource type
  [docker-image](https://hub.docker.com/r/concourse/git-resource/)
  to PULL a repo from github.
* `resource-slack-alert` uses a resource type
  [docker image](https://hub.docker.com/r/cfcommunity/slack-notification-resource)
  that will notify slack on your progress.
* `resource-repo-status` uses a resource type
  [docker image](https://hub.docker.com/r/dpb587/github-status-resource)
  that will update your git status for that particular commit.

For more information on using concourse for continuous integration,
refer to my cheat sheet on [concourse](https://github.com/JeffDeCola/my-cheat-sheets/tree/master/software/operations-tools/continuous-integration-continuous-deployment/concourse-cheat-sheet).
