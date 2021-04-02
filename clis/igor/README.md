@gpa-lab/igor
=============



[![oclif](https://img.shields.io/badge/cli-oclif-brightgreen.svg)](https://oclif.io)
[![Version](https://img.shields.io/npm/v/@gpa-lab/igor.svg)](https://npmjs.org/package/@gpa-lab/igor)
[![Downloads/week](https://img.shields.io/npm/dw/@gpa-lab/igor.svg)](https://npmjs.org/package/@gpa-lab/igor)
[![License](https://img.shields.io/npm/l/@gpa-lab/igor.svg)](https://github.com/IIP-Design/igor/blob/master/package.json)

<!-- toc -->
* [Usage](#usage)
* [Commands](#commands)
<!-- tocstop -->
# Usage
<!-- usage -->
```sh-session
$ npm install -g @gpa-lab/igor
$ igor COMMAND
running command...
$ igor (-v|--version|version)
@gpa-lab/igor/0.0.0 darwin-x64 node-v12.21.0
$ igor --help [COMMAND]
USAGE
  $ igor COMMAND
...
```
<!-- usagestop -->
# Commands
<!-- commands -->
* [`igor hello [FILE]`](#igor-hello-file)
* [`igor help [COMMAND]`](#igor-help-command)

## `igor hello [FILE]`

describe the command here

```
USAGE
  $ igor hello [FILE]

OPTIONS
  -f, --force
  -h, --help       show CLI help
  -n, --name=name  name to print

EXAMPLE
  $ igor hello
  hello world from ./src/hello.ts!
```

_See code: [src/commands/hello.ts](https://github.com/IIP-Design/igor/blob/v0.0.0/src/commands/hello.ts)_

## `igor help [COMMAND]`

display help for igor

```
USAGE
  $ igor help [COMMAND]

ARGUMENTS
  COMMAND  command to show help for

OPTIONS
  --all  see all commands in CLI
```

_See code: [@oclif/plugin-help](https://github.com/oclif/plugin-help/blob/v3.2.2/src/commands/help.ts)_
<!-- commandsstop -->
