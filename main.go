package main

import (
	"github.com/mislavzanic/container-updater/pkg/web"
	"github.com/mislavzanic/container-updater/pkg/container"
)

func main() {
	c := container.NewClient()
	l := web.NewListener(c)
	l.Run()
	// ctx := context.Background()
	// os.Setenv("DOCKER_HOST", "unix:///var/run/docker.sock")
	// // os.Setenv("DOCKER_TLS_VERIFY", "")
	// os.Setenv("DOCKER_API_VERSION", "1.40")
	// cli, err := sdkClient.NewClientWithOpts(sdkClient.FromEnv)
	// if err != nil {
	// 	panic(err)
	// }
	// defer cli.Close()

	// imageName := "mislavzanic/blog:dev"

	// out, err := cli.ImagePull(ctx, imageName, types.ImagePullOptions{})
	// if err != nil {
	// 	panic(err)
	// }
	// defer out.Close()
	// io.Copy(os.Stdout, out)

	// resp, err := cli.ContainerCreate(ctx, &container.Config{
	// 	Image: imageName,
	// }, nil, nil, nil, "")
	// if err != nil {
	// 	panic(err)
	// }

	// if err := cli.ContainerStart(ctx, resp.ID, types.ContainerStartOptions{}); err != nil {
	// 	panic(err)
	// }

	// fmt.Println(resp.ID)
}
