package store

// Task job to be done or completed
type Task struct {
	ID    int    // identifier of the task
	Title string // Title of the task
	Done  bool   // If task is completed or not
}

// Datastore manages a list of task in memory
type Datastore struct {
	tasks  []Task
	lastID int // lastID is incremented for each new stored task
}

// GetPendingTasks returns all the tasks which need to be done
func (ds *Datastore) GetPendingTasks() []Task {
	var pendingTasks []Task
	for _, task := range ds.tasks {
		if !task.Done {
			pendingTasks = append(pendingTasks, task)
		}
	}
	return pendingTasks
}

// SaveTask should store the task in the datastore
func (ds *Datastore) SaveTask(task Task) {
	ds.lastID++
	task.ID = ds.lastID
	ds.tasks = append(ds.tasks, task)
}

// DeleteTask should remove the task from the datastore
func (ds *Datastore) DeleteTask(task int) bool {
	delete(ds, task)
}
