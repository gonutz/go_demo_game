package ds

const (
	DIRECTSOUND_VERSION                           = 0x0900
	FX_LOCHARDWARE                                = 0x00000001
	FX_LOCSOFTWARE                                = 0x00000002
	CFX_LOCHARDWARE                               = 0x00000001
	CFX_LOCSOFTWARE                               = 0x00000002
	CFXR_LOCHARDWARE                              = 0x00000010
	CFXR_LOCSOFTWARE                              = 0x00000020
	KSPROPERTY_SUPPORT_GET                        = 0x00000001
	KSPROPERTY_SUPPORT_SET                        = 0x00000002
	FXGARGLE_WAVE_TRIANGLE                        = 0
	FXGARGLE_WAVE_SQUARE                          = 1
	FXGARGLE_RATEHZ_MIN                           = 1
	FXGARGLE_RATEHZ_MAX                           = 1000
	FXCHORUS_WAVE_TRIANGLE                        = 0
	FXCHORUS_WAVE_SIN                             = 1
	FXCHORUS_WETDRYMIX_MIN                        = 0.0
	FXCHORUS_WETDRYMIX_MAX                        = 100.0
	FXCHORUS_DEPTH_MIN                            = 0.0
	FXCHORUS_DEPTH_MAX                            = 100.0
	FXCHORUS_FEEDBACK_MIN                         = -99.0
	FXCHORUS_FEEDBACK_MAX                         = 99.0
	FXCHORUS_FREQUENCY_MIN                        = 0.0
	FXCHORUS_FREQUENCY_MAX                        = 10.0
	FXCHORUS_DELAY_MIN                            = 0.0
	FXCHORUS_DELAY_MAX                            = 20.0
	FXCHORUS_PHASE_MIN                            = 0
	FXCHORUS_PHASE_MAX                            = 4
	FXCHORUS_PHASE_NEG_180                        = 0
	FXCHORUS_PHASE_NEG_90                         = 1
	FXCHORUS_PHASE_ZERO                           = 2
	FXCHORUS_PHASE_90                             = 3
	FXCHORUS_PHASE_180                            = 4
	FXFLANGER_WAVE_TRIANGLE                       = 0
	FXFLANGER_WAVE_SIN                            = 1
	FXFLANGER_WETDRYMIX_MIN                       = 0.0
	FXFLANGER_WETDRYMIX_MAX                       = 100.0
	FXFLANGER_FREQUENCY_MIN                       = 0.0
	FXFLANGER_FREQUENCY_MAX                       = 10.0
	FXFLANGER_DEPTH_MIN                           = 0.0
	FXFLANGER_DEPTH_MAX                           = 100.0
	FXFLANGER_PHASE_MIN                           = 0
	FXFLANGER_PHASE_MAX                           = 4
	FXFLANGER_FEEDBACK_MIN                        = -99.0
	FXFLANGER_FEEDBACK_MAX                        = 99.0
	FXFLANGER_DELAY_MIN                           = 0.0
	FXFLANGER_DELAY_MAX                           = 4.0
	FXFLANGER_PHASE_NEG_180                       = 0
	FXFLANGER_PHASE_NEG_90                        = 1
	FXFLANGER_PHASE_ZERO                          = 2
	FXFLANGER_PHASE_90                            = 3
	FXFLANGER_PHASE_180                           = 4
	FXECHO_WETDRYMIX_MIN                          = 0.0
	FXECHO_WETDRYMIX_MAX                          = 100.0
	FXECHO_FEEDBACK_MIN                           = 0.0
	FXECHO_FEEDBACK_MAX                           = 100.0
	FXECHO_LEFTDELAY_MIN                          = 1.0
	FXECHO_LEFTDELAY_MAX                          = 2000.0
	FXECHO_RIGHTDELAY_MIN                         = 1.0
	FXECHO_RIGHTDELAY_MAX                         = 2000.0
	FXECHO_PANDELAY_MIN                           = 0
	FXECHO_PANDELAY_MAX                           = 1
	FXDISTORTION_GAIN_MIN                         = -60.0
	FXDISTORTION_GAIN_MAX                         = 0.0
	FXDISTORTION_EDGE_MIN                         = 0.0
	FXDISTORTION_EDGE_MAX                         = 100.0
	FXDISTORTION_POSTEQCENTERFREQUENCY_MIN        = 100.0
	FXDISTORTION_POSTEQCENTERFREQUENCY_MAX        = 8000.0
	FXDISTORTION_POSTEQBANDWIDTH_MIN              = 100.0
	FXDISTORTION_POSTEQBANDWIDTH_MAX              = 8000.0
	FXDISTORTION_PRELOWPASSCUTOFF_MIN             = 100.0
	FXDISTORTION_PRELOWPASSCUTOFF_MAX             = 8000.0
	FXCOMPRESSOR_GAIN_MIN                         = -60.0
	FXCOMPRESSOR_GAIN_MAX                         = 60.0
	FXCOMPRESSOR_ATTACK_MIN                       = 0.01
	FXCOMPRESSOR_ATTACK_MAX                       = 500.0
	FXCOMPRESSOR_RELEASE_MIN                      = 50.0
	FXCOMPRESSOR_RELEASE_MAX                      = 3000.0
	FXCOMPRESSOR_THRESHOLD_MIN                    = -60.0
	FXCOMPRESSOR_THRESHOLD_MAX                    = 0.0
	FXCOMPRESSOR_RATIO_MIN                        = 1.0
	FXCOMPRESSOR_RATIO_MAX                        = 100.0
	FXCOMPRESSOR_PREDELAY_MIN                     = 0.0
	FXCOMPRESSOR_PREDELAY_MAX                     = 4.0
	FXPARAMEQ_CENTER_MIN                          = 80.0
	FXPARAMEQ_CENTER_MAX                          = 16000.0
	FXPARAMEQ_BANDWIDTH_MIN                       = 1.0
	FXPARAMEQ_BANDWIDTH_MAX                       = 36.0
	FXPARAMEQ_GAIN_MIN                            = -15.0
	FXPARAMEQ_GAIN_MAX                            = 15.0
	FX_I3DL2REVERB_ROOM_MIN                       = (-10000)
	FX_I3DL2REVERB_ROOM_MAX                       = 0
	FX_I3DL2REVERB_ROOM_DEFAULT                   = (-1000)
	FX_I3DL2REVERB_ROOMHF_MIN                     = (-10000)
	FX_I3DL2REVERB_ROOMHF_MAX                     = 0
	FX_I3DL2REVERB_ROOMHF_DEFAULT                 = (-100)
	FX_I3DL2REVERB_ROOMROLLOFFFACTOR_MIN          = 0.0
	FX_I3DL2REVERB_ROOMROLLOFFFACTOR_MAX          = 10.0
	FX_I3DL2REVERB_ROOMROLLOFFFACTOR_DEFAULT      = 0.0
	FX_I3DL2REVERB_DECAYTIME_MIN                  = 0.1
	FX_I3DL2REVERB_DECAYTIME_MAX                  = 20.0
	FX_I3DL2REVERB_DECAYTIME_DEFAULT              = 1.49
	FX_I3DL2REVERB_DECAYHFRATIO_MIN               = 0.1
	FX_I3DL2REVERB_DECAYHFRATIO_MAX               = 2.0
	FX_I3DL2REVERB_DECAYHFRATIO_DEFAULT           = 0.83
	FX_I3DL2REVERB_REFLECTIONS_MIN                = (-10000)
	FX_I3DL2REVERB_REFLECTIONS_MAX                = 1000
	FX_I3DL2REVERB_REFLECTIONS_DEFAULT            = (-2602)
	FX_I3DL2REVERB_REFLECTIONSDELAY_MIN           = 0.0
	FX_I3DL2REVERB_REFLECTIONSDELAY_MAX           = 0.3
	FX_I3DL2REVERB_REFLECTIONSDELAY_DEFAULT       = 0.007
	FX_I3DL2REVERB_REVERB_MIN                     = (-10000)
	FX_I3DL2REVERB_REVERB_MAX                     = 2000
	FX_I3DL2REVERB_REVERB_DEFAULT                 = (200)
	FX_I3DL2REVERB_REVERBDELAY_MIN                = 0.0
	FX_I3DL2REVERB_REVERBDELAY_MAX                = 0.1
	FX_I3DL2REVERB_REVERBDELAY_DEFAULT            = 0.011
	FX_I3DL2REVERB_DIFFUSION_MIN                  = 0.0
	FX_I3DL2REVERB_DIFFUSION_MAX                  = 100.0
	FX_I3DL2REVERB_DIFFUSION_DEFAULT              = 100.0
	FX_I3DL2REVERB_DENSITY_MIN                    = 0.0
	FX_I3DL2REVERB_DENSITY_MAX                    = 100.0
	FX_I3DL2REVERB_DENSITY_DEFAULT                = 100.0
	FX_I3DL2REVERB_HFREFERENCE_MIN                = 20.0
	FX_I3DL2REVERB_HFREFERENCE_MAX                = 20000.0
	FX_I3DL2REVERB_HFREFERENCE_DEFAULT            = 5000.0
	FX_I3DL2REVERB_QUALITY_MIN                    = 0
	FX_I3DL2REVERB_QUALITY_MAX                    = 3
	FX_I3DL2REVERB_QUALITY_DEFAULT                = 2
	FX_WAVESREVERB_INGAIN_MIN                     = -96.0
	FX_WAVESREVERB_INGAIN_MAX                     = 0.0
	FX_WAVESREVERB_INGAIN_DEFAULT                 = 0.0
	FX_WAVESREVERB_REVERBMIX_MIN                  = -96.0
	FX_WAVESREVERB_REVERBMIX_MAX                  = 0.0
	FX_WAVESREVERB_REVERBMIX_DEFAULT              = 0.0
	FX_WAVESREVERB_REVERBTIME_MIN                 = 0.001
	FX_WAVESREVERB_REVERBTIME_MAX                 = 3000.0
	FX_WAVESREVERB_REVERBTIME_DEFAULT             = 1000.0
	FX_WAVESREVERB_HIGHFREQRTRATIO_MIN            = 0.001
	FX_WAVESREVERB_HIGHFREQRTRATIO_MAX            = 0.999
	FX_WAVESREVERB_HIGHFREQRTRATIO_DEFAULT        = 0.001
	CFX_AEC_MODE_PASS_THROUGH                     = 0x0
	CFX_AEC_MODE_HALF_DUPLEX                      = 0x1
	CFX_AEC_MODE_FULL_DUPLEX                      = 0x2
	CFX_AEC_STATUS_HISTORY_UNINITIALIZED          = 0x0
	CFX_AEC_STATUS_HISTORY_CONTINUOUSLY_CONVERGED = 0x1
	CFX_AEC_STATUS_HISTORY_PREVIOUSLY_DIVERGED    = 0x2
	CFX_AEC_STATUS_CURRENTLY_CONVERGED            = 0x8
	OK                                            = 0
	NO_VIRTUALIZATION                             = 142082058
	ERR_ALLOCATED                                 = -2005401590
	ERR_CONTROLUNAVAIL                            = -2005401570
	ERR_INVALIDPARAM                              = -2147024809
	ERR_INVALIDCALL                               = -2005401550
	ERR_GENERIC                                   = -2147467259
	ERR_PRIOLEVELNEEDED                           = -2005401530
	ERR_OUTOFMEMORY                               = -2147024882
	ERR_BADFORMAT                                 = -2005401500
	ERR_UNSUPPORTED                               = -2147467263
	ERR_NODRIVER                                  = -2005401480
	ERR_ALREADYINITIALIZED                        = -2005401470
	ERR_NOAGGREGATION                             = -2147221232
	ERR_BUFFERLOST                                = -2005401450
	ERR_OTHERAPPHASPRIO                           = -2005401440
	ERR_UNINITIALIZED                             = -2005401430
	ERR_NOINTERFACE                               = -2147467262
	ERR_ACCESSDENIED                              = -2147024891
	ERR_BUFFERTOOSMALL                            = -2005401420
	ERR_DS8_REQUIRED                              = -2005401410
	ERR_SENDLOOP                                  = -2005401400
	ERR_BADSENDBUFFERGUID                         = -2005401390
	ERR_OBJECTNOTFOUND                            = -2005397151
	ERR_FXUNAVAILABLE                             = -2005401380
	CAPS_PRIMARYMONO                              = 0x00000001
	CAPS_PRIMARYSTEREO                            = 0x00000002
	CAPS_PRIMARY8BIT                              = 0x00000004
	CAPS_PRIMARY16BIT                             = 0x00000008
	CAPS_CONTINUOUSRATE                           = 0x00000010
	CAPS_EMULDRIVER                               = 0x00000020
	CAPS_CERTIFIED                                = 0x00000040
	CAPS_SECONDARYMONO                            = 0x00000100
	CAPS_SECONDARYSTEREO                          = 0x00000200
	CAPS_SECONDARY8BIT                            = 0x00000400
	CAPS_SECONDARY16BIT                           = 0x00000800
	SCL_NORMAL                                    = 0x00000001
	SCL_PRIORITY                                  = 0x00000002
	SCL_EXCLUSIVE                                 = 0x00000003
	SCL_WRITEPRIMARY                              = 0x00000004
	SPEAKER_DIRECTOUT                             = 0x00000000
	SPEAKER_HEADPHONE                             = 0x00000001
	SPEAKER_MONO                                  = 0x00000002
	SPEAKER_QUAD                                  = 0x00000003
	SPEAKER_STEREO                                = 0x00000004
	SPEAKER_SURROUND                              = 0x00000005
	SPEAKER_5POINT1                               = 0x00000006 // obsolete 5.1 setting
	SPEAKER_7POINT1                               = 0x00000007 // obsolete 7.1 setting
	SPEAKER_7POINT1_SURROUND                      = 0x00000008 // correct 7.1 Home Theater setting
	SPEAKER_5POINT1_SURROUND                      = 0x00000009 // correct 5.1 setting
	SPEAKER_7POINT1_WIDE                          = SPEAKER_7POINT1
	SPEAKER_5POINT1_BACK                          = SPEAKER_5POINT1
	SPEAKER_GEOMETRY_MIN                          = 0x00000005 //   5 degrees
	SPEAKER_GEOMETRY_NARROW                       = 0x0000000A //  10 degrees
	SPEAKER_GEOMETRY_WIDE                         = 0x00000014 //  20 degrees
	SPEAKER_GEOMETRY_MAX                          = 0x000000B4 // 180 degrees
	BCAPS_PRIMARYBUFFER                           = 0x00000001
	BCAPS_STATIC                                  = 0x00000002
	BCAPS_LOCHARDWARE                             = 0x00000004
	BCAPS_LOCSOFTWARE                             = 0x00000008
	BCAPS_CTRL3D                                  = 0x00000010
	BCAPS_CTRLFREQUENCY                           = 0x00000020
	BCAPS_CTRLPAN                                 = 0x00000040
	BCAPS_CTRLVOLUME                              = 0x00000080
	BCAPS_CTRLPOSITIONNOTIFY                      = 0x00000100
	BCAPS_CTRLFX                                  = 0x00000200
	BCAPS_STICKYFOCUS                             = 0x00004000
	BCAPS_GLOBALFOCUS                             = 0x00008000
	BCAPS_GETCURRENTPOSITION2                     = 0x00010000
	BCAPS_MUTE3DATMAXDISTANCE                     = 0x00020000
	BCAPS_LOCDEFER                                = 0x00040000
	BCAPS_TRUEPLAYPOSITION                        = 0x00080000
	BPLAY_LOOPING                                 = 0x00000001
	BPLAY_LOCHARDWARE                             = 0x00000002
	BPLAY_LOCSOFTWARE                             = 0x00000004
	BPLAY_TERMINATEBY_TIME                        = 0x00000008
	BPLAY_TERMINATEBY_DISTANCE                    = 0x000000010
	BPLAY_TERMINATEBY_PRIORITY                    = 0x000000020
	BSTATUS_PLAYING                               = 0x00000001
	BSTATUS_BUFFERLOST                            = 0x00000002
	BSTATUS_LOOPING                               = 0x00000004
	BSTATUS_LOCHARDWARE                           = 0x00000008
	BSTATUS_LOCSOFTWARE                           = 0x00000010
	BSTATUS_TERMINATED                            = 0x00000020
	BLOCK_FROMWRITECURSOR                         = 0x00000001
	BLOCK_ENTIREBUFFER                            = 0x00000002
	BFREQUENCY_ORIGINAL                           = 0
	BFREQUENCY_MIN                                = 100
	BFREQUENCY_MAX                                = 200000
	BPAN_LEFT                                     = -10000
	BPAN_CENTER                                   = 0
	BPAN_RIGHT                                    = 10000
	BVOLUME_MIN                                   = -10000
	BVOLUME_MAX                                   = 0
	BSIZE_MIN                                     = 4
	BSIZE_MAX                                     = 0x0FFFFFFF
	BSIZE_FX_MIN                                  = 150 // NOTE ms, not bytes
	BNOTIFICATIONS_MAX                            = 100000
	_3DMODE_NORMAL                                = 0x00000000
	_3DMODE_HEADRELATIVE                          = 0x00000001
	_3DMODE_DISABLE                               = 0x00000002
	_3D_IMMEDIATE                                 = 0x00000000
	_3D_DEFERRED                                  = 0x00000001
	_3D_MINDISTANCEFACTOR                         = 1.175494351e-38
	_3D_MAXDISTANCEFACTOR                         = 3.402823466e+38
	_3D_DEFAULTDISTANCEFACTOR                     = 1.0
	_3D_MINROLLOFFFACTOR                          = 0.0
	_3D_MAXROLLOFFFACTOR                          = 10.0
	_3D_DEFAULTROLLOFFFACTOR                      = 1.0
	_3D_MINDOPPLERFACTOR                          = 0.0
	_3D_MAXDOPPLERFACTOR                          = 10.0
	_3D_DEFAULTDOPPLERFACTOR                      = 1.0
	_3D_DEFAULTMINDISTANCE                        = 1.0
	_3D_DEFAULTMAXDISTANCE                        = 1000000000.0
	_3D_MINCONEANGLE                              = 0
	_3D_MAXCONEANGLE                              = 360
	_3D_DEFAULTCONEANGLE                          = 360
	_3D_DEFAULTCONEOUTSIDEVOLUME                  = BVOLUME_MAX
	CCAPS_EMULDRIVER                              = CAPS_EMULDRIVER
	CCAPS_CERTIFIED                               = CAPS_CERTIFIED
	CCAPS_MULTIPLECAPTURE                         = 0x00000001
	CBCAPS_WAVEMAPPED                             = 0x80000000
	CBCAPS_CTRLFX                                 = 0x00000200
	CBLOCK_ENTIREBUFFER                           = 0x00000001
	CBSTATUS_CAPTURING                            = 0x00000001
	CBSTATUS_LOOPING                              = 0x00000002
	CBSTART_LOOPING                               = 0x00000001
	BPN_OFFSETSTOP                                = 0xFFFFFFFF
	CERTIFIED                                     = 0x00000000
	UNCERTIFIED                                   = 0x00000001
	FXR_PRESENT                                   = 0
	FXR_LOCHARDWARE                               = 1
	FXR_LOCSOFTWARE                               = 2
	FXR_UNALLOCATED                               = 3
	FXR_FAILED                                    = 4
	FXR_UNKNOWN                                   = 5
	FXR_SENDLOOP                                  = 6
	FX_I3DL2_ENVIRONMENT_PRESET_DEFAULT           = 0
	FX_I3DL2_ENVIRONMENT_PRESET_GENERIC           = 1
	FX_I3DL2_ENVIRONMENT_PRESET_PADDEDCELL        = 2
	FX_I3DL2_ENVIRONMENT_PRESET_ROOM              = 3
	FX_I3DL2_ENVIRONMENT_PRESET_BATHROOM          = 4
	FX_I3DL2_ENVIRONMENT_PRESET_LIVINGROOM        = 5
	FX_I3DL2_ENVIRONMENT_PRESET_STONEROOM         = 6
	FX_I3DL2_ENVIRONMENT_PRESET_AUDITORIUM        = 7
	FX_I3DL2_ENVIRONMENT_PRESET_CONCERTHALL       = 8
	FX_I3DL2_ENVIRONMENT_PRESET_CAVE              = 9
	FX_I3DL2_ENVIRONMENT_PRESET_ARENA             = 10
	FX_I3DL2_ENVIRONMENT_PRESET_HANGAR            = 11
	FX_I3DL2_ENVIRONMENT_PRESET_CARPETEDHALLWAY   = 12
	FX_I3DL2_ENVIRONMENT_PRESET_HALLWAY           = 13
	FX_I3DL2_ENVIRONMENT_PRESET_STONECORRIDOR     = 14
	FX_I3DL2_ENVIRONMENT_PRESET_ALLEY             = 15
	FX_I3DL2_ENVIRONMENT_PRESET_FOREST            = 16
	FX_I3DL2_ENVIRONMENT_PRESET_CITY              = 17
	FX_I3DL2_ENVIRONMENT_PRESET_MOUNTAINS         = 18
	FX_I3DL2_ENVIRONMENT_PRESET_QUARRY            = 19
	FX_I3DL2_ENVIRONMENT_PRESET_PLAIN             = 20
	FX_I3DL2_ENVIRONMENT_PRESET_PARKINGLOT        = 21
	FX_I3DL2_ENVIRONMENT_PRESET_SEWERPIPE         = 22
	FX_I3DL2_ENVIRONMENT_PRESET_UNDERWATER        = 23
	FX_I3DL2_ENVIRONMENT_PRESET_SMALLROOM         = 24
	FX_I3DL2_ENVIRONMENT_PRESET_MEDIUMROOM        = 25
	FX_I3DL2_ENVIRONMENT_PRESET_LARGEROOM         = 26
	FX_I3DL2_ENVIRONMENT_PRESET_MEDIUMHALL        = 27
	FX_I3DL2_ENVIRONMENT_PRESET_LARGEHALL         = 28
	FX_I3DL2_ENVIRONMENT_PRESET_PLATE             = 29

	WHDR_DONE      = 0x00000001 // done bit
	WHDR_PREPARED  = 0x00000002 // set if this header has been prepared
	WHDR_BEGINLOOP = 0x00000004 // loop start block
	WHDR_ENDLOOP   = 0x00000008 // loop end block
	WHDR_INQUEUE   = 0x00000010 // reserved for driver

	WAVECAPS_PITCH          = 0x0001 // supports pitch control
	WAVECAPS_PLAYBACKRATE   = 0x0002 // supports playback rate control
	WAVECAPS_VOLUME         = 0x0004 // supports volume control
	WAVECAPS_LRVOLUME       = 0x0008 // separate left-right volume control
	WAVECAPS_SYNC           = 0x0010
	WAVECAPS_SAMPLEACCURATE = 0x0020

	WAVE_FORMAT_PCM    = 1
	WAVE_INVALIDFORMAT = 0x00000000 // invalid format
	WAVE_FORMAT_1M08   = 0x00000001 // 11.025 kHz, Mono,   8-bit
	WAVE_FORMAT_1S08   = 0x00000002 // 11.025 kHz, Stereo, 8-bit
	WAVE_FORMAT_1M16   = 0x00000004 // 11.025 kHz, Mono,   16-bit
	WAVE_FORMAT_1S16   = 0x00000008 // 11.025 kHz, Stereo, 16-bit
	WAVE_FORMAT_2M08   = 0x00000010 // 22.05  kHz, Mono,   8-bit
	WAVE_FORMAT_2S08   = 0x00000020 // 22.05  kHz, Stereo, 8-bit
	WAVE_FORMAT_2M16   = 0x00000040 // 22.05  kHz, Mono,   16-bit
	WAVE_FORMAT_2S16   = 0x00000080 // 22.05  kHz, Stereo, 16-bit
	WAVE_FORMAT_4M08   = 0x00000100 // 44.1   kHz, Mono,   8-bit
	WAVE_FORMAT_4S08   = 0x00000200 // 44.1   kHz, Stereo, 8-bit
	WAVE_FORMAT_4M16   = 0x00000400 // 44.1   kHz, Mono,   16-bit
	WAVE_FORMAT_4S16   = 0x00000800 // 44.1   kHz, Stereo, 16-bit
	WAVE_FORMAT_44M08  = 0x00000100 // 44.1   kHz, Mono,   8-bit
	WAVE_FORMAT_44S08  = 0x00000200 // 44.1   kHz, Stereo, 8-bit
	WAVE_FORMAT_44M16  = 0x00000400 // 44.1   kHz, Mono,   16-bit
	WAVE_FORMAT_44S16  = 0x00000800 // 44.1   kHz, Stereo, 16-bit
	WAVE_FORMAT_48M08  = 0x00001000 // 48     kHz, Mono,   8-bit
	WAVE_FORMAT_48S08  = 0x00002000 // 48     kHz, Stereo, 8-bit
	WAVE_FORMAT_48M16  = 0x00004000 // 48     kHz, Mono,   16-bit
	WAVE_FORMAT_48S16  = 0x00008000 // 48     kHz, Stereo, 16-bit
	WAVE_FORMAT_96M08  = 0x00010000 // 96     kHz, Mono,   8-bit
	WAVE_FORMAT_96S08  = 0x00020000 // 96     kHz, Stereo, 8-bit
	WAVE_FORMAT_96M16  = 0x00040000 // 96     kHz, Mono,   16-bit
	WAVE_FORMAT_96S16  = 0x00080000 // 96     kHz, Stereo, 16-bit
)

