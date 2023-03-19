//go:build integration
// +build integration

package analyze_test

import (
	"context"
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/emil14/neva/internal/compiler/analyze"
	"github.com/emil14/neva/internal/compiler/src"
	ts "github.com/emil14/neva/pkg/types"
)

var h src.Helper

func TestAnalyzer(t *testing.T) {
	t.Parallel()

	type testcase struct {
		enabled bool
		name    string
		prog    src.Program
		wantErr error
	}

	tests := []testcase{
		{
			name: "root_pkg_refers_to_imported_pkg",
			prog: src.Program{
				Pkgs: map[string]src.Pkg{
					"pkg2": {
						Entities: map[string]src.Entity{
							"t1": h.TypeEntity(
								true,
								h.Def( // type t1<a> = vec<a>
									h.Inst("vec", h.Inst("a")),
									h.ParamWithNoConstr("a"),
								),
							),
							"c1": {
								Exported: true,
								Kind:     src.ComponentEntity,
							},
						},
					},
					"main": {
						Imports: h.Imports("pkg2"),
						Entities: map[string]src.Entity{
							"t1": h.TypeEntity(
								true,
								h.Def( // type t1 = pkg2.t1<int>
									h.Inst("pkg2.t1", h.Inst("int")),
								),
							),
							"main": h.RootComponentEntity(map[string]src.Node{
								"n1": h.ComponentNode("pkg1", "c1"),
							}),
						},
					},
				},
			},
			wantErr: nil,
		},
		{
			name: "root_pkg_refers_imported_pkg_that_refers_another_imported_pkg",
			prog: src.Program{
				Pkgs: map[string]src.Pkg{
					"pkg3": {
						Entities: map[string]src.Entity{
							"t1": h.TypeEntity(
								true,
								h.Def( // type t1<a> = vec<a>
									h.Inst("vec", h.Inst("a")),
									h.ParamWithNoConstr("a"),
								),
							),
						},
					},
					"pkg2": {
						Imports: h.Imports("pkg3"),
						Entities: map[string]src.Entity{
							"t1": h.TypeEntity(
								true,
								h.Def( // type t1<a> = t1<a>
									h.Inst("pkg3.t1", h.Inst("a")),
									h.ParamWithNoConstr("a"),
								),
							),
							"c1": {
								Exported: true,
								Kind:     src.ComponentEntity,
							},
						},
					},
					"main": {
						Imports: h.Imports("pkg2"),
						Entities: map[string]src.Entity{
							"t1": h.TypeEntity(
								true,
								h.Def( // type t1 = pkg2.t1<int>
									h.Inst("pkg2.t1", h.Inst("int")),
								),
							),
							"main": h.RootComponentEntity(map[string]src.Node{
								"n1": h.ComponentNode("pkg1", "c1"),
							}),
						},
					},
				},
			},
			wantErr: nil,
		},
		{ // FIXME false-positive
			name: "inassignable_message_and_4-step_import_chain",
			prog: src.Program{
				Pkgs: map[string]src.Pkg{
					"pkg1": {
						Entities: map[string]src.Entity{
							"m1": h.IntMsgEntity(true, 42),
							"c1": {
								Exported: true,
								Kind:     src.ComponentEntity,
							},
						},
					},
					"pkg2": {
						Imports: h.Imports("pkg1"),
						Entities: map[string]src.Entity{
							"m1": h.MsgWithRefEntity(true, &src.EntityRef{
								Pkg:  "pkg1",
								Name: "m1",
							}),
						},
					},
					"pkg3": {
						Imports: h.Imports("pkg1", "pkg2"),
						Entities: map[string]src.Entity{
							"m1": h.IntVecMsgEntity(
								true,
								[]src.Msg{
									{
										Ref: &src.EntityRef{
											Pkg:  "pkg1",
											Name: "m1",
										},
									},
									{
										Ref: &src.EntityRef{
											Pkg:  "pkg2",
											Name: "m1",
										},
									},
									{Value: h.IntMsgValue(43)},
								},
							),
						},
					},
					"main": {
						Imports: h.Imports("pkg1", "pkg2", "pkg3"),
						Entities: map[string]src.Entity{
							"m1": h.IntVecMsgEntity(
								true,
								[]src.Msg{
									{Value: h.IntMsgValue(44)},
									{
										Ref: &src.EntityRef{
											Pkg:  "pkg3",
											Name: "m1",
										},
									},
								},
							),
							"main": h.RootComponentEntity(map[string]src.Node{
								"n1": h.ComponentNode("pkg1", "c1"),
							}),
						},
					},
				},
			},
			wantErr: analyze.ErrVecEl,
		},
		{
			enabled: true,
			name:    "pkg1_imports_pkg2_and_pkg3_but_refers_to_only_pkg2_while_pkg2_actually_refers_pkg3",
			prog: src.Program{
				Pkgs: map[string]src.Pkg{
					"pkg1": {
						Imports: h.Imports("pkg2", "pkg3"), // pkg3 unused
						Entities: map[string]src.Entity{
							"main": h.RootComponentEntity(map[string]src.Node{
								"n1": h.ComponentNode("pkg2", "c1"),
							}),
						},
					},
					"pkg2": {
						Imports: map[string]string{
							"pkg3": "pkg3",
						},
						Entities: map[string]src.Entity{
							"c1": {
								Exported: true,
								Kind:     src.ComponentEntity,
							},
							"m1": h.MsgWithRefEntity(true, &src.EntityRef{
								Pkg:  "pkg3",
								Name: "m1",
							}),
						},
					},
					"pkg3": {
						Entities: map[string]src.Entity{
							"m1": h.IntMsgEntity(true, 42),
						},
					},
				},
			},
			wantErr: analyze.ErrUnusedImport,
		},
	}

	a := analyze.MustNew(
		ts.NewDefaultResolver(),
		ts.NewDefaultCompatChecker(),
		ts.Validator{},
	)

	for _, tt := range tests {
		if !tt.enabled {
			continue
		}

		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			_, err := a.Analyze(context.Background(), tt.prog)
			fmt.Println(err)
			assert.ErrorIs(t, err, tt.wantErr)
		})
	}
}
