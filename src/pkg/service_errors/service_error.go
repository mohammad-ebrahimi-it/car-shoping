package service_errors

type ServiceError struct {
	EndUserMessage   string `json:"endUserMessage"`
	TechnicalMessage string `json:"technicalMessage"`
	Err              error  `json:"err"`
}

func (s *ServiceError) Error() string {
	return s.EndUserMessage
}
