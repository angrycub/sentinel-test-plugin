# sentinel-test-plugin

This is a sample Sentinel plugin suitable for use with Nomad Enterprise. The
code is based on the [`Plugin Framework`][] documentation.

As of Nomad 1.7, the correct Sentinel SDK as the plugin base is [v0.3.13][].

## Required configuration

- Nomad Enterprise v1.7
- Go 1.21.4
- `make` (optional)

## Build the plugin

If you have `make` installed, you can build the binary by running `make`.
However, there is nothing `make`-specific about the build process and you
can also build it in the typical Go way.

```bash
mkdir -p bin
go build -o bin/sentinel-test-plugin ./...
```

## Install and use the plugin

Once built, you then need to install it on your Nomad Enterprise instances,
configure the servers to load it, upload a Sentinel policy that uses it, and
finally, run a test to observe it working.

### Install the plugin

Place the compiled plugin in a folder accessible to the user that the `nomad`
process is running as.

### Configure Nomad servers

Add the following configuration to each of your Nomad Enterprise servers.

```hcl
sentinel {
    import "sentinel-test-plugin" {
        path = "«full path to the compiled binary»"
    }
}
```

### Add a policy that uses the plugin

```hcl
# Test policy always fails for demonstration purposes

import "sentinel-test-plugin" as tp; # Since the plugin has a hyphenated name it must be aliased

print("The current minute is", tp.minute)

print(
  "Four months from now is",
  tp.add_month(4).string,
)

print(
  "Is four months from now January?",
  tp.add_month(4).string == "January",
)

main = rule { false }
```

### Plan a job

Finally, plan a Nomad job. I used the sample created by running
`nomad job init --short`.

```shell-session
$ nomad plan example.nomad.hcl
+ Job: "example"
+ Task Group: "cache" (1 create)
  + Task: "redis" (forces create)

Scheduler dry-run:
- All tasks successfully allocated.

Job Warnings:
1 warning:

* testPlugin : Result: false (allowed failure based on level)

Description:
  Test policy always fails for demonstration purposes

Print messages:

The current minute is 12
Four months from now is March
Is four months from now January? false

FALSE - testPlugin:17:1 - Rule "main"

Job Modify Index: 0
To submit the job with version verification run:

nomad job run -check-index 0 example.nomad.hcl

When running the job with the check-index flag, the job will only be run if the
job modify index given matches the server-side version. If the index has
changed, another user has modified the job and the plan's results are
potentially invalid.
```




<!-- Reference Links -->
[v0.3.13]: https://pkg.go.dev/github.com/hashicorp/sentinel-sdk@v0.3.13
[`Plugin Framework`]: https://docs.hashicorp.com/sentinel/extending/plugins#plugin-framework