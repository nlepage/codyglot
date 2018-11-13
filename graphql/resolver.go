//go:generate gorunpkg github.com/99designs/gqlgen

package graphql

import (
	"context"

	"github.com/nlepage/codyglot/service"
	"google.golang.org/grpc"
)

type Resolver struct{}

func (r *Resolver) Query() QueryResolver {
	return &queryResolver{r}
}

type queryResolver struct{ *Resolver }

func (r *queryResolver) ExecuteSnippet(ctx context.Context, language string, snippet string, executions []ExecutionInput) (*ExecuteResponse, error) {
	var sources []FileInput

	switch language {
	case "golang":
		sources = []FileInput{
			FileInput{Path: "main.go", Content: snippet},
		}
	case "javascript":
		sources = []FileInput{
			FileInput{Path: "index.js", Content: snippet},
		}
	case "typescript":
		sources = []FileInput{
			FileInput{Path: "index.ts", Content: snippet},
		}
	}

	return r.ExecuteSources(ctx, language, sources, executions)
}

func (r *queryResolver) ExecuteSources(ctx context.Context, language string, sources []FileInput, executions []ExecutionInput) (*ExecuteResponse, error) {
	opts := []grpc.DialOption{grpc.WithInsecure()}
	// FIXME configure
	conn, err := grpc.Dial("router:9090", opts...)
	if err != nil {
		return nil, err
	}
	defer conn.Close()

	client := service.NewCodyglotClient(conn)

	req := &service.ExecuteRequest{
		Language:   language,
		Sources:    mapSources(sources),
		Executions: mapExecutions(executions),
	}

	res, err := client.Execute(ctx, req)
	if err != nil {
		return nil, err
	}

	return &ExecuteResponse{
		Compilation: mapCommandResult(res.Compilation),
		Executions:  mapCommandResults(res.Executions),
	}, nil
}

func mapSources(in []FileInput) []*service.SourceFile {
	out := make([]*service.SourceFile, len(in))
	for i, v := range in {
		out[i] = &service.SourceFile{
			Path:    v.Path,
			Content: v.Content,
		}
	}
	return out
}

func mapExecutions(in []ExecutionInput) []*service.Execution {
	out := make([]*service.Execution, len(in))
	for i, v := range in {
		out[i] = &service.Execution{
			Stdin: v.Stdin,
		}
	}
	return out
}

func mapCommandResult(in *service.CommandResult) *CommandResult {
	return &CommandResult{
		Status:   int(in.Status),
		Stdout:   in.Stdout,
		Stderr:   in.Stderr,
		Duration: int(in.Duration),
	}
}

func mapCommandResults(in []*service.CommandResult) []*CommandResult {
	out := make([]*CommandResult, len(in))
	for i, v := range in {
		out[i] = mapCommandResult(v)
	}
	return out
}
