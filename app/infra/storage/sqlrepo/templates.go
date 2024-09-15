package sqlrepo

import (
	"bytes"
	"context"
	"database/sql"
	"errors"
	"io/fs"
	"path"
	"strings"
	"text/template"
	"time"

	"timeCardSimple/app/domain/id"
)

func ExecuteTemplateForQueryArgs(ctx context.Context, tq TemplateQueryer, t *template.Template, data any) (context.Context, string, []any, error) {
	if t == nil {
		return nil, "", nil, errors.New("sqlrepov2: template is nil")
	}

	templateData := &ExecuteTemplateData{
		Query: tq.TemplateQuery(t),
		Data:  data,
	}

	b := &bytes.Buffer{}
	err := t.Execute(b, templateData)
	if err != nil {
		return nil, "", nil, err
	}

	query, args := b.String(), templateData.Query.args
	return WithTemplateName(ctx, t.Name()), query, args, nil
}

type templateNameKey struct{}

func WithTemplateName(ctx context.Context, templateName string) context.Context {
	return context.WithValue(ctx, templateNameKey{}, templateName)
}

func TemplateNameFromContext(ctx context.Context) (templateName string, ok bool) {
	templateName, ok = ctx.Value(templateNameKey{}).(string)
	return
}

type ExecuteTemplateData struct {
	Query *TemplateQuery

	Data any
}

type TemplateQueryer interface {
	TemplateQuery(t *template.Template) *TemplateQuery
}

type PostgresNameTemplateQueryer struct{}

func (tq PostgresNameTemplateQueryer) TemplateQuery(t *template.Template) *TemplateQuery {
	return &TemplateQuery{
		Name:         t.Name(),
		queryDialect: Postgres{},
	}
}

type TemplateQuery struct {
	Name string

	args []any

	queryDialect QueryDialect

	placeholderIndex int
}

func (tq *TemplateQuery) Placeholders(vv ...any) string {
	var ps []string
	for _, v := range vv {
		ps = append(ps, tq.Placeholder(v))
	}
	return strings.Join(ps, ",")
}

func (tq *TemplateQuery) Placeholder(v any) string {
	v = TransformQueryPlaceholder(v)
	placeholderString := tq.queryDialect.Placeholder(tq.placeholderIndex)

	tq.placeholderIndex++
	tq.args = append(tq.args, v)

	return placeholderString
}

func TransformQueryPlaceholder(v any) any {
	switch vt := v.(type) {
	case id.ID:
		return vt.String()
	case *id.ID:
		var s string
		if vt != nil {
			s = vt.String()
		}
		return sql.NullString{
			Valid:  vt != nil,
			String: s,
		}

	case time.Time:
		return vt.UTC()
	case *time.Time:
		var t time.Time
		if vt != nil {
			t = *vt
		}
		return sql.NullTime{
			Valid: vt != nil,
			Time:  t.UTC(),
		}
	}

	return v
}

const (
	ExecutableTemplatePathPattern = "[A-Z]*"
	BaseTemplatePathPattern       = "_*"
)

// BuildQueryTemplates loads all Templates in fsys where the file name matches
// ExecutableTemplatePathPattern.
// Then, for each of those templates, it looks at the template definitions and
// parses those files in the the template.
//
// See ./testdata/singleReference for an example.
func BuildQueryTemplates(fsys fs.FS) (map[string]*template.Template, error) {
	templatePaths, err := collectExecutableTemplatePaths(fsys, ExecutableTemplatePathPattern)
	if err != nil {
		return nil, err
	}

	result := map[string]*template.Template{}
	for _, templatePath := range templatePaths {
		result[templatePath], err = template.ParseFS(fsys, templatePath)
		if err != nil {
			return nil, err
		}
	}

	for templatePath, tmpl := range result {
		basePaths := []string{}
		for _, tmplDep := range tmpl.Templates() {
			name := tmplDep.Name()
			if ok, err := path.Match(BaseTemplatePathPattern, name); err != nil {
				return nil, err
			} else if ok {
				baseName := path.Base(name)
				baseName, _, _ = strings.Cut(baseName, ".")
				joinedPath := path.Join(path.Dir(templatePath), baseName)
				basePaths = append(basePaths, joinedPath)
			}
		}
		if len(basePaths) > 0 {
			if _, err := tmpl.ParseFS(fsys, basePaths...); err != nil {
				return nil, err
			}
		}
	}

	return result, nil
}

func collectExecutableTemplatePaths(fsys fs.FS, pattern string) ([]string, error) {
	result := []string{}
	fs.WalkDir(
		fsys,
		".",
		func(filePath string, d fs.DirEntry, err error) error {
			if d != nil && d.IsDir() {
				return nil
			}
			if ok, err := path.Match(pattern, d.Name()); err != nil {
				return err
			} else if ok {
				result = append(result, filePath)
			}
			return nil
		},
	)
	return result, nil
}
