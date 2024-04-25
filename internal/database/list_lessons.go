package database

import (
	"context"
	"fmt"

	"github.com/georgysavva/scany/pgxscan"
	"github.com/romanzimoglyad/inquiry-backend/internal/database/model"

	sq "github.com/Masterminds/squirrel"
	"github.com/romanzimoglyad/inquiry-backend/internal/domain/domain"
)

func (d *Database) ListLessons(ctx context.Context, request *domain.ListLessonsRequest) (*domain.Lessons, error) {

	builder := sq.Select("lesson.*,"+
		"subject.name as subject_name,"+
		"unit.name as unit_name, "+
		"concept.name as concept_name, "+
		"skill.name as skill_name", " STRING_AGG(material_to_lesson.material_id, ',') AS material_ids").From(model.LessonTableName.String()).PlaceholderFormat(sq.Dollar).
		LeftJoin(model.SubjectTableName.String() + " ON " +
			model.LessonTableName.String() + ".subject_id" + " = " + model.SubjectTableName.String() + ".id").
		LeftJoin(model.UnitTableName.String() + " ON " +
			model.LessonTableName.String() + ".unit_id" + " = " + model.UnitTableName.String() + ".id").
		LeftJoin(model.ConceptTableName.String() + " ON " +
			model.LessonTableName.String() + ".concept_id" + " = " + model.ConceptTableName.String() + ".id").
		LeftJoin(model.SkillTableName.String() + " ON " +
			model.LessonTableName.String() + ".skill_id" + " = " + model.SkillTableName.String() + ".id").
		LeftJoin(model.MaterialToLessonTableName.String() + " ON " + model.LessonTableName.String() + ".id::TEXT" + " = " + model.MaterialToLessonTableName.String() + ".lesson_id").
		GroupBy("lesson.id," +
			"subject.name,unit.name," +
			"concept.name," +
			"skill.name").
		Where(sq.And{sq.Eq{"active": true}, sq.Eq{"user_id": request.UserId}})
	builder = withFilters(request.Filter, builder)
	builder = withOrderBy(builder)
	query, args, err := withOffset(request.Page, builder).ToSql()
	if err != nil {
		return nil, fmt.Errorf("error in selecting lessons : %w", err)
	}

	var lessons model.Lessons

	if err := pgxscan.Select(ctx, d.pool, &lessons, query, args...); err != nil {
		return nil, err
	}

	builder = sq.Select("count (*)").From(model.LessonTableName.String()).PlaceholderFormat(sq.Dollar).Where(sq.And{sq.Eq{"active": true}, sq.Eq{"user_id": request.UserId}})
	builder = withFilters(request.Filter, builder)
	query, args, err = builder.ToSql()
	if err != nil {
		return nil, fmt.Errorf("error in selecting lessons : %w", err)
	}
	var count int32
	if err := pgxscan.Get(ctx, d.pool, &count, query, args...); err != nil {
		return nil, err
	}

	return lessons.ToDomain(count), nil
}
func withOffset(page domain.Page, builder sq.SelectBuilder) sq.SelectBuilder {
	builder = builder.Offset(uint64((page.Page - 1) * page.Size)).Limit(uint64(page.Size))
	return builder
}

func withFilters(filter domain.Filter, builder sq.SelectBuilder) sq.SelectBuilder {
	if filter.SubjectId != 0 {
		builder = builder.Where(sq.Eq{"subject_id": filter.SubjectId})
	}
	if filter.ConceptId != 0 {
		builder = builder.Where(sq.Eq{"concept_id": filter.ConceptId})
	}
	if filter.SkillId != 0 {
		builder = builder.Where(sq.Eq{"skill_id": filter.SubjectId})
	}
	if filter.UnitId != 0 {
		builder = builder.Where(sq.Eq{"unit_id": filter.UnitId})
	}
	if filter.GradeId != 0 {
		builder = builder.Where(sq.Eq{"grade_id": filter.GradeId})
	}
	if filter.SearchText != "" {
		builder = builder.Where(sq.Like{"lesson.name": "%" + filter.SearchText + "%"})
	}

	return builder
}

func withOrderBy(builder sq.SelectBuilder) sq.SelectBuilder {
	return builder.OrderBy("updated_at desc")
}
