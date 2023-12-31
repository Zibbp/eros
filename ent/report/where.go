// Code generated by ent, DO NOT EDIT.

package report

import (
	"time"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"github.com/google/uuid"
	"github.com/zibbp/eros/ent/predicate"
	"github.com/zibbp/eros/internal/utils"
)

// ID filters vertices based on their ID field.
func ID(id uuid.UUID) predicate.Report {
	return predicate.Report(sql.FieldEQ(FieldID, id))
}

// IDEQ applies the EQ predicate on the ID field.
func IDEQ(id uuid.UUID) predicate.Report {
	return predicate.Report(sql.FieldEQ(FieldID, id))
}

// IDNEQ applies the NEQ predicate on the ID field.
func IDNEQ(id uuid.UUID) predicate.Report {
	return predicate.Report(sql.FieldNEQ(FieldID, id))
}

// IDIn applies the In predicate on the ID field.
func IDIn(ids ...uuid.UUID) predicate.Report {
	return predicate.Report(sql.FieldIn(FieldID, ids...))
}

// IDNotIn applies the NotIn predicate on the ID field.
func IDNotIn(ids ...uuid.UUID) predicate.Report {
	return predicate.Report(sql.FieldNotIn(FieldID, ids...))
}

// IDGT applies the GT predicate on the ID field.
func IDGT(id uuid.UUID) predicate.Report {
	return predicate.Report(sql.FieldGT(FieldID, id))
}

// IDGTE applies the GTE predicate on the ID field.
func IDGTE(id uuid.UUID) predicate.Report {
	return predicate.Report(sql.FieldGTE(FieldID, id))
}

// IDLT applies the LT predicate on the ID field.
func IDLT(id uuid.UUID) predicate.Report {
	return predicate.Report(sql.FieldLT(FieldID, id))
}

// IDLTE applies the LTE predicate on the ID field.
func IDLTE(id uuid.UUID) predicate.Report {
	return predicate.Report(sql.FieldLTE(FieldID, id))
}

// Name applies equality check predicate on the "name" field. It's identical to NameEQ.
func Name(v string) predicate.Report {
	return predicate.Report(sql.FieldEQ(FieldName, v))
}

// S3File applies equality check predicate on the "s3_file" field. It's identical to S3FileEQ.
func S3File(v string) predicate.Report {
	return predicate.Report(sql.FieldEQ(FieldS3File, v))
}

// CreatedAt applies equality check predicate on the "created_at" field. It's identical to CreatedAtEQ.
func CreatedAt(v time.Time) predicate.Report {
	return predicate.Report(sql.FieldEQ(FieldCreatedAt, v))
}

// UpdatedAt applies equality check predicate on the "updated_at" field. It's identical to UpdatedAtEQ.
func UpdatedAt(v time.Time) predicate.Report {
	return predicate.Report(sql.FieldEQ(FieldUpdatedAt, v))
}

// NameEQ applies the EQ predicate on the "name" field.
func NameEQ(v string) predicate.Report {
	return predicate.Report(sql.FieldEQ(FieldName, v))
}

// NameNEQ applies the NEQ predicate on the "name" field.
func NameNEQ(v string) predicate.Report {
	return predicate.Report(sql.FieldNEQ(FieldName, v))
}

// NameIn applies the In predicate on the "name" field.
func NameIn(vs ...string) predicate.Report {
	return predicate.Report(sql.FieldIn(FieldName, vs...))
}

// NameNotIn applies the NotIn predicate on the "name" field.
func NameNotIn(vs ...string) predicate.Report {
	return predicate.Report(sql.FieldNotIn(FieldName, vs...))
}

// NameGT applies the GT predicate on the "name" field.
func NameGT(v string) predicate.Report {
	return predicate.Report(sql.FieldGT(FieldName, v))
}

// NameGTE applies the GTE predicate on the "name" field.
func NameGTE(v string) predicate.Report {
	return predicate.Report(sql.FieldGTE(FieldName, v))
}

// NameLT applies the LT predicate on the "name" field.
func NameLT(v string) predicate.Report {
	return predicate.Report(sql.FieldLT(FieldName, v))
}

// NameLTE applies the LTE predicate on the "name" field.
func NameLTE(v string) predicate.Report {
	return predicate.Report(sql.FieldLTE(FieldName, v))
}

// NameContains applies the Contains predicate on the "name" field.
func NameContains(v string) predicate.Report {
	return predicate.Report(sql.FieldContains(FieldName, v))
}

// NameHasPrefix applies the HasPrefix predicate on the "name" field.
func NameHasPrefix(v string) predicate.Report {
	return predicate.Report(sql.FieldHasPrefix(FieldName, v))
}

// NameHasSuffix applies the HasSuffix predicate on the "name" field.
func NameHasSuffix(v string) predicate.Report {
	return predicate.Report(sql.FieldHasSuffix(FieldName, v))
}

