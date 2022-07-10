package internalhttp

//easyjson:json
type Event struct {
	// UUID - уникальный идентификатор события.
	UUID string `json:"uuid"`
	// Заголовок - короткий текст.
	Summary string `json:"summary"`
	// Дата и время начала события.
	StartedAt int `json:"startedAt"`
	// Дата и время начала события.
	FinishedAt int `json:"finishedAt"`
	// Описание события - длинный текст, опционально.
	Description string `json:"description"`
	// UUID пользователя, владельца события.
	UserUUID string `json:"userUuid"`
	// Дата и время уведомления о событии.
	NotificationAt int `json:"notificationAt"`
}

//easyjson:json
type ListOfEvents []Event

//easyjson:json
type EventUpdateAttributes struct {
	// Заголовок - короткий текст.
	Summary string `json:"summary"`
	// Unix timestamp даты и времени начала события.
	StartedAt int `json:"startedAt"`
	// Unix timestamp даты и времени завершения события.
	FinishedAt int `json:"finishedAt"`
	// Описание события - длинный текст, опционально.
	Description string `json:"description"`
	// UUID пользователя, владельца события.
	UserUUID string `json:"userUuid"`
	// Unix timestamp даты и времени уведомления о событии.
	NotificationAt int `json:"notificationAt"`
}
