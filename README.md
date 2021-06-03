[comment]: <> ( Generated 2021-06-03 14:15:48 by go-framework v1.17.0 )

# Go app framework generator

Tool for generating go applications

# Introduction

TODO: Give a short introduction of your project. Let this section explain the objectives, or the motivation behind this project.

# Getting Started

TODO: Guide users through getting your code up and running on their own system. In this section you can talk about:

1. Installation process
2. Software dependencies
3. Latest releases
4. API references

# Build and Test

TODO: Describe and show how to build your code and run the tests.

# Configuration

[comment]: <> ( SECTION-START: ConfigTable )

| Setting | Config file | Cmdline | Default Val | Applies to | Description |
| :------ | :---------- | :------ | :---------- | :--------- | :---------- |
| Single read configuration (bool) | Config.SingleReadConfig |  | true | all | Configuration is only read once (at startup) |
| Target directory (string) | Config.TargetDirectory | -d, --directory | . | generate | Target directory |
| Application title (string) | Application.Title | -t, --title |  | generate | Name of application |
| Application name (string) | Application.Name | -n, --name |  | generate | Name of application |
| Application summary (string) | Application.Summary | --summary |  | generate | Summary description of application |
| Application description (string) | Application.Description | --description |  | generate | Description of application |
| RepositoryPath (string) | Application.Repository | -r, --repository |  | generate | Path to repository |
| Commands ([]UserCommand) | Commands |  |  | generate | Commands to add |
| UserSettings ([]UserSetting) | UserSettings |  |  | generate | Settings to add |

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
