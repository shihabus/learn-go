In a go project, all the files in the root are part of the main package.
The `main` package will have a `main` function, which is the program's entry point.
While creating a new folder it creates a new package, with the folder name. 

We split our program into files and folders for easy management and scale. So the command that we usually use to run the go files is `go run file_name.go`. This command is unable to import external modules out of the box. So we have to initialize the go module using the command `go mod init module_name`. Module name would be the name of the project module. This creates a `go.mod` file with the module name and the minimum version of go required to run the project. It will also enlist external modules that we use in our project.

If you want to use function/values defined in other packages we have to import it. 
For instance, the module name be `myproject`.
Suppose you have a `getAge()` function in `/utils/util.go` file import will look like, `import "myproject/utils"`. To use it, you can use it as `utils.getAge()`.

All the functions and variables declared capitalized are public. Rest are private outside the package, and can only be accessed within the package. This includes files that are on the same level/in the same folder.
