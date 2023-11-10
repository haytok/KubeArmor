// SPDX-License-Identifier: Apache-2.0
// Copyright 2022 Authors of KubeArmor

//go:build arm64
// +build arm64

package monitor

// Syscall arm/aarch
const (
	// not present in arm but kept to not break x86-64 dependent code
	SysOpen   = 2
	SysUnlink = 87
	SysRmdir  = 84
	SysChown  = 92
	//

	SysOpenAt   = 56
	SysClose    = 57
	SysUnlinkAt = 35
	SysFChownAt = 54
	SysChroot   = 51

	SysSetuid = 146
	SysSetgid = 144

	SysMount  = 165
	SysUmount = 166

	SysSocket  = 198
	SysConnect = 203
	SysAccept  = 202
	SysBind    = 200
	SysListen  = 201

	SysExecve   = 221
	SysExecveAt = 281

	SysPtrace = 101

	DoExit            = 351
	SecurityBprmCheck = 352

	TCPConnect   = 400
	TCPAccept    = 401
	TCPConnectv6 = 402
	TCPAcceptv6  = 403

	FileOpen       = 450
	FilePermission = 451
	FileMknod      = 452
	FileUnlink     = 453
	FileMkdir      = 454

	FileRmdir   = 455
	FileSymlink = 456

	FileLink     = 457
	FileRename   = 458
	FileChmod    = 459
	FileTruncate = 460

	SocketCreate  = 461
	SocketConnect = 462
	SocketAccept  = 463
)

