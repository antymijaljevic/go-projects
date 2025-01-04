[Full Golang Tutorial - Learn Go by Building a TodoList App](https://www.youtube.com/watch?v=XCZWyN9ZbEQ)


- **Installing Go on MacOS**

1. Download package
[MacOS Arm](https://go.dev/dl/)

2. Verify
```bash
go version
```

- **Download Visual Studio Code Extensions**
GO by Go Team at Googlego.dev
TabOut by Albert Romkes

- **Block unused imports auto-remove**
1. Open VS code settings.json


2. Add this block to the config
```json
{
  ...
  "[go]": {
    "editor.codeActionsOnSave": {
        "source.organizeImports": false
    }
}
```

