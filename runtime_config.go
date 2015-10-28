package specs

// RuntimeSpec is the generic runtime state information on a running container
type RuntimeSpec struct {
	// Hostname is the container's host name.
	// the semantic is ???
	Hostname string `json:"hostname"` // optional b/c this would set the UTS namespace

	// Mounts is a mapping of names to mount configurations.
	// Which mounts will be mounted and where should be chosen with MountPoints
	// in Spec.
	Mounts map[string]Mount `json:"mounts"` // optional, default is empty object
	// Hooks are the commands run at various lifecycle events of the container.
	Hooks Hooks `json:"hooks"` // optional, default is empty object
}

// Hook specifies a command that is run at a particular event in the lifecycle of a container
type Hook struct {
	Path string   `json:"path"` // required
	Args []string `json:"args"` // optional,  empty array and non-existent field are equivalent
	Env  []string `json:"env"`  // optional,  empty array and non-existent field are equivalent
}

// Hooks for container setup and teardown
type Hooks struct {
	// Prestart is a list of hooks to be run before the container process is executed.
	// On Linux, they are run after the container namespaces are created.
	Prestart []Hook `json:"prestart"` // optional,  empty array and non-existent field are equivalent
	// Poststart is a list of hooks to be run after the container process is started.
	Poststart []Hook `json:"poststart"` // optional,  empty array and non-existent field are equivalent
	// Poststop is a list of hooks to be run after the container process exits.
	Poststop []Hook `json:"poststop"` // optional,  empty array and non-existent field are equivalent
}

// Mount specifies a mount for a container
type Mount struct {
	// Type specifies the mount kind.
	Type string `json:"type"` // required
	// Source specifies the source path of the mount.  In the case of bind mounts on
	// linux based systems this would be the file on the host.
	Source string `json:"source"` // required
	// Options are fstab style mount options.
	Options []string `json:"options"` // optional,  empty array and non-existent field are equivalent
}
