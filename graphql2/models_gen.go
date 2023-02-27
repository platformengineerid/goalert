// Code generated by github.com/99designs/gqlgen, DO NOT EDIT.

package graphql2

import (
	"fmt"
	"io"
	"strconv"
	"time"

	"github.com/target/goalert/alert"
	"github.com/target/goalert/alert/alertlog"
	"github.com/target/goalert/assignment"
	"github.com/target/goalert/escalation"
	"github.com/target/goalert/integrationkey"
	"github.com/target/goalert/label"
	"github.com/target/goalert/limit"
	"github.com/target/goalert/notification/slack"
	"github.com/target/goalert/notification/webhook"
	"github.com/target/goalert/override"
	"github.com/target/goalert/schedule"
	"github.com/target/goalert/schedule/rotation"
	"github.com/target/goalert/schedule/rule"
	"github.com/target/goalert/service"
	"github.com/target/goalert/user"
	"github.com/target/goalert/user/contactmethod"
	"github.com/target/goalert/util/timeutil"
)

type AlertConnection struct {
	Nodes    []alert.Alert `json:"nodes"`
	PageInfo *PageInfo     `json:"pageInfo"`
}

type AlertDataPoint struct {
	Timestamp  time.Time `json:"timestamp"`
	AlertCount int       `json:"alertCount"`
}

type AlertLogEntryConnection struct {
	Nodes    []alertlog.Entry `json:"nodes"`
	PageInfo *PageInfo        `json:"pageInfo"`
}

type AlertMetricsOptions struct {
	RInterval         timeutil.ISORInterval `json:"rInterval"`
	FilterByServiceID []string              `json:"filterByServiceID"`
}

type AlertPendingNotification struct {
	Destination string `json:"destination"`
}

type AlertRecentEventsOptions struct {
	Limit *int    `json:"limit"`
	After *string `json:"after"`
}

type AlertSearchOptions struct {
	FilterByStatus    []AlertStatus    `json:"filterByStatus"`
	FilterByServiceID []string         `json:"filterByServiceID"`
	Search            *string          `json:"search"`
	First             *int             `json:"first"`
	After             *string          `json:"after"`
	FavoritesOnly     *bool            `json:"favoritesOnly"`
	IncludeNotified   *bool            `json:"includeNotified"`
	Omit              []int            `json:"omit"`
	Sort              *AlertSearchSort `json:"sort"`
	CreatedBefore     *time.Time       `json:"createdBefore"`
	NotCreatedBefore  *time.Time       `json:"notCreatedBefore"`
	ClosedBefore      *time.Time       `json:"closedBefore"`
	NotClosedBefore   *time.Time       `json:"notClosedBefore"`
}

type AuthSubjectConnection struct {
	Nodes    []user.AuthSubject `json:"nodes"`
	PageInfo *PageInfo          `json:"pageInfo"`
}

type CalcRotationHandoffTimesInput struct {
	Handoff          time.Time  `json:"handoff"`
	From             *time.Time `json:"from"`
	TimeZone         string     `json:"timeZone"`
	ShiftLengthHours int        `json:"shiftLengthHours"`
	Count            int        `json:"count"`
}

type ChanWebhookConnection struct {
	Nodes    []webhook.ChanWebhook `json:"nodes"`
	PageInfo *PageInfo             `json:"pageInfo"`
}

type ChanWebhookSearchOptions struct {
	First              *int     `json:"first"`
	After              *string  `json:"after"`
	Search             *string  `json:"search"`
	Omit               []string `json:"omit"`
	EscalationPolicyID *string  `json:"escalationPolicyID"`
}

type ClearTemporarySchedulesInput struct {
	ScheduleID string    `json:"scheduleID"`
	Start      time.Time `json:"start"`
	End        time.Time `json:"end"`
}

type ConfigHint struct {
	ID    string `json:"id"`
	Value string `json:"value"`
}

type ConfigValue struct {
	ID          string     `json:"id"`
	Description string     `json:"description"`
	Value       string     `json:"value"`
	Type        ConfigType `json:"type"`
	Password    bool       `json:"password"`
	Deprecated  string     `json:"deprecated"`
}

