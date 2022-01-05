package session

import (
	"ca-reservaksin/businesses/session"
	"ca-reservaksin/drivers/database/healthFacilities"
	"ca-reservaksin/drivers/database/vaccine"

	"gorm.io/gorm"
)

type Session struct {
	gorm.Model
	Id                string                            `json:"id" gorm:"primaryKey"`
	HealthFacilitesID string                            `gorm:"size:191" json:"health_facilities_id"`
	NameSession       string                            `json:"name_session"`
	Capacity          int                               `json:"capacity"`
	VaccineID         string                            `gorm:"size:191" json:"vaccine_id"`
	Date              string                            `gorm:"type:Date" json:"date"`
	Tahap             string                            `json:"tahap"`
	StartSession      string                            `json:"start_session" gorm:"type:Datetime" time_format:"2006-01-02 15:04"`
	EndSession        string                            `json:"end_session" gorm:"type:Datetime"  time_format:"2006-01-02 15:04"`
	Vaccine           vaccine.Vaccine                   `gorm:"foreignKey:VaccineID;constraint:OnUpdate:CASCADE,OnDelete:RESTRICT;"`
	HealthFacilites   healthFacilities.HealthFacilities `gorm:"foreignKey:HealthFacilitesID;constraint:OnUpdate:CASCADE,OnDelete:RESTRICT;"`
}

func (rec *Session) ToDomain() session.Domain {
	return session.Domain{
		Id:                rec.Id,
		HealthFacilitesID: rec.HealthFacilitesID,
		HealthFacilites:   rec.HealthFacilites.ToDomain(),
		NameSession:       rec.NameSession,
		Capacity:          rec.Capacity,
		VaccineID:         rec.VaccineID,
		Vaccine:           rec.Vaccine.ToDomain(),
		Date:              rec.Date,
		Tahap:             rec.Tahap,
		StartSession:      rec.StartSession,
		EndSession:        rec.EndSession,
		CreatedAt:         rec.CreatedAt,
		UpdatedAt:         rec.UpdatedAt,
	}
}

func FromDomain(dataSession *session.Domain) *Session {
	return &Session{
		Id:                dataSession.Id,
		HealthFacilitesID: dataSession.HealthFacilitesID,
		NameSession:       dataSession.NameSession,
		Capacity:          dataSession.Capacity,
		VaccineID:         dataSession.VaccineID,
		Date:              dataSession.Date,
		Tahap:             dataSession.Tahap,
		StartSession:      dataSession.StartSession,
		EndSession:        dataSession.EndSession,
	}
}
