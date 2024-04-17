exit

A cli based time tracking tool for freelancers that makes tracking your working time easy and exports it to excel - all
through your favorite terminal.

## Commands

- `clitt start` - starts a new work session
- `clitt stop` - ends the current active work session
- `clitt continue`- continues the last ended session, only works if there is no active session
- `clitt abort`- aborts and deletes irreversively the current session
- `clitt edit <id> --description --start --end` - edits an existing session
- `clitt delete  <id>` - deletes an existing session
- `clitt task add` - adds a description to the current session
- `clitt task list` shows all descriptions
- `clitt describe` shows all sessions
- `clitt export)` generates an excel file with all recorded sessions

## Exporting

When doing `clitt export` then all recorded work sessions are exported into an excel file where each month is stored in
a
separate sheet.