type ConfigValueInput struct {
	ID    string `json:"id"`
	Value string `json:"value"`
}

type CreateAlertInput struct {
	Summary   string  `json:"summary"`
	Details   *string `json:"details"`
	ServiceID string  `json:"serviceID"`
	Sanitize  *bool   `json:"sanitize"`
}

type CreateEscalationPolicyInput struct {
	Name        string                            `json:"name"`
	Description *string                           `json:"description"`
	Repeat      *int                              `json:"repeat"`
	Favorite    *bool                             `json:"favorite"`
	Steps       []CreateEscalationPolicyStepInput `json:"steps"`
}

type CreateEscalationPolicyStepInput struct {
	EscalationPolicyID *string                `json:"escalationPolicyID"`
	DelayMinutes       int                    `json:"delayMinutes"`
	Targets            []assignment.RawTarget `json:"targets"`
	NewRotation        *CreateRotationInput   `json:"newRotation"`
	NewSchedule        *CreateScheduleInput   `json:"newSchedule"`
}

type CreateHeartbeatMonitorInput struct {
	ServiceID      *string `json:"serviceID"`
	Name           string  `json:"name"`
	TimeoutMinutes int     `json:"timeoutMinutes"`
}

type CreateIntegrationKeyInput struct {
	ServiceID *string            `json:"serviceID"`
	Type      IntegrationKeyType `json:"type"`
	Name      string             `json:"name"`
}

type CreateRotationInput struct {
	Name        string        `json:"name"`
	Description *string       `json:"description"`
	TimeZone    string        `json:"timeZone"`
	Start       time.Time     `json:"start"`
	Favorite    *bool         `json:"favorite"`
	Type        rotation.Type `json:"type"`
	ShiftLength *int          `json:"shiftLength"`
	UserIDs     []string      `json:"userIDs"`
}

type CreateScheduleInput struct {
	Name             string                    `json:"name"`
	Description      *string                   `json:"description"`
	TimeZone         string                    `json:"timeZone"`
	Favorite         *bool                     `json:"favorite"`
	Targets          []ScheduleTargetInput     `json:"targets"`
	NewUserOverrides []CreateUserOverrideInput `json:"newUserOverrides"`
}

type CreateServiceInput struct {
	Name                 string                        `json:"name"`
	Description          *string                       `json:"description"`
	Favorite             *bool                         `json:"favorite"`
	EscalationPolicyID   *string                       `json:"escalationPolicyID"`
	NewEscalationPolicy  *CreateEscalationPolicyInput  `json:"newEscalationPolicy"`
	NewIntegrationKeys   []CreateIntegrationKeyInput   `json:"newIntegrationKeys"`
	Labels               []SetLabelInput               `json:"labels"`
	NewHeartbeatMonitors []CreateHeartbeatMonitorInput `json:"newHeartbeatMonitors"`
}

type CreateUserCalendarSubscriptionInput struct {
	Name            string `json:"name"`
	ReminderMinutes []int  `json:"reminderMinutes"`
	ScheduleID      string `json:"scheduleID"`
	Disabled        *bool  `json:"disabled"`
}

type CreateUserContactMethodInput struct {
	UserID                  string                           `json:"userID"`
	Type                    contactmethod.Type               `json:"type"`
	Name                    string                           `json:"name"`
	Value                   string                           `json:"value"`
	NewUserNotificationRule *CreateUserNotificationRuleInput `json:"newUserNotificationRule"`
}

type CreateUserInput struct {
	Username string    `json:"username"`
	Password string    `json:"password"`
	Name     *string   `json:"name"`
	Email    *string   `json:"email"`
	Role     *UserRole `json:"role"`
	Favorite *bool     `json:"favorite"`
}

type CreateUserNotificationRuleInput struct {
	UserID          *string `json:"userID"`
	ContactMethodID *string `json:"contactMethodID"`
	DelayMinutes    int     `json:"delayMinutes"`
}

