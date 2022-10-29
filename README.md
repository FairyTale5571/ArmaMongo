# Arma Mongo

## The extension only works in write mode
### It was provided for logging only, but reads from the database can be scheduled if there is demand.

## Installation

#### Setup environment variables

```bash
LOG_SNOWFLAKE=your_snowflake
LOG_HOOK=your_webhook
MONGO_URL=mongodb://user:password@host:port
DATABASE_MONGO=database_name
```

#### In Arma:
    
Place RRPServer_Logger.pbo to Server Mods

For writing log use function:
```sqf
RRPServer_logger_write
```

Example:
```sqf
[
    "MyCollection", // Collection name   
    "LogType", // Log type
    "PlayerUID", // Player UID
    // To write fields, you must use an array with a key-value value, an example is below
    // Each key must have a value, otherwise it will not be written
    [
        ["key1", "value1"],
        ["money", 100000],
        ["vehicle", "BMW"]
    ]
] call RRPServer_logger_write
```

#### Result in MongoDB:
```bson
    {
      "_id": {
        "$oid": "635c6a5ff3093c2155757b48"
      },
      "log": {
        "money": "100000",
        "vehicle": "BMW",
        "key1": "value1"
      },
      "player_uid": "PlayerUID",
      "type_log": "LogType",
      "created_at": {
        "$numberLong": "1667000927349144500"
      }
    }
```