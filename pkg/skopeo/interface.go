package skopeo

// Skopeo interface
type Skopeo interface {
	// Displays the valid architecture values an image can have
	ValidArchitectures() string
	// ImageInspect displays the low-level information on images identified by the ID or name
	ImageInspect(name string) (string, error)
	// ImageInspect displays the low-level information on images identified by the ID or name
	ImageInspectWithTags(name string) (string, error)
	// ImageInspect displays the low-level information on images identified by the ID or name
	ImageInspectWithOSOverride(name string, arch string) (string, error)
	// ImageCopy copies all the manifests of the image to a registry or ociarchive
	//ImageCopy() (string, error)
}

func NewSkopeo() (Skopeo, error) {
	return newSkopeoCli()
}
