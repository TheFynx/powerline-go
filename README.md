# powerline-go

A beautiful and useful low-latency prompt for your shell, written in go.

## Features

- Easy to extend and customize
- Cross-platform support
- Fast execution (written in Go)

## Segment Configuration

### Git (`git`)

The git segment has several display modes and configuration options:

```
-git-mode string
    How to display git status (default "fancy")
    - "fancy": Shows detailed stats in separate segments
    - "compact": Shows stats in a condensed format
    - "simple": Shows minimal status indicators

-git-disable-stats string
    Comma-separated list to disable individual git statuses:
    - ahead: Commits ahead of remote
    - behind: Commits behind remote
    - staged: Staged changes
    - notStaged: Unstaged changes
    - untracked: Untracked files
    - conflicted: Merge conflicts
    - stashed: Stashed changes

-git-assume-unchanged-size int
    Disable checking for changes in git repositories where the index is larger
    than this size (in KB) to improve performance (default 2048)

-ignore-repos string
    A list of git repos to ignore (comma-separated)
    Repos are identified by their root directory
```

### Kubernetes (`kube`)

The Kubernetes segment offers name shortening options:

```
-shorten-eks-names
    Shortens names for EKS Kube clusters
    Example: "arn:aws:eks:region:account:cluster/cluster-name" -> "cluster-name"

-shorten-gke-names
    Shortens names for GKE Kube clusters
    Example: "gke_project_zone_cluster-name" -> "cluster-name"

-shorten-openshift-names
    Shortens names for OpenShift clusters
    Example: "namespace/api-url:port/user" -> "api-url"
```

### Duration (`duration`)

The duration segment shows command execution time with configurable precision:

```
-duration string
    The elapsed clock-time of the previous command

-duration-min string
    The minimal time a command has to take before the duration segment is shown
    (default "0")

-duration-low-precision
    Display duration with lower precision
    Example: "1.234ms" -> "1ms"
```

### Current Working Directory (`cwd`)

The cwd segment has multiple display modes and options:

```
-cwd-mode string
    How to display the current directory (default "fancy")
    - "fancy": Full path with shortcuts
    - "semifancy": Abbreviated path
    - "plain": Simple path
    - "dironly": Current directory only

-cwd-max-depth int
    Maximum number of directories to show in path (default 5)

-cwd-max-dir-size int
    Maximum number of letters displayed for each directory in the path
    (default -1)

-path-aliases string
    One or more aliases from a path to a short name
    Format: "foo/bar/baz=FBB,other/path=OP"
```

### Docker Context (`docker-context`)

Shows the current Docker context, hiding the default context:

- Reads from `DOCKER_CONTEXT` environment variable
- Falls back to context from `~/.docker/config.json`
- Doesn't show anything when using the "default" context

### Virtual Environment (`venv`)

Python virtual environment display options:

```
-venv-name-size-limit int
    Show indicator instead of virtualenv name if name is longer than this limit
    (default 0 = unlimited)
```

### Shell Variable (`shell-var`)

Custom shell variable display options:

```
-shell-var string
    A shell variable to add to the segments

-shell-var-no-warn-empty
    Disables warning for empty shell variable
```

### Host (`host`)

Hostname display options:

```
-colorize-hostname
    Colorize the hostname based on a hash of itself

-hostname-only-if-ssh
    Show hostname only for SSH connections

-alternate-ssh-icon
    Show the older, original icon for SSH connections
```

### AWS (`aws`)

Shows AWS profile and region information:

- Reads from environment variables:
  - `AWS_PROFILE` or `AWS_VAULT` for profile name
  - `AWS_DEFAULT_REGION` for region
- Displays in format: `profile (region)`
- Hidden when no profile is set

### Node.js (`node`)

Shows Node.js version and package information:

- Displays Node.js version when `node` is available
- Shows package version from `package.json` when in a Node.js project
- Uses custom indicators configurable via theme

### Time (`time`)

Shows current time with custom format:

Custom time format (uses Go's time formatting) Example: "15:04:05" for 24-hour time

```bash
-time string
    Custom time format (uses Go's time formatting)
    Example: "15:04:05" for 24-hour time
```

### Load (`load`)

Shows system load average with configurable thresholds:

- Displays the 5-minute load average by default
- Changes background color when load exceeds threshold
- Configurable via theme:
  - `LoadAvgValue`: Load average period (1, 5, or 15 minutes)
  - `LoadThresholdBad`: Threshold relative to CPU count for high load warning
  - `LoadHighBg`: Background color for high load
  - `LoadBg`: Normal background color
  - `LoadFg`: Foreground color

## General Configuration

```
-modules string
    The list of modules to load, separated by ','
    (default "venv,user,host,ssh,cwd,perms,git,hg,jobs,exit,root")

-modules-right string
    The list of modules to load anchored to the right
    (requires -eval mode)

-theme string
    Set this to the theme you want to use
    (default, low-contrast, gruvbox, solarized-dark16, solarized-light16)

-shell string
    Set this to your shell type
    (autodetect, bare, bash, zsh)

-priority string
    Segments sorted by priority, if not enough space exists
```

## Shell Configuration

### Bash

```bash
function _update_ps1() {
    PS1="$($GOPATH/bin/powerline-go -error $? -jobs $(jobs -p | wc -l))"
}

if [ "$TERM" != "linux" ] && [ -f "$GOPATH/bin/powerline-go" ]; then
    PROMPT_COMMAND="_update_ps1; $PROMPT_COMMAND"
fi
```

### ZSH

```bash
function powerline_precmd() {
    PS1="$($GOPATH/bin/powerline-go -error $? -jobs ${${(%):%j}:-0})"
}

function install_powerline_precmd() {
    for s in "${precmd_functions[@]}"; do
        if [ "$s" = "powerline_precmd" ]; then
            return
        fi
    done
    precmd_functions+=(powerline_precmd)
}

if [ "$TERM" != "linux" ]; then
    install_powerline_precmd
fi
```

### Fish

```fish
function fish_prompt
    eval $GOPATH/bin/powerline-go -error $status -jobs (count (jobs -p))
end
```

## Nix Shell Configuration

When using nix-shell --pure, add this to your .bashrc:

```bash
if [ "$IN_NIX_SHELL" == "pure" ]; then
    if [ -x "$HOME/.nix-profile/bin/powerline-go" ]; then
        alias powerline-go="$HOME/.nix-profile/bin/powerline-go"
    fi
fi
```

## Acknowledgments

- Original powerline-go by [@justjanne](https://github.com/justjanne)
- Based on Powerline-Shell by [@banga](https://github.com/banga)
- Inspired by [Powerline](https://github.com/Lokaltog/vim-powerline)
