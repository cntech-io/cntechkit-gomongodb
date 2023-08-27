### mongodb helper methods for go projects

```go
mongodb := NewMongoDB().
    Connect().
    AttachCollection("first_collection").
    AttachCollection("second_collection")

mongodb.Do("first_collection").find()

mongodb.Disconnect()
```
