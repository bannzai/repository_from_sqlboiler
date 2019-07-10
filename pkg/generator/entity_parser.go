package generator

import "github.com/bannzai/repository_from_sqlboiler/pkg/model"

type EntityParser interface {
	Parse() model.Entity
}
