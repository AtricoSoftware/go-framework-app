# Introduction 
Framework for go application

# Getting Started
1. Copy this code to a new folder
1. Rename the project
1. Change the module url as appropriate (go.mod)
1. Change the import statements to match new module name
    * replace all 'ma.ci.go-framework-app' with new repo name (Make sure you include all files, not just code)
    * replace all 'go-framework-app' with new app name (Make sure you include all files, not just code)
1. Update pkg/Info details
1. Check build.sh values (MODULE and OUTPUT)

# Features
## Commands
cobra is used for commands<br>
Root command is implemented idiomatically<br>
### Version command
'version' is implemented, by default this simply displays the version<br>
 using the '--full' option displays more information (see [Build details](#builddetails))

## <a name="builddetails"/> Build details
The version command can display full application details, these are read from the pkg.BuildDetails variable which is a JSON string.  This is set using ldflags as can be seen in build.sh<br>
#### example

```bash
> go-framework-app version --full

go-framework-app
Here is the describtion
v1.0.0

Details
-------
Built:
  Built on: Fri Jul 31 12:15:59 BST 2020
  Built by: rob
Git:
  Repository: MODULE
  Branch: master
  Commit: e916cd3dccf2a7b8226f3fa2b19e5b4a66761d63
```

## <a name="pipeline"/> Build pipeline
There is a sample pipeline.yml file for building your app on ADO<br>
The build script will set the build details and build for each platform (windows, Mac and linux)
The apps will then be zipped and pushed to artifactory

## <a name="settings"/> Settings
Settings are captured in an interface in settings.Settings - This allows them to be injected in testing<br>
For normal use (ie not testing) use
```go
settings := settings.GetSettings()
```
Settings can be added as both a command line option and a config file entry with a single function call.  Viper is used to read the config file
The config file is read from the following locations (the first successfully read file will be used)
1. User specified file if present on commandline (-- config file)
1. File present in current directory
1. File present in user home dir
The file will be name .app-name.EXT where EXT is json, yaml, yml or ini (the extension determines the format of the file)
If user specified, the file may have any name but the extension is used as above

## Links
[cobra (commandline parsing)](https://github.com/spf13/cobra)
[viper (config files)](https://github.com/spf13/viper)