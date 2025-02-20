import "@typespec/http";

using TypeSpec.Http;

@service({
  title: "Moonlogs",
  #suppress "deprecated" "Use global version atm"
  version: "v1.18.6"
})
namespace Moonlogs;

model Meta {
  page: integer;
  count: integer;
  pages: integer;
}

model Response<Data> {
  success: boolean;
  code: string;
  error: string;
  data: Data;
  meta: Meta;
}

model Schema {
  id: integer;
  title: string;
  description: string;
  name: string;
  fields: SchemaField[];
  kinds: SchemaKind[];
  tag_id: integer;
  renetention_days: integer;
}

model SchemaField {
  title: string;
  name: string;
}

model SchemaKind {
  title: string;
  name: string;
}

@route("/api/schemas")
@tag("Schemas")
interface Schemas {
  /** List all schemas */
  @get list(): Response<Schema[]>;

  /** Get schema by id */
  @get read(@path id: string): Response<Schema>;

  // /** Create a widget */
  // @post create(@body body: Widget): Widget | Error;

  // /** Update a widget */
  // @patch update(@path id: string, @body body: Widget): Widget | Error;

  // /** Delete a widget */
  // @delete delete(@path id: string): void | Error;

  // /** Analyze a widget */
  // @route("{id}/analyze") @post analyze(@path id: string): AnalyzeResult | Error;
}
