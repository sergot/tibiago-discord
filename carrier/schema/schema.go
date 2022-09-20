package schema

import (
	"github.com/Yiling-J/carrier"
	"github.com/sergot/tibiago/ent"
)

var (
	Schemas = []carrier.Schema{
		&carrier.EntSchema{
			To: &ent.BossCreate{},
		},
		&carrier.EntSchema{
			To: &ent.BosslistCreate{},
		},
		&carrier.EntSchema{
			To: &ent.ParticipantCreate{},
		},
		&carrier.EntSchema{
			To: &ent.InstanceCreate{},
		},
		&carrier.EntSchema{
			To: &ent.InstanceConfigCreate{},
		},
	}
)
