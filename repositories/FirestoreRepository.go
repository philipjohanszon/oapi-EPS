package repositories

import (
	"cloud.google.com/go/firestore"
	"context"
	"endpointsService/structures"
	"errors"
	"github.com/google/uuid"
	"google.golang.org/api/iterator"
)

type EndpointRepository struct {
	Client *firestore.Client
}

type QueryParameter struct {
	Path    string
	Operand string
	Value   interface{}
}

func NewQueryParameter(path string, operand string, value interface{}) QueryParameter {
	param := QueryParameter{Path: path, Operand: operand, Value: value}

	return param
}

const CollectionName = "endpoints"
const IdServicePrefix = "EPS:"

func (repository EndpointRepository) Get(ctx context.Context, parameters []QueryParameter) ([]*structures.EndpointEntity, error) {
	query := repository.Client.Collection(CollectionName).Query

	for _, parameter := range parameters {
		query = query.Where(parameter.Path, parameter.Operand, parameter.Value)
	}

	documents := query.Documents(ctx)

	endpoints := make([]*structures.EndpointEntity, 0)

	for {
		document, err := documents.Next()

		if errors.Is(err, iterator.Done) {
			break
		}

		if err != nil {
			return nil, err
		}

		endpoint := new(structures.EndpointEntity)

		data := document.Data()

		convertErr := endpoint.FromMap(document.Ref.ID, data)

		if convertErr != nil {
			return nil, convertErr
		}

		endpoints = append(endpoints, endpoint)
	}

	return endpoints, nil
}

func (repository EndpointRepository) GetById(ctx context.Context, id string) (*structures.EndpointEntity, error) {
	document, err := repository.Client.Collection(CollectionName).Doc(id).Get(ctx)

	if err != nil {
		return nil, err
	}

	endpoint := new(structures.EndpointEntity)

	convertErr := endpoint.FromMap(id, document.Data())

	if convertErr != nil {
		return nil, convertErr
	}

	return endpoint, nil
}

func (repository EndpointRepository) Create(ctx context.Context, body *structures.CreateBody) (*structures.EndpointEntity, error) {
	endpoint := structures.NewEndpointEntity(&body.ProjectId, &body.Method)

	id := IdServicePrefix + uuid.New().String()

	_, err := repository.Client.Collection(CollectionName).Doc(id).Set(ctx, endpoint.ToMap())

	if err != nil {
		return nil, err
	}

	return endpoint, nil
}

func (repository EndpointRepository) Delete(ctx context.Context, id string) error {
	_, err := repository.Client.Collection(CollectionName).Doc(id).Delete(ctx)

	if err != nil {
		return err
	}

	return nil
}

func (repository EndpointRepository) Update(ctx context.Context, endpoint *structures.EndpointEntity) error {
	data := endpoint.ToMap()

	_, err := repository.Client.Collection(CollectionName).Doc(endpoint.Id).Set(ctx, data)

	if err != nil {
		return err
	}

	return nil
}
