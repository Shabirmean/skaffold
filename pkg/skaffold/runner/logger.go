/*
Copyright 2019 The Skaffold Authors

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

package runner

import (
	"io"

	"github.com/GoogleContainerTools/skaffold/pkg/skaffold/kubernetes"
	"github.com/GoogleContainerTools/skaffold/pkg/skaffold/schema/latest"
)

func (r *SkaffoldRunner) createLogger(out io.Writer, artifacts []*latest.Artifact) {
	var imageNames []string
	for _, artifact := range artifacts {
		imageNames = append(imageNames, artifact.ImageName)
	}
	r.logger = r.newLoggerForImages(out, imageNames)
}

func (r *SkaffoldRunner) newLoggerForImages(out io.Writer, images []string) *kubernetes.LogAggregator {
	return kubernetes.NewLogAggregator(out, images, r.imageList, r.runCtx.Namespaces)
}
