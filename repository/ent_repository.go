package repository

import (
	"context"
	"fmt"
	"log"
	"os"
	"todogolist/ent"
	"todogolist/ent/migrate"
	"todogolist/ent/task"
	"todogolist/models"

	//Add this module to be able to connect to postgreSQL.
	_ "github.com/lib/pq"
)

type entApp struct {
	db *ent.Client
}

func entRepositoryDomainWrapper(repositoryTasks []*ent.Task) []models.TaskDomainResponse {
	tasksResponse := []models.TaskDomainResponse{}
	for _, task := range repositoryTasks {
		taskInstance := models.TaskDomainResponse{}
		taskInstance.TaskID = fmt.Sprint(task.ID)
		taskInstance.TaskName = task.TaskName
		taskInstance.TaskDate = task.TaskDate
		if task.Periodicity != nil {
			taskInstance.Periodicity = *task.Periodicity
		}
		taskInstance.RecurrentTask = task.RecurrentTask
		taskInstance.TaskLabel = task.TaskLabel

		if task.TaskTimeInvested != nil {
			taskInstance.TaskTimeInvested = *task.TaskTimeInvested
		}

		tasksResponse = append(tasksResponse, taskInstance)

	}

	return tasksResponse

}

func (e *entApp) Setup(dryRun bool) (string, error) {
	ctx := context.Background()
	defer e.db.Close()

	switch dryRun {
	case true:

		if err := e.db.Schema.WriteTo(ctx, os.Stdout); err != nil {
			return "", nil
		}
		return "", nil

	default:
		if err := e.db.Schema.Create(
			ctx,
			migrate.WithDropIndex(true),
			migrate.WithDropColumn(true),
		); err != nil {
			return "", err
		}
		return "DB Setup Succesfully", nil

	}

}

func (e *entApp) GetTask(taskFilterDomain models.TaskFilterDomain) ([]models.TaskDomainResponse, error) {
	ctx := context.Background()
	defer e.db.Close()

	queryBuilder := e.db.Task.Query()

	if taskFilterDomain.TaskName != "" {

		queryBuilder.Where(task.TaskName(taskFilterDomain.TaskName))
	}

	if !taskFilterDomain.StartDate.IsZero() {
		queryBuilder.Where(task.TaskDateGTE(taskFilterDomain.StartDate))

	}

	if !taskFilterDomain.EndDate.IsZero() {

		queryBuilder.Where(task.TaskDateLTE(taskFilterDomain.EndDate))
	}

	if taskFilterDomain.TaskLabel != "" {
		queryBuilder.Where(task.TaskLabel(taskFilterDomain.TaskLabel))
	}

	if taskFilterDomain.Limit != 0 {
		queryBuilder.Limit(taskFilterDomain.Limit)
	}

	tasks, err := queryBuilder.Order(ent.Desc("task_date")).All(ctx)

	if err != nil {
		return nil, err
	}

	return entRepositoryDomainWrapper(tasks), nil

}

func (e *entApp) CreateTask(taskDomains []models.TaskDomain, debugMode bool) (string, error) {
	ctx := context.Background()
	defer e.db.Close()
	bulk := make([]*ent.TaskCreate, len(taskDomains))

	for i, taskDomain := range taskDomains {
		bulk[i] = e.db.Task.Create().
			SetTaskName(taskDomain.TaskName).
			SetTaskLabel(taskDomain.TaskLabel).
			SetTaskDate(taskDomain.TaskDate).
			SetPeriodicity(taskDomain.Periodicity).
			SetTaskTimeInvested(taskDomain.TaskTimeInvested)

		if taskDomain.Periodicity != "" {
			bulk[i].SetRecurrentTask(true)
		}

	}

	taskCreation, err := e.db.Task.CreateBulk(bulk...).Save(ctx)

	if err != nil {
		return "", err
	}

	if debugMode == true {
		return fmt.Sprint("Task Created %s", taskCreation), nil
	}

	return "Task Created Succesfully", nil
}

func (e *entApp) ListTask(taskFilterListDomain models.TaskFilterListDomain) ([]models.TaskDomainResponse, error) {
	ctx := context.Background()
	defer e.db.Close()

	tasks, err := e.db.Task.Query().Order(ent.Desc("task_date")).Limit(taskFilterListDomain.Limit).All(ctx)

	if err != nil {
		return nil, err
	}

	return entRepositoryDomainWrapper(tasks), nil

}

func (e *entApp) DeleteTask(taskDeleteDomain models.TaskDeleteDomain) (string, error) {
	ctx := context.Background()
	defer e.db.Close()

	_, err := e.db.Task.Delete().Where(task.TaskName(taskDeleteDomain.TaskName)).Exec(ctx)

	if err != nil {
		return "", err
	}

	return "TaskDeletedSuccesfully", nil

}

func newEntRepository(configpath string) Repository {

	config, err := loadConfiguration(configpath)

	if err != nil {
		log.Fatalf("Unable to proccess repository configuration %s", err)
	}

	client, err := ent.Open(config.DbDialect, config.DbURI)

	if err != nil {
		log.Fatalf("Unable to create client connection err %s", err)
	}

	return &entApp{db: client} // nil

}
