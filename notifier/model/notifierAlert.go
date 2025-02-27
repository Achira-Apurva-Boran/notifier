package model

import "time"

type NotifierAlert struct {
	AlertID string
	Zipcode string
	WatchID string
	UserID  string

	AlertTriggered  bool
	TriggerUpdateTS time.Time
	AlertStatus     string
	//ENUM('NOT_SENT', 'ALERT_SEND', 'ALERT_IGNORED_DUPLICATE', 'ALERT_IGNORED_TRESHOLD_REACHED') COLLATE utf8_unicode_ci NOT NULL default 'NOT_TRIGGERED',

	FieldType string
	//ENUM('temp', 'feels_like', 'temp_min', 'temp_max', 'pressure','humidity') COLLATE utf8_unicode_ci NOT NULL,

	Operator string
	//ENUM('gt', 'gte', 'eq', 'lt', 'lte') COLLATE utf8_unicode_ci NOT NULL,

	Value float32
}
