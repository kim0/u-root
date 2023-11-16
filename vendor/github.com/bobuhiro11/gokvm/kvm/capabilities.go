package kvm

// Capability is a virtual machine capability type.
//
//go:generate stringer -type=Capability
type Capability uint8

const (
	CapIRQChip                  Capability = 0
	CapHLT                      Capability = 1
	CapMMUShadowCacheControl    Capability = 2
	CapUserMemory               Capability = 3
	CapSetTSSAddr               Capability = 4
	CapVAPIC                    Capability = 6
	CapEXTCPUID                 Capability = 7
	CapClockSource              Capability = 8
	CapNRVCPUS                  Capability = 9  /* returns recommended max vcpus per vm */
	CapNRMemSlots               Capability = 10 /* returns max memory slots per vm */
	CapPIT                      Capability = 11
	CapNopIODelay               Capability = 12
	CapPVMMU                    Capability = 13
	CapMPState                  Capability = 14
	CapCoalescedMMIO            Capability = 15
	CapSyncMMU                  Capability = 16 /* Changes to host mmap are reflected in guest */
	CapIOMMU                    Capability = 18
	CapDestroyMemoryRegionWorks Capability = 21
	CapUserNMI                  Capability = 22
	CapSetGuestDebug            Capability = 23
	CapReinjectControl          Capability = 24
	CapIRQRouting               Capability = 25
	CapIRQInjectStatus          Capability = 26
	CapAssignDevIRQ             Capability = 29
	CapJoinMemoryRegionsWorks   Capability = 30
	CapMCE                      Capability = 31
	CapIRQFD                    Capability = 32
	CapPIT2                     Capability = 33
	CapSetBootCPUID             Capability = 34
	CapPITState2                Capability = 35
	CapIOEventFD                Capability = 36
	CapSetIdentityMapAddr       Capability = 37
	CapXENHVM                   Capability = 38
	CapAdjustClock              Capability = 39
	CapInternalErrorData        Capability = 40
	CapVCPUEvents               Capability = 41
	CapS390PSW                  Capability = 42
	CapPPCSegState              Capability = 43
	CapHyperV                   Capability = 44
	CapHyperVVAPIC              Capability = 45
	CapHyperVSPIN               Capability = 46
	CapPCISEgment               Capability = 47
	CapPPCPairedSingles         Capability = 48
	CapINTRShadow               Capability = 49
	CapDebugRegs                Capability = 50
	CapX86RobustSinglestep      Capability = 51
	CapPPCOSI                   Capability = 52
	CapPPCUnsetIRQ              Capability = 53
	CapEnableCap                Capability = 54
	CapXSave                    Capability = 55
	CapXCRS                     Capability = 56
	CapPPCGetPVInfo             Capability = 57
	CapPPCIRQLevel              Capability = 58
	CapASYNCPF                  Capability = 59
	CapTSCControl               Capability = 60
	CapGetTSCkHz                Capability = 61
	CapPPCBookeSREGS            Capability = 62
	CapSPAPRTCE                 Capability = 63
	CapPPCSMT                   Capability = 64
	CapPPCRMA                   Capability = 65
	CapMAXVCPUS                 Capability = 66 /* returns max vcpus per vm */
	CapPPCHIOR                  Capability = 67
	CapPPCPAPR                  Capability = 68
	CapSWTLB                    Capability = 69
	CapONEREG                   Capability = 70
	CapS390GMap                 Capability = 71
	CapTSCDeadlineTimer         Capability = 72
	CapS390UControl             Capability = 73
	CapSyncRegs                 Capability = 74
	CapPCI23                    Capability = 75
	CapKVMClockCtrl             Capability = 76
	CapSignalMSI                Capability = 77
	CapPPCGetSMMUInfo           Capability = 78
	CapS390COW                  Capability = 79
	CapPPCAllocHTAB             Capability = 80
	CapReadOnlyMEM              Capability = 81
	CapIRQFDResample            Capability = 82
	CapPPCBokkeWatchdog         Capability = 83
	CapPPCHTABFD                Capability = 84
	CapS390CSSSupport           Capability = 85
	CapPPCEPR                   Capability = 86
	CapARMPSCI                  Capability = 87
	CapARMSetDeviceAddr         Capability = 88
	CapDeviceCtrl               Capability = 89
	CapIRQMPIC                  Capability = 90
	CapPPCRTAS                  Capability = 91
	CapIRQXICS                  Capability = 92
	CapARMEL132BIT              Capability = 93
	CapSPAPRMultiTCE            Capability = 94
	CapEXTEmulCPUID             Capability = 95
	CapHyperVTIME               Capability = 96
	CapIOAPICPolarityIgnored    Capability = 97
	CapEnableCAPVM              Capability = 98
	CapS390IRQCHIP              Capability = 99
	CapIOEVENTFDNoLength        Capability = 100
	CapVMAttributes             Capability = 101
	CapARMPSCI02                Capability = 102
	CapPPCFixupHCALL            Capability = 103
	CapPPCEnableHCALL           Capability = 104
	CapCheckExtentionVM         Capability = 105
	CapS390UserSIGP             Capability = 106
	CapS390VectorRegisters      Capability = 107
	CapS390MemOp                Capability = 108
	CapS390UserSTSI             Capability = 109
	CapS390SKEYS                Capability = 110
	CapMIPSFPU                  Capability = 111
	CapMIPSMSA                  Capability = 112
	CapS390InjectIRQ            Capability = 113
	CapS390IRQState             Capability = 114
	CapPPCHWRNG                 Capability = 115
	CapDisableQuirks            Capability = 116
	CapX86SMM                   Capability = 117
	CapMultiAddressSpace        Capability = 118
	CapGuestDebugHWBPS          Capability = 119
	CapGuestDebugHWWPS          Capability = 120
	CapSplitIRQChip             Capability = 121
	CapIOEventFDAnyLength       Capability = 122
	CapHyperVSYNIC              Capability = 123
	CapS390RI                   Capability = 124
	CapSPAPRTCE64               Capability = 125
	CapARMPMUV3                 Capability = 126
	CapVCPUAttributes           Capability = 127
	CapMAXVCPUID                Capability = 128
	CapX2APICAPI                Capability = 129
	CapS390UserINSTR0           Capability = 130
	CapMSIDEVID                 Capability = 131
	CapPPCHTM                   Capability = 132
	CapSPAPRResizeHPT           Capability = 133
	CapPPCMMURADIX              Capability = 134
	CapPPCMMUHASHV3             Capability = 135
	CapImmediateExit            Capability = 136
	CapMIPSVZ                   Capability = 137
	CapMIPSTE                   Capability = 138
	CapMIPS64BIT                Capability = 139
	CapS390GS                   Capability = 140
	CapS390AIS                  Capability = 141
	CapSPAPRTCEVFIO             Capability = 142
	CapX86DisableExits          Capability = 143
	CapARMUserIRQ               Capability = 144
	CapS390CMMAMigration        Capability = 145
	CapPPCFWNMI                 Capability = 146
	CapPPCSMTPossible           Capability = 147
	CapHyperVSYNIC2             Capability = 148
	CapHyperVVPIndex            Capability = 149
	CapS390AISMigration         Capability = 150
	CapPPCGetCPUChar            Capability = 151
	CapS390BPB                  Capability = 152
	CapGETMSRFeatures           Capability = 153
	CapHyperVEventFD            Capability = 154
	CapHyperVTLBFlush           Capability = 155
	CapS390HPage1M              Capability = 156
	CapNestedState              Capability = 157
	CapARMInjectSErrorESR       Capability = 158
	CapMSRPlatformInfo          Capability = 159
	CapPPCNestedHV              Capability = 160
	CapHyperVSendIPI            Capability = 161
	CapCoalescedPIO             Capability = 162
	CapHyperVEnlightenedVMCS    Capability = 163
	CapExceptionPayload         Capability = 164
	CapARMVMIPASize             Capability = 165
	CapManualDirtyLogProtect    Capability = 166 /* Obsolete */
	CapHyerVCPUID               Capability = 167
	CapManualDirtyLogProtect2   Capability = 168
	CapPPCIRQXive               Capability = 169
	CapARMSVE                   Capability = 170
	CapARMPTRAuthAddress        Capability = 171
	CapARMPTRAuthGeneric        Capability = 172
	CapPMUEventFilter           Capability = 173
	CapARMIRQLineLayout2        Capability = 174
	CapHyperVDirectTLBFlush     Capability = 175
	CapPPCGuestDebugSStep       Capability = 176
	CapARMNISVToUser            Capability = 177
	CapARMInjectEXTDABT         Capability = 178
	CapS390VCPUResets           Capability = 179
	CapS390Protected            Capability = 180
	CapPPCSecureGuest           Capability = 181
	CapHALTPoll                 Capability = 182
	CapASYNCPFInt               Capability = 183
	CapLastCPU                  Capability = 184
	CapSmallerMaxPhyAddr        Capability = 185
	CapS390DIAG318              Capability = 186
	CapStealTime                Capability = 187
	CapX86UserSpaceMSR          Capability = 188
	CapX86MSRFilter             Capability = 189
	CapEnforcePVFeatureCPUID    Capability = 190
	CapSysHyperVCPUID           Capability = 191
	CapDirtyLogRing             Capability = 192
	CapX86BusLockExit           Capability = 193
	CapPPCDAWR1                 Capability = 194
	CapSetGuestDebug2           Capability = 195
	CapSGXAttribute             Capability = 196
	CapVMCopyEncContextFrom     Capability = 197
	CapPTPKVM                   Capability = 198
	CapHyperVEnforceCPUID       Capability = 199
	CapSREGS2                   Capability = 200
	CapEXitHyperCall            Capability = 201
	CapPPCRPTInvalidate         Capability = 202
	CapBinaryStatsFD            Capability = 203
	CapExitOnEmulationFailure   Capability = 204
	CapARMMTE                   Capability = 205
	CapVMMoveEncContextFrom     Capability = 206
	CapVMGPABits                Capability = 207
	CapXSave2                   Capability = 208
	CapSysAttributes            Capability = 209
	CapPPCAILMode3              Capability = 210
	CapS390MemOpExtention       Capability = 211
	CapPMUCapability            Capability = 212
	CapDisableQuirks2           Capability = 213
	CapVMTSCControl             Capability = 214
	CapSystemEventData          Capability = 215
	CapARMSystemSuspend         Capability = 216
	CapS390ProtectedDump        Capability = 217
	CapX86TripleFaultEvent      Capability = 218
	CapX86NotifyVMExit          Capability = 219
	CapVMDisableNXHugePages     Capability = 220
	CapS390ZPCIOP               Capability = 221
	CapS390CPUTOPOLOGY          Capability = 222
	CapDirtyLogRingACQRel       Capability = 223
)

func CheckExtension(kvmfd uintptr, c Capability) (uintptr, error) {
	return Ioctl(kvmfd, IIO(kvmCheckExtension), uintptr(c))
}
