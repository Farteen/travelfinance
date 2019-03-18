package event

const (
	MongoEventCollection = "events"
)

const (
	EventCreationFailedErr = iota + 1000
)

const (
	EventCreationFailedErrMsg = "创建事件失败"
	EventMongoQueryBatchSize = 10
)