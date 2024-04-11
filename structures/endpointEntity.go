package structures

import (
	"errors"
	"time"
)

type QueryParameterEntity struct {
	Name  string `json:"name"`
	Value string `json:"value"`
}

type HeaderEntity struct {
	Name  string `json:"name"`
	Value string `json:"value"`
}

type EndpointEntity struct {
	Id              string                 `json:"id"`
	ProjectId       string                 `json:"projectId"`
	Name            string                 `json:"name"`
	Url             string                 `json:"url"`
	Path            string                 `json:"path"`
	Method          string                 `json:"method"`
	Body            string                 `json:"body"`
	QueryParameters []QueryParameterEntity `json:"queryParameters"`
	Headers         []HeaderEntity         `json:"headers"`

	CreatedAt time.Time
	UpdatedAt time.Time
}

func NewEndpointEntity(projectId *string, method *string) *EndpointEntity {
	entity := new(EndpointEntity)
	entity.ProjectId = *projectId

	entity.Headers = make([]HeaderEntity, 0)
	entity.QueryParameters = make([]QueryParameterEntity, 0)
	entity.Method = *method
	entity.CreatedAt = time.Now()
	entity.UpdatedAt = time.Now()
	entity.Name = "New endpoint"

	return entity
}

func (entity *EndpointEntity) ToMap() map[string]interface{} {
	data := make(map[string]interface{})

	data["projectId"] = entity.ProjectId
	data["name"] = entity.Name
	data["url"] = entity.Url
	data["path"] = entity.Path
	data["method"] = entity.Method
	data["body"] = entity.Body
	data["createdAt"] = entity.CreatedAt
	data["updatedAt"] = entity.UpdatedAt

	queryParameters := make([]interface{}, 0)
	headers := make([]interface{}, 0)

	for _, parameter := range entity.QueryParameters {
		parameterData := make(map[string]interface{})

		parameterData["name"] = parameter.Name
		parameterData["value"] = parameter.Value

		queryParameters = append(queryParameters, parameterData)
	}

	for _, header := range entity.Headers {
		headerData := make(map[string]interface{})

		headerData["name"] = header.Name
		headerData["value"] = header.Value

		headers = append(headers, headerData)
	}

	data["queryParameters"] = queryParameters
	data["headers"] = headers

	return data
}

func (entity *EndpointEntity) FromMap(id string, data map[string]interface{}) error {
	entity.Id = id
	entity.ProjectId = data["projectId"].(string)
	entity.Name = data["name"].(string)
	entity.Url = data["url"].(string)
	entity.Path = data["path"].(string)
	entity.Method = data["method"].(string)
	entity.Body = data["body"].(string)
	entity.UpdatedAt = data["updatedAt"].(time.Time)
	entity.CreatedAt = data["createdAt"].(time.Time)
	entity.QueryParameters = make([]QueryParameterEntity, 0)
	entity.Headers = make([]HeaderEntity, 0)

	headersData, headersOk := data["headers"].([]map[string]interface{})
	queryParametersData, queryParametersOk := data["queryParameters"].([]map[string]interface{})

	if !headersOk {
		if len(headersData) != 0 {
			return errors.New("No headers")
		}

	}

	if !queryParametersOk {
		if len(queryParametersData) != 0 {
			return errors.New("No queryParameters")
		}
	}

	for _, headerData := range headersData {
		header := new(HeaderEntity)

		header.Name = headerData["name"].(string)
		header.Value = headerData["value"].(string)

		entity.Headers = append(entity.Headers, *header)
	}

	for _, queryParametersData := range queryParametersData {
		queryParameter := new(QueryParameterEntity)

		queryParameter.Name = queryParametersData["name"].(string)
		queryParameter.Value = queryParametersData["value"].(string)

		entity.QueryParameters = append(entity.QueryParameters, *queryParameter)
	}

	return nil
}

func (entity *EndpointEntity) UpdateFromBody(body *UpdateBodyDTO) {
	entity.Name = body.Name
	entity.Url = body.Url
	entity.Path = body.Path
	entity.Method = body.Method
	entity.Body = body.Body
	entity.UpdatedAt = time.Now()

	entity.Headers = body.Headers
	entity.QueryParameters = body.QueryParameters
}
