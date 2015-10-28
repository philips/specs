package specs

import "os"

// LinuxStateDirectory holds the container's state information
const LinuxStateDirectory = "/run/opencontainer/containers"

// LinuxRuntimeSpec is the full specification for linux containers.
type LinuxRuntimeSpec struct {
	RuntimeSpec
	// LinuxRuntime is platform specific configuration for linux based containers.
	Linux LinuxRuntime `json:"linux"`
}

// LinuxRuntime hosts the Linux-only runtime information
type LinuxRuntime struct {
	// UIDMapping specifies user mappings for supporting user namespaces on linux.
	UIDMappings []IDMapping `json:"uidMappings"` // optional,  empty array and non-existent field are equivalent
	// GIDMapping specifies group mappings for supporting user namespaces on linux.
	GIDMappings []IDMapping `json:"gidMappings"` // optional,  empty array and non-existent field are equivalent
	// Rlimits specifies rlimit options to apply to the container's process.
	Rlimits []Rlimit `json:"rlimits"` // optional,  empty array and non-existent field are equivalent
	// Sysctl are a set of key value pairs that are set for the container on start
	Sysctl map[string]string `json:"sysctl"` // optional,  empty array and non-existent field are equivalent
	// CgroupsPath specifies the path to cgroups that are created and/or joined by the container.
	// The path is expected to be relative to the cgroups mountpoint.
	// If resources are specified, the cgroups at CgroupsPath will be updated based on resources.
	CgroupsPath string `json:"cgroupsPath"` // optional, see spec
	// Namespaces contains the namespaces that are created and/or joined by the container
	Namespaces []Namespace `json:"namespaces"` // optional, empty array and non-existent field are equivalent
	// Devices are a list of device nodes that are created and enabled for the container
	Devices []Device `json:"devices"` // optional, empty array and non-existent field are equivalent
	// RootfsPropagation is the rootfs mount propagation mode for the container
	RootfsPropagation string            `json:"rootfsPropagation"` // optional, empty array and non-existent field are equivalent
	Extensions        ExtensionUnicorns // optional
}

type ExtensionUnicorns struct {
	// ApparmorProfile specified the apparmor profile for the container.
	ApparmorProfile string `json:"apparmorProfile"` // optional TODO: optional to implement
	// SelinuxProcessLabel specifies the selinux context that the container process is run as.
	SelinuxProcessLabel string `json:"selinuxProcessLabel"` // optional TODO: optional to implement
	// Seccomp specifies the seccomp security settings for the container.
	Seccomp Seccomp `json:"seccomp"` // optional, empty array and non-existent field are equivalent
}

// Namespace is the configuration for a linux namespace
type Namespace struct {
	// Type is the type of Linux namespace
	Type NamespaceType `json:"type"` // required
	// Path is a path to an existing namespace persisted on disk that can be joined
	// and is of the same type
	Path string `json:"path"` // optional
}

// NamespaceType is one of the linux namespaces
type NamespaceType string

const (
	// PIDNamespace for isolating process IDs
	PIDNamespace NamespaceType = "pid"
	// NetworkNamespace for isolating network devices, stacks, ports, etc
	NetworkNamespace = "network"
	// MountNamespace for isolating mount points
	MountNamespace = "mount"
	// IPCNamespace for isolating System V IPC, POSIX message queues
	IPCNamespace = "ipc"
	// UTSNamespace for isolating hostname and NIS domain name
	UTSNamespace = "uts"
	// UserNamespace for isolating user and group IDs
	UserNamespace = "user"
)

// IDMapping specifies UID/GID mappings
type IDMapping struct {
	// HostID is the UID/GID of the host user or group
	HostID uint32 `json:"hostID"` // required
	// ContainerID is the UID/GID of the container's user or group
	ContainerID uint32 `json:"containerID"` // required
	// Size is the length of the range of IDs mapped between the two namespaces
	Size uint32 `json:"size"` // required
}

// Rlimit type and restrictions
type Rlimit struct {
	// Type of the rlimit to set
	Type string `json:"type"` // required
	// Hard is the hard limit for the specified type
	Hard uint64 `json:"hard"` // required
	// Soft is the soft limit for the specified type
	Soft uint64 `json:"soft"` // required
}

// Seccomp represents syscall restrictions
type Seccomp struct {
	DefaultAction Action     `json:"defaultAction"`
	Architectures []Arch     `json:"architectures"`
	Syscalls      []*Syscall `json:"syscalls"`
}

// Additional architectures permitted to be used for system calls
// By default only the native architecture of the kernel is permitted
type Arch string

const (
	ArchX86         Arch = "SCMP_ARCH_X86"
	ArchX86_64      Arch = "SCMP_ARCH_X86_64"
	ArchX32         Arch = "SCMP_ARCH_X32"
	ArchARM         Arch = "SCMP_ARCH_ARM"
	ArchAARCH64     Arch = "SCMP_ARCH_AARCH64"
	ArchMIPS        Arch = "SCMP_ARCH_MIPS"
	ArchMIPS64      Arch = "SCMP_ARCH_MIPS64"
	ArchMIPS64N32   Arch = "SCMP_ARCH_MIPS64N32"
	ArchMIPSEL      Arch = "SCMP_ARCH_MIPSEL"
	ArchMIPSEL64    Arch = "SCMP_ARCH_MIPSEL64"
	ArchMIPSEL64N32 Arch = "SCMP_ARCH_MIPSEL64N32"
)

// Action taken upon Seccomp rule match
type Action string

const (
	ActKill  Action = "SCMP_ACT_KILL"
	ActTrap  Action = "SCMP_ACT_TRAP"
	ActErrno Action = "SCMP_ACT_ERRNO"
	ActTrace Action = "SCMP_ACT_TRACE"
	ActAllow Action = "SCMP_ACT_ALLOW"
)

// Operator used to match syscall arguments in Seccomp
type Operator string

const (
	OpNotEqual     Operator = "SCMP_CMP_NE"
	OpLessThan     Operator = "SCMP_CMP_LT"
	OpLessEqual    Operator = "SCMP_CMP_LE"
	OpEqualTo      Operator = "SCMP_CMP_EQ"
	OpGreaterEqual Operator = "SCMP_CMP_GE"
	OpGreaterThan  Operator = "SCMP_CMP_GT"
	OpMaskedEqual  Operator = "SCMP_CMP_MASKED_EQ"
)

// Arg used for matching specific syscall arguments in Seccomp
type Arg struct {
	Index    uint     `json:"index"`
	Value    uint64   `json:"value"`
	ValueTwo uint64   `json:"valueTwo"`
	Op       Operator `json:"op"`
}

// Syscall is used to match a syscall in Seccomp
type Syscall struct {
	Name   string `json:"name"`
	Action Action `json:"action"`
	Args   []*Arg `json:"args"`
}
