package things

import (
	"bytes"
	"database/sql"
	"fmt"
	"path/filepath"
	"text/template"
	"time"

	"github.com/crockeo/dotfiles/cmd/migrate-to-obsidian/util"
)

// Things database schema as of 2023-06-01,
// this may have changed if you're reading this from
// ~ the future ~.
type Area struct {
	Uuid         string `sql:"uuid"`
	Title        string `sql:"title"`
	Visible      *int   `sql:"visible"`
	Index        int    `sql:"index"`
	CachedTags   []byte `sql:"cachedTags"`
	Experimental []byte `sql:"experimental"`
}

func GetAreas(conn *sql.DB) (map[string]*Area, error) {
	rows, err := conn.Query("SELECT * FROM TMArea")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	areas := map[string]*Area{}
	for rows.Next() {
		area := &Area{}
		err := rows.Scan(
			&area.Uuid,
			&area.Title,
			&area.Visible,
			&area.Index,
			&area.CachedTags,
			&area.Experimental,
		)
		if err != nil {
			return nil, err
		}
		areas[area.Uuid] = area
	}
	return areas, nil
}

type Task struct {
	Uuid                             string   `sql:"uuid"`
	LeavesTombstone                  int      `sql:"leavesTombstone"`
	CreationDate                     float32  `sql:"creationDate"`
	UserModificationDate             *float32  `sql:"userModificationDate"`
	Type                             int      `sql:"type"`
	Status                           int      `sql:"status"`
	StopDate                         *float32 `sql:"stopDate"`
	Trashed                          int      `sql:"trashed"`
	Title                            string   `sql:"title"`
	Notes                            string   `sql:"notes"`
	NotesSync                        int      `sql:"notesSync"`
	CachedTags                       []byte   `sql:"cachedTags"`
	Start                            int      `sql:"start"`
	StartDate                        *int     `sql:"startDate"`
	StartBucket                      int      `sql:"startBucket"`
	ReminderTime                     *int     `sql:"reminderTime"`
	LastReminderInteractionDate      *float32 `sql:"lastReminderInteractionDate"`
	Deadline                         *int     `sql:"deadline"`
	DeadlineSuppressionDate          *int     `sql:"deadlineSuppressionDate"`
	T2_deadlineOffset                int      `sql:"t2_deadlineOffset"`
	Index                            int      `sql:"index"`
	TodayIndex                       int      `sql:"todayIndex"`
	TodayIndexReferenceDate          *int     `sql:"todayIndexReferenceDate"`
	Area                             *string  `sql:"area"`
	Project                          *string  `sql:"project"`
	Heading                          *string  `sql:"heading"`
	Contact                          *string  `sql:"contact"`
	UntrashedLeafActionsCount        int      `sql:"untrashedLeafActionsCount"`
	OpenUntrashedLeafActionsCount    int      `sql:"openUntrashedLeafActionsCount"`
	ChecklistItemsCount              int      `sql:"checklistItemsCount"`
	OpenChecklistItemsCount          int      `sql:"openChecklistItemsCount"`
	Rt1_RepeatingTemplate            *string  `sql:"rt1_repeatingTemplate"`
	Rt1_RecurrenceRule               []byte   `sql:"rt1_recurrenceRule"`
	Rt1_InstanceCreationStartDate    *int     `sql:"rt1_instanceCreationStartDate"`
	Rt1_InstanceCreationPaused       *int     `sql:"rt1_instanceCreationPaused"`
	Rt1_InstanceCreationCount        *int     `sql:"rt1_instanceCreationCount"`
	Rt1_AfterCompletionReferenceDate *int     `sql:"rt1_afterCompletionReferenceDate"`
	Rt1_NextInstanceStartDate        *int     `sql:"rt1_nextInstanceStartDate"`
	Experimental                     []byte   `sql:"experimental"`
	Repeater                         []byte   `sql:"repeater"`
	RepeaterMigrationDate            *float32 `sql:"repeaterMigrationDate"`
}

func GetTasks(conn *sql.DB) (map[string]*Task, error) {
	rows, err := conn.Query("SELECT * FROM TMTask")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	tasks := map[string]*Task{}
	for rows.Next() {
		task := &Task{}
		err := rows.Scan(
			&task.Uuid,
			&task.LeavesTombstone,
			&task.CreationDate,
			&task.UserModificationDate,
			&task.Type,
			&task.Status,
			&task.StopDate,
			&task.Trashed,
			&task.Title,
			&task.Notes,
			&task.NotesSync,
			&task.CachedTags,
			&task.Start,
			&task.StartDate,
			&task.StartBucket,
			&task.ReminderTime,
			&task.LastReminderInteractionDate,
			&task.Deadline,
			&task.DeadlineSuppressionDate,
			&task.T2_deadlineOffset,
			&task.Index,
			&task.TodayIndex,
			&task.TodayIndexReferenceDate,
			&task.Area,
			&task.Project,
			&task.Heading,
			&task.Contact,
			&task.UntrashedLeafActionsCount,
			&task.OpenUntrashedLeafActionsCount,
			&task.ChecklistItemsCount,
			&task.OpenChecklistItemsCount,
			&task.Rt1_RepeatingTemplate,
			&task.Rt1_RecurrenceRule,
			&task.Rt1_InstanceCreationStartDate,
			&task.Rt1_InstanceCreationPaused,
			&task.Rt1_InstanceCreationCount,
			&task.Rt1_AfterCompletionReferenceDate,
			&task.Rt1_NextInstanceStartDate,
			&task.Experimental,
			&task.Repeater,
			&task.RepeaterMigrationDate,
		)
		if err != nil {
			return nil, err
		}
		tasks[task.Uuid] = task
	}
	return tasks, nil
}