type CreateUserOverrideInput struct {
	ScheduleID   *string   `json:"scheduleID"`
	Start        time.Time `json:"start"`
	End          time.Time `json:"end"`
	AddUserID    *string   `json:"addUserID"`
	RemoveUserID *string   `json:"removeUserID"`
}

type DebugCarrierInfoInput struct {
	Number string `json:"number"`
}

type DebugMessage struct {
	ID          string     `json:"id"`
	CreatedAt   time.Time  `json:"createdAt"`
	UpdatedAt   time.Time  `json:"updatedAt"`
	Type        string     `json:"type"`
	Status      string     `json:"status"`
	UserID      *string    `json:"userID"`
	UserName    *string    `json:"userName"`
	Source      *string    `json:"source"`
	Destination string     `json:"destination"`
	ServiceID   *string    `json:"serviceID"`
	ServiceName *string    `json:"serviceName"`
	AlertID     *int       `json:"alertID"`
	ProviderID  *string    `json:"providerID"`
	SentAt      *time.Time `json:"sentAt"`
	RetryCount  int        `json:"retryCount"`
}

type DebugMessageStatusInfo struct {
	State *NotificationState `json:"state"`
}

type DebugMessageStatusInput struct {
	ProviderMessageID string `json:"providerMessageID"`
}

type DebugMessagesInput struct {
	First         *int       `json:"first"`
	CreatedBefore *time.Time `json:"createdBefore"`
	CreatedAfter  *time.Time `json:"createdAfter"`
}

type DebugSendSMSInfo struct {
	ID          string `json:"id"`
	ProviderURL string `json:"providerURL"`
	FromNumber  string `json:"fromNumber"`
}

type DebugSendSMSInput struct {
	From string `json:"from"`
	To   string `json:"to"`
	Body string `json:"body"`
}

type EscalationPolicyConnection struct {
	Nodes    []escalation.Policy `json:"nodes"`
	PageInfo *PageInfo           `json:"pageInfo"`
}

type EscalationPolicySearchOptions struct {
	First          *int     `json:"first"`
	After          *string  `json:"after"`
	Search         *string  `json:"search"`
	Omit           []string `json:"omit"`
	FavoritesOnly  *bool    `json:"favoritesOnly"`
	FavoritesFirst *bool    `json:"favoritesFirst"`
}

type IntegrationKeyConnection struct {
	Nodes    []integrationkey.IntegrationKey `json:"nodes"`
	PageInfo *PageInfo                       `json:"pageInfo"`
}

type IntegrationKeySearchOptions struct {
	First  *int     `json:"first"`
	After  *string  `json:"after"`
	Search *string  `json:"search"`
	Omit   []string `json:"omit"`
}

type LabelConnection struct {
	Nodes    []label.Label `json:"nodes"`
	PageInfo *PageInfo     `json:"pageInfo"`
}

type LabelKeySearchOptions struct {
	First  *int     `json:"first"`
	After  *string  `json:"after"`
	Search *string  `json:"search"`
	Omit   []string `json:"omit"`
}

type LabelSearchOptions struct {
	First      *int     `json:"first"`
	After      *string  `json:"after"`
	Search     *string  `json:"search"`
	UniqueKeys *bool    `json:"uniqueKeys"`
	Omit       []string `json:"omit"`
}

type LabelValueSearchOptions struct {
	Key    string   `json:"key"`
	First  *int     `json:"first"`
	After  *string  `json:"after"`
	Search *string  `json:"search"`
	Omit   []string `json:"omit"`
}

type LinkAccountInfo struct {
	UserDetails    string       `json:"userDetails"`
	AlertID        *int         `json:"alertID"`
	AlertNewStatus *AlertStatus `json:"alertNewStatus"`
}

type MessageLogConnection struct {
	Nodes    []DebugMessage `json:"nodes"`
	PageInfo *PageInfo      `json:"pageInfo"`
}

type MessageLogSearchOptions struct {
	First         *int       `json:"first"`
	After         *string    `json:"after"`
	CreatedBefore *time.Time `json:"createdBefore"`
	CreatedAfter  *time.Time `json:"createdAfter"`
	Search        *string    `json:"search"`
	Omit          []string   `json:"omit"`
}

