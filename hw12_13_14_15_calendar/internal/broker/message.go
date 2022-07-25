package broker

//easyjson:json
type Event struct {
	// UUID - уникальный идентификатор события.
	UUID string `json:"uuid"`
	// Заголовок - короткий текст.
	Summary string `json:"summary"`
	// Дата и время начала события.
	StartedAt int32 `json:"startedAt"`
	// Дата и время начала события.
	FinishedAt int32 `json:"finishedAt"`
	// Описание события - длинный текст, опционально.
	Description string `json:"description"`
	// UUID пользователя, владельца события.
	UserUUID string `json:"userUuid"`
	// Дата и время уведомления о событии.
	NotificationAt int32 `json:"notificationAt"`
}
