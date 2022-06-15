具体架构请参考文档：[Design](design.md)（[设计](design_cn.md)）

**ATTENTION: This project is a Work-in-Progress.**

**注意，目前尚在开发，暂时无法运行，仅供代码参考。**

## Alien Structure
```
.
├── api
│   ├── cluster
│   │   └── v1
│   ├── component
│   │   └── v1
│   ├── diagnosis
│   │   └── v1
│   ├── metadata
│   │   └── v1
│   ├── migration
│   │   └── v1
│   ├── scale
│   │   └── v1
│   └── upgrade
│       └── v1
├── cmd
│   └── multicluster
├── configs
├── deploy
│   ├── build
│   ├── docker-compose
│   └── kubernetes
├── docs
├── example
├── internal
│   ├── biz
│   │   └── constant
│   ├── conf
│   ├── data
│   │   └── ent
│   │       ├── cluster
│   │       ├── component
│   │       ├── enttest
│   │       ├── hook
│   │       ├── migrate
│   │       ├── predicate
│   │       ├── runtime
│   │       └── schema
│   ├── errors
│   ├── middleware
│   │   ├── auth
│   │   ├── logging
│   │   └── metric
│   ├── server
│   └── service
└── third_party
    ├── errors
    ├── google
    │   ├── api
    │   └── protobuf
    ├── page
    └── validate

```