type NotificationState struct {
	Details           string              `json:"details"`
	Status            *NotificationStatus `json:"status"`
	FormattedSrcValue string              `json:"formattedSrcValue"`
}

type PageInfo struct {
	EndCursor   *string `json:"endCursor"`
	HasNextPage bool    `json:"hasNextPage"`
}

type PhoneNumberInfo struct {
	ID          string `json:"id"`
	CountryCode string `json:"countryCode"`
	RegionCode  string `json:"regionCode"`
	Formatted   string `json:"formatted"`
	Valid       bool   `json:"valid"`
	Error       string `json:"error"`
}

type RotationConnection struct {
	Nodes    []rotation.Rotation `json:"nodes"`
	PageInfo *PageInfo           `json:"pageInfo"`
}

type RotationSearchOptions struct {
	First          *int     `json:"first"`
	After          *string  `json:"after"`
	Search         *string  `json:"search"`
	Omit           []string `json:"omit"`
	FavoritesOnly  *bool    `json:"favoritesOnly"`
	FavoritesFirst *bool    `json:"favoritesFirst"`
}

type SWOConnection struct {
	Name    string `json:"name"`
	Version string `json:"version"`
	Type    string `json:"type"`
	IsNext  bool   `json:"isNext"`
	Count   int    `json:"count"`
}

type SWONode struct {
	ID       string `json:"id"`
	CanExec  bool   `json:"canExec"`
	IsLeader bool   `json:"isLeader"`
	// The uptime of the node in seconds. Empty if the node/connection is *not* a GoAlert instance in SWO mode.
	Uptime      string          `json:"uptime"`
	ConfigError string          `json:"configError"`
	Connections []SWOConnection `json:"connections"`
}

type SWOStatus struct {
	State         SWOState  `json:"state"`
	LastStatus    string    `json:"lastStatus"`
	LastError     string    `json:"lastError"`
	Nodes         []SWONode `json:"nodes"`
	MainDBVersion string    `json:"mainDBVersion"`
	NextDBVersion string    `json:"nextDBVersion"`
}

type ScheduleConnection struct {
	Nodes    []schedule.Schedule `json:"nodes"`
	PageInfo *PageInfo           `json:"pageInfo"`
}

type ScheduleRuleInput struct {
	ID            *string                 `json:"id"`
	Start         *timeutil.Clock         `json:"start"`
	End           *timeutil.Clock         `json:"end"`
	WeekdayFilter *timeutil.WeekdayFilter `json:"weekdayFilter"`
}

type ScheduleSearchOptions struct {
	First          *int     `json:"first"`
	After          *string  `json:"after"`
	Search         *string  `json:"search"`
	Omit           []string `json:"omit"`
	FavoritesOnly  *bool    `json:"favoritesOnly"`
	FavoritesFirst *bool    `json:"favoritesFirst"`
}

type ScheduleTarget struct {
	ScheduleID string                `json:"scheduleID"`
	Target     *assignment.RawTarget `json:"target"`
	Rules      []rule.Rule           `json:"rules"`
}

type ScheduleTargetInput struct {
	ScheduleID  *string               `json:"scheduleID"`
	Target      *assignment.RawTarget `json:"target"`
	NewRotation *CreateRotationInput  `json:"newRotation"`
	Rules       []ScheduleRuleInput   `json:"rules"`
}

type SendContactMethodVerificationInput struct {
	ContactMethodID string `json:"contactMethodID"`
}

type ServiceConnection struct {
	Nodes    []service.Service `json:"nodes"`
	PageInfo *PageInfo         `json:"pageInfo"`
}

type ServiceSearchOptions struct {
	First          *int     `json:"first"`
	After          *string  `json:"after"`
	Search         *string  `json:"search"`
	Omit           []string `json:"omit"`
	FavoritesOnly  *bool    `json:"favoritesOnly"`
	FavoritesFirst *bool    `json:"favoritesFirst"`
}

