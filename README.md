# dodo daemon plugin

Adds support for long-running dodo backdrops ("daemons").

The idea of this plugin is a bit silly, because it does exactly the same as
docker-compose, except with less features. The only reason this exists is to
take advantage of the rest of the dodo ecosystem (e.g. using builders or other
plugins) for long-running services.

## usage

```
$ dodo daemon --help
run backdrops in daemon mode

Usage:
  dodo daemon [command]

Available Commands:
  restart     restart a daemon backdrop
  start       run a backdrop in daemon mode
  stop        stop a daemon backdrop

Flags:
  -h, --help   help for daemon

Use "dodo daemon [command] --help" for more information about a command.
```

## installation

Install this plugin by downloading the correct file for your system from the
[releases page](https://github.com/wabenet/dodo-daemon/releases),
then copy it into the dodo plugin directory (`${HOME}/.dodo/plugins`).

Alternatively, if you want to compile your own dodo distribution, you can add
this plugin with the following generate config:

```yaml
plugins:
  - import: github.com/wabenet/dodo-daemon/pkg/plugin
```

## license & authors

```text
Copyright 2021 Ole Claussen

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
```
