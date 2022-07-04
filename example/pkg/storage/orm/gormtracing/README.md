# gorm-opentracing

[![Go Report Card](https://goreportcard.com/badge/github.com/go-gorm/opentracing)](https://goreportcard.com/report/github.com/go-gorm/opentracing) [![go.dev reference](https://img.shields.io/badge/go.dev-reference-007d9c?logo=go&logoColor=white&style=flat-square)](https://pkg.go.dev/gorm.io/plugin/opentracing)

opentracing support for gorm2.

### Features

- [x] Record `SQL` in `span` logs.
- [x] Record `Result` in `span` logs.
- [x] Record `Table` in `span` tags.
- [x] Record `Error` in `span` tags and logs.
- [x] Register `Create` `Query` `Delete` `Update` `Row` `Raw` tracing callbacks. 

### Get Started

I assume that you already have an opentracing Tracer client started in your project.

```go
func main() {
	var db *gorm.DB
	
	db.Use(gormtracing.New())
	
	// if you want to use customized tracer instead of opentracing.GlobalTracer() which is default,
	// you can use the option WithTracer(yourTracer)
}
```

### Plugin options

```go
// WithLogResult log result into span log, default: disabled.
func WithLogResult(logResult bool)

// WithTracer allows to use customized tracer rather than the global one only.
func WithTracer(tracer opentracing.Tracer)

// WithSqlParameters is a switch to control that whether record parameters in sql or not.  
func WithSqlParameters(logSqlParameters bool)

// WithErrorTagHook allows to customize error tag on opentracing.Span.
func WithErrorTagHook(errorTagHook errorTagHook)
```
