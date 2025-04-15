package response

const (
	ErrMethodNotAllowed      = "Метод не поддерживается"
	ErrTokenRequired         = "Требуется токен авторизации"
	ErrInvalidTokenFormat    = "Невалидный формат токена"
	ErrInvalidToken          = "Невалидный токен"
	ErrInvalidJSON           = "Невалидный JSON"
	ErrPostCreationFailed    = "Ошибка при создании поста"
	ErrPostRetrievalFailed   = "Ошибка при получении постов"
	ErrCommentCreationFailed = "Ошибка при создании комментария"
	ErrCommentsDisabled      = "Комментарии под этим постом отключены"
	ErrEncodResponse         = "Ошибка при кодировании ответа"
)
