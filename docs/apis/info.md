# Info API Methods

Info API returns basic info about the node

The API provides the following functions and endpoints:

* [/info](#getinfo)


Client lib APIs:
* [Info()](#getinfo)



##  `/info`

Returns basic info about the node.


### Parameters
None.

### Examples

#### cURL

```shell
curl --location 'http://localhost:8080/info'
```

#### Client lib

Information of a node can be retrieved via `Info() (*jsonmodels.InfoResponse, error)`
```go
info, err := goshimAPI.Info()
if err != nil {
    // return error
}

// will print the response
fmt.Println(string(info))
```

#### Response examples
```json
{
  "version": "v0.6.2",
  "networkVersion": 30,
  "tangleTime": {
    "messageID": "6ndfmfogpH9H8C9X9Fbb7Jmuf8RJHQgSjsHNPdKUUhoJ",
    "time": 1621879864032595415,
    "synced": true
  },
  "identityID": "D9SPFofAGhA5V9QRDngc1E8qG9bTrnATmpZMdoyRiBoW",
  "identityIDShort": "XBgY5DsUPng",
  "publicKey": "9DB3j9cWYSuEEtkvanrzqkzCQMdH1FGv3TawJdVbDxkd",
  "solidMessageCount": 74088,
  "totalMessageCount": 74088,
  "enabledPlugins": [
    "Activity",
    "Analysis-Client",
    "Autopeering",
    "Banner",
    "CLI",
    "Clock",
    "Config",
    "Consensus",
    "DRNG",
    "Dashboard",
    "Database",
    "Gossip",
    "Graceful Shutdown",
    "Logger",
    "Mana",
    "ManaRefresher",
    "Manualpeering",
    "MessageLayer",
    "Metrics",
    "NetworkDelay",
    "PoW",
    "PortCheck",
    "Profiling",
    "Prometheus",
    "RemoteLog",
    "RemoteLogMetrics",
    "WebAPI",
    "WebAPI DRNG Endpoint",
    "WebAPI Mana Endpoint",
    "WebAPI WeightProvider Endpoint",
    "WebAPI autopeering Endpoint",
    "WebAPI data Endpoint",
    "WebAPI faucet Endpoint",
    "WebAPI healthz Endpoint",
    "WebAPI info Endpoint",
    "WebAPI ledgerstate Endpoint",
    "WebAPI message Endpoint",
    "WebAPI tools Endpoint",
    "snapshot"
  ],
  "disabledPlugins": [
    "Analysis-Dashboard",
    "Analysis-Server",
    "Faucet",
    "ManaEventLogger",
    "Spammer",
    "TXStream"
  ],
  "mana": {
    "access": 1,
    "accessTimestamp": "2021-05-24T20:11:05.451224937+02:00",
    "consensus": 10439991680906,
    "consensusTimestamp": "2021-05-24T20:11:05.451228137+02:00"
  },
  "manaDelegationAddress": "1HMQic52dz3xLY2aeDXcDhX53LgbsHghdfD8eGXR1qVHy",
  "mana_decay": 0.00003209,
  "scheduler": {
    "running": true,
    "rate": "5ms",
    "nodeQueueSizes": {}
  },
  "rateSetter": {
    "rate": 20000,
    "size": 0
  }
}
```

#### Results
|Return field | Type | Description|
|:-----|:------|:------|
| `version`  | `String` | Version of GoShimmer. |
| `networkVersion`  | `uint32` | Network Version of the autopeering. |
| `tangleTime`  | `TangleTime` | TangleTime sync status |
| `identityID`  | `string` | Identity ID of the node encoded in base58. |
| `identityIDShort`  | `string` | Identity ID of the node encoded in base58 and truncated to its first 8 bytes. |
| `publicKey`  | `string` | Public key of the node encoded in base58 |
| `messageRequestQueueSize`  | `int` | The number of messages a node is trying to request from neighbors. |
| `solidMessageCount`  | `int` | The number of solid messages in the node's database. |
| `totalMessageCount`  | `int` | The number of messages in the node's database. |
| `enabledPlugins`  | `[]string` | List of enabled plugins. |
| `disabledPlugins`  | `[]string` | List if disabled plugins. |
| `mana`  | `Mana` | Mana values. |
| `manaDelegationAddress`  | `string` | Mana Delegation Address. |
| `mana_decay`  | `float64` | The decay coefficient of `bm2`. |
| `scheduler`  | `Scheduler` |  Scheduler is the scheduler used.|
| `rateSetter`  | `RateSetter` | RateSetter is the rate setter used. |
| `error` | `string` | Error message. Omitted if success.     |