{{- $root := . -}}

# Auto generated by kcl-test; DO NOT EDIT!!!

__kcl_test_run = option("__kcl_test_run", type="str", default="")
__kcl_test_debug = option("__kcl_test_debug", type="bool", default=False)

{{range $i, $testSchemaName := .}}
# {{$testSchemaName}} begin
if not __kcl_test_run or __kcl_test_run == '{{$testSchemaName}}':
    __kcl_test_{{$testSchemaName}} = {{$testSchemaName}}{}
# {{$testSchemaName}} end
{{end}}
