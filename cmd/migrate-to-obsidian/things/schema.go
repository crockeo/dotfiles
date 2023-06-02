package things

// Things database schema as of 2023-06-01,
// this may have changed if you're reading this from
// ~ the future ~.
type Area struct {
	Uuid         string `sql:"uuid"`
	Title        string `sql:"title"`
	Visible      int    `sql:"visible"`
	Index        int    `sql:"index"`
	CachedTags   []byte `sql:"cachedTags"`
	Experimental []byte `sql:"experimental"`
}

type Task struct {
	Uuid                             string  `sql:"uuid"`
	LeavesTombstone                  int     `sql:"leavesTombstone"`
	CreationDate                     float32 `sql:"creationDate"`
	UserModificationDate             float32 `sql:"userModificationDate"`
	Type                             int     `sql:"type"`
	Status                           int     `sql:"status"`
	StopDate                         float32 `sql:"stopDate"`
	Trashed                          int     `sql:"trashed"`
	Title                            string  `sql:"title"`
	Notes                            string  `sql:"notes"`
	NotesSync                        int     `sql:"notesSync"`
	CachedTags                       []byte  `sql:"cachedTags"`
	Start                            int     `sql:"start"`
	StartDate                        int     `sql:"startDate"`
	StartBucket                      int     `sql:"startBucket"`
	ReminderTime                     int     `sql:"reminderTime"`
	LastReminderInteractionDate      float32 `sql:"lastReminderInteractionDate"`
	Deadline                         int     `sql:"deadline"`
	DeadlineSuppressionDate          int     `sql:"deadlineSuppressionDate"`
	T2_deadlineOffset                int     `sql:"t2_deadlineOffset"`
	Index                            int     `sql:"index"`
	TodayIndex                       int     `sql:"todayIndex"`
	TodayIndexReferenceDate          int     `sql:"todayIndexReferenceDate"`
	Area                             string  `sql:"area"`
	Project                          string  `sql:"project"`
	Heading                          string  `sql:"heading"`
	Contact                          string  `sql:"contact"`
	UntrashedLeafActionsCount        int     `sql:"untrashedLeafActionsCount"`
	OpenUntrashedLeafActionsCount    int     `sql:"openUntrashedLeafActionsCount"`
	ChecklistItemsCount              int     `sql:"checklistItemsCount"`
	OpenChecklistItemsCount          int     `sql:"openChecklistItemsCount"`
	Rt1_repeatingTemplate            string  `sql:"rt1_repeatingTemplate"`
	Rt1_recurrenceRule               []byte  `sql:"rt1_recurrenceRule"`
	Rt1_instanceCreationStartDate    int     `sql:"rt1_instanceCreationStartDate"`
	Rt1_instanceCreationPaused       int     `sql:"rt1_instanceCreationPaused"`
	Rt1_instanceCreationCount        int     `sql:"rt1_instanceCreationCount"`
	Rt1_afterCompletionReferenceDate int     `sql:"rt1_afterCompletionReferenceDate"`
	Rt1_nextInstanceStartDate        int     `sql:"rt1_nextInstanceStartDate"`
	Experimental                     []byte  `sql:"experimental"`
	Repeater                         []byte  `sql:"repeater"`
	RepeaterMigrationDate            float32 `sql:"repeaterMigrationDate"`
}

type AreaTag struct {
	Areas string `sql:"areas"`
	Tags  string `sql:"tags"`
}

type TaskTag struct {
	Tasks string `sql:"tasks"`
	Tags  string `sql:"tags"`
}
