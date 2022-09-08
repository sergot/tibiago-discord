package models

// type BaseModel struct {
// 	ID uuid.UUID `bun:",pk,type:uuid default uuid_generate_v4()"`
// }

// type Boss struct {
// 	BaseModel
// 	Name string
// }

// type Bosslist struct {
// 	BaseModel
// 	BossID       uuid.UUID
// 	Boss         *Boss `bun:"rel:has-one"`
// 	When         string
// 	Participants []Participant `bun:"many2many:bosslist_to_participant"`
// }

// type Participant struct {
// 	BaseModel
// 	BosslistID uuid.UUID
// }

// type BosslistToParticipant struct {
// 	ParticipantID uuid.UUID
// 	BosslistID    uuid.UUID
// }