// NameEqualFold applies the EqualFold predicate on the "name" field.
func NameEqualFold(v string) predicate.Report {
	return predicate.Report(sql.FieldEqualFold(FieldName, v))
}

// NameContainsFold applies the ContainsFold predicate on the "name" field.
func NameContainsFold(v string) predicate.Report {
	return predicate.Report(sql.FieldContainsFold(FieldName, v))
}

// StatusEQ applies the EQ predicate on the "status" field.
func StatusEQ(v utils.ReportStatus) predicate.Report {
	vc := v
	return predicate.Report(sql.FieldEQ(FieldStatus, vc))
}

// StatusNEQ applies the NEQ predicate on the "status" field.
func StatusNEQ(v utils.ReportStatus) predicate.Report {
	vc := v
	return predicate.Report(sql.FieldNEQ(FieldStatus, vc))
}

// StatusIn applies the In predicate on the "status" field.
func StatusIn(vs ...utils.ReportStatus) predicate.Report {
	v := make([]any, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Report(sql.FieldIn(FieldStatus, v...))
}

// StatusNotIn applies the NotIn predicate on the "status" field.
func StatusNotIn(vs ...utils.ReportStatus) predicate.Report {
	v := make([]any, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Report(sql.FieldNotIn(FieldStatus, v...))
}

// S3FileEQ applies the EQ predicate on the "s3_file" field.
func S3FileEQ(v string) predicate.Report {
	return predicate.Report(sql.FieldEQ(FieldS3File, v))
}

// S3FileNEQ applies the NEQ predicate on the "s3_file" field.
func S3FileNEQ(v string) predicate.Report {
	return predicate.Report(sql.FieldNEQ(FieldS3File, v))
}

// S3FileIn applies the In predicate on the "s3_file" field.
func S3FileIn(vs ...string) predicate.Report {
	return predicate.Report(sql.FieldIn(FieldS3File, vs...))
}

// S3FileNotIn applies the NotIn predicate on the "s3_file" field.
func S3FileNotIn(vs ...string) predicate.Report {
	return predicate.Report(sql.FieldNotIn(FieldS3File, vs...))
}

// S3FileGT applies the GT predicate on the "s3_file" field.
func S3FileGT(v string) predicate.Report {
	return predicate.Report(sql.FieldGT(FieldS3File, v))
}

// S3FileGTE applies the GTE predicate on the "s3_file" field.
func S3FileGTE(v string) predicate.Report {
	return predicate.Report(sql.FieldGTE(FieldS3File, v))
}

// S3FileLT applies the LT predicate on the "s3_file" field.
func S3FileLT(v string) predicate.Report {
	return predicate.Report(sql.FieldLT(FieldS3File, v))
}

// S3FileLTE applies the LTE predicate on the "s3_file" field.
func S3FileLTE(v string) predicate.Report {
	return predicate.Report(sql.FieldLTE(FieldS3File, v))
}

// S3FileContains applies the Contains predicate on the "s3_file" field.
func S3FileContains(v string) predicate.Report {
	return predicate.Report(sql.FieldContains(FieldS3File, v))
}

// S3FileHasPrefix applies the HasPrefix predicate on the "s3_file" field.
func S3FileHasPrefix(v string) predicate.Report {
	return predicate.Report(sql.FieldHasPrefix(FieldS3File, v))
}

// S3FileHasSuffix applies the HasSuffix predicate on the "s3_file" field.
func S3FileHasSuffix(v string) predicate.Report {
	return predicate.Report(sql.FieldHasSuffix(FieldS3File, v))
}

// S3FileIsNil applies the IsNil predicate on the "s3_file" field.
func S3FileIsNil() predicate.Report {
	return predicate.Report(sql.FieldIsNull(FieldS3File))
}

// S3FileNotNil applies the NotNil predicate on the "s3_file" field.
func S3FileNotNil() predicate.Report {
	return predicate.Report(sql.FieldNotNull(FieldS3File))
}

// S3FileEqualFold applies the EqualFold predicate on the "s3_file" field.
func S3FileEqualFold(v string) predicate.Report {
	return predicate.Report(sql.FieldEqualFold(FieldS3File, v))
}

// S3FileContainsFold applies the ContainsFold predicate on the "s3_file" field.
func S3FileContainsFold(v string) predicate.Report {
	return predicate.Report(sql.FieldContainsFold(FieldS3File, v))
}

// CreatedAtEQ applies the EQ predicate on the "created_at" field.
func CreatedAtEQ(v time.Time) predicate.Report {
	return predicate.Report(sql.FieldEQ(FieldCreatedAt, v))
}

// CreatedAtNEQ applies the NEQ predicate on the "created_at" field.
func CreatedAtNEQ(v time.Time) predicate.Report {
	return predicate.Report(sql.FieldNEQ(FieldCreatedAt, v))
}

// CreatedAtIn applies the In predicate on the "created_at" field.
func CreatedAtIn(vs ...time.Time) predicate.Report {
	return predicate.Report(sql.FieldIn(FieldCreatedAt, vs...))
}

// CreatedAtNotIn applies the NotIn predicate on the "created_at" field.
func CreatedAtNotIn(vs ...time.Time) predicate.Report {
	return predicate.Report(sql.FieldNotIn(FieldCreatedAt, vs...))
}

// CreatedAtGT applies the GT predicate on the "created_at" field.
func CreatedAtGT(v time.Time) predicate.Report {
	return predicate.Report(sql.FieldGT(FieldCreatedAt, v))
}

// CreatedAtGTE applies the GTE predicate on the "created_at" field.
func CreatedAtGTE(v time.Time) predicate.Report {
	return predicate.Report(sql.FieldGTE(FieldCreatedAt, v))
}

// CreatedAtLT applies the LT predicate on the "created_at" field.
func CreatedAtLT(v time.Time) predicate.Report {
	return predicate.Report(sql.FieldLT(FieldCreatedAt, v))
}

// CreatedAtLTE applies the LTE predicate on the "created_at" field.
func CreatedAtLTE(v time.Time) predicate.Report {
	return predicate.Report(sql.FieldLTE(FieldCreatedAt, v))
}

// UpdatedAtEQ applies the EQ predicate on the "updated_at" field.
func UpdatedAtEQ(v time.Time) predicate.Report {
	return predicate.Report(sql.FieldEQ(FieldUpdatedAt, v))
}

// UpdatedAtNEQ applies the NEQ predicate on the "updated_at" field.
func UpdatedAtNEQ(v time.Time) predicate.Report {
	return predicate.Report(sql.FieldNEQ(FieldUpdatedAt, v))
}

// UpdatedAtIn applies the In predicate on the "updated_at" field.
func UpdatedAtIn(vs ...time.Time) predicate.Report {
	return predicate.Report(sql.FieldIn(FieldUpdatedAt, vs...))
}

// UpdatedAtNotIn applies the NotIn predicate on the "updated_at" field.
func UpdatedAtNotIn(vs ...time.Time) predicate.Report {
	return predicate.Report(sql.FieldNotIn(FieldUpdatedAt, vs...))
}

// UpdatedAtGT applies the GT predicate on the "updated_at" field.
func UpdatedAtGT(v time.Time) predicate.Report {
	return predicate.Report(sql.FieldGT(FieldUpdatedAt, v))
}

// UpdatedAtGTE applies the GTE predicate on the "updated_at" field.
func UpdatedAtGTE(v time.Time) predicate.Report {
	return predicate.Report(sql.FieldGTE(FieldUpdatedAt, v))
}

// UpdatedAtLT applies the LT predicate on the "updated_at" field.
func UpdatedAtLT(v time.Time) predicate.Report {
	return predicate.Report(sql.FieldLT(FieldUpdatedAt, v))
}

// UpdatedAtLTE applies the LTE predicate on the "updated_at" field.
func UpdatedAtLTE(v time.Time) predicate.Report {
	return predicate.Report(sql.FieldLTE(FieldUpdatedAt, v))
}

// HasScript applies the HasEdge predicate on the "script" edge.
func HasScript() predicate.Report {
	return predicate.Report(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.Edge(sqlgraph.M2O, true, ScriptTable, ScriptColumn),
		)
		sqlgraph.HasNeighbors(s, step)
	})
}

// HasScriptWith applies the HasEdge predicate on the "script" edge with a given conditions (other predicates).
func HasScriptWith(preds ...predicate.Script) predicate.Report {
	return predicate.Report(func(s *sql.Selector) {
		step := newScriptStep()
		sqlgraph.HasNeighborsWith(s, step, func(s *sql.Selector) {
			for _, p := range preds {
				p(s)
			}
		})
	})
}

// And groups predicates with the AND operator between them.
func And(predicates ...predicate.Report) predicate.Report {
	return predicate.Report(func(s *sql.Selector) {
		s1 := s.Clone().SetP(nil)
		for _, p := range predicates {
			p(s1)
		}
		s.Where(s1.P())
	})
}

// Or groups predicates with the OR operator between them.
func Or(predicates ...predicate.Report) predicate.Report {
	return predicate.Report(func(s *sql.Selector) {
		s1 := s.Clone().SetP(nil)
		for i, p := range predicates {
			if i > 0 {
				s1.Or()
			}
			p(s1)
		}
		s.Where(s1.P())
	})
}

// Not applies the not operator on the given predicate.
func Not(p predicate.Report) predicate.Report {
	return predicate.Report(func(s *sql.Selector) {
		p(s.Not())
	})
}