func (t *Task) IsActive() bool {
	return t.StopDate == nil && t.Trashed == 0
}

func (t *Task) Hierarchy(areas map[string]*Area, tasks map[string]*Task) TaskHierarchy {
	var heading *Task
	var project *Task

	lookup := t
	if lookup.Heading != nil {
		heading = tasks[*lookup.Heading]
		lookup = heading
	}

	if lookup.Project != nil {
		project = tasks[*lookup.Project]
		lookup = project
	}

	var area *Area
	if lookup.Area != nil {
		area = areas[*lookup.Area]
	}

	return TaskHierarchy{
		area,
		project,
		heading,
	}
}

func (t *Task) ScheduledBlock() string {
	_ = time.ANSIC
	if t.StartDate == nil {
		return ""
	}

	// This is some cursed nonsense...
	//
	// Things uses a date format that I haven't seen before.
	// It is an integer where each addition of `128` is a day.
	// We know that `132604160` is today based on cross-referencing
	// Things startDate w/ its rendered scheduled dates.
	//
	// So we just:
	// - Find the # of days between today and the target date.
	// - Use Go's date library to do the day offset calculation correctly.
	// - Hope that this is actually a linear function :)
	dayRate := 128
	today := 132604160
	dayDiff := (*t.StartDate / dayRate) - (today / dayRate)
	todayDate := time.Date(2023, 6, 2, 0, 0, 0, 0, time.UTC)
	targetDate := todayDate.AddDate(0, 0, dayDiff)

	return fmt.Sprintf(
		"[scheduled:: %s]",
		targetDate.Format("2006-01-02"),
	)
}

func (t *Task) Render() ([]byte, error) {
	contents := bytes.Buffer{}
	if err := taskTemplate.Execute(&contents, t); err != nil {
		return nil, err
	}
	return contents.Bytes(), nil
}

var taskTemplate = template.Must(template.New("task").Parse(`
- [ ] #task {{ .Title }} {{ .ScheduledBlock }}

{{ .Notes }}
`))

type TaskHierarchy struct {
	Area    *Area
	Project *Task
	Heading *Task
}

func (t TaskHierarchy) Path() string {
	path := ""
	if t.Area != nil {
		path = filepath.Join(path, util.EscapePath(t.Area.Title))
	}
	if t.Project != nil {
		path = filepath.Join(path, util.EscapePath(t.Project.Title))
	}
	if t.Heading != nil {
		path = filepath.Join(path, util.EscapePath(t.Heading.Title))
	}
	return path
}

type Tag struct {
	Uuid         string  `sql:"uuid"`
	Title        string  `sql:"title"`
	Shortcut     string  `sql:"shortcut"`
	UsedDate     float32 `sql:"usedDate"`
	Parent       string  `sql:"parent"`
	Index        int     `sql:"index"`
	Experimental []byte  `sql:"experimental"`
}

func GetTags(conn *sql.DB) ([]Tag, error) {
	rows, err := conn.Query("SELECT * FROM TMTag")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	tags := []Tag{}
	for rows.Next() {
		tag := Tag{}
		err := rows.Scan(
			&tag.Uuid,
			&tag.Title,
			&tag.Shortcut,
			&tag.UsedDate,
			&tag.Parent,
			&tag.Index,
			&tag.Experimental,
		)
		if err != nil {
			return nil, err
		}
		tags = append(tags, tag)
	}
	return tags, nil
}

type AreaTag struct {
	Area string `sql:"areas"`
	Tags string `sql:"tags"`
}

func GetAreaTags(conn *sql.DB) ([]AreaTag, error) {
	rows, err := conn.Query("SELECT * FROM TMAreaTag")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	areaTags := []AreaTag{}
	for rows.Next() {
		areaTag := AreaTag{}
		err := rows.Scan(
			&areaTag.Area,
			&areaTag.Tags,
		)
		if err != nil {
			return nil, err
		}
		areaTags = append(areaTags, areaTag)
	}
	return areaTags, nil
}

type TaskTag struct {
	Task string `sql:"tasks"`
	Tags string `sql:"tags"`
}

func GetTaskTags(conn *sql.DB) ([]TaskTag, error) {
	rows, err := conn.Query("SELECT * FROM TMTaskTag")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	taskTags := []TaskTag{}
	for rows.Next() {
		taskTag := TaskTag{}
		err := rows.Scan(
			&taskTag.Task,
			&taskTag.Tags,
		)
		if err != nil {
			return nil, err
		}
		taskTags = append(taskTags, taskTag)
	}
	return taskTags, nil
}