type SetFavoriteInput struct {
	Target   *assignment.RawTarget `json:"target"`
	Favorite bool                  `json:"favorite"`
}

type SetLabelInput struct {
	Target *assignment.RawTarget `json:"target"`
	Key    string                `json:"key"`
	Value  string                `json:"value"`
}

type SetScheduleOnCallNotificationRulesInput struct {
	ScheduleID string                        `json:"scheduleID"`
	Rules      []OnCallNotificationRuleInput `json:"rules"`
}

type SetTemporaryScheduleInput struct {
	ScheduleID string                `json:"scheduleID"`
	ClearStart *time.Time            `json:"clearStart"`
	ClearEnd   *time.Time            `json:"clearEnd"`
	Start      time.Time             `json:"start"`
	End        time.Time             `json:"end"`
	Shifts     []schedule.FixedShift `json:"shifts"`
}

type SlackChannelConnection struct {
	Nodes    []slack.Channel `json:"nodes"`
	PageInfo *PageInfo       `json:"pageInfo"`
}

type SlackChannelSearchOptions struct {
	First  *int     `json:"first"`
	After  *string  `json:"after"`
	Search *string  `json:"search"`
	Omit   []string `json:"omit"`
}

type StringConnection struct {
	Nodes    []string  `json:"nodes"`
	PageInfo *PageInfo `json:"pageInfo"`
}

type SystemLimit struct {
	ID          limit.ID `json:"id"`
	Description string   `json:"description"`
	Value       int      `json:"value"`
}

type SystemLimitInput struct {
	ID    limit.ID `json:"id"`
	Value int      `json:"value"`
}

type TimeZone struct {
	ID string `json:"id"`
}

type TimeZoneConnection struct {
	Nodes    []TimeZone `json:"nodes"`
	PageInfo *PageInfo  `json:"pageInfo"`
}

type TimeZoneSearchOptions struct {
	First  *int     `json:"first"`
	After  *string  `json:"after"`
	Search *string  `json:"search"`
	Omit   []string `json:"omit"`
}

type UpdateAlertsByServiceInput struct {
	ServiceID string      `json:"serviceID"`
	NewStatus AlertStatus `json:"newStatus"`
}

type UpdateAlertsInput struct {
	AlertIDs  []int       `json:"alertIDs"`
	NewStatus AlertStatus `json:"newStatus"`
}

type UpdateEscalationPolicyInput struct {
	ID          string   `json:"id"`
	Name        *string  `json:"name"`
	Description *string  `json:"description"`
	Repeat      *int     `json:"repeat"`
	StepIDs     []string `json:"stepIDs"`
}

type UpdateEscalationPolicyStepInput struct {
	ID           string                 `json:"id"`
	DelayMinutes *int                   `json:"delayMinutes"`
	Targets      []assignment.RawTarget `json:"targets"`
}

type UpdateHeartbeatMonitorInput struct {
	ID             string  `json:"id"`
	Name           *string `json:"name"`
	TimeoutMinutes *int    `json:"timeoutMinutes"`
}

type UpdateRotationInput struct {
	ID              string         `json:"id"`
	Name            *string        `json:"name"`
	Description     *string        `json:"description"`
	TimeZone        *string        `json:"timeZone"`
	Start           *time.Time     `json:"start"`
	Type            *rotation.Type `json:"type"`
	ShiftLength     *int           `json:"shiftLength"`
	ActiveUserIndex *int           `json:"activeUserIndex"`
	UserIDs         []string       `json:"userIDs"`
}

type UpdateScheduleInput struct {
	ID          string  `json:"id"`
	Name        *string `json:"name"`
	Description *string `json:"description"`
	TimeZone    *string `json:"timeZone"`
}

type UpdateServiceInput struct {
	ID                   string     `json:"id"`
	Name                 *string    `json:"name"`
	Description          *string    `json:"description"`
	EscalationPolicyID   *string    `json:"escalationPolicyID"`
	MaintenanceExpiresAt *time.Time `json:"maintenanceExpiresAt"`
}