var syscalls = map[int32]string{
	0:   "SYS_IO_SETUP",
	1:   "SYS_IO_DESTROY",
	2:   "SYS_IO_SUBMIT",
	3:   "SYS_IO_CANCEL",
	4:   "SYS_IO_GETEVENTS",
	5:   "SYS_SETXATTR",
	6:   "SYS_LSETXATTR",
	7:   "SYS_FSETXATTR",
	8:   "SYS_GETXATTR",
	9:   "SYS_LGETXATTR",
	10:  "SYS_FGETXATTR",
	11:  "SYS_LISTXATTR",
	12:  "SYS_LLISTXATTR",
	13:  "SYS_FLISTXATTR",
	14:  "SYS_REMOVEXATTR",
	15:  "SYS_LREMOVEXATTR",
	16:  "SYS_FREMOVEXATTR",
	17:  "SYS_GETCWD",
	18:  "SYS_LOOKUP_DCOOKIE",
	19:  "SYS_EVENTFD2",
	20:  "SYS_EPOLL_CREATE1",
	21:  "SYS_EPOLL_CTL",
	22:  "SYS_EPOLL_PWAIT",
	23:  "SYS_DUP",
	24:  "SYS_DUP3",
	25:  "SYS_FCNTL",
	26:  "SYS_INOTIFY_INIT1",
	27:  "SYS_INOTIFY_ADD_WATCH",
	28:  "SYS_INOTIFY_RM_WATCH",
	29:  "SYS_IOCTL",
	30:  "SYS_IOPRIO_SET",
	31:  "SYS_IOPRIO_GET",
	32:  "SYS_FLOCK",
	33:  "SYS_MKNODAT",
	34:  "SYS_MKDIRAT",
	35:  "SYS_UNLINKAT",
	36:  "SYS_SYMLINKAT",
	37:  "SYS_LINKAT",
	38:  "SYS_RENAMEAT",
	39:  "SYS_UMOUNT2",
	40:  "SYS_MOUNT",
	41:  "SYS_PIVOT_ROOT",
	42:  "SYS_NFSSERVCTL",
	43:  "SYS_STATFS",
	44:  "SYS_FSTATFS",
	45:  "SYS_TRUNCATE",
	46:  "SYS_FTRUNCATE",
	47:  "SYS_FALLOCATE",
	48:  "SYS_FACCESSAT",
	49:  "SYS_CHDIR",
	50:  "SYS_FCHDIR",
	51:  "SYS_CHROOT",
	52:  "SYS_FCHMOD",
	53:  "SYS_FCHMODAT",
	54:  "SYS_FCHOWNAT",
	55:  "SYS_CHOWN",
	56:  "SYS_OPENAT",
	57:  "SYS_CLOSE",
	58:  "SYS_VHANGUP",
	59:  "SYS_PIPE2",
	60:  "SYS_QUOTACTL",
	61:  "SYS_GETDENTS64",
	62:  "SYS_LSEEK",
	63:  "SYS_READ",
	64:  "SYS_WRITE",
	65:  "SYS_READV",
	66:  "SYS_WRITEV",
	67:  "SYS_PREAD64",
	68:  "SYS_PWRITE64",
	69:  "SYS_PREADV",
	70:  "SYS_PWRITEV",
	71:  "SYS_SENDFILE",
	72:  "SYS_PSELECT6",
	73:  "SYS_PPOLL",
	74:  "SYS_SIGNALFD4",
	75:  "SYS_VMSPLICE",
	76:  "SYS_SPLICE",
	77:  "SYS_TEE",
	78:  "SYS_READLINKAT",
	79:  "SYS_NEWFSTATAT",
	80:  "SYS_FSTAT",
	81:  "SYS_SYNC",
	82:  "SYS_FSYNC",
	83:  "SYS_FDATASYNC",
	84:  "SYS_SYNC_FILE_RANGE",
	85:  "SYS_TIMERFD_CREATE",
	86:  "SYS_TIMERFD_SETTIME",
	87:  "SYS_TIMERFD_GETTIME",
	88:  "SYS_UTIMENSAT",
	89:  "SYS_ACCT",
	90:  "SYS_CAPGET",
	91:  "SYS_CAPSET",
	92:  "SYS_PERSONALITY",
	93:  "SYS_EXIT",
	94:  "SYS_EXIT_GROUP",
	95:  "SYS_WAITID",
	96:  "SYS_SET_TID_ADDRESS",
	97:  "SYS_UNSHARE",
	98:  "SYS_FUTEX",
	99:  "SYS_SET_ROBUST_LIST",
	100: "SYS_GET_ROBUST_LIST",
	101: "SYS_NANOSLEEP",
	102: "SYS_GETITIMER",
	103: "SYS_SETITIMER",
	104: "SYS_KEXEC_LOAD",
	105: "SYS_INIT_MODULE",
	106: "SYS_DELETE_MODULE",
	107: "SYS_TIMER_CREATE",
	108: "SYS_TIMER_GETTIME",
	109: "SYS_TIMER_GETOVERRUN",
	110: "SYS_TIMER_SETTIME",
	111: "SYS_TIMER_DELETE",
	112: "SYS_CLOCK_SETTIME",
	113: "SYS_CLOCK_GETTIME",
	114: "SYS_CLOCK_GETRES",
	115: "SYS_CLOCK_NANOSLEEP",
	116: "SYS_SYSLOG",
	117: "SYS_PTRACE",
	118: "SYS_SCHED_SETPARAM",
	119: "SYS_SCHED_SETSCHEDULER",
	120: "SYS_SCHED_GETSCHEDULER",
	121: "SYS_SCHED_GETPARAM",
	122: "SYS_SCHED_SETAFFINITY",
	123: "SYS_SCHED_GETAFFINITY",
	124: "SYS_SCHED_YIELD",
	125: "SYS_SCHED_GET_PRIORITY_MAX",
	126: "SYS_SCHED_GET_PRIORITY_MIN",
	127: "SYS_SCHED_RR_GET_INTERVAL",
	128: "SYS_RESTART_SYSCALL",
	129: "SYS_KILL",
	130: "SYS_TKILL",
	131: "SYS_TGKILL",
	132: "SYS_SIGALTSTACK",
	133: "SYS_RT_SIGSUSPEND",
	134: "SYS_RT_SIGACTION",
	135: "SYS_RT_SIGPROCMASK",
	136: "SYS_RT_SIGPENDING",
	137: "SYS_RT_SIGTIMEDWAIT",
	138: "SYS_RT_SIGQUEUEINFO",
	139: "SYS_SIGRETURN",
	140: "SYS_SETPRIORITY",
	141: "SYS_GETPRIORITY",
	142: "SYS_REBOOT",
	143: "SYS_SETREGID",
	144: "SYS_SETGID",
	145: "SYS_SETREUID",
	146: "SYS_SETUID",
	147: "SYS_SETRESUID",
	148: "SYS_GETRESUID",
	149: "SYS_SETRESGID",
	150: "SYS_GETRESGID",
	151: "SYS_SETFSUID",
	152: "SYS_SETFSGID",
	153: "SYS_TIMES",
	154: "SYS_SETPGID",
	155: "SYS_GETPGID",
	156: "SYS_GETSID",
	157: "SYS_SETSID",
	158: "SYS_GETGROUPS",
	159: "SYS_SETGROUPS",
	160: "SYS_UNAME",
	161: "SYS_SETHOSTNAME",
	162: "SYS_SETDOMAINNAME",
	163: "SYS_GETRLIMIT",
	164: "SYS_SETRLIMIT",
	165: "SYS_GETRUSAGE",
	166: "SYS_UMASK",
	167: "SYS_PRCTL",
	168: "SYS_GETCPU",
	169: "SYS_GETTIMEOFDAY",
	170: "SYS_SETTIMEOFDAY",
	171: "SYS_ADJTIMEX",
	172: "SYS_GETPID",
	173: "SYS_GETPPID",
	174: "SYS_GETUID",
	175: "SYS_GETEUID",
	176: "SYS_GETGID",
	177: "SYS_GETEGID",
	178: "SYS_GETTID",
	179: "SYS_SYSINFO",
	180: "SYS_MQ_OPEN",
	181: "SYS_MQ_UNLINK",
	182: "SYS_MQ_TIMEDSEND",
	183: "SYS_MQ_TIMEDRECEIVE",
	184: "SYS_MQ_NOTIFY",
	185: "SYS_MQ_GETSETATTR",
	186: "SYS_MSGGET",
	187: "SYS_MSGCTL",
	188: "SYS_MSGRCV",
	189: "SYS_MSGSND",
	190: "SYS_SEMGET",
	191: "SYS_SEMCTL",
	192: "SYS_SEMTIMEDOP",
	193: "SYS_SEMOP",
	194: "SYS_SHMGET",
	195: "SYS_SHMCTL",
	196: "SYS_SHMAT",
	197: "SYS_SHMDT",
	198: "SYS_SOCKET",
	199: "SYS_SOCKETPAIR",
	200: "SYS_BIND",
	201: "SYS_LISTEN",
	202: "SYS_ACCEPT",
	203: "SYS_CONNECT",
	204: "SYS_GETSOCKNAME",
	205: "SYS_GETPEERNAME",
	206: "SYS_SENDTO",
	207: "SYS_RECVFROM",
	208: "SYS_SETSOCKOPT",
	209: "SYS_GETSOCKOPT",
	210: "SYS_SHUTDOWN",
	211: "SYS_SENDMSG",
	212: "SYS_RECVMSG",
	213: "SYS_READAHEAD",
	214: "SYS_BRK",
	215: "SYS_MUNMAP",
	216: "SYS_MREMAP",
	217: "SYS_ADD_KEY",
	218: "SYS_REQUEST_KEY",
	219: "SYS_KEYCTL",
	220: "SYS_CLONE",
	221: "SYS_EXECVE",
	222: "SYS_MMAP",
	223: "SYS_FADVISE64",
	224: "SYS_SWAPON",
	225: "SYS_SWAPOFF",
	226: "SYS_MPROTECT",
	227: "SYS_MSYNC",
	228: "SYS_MLOCK",
	229: "SYS_MUNLOCK",
	230: "SYS_MLOCKALL",
	231: "SYS_MUNLOCKALL",
	232: "SYS_MINCORE",
	233: "SYS_MADVISE",
	234: "SYS_REMAP_FILE_PAGES",
	235: "SYS_MBIND",
	236: "SYS_GET_MEMPOLICY",
	237: "SYS_SET_MEMPOLICY",
	238: "SYS_MIGRATE_PAGES",
	239: "SYS_MOVE_PAGES",
	240: "SYS_RT_TGSIGQUEUEINFO",
	241: "SYS_PERF_EVENT_OPEN",
	242: "SYS_ACCEPT4",
	243: "SYS_RECVMMSG",
	// syscalls not defined for the following numbers
	244: "",
	245: "",
	246: "",
	247: "",
	248: "",
	249: "",
	250: "",
	251: "",
	252: "",
	253: "",
	254: "",
	255: "",
	256: "",
	257: "",
	258: "",
	259: "",
	260: "SYS_WAIT4",
	261: "SYS_PRLIMIT64",
	262: "SYS_FANOTIFY_INIT",
	263: "SYS_FANOTIFY_MARK",
	264: "SYS_NAME_TO_HANDLE_AT",
	265: "SYS_OPEN_BY_HANDLE_AT",
	266: "SYS_CLOCK_ADJTIME",
	267: "SYS_SYNCFS",
	268: "SYS_SETNS",
	269: "SYS_SENDMMSG",
	270: "SYS_PROCESS_VM_READV",
	271: "SYS_PROCESS_VM_WRITEV",
	272: "SYS_KCMP",
	273: "SYS_FINIT_MODULE",
	274: "SYS_SCHED_SETATTR",
	275: "SYS_SCHED_GETATTR",
	276: "SYS_RENAMEAT2",
	277: "SYS_SECCOMP",
	278: "SYS_GETRANDOM",
	279: "SYS_MEMFD_CREATE",
	280: "SYS_BPF",
	281: "SYS_EXECVEAT",
	282: "SYS_USERFAULTFD",
	283: "SYS_MEMBARRIER",
	284: "SYS_MLOCK2",
	285: "SYS_COPY_FILE_RANGE",
	286: "SYS_PREADV2",
	287: "SYS_PWRITEV2",
	288: "SYS_PKEY_MPROTECT",
	289: "SYS_PKEY_ALLOC",
	290: "SYS_PKEY_FREE",
	291: "SYS_STATX",

	352: "SECURITY_BPRM_CHECK",
	450: "FILE_OPEN",
	451: "FILE_PERMISSION",
	452: "FILE_MKNOD",
	453: "FILE_UNLINK",
	454: "FILE_MKDIR",
	455: "FILE_RMDIR",
	456: "FILE_SYMLINK",
	457: "FILE_LINK",
	458: "FILE_RENAME",
	459: "FILE_CHMOD",
	460: "FILE_TRUNCATE",
	461: "SOCKET_CREATE",
	462: "SOCKET_CONNECT",
	463: "SOCKET_ACCEPT",
}
