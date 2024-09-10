package models

import (
	"time"

	"github.com/sev-2/raiden"
)

type Doctors struct {
	raiden.ModelBase

	Id        int64      `json:"id,omitempty" column:"name:id;type:bigint;primaryKey;autoIncrement;nullable:false"`
	Name      *string    `json:"name,omitempty" column:"name:name;type:varchar;nullable;default:'255'::character varying"`
	Location  *string    `json:"location,omitempty" column:"name:location;type:text;nullable"`
	CreatedAt *time.Time `json:"created_at,omitempty" column:"name:created_at;type:timestampz;nullable;default:now()"`

	// Table information
	Metadata string `json:"-" schema:"public"`

	// Access control
	Acl string `json:"-" read:"" write:""`
}
