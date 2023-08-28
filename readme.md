### mongodb helper methods for go projects

```go
mongodb := NewMongoDB(false). // if true, creates logs collection by default on Connect method
    Connect().
    AttachCollection("first_collection").
    AttachCollection("second_collection")

mongodb.Do("first_collection").find()

// if logs collection created, it pushes predefined log records to logs collection
mongodb.PushLog("app-name","log description") 

mongodb.Disconnect()
```
