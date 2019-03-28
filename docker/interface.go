package docker

import "kim/config"

// Dockerimages handles the image repository creation
type Dockerimages interface {
	CreateDockerRepository(config.Configstructure, string) string
}

// CreateDockerRepository creates the docker repository
func CreateDockerRepository(config config.Configstructure, repoPath string) string {
	switch imageRepoProvider := config.ImageRepositoryProvider; imageRepoProvider {
	case "aws":
		return awsImageRepositoryCreator(config, repoPath)
	}

	return ""
}
