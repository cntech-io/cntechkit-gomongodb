### mongodb helper methods for go projects

```go
mongodb := NewMongoDB().
    Connect().
    AttachCollection("fist_collection").
    AttachCollection("second_collection")

mongodb.Do("first_collection").find()

mongodb.Disconnect()
```
