//go:build ignore

package main

import (
	"log"

	"ariga.io/ogent"
	"entgo.io/contrib/entoas"
	"entgo.io/ent/entc"
	"entgo.io/ent/entc/gen"
	"github.com/go-faster/yaml"
	"github.com/ogen-go/ogen"
)

func main() {
	spec := new(ogen.Spec)
	oas, err := entoas.NewExtension(
		entoas.Spec(spec),
		entoas.Mutations(func(graph *gen.Graph, spec *ogen.Spec) error {
			i := ogen.NewInfo().
				SetTitle("Konyanko Schema API").
				SetVersion("0.1.0").
				SetDescription("This is an auto generated API description for Konyanko server made out of an Ent schema definition")
			spec.SetInfo(i)

			p := make(ogen.Paths)
			for k, v := range spec.Paths {
				p["/api/v1"+k] = v
			}
			spec.Paths = p

			respSchema := ogen.NewSchema().AsArray()
			respSchema.Common = ogen.OpenAPICommon{
				Extensions: ogen.Extensions(map[string]yaml.Node{
					"x-ogen-name": yaml.Node{
						Kind:  yaml.ScalarNode,
						Value: "ItemByDate",
					},
				}),
			}
			respPropsSchemas := []*ogen.Schema{
				ogen.NewSchema().SetRef("#/components/schemas/ItemList"),
				ogen.NewSchema().
					SetType("object").
					SetProperties(&ogen.Properties{
						ogen.Property{
							Name: "episode",
							Schema: ogen.NewSchema().
								SetRef("#/components/schemas/EpisodeList"),
						},
					}),
				ogen.NewSchema().
					SetType("object").
					SetProperties(&ogen.Properties{
						ogen.Property{
							Name: "anime",
							Schema: ogen.NewSchema().
								SetRef("#/components/schemas/AnimeList"),
						},
					}),
				ogen.NewSchema().
					SetType("object").
					SetProperties(&ogen.Properties{
						ogen.Property{
							Name: "release_group",
							Schema: ogen.NewSchema().
								SetRef("#/components/schemas/ReleaseGroupList"),
						},
					}),
			}
			respOK := ogen.NewResponse().
				SetDescription("result of listing items by date").
				SetJSONContent(respSchema.SetItems(ogen.NewSchema().
					SetAllOf(respPropsSchemas)))

			spec.AddPathItem("/api/v1/items/by_date/{day}", ogen.NewPathItem().
				SetGet(ogen.NewOperation().
					SetOperationID("listItemByDate").
					SetSummary("List items by published date.").
					AddTags("Item").
					AddParameters(ogen.NewParameter().
						InPath().
						SetName("day").
						SetDescription("Item published date").
						SetRequired(true).
						SetSchema(ogen.Date()),
					).
					AddResponse("200", respOK).
					AddResponse("400", ogen.NewResponse().
						SetRef("#/components/responses/400")).
					AddResponse("404", ogen.NewResponse().
						SetRef("#/components/responses/404")).
					AddResponse("409", ogen.NewResponse().
						SetRef("#/components/responses/409")).
					AddResponse("500", ogen.NewResponse().
						SetRef("#/components/responses/500")),
				),
			)

			return nil
		}),
	)
	if err != nil {
		log.Fatalf("creating entoas extension: %v", err)
	}
	ogent, err := ogent.NewExtension(spec)
	if err != nil {
		log.Fatalf("creating ogent extension: %v", err)
	}
	err = entc.Generate("./schema", &gen.Config{}, entc.Extensions(ogent, oas))
	if err != nil {
		log.Fatalf("running ent codegen: %v", err)
	}
}
