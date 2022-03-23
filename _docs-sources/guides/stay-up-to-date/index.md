# Update Guide

This guide will show you how to use Terragrunt Pro to keep your infrastructure up to date.

## Terragrunt Pro

### About Terragrunt Pro

Terragrunt Pro is a full-stack framework for Terraform that ships with all the tools you need to implement the full
Terraform lifecycle, including creating modules, writing automated tests, managing versions, managing environments,
automatically updating dependencies, and much more.

### Install Terragrunt Pro

Install Terragrunt using any of the methods at the Terragrunt Installation page.

For macOS users, the most common approach is:

```
brew install terragrunt
```

For Windows users, the most common approach is:

```
choco install terragrunt
```

### Authenticate with Terragrunt Pro

Some features in Terragrunt Pro, such as automatic updates, require authentication to Gruntwork. To authenticate, run:

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

### A note on _patching_

Note that Terragrunt Pro can update not only the version numbers of your modules, but it can also discover and apply
_patches_ which can automatically update other aspects of your code. This can be especially useful for:

- **Backwards incompatible releases**: If a module maintainer made a backwards incompatible change in the latest
  release, such as renaming an input variable, that maintainer can create a patch that automatically updates your code
  to use the new input variable name. This makes upgrading across major versions much easier.

- **Rolling out new functionality**: Sometimes, you want to roll out a change across multiple repos, such as updating
  the copyright year or adding a license file. You can create a patch file to do this.

### The update command

To update a module, you can use the `update` command:

```
terragrunt update <ENV> [<MODULE>]
```

Where:

- `ENV`: the environment in which to update the module (e.g., dev, stage, prod). This corresponds to a folder name in
  your `infrastructure-live` repo.
- `MODULE`: the module to update in that environment (e.g., `eks`, `rds`). The `live list` command shows you all the
  module and environment names.

For example, to update the `eks` module in `dev`, you'd run:

```
terragrunt update dev eks
```

By default, the `update` command walks you through an interactive process where it:

- Scans for dependencies.
- Looks for new versions of the dependencies it found.
- If new versions are found, prompts you if you want to update to the new versions.
- If you choose to update to a new version, looks for patches for that new version, and prompts you if you want to
  apply one of the patches. NOTE: Patches will be executed in a Docker container on your computer, so make sure to read
  the patch source code carefully to be sure you're comfortable with it!
- After updating the version and applying patches, runs `terragrunt plan` as a quick check that everything is still
  working.
- Runs `git diff` to show you the code changes that resulted from the update.

### Committing the changes

If you're happy with the changes from the `update` command, commit and push them to your repository:

```
git add .
git commit -m “Update EKS in dev”
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
  --dependency-per-branch
```

Where:

- `ENV` is the name of the environment to update (e.g., dev, stage, prod).
- `--pr` tells the CLI to automatically commit the changes and open a pull request (PR).
- `--dependency-per-branch` tells the CLI to update each dependency in a separate branch, and therefore, open a
  separate pull request for each one.

For example, to update the entire `dev` environment, run:

```
terragrunt update \
  dev \
  --pr \
  --dependency-per-branch
```

This will take you through the same interactive process as the `update` command you ran earlier—scanning for new
versions and patches and prompting you if you want to update to the new versions and apply the patches—except this
time:

- You'll go through this process multiple times, once for each dependency in the environment that needs to be updated.
- At the end, you'll have a PR opened for each dependency that was updated. You'll want to go through each PR, review
  the changes, and if everything looks good, merge them in.

## Configure automated upgrades in CI

Now that you've gone through the process of updating a single module and a whole environment using the CLI, you're in a
great place to learn how to automate it by adding it to your CI / CD pipeline on an automated schedule (similar to a
cron job). You can still use the CLI on a one-off basis whenever you want, but by running the `update` command on a
regular schedule, you can ensure that your infrastructure is kept up to date automatically, on a regular cadence.

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
          --in-order <ENV>[,<ENV>...] \     # Open 1 PR per env, in this order
          --pr \                            # Open PRs with updates
          --branch-per-dependency \         # Open 1 PR per dependency
          --non-interactive \               # Don’t prompt for input
          --patch-via-comments              # Use PR comments to select patches
```

You'll notice it's the same `update` command again, with just a handful of new flags:

- `--in-order`: This flags, combined with `--pr`, tells the `update` command to open pull requests for one environment
  at a time, in the order specified. For example, if you used `--in-order dev,stage,prod`, and a new version of the
  `eks` module came out, the `update` command would first open a PR to update it in the `dev` environment; only after
  that PR has been merged, next time you run `update`, it will open a new PR to update `eks` in the `stage` environment;
  and when that PR has been merged, the next time you run `update`, it will update `eks` in the `prod` environment.
  This allows you to ensure updates are "promoted" from environment to environment in the order you expect.
- `--non-interactive`: This ensures the CLI does not prompt you for any interactive inputs. The default behavior is to
  update the code to any new versions that are available, but NOT to apply patches for those new versions automatically
  (see the next flag for how patches are handled in a non-interactive setting).
- `--patch-via-comments`: If a PR is opened to bump to a new version number, and a patch is available for that version,
  the `update` command will add a comment to the PR with information about that patch, and the ability to click a
  checkbox in that comment to apply the patch. This allows you to preview patches and apply the ones you want all from
  the GitHub PR UI.

Now, to use this new `auto-update` job, add an `auto-update` entry under `workflows`:

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
git commit -m “Configure auto update in CI”
git push origin main
```

That's it! Your CI / CD pipeline will now automatically update your infrastructure. Check your `infrastructure-live`
repo and inbox periodically for notifications about new PRs, and when things look good, merge them in.

## Success!

🎉 Congratulations! 🎉

You now understand how to keep your infrastructure up to date with Terragrunt Pro!
