package tests

import (
	"context"
	"testing"

	"github.com/aclivo/olap"
)

func StorageTestSuit(factory func() olap.Storage, t *testing.T) {
	storage := factory()
	ctx := context.Background()
	dim := olap.Dimension{Name: "Product"}
	ele1 := olap.Element{Dimension: dim.Name, Name: "car"}
	ele2 := olap.Element{Dimension: dim.Name, Name: "motorcycle"}
	cub := olap.Cube{Name: "Sales", Dimensions: []string{dim.Name}}
	cel1 := olap.Cell{Cube: cub.Name, Elements: []string{ele1.Name}, Value: 101}
	cel2 := olap.Cell{Cube: cub.Name, Elements: []string{ele2.Name}, Value: 202}

	if err := storage.AddDimension(ctx, dim); err != nil {
		t.Fatal(err)
	}

	if err := storage.AddElement(ctx, ele1); err != nil {
		t.Fatal(err)
	}

	if err := storage.AddElement(ctx, ele2); err != nil {
		t.Fatal(err)
	}

	if err := storage.AddCube(ctx, cub); err != nil {
		t.Fatal(err)
	}

	if err := storage.AddCell(ctx, cel1); err != nil {
		t.Fatal(err)
	}

	if err := storage.AddCell(ctx, cel2); err != nil {
		t.Fatal(err)
	}
}
