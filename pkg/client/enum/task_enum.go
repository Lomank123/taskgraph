package enum

type TaskStatus string

const (
	TaskStatusPending   TaskStatus = "pending"
	TaskStatusRunning   TaskStatus = "running"
	TaskStatusSuccess   TaskStatus = "success"
	TaskStatusFailed    TaskStatus = "failed"
	TaskStatusBlocked   TaskStatus = "blocked"
	TaskStatusCancelled TaskStatus = "cancelled"
	TaskStatusTimeout   TaskStatus = "timeout"
)

var AllTaskStatuses = []TaskStatus{
	TaskStatusPending,
	TaskStatusRunning,
	TaskStatusSuccess,
	TaskStatusFailed,
	TaskStatusCancelled,
	TaskStatusTimeout,
	TaskStatusBlocked,
}

type TaskType string

const (
	TaskTypeHttp TaskType = "http"
	TaskTypeMq   TaskType = "mq"
	TaskTypeScript TaskType = "script"
)

var AllTaskTypes = []TaskType{
	TaskTypeHttp,
	TaskTypeMq,
	TaskTypeScript,
}
