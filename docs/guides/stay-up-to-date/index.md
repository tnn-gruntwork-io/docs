# Update Guide

This guide will show you how to use Terragrunt 2.0 to keep your infrastructure up to date.

## Terragrunt 2.0

Terragrunt 2.0 is the new version of Terragrunt! It's a full-stack framework for Terraform that takes all the
functionality you loved from Terragrunt 1.0 and adds all the features you need to implement the full Terraform
lifecycle.

### Terragrunt 2.0 Features

:::info

Terragrunt 2.0 is a CLI tool. Just as with Terragrunt 1.0, most of the features are part of the open source
**Terragrunt 2.0 Free Edition**, but a few are part of **Terragrunt 2.0 Pro**, which requires a subscription (similar
to other developer tools, such as IntelliJ Community and IntelliJ Ultimate).

:::


| Feature           | Description                               | 1.0 | 2.0 Free | 2.0 Pro |
|-------------------|-------------------------------------------|-----|----------|---------|
| `source`          | DRY Terraform modules                     | ✅   | ✅        | ✅       |
| `remote_state`    | DRY backend configuration                 | ✅   | ✅        | ✅       |
| `extra_arguments` | DRY CLI arguments                         | ✅   | ✅        | ✅       |
| `run-all`         | Multi-module commands                     | ✅   | ✅        | ✅       |
| `dependency`      | Multi-module dependencies                 | ✅   | ✅        | ✅       |
| `hook`            | Before and after hooks                    | ✅   | ✅        | ✅       |
| `generate`        | Scaffold repos, modules, and environments | ❌   | ✅        | ✅       |
| `copy`            | Copy repos, modules, and environments     | ❌   | ✅        | ✅       |
| `test`            | Run automated tests                       | ❌   | ✅        | ✅       |
| `live`            | Manage live environments                  | ❌   | ❌        | ✅       |
| `catalog`         | Manage your module catalog                | ❌   | ❌        | ✅       |
| `update`          | Automatically update dependencies         | ❌   | ❌        | ✅       |

This guide mostly focuses on the automatic updates functionality in Terragrunt 2.0.

### Install Terragrunt 2.0

For macOS users, the easiest way to install Terragrunt is to use HomeBrew:

```
brew install terragrunt
```

For Windows users, the easiest option is to use Chocolatey:

```
choco install terragrunt
```

### Authenticate with Terragrunt

Some features in Terragrunt 2.0, such as automatic updates, require authentication to Gruntwork. To authenticate, run:

```
terragrunt auth
```

## See the current status of your infrastructure

Now that you're logged into Terragrunt, let's see the current state of your infrastructure.

### Check out your infrastructure-live repo

The first step is to check out the code for your infrastructure. Typically, this will be in an `infrastructure-live`
repo, which you'll want to check out to your computer using Git:

```
git clone "<URL OF YOUR INFRA LIVE REPO>"
```

Next, jump into the folder you just checked out:

```
cd infrastructure-live
```

### See what's deployed

To see what's deployed, run:

```
terragrunt live list
```

You should see a table that looks something like this:

```
Module Name   | dev     | stage    | prod    | (latest version)
vpc           | v1.1.0  | v1.1.0   | v1.1.0  | v1.1.0
eks           | v4.1.1* | v4.1.1*  | v4.1.1* | v5.0.0
rds           | v0.2.3  | v0.2.2*  | v0.2.2* | v0.2.3

* indicates an update is available
```

This table shows:

- Each of your environments (e.g., dev, stage, prod).
- The modules deployed (e.g., vpc, eks, rds).
- The version of each module deployed in each environment (e.g., v0.2.3, v0.2.2).
- The latest version available of each module (e.g., v0.2.3).

An asterisk next to any version number indicates a newer version is available and that you may want to update: in the
example above, you can see that the `eks` and `rds` modules need an update. The next several sections will walk you
through how to update these modules.

## Update a single module with the CLI

Now that you have seen which modules have newer versions available, let's update one of them to the latest version!

### Patches: going beyond version number bumps

