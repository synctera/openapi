# OpenAPI

This repository contains shared OpenAPI schemas (as well as our standard errorcodes configuration).

A recent survey found that 6/10 people find this diagram helpful:

```mermaid
graph LR
O1(openapi) -->|openapi/v*/common/*.yml| M(mainapi)
O1 -->|openapi/v*/common/*.yml| W(webhook)
O1 -->|openapi/v*/common/*.yml| C(cards)
O1 -->|openapi/v*/common/*.yml| X(foo-service)
M -->|CI:openapi trigger| O2(openapi)
W -->|CI:openapi trigger| O2(openapi)
C -->|CI:openapi trigger| O2(openapi)
X -->|CI:openapi trigger| O2(openapi)
O2 -->|CI:spec| ME(external-api-merged-bundled.yml)
```

## Also see

- [REST API Guidelines](https://gitlab.com/synctera/architecture/-/blob/main/api/REST-API-Guidelines.md)
- [OpenAPI tooling at Synctera](https://gitlab.com/synctera/architecture/-/tree/main/development/practices/openapi)