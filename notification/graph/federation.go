// Code generated by github.com/99designs/gqlgen, DO NOT EDIT.

package graph

import (
	"context"
	"errors"
	"fmt"
	"strings"
	"sync"

	"github.com/99designs/gqlgen/plugin/federation/fedruntime"
)

var (
	ErrUnknownType  = errors.New("unknown type")
	ErrTypeNotFound = errors.New("type not found")
)

func (ec *executionContext) __resolve__service(ctx context.Context) (fedruntime.Service, error) {
	if ec.DisableIntrospection {
		return fedruntime.Service{}, errors.New("federated introspection disabled")
	}

	var sdl []string

	for _, src := range sources {
		if src.BuiltIn {
			continue
		}
		sdl = append(sdl, src.Input)
	}

	return fedruntime.Service{
		SDL: strings.Join(sdl, "\n"),
	}, nil
}

func (ec *executionContext) __resolve_entities(ctx context.Context, representations []map[string]interface{}) []fedruntime.Entity {
	list := make([]fedruntime.Entity, len(representations))

	repsMap := map[string]struct {
		i []int
		r []map[string]interface{}
	}{}

	// We group entities by typename so that we can parallelize their resolution.
	// This is particularly helpful when there are entity groups in multi mode.
	buildRepresentationGroups := func(reps []map[string]interface{}) {
		for i, rep := range reps {
			typeName, ok := rep["__typename"].(string)
			if !ok {
				// If there is no __typename, we just skip the representation;
				// we just won't be resolving these unknown types.
				ec.Error(ctx, errors.New("__typename must be an existing string"))
				continue
			}

			_r := repsMap[typeName]
			_r.i = append(_r.i, i)
			_r.r = append(_r.r, rep)
			repsMap[typeName] = _r
		}
	}

	isMulti := func(typeName string) bool {
		switch typeName {
		default:
			return false
		}
	}

	resolveEntity := func(ctx context.Context, typeName string, rep map[string]interface{}, idx []int, i int) (err error) {
		// we need to do our own panic handling, because we may be called in a
		// goroutine, where the usual panic handling can't catch us
		defer func() {
			if r := recover(); r != nil {
				err = ec.Recover(ctx, r)
			}
		}()

		switch typeName {
		case "Comment":
			resolverName, err := entityResolverNameForComment(ctx, rep)
			if err != nil {
				return fmt.Errorf(`finding resolver for Entity "Comment": %w`, err)
			}
			switch resolverName {

			case "findCommentByID":
				id0, err := ec.unmarshalNID2string(ctx, rep["id"])
				if err != nil {
					return fmt.Errorf(`unmarshalling param 0 for findCommentByID(): %w`, err)
				}
				entity, err := ec.resolvers.Entity().FindCommentByID(ctx, id0)
				if err != nil {
					return fmt.Errorf(`resolving Entity "Comment": %w`, err)
				}

				list[idx[i]] = entity
				return nil
			}
		case "File":
			resolverName, err := entityResolverNameForFile(ctx, rep)
			if err != nil {
				return fmt.Errorf(`finding resolver for Entity "File": %w`, err)
			}
			switch resolverName {

			case "findFileByID":
				id0, err := ec.unmarshalNID2string(ctx, rep["id"])
				if err != nil {
					return fmt.Errorf(`unmarshalling param 0 for findFileByID(): %w`, err)
				}
				entity, err := ec.resolvers.Entity().FindFileByID(ctx, id0)
				if err != nil {
					return fmt.Errorf(`resolving Entity "File": %w`, err)
				}

				list[idx[i]] = entity
				return nil
			}
		case "Follow":
			resolverName, err := entityResolverNameForFollow(ctx, rep)
			if err != nil {
				return fmt.Errorf(`finding resolver for Entity "Follow": %w`, err)
			}
			switch resolverName {

			case "findFollowByID":
				id0, err := ec.unmarshalNID2string(ctx, rep["id"])
				if err != nil {
					return fmt.Errorf(`unmarshalling param 0 for findFollowByID(): %w`, err)
				}
				entity, err := ec.resolvers.Entity().FindFollowByID(ctx, id0)
				if err != nil {
					return fmt.Errorf(`resolving Entity "Follow": %w`, err)
				}

				list[idx[i]] = entity
				return nil
			}

		}
		return fmt.Errorf("%w: %s", ErrUnknownType, typeName)
	}

	resolveManyEntities := func(ctx context.Context, typeName string, reps []map[string]interface{}, idx []int) (err error) {
		// we need to do our own panic handling, because we may be called in a
		// goroutine, where the usual panic handling can't catch us
		defer func() {
			if r := recover(); r != nil {
				err = ec.Recover(ctx, r)
			}
		}()

		switch typeName {

		default:
			return errors.New("unknown type: " + typeName)
		}
	}

	resolveEntityGroup := func(typeName string, reps []map[string]interface{}, idx []int) {
		if isMulti(typeName) {
			err := resolveManyEntities(ctx, typeName, reps, idx)
			if err != nil {
				ec.Error(ctx, err)
			}
		} else {
			// if there are multiple entities to resolve, parallelize (similar to
			// graphql.FieldSet.Dispatch)
			var e sync.WaitGroup
			e.Add(len(reps))
			for i, rep := range reps {
				i, rep := i, rep
				go func(i int, rep map[string]interface{}) {
					err := resolveEntity(ctx, typeName, rep, idx, i)
					if err != nil {
						ec.Error(ctx, err)
					}
					e.Done()
				}(i, rep)
			}
			e.Wait()
		}
	}
	buildRepresentationGroups(representations)

	switch len(repsMap) {
	case 0:
		return list
	case 1:
		for typeName, reps := range repsMap {
			resolveEntityGroup(typeName, reps.r, reps.i)
		}
		return list
	default:
		var g sync.WaitGroup
		g.Add(len(repsMap))
		for typeName, reps := range repsMap {
			go func(typeName string, reps []map[string]interface{}, idx []int) {
				resolveEntityGroup(typeName, reps, idx)
				g.Done()
			}(typeName, reps.r, reps.i)
		}
		g.Wait()
		return list
	}
}

func entityResolverNameForComment(ctx context.Context, rep map[string]interface{}) (string, error) {
	for {
		var (
			m   map[string]interface{}
			val interface{}
			ok  bool
		)
		_ = val
		m = rep
		if _, ok = m["id"]; !ok {
			break
		}
		return "findCommentByID", nil
	}
	return "", fmt.Errorf("%w for Comment", ErrTypeNotFound)
}

func entityResolverNameForFile(ctx context.Context, rep map[string]interface{}) (string, error) {
	for {
		var (
			m   map[string]interface{}
			val interface{}
			ok  bool
		)
		_ = val
		m = rep
		if _, ok = m["id"]; !ok {
			break
		}
		return "findFileByID", nil
	}
	return "", fmt.Errorf("%w for File", ErrTypeNotFound)
}

func entityResolverNameForFollow(ctx context.Context, rep map[string]interface{}) (string, error) {
	for {
		var (
			m   map[string]interface{}
			val interface{}
			ok  bool
		)
		_ = val
		m = rep
		if _, ok = m["id"]; !ok {
			break
		}
		return "findFollowByID", nil
	}
	return "", fmt.Errorf("%w for Follow", ErrTypeNotFound)
}