package web

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/rs/zerolog/log"
	"go-svelte/auth"
	"os"
	"path/filepath"
	"strings"
)

func GetStaticFiles(staticDir string) (map[string]string, error) {
	staticFiles := make(map[string]string)
	err := filepath.Walk(staticDir, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}
		if !info.IsDir() {
			relativePath := filepath.ToSlash(strings.TrimPrefix(path, staticDir))
			url := "/" + strings.TrimPrefix(relativePath, "/")
			log.Info().Msgf("Found static file: %s, setting url: %s", path, url)
			staticFiles[url] = path
			if strings.HasSuffix(url, ".html") {
				staticFiles[strings.TrimSuffix(url, ".html")] = path
			}
		}
		return nil
	})

	if err != nil {
		return nil, err
	}

	return staticFiles, nil
}

func HandleRoutes(r *gin.Engine, staticDir string) *gin.Engine {

	// Add the auth routes
	authRoutes := r.Group("/api/auth")
	auth.AuthRoutes(authRoutes)

	r.StaticFile("/", fmt.Sprintf("%s/index.html", staticDir))
	/* Now we handle all the static files on the filesystem */
	staticFiles, err := GetStaticFiles(staticDir)
	if err != nil {
		log.Fatal().Msgf("Error getting static files: %s", err)
	}
	for url, path := range staticFiles {
		r.StaticFile(url, path)
	}

	return r
}