type UpdateUserCalendarSubscriptionInput struct {
	ID              string  `json:"id"`
	Name            *string `json:"name"`
	ReminderMinutes []int   `json:"reminderMinutes"`
	Disabled        *bool   `json:"disabled"`
}

type UpdateUserContactMethodInput struct {
	ID    string  `json:"id"`
	Name  *string `json:"name"`
	Value *string `json:"value"`
}

type UpdateUserInput struct {
	ID                          string    `json:"id"`
	Name                        *string   `json:"name"`
	Email                       *string   `json:"email"`
	Role                        *UserRole `json:"role"`
	StatusUpdateContactMethodID *string   `json:"statusUpdateContactMethodID"`
}

type UpdateUserOverrideInput struct {
	ID           string     `json:"id"`
	Start        *time.Time `json:"start"`
	End          *time.Time `json:"end"`
	AddUserID    *string    `json:"addUserID"`
	RemoveUserID *string    `json:"removeUserID"`
}

type UserConnection struct {
	Nodes    []user.User `json:"nodes"`
	PageInfo *PageInfo   `json:"pageInfo"`
}

type UserOverrideConnection struct {
	Nodes    []override.UserOverride `json:"nodes"`
	PageInfo *PageInfo               `json:"pageInfo"`
}

type UserOverrideSearchOptions struct {
	First              *int       `json:"first"`
	After              *string    `json:"after"`
	Omit               []string   `json:"omit"`
	ScheduleID         *string    `json:"scheduleID"`
	FilterAddUserID    []string   `json:"filterAddUserID"`
	FilterRemoveUserID []string   `json:"filterRemoveUserID"`
	FilterAnyUserID    []string   `json:"filterAnyUserID"`
	Start              *time.Time `json:"start"`
	End                *time.Time `json:"end"`
}

type UserSearchOptions struct {
	First          *int                `json:"first"`
	After          *string             `json:"after"`
	Search         *string             `json:"search"`
	Omit           []string            `json:"omit"`
	CMValue        *string             `json:"CMValue"`
	CMType         *contactmethod.Type `json:"CMType"`
	FavoritesOnly  *bool               `json:"favoritesOnly"`
	FavoritesFirst *bool               `json:"favoritesFirst"`
}

type VerifyContactMethodInput struct {
	ContactMethodID string `json:"contactMethodID"`
	Code            int    `json:"code"`
}

type AlertSearchSort string

const (
	AlertSearchSortStatusID      AlertSearchSort = "statusID"
	AlertSearchSortDateID        AlertSearchSort = "dateID"
	AlertSearchSortDateIDReverse AlertSearchSort = "dateIDReverse"
)

var AllAlertSearchSort = []AlertSearchSort{
	AlertSearchSortStatusID,
	AlertSearchSortDateID,
	AlertSearchSortDateIDReverse,
}

func (e AlertSearchSort) IsValid() bool {
	switch e {
	case AlertSearchSortStatusID, AlertSearchSortDateID, AlertSearchSortDateIDReverse:
		return true
	}
	return false
}

func (e AlertSearchSort) String() string {
	return string(e)
}

func (e *AlertSearchSort) UnmarshalGQL(v interface{}) error {
	str, ok := v.(string)
	if !ok {
		return fmt.Errorf("enums must be strings")
	}

	*e = AlertSearchSort(str)
	if !e.IsValid() {
		return fmt.Errorf("%s is not a valid AlertSearchSort", str)
	}
	return nil
}

func (e AlertSearchSort) MarshalGQL(w io.Writer) {
	fmt.Fprint(w, strconv.Quote(e.String()))
}

type AlertStatus string

const (
	AlertStatusStatusAcknowledged   AlertStatus = "StatusAcknowledged"
	AlertStatusStatusClosed         AlertStatus = "StatusClosed"
	AlertStatusStatusUnacknowledged AlertStatus = "StatusUnacknowledged"
)

var AllAlertStatus = []AlertStatus{
	AlertStatusStatusAcknowledged,
	AlertStatusStatusClosed,
	AlertStatusStatusUnacknowledged,
}

