# go-rest-api

Go Rest API

```github.com/kanhaiya15/go-rest-api/
├── pkg/
|   ├── errors/
|   ├── log/
├── services/
|   ├── auth/
|   └── user/
├── api/ (OpenAPI/Swagger specs, JSON schema files, protocol definition files.)
|   ├── proto/  (protocol buffer files)
|   |   ├── v1/
|   |   |   ├── account.proto
|   |   |   ├── account.pb.go
|   |   └── v2/
|   └── rest/   (json files)
|       ├── v1/
|       |   └── account.json
|       └── v2/
├── configs/ (project config settings, default configs, file templates)
├── scripts/ (Scripts to perform various build, install, analysis, etc operations.)
├── build/ (Packaging and Continuous Integration.)
├── test / (system and module level tests)
├── docs/ (project documents folder)
├── examples/ (project examples for service interactions)
├── githooks/ (project git hooks)
├── assets/ (common assests for all services)
├── Makefile
├── README.md
└── docker-compose.yml
```
