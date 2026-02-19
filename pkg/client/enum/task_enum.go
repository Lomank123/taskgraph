package enum

type TaskStatus string

const (
	TaskStatusPending   TaskStatus = "pending"
	TaskStatusRunning   TaskStatus = "running"
	TaskStatusSuccess   TaskStatus = "success"
	TaskStatusFailed    TaskStatus = "failed"
	TaskStatusCancelled TaskStatus = "cancelled"
	TaskStatusTimeout   TaskStatus = "timeout"
	// TODO: Not sure if we need this
	// TaskStatusRetrying  TaskStatus = "retrying"
)

var AllTaskStatuses = []TaskStatus{
	TaskStatusPending,
	TaskStatusRunning,
	TaskStatusSuccess,
	TaskStatusFailed,
	TaskStatusCancelled,
	TaskStatusTimeout,
}

type TaskType string

const (
	TaskTypeHttp TaskType = "http"
	TaskTypeMq   TaskType = "mq"
)

var AllTaskTypes = []TaskType{
	TaskTypeHttp,
	TaskTypeMq,
}
