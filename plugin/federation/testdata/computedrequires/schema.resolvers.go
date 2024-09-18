package computedrequires

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.
// Code generated by github.com/99designs/gqlgen version v0.17.53-dev

import (
	"context"
	"encoding/json"
	"fmt"

	explicitrequires "github.com/99designs/gqlgen/plugin/federation/testdata/computedrequires/generated"
	model "github.com/99designs/gqlgen/plugin/federation/testdata/computedrequires/generated/models"
)

// Key3 is the resolver for the key3 field.
func (r *multiHelloMultipleRequiresResolver) Key3(ctx context.Context, obj *model.MultiHelloMultipleRequires, federationRequires map[string]interface{}) (string, error) {
	key1 := federationRequires["key1"].(string)
	key2 := federationRequires["key2"].(string)
	return key1 + ":" + key2, nil
}

// Key2 is the resolver for the key2 field.
func (r *multiHelloRequiresResolver) Key2(ctx context.Context, obj *model.MultiHelloRequires, federationRequires map[string]interface{}) (string, error) {
	key1 := federationRequires["key1"].(string)
	return key1, nil
}

// Size is the resolver for the size field.
func (r *multiPlanetRequiresNestedResolver) Size(ctx context.Context, obj *model.MultiPlanetRequiresNested, federationRequires map[string]interface{}) (int, error) {
	foo := federationRequires["world"].(map[string]interface{})["foo"].(string)
	return len(foo), nil
}

// Weight is the resolver for the weight field.
func (r *planetMultipleRequiresResolver) Weight(ctx context.Context, obj *model.PlanetMultipleRequires, foo *string, federationRequires map[string]interface{}) (int, error) {
	diameter, err := federationRequires["diameter"].(json.Number).Int64()
	if err != nil {
		return 0, err
	}

	density, err := federationRequires["density"].(json.Number).Int64()
	if err != nil {
		return 0, err
	}

	return int(diameter) + int(density), nil
}

// Size is the resolver for the size field.
func (r *planetRequiresResolver) Size(ctx context.Context, obj *model.PlanetRequires, federationRequires map[string]interface{}) (int, error) {
	diameter, err := federationRequires["diameter"].(json.Number).Int64()
	if err != nil {
		return 0, err
	}

	return int(diameter), nil
}

// Size is the resolver for the size field.
func (r *planetRequiresNestedResolver) Size(ctx context.Context, obj *model.PlanetRequiresNested, federationRequires map[string]interface{}) (int, error) {
	foo := federationRequires["world"].(map[string]interface{})["foo"].(string)
	return len(foo), nil
}

// Sizes is the resolver for the sizes field.
func (r *planetRequiresNestedResolver) Sizes(ctx context.Context, obj *model.PlanetRequiresNested, federationRequires map[string]interface{}) ([]int, error) {
	foo := federationRequires["world"].(map[string]interface{})["foo"].(string)
	return []int{len(foo)}, nil
}

// Test is the resolver for the test field.
func (r *queryResolver) Test(ctx context.Context) (*string, error) {
	panic(fmt.Errorf("not implemented: Test - test"))
}

// MultiHelloMultipleRequires returns explicitrequires.MultiHelloMultipleRequiresResolver implementation.
func (r *Resolver) MultiHelloMultipleRequires() explicitrequires.MultiHelloMultipleRequiresResolver {
	return &multiHelloMultipleRequiresResolver{r}
}

// MultiHelloRequires returns explicitrequires.MultiHelloRequiresResolver implementation.
func (r *Resolver) MultiHelloRequires() explicitrequires.MultiHelloRequiresResolver {
	return &multiHelloRequiresResolver{r}
}

// MultiPlanetRequiresNested returns explicitrequires.MultiPlanetRequiresNestedResolver implementation.
func (r *Resolver) MultiPlanetRequiresNested() explicitrequires.MultiPlanetRequiresNestedResolver {
	return &multiPlanetRequiresNestedResolver{r}
}

// PlanetMultipleRequires returns explicitrequires.PlanetMultipleRequiresResolver implementation.
func (r *Resolver) PlanetMultipleRequires() explicitrequires.PlanetMultipleRequiresResolver {
	return &planetMultipleRequiresResolver{r}
}

// PlanetRequires returns explicitrequires.PlanetRequiresResolver implementation.
func (r *Resolver) PlanetRequires() explicitrequires.PlanetRequiresResolver {
	return &planetRequiresResolver{r}
}

// PlanetRequiresNested returns explicitrequires.PlanetRequiresNestedResolver implementation.
func (r *Resolver) PlanetRequiresNested() explicitrequires.PlanetRequiresNestedResolver {
	return &planetRequiresNestedResolver{r}
}

// Query returns explicitrequires.QueryResolver implementation.
func (r *Resolver) Query() explicitrequires.QueryResolver { return &queryResolver{r} }

type multiHelloMultipleRequiresResolver struct{ *Resolver }
type multiHelloRequiresResolver struct{ *Resolver }
type multiPlanetRequiresNestedResolver struct{ *Resolver }
type planetMultipleRequiresResolver struct{ *Resolver }
type planetRequiresResolver struct{ *Resolver }
type planetRequiresNestedResolver struct{ *Resolver }
type queryResolver struct{ *Resolver }
