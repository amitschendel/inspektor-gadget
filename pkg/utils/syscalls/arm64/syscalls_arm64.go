package arm64

// This updated to kernel 6.6-rc2
var SyscallsNameToNumber = map[string]int{
	"_llseek":                      -1,
	"_newselect":                   -1,
	"_sysctl":                      -1,
	"accept":                       202,
	"accept4":                      242,
	"access":                       -1,
	"acct":                         89,
	"add_key":                      217,
	"adjtimex":                     171,
	"alarm":                        -1,
	"arc_gettls":                   -1,
	"arc_settls":                   -1,
	"arc_usr_cmpxchg":              -1,
	"arch_prctl":                   -1,
	"arm_fadvise64_64":             -1,
	"atomic_barrier":               -1,
	"atomic_cmpxchg_32":            -1,
	"bdflush":                      -1,
	"bind":                         200,
	"bpf":                          280,
	"brk":                          214,
	"cachectl":                     -1,
	"cacheflush":                   -1,
	"cachestat":                    451,
	"capget":                       90,
	"capset":                       91,
	"chdir":                        49,
	"chmod":                        -1,
	"chown":                        -1,
	"chown32":                      -1,
	"chroot":                       51,
	"clock_adjtime":                266,
	"clock_adjtime64":              -1,
	"clock_getres":                 114,
	"clock_getres_time64":          -1,
	"clock_gettime":                113,
	"clock_gettime64":              -1,
	"clock_nanosleep":              115,
	"clock_nanosleep_time64":       -1,
	"clock_settime":                112,
	"clock_settime64":              -1,
	"clone":                        220,
	"clone2":                       -1,
	"clone3":                       435,
	"close":                        57,
	"close_range":                  436,
	"connect":                      203,
	"copy_file_range":              285,
	"creat":                        -1,
	"create_module":                -1,
	"delete_module":                106,
	"dipc":                         -1,
	"dup":                          23,
	"dup2":                         -1,
	"dup3":                         24,
	"epoll_create":                 -1,
	"epoll_create1":                20,
	"epoll_ctl":                    21,
	"epoll_ctl_old":                -1,
	"epoll_pwait":                  22,
	"epoll_pwait2":                 441,
	"epoll_wait":                   -1,
	"epoll_wait_old":               -1,
	"eventfd":                      -1,
	"eventfd2":                     19,
	"exec_with_loader":             -1,
	"execv":                        -1,
	"execve":                       221,
	"execveat":                     281,
	"exit":                         93,
	"exit_group":                   94,
	"faccessat":                    48,
	"faccessat2":                   439,
	"fadvise64":                    223,
	"fadvise64_64":                 -1,
	"fallocate":                    47,
	"fanotify_init":                262,
	"fanotify_mark":                263,
	"fchdir":                       50,
	"fchmod":                       52,
	"fchmodat":                     53,
	"fchown":                       55,
	"fchown32":                     -1,
	"fchownat":                     54,
	"fcntl":                        25,
	"fcntl64":                      -1,
	"fdatasync":                    83,
	"fgetxattr":                    10,
	"finit_module":                 273,
	"flistxattr":                   13,
	"flock":                        32,
	"fork":                         -1,
	"fp_udfiex_crtl":               -1,
	"fremovexattr":                 16,
	"fsconfig":                     431,
	"fsetxattr":                    7,
	"fsmount":                      432,
	"fsopen":                       430,
	"fspick":                       433,
	"fstat":                        80,
	"fstat64":                      -1,
	"fstatat64":                    -1,
	"fstatfs":                      44,
	"fstatfs64":                    -1,
	"fsync":                        82,
	"ftruncate":                    46,
	"ftruncate64":                  -1,
	"futex":                        98,
	"futex_time64":                 -1,
	"futex_waitv":                  449,
	"futimesat":                    -1,
	"get_kernel_syms":              -1,
	"get_mempolicy":                236,
	"get_robust_list":              100,
	"get_thread_area":              -1,
	"getcpu":                       168,
	"getcwd":                       17,
	"getdents":                     -1,
	"getdents64":                   61,
	"getdomainname":                -1,
	"getdtablesize":                -1,
	"getegid":                      177,
	"getegid32":                    -1,
	"geteuid":                      175,
	"geteuid32":                    -1,
	"getgid":                       176,
	"getgid32":                     -1,
	"getgroups":                    158,
	"getgroups32":                  -1,
	"gethostname":                  -1,
	"getitimer":                    102,
	"getpagesize":                  -1,
	"getpeername":                  205,
	"getpgid":                      155,
	"getpgrp":                      -1,
	"getpid":                       172,
	"getpmsg":                      -1,
	"getppid":                      173,
	"getpriority":                  141,
	"getrandom":                    278,
	"getresgid":                    150,
	"getresgid32":                  -1,
	"getresuid":                    148,
	"getresuid32":                  -1,
	"getrlimit":                    163,
	"getrusage":                    165,
	"getsid":                       156,
	"getsockname":                  204,
	"getsockopt":                   209,
	"gettid":                       178,
	"gettimeofday":                 169,
	"getuid":                       174,
	"getuid32":                     -1,
	"getunwind":                    -1,
	"getxattr":                     8,
	"getxgid":                      -1,
	"getxpid":                      -1,
	"getxuid":                      -1,
	"idle":                         -1,
	"init_module":                  105,
	"inotify_add_watch":            27,
	"inotify_init":                 -1,
	"inotify_init1":                26,
	"inotify_rm_watch":             28,
	"io_cancel":                    3,
	"io_destroy":                   1,
	"io_getevents":                 4,
	"io_pgetevents":                292,
	"io_pgetevents_time64":         -1,
	"io_setup":                     0,
	"io_submit":                    2,
	"io_uring_enter":               426,
	"io_uring_register":            427,
	"io_uring_setup":               425,
	"ioctl":                        29,
	"ioperm":                       -1,
	"iopl":                         -1,
	"ioprio_get":                   31,
	"ioprio_set":                   30,
	"ipc":                          -1,
	"kcmp":                         272,
	"kern_features":                -1,
	"kexec_file_load":              294,
	"kexec_load":                   104,
	"keyctl":                       219,
	"kill":                         129,
	"landlock_add_rule":            445,
	"landlock_create_ruleset":      444,
	"landlock_restrict_self":       446,
	"lchown":                       -1,
	"lchown32":                     -1,
	"lgetxattr":                    9,
	"link":                         -1,
	"linkat":                       37,
	"listen":                       201,
	"listxattr":                    11,
	"llistxattr":                   12,
	"lookup_dcookie":               18,
	"lremovexattr":                 15,
	"lseek":                        62,
	"lsetxattr":                    6,
	"lstat":                        -1,
	"lstat64":                      -1,
	"madvise":                      233,
	"mbind":                        235,
	"membarrier":                   283,
	"memfd_create":                 279,
	"memfd_secret":                 447,
	"memory_ordering":              -1,
	"migrate_pages":                238,
	"mincore":                      232,
	"mkdir":                        -1,
	"mkdirat":                      34,
	"mknod":                        -1,
	"mknodat":                      33,
	"mlock":                        228,
	"mlock2":                       284,
	"mlockall":                     230,
	"mmap":                         222,
	"mmap2":                        -1,
	"modify_ldt":                   -1,
	"mount":                        40,
	"mount_setattr":                442,
	"move_mount":                   429,
	"move_pages":                   239,
	"mprotect":                     226,
	"mq_getsetattr":                185,
	"mq_notify":                    184,
	"mq_open":                      180,
	"mq_timedreceive":              183,
	"mq_timedreceive_time64":       -1,
	"mq_timedsend":                 182,
	"mq_timedsend_time64":          -1,
	"mq_unlink":                    181,
	"mremap":                       216,
	"msgctl":                       187,
	"msgget":                       186,
	"msgrcv":                       188,
	"msgsnd":                       189,
	"msync":                        227,
	"multiplexer":                  -1,
	"munlock":                      229,
	"munlockall":                   231,
	"munmap":                       215,
	"name_to_handle_at":            264,
	"nanosleep":                    101,
	"newfstatat":                   79,
	"nfsservctl":                   42,
	"nice":                         -1,
	"old_adjtimex":                 -1,
	"old_getpagesize":              -1,
	"oldfstat":                     -1,
	"oldlstat":                     -1,
	"oldolduname":                  -1,
	"oldstat":                      -1,
	"oldumount":                    -1,
	"olduname":                     -1,
	"open":                         -1,
	"open_by_handle_at":            265,
	"open_tree":                    428,
	"openat":                       56,
	"openat2":                      437,
	"or1k_atomic":                  -1,
	"osf_adjtime":                  -1,
	"osf_afs_syscall":              -1,
	"osf_alt_plock":                -1,
	"osf_alt_setsid":               -1,
	"osf_alt_sigpending":           -1,
	"osf_asynch_daemon":            -1,
	"osf_audcntl":                  -1,
	"osf_audgen":                   -1,
	"osf_chflags":                  -1,
	"osf_execve":                   -1,
	"osf_exportfs":                 -1,
	"osf_fchflags":                 -1,
	"osf_fdatasync":                -1,
	"osf_fpathconf":                -1,
	"osf_fstat":                    -1,
	"osf_fstatfs":                  -1,
	"osf_fstatfs64":                -1,
	"osf_fuser":                    -1,
	"osf_getaddressconf":           -1,
	"osf_getdirentries":            -1,
	"osf_getdomainname":            -1,
	"osf_getfh":                    -1,
	"osf_getfsstat":                -1,
	"osf_gethostid":                -1,
	"osf_getitimer":                -1,
	"osf_getlogin":                 -1,
	"osf_getmnt":                   -1,
	"osf_getrusage":                -1,
	"osf_getsysinfo":               -1,
	"osf_gettimeofday":             -1,
	"osf_kloadcall":                -1,
	"osf_kmodcall":                 -1,
	"osf_lstat":                    -1,
	"osf_memcntl":                  -1,
	"osf_mincore":                  -1,
	"osf_mount":                    -1,
	"osf_mremap":                   -1,
	"osf_msfs_syscall":             -1,
	"osf_msleep":                   -1,
	"osf_mvalid":                   -1,
	"osf_mwakeup":                  -1,
	"osf_naccept":                  -1,
	"osf_nfssvc":                   -1,
	"osf_ngetpeername":             -1,
	"osf_ngetsockname":             -1,
	"osf_nrecvfrom":                -1,
	"osf_nrecvmsg":                 -1,
	"osf_nsendmsg":                 -1,
	"osf_ntp_adjtime":              -1,
	"osf_ntp_gettime":              -1,
	"osf_old_creat":                -1,
	"osf_old_fstat":                -1,
	"osf_old_getpgrp":              -1,
	"osf_old_killpg":               -1,
	"osf_old_lstat":                -1,
	"osf_old_open":                 -1,
	"osf_old_sigaction":            -1,
	"osf_old_sigblock":             -1,
	"osf_old_sigreturn":            -1,
	"osf_old_sigsetmask":           -1,
	"osf_old_sigvec":               -1,
	"osf_old_stat":                 -1,
	"osf_old_vadvise":              -1,
	"osf_old_vtrace":               -1,
	"osf_old_wait":                 -1,
	"osf_oldquota":                 -1,
	"osf_pathconf":                 -1,
	"osf_pid_block":                -1,
	"osf_pid_unblock":              -1,
	"osf_plock":                    -1,
	"osf_priocntlset":              -1,
	"osf_profil":                   -1,
	"osf_proplist_syscall":         -1,
	"osf_reboot":                   -1,
	"osf_revoke":                   -1,
	"osf_sbrk":                     -1,
	"osf_security":                 -1,
	"osf_select":                   -1,
	"osf_set_program_attributes":   -1,
	"osf_set_speculative":          -1,
	"osf_sethostid":                -1,
	"osf_setitimer":                -1,
	"osf_setlogin":                 -1,
	"osf_setsysinfo":               -1,
	"osf_settimeofday":             -1,
	"osf_shmat":                    -1,
	"osf_signal":                   -1,
	"osf_sigprocmask":              -1,
	"osf_sigsendset":               -1,
	"osf_sigstack":                 -1,
	"osf_sigwaitprim":              -1,
	"osf_sstk":                     -1,
	"osf_stat":                     -1,
	"osf_statfs":                   -1,
	"osf_statfs64":                 -1,
	"osf_subsys_info":              -1,
	"osf_swapctl":                  -1,
	"osf_swapon":                   -1,
	"osf_syscall":                  -1,
	"osf_sysinfo":                  -1,
	"osf_table":                    -1,
	"osf_uadmin":                   -1,
	"osf_usleep_thread":            -1,
	"osf_uswitch":                  -1,
	"osf_utc_adjtime":              -1,
	"osf_utc_gettime":              -1,
	"osf_utimes":                   -1,
	"osf_utsname":                  -1,
	"osf_wait4":                    -1,
	"osf_waitid":                   -1,
	"pause":                        -1,
	"pciconfig_iobase":             -1,
	"pciconfig_read":               -1,
	"pciconfig_write":              -1,
	"perf_event_open":              241,
	"perfctr":                      -1,
	"personality":                  92,
	"pidfd_getfd":                  438,
	"pidfd_open":                   434,
	"pidfd_send_signal":            424,
	"pipe":                         -1,
	"pipe2":                        59,
	"pivot_root":                   41,
	"pkey_alloc":                   289,
	"pkey_free":                    290,
	"pkey_mprotect":                288,
	"poll":                         -1,
	"ppoll":                        73,
	"ppoll_time64":                 -1,
	"prctl":                        167,
	"pread64":                      67,
	"preadv":                       69,
	"preadv2":                      286,
	"prlimit64":                    261,
	"process_madvise":              440,
	"process_mrelease":             448,
	"process_vm_readv":             270,
	"process_vm_writev":            271,
	"pselect6":                     72,
	"pselect6_time64":              -1,
	"ptrace":                       117,
	"pwrite64":                     68,
	"pwritev":                      70,
	"pwritev2":                     287,
	"query_module":                 -1,
	"quotactl":                     60,
	"quotactl_fd":                  443,
	"read":                         63,
	"readahead":                    213,
	"readdir":                      -1,
	"readlink":                     -1,
	"readlinkat":                   78,
	"readv":                        65,
	"reboot":                       142,
	"recv":                         -1,
	"recvfrom":                     207,
	"recvmmsg":                     243,
	"recvmmsg_time64":              -1,
	"recvmsg":                      212,
	"remap_file_pages":             234,
	"removexattr":                  14,
	"rename":                       -1,
	"renameat":                     38,
	"renameat2":                    276,
	"request_key":                  218,
	"restart_syscall":              128,
	"riscv_flush_icache":           -1,
	"riscv_hwprobe":                -1,
	"rmdir":                        -1,
	"rseq":                         293,
	"rt_sigaction":                 134,
	"rt_sigpending":                136,
	"rt_sigprocmask":               135,
	"rt_sigqueueinfo":              138,
	"rt_sigreturn":                 139,
	"rt_sigsuspend":                133,
	"rt_sigtimedwait":              137,
	"rt_sigtimedwait_time64":       -1,
	"rt_tgsigqueueinfo":            240,
	"rtas":                         -1,
	"s390_guarded_storage":         -1,
	"s390_pci_mmio_read":           -1,
	"s390_pci_mmio_write":          -1,
	"s390_runtime_instr":           -1,
	"s390_sthyi":                   -1,
	"sched_get_affinity":           -1,
	"sched_get_priority_max":       125,
	"sched_get_priority_min":       126,
	"sched_getaffinity":            123,
	"sched_getattr":                275,
	"sched_getparam":               121,
	"sched_getscheduler":           120,
	"sched_rr_get_interval":        127,
	"sched_rr_get_interval_time64": -1,
	"sched_set_affinity":           -1,
	"sched_setaffinity":            122,
	"sched_setattr":                274,
	"sched_setparam":               118,
	"sched_setscheduler":           119,
	"sched_yield":                  124,
	"seccomp":                      277,
	"select":                       -1,
	"semctl":                       191,
	"semget":                       190,
	"semop":                        193,
	"semtimedop":                   192,
	"semtimedop_time64":            -1,
	"send":                         -1,
	"sendfile":                     71,
	"sendfile64":                   -1,
	"sendmmsg":                     269,
	"sendmsg":                      211,
	"sendto":                       206,
	"set_mempolicy":                237,
	"set_mempolicy_home_node":      450,
	"set_robust_list":              99,
	"set_thread_area":              -1,
	"set_tid_address":              96,
	"setdomainname":                162,
	"setfsgid":                     152,
	"setfsgid32":                   -1,
	"setfsuid":                     151,
	"setfsuid32":                   -1,
	"setgid":                       144,
	"setgid32":                     -1,
	"setgroups":                    159,
	"setgroups32":                  -1,
	"sethae":                       -1,
	"sethostname":                  161,
	"setitimer":                    103,
	"setns":                        268,
	"setpgid":                      154,
	"setpgrp":                      -1,
	"setpriority":                  140,
	"setregid":                     143,
	"setregid32":                   -1,
	"setresgid":                    149,
	"setresgid32":                  -1,
	"setresuid":                    147,
	"setresuid32":                  -1,
	"setreuid":                     145,
	"setreuid32":                   -1,
	"setrlimit":                    164,
	"setsid":                       157,
	"setsockopt":                   208,
	"settimeofday":                 170,
	"setuid":                       146,
	"setuid32":                     -1,
	"setxattr":                     5,
	"sgetmask":                     -1,
	"shmat":                        196,
	"shmctl":                       195,
	"shmdt":                        197,
	"shmget":                       194,
	"shutdown":                     210,
	"sigaction":                    -1,
	"sigaltstack":                  132,
	"signal":                       -1,
	"signalfd":                     -1,
	"signalfd4":                    74,
	"sigpending":                   -1,
	"sigprocmask":                  -1,
	"sigreturn":                    -1,
	"sigsuspend":                   -1,
	"socket":                       198,
	"socketcall":                   -1,
	"socketpair":                   199,
	"splice":                       76,
	"spu_create":                   -1,
	"spu_run":                      -1,
	"ssetmask":                     -1,
	"stat":                         -1,
	"stat64":                       -1,
	"statfs":                       43,
	"statfs64":                     -1,
	"statx":                        291,
	"stime":                        -1,
	"subpage_prot":                 -1,
	"swapcontext":                  -1,
	"swapoff":                      225,
	"swapon":                       224,
	"switch_endian":                -1,
	"symlink":                      -1,
	"symlinkat":                    36,
	"sync":                         81,
	"sync_file_range":              84,
	"sync_file_range2":             -1,
	"syncfs":                       267,
	"sys_debug_setcontext":         -1,
	"syscall":                      -1,
	"sysfs":                        -1,
	"sysinfo":                      179,
	"syslog":                       116,
	"sysmips":                      -1,
	"tee":                          77,
	"tgkill":                       131,
	"time":                         -1,
	"timer_create":                 107,
	"timer_delete":                 111,
	"timer_getoverrun":             109,
	"timer_gettime":                108,
	"timer_gettime64":              -1,
	"timer_settime":                110,
	"timer_settime64":              -1,
	"timerfd":                      -1,
	"timerfd_create":               85,
	"timerfd_gettime":              87,
	"timerfd_gettime64":            -1,
	"timerfd_settime":              86,
	"timerfd_settime64":            -1,
	"times":                        153,
	"tkill":                        130,
	"truncate":                     45,
	"truncate64":                   -1,
	"ugetrlimit":                   -1,
	"umask":                        166,
	"umount":                       -1,
	"umount2":                      39,
	"uname":                        160,
	"unlink":                       -1,
	"unlinkat":                     35,
	"unshare":                      97,
	"uselib":                       -1,
	"userfaultfd":                  282,
	"ustat":                        -1,
	"utime":                        -1,
	"utimensat":                    88,
	"utimensat_time64":             -1,
	"utimes":                       -1,
	"utrap_install":                -1,
	"vfork":                        -1,
	"vhangup":                      58,
	"vm86":                         -1,
	"vm86old":                      -1,
	"vmsplice":                     75,
	"wait4":                        260,
	"waitid":                       95,
	"waitpid":                      -1,
	"write":                        64,
	"writev":                       66,
}

