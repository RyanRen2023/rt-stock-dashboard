# CST 8916 Assignment 1

## Real-time Application with REST, GraphQL, and WebSockets**

## Team Members and Responsibilities
| Member | Responsibilities |
|--------|-------------|
| **Shaoxian Duan** | Section 2 and section 3 draft,REST API Implementation |
| **Xihai Ren** | Section 1 and Section 3 draft, GraphQL Real-time Communication & Frontend Development,Wrap the whole project |

## Requirments

- [Assignment Requirements](docs/Requirements.md)

## Real-time Stock Monitoring Services
- **Stocks Server**: [stocks-server](stocks-server/README.md)
  - The **REST API server** that provides stock market data retrieval.  
  - Handles **data fetching** from external stock APIs.  
- **GraphQL Server**: [graphql-server](graphql-server/README.md)
  - Exposes a **GraphQL API** for querying stock market data.  
  - Allows clients to request **only the specified data fields**.  
  - Supports **GraphQL Subscriptions** via WebSockets for real-time updates.  
- **Stock Dashboard**: [stock-dashboard](stock-dashboard/README.md)
  - Provides a **user-friendly UI** to search for stocks and view real-time price changes.  

### Run the service 

```sh
   docker compose up -d --build
```

## Submistion 

1. [CST8916 Assignment Report](docs/Report.md)
2. ![Video Presentation]()
3. [Persentation doc](docs/Presentation.pdf)

