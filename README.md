[comment]: <> ( Generated 2021-06-23 15:07:34 by go-framework v1.21.0 )

# Go app framework generator

Tool for generating go applications

# Introduction

A utility to generate a framework (command line) app in Go, so we don't have to code all the useful boiler-plate each time.

# Commands

## generate

This will create or modify an app framework according to the config. The new app will have a standard structure with commands, options, and a variety of
boiler-plate code.<br>
#### Libraries used in generated code

| Library | Description |
| :------ | :---------- |
| [cobra](https://github.com/spf13/cobra) | commands |
| [viper](https://github.com/spf13/viper) | configuration |
| [container](https://github.com/atrico-go/container) | IoC container |
| [testing](https://github.com/atrico-go/testing) | Testing utilities |

#### Recommendations

* Any code changes that can be made with a change to the framework config (e.g. adding commands, option, etc) should be done so and the framework regenerated.
* Bespoke code should be placed within the api package, in a sub package if extensive
* Very limited changes should be made to generated files and only outside commented blocks (see below)

#### Comments

###### Start of file comment

```go
// Generated <date>> by go-framework <version>
```

Generated files all start with a comment that identifies them as generated.

###### Block comments

```go
// SECTION-START: <name>
....
// SECTION-END
```

Block comments mark sections of the file that will be overwritten by the framework if the app is regenerated. Code within these blocks should not be modified.
Any generated file with no block comments will be ***entirely*** overwritten. Removing a commented section will prevent it from being written in the future (not
recommended)<br>
***note*** - If a file does not exist, the generated file may have code outside of commented blocks.  This is to provide a stub for the user to work with.  Should the generator change, these won't be guaranteed to compile (as they will not be overwritten to preserve user code)  The changes required should be obvious enough.

## list skeletons

This will list the available skeletons. This will include those built in and any read from files as specified in the configuration.

# Configuration

Configuration can be read from a variety of file formats (json, yaml, etc) See [viper library](https://github.com/spf13/viper)
<br>
The following table shows the configuration options available. Some options can be configured from the commandline and/or environment variables and all can be
configured from the config file. The default config file is called .go-framework.<ext> where <ext> is json,yaml, etc and determines how it will be parsed.

[comment]: <> ( SECTION-START: ConfigTable )

| Setting | Config file | Cmdline | Default Val | Applies to | Description |
| :------ | :---------- | :------ | :---------- | :--------- | :---------- |
| Single read configuration (bool) | Config.SingleReadConfig |  | true | generate | Configuration is only read once (at startup) |
| Target directory (string) | Config.TargetDirectory | -d, --directory | . | generate | Target directory |
| Application title (string) | Application.Title | -t, --title |  | generate | Name of application |
| Application name (string) | Application.Name | -n, --name |  | generate | Name of application |
| Application summary (string) | Application.Summary | --summary |  | generate | Summary description of application |
| Application description (string) | Application.Description | --description |  | generate | Description of application |
| RepositoryPath (string) | Application.Repository | -r, --repository |  | generate | Path to repository |
| Commands ([]UserCommand) | Commands |  |  | generate | Commands to add |
| UserSettings ([]UserSetting) | UserSettings |  |  | generate | Settings to add |
| Skeleton Files ([]string) | SkeletonFiles | -s, --skeleton |  | generate, list/skeletons | File(s) with skeleton definitions |
| ConfigFile (string) | ConfigFile | --config-file |  | all | Alternate config file |
| Verbose (bool) | Verbose | -v, --verbose |  | all | Generate more detailed output |

<details>
  <summary>Further details</summary>
Config file ids with a dot can be "scoped"<br>
e.g. "a.b.c" can be added to config file as:<br>

### yaml

```yaml
a:
  b:
    c: "value"
```

### json

```json
{
  "a": {
    "b": {
      "c": "value"
    }
  }
}
```

</details>

[comment]: <> ( SECTION-END )

## UserCommand

```json
    {
  "Name": "<name of command>",
  "Description": "<description>",
  "NoImplementation": "<true/false>",
  "Args": [
    "<mandatory cmdline arg>",
    "mandatory cmdline arg"
  ],
  "OptionalArgs": [
    "optional cmdline arg"
  ]
}
```

If this is a sub command, the name should be in the form "base/sub"<br>
"pure" base commands should have NoImplementation set to true

## UserSetting

```json
{
  "Skeleton": "<skeleton on which to base this setting>",
  "Name": "<human readable name>",
  "Id": "<scoped id in config file>",
  "Description": "<description>>",
  "Type": "<variable type - Custom types are allowed>",
  "DefaultVal": "<default value if not specified>",
  "Cmdline": "<long cmd line option, e.g. 'abc' will give an option of '--abc'>",
  "CmdlineShortcut": "<short cmdline option, e.g. 'a' will give option '-a'>",
  "EnvVar": "<environment variable that can be used>",
  "AppliesTo": [
    "<command that this option applies to>"
  ]
}
```

If Skeleton is set, the skeleton values are used by default. Any other values set will override the skeleton.<br>
To apply the setting to all commands, AppliesTo can be omitted

## Skeleton file

```json
{
  "<name>": {
    <as
    for
    userSetting>
  }
}
```

The "Skeleton" field in UserSetting is ignored in this case. It is recommended that "AppliesTo" is left blank