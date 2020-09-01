# Introduction 
Framework for go application

# Getting Started
## Run the generator
```none
go-framework generate --name <application name> --repository <url to repo where app will be stored> --directory <directory into which to write project> --summary "summary of app" --description "description of app"
```
It is recommended that you perform a code cleanup to ensure the files conform to your project's guidelines.  Some files will contain extraneous newlines due to the templating library used ins generation (this will not affect function)

## Upgrading existing projects
The generator will read some settings from the config file in order to create a framework closer to your existing application.  This can be used to generate fuller fromeworks but is primarily intended to allow upgrades to improeved versions of the framework without rewriting all of your code)
### Config file name and location
Config file should be named 
```none
.go-framework.<type>
```
where the type is json, yaml, etc which indicates the format of the file
### Config file structure 
#### Example
```json
{
   "commands": [
     "cmd1",
     "cmd2",
     "cmd3"
   ],
   "settings": [
     {
       "name": "String1",
       "description": "This is setting 1 (string, shortcut)",
       "type": "string",
       "cmdline": "setting1",
       "cmdlineShortcut": "a",
       "appliesTo": [
         "root",
         "cmd1"
       ]
     },
     {
       "name": "String2",
       "description": "This is setting 2 (string, no-shortcut)",
       "type": "string",
       "cmdline": "setting2",
       "appliesTo": [
         "cmd2",
         "cmd3"
       ]
     }
}
``` 
commands is a simple list of commands as text that will be created<br>
settings is a list of settings objects as follows: <br>
| attribute       | Meaning                                                                                                                         |
|-----------------|---------------------------------------------------------------------------------------------------------------------------------|
| name            | Name of setting, will be used as the function name in Setting interface                                                         |
| description     | Description, this will be used in comments and for the usage information                                                        |
| type            | Type of variable, currently supported are [string, []string, bool, int]                                                         |
| cmdline         | Long cmd line string, can be set on command line with --<cmdline>, if not present then setting can only be set from config file |
| cmdlineShortcut | Single char cmdline,  can be set on command line with -<cmdlineShortcut>, if cmdline is not present, this is ignored            |
| appliesTo       | List of commands to which this commandline is applied, if cmdline is not present, this is ignored                               |

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