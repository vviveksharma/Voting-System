# Voting-System

Hashing Based Voting System

## Specs

For Reading the specs copy the specs.yml to Swagger Editor to see the full yaml

## Flow Diagram

- Image for the User Flow

  ![Alt Text](/flow-images/User-Flow.png)

- BlockChain Flow
  ![Alt Text](/flow-images/Voting%20System%20-%20Block%20Chain%20Flow.jpg)

- Admin Flow 
    ![Alt Text](/flow-images/Voting%20System%20-%20Admin%20Flow.jpeg)

### For complete Flows and Diagrm access to this Miro Board

- https://miro.com/app/board/uXjVN3Lvs34=/?share_link_id=743498886238


## How to run this on your local

- Install Docker on your Local
  - `Make compose-build` (To build your golang build)
  - `Make compose-with-debug` (To run the backend in the log format).
  - `Make compose-without-app` (To run the backend in the Debug mode).
  - `Make compose-up` (To run the backend in foreground).