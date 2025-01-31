- **Go Workspaces**
[Go Workspaces](https://go.dev/doc/tutorial/workspaces)

## Create a module for your code

### 1. From the command prompt, create a directory for your code called workspace.
```
mkdir -p the-complete-developers-guide/workspace/hello-world
touch the-complete-developers-guide/workspace/hello.go
```

### 2. Initialize the module and add a dependency
```
go mod init example.com/hello
go get golang.org/x/example/hello/reverse
```

### 3. Run program
```
go run .
```

## Create the workspace

### 1. Initialize the workspace (in workspace dir)
```
go work init ./hello-world 
go run ./hello-world
```

## Adding another project

### 1. Clone the repository
```
git clone https://go.googlesource.com/example
```

### 2. Add the module to the workspace
```
go work use ./example/hello
```

### 3. Adding the code
```
code example/hello/reverse/int.go
code workspace/hello/hello.go
```

### 4. Run the code in the workspace
```
go run ./hello-world
```

## Additional
go work use [-r] [dir] adds a use directive to the go.work file for dir, if it exists, and removes the use directory if the argument directory doesn’t exist. The -r flag examines subdirectories of dir recursively.
go work edit edits the go.work file similarly to go mod edit
go work sync syncs dependencies from the workspace’s build list into each of the workspace modules.