Note that Terragrunt can update not only the version numbers of your modules, but it can also discover and apply
_patches_ which can automatically update other aspects of your code: e.g., rename a variable, write a new file, etc.
Under the hood, a patch is a YAML file that defines (a) what module and versions to target, (b) a series of shell
commands to execute, and (c) a Docker container in which to run those commands.

Any library maintainer can publish a patch for a specific release, and Terragrunt can automatically find these
patches, show them to you for review, and, with your approval, apply the patches as part of the update process. This is
useful for:

- **Backwards incompatible releases**: If a module maintainer made a backwards incompatible change in the latest
  release, such as renaming an input variable, that maintainer can create a patch that automatically updates your code
  to use the new input variable name. This makes upgrading across major versions much easier.

- **Rolling out new functionality**: Sometimes, you want to roll out a change across multiple repos, such as updating
  the copyright year, or adding a license file, or enabling a new feature for users of your library. You can create a
  patch file to automate this process.

For more information on creating patches, check out the <a>patch file format documentation</a>.

### The update command

To update a module to the latest version (including discovering and applying any patches for that version), you can use
the `update` command:

```
terragrunt update <ENV> [<MODULE>]
```

Where:

- `ENV`: the environment in which to update the module (e.g., dev, stage, prod). This corresponds to a folder name in
  your `infrastructure-live` repo.
- `MODULE`: the module to update in that environment (e.g., `eks`, `rds`). The `live list` command shows you all the
  module and environment names.

:::tip

By default, the `update` command will solely make modifications to your local code (e.g., your local checkout of the
`infrastructure-live` repo). It will **not** deploy any changes, and the command is idempotent, so it's always safe to
run it, as you'll always have a chance to review the changes before they affect anything.

:::

For example, to update the `eks` module in `dev`, you'd run:

```
terragrunt update dev eks
```

By default, the `update` command walks you through an interactive process where it performs the following steps:

1. **Find new versions**: Scan for dependencies and show if new versions are available.
1. **Prompt to update versions**: Interactively prompt the user whether to update to a new version.
1. **Find patches**: If you chose to update to a new version, show any patches available for that version.
1. **Prompt to apply patches**: Interactively prompt the user whether to apply a patch. You should always review the
   patch source code before applying to ensure the updates are applicable and safe for your environment.
1. **Test**: After updating the version and applying patches, run `terragrunt plan` as a quick test that everything is
   still working.

:::tip

After `update` has finished, try running `git status` and `git diff` to see the changes it made! You should see local
modifications with version numbers being modified, and, if you applied a patch, possibly other code changes.

:::

By default, the `update` command makes changes to the code you have checked out locally, so you'll need to commit those
changes as described in the next section.

### Committing the changes

If you're happy with the changes from the `update` command, commit and push them to your repository:

```
git add .
git commit -m "Update EKS in dev"
git push origin main
```

## Update an entire environment with the CLI

You can use the `update` command to automatically update not only a single module, but an entire environment (e.g.,
dev, stage, prod).

The recommended way to do this is to run:

```
terragrunt update \
  <ENV> \
  --pr \
  --branch-per-dependency
```

Where:

- `ENV` is the name of the environment to update (e.g., dev, stage, prod).
- `--pr` tells the CLI to automatically commit the changes and open a pull request (PR).
- `--branch-per-dependency` tells the CLI to update each dependency in a separate branch, and therefore, open a
  separate pull request for each one.

:::tip

With the `--pr` flag, the `update` command will solely make modifications in branches and open PRs with the changes. It
will **not** deploy any changes, and the command is idempotent, so it's always safe to run it, as you'll always have a
chance to review the changes before they affect anything.

:::

For example, to update the entire `dev` environment, run:

```
terragrunt update \
  dev \
  --pr \
  --branch-per-dependency
```

This will take you through the same interactive process as the `update` command you ran earlier—scanning for new
versions and patches and prompting you if you want to update to the new versions and apply the patches—except this
time:

- You'll go through this process multiple times, once for each dependency in the environment that needs to be updated.
- At the end, you'll have a PR opened for each dependency that was updated.

:::tip

When this `update` command finishes running, it'll print out a list of all the PRs that were opened. Remember to go
through each PR, review the changes, and if everything looks good, merge them in!

:::

