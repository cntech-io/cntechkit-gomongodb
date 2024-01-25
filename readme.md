### install

```bash
go get github.com/cntech-io/cntechkit-gogin/v2
```

### Methods

| Method                                                           | Description                                    |
| ---------------------------------------------------------------- | ---------------------------------------------- |
| mongodb.NewMongoDB(enableLogger bool)                            | Creates mongodb instance                       |
| &nbsp;&nbsp;&nbsp;&nbsp;.Connect()                               | Connects to mongodb                            |
| &nbsp;&nbsp;&nbsp;&nbsp;.AttachCollection(collectionName string) | Attachs collection                             |
| &nbsp;&nbsp;&nbsp;&nbsp;.PushLog()                               | add logs if logger enabled                     |
| &nbsp;&nbsp;&nbsp;&nbsp;.Run()                                   | Runs server                                    |
| env.NewMongoDBEnv()                                              | Loads predefined mongodb environment variables |
