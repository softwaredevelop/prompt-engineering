//revive:disable:package-comments,exported
package main

import (
	"github.com/pulumi/pulumi-github/sdk/v6/go/github"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
	standardRepo "github.com/softwaredevelop/pulumi-go-components/components/github"
)

// defineInfrastructure hozza létre a projekt erőforrásait.
// A teszteléshez és a fő programhoz is használható.
func defineInfrastructure(ctx *pulumi.Context) (*standardRepo.StandardRepo, *github.IssueLabel, error) {
	repository, err := standardRepo.NewStandardRepo(ctx, "prompt-engineering", &standardRepo.StandardRepoArgs{
		RepositoryName: pulumi.String("prompt-engineering"),
		Description:    pulumi.String("This is a repository for the prompt-engineering projects."),
		Topics: pulumi.StringArray{
			pulumi.String("dagger"),
			pulumi.String("docker"),
			pulumi.String("github"),
			pulumi.String("gitlab"),
			pulumi.String("go"),
			pulumi.String("golang"),
			pulumi.String("prompt-engineering"),
			pulumi.String("pulumi"),
			pulumi.String("vscode"),
		},
	})
	if err != nil {
		return nil, nil, err
	}

	goModulesLabel, err := github.NewIssueLabel(ctx, "label-go-modules", &github.IssueLabelArgs{
		Repository:  repository.Repository.Name,
		Name:        pulumi.String("go-modules dependencies"),
		Color:       pulumi.String("9BE688"),
		Description: pulumi.String("This issue is related to go modules dependencies"),
	}, pulumi.Parent(repository))
	if err != nil {
		return nil, nil, err
	}

	ctx.Export("repositoryUrl", repository.RepositoryURL)

	return repository, goModulesLabel, nil
}

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		_, _, err := defineInfrastructure(ctx)
		return err
	})
}
