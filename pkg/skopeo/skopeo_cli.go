package skopeo

import (
	"errors"
	"os/exec"
	"strings"
)

type skopeoCli struct {
	filePath string
}

var architectures = []string{
	"386",
	"amd64",
	"amd64p32",
	"arm",
	"arm64",
	"loong64",
	"mips",
	"mipsle",
	"mips64",
	"mips64le",
	"ppc64",
	"ppc64le",
	"riscv64",
	"s390x",
	"wasm",
}

// ValidArchitectures - list of all the valid architecture values in an image
func (s *skopeoCli) ValidArchitectures() string {
	return strings.Join(architectures, ", ")
}

// ImageInspect - inspects a docker container image
// TODO: oci-archive needs disk access which we will do later - need the filesystem mcp
func (s *skopeoCli) ImageInspect(name string) (string, error) {
	return s.exec("inspect", "--no-tags", "docker://"+name)
}

// ImageInspectWithTags - inspects a docker contaier image with repo tags included
func (s *skopeoCli) ImageInspectWithTags(name string) (string, error) {
	return s.exec("inspect", "docker://"+name)
}

// ImageInspectWithTags - inspects a docker contaier image with repo tags included
func (s *skopeoCli) ImageInspectWithOSOverride(name string, arch string) (string, error) {
	return s.exec("inspect", "--no-tags", "--override-arch="+arch, "docker://"+name)
}

func (s *skopeoCli) exec(args ...string) (string, error) {
	output, err := exec.Command(s.filePath, args...).CombinedOutput()
	return string(output), err
}

// ImageInspectForAllArches displays the low-level information for all suppported architecture variants of the image
func (s *skopeoCli) ImageInspectForAllArches(name string) ([]string, error) {
	var archInfo []string
	for _, arch := range architectures {
		output, err := s.exec("inspect", "--no-tags", "--override-arch="+arch, "docker://"+name)
		if err != nil {
			return archInfo, err
		}
		archInfo = append(archInfo, output)
	}
	return archInfo, nil
}

func newSkopeoCli() (*skopeoCli, error) {
	for _, cmd := range []string{"skopeo", "skopeo.exe"} {
		filePath, err := exec.LookPath(cmd)
		if err != nil {
			continue
		}
		if _, err = exec.Command(filePath, "--version").CombinedOutput(); err == nil {
			return &skopeoCli{filePath}, nil
		}
	}
	return nil, errors.New("skopeo CLI not found")
}