func (e AlertStatus) IsValid() bool {
	switch e {
	case AlertStatusStatusAcknowledged, AlertStatusStatusClosed, AlertStatusStatusUnacknowledged:
		return true
	}
	return false
}

func (e AlertStatus) String() string {
	return string(e)
}

func (e *AlertStatus) UnmarshalGQL(v interface{}) error {
	str, ok := v.(string)
	if !ok {
		return fmt.Errorf("enums must be strings")
	}

	*e = AlertStatus(str)
	if !e.IsValid() {
		return fmt.Errorf("%s is not a valid AlertStatus", str)
	}
	return nil
}

func (e AlertStatus) MarshalGQL(w io.Writer) {
	fmt.Fprint(w, strconv.Quote(e.String()))
}

type ConfigType string

const (
	ConfigTypeString     ConfigType = "string"
	ConfigTypeStringList ConfigType = "stringList"
	ConfigTypeInteger    ConfigType = "integer"
	ConfigTypeBoolean    ConfigType = "boolean"
)

var AllConfigType = []ConfigType{
	ConfigTypeString,
	ConfigTypeStringList,
	ConfigTypeInteger,
	ConfigTypeBoolean,
}

func (e ConfigType) IsValid() bool {
	switch e {
	case ConfigTypeString, ConfigTypeStringList, ConfigTypeInteger, ConfigTypeBoolean:
		return true
	}
	return false
}

func (e ConfigType) String() string {
	return string(e)
}

func (e *ConfigType) UnmarshalGQL(v interface{}) error {
	str, ok := v.(string)
	if !ok {
		return fmt.Errorf("enums must be strings")
	}

	*e = ConfigType(str)
	if !e.IsValid() {
		return fmt.Errorf("%s is not a valid ConfigType", str)
	}
	return nil
}

func (e ConfigType) MarshalGQL(w io.Writer) {
	fmt.Fprint(w, strconv.Quote(e.String()))
}

type IntegrationKeyType string

const (
	IntegrationKeyTypeGeneric                IntegrationKeyType = "generic"
	IntegrationKeyTypeGrafana                IntegrationKeyType = "grafana"
	IntegrationKeyTypeSite24x7               IntegrationKeyType = "site24x7"
	IntegrationKeyTypePrometheusAlertmanager IntegrationKeyType = "prometheusAlertmanager"
	IntegrationKeyTypeEmail                  IntegrationKeyType = "email"
)

var AllIntegrationKeyType = []IntegrationKeyType{
	IntegrationKeyTypeGeneric,
	IntegrationKeyTypeGrafana,
	IntegrationKeyTypeSite24x7,
	IntegrationKeyTypePrometheusAlertmanager,
	IntegrationKeyTypeEmail,
}

func (e IntegrationKeyType) IsValid() bool {
	switch e {
	case IntegrationKeyTypeGeneric, IntegrationKeyTypeGrafana, IntegrationKeyTypeSite24x7, IntegrationKeyTypePrometheusAlertmanager, IntegrationKeyTypeEmail:
		return true
	}
	return false
}

func (e IntegrationKeyType) String() string {
	return string(e)
}

func (e *IntegrationKeyType) UnmarshalGQL(v interface{}) error {
	str, ok := v.(string)
	if !ok {
		return fmt.Errorf("enums must be strings")
	}

	*e = IntegrationKeyType(str)
	if !e.IsValid() {
		return fmt.Errorf("%s is not a valid IntegrationKeyType", str)
	}
	return nil
}

func (e IntegrationKeyType) MarshalGQL(w io.Writer) {
	fmt.Fprint(w, strconv.Quote(e.String()))
}

type NotificationStatus string

const (
	NotificationStatusOk    NotificationStatus = "OK"
	NotificationStatusWarn  NotificationStatus = "WARN"
	NotificationStatusError NotificationStatus = "ERROR"
)

var AllNotificationStatus = []NotificationStatus{
	NotificationStatusOk,
	NotificationStatusWarn,
	NotificationStatusError,
}

