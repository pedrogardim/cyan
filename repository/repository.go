package repository

import "go.mongodb.org/mongo-driver/mongo"

type RepositoryInterface interface {
	List()
	Get()
	Create()
	Delete()
	Update()
}

type Repository struct {
	entity string
	client *mongo.Client
	dbName string
}

func (r *Repository) List() {
	return
}

func (r *Repository) Get() {

}

func (r *Repository) Create() {

}

func (r *Repository) Delete() {

}

func (r *Repository) Update() {

}

func NewRepository(entity string, model interface{}, client *mongo.Client, dbName string) RepositoryInterface {

	return &Repository{
		entity: entity,
		client: client,
		dbName: dbName,
	}

}
