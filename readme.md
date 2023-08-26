### mongodb helper methods for go projects

```go
mongodb := NewMongoDB().
    Connect().
    AttachCollection("fist_collection").
    AttachCollection("second_collection")

mongodb.Collection["first_collection"].find()

mongodb.Disconnect()
```
