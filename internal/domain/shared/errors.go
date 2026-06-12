package shared

import "errors"

var ErrEventNotFound error = errors.New("event not found")
var ErrEventBadRequest error = errors.New("bad request")
var ErrInvalidEventId error = errors.New("invalid event id")
var ErrInvalidEventName error = errors.New("invalid event name")
var ErrInvalidPageNumber error = errors.New("invalid page number")
var ErrInvalidPageSize error = errors.New("invalid page size")

var ErrQueueMessageNotFound error = errors.New("QueueMessage not found")
