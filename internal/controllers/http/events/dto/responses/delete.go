package responses

import "github.com/aevula/interview-hustlers-calendar/internal/controllers/http/httputils"

type DeleteEventResponse struct {
}

func ToDeleteEventResponse(_ httputils.EmptyResponse) DeleteEventResponse {
	return DeleteEventResponse{}
}
