import React, { useContext, useState } from 'react'
import Button from '@material-ui/core/Button'
import Grid from '@material-ui/core/Grid'
import Popover from '@material-ui/core/Popover'
import Typography from '@material-ui/core/Typography'
import { makeStyles } from '@material-ui/core'
import { DateTime } from 'luxon'
import { ScheduleCalendarContext } from '../ScheduleDetails'
import CardActions from '../../details/CardActions'
import { Edit as EditIcon, Delete as DeleteIcon } from '@material-ui/icons'
import ScheduleOverrideEditDialog from '../ScheduleOverrideEditDialog'
import ScheduleOverrideDeleteDialog from '../ScheduleOverrideDeleteDialog'
import { User, UserOverride } from '../../../schema'

const useStyles = makeStyles({
  cardActionContainer: {
    width: '100%',
  },
  button: {
    padding: '4px',
    minHeight: 0,
    fontSize: 12,
  },
  buttonContainer: {
    display: 'flex',
    alignItems: 'center',
  },
  flexGrow: {
    flexGrow: 1,
  },
  paper: {
    padding: 8,
    maxWidth: 275,
  },
})

export interface ScheduleCalendarEvent {
  title: string | JSX.Element
  userID: any
  start: Date
  end: Date
  fixed: boolean
  isTempSchedShift: any
  tempSched: any
  isOverride: boolean
  override: UserOverride
}

interface ScheduleCalendarEventWrapperProps {
  event: ScheduleCalendarEvent
  children: JSX.Element
}

