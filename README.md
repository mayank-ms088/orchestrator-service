# orchestrator-service

## Services

### 1. Orchestrator service Interface:

```
    // User Schema which is returned by the service.
    type User struct {
        name string
        class string
        roll int64
    }

    // Request Schema which is given input to the service.
    type RPCRequest struct {
        name string
    }

```

```
    // Function which is called by the client
    // This Function makes forwards the request to the other orchestrator instance exposed at port 9001
    func (server *OrchestratorServer) GetUserByName(
        ctx context.Context,
        req *orchestrator.RPCRequest
    )(*orchestrator.User,error)

```

```
   // Function which is called by the orchestrator 1
    // This Function makes forwards the request to the Dummy Data Service exposed at port 9001
    func (server *OrchestratorServer) GetUser(
        ctx context.Context,
        req *orchestrator.RPCRequest
    )(*orchestrator.User,error)

```

### 2. Dummy Data service Interface:

```
    //Function That returns the Hard Coded data object of type user.
    //Hosted at port 10000
    func (server *DummyDataServer) GetMockUserData(
	    ctx context.Context,
	    req *datamock.RPCRequest,
    )(*datamock.User,error)

```