var SyscallsNumberToName = map[int]string{
	202: "accept",
	242: "accept4",
	89:  "acct",
	217: "add_key",
	171: "adjtimex",
	200: "bind",
	280: "bpf",
	214: "brk",
	451: "cachestat",
	90:  "capget",
	91:  "capset",
	49:  "chdir",
	51:  "chroot",
	266: "clock_adjtime",
	114: "clock_getres",
	113: "clock_gettime",
	115: "clock_nanosleep",
	112: "clock_settime",
	220: "clone",
	435: "clone3",
	57:  "close",
	436: "close_range",
	203: "connect",
	285: "copy_file_range",
	106: "delete_module",
	23:  "dup",
	24:  "dup3",
	20:  "epoll_create1",
	21:  "epoll_ctl",
	22:  "epoll_pwait",
	441: "epoll_pwait2",
	19:  "eventfd2",
	221: "execve",
	281: "execveat",
	93:  "exit",
	94:  "exit_group",
	48:  "faccessat",
	439: "faccessat2",
	223: "fadvise64",
	47:  "fallocate",
	262: "fanotify_init",
	263: "fanotify_mark",
	50:  "fchdir",
	52:  "fchmod",
	53:  "fchmodat",
	55:  "fchown",
	54:  "fchownat",
	25:  "fcntl",
	83:  "fdatasync",
	10:  "fgetxattr",
	273: "finit_module",
	13:  "flistxattr",
	32:  "flock",
	16:  "fremovexattr",
	431: "fsconfig",
	7:   "fsetxattr",
	432: "fsmount",
	430: "fsopen",
	433: "fspick",
	80:  "fstat",
	44:  "fstatfs",
	82:  "fsync",
	46:  "ftruncate",
	98:  "futex",
	449: "futex_waitv",
	236: "get_mempolicy",
	100: "get_robust_list",
	168: "getcpu",
	17:  "getcwd",
	61:  "getdents64",
	177: "getegid",
	175: "geteuid",
	176: "getgid",
	158: "getgroups",
	102: "getitimer",
	205: "getpeername",
	155: "getpgid",
	172: "getpid",
	173: "getppid",
	141: "getpriority",
	278: "getrandom",
	150: "getresgid",
	148: "getresuid",
	163: "getrlimit",
	165: "getrusage",
	156: "getsid",
	204: "getsockname",
	209: "getsockopt",
	178: "gettid",
	169: "gettimeofday",
	174: "getuid",
	8:   "getxattr",
	105: "init_module",
	27:  "inotify_add_watch",
	26:  "inotify_init1",
	28:  "inotify_rm_watch",
	3:   "io_cancel",
	1:   "io_destroy",
	4:   "io_getevents",
	292: "io_pgetevents",
	0:   "io_setup",
	2:   "io_submit",
	426: "io_uring_enter",
	427: "io_uring_register",
	425: "io_uring_setup",
	29:  "ioctl",
	31:  "ioprio_get",
	30:  "ioprio_set",
	272: "kcmp",
	294: "kexec_file_load",
	104: "kexec_load",
	219: "keyctl",
	129: "kill",
	445: "landlock_add_rule",
	444: "landlock_create_ruleset",
	446: "landlock_restrict_self",
	9:   "lgetxattr",
	37:  "linkat",
	201: "listen",
	11:  "listxattr",
	12:  "llistxattr",
	18:  "lookup_dcookie",
	15:  "lremovexattr",
	62:  "lseek",
	6:   "lsetxattr",
	233: "madvise",
	235: "mbind",
	283: "membarrier",
	279: "memfd_create",
	447: "memfd_secret",
	238: "migrate_pages",
	232: "mincore",
	34:  "mkdirat",
	33:  "mknodat",
	228: "mlock",
	284: "mlock2",
	230: "mlockall",
	222: "mmap",
	40:  "mount",
	442: "mount_setattr",
	429: "move_mount",
	239: "move_pages",
	226: "mprotect",
	185: "mq_getsetattr",
	184: "mq_notify",
	180: "mq_open",
	183: "mq_timedreceive",
	182: "mq_timedsend",
	181: "mq_unlink",
	216: "mremap",
	187: "msgctl",
	186: "msgget",
	188: "msgrcv",
	189: "msgsnd",
	227: "msync",
	229: "munlock",
	231: "munlockall",
	215: "munmap",
	264: "name_to_handle_at",
	101: "nanosleep",
	79:  "newfstatat",
	42:  "nfsservctl",
	265: "open_by_handle_at",
	428: "open_tree",
	56:  "openat",
	437: "openat2",
	241: "perf_event_open",
	92:  "personality",
	438: "pidfd_getfd",
	434: "pidfd_open",
	424: "pidfd_send_signal",
	59:  "pipe2",
	41:  "pivot_root",
	289: "pkey_alloc",
	290: "pkey_free",
	288: "pkey_mprotect",
	73:  "ppoll",
	167: "prctl",
	67:  "pread64",
	69:  "preadv",
	286: "preadv2",
	261: "prlimit64",
	440: "process_madvise",
	448: "process_mrelease",
	270: "process_vm_readv",
	271: "process_vm_writev",
	72:  "pselect6",
	117: "ptrace",
	68:  "pwrite64",
	70:  "pwritev",
	287: "pwritev2",
	60:  "quotactl",
	443: "quotactl_fd",
	63:  "read",
	213: "readahead",
	78:  "readlinkat",
	65:  "readv",
	142: "reboot",
	207: "recvfrom",
	243: "recvmmsg",
	212: "recvmsg",
	234: "remap_file_pages",
	14:  "removexattr",
	38:  "renameat",
	276: "renameat2",
	218: "request_key",
	128: "restart_syscall",
	293: "rseq",
	134: "rt_sigaction",
	136: "rt_sigpending",
	135: "rt_sigprocmask",
	138: "rt_sigqueueinfo",
	139: "rt_sigreturn",
	133: "rt_sigsuspend",
	137: "rt_sigtimedwait",
	240: "rt_tgsigqueueinfo",
	125: "sched_get_priority_max",
	126: "sched_get_priority_min",
	123: "sched_getaffinity",
	275: "sched_getattr",
	121: "sched_getparam",
	120: "sched_getscheduler",
	127: "sched_rr_get_interval",
	122: "sched_setaffinity",
	274: "sched_setattr",
	118: "sched_setparam",
	119: "sched_setscheduler",
	124: "sched_yield",
	277: "seccomp",
	191: "semctl",
	190: "semget",
	193: "semop",
	192: "semtimedop",
	71:  "sendfile",
	269: "sendmmsg",
	211: "sendmsg",
	206: "sendto",
	237: "set_mempolicy",
	450: "set_mempolicy_home_node",
	99:  "set_robust_list",
	96:  "set_tid_address",
	162: "setdomainname",
	152: "setfsgid",
	151: "setfsuid",
	144: "setgid",
	159: "setgroups",
	161: "sethostname",
	103: "setitimer",
	268: "setns",
	154: "setpgid",
	140: "setpriority",
	143: "setregid",
	149: "setresgid",
	147: "setresuid",
	145: "setreuid",
	164: "setrlimit",
	157: "setsid",
	208: "setsockopt",
	170: "settimeofday",
	146: "setuid",
	5:   "setxattr",
	196: "shmat",
	195: "shmctl",
	197: "shmdt",
	194: "shmget",
	210: "shutdown",
	132: "sigaltstack",
	74:  "signalfd4",
	198: "socket",
	199: "socketpair",
	76:  "splice",
	43:  "statfs",
	291: "statx",
	225: "swapoff",
	224: "swapon",
	36:  "symlinkat",
	81:  "sync",
	84:  "sync_file_range",
	267: "syncfs",
	179: "sysinfo",
	116: "syslog",
	77:  "tee",
	131: "tgkill",
	107: "timer_create",
	111: "timer_delete",
	109: "timer_getoverrun",
	108: "timer_gettime",
	110: "timer_settime",
	85:  "timerfd_create",
	87:  "timerfd_gettime",
	86:  "timerfd_settime",
	153: "times",
	130: "tkill",
	45:  "truncate",
	166: "umask",
	39:  "umount2",
	160: "uname",
	35:  "unlinkat",
	97:  "unshare",
	282: "userfaultfd",
	88:  "utimensat",
	58:  "vhangup",
	75:  "vmsplice",
	260: "wait4",
	95:  "waitid",
	64:  "write",
	66:  "writev",
}
