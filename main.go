package main

import (
	"context"
	"fmt"
	"os"
	"runtime"
	"time"

	"github.com/docker/docker/api/types"
	"github.com/docker/docker/client"
	"github.com/docker/docker/pkg/archive"
)

func main() {
	dir, err := os.Getwd()
	if err != nil {
		panic(err)
	}

	fmt.Printf("this is my current directory: %s", dir)

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	client, err := client.NewEnvClient()
	if err != nil {
		panic(err)
	}

	buildContext, err := archive.TarWithOptions(".", &archive.TarOptions{})
	if err != nil {
		panic(err)
	}

	if runtime.GOOS == "windows" {
		_, err = client.ImageBuild(ctx, buildContext, types.ImageBuildOptions{
			Context:    buildContext,
			Dockerfile: "dockerfiles/sayhi.Dockerfile",
		})
	} else {
		_, err = client.ImageBuild(ctx, buildContext, types.ImageBuildOptions{
			Context:    buildContext,
			Dockerfile: "dockerfiles/sayhi.Dockerfile",
		})
	}

	if err != nil {
		panic(err)
	}
}
