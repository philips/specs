package specs

// Spec is the base configuration for the container.  It specifies platform
// independent configuration.
type Spec struct {
	// Version is the version of the specification that is supported.
	Version string `json:"version"` // required
	// Platform is the host information for OS and Arch.
	Platform Platform `json:"platform"` // required
	// Process is the container's main process.
	Process Process `json:"process"` // required
	// TODO(brandon): revisit with lifecycle
	// Root is the root information for the container's filesystem.
	Root Root `json:"root"` // required
	// Mounts profile configuration for adding mounts to the container's filesystem.
	Mounts []MountPoint `json:"mounts"` // optional, empty array and non-existent field are equivalent
	// TODO: runc uses mounts today to make /dev/ happen
	// TODO: linux spec specifies that the /dev/ mount happens BEFORE user mounts
}

// Process contains information to start a specific application inside the container.
type Process struct {
	// Terminal creates an interactive terminal for the container.
	Terminal bool `json:"terminal"` // optional, defaults to false
	// User specifies user information for the process.
	User User `json:"user"` // required
	// TODO(brandon): recommendation: maybe modified by runtime (e.g. xdg-app)
	// Args specifies the binary and arguments for the application to execute.
	Args []string `json:"args"` // required, must have at least arg[0]
	// Env populates the process environment for the process.
	Env []string `json:"env"` // optional, empty array and non-existent field are equivalent
	// Cwd is the current working directory for the process and must be
	// relative to the container's root.
	CWD string `json:"cwd"` // optional, defaults to root on Linux (microsoft can figure it out for themselves)
}

// Root contains information about the container's root filesystem on the host.
type Root struct {
	Path      string `json:"path"`      // required
	Writeable bool   `json:"writeable"` // optional, default to false
}

// Platform specifies OS and arch information for the host system that the container
// is created for.
type Platform struct {
	OS   string `json:"os"`   // required
	Arch string `json:"arch"` // required
}

// MountPoint describes a directory that may be fullfilled by a mount in the runtime.json.
type MountPoint struct {
	Name string `json:"name"` // required
	Path string `json:"path"` // required
}
