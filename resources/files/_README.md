{"Type":"Mixed"}
[comment]: <> ( {{.Comment}} )

# {{.ApplicationTitle}}

{{.ApplicationDescription}}

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
{{- range .UserSettings}}
| {{.Name}} ({{.Type}}) | {{.Id}} | {{if (ne .CmdlineShortcut "")}}-{{.CmdlineShortcut}}, {{end}}{{if (ne .Cmdline "")}}--{{.Cmdline}}{{end}} | {{.DefaultVal}} | {{.AppliesToCSL}} | {{.Description}} |
{{- end}}

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

