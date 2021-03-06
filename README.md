# Cloud Native Development (CND)

[![CircleCI](https://circleci.com/gh/okteto/cnd.svg?style=svg)](https://circleci.com/gh/okteto/cnd)

**Cloud Native Development** (CND) is about running your development flow entirely in kubernetes, avoiding the time-consuming `docker build/push/pull/redeploy` cycle. 

CND helps you achieve this with a mix of kubernetes automation, file synching between your local file system and kubernetes and hot reloading of containers.

## How does it work

This is how a standard dev environment looks like:

<img align="left" src="docs/env.png">

&nbsp;

And this how it looks after converting it into a cloud native environment:

<img align="left" src="docs/cnd.png">
&nbsp;

The **cnd** container duplicates the manifest of the **api** pod, so it is fully integrated with every Kubernetes feature.

Local changes are synched to the **cnd** container via `ksync`. As you save locally, it will be automatically synched in your **cnd** container in seconds.

Once you're ready to integrate, you can revert back to your original configuration for general end-to-end testing before sending a PR or pushing to production.


## Installation

The synching functionality of **cnd** is provided by [ksync](https://github.com/vapor-ware/ksync).

To install `ksync`, execute:

```bash
curl https://vapor-ware.github.io/gimme-that/gimme.sh | bash
```

and:

```
ksync init --image=vaporio/ksync:0.3.2-hotfix
```

check the `ksync` installation by executing:

```bash
ksync doctor
```

If `ksync` is successfully installed, install **cnd** from by executing:

```bash
go get github.com/okteto/cnd
```

## Usage

Note: these instructions assume that you already have a kubernetes-based application running. 

Define your Cloud Native Development file (`cnd.yml`). A `cnd.yml` looks like this:

```yaml
name: dev
swap:
  deployment:
    file: nginx-deployment.yml
    container: nginx
    image: ubuntu
  service:
    file: nginx-service.yml
mount:
  source: .
  target: /src
```

For more information about the Cloud Native Development file, see its [reference](/docs/cnd-file.md).

To convert your dev environment to a cloud native environment, execute:

```bash
cnd up
```

by default, it uses a `cnd.yml` in your current folder. For using a different file, execute:

```bash
cnd up -f path-to-cnd-file
```

To create a long-running session to your cloud native environment, execute:

```bash
cnd exec sh
```

You can also execute standalone commands like:

```bash
cnd exec go test
```

In order to revert back to your original configuration, execute:

```bash
cnd down
```

Switch back and forth your development mode as you wish by executing `cnd up` and `cnd down`. 
Finally, in order to permanently remove your cloud native development, execute:

```bash
cnd rm
```

For a full demo of Cloud Native Development, check the [Voting App demo](https://github.com/okteto/cnd-voting-demo).

## Troubleshooting

### Files are not syncing
cnd uses [ksync](https://github.com/vapor-ware/ksync) and [syncthing](https://docs.syncthing.net) to sync files between your environments. If your cloud native environment is not being updated correctly, review the following:

1. Verify that ksync is functioning correctly by running `ksync doctor`
1. Review ksync's logs (~/.ksync/daemon.log) for errors
1. Browse to http://localhost:8384 to see if there are any syncthing-specific issues.

### Files syncing is slow
Please follow [syncthing's docs](https://docs.syncthing.net/users/faq.html#why-is-the-sync-so-slow) to troubleshoot this.
