package typeregistrypkg

import (
	codecpkg "engine/modules/codec/pkg"
	prototypepkg "engine/modules/prototype/pkg"
	smoothpkg "engine/modules/smooth/pkg"
	"engine/modules/transition"
	transitionpkg "engine/modules/transition/pkg"
	"reflect"
	"strings"

	"github.com/ogiusek/ioc/v2"
)

func isComponent[T any]() bool {
	t := reflect.TypeFor[T]()
	typeName := t.String()
	return strings.HasSuffix(typeName, "Component")
}

func isComparable[T any]() bool {
	return reflect.TypeFor[T]().Comparable()
}

func PkgT[T any](b ioc.Builder) {
	var zero T
	pkgs := []ioc.Pkg{
		codecpkg.PkgT[T],
	}

	if !isComponent[T]() {
		goto register
	}

	if isComparable[T]() {
		pkgs = append(pkgs,
			prototypepkg.PkgT[T],
		)
	}

	if _, ok := any(zero).(transition.LerpConstraint[T]); ok {
		pkgs = append(pkgs,
			transitionpkg.PkgT[T],
			smoothpkg.PkgT[T],
		)
	}

register:
	for _, pkg := range pkgs {
		pkg(b)
	}
}