## Configure automated upgrades in CI

Now that you've gone through the process of updating a single module and a whole environment using the CLI, you're in a
great place to learn how to automate it by adding it to your CI / CD pipeline on a regular schedule (similar to a
cron job). You can still use the CLI on a one-off basis whenever you want, but by running the `update` command on a
regular schedule in CI, you can ensure that your infrastructure is kept up to date automatically, on a regular cadence.

In a CI environment, you will typically run the `update` command with the following flags:

```
terragrunt update \
  --pr \
  --branch-per-dependency \
  --in-order <ENV1>[,<ENV2>,<ENV3>...] \
  --non-interactive \
  --patch-via-comments
```

Where:

- `--pr` tells the CLI to automatically commit the changes and open a pull request (PR).
- `--branch-per-dependency` tells the CLI to update each dependency in a separate branch, and therefore, open a
  separate pull request for each one.
- `--in-order <ENV1>[,<ENV2>,<ENV3>...]` tells the `update` command to open pull requests for one environment at a time,
  in the order specified. That is, initially, only open a PR for `ENV1`; only when that PR is merged, open another PR
  for `ENV2`; only when that PR is merged, opened another PR for `ENV3`; and so on. This allows you to ensure updates
  are "promoted" from environment to environment in the order you expect.
- `--non-interactive`: This ensures the CLI does not prompt you for any interactive inputs. When this flag is set, the
  `update` command, by default, will automatically update modules to the latest versions, but it will NOT apply patches
  for those new versions automatically. You can instead apply patches in the PR UI—only after reviewing the patch
  source code—as explained in the next flag.
- `--patch-via-comments`: Add a comment to each update PR that shows the list of patches available. Next to each patch
  is a checkbox that, if you check, will automatically apply that patch and update the PR with the results. This allows
  you to review and apply patches entirely from the GitHub PR UI.

:::tip

With the `--pr` flag, the `update` command will solely make modifications in branches and open PRs with the changes. It
will **not** deploy any changes, and the command is idempotent, so it's always safe to run it, as you'll always have a
chance to review the changes before they affect anything.

:::

The next several sections show you how to configure the `update` command to run on a regular schedule in various CI
servers.

### GitLab instructions

_Instructions coming soon!_

### GitHub Actions instructions

_Instructions coming soon!_

### CircleCI instructions

To configure CircleCI to run automated upgrades, edit your `.circleci/config.yml` file to add a new entry under `jobs`
called `auto-update`:

```yaml
jobs:
  auto-update:
    steps:
      - checkout
      - run: |
          terragrunt update \
            --pr \                            # Open PRs with updates
            --branch-per-dependency \         # Open 1 PR per dependency
            --in-order dev,stage,prod \       # Open 1 PR per env, in this order
            --non-interactive \               # Don’t prompt for input
            --patch-via-comments              # Use PR comments to select patches
```

:::tip

Make sure to update the `--in-order` flag with the names of the environments you wish to update!

:::


Next, tell CircleCI to run the `auto-update` job on a regular schedule by adding a new `auto-update` entry under
`workflows`:

```yaml
workflows:
  auto-update:
    jobs:
      - auto-update
    triggers:
      - schedule:
        cron: "*/5 * * * *" # every 5 min
        filters:
        branches:
        only:
          - main
```

This configures the `auto-update` job to run every 5 minutes. So, every 5 minutes, Terragrunt will check for updates,
and open PRs with any new versions and patches it finds.

Save the changes to the `.circleci/config.yml` file, commit and push it:

```
git add .
git commit -m "Configure auto update in CI"
git push origin main
```

That's it! Your CI / CD pipeline will now automatically update your infrastructure.

:::tip

Terragrunt will now open PRs automatically when there are new updates. **Remember to periodically check your inbox for
notifications** of these PRs, and to review and merge these in if the changes look good.

:::

## Success!

🎉 Congratulations! 🎉

You now understand how to keep your infrastructure up to date with Terragrunt!


<!-- ##DOCS-SOURCER-START
{"sourcePlugin":"local-copier","hash":"6fba8043573f4c7ba50d3bea57e348fc"}
##DOCS-SOURCER-END -->
