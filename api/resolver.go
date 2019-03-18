package api

import (
	"github.com/projects/go-graphql-jobs/graph"
	"context"
	"fmt"
	"time"
	"github.com/99designs/gqlgen/graphql"
	"github.com/projects/go-graphql-jobs/dal"
	"github.com/projects/go-graphql-jobs/models"
) // THIS CODE IS A STARTING POINT ONLY. IT WILL NOT BE UPDATED WITH SCHEMA CHANGES.
//Resolver is a thing
type Resolver struct{}

//Application is a thinkg
func (r *Resolver) Application() ApplicationResolver {
	return &applicationResolver{r}
}

//Job is a thing
func (r *Resolver) Job() graph.jobResolver{
	return &jobResolver{r}
}

//Mutation is a thing
func (r *Resolver) Mutation() graph.mutationResolver {
	return &mutationResolver{r}
}

//Query is a thing
func (r *Resolver) Query() graph.queryResolver {
	return &queryResolver{r}
}

type applicationResolver struct{ *Resolver }

func (r *applicationResolver) CreatedAt(ctx context.Context, obj *Application) (string, error) {
	panic("not implemented")
}

type jobResolver struct{ *Resolver }

func (r *jobResolver) CreatedAt(ctx context.Context, obj *Job) (string, error) {
	panic("not implemented")
}
func (r *jobResolver) DeletedAt(ctx context.Context, obj *Job) (*string, error) {
	panic("not implemented")
}

type mutationResolver struct{ *Resolver }

func (r *mutationResolver) CreateJob(ctx context.Context, models models.NewJob) (dal.Job, error) {
	jobRepository, err := dal.NewJobFirebaseRepository(r.db.Conn, r.db.Context)
	if err != nil {
		fmt.Printf("firebase error: ", err)
		return dal.Job{}, err
	}

	var jobID string
	if jobID, err = jobRepository.GetID(); err != nil {
		fmt.Printf("jobRepository GetID error: ", err)
		return dal.Job{}, err
	}
	// Create job object from request
	job := dal.Job{
		ID:          jobID,
		Name:        models.Name
		Description: models.Description
		Location:    models.Location
		CreatedBy:   models.CreatedBy
		CreatedAt:   time.Now().UTC(),
	}

	// Set the values in the DB
	if err = jobRepository.Insert(job); err != nil {
		fmt.Printf("firebase error: ", err)
		return dal.Job{}, err
	}

	return job, nil
}
func (r *mutationResolver) DeleteJob(ctx context.Context, id string) (string, error) {
	panic("not implemented")
}
func (r *mutationResolver) CreateApplication(ctx context.Context, models NewApplication) (*Application, error) {
	panic("not implemented")
}

type queryResolver struct{ *Resolver }

func (r *queryResolver) Jobs(ctx context.Context) ([]dal.Job, error) {
	var allJobs []dal.Job

	jobRepository, err := dal.NewJobFirebaseRepository(r.db.Conn, r.db.Context)

	if allJobs, err = jobRepository.GetAll(); err != nil {
		fmt.Printf("firebase error", err)
	}

	return allJobs, nil
}
func (r *queryResolver) Applications(ctx context.Context, jobID string) ([]Application, error) {
	panic("not implemented")
}