var (
	CLSID_DirectSound                      = GUID{0x47d4d946, 0x62e8, 0x11cf, [8]byte{0x93, 0xbc, 0x44, 0x45, 0x53, 0x54, 0x0, 0x0}}
	CLSID_DirectSound8                     = GUID{0x3901cc3f, 0x84b5, 0x4fa4, [8]byte{0xba, 0x35, 0xaa, 0x81, 0x72, 0xb8, 0xa0, 0x9b}}
	CLSID_DirectSoundCapture               = GUID{0xb0210780, 0x89cd, 0x11d0, [8]byte{0xaf, 0x8, 0x0, 0xa0, 0xc9, 0x25, 0xcd, 0x16}}
	CLSID_DirectSoundCapture8              = GUID{0xe4bcac13, 0x7f99, 0x4908, [8]byte{0x9a, 0x8e, 0x74, 0xe3, 0xbf, 0x24, 0xb6, 0xe1}}
	CLSID_DirectSoundFullDuplex            = GUID{0xfea4300c, 0x7959, 0x4147, [8]byte{0xb2, 0x6a, 0x23, 0x77, 0xb9, 0xe7, 0xa9, 0x1d}}
	DEVID_DefaultPlayback                  = GUID{0xdef00000, 0x9c6d, 0x47ed, [8]byte{0xaa, 0xf1, 0x4d, 0xda, 0x8f, 0x2b, 0x5c, 0x03}}
	DEVID_DefaultCapture                   = GUID{0xdef00001, 0x9c6d, 0x47ed, [8]byte{0xaa, 0xf1, 0x4d, 0xda, 0x8f, 0x2b, 0x5c, 0x03}}
	DEVID_DefaultVoicePlayback             = GUID{0xdef00002, 0x9c6d, 0x47ed, [8]byte{0xaa, 0xf1, 0x4d, 0xda, 0x8f, 0x2b, 0x5c, 0x03}}
	DEVID_DefaultVoiceCapture              = GUID{0xdef00003, 0x9c6d, 0x47ed, [8]byte{0xaa, 0xf1, 0x4d, 0xda, 0x8f, 0x2b, 0x5c, 0x03}}
	IID_IReferenceClock                    = GUID{0x56a86897, 0x0ad4, 0x11ce, [8]byte{0xb0, 0x3a, 0x00, 0x20, 0xaf, 0x0b, 0xa7, 0x70}}
	IID_IDirectSound                       = GUID{0x279AFA83, 0x4981, 0x11CE, [8]byte{0xA5, 0x21, 0x00, 0x20, 0xAF, 0x0B, 0xE5, 0x60}}
	IID_IDirectSound8                      = GUID{0xC50A7E93, 0xF395, 0x4834, [8]byte{0x9E, 0xF6, 0x7F, 0xA9, 0x9D, 0xE5, 0x09, 0x66}}
	IID_IDirectSoundBuffer                 = GUID{0x279AFA85, 0x4981, 0x11CE, [8]byte{0xA5, 0x21, 0x00, 0x20, 0xAF, 0x0B, 0xE5, 0x60}}
	IID_IDirectSoundBuffer8                = GUID{0x6825a449, 0x7524, 0x4d82, [8]byte{0x92, 0x0f, 0x50, 0xe3, 0x6a, 0xb3, 0xab, 0x1e}}
	GUID_All_Objects                       = GUID{0xaa114de5, 0xc262, 0x4169, [8]byte{0xa1, 0xc8, 0x23, 0xd6, 0x98, 0xcc, 0x73, 0xb5}}
	IID_IDirectSound3DListener             = GUID{0x279AFA84, 0x4981, 0x11CE, [8]byte{0xA5, 0x21, 0x00, 0x20, 0xAF, 0x0B, 0xE5, 0x60}}
	IID_IDirectSound3DBuffer               = GUID{0x279AFA86, 0x4981, 0x11CE, [8]byte{0xA5, 0x21, 0x00, 0x20, 0xAF, 0x0B, 0xE5, 0x60}}
	IID_IDirectSoundCapture                = GUID{0xb0210781, 0x89cd, 0x11d0, [8]byte{0xaf, 0x8, 0x0, 0xa0, 0xc9, 0x25, 0xcd, 0x16}}
	IID_IDirectSoundCaptureBuffer          = GUID{0xb0210782, 0x89cd, 0x11d0, [8]byte{0xaf, 0x8, 0x0, 0xa0, 0xc9, 0x25, 0xcd, 0x16}}
	IID_IDirectSoundCaptureBuffer8         = GUID{0x990df4, 0xdbb, 0x4872, [8]byte{0x83, 0x3e, 0x6d, 0x30, 0x3e, 0x80, 0xae, 0xb6}}
	IID_IDirectSoundNotify                 = GUID{0xb0210783, 0x89cd, 0x11d0, [8]byte{0xaf, 0x8, 0x0, 0xa0, 0xc9, 0x25, 0xcd, 0x16}}
	IID_IKsPropertySet                     = GUID{0x31efac30, 0x515c, 0x11d0, [8]byte{0xa9, 0xaa, 0x00, 0xaa, 0x00, 0x61, 0xbe, 0x93}}
	IID_IDirectSoundFXGargle               = GUID{0xd616f352, 0xd622, 0x11ce, [8]byte{0xaa, 0xc5, 0x00, 0x20, 0xaf, 0x0b, 0x99, 0xa3}}
	IID_IDirectSoundFXChorus               = GUID{0x880842e3, 0x145f, 0x43e6, [8]byte{0xa9, 0x34, 0xa7, 0x18, 0x06, 0xe5, 0x05, 0x47}}
	IID_IDirectSoundFXFlanger              = GUID{0x903e9878, 0x2c92, 0x4072, [8]byte{0x9b, 0x2c, 0xea, 0x68, 0xf5, 0x39, 0x67, 0x83}}
	IID_IDirectSoundFXEcho                 = GUID{0x8bd28edf, 0x50db, 0x4e92, [8]byte{0xa2, 0xbd, 0x44, 0x54, 0x88, 0xd1, 0xed, 0x42}}
	IID_IDirectSoundFXDistortion           = GUID{0x8ecf4326, 0x455f, 0x4d8b, [8]byte{0xbd, 0xa9, 0x8d, 0x5d, 0x3e, 0x9e, 0x3e, 0x0b}}
	IID_IDirectSoundFXCompressor           = GUID{0x4bbd1154, 0x62f6, 0x4e2c, [8]byte{0xa1, 0x5c, 0xd3, 0xb6, 0xc4, 0x17, 0xf7, 0xa0}}
	IID_IDirectSoundFXParamEq              = GUID{0xc03ca9fe, 0xfe90, 0x4204, [8]byte{0x80, 0x78, 0x82, 0x33, 0x4c, 0xd1, 0x77, 0xda}}
	IID_IDirectSoundFXI3DL2Reverb          = GUID{0x4b166a6a, 0x0d66, 0x43f3, [8]byte{0x80, 0xe3, 0xee, 0x62, 0x80, 0xde, 0xe1, 0xa4}}
	IID_IDirectSoundFXWavesReverb          = GUID{0x46858c3a, 0x0dc6, 0x45e3, [8]byte{0xb7, 0x60, 0xd4, 0xee, 0xf1, 0x6c, 0xb3, 0x25}}
	IID_IDirectSoundCaptureFXAec           = GUID{0xad74143d, 0x903d, 0x4ab7, [8]byte{0x80, 0x66, 0x28, 0xd3, 0x63, 0x03, 0x6d, 0x65}}
	IID_IDirectSoundCaptureFXNoiseSuppress = GUID{0xed311e41, 0xfbae, 0x4175, [8]byte{0x96, 0x25, 0xcd, 0x8, 0x54, 0xf6, 0x93, 0xca}}
	IID_IDirectSoundFullDuplex             = GUID{0xedcb4c7a, 0xdaab, 0x4216, [8]byte{0xa4, 0x2e, 0x6c, 0x50, 0x59, 0x6d, 0xdc, 0x1d}}
	_3DALG_DEFAULT                         = GUID{}
	_3DALG_NO_VIRTUALIZATION               = GUID{0xc241333f, 0x1c1b, 0x11d2, [8]byte{0x94, 0xf5, 0x0, 0xc0, 0x4f, 0xc2, 0x8a, 0xca}}
	_3DALG_HRTF_FULL                       = GUID{0xc2413340, 0x1c1b, 0x11d2, [8]byte{0x94, 0xf5, 0x0, 0xc0, 0x4f, 0xc2, 0x8a, 0xca}}
	_3DALG_HRTF_LIGHT                      = GUID{0xc2413342, 0x1c1b, 0x11d2, [8]byte{0x94, 0xf5, 0x0, 0xc0, 0x4f, 0xc2, 0x8a, 0xca}}
	GUID_DSFX_STANDARD_GARGLE              = GUID{0xdafd8210, 0x5711, 0x4b91, [8]byte{0x9f, 0xe3, 0xf7, 0x5b, 0x7a, 0xe2, 0x79, 0xbf}}
	GUID_DSFX_STANDARD_CHORUS              = GUID{0xefe6629c, 0x81f7, 0x4281, [8]byte{0xbd, 0x91, 0xc9, 0xd6, 0x04, 0xa9, 0x5a, 0xf6}}
	GUID_DSFX_STANDARD_FLANGER             = GUID{0xefca3d92, 0xdfd8, 0x4672, [8]byte{0xa6, 0x03, 0x74, 0x20, 0x89, 0x4b, 0xad, 0x98}}
	GUID_DSFX_STANDARD_ECHO                = GUID{0xef3e932c, 0xd40b, 0x4f51, [8]byte{0x8c, 0xcf, 0x3f, 0x98, 0xf1, 0xb2, 0x9d, 0x5d}}
	GUID_DSFX_STANDARD_DISTORTION          = GUID{0xef114c90, 0xcd1d, 0x484e, [8]byte{0x96, 0xe5, 0x09, 0xcf, 0xaf, 0x91, 0x2a, 0x21}}
	GUID_DSFX_STANDARD_COMPRESSOR          = GUID{0xef011f79, 0x4000, 0x406d, [8]byte{0x87, 0xaf, 0xbf, 0xfb, 0x3f, 0xc3, 0x9d, 0x57}}
	GUID_DSFX_STANDARD_PARAMEQ             = GUID{0x120ced89, 0x3bf4, 0x4173, [8]byte{0xa1, 0x32, 0x3c, 0xb4, 0x06, 0xcf, 0x32, 0x31}}
	GUID_DSFX_STANDARD_I3DL2REVERB         = GUID{0xef985e71, 0xd5c7, 0x42d4, [8]byte{0xba, 0x4d, 0x2d, 0x07, 0x3e, 0x2e, 0x96, 0xf4}}
	GUID_DSFX_WAVES_REVERB                 = GUID{0x87fc0268, 0x9a55, 0x4360, [8]byte{0x95, 0xaa, 0x00, 0x4a, 0x1d, 0x9d, 0xe2, 0x6c}}
	GUID_DSCFX_CLASS_AEC                   = GUID{0xBF963D80, 0xC559, 0x11D0, [8]byte{0x8A, 0x2B, 0x00, 0xA0, 0xC9, 0x25, 0x5A, 0xC1}}
	GUID_DSCFX_MS_AEC                      = GUID{0xcdebb919, 0x379a, 0x488a, [8]byte{0x87, 0x65, 0xf5, 0x3c, 0xfd, 0x36, 0xde, 0x40}}
	GUID_DSCFX_SYSTEM_AEC                  = GUID{0x1c22c56d, 0x9879, 0x4f5b, [8]byte{0xa3, 0x89, 0x27, 0x99, 0x6d, 0xdc, 0x28, 0x10}}
	GUID_DSCFX_CLASS_NS                    = GUID{0xe07f903f, 0x62fd, 0x4e60, [8]byte{0x8c, 0xdd, 0xde, 0xa7, 0x23, 0x66, 0x65, 0xb5}}
	GUID_DSCFX_MS_NS                       = GUID{0x11c5c73b, 0x66e9, 0x4ba1, [8]byte{0xa0, 0xba, 0xe8, 0x14, 0xc6, 0xee, 0xd9, 0x2d}}
	GUID_DSCFX_SYSTEM_NS                   = GUID{0x5ab0882e, 0x7274, 0x4516, [8]byte{0x87, 0x7d, 0x4e, 0xee, 0x99, 0xba, 0x4f, 0xd0}}

	I3DL2_ENVIRONMENT_PRESET_DEFAULT = FXI3DL2Reverb{
		Room:              -1000,
		RoomHF:            -100,
		RoomRolloffFactor: 0,
		DecayTime:         1.49,
		DecayHFRatio:      0.83,
		Reflections:       -2602,
		ReflectionsDelay:  0.007,
		Reverb:            200,
		ReverbDelay:       0.011,
		Diffusion:         100,
		Density:           100,
		HFReference:       5000,
	}

	I3DL2_ENVIRONMENT_PRESET_GENERIC = FXI3DL2Reverb{
		Room:              -1000,
		RoomHF:            -100,
		RoomRolloffFactor: 0,
		DecayTime:         1.49,
		DecayHFRatio:      0.83,
		Reflections:       -2602,
		ReflectionsDelay:  0.007,
		Reverb:            200,
		ReverbDelay:       0.011,
		Diffusion:         100,
		Density:           100,
		HFReference:       5000,
	}

	I3DL2_ENVIRONMENT_PRESET_PADDEDCELL = FXI3DL2Reverb{
		Room:              -1000,
		RoomHF:            -6000,
		RoomRolloffFactor: 0,
		DecayTime:         0.17,
		DecayHFRatio:      0.1,
		Reflections:       -1204,
		ReflectionsDelay:  0.001,
		Reverb:            207,
		ReverbDelay:       0.002,
		Diffusion:         100,
		Density:           100,
		HFReference:       5000,
	}

	I3DL2_ENVIRONMENT_PRESET_ROOM = FXI3DL2Reverb{
		Room:              -1000,
		RoomHF:            -454,
		RoomRolloffFactor: 0,
		DecayTime:         0.4,
		DecayHFRatio:      0.83,
		Reflections:       -1646,
		ReflectionsDelay:  0.002,
		Reverb:            53,
		ReverbDelay:       0.003,
		Diffusion:         100,
		Density:           100,
		HFReference:       5000,
	}

	I3DL2_ENVIRONMENT_PRESET_BATHROOM = FXI3DL2Reverb{
		Room:              -1000,
		RoomHF:            -1200,
		RoomRolloffFactor: 0,
		DecayTime:         1.49,
		DecayHFRatio:      0.54,
		Reflections:       -370,
		ReflectionsDelay:  0.007,
		Reverb:            1030,
		ReverbDelay:       0.011,
		Diffusion:         100,
		Density:           60,
		HFReference:       5000,
	}

	I3DL2_ENVIRONMENT_PRESET_LIVINGROOM = FXI3DL2Reverb{
		Room:              -1000,
		RoomHF:            -6000,
		RoomRolloffFactor: 0,
		DecayTime:         0.5,
		DecayHFRatio:      0.1,
		Reflections:       -1376,
		ReflectionsDelay:  0.003,
		Reverb:            -1104,
		ReverbDelay:       0.004,
		Diffusion:         100,
		Density:           100,
		HFReference:       5000,
	}

	I3DL2_ENVIRONMENT_PRESET_STONEROOM = FXI3DL2Reverb{
		Room:              -1000,
		RoomHF:            -300,
		RoomRolloffFactor: 0,
		DecayTime:         2.31,
		DecayHFRatio:      0.64,
		Reflections:       -711,
		ReflectionsDelay:  0.012,
		Reverb:            83,
		ReverbDelay:       0.017,
		Diffusion:         100,
		Density:           100,
		HFReference:       5000,
	}

	I3DL2_ENVIRONMENT_PRESET_AUDITORIUM = FXI3DL2Reverb{
		Room:              -1000,
		RoomHF:            476,
		RoomRolloffFactor: 0,
		DecayTime:         4.32,
		DecayHFRatio:      0.59,
		Reflections:       -789,
		ReflectionsDelay:  0.02,
		Reverb:            -289,
		ReverbDelay:       0.03,
		Diffusion:         100,
		Density:           100,
		HFReference:       5000,
	}

	I3DL2_ENVIRONMENT_PRESET_CONCERTHALL = FXI3DL2Reverb{
		Room:              -1000,
		RoomHF:            -500,
		RoomRolloffFactor: 0,
		DecayTime:         3.92,
		DecayHFRatio:      0.7,
		Reflections:       -1230,
		ReflectionsDelay:  0.02,
		Reverb:            -2,
		ReverbDelay:       0.029,
		Diffusion:         100,
		Density:           100,
		HFReference:       5000,
	}

	I3DL2_ENVIRONMENT_PRESET_CAVE = FXI3DL2Reverb{
		Room:              -1000,
		RoomHF:            0,
		RoomRolloffFactor: 0,
		DecayTime:         2.91,
		DecayHFRatio:      1.3,
		Reflections:       -602,
		ReflectionsDelay:  0.015,
		Reverb:            -302,
		ReverbDelay:       0.022,
		Diffusion:         100,
		Density:           100,
		HFReference:       5000,
	}

	I3DL2_ENVIRONMENT_PRESET_ARENA = FXI3DL2Reverb{
		Room:              -1000,
		RoomHF:            -698,
		RoomRolloffFactor: 0,
		DecayTime:         7.24,
		DecayHFRatio:      0.33,
		Reflections:       -1166,
		ReflectionsDelay:  0.02,
		Reverb:            16,
		ReverbDelay:       0.03,
		Diffusion:         100,
		Density:           100,
		HFReference:       5000,
	}

	I3DL2_ENVIRONMENT_PRESET_HANGAR = FXI3DL2Reverb{
		Room:              -1000,
		RoomHF:            -1000,
		RoomRolloffFactor: 0,
		DecayTime:         10.05,
		DecayHFRatio:      0.23,
		Reflections:       -602,
		ReflectionsDelay:  0.02,
		Reverb:            198,
		ReverbDelay:       0.03,
		Diffusion:         100,
		Density:           100,
		HFReference:       5000,
	}

	I3DL2_ENVIRONMENT_PRESET_CARPETEDHALLWAY = FXI3DL2Reverb{
		Room:              -1000,
		RoomHF:            -4000,
		RoomRolloffFactor: 0,
		DecayTime:         0.3,
		DecayHFRatio:      0.1,
		Reflections:       -1831,
		ReflectionsDelay:  0.002,
		Reverb:            -1630,
		ReverbDelay:       0.03,
		Diffusion:         100,
		Density:           100,
		HFReference:       5000,
	}

	I3DL2_ENVIRONMENT_PRESET_HALLWAY = FXI3DL2Reverb{
		Room:              -1000,
		RoomHF:            -300,
		RoomRolloffFactor: 0,
		DecayTime:         1.49,
		DecayHFRatio:      0.59,
		Reflections:       -1219,
		ReflectionsDelay:  0.007,
		Reverb:            441,
		ReverbDelay:       0.011,
		Diffusion:         100,
		Density:           100,
		HFReference:       5000,
	}

	I3DL2_ENVIRONMENT_PRESET_STONECORRIDOR = FXI3DL2Reverb{
		Room:              -1000,
		RoomHF:            -237,
		RoomRolloffFactor: 0,
		DecayTime:         2.7,
		DecayHFRatio:      0.79,
		Reflections:       -1214,
		ReflectionsDelay:  0.013,
		Reverb:            395,
		ReverbDelay:       0.02,
		Diffusion:         100,
		Density:           100,
		HFReference:       5000,
	}

	I3DL2_ENVIRONMENT_PRESET_ALLEY = FXI3DL2Reverb{
		Room:              -1000,
		RoomHF:            -270,
		RoomRolloffFactor: 0,
		DecayTime:         1.49,
		DecayHFRatio:      0.86,
		Reflections:       -1204,
		ReflectionsDelay:  0.007,
		Reverb:            -4,
		ReverbDelay:       0.011,
		Diffusion:         100,
		Density:           100,
		HFReference:       5000,
	}

	I3DL2_ENVIRONMENT_PRESET_FOREST = FXI3DL2Reverb{
		Room:              -1000,
		RoomHF:            -3300,
		RoomRolloffFactor: 0,
		DecayTime:         1.49,
		DecayHFRatio:      0.54,
		Reflections:       -2560,
		ReflectionsDelay:  0.162,
		Reverb:            -613,
		ReverbDelay:       0.088,
		Diffusion:         79,
		Density:           100,
		HFReference:       5000,
	}

	I3DL2_ENVIRONMENT_PRESET_CITY = FXI3DL2Reverb{
		Room:              -1000,
		RoomHF:            -800,
		RoomRolloffFactor: 0,
		DecayTime:         1.49,
		DecayHFRatio:      0.67,
		Reflections:       -2273,
		ReflectionsDelay:  0.007,
		Reverb:            -2217,
		ReverbDelay:       0.011,
		Diffusion:         50,
		Density:           100,
		HFReference:       5000,
	}

	I3DL2_ENVIRONMENT_PRESET_MOUNTAINS = FXI3DL2Reverb{
		Room:              -1000,
		RoomHF:            -2500,
		RoomRolloffFactor: 0,
		DecayTime:         1.49,
		DecayHFRatio:      0.21,
		Reflections:       -2780,
		ReflectionsDelay:  0.3,
		Reverb:            -2014,
		ReverbDelay:       0.1,
		Diffusion:         27,
		Density:           100,
		HFReference:       5000,
	}

	I3DL2_ENVIRONMENT_PRESET_QUARRY = FXI3DL2Reverb{
		Room:              -1000,
		RoomHF:            -1000,
		RoomRolloffFactor: 0,
		DecayTime:         1.49,
		DecayHFRatio:      0.83,
		Reflections:       -10000,
		ReflectionsDelay:  0.061,
		Reverb:            500,
		ReverbDelay:       0.025,
		Diffusion:         100,
		Density:           100,
		HFReference:       5000,
	}

	I3DL2_ENVIRONMENT_PRESET_PLAIN = FXI3DL2Reverb{
		Room:              -1000,
		RoomHF:            -2000,
		RoomRolloffFactor: 0,
		DecayTime:         1.49,
		DecayHFRatio:      0.5,
		Reflections:       -2466,
		ReflectionsDelay:  0.179,
		Reverb:            -2514,
		ReverbDelay:       0.1,
		Diffusion:         21,
		Density:           100,
		HFReference:       5000,
	}

	I3DL2_ENVIRONMENT_PRESET_PARKINGLOT = FXI3DL2Reverb{
		Room:              -1000,
		RoomHF:            0,
		RoomRolloffFactor: 0,
		DecayTime:         1.65,
		DecayHFRatio:      1.5,
		Reflections:       -1363,
		ReflectionsDelay:  0.008,
		Reverb:            -1153,
		ReverbDelay:       0.012,
		Diffusion:         100,
		Density:           100,
		HFReference:       5000,
	}

	I3DL2_ENVIRONMENT_PRESET_SEWERPIPE = FXI3DL2Reverb{
		Room:              -1000,
		RoomHF:            -1000,
		RoomRolloffFactor: 0,
		DecayTime:         2.81,
		DecayHFRatio:      0.14,
		Reflections:       429,
		ReflectionsDelay:  0.014,
		Reverb:            648,
		ReverbDelay:       0.021,
		Diffusion:         80,
		Density:           60,
		HFReference:       5000,
	}

	I3DL2_ENVIRONMENT_PRESET_UNDERWATER = FXI3DL2Reverb{
		Room:              -1000,
		RoomHF:            -4000,
		RoomRolloffFactor: 0,
		DecayTime:         1.49,
		DecayHFRatio:      0.1,
		Reflections:       -449,
		ReflectionsDelay:  0.007,
		Reverb:            1700,
		ReverbDelay:       0.011,
		Diffusion:         100,
		Density:           100,
		HFReference:       5000,
	}

	I3DL2_ENVIRONMENT_PRESET_SMALLROOM = FXI3DL2Reverb{
		Room:              -1000,
		RoomHF:            -600,
		RoomRolloffFactor: 0,
		DecayTime:         1.1,
		DecayHFRatio:      0.83,
		Reflections:       -400,
		ReflectionsDelay:  0.005,
		Reverb:            500,
		ReverbDelay:       0.01,
		Diffusion:         100,
		Density:           100,
		HFReference:       5000,
	}

	I3DL2_ENVIRONMENT_PRESET_MEDIUMROOM = FXI3DL2Reverb{
		Room:              -1000,
		RoomHF:            -600,
		RoomRolloffFactor: 0,
		DecayTime:         1.3,
		DecayHFRatio:      0.83,
		Reflections:       -1000,
		ReflectionsDelay:  0.01,
		Reverb:            -200,
		ReverbDelay:       0.02,
		Diffusion:         100,
		Density:           100,
		HFReference:       5000,
	}

	I3DL2_ENVIRONMENT_PRESET_LARGEROOM = FXI3DL2Reverb{
		Room:              -1000,
		RoomHF:            -600,
		RoomRolloffFactor: 0,
		DecayTime:         1.5,
		DecayHFRatio:      0.83,
		Reflections:       -1600,
		ReflectionsDelay:  0.03,
		Reverb:            -1000,
		ReverbDelay:       0.04,
		Diffusion:         100,
		Density:           100,
		HFReference:       5000,
	}

	I3DL2_ENVIRONMENT_PRESET_MEDIUMHALL = FXI3DL2Reverb{
		Room:              -1000,
		RoomHF:            -600,
		RoomRolloffFactor: 0,
		DecayTime:         1.8,
		DecayHFRatio:      0.7,
		Reflections:       -1300,
		ReflectionsDelay:  0.015,
		Reverb:            -800,
		ReverbDelay:       0.03,
		Diffusion:         100,
		Density:           100,
		HFReference:       5000,
	}

	I3DL2_ENVIRONMENT_PRESET_LARGEHALL = FXI3DL2Reverb{
		Room:              -1000,
		RoomHF:            -600,
		RoomRolloffFactor: 0,
		DecayTime:         1.8,
		DecayHFRatio:      0.7,
		Reflections:       -2000,
		ReflectionsDelay:  0.03,
		Reverb:            -1400,
		ReverbDelay:       0.06,
		Diffusion:         100,
		Density:           100,
		HFReference:       5000,
	}

	I3DL2_ENVIRONMENT_PRESET_PLATE = FXI3DL2Reverb{
		Room:              -1000,
		RoomHF:            -200,
		RoomRolloffFactor: 0,
		DecayTime:         1.3,
		DecayHFRatio:      0.9,
		Reflections:       0,
		ReflectionsDelay:  0.002,
		Reverb:            0,
		ReverbDelay:       0.01,
		Diffusion:         100,
		Density:           75,
		HFReference:       5000,
	}
)