func (e NotificationStatus) IsValid() bool {
	switch e {
	case NotificationStatusOk, NotificationStatusWarn, NotificationStatusError:
		return true
	}
	return false
}

func (e NotificationStatus) String() string {
	return string(e)
}

func (e *NotificationStatus) UnmarshalGQL(v interface{}) error {
	str, ok := v.(string)
	if !ok {
		return fmt.Errorf("enums must be strings")
	}

	*e = NotificationStatus(str)
	if !e.IsValid() {
		return fmt.Errorf("%s is not a valid NotificationStatus", str)
	}
	return nil
}

func (e NotificationStatus) MarshalGQL(w io.Writer) {
	fmt.Fprint(w, strconv.Quote(e.String()))
}

type SWOAction string

const (
	SWOActionReset   SWOAction = "reset"
	SWOActionExecute SWOAction = "execute"
)

var AllSWOAction = []SWOAction{
	SWOActionReset,
	SWOActionExecute,
}

func (e SWOAction) IsValid() bool {
	switch e {
	case SWOActionReset, SWOActionExecute:
		return true
	}
	return false
}

func (e SWOAction) String() string {
	return string(e)
}

func (e *SWOAction) UnmarshalGQL(v interface{}) error {
	str, ok := v.(string)
	if !ok {
		return fmt.Errorf("enums must be strings")
	}

	*e = SWOAction(str)
	if !e.IsValid() {
		return fmt.Errorf("%s is not a valid SWOAction", str)
	}
	return nil
}

func (e SWOAction) MarshalGQL(w io.Writer) {
	fmt.Fprint(w, strconv.Quote(e.String()))
}

type SWOState string

const (
	SWOStateUnknown   SWOState = "unknown"
	SWOStateResetting SWOState = "resetting"
	SWOStateIdle      SWOState = "idle"
	SWOStateSyncing   SWOState = "syncing"
	SWOStatePausing   SWOState = "pausing"
	SWOStateExecuting SWOState = "executing"
	SWOStateDone      SWOState = "done"
)

var AllSWOState = []SWOState{
	SWOStateUnknown,
	SWOStateResetting,
	SWOStateIdle,
	SWOStateSyncing,
	SWOStatePausing,
	SWOStateExecuting,
	SWOStateDone,
}

func (e SWOState) IsValid() bool {
	switch e {
	case SWOStateUnknown, SWOStateResetting, SWOStateIdle, SWOStateSyncing, SWOStatePausing, SWOStateExecuting, SWOStateDone:
		return true
	}
	return false
}

func (e SWOState) String() string {
	return string(e)
}

func (e *SWOState) UnmarshalGQL(v interface{}) error {
	str, ok := v.(string)
	if !ok {
		return fmt.Errorf("enums must be strings")
	}

	*e = SWOState(str)
	if !e.IsValid() {
		return fmt.Errorf("%s is not a valid SWOState", str)
	}
	return nil
}

func (e SWOState) MarshalGQL(w io.Writer) {
	fmt.Fprint(w, strconv.Quote(e.String()))
}

type UserRole string

const (
	UserRoleUnknown UserRole = "unknown"
	UserRoleUser    UserRole = "user"
	UserRoleAdmin   UserRole = "admin"
)

var AllUserRole = []UserRole{
	UserRoleUnknown,
	UserRoleUser,
	UserRoleAdmin,
}

func (e UserRole) IsValid() bool {
	switch e {
	case UserRoleUnknown, UserRoleUser, UserRoleAdmin:
		return true
	}
	return false
}

func (e UserRole) String() string {
	return string(e)
}

func (e *UserRole) UnmarshalGQL(v interface{}) error {
	str, ok := v.(string)
	if !ok {
		return fmt.Errorf("enums must be strings")
	}

	*e = UserRole(str)
	if !e.IsValid() {
		return fmt.Errorf("%s is not a valid UserRole", str)
	}
	return nil
}

func (e UserRole) MarshalGQL(w io.Writer) {
	fmt.Fprint(w, strconv.Quote(e.String()))
}
