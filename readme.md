# INFO
- Baseline for the back-end part of the aluminium garbage richard project 🙈
- Examples regarding needs of the project will be here 😏

# How to mock
- Make sure you install mockgen globally
```
go install github.com/golang/mock/mockgen@latest
```
- Generate the mock and the rest is ok (down there is an example)
```
mockgen --source hello_world_repository.go --destination ../mocks/hello_world_repository_mock.go --package mocks
```