export default function ScheduleCalendarEventWrapper({
  children,
  event,
}: ScheduleCalendarEventWrapperProps): JSX.Element {
  const classes = useStyles()
  const [anchorEl, setAnchorEl] = useState<EventTarget | null>(null)

  const [showEditDialog, setShowEditDialog] = useState(null)
  const [showDeleteDialog, setShowDeleteDialog] = useState(null)

  const { setOverrideDialog, onEditTempSched, onDeleteTempSched } = useContext(
    ScheduleCalendarContext,
  )
  const open = Boolean(anchorEl)
  const id = open ? 'shift-popover' : undefined

  function handleClick(_event: MouseEvent): void {
    setAnchorEl(_event.currentTarget)
  }

  function handleCloseShiftInfo(): void {
    setAnchorEl(null)
  }

  function handleKeyDown(_event: KeyboardEvent): void {
    const code = _event.key
    if (code === 'Enter' || code === ' ') {
      setAnchorEl(_event.currentTarget)
    }
  }

  function handleShowOverrideForm(): void {
    handleCloseShiftInfo()

    setOverrideDialog({
      variantOptions: ['replace', 'remove'],
      removeUserReadOnly: true,
      defaultValue: {
        start: event.start.toISOString(),
        end: event.end.toISOString(),
        removeUserID: event.userID,
      },
    })
  }

  function renderTempSchedButtons(): JSX.Element {
    if (DateTime.fromISO(event.end) <= DateTime.utc()) {
      // no actions on past events
      return <React.Fragment />
    }
    return (
      <React.Fragment>
        <Grid item>
          <Button
            data-cy='edit-temp-sched'
            size='small'
            onClick={() => onEditTempSched(event.tempSched)}
            variant='contained'
            color='primary'
            title='Edit this temporary schedule'
          >
            Edit
          </Button>
        </Grid>
        {!event.isTempSchedShift && (
          <React.Fragment>
            <Grid item className={classes.flexGrow} />
            <Grid item>
              <Button
                data-cy='delete-temp-sched'
                size='small'
                onClick={() => onDeleteTempSched(event.tempSched)}
                variant='contained'
                color='primary'
                title='Delete this temporary schedule'
              >
                Delete
              </Button>
            </Grid>
          </React.Fragment>
        )}
      </React.Fragment>
    )
  }

  function renderOverrideButtons(): JSX.Element {
    return (
      <div className={classes.cardActionContainer}>
        <CardActions
          secondaryActions={[
            {
              icon: <EditIcon fontSize='small' />,
              label: 'Edit',
              handleOnClick: () => {
                handleCloseShiftInfo()
                setShowEditDialog(event?.override?.id)
              },
            },
            {
              icon: <DeleteIcon fontSize='small' />,
              label: 'Delete',
              handleOnClick: () => {
                handleCloseShiftInfo()
                setShowDeleteDialog(event?.override?.id)
              },
            },
          ]}
        />
      </div>
    )
  }

  function renderShiftButtons(): JSX.Element {
    return (
      <React.Fragment>
        <Grid item className={classes.flexGrow} />
        <Grid item>
          <Button
            data-cy='override'
            size='small'
            onClick={handleShowOverrideForm}
            variant='contained'
            color='primary'
            title={`Temporarily remove ${event.title} from this schedule`}
          >
            Override Shift
          </Button>
        </Grid>
      </React.Fragment>
    )
  }

  function renderButtons(): JSX.Element {
    if (DateTime.fromJSDate(event.end) <= DateTime.utc())
      return <React.Fragment />
    if (event.tempSched) return renderTempSchedButtons()
    if (event.fixed) return <React.Fragment />
    if (event.isOverride) return renderOverrideButtons()

    return renderShiftButtons()
  }

  function renderOverrideDescription(): JSX.Element {
    if (!event.isOverride) return <React.Fragment />

    const getDesc = (
      addUser: User | undefined,
      removeUser: User | undefined,
    ): JSX.Element => {
      if (addUser && removeUser)
        return (
          <React.Fragment>
            <b>{addUser.name}</b> replaces <b>{removeUser.name}</b>.
          </React.Fragment>
        )
      if (addUser)
        return (
          <React.Fragment>
            Adds <b>{addUser.name}</b>.
          </React.Fragment>
        )
      if (removeUser)
        return (
          <React.Fragment>
            Removes <b>{removeUser.name}</b>.
          </React.Fragment>
        )

      return <React.Fragment />
    }

    return (
      <Grid item xs={12}>
        <Typography variant='body2'>
          {getDesc(event.override.addUser, event.override.removeUser)}
        </Typography>
      </Grid>
    )
  }

  /*
   * Renders an interactive tooltip when selecting
   * an event in the calendar that will show
   * the full shift start and end date times, as
   * well as the controls relevant to the event.
   */
  function renderShiftInfo(): JSX.Element {
    const fmt = (date: Date) =>
      DateTime.fromJSDate(date).toLocaleString(DateTime.DATETIME_FULL)

    return (
      <Grid container spacing={1}>
        {renderOverrideDescription()}
        <Grid item xs={12}>
          <Typography variant='body2'>
            {`${fmt(event.start)}  –  ${fmt(event.end)}`}
          </Typography>
        </Grid>
        {renderButtons()}
      </Grid>
    )
  }

  if (!children) return <React.Fragment />
  return (
    <React.Fragment>
      <Popover
        id={id}
        open={open}
        anchorEl={anchorEl}
        onClose={handleCloseShiftInfo}
        anchorOrigin={{
          vertical: 'bottom',
          horizontal: 'left',
        }}
        transformOrigin={{
          vertical: 'top',
          horizontal: 'left',
        }}
        PaperProps={{
          'data-cy': 'shift-tooltip',
        }}
        classes={{
          paper: classes.paper,
        }}
      >
        {renderShiftInfo()}
      </Popover>
      {React.cloneElement(children, {
        tabIndex: 0,
        onClick: handleClick,
        onKeyDown: handleKeyDown,
        role: 'button',
        'aria-pressed': open,
        'aria-describedby': id,
      })}
      {showEditDialog && (
        <ScheduleOverrideEditDialog
          overrideID={showEditDialog}
          onClose={() => setShowEditDialog(null)}
        />
      )}
      {showDeleteDialog && (
        <ScheduleOverrideDeleteDialog
          overrideID={showDeleteDialog}
          onClose={() => setShowDeleteDialog(null)}
        />
      )}
    </React.Fragment>
  )
}
