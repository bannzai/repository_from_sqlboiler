{{- $entity := .Entity -}}
{{- $entityName := $entity.Name -}}
{{- $entityLocalVarName := golangVariableCase $entityName -}}
{{- $entityStructName := golangStructNameCase $entityName -}}
{{- $entitySelectorName := entitySelectorName $entityName -}}

package repository 

type {{$entityStructName}}Repository interface {
  {{ .FetchByPrimaryKey }}
  FetchBy(ctx context.Context, parameter parameter.ParameterType) (entity.{{$entityStructName}}, error)
  FetchList(ctx context.Context, parameter parameter.ParameterType) ([]entity.{{$entityStructName}}, error)
  FetchAll(ctx context.Context) ([]entity.{{$entityStructName}} , error)
  Exists(ctx context.Context, parameter parameter.ParameterType) (bool, error)
}

// {{$entityStructName}} is interface of datastore (e.g RDB).
type {{$entityStructName}} struct { }
func ({{$entityStructName}}) {{ .FetchByPrimaryKey }} {
	e, err := entity.{{$entitySelectorName}}({{.SQLQueryArguments}}).OneG(ctx)

	if err != nil {
		return entity.{{$entityStructName}}{}, errors.Wrap(err, "")
	}

	return *e, nil
}

 // FetchBy is find {{$entityStructName}} entity from primary key({{.PrimaryKeys}}). When not found `ResourceNotFound` error.
func ({{$entityStructName}}) FetchBy(ctx context.Context, parameter parameter.ParameterType) (entity.{{$entityStructName}}, error) {
	{{$entityLocalVarName}}, err := entity.{{$entitySelectorName}}(parameter).OneG(ctx)

	if err != nil {
		if err == sql.ErrNoRows {
			return entity.{{$entityStructName}}{}, errors.Wrap(ResourceNotFound{err}, "")
		}
		return entity.{{$entityStructName}}{}, errors.Wrap(err, "")
	}

	return *{{$entityLocalVarName}}, nil
}

// FetchAll is fetch all for {{$entityStructName}} entity.
func ({{$entityStructName}}) FetchList(ctx context.Context, parameter parameter.ParameterType) ([]entity.{{$entityStructName}}, error) {
	{{$entityLocalVarName}}, err := entity.{{$entitySelectorName}}(parameter).AllG(ctx)

	if err != nil {
		return []entity.{{$entityStructName}}{}, errors.Wrap(err, "")
	}

	list := make([]entity.{{$entityStructName}}, len({{$entityLocalVarName}}))
	for i, e := range {{$entityLocalVarName}} {
		list[i] = *e
	}

	return list, nil
}

// FetchAll is fetch all for {{$entityStructName}} entity.
func ({{$entityStructName}}) FetchAll(ctx context.Context) ([]entity.{{$entityStructName}}, error) {
	{{$entityLocalVarName}}, err := entity.{{$entitySelectorName}}().AllG(ctx)

	if err != nil {
		return []entity.{{$entityStructName}}{}, errors.Wrap(err, "")
	}

	list := make([]entity.{{$entityStructName}}, len({{$entityLocalVarName}}))
	for i, e := range {{$entityLocalVarName}} {
		list[i] = *e
	}

	return list, nil
}

// Exists is fetch all for {{$entityStructName}} entity.
func ({{$entityStructName}}) Exists(ctx context.Context, parameter parameter.ParameterType) (bool, error) {
	exists, err := entity.{{$entitySelectorName}}(parameter).ExistsG(ctx)

	if err != nil {
		return false, errors.Wrap(err, "")
	}

	return exists, nil
}

func New{{$entityStructName}}Repository() {{$entityStructName}}Repository {
	return {{$entityStructName}}{}
}

