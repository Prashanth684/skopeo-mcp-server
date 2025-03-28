package mcp

import (
	"context"
	"github.com/mark3labs/mcp-go/mcp"
	"github.com/mark3labs/mcp-go/server"
)

func (s *Server) initSkopeo() []server.ServerTool {
	return []server.ServerTool{
		{mcp.NewTool("valid_architectures",
			mcp.WithDescription("Displays the valid architecture values that the image information can contain"),
		), s.validArchitectures},
		{mcp.NewTool("image_inspect",
			mcp.WithDescription("Displays low level information about a container image, an oci image or an oci archive."),
			mcp.WithString("imageURL", mcp.Description("The pull spec of an image or a path to an oci layout or an oci archive"), mcp.Required()),
		), s.imageInspect},
		{mcp.NewTool("image_inspect_with_tags",
			mcp.WithDescription("Displays low level information about a container image, an oci image or an oci archive.It includes repo tags information as well which might be slow"),
			mcp.WithString("imageURL", mcp.Description("The pull spec of an image or a path to an oci layout or an oci archive"), mcp.Required()),
		), s.imageInspectWithTags},
		{mcp.NewTool("image_inspect_with_os_override",
			mcp.WithDescription("This checks if an image of that particular os and architecture exists. If so, it returns the data of that image"),
			mcp.WithString("imageURL", mcp.Description("The pull spec of an image or a path to an oci layout or an oci archive"), mcp.Required()),
			mcp.WithString("arch", mcp.Description("The architecture the image manifest that we want to check to see if it is present"), mcp.Required()),
		), s.imageInspectWithOSOverride},
	}
}

func (s *Server) imageInspect(_ context.Context, ctr mcp.CallToolRequest) (*mcp.CallToolResult, error) {
	return NewTextResult(s.skopeo.ImageInspect(ctr.Params.Arguments["imageURL"].(string))), nil
}

func (s *Server) imageInspectWithTags(_ context.Context, ctr mcp.CallToolRequest) (*mcp.CallToolResult, error) {
	return NewTextResult(s.skopeo.ImageInspectWithTags(ctr.Params.Arguments["imageURL"].(string))), nil
}

func (s *Server) imageInspectWithOSOverride(_ context.Context, ctr mcp.CallToolRequest) (*mcp.CallToolResult, error) {
	return NewTextResult(s.skopeo.ImageInspectWithOSOverride(ctr.Params.Arguments["imageURL"].(string), ctr.Params.Arguments["arch"].(string))), nil
}

func (s *Server) validArchitectures(_ context.Context, ctr mcp.CallToolRequest) (*mcp.CallToolResult, error) {
	return NewTextResult(s.skopeo.ValidArchitectures(), nil), nil
}
