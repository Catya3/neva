package analyzer

import (
	"errors"
	"fmt"

	src "github.com/nevalang/neva/internal/compiler/sourcecode"
)

var (
	ErrEmptyConst                  = errors.New("const must have value or reference to another const")
	ErrEntityNotConst              = errors.New("entity is not a const")
	ErrResolveConstType            = errors.New("cannot resolve constant type")
	ErrConstValuesOfDifferentTypes = errors.New("constant cannot have values of different types at once")
)

func (a Analyzer) analyzeConst(cnst src.Const, scope src.Scope) (src.Const, error) {
	if cnst.Value == nil && cnst.Ref == nil {
		return src.Const{}, ErrEmptyConst
	}

	if cnst.Value == nil {
		entity, _, err := scope.Entity(*cnst.Ref)
		if err != nil {
			return src.Const{}, fmt.Errorf("entity: %w", err)
		}
		if entity.Kind != src.ConstEntity {
			return src.Const{}, fmt.Errorf("%w: %v", ErrEntityNotConst, entity.Kind)
		}
	}

	resolvedType, err := a.analyzeTypeExpr(cnst.Value.TypeExpr, scope)
	if err != nil {
		return src.Const{}, fmt.Errorf("%w: %v", ErrResolveConstType, err)
	}

	switch resolvedType.Inst.Ref.String() {
	case "bool":
		if cnst.Value.Int != 0 || cnst.Value.Float != 0 || cnst.Value.Str != "" {
			return src.Const{}, fmt.Errorf("%w: %v", ErrConstValuesOfDifferentTypes, cnst.Value)
		}
	case "int":
		if cnst.Value.Bool != false || cnst.Value.Float != 0 || cnst.Value.Str != "" {
			return src.Const{}, fmt.Errorf("%w: %v", ErrConstValuesOfDifferentTypes, cnst.Value)
		}
	case "float":
		if cnst.Value.Bool != false || cnst.Value.Int != 0 || cnst.Value.Str != "" {
			return src.Const{}, fmt.Errorf("%w: %v", ErrConstValuesOfDifferentTypes, cnst.Value)
		}
	case "str":
		if cnst.Value.Bool != false || cnst.Value.Int != 0 || cnst.Value.Float != 0 {
			return src.Const{}, fmt.Errorf("%w: %v", ErrConstValuesOfDifferentTypes, cnst.Value)
		}
	}

	valueCopy := *cnst.Value
	valueCopy.TypeExpr = resolvedType

	return src.Const{
		Value: &valueCopy,
	}, nil
}