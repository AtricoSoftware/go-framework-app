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
TODO 